package main

import (
	"context"
	"log"
	"longsetstring/handler"
	"longsetstring/stringService"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s := stringService.NewStringService()

	handler := handler.Newhandler(s)

	RunServer(handler)
}

func RunServer(handler handler.Handler) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	server := http.Server{
		Addr:    ":8080",
		Handler: handler.InitRouter(),
	}

	go func() {
		log.Println("Server start")
		err := server.ListenAndServe()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	<-ch

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := server.Shutdown(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Server stoped")
}
