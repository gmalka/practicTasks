.PHONY: all up


all		:	down up
		@sleep 2;
		go run main.go;

up		:
		docker-compose up -d

down	:
		@docker-compose down

run		:
		go run main.go