package main

import (
	"os"
)

func main() {
	bytes, err := os.ReadFile("Test.class")

	if err != nil {
		panic(err)
	}

	reader := ClassReader{
		data: bytes,
	}

	cf := reader.ReadClassFile()

	resources := make(map[string]ClassFile)

	resources["Test"] = *cf

	runtime := CreateRuntime(resources)

	runtime.Run("Test")

	//fmt.Printf("%#v\n", cf)
}
