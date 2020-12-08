package main

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

//Enrollee struct
type Enrollee struct {
	enrolleeID    string `json:"enrolleeId"`
	lastName      string `json:"lastName"`
	firstName     string `json:"firstName"`
	middleInitial string `json:"middleInitial"`
	birthDate     string `json:"birthDate"`
	phoneNumber   string `json:"phoneNumber"`
	sex           string `json:"sex"`
	activeStatus  bool   `json:"activeStatus"`
}

// EnrolleeService interface
type EnrolleeService interface {
	CreateEnrollee(ctx context.Context, enrollee Enrollee) (string, error)
	GetEnrolleeById(ctx context.Context, id string) (interface{}, error)
	GetAllEnrollees(ctx context.Context) (interface{}, error)
	UpdateEnrollee(ctx context.Context, enrollee Enrollee) (string, error)
	DeleteEnrollee(ctx context.Context, id string) (string, error)
}

// service implements the Enrollee Service
type enrolleeservice struct {
	repository Repository
	logger     log.Logger
}

//Repository describes the enrollee service for repository interaction
type Repository interface {
	CreateEnrollee(ctx context.Context, enrollee Enrollee) error
	GetEnrolleeById(ctx context.Context, id string) (interface{}, error)
	GetAllEnrollees(ctx context.Context) (interface{}, error)
	UpdateEnrollee(ctx context.Context, enrollee Enrollee) (string, error)
	DeleteEnrollee(ctx context.Context, id string) (string, error)
}

// NewService creates and returns a new Enrollee service instance
func NewService(rep Repository, logger log.Logger) EnrolleeService {
	return &enrolleeservice{
		repository: rep,
		logger:     logger,
	}
}
func (s enrolleeservice) CreateEnrollee(ctx context.Context, enrollee Enrollee) (string, error) {
	logger := log.With(s.logger, "method", "CreateEnrollee")
	var msg = "success"
	enrolleeDetails := Enrollee{
		enrolleeID:    enrollee.enrolleeID,
		lastName:      enrollee.lastName,
		firstName:     enrollee.firstName,
		middleInitial: enrollee.middleInitial,
		birthDate:     enrollee.birthDate,
		phoneNumber:   enrollee.phoneNumber,
		sex:           enrollee.sex,
		activeStatus:  enrollee.activeStatus,
	}
	if err := s.repository.CreateEnrollee(ctx, enrolleeDetails); err != nil {
		level.Error(logger).Log("err from repo is ", err)
		return "", err
	}
	return msg, nil
}

func (s enrolleeservice) GetEnrolleeByID(ctx context.Context, id string) (interface{}, error) {
	logger := log.With(s.logger, "method", "GetEnrolleeById")
	var enrollee interface{}
	var empty interface{}
	enrollee, err := s.repository.GetEnrolleeById(ctx, id)
	if err != nil {
		level.Error(logger).Log("err ", err)
		return empty, err
	}
	return enrollee, nil
}
func (s enrolleeservice) GetAllEnrollees(ctx context.Context) (interface{}, error) {
	logger := log.With(s.logger, "method", "GetAllEnrollees")
	var enrollee interface{}
	var empty interface{}
	enrollee, err := s.repository.GetAllEnrollees(ctx)
	if err != nil {
		level.Error(logger).Log("err ", err)
		return empty, err
	}
	return enrollee, nil
}
func (s enrolleeservice) DeleteEnrollee(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "DeleteEnrollee")
	msg, err := s.repository.DeleteEnrollee(ctx, id)
	if err != nil {
		level.Error(logger).Log("err ", err)
		return "", err
	}
	return msg, nil
}
func (s enrolleeservice) UpdateEnrollee(ctx context.Context, enrollee Enrollee) (string, error) {
	logger := log.With(s.logger, "method", "CreateEnrollee")
	var msg = "success"
	enrolleeDetails := Enrollee{
		enrolleeID:    enrollee.enrolleeID,
		lastName:      enrollee.lastName,
		firstName:     enrollee.firstName,
		middleInitial: enrollee.middleInitial,
		birthDate:     enrollee.birthDate,
		phoneNumber:   enrollee.phoneNumber,
		sex:           enrollee.sex,
		activeStatus:  enrollee.activeStatus,
	}
	msg, err := s.repository.UpdateEnrollee(ctx, enrolleeDetails)
	if err != nil {
		level.Error(logger).Log("err from repo is ", err)
		return "", err
	}
	return msg, nil
}
