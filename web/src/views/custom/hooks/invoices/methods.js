import axios from 'axios'
import store from '@/store/index.js'

function getAllList() {
  const token = store.getters.token

  return axios({
    method: 'get',
    url: '/invoices',
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}

function postCreateOne(data) {
  const token = store.getters.token

  return axios({
    method: 'post',
    url: '/invoices',
    data: data,
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}

function postDeleteOneByID(id) {
  const token = store.getters.token

  const url = `/invoices/${id}`

  return axios({
    method: 'delete',
    url: url,
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
