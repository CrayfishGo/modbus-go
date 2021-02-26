package response

import "modbus-go/modbus"

type ModbusResponsePdu interface {
	modbus.ModbusPdu
}
