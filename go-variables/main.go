package main

import (
	"fmt"
)

var (
	nome                   = "Fernando"
	idade                  = 42
	url, path, query, page = "menyoo.com.br", "/search", "&q=golang", 1
)

const PI = 3.141592

func main() {
	var x, y int = 1, 2
	var a string
	var s string

    a = "texto 1"
    b := "texto 2"

	ola := func() {
		fmt.Printf("Olá da função anômnima!\r\n")
	}

	fmt.Printf("a tipo: %T\r\n", a)
	fmt.Printf("b tipo:%T\r\n", b)
	fmt.Printf("PI tipo: %T\r\n", PI)
	fmt.Printf("ola tipo: %T\r\n", ola)

	fmt.Printf("valor de a = %v\r\n", a)
	fmt.Printf("valor de b = %v\r\n", b)
	fmt.Printf("valor de PI = %v\r\n", PI)

    fmt.Printf("valor de x = %v\r\n", x)
    fmt.Printf("valor de y = %v\r\n", y)

    fmt.Printf("valor de s = %\r\n", s)

    ola()

}
