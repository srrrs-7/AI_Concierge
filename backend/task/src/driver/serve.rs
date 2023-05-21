package driver

import (
	"fmt"
	"net/http"
	"template/pkg"
	"template/util/env"
	"template/util/log"

	"github.com/gorilla/mux"
)

func NewRouter(
	env *env.Env,
	logicRepo pkg.Repositories,

) {
	r := mux.NewRouter()
	r.Use(log.NewLogging)

	r.HandleFunc("/logicRepo/", logicRepo.MainLogic).Methods("GET")
	r.HandleFunc("/second/", logicRepo.SecondLogic).Methods("GET")
	r.HandleFunc("/sub/", logicRepo.SubLogic).Methods("GET")

	fmt.Println("start server on port:", env.HttpPort)
	err := http.ListenAndServe(":"+env.HttpPort, r)
	if err != nil {
		panic(err)
	}
}
