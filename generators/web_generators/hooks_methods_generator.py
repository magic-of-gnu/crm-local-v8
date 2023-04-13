import argparse

tmpl_get_all = '''
async function getAllList() {{
  return await axios({{
    method: 'get',
    url: '{route_get_all}',
    baseURL: '{base_url}',
    crossOrigin: true,
    responseType: 'json',
  }})
    .then((response) => {{
      if (response.status === 200) {{
        return response.data
      }}
    }})
    .catch(function (error) {{
    }})
    .finally(function () {{
      // always executed
    }})
}}'''

tmpl_create_one = '''
async function postCreateOne(data) {{
  return await axios({{
    method: 'post',
    url: '{route_create_one}',
    baseURL: '{base_url}',
    crossOrigin: true,
    responseType: 'json',
    data: data,
  }})
    .then((response) => {{
      if (response.status === 200) {{
        return response.data
      }}
    }})
    .catch(function (error) {{
    }})
    .finally(function () {{
      // always executed
    }})
}}'''

tmpl_delete_by_id = '''
async function postDeleteByID(data) {{
  return await axios({{
    method: 'delete',
    url: '{route_delete}',
    baseURL: '{base_url}',
    crossOrigin: true,
    responseType: 'json',
    data: data,
  }})
    .then((response) => {{
      if (response.status === 200) {{
        return response.data
      }}
    }})
    .catch(function (error) {{
      console.log('error: ', error)
    }})
    .finally(function () {{
      // always executed
    }})
}}
'''

tmpl_export = '''
export default {
  getAllList,
  postCreateOne,
  postDeleteByID,
}
'''

if __name__ == "__main__":
  _description = '''Usage:

python3 hooks_methods_generator.py --base-url "http://malcorp.test/api/server" --route-get-all "/attendances/list" --route-create-one "/attendances/create_one" --route-delete-by-id "/attendances" --fname-out methods.
  '''

  parser = argparse.ArgumentParser(_description)
  parser.add_argument("--base-url", dest="base_url", help="Base URL path")
  parser.add_argument("--route-get-all", dest="route_get_all", help="Route path to get all")
  parser.add_argument("--route-create-one", dest="route_create_one", help="Route path to create_one")
  parser.add_argument("--route-delete-by-id", dest="route_delete_by_id", help="Route path to delete_by_id")
  parser.add_argument('--fname-out', dest="fname_out", help="Filename of the output .js file")

  # fname_out = 'methods.js'
  # route_get_all = '/attendances/list'
  # base_url = 'http://malcorp.test/api/server'
  # route_create_one = '/attendances/create_one'
  # route_delete_by_id = '/attendances'

  args = parser.parse_args()
  base_url = args.base_url
  route_get_all = args.route_get_all
  route_create_one = args.route_create_one
  route_delete_by_id = args.route_delete_by_id
  fname_out = args.fname_out

  str_get_all = tmpl_get_all.format(**{
    'route_get_all': route_get_all,
    'base_url': base_url,
  })

  str_create_one = tmpl_create_one.format(**{
    'route_create_one': route_create_one,
    'base_url': base_url
  })

  str_delete_by_id = tmpl_delete_by_id.format(**{
    'route_delete': route_delete_by_id,
    'base_url': base_url
  })

  with open(fname_out, 'w') as f:
    f.write("\n".join([
      "import axios from 'axios'\n",
      str_get_all,
      str_create_one,
      str_delete_by_id,
      tmpl_export
    ]))


