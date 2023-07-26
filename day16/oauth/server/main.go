package main

import (
	"log"
	"mytesting/pkg/handler"
	authrepository "mytesting/pkg/repository/auth_repository"
	clientrepository "mytesting/pkg/repository/client_repository"
	codesrepository "mytesting/pkg/repository/codes_repository"
	userrepository "mytesting/pkg/repository/user_repository"
	"net/http"
)



func main() {
	a := authrepository.NewAuth([]byte("vnpanrgiuhrugnaernglr;g;oirjJpHPFH"), []byte("g4RO#rp3krpokr2kRPK3k[pC"))
	c := clientrepository.NewClient()
	u := userrepository.NewUserRepository()
	codes, err := codesrepository.NewCodes("localhost:6379")
	if err != nil {
		log.Fatalln(err)
	}

	h := handler.NewHandler(a, c, u, codes)

	http.ListenAndServe(":8080", h.InitRouter())
}
// 820378