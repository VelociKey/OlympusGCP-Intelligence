package main

import (
	"context"
	"testing"

	intelligencev1 "OlympusGCP-Intelligence/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/intelligence/v1"
	"connectrpc.com/connect"
)

func TestIntelligenceServer(t *testing.T) {
	server := &IntelligenceServer{}
	ctx := context.Background()

	// Test Predict
	req := connect.NewRequest(&intelligencev1.PredictRequest{
		Model:  "gemini-pro",
		Prompt: "What is the meaning of life?",
	})
	res, err := server.Predict(ctx, req)
	if err != nil {
		t.Fatalf("Predict failed: %v", err)
	}
	if res.Msg.Prediction == "" {
		t.Error("Expected prediction, got empty string")
	}
}
