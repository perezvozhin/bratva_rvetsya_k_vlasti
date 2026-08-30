package gem_service

import (
	"errors"
	"io/fs"
	"path/filepath"
	"strings"
	"time"
	"unicode"

	"google.golang.org/genai"
)

const maxSlugLen = 60

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

		//concatenate model response(if streamed in chunks)
		//continue skips appending it to msgs until whole msg is sent
		if n := len(msgs); n > 0 && c.Role == "model" && msgs[n-1].Role == "model" {
			msgs[n-1].Text += m.Text
			continue
		}

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
	return filepath.Join(g.chatsDir, "chat_"+slugify(chatName)+"_history.json")
}

// slugify turns a user-typed chat name into a stable, filesystem-safe key.
func slugify(s string) string {
	var b strings.Builder
	dash := false

	for _, r := range strings.ToLower(strings.TrimSpace(s)) {
		switch {
		case unicode.IsLetter(r), unicode.IsDigit(r):
			b.WriteRune(r)
			dash = false
		default:
			if !dash && b.Len() > 0 {
				b.WriteByte('-')
				dash = true
			}
		}
	}

	out := strings.Trim(b.String(), "-")

	if r := []rune(out); len(r) > maxSlugLen {
		out = strings.Trim(string(r[:maxSlugLen]), "-")
	}

	return out
}
