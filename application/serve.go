package main

import (
	"curve-tech-test/application/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/health", controllers.Health)
	http.HandleFunc("/v1/rates", controllers.Rates)
	http.HandleFunc("/v1/predict", controllers.Predict)

	println("Server Started, please visit: http://127.0.0.1:8091/v1/health")
	err := http.ListenAndServe(":8091", nil)
	if err != nil {
		panic(err)
	}
}
