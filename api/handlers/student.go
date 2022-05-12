package httphandler

import (
	"fmt"
	"net/http"
	dto "root/domain"
	interfaces "root/interface"
	"strconv"

	"github.com/gin-gonic/gin"
)

type studentHandler struct {
	studentUsecase interfaces.StudentUsecase
}

func NewStudentHandler(r *gin.Engine, l interfaces.StudentUsecase) {
	handler := studentHandler{studentUsecase: l}
	r.GET("/students", handler.GetStudent)
	r.POST("/student", handler.create)

	r.PUT("/student/:id", handler.update)
	r.DELETE("/student/:id", handler.delete)

}

func (l *studentHandler) GetStudent(ctx *gin.Context) {
	//fmt.Println(name, book)
	author, err := l.studentUsecase.GetStudent(ctx)
	if err != nil {

		return
	}

	ctx.JSON(http.StatusOK, author)
}

func (l *studentHandler) delete(ctx *gin.Context) {
	id := ctx.Param("id")
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("invalid id")
		return
	}
	err = l.studentUsecase.DeleteStudent(ctx, id64)
	if err != nil {
		fmt.Println("Unable to Delete")
	}
}

func (l *studentHandler) update(ctx *gin.Context) {
	req := new(dto.UpdateStudentRequest)
	if err := ctx.Bind(req); err != nil {
		fmt.Println("error", err)
		return
	}
	err := l.studentUsecase.UpdateStudent(ctx, req.Firstname, req.Lastname)
	if err != nil {
		fmt.Println("Unable to Update")
		return
	}

}

func (l *studentHandler) create(ctx *gin.Context) {
	req := new(dto.Student)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {

		fmt.Printf("%+v", err)
		ctx.JSON(http.StatusOK, err)
		return
	}
	//fmt.Printf("2......\n")
	err = l.studentUsecase.CreateStudent(ctx, req)
	if err != nil {
		fmt.Printf("%+v\n", err)
		ctx.JSON(400, gin.H{"status": "cannot create student data"})
		return
	}

	ctx.JSON(201, "Created")
}
