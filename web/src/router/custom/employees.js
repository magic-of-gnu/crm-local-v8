const routes = {
  path: '/employees',
  name: 'Employees List',
  redirect: '/employee/list',
  children: [
    {
      path: '/employee/create_one',
      name: 'employees_create_one',
      component: () =>
        import('@/views/custom/views/employees/EmployeesCreateOneView.vue'),
      meta: {
        requiresAuth: true
      },
    },
    {
      path: '/employee/list',
      name: 'employees_list',
      component: () => 
        import('@/views/custom/views/employees/EmployeesListView.vue'),
      meta: {
        requiresAuth: true
      },
    },
  ],
}

export default routes;