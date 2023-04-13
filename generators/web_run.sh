#!bin/bash

column_labels="ID,First Name,Last Name,Username,Last Login,Is Admin,User Type,Created At,Updated At"
column_values="id,first_name,last_name,username,last_login,is_admin,user_type,created_at,updated_at"

base_name="users"
base_name_upper=${base_name^}
base_name_lower=${base_name,}
route_base="/${base_name_lower}"
route_base_name="${base_name_upper}Base"

base_url="http://malcorp.test/api/server"


route_list="/${base_name_lower}/list"
route_list_name="${base_name_lower}_list"
route_list_view="/views/custom/views/${base_name_lower}/${base_name_upper}ListView.vue"
route_list_view_relative="@${routes_list_view}"

route_redirect=${route_list}

route_create_one="/${base_name_lower}/create_one"
route_create_one_name="${base_name_lower}_create_one"
route_create_one_view="/views/custom/views/${base_name_lower}/${base_name_upper}CreateOneView.vue"
route_create_one_view_relative="@${routes_create_one_view}"

route_delete_by_id="/${base_name_lower}"

base_dir="/home/danik/projects/danik/mycrm/local_v8/web/src"
fname_routes="${base_dir}/router/custom/${base_name_lower}.js"
fname_hooks_methods="${base_dir}/views/custom/hooks/${base_name_lower}/methods.js"
fname_hooks_methods_relative="@/views/custom/hooks/${base_name_lower}/methods.js"
fname_list_view="${base_dir}/views/custom/views/${base_name_lower}/${base_name_upper}ListView.vue"

mkdir -p ${fname_routes%\/*}
mkdir -p ${fname_hooks_methods%\/*}
mkdir -p ${fname_list_view%\/*}

python3 web_generators/routes_generator.py --base-route "${route_base}" --base-route-name "${route_base_name}" --base-redirect-route ${route_redirect} --routes "${route_list},${route_create_one}" --routes-names "${route_list_name},${route_create_one_name}" --routes-view-paths "@${route_list_view},@${route_create_one_view}" --fname-out ${fname_routes}
python3 web_generators/hooks_methods_generator.py --base-url "${base_url}" --route-get-all "${route_list}" --route-create-one "${route_create_one}" --route-delete-by-id "${route_delete_by_id}" --fname-out ${fname_hooks_methods}
python3 web_generators/views_list_view.py --route-create-one-name ${route_create_one_name} --column-labels "${column_labels}" --column-values "${column_values}" --fname-methods ${fname_hooks_methods_relative} --fname-out ${fname_list_view}