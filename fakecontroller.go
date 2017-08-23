package main

import "fmt"

// FakeController is a demo purpose controller. Just for simulation purpose.
type FakeController struct {
	name   string
	motors []float32
}

// NewFakeController constructor
func NewFakeController() *FakeController {
	return &FakeController{
		name:   "Fake Controller",
		motors: make([]float32, 2),
	}
}

// Info returns a controller description
func (c *FakeController) Info() *IOControllerInfo {
	return &IOControllerInfo{
		Name:              c.name,
		MotorChannelCount: uint(len(c.motors)),
	}
}

// SetMotorSpeed changes the motor on a given channel speed. The speed
// should be comprised between -1.0 and 1.0. Setting the speed to 0.0 will
// poweroff the motor.
func (c *FakeController) SetMotorSpeed(channel uint, speed float32) error {
	fmt.Printf("Setting motor %d speed to %f\n", channel+1, speed)
	c.motors[channel] = speed
	return nil
}

// GetMotorSpeed returns the current speed of a motor on a given channel.
func (c *FakeController) GetMotorSpeed(channel uint) (float32, error) {
	return c.motors[channel], nil
}
