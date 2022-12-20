package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

// Uma funcao que será executada uma vez só, antes do main, pra gerar as secret keys
// func init() {
// 	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(hashKey)

// 	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(blockKey)
// }

func main() {

	//Load html templates
	utils.LoadTemplates()
	config.Load()
	cookies.Configure()

	r := router.Generate()

	fmt.Printf("Starting webapp. port %d!", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))

}
