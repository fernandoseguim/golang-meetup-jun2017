package main

import (
    "encoding/json"
    "fmt"
)

type Aluno struct {
    Nome string
    Idade int
}

type Localidade struct {
    X int `json:"valor_x"`
    Y int `json:"valor_x"`
    Nome string `json:"nome_da_localidade"`
    valor int
}

func (p *Localidade) Soma(l Localidade){
    p.X = l.X
    p.Y = l.Y
}

func main()  {

    aluno := Aluno{
        Nome: "Fernando",
        Idade: 28,
    }

    fmt.Println("Aluno:", aluno)

    minhaCasa := Localidade{3, 4, "casa", 500}

    estruturaVazia := Aluno{}

    escolaLocalidade := Localidade{
        X: 100,
        Y: 200,
        Nome: "escola",
    }

    var outraLocalidade Localidade
    outraLocalidade.X = 10
    outraLocalidade.Y = 20
    outraLocalidade.Nome = "Trabalho"

    fmt.Println("Minha Casa: ", minhaCasa)
    fmt.Println("Outra Localidade", outraLocalidade)
    fmt.Println("Escola", escolaLocalidade)
    fmt.Println("Aluno", aluno)
    fmt.Printf("Estrutura vazia %q\r\n", estruturaVazia)
    fmt.Printf("Estrutura vazia %+v\r\n", estruturaVazia)

    minhaCasa.Soma(outraLocalidade)

    fmt.Println("Localidade Minha Casa somada com outra Localidade", minhaCasa)

    fmt.Printf("Minha Casa %v\r\n", minhaCasa)
    fmt.Printf("Minha Casa %+v\r\n", minhaCasa)

    //Brincando com json
    j, err := json.Marshal(minhaCasa)
    if err != nil{
        panic(err)
    }

    fmt.Println("Minha casa json", string(j))

    novaLocalidade := Localidade{}
    err = json.Unmarshal(j, &novaLocalidade)
    if err != nil{
        panic(err)
    }

    fmt.Println("Pondo depois do Unmarshal: ", novaLocalidade)
}
