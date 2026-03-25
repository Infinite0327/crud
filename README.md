# crud
# Go 语言 CRUD 基础演示项目

这是一个基于 Gin 和 GORM 框架编写的后端项目，实现了用户的增删改查功能。

## 🛠️ 技术栈
- 语言：Go
- 框架：Gin,Gorm
- 数据库：MySQL

## 🚀 如何运行
1. 在本地 MySQL 中创建一个名为 `crud_demo` 的数据库。
2. 修改 `main.go` 中的 DSN 字符串，将“你的密码”改成你本地的 MySQL 密码
具体原代码是：
   `dsn := "root:你的密码@tcp(127.0.0.1:3306)/crud_demo"`

