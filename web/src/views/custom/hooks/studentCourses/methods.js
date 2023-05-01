import axios from 'axios'
import store from '@/store/index.js'

function getAllList() {
  const token = store.getters.token

  return axios({
    method: 'get',
    url: '/student_courses',
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}

function postCreateOne(data) {
  const token = store.getters.token

  return axios({
    method: 'post',
    url: '/student_courses',
    data: data,
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}

function getManyByCustomID(params) {
  const token = store.getters.token

  return axios({
    method: 'get',
    url: '/student_courses/custom_id',
    headers: {
      "Authorization": "Bearer " + token
    },
    params: params,
  })
}

export default {
  getAllList,
  postCreateOne,
  getManyByCustomID
}
