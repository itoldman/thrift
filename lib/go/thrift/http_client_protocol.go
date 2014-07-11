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
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	//"encoding/binary"
)

type THTTPProtocol struct {
	trans         TRichTransport
	origTransport TTransport
	reader        io.Reader
	writer        io.Writer
	strictRead    bool
	strictWrite   bool
	buffer        []byte
	count         int32
}

type THTTPProtocolFactory struct {
	strictRead  bool
	strictWrite bool
}

func NewTHTTPProtocolTransport(t TTransport) *THTTPProtocol {
	return NewTHTTPProtocol(t, false, true)
}

func NewTHTTPProtocol(t TTransport, strictRead, strictWrite bool) *THTTPProtocol {
	fmt.Println("New HTTP protocol")
	p := &THTTPProtocol{origTransport: t, strictRead: strictRead, strictWrite: strictWrite}
	if et, ok := t.(TRichTransport); ok {
		p.trans = et
	} else {
		p.trans = NewTRichTransport(t)
	}
	p.reader = p.trans
	p.writer = p.trans
	p.buffer = make([]byte, 1024, 1024)

	return p
}

func NewTHTTPProtocolFactoryDefault() *THTTPProtocolFactory {
	return NewTHTTPProtocolFactory(false, true)
}

func NewTHTTPProtocolFactory(strictRead, strictWrite bool) *THTTPProtocolFactory {
	return &THTTPProtocolFactory{strictRead: strictRead, strictWrite: strictWrite}
}

func (p *THTTPProtocolFactory) GetProtocol(t TTransport) TProtocol {
	return NewTHTTPProtocol(t, p.strictRead, p.strictWrite)
}

/**
 * Writing Methods
 */

func (p *THTTPProtocol) WriteMessageBegin(name string, typeId TMessageType, seqId int32) error {
	if value, ok := p.origTransport.(*THttpClient); ok {
		method, err := p.getMethod(name)
		if err != nil {
			return err
		}
		value.method = method
		value.SetUrl(p.buildUrl(name))
		_, err = p.trans.WriteString(fmt.Sprintf("method=%s&seq_id=%d&", name, seqId))
		return err
	}

	return errors.New("THTTPProtocol can only work with THttpClient transport")

}

func (p *THTTPProtocol) WriteMessageEnd() error {
	return nil
}

func (p *THTTPProtocol) WriteStructBegin(name string) error {
	return nil
}

func (p *THTTPProtocol) WriteStructEnd() error {
	return nil
}

func (p *THTTPProtocol) WriteFieldBegin(name string, typeId TType, id int16) error {
	_, err := p.trans.WriteString(fmt.Sprintf("%d=%s&%s=", id, name, name))
	return err
}

func (p *THTTPProtocol) WriteFieldEnd() error {
	_, err := p.trans.WriteString("&")
	return err
}

func (p *THTTPProtocol) WriteFieldStop() error {
	return nil
}

func (p *THTTPProtocol) WriteMapBegin(keyType TType, valueType TType, size int) error {
	return nil
}

func (p *THTTPProtocol) WriteMapEnd() error {
	return nil
}

func (p *THTTPProtocol) WriteListBegin(elemType TType, size int) error {
	return nil
}

func (p *THTTPProtocol) WriteListEnd() error {
	return nil
}

func (p *THTTPProtocol) WriteSetBegin(elemType TType, size int) error {
	return nil
}

func (p *THTTPProtocol) WriteSetEnd() error {
	return nil
}

func (p *THTTPProtocol) WriteBool(value bool) error {
	return nil
}

func (p *THTTPProtocol) WriteByte(value byte) error {
	return nil
}

func (p *THTTPProtocol) WriteI16(value int16) error {
	return nil
}

func (p *THTTPProtocol) WriteI32(value int32) error {
	_, err := p.trans.WriteString(fmt.Sprintf("%d", value))
	return err
}

func (p *THTTPProtocol) WriteI64(value int64) error {
	return nil
}

func (p *THTTPProtocol) WriteDouble(value float64) error {
	return nil
}

func (p *THTTPProtocol) WriteString(value string) error {
	_, err := p.trans.WriteString(value)
	return err
}

func (p *THTTPProtocol) WriteBinary(value []byte) error {
	return nil
}

/**
 * Reading methods
 */
func (p *THTTPProtocol) ReadMessageBegin2(map[string][]string) (name string, typeId TMessageType, seqid int32, err error) {
	return "", 1, 1, nil
}

