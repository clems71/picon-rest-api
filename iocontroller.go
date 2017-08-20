package main

type IOControllerInfo struct {
	Name               string
	MotorChannelCount  uint
	OutputChannelCount uint
}

type IOController interface {
	Info() *IOControllerInfo

	SetMotorSpeed(uint, float32) error
	GetMotorSpeed(uint) (float32, error)

	SetOutputValue(uint, float32) error
	GetOutputValue(uint) (float32, error)
}
