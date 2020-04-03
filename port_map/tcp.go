package port_map

import (
    "context"
    "io"
    "log"
    "net"
    "reflect"
)

type Forwarder interface {
    IsMatch(buf []byte) bool
    Name() string
    MinLen() int
    Forward(ctx context.Context, buf []byte, conn net.Conn) error
}

// 一般用 tcp 兜底 即可
type TcpForwarder struct {
    Addr string
}

func NewTcpForwarder(addr string) Forwarder {
    return &TcpForwarder{Addr: addr}
}

func (f *TcpForwarder) MinLen() int {
    return 0
}

func (f *TcpForwarder) Name() string {
    return reflect.TypeOf(*f).Name()
}

func (f *TcpForwarder) IsMatch(buf []byte) bool {
    //tcp 不需要特征
    return true
}

func (f *TcpForwarder) Forward(ctx context.Context, buf []byte, conn net.Conn) error {
    //开始执行转发
    //建立 一个远程连接
    client, err := net.Dial("tcp", f.Addr)
    if err != nil {
        log.Println("err", err)
        _ = conn.Close()
        return err
    }
    
    log.Println("connect from", conn.RemoteAddr())
    closeConn := func() {
        log.Println("close", conn.RemoteAddr())
        _ = conn.Close()
        _ = client.Close()
    }
    _, err = client.Write(buf)
    if err != nil {
        log.Println("write error", err)
        closeConn()
        return err
    }
    go func() {
        defer closeConn()
        _, _ = io.Copy(client, conn)
    }()
    go func() {
        defer closeConn()
        _, _ = io.Copy(conn, client)
    }()
    return nil
}
