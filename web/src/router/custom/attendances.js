const routes = {
  path: '/attendances',
  name: 'AttendancesBase',
  redirect: '/attendances/list',
  children: [
    {
      path: '/attendances/list',
      name: 'attendances_list',
      component: () =>
        import('@/views/custom/views/attendances/AttendancesListView.vue'),
      meta: {
        requiresAuth: true
      },
    },
    {
      path: '/attendances/create_by_course_name',
      name: 'attendances_create_by_course_name',
      component: () =>
        import('@/views/custom/views/attendances/AttendancesCreateByCourseName.vue'),
      meta: {
        requiresAuth: true
      },
    },
    {
      path: '/attendances/create_one/:id',
      name: 'attendances_create_one_id',
      component: () =>
        import('@/views/custom/views/attendances/AttendancesCreateOneView.vue'),
      meta: {
        requiresAuth: true
      },
    }
  ],
}

export default routes;