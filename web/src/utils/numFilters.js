function numFilterEqual(data, col, val) {
  return data.filter((row) => {
    return row[col] === val
  })
}

function numFilterNotEqual(data, col, val) {
  return data.filter((row) => {
    return row[col] !== val
  })
}

function numFilterGt(data, col, val) {
  return data.filter((row) => {
    return row[col] > val
  })
}

function numFilterGe(data, col, val) {
  return data.filter((row) => {
    return row[col] >= val
  })
}

function numFilterLt(data, col, val) {
  return data.filter((row) => {
    return row[col] < val
  })
}

function numFilterLe(data, col, val) {
  return data.filter((row) => {
    return row[col] <= val
  })
}


// function numberOperators() {return ["=", "!=", ">", "<", ">=", "<="]}

export {
  numFilterEqual,
  numFilterNotEqual,
  numFilterGt,
  numFilterGe,
  numFilterLt,
  numFilterLe,
}