package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type usuario struct {
	Id          int    `json:"id"`
	Nome        string `json:"nome"`
	Sobrenome   string `json:"sobrenome"`
	Email       string `json:"email"`
	Idade       int    `json:"idade"`
	Altura      int    `json:"altura"`
	Ativo       bool   `json:"ativo"`
	DataCriação string `json:"dataCriacao"`
}

var todos = []usuario{}

func main() {

	todos = append(todos, usuario{
		Id:          1,
		Nome:        "cesar",
		Sobrenome:   "silva",
		Email:       "email",
		Idade:       15,
		Altura:      182,
		Ativo:       true,
		DataCriação: "2024-04-08"})

	todos = append(todos, usuario{
		Id:          2,
		Nome:        "augusto",
		Sobrenome:   "silva",
		Email:       "email",
		Idade:       15,
		Altura:      182,
		Ativo:       true,
		DataCriação: "2024-04-08"})

	fmt.Println("lista todos: ", todos)

	router := gin.Default()

	// exercicio 2 lista manha
	router.GET("/exercicio2", func(c *gin.Context) {
		c.JSON(200, gin.H{"mensagem": "Olá Cesar"})
	})
	// exercicio 3 lista manha
	router.GET("/listaUsuarios/GetAll", func(c *gin.Context) {
		c.JSON(200, todos)
	})

	// exercicio 1 lista tarde
	router.GET("/filtraUsuario/", filtraUsuario)

	//exercicio 2 lista tarde
	router.GET("/usuario/:id", buscaUsuario)

	router.Run()

}

// exercicio 2 lista tarde
func buscaUsuario(ctxt *gin.Context) {

	id := ctxt.Param("id")
	fmt.Printf("O id fornecido no path param é: %s\n", id)
	ide, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	var idUser int
	var idNome string

	for i, user := range todos {
		if ide == user.Id {
			idUser = todos[i].Id
			idNome = todos[i].Nome
			continue
		}
	}

	if idUser != 0 {
		ctxt.String(200, "informação do empregado %v, nome: %s", idUser, idNome)
	} else {
		ctxt.String(404, "informação do empregado não existe!")
	}
}

// exercicio 1 lista tarde
func filtraUsuario(ctxt *gin.Context) {

	//http://localhost:8080/filtraUsuario/?id=0
	queryParamId := ctxt.Query("id")

	//http://localhost:8080/filtraUsuario/?nome=cesar
	queryParamNome := ctxt.Query("nome")
	// queryParamSobrenome := ctxt.Query("sobrenome")
	// queryParamEmail := ctxt.Query("email")
	// queryParamIdade := ctxt.Query("idade")
	// queryParamAltura := ctxt.Query("altura")
	// queryParamAtivo := ctxt.Query("ativo")
	// queryParamData := ctxt.Query("dataCriacao")

	switch {
	case queryParamId != "":
		ide, err := strconv.Atoi(queryParamId)
		if err != nil {
			panic(err)
		}
		ctxt.JSON(200, todos[ide])
	case queryParamNome != "":
		for i, user := range todos {
			if queryParamNome == user.Nome {
				ctxt.JSON(200, todos[i])
				continue
			}
		}
	default:
		//http://localhost:8080/filtraUsuario/?sobrenomenome=cesar
		fmt.Println("Identificador query nao encontrado")
	}

}

/*
	Id          int     `json:"id"`
	Nome        string  `json:"nome"`
	Sobrenome   string  `json:"sobrenome"`
	Email       string  `json:"email"`
	Idade       int     `json:"idade"`
	Altura      float64 `json:"altura"`
	Ativo       bool    `json:"ativo"`
	DataCriação string  `json:"dataCriacao"`
*/
