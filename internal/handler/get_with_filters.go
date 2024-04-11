package handler

import (
	"car-catalog/internal/db"
	"car-catalog/internal/server"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetWFilters(w http.ResponseWriter, r *http.Request) {
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

	limit, page := filters.Get("limit"), filters.Get("page")
	var limitInd, pageInd int
	if limit == "" {
		limitInd = 10
	} else {
		limitInd, err = strconv.Atoi(limit)
		if err != nil {
			server.Error(map[string]interface{}{"message": "Invalid limit"}, w)
			return
		}
	}

	if page == "" {
		pageInd = 0
	} else {
		pageInd, err = strconv.Atoi(page)
		if err != nil {
			server.Error(map[string]interface{}{"message": "Invalid page"}, w)
			return
		}
	}

	result, err := db.GetCars(changed, limitInd, pageInd)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		server.Error(map[string]interface{}{"message": err.Error()}, w)
		return
	}
	w.Write(resultJson)
}
