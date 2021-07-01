# Project's Info
Save data of book to MySQL database and how to use eleasticsearch to store book's information and perform some search, delete actions

## Project's tree
```
.
├── data
│   ├── train.csv
├── handler
│   ├── handler.go
├── storage
│   ├── book_manager.go
│   ├── book.go
│   ├── es_client.go
│   └── mysql_storage.go

├── README.md
├── go.mod
├── go.sum
└── main.go
```

## Note
Download libraries by running
```
go get github.com/olivere/elastic/v7

```

