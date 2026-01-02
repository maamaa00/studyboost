package repository

import "studyboost/internal/domain"

type SessionMemoryRepository struct {
    sessions []domain.Session
}

func NewSessionMemoryRepository() *SessionMemoryRepository {
    return &SessionMemoryRepository{
        sessions: []domain.Session{},
    }
}

func (r *SessionMemoryRepository) Save(session domain.Session) error {
    r.sessions = append(r.sessions, session)
    return nil
}