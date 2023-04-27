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

function utcTimeToUpdateTime(time_string) {
  const date = new Date(time_string);
  const [yyyy, month, day] = [date.getFullYear(), date.toLocaleString('default', { month: 'long' }), date.getDate()]
  const [hh, mm, ss] = [date.getHours().toString().padStart(2, '0'), date.getMinutes().toString().padStart(2, '0'), date.getSeconds().toString().padStart(2, '0')]
  return `${yyyy} ${month} ${day} ${hh}:${mm}:${ss}`
}

function showUUID(uuid_str, l=4) {
  return uuid_str.slice(0, l)
}

function sortList(data, col, asc) {
  if (asc[col]) {
    data.sort((x, y) => (x[col] > y[col] ? -1 : 1));
  } else {
    data.sort((x, y) => (x[col] < y[col] ? -1 : 1));
  }
  console.log("before asc[col]: ", asc[col])
  asc[col] = !asc[col]
  console.log("after asc[col]: ", asc[col])
}

function stringOperators() {return ["in", "=", "!="]}
function numberOperators() {return ["=", "!=", ">", "<", ">=", "<="]}

export {
  epochToDatetime,
  utcTimeToDate,
  utcTimeToLocaleTime,
  utcTimeToTimeHoursMinutes,
  utcTimeToUpdateTime,
  sortList,
  stringOperators,
  numberOperators,
  showUUID,
}