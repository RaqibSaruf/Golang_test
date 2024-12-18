package reqValidator

import (
	"database/sql"
	"fmt"
	"log/slog"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ReqValidator struct {
	validator *validator.Validate
	db        *sql.DB
}

func New(db *sql.DB) *ReqValidator {
	v := validator.New()

	// Register "uniqueEmail" validation rule
	v.RegisterValidation("unique", uniqueValidation(db))

	return &ReqValidator{
		validator: v,
		db:        db,
	}
}

func (cv *ReqValidator) Struct(s interface{}) error {
	return cv.validator.Struct(s)
}

func uniqueValidation(db *sql.DB) func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		email := fl.Field().String()
		params := strings.Split(fl.Param(), ";") // Extract the table name from the tag parameter
		if len(params) < 2 {
			slog.Info("Missing required parameters for unique validation")
			return false
		}

		tableName := params[0]
		columnName := params[1]
		// Query the database to check if the email exists
		query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s = ?)", tableName, columnName)
		var exists bool
		err := db.QueryRow(query, email).Scan(&exists)
		if err != nil {
			slog.Info("Database query error:", "error", err.Error())
			return false
		}

		// Return true if email does NOT exist (is unique)
		return !exists
	}
}
