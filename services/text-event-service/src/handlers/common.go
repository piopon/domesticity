package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func readEventID(request *http.Request) primitive.ObjectID {
	params := mux.Vars(request)
	id, error := primitive.ObjectIDFromHex(params["id"])
	if error != nil {
		panic(error)
	}
	return id
}
