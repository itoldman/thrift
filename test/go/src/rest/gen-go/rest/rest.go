// Autogenerated by Thrift Compiler (0.9.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package rest

import (
	"fmt"
	"math"
	"thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = math.MinInt32
var _ = thrift.ZERO
var _ = fmt.Printf

type Rest interface {
	// Parameters:
	//  - ClientId
	Config(client_id string) (r string, err error)
	// Parameters:
	//  - Num1
	//  - Num2
	Add(num1 int32, num2 int32) (r int32, err error)
	// Parameters:
	//  - Value1
	//  - Value2
	//  - Value3
	//  - Value4
	//  - Value5
	//  - Value6
	//  - Value7
	Test1(value1 bool, value2 int16, value3 int64, value4 float64, value5 []byte, value6 map[string]string, value7 []int32) (r map[string]bool, err error)
}

type RestClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewRestClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *RestClient {
	return &RestClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewRestClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *RestClient {
	return &RestClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - ClientId
func (p *RestClient) Config(client_id string) (r string, err error) {
	if err = p.sendConfig(client_id); err != nil {
		return
	}
	return p.recvConfig()
}

func (p *RestClient) sendConfig(client_id string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("config", thrift.CALL, p.SeqId)
	args0 := NewConfigArgs()
	args0.ClientId = client_id
	err = args0.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *RestClient) recvConfig() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result1 := NewConfigResult()
	err = result1.Read(iprot)
	iprot.ReadMessageEnd()
	value = result1.Success
	return
}

// Parameters:
//  - Num1
//  - Num2
func (p *RestClient) Add(num1 int32, num2 int32) (r int32, err error) {
	if err = p.sendAdd(num1, num2); err != nil {
		return
	}
	return p.recvAdd()
}

func (p *RestClient) sendAdd(num1 int32, num2 int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("add", thrift.CALL, p.SeqId)
	args4 := NewAddArgs()
	args4.Num1 = num1
	args4.Num2 = num2
	err = args4.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *RestClient) recvAdd() (value int32, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error6 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error7 error
		error7, err = error6.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error7
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result5 := NewAddResult()
	err = result5.Read(iprot)
	iprot.ReadMessageEnd()
	value = result5.Success
	return
}

// Parameters:
//  - Value1
//  - Value2
//  - Value3
//  - Value4
//  - Value5
//  - Value6
//  - Value7
func (p *RestClient) Test1(value1 bool, value2 int16, value3 int64, value4 float64, value5 []byte, value6 map[string]string, value7 []int32) (r map[string]bool, err error) {
	if err = p.sendTest1(value1, value2, value3, value4, value5, value6, value7); err != nil {
		return
	}
	return p.recvTest1()
}

func (p *RestClient) sendTest1(value1 bool, value2 int16, value3 int64, value4 float64, value5 []byte, value6 map[string]string, value7 []int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("test1", thrift.CALL, p.SeqId)
	args8 := NewTest1Args()
	args8.Value1 = value1
	args8.Value2 = value2
	args8.Value3 = value3
	args8.Value4 = value4
	args8.Value5 = value5
	args8.Value6 = value6
	args8.Value7 = value7
	err = args8.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *RestClient) recvTest1() (value map[string]bool, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error10 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error11 error
		error11, err = error10.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error11
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result9 := NewTest1Result()
	err = result9.Read(iprot)
	iprot.ReadMessageEnd()
	value = result9.Success
	return
}

type RestProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      Rest
}

func (p *RestProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *RestProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *RestProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewRestProcessor(handler Rest) *RestProcessor {

	self12 := &RestProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self12.processorMap["config"] = &restProcessorConfig{handler: handler}
	self12.processorMap["add"] = &restProcessorAdd{handler: handler}
	self12.processorMap["test1"] = &restProcessorTest1{handler: handler}
	return self12
}

func (p *RestProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x13 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x13.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x13

}

type restProcessorConfig struct {
	handler Rest
}

func (p *restProcessorConfig) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewConfigArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("config", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewConfigResult()
	if result.Success, err = p.handler.Config(args.ClientId); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing config: "+err.Error())
		oprot.WriteMessageBegin("config", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("config", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type restProcessorAdd struct {
	handler Rest
}

func (p *restProcessorAdd) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewAddArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("add", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewAddResult()
	if result.Success, err = p.handler.Add(args.Num1, args.Num2); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing add: "+err.Error())
		oprot.WriteMessageBegin("add", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("add", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type restProcessorTest1 struct {
	handler Rest
}

func (p *restProcessorTest1) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewTest1Args()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("test1", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewTest1Result()
	if result.Success, err = p.handler.Test1(args.Value1, args.Value2, args.Value3, args.Value4, args.Value5, args.Value6, args.Value7); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing test1: "+err.Error())
		oprot.WriteMessageBegin("test1", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("test1", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type ConfigArgs struct {
	ClientId string `thrift:"client_id,1,required"`
}

func NewConfigArgs() *ConfigArgs {
	return &ConfigArgs{}
}

func (p *ConfigArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *ConfigArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.ClientId = v
	}
	return nil
}

func (p *ConfigArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("config_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *ConfigArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("client_id", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:client_id: %s", p, err)
	}
	if err := oprot.WriteString(string(p.ClientId)); err != nil {
		return fmt.Errorf("%T.client_id (1) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:client_id: %s", p, err)
	}
	return err
}

func (p *ConfigArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ConfigArgs(%+v)", *p)
}

type ConfigResult struct {
	Success string `thrift:"success,0"`
}

func NewConfigResult() *ConfigResult {
	return &ConfigResult{}
}

func (p *ConfigResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *ConfigResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 0: %s")
	} else {
		p.Success = v
	}
	return nil
}

func (p *ConfigResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("config_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	switch {
	default:
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *ConfigResult) writeField0(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
		return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Success)); err != nil {
		return fmt.Errorf("%T.success (0) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 0:success: %s", p, err)
	}
	return err
}

func (p *ConfigResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ConfigResult(%+v)", *p)
}

type AddArgs struct {
	Num1 int32 `thrift:"num1,1,required"`
	Num2 int32 `thrift:"num2,2,required"`
}

func NewAddArgs() *AddArgs {
	return &AddArgs{}
}

func (p *AddArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *AddArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.Num1 = v
	}
	return nil
}

func (p *AddArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s")
	} else {
		p.Num2 = v
	}
	return nil
}

func (p *AddArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("add_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *AddArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("num1", thrift.I32, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:num1: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Num1)); err != nil {
		return fmt.Errorf("%T.num1 (1) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:num1: %s", p, err)
	}
	return err
}

func (p *AddArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("num2", thrift.I32, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:num2: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Num2)); err != nil {
		return fmt.Errorf("%T.num2 (2) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:num2: %s", p, err)
	}
	return err
}

func (p *AddArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddArgs(%+v)", *p)
}

type AddResult struct {
	Success int32 `thrift:"success,0"`
}

func NewAddResult() *AddResult {
	return &AddResult{}
}

func (p *AddResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *AddResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 0: %s")
	} else {
		p.Success = v
	}
	return nil
}

func (p *AddResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("add_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	switch {
	default:
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *AddResult) writeField0(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("success", thrift.I32, 0); err != nil {
		return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Success)); err != nil {
		return fmt.Errorf("%T.success (0) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 0:success: %s", p, err)
	}
	return err
}

func (p *AddResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddResult(%+v)", *p)
}

type Test1Args struct {
	Value1 bool              `thrift:"value1,1,required"`
	Value2 int16             `thrift:"value2,2,required"`
	Value3 int64             `thrift:"value3,3,required"`
	Value4 float64           `thrift:"value4,4,required"`
	Value5 []byte            `thrift:"value5,5,required"`
	Value6 map[string]string `thrift:"value6,6,required"`
	Value7 []int32           `thrift:"value7,7,required"`
}

func NewTest1Args() *Test1Args {
	return &Test1Args{}
}

func (p *Test1Args) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.readField7(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Test1Args) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.Value1 = v
	}
	return nil
}

func (p *Test1Args) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI16(); err != nil {
		return fmt.Errorf("error reading field 2: %s")
	} else {
		p.Value2 = v
	}
	return nil
}

