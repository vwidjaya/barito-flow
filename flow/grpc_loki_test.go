package flow

import (
	. "github.com/BaritoLog/go-boilerplate/testkit"
	"testing"
)

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
