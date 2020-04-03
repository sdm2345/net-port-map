package port_map

import (
	"net/http"
	"reflect"
)

var requests []string

func init() {
	requests = []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace,
	}
}

type HttpForward struct {
	TcpForwarder
}

func NewHttpForward(addr string) Forwarder {
	return &HttpForward{
		TcpForwarder{Addr: addr},
	}
}

func (f *HttpForward) Name() string {
	return reflect.TypeOf(*f).Name()
}

//head 匹配的最小长度
func (f *HttpForward) MinLen() int {
	var v0 int
	//取最大值
	for _, v := range requests {
		v0=getMax(v0,len(v))
	}
	return v0 + 1
}

func (f *HttpForward) IsMatch(buf []byte) bool {
	if len(buf)<f.MinLen(){
		return false
	}
	for _, f := range requests {
		//tcp 数据流 的开始  "GET " OR "POST "  ...
		if string(buf[:len(f)+1]) == f+" " {
			//log.Println("match http ok",f)
			return true
		}
	}

	return false

}

// ╰─$ nc -l 0.0.0.0 1234 | xxd                                                                                                                 130 ↵
//00000000: 4745 5420 2f64 666b 6c6a 6c33 3435 2048  GET /dfkljl345 H
//00000010: 5454 502f 312e 310d 0a48 6f73 743a 206c  TTP/1.1..Host: l
//00000020: 6f63 616c 686f 7374 3a31 3233 340d 0a55  ocalhost:1234..U
//00000030: 7365 722d 4167 656e 743a 2063 7572 6c2f  ser-Agent: curl/
//00000040: 372e 3634 2e31 0d0a 4163 6365 7074 3a20  7.64.1..Accept:
// http 的 头部
//curl http://localhost:1234/dfkljl345
// 代理服务器 模式
// curl -x http://0:1234/  http://localhost:1234/dfkljl3
//

// https 请求包分析
//
//测试 服务端
// tcp.port == 1234 and ip.dst == 127.0.0.1
//
// nc -l 0.0.0.0 1234 | xxd
// 00000000: 1603 0100 cd01 0000 c903 03df 42ca 5f9e  ............B._.
// 00000010: 5d0e a86e affa fc44 5f25 20e4 62ef 511e  ]..n...D_% .b.Q.
// 00000020: 2558 bf2b 74bd c375 d47d 1a00 005c c030  %X.+t..u.}...\.0
// 00000030: c02c c028 c024 c014 c00a 009f 006b 0039  .,.(.$.......k.9
// 00000040: cca9 cca8 ccaa ff85 00c4 0088 0081 009d  ................
// 00000050: 003d 0035 00c0 0084 c02f c02b c027 c023  .=.5...../.+.'.#
// 00000060: c013 c009 009e 0067 0033 00be 0045 009c  .......g.3...E..
// 00000070: 003c 002f 00ba 0041 c011 c007 0005 0004  .<./...A........
// 00000080: c012 c008 0016 000a 00ff 0100 0044 000b  .............D..
// 00000090: 0002 0100 000a 0008 0006 001d 0017 0018  ................
// 000000a0: 000d 001c 001a 0601 0603 efef 0501 0503  ................
// 000000b0: 0401 0403 eeee eded 0301 0303 0201 0203  ................
// 000000c0: 0010 000e 000c 0268 3208 6874 7470 2f31  .......h2.http/1
//
// 本地打开
//
// wireShark
// 监听 loopback
// 过滤器
//  tcp.port == 1234 and ip.dst == 127.0.0.1
//
// 测试
//
//  curl https://127.0.0.1:1234/
//
//
//   nc -l 0.0.0.0 1234 | xxd                                                                                                                 130 ↵
//   00000000: 1603 0100 cd01 0000 c903 03a6 351a dfd4  ............5...
//   00000010: 4013 e4a7 3b0d 1bca f60c ee87 9760 7b51  @...;........`{Q
//   00000020: 70e6 89b0 96a1 3379 63b3 1600 005c c030  p.....3yc....\.0
//   00000030: c02c c028 c024 c014 c00a 009f 006b 0039  .,.(.$.......k.9
//   00000040: cca9 cca8 ccaa ff85 00c4 0088 0081 009d  ................
//   00000050: 003d 0035 00c0 0084 c02f c02b c027 c023  .=.5...../.+.'.#
//   00000060: c013 c009 009e 0067 0033 00be 0045 009c  .......g.3...E..
//   00000070: 003c 002f 00ba 0041 c011 c007 0005 0004  .<./...A........
//   00000080: c012 c008 0016 000a 00ff 0100 0044 000b  .............D..
//   00000090: 0002 0100 000a 0008 0006 001d 0017 0018  ................
//   000000a0: 000d 001c 001a 0601 0603 efef 0501 0503  ................
//   000000b0: 0401 0403 eeee eded 0301 0303 0201 0203  ................
//   000000c0: 0010 000e 000c 0268 3208 6874 7470 2f31  .......h2.http/1
//   000000d0: 2e31                                     .1
//
//
//https 请求特征
//
//16 表示 协商 会话
//03 01 表示版本      可变的内容
//00 cd 表示 header 长度 可变
//01 表示 hello
//第1个字节 =16
//第6个字节= 01
//那么就是可以 认为是https 协议

type HttpsForward struct {
	TcpForwarder
}

func NewHttpsForward(addr string) Forwarder {
	return &HttpsForward{
		TcpForwarder{Addr: addr},
	}
}

func (f *HttpsForward) Name() string {
	return reflect.TypeOf(*f).Name()
}


func (f *HttpsForward) MinLen() int {
	return 10
}

func (f *HttpsForward) IsMatch(buf []byte) bool {

	if len(buf)<6{
		return false
	}
	if buf[0] != 0x16 || buf[5] != 1 {
		return false
	}
	return true
}
