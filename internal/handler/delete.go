package handler

import (
	"car-catalog/internal/db"
	"car-catalog/internal/server"
	"net/http"
	"strconv"
)

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		server.Error(map[string]interface{}{"message": "id it not provided"}, w)
		return
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		server.Error(map[string]interface{}{"message": err}, w)
		return
	}
	err = db.DeleteCar(int64(intId))
	if err != nil {
		server.Error(map[string]interface{}{"message": err}, w)
		return
	}

	server.Ok(map[string]interface{}{"message": "delete successful"}, w)
}
