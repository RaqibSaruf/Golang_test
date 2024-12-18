package storage

import (
	"database/sql"

	"github.com/RaqibSaruf/student-management/internal/types"
)

type Storage interface {
	GetDB() *sql.DB
	CreateStudent(student *types.Student) (int64, error)
	GetStudentById(id int64) (*types.Student, error)
	UpdateStudent(student *types.Student) (*types.Student, error)
	GetStudentList() ([]types.Student, error)
	DeleteStudentById(id int64) (bool, error)
}
