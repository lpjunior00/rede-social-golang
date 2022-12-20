package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// jwt token é uma string que contem as informacoes de uma autorizacao (id, tempo que dura, etc). que serve pra autenticar uma requisicao
func CreateToken(userId uint64) (string, error) {

	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix() //O .Unix devolve a quantidade de milisegundos desde 1 de janeiro de 1970 (era unix)
	permissions["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions) //Aqui vai gerar um novo token usando esse mtodo informado de assinatura
	return token.SignedString([]byte(config.SecretKey))

}

// Verifica se o token passado na requisição é valido
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, erro := jwt.Parse(tokenString, returnSecretKey)
	if erro != nil {
		return erro
	}

	//Pego as permissoes do token e valido se ele é valido
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")

}

func extractToken(r *http.Request) string {

	//Vai recuperar o bearer token (portador). Exemplo: Bearer UYQWIYUQWYUIQWYUQWUIQYWUI . Vai precisar dar um split pra pegar o token
	token := r.Header.Get("Authorization")

	//verifico se veio no padrao (BEARER XXX). duas palavaras separadas por um espaço
	stringSplit := strings.Split(token, " ")
	if len(stringSplit) == 2 {
		return stringSplit[1]
	}

	return ""
}

// Funcao que vai retornar a secret key, numa interface generica e um error
func returnSecretKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Signture method not exists %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

func ExtractUserId(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, erro := jwt.Parse(tokenString, returnSecretKey)
	if erro != nil {
		return 0, erro
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return userId, nil
	}

	return 0, errors.New("invalid token")
}
