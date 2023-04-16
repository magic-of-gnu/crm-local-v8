import axios from 'axios'
import store from '@/store/index.js'

async function getAllList() {
  const token = store.getters.token

  return await axios({
    method: 'get',
    url: '/students/list',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    headers: {
      "Authorization": "Bearer " + token
    }
  })
    .then((response) => {
      if (response.status === 200) {
        console.log('students then; ', response)
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

async function postStudentsCreateOne(data) {
  const token = store.getters.token

  return await axios({
    method: 'post',
    url: '/students/create_one',
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
  getAllList,
  postStudentsCreateOne,
}
