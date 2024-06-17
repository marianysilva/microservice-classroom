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

func (r ClassroomRepository) DeleteClassroom(classroom *domain.DeletedClassroom) error {
	stmt, err := r.statement(deleteClassroomByUUID)
	if err != nil {
		return err
	}

	args := []interface{}{
		classroom.UUID,
	}
	if err := stmt.Get(classroom, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting classroom")
	}
	return nil
}

func (r ClassroomRepository) DeleteClassroomsByCourse(classroom *domain.DeletedClassroom) ([]domain.DeletedClassroom, error) {
	stmt, err := r.statement(deleteClassroomsByCourseUUID)
	if err != nil {
		return []domain.DeletedClassroom{}, err
	}

	args := []interface{}{
		classroom.CourseUUID,
	}
	var deleted []domain.DeletedClassroom
	if err := stmt.Get(&deleted, args...); err != nil {
		return []domain.DeletedClassroom{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting classroom by course")
	}
	return deleted, nil
}

func (r ClassroomRepository) DeleteClassroomsBySubject(classroom *domain.DeletedClassroom) ([]domain.DeletedClassroom, error) {
	stmt, err := r.statement(deleteClassroomsBySubjectUUID)
	if err != nil {
		return []domain.DeletedClassroom{}, err
	}

	args := []interface{}{
		classroom.SubjectUUID,
	}
	var deleted []domain.DeletedClassroom
	if err := stmt.Get(&deleted, args...); err != nil {
		return []domain.DeletedClassroom{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting classroom by subject")
	}
	return deleted, nil
}

func (r ClassroomRepository) DeleteClassrooms(classroom *domain.DeletedClassroom) ([]domain.DeletedClassroom, error) {
	if classroom.SubjectUUID != nil {
		return r.DeleteClassroomsBySubject(classroom)
	}

	return r.DeleteClassroomsByCourse(classroom)
}
