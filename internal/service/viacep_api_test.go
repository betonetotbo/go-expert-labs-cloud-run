package service

import (
	"github.com/betonetotbo/go-expert-labs-cloud-run/internal/entity"
	"github.com/betonetotbo/go-expert-labs-cloud-run/internal/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_NewViacepApi(t *testing.T) {
	assert.NotNil(t, NewViacepApi())
}

func Test_ConsultaCep_Success(t *testing.T) {
	svc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/ws/01001000/json/", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
  "cep": "01001-000",
  "logradouro": "Praça da Sé",
  "complemento": "lado ímpar",
  "unidade": "",
  "bairro": "Sé",
  "localidade": "São Paulo",
  "uf": "SP",
  "estado": "São Paulo",
  "regiao": "Sudeste",
  "ibge": "3550308",
  "gia": "1004",
  "ddd": "11",
  "siafi": "7107"
}`))
	}))
	defer svc.Close()

	vc := &viacepApi{
		baseUrl: svc.URL,
	}

	// Act
	resp, err := vc.ConsultaCep(entity.CEP("01001-000"))

	// Verify
	assert.NoError(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, ViacepResponse{
			Cep:         "01001-000",
			Logradouro:  "Praça da Sé",
			Complemento: "lado ímpar",
			Unidade:     "",
			Bairro:      "Sé",
			Localidade:  "São Paulo",
			Uf:          "SP",
			Estado:      "São Paulo",
			Regiao:      "Sudeste",
			Ibge:        "3550308",
			Gia:         "1004",
			Ddd:         "11",
			Siafi:       "7107",
		}, *resp)
	}
}

func Test_ConsultaCep_Error(t *testing.T) {
	svc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer svc.Close()

	vc := &viacepApi{
		baseUrl: svc.URL,
	}

	// Act
	resp, err := vc.ConsultaCep(entity.CEP("01001-000"))

	// Verify
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, utils.GetHttpErrorCode(err))
	assert.Nil(t, resp)
}
