package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"OlympusGCP-Intelligence/gen/v1/intelligence/intelligencev1connect"
	"OlympusGCP-Intelligence/10000-Autonomous-Actors/10700-Processing-Engines/10710-Reasoning-Inference/inference"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	server := &inference.IntelligenceServer{}
	mux := http.NewServeMux()
	path, handler := intelligencev1connect.NewIntelligenceServiceHandler(server)
	mux.Handle(path, handler)

	// Health Check / Pulse
	mux.HandleFunc("/pulse", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"HEALTHY", "workspace":"OlympusGCP-Intelligence", "time":"%s"}`, time.Now().Format(time.RFC3339))
	})

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
