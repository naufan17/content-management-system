package service

import (
	"errors"

	"github.com/google/uuid"

	"github.com/naufan17/content-management-system/internal/dto"
	"github.com/naufan17/content-management-system/internal/repository"
)

func GetNews() ([]dto.NewsDto, error) {
	news, err := repository.FindAllNews()

	if err != nil {
		return nil, err
	}

	var newsDtos []dto.NewsDto

	for _, n := range news {
		newsDtos = append(newsDtos, dto.NewsModelToDto(n))
	}

	return newsDtos, nil
}

func GetNewsByID(id uuid.UUID) (dto.NewsDto, error) {
	news, err := repository.FindByIDNews(id)

	if err != nil {
		return dto.NewsDto{}, err
	}

	return dto.NewsModelToDto(news), nil
}

func CreateNews(news dto.CreateNewsRequest) error {
	err := repository.CreateNews(dto.CreateNewsDtoToModel(news))

	if err != nil {
		return err
	}

	return nil
}

func UpdateNews(id uuid.UUID, news dto.UpdateNewsRequest) error {
	newsFromDb, err := repository.FindByIDNews(id)

	if err != nil {
		return errors.New("not found")
	}

	if newsFromDb.UserID != news.UserID {
		return errors.New("not authorized")
	}

	err = repository.UpdateNews(id, dto.UpdateNewsDtoToModel(news))

	if err != nil {
		return err
	}

	return nil
}

func DeleteNews(id uuid.UUID, userID uuid.UUID) error {
	newsFromDb, err := repository.FindByIDNews(id)

	if err != nil {
		return errors.New("not found")
	}

	if newsFromDb.UserID != userID {
		return errors.New("not authorized")
	}

	err = repository.DeleteNews(id)

	if err != nil {
		return err
	}

	return nil
}
