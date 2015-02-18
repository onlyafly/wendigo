package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/robertkrimen/otto"
)

func readFile(fileName string) (string, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	content := bytes.NewBuffer(data).String()
	return content, nil
}

func main() {
	vm := otto.New()

	apihub, _ := vm.Object(`apihub = {}`)
	apihub.Set("provide", func(call otto.FunctionCall) otto.Value {
		right, _ := call.Argument(0).ToInteger()
		result, _ := vm.ToValue(2 + right)
		return result
	})

	content, readErr := readFile("example.js")
	if readErr != nil {
		fmt.Printf("Problem: %v\n", readErr)
	}

	vm.Run(content)

}
