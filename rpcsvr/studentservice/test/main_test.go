package test

import (
	"bytes"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/njuer/course/cloudwego/rpcstusvr/kitex_gen/student"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

const (
	queryURLFmt = "http://127.0.0.1:8888/get/student/query-student-info?id="
	registerURL = "http://127.0.0.1:8888/post/student/add-student-info"
)

var httpCli = &http.Client{Timeout: 3 * time.Second}

func TestStudentService(t *testing.T) {
	for i := 1; i <= 100; i++ {
		newStu := genStudent(i)
		resp, err := register(newStu)
		Assert(t, err == nil, err)
		Assert(t, resp.Success)

		stu, err := query(i)
		Assert(t, err == nil, err)
		Assert(t, stu.Id == newStu.Id)
		Assert(t, stu.Name == newStu.Name)
		Assert(t, stu.College.Name == newStu.College.Name)
	}
}

func BenchmarkStudentService(b *testing.B) {
	for i := 1; i < b.N; i++ {
		newStu := genStudent(i)
		resp, err := register(newStu)
		Assert(b, err == nil, err)
		Assert(b, resp.Success, resp.Message)

		stu, err := query(i)
		Assert(b, err == nil, err)
		Assert(b, stu.Id == newStu.Id)
		Assert(b, stu.Name == newStu.Name, newStu.Id, stu.Name, newStu.Name)
		Assert(b, stu.College.Name == newStu.College.Name)
	}
}

func register(stu *student.Student) (rResp *student.RegisterResp, err error) {
	reqBody, err := json.Marshal(stu)
	if err != nil {
		return nil, fmt.Errorf("marshal request failed: err=%v", err)
	}
	var resp *http.Response
	req, err := http.NewRequest(http.MethodPost, registerURL, bytes.NewBuffer(reqBody))
	resp, err = httpCli.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	if err = json.Unmarshal(body, &rResp); err != nil {
		return
	}
	return
}

func query(id int) (student student.Student, err error) {
	var resp *http.Response
	resp, err = httpCli.Get(fmt.Sprint(queryURLFmt, id))
	defer resp.Body.Close()
	if err != nil {
		return
	}
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	if err = json.Unmarshal(body, &student); err != nil {
		return
	}
	return
}

func genStudent(id int) *student.Student {
	return &student.Student{
		Id:   int32(id),
		Name: fmt.Sprintf("student-%d", id),
		College: &student.College{
			Name:    "",
			Address: "",
		},
	}
}

// Assert asserts cond is true, otherwise fails the test.
func Assert(t testingTB, cond bool, val ...interface{}) {
	t.Helper()
	if !cond {
		if len(val) > 0 {
			val = append([]interface{}{"assertion failed:"}, val...)
			t.Fatal(val...)
		} else {
			t.Fatal("assertion failed")
		}
	}
}

// testingTB is a subset of common methods between *testing.T and *testing.B.
type testingTB interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Helper()
}
