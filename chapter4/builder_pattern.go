package main

import (
	"log"
	"os"
	"time"
)

type Server struct {
	param serverParam
}

type serverParam struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

func NewBuilder(host string, port int) *serverParam {
	return &serverParam{
		host: host,
		port: port,
	}
}

func (sp *serverParam) Timeout(timeout time.Duration) *serverParam {
	sp.timeout = timeout
	return sp
}

func (sp *serverParam) Logger(logger *log.Logger) *serverParam {
	sp.logger = logger
	return sp
}

func (sp *serverParam) Build() *Server {
	return &Server{
		param: *sp,
	}
}

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)

	s := NewBuilder("localhost", 8080).Timeout(10 * time.Second).Logger(logger).Build()

	log.Printf("%#v\n", s)
}
