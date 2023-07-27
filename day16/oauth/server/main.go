package main

import (
	"mytesting/pkg/handler"
	authrepository "mytesting/pkg/repository/auth_repository"
	clientrepository "mytesting/pkg/repository/client_repository"
	userrepository "mytesting/pkg/repository/user_repository"
	"net/http"
)

func main() {
	a := authrepository.NewAuth([]byte("vnpanrgiuhrugnaernglr;g;oirjJpHPFH"), []byte("g4RO#rp3krpokr2kRPK3k[pC"))
	c := clientrepository.NewClient()
	u := userrepository.NewUserRepository()

	h := handler.NewHandler(a, c, u)

	http.ListenAndServe(":8080", h.InitRouter())
}

// 820378
