package main

import (
	"fmt"
	"github.com/mike5xe/bonus-service/internal/config"
	"net/http"
)

func main() {
	cfg, err := config.LoadMainConfig()
	if err != nil {
		panic("failed to load configuration: " + err.Error())
	}

	fmt.Printf("Hello,zaebal. We are startanuli on port :%d\n", cfg.Service.Port)
	if err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Service.Port), http.HandlerFunc(helloHandler)); err != nil {
		panic("failed to run server: " + err.Error())
	}
}

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"text":"hello"}`))
}
