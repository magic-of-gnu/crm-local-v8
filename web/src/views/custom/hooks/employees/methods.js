import axios from 'axios'
import store from '@/store/index.js'


function getAllList() {
  const token = store.getters.token

  return axios({
    method: 'get',
    url: '/employees',
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}

function postCreateOne(data) {
  const token = store.getters.token

  return axios({
    method: 'post',
    url: '/employees',
    data: data,
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}

export default {
  getAllList,
  postCreateOne,
}
