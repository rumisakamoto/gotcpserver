package main
import (
    "fmt"
    "net"
    "os"
)
const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)
func main() {
    // listen incoming connections.
    l, err := net.Listen(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)
    if err != nil {
        fmt.Println("Listening error: ", err.Error())
        os.Exit(1)
    }
    // close the listener when the application closes.
    defer l.Close()

    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
    for {
        // listen an incoming connection
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Accepting error: ", err.Error())
            os.Exit(1)
        }
        // handle connections in a new goroutine
        go handleRequest(conn)
    }
}
func handleRequest(conn net.Conn) {
    // make a buffer to hold incoming data
    buf := make([]byte, 1024)

    // read the incoming connection into the buffer
    reqLen, err := conn.Read(buf)
    if err != nil {
        fmt.Println("Incoming connection reading error :", err.Error())
    }
    receivedData := string(buf[:reqLen])
    fmt.Println(receivedData)

    // send a response
    conn.Write([]byte("Message received."))

    conn.Close()
}
