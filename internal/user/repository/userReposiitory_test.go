package repository

import (
	"context"
	wrapper "github.com/go-park-mail-ru/2021_2_SaberDevs/internal"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
	"regexp"
	"testing"
)

func TestGetByName(t *testing.T) {
	log := wrapper.NewLogger()
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	urepo := NewUserRepository(db, log)

	rows := sqlxmock.NewRows([]string{"login", "name", "surname", "email", "password", "score", "avatarurl", "description"}).
		AddRow("1", "", "", "", "", 1, "", "")

	query := `SELECT Login, Name, Surname, Email, Password, Score, AvatarUrl, Description FROM author WHERE Name = $1`

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("oo").WillReturnRows(rows)

	comm, err := urepo.GetByName(context.TODO(), "oo")
	assert.NoError(t, err)
	assert.NotNil(t, comm)
}

func TestUpdateUser(t *testing.T) {
	log := wrapper.NewLogger()
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	urepo := NewUserRepository(db, log)
	user := umodels.User{
		Login: "1",
		Description: "1",
		Name: "1",
		Surname: "1",
		Password: "1",
		AvatarURL: "1",
	}


	result := sqlxmock.NewResult(0, 0)

	query1 := `UPDATE author SET Description = $1 WHERE Login = $2`
	query2 := `UPDATE author SET NAME = $1 WHERE Login = $2`
	query3 := `UPDATE author SET SURNAME = $1 WHERE Login = $2`
	query4 := `UPDATE author SET PASSWORD = $1 WHERE Login = $2`
	query5 := `UPDATE author SET AvatarUrl = $1 WHERE Login = $2`

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query1)).WithArgs(user.Description, user.Login).WillReturnResult(result)
	mock.ExpectExec(regexp.QuoteMeta(query2)).WithArgs(user.Name, user.Login).WillReturnResult(result)
	mock.ExpectExec(regexp.QuoteMeta(query3)).WithArgs(user.Surname, user.Login).WillReturnResult(result)
	mock.ExpectExec(regexp.QuoteMeta(query4)).WithArgs(user.Password, user.Login).WillReturnResult(result)
	mock.ExpectExec(regexp.QuoteMeta(query5)).WithArgs(user.AvatarURL, user.Login).WillReturnResult(result)
	mock.ExpectCommit()

	comm, err := urepo.UpdateUser(context.TODO(), &user)
	assert.NoError(t, err)
	assert.NotNil(t, comm)
}

func TestGetByLogin(t *testing.T) {
	log := wrapper.NewLogger()
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	urepo := NewUserRepository(db, log)

	rows := sqlxmock.NewRows([]string{"login", "name", "surname", "email", "password", "score", "avatarurl", "description"}).
		AddRow("1", "", "", "", "", 1, "", "")

	query := `SELECT Login, Name, Surname, Email, Password, Score, AvatarUrl, Description FROM author WHERE Login = $1`

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("oo").WillReturnRows(rows)

	comm, err := urepo.GetByLogin(context.TODO(), "oo")
	assert.NoError(t, err)
	assert.NotNil(t, comm)
}

func TestStore(t *testing.T) {
	log := wrapper.NewLogger()
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	urepo := NewUserRepository(db, log)

	user := umodels.User{
	}

	result := sqlxmock.NewResult(0, 0)

	// rows := sqlxmock.NewRows([]string{"login"}).
	// 	AddRow("1")

	query := "SELECT login FROM author WHERE login = $1"

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("")

	// rows1 := sqlxmock.NewRows([]string{""})

	query1 := `INSERT INTO author (Login, Name, Surname, Email, Password, Score, AvatarUrl, Description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	mock.ExpectExec(regexp.QuoteMeta(query1)).WithArgs("", "", "", "", "", 0, "", "").WillReturnResult(result)

	comm, err := urepo.Store(context.TODO(), &user)
	assert.NoError(t, err)
	assert.NotNil(t, comm)
}
