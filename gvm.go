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
	runtime := CreateRuntime()

	runtime.Run(*cf)

	//fmt.Printf("%#v\n", cf)
}
