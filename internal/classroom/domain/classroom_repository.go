package domain

type ClassroomRepository interface {
	Classrooms() ([]Classroom, error)
	CreateClassroom(classroom *Classroom) error
}
