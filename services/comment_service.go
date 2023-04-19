package services

import (
	"Project-Akhir/helpers"
	"Project-Akhir/models"
	"Project-Akhir/repositori"
)

type CommentService struct {
	CommentRepository repositori.CommentRepository
	PhotoRepository   repositori.PhotoRepository
}

func NewCommentService(commentRepository repositori.CommentRepository, photoRepository repositori.PhotoRepository) *CommentService {
	return &CommentService{
		CommentRepository: commentRepository,
		PhotoRepository:   photoRepository,
	}
}

func (cs *CommentService) Creat(request models.CommentCreateRequest, userId string, photoId string) (models.CommentCreateResponse, error) {
	id := helpers.GenerateID()

	_, err := cs.PhotoRepository.GetById(photoId)
	if err != nil {
		return models.CommentCreateResponse{}, err
	}

	comment := models.Comment{
		ID:      id,
		UserID:  userId,
		PhotoID: photoId,
		Massage: request.Massage,
	}

	res, err := cs.CommentRepository.Save(comment)
	if err != nil {
		return models.CommentCreateResponse{}, err
	}

	return models.CommentCreateResponse{
		ID:        res.ID,
		UserID:    res.UserID,
		PhotoID:   res.PhotoID,
		Message:   res.Massage,
		CreatedAt: res.CreatedAt,
	}, nil
}

func (cs *CommentService) GetAll() ([]models.CommentResponse, error) {
	cResponse := make([]models.CommentResponse, 0)

	response, err := cs.CommentRepository.GetAll()

	if err != nil {
		return []models.CommentResponse{}, err
	}

	for _, val := range response {
		cResponse = append(cResponse, models.CommentResponse{
			ID:        val.ID,
			UserID:    val.UserID,
			PhotoID:   val.PhotoID,
			Message:   val.Massage,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}

	return cResponse, err
}

func (cs *CommentService) GetById(id string) (models.CommentResponse, error) {
	response, err := cs.CommentRepository.GetById(id)

	if err != nil {
		if err == models.ErrorNotFound {
			return models.CommentResponse{}, models.ErrorNotFound
		}
		return models.CommentResponse{}, err
	}

	return models.CommentResponse{
		ID:        response.ID,
		UserID:    response.UserID,
		PhotoID:   response.PhotoID,
		Message:   response.Massage,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	}, nil
}

func (cs *CommentService) UpdateById(request models.CommentUpdateRequest, userId string, id string) (models.CommentUpdateResponse, error) {
	comment, err := cs.CommentRepository.GetById(id)

	if err != nil {
		if err == models.ErrorNotFound {
			return models.CommentUpdateResponse{}, models.ErrorNotFound
		}
		return models.CommentUpdateResponse{}, err
	}

	if comment.UserID != userId {
		return models.CommentUpdateResponse{}, models.ErrorForbiddenAccess
	}

	commentUpdate := models.Comment{
		Massage: request.Message,
	}

	res, err := cs.CommentRepository.Update(commentUpdate, id)
	if err != nil {
		return models.CommentUpdateResponse{}, err
	}

	return models.CommentUpdateResponse{
		ID:        res.ID,
		UserID:    res.UserID,
		PhotoID:   res.PhotoID,
		Message:   res.Massage,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil

}

func (cs *CommentService) DeleteById(userId string, id string) error {

	comment, err := cs.CommentRepository.GetById(id)

	if err != nil {
		if err == models.ErrorNotFound {
			return models.ErrorNotFound
		}
		return err
	}

	if comment.UserID != userId {
		return models.ErrorForbiddenAccess
	}

	err = cs.CommentRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
