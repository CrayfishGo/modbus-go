package modbus

import "modbus-go/funcode"

type ModbusPdu interface {
	GetFunctionCode() funcode.FunctionCode
}
