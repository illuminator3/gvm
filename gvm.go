package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	file := *flag.String("file", "", "file to run")

	flag.Parse()

	if len(file) == 0 {
		fmt.Println("Usage: gvm -file <file>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	bytes, err := os.ReadFile(file)

	if err != nil {
		panic(err)
	}

	reader := ClassReader{
		data: bytes,
	}

	cf := reader.ReadClassFile()

	resources := make(map[string]ClassFile)

	resources[file] = *cf

	runtime := CreateRuntime(resources)

	runtime.Run(file)
}
