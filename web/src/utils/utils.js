function epochToDatetime(epochTime) {
  const date = new Date(epochTime * 1000);
  return new Intl.DateTimeFormat('default', {dateStyle: 'long'}).format(date)
}

function utcTimeToDate(time_string) {
  const date = new Date(time_string);
  return date.toDateString()
}

function utcTimeToLocaleTime(time_string) {
  const date = new Date(time_string);
  return date.toLocaleTimeString()
}

function utcTimeToTimeHoursMinutes(time_string) {
  const date = new Date(time_string);
  return date.getHours().toString().padStart(2, '0') + ":" + date.getMinutes().toString().padStart(2, '0')
}

export {
  epochToDatetime,
  utcTimeToDate,
  utcTimeToLocaleTime,
  utcTimeToTimeHoursMinutes,
}