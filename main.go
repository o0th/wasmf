package main

import (
	"github.com/wasmerio/wasmer-go/wasmer"
	"io/ioutil"
	"fmt"
)

type response struct {
	status int
}

func main() {
	wasmBytes, _ := ioutil.ReadFile("./assemblyscript/exports.wasm")

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	module, err := wasmer.NewModule(store, wasmBytes)
	if err != nil {
		panic(fmt.Sprintln("Failed to compile module;", err))
	}

	importObject := wasmer.NewImportObject()

	instance, err := wasmer.NewInstance(module, importObject)
	if err != nil {
		panic(fmt.Sprintln("Failed to instantiate the module;", err))
	}

	memory, err := instance.Exports.GetMemory("memory")
	if err != nil {
		panic(fmt.Sprintln("failed to get memory;", err))
	}

	data := memory.Data()
	fmt.Println(string(data))
}
