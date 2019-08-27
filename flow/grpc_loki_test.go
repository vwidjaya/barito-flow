package flow

import (
	"fmt"
	"testing"

	. "github.com/BaritoLog/go-boilerplate/testkit"
	pb "github.com/vwidjaya/barito-proto/producer"
)

func NewTestLokiService() *baritoLokiService {
	lkConfig := lokiConfig{
		bulkSize: 1,
	}

	return &baritoLokiService{
		lkClient: NewLoki(lkConfig),
	}
}

func TestLokiService_Produce_OnStoreError(t *testing.T) {
	srv := NewTestLokiService()
	nilContentTimber := &pb.Timber{}

	_, err := srv.Produce(nil, nilContentTimber)
	FatalIfWrongGrpcError(t, onStoreErrorGrpc(fmt.Errorf("")), err)
}

func TestLokiService_Produce_OnSuccess(t *testing.T) {
	srv := NewTestLokiService()
	_, err := srv.Produce(nil, pb.SampleTimberProto())
	FatalIfError(t, err)
}

func TestLokiService_ProduceBatch_OnStoreError(t *testing.T) {
	srv := NewTestLokiService()
	nilContentTimberCol := &pb.TimberCollection{
		Items: []*pb.Timber{
			&pb.Timber{},
		},
	}

	_, err := srv.ProduceBatch(nil, nilContentTimberCol)
	FatalIfWrongGrpcError(t, onStoreErrorGrpc(fmt.Errorf("")), err)
}

func TestLokiService_ProduceBatch_OnSuccess(t *testing.T) {
	srv := NewTestLokiService()
	_, err := srv.ProduceBatch(nil, pb.SampleTimberCollectionProto())
	FatalIfError(t, err)
}

func TestLokiService_Start(t *testing.T) {
	service := &baritoLokiService{
		grpcAddr: ":24400",
	}

	var err error
	go func() {
		err = service.Start()
	}()
	defer service.Close()

	FatalIfError(t, err)
}
