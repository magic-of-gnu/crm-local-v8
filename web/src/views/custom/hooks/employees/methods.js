import axios from 'axios'
import store from '@/store/index.js'


async function getEmployeesList() {
  const token = store.getters.token

  return await axios({
    method: 'get',
    url: '/employees/list',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    headers: {
      "Authorization": "Bearer " + token
    }
  })
    .then((response) => {
      if (response.status === 200) {
        console.log('employees then; ', response)
        return response.data
      }
    })
    .catch(function (error) {
      console.log('error: ', error)
      console.log('i got error here: ', error)
    })
    .finally(function () {
      console.log('finally')
      // always executed
    })
}

async function postEmployeesCreateOne(data) {
  const token = store.getters.token

  return await axios({
    method: 'post',
    url: '/employees/create_one',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    data: data,
    headers: {
      "Authorization": "Bearer " + token
    }
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
