package main

import (
	"context"
	"testing"

	intelligencev1 "OlympusGCP-Intelligence/gen/v1/intelligence"
	"OlympusGCP-Intelligence/10000-Autonomous-Actors/10700-Processing-Engines/10710-Reasoning-Inference/inference"
	"connectrpc.com/connect"
)

func TestIntelligenceServer(t *testing.T) {
	server := &inference.IntelligenceServer{}
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
