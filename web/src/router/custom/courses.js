const routes = {
  path: '/courses',
  name: 'Courses List',
  redirect: '/courses/list',
  children: [
    {
      path: '/courses/create_one',
      name: 'courses_create_one',
      component: () =>
        import('@/views/custom/views/courses/CoursesCreateOneView.vue'),
    },
    {
      path: '/courses/list',
      name: 'courses_list',
      component: () => 
        import('@/views/custom/views/courses/CoursesListView.vue'),
    },
  ],
}

export default routes;