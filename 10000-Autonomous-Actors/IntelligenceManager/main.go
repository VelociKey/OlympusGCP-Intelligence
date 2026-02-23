package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	intelligencev1 "OlympusGCP-Intelligence/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/intelligence/v1"
	"OlympusGCP-Intelligence/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/intelligence/v1/intelligencev1connect"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type IntelligenceServer struct{}

func (s *IntelligenceServer) Predict(ctx context.Context, req *connect.Request[intelligencev1.PredictRequest]) (*connect.Response[intelligencev1.PredictResponse], error) {
	slog.Info("Predict", "model", req.Msg.Model, "prompt", req.Msg.Prompt)

	prediction := fmt.Sprintf("[Sovereign Intelligence] Model: %s acknowledged prompt. Result: Success.", req.Msg.Model)
	return connect.NewResponse(&intelligencev1.PredictResponse{Prediction: prediction}), nil
}

func main() {
	server := &IntelligenceServer{}
	mux := http.NewServeMux()
	path, handler := intelligencev1connect.NewIntelligenceServiceHandler(server)
	mux.Handle(path, handler)

	port := "8096" // From genesis.json
	slog.Info("IntelligenceManager starting", "port", port)

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := srv.ListenAndServe()
	if err != nil {
		slog.Error("Server failed", "error", err)
	}
}
