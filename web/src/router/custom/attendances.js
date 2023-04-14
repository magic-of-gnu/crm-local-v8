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
      path: '/attendances/create_one',
      name: 'attendances_create_one',
      component: () =>
        import('@/views/custom/views/attendances/AttendancesCreateOneView.vue'),
      meta: {
        requiresAuth: true
      },
    }
  ],
}

export default routes;