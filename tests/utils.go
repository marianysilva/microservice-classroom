package tests

import (
	"fmt"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-classroom/tests/database"
)

var (
	Now                       = time.Now()
	ClassroomUUID             = uuid.MustParse("00000000-0000-0000-0000-aaaaaaaaaaaa")
	ClassroomSubscriptionUUID = uuid.MustParse("00000000-0000-0000-0000-bbbbbbbbbbbb")
	ClassroomLessonUUID       = uuid.MustParse("00000000-0000-0000-0000-cccccccccccc")
	ClassroomTimetable        = uuid.MustParse("00000000-0000-0000-0000-dddddddddddd")
	UserUUID                  = uuid.MustParse("00000000-0000-0000-0000-111111111111")
	CourseUUID                = uuid.MustParse("00000000-0000-0000-0000-222222222222")
	SubjectUUID               = uuid.MustParse("00000000-0000-0000-0000-333333333333")
	EmptyRows                 = sqlmock.NewRows([]string{})
)

func NewTestDB(queries map[string]string) (*sqlx.DB, sqlmock.Sqlmock, map[string]*sqlmock.ExpectedPrepare) {
	db, mock := database.NewDBMock()

	sqlStatements := make(map[string]*sqlmock.ExpectedPrepare)
	for queryName, query := range queries {
		stmt := mock.ExpectPrepare(fmt.Sprintf("^%s$", regexp.QuoteMeta(string(query))))
		sqlStatements[queryName] = stmt
	}

	mock.MatchExpectationsInOrder(false)
	return db, mock, sqlStatements
}
