package viacep

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type ViaCEPResponse struct {
	Localidade string `json:"localidade"`
	Erro       string `json:"erro"`
}

var ErrCepNotFound = errors.New("can not find zipcode")

func GetLocationByCEP(cep string) (string, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("viacep api error: status %d", resp.StatusCode)
	}

	var viaCepData ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&viaCepData); err != nil {
		return "", err
	}

	if viaCepData.Erro != "" || viaCepData.Localidade == "" {
		return "", ErrCepNotFound
	}

	return viaCepData.Localidade, nil
}