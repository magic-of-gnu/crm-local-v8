import axios from 'axios'
import store from '@/store/index.js'


function getAllList() {
  const token = store.getters.token

  return axios({
    method: 'get',
    url: '/attendances',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    headers: {
      "Authorization": "Bearer ${token}"
    }
  })
}

function postCreateOne(data) {
  const token = store.getters.token

  return axios({
    method: 'post',
    url: '/attendances',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    data: data,
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}

function postCreateMany(data) {
  const token = store.getters.token

  return axios({
    method: 'post',
    url: '/attendances/create_many',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    data: data,
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}

async function postDeleteByID(data) {
  const token = store.getters.token

  return await axios({
    method: 'delete',
    url: '/attendances',
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
      // always executed
    })
}


export default {
  getAllList,
  postCreateOne,
  postCreateMany,
  postDeleteByID,
}
