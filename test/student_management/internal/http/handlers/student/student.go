package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/RaqibSaruf/student-management/internal/storage"
	"github.com/RaqibSaruf/student-management/internal/types"
	"github.com/RaqibSaruf/student-management/internal/utils/response"
	reqValidator "github.com/RaqibSaruf/student-management/internal/utils/validator"
	"github.com/go-playground/validator/v10"
)

func Create(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			// to send direct error message
			// response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			//to send custom error message
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// request validation
		if err := reqValidator.New(storage.GetDB()).Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusUnprocessableEntity, response.ValidationError(student, validateErrs))
			return
		}

		lastId, err := storage.CreateStudent(&student)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		student.ID = lastId
		slog.Info("creating a student")

		response.WriteJson(w, http.StatusCreated, response.Success(student, "Student created successfully"))
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		student, err := storage.GetStudentById(intId)

		if err != nil {
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, response.Success(student, "Student found successfully"))
	}
}

func Update(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		var student types.Student

		parseErr := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(parseErr, io.EOF) {
			// to send direct error message
			// response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			//to send custom error message
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if parseErr != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(parseErr))
			return
		}

		// request validation
		if parseErr := reqValidator.New(storage.GetDB()).Struct(student); parseErr != nil {
			validateErrs := parseErr.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusUnprocessableEntity, response.ValidationError(student, validateErrs))
			return
		}

		student.ID = intId

		std, err := storage.UpdateStudent(&student)

		if err != nil {
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, response.Success(std, "Student updated successfully"))
	}
}

func Delete(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		_, err = storage.DeleteStudentById(intId)

		if err != nil {
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, response.Success(nil, "Student deleted successfully"))
	}
}

func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		students, err := storage.GetStudentList()

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, response.Success(students, "Students found successfully"))
	}
}
