package services

import (
	"Project-Akhir/helpers"
	"Project-Akhir/models"
	"Project-Akhir/repositori"
	"fmt"
)

type SocialMediaService struct {
	SocialMediaRepository repositori.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepository repositori.SocialMediaRepository) *SocialMediaService {
	return &SocialMediaService{
		SocialMediaRepository: socialMediaRepository,
	}
}

func (sms *SocialMediaService) Add(req models.SocialMediaCreateRequest, userId string) (models.SocialMediaCreateResponse, error) {
	id := helpers.GenerateID()

	socialMedia := models.SocialMedia{
		ID:             id,
		Name:           req.Name,
		SocialMediaURL: req.SocialMediaURL,
		UserID:         userId,
	}

	response, err := sms.SocialMediaRepository.Save(socialMedia)
	if err != nil {
		if err != models.ErrorNotFound {
			return models.SocialMediaCreateResponse{}, models.ErrorNotFound
		}
		return models.SocialMediaCreateResponse{}, err
	}
	return models.SocialMediaCreateResponse{
		ID:             response.ID,
		UserID:         response.ID,
		Name:           response.Name,
		SocialMediaURL: response.SocialMediaURL,
		CreatedAt:      response.CreatedAt,
	}, nil

}

func (sms *SocialMediaService) GetAll() ([]models.SocialMediaResponse, error) {
	sosmed := make([]models.SocialMediaResponse, 0)

	ress, err := sms.SocialMediaRepository.Get()

	if err != nil {
		return []models.SocialMediaResponse{}, err
	}

	for _, yes := range ress {
		sosmed = append(sosmed, models.SocialMediaResponse{
			ID:             yes.ID,
			UserID:         yes.UserID,
			Name:           yes.Name,
			SocialMediaURL: yes.SocialMediaURL,
			CreatedAt:      yes.CreatedAt,
			UpdatedAt:      yes.UpdatedAt,
		})
	}

	fmt.Println("response", sosmed)
	return sosmed, nil

}

func (sms *SocialMediaService) GetById(id string) (models.SocialMediaResponse, error) {
	sosmed, err := sms.SocialMediaRepository.GetById(id)

	if err != nil {
		if err != models.ErrorNotFound {
			return models.SocialMediaResponse{}, err
		}
		return models.SocialMediaResponse{}, models.ErrorNotFound
	}

	return models.SocialMediaResponse{
		ID:             sosmed.ID,
		UserID:         sosmed.UserID,
		Name:           sosmed.Name,
		SocialMediaURL: sosmed.SocialMediaURL,
		CreatedAt:      sosmed.CreatedAt,
		UpdatedAt:      sosmed.UpdatedAt,
	}, nil
}

func (sms *SocialMediaService) UpdateById(request models.SocialMediaUpdateRequest, id string, userId string) (models.SocialMediaUpdateResponse, error) {
	getById, err := sms.SocialMediaRepository.GetById(id)
	if err != nil {
		if err != models.ErrorNotFound {
			return models.SocialMediaUpdateResponse{}, err
		}
		return models.SocialMediaUpdateResponse{}, models.ErrorNotFound
	}

	if getById.UserID != userId {
		return models.SocialMediaUpdateResponse{}, models.ErrorForbiddenAccess
	}

	socialMedia := models.SocialMedia{
		Name:           request.Name,
		SocialMediaURL: request.SocialMediaURL,
	}

	res, err := sms.SocialMediaRepository.Update(socialMedia, id)
	if err != nil {
		return models.SocialMediaUpdateResponse{}, err
	}

	return models.SocialMediaUpdateResponse{
		ID:             res.ID,
		UserID:         res.UserID,
		Name:           res.Name,
		SocialMediaURL: res.SocialMediaURL,

		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (sms *SocialMediaService) DeleteById(id string, userId string) error {
	getById, err := sms.SocialMediaRepository.GetById(id)

	if err != nil {
		if err != models.ErrorNotFound {
			return err
		}
		return models.ErrorNotFound
	}

	if getById.UserID != userId {
		return models.ErrorForbiddenAccess
	}

	err = sms.SocialMediaRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil

}
