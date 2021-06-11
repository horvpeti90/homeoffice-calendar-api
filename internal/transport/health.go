package transport

import (
	"net/http"

	"github.com/unrolled/render"
)

type HealthResponse struct {
	Status bool `json:"status" example:"true"`
}

func NewHealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := HealthResponse{
			Status: true,
		}

		renderer := render.New()

		err := renderer.JSON(w, http.StatusOK, status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
