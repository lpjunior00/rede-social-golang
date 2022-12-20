package requests

import (
	"fmt"
	"io"
	"net/http"
	"webapp/src/cookies"
)

func RequestWithAutentication(r *http.Request, method, url string, data io.Reader) (*http.Response, error) {
	//Crio uma nova requisicao. A antiga serve só pra trazer o cookie
	request, erro := http.NewRequest(method, url, data)

	if erro != nil {
		return nil, erro
	}

	//Recupero os dados do cookie e seto na nova requisicao
	cookie, _ := cookies.ReadCookie(r)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cookie["token"]))

	//Crio um novo httpClient e chamo o methodo DO para fazer a requisicao
	client := &http.Client{}
	response, erro := client.Do(request)
	if erro != nil {
		return nil, erro
	}

	return response, nil
}
