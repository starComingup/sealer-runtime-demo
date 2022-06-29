package main

import (
	"fmt"
	"go-learn/apply"
)

func main() {
	processor := apply.NewProcessor()
	err := processor.ApplyClusterFile()
	if err != nil {
		fmt.Println(" handler a error")
	}
}
