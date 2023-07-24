package idlProvider

import (
	"os"
	"time"

	"github.com/cloudwego/kitex/pkg/generic"
)

// map serviceName to idl path
var serviceIdlMap = make(map[string]string)

// map idlPath to content
var idlContentMap = make(map[string]*generic.ThriftContentProvider)

func Init() {
	// create goroutine to update idlContentMap
	go func() {
		for {
			// update idlContentMap
			// attention: provider out of date
			for k := range idlContentMap {
				err := idlContentMap[k].UpdateIDL(getIdlFileContent(k), map[string]string{})
				if err != nil {
					panic("Error: fail to update idl " + err.Error())
				}
			}
			// time.Sleep(time.Second)
			time.Sleep(10 * time.Second)
		}
	}()
	// init serviceIdlMap
	serviceIdlMap["student"] = "../idl/student.thrift"
	var err error
	idlContentMap["../idl/student.thrift"], err = generic.NewThriftContentProvider(getIdlFileContent("../idl/student.thrift"), map[string]string{})
	if err != nil {
		panic("error when init idlContentMap")
	}
	serviceIdlMap["teacher"] = "../idl/teacher.thrift"
	idlContentMap["../idl/teacher.thrift"], err = generic.NewThriftContentProvider(getIdlFileContent("../idl/teacher.thrift"), map[string]string{})
	if err != nil {
		panic("error when init idlContentMap")
	}
}
func GetIdlByServiceName(serviceName string) *generic.ThriftContentProvider {
	res := idlContentMap[serviceIdlMap[serviceName]]
	var err error
	idlContentMap[serviceIdlMap[serviceName]], err = generic.NewThriftContentProvider(getIdlFileContent(serviceIdlMap[serviceName]), map[string]string{})
	if err != nil {
		panic("error when init idlContentMap")
	}
	return res
}
func getIdlFileContent(idlPath string) string {
	content, err := os.ReadFile(idlPath)
	if err != nil {
		panic(err)
	}
	// 将内容转换为字符串
	//fmt.Println(string(content))
	return string(content)
}
