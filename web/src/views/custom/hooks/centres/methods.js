import axios from 'axios'
import store from '@/store/index.js'

function getCentresList() {
  const token = store.getters.token

  return axios({
    method: 'get',
    url: '/centres',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}

function postCreateOne(data) {
  const token = store.getters.token

  return axios({
    method: 'post',
    url: '/centres',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    data: data,
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}

export default {
  getCentresList,
  postCreateOne,
}
