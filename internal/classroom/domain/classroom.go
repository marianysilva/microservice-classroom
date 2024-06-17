package domain

import (
	"time"

	"github.com/google/uuid"
)

type Classroom struct {
	UUID         uuid.UUID  `json:"uuid"`
	Code         string     `json:"code"`
	SubjectUUID  *uuid.UUID `db:"subject_uuid" json:"subject_uuid"`
	CourseUUID   uuid.UUID  `db:"course_uuid" json:"course_uuid"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Format       string     `json:"format"`
	CanSubscribe bool       `db:"can_subscribe" json:"can_subscribe"`
	StartsAt     time.Time  `db:"starts_at" json:"starts_at"`
	EndsAt       *time.Time `db:"ends_at" json:"ends_at"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updated_at"`
}

type DeletedClassroom struct {
	UUID        *uuid.UUID `json:"uuid"`
	SubjectUUID *uuid.UUID `db:"subject_uuid" json:"subject_uuid"`
	CourseUUID  *uuid.UUID `db:"course_uuid" json:"course_uuid"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at"`
}
