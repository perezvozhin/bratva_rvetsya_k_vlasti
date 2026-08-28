package gem_service

import (
	"JOB_FINDER/internals/domain"
	"context"

	"google.golang.org/genai"
)

func (g *GeminiService) UploadVideo(ctx context.Context, path string) (UploadFileStatus, error) {
	//return *os.File(to read and close) for video file, error if not found
	file, err := g.caller.OpenFile(path)
	if err != nil {
		return UploadFileStatus{}, err
	}
	defer file.Close()

	genaiFile, err := g.client.Files.Upload(ctx, file, &genai.UploadFileConfig{MIMEType: "video/mp4"})
	if err != nil {
		return UploadFileStatus{}, err
	}

	res := UploadFileStatus{
		Name:           genaiFile.Name,
		MIMEType:       genaiFile.MIMEType,
		SizeBytes:      genaiFile.SizeBytes,
		CreateTime:     genaiFile.CreateTime,
		ExpirationTime: genaiFile.ExpirationTime,
		State:          FileState(genaiFile.State),
		Message:        "status: success",
		FileURI:        genaiFile.URI,
	}
	//TODO: handle err message and happy path one
	//Message:        genaiFile.Error.Message, not nil on err

	return res, nil
}

// loads a chat session(if active - from map, or last saved from json, if json empty - empty history)
func (g *GeminiService) session(ctx context.Context, chatName string) (*session, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if s, ok := g.chats[chatName]; ok {
		return s, nil
	}

	history, err := g.loadHistory(chatName)
	if err != nil {
		return nil, err
	}

	model := history.Model
	if model == "" {
		model = g.defaultModel
	}

	chat, err := g.client.Chats.Create(ctx, model, nil, toContents(history.Messages))
	if err != nil {
		g.logger.Error("Couldn't create chat "+chatName+": ", err)
		return nil, err
	}

	s := &session{chat: chat, name: chatName, model: model}
	g.chats[chatName] = s
	return s, nil
}

func (g *GeminiService) Ask(ctx context.Context, chatName string, req domain.Request) (<-chan StreamResponse, error) {
	s, err := g.session(ctx, chatName)
	if err != nil {
		return nil, err
	}

	ch := make(chan StreamResponse)

	go func() {
		defer close(ch)

		s.mu.Lock()
		defer s.mu.Unlock()

		ok := true
		parts := []genai.Part{{Text: req.Text}}
		if req.FileURI != "" {
			parts = append([]genai.Part{{FileData: &genai.FileData{
				FileURI: req.FileURI, MIMEType: req.MIMEType,
			}}}, parts...)
		}
		for chunk, err := range s.chat.SendMessageStream(ctx, parts...) {
			var resp StreamResponse
			if err != nil {
				ok = false
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

		if !ok {
			return
		}
		// TODO: manual injection after each successful prompt...
		// is it possible to automate this(interactions API)
		msgs := fromContents(s.chat.History(false)) // comprehensive, not curated
		s.chat, _ = g.client.Chats.Create(ctx, s.model, nil, toContents(msgs))

		if err := g.saveHistory(s); err != nil {
			g.logger.Error("couldn't save history for "+s.name+": ", err)
		}
		// g.logger.Infow("history after turn",
		//         "chat", s.name,
		//         "curated", len(s.chat.History(true)),
		//         "comprehensive", len(s.chat.History(false)))
	}()

	return ch, nil
}
