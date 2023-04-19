package controllers

import (
	"Project-Akhir/models"
	"Project-Akhir/services"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	CommentService services.CommentService
}

func NewCommentControoler(commentService services.CommentService) *CommentController {
	return &CommentController{
		CommentService: commentService,
	}
}

func (cc *CommentController) CreateCommentByPhotoID(ctx *gin.Context) {
	commentRequest := models.CommentCreateRequest{}

	if err := ctx.ShouldBindJSON(&commentRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	valid, err := valid.ValidateStruct(commentRequest)
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

	photoId := ctx.Param("id_photo")
	result, err := cc.CommentService.Creat(commentRequest, userId.(string), photoId)

	if err != nil {
		if err == models.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ResponseFailed{
				Meta: models.Meta{
					Code:    http.StatusInternalServerError,
					Message: http.StatusText(http.StatusInternalServerError),
				},
				Error: err.Error(),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusNotFound, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusNotFound,
				Message: http.StatusText(http.StatusNotFound),
			},
			Error: "Comment " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.ResponseSuccess{
		Meta: models.Meta{
			Code:    http.StatusCreated,
			Message: http.StatusText(http.StatusCreated),
		},
		Data: result,
	})

}

func (cc *CommentController) GetAllCommnet(ctx *gin.Context) {
	comments, err := cc.CommentService.GetAll()

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
		Data: comments,
	})
}

func (cc *CommentController) GetCommentsByID(ctx *gin.Context) {
	id := ctx.Param("id")

	comment, err := cc.CommentService.GetById(id)
	if err != nil {
		if err == models.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, models.ResponseFailed{
				Meta: models.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},
				Error: "Comment " + err.Error(),
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
		Data: comment,
	})
}

func (cc *CommentController) UpdateComment(ctx *gin.Context) {
	commentRequest := models.CommentUpdateRequest{}

	if err := ctx.ShouldBindJSON(&commentRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ResponseFailed{
			Meta: models.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	valid, err := valid.ValidateStruct(commentRequest)
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

	id := ctx.Param("id")
	result, err := cc.CommentService.UpdateById(commentRequest, userId.(string), id)

	if err != nil {
		if err == models.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, models.ResponseFailed{
				Meta: models.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},
				Error: "Comment " + err.Error(),
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
		Data: result,
	})
}

func (cc *CommentController) DeleteComment(ctx *gin.Context) {

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

	id := ctx.Param("id")
	err := cc.CommentService.DeleteById(userId.(string), id)

	if err != nil {
		if err == models.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, models.ResponseFailed{
				Meta: models.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},
				Error: "Comment " + err.Error(),
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
		Data: "Delete comment success.",
	})
}
