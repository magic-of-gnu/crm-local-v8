import argparse

tmpl_repo_type = """type {repo_name} struct {{
	dbpool *pgxpool.Pool
}}
"""

tmpl_repo_type_init = """func {create_repo_name} (dbpool *pgxpool.Pool) *{repo_name} {{
	return &{repo_name}{{
		dbpool: dbpool,
	}}
}}"""

tmpl_get_all = """func (rr *{repo_name}) GetAll() ([]{model_name}, error) {{

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `SELECT ... FROM {table_name}`

	var items []{model_name}

	rows, err := rr.dbpool.Query(ctx, query)
	if err != nil {{
		return items, err
	}}
	defer rows.Close()

	for rows.Next() {{
		var item {model_name}
		err = rows.Scan(...)

		if err != nil {{
			return items, err
		}}

		items = append(items, item)
	}}

	if err = rows.Err(); err != nil {{
		return items, err
	}}

	return items, nil
}}
"""

tmpl_delete_by_id = '''
func (rr *{repo_name}) DeleteOneByID(uid uuid.UUID) error {{

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := "DELETE FROM {table_name} WHERE id = $1"

	row, err := rr.dbpool.Exec(ctx, query, uid)

	if err != nil {{
		fmt.Println("err: ", err)
		return err
	}}

	if row.RowsAffected() == 0 {{
		fmt.Println("err: data was not deleted")
		return fmt.Errorf("data was not deleted")
	}}

	return nil

}}'''

tmpl_create_one = '''func (rr *{repo_name}) CreateOne(...) (*{model_name}, error) {{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `INSERT INTO {table_name} (...) VALUES ()`

	var item {model_name}

	_, err := rr.dbpool.Exec(ctx, query,
		...
	)

	if err != nil {{
		fmt.Println("err: ", err)
		return &item, err
	}}

	item = {model_name}{{...}}

	return &item, nil

}}
'''

if __name__ == "__main__":

	_description = """E.g.
	python3 {0} --package-name repos --repo-name lectureCalendarPostgresRepo --create-repo-name NewLectureCalendarPostgresRepo
	--model-name models.LecturesCalendar --table-name lectures_calendar --fname-out lectures_calendar.go
	"""

	_description = """E.g.
	python3 {0} --package-name repos --repo-name lectureCalendarPostgresRepo --model-name models.LecturesCalendar --table-name lectures_calendar --fname-out lectures_calendar.go
	
"""

	parser = argparse.ArgumentParser(_description)
	parser.add_argument("--package-name", dest="package_name", help='Name of the package')
	parser.add_argument("--repo-name", dest="repo_name", help='Name of the repo object')
	parser.add_argument("--model-name", dest="model_name", help='Name of the model')
	parser.add_argument("--table-name", dest="table_name", help='Name of the table in db')
	parser.add_argument("--fname-out", dest="fname_out", help='Filename of output .go file')

	args = parser.parse_args()
	package_name = args.package_name
	repo_name = args.repo_name
	# create_repo_name = args.create_repo_name
	model_name = args.model_name
	table_name = args.table_name
	fname_out = args.fname_out

	repo_name_list = list(repo_name)
	repo_name_capital = repo_name_list[:]

	repo_name_list[0] = repo_name_list[0].lower()
	repo_name_capital[0] = repo_name_capital[0].upper()

	repo_name = ''.join(repo_name_list)
	create_repo_name = 'New' + ''.join(repo_name_capital)

	# package_name = "repos"
	# repo_name = 'lectureCalendarPostgresRepo'
	# create_repo_name = 'NewLectureCalendarPostgresRepo'
	# model_name = "models.LecturesCalendar"
	# table_name = "lectures_calendar"

	str_repo_type = tmpl_repo_type.format(repo_name=repo_name)
	str_repo_type_init = tmpl_repo_type_init.format(**{
		'create_repo_name': create_repo_name,
		'repo_name': repo_name,
	})
	str_get_all = tmpl_get_all.format(**{
		'repo_name': repo_name,
		'model_name': model_name,
		'table_name': table_name
	})
	str_delete_by_id = tmpl_delete_by_id.format(**{
		'table_name': table_name,
		'repo_name': repo_name,
		'table_name': table_name
	})

	str_create_one = tmpl_create_one.format(**{
		'repo_name': repo_name,
		'table_name': table_name,
		'model_name': model_name,
	})

	# print(str_repo_type)
	# print(str_repo_type_init)
	# print(str_get_all)
	# print(str_delete_by_id)
	# print(str_create_one)

	with open(fname_out, 'w') as f:
		f.write("\n".join([
			f"package {package_name}",
			"\n",
			str_repo_type,
			"\n",
			str_repo_type_init,
			"\n",
			str_get_all,
			"\n",
			str_create_one,
			"\n",
			str_delete_by_id
		]))


