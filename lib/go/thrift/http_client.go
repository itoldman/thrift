/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package thrift

import (
	"bytes"
	//"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	//"os"
	"strconv"

	//"strings"
)

type THttpClient struct {
	schema             string
	host               string
	port               int
	response           *http.Response
	body               *bytes.Buffer
	url                string
	requestBuffer      *bytes.Buffer
	header             http.Header
	nsecConnectTimeout int64
	nsecReadTimeout    int64
	SeqId              int32
	method             string
}

type THttpClientTransportFactory struct {
	url    string
	isPost bool
}

func (p *THttpClientTransportFactory) GetTransport(trans TTransport) TTransport {
	if trans != nil {
		t, ok := trans.(*THttpClient)
		if ok && t.url != "" {
			if t.requestBuffer != nil {
				t2, _ := NewTHttpPostClient(t.url)
				return t2
			}
			t2, _ := NewTHttpClient(t.url)
			return t2
		}
	}
	if p.isPost {
		s, _ := NewTHttpPostClient(p.url)
		return s
	}
	s, _ := NewTHttpClient(p.url)
	return s
}

func NewTHttpClientTransportFactory(url string) *THttpClientTransportFactory {
	return &THttpClientTransportFactory{url: url, isPost: false}
}

func NewTHttpPostClientTransportFactory(url string) *THttpClientTransportFactory {
	return &THttpClientTransportFactory{url: url, isPost: true}
}

func NewTHttpClient(urlstr string) (TTransport, error) {
	_, err := url.Parse(urlstr)
	if err != nil {
		return nil, err
	}
	response, err := http.Get(urlstr)
	if err != nil {
		return nil, err
	}
	return &THttpClient{response: response, url: urlstr}, nil
}

func NewTHttpPostClient(urlstr string) (TTransport, error) {
	_, err := url.Parse(urlstr)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, 0, 1024)
	return &THttpClient{url: urlstr, requestBuffer: bytes.NewBuffer(buf), header: http.Header{}}, nil
}

func NewTHttpRPCClient(schema string, host string, port int) (TTransport, error) {
	buf := make([]byte, 0, 1024)
	body := make([]byte, 0, 1024)

	return &THttpClient{schema: schema, host: host, port: port, requestBuffer: bytes.NewBuffer(buf), body: bytes.NewBuffer(body), header: http.Header{}}, nil
}

func (p *THttpClient) SetUrl(path string) error {
	urlstr := fmt.Sprint(p.schema, "://", p.host, ":", p.port, path)
	_, err := url.Parse(urlstr)
	if err != nil {
		return err
	}
	p.url = urlstr
	return nil
}

// Set the HTTP Header for this specific Thrift Transport
// It is important that you first assert the TTransport as a THttpClient type
// like so:
//
// httpTrans := trans.(THttpClient)
// httpTrans.SetHeader("User-Agent","Thrift Client 1.0")
func (p *THttpClient) SetHeader(key string, value string) {
	p.header.Add(key, value)
}

// Get the HTTP Header represented by the supplied Header Key for this specific Thrift Transport
// It is important that you first assert the TTransport as a THttpClient type
// like so:
//
// httpTrans := trans.(THttpClient)
// hdrValue := httpTrans.GetHeader("User-Agent")
func (p *THttpClient) GetHeader(key string) string {
	return p.header.Get(key)
}

// Deletes the HTTP Header given a Header Key for this specific Thrift Transport
// It is important that you first assert the TTransport as a THttpClient type
// like so:
//
// httpTrans := trans.(THttpClient)
// httpTrans.DelHeader("User-Agent")
func (p *THttpClient) DelHeader(key string) {
	p.header.Del(key)
}

func (p *THttpClient) Open() error {
	// do nothing
	return nil
}

func (p *THttpClient) IsOpen() bool {
	return p.response != nil || p.requestBuffer != nil
}

func (p *THttpClient) Peek() bool {
	return p.IsOpen()
}

func (p *THttpClient) Close() error {
	if p.response != nil && p.response.Body != nil {
		err := p.response.Body.Close()
		p.response = nil
		return err
	}
	if p.requestBuffer != nil {
		p.requestBuffer.Reset()
		p.requestBuffer = nil
	}
	return nil
}

func (p *THttpClient) Read(buf []byte) (int, error) {
	fmt.Printf("http client read :%s\n", string(p.body.Bytes()))
	if p.response == nil {
		return 0, NewTTransportException(NOT_OPEN, "Response buffer is empty, no request.")
	}
	length := p.body.Len()
	copy(buf, p.body.Bytes())
	fmt.Printf("read to buf:%s\n", string(buf))
	return length, nil
}

func (p *THttpClient) ReadBody() ([]byte, error) {
	body, err := ioutil.ReadAll(p.response.Body)
	return body, err
}

func (p *THttpClient) ReadByte() (c byte, err error) {
	return readByte(p.response.Body)
}

func (p *THttpClient) Write(buf []byte) (int, error) {
	fmt.Printf("http client Write:%v\n", buf)
	n, err := p.requestBuffer.Write(buf)
	return n, err
}

func (p *THttpClient) WriteByte(c byte) error {
	return p.requestBuffer.WriteByte(c)
}

func (p *THttpClient) WriteString(s string) (n int, err error) {
	return p.requestBuffer.WriteString(s)
}

func (p *THttpClient) Flush() error {
	fmt.Println("Http client flushing")
	client := &http.Client{}
	var req *http.Request
	var err error
	if p.method == "POST" {
		req, err = http.NewRequest("POST", p.url, p.requestBuffer)
	} else if p.method == "GET" {
		req, err = http.NewRequest("GET", p.buildGetUrl(), nil)
	}

	//req, err := http.NewRequest("POST", "http://localhost:9090/config", strings.NewReader("client_id=gl"))

	fmt.Printf("Http request is:%v\n", req)
	if err != nil {
		return NewTTransportExceptionFromError(err)
	}
	//p.header.Add("Content-Type", "application/x-thrift")

	p.header.Add("Content-Type", "application/x-www-form-urlencoded")

	req.Header = p.header
	fmt.Println("Do http request")
	response, err := client.Do(req)
	fmt.Printf("Http response is:%v\n", response)
	if err != nil {
		return NewTTransportExceptionFromError(err)
	}
	if response.StatusCode != http.StatusOK {
		// TODO(pomack) log bad response
		return NewTTransportException(UNKNOWN_TRANSPORT_EXCEPTION, "HTTP Response code: "+strconv.Itoa(response.StatusCode))
	}

	body, err := ioutil.ReadAll(response.Body)
	fmt.Printf("Response body is:%s\n", string(body))

	if err != nil {
		fmt.Printf("Read Response body err:%v\n", err)

		return err
	}
	p.body.Write(body)
	p.response = response
	return nil
}

func (p *THttpClient) buildGetUrl() string {
	return p.url + "?" + p.requestBuffer.String()
}
