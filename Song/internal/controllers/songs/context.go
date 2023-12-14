package songs

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"Projet_Middleware/Song/internal/models"
)

// Ctx is a middleware to extract song ID from URL parameters and add it to the request context.
func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		songID, err := uuid.FromString(chi.URLParam(r, "id"))
		if err != nil {
			logrus.Errorf("failed to parse song ID: %s", err)
			customError := &models.CustomError{
				Message: fmt.Sprintf("invalid song ID format: %s", chi.URLParam(r, "id")),
				Code:    http.StatusUnprocessableEntity,
			}
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
			return
		}

		ctx := context.WithValue(r.Context(), "songID", songID)
		next.ServeHTTP(w, r.WithContext(ctx))

		logrus.Infof("Song ID %s extracted and added to the context", songID)
	})
}