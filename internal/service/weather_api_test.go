package service

import (
	"github.com/betonetotbo/go-expert-labs-cloud-run/internal/config"
	"github.com/betonetotbo/go-expert-labs-cloud-run/internal/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_NewWeatherApi(t *testing.T) {
	assert.NotNil(t, NewWeatherAPI(&config.Config{}))
}

func Test_ConsultaClima_Success(t *testing.T) {
	svc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/current.json", r.URL.Path)
		assert.Equal(t, "key=key123&q=localidade&lang=pt", r.URL.RawQuery)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
  "location": {
    "name": "Taubate",
    "region": "Sao Paulo",
    "country": "Brazil",
    "lat": -23.03,
    "lon": -45.55,
    "tz_id": "America/Sao_Paulo",
    "localtime_epoch": 1726619330,
    "localtime": "2024-09-17 21:28"
  },
  "current": {
    "last_updated_epoch": 1726618500,
    "last_updated": "2024-09-17 21:15",
    "temp_c": 17.3,
    "temp_f": 63.1,
    "is_day": 0,
    "condition": {
      "text": "Partly cloudy",
      "icon": "//cdn.weatherapi.com/weather/64x64/night/116.png",
      "code": 1003
    },
    "wind_mph": 3.4,
    "wind_kph": 5.4,
    "wind_degree": 104,
    "wind_dir": "ESE",
    "pressure_mb": 1022.0,
    "pressure_in": 30.18,
    "precip_mm": 0.01,
    "precip_in": 0.0,
    "humidity": 72,
    "cloud": 75,
    "feelslike_c": 17.3,
    "feelslike_f": 63.1,
    "windchill_c": 16.2,
    "windchill_f": 61.2,
    "heatindex_c": 16.2,
    "heatindex_f": 61.2,
    "dewpoint_c": 12.3,
    "dewpoint_f": 54.2,
    "vis_km": 10.0,
    "vis_miles": 6.0,
    "uv": 1.0,
    "gust_mph": 4.2,
    "gust_kph": 6.7
  }
}`))
	}))
	defer svc.Close()

	w := &weatherAPI{
		cfg: &config.Config{
			WeatherApiKey: "key123",
		},
		baseUrl: svc.URL,
	}

	// Act
	resp, err := w.ConsultaClima("localidade")

	// Verify
	assert.NoError(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, WeatherAPIResult{
			Location: WeatherLocation{
				Name:           "Taubate",
				Region:         "Sao Paulo",
				Country:        "Brazil",
				Lat:            -23.03,
				Lon:            -45.55,
				TzId:           "America/Sao_Paulo",
				LocaltimeEpoch: 1726619330,
				Localtime:      "2024-09-17 21:28",
			},
			Current: WeatherCurrent{
				LastUpdatedEpoch: 1726618500,
				LastUpdated:      "2024-09-17 21:15",
				TempC:            17.3,
				TempF:            63.1,
				IsDay:            0,
				Condition: WeatherCurrentCondition{
					Text: "Partly cloudy",
					Icon: "//cdn.weatherapi.com/weather/64x64/night/116.png",
					Code: 1003,
				},
				WindMph:    3.4,
				WindKph:    5.4,
				WindDegree: 104,
				WindDir:    "ESE",
				PressureMb: 1022.0,
				PressureIn: 30.18,
				PrecipMm:   0.01,
				PrecipIn:   0.0,
				Humidity:   72,
				Cloud:      75,
				FeelslikeC: 17.3,
				FeelslikeF: 63.1,
				WindchillC: 16.2,
				WindchillF: 61.2,
				HeatindexC: 16.2,
				HeatindexF: 61.2,
				DewpointC:  12.3,
				DewpointF:  54.2,
				VisKm:      10.0,
				VisMiles:   6.0,
				Uv:         1.0,
				GustMph:    4.2,
				GustKph:    6.7,
			},
		}, *resp)
	}
}

func Test_ConsultaClima_Error(t *testing.T) {
	svc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer svc.Close()

	w := &weatherAPI{
		cfg: &config.Config{
			WeatherApiKey: "key123",
		},
		baseUrl: svc.URL,
	}

	// Act
	resp, err := w.ConsultaClima("localidade")

	// Verify
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, utils.GetHttpErrorCode(err))
	assert.Nil(t, resp)
}
