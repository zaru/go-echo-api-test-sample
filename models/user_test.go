package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zaru/go-echo-api-test-sample/test"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFindByID(t *testing.T) {
	mockDB, mock, sqlxDB := test.MockDB(t)
	defer mockDB.Close()

	var cols []string = []string{"id", "name"}
	mock.ExpectQuery("SELECT *").WillReturnRows(sqlmock.NewRows(cols).
		AddRow(1, "foobar"))

	um := NewUserModel(sqlxDB)
	u := um.FindByID("1")

	expect := User{
		ID:   1,
		Name: "foobar",
	}
	assert.Equal(t, expect, u)
}

func TestFindAll(t *testing.T) {
	mockDB, mock, sqlxDB := test.MockDB(t)
	defer mockDB.Close()

	u1 := User{ID: 1, Name: "foobar"}
	u2 := User{ID: 2, Name: "barbaz"}

	var cols []string = []string{"id", "name"}
	mock.ExpectQuery("SELECT *").WillReturnRows(sqlmock.NewRows(cols).
		AddRow(u1.ID, u1.Name).
		AddRow(u2.ID, u2.Name))

	um := NewUserModel(sqlxDB)
	u := um.FindAll()

	expect := []User{}
	expect = append(expect, u1, u2)
	assert.Equal(t, expect, u)
}
