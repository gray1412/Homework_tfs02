# Bài tập 2
- Thiết kế database cho chủ đề khoá học OCG TFS 
- Viết các lệnh CRUD (Create/Read/Update/Delete) với schema trên sử dụng golang + gorm + mysql

# Project's tree
```
D:.
│   design_database.PNG
│   go.mod
│   go.sum
│   main.go
│   README.MD
│
├───.idea
│       .gitignore
│       courseocg.iml
│       modules.xml
│       vcs.xml
│       workspace.xml
│
├───.vscode
│       settings.json
│
├───database
│       connectdb.go
│
├───handler
│       user_handler.go
│
├───service
│       user_service.go
│
├───storage
│   │   model.go
│   │
│   ├───request
│   │       user_request.go
│   │
│   └───respone
│           response_with_json.go
│           user_response.go
│
└───validator
        validator.go
```