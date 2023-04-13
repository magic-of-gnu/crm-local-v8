import argparse

tmpl_get_all = '''func Get{name}All(c *gin.Context) {{
	title := "{name} List"

	items, err := App.{service_name}.GetAll()

	if err != nil {{

		c.JSON(http.StatusOK, gin.H{{
			"title":   title,
			"message": "error during getting data from postgres",
			"error":   err.Error(),
		}})

		return
	}}

	data := make(map[string]interface{{}})
	data["data"] = items

	c.JSON(http.StatusOK, gin.H{{
		"title": title,
		"data":  data,
	}})
}}
'''

tmpl_create_one = '''func Post{name}CreateOne(c *gin.Context) {{
	title := "{name} CreateOne"

	var req *{create_one_request_name}

	if err := c.ShouldBind(&req); err != nil {{
		c.JSON(http.StatusInternalServerError, gin.H{{
			"title":   title,
			"error":   err.Error(),
			"message": "error during bidning data",
		}})
		return
	}}

	err := App.Validator.Struct(req)
	if err != nil {{
		c.JSON(http.StatusInternalServerError, gin.H{{
			"title":   title,
			"message": "error during validation",
			"error":   err.Error(),
		}})
		fmt.Println("error during validation; err: ", err)
		return
	}}

	item, err := App.{service_name}.CreateOne(...)
	if err != nil {{
		c.JSON(http.StatusInternalServerError, gin.H{{
			"title":   title,
			"message": "error during validation",
			"error":   err.Error(),
		}})
		fmt.Println("error during validation; err: ", err)
		return
	}}

	c.JSON(http.StatusCreated, gin.H{{
		"title": title,
		"data":  item,
	}})
}}
'''

tmpl_delete_by_id = '''
func Delete{name}ByID(c *gin.Context) {{
	title := "{name} Delete"

	var req *{delete_by_id_request_name}

	if err := c.ShouldBind(&req); err != nil {{
		c.JSON(http.StatusInternalServerError, gin.H{{
			"title":   title,
			"error":   err.Error(),
			"message": "error during bidning data",
		}})
		return
	}}

	err := App.Validator.Struct(req)
	if err != nil {{
		c.JSON(http.StatusInternalServerError, gin.H{{
			"title":   title,
			"message": "error during validation",
			"error":   err.Error(),
		}})
		return
	}}

	err = App.{service_name}.DeleteOneByID(req.ID)
	if err != nil {{
		c.JSON(http.StatusInternalServerError, gin.H{{
			"title":   title,
			"message": "error delete in db",
			"error":   err.Error(),
		}})
		return
	}}

	c.JSON(http.StatusOK, gin.H{{
		"title":   title,
		"message": "success",
	}})
}}
'''

if __name__ == "__main__":
  _description = """python3 handlers_generator.py --name Attendances --service-name AttendancesServices --create-one-request-name models.AttendancesCreateOneRequest --delete-by-id-request-name models.AttendancesByIDRequest --fname-out attendances.go"""

  parser = argparse.ArgumentParser(_description)
  parser.add_argument("--name", help="Name of the Object")
  parser.add_argument("--service-name", dest="service_name", help="Name of the Service")
  parser.add_argument("--create-one-request-name", dest="create_one_request_name", help="Create One Request Name")
  parser.add_argument("--delete-by-id-request-name", dest="delete_by_id_request_name", help="Delete By ID Request Name")
  parser.add_argument("--fname-out", dest="fname_out", help="Filename of the output .go file")

  args = parser.parse_args()
  name = args.name
  service_name = args.service_name
  create_one_request_name = args.create_one_request_name
  delete_by_id_request_name = args.delete_by_id_request_name
  fname_out = args.fname_out

  # name = "Attendances"
  # service_name = "AttendancesServices"
  # create_one_request_name = "models.AttendancesCreateOneRequest"
  # delete_by_id_request_name = "models.AttendancesDeleteByIDRequest"

  str_delete_by_id = tmpl_delete_by_id.format(**{
    'name': name,
    'delete_by_id_request_name': delete_by_id_request_name,
    'service_name': service_name
  })

  str_create_one = tmpl_create_one.format(**{
    'name': name,
    'service_name': service_name,
    'create_one_request_name': create_one_request_name,
  })

  str_get_all = tmpl_get_all.format(**{
    'name': name,
    'service_name': service_name
  })

  with open(fname_out, 'w') as f:
    f.write("\n".join([
      "package handlers\n",
      str_get_all,
      "\n",
      str_create_one,
      "\n",
      str_delete_by_id
    ]))
