package gem_service

import (
	"errors"
	"io/fs"
	"path/filepath"
	"strings"
	"time"

	"google.golang.org/genai"
)

func toContents(messages []Message) []*genai.Content {
	var content []*genai.Content

	for _, m := range messages {
		c := &genai.Content{Role: m.Role, Parts: []*genai.Part{{Text: m.Text}}}
		content = append(content, c)
	}

	return content
}

func fromContents(content []*genai.Content) []Message {
	var msgs []Message

	for _, c := range content {
		m := Message{Role: c.Role}
		var sb strings.Builder
		for _, part := range c.Parts {
			sb.WriteString(part.Text)
			//TODO: add check if its not a text answer(mimetype)
		}
		m.Text = sb.String()
		msgs = append(msgs, m)
	}

	return msgs
}

func (g *GeminiService) loadHistory(chatName string) (ChatHistory, error) {
	var h ChatHistory
	err := g.caller.LoadJSON(g.historyFile(chatName), &h)
	if errors.Is(err, fs.ErrNotExist) {
		return ChatHistory{}, nil
	}
	if err != nil {
		g.logger.Error("couldn't load history for "+chatName+": ", err)
		return ChatHistory{}, err
	}
	return h, nil
}

func (g *GeminiService) saveHistory(s *session) error {
	h := ChatHistory{
		Model:     s.model,
		ChatName:  s.name,
		Messages:  fromContents(s.chat.History(true)),
		UpdatedAt: time.Now(),
	}
	return g.caller.SaveJSON(g.historyFile(s.name), h)
}

func (g *GeminiService) historyFile(chatName string) string {
	return filepath.Join(g.chatsDir, "chat_"+chatName+"_history.json")
}
