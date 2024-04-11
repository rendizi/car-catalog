package main_test

import (
	. "github.com/Eun/go-hit"
	"net/http"
	"testing"
)

func TestUpdateTests(t *testing.T) {
	Test(t,
		Description("Update"),
		Put("http://localhost:8080/cars?id=1&mark=lada"),
		Expect().Status().Equal(http.StatusOK),
	)
	Test(t,
		Description("Update"),
		Put("http://localhost:8080/cars?id=2&name=Ktoto"),
		Expect().Status().Equal(http.StatusOK),
	)
}
