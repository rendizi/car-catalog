package main_test

import (
	. "github.com/Eun/go-hit"
	"net/http"
	"testing"
)

func TestDeleteTests(t *testing.T) {
	Test(t,
		Description("DELETE"),
		Delete("http://localhost:8080/cars?id=1"),
		Expect().Status().Equal(http.StatusOK),
	)
	Test(t,
		Description("DELETE"),
		Delete("http://localhost:8080/cars?id=2"),
		Expect().Status().Equal(http.StatusOK),
	)
}
