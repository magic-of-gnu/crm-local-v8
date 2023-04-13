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

repos_package="repos"
services_package="services"

# base_name="users"
# model_name="models.User"

base_name="tokens"
model_name="models.Token"
fname_base=tokens.go

base_name_lower=${base_name,}
base_name_upper=${base_name^}
table_name=${base_name}
repo_postgres_name="${base_name_upper}PostgresRepo"
repo_name="${base_name_upper}Repo"
service_name="${base_name_lower}Service"
handler_name="${base_name_upper}"
service_name_capital="${base_name_upper}Service"
create_one_request_model="models.${base_name_upper}CreateOneRequest"
delete_by_id_request_model="models.${baes_name_upper}DeleteByIDRequest"

base_dir="/home/danik/projects/danik/mycrm/local_v8/server/internal"


fname_repo=${base_dir}/repos/${fname_base}
fname_service=${base_dir}/services/${fname_base}
fname_handler=${base_dir}/handlers/${fname_base}

python3 server_generators/repos_generator.py --package-name ${repos_package} --repo-name ${repo_postgres_name} --model-name ${model_name} --table-name ${table_name} --fname-out ${fname_repo}
python3 server_generators/service_generator.py --service-name ${service_name} --repo-name ${repo_name} --repo-package ${repos_package} --service-package ${services_package} --model-name ${model_name} --fname-out ${fname_service}
python3 server_generators/handlers_generator.py --name ${handler_name} --service-name ${service_name_capital} --create-one-request-name ${create_one_request_model} --delete-by-id-request-name ${delete_by_id_request_model} --fname-out ${fname_handler}