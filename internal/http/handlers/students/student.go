package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/Aditya7880900936/Students_API/internal/types"
	"github.com/Aditya7880900936/Students_API/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating New Student")
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		// if err != nil {
		// 	slog.Error("Failed to Decode the Request Body", slog.Any("Error", err))

		// 	w.WriteHeader(http.StatusBadRequest)

		// 	return
		// }

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// request validation

		if err := validator.New().Struct(student); err!= nil {
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(err.(validator.ValidationErrors)))
			return
		}

		// w.Write([]byte("Welcome to Students API"))

		response.WriteJson(w, http.StatusCreated, map[string]string{"message": "Student Created Successfully"})
	}
}
