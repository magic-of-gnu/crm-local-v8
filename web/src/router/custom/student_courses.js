const routes = {
  path: '/student_courses',
  name: 'StudentCourses List',
  redirect: '/student_courses/list',
  children: [
    {
      path: '/student_courses/create_one',
      name: 'student_courses_create_one',
      component: () =>
        import('@/views/custom/views/student_courses/StudentCoursesCreateOneView.vue'),
      meta: {
        requiresAuth: true
      },
    },
    {
      path: '/student_courses/list',
      name: 'student_courses_list',
      component: () => 
        import('@/views/custom/views/student_courses/StudentCoursesListView.vue'),
      meta: {
        requiresAuth: true
      },
    },
  ],
}

export default routes;