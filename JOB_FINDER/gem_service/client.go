package gem_service

import (
	"context"

	"google.golang.org/genai"
)

// one method per Gemini endpoint, uses caller

func (g *GeminiService) UploadVideo(ctx context.Context, path string) (*genai.File, error) {
	//return io.Reader for video file, error if not found
	file, err := g.caller.OpenFile(path)
	if err != nil {
		//err should be logged in openfile
	}
	defer file.Close()

	return g.client.Files.Upload(ctx, file, &genai.UploadFileConfig{MIMEType: "video/mp4"})
}

func (g *GeminiService) CreateChat(ctx context.Context, model string) (*genai.Chat, error) {
	history := []*genai.Content{}
	chat, err := g.client.Chats.Create(ctx, model, nil, history)
	if err != nil {
		g.logger.Error("Couldn't create a new chat: ", err)
	}
	return chat, err
}

// SendMessageStream accepts chat(for history), propmt(new Q)
// return channel to stream response to
func (g *GeminiService) AskStream(ctx context.Context, chat *genai.Chat, text string) <-chan StreamResponse {
	ch := make(chan StreamResponse)

	go func() {
		defer close(ch)
		for chunk, err := range chat.SendMessageStream(ctx, genai.Part{Text: text}) {
			var resp StreamResponse
			if err != nil {
				resp = StreamResponse{Error: err}
			} else {
				resp = StreamResponse{Chunk: chunk.Text()}
			}

			select {
			case ch <- resp:
			case <-ctx.Done():
				return
			}
		}
	}()
	return ch
}
