
export default {
  path: '/invoices',
  name: 'Invoices',
  redirect: '/invoices/list',
  children: [
    {
      path: '/invoices/create_one',
      name: 'invoices_create_one',
      component: () =>
        import('@/views/custom/views/invoices/InvoicesCreateOneView.vue'),
      meta: {
        requiresAuth: true
      },
    },
    {
      path: '/invoices/list',
      name: 'invoices_list_all',
      component: () => import('@/views/custom/views/invoices/InvoicesListView.vue'),
      meta: {
        requiresAuth: true
      },
    },
  ],
}