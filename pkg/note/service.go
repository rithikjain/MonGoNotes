package note

import "github.com/rithikjain/MongoNotes/pkg/entities"

type Service interface {
	CreateNote(note *entities.Note) (*entities.Note, error)

	GetAllNotes() (*[]entities.Note, error)

	UpdateNote(note *entities.Note) error

	DeleteNote(id string) error
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

func (s *service) GetAllNotes() (*[]entities.Note, error) {
	return s.repo.GetAllNotes()
}

func (s *service) UpdateNote(note *entities.Note) error {
	return s.repo.UpdateNote(note)
}

func (s *service) DeleteNote(id string) error {
	return s.repo.DeleteNote(id)
}
