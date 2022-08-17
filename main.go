package main

import (
	"github.com/wasmerio/wasmer-go/wasmer"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"strconv"
)

func main() {
	wasmBytes, _ := ioutil.ReadFile("exports.wasm")
	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)
	module, _ := wasmer.NewModule(store, wasmBytes)

	importObject := wasmer.NewImportObject()
	instance, _ := wasmer.NewInstance(module, importObject)

	sum, _ := instance.Exports.GetFunction("sum")

	app := fiber.New()

	app.Get("/:n1/:n2", func(c *fiber.Ctx) error {
		n1 := c.Params("n1")
		n2 := c.Params("n2")

		n1int64, _ := strconv.ParseInt(n1, 10, 32)
		n2int64, _ := strconv.ParseInt(n2, 10, 32)

		n1int32 := int32(n1int64)
		n2int32 := int32(n2int64)

		result, _ := sum(n1int32, n2int32)
		return c.SendString(strconv.Itoa(int(result.(int32))))
	})

	app.Listen(":3000")
}
