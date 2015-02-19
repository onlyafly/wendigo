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

	fmt.Printf("Start\n")

	apihub, _ := vm.Object(`apihub = {}`)
	apihub.Set("provide", func(call otto.FunctionCall) otto.Value {

		fmt.Printf("provide Start\n")

		arg := call.Argument(0)
		switch arg.Class() {
		case "Array":
			elemsAny, _ := arg.Export()
			switch elems := elemsAny.(type) {
			case []interface{}:
				for elem := range elems {
					fmt.Printf("elem: %v\n", elem)
				}
			default:
				panic("Arg to provide not an array")
			}
		default:
			panic("Arg to provide not an array")
		}

		return otto.NullValue()
	})

	content, readErr := readFile("example.js")
	if readErr != nil {
		fmt.Printf("Problem: %v\n", readErr)
	}

	_, err := vm.Run(content)
	if err != nil {
		fmt.Printf("Error during script execution: %v\n", err)
	}

	vm.Run(`
			console.log("The value of baz.a is " + baz.a);
		`)

	fmt.Printf("End\n")

}
