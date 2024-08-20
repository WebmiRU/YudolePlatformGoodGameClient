package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
	"time"
)

var WriteChan = make(chan any, 999)

type Connection struct {
	Conn           net.Conn
	IsConnected    bool
	stopWriterChan chan bool
}

func (conn *Connection) startWriter() {
	conn.stopWriterChan = make(chan bool)

loop:
	for {
		select {
		case message := <-WriteChan:
			if err := json.NewEncoder(conn.Conn).Encode(message); err != nil {
				log.Println("Error encoding message:", err, err.Error(), message)
			}
		case <-conn.stopWriterChan:
			log.Println("Stop Writer")
			break loop
		}
	}
}

func (conn *Connection) Open() {
	go conn.startWriter()
}

func (conn *Connection) Close() {
	conn.stopWriterChan <- true
}

func reconnect(server string, port int, message any) {
	fmt.Println("Reconnection after 5 seconds...")
	time.Sleep(5 * time.Second)
	Connect(server, port, message)
}

func Connect(server string, port int, message any) {
	defer reconnect(server, port, message)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", server, port))

	if err != nil {
		log.Println("YudolePlatform server connection error:", err.Error())
		return
	}

	c := &Connection{
		Conn: conn,
	}

	c.Open()

loop:
	for {
		if err := json.NewDecoder(conn).Decode(&message); err != nil {
			if _, ok := err.(net.Error); ok {
				break loop
			}

			switch err {
			case io.EOF, net.ErrClosed:
				break loop
			default:
				log.Println(reflect.TypeOf(err), err)
				log.Println("YudolePlatform server message encode error:", err.Error())
			}
		}

		fmt.Println(message)
	}

	log.Println("Disconnected from YudolePlatform server")

	c.Close()
	conn.Close()
}
