package todo_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/ifreddyrondon/go-workshop/santiago-nov2018/resources/src/21_chi/3/todo"
)

func setup() (*httptest.Server, func()) {
	handler := todo.Routes()

	// run server using httptest
	server := httptest.NewServer(handler)
	teardown := func() {
		server.Close()
	}
	return server, teardown
}

func TestGetAllTodos(t *testing.T) {
	t.Parallel()

	server, teardown := setup()
	defer teardown()

	e := httpexpect.New(t, server.URL)

	// is it working?
	obj := e.GET("/").
		Expect().
		Status(http.StatusOK).JSON().Array().NotEmpty().First()

	obj.Object().ContainsKey("slug")
}