func (p *THTTPProtocol) ReadMessageBegin() (name string, typeId TMessageType, seqId int32, err error) {

	if value, ok := p.origTransport.(*THttpClient); ok {
		seq1 := value.response.Header.Get("seq_id")
		seq2, _ := strconv.Atoi(seq1)
		fmt.Printf("response seq_id is:%d\n", seq2)
		if seq2 == 0 {
			seq2++
		}
		return "", 1, int32(seq2), nil
	}

	return "", 1, 0, errors.New("THTTPProtocol can only work with THttpClient transport")

}

func (p *THTTPProtocol) ReadMessageEnd() error {
	return nil
}

func (p *THTTPProtocol) ReadStructBegin() (name string, err error) {
	return
}

func (p *THTTPProtocol) ReadStructEnd() error {
	return nil
}

func (p *THTTPProtocol) ReadFieldBegin() (name string, typeId TType, fieldId int16, err error) {
	p.count++
	//fmt.Printf("buffer length is:%d\n", len(p.buffer))
	if p.count > 1 {
		return "", STOP, 0, nil
	}
	// if len(p.buffer) > 0 {
	// 	return "", 0, 0, nil
	// }
	return "", VOID, 0, nil
}

func (p *THTTPProtocol) ReadFieldEnd() error {
	return nil
}

//var invalidDataLength = NewTProtocolExceptionWithType(INVALID_DATA, errors.New("Invalid data length"))

func (p *THTTPProtocol) ReadMapBegin() (kType, vType TType, size int, err error) {
	return kType, vType, size, nil
}

func (p *THTTPProtocol) ReadMapEnd() error {
	return nil
}

func (p *THTTPProtocol) ReadListBegin() (elemType TType, size int, err error) {
	return
}

func (p *THTTPProtocol) ReadListEnd() error {
	return nil
}

func (p *THTTPProtocol) ReadSetBegin() (elemType TType, size int, err error) {
	return elemType, size, nil
}

func (p *THTTPProtocol) ReadSetEnd() error {
	return nil
}

func (p *THTTPProtocol) ReadBool() (bool, error) {
	return true, nil
}

func (p *THTTPProtocol) ReadByte() (value byte, err error) {
	return
}

func (p *THTTPProtocol) ReadI16() (value int16, err error) {
	return 0, nil
}

func (p *THTTPProtocol) ReadI32() (value int32, err error) {
	len, _ := p.trans.Read(p.buffer)
	//fmt.Printf("Http protocol ReadI32, p.buffer:%v\n", p.buffer)

	s := string(p.buffer[:len])
	fmt.Printf("Http protocol ReadI32, s:%v\n", s)

	i, _ := strconv.Atoi(s)
	fmt.Printf("Http protocol ReadI32, i:%d\n", i)

	return int32(i), nil

}

func (p *THTTPProtocol) ReadI64() (value int64, err error) {
	return 0, nil
}

func (p *THTTPProtocol) ReadDouble() (value float64, err error) {
	return 0, nil
}

func (p *THTTPProtocol) ReadString() (value string, err error) {
	//_, e := io.ReadFull(p.trans, p.buffer)

	p.trans.Read(p.buffer)
	fmt.Printf("Http protocol ReadString:%s\n", string(p.buffer))

	return string(p.buffer), nil
}

func (p *THTTPProtocol) ReadBinary() ([]byte, error) {
	return []byte(""), nil
}

func (p *THTTPProtocol) Flush() (err error) {
	return NewTProtocolException(p.trans.Flush())
}

func (p *THTTPProtocol) Skip(fieldType TType) (err error) {
	return SkipDefaultDepth(p, fieldType)
}

func (p *THTTPProtocol) Transport() TTransport {
	return p.origTransport
}

func (p *THTTPProtocol) readAll(buf []byte) error {
	_, err := io.ReadFull(p.reader, buf)
	return NewTProtocolException(err)
}

func (p *THTTPProtocol) checkMethod(s string) bool {
	l := []string{"get", "post", "del", "put"}
	for i := 0; i < len(l); i++ {
		if strings.HasSuffix(s, l[i]) {
			return true
		}

	}
	return false

}

func (p *THTTPProtocol) getMethod(s string) (string, error) {
	l := []string{"get", "post", "del", "put"}
	for i := 0; i < len(l); i++ {
		if strings.HasSuffix(s, l[i]) {
			return strings.ToUpper(l[i]), nil
		}

	}
	return "", errors.New("Not end with get, post, del or put")
}

func (p *THTTPProtocol) buildUrl(s string) string {
	l := strings.Split(s, "_")
	return "/" + strings.Join(l[:len(l)-1], "/")
}
