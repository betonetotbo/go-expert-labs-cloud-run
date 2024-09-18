package service

import (
	"encoding/json"
	"fmt"
	"github.com/betonetotbo/go-expert-labs-cloud-run/internal/entity"
	"github.com/betonetotbo/go-expert-labs-cloud-run/internal/utils"
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

	viacepApi struct {
		baseUrl string
	}
)

func NewViacepApi() ViacepApi {
	return &viacepApi{baseUrl: "https://viacep.com.br"}
}

func (v *viacepApi) ConsultaCep(cep entity.CEP) (*ViacepResponse, error) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/ws/%s/json/", v.baseUrl, cep.GetDigits()), nil)
	resp, e := http.DefaultClient.Do(req)
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, utils.NewHttpError(resp.StatusCode)
	}

	var respObj ViacepResponse
	e = json.NewDecoder(resp.Body).Decode(&respObj)
	if e != nil {
		return nil, e
	}

	if respObj.Uf == "" {
		return nil, utils.NewHttpError(http.StatusNotFound)
	}

	return &respObj, nil
}
