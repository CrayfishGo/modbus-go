package funcode

type FunctionCode int64

const (
	ReadCoils                      FunctionCode = 0x01
	ReadDiscreteInputs             FunctionCode = 0x02
	ReadHoldingRegisters           FunctionCode = 0x03
	ReadInputRegisters             FunctionCode = 0x04
	WriteSingleCoil                FunctionCode = 0x05
	WriteSingleRegister            FunctionCode = 0x06
	ReadExceptionStatus            FunctionCode = 0x07
	Diagnostics                    FunctionCode = 0x08
	GetCommEventCounter            FunctionCode = 0x0B
	GetCommEventLog                FunctionCode = 0x0C
	WriteMultipleCoils             FunctionCode = 0x0F
	WriteMultipleRegisters         FunctionCode = 0x10
	ReportSlaveId                  FunctionCode = 0x11
	ReadFileRecord                 FunctionCode = 0x14
	WriteFileRecord                FunctionCode = 0x15
	MaskWriteRegister              FunctionCode = 0x16
	ReadWriteMultipleRegisters     FunctionCode = 0x17
	ReadFifoQueue                  FunctionCode = 0x18
	EncapsulatedInterfaceTransport FunctionCode = 0x2B
)

type ExceptionCode int64

const (
	IllegalFunction                     ExceptionCode = 0x01
	IllegalDataAddress                  ExceptionCode = 0x02
	IllegalDataValue                    ExceptionCode = 0x03
	SlaveDeviceFailure                  ExceptionCode = 0x04
	Acknowledge                         ExceptionCode = 0x05
	SlaveDeviceBusy                     ExceptionCode = 0x06
	MemoryParityError                   ExceptionCode = 0x08
	GatewayPathUnavailable              ExceptionCode = 0x0A
	GatewayTargetDeviceFailedToResponse ExceptionCode = 0x0B
)
