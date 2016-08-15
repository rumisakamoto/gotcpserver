package main
import (
    "fmt"
    "net"
    "os"
    _ "github.com/lib/pq"
    "database/sql"
    "strconv"
)
const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
    DB_TYPE = "postgres"
    DB_NAME = "testdb" // TODO change
    DB_USER = "postgres"
    DB_PSWD = "postgres"
    DB_HOST = "localhost"
    DB_QUERY = "SELECT id FROM testtable WHERE id=$1" // TODO change
)
type Record struct {
    id int
}
func main() {

    // listen incoming connections.
    l, err := net.Listen(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)
    checkError(err, "Listening error: %s")
    // close the listener when the application closes.
    defer l.Close()

    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
    for {
        // listen an incoming connection
        conn, err := l.Accept()
        checkError(err, "Accepting error: %s")
        // handle connections in a new goroutine
        go handleRequest(conn)
    }
}
func handleRequest(conn net.Conn) {
    // make a buffer to hold incoming data
    buf := make([]byte, 1024)

    // read the incoming connection into the buffer
    reqLen, err := conn.Read(buf)
    checkError(err, "Incoming connection reading error: %s")

    receivedData := string(buf[:reqLen])
    fmt.Println(receivedData)

    data := getDeviceData("1")

    // send a response
    conn.Write([]byte(strconv.Itoa(data.id)))

    conn.Close()
}
func getDeviceData(id string) Record {
    // open databse connection.
    db, err := sql.Open(DB_TYPE, "user="+DB_USER+" dbname="+DB_NAME+" password="+DB_PSWD+" host="+DB_HOST+" sslmode=disable")
    checkError(err, "Opening database connection is failed: %s")

    err = db.Ping()
    checkError(err, "Establishing database connection is failed: %s")

    query, err := db.Prepare(DB_QUERY)
    checkError(err, "Prepareing query statement is failed: %s")

    var r Record
    err = query.QueryRow(id).Scan(&r.id)
    checkError(err, "Getting data is failed: %s")

    db.Close()

    return r
}
func checkError(err error, msg string) {
    if err != nil {
      //fmt.Fprintf(os.Stderr, msg, err.Error())
      fmt.Println(msg, err.Error())
      os.Exit(1)
    }
}
