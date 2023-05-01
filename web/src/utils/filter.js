import { strFilterIn, strFilterEqual, strFilterNotEqual } from "./strFilters"
import { numFilterEqual, numFilterGe, numFilterGt, numFilterLe, numFilterLt, numFilterNotEqual } from "./numFilters"

function applyFilter(data, columns, tos, iss, outputData) {
  const n = columns.length

  let newdata = data.slice()
  console.log('before data: ', newdata)

  for (let ii = 0; ii < n; ii++ ) {
    if (iss[ii] === 'in' && (tos[ii] !== "" && tos[ii] !== null) ) {
      newdata = strFilterIn(newdata, columns[ii], tos[ii])
    } else
    if (iss[ii] === '=' && (tos[ii] !== "" && tos[ii] !== null) ) {
      newdata = strFilterEqual(newdata, columns[ii], tos[ii])
    } else
    if (iss[ii] === '!=' && (tos[ii] !== "" && tos[ii] !== null) ) {
      newdata = strFilterNotEqual(newdata, columns[ii], tos[ii])
    }
  }

  outputData.value = newdata
  return newdata

}

function applyAllFilters(data, columns, tos, iss, outputData) {
  const n = columns.length

  let newdata = data.slice()

  for (let ii = 0; ii < n; ii++ ) {
    if (iss[ii] === 'sin' && (tos[ii] !== "" && tos[ii] !== null) ) {
      newdata = strFilterIn(newdata, columns[ii], tos[ii])
    } else
    if (iss[ii] === 's==' && (tos[ii] !== "" && tos[ii] !== null) ) {
      newdata = strFilterEqual(newdata, columns[ii], tos[ii])
    } else
    if (iss[ii] === 's!=' && (tos[ii] !== "" && tos[ii] !== null) ) {
      newdata = strFilterNotEqual(newdata, columns[ii], tos[ii])
    }
    if (iss[ii] === 'neq' && (tos[ii] !== "" && tos[ii] !== null) ) {
      newdata = numFilterEqual(newdata, columns[ii], Number(tos[ii]))
    } else
    if (iss[ii] === 'nneq' && (tos[ii] !== "" && tos[ii] !== null) ) {
      newdata = numFilterNotEqual(newdata, columns[ii], Number(tos[ii]))
    } else
    if (iss[ii] === 'ngt' && (tos[ii] !== "" && tos[ii] !== null) ) {
      newdata = numFilterGt(newdata, columns[ii], Number(tos[ii]))
    } else
    if (iss[ii] === 'nge' && (tos[ii] !== "" && tos[ii] !== null) ) {
      newdata = numFilterGe(newdata, columns[ii], Number(tos[ii]))
    } else
    if (iss[ii] === 'nlt' && (tos[ii] !== "" && tos[ii] !== null) ) {
      newdata = numFilterLt(newdata, columns[ii], Number(tos[ii]))
    } else
    if (iss[ii] === 'nle' && (tos[ii] !== "" && tos[ii] !== null) ) {
      newdata = numFilterLe(newdata, columns[ii], Number(tos[ii]))
    }
  }

  outputData.value = newdata
  return newdata

}

function applyNumFilter(data, columns, tos, iss, outputData) {
  const n = columns.length

  let newdata = data.slice()

  for (let ii = 0; ii < n; ii++ ) {
    if (iss[ii] === '=' && (tos[ii] !== "" && tos[ii] !== null) ) {
      newdata = numFilterEqual(newdata, columns[ii], tos[ii])
    } else
    if (iss[ii] === '!=' && (tos[ii] !== "" && tos[ii] !== null) ) {
      newdata = numFilterNotEqual(newdata, columns[ii], tos[ii])
    } else
    if (iss[ii] === '>' && (tos[ii] !== "" && tos[ii] !== null) ) {
      newdata = numFilterGt(newdata, columns[ii], tos[ii])
    } else
    if (iss[ii] === '>=' && (tos[ii] !== "" && tos[ii] !== null) ) {
      newdata = numFilterGe(newdata, columns[ii], tos[ii])
    } else
    if (iss[ii] === '<' && (tos[ii] !== "" && tos[ii] !== null) ) {
      newdata = numFilterLt(newdata, columns[ii], tos[ii])
    } else
    if (iss[ii] === '<=' && (tos[ii] !== "" && tos[ii] !== null) ) {
      newdata = numFilterLe(newdata, columns[ii], tos[ii])
    }
  }

  outputData.value = newdata
  return newdata

}

export {
  applyFilter,
  applyNumFilter,
  applyAllFilters,
}