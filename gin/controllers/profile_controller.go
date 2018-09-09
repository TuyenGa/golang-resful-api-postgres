package controllers

import (
	models "gin-app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPeople(c *gin.Context) {
	var people models.People
	query := "SELECT * FROM profile"

	db, err := models.GetPostgresDb()
	if err != nil {
		panic(err)
	}
	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Age, &user.Name)

		if err != nil {
			panic(err)
		}

		people = append(people, user)
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "result": people})
}

func Create(c *gin.Context) {
	user := new(models.User)

	query := "INSERT INTO profile (name, age) VALUES ($1, $2) returning *"

	db, err := models.GetPostgresDb()

	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare(query)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	name := c.PostForm("name")
	age, _ := strconv.Atoi(c.PostForm("age"))
	err = stmt.QueryRow(name, age).Scan(&user.ID, &user.Age, &user.Name)

	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "result": user})
}

func Update(c *gin.Context) {
	query := `UPDATE profile SET name=$1, age=$2 WHERE id=$3`
	db, err := models.GetPostgresDb()
	if err != nil {
		panic(err)
	}
	id := c.Param("id")
	name := c.PostForm("name")
	age, _ := strconv.Atoi(c.PostForm("age"))
	_, err = db.Exec(query, name, age, id)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "result": "ok"})
}

func Delete(c *gin.Context) {
	query := "DELETE FROM profile WHERE id = $1"

	db, err := models.GetPostgresDb()
	if err != nil {
		panic(err)
	}
	id := c.Param("id")
	_, err = db.Exec(query, id)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "result": "Delete Done!"})
}
