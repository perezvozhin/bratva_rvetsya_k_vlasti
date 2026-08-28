package gem_service

import (
	"JOB_FINDER/caller"
	"context"
	"sync"

	"go.uber.org/zap"
	"google.golang.org/genai"
)

type session struct {
	chat  *genai.Chat
	mu    sync.Mutex
	model string
	name  string
}

type GeminiService struct {
	caller *caller.Caller
	client *genai.Client
	logger *zap.SugaredLogger
	mu     sync.Mutex
	chats  map[string]*session
}

func Init(ctx context.Context, caller *caller.Caller, apiKey string) *GeminiService {
	cClient := caller.Client()
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
		chats:  make(map[string]*session),
	}
}
