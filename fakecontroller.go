package main

import "fmt"

// FakeController is a demo purpose controller. Just for simulation purpose.
type FakeController struct {
	name    string
	motors  []float32
	outputs []float32
}

// NewFakeController constructor
func NewFakeController(name string, motorChannels uint, outputChannels uint) *FakeController {
	return &FakeController{
		name:    name,
		motors:  make([]float32, motorChannels),
		outputs: make([]float32, outputChannels),
	}
}

// Info returns a controller description
func (c *FakeController) Info() *IOControllerInfo {
	return &IOControllerInfo{
		Name:               c.name,
		MotorChannelCount:  uint(len(c.motors)),
		OutputChannelCount: uint(len(c.outputs)),
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

// SetOutputValue sets output data for a given channel.
func (c *FakeController) SetOutputValue(channel uint, value float32) error {
	fmt.Printf("Setting output %d value to %f\n", channel+1, value)
	c.outputs[channel] = value
	return nil
}

// GetOutputValue returns the current value of an output on a given channel.
func (c *FakeController) GetOutputValue(channel uint) (float32, error) {
	return c.outputs[channel], nil
}
