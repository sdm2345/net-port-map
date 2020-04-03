package port_map

import (
	"context"

	"io"
	"log"
	"net"
)

type Server struct {
	matches []Forwarder
}

func (s *Server) Add(match Forwarder) {

	s.matches = append(s.matches, match)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (s *Server) Listen(ctx context.Context, addr string) {

	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("error", err)
	}
	go func() {
		<-ctx.Done()
		log.Println("close add",addr)
		_=conn.Close()
	}()
	//最大长度即可
	max := 0
	for _, f := range s.matches {
		max = getMax(max, f.MinLen())
	}
	buf := make([]byte, max)
	for {
		conn, err := conn.Accept()
		if err != nil {
			log.Println("accept error", err)
			continue
		}

		n, err := io.ReadFull(conn, buf)
		if err != nil {
			log.Println("read error", err)
			continue
		}

		for _, f := range s.matches {

			if n >= f.MinLen() && f.IsMatch(buf[:n]) {
				//log.Println("match ok", f.Name())
				go func() {
					err := f.Forward(ctx, buf[:n], conn)
					if err != nil {
						log.Println("forward get error", err)
					}
				}()
				break
			}
		}
	}

}
