package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"example.com/monorepo-backend/responses"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleTodoCreate(t *testing.T) {
	tests := []Case{}
	func() {
		formData := make(url.Values)
		formData["name"] = []string{"Test"}
		request, err := http.NewRequest("POST", "/api/v1/todo/create", bytes.NewBuffer([]byte(formData.Encode())))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		tests = append(tests, Case{
			name:     "Create task",
			r:        request,
			wantCode: http.StatusOK,
			wantData: map[string]interface{}{
				"name":   "Test",
				"status": "undone",
			},
		})
	}()

	func() {
		formData := make(url.Values)
		formData["name"] = []string{"Test"}
		request, err := http.NewRequest("POST", "/api/v1/todo/create", bytes.NewBuffer([]byte(formData.Encode())))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		tests = append(tests, Case{
			name:     "Conflict",
			r:        request,
			wantCode: http.StatusConflict,
			wantData: nil,
		})
	}()

	func() {
		formData := make(url.Values)
		formData["name"] = []string{""}
		request, err := http.NewRequest("POST", "/api/v1/todo/create", bytes.NewBuffer([]byte(formData.Encode())))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		tests = append(tests, Case{
			name:     "Bad request",
			r:        request,
			wantCode: http.StatusBadRequest,
			wantData: nil,
		})
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewServer()
			rr := httptest.NewRecorder()
			s.HandleCreateToDo(rr, tt.r)
			res := rr.Result()
			assert.Equal(t, tt.wantCode, res.StatusCode)
			if tt.wantCode == http.StatusOK {
				decoder := json.NewDecoder(res.Body)
				response := responses.GeneralResponse{}
				err := decoder.Decode(&response)
				assert.Empty(t, err)
				mp, ok := response.Data.(map[string]interface{})
				assert.Equal(t, true, ok)
				assert.Equal(t, tt.wantData["name"], mp["name"])
				assert.Equal(t, tt.wantData["status"], mp["status"])
				assert.Equal(t, true, ok)
			}
		})
	}
}
