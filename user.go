package main

import(
    "net"
)

type User struct{
    Name string
    Addr string
    C chan string
    conn net.Conn
}

func (u *User)ListenMessage() {
    for{
        msg := <-u.C
        u.conn.Write([]byte(msg + "\n"))
    }
}

func NewUser(conn net.Conn) *User{
    userAddr := conn.RemoteAddr().String()
    user := &User{
        Name: userAddr,
        Addr: userAddr,
        C: make(chan string),
        conn: conn,
    }
    go user.ListenMessage()
    return user
}


