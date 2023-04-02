package main

import (
	"log"
	"os"
	"time"
)

type Server struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

type Option func(*Server)

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithLogger(logger *log.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

func New(host string, port int, options ...Option) *Server {
	s := &Server{
		host: host,
		port: port,
	}

	for _, option := range options {
		option(s)
	}

	return s
}

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)

	s := New("localhost", 8080, WithTimeout(10*time.Second), WithLogger(logger))
	log.Printf("%#v\n", s)
}
