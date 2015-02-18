package main

import (
	"fmt"

	"github.com/robertkrimen/otto"
)

func main() {
	vm := otto.New()
	vm.Run(`
      abc = 2 + 2;
      console.log("The value of abc is " + abc); // 4
  `)

	stringVal, _ := vm.Get("abc")
	intVal, _ := stringVal.ToInteger()

	fmt.Printf("abc = %v\n", intVal)

	vm.Set("def", 11)
	vm.Run(`
    console.log("The value of def is " + def);
    // The value of def is 11
  `)

	vm.Set("xyzzy", "Nothing happens.")
	vm.Run(`
      console.log(xyzzy.length); // 16
  `)

	xyzzyVmVal, _ := vm.Run("xyzzy.length")
	{
		// value is an int64 with a value of 16
		xyzzyInt, _ := xyzzyVmVal.ToInteger()
		fmt.Printf("xyzzy = %v\n", xyzzyInt)
	}

	undefinedVmVal, err := vm.Run("abcdefghijlmnopqrstuvwxyz.length")
	if err != nil {
		// err = ReferenceError: abcdefghijlmnopqrstuvwxyz is not defined
		// If there is an error, then value.IsUndefined() is true
		fmt.Printf("err: %v, %v\n", err, undefinedVmVal)
	}

	vm.Set("sayHello", func(call otto.FunctionCall) otto.Value {
		fmt.Printf("Hello, %s.\n", call.Argument(0).String())
		return otto.Value{}
	})

	vm.Set("twoPlus", func(call otto.FunctionCall) otto.Value {
		right, _ := call.Argument(0).ToInteger()
		result, _ := vm.ToValue(2 + right)
		return result
	})

	result, _ := vm.Run(`
    sayHello("Xyzzy");      // Hello, Xyzzy.
    sayHello();             // Hello, undefined

    result = twoPlus(2.0); // 4
  `)
	fmt.Printf("result: %v\n", result)

	object, _ := vm.Object(`baz = {}`)
	object.Set("a", 11)

	vm.Run(`
    console.log("The value of baz.a is " + baz.a);
  `)
}
