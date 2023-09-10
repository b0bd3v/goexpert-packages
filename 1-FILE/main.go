package main

import "os"

func main() {
	f, err := os.Create("file.txt")

	if err != nil {
		panic(err)
	}

	//size, err := file.WriteString("Hello World")

	size, err := f.Write([]byte("Helo World"))

	if err != nil {
		panic(err)
	}

	println("Size: ", size, "bytes")

	f.Close()

	// Read file
	file, err := os.ReadFile("file.txt")

	if err != nil {
		panic(err)
	}

	println("File content: ", string(file))
}
