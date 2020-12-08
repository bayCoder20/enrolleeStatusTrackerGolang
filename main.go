package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	db := GetDBconn()
	r := mux.NewRouter()
	var svc EnrolleeService
	svc = enrolleeservice{}
	{
		repository, err := NewRepo(db, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		svc = NewService(repository, logger)
	}
	CreateEnrolleeHandler := httptransport.NewServer(
		makeCreateEnrolleeEndpoint(svc),
		decodeCreateEnrolleeRequest,
		encodeResponse,
	)
	GetByEnrolleeIdHandler := httptransport.NewServer(
		makeGetEnrolleeByIdEndpoint(svc),
		decodeGetEnrolleeByIdRequest,
		encodeResponse,
	)
	GetAllEnrolleesHandler := httptransport.NewServer(
		makeGetAllEnrolleesEndpoint(svc),
		decodeGetAllEnrolleesRequest,
		encodeResponse,
	)
	DeleteEnrolleeHandler := httptransport.NewServer(
		makeDeleteEnrolleeEndpoint(svc),
		decodeDeleteEnrolleeRequest,
		encodeResponse,
	)
	UpdateEnrolleeHandler := httptransport.NewServer(
		makeUpdateEnrolleeendpoint(svc),
		decodeUpdateEnrolleeRequest,
		encodeResponse,
	)
	http.Handle("/", r)
	http.Handle("/enrollee", CreateEnrolleeHandler)

	http.Handle("/enrollee/update", UpdateEnrolleeHandler)

	r.Handle("/enrollee/getAll", GetAllEnrolleesHandler).Methods("GET")
	r.Handle("/enrollee/{enrolleeid}", GetByEnrolleeIdHandler).Methods("GET")
	r.Handle("/enrollee/{enrolleeid}", DeleteEnrolleeHandler).Methods("DELETE")
	logger.Log("msg", "HTTP", "addr", ":8000")
	logger.Log("err", http.ListenAndServe(":8000", nil))
}
