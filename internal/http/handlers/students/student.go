package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/Aditya7880900936/Students_API/internal/types"
	"github.com/Aditya7880900936/Students_API/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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
		slog.Info("Creating New Student")

		// w.Write([]byte("Welcome to Students API"))

		response.WriteJson(w, http.StatusCreated, map[string]string{"message": "Student Created Successfully"})
	}
}
