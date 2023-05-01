import axios from 'axios'
import store from '@/store/index.js'

function getAllList() {
  const token = store.getters.token

  return axios({
    method: 'get',
    url: '/payment_statuses',
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}

function postCreateOne(data) {
  const token = store.getters.token

  return axios({
    method: 'post',
    url: '/payment_statuses',
    data: data,
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}

function postDeleteOneByID(id) {
  const token = store.getters.token

  const url = `/payment_statuses/${id}`
  console.log("delete url: ", url)

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
