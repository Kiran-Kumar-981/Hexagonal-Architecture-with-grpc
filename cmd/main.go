package main

import (
	"fmt"
	"hexarch/internal/adapter/core/arthimetic"
)

func main() {
	arthAdapter := arthimetic.NewAdapter()
	result, err := arthAdapter.Addition(1, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
