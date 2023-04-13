import argparse

tmpl_column_labels = '                <CTableHeaderCell scope="col">{val}</CTableHeaderCell>'
tmpl_column_values = '                <CTableDataCell>{{{{ item.{val} }}}}</CTableDataCell>'

tmpl = '''
<template>
  <CRow>
    <CCol :xs="12">
      <CCard class="mb-4">
        <CCardHeader>
          <CRow class="justify-content-between">
            <CCol>
              <strong>Vue Progress</strong> <small>Basic example</small>
            </CCol>
            <CCol class="align-self-end">
              <router-link
                :to="{{ name: '{route_create_one_name}' }}"
                component="CButton"
                color="primary"
                disabled
                >Create One
              </router-link>
            </CCol>
          </CRow>
        </CCardHeader>
        <CCardBody>
          <CTable hover>
            <CTableHead>
              <CTableRow>
{table_column_labels}
              </CTableRow>
            </CTableHead>
            <CTableBody>
              <CTableRow v-for="item in itemsList" :key="item.id">
{table_column_values}
              </CTableRow>
            </CTableBody>
          </CTable>
        </CCardBody>
      </CCard>
    </CCol>
  </CRow>
</template>

<script setup>
import m from '{fname_methods}'
import {{ ref, onMounted }} from 'vue'

const itemsList = ref('')

onMounted(() => {{
  m.getAllList().then((d) => {{
    itemsList.value = d.data.data
  }})
}})
</script>
'''

if __name__ == "__main__":
  _description = '''Example with Operations Column

  python3 views_list_view.py --route-create-one-name attendances_create_one --column-labels "ID,Value,Name,Description,Created At,Updated At" --column-values 'id,values,name,description,created_at,updated_at' --fname-methods @/views/custom/hooks/attendanceValues/methods.js --fname-out AttendancesListView.vue

  Example excluding Operations Column

  python3 views_list_view.py --route-create-one-name attendances_create_one --column-labels "ID,Value,Name,Description,Created At,Updated At" --column-values 'id,values,name,description,created_at,updated_at' --fname-methods @/views/custom/hooks/attendanceValues/methods.js --fname-out AttendancesListView.vue --exclude-operations
  '''

  parser = argparse.ArgumentParser(_description)
  parser.add_argument('--route-create-one-name', help="Name of the create_one route")
  parser.add_argument('--column-labels', help="Labels of the Columns")
  parser.add_argument('--exclude-operations', help="Whether to include operations", action="store_true")
  parser.add_argument('--column-values', help="Values of the Columns")
  parser.add_argument('--fname-methods', help="Filename of the methods file")
  parser.add_argument('--fname-out', help="Filename of the output .vue file")

  args = parser.parse_args()
  route_create_one_name = args.route_create_one_name
  column_labels = args.column_labels
  column_values = args.column_values
  fname_methods = args.fname_methods
  fname_out = args.fname_out

  exclude_operations = args.exclude_operations


# route_create_one_name = 'attendances_create_one'
# column_labels = "ID,Value,Name,Description,Created At,Updated At"
# column_values = 'id,value,name,description,created_at,updated_at'
# fname_methods = '@/views/custom/hooks/attendanceValues/methods.js'



  str_table_column_labels = [tmpl_column_labels.format(val=item.strip()) for item in column_labels.split(",")]
  str_table_column_values = [tmpl_column_values.format(val=item.strip()) for item in column_values.split(",")]

  if not exclude_operations:
    str_table_column_labels.append(tmpl_column_labels.format(val="Operations"))

  str_table_column_labels = '\n'.join(str_table_column_labels)
  str_table_column_values = '\n'.join(str_table_column_values)

  str_all = tmpl.format(**{
    'route_create_one_name': route_create_one_name,
    'table_column_labels': str_table_column_labels,
    'table_column_values': str_table_column_values,
    'fname_methods': fname_methods,
  })

  with open(fname_out, 'w') as f:
    f.write(str_all)
