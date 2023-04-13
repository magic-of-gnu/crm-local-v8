# Simple Generators

Consists of 2 directory:
- `server_generators`
- `web_generators`

## Server Generators


### Preresequites
- generated models


Server generators can generate:
- `repo`
- `service`
- `handlers`

After generating repo, services and handlers, you have to manually fix them:
- input and output arguments of some methods in `repos` and `services`
- more advanced queries in `CreateOne` and `ListAll` queries
- add interfaces for `repo.go` and `service.go`
- add new services and repos into the `App`
- add handler routes into `handlers`
- moves files to corresponding directories


For usage, please view `run.sh`:
```bash
repos_package="repos"
repo_postgres_name="AttendancesPostgresRepo"
repo_name="AttendancesRepo"
service_name="attendancesService"
services_package="services"
table_name="attendances"
model_name="models.Attenance"
handler_name="Attendances"
service_name_capital="AttendancesService"
create_one_request_model="models.AttendancesCreateOneRequest"
delete_by_id_request_model="models.AttendancesDeleteByIDRequest"

base_dir="server/internals"
fname_repo=${base_dir}/repos/attendances.go
fname_service=${base_dir}/services/attendances.go
fname_handler=${base_dir}/handlers/attendances.go

python3 repos_generator.py --package-name ${repos_package} --repo-name ${repo_postgres_name} --model-name ${model_name} --table-name ${table_name} --fname-out ${fname_repo}

python3 service_generator.py --service-name ${service_name} --repo-name ${repo_name} --repo-package ${repos_package} --service-package ${services_package} --model-name ${model_name} --fname-out ${fname_service}

python3 handlers_generator.py --name ${handler_name} --service-name ${service_name_capital} --create-one-request-name ${create_one_request_model} --delete-by-id-request-name ${delete_by_id_request_model} --fname-out ${fname_handler}
```

The code should look like this:
```bash
python3 repos_generator.py --package-name repos --repo-name AttendnacesPostgresRepo --model-name models.LecturesCalendar --table-name lectures_calendar --fname-out lectures_calendar.go

python3 service_generator.py --service-name attendancesService --repo-name AttendancesRepo --repo-package repo --service-package services --model-name models.Attendance --fname-out attendances.go

python3 handlers_generator.py --name Attendances --service-name AttendancesServices --create-one-request-name models.AttendancesCreateOneRequest --delete-by-id-request-name models.AttendancesByIDRequest --fname-out attendances.go
```


Usage for web generators:
```bash
base_name="attendances"
base_name_upper=${base_name^}
base_name_lower=${base_name,}
route_base="/${base_name_lower}"
route_base_name="${base_name_upper}Base"
route_redirect=${route_list}

base_url="http://malcorp.test/api/server"


route_list="/${base_name_lower}/list"
route_list_name="${base_name_lowre}_list"
route_list_view="/views/custom/views/${base_name_lower}/${base_name_upper}ListView.vue"
route_list_view_relative="@${routes_list_view}"

route_create_one="/${base_name_lower}/create_one"
route_create_one_name="${base_name_lower}_create_one"
route_create_one_view="/views/custom/views/${base_name_lower}/${base_name_upper}CreateOneView.vue"
route_create_one_view_relative="@${routes_create_one_view}"

route_delete_by_id="/${base_name_lower}"

column_labels="ID,Value,Name,Description,Created At,Updated At"
column_values="id,value,name,description,created_at,updated_at"


base_dir="/home/danik/projects/danik/mycrm/local_v8/web/src"
fname_routes="${base_dir}/router/custom/${base_name_lower}.js"
fname_hooks_methods="${base_dir}/views/custom/hooks/${base_name_lower}/methods.js"
fname_list_view="${base_dir}/views/custom/views/${base_name_upper}ListView.vue"


python3 routes_generator.py --base-route "${route_baes}" --base-route-name "${route_base_name}" --base-redirect-route ${route_redirect} --routes "${route_list},${route_create_one}" --routes-names "${route_list_name},${route_create_one_name}" --routes-view-paths "${route_list_view},${route_create_one_view}" --fname-out ${fname_routes}

python3 hooks_methods_generator.py --base-url "${base_url}" --route-get-all "${route_list}" --route-create-one "${route_create_one}" --route-delete-by-id "${route_delete_by_id}" --fname-out ${fname_hooks_methods}

python3 views_list_view.py --route-create-one-name ${route_create_one_name} --column-labels "${column_labels}" --column-values "${column_values}" --fname-methods ${fname_hooks_methods} --fname-out ${fname_list_view}
```

Usage for web generators:
```bash

python3 routes_generator.py --base-route "/attendances" --base-route-name AttendancesBase --base-redirect-route "/attendances/list" --routes "/attendances/list,/attendances/create_one" --routes-names "attendances_list" --routes-view-paths "@/views/custom/views/attendances/AttendancesListView.vue,@/views/custom/views/attendances/AttendancesCreateOneView.vue" --fname-out attendances.js

python3 hooks_methods_generator.py --base-url "http://malcorp.test/api/server" --route-get-all "/attendances/list" --route-create-one "/attendances/create_one" --route-delete-by-id "/attendances" --fname-out methods.

python3 views_list_view.py --route-create-one-name attendances_create_one --column-labels "ID,Value,Name,Description,Created At,Updated At" --column-values 'id,values,name,description,created_at,updated_at' --fname-methods @/views/custom/hooks/attendanceValues/methods.js --fname-out AttendancesListView.vue
```