func (p *Test1Args) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 3: %s")
	} else {
		p.Value3 = v
	}
	return nil
}

func (p *Test1Args) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return fmt.Errorf("error reading field 4: %s")
	} else {
		p.Value4 = v
	}
	return nil
}

func (p *Test1Args) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return fmt.Errorf("error reading field 5: %s")
	} else {
		p.Value5 = v
	}
	return nil
}

func (p *Test1Args) readField6(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return fmt.Errorf("error reading map begin: %s")
	}
	p.Value6 = make(map[string]string, size)
	for i := 0; i < size; i++ {
		var _key14 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s")
		} else {
			_key14 = v
		}
		var _val15 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s")
		} else {
			_val15 = v
		}
		p.Value6[_key14] = _val15
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return fmt.Errorf("error reading map end: %s")
	}
	return nil
}

func (p *Test1Args) readField7(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list being: %s")
	}
	p.Value7 = make([]int32, 0, size)
	for i := 0; i < size; i++ {
		var _elem16 int32
		if v, err := iprot.ReadI32(); err != nil {
			return fmt.Errorf("error reading field 0: %s")
		} else {
			_elem16 = v
		}
		p.Value7 = append(p.Value7, _elem16)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s")
	}
	return nil
}

func (p *Test1Args) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("test1_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *Test1Args) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("value1", thrift.BOOL, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:value1: %s", p, err)
	}
	if err := oprot.WriteBool(bool(p.Value1)); err != nil {
		return fmt.Errorf("%T.value1 (1) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:value1: %s", p, err)
	}
	return err
}

