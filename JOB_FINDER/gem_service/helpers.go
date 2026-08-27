package gem_service

import (
	"errors"
	"io/fs"
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
		for _, part := range c.Parts {
			m.Text += part.Text
			//TODO: add check if its not a text answer(mimetype)
		}
		msgs = append(msgs, m)
	}

	return msgs
}

func (g *GeminiService) loadHistory(chatName string) (ChatHistory, error) {
	var h ChatHistory
	err := g.caller.LoadJSON(historyFile(chatName), &h)
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
	return g.caller.SaveJSON(historyFile(s.name), h)
}

func historyFile(chatName string) string {
	return "chat_" + chatName + "_history.json"
}
