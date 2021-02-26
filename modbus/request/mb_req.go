package request

import "modbus-go/funcode"

type SimpleModbusRequest struct {
	functionCode funcode.FunctionCode
}

func (r *SimpleModbusRequest) GetFunctionCode() funcode.FunctionCode {
	return r.functionCode
}

func NewSimpleModbusRequest(code funcode.FunctionCode) *SimpleModbusRequest {
	return &SimpleModbusRequest{
		functionCode: code,
	}
}

type ReadCoilsRequest struct {
	SimpleModbusRequest
	address  int64
	quantity int64
}

func NewReadCoilsRequest(address int64, quantity int64) *ReadCoilsRequest {
	rc := &ReadCoilsRequest{
		address:  address,
		quantity: quantity,
	}
	rc.functionCode = funcode.ReadCoils
	return rc
}

type ReadDiscreteInputsRequest struct {
	SimpleModbusRequest
	address  int64
	quantity int64
}

func NewReadDiscreteInputsRequest(address int64, quantity int64) *ReadDiscreteInputsRequest {
	rc := &ReadDiscreteInputsRequest{
		address:  address,
		quantity: quantity,
	}
	rc.functionCode = funcode.ReadDiscreteInputs
	return rc
}

type ReadHoldingRegistersRequest struct {
	SimpleModbusRequest
	address  int64
	quantity int64
}

func NewReadHoldingRegistersRequest(address int64, quantity int64) *ReadHoldingRegistersRequest {
	rc := &ReadHoldingRegistersRequest{
		address:  address,
		quantity: quantity,
	}
	rc.functionCode = funcode.ReadHoldingRegisters
	return rc
}

type ReadInputRegistersRequest struct {
	SimpleModbusRequest
	address  int64
	quantity int64
}

func NewReadInputRegistersRequest(address int64, quantity int64) *ReadInputRegistersRequest {
	rc := &ReadInputRegistersRequest{
		address:  address,
		quantity: quantity,
	}
	rc.functionCode = funcode.ReadInputRegisters
	return rc
}

type WriteSingleCoilRequest struct {
	SimpleModbusRequest
	address int64
	value   int64
}

/**
 *  address 0x0000 to 0xFFFF (0 to 65535)
 *  value   true or false (0xFF00 or 0x0000)
 */
func NewWriteSingleCoilRequest(address int64, value bool) *WriteSingleCoilRequest {
	rc := &WriteSingleCoilRequest{
		address: address,
	}
	rc.value = 0x0000
	if value {
		rc.value = 0xFF00
	}
	rc.functionCode = funcode.WriteSingleCoil
	return rc
}

type WriteSingleRegisterRequest struct {
	SimpleModbusRequest
	address int64
	value   int64
}

// address：  	0x0000 to 0xFFFF (0 to 65535)
// value: 	 	0x0000 to 0xFFFF (0 to 65535)
func NewWriteSingleRegisterRequest(address int64, value int64) *WriteSingleRegisterRequest {
	rc := &WriteSingleRegisterRequest{
		address: address,
		value:   value,
	}
	rc.functionCode = funcode.WriteSingleRegister
	return rc
}
