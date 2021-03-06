//并发时钟客户端
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "120.77.220.239:8005")
	if err != nil {
		log.Fatal(err)
	}
	done:=make(chan struct{})
	go func(){
		io.Copy(os.Stdout,conn)
		
		done<-struct{}{}
	}()
	mustCopy(conn,os.Stdin)
	conn.Close()	
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
