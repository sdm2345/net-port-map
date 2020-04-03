package port_map

import (
    "reflect"
)

/**
ssh 协议
╰─$ nc -l 0.0.0.0 1235 | xxd                                                                                                                 130 ↵
00000000: 5353 482d 322e 302d 4f70 656e 5353 485f  SSH-2.0-OpenSSH_
00000010: 372e 390d 0a                             7.9..

*/

type SshForward struct {
    TcpForwarder
}

func NewSshForward(addr string) Forwarder {
    return &SshForward{
        TcpForwarder: TcpForwarder{Addr: addr},
    }
}

func (f *SshForward) Name() string {
    return reflect.TypeOf(*f).Name()
}

func (f *SshForward) MinLen() int {
    return 4
}

func (f *SshForward) IsMatch(buf []byte) bool {
    return string(buf[:4]) == "SSH-"
}
