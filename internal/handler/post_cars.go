package handler

import (
	"car-catalog/internal/additional_info"
	"car-catalog/internal/db"
	"car-catalog/internal/server"
	"encoding/json"
	"log"
	"net/http"
)

func InsertCars(w http.ResponseWriter, r *http.Request) {
	var regNums db.Car
	err := json.NewDecoder(r.Body).Decode(&regNums)
	if err != nil {
		server.Error(map[string]interface{}{"message": "error decoding"}, w)
		return
	}
	var ids []int64

	for _, regNum := range regNums.RegNums {
		log.Println(1)
		carInfo, err := additional_info.Get(regNum)
		if err != nil {
			server.Error(map[string]interface{}{"message": err.Error()}, w)
			return
		}
		log.Println(2)

		id, err := db.InsertCar(carInfo)
		if err != nil {
			server.Error(map[string]interface{}{"message": err.Error()}, w)
			return
		}
		log.Println(3)

		ids = append(ids, id)
	}
	idsJSON, err := json.Marshal(map[string][]int64{"ids": ids})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(idsJSON)
}
