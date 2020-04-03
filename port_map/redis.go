package port_map

import (
	"log"
	"reflect"
	"regexp"
)

/**
redis 协议的 头部
命令行 链接
*1
$7
COMMAND

认证
*2
$4
AUTH

特征很明显 , 行协议
*2
$3
GET
$3
abc

特征明显
*<参数数量> CR LF
$<参数1 的字节数量> CR LF

第一次 链接

*/

type RedisForward struct {
	TcpForwarder
}

func NewRedisForward(addr string) Forwarder {
	return &RedisForward{
		TcpForwarder: TcpForwarder{Addr: addr},
	}
}

var redisReg *regexp.Regexp

func init() {
	var err error
	redisReg, err = regexp.Compile("^\\*[0-9]+\r\n\\$[0-9]+")
	if err != nil {
		log.Fatal(err)
	}
}


func (f *RedisForward) Name() string {
	return reflect.TypeOf(*f).Name()
}

func (f *RedisForward) MinLen() int {
	//8 字节 可以探测出来了
	return 8
}

func (f *RedisForward) IsMatch(buf []byte) bool {

	return redisReg.Match(buf)
}
