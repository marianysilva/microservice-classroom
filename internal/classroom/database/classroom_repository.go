package database

import (
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-classroom/internal/classroom/domain"
	"github.com/sumelms/microservice-classroom/pkg/errors"
)

// NewClassroomRepository creates the subject subjectRepository
func NewClassroomRepository(db *sqlx.DB) (ClassroomRepository, error) { //nolint: revive
	sqlStatements := make(map[string]*sqlx.Stmt)

	for queryName, query := range queriesClassroom() {
		stmt, err := db.Preparex(query)
		if err != nil {
			return ClassroomRepository{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error preparing statement %s", queryName)
		}
		sqlStatements[queryName] = stmt
	}

	return ClassroomRepository{
		statements: sqlStatements,
	}, nil
}

type ClassroomRepository struct {
	statements map[string]*sqlx.Stmt
}

func (r ClassroomRepository) statement(s string) (*sqlx.Stmt, error) {
	stmt, ok := r.statements[s]
	if !ok {
		return nil, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", s)
	}
	return stmt, nil
}

func (r ClassroomRepository) Classrooms() ([]domain.Classroom, error) {
	stmt, err := r.statement(listClassrooms)
	if err != nil {
		return []domain.Classroom{}, err
	}

	var cc []domain.Classroom
	if err := stmt.Select(&cc); err != nil {
		return []domain.Classroom{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting classroom")
	}
	return cc, nil
}

func (r ClassroomRepository) CreateClassroom(classroom *domain.Classroom) error {
	stmt, err := r.statement(createClassroom)
	if err != nil {
		return err
	}

	args := []interface{}{
		classroom.Code,
		classroom.CourseUUID,
		classroom.SubjectUUID,
		classroom.Name,
		classroom.Description,
		classroom.CanSubscribe,
		classroom.Format,
		classroom.StartsAt,
		classroom.EndsAt,
	}
	if err := stmt.Get(classroom, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating classroom")
	}
	return nil
}
