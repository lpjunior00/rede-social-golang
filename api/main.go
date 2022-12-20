package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// Criar uma funcao para preencher um slice de chave para montar a chave do secret
//Lembrando que init roda sempre antes do main
/*
func init() {
	chave := make([]byte, 64)

	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(stringBase64)
}
*/

func main() {

	config.LoadConfigs()

	r := router.Gerar()

	fmt.Printf("Listening on port: %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))

}
