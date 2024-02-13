package http

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/miguel-rosa/go-city-server/internal"
	"github.com/miguel-rosa/go-city-server/internal/database"
	"github.com/miguel-rosa/go-city-server/internal/post"
)

var service post.Service

func Configure() {
	service = post.Service{
		Repository: &post.RepositoryPostgres{
			Conn: database.Conn,
		},
	}
}

func PostPosts(ctx *gin.Context) {
	var post internal.Post
	if err := ctx.BindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	response, err := service.Create(ctxTimeout, post)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func GetPost(ctx *gin.Context) {
	param := ctx.Param("id")

	if param == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": post.ErrIdEmpty,
		})
		return
	}

	parsedID, err := uuid.Parse(param)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": post.ErrUUIDInvalid,
		})
		return
	}

	p, err := service.FindOneByID(ctx, parsedID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == post.ErrPostNotFound {
			statusCode = http.StatusNotFound
		}

		ctx.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, p)
}

func GetAll(ctx *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var pagination internal.Pagination

	pagination.Limit = 10
	pagination.Offset = 0

	ctx.ShouldBindQuery(&pagination)

	posts, err := service.FindAll(ctxTimeout, pagination)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"posts": posts})
}
