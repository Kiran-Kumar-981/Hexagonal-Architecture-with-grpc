package api

import (
	"fmt"
	"hexarch/internal/ports"
)

type Adapter struct {
	arthi ports.ArthimeticPort
}

func NewAdapter(arith ports.ArthimeticPort) *Adapter {
	return &Adapter{arthi: arith}
}
func (apia Adapter) GetAddition(a, b int) (int, error) {
	result, err := apia.arthi.Addition(a, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	return result, nil
}
func (apia Adapter) GetSubtraction(a, b int) (int, error) {
	result, err := apia.arthi.Addition(a, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	return result, nil
}
func (apia Adapter) GetMultiplication(a, b int) (int, error) {
	result, err := apia.arthi.Addition(a, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	return result, nil
}
func (apia Adapter) GetDivition(a, b int) (int, error) {
	result, err := apia.arthi.Addition(a, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	return result, nil
}
