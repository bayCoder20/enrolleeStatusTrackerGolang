package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

// Endpoint for the Enrollee service.
func makeCreateEnrolleeEndpoint(s EnrolleeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateEnrolleeRequest)
		msg, err := s.CreateEnrollee(ctx, req.enrollee)
		return CreateEnrolleeResponse{Msg: msg, Err: err}, nil
	}
}

func makeGetEnrolleeByIDEndpoint(s EnrolleeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetEnrolleeByIdRequest)
		enrolleeDetails, err := s.GetEnrolleeById(ctx, req.Id)
		if err != nil {
			return GetEnrolleeByIdResponse{
				Enrollee: enrolleeDetails, Err: "Id not found"}, nil
		}
		return GetEnrolleeByIdResponse{Enrollee: enrolleeDetails, Err: ""}, nil
	}
}

func makeGetAllEnrolleesEndpoint(s EnrolleeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		enrolleeDetails, err := s.GetAllEnrollees(ctx)
		if err != nil {
			return GetAllEnrolleesResponse{Enrollee: enrolleeDetails, Err: "no data found"}, nil
		}
		return GetAllEnrolleesResponse{Enrollee: enrolleeDetails, Err: ""}, nil
	}
}

func makeDeleteEnrolleeEndpoint(s EnrolleeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteEnrolleeRequest)
		msg, err := s.DeleteEnrollee(ctx, req.enrolleeId)
		if err != nil {
			return DeleteEnrolleeResponse{Msg: msg, Err: err}, nil
		}
		return DeleteEnrolleeResponse{Msg: msg, Err: nil}, nil
	}
}

func makeUpdateEnrolleeendpoint(s EnrolleeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateEnrolleeRequest)
		msg, err := s.UpdateEnrollee(ctx, req.enrollee)
		return msg, err
	}
}

func decodeCreateEnrolleeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateEnrolleeRequest
	fmt.Println("into Decoding")
	if err := json.NewDecoder(r.Body).Decode(&req.enrollee); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetEnrolleeByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetEnrolleeByIDRequest
	vars := mux.Vars(r)
	req = GetEnrolleeByIDRequest{ID: vars["enrolleeId"]}
	return req, nil
}

func decodeGetAllEnrolleesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetAllEnrolleesRequest
	return req, nil
}

func decodeDeleteEnrolleeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req DeleteEnrolleeRequest
	vars := mux.Vars(r)
	req = DeleteEnrolleeRequest{enrolleeID: vars["enrolleeId"]}
	return req, nil
}
func decodeUpdateEnrolleeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req UpdateEnrolleeRequest
	if err := json.NewDecoder(r.Body).Decode(&req.enrollee); err != nil {
		return nil, err
	}
	return req, nil
}
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type (
	// CreateEnrolleeRequest struct statement
	CreateEnrolleeRequest struct {
		enrollee Enrollee
	}
	// CreateEnrolleeResponse struct statement
	CreateEnrolleeResponse struct {
		Msg string `json:"msg"`
		Err error  `json:"error,omitempty"`
	}
	// GetEnrolleeByIDRequest struct statement
	GetEnrolleeByIDRequest struct {
		ID string `json:"enrolleeID"`
	}
	// GetEnrolleeByIDResponse struct statement
	GetEnrolleeByIDResponse struct {
		Enrollee interface{} `json:"enrollee,omitempty"`
		Err      string      `json:"error,omitempty"`
	}
	// GetAllEnrolleesRequest struct statement
	GetAllEnrolleesRequest struct {
		Enrollee interface{} `json:"enrollee,omitempty"`
		Err      string      `json:"error,omitempty"`
	}
	// GetAllEnrolleesResponse struct statement
	GetAllEnrolleesResponse struct {
		Enrollee interface{} `json:"enrollee,omitempty"`
		Err      string      `json:"error,omitempty"`
	}
	// DeleteEnrolleeRequest struct statement
	DeleteEnrolleeRequest struct {
		enrolleeID string `json:"enrolleeId"`
	}
	// DeleteEnrolleeResponse struct statement
	DeleteEnrolleeResponse struct {
		Msg string `json:"response"`
		Err error  `json:"error,omitempty"`
	}
	// UpdateEnrolleeRequest struct statement
	UpdateEnrolleeRequest struct {
		enrollee Enrollee
	}
	// UpdateEnrolleeResponse struct statement
	UpdateEnrolleeResponse struct {
		Msg string `json:"status,omitempty"`
		Err error  `json:"error,omitempty"`
	}
)
