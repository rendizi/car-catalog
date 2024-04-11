package db

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
)

type Car struct {
	RegNums []string `json:"regNums"`
}

type CarInfo struct {
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int    `json:"year"`
	Owner  People `json:"owner"`
}

type People struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

func InsertCar(info CarInfo) (int64, error) {
	if db == nil {
		return 0, errors.New("database connection is nil")
	}

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM cars WHERE regNum = $1)", info.RegNum).Scan(&exists)
	if err != nil {
		return 0, err
	}
	if exists {
		return 0, errors.New(info.RegNum + " already exists in db")
	}

	var id int64
	query := `
        INSERT INTO cars(regNum, mark, model, carYear, ownerName, ownerSurname, ownerPatronymic)
        VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id
    `
	err = db.QueryRow(query, info.RegNum, info.Mark, info.Model, info.Year, info.Owner.Name, info.Owner.Surname, info.Owner.Patronymic).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func UpdateRow(info CarInfo, id int64) error {
	changed := Changed(info)
	query := "UPDATE cars SET  "
	i := 2
	vals := make([]interface{}, 0)

	for key, val := range changed {
		vals = append(vals, val)
		query += key + " = $" + strconv.Itoa(i) + ","
		i++
	}

	if query == "UPDATE cars SET  " {
		return nil
	}

	query = query[:len(query)-1]
	query += " WHERE id = $1 "
	log.Println(query)
	vals = append([]interface{}{id}, vals...)
	_, err := db.Exec(query, vals...)
	if err != nil {
		return err
	}
	return nil
}

func GetCars(info CarInfo, limit int, offset int) ([]CarInfo, error) {
	changed := Changed(info)
	query := "SELECT regNum, mark, model, carYear, ownerName,ownerSurname,ownerPatronymic FROM cars WHERE  "
	i := 3
	vals := make([]interface{}, 0)
	for key, val := range changed {
		vals = append(vals, val)
		query += key + " = $" + strconv.Itoa(i) + " AND "
		i++
	}
	query = query[:len(query)-5]
	query += " LIMIT $1 OFFSET $2"
	vals = append([]interface{}{limit, offset}, vals...)

	rows, err := db.Query(query, vals...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []CarInfo
	for rows.Next() {
		var car CarInfo
		var patronymic sql.NullString
		err = rows.Scan(&car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner.Name, &car.Owner.Surname, &patronymic)
		if err != nil {
			return nil, err
		}
		if patronymic.Valid {
			car.Owner.Patronymic = patronymic.String
		}
		cars = append(cars, car)
	}

	return cars, nil
}

func DeleteCar(id int64) error {
	query := `DELETE FROM cars WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func Changed(info CarInfo) map[string]interface{} {
	res := make(map[string]interface{})
	if info.RegNum != "" {
		res["regNum"] = info.RegNum
	}
	if info.Mark != "" {
		res["mark"] = info.Mark
	}
	if info.Model != "" {
		res["model"] = info.Model
	}
	if info.Year != 0 {
		res["carYear"] = info.Year
	}
	if info.Owner.Name != "" {
		res["ownerName"] = info.Owner.Name
	}
	if info.Owner.Surname != "" {
		res["ownerSurname"] = info.Owner.Surname
	}
	if info.Owner.Patronymic != "" {
		res["ownerPatronymic"] = info.Owner.Patronymic
	}

	return res
}
