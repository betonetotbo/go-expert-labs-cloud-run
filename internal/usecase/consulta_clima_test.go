package usecase

import (
	"github.com/betonetotbo/go-expert-labs-cloud-run/internal/entity"
	"github.com/betonetotbo/go-expert-labs-cloud-run/internal/service"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	_ "github.com/stretchr/testify/mock"
	"testing"
)

type (
	mockViacepApi struct {
		mock.Mock
	}

	mockWeatherApi struct {
		mock.Mock
	}
)

func (m *mockWeatherApi) ConsultaClima(consulta string) (*service.WeatherAPIResult, error) {
	returns := m.Called(consulta)
	r := returns.Get(0).(*service.WeatherAPIResult)
	return r, returns.Error(1)
}

func (m *mockViacepApi) ConsultaCep(cep entity.CEP) (*service.ViacepResponse, error) {
	returns := m.Called(cep)
	r := returns.Get(0).(*service.ViacepResponse)
	return r, returns.Error(1)
}

func Test_ConsultaClima(t *testing.T) {
	// Assemble
	viacep := &mockViacepApi{}
	weather := &mockWeatherApi{}
	cc := NewConsultaClima(viacep, weather)

	viacep.On("ConsultaCep", entity.CEP("89120-000")).Once().Return(&service.ViacepResponse{
		Localidade: "localidade",
		Uf:         "uf",
	}, nil)
	weather.On("ConsultaClima", "localidade,uf,Brasil").Once().Return(&service.WeatherAPIResult{
		Current: service.WeatherCurrent{
			TempC: 1.12,
			TempF: 3.45,
		},
	}, nil)

	// Act
	out, err := cc.Execute(&ConsultaClimaInputDTO{
		Cep: entity.CEP("89120-000"),
	})

	// Verify
	assert.NoError(t, err)
	if assert.NotNil(t, out) {
		assert.Equal(t, ConsultaClimaOutputDTO{
			TempC: 1.12,
			TempF: 3.45,
			TempK: 274.12,
		}, *out)
	}
}

func Test_ConsultaClima_InvalidCep(t *testing.T) {
	// Assemble
	viacep := &mockViacepApi{}
	weather := &mockWeatherApi{}
	cc := NewConsultaClima(viacep, weather)

	// Act
	out, err := cc.Execute(&ConsultaClimaInputDTO{
		Cep: entity.CEP("666"),
	})

	// Verify
	assert.Error(t, err)
	assert.ErrorContains(t, err, "cep is not valid: 666")
	assert.Nil(t, out)
}
