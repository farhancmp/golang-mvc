package latihancontroller

import (
	"latihan/app/models"
	"latihan/app/models/structs"
	"latihan/app/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	db, sql, err := models.DbConnect()
	if err != nil {
		log.Println(err.Error())
	}
	defer sql.Close()

	// dengan map string
	result := []map[string]interface{}{}

	// dengan map string
	db.Table("mahasiswa").Select("*").Scan(&result)

	// log.Println(result)
	for i := 0; i < len(result); i++ {
		result[i]["id_mahasiswa"] = result[i]["id"]
		result[i]["nama_mahasiswa"] = result[i]["name"]
		delete(result[i], "id")
		delete(result[i], "name")
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   result,
	})
}

func Post(c *gin.Context) {
	var mahasiswa structs.Mahasiswa

	err := c.ShouldBind(&mahasiswa)
	if err != nil {
		log.Println(err.Error())
	}

	db, sql, err := models.DbConnect()
	if err != nil {
		log.Println(err.Error())
	}

	defer sql.Close()

	db.Debug().Table("mahasiswa").Create(&mahasiswa)

	c.JSON(200, gin.H{
		"status": true,
		"data":   mahasiswa,
	})
}

func Update(c *gin.Context) {
	var mahasiswa structs.Mahasiswa

	id := c.Param("id")

	err := c.ShouldBind(&mahasiswa)
	if err != nil {
		c.JSON(404, gin.H{
			"status": false,
			"data":   err.Error(),
		})
		return
	}

	db, sql, err := models.DbConnect()
	if err != nil {
		log.Println(err.Error())
	}

	defer sql.Close()
	query := db.Debug().Table("mahasiswa").Where("idd=?", id).Updates(&mahasiswa).Scan(&mahasiswa)
	if queryError := query.Error; queryError != nil {
		utils.ResponseHandler(c, 500, false, queryError.Error(), []interface{}{})
		return
	}

	if query.RowsAffected == 0 {
		utils.ResponseHandler(c, 404, false, "Resource not found", []interface{}{})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   mahasiswa,
	})
}

func Delete(c *gin.Context) {
	var mahasiswa structs.Mahasiswa

	id := c.Param("id")

	db, sql, err := models.DbConnect()
	if err != nil {
		log.Println(err.Error())
	}

	defer sql.Close()
	query := db.Debug().Table("mahasiswa").Where("id=?", id).Delete(&mahasiswa)
	if queryError := query.Error; queryError != nil {
		utils.ResponseHandler(c, 500, false, queryError.Error(), []interface{}{})
		return
	}

	if query.RowsAffected == 0 {
		utils.ResponseHandler(c, 404, false, "Resource not found", []interface{}{})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   []map[string]interface{}{},
	})
}
