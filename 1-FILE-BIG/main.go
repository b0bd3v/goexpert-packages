package main

import (
	"bufio"
	"os"
)

func main() {
	f, err := os.Create("file-big.txt")

	if err != nil {
		panic(err)
	}

	_, err2 := f.WriteString("Helo World, this is a big file, with a lot of text, to test the performance of the file creation and reading.")

	if err2 != nil {
		panic(err2)
	}

	file2, err := os.Open("file-big.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file2)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			break
		}

		println(string(buffer[:n])
	}
}
