
export default {
  path: '/centres',
  name: 'Centres List',
  redirect: '/centres/list',
  children: [
    {
      path: '/centres/create_one',
      name: 'centre_create_one',
      component: () =>
        import('@/views/custom/views/centres/CentresCreateOneView.vue'),
      meta: {
        requiresAuth: true
      },
    },
    {
      path: '/centres/list',
      name: 'List Centres',
      component: () => import('@/views/custom/views/centres/CentresListView.vue'),
      meta: {
        requiresAuth: true
      },
    },
  ],
}