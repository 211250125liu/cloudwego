package main

import (
	"context"
	"errors"
	"time"

	teacher "github.com/njuer/course/cloudwego/rpcteasvr/kitex_gen/teacher"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TeacherServiceImpl implements the last service interface defined in the IDL.
type TeacherServiceImpl struct {
	db *gorm.DB
}
type dbTea struct {
	gorm.Model
	Id          int32
	Name        string
	Collegename string
	Collegeaddr string
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (s *TeacherServiceImpl) InitDB() *TeacherServiceImpl {
	db, err := gorm.Open(sqlite.Open("info.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// create table
	err = db.AutoMigrate(&dbTea{})
	if err != nil {
		panic(err)
	}
	s.db = db
	return s
}

func ToDbTea(tea *teacher.Teacher) *dbTea {
	return &dbTea{
		Id:          tea.Id,
		Name:        tea.Name,
		Collegename: tea.College.Name,
		Collegeaddr: tea.College.Address,
	}
}
func ToTea(dbTea *dbTea) *teacher.Teacher {
	return &teacher.Teacher{
		Id:   dbTea.Id,
		Name: dbTea.Name,
		College: &teacher.College{
			Name:    dbTea.Collegename,
			Address: dbTea.Collegeaddr,
		},
	}
}
func assign(s *teacher.Teacher, dbTea *dbTea) {
	dbTea.Id = s.Id
	dbTea.Name = s.Name
	dbTea.Collegename = s.College.Name
	dbTea.Collegeaddr = s.College.Address
}

// Register implements the TeacherServiceImpl interface.
func (s *TeacherServiceImpl) Register(ctx context.Context, theTeacher *teacher.Teacher) (resp *teacher.RegisterResp, err error) {
	// TODO: Your code here...
	tteacher := ToDbTea(theTeacher)
	var dbTeaInstance dbTea
	res := s.db.First(&dbTeaInstance, tteacher.Id)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		s.db.Create(tteacher)
	} else {
		assign(theTeacher, &dbTeaInstance)
		s.db.Save(&dbTeaInstance)
	}
	resp = &teacher.RegisterResp{
		Success: true,
		Message: "success",
	}
	return
}

// Query implements the TeacherServiceImpl interface.
func (s *TeacherServiceImpl) Query(ctx context.Context, req *teacher.QueryReq) (resp *teacher.Teacher, err error) {
	// TODO: Your code here...
	teaId := req.Id
	var dbTeaInstance dbTea
	res := s.db.First(&dbTeaInstance, "id = ?", teaId)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	} else {
		resp = ToTea(&dbTeaInstance)
	}
	return
}
