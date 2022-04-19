package main

import(
    "fmt"
    "net"
    "sync"
)

type Server struct{
    Ip string
    Port int
    OnlineMap map[string]*User
    Message chan string
    MapLock sync.RWMutex
}

func NewServer(ip string, port int) *Server{
    server := &Server{
        Ip: ip,
        Port: port,
        OnlineMap: make(map[string]*User),
        Message: make(chan string),
    }
    return server
}

func (s *Server) Handler(conn net.Conn){
    user := NewUser(conn)
    s.MapLock.Lock()
    s.OnlineMap[user.Name] = user
    s.MapLock.Unlock()

    s.BoardCast(user, "已上线")
    select{}
}

func (s *Server) BoardCast(user *User, msg string){
    sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

    s.Message <- sendMsg
}

func (s *Server) ListenMessage (){
    for{
        msg := <-s.Message
        s.MapLock.Lock()
        for _, cli := range s.OnlineMap{
            cli.C <- msg
        }
        s.MapLock.Unlock()
    }
}

func (s *Server) Start(){
    listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d",s.Ip,s.Port))
    if err != nil {
        fmt.Println("net.Listen err:",err)
        return 
    }
    defer listener.Close()
    go s.ListenMessage()
    for{
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("listener accept err:",err)
            continue
        }

        go s.Handler(conn)
    }
}
