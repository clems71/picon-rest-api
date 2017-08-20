package main

type IOControllerInfo struct {
	Name              string
	MotorChannelCount uint
}

type IOController interface {
	Info() *IOControllerInfo
	SetMotorSpeed(uint, float32) error
	GetMotorSpeed(uint) (float32, error)
}
