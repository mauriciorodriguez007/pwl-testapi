package main

import (
	"fmt"
	"github.com/Microsoft/go-winio"
	"log"
	"net"
	"os"
	"time"
)

var PipeConReader net.Conn
var PipeConWriter net.Conn

func main() {
	//go WritePipe("My message")
	//go ExampleDial()
	requestPath := `\\.\pipe\CPE_REQUEST`
	responsePath := `\\.\pipe\CPE_RESPONSE`

	PipeConWriter = NamedPipeServer(requestPath, "writer")
	PipeConReader = NamedPipeServer(responsePath, "reader")
	
	go PipeReader(PipeConReader)	
	time.Sleep(10 * time.Second)

	n, err := PipeConWriter.Write([]byte("this is my message 23423 2434 234 234 234 234 234 2343 2423 423 432432432 4"))
	if err != nil {
		fmt.Printf("\n%s PipeConWriter.write failed : %s\n",time.Now(),err)
		return
	}
	fmt.Printf("\n%s PipeConWriter Snd: %d bytes",time.Now(),n)

	wait := make(chan string)
	<-wait
	fmt.Printf("exiting main function")
}

func PipeReader(c net.Conn) {
	fmt.Printf("\n%s PipeReader client CONNECTED [%s]",time.Now() ,c.RemoteAddr().Network())
	buf := make([]byte, 4196)
	for {
		n, err := c.Read(buf)
		fmt.Printf("\n nPipeReader rcvd: %d bytes\n", n)
		if err != nil {
			str := string(buf)
			fmt.Printf(".read failed error: %s\n", str)
			return
		}
	}
}

// SERVER role - it listens for 'clients' to connet to it
func NamedPipeServer(pipePath string, role string) net.Conn {

	me := fmt.Sprintf("NamedPipeServer(%s)",pipePath)
	fmt.Printf("\n%s %s listening role: %s", time.Now(),me,role)
	if err := os.RemoveAll(pipePath); err != nil {
		log.Fatal(err)
	}
	pc := &winio.PipeConfig{
		SecurityDescriptor: "D:P(A;;GA;;;AU)",
		InputBufferSize:    512,
		OutputBufferSize:   512,
		MessageMode:        true,
	}
	l, err := winio.ListenPipe(pipePath, pc)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()
	fmt.Printf("\n%s %s waiting for client", time.Now(),me)
	conn, err := l.Accept()
	if err != nil {
		log.Fatal("\n%s %s Accept failed err: ", time.Now(),me,err)
	}
	fmt.Printf("\n%s %s received new connection", time.Now(),me)
	return conn
}
