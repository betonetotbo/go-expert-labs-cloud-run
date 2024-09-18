package service

import (
	"encoding/json"
	"fmt"
	"github.com/betonetotbo/go-expert-labs-cloud-run/internal/config"
	"github.com/betonetotbo/go-expert-labs-cloud-run/internal/utils"
	"net/http"
	"net/url"
)

type (
	WeatherAPI interface {
		ConsultaClima(consulta string) (*WeatherAPIResult, error)
	}

	weatherAPI struct {
		cfg     *config.Config
		baseUrl string
	}

	WeatherLocation struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzId           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	}

	WeatherCurrentCondition struct {
		Text string `json:"text"`
		Icon string `json:"icon"`
		Code int    `json:"code"`
	}

	WeatherCurrent struct {
		LastUpdatedEpoch int                     `json:"last_updated_epoch"`
		LastUpdated      string                  `json:"last_updated"`
		TempC            float64                 `json:"temp_c"`
		TempF            float64                 `json:"temp_f"`
		IsDay            int                     `json:"is_day"`
		Condition        WeatherCurrentCondition `json:"condition"`
		WindMph          float64                 `json:"wind_mph"`
		WindKph          float64                 `json:"wind_kph"`
		WindDegree       int                     `json:"wind_degree"`
		WindDir          string                  `json:"wind_dir"`
		PressureMb       float64                 `json:"pressure_mb"`
		PressureIn       float64                 `json:"pressure_in"`
		PrecipMm         float64                 `json:"precip_mm"`
		PrecipIn         float64                 `json:"precip_in"`
		Humidity         int                     `json:"humidity"`
		Cloud            int                     `json:"cloud"`
		FeelslikeC       float64                 `json:"feelslike_c"`
		FeelslikeF       float64                 `json:"feelslike_f"`
		WindchillC       float64                 `json:"windchill_c"`
		WindchillF       float64                 `json:"windchill_f"`
		HeatindexC       float64                 `json:"heatindex_c"`
		HeatindexF       float64                 `json:"heatindex_f"`
		DewpointC        float64                 `json:"dewpoint_c"`
		DewpointF        float64                 `json:"dewpoint_f"`
		VisKm            float64                 `json:"vis_km"`
		VisMiles         float64                 `json:"vis_miles"`
		Uv               float64                 `json:"uv"`
		GustMph          float64                 `json:"gust_mph"`
		GustKph          float64                 `json:"gust_kph"`
	}

	WeatherAPIResult struct {
		Location WeatherLocation `json:"location"`
		Current  WeatherCurrent  `json:"current"`
	}
)

func NewWeatherAPI(cfg *config.Config) WeatherAPI {
	return &weatherAPI{
		cfg:     cfg,
		baseUrl: "https://api.weatherapi.com",
	}
}

func (w *weatherAPI) ConsultaClima(consulta string) (*WeatherAPIResult, error) {
	consulta = url.QueryEscape(consulta)
	u := fmt.Sprintf("%s/v1/current.json?key=%s&q=%s&lang=pt", w.baseUrl, w.cfg.WeatherApiKey, consulta)
	req, _ := http.NewRequest(http.MethodGet, u, nil)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, utils.NewHttpError(resp.StatusCode)
	}

	var result WeatherAPIResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
