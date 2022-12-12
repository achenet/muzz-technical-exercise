package handlers

import (
	"encoding/json"
	"fmt"
	td "github.com/maxatome/go-testdeep"
	"github.com/maxatome/go-testdeep/helpers/tdhttp"
	"muzz/db"
	"net/http"
	"testing"
)

func TestBasic(t *testing.T) {
	testdb, err := db.Connect()
	td.CmpNoError(t, err)
	ta := tdhttp.NewTestAPI(t, NewAPI(testdb))
	ta.PostJSON("/user/create", json.RawMessage(`{}`)).
		CmpStatus(http.StatusCreated).OrDumpResponse()

	row := testdb.QueryRow("SELECT id FROM profiles;")
	td.CmpNoError(t, err)
	var id int
	err = row.Scan(&id)
	td.CmpNoError(t, err)

	ta.Get(fmt.Sprintf("/profiles?user-id=%d", id)).CmpStatus(http.StatusOK)

	ta.PostJSON("/swipe", json.RawMessage(`{}`)).
		CmpStatus(http.StatusOK).
		CmpResponse(td.JSON(`{}`)).
		OrDumpResponse()
}
