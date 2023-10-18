package account_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Lucas-Linhar3s/TheBarber/interfaces"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func TestCreateAccount(t *testing.T) {
	r := gin.Default()
	n := r.Group("/v1")
	interfaces.Router(n)

	req := map[string]interface{}{
		"name":          "Lucas",
		"last_name":     "Linhares",
		"phone":         "11999999999",
		"date_of_birth": "1999-01-01",
		"photo":         "https://avatars.githubusercontent.com/u/60052506?v=4",
		"is_barber":     false,
		"email":         "test@teste.com",
		"password":      "123456789",
	}

	body, err := json.Marshal(req)
	if err != nil {
		t.Error(err)
	}

	t.Run("Create account", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/v1/account/create", bytes.NewBuffer(body))
		if err != nil {
			t.Error(err)
		}

		r.ServeHTTP(w, req)
		assert.Equal(t, 201, w.Code)
	})

}
