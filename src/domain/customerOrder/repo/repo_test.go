package repo

import (
	"testing"

	"gitlab.com/altiano/goreen-tea/src/mocks"
)

func setupRepo(t *testing.T) (m struct {
	mocks.RepoBase
	sut Repo
}) {
	m.RepoBase = mocks.SetupRepoBase(t)
	m.Database.Db.EXPECT().Collection("customerOrder").Return(m.Database.Coll)

	m.sut = NewRepo(m.Trace.Tracer, m.Database.Db)
	return m
}
