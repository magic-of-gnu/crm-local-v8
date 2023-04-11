const routes = {
  path: '/users',
  name: 'UsersBase',
  redirect: '/users/list',
  children: [
    {
      path: '/users/list',
      name: 'users_list',
      component: () =>
        import('@/views/custom/views/users/UsersListView.vue'),
    },
    {
      path: '/users/create_one',
      name: 'users_create_one',
      component: () =>
        import('@/views/custom/views/users/UsersCreateOneView.vue'),
    }
  ],
}

export default routes;