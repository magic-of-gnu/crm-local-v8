import argparse

tmpl_service_type = """type {service_name} struct {{
	{repo_name} {repo_package}.{repo_name}
}}

func New{service_name_capital}(repo {repo_package}.{repo_name}) *{service_name} {{
	return &{service_name}{{
		{repo_name}: repo,
	}}
}}"""


tmpl_get_all = """func (ss *{service_name}) GetAll() ([]{model_name}, error) {{
	return ss.{repo_name}.GetAll()
}}"""

tmpl_delete_by_id = """func (ss *{service_name}) DeleteOneByID(uid uuid.UUID) error {{

	err := ss.{repo_name}.DeleteOneByID(uid)
	if err != nil {{
		return err
	}}

	return nil
}}"""

tmpl_create_one = """func (ss *{service_name}) CreateOne(...) (*{model_name}, error) {{
	var item *{model_name}

	uid, err := uuid.NewRandom()
	if err != nil {{
		return item, err
	}}

	time_now := time.Now()

	item, err = ss.{repo_name}.CreateOne(...)
	if err != nil {{
		return item, err
	}}

	return item, nil
}}
"""

if __name__ == "__main__":
  _description = '''python3 service_generator.py --service-name attendancesService --repo-name AttendancesRepo --repo-package repos --service-package services --model-name models.Attendance --fname-out attendances.go'''

  parser = argparse.ArgumentParser(_description)
  parser.add_argument('--service-name', dest="service_name", help="Name of the Service")
  parser.add_argument('--repo-name', dest="repo_name", help="Name of the Repo")
  parser.add_argument('--repo-package', dest="repo_package", help="Name of the Repos package")
  parser.add_argument('--service-package', dest="service_package", help="Name of the Service package")
  parser.add_argument("--model-name", dest="model_name", help="Name of the Model")
  parser.add_argument("--fname-out", dest="fname_out", help="Name of the output .go file")

  args = parser.parse_args()
  service_name = list(args.service_name)
  repo_name = args.repo_name
  repo_package = args.repo_package
  service_package = args.service_package
  model_name = args.model_name
  fname_out = args.fname_out

  service_name[0] = service_name[0].lower()
  service_name_capital = service_name[:]
  service_name_capital[0] = service_name_capital[0].upper()

  service_name = ''.join(service_name)
  service_name_capital = ''.join(service_name_capital)

  # service_name = 'studentCourses'
  # service_name_capital = 'StudentCourses'
  # repo_name = 'StudentCoursesRepo'
  # repo_package = 'repos'
  # model_name = "models.StudentCourses"

  str_service_type = tmpl_service_type.format(**{
    'service_name': service_name,
    'repo_name': repo_name,
    'repo_package': repo_package,
    'service_name_capital': service_name_capital,
  })

  str_get_all = tmpl_get_all.format(**{
    'service_name': service_name,
    'model_name': model_name,
    'repo_name': repo_name
  })

  str_delete_by_id = tmpl_delete_by_id.format(**{
    'service_name': service_name,
    'repo_name': repo_name,
  })

  str_create_one = tmpl_create_one.format(**{
    'service_name': service_name,
    'model_name': model_name,
    'repo_name': repo_name
  })

  with open(fname_out, 'w') as f:
    f.write("\n".join([
      f"package {service_package}\n",
      str_service_type,
      "\n",
      str_get_all,
      "\n",
      str_create_one,
      "\n",
      str_delete_by_id
    ]))
