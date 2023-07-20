package main

import (
	"echohandlertest/application"
	customermemoryrepository "echohandlertest/domain/customer/customerMemoryRepository"
	productmemoryrepository "echohandlertest/domain/product/productMemoryRepository"
	orderservice "echohandlertest/services/orderService"
	tavernservice "echohandlertest/services/tavernService"
	"log"
	"net/http"
)

func main() {
	pr := productmemoryrepository.New()
	ord := customermemoryrepository.New()
	service, err := orderservice.NewOrderService(orderservice.WithMemoryProductRepository(pr),
		orderservice.WithRepository(ord))

	if err != nil {
		log.Fatalln(err.Error())
	}

	t, err := tavernservice.NewTavern(tavernservice.WithProductRepository(pr),
		tavernservice.WithOrderService(service), tavernservice.WithCustomerRepository(ord))
	if err != nil {
		log.Fatalln(err.Error())
	}

	h := application.NewHandler(t)

	http.ListenAndServe("localhost:8080", h.InitRouter())
}
