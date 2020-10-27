package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func readEventID(request *http.Request) int {
	id, error := strconv.Atoi(mux.Vars(request)["id"])
	if error != nil {
		panic(error)
	}
	return id
}