func (p *Test1Args) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("value2", thrift.I16, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:value2: %s", p, err)
	}
	if err := oprot.WriteI16(int16(p.Value2)); err != nil {
		return fmt.Errorf("%T.value2 (2) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:value2: %s", p, err)
	}
	return err
}

func (p *Test1Args) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("value3", thrift.I64, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:value3: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Value3)); err != nil {
		return fmt.Errorf("%T.value3 (3) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:value3: %s", p, err)
	}
	return err
}

func (p *Test1Args) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("value4", thrift.DOUBLE, 4); err != nil {
		return fmt.Errorf("%T write field begin error 4:value4: %s", p, err)
	}
	if err := oprot.WriteDouble(float64(p.Value4)); err != nil {
		return fmt.Errorf("%T.value4 (4) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 4:value4: %s", p, err)
	}
	return err
}

func (p *Test1Args) writeField5(oprot thrift.TProtocol) (err error) {
	if p.Value5 != nil {
		if err := oprot.WriteFieldBegin("value5", thrift.BINARY, 5); err != nil {
			return fmt.Errorf("%T write field begin error 5:value5: %s", p, err)
		}
		if err := oprot.WriteBinary(p.Value5); err != nil {
			return fmt.Errorf("%T.value5 (5) field write error: %s", p)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 5:value5: %s", p, err)
		}
	}
	return err
}

func (p *Test1Args) writeField6(oprot thrift.TProtocol) (err error) {
	if p.Value6 != nil {
		if err := oprot.WriteFieldBegin("value6", thrift.MAP, 6); err != nil {
			return fmt.Errorf("%T write field begin error 6:value6: %s", p, err)
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Value6)); err != nil {
			return fmt.Errorf("error writing map begin: %s")
		}
		for k, v := range p.Value6 {
			if err := oprot.WriteString(string(k)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p)
			}
			if err := oprot.WriteString(string(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return fmt.Errorf("error writing map end: %s")
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 6:value6: %s", p, err)
		}
	}
	return err
}

func (p *Test1Args) writeField7(oprot thrift.TProtocol) (err error) {
	if p.Value7 != nil {
		if err := oprot.WriteFieldBegin("value7", thrift.LIST, 7); err != nil {
			return fmt.Errorf("%T write field begin error 7:value7: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.I32, len(p.Value7)); err != nil {
			return fmt.Errorf("error writing list begin: %s")
		}
		for _, v := range p.Value7 {
			if err := oprot.WriteI32(int32(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return fmt.Errorf("error writing list end: %s")
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 7:value7: %s", p, err)
		}
	}
	return err
}

func (p *Test1Args) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Test1Args(%+v)", *p)
}

type Test1Result struct {
	Success map[string]bool `thrift:"success,0"`
}

func NewTest1Result() *Test1Result {
	return &Test1Result{}
}

func (p *Test1Result) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Test1Result) readField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadSetBegin()
	if err != nil {
		return fmt.Errorf("error reading set being: %s")
	}
	p.Success = make(map[string]bool, size)
	for i := 0; i < size; i++ {
		var _elem17 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s")
		} else {
			_elem17 = v
		}
		p.Success[_elem17] = true
	}
	if err := iprot.ReadSetEnd(); err != nil {
		return fmt.Errorf("error reading set end: %s")
	}
	return nil
}

func (p *Test1Result) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("test1_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	switch {
	default:
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *Test1Result) writeField0(oprot thrift.TProtocol) (err error) {
	if p.Success != nil {
		if err := oprot.WriteFieldBegin("success", thrift.SET, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := oprot.WriteSetBegin(thrift.STRING, len(p.Success)); err != nil {
			return fmt.Errorf("error writing set begin: %s")
		}
		for v, _ := range p.Success {
			if err := oprot.WriteString(string(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p)
			}
		}
		if err := oprot.WriteSetEnd(); err != nil {
			return fmt.Errorf("error writing set end: %s")
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *Test1Result) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Test1Result(%+v)", *p)
}