const routes = {
  path: '/attendance_values',
  name: 'AttendanceValues List',
  redirect: '/attendance_values/list',
  children: [
    {
      path: '/attendance_values/create_one',
      name: 'attendance_values_create_one',
      component: () =>
        import('@/views/custom/views/attendance_values/AttendanceValuesCreateOneView.vue'),
      meta: {
        requiresAuth: true
      },
    },
    {
      path: '/attendance_values/list',
      name: 'attendance_values_list',
      component: () => 
        import('@/views/custom/views/attendance_values/AttendanceValuesListView.vue'),
      meta: {
        requiresAuth: true
      },
    },
  ],
}

export default routes;