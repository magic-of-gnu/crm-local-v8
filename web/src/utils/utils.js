function epochToDatetime(epochTime) {
  const date = new Date(epochTime * 1000);
  return new Intl.DateTimeFormat('default', {dateStyle: 'long'}).format(date)
}

function jsEpochToGoEpoch(jsEpochTime) {
  return Math.round(jsEpochTime / 1000) // milliseconds -> seconds
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

function stringOperatorsAsOptions() {
  return [
    {
      label: "in",
      value: "sin",
    },
    {
      label: "=",
      value: "s==",
    },
    {
      label: "!=",
      value: "s!=",
    }
  ]
}

function numberOperatorsAsOptions() {
  return [
    {
      label: "=",
      value: "neq",
    },
    {
      label: "!=",
      value: "nneq",
    },
    {
      label: ">",
      value: "ngt",
    },
    {
      label: "<",
      value: "nlt",
    },
    {
      label: ">=",
      value: "nge",
    },
    {
      label: "<=",
      value: "nle",
    }
  ]
}

export {
  epochToDatetime,
  jsEpochToGoEpoch,
  numberOperators,
  numberOperatorsAsOptions,
  showUUID,
  sortList,
  stringOperators,
  stringOperatorsAsOptions,
  utcTimeToDate,
  utcTimeToLocaleTime,
  utcTimeToTimeHoursMinutes,
  utcTimeToUpdateTime,
}