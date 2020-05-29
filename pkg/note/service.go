package note

import "github.com/rithikjain/MongoNotes/pkg/entities"

type Service interface {
	CreateNote(note *entities.Note) (*entities.Note, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) CreateNote(note *entities.Note) (*entities.Note, error) {
	return s.repo.CreateNote(note)
}
