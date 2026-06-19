package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"regexp"

	"github.com/joaojorgedosanjos/clima-cep/internal/temperature"
	"github.com/joaojorgedosanjos/clima-cep/internal/viacep"
	"github.com/joaojorgedosanjos/clima-cep/internal/weather"
)

type Response struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func ClimaHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.PathValue("cep")

	if matched, _ := regexp.MatchString(`^\d{8}$`, cep); !matched {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	city, err := viacep.GetLocationByCEP(cep)
	if err != nil {
		if err == viacep.ErrCepNotFound {
			http.Error(w, "can not find zipcode", http.StatusNotFound)
			return
		}
		http.Error(w, "erro ao consultar via cep", http.StatusInternalServerError)
		return
	}

	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		http.Error(w, "api key da weatherapi não configurada", http.StatusInternalServerError)
		return
	}

	tempC, err := weather.GetCurrentTemperature(city, apiKey)
	if err != nil {
		http.Error(w, "erro ao consultar weather api", http.StatusInternalServerError)
		return
	}

	resp := Response{
		TempC: tempC,
		TempF: temperature.CelsiusToFahrenheit(tempC),
		TempK: temperature.CelsiusToKelvin(tempC),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
