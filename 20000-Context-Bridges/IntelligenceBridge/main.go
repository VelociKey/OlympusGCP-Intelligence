package main

import (
	"context"
	"net/http"

	"mcp-go/mcp"

	"connectrpc.com/connect"

	mcpbridge "Olympus2/90000-Enablement-Labs/P0000-pkg/000-mcp-bridge"
	intelligencev1 "OlympusGCP-Intelligence/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/intelligence/v1"
	"OlympusGCP-Intelligence/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/intelligence/v1/intelligencev1connect"
)

func main() {
	s := mcpbridge.NewBridgeServer("OlympusIntelligenceBridge", "1.0.0")

	client := intelligencev1connect.NewIntelligenceServiceClient(
		http.DefaultClient,
		"http://localhost:8096",
	)

	s.AddTool(mcp.NewTool("intelligence_predict",
		mcp.WithDescription("Request a prediction from a local model. Args: {model: string, prompt: string}"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		m, err := mcpbridge.ExtractMap(request)
		if err != nil {
			return mcpbridge.HandleError(err)
		}

		model, _ := m["model"].(string)
		prompt, _ := m["prompt"].(string)

		resp, err := client.Predict(ctx, connect.NewRequest(&intelligencev1.PredictRequest{
			Model:  model,
			Prompt: prompt,
		}))
		if err != nil {
			return mcpbridge.HandleError(err)
		}

		return mcp.NewToolResultText(resp.Msg.Prediction), nil
	})

	s.Run()
}
