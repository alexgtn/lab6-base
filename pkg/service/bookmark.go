package service

import "github.com/alexgtn/esi2021-lab4/pkg/domain"

type bookRepository interface {
	Create(s *domain.Bookmark) (*domain.Bookmark, error)
	GetAll() ([]*domain.Bookmark, error)
}

type BookmarkService struct {
	bookRepository bookRepository
}

func NewBookmarkService(bR bookRepository) *BookmarkService {
	return &BookmarkService{
		bookRepository: bR,
	}
}

func (s *BookmarkService) Create(book *domain.Bookmark) (*domain.Bookmark, error) {
	return s.bookRepository.Create(book)
}

func (s *BookmarkService) GetAll() ([]*domain.Bookmark, error){
	return s.bookRepository.GetAll()
}