# Note

## Introduction to Project

### Programming Languages

* Golang

### Framework/Library

* https://github.com/gorilla/mux
* https://github.com/go-sql-driver/mysql
* net/http
* fmt
* log 
* encoding/json
* database/sql

### Database

* MYSQL

### File Structure

```bash
todos
├── README.md
├── api
│   └── api.go          // API implementation
├── db
│   └── db.go           // database implementation
├── go.mod              // Go pkg configuration
├── go.sum              // Go pkg configuration
├── main.go             // main entry of program
└── utility
    └── utility.go      // Utility implementation
```

### Installation/Launch program

#### With Go Install

```bash
-- WorkSpace/
 |___ bin/
    |___ executable files
 |___ pkg/
 |___ src/
    |___ todos/
        |__ api/
        ...

$ cd WorkSpace/src/todos
$ go install
$ cd WorkSpace/bin/
$ ./todos
```

#### With Go run

```bash
-- todos/
    |__ main.go
    ...

$ cd todos
$ go run main.go
```

### API Usage

#### GET /todos

從資料庫取得所有 todos，並把所有的todos list轉為json發送response

##### commands

```bash
$ curl -X GET localhost:9090/todos
```

##### HTTP status code

* 200 Status OK
    * 成功取得資料並發送 JSON response
* 500 Status Internal Server Error
    * 無法從資料庫取得資料
    * 資料庫可能發送錯誤

##### Response data

格式為 JSON，以下為舉例
```json
[
    {
        "id":1,
        "Title":"This is title",
        "Description":"This is description",
        "Completed":false,
        "CreatedAt":"2023-04-13 14:52:06"
    },
    {
        "id":2,
        "Title":"This is title2",
        "Description":"This is description2",
        "Completed":false,
        "CreatedAt":"2023-04-13 14:52:06"
    },
    ...
]
```



#### GET /todos/:id

從資料庫取得該id對應的todos，並把該todos轉為json發送response

##### commands

```bash
$ curl -X GET localhost:9090/todos/1
```

##### HTTP status code

* 200 Status OK
    * 成功取得資料並發送 JSON response
* 400 Status Bad Request
    * params中的 ID 不為 integer或不合法 
* 404 Status Not Found
    * 找不到對應 ID 的 todos
    * 不存在該 todos

##### Response data

格式為 JSON，以下為舉例
```json
{
    "id":1,
    "Title":"This is title",
    "Description":"This is description",
    "Completed":false,
    "CreatedAt":"2023-04-13 14:52:06"
}
```

#### POST /todos

取得 http request中的 POST request Data (格式為 JSON)，並將該資料存入 database中。 id, Completed, CreatedAt 皆為預設值

##### commands

```bash
$ curl -X POST "localhost:9090/todos" -H 'Content-Type: application/json' -d '{"Title":"This is test title", "Description": "Test Description"}'
```

##### HTTP status code

* 200 Status OK
    * 成功存入資料並發送 JSON response
* 400 Status Bad Request
    * POST Request Data 不為合法的 JSON 格式
* 500 Status Internal Server Error
    * 無法存入database中
    * 資料庫可能發生錯誤

##### Response data

格式為 JSON，以下為舉例
```json
{
    "id":7,
    "Title":"This is test title",
    "Description":"Test Description",
    "Completed":false,
    "CreatedAt":"2023-04-13 15:23:24"
}
```

#### PUT /todos/:id

取得 http request 中的 PUT request Data (格式為 JSON)，並將該資料更新對應 id 的 todos。id 不能被更改

##### commands

```bash
$ curl -X PUT "localhost:9090/todos/7" -H 'Content-Type: applicaiton/json' -d '{"Title":"ChangedTitle","Description":"ChangedDescription","Completed":true,"CreatedAt":"2023-04-13 15:00:00"}'
```

##### HTTP status code

* 200 Status OK
    * 成功更新資料並發送 JSON response
* 400 Status Bad Request
    * PUT Request Data 不為合法的 JSON 格式
    * ID 不合法
* 500 Status Internal Server Error
    * 無法更新database
    * 無法重新 Refetch 更新的資料
    * 資料庫可能發生錯誤


##### Response data

格式為 JSON，以下為舉例
```json
{
    "id":7,
    "Title":"ChangedTitle",
    "Description":"ChangedDescription",
    "Completed":true,
    "CreatedAt":"2023-04-13 15:00:00"
}
```

#### DELETE /todos/:id

刪除對應 id 的 todos


##### commands

```bash
$ curl -X DELETE localhost:9090/todos/7
```

##### HTTP status code

* 200 Status OK
    * 成功刪除資料
* 400 Status Bad Request
    * ID 不合法
* 500 Status Internal Server Error
    * 無法刪除資料
    * 資料庫可能發生錯誤


##### Response data

格式為 JSON，此 API 預設回傳空json
```json
{}
```