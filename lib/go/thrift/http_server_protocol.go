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
	//"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	//"encoding/binary"
)

type THTTPServerProtocol struct {
	w          http.ResponseWriter
	r          *http.Request
	fieldIndex int
	SeqId      int32
	buffer     []byte
}

type THTTPServerProtocolFactory struct {
}

func NewTHTTPServerProtocolTransport() *THTTPServerProtocol {
	return NewTHTTPServerProtocol()
}

func NewTHTTPServerProtocol() *THTTPServerProtocol {
	fmt.Println("New HTTP Server protocol")
	p := &THTTPServerProtocol{}
	return p
}

func NewTHTTPServerProtocolFactoryDefault() *THTTPServerProtocolFactory {
	return NewTHTTPServerProtocolFactory()
}

func NewTHTTPServerProtocolFactory() *THTTPServerProtocolFactory {
	return &THTTPServerProtocolFactory{}
}

func (p *THTTPServerProtocolFactory) GetProtocol(t TTransport) TProtocol {
	return NewTHTTPServerProtocol()
}

/**
 * Writing Methods
 */

func (p *THTTPServerProtocol) WriteMessageBegin(name string, typeId TMessageType, seqId int32) error {
	fmt.Println("HTTP server protocol WriteMessageBegin")
	p.fieldIndex = 0
	//p.buffer = p.buffer[:0]
	return nil
}

func (p *THTTPServerProtocol) WriteMessageEnd() error {
	return p.Flush()
	return nil
}

func (p *THTTPServerProtocol) WriteStructBegin(name string) error {
	return nil
}

func (p *THTTPServerProtocol) WriteStructEnd() error {
	return nil
}

func (p *THTTPServerProtocol) WriteFieldBegin(name string, typeId TType, id int16) error {
	//_, err := p.trans.WriteString(name + "=")
	//return err
	return nil
}

func (p *THTTPServerProtocol) WriteFieldEnd() error {
	//_, err := p.trans.WriteString("&")
	//return err
	return nil
}

func (p *THTTPServerProtocol) WriteFieldStop() error {
	return nil
}

func (p *THTTPServerProtocol) WriteMapBegin(keyType TType, valueType TType, size int) error {
	return nil
}

func (p *THTTPServerProtocol) WriteMapEnd() error {
	return nil
}

func (p *THTTPServerProtocol) WriteListBegin(elemType TType, size int) error {
	return nil
}

func (p *THTTPServerProtocol) WriteListEnd() error {
	return nil
}

func (p *THTTPServerProtocol) WriteSetBegin(elemType TType, size int) error {
	return nil
}

func (p *THTTPServerProtocol) WriteSetEnd() error {
	return nil
}

func (p *THTTPServerProtocol) WriteBool(value bool) error {
	return nil
}

func (p *THTTPServerProtocol) WriteByte(value byte) error {
	return nil
}

func (p *THTTPServerProtocol) WriteI16(value int16) error {
	return nil
}

func (p *THTTPServerProtocol) WriteI32(value int32) error {
	fmt.Printf("WriteI32:%d\n", value)
	p.buffer = append(p.buffer, fmt.Sprintf("%d", value)...)
	return nil
}

func (p *THTTPServerProtocol) WriteI64(value int64) error {
	return nil
}

func (p *THTTPServerProtocol) WriteDouble(value float64) error {
	return nil
}

func (p *THTTPServerProtocol) WriteString(value string) error {
	p.buffer = append(p.buffer, value...)
	return nil
}

func (p *THTTPServerProtocol) WriteBinary(value []byte) error {
	return nil
}

/**
 * Reading methods
 */

func (p *THTTPServerProtocol) ReadMessageBegin() (name string, typeId TMessageType, seqId int32, err error) {
	method := p.r.PostFormValue("method")
	seqId1 := p.r.PostFormValue("seq_id")
	seqId2, _ := strconv.Atoi(seqId1)
	seqId = int32(seqId2)
	fmt.Printf("Got seq_id:%d\n", seqId)

	p.SeqId = seqId
	return method, CALL, seqId, nil
}

func (p *THTTPServerProtocol) ReadMessageEnd() error {
	return nil
}

func (p *THTTPServerProtocol) ReadStructBegin() (name string, err error) {
	return
}

func (p *THTTPServerProtocol) ReadStructEnd() error {
	return nil
}

func (p *THTTPServerProtocol) ReadFieldBegin() (name string, typeId TType, fieldId int16, err error) {
	p.fieldIndex++
	name = strconv.Itoa(p.fieldIndex)
	value := p.r.PostFormValue(name)
	if value == "" {
		return "", STOP, 0, nil
	}
	return name, VOID, int16(p.fieldIndex), nil
}

func (p *THTTPServerProtocol) ReadFieldEnd() error {
	return nil
}

//var invalidDataLength = NewTProtocolExceptionWithType(INVALID_DATA, errors.New("Invalid data length"))

func (p *THTTPServerProtocol) ReadMapBegin() (kType, vType TType, size int, err error) {
	return kType, vType, size, nil
}

func (p *THTTPServerProtocol) ReadMapEnd() error {
	return nil
}

func (p *THTTPServerProtocol) ReadListBegin() (elemType TType, size int, err error) {
	return
}

func (p *THTTPServerProtocol) ReadListEnd() error {
	return nil
}

func (p *THTTPServerProtocol) ReadSetBegin() (elemType TType, size int, err error) {
	return elemType, size, nil
}

func (p *THTTPServerProtocol) ReadSetEnd() error {
	return nil
}

func (p *THTTPServerProtocol) ReadBool() (bool, error) {
	return true, nil
}

func (p *THTTPServerProtocol) ReadByte() (value byte, err error) {
	return
}

func (p *THTTPServerProtocol) ReadI16() (value int16, err error) {
	return 0, nil
}

func (p *THTTPServerProtocol) ReadI32() (value int32, err error) {
	i := strconv.Itoa(p.fieldIndex)
	value1 := p.r.PostFormValue(i)
	value2, _ := strconv.Atoi(value1)
	return int32(value2), nil

}

func (p *THTTPServerProtocol) ReadI64() (value int64, err error) {
	return 0, nil
}

func (p *THTTPServerProtocol) ReadDouble() (value float64, err error) {
	return 0, nil
}

func (p *THTTPServerProtocol) ReadString() (value string, err error) {
	i := strconv.Itoa(p.fieldIndex)
	value = p.r.PostFormValue(i)
	return value, nil
}

func (p *THTTPServerProtocol) ReadBinary() ([]byte, error) {
	return []byte(""), nil
}

func (p *THTTPServerProtocol) Flush() (err error) {
	p.w.Header().Set("seq_id", fmt.Sprintf("%d", p.SeqId))
	p.w.Write(p.buffer)
	p.buffer = p.buffer[:0]
	return nil
}

func (p *THTTPServerProtocol) Skip(fieldType TType) (err error) {
	return SkipDefaultDepth(p, fieldType)
}

func (p *THTTPServerProtocol) Transport() TTransport {
	trans, _ := NewTSocket(net.JoinHostPort("localhost", "9090"))
	return trans
}
