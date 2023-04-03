const routes = {
  path: '/students',
  name: 'Students List',
  redirect: '/students/list',
  children: [
    {
      path: '/students/create_one',
      name: 'students_create_one',
      component: () =>
        import('@/views/custom/views/students/StudentsCreateOneView.vue'),
    },
    {
      path: '/students/list',
      name: 'students_list',
      component: () => 
        import('@/views/custom/views/students/StudentsListView.vue'),
    },
  ],
}

export default routes;