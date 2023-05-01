import axios from 'axios'
import store from '@/store/index.js'

function getAllList() {
  const token = store.getters.token

  return axios({
    method: 'get',
    url: '/centres',
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
    data: data,
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}

function postDeleteOneByID(id) {
  const token = store.getters.token

  return axios({
    method: 'delete',
    url: `/centres/${id}`,
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}

export default {
  getAllList,
  postCreateOne,
  postDeleteOneByID,
}
