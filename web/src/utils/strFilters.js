function strFilterIn(data, col, val) {
  return data.filter((row) => {
    return row[col].toLowerCase().includes(val.toLowerCase())
  })
}

function strFilterEqual(data, col, val) {
  return data.filter((row) => {
    return row[col].toLowerCase() === val.toLowerCase()
  })
}

function strFilterNotEqual(data, col, val) {
  return data.filter((row) => {
    return row[col].toLowerCase() !== val.toLowerCase()
  })
}

export {
  strFilterIn,
  strFilterEqual,
  strFilterNotEqual
}