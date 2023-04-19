package controllers

import (
	"Project-Akhir/models"
	"Project-Akhir/services"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type SocialMediaController struct {
	SocialMediaService services.SocialMediaService
}

func NewSocialMediaController(socialMediaService services.SocialMediaService) *SocialMediaController {
	return &SocialMediaController{
		SocialMediaService: socialMediaService,
	}
}

func (smc *SocialMediaController) CreateSocialMedia(ctx *gin.Context) {
	newSocialMedia := models.SocialMediaCreateRequest{}

	if err := ctx.ShouldBindJSON(&newSocialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}
	valid, err := valid.ValidateStruct(newSocialMedia)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	userId, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: models.ErrorInvalidToken.Err,
		})
		return
	}

	res, err := smc.SocialMediaService.Add(newSocialMedia, userId.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.ResponseSuccess{
		Meta: models.Meta{
			Code:    http.StatusCreated,
			Message: http.StatusText(http.StatusCreated),
		},
		Data: res,
	})
}

func (smc *SocialMediaController) GetAllSocialMedia(ctx *gin.Context) {
	sosmed, err := smc.SocialMediaService.GetAll()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseSuccess{
		Meta: models.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: sosmed,
	})
}

func (smc *SocialMediaController) GetByIdSocialMedia(ctx *gin.Context) {
	id := ctx.Param("id") // ini berdasarkan id colom social media ?

	sosmed, err := smc.SocialMediaService.GetById(id)

	if err != nil {
		if err == models.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, models.ResponseFailed{
				Meta: models.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},
				Error: err.Error(),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseSuccess{
		Meta: models.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: sosmed,
	})
}

func (smc *SocialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	update := models.SocialMediaUpdateRequest{}

	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&update); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}
	valid, err := valid.ValidateStruct(update)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	userId, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: models.ErrorInvalidToken.Err,
		})
		return
	}

	res, err := smc.SocialMediaService.UpdateById(update, id, userId.(string))
	if err != nil {
		if err == models.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, models.ResponseFailed{
				Meta: models.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},
				Error: err.Error(),
			})
			return

		} else if err == models.ErrorForbiddenAccess {
			ctx.AbortWithStatusJSON(http.StatusForbidden, models.ResponseFailed{
				Meta: models.Meta{
					Code:    http.StatusForbidden,
					Message: http.StatusText(http.StatusForbidden),
				},
				Error: err.Error(),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, models.ResponseSuccess{
		Meta: models.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: res,
	})
}

func (smc *SocialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	id := ctx.Param("id")

	userId, isExist := ctx.Get("user_id")

	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: models.ErrorInvalidToken.Err,
		})
		return
	}

	err := smc.SocialMediaService.DeleteById(id, userId.(string))

	if err != nil {
		if err == models.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, models.ResponseFailed{
				Meta: models.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},

				Error: err.Error(),
			})
			return

		} else if err == models.ErrorForbiddenAccess {
			ctx.AbortWithStatusJSON(http.StatusForbidden, models.ResponseFailed{
				Meta: models.Meta{
					Code:    http.StatusForbidden,
					Message: http.StatusText(http.StatusForbidden),
				},
				Error: err.Error(),
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, "berhasil menghapus social media")
}
