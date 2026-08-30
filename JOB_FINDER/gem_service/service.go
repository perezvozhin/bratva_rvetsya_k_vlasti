package gem_service

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var (
	ErrInvalidChatName = errors.New("invalid chat name")
	ErrChatExists      = errors.New("chat already exists")
)

func (g *GeminiService) ListChats() ([]ChatSummary, error) {
	entries, err := os.ReadDir(g.chatsDir)
	if err != nil {
		g.logger.Error("couldn't read chats dir: ", err)
		return nil, err
	}

	out := []ChatSummary{}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		name := e.Name()
		if !strings.HasPrefix(name, "chat_") || !strings.HasSuffix(name, "_history.json") {
			continue
		}

		var h ChatHistory
		if err := g.caller.LoadJSON(filepath.Join(g.chatsDir, name), &h); err != nil {
			g.logger.Warn("skipping unreadable history "+name+": ", err)
			continue
		}

		// fallback for files written before ChatName was persisted
		if h.ChatName == "" {
			h.ChatName = strings.TrimSuffix(strings.TrimPrefix(name, "chat_"), "_history.json")
		}

		out = append(out, ChatSummary{ChatName: h.ChatName, UpdatedAt: h.UpdatedAt})
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].UpdatedAt.After(out[j].UpdatedAt)
	})

	return out, nil
}

func (g *GeminiService) Messages(chatName string) ([]Message, error) {
	if slugify(chatName) == "" {
		return nil, ErrInvalidChatName
	}

	h, err := g.loadHistory(chatName)
	if err != nil {
		return nil, err
	}
	if h.Messages == nil {
		return []Message{}, nil
	}
	return h.Messages, nil
}

// CreateChat writes an empty history file so the chat shows up in the
// sidebar before its first message.
func (g *GeminiService) CreateChat(chatName string) (ChatSummary, error) {
	if slugify(chatName) == "" {
		return ChatSummary{}, ErrInvalidChatName
	}

	path := g.historyFile(chatName)

	_, err := os.Stat(path)
	switch {
	case err == nil:
		return ChatSummary{}, ErrChatExists
	case !errors.Is(err, fs.ErrNotExist):
		return ChatSummary{}, err
	}

	h := ChatHistory{
		ChatName:  strings.TrimSpace(chatName),
		Model:     g.defaultModel,
		Messages:  []Message{},
		UpdatedAt: time.Now(),
	}

	if err := g.caller.SaveJSON(path, h); err != nil {
		g.logger.Error("couldn't create chat "+chatName+": ", err)
		return ChatSummary{}, err
	}

	return ChatSummary{ChatName: h.ChatName, UpdatedAt: h.UpdatedAt}, nil
}
