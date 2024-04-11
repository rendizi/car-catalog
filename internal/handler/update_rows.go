package handler

import (
	"car-catalog/internal/db"
	"car-catalog/internal/server"
	"net/http"
	"strconv"
)

func UpdateCars(w http.ResponseWriter, r *http.Request) {
	filters := r.URL.Query()
	var year int
	var err error
	if filters.Get("year") != "" {
		year, err = strconv.Atoi(filters.Get("year"))
		if err != nil {
			server.Error(map[string]interface{}{"message": "Invalid year"}, w)
			return
		}
	}
	year = 0
	changed := db.CarInfo{
		RegNum: filters.Get("regNum"),
		Model:  filters.Get("model"),
		Mark:   filters.Get("mark"),
		Year:   year,
		Owner: db.People{
			Name:       filters.Get("name"),
			Surname:    filters.Get("surname"),
			Patronymic: filters.Get("patronymic"),
		},
	}

	if filters.Get("id") == "" {
		server.Error(map[string]interface{}{"message": "Id is not provided"}, w)
		return
	}

	id, err := strconv.Atoi(filters.Get("id"))
	if err != nil {
		server.Error(map[string]interface{}{"message": "Invalid id"}, w)
		return
	}

	err = db.UpdateRow(changed, int64(id))
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}

	server.Ok(map[string]interface{}{"message": "update successful"}, w)

}
