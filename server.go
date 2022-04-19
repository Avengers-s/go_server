package main

import(
    "fmt"
    "net"
)

type Server struct{
    Ip string
    Port int
}

func NewServer(ip string, port int) *Server{
    server := &Server{ip, port}
    return server
}

func (s *Server) Handler(conn net.Conn){
    fmt.Println("连接成功!!!")
}

func (s *Server) Start(){
    listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d",s.Ip,s.Port))
    if err != nil {
        fmt.Println("net.Listen err:",err)
        return 
    }
    defer listener.Close()

    for{
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("listener accept err:",err)
            continue
        }

        go s.Handler(conn)
    }
}
