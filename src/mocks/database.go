package mocks

import (
	"github.com/golang/mock/gomock"
	databaseMocks "gitlab.com/altiano/goreen-tea/src/frameworks/database/mocks"
)

type (
	Database struct {
		Db           *databaseMocks.MockDb
		Coll         *databaseMocks.MockColl
		SingleResult *databaseMocks.MockSingleResult
		Cursor       *databaseMocks.MockCursor
	}
)

/*
	Database
*/
func setupDatabase(ctrl *gomock.Controller) Database {
	mockDb := databaseMocks.NewMockDb(ctrl)
	mockColl := databaseMocks.NewMockColl(ctrl)
	mockCursor := databaseMocks.NewMockCursor(ctrl)

	mockSr := databaseMocks.NewMockSingleResult(ctrl)

	return Database{
		Db:           mockDb,
		Coll:         mockColl,
		Cursor:       mockCursor,
		SingleResult: mockSr,
	}
}
