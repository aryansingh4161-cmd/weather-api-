package main 

import (
	"encoding/json" 
	"net/http"      
	"os"            
	"strings"       
)


type apiConfigdata struct {
	
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`

}


type weatherData struct {
	Name string `json:"name"` 
	Main struct { 

		Kelvin float64 `json:"temp"` 
		
	} `json:"main"` 

}


func loadApiKey(filename string) (apiConfigdata, error) {

	data, err := os.ReadFile(filename)

	if err != nil {                   
		return apiConfigdata{}, err 
	}

	var config apiConfigdata       
	err = json.Unmarshal(data, &config) 

	if err != nil {                 
		return apiConfigdata{}, err 
	}
	return config, nil 
}


func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from go!\n")) 
}


func query(city string) (weatherData, error) {
	apiConfig, err := loadApiKey(".apiConfig") 

	if err != nil {                           
		return weatherData{}, err 
	}

	
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiConfig.OpenWeatherMapApiKey)
	if err != nil { 
		
		return weatherData{}, err 
	}
	defer resp.Body.Close() 

	var d weatherData 

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err 
	}
	return d, nil 
	
}

func main() {
	http.HandleFunc("/hello", hello) 


	
	http.HandleFunc("/weather/",
		func(w http.ResponseWriter, r *http.Request) {
			
			city := strings.SplitN(r.URL.Path, "/", 3)[2]
			data, err := query(city) 
			if err != nil {          

				http.Error(w, err.Error(), http.StatusInternalServerError) 
				return  

			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8") 

			json.NewEncoder(w).Encode(data) 
		})

	http.ListenAndServe(":8080", nil) 
}
// run this code in terminal and then 
//http://localhost:8080/weather/bhopal type this in your browser 
// change the city to get its temperature 
