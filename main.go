package main

import . "fmt"

//go:generate mockgen -source=main.go -destination=mocks/main.go
type I interface{}

func main() {
	/* fmt. */ Println("Hello world!")
}
