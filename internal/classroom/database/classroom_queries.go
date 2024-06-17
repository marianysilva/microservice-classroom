package database

import "fmt"

const (
	returningColumns = `uuid, code, course_uuid, subject_uuid,
		name, description, can_subscribe, format,
		starts_at, ends_at, created_at, updated_at`
	// CREATE.
	createClassroom = "create classroom"
	// READ.
	listClassrooms = "list classrooms"
	// UPDATE.
	// DELETE.
	deleteClassroomByUUID         = "delete classroom by UUID"
	deleteClassroomsByCourseUUID  = "delete classrooms by courseUUID"
	deleteClassroomsBySubjectUUID = "delete classrooms by subjectUUID"
)

func queriesClassroom() map[string]string {
	return map[string]string{
		// CREATE.
		createClassroom: fmt.Sprintf(`INSERT INTO classrooms (
				code, course_uuid, subject_uuid,
				name, description, can_subscribe, format,
				starts_at, ends_at
			) VALUES (
				$1, $2, $3,
				$4, $5, $6, $7,
				$8, $9)
			RETURNING %s`, returningColumns),
		// READ.
		listClassrooms: fmt.Sprintf("SELECT %s FROM classrooms WHERE deleted_at IS NULL", returningColumns),
		// DELETE.
		deleteClassroomByUUID: `UPDATE classrooms
				SET deleted_at = NOW()
			WHERE uuid = $1
				AND deleted_at IS NULL
			RETURNING uuid, deleted_at`,
		deleteClassroomsByCourseUUID: `UPDATE classrooms
				SET deleted_at = NOW()
			WHERE course_uuid = $1
				AND deleted_at IS NULL
			RETURNING uuid, deleted_at`,
		deleteClassroomsBySubjectUUID: `UPDATE classrooms
				SET deleted_at = NOW()
			WHERE subject_uuid = $1
				AND deleted_at IS NULL
			RETURNING uuid, deleted_at`,
	}
}
