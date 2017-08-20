package main

import (
	i2c "github.com/davecheney/i2c"
)

// PiconController is the RPIzero hat controller
type PiconController struct {
	bus    *i2c.I2C
	motors []float32
}

// NewPiconController is the PiconController constructor
func NewPiconController() (*PiconController, error) {
	bus, err := i2c.New(0x22, 1)
	if err != nil {
		return nil, err
	}
	return &PiconController{
		bus:    bus,
		motors: make([]float32, 2),
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
