package usecase

import (
	"fmt"
	"github.com/betonetotbo/go-expert-labs-cloud-run/internal/entity"
	"github.com/betonetotbo/go-expert-labs-cloud-run/internal/service"
	"github.com/betonetotbo/go-expert-labs-cloud-run/internal/utils"
	"log"
)

type (
	ConsultaClimaInputDTO struct {
		Cep entity.CEP `json:"cep"`
	}

	ConsultaClimaOutputDTO struct {
		TempC float64 `json:"temp_c"`
		TempF float64 `json:"temp_f"`
		TempK float64 `json:"temp_k"`
	}

	ConsultaClima struct {
		viacep  service.ViacepApi
		weather service.WeatherAPI
	}
)

func NewConsultaClima(viacep service.ViacepApi, weather service.WeatherAPI) *ConsultaClima {
	return &ConsultaClima{
		viacep:  viacep,
		weather: weather,
	}
}

func (c *ConsultaClima) Execute(input *ConsultaClimaInputDTO) (*ConsultaClimaOutputDTO, error) {
	if !input.Cep.IsValid() {
		return nil, fmt.Errorf("cep is not valid: %s", input.Cep)
	}

	log.Printf("Consultando CEP: %s", input.Cep)
	cc, err := c.viacep.ConsultaCep(input.Cep)
	if err != nil {
		return nil, err
	}
	log.Printf("CEP encontrado: %+v", cc)

	q := utils.ConcatFields(cc.Logradouro, cc.Bairro, cc.Localidade, cc.Uf, "Brasil")
	log.Printf("Consulta clima para: %s", q)
	wr, err := c.weather.ConsultaClima(q)
	if err != nil {
		return nil, err
	}

	return &ConsultaClimaOutputDTO{
		TempC: wr.Current.TempC,
		TempF: wr.Current.TempF,
		TempK: wr.Current.TempC + 273.0,
	}, nil
}
