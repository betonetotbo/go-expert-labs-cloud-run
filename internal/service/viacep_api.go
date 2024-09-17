package service

import (
	"encoding/json"
	"fmt"
	"github.com/betonetotbo/go-expert-labs-cloud-run/internal/entity"
	"net/http"
)

type (
	ViacepResponse struct {
		Cep         string `json:"cep"`
		Logradouro  string `json:"logradouro"`
		Complemento string `json:"complemento"`
		Unidade     string `json:"unidade"`
		Bairro      string `json:"bairro"`
		Localidade  string `json:"localidade"`
		Uf          string `json:"uf"`
		Estado      string `json:"estado"`
		Regiao      string `json:"regiao"`
		Ibge        string `json:"ibge"`
		Gia         string `json:"gia"`
		Ddd         string `json:"ddd"`
		Siafi       string `json:"siafi"`
	}

	ViacepApi interface {
		ConsultaCep(cep entity.CEP) (*ViacepResponse, error)
	}

	viacepApi struct{}
)

func NewViacepApi() ViacepApi {
	return &viacepApi{}
}

func (v *viacepApi) ConsultaCep(cep entity.CEP) (*ViacepResponse, error) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep.GetDigits()), nil)
	resp, e := http.DefaultClient.Do(req)
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	var respObj ViacepResponse
	e = json.NewDecoder(resp.Body).Decode(&respObj)
	if e != nil {
		return nil, e
	}
	return &respObj, nil
}
