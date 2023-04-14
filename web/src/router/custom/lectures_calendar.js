const routes = {
  path: '/lectures_calendar',
  name: 'LecturesCalendar List',
  redirect: '/lectures_calendar/list',
  children: [
    {
      path: '/lectures_calendar/create_one',
      name: 'lectures_calendar_create_one',
      component: () =>
        import('@/views/custom/views/lectures_calendar/LecturesCalendarCreateOneView.vue'),
      meta: {
        requiresAuth: true
      },
    },
    {
      path: '/lectures_calendar/list',
      name: 'lectures_calendar_list',
      component: () => 
        import('@/views/custom/views/lectures_calendar/LecturesCalendarListView.vue'),
      meta: {
        requiresAuth: true
      },
    },
  ],
}

export default routes;