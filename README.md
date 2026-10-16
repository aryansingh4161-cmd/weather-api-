# 🌤️ Weather API (Go)

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-green)
![Status](https://img.shields.io/badge/status-learning%20project-blue)

A small weather server written in Go. It takes a city name, asks the OpenWeatherMap API for the current weather, and returns the result as JSON.

## ✨ Features
- `/hello` endpoint to check the server is running
- `/weather/<city>` endpoint that returns the city name and temperature
- API key kept in a separate config file, not in the code

## 📋 Table of Contents
- [Requirements](#requirements)
- [Setup](#setup)
- [Usage](#usage)
- [Example response](#example-response)
- [Project structure](#project-structure)

## Requirements
- [Go](https://go.dev/dl/) installed
- A free API key from [OpenWeatherMap](https://openweathermap.org/api)

## Setup
1. Clone the repo:
```bash
   git clone https://github.com/aryansingh4161-cmd/weather-api-.git
   cd weather-api-
```
2. Create a file named `.apiConfig` in the project folder:
```json
   {"OpenWeatherMapApiKey": "your_key_here"}
```
3. Run the server:
```bash
   go run main.go
```

## Usage
Open these in your browser:
- `http://localhost:8080/hello`
- `http://localhost:8080/weather/Patiala`

## Example response
```json
{
  "name": "Patiala",
  "main": { "Kelvin": 301.2 }
}
```
Temperature is in Kelvin. Celsius = Kelvin − 273.15.

## Project structure
```
.
├── main.go          # server and API logic
├── go.mod           # Go module file
├── .gitignore       # keeps .apiConfig out of git
└── README.md
```

## 🔒 Note
`.apiConfig` is listed in `.gitignore`. Never commit your API key.

## 📄 License
MIT