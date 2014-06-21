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
	//"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

const (
	DEFAULT_URL = "/_rpc_handler_"
)

// Simple, non-concurrent server for testing.
type THTTPServer struct {
	quit           chan struct{}
	networkAddr    string
	processor      TProcessor
	inputProtocol  *THTTPServerProtocol
	outputProtocol *THTTPServerProtocol
}

func NewTHTTPServer4(networkAddr string, processor TProcessor, inputProtocol *THTTPServerProtocol, outputProtocol *THTTPServerProtocol) *THTTPServer {
	return &THTTPServer{
		networkAddr:    networkAddr,
		processor:      processor,
		inputProtocol:  inputProtocol,
		outputProtocol: outputProtocol,
		quit:           make(chan struct{}, 1),
	}
}

func (p *THTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.rpcHandler(w, r)
}

func (p *THTTPServer) Serve() error {
	http.Handle(DEFAULT_URL, p)
	http.ListenAndServe(p.networkAddr, nil)
	return nil
}

func (p *THTTPServer) Stop() error {
	p.quit <- struct{}{}
	//p.serverTransport.Interrupt()
	return nil
}

func (p *THTTPServer) rpcHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	defer func() {
		if e := recover(); e != nil {
			log.Printf("panic in processor: %s: %s", e, debug.Stack())
		}
	}()

	p.inputProtocol.r = r
	p.outputProtocol.w = w
	// handle reqeust
	ok, err := p.processor.Process(p.inputProtocol, p.outputProtocol)
	if err, ok := err.(TTransportException); ok && err.TypeId() == END_OF_FILE {
		return
	} else if err != nil {
		log.Printf("error processing request: %s", err)
		return
	}
	if !ok {
		return
	}

	return
}
