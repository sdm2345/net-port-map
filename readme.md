端口复用的服务
============
简单同一个端口,提供多种服务,根据协议特征
把数据流转发到相应的后端



比如监听端口80 
转发到 后端 redis 6379 
和 http 服务 :8080
和 https 服务 :1443
```bash
go get github.com/sdm2345/net-port-map
net-port-map \
    -l tcp://0.0.0.0:80 \
    -f redis://127.0.0.1:6379 \
    -f http://127.0.0.1:8080 
    -f https://127.0.0.1:1443 

# 下列测试用例
redis-cli -h 127.0.0.1 -p 80 incr a
curl http://locahost:80/hello
curl https://domain:80/hello
    
#未适配的tcp 协议放到最后即可

 
net-port-map \
    -l tcp://0.0.0.0:80 \
    -f redis://127.0.0.1:6379 \
    -f http://127.0.0.1:8080 
    -f https://127.0.0.1:1443 
    -f tcp://127.0.0.1:1080 

go run main.go  -l tcp://0.0.0.0:80 -f http://127.0.0.1:8080 -f https://127.0.0.1:7788 -f redis://127.0.0.1:6379  -f tcp://127.0.0.1:4456
```    

```bash 
使用 nmap 探测协议版本行为
sudo nmap -sV 127.0.0.1 -p 80                                                                                                            130 ↵
Starting Nmap 7.80 ( https://nmap.org ) at 2020-04-03 16:52 CST
Nmap scan report for localhost (127.0.0.1)
Host is up (0.00015s latency).

PORT   STATE SERVICE VERSION
80/tcp open  http    nginx 1.17.9
sudo nmap -sV 127.0.0.1 -p 6379
Starting Nmap 7.80 ( https://nmap.org ) at 2020-04-03 16:52 CST
Nmap scan report for localhost (127.0.0.1)
Host is up (0.00013s latency).

PORT     STATE SERVICE VERSION
6379/tcp open  redis   Redis key-value store 5.0.8



```