package main_test

import (
	. "github.com/Eun/go-hit"
	"net/http"
	"testing"
)

func TestGetTests(t *testing.T) {
	Test(t,
		Description("Get"),
		Get("http://localhost:8080/cars"),
		Expect().Status().Equal(http.StatusOK),
	)
	Test(t,
		Description("Get"),
		Get("http://localhost:8080/cars?mark=Toyota"),
		Expect().Status().Equal(http.StatusOK),
	)
	Test(t,
		Description("Get"),
		Get("http://localhost:8080/cars?mark=lada"),
		Expect().Status().Equal(http.StatusOK),
	)
	Test(t,
		Description("Get"),
		Get("http://localhost:8080/cars?name=Elnara"),
		Expect().Status().Equal(http.StatusOK),
	)
}
