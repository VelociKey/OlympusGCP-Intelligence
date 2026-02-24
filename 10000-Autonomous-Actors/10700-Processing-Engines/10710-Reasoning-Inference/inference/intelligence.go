package inference

import (
	"context"
	"fmt"
	"log/slog"

	intelligencev1 "OlympusGCP-Intelligence/gen/v1/intelligence"
	"connectrpc.com/connect"
)

type IntelligenceServer struct{}

func (s *IntelligenceServer) Predict(ctx context.Context, req *connect.Request[intelligencev1.PredictRequest]) (*connect.Response[intelligencev1.PredictResponse], error) {
	slog.Info("Predict", "model", req.Msg.Model, "prompt", req.Msg.Prompt)

	prediction := fmt.Sprintf("[Sovereign Intelligence] Model: %s acknowledged prompt. Result: Success.", req.Msg.Model)
	return connect.NewResponse(&intelligencev1.PredictResponse{Prediction: prediction}), nil
}
