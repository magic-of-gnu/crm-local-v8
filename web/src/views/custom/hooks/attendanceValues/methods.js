import axios from 'axios'
import store from '@/store/index.js'

async function getAllList() {
  const token = store.getters.token

  return await axios({
    method: 'get',
    url: '/attendance_values',
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}

function postCreateOne(data) {
  const token = store.getters.token

  return axios({
    method: 'post',
    url: '/attendance_values',
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
