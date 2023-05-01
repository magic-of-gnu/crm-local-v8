
export default {
  path: '/payment_statuses',
  name: 'Payment Statuses',
  redirect: '/payment_statuses/list',
  children: [
    {
      path: '/payment_statuses/create_one',
      name: 'payment_statuses_create_one',
      component: () =>
        import('@/views/custom/views/payment_statuses/PaymentStatusesCreateOneView.vue'),
      meta: {
        requiresAuth: true
      },
    },
    {
      path: '/payment_statuses/list',
      name: 'payment_statuses_list_all',
      component: () => import('@/views/custom/views/payment_statuses/PaymentStatusesListView.vue'),
      meta: {
        requiresAuth: true
      },
    },
  ],
}