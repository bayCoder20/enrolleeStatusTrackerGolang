package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
)

var (
	RepoErr                   = errors.New("Unable to handle Repo Request")
	ErrIdNotFound             = errors.New("Id not found")
	ErrLastNameNotFound       = errors.New("Last name is not found")
	ErrFirstNameNotFound      = errors.New("First name is not found")
	ErrMiddileInitialNotFound = errors.New("Middle initial is not found")
	ErrBirthDateNotFound      = errors.New("Birth date is not found")
	ErrSexNotFound            = errors.New("Sex is not found")
	ErrActiveStatusNotFound   = errors.New("Active status is not found")
)

type repo struct {
	db     *sql.DB
	logger log.Logger
}

//Creates and returns an instance
func NewRepo(db *sql.DB, logger log.Logger) (Repository, error) {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "mongodb"),
	}, nil
}
func (repo *repo) CreateEnrollee(ctx context.Context, enrollee Enrollee) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO tbl_enrolleego (enrolleeId, lastName, firstName, middleInitial, birthDate, phoneNumber, sex, activeStatus) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", enrollee.enrolleeId, enrollee.lastName, enrollee.firstName, enrollee.middleInitial, enrollee.birthDate, enrollee.phoneNumber, enrollee.sex, enrollee.activeStatus)
	if err != nil {
		fmt.Println("Error occured inside CreateEnrollee in repo")
		return err
	} else {
		fmt.Println("User Created:", enrollee.lastName)
	}
	return nil
}
func (repo *repo) GetEnrolleeById(ctx context.Context, id string) (interface{}, error) {
	enrollee := Enrollee{}
	err := repo.db.QueryRowContext(ctx, "SELECT	c.enrolleeid,c.lastname,c.firstname, c.middleinitial, c.birthdate, c.phonenumber, c.sex, c.activestatus FROM tbl_enrolleego as c where c.enrolleeid = ?", id).Scan(
		&enrollee.enrolleeId,
		&enrollee.lastName, &enrollee.firstName, &enrollee.middleInitial, &enrollee.birthDate, &enrollee.phoneNumber, &enrollee.sex, &enrollee.activeStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			return enrollee, ErrIdNotFound
		}
		return enrollee, err
	}
	return enrollee, nil
}
func (repo *repo) GetAllEnrollees(ctx context.Context) (interface{}, error) {
	enrollee := Enrollee{}
	var res []interface{}
	rows, err := repo.db.QueryContext(ctx, "SELECT c.enrolleeid,c.lastname,c.firstname, c.middleinitial, c.birthdate, c.phonenumber, c.sex, c.activestatus FROM tbl_enrolleego as c ")
	if err != nil {
		if err == sql.ErrNoRows {
			return enrollee, ErrIdNotFound
		}
		return enrollee, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&enrollee.enrolleeId,
			&enrollee.lastName, &enrollee.firstName, &enrollee.middleInitial, &enrollee.birthDate, &enrollee.phoneNumber, &enrollee.sex, &enrollee.activeStatus)
		res = append([]interface{}{enrollee}, res...)
	}
	return res, nil
}
func (repo *repo) DeleteEnrollee(ctx context.Context, id string) (string, error) {
	res, err := repo.db.ExecContext(ctx, "DELETE FROM tbl_enrolleego WHERE enrolleeId = ? ", id)
	if err != nil {
		return "", err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return "", err
	} else if rowCnt == 0 {
		return "", ErrIdNotFound
	}
	return "Enrollee successfully deleted", nil
}
func (repo *repo) UpdateEnrollee(ctx context.Context, enrollee Enrollee) (string, error) {
	res, err := repo.db.ExecContext(ctx, "UPDATE tbl_enrolleego as c SET c.lastname = ?,c.firstname = ?, c.middleinitial = ?, c.birthdate = ?, c.phonenumber = ?, c.sex = ?, c.activestatus = ?", enrollee.lastName, enrollee.firstName, enrollee.middleInitial, enrollee.birthDate, enrollee.phoneNumber, enrollee.sex, enrollee.activeStatus)
	if err != nil {
		return "", err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return "", err
	}
	if rowCnt == 0 {
		return "", ErrIdNotFound
	}
	return "Enrollee successfully updated", err
}
