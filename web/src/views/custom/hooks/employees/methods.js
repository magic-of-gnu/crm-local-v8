import axios from 'axios'
import store from '@/store/index.js'


function getEmployeesList() {
  const token = store.getters.token

  return axios({
    method: 'get',
    url: '/employees/list',
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
    url: '/employees/create_one',
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
  getEmployeesList,
  postCreateOne,
}
