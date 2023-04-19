package services

import (
	"Project-Akhir/helpers"
	"Project-Akhir/models"
	"Project-Akhir/repositori"
)

type PhotoService struct {
	PhotoRepository repositori.PhotoRepository
}

func NewPhotoService(photoRepository repositori.PhotoRepository) *PhotoService {
	return &PhotoService{
		PhotoRepository: photoRepository,
	}
}

func (ps *PhotoService) Add(request models.PhotoCreateRequest, userID string) (models.PhotoCreateResponse, error) {
	id := helpers.GenerateID()

	photo := models.Photo{
		ID:       id,
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoURL: request.PhotoURL,
		UserID:   userID,
	}

	response, err := ps.PhotoRepository.Save(photo)

	if err != nil {
		if err != models.ErrorNotFound {
			return models.PhotoCreateResponse{}, models.ErrorNotFound
		}
		return models.PhotoCreateResponse{}, err
	}

	return models.PhotoCreateResponse{
		ID:        response.ID,
		UserID:    response.UserID,
		Title:     response.Title,
		Caption:   response.Caption,
		PhotoURL:  response.PhotoURL,
		CreatedAt: response.CreatedAt,
	}, nil
}

func (ps *PhotoService) GetAll() ([]models.PhotoResponse, error) {
	response := make([]models.PhotoResponse, 0)

	res, err := ps.PhotoRepository.Get()

	if err != nil {
		return []models.PhotoResponse{}, err
	}

	for _, ok := range res {
		response = append(response, models.PhotoResponse{
			ID:        ok.ID,
			UserID:    ok.UserID,
			Title:     ok.Title,
			Caption:   ok.Caption,
			PhotoURL:  ok.PhotoURL,
			CreatedAt: ok.CreatedAt,
			UpdatedAt: ok.UpdatedAt,
		})
	}

	return response, nil
}

func (ps *PhotoService) GetById(id string) (models.PhotoResponse, error) {
	photo, err := ps.PhotoRepository.GetById(id)

	if err != nil {
		if err != models.ErrorNotFound {
			return models.PhotoResponse{}, err
		}
		return models.PhotoResponse{}, models.ErrorNotFound
	}

	return models.PhotoResponse{
		ID:        photo.ID,
		UserID:    photo.UserID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		CreatedAt: photo.CreatedAt,
		UpdatedAt: photo.UpdatedAt,
	}, nil
}

func (ps *PhotoService) UpdateById(request models.PhotoUpdateRequest, id string, userId string) (models.PhotoUpdateResponse, error) {
	byID, err := ps.PhotoRepository.GetById(id)

	if err != nil {
		if err != models.ErrorNotFound {
			return models.PhotoUpdateResponse{}, err
		}
		return models.PhotoUpdateResponse{}, models.ErrorNotFound
	}

	if byID.UserID != userId {
		return models.PhotoUpdateResponse{}, models.ErrorForbiddenAccess
	}

	photo := models.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoURL: request.PhotoURL,
	}

	ress, err := ps.PhotoRepository.Update(photo, id)

	if err != nil {
		return models.PhotoUpdateResponse{}, nil
	}

	return models.PhotoUpdateResponse{
		ID:        ress.ID,
		UserID:    ress.UserID,
		Title:     ress.Title,
		Caption:   ress.Caption,
		PhotoURL:  ress.Caption,
		CreatedAt: ress.CreatedAt,
		UpdatedAt: ress.UpdatedAt,
	}, nil
}

func (ps *PhotoService) DeleteById(id string, userId string) error {
	getById, err := ps.PhotoRepository.GetById(id)
	if err != nil {
		if err != models.ErrorNotFound {
			return err
		}
		return models.ErrorNotFound
	}

	if getById.UserID != userId {
		return models.ErrorForbiddenAccess
	}

	err = ps.PhotoRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}