import axios from 'axios'
import store from '@/store/index.js'

async function getCentresList() {
  const token = store.getters.token

  return await axios({
    method: 'get',
    url: '/centres/list',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
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
      console.log('finally')
      // always executed
    })
}

async function postCentresCreateOne(data) {
  const token = store.getters.token

  return await axios({
    method: 'post',
    url: '/centres/create_one',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    data: data,
    headers: {
      "Authorization": "Bearer " + token
    }
  })
    .then((response) => {
      console.log('create_one then')
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
  getCentresList,
  postCentresCreateOne,
}
