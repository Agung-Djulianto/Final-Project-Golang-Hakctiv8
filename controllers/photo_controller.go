package controllers

import (
	"Project-Akhir/models"
	"Project-Akhir/services"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type PhotoControlller struct {
	PhotoService services.PhotoService
}

func NewPhotoController(photoService services.PhotoService) *PhotoControlller {
	return &PhotoControlller{
		PhotoService: photoService,
	}
}

func (pc *PhotoControlller) CreatePhoto(ctx *gin.Context) {
	photo := models.PhotoCreateRequest{}

	if err := ctx.ShouldBindJSON(&photo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	valid, err := valid.ValidateStruct(photo)

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

	userID, isExist := ctx.Get("user_id")

	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	response, err := pc.PhotoService.Add(photo, userID.(string))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (pc *PhotoControlller) GetAllPhoto(ctx *gin.Context) {
	photo, err := pc.PhotoService.GetAll()

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
		Data: photo,
	})
}

func (pc *PhotoControlller) GetPhotoById(ctx *gin.Context) {
	id := ctx.Param("id")

	photo, err := pc.PhotoService.GetById(id)

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
		Data: photo,
	})
}

func (pc *PhotoControlller) UpdatePhoto(ctx *gin.Context) {
	updatePhoto := models.PhotoUpdateRequest{}
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&updatePhoto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}
	valid, err := valid.ValidateStruct(updatePhoto)

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
			Error: err.Error(),
		})
		return
	}

	updated, err := pc.PhotoService.UpdateById(updatePhoto, id, userId.(string))
	if err != nil {
		if err == models.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, models.ResponseFailed{
				Meta: models.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},
				Error: "Photo " + err.Error(),
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
		Data: updated,
	})
}

func (pc *PhotoControlller) DeletePhoto(ctx *gin.Context) {
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

	err := pc.PhotoService.DeleteById(id, userId.(string))

	if err != nil {
		if err == models.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, models.ResponseFailed{
				Meta: models.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},
				Error: "Photo " + err.Error(),
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
		Data: "Delete photo success.",
	})

}
