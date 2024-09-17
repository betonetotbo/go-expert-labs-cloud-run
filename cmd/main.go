package main

import (
	"encoding/json"
	"fmt"
	"github.com/betonetotbo/go-expert-labs-cloud-run/internal/config"
	"github.com/betonetotbo/go-expert-labs-cloud-run/internal/entity"
	"github.com/betonetotbo/go-expert-labs-cloud-run/internal/service"
	"github.com/betonetotbo/go-expert-labs-cloud-run/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	vc := service.NewViacepApi()
	wa := service.NewWeatherAPI(cfg)
	uc := usecase.NewConsultaClima(vc, wa)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		cep := r.URL.Query().Get("cep")
		output, err := uc.Execute(&usecase.ConsultaClimaInputDTO{
			Cep: entity.CEP(cep),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			_ = json.NewEncoder(w).Encode(output)
		}
	})

	log.Printf("HTTP server started on port %d", cfg.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r)
}
