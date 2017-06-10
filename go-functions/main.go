package main

import (
    "fmt"
)

func soma(x int, y int) int{
    return y + x
}

func troca(x string, y string) (string, string)  {
    return y, x
}

func divide(x, y int) (resultado, resto int)  {
    resto = x % y
    resultado = x / y
    return
}

func executaFuncao(f func(string) string, valor string){
    aux := f(valor)
    fmt.Println(aux)
}

func printValorByRef(valor *string){
    fmt.Printf("Valor por referencia = %v\r\n", *valor)
}

func main(){
    fmt.Printf("+++++++++++++++ Funções! +++++++++++++++++\r\n")

    fmt.Printf("Soma 1+1 = %v\r\n", soma(1,1))

    b, a := troca("A", "B")
    fmt.Printf("Troca A, B = %v, %v\r\n", b, a)

    resu, rest := divide(5,2)
    fmt.Printf("A divisão de 5 por 2 é = %v\r\n", resu)
    fmt.Printf("O resto da divisão de 5 por 2 é = %v\r\n", rest)

    ola := func (v string) string  {
        return "Olá " + v + "!\r\n"
    }

    executaFuncao(ola, "Cesar")

    valor := "Esse valor não vai ser copiado, só estamos passando o ponteiro"
    printValorByRef(&valor)

}
