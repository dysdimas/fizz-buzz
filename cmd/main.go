package main

import (
	"github.com/dysdimas/internal/handler"
	"github.com/dysdimas/internal/server"
	"github.com/dysdimas/internal/usecase"
)

func main() {
	usecase := usecase.NewFizzBuzzUsecase()
	handler := handler.NewFizzBuzzHandler(usecase)
	server.StartServer(handler)
}
