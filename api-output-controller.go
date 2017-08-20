package main

import "fmt"

type OutputDesc struct {
	ID    uint    `json:"id"`
	Value float32 `json:"value"`
}

type OutputDescInput struct {
	Value float32 `form:"value" json:"value"`
}

func outputDescGetOne(c IOController, id uint) (OutputDesc, error) {
	if id >= c.Info().OutputChannelCount {
		return OutputDesc{}, fmt.Errorf("output index out of bounds")
	}

	value, err := c.GetOutputValue(id)
	return OutputDesc{
		ID:    id,
		Value: value,
	}, err
}

func outputDescGetAll(c IOController) ([]OutputDesc, error) {
	outputCount := c.Info().OutputChannelCount
	result := make([]OutputDesc, 0, outputCount)

	for i := uint(0); i < outputCount; i++ {
		desc, err := outputDescGetOne(c, i)
		if err != nil {
			return []OutputDesc{}, err
		}
		result = append(result, desc)
	}

	return result, nil
}

func outputDescSetOne(c IOController, id uint, value float32) (OutputDesc, error) {
	if id >= c.Info().OutputChannelCount {
		return OutputDesc{}, fmt.Errorf("output index out of bounds")
	}

	err := c.SetOutputValue(id, clamp(value))
	if err != nil {
		return OutputDesc{}, err
	}

	return outputDescGetOne(c, id)
}
