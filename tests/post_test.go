package main_test

import (
	. "github.com/Eun/go-hit"
	"net/http"
	"testing"
)

func TestPostTests(t *testing.T) {
	Test(t,
		Description("Post"),
		Post("http://localhost:8080/cars"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().JSON(map[string][]string{"regNums": []string{"X123XX119", "X123XX122", "X123XX121", "X123XX120"}}),
		Expect().Status().Equal(http.StatusOK),
	)
	Test(t,
		Description("Post"),
		Post("http://localhost:8080/cars"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().JSON(map[string][]string{"regNums": []string{"X123XX124"}}),
		Expect().Status().Equal(http.StatusOK),
	)
}
