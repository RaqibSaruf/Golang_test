package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/RaqibSaruf/student-management/internal/config"
	"github.com/RaqibSaruf/student-management/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
	 id INTEGER PRIMARY KEY AUTOINCREMENT,
     name TEXT NOT NULL,
     age INTEGER NOT NULL,
     email TEXT NOT NULL UNIQUE
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{Db: db}, nil

}

func (s *Sqlite) GetDB() *sql.DB {
	return s.Db
}

func (s *Sqlite) GetStudentById(id int64) (*types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT * FROM students WHERE id = ? Limit 1")

	if err != nil {
		return &types.Student{}, err
	}

	defer stmt.Close()

	var student types.Student
	err = stmt.QueryRow(id).Scan(&student.ID, &student.Name, &student.Age, &student.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return &types.Student{}, fmt.Errorf("no student found with id %d", id)
		}

	}

	return &student, nil
}

func (s *Sqlite) GetStudentList() ([]types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM students")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []types.Student

	for rows.Next() {
		var student types.Student
		err := rows.Scan(&student.ID, &student.Name, &student.Email, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}

func (s *Sqlite) CreateStudent(student *types.Student) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO students (name, age, email) VALUES (?, ?, ?)")

	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(student.Name, student.Age, student.Email)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Sqlite) UpdateStudent(student *types.Student) (*types.Student, error) {
	stmt, err := s.Db.Prepare("UPDATE students SET name = ?, age = ?, email = ? WHERE id = ?")

	if err != nil {
		return student, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(student.Name, student.Age, student.Email, student.ID)
	if err != nil {
		return student, err
	}

	return student, nil
}

func (s *Sqlite) DeleteStudentById(id int64) (bool, error) {
	stmt, err := s.Db.Prepare("DELETE FROM students WHERE id = ?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return false, err
	}

	return true, nil
}
