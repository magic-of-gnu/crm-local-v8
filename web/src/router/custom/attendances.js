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
    },
    {
      path: '/attendances/create_one',
      name: 'attendances_create_one',
      component: () =>
        import('@/views/custom/views/attendances/AttendancesCreateOneView.vue'),
    }
  ],
}

export default routes;