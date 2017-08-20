package main

import (
	i2c "github.com/davecheney/i2c"
)

// PiconController is the RPIzero hat controller
type PiconController struct {
	bus     *i2c.I2C
	motors  []float32
	outputs []float32
}

// NewPiconController is the PiconController constructor
func NewPiconController() (*PiconController, error) {
	bus, err := i2c.New(0x22, 1)
	if err != nil {
		return nil, err
	}
	return &PiconController{
		bus:     bus,
		motors:  make([]float32, 2),
		outputs: make([]float32, 6),
	}, nil
}

// Reset the PiconController to initial state
func (c *PiconController) Reset() error {
	_, err := c.bus.Write([]byte{0x20, 0x00})
	return err
}

// Info returns a controller description
func (c *PiconController) Info() *IOControllerInfo {
	return &IOControllerInfo{
		Name:              "Picon Controller",
		MotorChannelCount: uint(2),
	}
}

// SetMotorSpeed changes the motor on a given channel speed. The speed
// should be comprised between -1.0 and 1.0. Setting the speed to 0.0 will
// poweroff the motor.
func (c *PiconController) SetMotorSpeed(channel uint, speed float32) error {
	motorAddress := byte(channel)
	motorSpeed := byte(speed * 127.0)
	// fmt.Printf("MOTOR_SPEED = 0x%02x\n", motorSpeed)
	if _, err := c.bus.Write([]byte{motorAddress, motorSpeed}); err != nil {
		return err
	}
	c.motors[channel] = speed
	return nil
}

// GetMotorSpeed returns the current speed of a motor on a given channel.
func (c *PiconController) GetMotorSpeed(channel uint) (float32, error) {
	return c.motors[channel], nil
}

// SetOutputValue sets output data for a given channel (6 avalaible).
// Channel  Name    Type    Values
// 0     	On/Off  Byte    0 is OFF, 1 is ON
// 1     	PWM     Byte    0 to 100 percentage of ON time
// 2     	Servo   Byte    -100 to + 100 Position in degrees
// 3     	WS2812B 4 Bytes 0:Pixel ID, 1:Red, 2:Green, 3:Blue
func (c *PiconController) SetOutputValue(channel uint, value float32) error {
	outputAddress := byte(channel)
	byteValue := byte(value)
	if _, err := c.bus.Write([]byte{outputAddress, byteValue}); err != nil {
		return err
	}
	c.outputs[channel] = value
	return nil
}

// GetOutputValue returns the current value of an output on a given channel.
func (c *PiconController) GetOutputValue(channel uint) (float32, error) {
	return c.outputs[channel], nil
}
