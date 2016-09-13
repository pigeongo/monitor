package monitor

import (
    "net"
    "time"
)

func Diagnose(conn *net.TCPConn, buffer []byte, timeout time.Duration) {
    reader := make(chan byte)
    go func(c chan byte, t time.Duration) {
        select {
        case <-c:
            conn.SetDeadline(time.Now().Add(t))
        case <-time.After(t):
            conn.Close()
        }
    }(reader, timeout)
    go func(buffer []byte, c chan byte) {
        for _, b := range buffer {
            c <- b
        }
        close(c)
    }(buffer, reader)
}