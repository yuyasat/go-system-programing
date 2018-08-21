package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

var contents = []string{
	"これは、わたしが小さいときに、村の茂兵（もへい）というおじいさんからきいたお話です。",
	"むかしは、わたしたちの村のちかくの、中山というところに小さなお城（しろ）があって、",
	"中山さまというおとのさまがおられたそうです。",
	"その中山から、すこしはなれた山の中に、「ごんぎつね」というきつねがいました。",
	"ごんは、ひとりぼっちの小ぎつねで、しだのいっぱいしげった森の中に穴（あな）をほって住んでいました。",
	"そして、夜でも昼でも、あたりの村へ出ていって、いたずらばかりしました。",
}

func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()
	for {
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain",
			"Transfer-Encoding: chunked",
			"", ""}, "\r\n"))
		for _, content := range contents {
			bytes := []byte(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
		}
		fmt.Fprintf(conn, "0\r\n\r\n")
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(conn)
	}
}
