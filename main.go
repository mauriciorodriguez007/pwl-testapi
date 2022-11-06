package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	//"github.com/Microsoft/go-winio"
	"github.com/Microsoft/go-winio"
	"gopkg.in/natefinch/npipe.v2"
)

func main() {
	go WritePipe("My message")
	//go ExampleDial()

	// pipePath := `\\.\pipe\mypipename`

	// if err := os.RemoveAll(pipePath); err != nil {
	// 	log.Fatal(err)
	// }

	// pc := &winio.PipeConfig{
	// 	SecurityDescriptor: "D:P(A;;GA;;;AU)",
	// 	InputBufferSize:    512,
	// 	OutputBufferSize:   512,
	// 	MessageMode: true,
	// }

	// l, err := winio.ListenPipe(pipePath, pc)

	// if err != nil {
	// 	log.Fatal("listen error:", err)
	// }
	// defer l.Close()

	// for {
	// 	fmt.Printf("waiting on message")
	// 	conn, err := l.Accept()
	// 	if err != nil {
	// 		log.Fatal("accept error:", err)
	// 	}
	// 	fmt.Println("got a connection - dispatching to handler")
	// 	go printMsgServer(conn)
	// }
	//ConnectPipe()

	wait := make(chan string)
	<-wait
	//fmt.Printf("exiting main function")
}

// Use Dial to connect to a server and read messages from it.
func ExampleDial() {
	pipePath := "\\\\.\\pipe\\DIGI3A"
	//npipe.OpenFile()
	//conn, err := npipe.(pipePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	conn, err := npipe.Dial(pipePath)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		// handle error
	}
	for {
		if _, err := fmt.Fprintln(conn, "Hi server!"); err != nil {
			// handle error
			log.Fatalf("error writing  file: %v", err)
		}
	}

	// r := bufio.NewReader(conn)
	// msg, err := r.ReadString('\n')
	// if err != nil {
	// 	// handle eror
	// }
	//	fmt.Println(msg)
}

func WritePipe(msg string) {
	fmt.Printf("waiting 2 seconds")
	time.Sleep(2 * time.Second)
	pipePath := `\\.\pipe\DIGI3A`
	fmt.Printf("flag: %d", os.ModeNamedPipe)
	//f, err := os.OpenFile(pipePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	f,err := winio.DialPipe(pipePath,nil)
	//f, err := os.OpenFile(pipePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	buffer := []byte{0x30, 0x31, 0x33, 0x34}
	for {
		fmt.Printf("\nwriting to pipe")
		//f.WriteString("message from client\n")
		b, err := f.Write(buffer)
		fmt.Printf("\n wrote bytes: %d", b)
		if err != nil {
			fmt.Printf("\n error writing error: %s", err)
		}

		time.Sleep(1 * time.Second)
	}
	f.Close()
	fmt.Println("done writing to pipe")
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
