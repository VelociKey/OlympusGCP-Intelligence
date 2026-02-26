package main

import (
	"context"
	"dagger/olympusgcp-intelligence/internal/dagger"
)

type OlympusGCPIntelligence struct{}

func (m *OlympusGCPIntelligence) HelloWorld(ctx context.Context) string {
	return "Hello from OlympusGCP-Intelligence!"
}

func main() {
	dagger.Serve()
}
