package flow

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/BaritoLog/go-boilerplate/errkit"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/vwidjaya/barito-proto/producer"
	"google.golang.org/grpc"
)

type BaritoLokiService interface {
	pb.ProducerServer
	Start() error
	LaunchREST() error
	Close()
}

type baritoLokiService struct {
	grpcAddr string
	restAddr string
	lkClient Loki

	grpcServer   *grpc.Server
	reverseProxy *http.Server
}

func NewBaritoLokiService(params map[string]interface{}) BaritoLokiService {
	lkConfig := params["lkConfig"].(lokiConfig)

	return &baritoLokiService{
		grpcAddr: params["grpcAddr"].(string),
		restAddr: params["restAddr"].(string),
		lkClient: NewLoki(lkConfig),
	}
}

func (s *baritoLokiService) initGrpcServer() (lis net.Listener, srv *grpc.Server, err error) {
	lis, err = net.Listen("tcp", s.grpcAddr)
	if err != nil {
		return
	}

	srv = grpc.NewServer()
	pb.RegisterProducerServer(srv, s)

	s.grpcServer = srv
	return
}

func (s *baritoLokiService) Start() (err error) {
	lis, grpcSrv, err := s.initGrpcServer()
	if err != nil {
		err = errkit.Concat(ErrInitGrpc, err)
		return
	}

	return grpcSrv.Serve(lis)
}

func (s *baritoLokiService) LaunchREST() (err error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = pb.RegisterProducerHandlerFromEndpoint(ctx, mux, "localhost"+s.grpcAddr, opts)
	if err != nil {
		err = errkit.Concat(ErrRegisterGrpc, err)
		return
	}

	s.reverseProxy = &http.Server{
		Addr:    s.restAddr,
		Handler: mux,
	}

	err = s.reverseProxy.ListenAndServe()
	if err != nil {
		err = errkit.Concat(ErrReverseProxy, err)
	}
	return
}

func (s *baritoLokiService) Close() {
	if s.reverseProxy != nil {
		s.reverseProxy.Close()
	}

	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}
}

func (s *baritoLokiService) Produce(_ context.Context, timber *pb.Timber) (resp *pb.ProduceResult, err error) {
	labels := generateLokiLabels(timber.GetContext())

	timber.Timestamp = time.Now().UTC().Format(time.RFC3339)
	err = s.lkClient.Store(labels, *timber)
	if err != nil {
		err = onStoreErrorGrpc(err)
		return
	}

	resp = &pb.ProduceResult{
		Topic: labels,
	}
	return
}

func (s *baritoLokiService) ProduceBatch(_ context.Context, timberCollection *pb.TimberCollection) (resp *pb.ProduceResult, err error) {
	labels := generateLokiLabels(timberCollection.GetContext())

	for _, timber := range timberCollection.GetItems() {
		timber.Context = timberCollection.GetContext()
		timber.Timestamp = time.Now().UTC().Format(time.RFC3339)
		err = s.lkClient.Store(labels, *timber)
		if err != nil {
			err = onStoreErrorGrpc(err)
			return
		}
	}

	resp = &pb.ProduceResult{
		Topic: labels,
	}
	return
}

func generateLokiLabels(tCtx *pb.TimberContext) string {
	esIndexPrefix := tCtx.GetEsIndexPrefix()
	currDate := time.Now().Format("2006.01.02")
	return fmt.Sprintf("{app_name=\"%s-%s\"}", esIndexPrefix, currDate)
}
