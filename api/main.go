package main

import (
	"context"
	"flag"
	"net/http"

	"git.neds.sh/matty/entain/api/proto/racing"
	"git.neds.sh/matty/entain/api/proto/sports"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	apiEndpoint           = flag.String("api-endpoint", "localhost:8000", "API endpoint")
	racingServiceEndpoint = flag.String("racingService-endpoint", "localhost:9000", "racing service endpoint")
	sportServiceEndpoint  = flag.String("sportService-endpoint", "localhost:9001", "sports service endpoint")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatalf("failed running api server: %s", err)
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	if racingErr := racing.RegisterRacingHandlerFromEndpoint(
		ctx,
		mux,
		*racingServiceEndpoint,
		[]grpc.DialOption{grpc.WithInsecure()},
	); racingErr != nil {
		return racingErr
	}

	if sportsErr := sports.RegisterSportsHandlerFromEndpoint(
		ctx,
		mux,
		*sportServiceEndpoint,
		[]grpc.DialOption{grpc.WithInsecure()},
	); sportsErr != nil {
		return sportsErr
	}

	log.Infof("API server listening on: %s", *apiEndpoint)

	return http.ListenAndServe(*apiEndpoint, mux)
}
