package flow

import (
	"github.com/BaritoLog/go-boilerplate/errkit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	ErrMakeSyncProducer       = errkit.Error("Make sync producer failed")
	ErrKafkaRetryLimitReached = errkit.Error("Error connecting to kafka, retry limit reached")
	ErrInitGrpc               = errkit.Error("Failed to listen to gRPC address")
	ErrRegisterGrpc           = errkit.Error("Error registering gRPC server endpoint into reverse proxy")
	ErrReverseProxy           = errkit.Error("Error serving REST reverse proxy")
)

func onLimitExceededGrpc() error {
	return status.Errorf(codes.ResourceExhausted, "Bandwidth Limit Exceeded")
}

func onBadRequestGrpc(err error) error {
	return status.Errorf(codes.InvalidArgument, err.Error())
}

func onStoreErrorGrpc(err error) error {
	return status.Errorf(codes.FailedPrecondition, err.Error())
}

func onCreateTopicErrorGrpc(err error) error {
	return status.Errorf(codes.Unavailable, err.Error())
}

func onSendCreateTopicErrorGrpc(err error) error {
	return status.Errorf(codes.Unavailable, err.Error())
}
