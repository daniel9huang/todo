package main

import (
	"database/sql"
	"net/http"

	//    "encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	HOST = "0.0.0.0"
	PORT = 5432

	DB_USER     = "postgres"
	DB_PASSWORD = "123456"
	DB_NAME     = "todo"
)

type Database struct {
	Conn *sql.DB
}

func Initialize() (Database, error) {
	db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, DB_USER, DB_PASSWORD, DB_NAME)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	return db, nil
}

func getAllTasks(context *gin.Context) {
	db, _ := Initialize()
	rows, _ := db.Conn.Query("select * from task")

	var tasks []task

	for rows.Next() {
		var id int
		var title string
		var body string
		rows.Scan(&id, &title, &body)

		tasks = append(tasks, task{Title: title, Body: body})

		fmt.Println(id)
	}
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": tasks})
}

func getOneTaskByTitle(context *gin.Context) {
	title := context.Param("title")
	db, _ := Initialize()
	oneTaskQueryString := fmt.Sprintf("SELECT * FROM task where title='%s'", title)
	rows, _ := db.Conn.Query(oneTaskQueryString)

	var tasks []task

	for rows.Next() {
		var id int
		var title string
		var body string
		rows.Scan(&id, &title, &body)

		tasks = append(tasks, task{Title: title, Body: body})
	}
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": tasks[0]})
}

func insertTask(context *gin.Context) {
	title := context.PostForm("title")
	body := context.PostForm("body")
	db, _ := Initialize()
	oneTaskQueryString := fmt.Sprintf("insert into task values(DEFAULT, '%s', '%s');", title, body)
	db.Conn.QueryRow(oneTaskQueryString)
}

func updateTaskBodybyTitle(context *gin.Context) {
	title := context.PostForm("title")
	newBody := context.PostForm("body")
	db, _ := Initialize()
	oneTaskQueryString := fmt.Sprintf("update task set body='%s' where title='%s'", newBody, title)
	db.Conn.QueryRow(oneTaskQueryString)
}

func deleteTaskByTitle(context *gin.Context) {
	title := context.Param("title")
	db, _ := Initialize()
	oneTaskQueryString := fmt.Sprintf("delete from task where title='%s'", title)
	db.Conn.QueryRow(oneTaskQueryString)
}
