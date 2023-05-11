package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

func GetAllStudents(c *gin.Context) {

	fmt.Println("handler students get all")

	students, err := App.StudentsService.GetAll()

	if err != nil {

		fmt.Println("error here")
		c.JSON(http.StatusOK, gin.H{
			"title": "Students List",
			"error": err.Error(),
		})

		return
	}

	data := make(map[string]interface{})
	students_data := make([]map[string]string, len(students))

	for ii, item := range students {
		students_data[ii] = map[string]string{
			"id":         item.ID.String(),
			"first_name": item.FirstName,
			"last_name":  item.LastName,
			"username":   item.Username,
			"created_at": item.CreatedAt.Format(time.DateTime),
			"updated_at": item.UpdatedAt.Format(time.DateTime),
		}
	}
	data["data"] = students_data

	c.JSON(http.StatusOK, gin.H{
		"title": "Students List",
		"data":  data,
	})
}

func PostStudentsCreateOne(c *gin.Context) {

	var student *models.Student

	if err := c.ShouldBind(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"title": "Student",
			"error": err.Error(),
			"toasts": []map[string]string{
				{
					"content": "Error: incorrect request",
					"color":   "warning",
				},
			},
		})
		return
	}

	student, err := App.StudentsService.CreateOne(
		student.FirstName,
		student.LastName,
		student.Username,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title": "Students",
			"error": err.Error(),
			"toasts": []map[string]string{
				{
					"content": "Error: internal error",
					"color":   "warning",
				},
			},
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"title": "Students",
		"toasts": []map[string]string{
			{
				"content": "Created",
				"color":   "success",
			},
		},
		"data": student,
	})
}

func DeleteStudentsByID(c *gin.Context) {
	title := "Students Delete"

	id_str := c.Param("id")
	uid, err := uuid.Parse(id_str)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"title":   title,
			"error":   err.Error(),
			"message": "error during bidning data",
			"toasts": []map[string]string{
				{
					"content": "Error: incorrect request",
					"color":   "warning",
				},
			},
		})
		return
	}

	err = App.StudentsService.DeleteOneByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"message": "error delete in db",
			"error":   err.Error(),
			"toasts": []map[string]string{
				{
					"content": "Error: internal error",
					"color":   "warning",
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":   title,
		"message": "success",
		"toasts": []map[string]string{
			{
				"content": "Deleted",
				"color":   "success",
			},
		},
	})
}

func GetOneStudentsByID(c *gin.Context) {
	title := "Students GetOne"

	id_str := c.Param("id")
	uid, err := uuid.Parse(id_str)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"title":   title,
			"error":   err.Error(),
			"message": "error during bidning data",
			"toasts": []map[string]string{
				{
					"content": "Error: incorrect request",
					"color":   "warning",
				},
			},
		})
		return
	}

	item, err := App.StudentsService.GetOneByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"message": "error delete in db",
			"error":   err.Error(),
			"toasts": []map[string]string{
				{
					"content": "Error: internal error",
					"color":   "warning",
				},
			},
		})
		return
	}

	data := make(map[string]interface{})
	data["data"] = item

	c.JSON(http.StatusOK, gin.H{
		"title":   title,
		"message": "success",
		"toasts": []map[string]string{
			{
				"content": "Created",
				"color":   "success",
			},
		},
		"data": data,
	})
}

func PatchStudentsByID(c *gin.Context) {
	title := "Rooms Patch"

	// get the item from db and update it by the request
	id_str := c.Param("id")
	uid, err := uuid.Parse(id_str)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"error":   err.Error(),
			"message": "error during bidning data",
			"toasts": []map[string]string{
				{
					"content": "Error: incorrect request",
					"color":   "warning",
				},
			},
		})
		return
	}

	req, err := App.StudentsService.GetOneByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"error":   err.Error(),
			"message": "error during bidning data",
			"toasts": []map[string]string{
				{
					"content": "Error: incorrect request",
					"color":   "warning",
				},
			},
		})
		return
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"error":   err.Error(),
			"message": "error during bidning data",
			"toasts": []map[string]string{
				{
					"content": "Error: incorrect request",
					"color":   "warning",
				},
			},
		})
		return
	}

	err = App.StudentsService.UpdateOneByID(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"title":   title,
			"message": "error delete in db",
			"error":   err.Error(),
			"toasts": []map[string]string{
				{
					"content": "Error: internal error",
					"color":   "warning",
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":   title,
		"message": "success",
		"toasts": []map[string]string{
			{
				"content": "Created",
				"color":   "success",
			},
		},
	})
}
