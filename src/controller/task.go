package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/FukuHiro12111/go-todo-api/src/model"
	"github.com/gin-gonic/gin"
)

// タスク一覧
func TaskGET(c *gin.Context) {
	db := model.DBConnect()
	result, err := db.Query("SELECT * FROM task ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	// json返却用
	tasks := []model.Task{}
	for result.Next() {
		task := model.Task{}
		var id uint
		var createAt, updateAt time.Time
		var title string

		err = result.Scan(&id, &createAt, &updateAt, &title)
		if err != nil {
			panic(err.Error())
		}

		task.ID = id
		task.CreateAt = createAt
		task.UpdateAt = updateAt
		task.Title = title
		tasks = append(tasks, task)
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// タスク検索
func FindByID(id uint) model.Task {
	db := model.DBConnect()
	result, err := db.Query("SELECT * FROM task WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	task := model.Task{}
	for result.Next() {
		var createAt, updateAt time.Time
		var title string

		err = result.Scan(&id, &createAt, &updateAt, &title)
		if err != nil {
			panic(err.Error())
		}

		task.ID = id
		task.CreateAt = createAt
		task.UpdateAt = updateAt
		task.Title = title
	}
	return task
}

// タスク登録
func TaskPOST(c *gin.Context) {
	db := model.DBConnect()

	title := c.PostForm("title")
	now := time.Now()

	_, err := db.Exec("INSERT INTO task (title, create_at, update_at) VALUES(?, ?, ?)", title, now, now)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("post sent. title: %s", title)
}

// タスク更新
func TaskPATCH(c *gin.Context) {
	db := model.DBConnect()

	id, _ := strconv.Atoi(c.Param("id"))
	title := c.PostForm("title")
	now := time.Now()

	_, err := db.Exec("UPDATE task SET title = ?, update_at=? WHERE id = ?", title, now, id)
	if err != nil {
		panic(err.Error())
	}

	task := FindByID(uint(id))

	fmt.Println(task)
	c.JSON(http.StatusOK, gin.H{"task": task})
}

// タスク削除
func TaskDELETE(c *gin.Context) {
	db := model.DBConnect()

	id, _ := strconv.Atoi(c.Param("id"))

	_, err := db.Query("DELETE FROM task WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, "deleted")

}