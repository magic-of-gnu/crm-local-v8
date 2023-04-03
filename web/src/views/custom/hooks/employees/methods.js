// import axios from 'axios'
// const axios = require('axios')
import axios from 'axios'

async function getEmployeesList() {
  return await axios({
    method: 'get',
    url: '/employees/list',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
  })
    .then((response) => {
      if (response.status === 200) {
        console.log('employees then; ', response)
        return response.data
      }
    })
    .catch(function (error) {
      console.log('error: ', error)
    })
    .finally(function () {
      console.log('finally')
      // always executed
    })
}

async function postEmployeesCreateOne(data) {
  return await axios({
    method: 'post',
    url: '/employees/create_one',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    data: data,
  })
    .then((response) => {
      if (response.status === 200) {
        return response.data
      }
    })
    .catch(function (error) {
      console.log('error: ', error)
    })
    .finally(function () {
      console.log('create one finally')
      // always executed
    })
}

export default {
  getEmployeesList,
  postEmployeesCreateOne,
}
