package main

import (
	"fmt"
	"html/template"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

type http_client struct {
	buffer []byte
}

type http_protocol struct {
	client *http_client
	buffer []byte
}

type Json struct {
	Status  string
	Message string
}

func (c *http_client) read(buffer []byte) {
	newbuffer := []byte("12")
	c.buffer = newbuffer
	fmt.Printf("client buffer is:%v\n", c.buffer)

	//buffer = c.buffer
	copy(buffer, c.buffer)
}

func (p *http_protocol) read() {
	p.client.read(p.buffer)
	fmt.Printf("protocol buffer is:%v\n", p.buffer)
	i, _ := strconv.Atoi(string(p.buffer))
	fmt.Printf("i is:%d\n", i)

}

func main() {
	test_copy()
}

func test_copy() {
	client := http_client{buffer: make([]byte, 0, 10)}
	protocol := http_protocol{client: &client, buffer: make([]byte, 10, 10)}
	protocol.read()

}

func enc_dec_url() {
	urlstr := "test encode and decode url"
	s, _ := url.QueryUnescape(urlstr)
	fmt.Printf("unescaped:%s\n", s)
	s2 := url.QueryEscape(urlstr)
	fmt.Printf("escaped:%s\n", s2)
	s3, _ := url.QueryUnescape(s2)
	fmt.Printf("unescaped:%s\n", s3)

}

func test_template() {
	data := Json{"fuck", "fuck"}
	t := template.Must(template.ParseFiles("test.html"))
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}

func test_template2() {
	data := "hello world"
	t := template.Must(template.ParseFiles("test2.html"))
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}

func test_convert() {
	s := string("3")
	fmt.Printf("s:%v\n", s)
	i, _ := strconv.Atoi(s)
	fmt.Printf("s to i:%d\n", i)
}

func test_path() {
	absPath, _ := filepath.Abs("test.go")
	fmt.Printf("abspath:%s\n", absPath)

}

// func test_map() {
// 	m1 := make([]string, 0, 1024)
// 	m1 = append(m1, "hello")
// 	fmt.Printf("m1 is:%v\n", m1)

// 	c1 := http_client{}
// 	c1.buffer = append(c1.buffer, "wonderful")
// 	fmt.Printf("c1.buffer is:%v\n", c1.buffer)

// }
