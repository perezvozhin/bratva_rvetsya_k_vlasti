package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"JOB_FINDER/gem_service"
	"JOB_FINDER/storage"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type ChatHandler struct {
	gemService *gem_service.GeminiService
	store      *storage.Store
	logger     *zap.SugaredLogger
}

func NewChatHandler(g *gem_service.GeminiService, s *storage.Store, l *zap.SugaredLogger) *ChatHandler {
	return &ChatHandler{
		gemService: g,
		store:      s,
		logger:     l,
	}
}

func (h *ChatHandler) RegisterRoutes(r chi.Router) {
	r.Get("/api/chats", h.getChats)
	r.Post("/api/chats", h.postChats)
	r.Get("/api/chats/{id}/messages", h.getMessages)
	r.Delete("/api/chats/{id}", h.deleteChat)
	r.Post("/api/chats/{chatId}/interview", h.postInterview)
	r.Post("/api/chats/{chatId}/messages", h.postMessage)
}

// GET /api/chats
func (h *ChatHandler) getChats(w http.ResponseWriter, r *http.Request) {
	// TODO: Убрать заглушку. Реализовать чтение JSON-файлов историй
	// из папки chats/ и возврат реального массива созданных чатов.
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("[]"))
}

// POST /api/chats
func (h *ChatHandler) postChats(w http.ResponseWriter, r *http.Request) {
	// TODO: Убрать заглушку. Реализовать создание нового чата:
	// генерация ID, создание пустого JSON-файла истории и возврат объекта чата клиенту.
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "ok"}`))
}

// GET /api/chats/{id}/messages
func (h *ChatHandler) getMessages(w http.ResponseWriter, r *http.Request) {
	// TODO: Убрать заглушку. Извлечь {id} из URL, прочитать 
	// соответствующий файл из папки chats/ и вернуть массив истории сообщений.
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("[]"))
}

// DELETE /api/chats/{id}
func (h *ChatHandler) deleteChat(w http.ResponseWriter, r *http.Request) {
	// TODO: Убрать заглушку. Реализовать удаление 
	// JSON-файла истории чата из папки chats/ по {id}.
	w.WriteHeader(http.StatusOK)
}

// POST /api/chats/{chatId}/interview
func (h *ChatHandler) postInterview(w http.ResponseWriter, r *http.Request) {
	_ = chi.URLParam(r, "chatId")

	// Лимит 500 МБ
	err := r.ParseMultipartForm(500 << 20)
	if err != nil {
		h.logger.Errorw("ParseMultipartForm failed", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		h.logger.Errorw("FormFile failed", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// TODO 
	// 1. Убрать os.CreateTemp. Видео нужно сохранять напрямую в папку interviews/ 
	//    через методы пакета storage (чтобы создался metadata.json).
	// 2. Полученный путь передать в UploadVideo.
	// 3. После успешного UploadVideo вызвать h.store.UpdateGeminiID(UUID_видео, status.Name),
	//    чтобы привязать файл на диске к ID в облаке Гугла.

	tempFile, err := os.CreateTemp("", "upload-*.mp4")
	if err != nil {
		h.logger.Errorw("CreateTemp failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tempFilePath := tempFile.Name()
	defer os.Remove(tempFilePath)

	_, err = io.Copy(tempFile, file)
	if err != nil {
		h.logger.Errorw("Copy failed", "error", err)
		tempFile.Close()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tempFile.Close()

	status, err := h.gemService.UploadVideo(r.Context(), tempFilePath)
	if err != nil {
		h.logger.Errorw("UploadVideo failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		h.logger.Errorw("Encode status failed", "error", err)
	}
}

// POST /api/chats/{chatId}/messages
func (h *ChatHandler) postMessage(w http.ResponseWriter, r *http.Request) {
	chatID := chi.URLParam(r, "chatId")

	var req struct {
		Text string `json:"text"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Errorw("Decode JSON failed", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		h.logger.Error("Streaming unsupported")
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ch, err := h.gemService.Ask(r.Context(), chatID, req.Text)
	if err != nil {
		h.logger.Errorw("Ask failed", "error", err)
		return
	}

	for streamResp := range ch {
		if streamResp.Error != nil {
			h.logger.Errorw("Stream response error", "error", streamResp.Error)
			break
		}

		chunkBytes, err := json.Marshal(map[string]string{"text": streamResp.Chunk})
		if err != nil {
			h.logger.Errorw("Marshal chunk failed", "error", err)
			continue
		}

		fmt.Fprintf(w, "data: %s\n\n", string(chunkBytes))
		flusher.Flush()
	}

	fmt.Fprintf(w, "data: [DONE]\n\n")
	flusher.Flush()
}