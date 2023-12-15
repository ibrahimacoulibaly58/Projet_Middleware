package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"Projet_Middleware/User/internal/models"
)

// Ctx is a middleware to extract user ID from URL parameters and add it to the request context.
func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := uuid.FromString(chi.URLParam(r, "id"))
		if err != nil {
			logrus.Errorf("failed to parse user ID: %s", err)
			customError := &models.CustomError{
				Message: fmt.Sprintf("invalid user ID format: %s", chi.URLParam(r, "id")),
				Code:    http.StatusUnprocessableEntity,
			}
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))

		logrus.Infof("User ID %s extracted and added to the context", userID)
	})
}
