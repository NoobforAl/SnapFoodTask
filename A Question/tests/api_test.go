package routes

import (
	"A-Question/Api/router"
	"A-Question/NodeFoodAdder/core"
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type testCaseSample struct {
	detail       string
	route        string
	expectedCode int
}

func TestOrderRoute(t *testing.T) {
	test := testCaseSample{
		detail:       "api not work!",
		route:        "/api/order",
		expectedCode: 200,
	}

	app := fiber.New()
	router.Router(app)

	req := httptest.NewRequest("POST", test.route, bytes.NewBufferString(`{
		"order_id":1,
		"price": 2000,
		"title":"sample"
	}`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, 100)
	if err != nil {
		t.Error(err.Error())
		return
	}

	assert.Equalf(t, test.expectedCode, resp.StatusCode, test.detail)
}

func TestNodeFoodAdder(t *testing.T) {
	_, err := core.GetFood()
	if err != nil {
		t.Error("can't get data, err: ", err.Error())
	}
}

func TestNotFoundRoute(t *testing.T) {
	test := testCaseSample{
		detail:       "not found page api not work!",
		route:        "/not-found",
		expectedCode: 404,
	}

	app := fiber.New()
	router.Router(app)

	req := httptest.NewRequest("POST", test.route, nil)
	resp, _ := app.Test(req, 1)
	assert.Equalf(t, test.expectedCode, resp.StatusCode, test.detail)

}
