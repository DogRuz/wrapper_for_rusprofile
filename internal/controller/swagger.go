package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SwaggerHandler struct{}

func (s SwaggerHandler) Read(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("swagger/api.swagger.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error":  err.Error(),
			"result": nil,
		})
		if err != nil {
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		return
	}
}
