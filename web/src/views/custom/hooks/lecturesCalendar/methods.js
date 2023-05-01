import axios from 'axios'
import store from '@/store/index.js'

function getAllList() {
  const token = store.getters.token

  return axios({
    method: 'get',
    url: '/lectures_calendar',
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
    url: '/lectures_calendar',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    data: data,
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}

function postDeleteByID(id) {
  const token = store.getters.token

  const url = `/lectures_calendar/${id}`
  console.log("delete url: ", url)

  return axios({
    method: 'delete',
    url: '/lectures_calendar/',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    headers: {
      "Authorization": "Bearer " + token
    }
  })
}

async function getManyByCourseID(params) {
  const token = store.getters.token

  return await axios({
    method: 'get',
    url: '/lectures_calendar/getByCourseID',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    headers: {
      "Authorization": "Bearer " + token
    },
    params: params,
  })
}

function getOneByID(id) {
  const token = store.getters.token

  return axios({
    method: 'get',
    url: `/lectures_calendar/${id}`,
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    headers: {
      "Authorization": "Bearer " + token
    },
  })
}

export default {
  getAllList,
  postCreateOne,
  postDeleteByID,
  getManyByCourseID,
  getOneByID
}
