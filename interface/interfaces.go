package interfaces

import (
	dto "root/domain"

	"github.com/gin-gonic/gin"
)

type StudentUsecase interface {
	GetStudent(c *gin.Context) ([]*dto.Student, error)
	DeleteStudent(c *gin.Context, id int64) error
	CreateStudent(c *gin.Context, request *dto.Student) error
	UpdateStudent(c *gin.Context, name string, name1 string) error
}
