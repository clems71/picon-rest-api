package main

import "fmt"

type MotorDescInput struct {
	Speed float32 `form:"speed" json:"speed"`
}

type MotorDesc struct {
	ID uint `json:"id"`
	MotorDescInput
}

func motorDescGetOne(c IOController, id uint) (MotorDesc, error) {
	if id >= c.Info().MotorChannelCount {
		return MotorDesc{}, fmt.Errorf("motor index out of bounds")
	}

	speed, err := c.GetMotorSpeed(id)
	return MotorDesc{
		ID: id,
		MotorDescInput: MotorDescInput{
			Speed: speed,
		},
	}, err
}

func motorDescGetAll(c IOController) ([]MotorDesc, error) {
	motorCount := c.Info().MotorChannelCount
	result := make([]MotorDesc, 0, motorCount)

	for i := uint(0); i < motorCount; i++ {
		desc, err := motorDescGetOne(c, i)
		if err != nil {
			return []MotorDesc{}, err
		}
		result = append(result, desc)
	}

	return result, nil
}

func motorDescSetOne(c IOController, id uint, speed float32) (MotorDesc, error) {
	if id >= c.Info().MotorChannelCount {
		return MotorDesc{}, fmt.Errorf("motor index out of bounds")
	}

	err := c.SetMotorSpeed(id, clamp(speed))
	if err != nil {
		return MotorDesc{}, err
	}

	return motorDescGetOne(c, id)
}
