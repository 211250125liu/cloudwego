package main

import (
	"context"
	"errors"
	"time"

	student "github.com/njuer/course/cloudwego/rpcstusvr/kitex_gen/student"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct {
	db *gorm.DB
}
type dbStu struct {
	gorm.Model
	Id          int32
	Name        string
	Collegename string
	Collegeaddr string
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (s *StudentServiceImpl) InitDB() *StudentServiceImpl {
	db, err := gorm.Open(sqlite.Open("info.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// create table
	err = db.AutoMigrate(&dbStu{})
	if err != nil {
		panic(err)
	}
	s.db = db
	return s
}

func ToDbStu(stu *student.Student) *dbStu {
	return &dbStu{
		Id:          stu.Id,
		Name:        stu.Name,
		Collegename: stu.College.Name,
		Collegeaddr: stu.College.Address,
	}
}
func ToStu(dbStu *dbStu) *student.Student {
	return &student.Student{
		Id:   dbStu.Id,
		Name: dbStu.Name,
		College: &student.College{
			Name:    dbStu.Collegename,
			Address: dbStu.Collegeaddr,
		},
	}
}
func assign(s *student.Student, dbStu *dbStu) {
	dbStu.Id = s.Id
	dbStu.Name = s.Name
	dbStu.Collegename = s.College.Name
	dbStu.Collegeaddr = s.College.Address
}

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, theStudent *student.Student) (resp *student.RegisterResp, err error) {
	// TODO: Your code here...
	tstudent := ToDbStu(theStudent)
	var dbStuInstance dbStu
	res := s.db.First(&dbStuInstance, tstudent.Id)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		s.db.Create(tstudent)
	} else {
		assign(theStudent, &dbStuInstance)
		s.db.Save(&dbStuInstance)
	}
	resp = &student.RegisterResp{
		Success: true,
		Message: "success",
	}
	return
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *student.QueryReq) (resp *student.Student, err error) {
	// TODO: Your code here...
	stuId := req.Id
	var dbStuInstance dbStu
	res := s.db.First(&dbStuInstance, "id = ?", stuId)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	} else {
		resp = ToStu(&dbStuInstance)
	}
	return
}
