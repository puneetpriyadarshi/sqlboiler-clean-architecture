package usecase

import (
	"database/sql"
	"fmt"

	dto "root/domain"
	interfaces "root/interface"
	"root/models"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type studentUsecase struct {
	db *sql.DB
}

func NewStudentUsecase(db *sql.DB) interfaces.StudentUsecase {
	return &studentUsecase{
		db: db,
	}
}
func ToStudentDTO(student *models.Student) *dto.Student {
	return &dto.Student{
		ID:        student.ID,
		Firstname: student.Firstname.String,
		Lastname:  student.Lastname.String,
	}
}

func ToStudentSliceDTO(adminUserSlice models.StudentSlice) []*dto.Student {
	adminUsers := make([]*dto.Student, 0, len(adminUserSlice))
	for _, adminUser := range adminUserSlice {
		adminUsers = append(adminUsers, ToStudentDTO(adminUser))
	}
	return adminUsers
}

func (c *studentUsecase) GetStudent(ctx *gin.Context) ([]*dto.Student, error) {

	users, err := models.Students().All(ctx, c.db)
	var userList []*dto.Student

	fmt.Printf("users--->>>%+v\n", userList)
	return ToStudentSliceDTO(users), err

}

func (c *studentUsecase) DeleteStudent(ctx *gin.Context, id int64) error {

	del, err := models.Students(qm.Where("id=?", id)).DeleteAll(ctx, c.db)
	if err != nil {
		fmt.Println("unable to delete")
		return err
	}
	fmt.Printf("%+v deleted", del)
	return err
}

func (c *studentUsecase) UpdateStudent(ctx *gin.Context, name string, name1 string) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	fmt.Printf("id ----%+v\n", id)
	student, err := models.FindStudent(ctx, c.db, int64(id))
	if err != nil {
		fmt.Println("helooo-------")
		fmt.Printf("%+v", err)
	}
	fmt.Printf("%+v", student)
	if name == "" && name1 != "" {
		student.Lastname = null.StringFrom(name1)
	}
	if name != "" && name1 == "" {
		student.Firstname = null.StringFrom(name)
	}
	// if name1 == "" {
	// 	log.Println("CANNOT BE UPDATED")
	// } else {
	// 	student.Lastname = null.StringFrom(name)
	// }
	_, err = student.Update(ctx, c.db, boil.Infer())

	if err != nil {
		return err
	}
	return nil
}

func (c *studentUsecase) CreateStudent(ctx *gin.Context, request *dto.Student) error {
	//err := ctx.BindJSON(&request)
	//fmt.Printf("%+v, %+v", request, err)
	//if err != nil {
	//	return dto.CreateAuthorResponse{}, err
	//}
	studentData := models.Student{
		ID:        request.ID,
		Firstname: null.StringFrom(request.Firstname),
		Lastname:  null.StringFrom(request.Lastname),
	}

	err := studentData.Insert(ctx, c.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
