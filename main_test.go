package main_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/lucasstettner/api-boilerplate/app"
)

var a app.App

// This tests the main function of the app, it is necessary that this passes
// Therefore we use the testing.Main package
func TestMain(m *testing.M) {
	a = app.App{}

	a.Start(false)

	if err := a.Config.DB.Ping(); err != nil {
		log.Fatalf("Error pinging conn: %s", err)
		return
	}

	os.Exit(m.Run())
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}
