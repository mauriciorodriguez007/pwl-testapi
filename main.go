package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/Microsoft/go-winio"
)

func main() {
	pipePath := `\\.\pipe\mypipename`

	if err := os.RemoveAll(pipePath); err != nil {
		log.Fatal(err)
	}

	pc := &winio.PipeConfig{
		SecurityDescriptor: "D:P(A;;GA;;;AU)",
		InputBufferSize:    512,
		OutputBufferSize:   512,
		MessageMode: true,
	}

	l, err := winio.ListenPipe(pipePath, pc)

	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()

	go WritePipe("My message")
	for {
		fmt.Printf("waiting on message")
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		fmt.Println("got a connection - dispatching to handler")
		go printMsgServer(conn)
	}
	//ConnectPipe()

	wait := make(chan string)
	<-wait
	//fmt.Printf("exiting main function")
}

func printMsgServer(c net.Conn) {
	fmt.Println("in the handler")
	log.Printf("Client connected [%s]", c.RemoteAddr().Network())

	buf := make([]byte, 512)
	for true {
		fmt.Printf("read blocking...")
		n, err := c.Read(buf)
		fmt.Printf("read %d bytes\n", n)
		if err != nil {
			str := string(buf)
			fmt.Printf("got message: %s\n", str)
		}
	}

}
func ConnectPipe() {

}

func ReadPipe() {

}

func WritePipe(msg string) {
	fmt.Printf("waiting 10 seconds")
	time.Sleep(10 * time.Second)
	pipePath := `\\.\pipe\mypipename`
	f, err := os.OpenFile(pipePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	for true {
		fmt.Printf("writing to pipe")
		f.WriteString("message from client\n")		
		time.Sleep(1 * time.Second)		
	}
	f.Close()
	fmt.Println("done writing to pipe")
}
