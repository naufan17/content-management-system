package service

import (
	"errors"

	"github.com/google/uuid"

	"github.com/naufan17/content-management-system/internal/dto"
	"github.com/naufan17/content-management-system/internal/repository"
)

func GetPages() ([]dto.PageDto, error) {
	pages, err := repository.FindAllPage()

	if err != nil {
		return nil, err
	}

	var pagesDtos []dto.PageDto

	for _, page := range pages {
		pagesDtos = append(pagesDtos, dto.PageModelToDto(page))
	}

	return pagesDtos, nil
}

func GetPage(id uuid.UUID) (dto.PageDto, error) {
	page, err := repository.FindByIDPage(id)

	if err != nil {
		return dto.PageDto{}, err
	}

	return dto.PageModelToDto(page), nil
}

func CreatePage(page dto.CreatePageRequest) error {
	err := repository.CreatePage(dto.CreatePageDtoToModel(page))

	if err != nil {
		return err
	}

	return nil
}

func UpdatePage(id uuid.UUID, page dto.UpdatePageRequest) error {
	pageFromDb, err := repository.FindByIDPage(id)

	if err != nil {
		return errors.New("not found")
	}

	if pageFromDb.UserID != page.UserID {
		return errors.New("not authorized")
	}

	err = repository.UpdatePage(id, dto.UpdatePageDtoToModel(page))

	if err != nil {
		return err
	}

	return nil
}

func DeletePage(id uuid.UUID, userID uuid.UUID) error {
	pageFromDb, err := repository.FindByIDPage(id)

	if err != nil {
		return errors.New("not found")
	}

	if pageFromDb.UserID != userID {
		return errors.New("not authorized")
	}

	err = repository.DeletePage(id)

	if err != nil {
		return err
	}

	return nil
}
