package controllers

import (
	"net/http"

	"github.com/muhajirmuhajir/myapi/structs"
	"github.com/gin-gonic/gin"
)

// get one data by id
func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&person).Error

	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// get all data
func (idb *InDB) GetPersons(c *gin.Context) {
	var (
		persons []structs.Person
		result  gin.H
	)

	idb.DB.Find(&persons)
	if len(persons) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": persons,
			"count":  len(persons),
		}
	}

	c.JSON(http.StatusOK, result)
}

// create new data
func (idb *InDB) CreatePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")
	person.FirstName = firstName
	person.LastName = lastName

	idb.DB.Create(&person)
	result = gin.H{
		"message" : "User Created",
		"result": person,
	}

	c.JSON(http.StatusOK, result)
}

// update data by id
func (idb *InDB) UpdatePerson(c *gin.Context) {
	id := c.Query("id")
	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")
	var (
		person    structs.Person
		newPerson structs.Person
		result    gin.H
	)

	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	newPerson.FirstName = firstName
	newPerson.LastName = lastName
	err = idb.DB.Model(&person).Updates(newPerson).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "success update",
		}
	}

	c.JSON(http.StatusOK, result)
}

// delete data by id
func (idb *InDB) DeletePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.First(&person, id).Error

	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	err = idb.DB.Delete(&person).Error

	if err != nil {
		result = gin.H{
			"result": "deleted failed",
		}
	} else {
		result = gin.H{
			"result": "data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
