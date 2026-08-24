package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
)

type Store struct {
	dir    string
	mu     sync.RWMutex
	byID   map[string]Interview // TODO: when deletion - map keeps its size...use other data structure(not urgent)
	logger *zap.SugaredLogger
}

type Interview struct {
	UUID         string    `json:"id"`
	Title        string    `json:"title"` // "OZON, T-Bank etc."
	FileName     string    `json:"file"`
	CreatedAt    time.Time `json:"date"` // "2026-08-21"
	SizeBytes    int64     `json:"size_bytes"`
	GeminiChatID string    `json:"gemini_file_id,omitempty"`
	Status       string    `json:"status"` // "new", "in progress", "done"
}

// TODO: use Storage interface - dependency injection(testing)
func NewStore(dir string, logger *zap.SugaredLogger) (*Store, error) {
	// If no such directory - create with rw permissions
	if err := os.MkdirAll(dir, 0o755); err != nil {
		logger.Errorw("Error creating directory", "path", dir, "error", err)
		return nil, err
	}

	return &Store{
		dir:    dir,
		byID:   make(map[string]Interview),
		logger: logger,
	}, nil
}

func (s *Store) Show() {
	s.logger.Info(s.byID)
}

// Scan scans the directory for files and updates the byID map with the found interviews.
func (s *Store) Scan() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	files, err := os.ReadDir(s.dir)
	if err != nil {
		s.logger.Errorw("Error reading directory", "path", s.dir, "error", err)
		return err
	}

	for _, file := range files {
		extension := filepath.Ext(file.Name())
		metadata := strings.TrimSuffix(file.Name(), extension) + ".json"
		metadataPath := filepath.Join(s.dir, metadata)
		if extension == ".mp4" {
			_, err = os.Stat(metadataPath)
			if os.IsNotExist(err) {
				//create metadata file
				info, err := file.Info()
				if err != nil {
					return err
				}
				size := info.Size()
				path := filepath.Join(s.dir, file.Name())
				err = s.ingest(path, size)
				if err != nil {
					return err
				}
			}
		}
		if extension == ".json" {
			//call load metadata here
			err = s.load(metadataPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Store) ingest(filePath string, size int64) error {
	uuid := NewUUID()
	name := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	jsonPath := strings.TrimSuffix(filePath, filepath.Ext(filePath)) + ".json"

	interview := Interview{
		UUID:      uuid,
		Title:     name,
		FileName:  filepath.Base(filePath),
		CreatedAt: time.Now(),
		SizeBytes: size,
		Status:    "new",
	}

	data, err := json.MarshalIndent(interview, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(jsonPath, data, 0644)
	if err != nil {
		return err
	}

	//lock is used in parent call of Scan
	s.byID[uuid] = interview

	return nil
}

func (s *Store) load(file string) error {
	var interview Interview

	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &interview)
	if err != nil {
		return err
	}

	//lock in Scan
	s.byID[interview.UUID] = interview

	return err
}
