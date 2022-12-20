package cookies

import (
	"net/http"
	"time"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Utiliza as variaveis de ambiente para criar o secure cookie
func Configure() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func SaveCookie(w http.ResponseWriter, userId, token string) error {
	data := map[string]string{
		"userId": userId,
		"token":  token,
	}

	//Codifica o conteudo do json
	codificatedData, erro := s.Encode("data", data)
	if erro != nil {
		return erro
	}

	//Vai setar o cookie dentro do brwoser, passando o nome, o conteudo codificado
	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    codificatedData,
		Path:     "/",  //diz que vai estar acessivel pra toda a apicação
		HttpOnly: true, //Só via client.
	})

	return nil
}

// Retorna os valores armazenados no cookie
func ReadCookie(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("data")

	if erro != nil {
		return nil, erro
	}

	values := make(map[string]string)
	if erro = s.Decode("data", cookie.Value, &values); erro != nil {
		return nil, erro
	}

	return values, nil
}

func DeleteCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0), //Definir que o cookie ta expirado
	})
}
