import argparse


router_template = '''const routes = {{
  path: '{base_route}',
  name: '{base_route_name}',
  redirect: '{base_redirect_route}',
  children: [
{child_routes}
  ],
}}

export default routes;'''

child_route_template = '''    {{
      path: '{route_path}',
      name: '{route_name}',
      component: () =>
        import('{route_view_path}'),
    }}'''


if __name__ == "__main__":
  _description = '''Example:

  python3 routes_generator.py --base-route "/attendances" --base-route-name AttendancesBase --base-redirect-route "/attendances/list" --routes "/attendances/list,/attendances/create_one" --routes-names "attendances_list" --routes-view-paths "@/views/custom/views/attendances/AttendancesListView.vue,@/views/custom/views/attendances/AttendancesCreateOneView.vue" --fname-out attendances.js'''

  parser = argparse.ArgumentParser(_description)
  parser.add_argument('--base-route', dest='base_route', help='Base route path')
  parser.add_argument('--base-route-name', dest='base_route_name', help='Name of the Base route in the Named Routes')
  parser.add_argument('--base-redirect-route', dest="base_redirect_route", help='Route Path to redirect')
  parser.add_argument('--routes', help="List of routes in child routes, pass via ',' as seperator")
  parser.add_argument('--routes-names', dest="routes_names", help="List of routes'' names in child routes, pass via ',' as seperator")
  parser.add_argument('--routes-view-paths', dest="routes_view_paths", help="Path of the views, pass via ',' as seperator")
  parser.add_argument('--fname-out', dest="fname_out", help='Filename of the output .js file')

  args = parser.parse_args()

  base_route = args.base_route
  base_route_name = args.base_route_name
  base_redirect_route = args.base_redirect_route
  child_routes = args.routes.split(',')
  child_routes_names = args.routes_names.split(",")
  child_routes_view_paths = args.routes_view_paths.split(",")
  fname_out = args.fname_out

  child_routes_args = [{
    'route_path': route.strip(),
    'route_name': name.strip(),
    'route_view_path': view_path.strip()
  } for (route, name, view_path) in zip(child_routes, child_routes_names, child_routes_view_paths)]


  # base_route = "/student_courses"
  # base_rout_name = "StudentCoursesBase"
  # base_redirect_route = "/student_courses/list"
# routes = [{
#   "route_path": "/student_courses/list",
#   "route_name": "student_courses_list",
#   "route_view_path": "@/views/custom/views/student_courses/StudentCoursesListView.vue"
# }, {
#   "route_path": "/student_courses/create_one",
#   "route_name": "student_courses_create_one",
#   "route_view_path": "@/views/custom/views/student_courses/StudentCoursesCreateOneView.vue"
# }]



  child_routes_generated = ",\n".join([child_route_template.format(**item) for item in child_routes_args])
  # print(child_routes_generated)

  routes_generated = router_template.format(
    base_route=base_route,
    base_route_name=base_route_name,
    base_redirect_route=base_redirect_route,
    child_routes=child_routes_generated
  )

  # print("routes_generated: ", routes_generated)
  with open(fname_out, 'w') as f:
    f.write(routes_generated)