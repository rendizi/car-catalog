package additional_info

import (
	"car-catalog/internal/db"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"regexp"
)

func Get(regNum string) (db.CarInfo, error) {
	pattern := `^[A-Z]\d{3}[A-Z]{2}\d{2,3}$`
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(regNum) {
		return db.CarInfo{}, errors.New("wrong format of regional number")
	}
	err := godotenv.Load(".env")
	if err != nil {
		return db.CarInfo{}, err
	}
	url := os.Getenv("API_URL")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return db.CarInfo{}, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return db.CarInfo{}, err
	}

	if resp.StatusCode != 200 {
		return db.CarInfo{}, fmt.Errorf("status code is not 200")
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return db.CarInfo{}, err
	}

	var carInfo db.CarInfo
	if err := json.Unmarshal(bodyBytes, &carInfo); err != nil {
		return db.CarInfo{}, err
	}

	return carInfo, nil
}
