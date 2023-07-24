# CloudWeGo

| 更新人 | 更新时间  | 更新内容                                              |
| ------ | --------- | ----------------------------------------------------- |
| 张淞钦 | 2023.7.20 | 添加项目情景、go module部分的命名统一、数据库字段初设 |
| 张淞钦 | 2023.7.21 | 项目基本功能完成                                      |
|        |           |                                                       |
|        |           |                                                       |
|        |           |                                                       |
|        |           |                                                       |
|        |           |                                                       |
|        |           |                                                       |
|        |           |                                                       |



## 项目情景

实现**学院与教师**、**学生与学院**之间的查询、插入（、修改、删除）管理。

服务分为

- 学院与教师（teacherService）
- 学生与学院（studentService）

数据库表描述：

1. 学院（college）
   - collegeId（学员编号）
   - address（学院地址）
   - collegeName（学院名称）
2. 教师
   - college（所属学院）
   - teacherName
   - teacherId
   - position（职位，此项作为热更新idl，初始阶段不设置该字段）
3. 学生
   - stuId
   - stuName
   - college（所属学院）

## 项目统一约定

### goModule名称：

hertzSvr部分：github.com/njuer/course/cloudwego/httpsvr

rpcSvr的student部分：github.com/njuer/course/cloudwego/rpcstusvr

rpcSvr的teacher部分：github.com/njuer/course/cloudwego/rpcteasvr

### 端口

etcd 2379

hertzsvr 

rpcstudent 9998

rpcteacher 9999

## 项目测试

请使用Day3中给出的文档中的测试用例，按照代码设计去更改测试用例去测试。

项目运行请参见Day4完成的项目的运行方法。