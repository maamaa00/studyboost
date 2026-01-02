package usecase

import (
    "time"

    "studyboost/internal/domain"
)

type CreateSessionInput struct {
    Title     string
    StartedAt time.Time
    EndedAt   time.Time
}

type CreateSessionOutput struct {
    Session domain.Session
}

type SessionRepository interface {
    Save(session domain.Session) error
}

type CreateSessionUsecase struct {
    repo SessionRepository
}

func NewCreateSessionUsecase(repo SessionRepository) *CreateSessionUsecase {
    return &CreateSessionUsecase{repo: repo}
}

func (uc *CreateSessionUsecase) Execute(input CreateSessionInput) (CreateSessionOutput, error) {
    session := domain.Session{
        Title:     input.Title,
        StartedAt: input.StartedAt,
        EndedAt:   input.EndedAt,
        CreatedAt: time.Now(),
    }

    if err := uc.repo.Save(session); err != nil {
        return CreateSessionOutput{}, err
    }

    return CreateSessionOutput{Session: session}, nil
}