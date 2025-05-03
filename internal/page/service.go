package page

import (
	"errors"

	"github.com/google/uuid"
)

type PageService interface {
	GetPages() ([]PageDto, error)
	GetPage(id uuid.UUID) (PageDto, error)
	CreatePage(page CreatePageRequest) error
	UpdatePage(id uuid.UUID, page UpdatePageRequest) error
	DeletePage(id uuid.UUID, userID uuid.UUID) error
}

type pageService struct {
	pageRepository PageRepository
}

func NewPageService(pageRepository PageRepository) PageService {
	return &pageService{pageRepository: pageRepository}
}

func (s *pageService) GetPages() ([]PageDto, error) {
	pages, err := s.pageRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var pagesDtos []PageDto

	for _, page := range pages {
		pagesDtos = append(pagesDtos, PageModelToDto(page))
	}

	return pagesDtos, nil
}

func (s *pageService) GetPage(id uuid.UUID) (PageDto, error) {
	page, err := s.pageRepository.FindByID(id)

	if err != nil {
		return PageDto{}, err
	}

	return PageModelToDto(page), nil
}

func (s *pageService) CreatePage(page CreatePageRequest) error {
	err := s.pageRepository.Create(CreatePageDtoToModel(page))

	if err != nil {
		return err
	}

	return nil
}

func (s *pageService) UpdatePage(id uuid.UUID, page UpdatePageRequest) error {
	pageFromDb, err := s.pageRepository.FindByID(id)

	if err != nil {
		return errors.New("not found")
	}

	if pageFromDb.UserID != page.UserID {
		return errors.New("not authorized")
	}

	err = s.pageRepository.Update(id, UpdatePageDtoToModel(page))

	if err != nil {
		return err
	}

	return nil
}

func (s *pageService) DeletePage(id uuid.UUID, userID uuid.UUID) error {
	pageFromDb, err := s.pageRepository.FindByID(id)

	if err != nil {
		return errors.New("not found")
	}

	if pageFromDb.UserID != userID {
		return errors.New("not authorized")
	}

	err = s.pageRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
