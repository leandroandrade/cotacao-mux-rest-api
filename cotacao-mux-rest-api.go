package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"regexp"
	"errors"
	"io"
	"bytes"
)

const cotacaoEndpoint = "https://ptax.bcb.gov.br/ptax_internet/consultarUltimaCotacaoDolar.do"

type cotacao struct {
	Compra string `json:"compra,omitempty"`
	Venda  string `json:"venda,omitempty"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	sub := router.PathPrefix("/api/v1").Subrouter()
	sub.HandleFunc("/info", hello).Methods("GET")
	sub.HandleFunc("/cotacao", getCotacao).Methods("GET")

	fmt.Println("Initializing server...")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func hello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("API para retornar a cotação do Dolar."))
}

func getCotacao(writer http.ResponseWriter, request *http.Request) {
	content := retornaCotacaoEndpoint()
	cotacao := formatResponse(content)

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(cotacao)
}

func retornaCotacaoEndpoint() (content []byte) {
	response, err := http.Get(cotacaoEndpoint)
	checkError(err)

	defer response.Body.Close()

	var buffer bytes.Buffer
	io.Copy(&buffer, response.Body)

	return buffer.Bytes()
}

func formatResponse(content []byte) (value cotacao) {
	regex := regexp.MustCompile("[1-9],[0-9][0-9][0-9][0-9]")
	valores := regex.FindAll(content, -1)

	if len(valores) == 2 {
		value = cotacao{string(valores[0][:]), string(valores[1][:])}
	} else {
		panic(errors.New("não foi possivel encontrar valores"))
	}
	return
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
