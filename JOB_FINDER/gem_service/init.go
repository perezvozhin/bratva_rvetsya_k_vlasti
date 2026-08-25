package gem_service

import (
	"JOB_FINDER/caller"
	"context"

	"go.uber.org/zap"
	"google.golang.org/genai"
)

type GeminiService struct {
	caller *caller.Caller
	client *genai.Client
	logger *zap.SugaredLogger
}

func Init(ctx context.Context, caller *caller.Caller, apiKey string) *GeminiService {
	cClient := caller.HTTPClient()
	cLogger := caller.Logger()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:     apiKey,
		Backend:    genai.BackendGeminiAPI,
		HTTPClient: cClient,
	})
	if err != nil {
		cLogger.Error("Failed to Initialize geminiService: ", err)
	}
	return &GeminiService{
		caller: caller,
		client: client,
		logger: cLogger,
	}
}
