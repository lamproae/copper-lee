package decoding

import (
	"github.com/cheekybits/is"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TesthandleCreateUser(t *testing.T) {
	is := is.New(t)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/user", strings.NewReader(`{
		"email":"liwei@google.com"
	}`))

	is.NoErr(err)

	handleCreateUser(w, r)
	is.Equal(w.Code, http.StatusCreated)
}
