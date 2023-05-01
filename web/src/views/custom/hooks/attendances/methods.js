import axios from 'axios'
import store from '@/store/index.js'


function getAllList() {
  const token = store.getters.token

  return axios({
    method: 'get',
    url: '/attendances',
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
    data: data,
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}


export default {
  getAllList,
  postCreateOne,
  postCreateMany,
  postDeleteByID,
}
