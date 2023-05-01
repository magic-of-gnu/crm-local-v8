import axios from 'axios'
import store from '@/store/index.js'

function getAllList() {
  const token = store.getters.token

  return axios({
    method: 'get',
    url: '/users/list',
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}

function postCreateOne(data) {
  const token = store.getters.token

  return axios({
    method: 'post',
    url: '/users/create_one',
    data: data,
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}

function postDeleteByID(data) {
  const token = store.getters.token

  return axios({
    method: 'delete',
    url: '/users',
    data: data,
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}


export default {
  getAllList,
  postCreateOne,
  postDeleteByID,
}
