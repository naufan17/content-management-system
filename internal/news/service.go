package news

import (
	"errors"

	"github.com/google/uuid"
)

type NewsService interface {
	GetNews() ([]NewsDto, error)
	GetNewsByID(id uuid.UUID) (NewsDto, error)
	CreateNews(news CreateNewsRequest) error
	UpdateNews(id uuid.UUID, news UpdateNewsRequest) error
	DeleteNews(id uuid.UUID, userID uuid.UUID) error
}

type newsService struct {
	newsRepository NewsRepository
}

func NewNewsService(newsRepository NewsRepository) NewsService {
	return &newsService{
		newsRepository: newsRepository,
	}
}

func (s *newsService) GetNews() ([]NewsDto, error) {
	news, err := s.newsRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var newsDtos []NewsDto

	for _, n := range news {
		newsDtos = append(newsDtos, NewsModelToDto(n))
	}

	return newsDtos, nil
}

func (s *newsService) GetNewsByID(id uuid.UUID) (NewsDto, error) {
	news, err := s.newsRepository.FindByID(id)

	if err != nil {
		return NewsDto{}, err
	}

	return NewsModelToDto(news), nil
}

func (s *newsService) CreateNews(news CreateNewsRequest) error {
	err := s.newsRepository.Create(CreateNewsDtoToModel(news))

	if err != nil {
		return err
	}

	return nil
}

func (s *newsService) UpdateNews(id uuid.UUID, news UpdateNewsRequest) error {
	newsFromDb, err := s.newsRepository.FindByID(id)

	if err != nil {
		return errors.New("not found")
	}

	if newsFromDb.UserID != news.UserID {
		return errors.New("not authorized")
	}

	err = s.newsRepository.Update(id, UpdateNewsDtoToModel(news))

	if err != nil {
		return err
	}

	return nil
}

func (s *newsService) DeleteNews(id uuid.UUID, userID uuid.UUID) error {
	newsFromDb, err := s.newsRepository.FindByID(id)

	if err != nil {
		return errors.New("not found")
	}

	if newsFromDb.UserID != userID {
		return errors.New("not authorized")
	}

	err = s.newsRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
