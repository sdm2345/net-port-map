package main

import (
    "context"
    "flag"
    "fmt"
    "log"
    "net/url"
    "os"
    "os/signal"
    "runtime"
    "strings"
    "syscall"
    "time"
)
import "github.com/sdm2345/net-port-map/port_map"

type ForwardFlag []string

func (i *ForwardFlag) String() string {
    return fmt.Sprint(*i)
}

func (i *ForwardFlag) Set(value string) error {
    *i = append(*i, value)
    return nil
}

func usage() {
    flag.Usage()
    fmt.Println(strings.TrimSpace(`
net-port-map -l tcp://0.0.0.0:80 -f http://127.0.0.1:8080 -f tcp://127.0.0.1:4456
supported forward schema 
http,https,redis,tcp
	`))
}

func main() {
    
    log.SetFlags(log.Lshortfile)
    addr := flag.String("l", "", "-l tcp://0.0.0.0:80 local listen addr")
    var rules ForwardFlag
    flag.Var(&rules, "f", "-f tcp://127.0.0.1:3306 forward target")
    flag.Parse()
    
    if len(rules) == 0 || *addr == "" {
        usage()
        return
    }
    
    listenAddr, err := url.Parse(*addr)
    if err != nil {
        log.Fatal("error listen addr", *addr, err)
    }
    if listenAddr.Scheme != "tcp" {
        log.Fatal("unsupported schema", listenAddr.Scheme)
    }
    
    s := port_map.Server{}
    for _, rule := range rules {
        
        info, err := url.Parse(rule)
        if err != nil {
            log.Fatal("error url", err)
        }
        switch info.Scheme {
        case "http":
            s.Add(port_map.NewHttpForward(info.Host))
        case "ssh":
            s.Add(port_map.NewSshForward(info.Host))
        case "https":
            s.Add(port_map.NewHttpsForward(info.Host))
        case "redis":
            s.Add(port_map.NewRedisForward(info.Host))
        case "tcp":
            s.Add(port_map.NewTcpForwarder(info.Host))
        // todo:需要适配其他协议,加接口实现
        default:
            fmt.Println("unsupported forward schema", rule)
            usage()
            return
        }
    }
    
    ctx, cancel := context.WithCancel(context.Background())
    
    c := make(chan os.Signal, 2)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    
    go s.Listen(ctx, listenAddr.Host)
    //go monitor(ctx)
    <-c
    cancel()
    os.Exit(0)
}

func monitor(ctx context.Context) {
    //test
    for {
        select {
        case <-time.After(time.Second):
            time.Sleep(time.Second)
            log.Println("NumGoroutine:", runtime.NumGoroutine())
        
        case <-ctx.Done():
            {
                break
            }
            
        }
    }
}
