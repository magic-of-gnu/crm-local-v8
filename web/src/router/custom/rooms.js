const routes = {
  path: '/rooms',
  name: 'Rooms List',
  redirect: '/rooms/list',
  children: [
    {
      path: '/rooms/create_one',
      name: 'rooms_create_one',
      component: () =>
        import('@/views/custom/views/rooms/RoomsCreateOneView.vue'),
      meta: {
        requiresAuth: true
      },
    },
    {
      path: '/rooms/list',
      name: 'rooms_list',
      component: () => 
        import('@/views/custom/views/rooms/RoomsListView.vue'),
      meta: {
        requiresAuth: true
      },
    },
  ],
}

export default routes;