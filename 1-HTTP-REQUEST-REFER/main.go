package main

import "fmt"

func main() {
	//req, err := http.Get("https://www.google.com.br")
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer req.Body.Close()
	//res, err := io.ReadAll(req.Body)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//println(string(res))

	defer fmt.Println("Primeira Linha")
	fmt.Println("Segunda Linha")
	fmt.Println("Terceira Linha")
}
