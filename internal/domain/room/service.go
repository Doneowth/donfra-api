package room

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"strings"
)

type Service struct {
	store    Store
	passcode string
	baseURL  string
}

func NewService(store Store, passcode, baseURL string) *Service {
	return &Service{store: store, passcode: passcode, baseURL: baseURL}
}

func (s *Service) Init(pass string) (string, error) {
	if strings.TrimSpace(pass) != s.passcode {
		return "", errors.New("invalid passcode")
	}
	b := make([]byte, 24)
	rand.Read(b)
	tok := base64.RawURLEncoding.EncodeToString(b)
	s.store.SetOpen(tok)
	path := "/coding?invite=" + tok + "&role=agent"
	if s.baseURL == "" {
		return path, nil
	}
	return strings.TrimRight(s.baseURL, "/") + path, nil
}

func (s *Service) IsOpen() bool           { return s.store.IsOpen() }
func (s *Service) Validate(t string) bool { return s.store.Validate(strings.TrimSpace(t)) }
func (s *Service) Close() error           { return s.store.Close() }
