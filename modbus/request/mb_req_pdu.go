package request

import "modbus-go/modbus"

type ModbusRequestPdu interface {
	modbus.ModbusPdu
}
