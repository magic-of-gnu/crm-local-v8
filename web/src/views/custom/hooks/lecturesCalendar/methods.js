import axios from 'axios'
import store from '@/store/index.js'

async function getAllList() {
  const token = store.getters.token

  return await axios({
    method: 'get',
    url: '/lectures_calendar/list',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    headers: {
      "Authorization": "Bearer " + token
    }
  })
    .then((response) => {
      if (response.status === 200) {
        return response.data
      }
    })
    .catch(function (error) {
      console.log('error: ', error)
    })
    .finally(function () {
      console.log('finally')
      // always executed
    })
}

async function postCreateOne(data) {
  const token = store.getters.token

  return await axios({
    method: 'post',
    url: '/lectures_calendar/create_one',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    data: data,
    headers: {
      "Authorization": "Bearer " + token
    }
  })
    .then((response) => {
      if (response.status === 200) {
        return response.data
      }
    })
    .catch(function (error) {
      console.log('error: ', error)
    })
    .finally(function () {
      // always executed
    })
}

async function postDeleteByID(data) {
  const token = store.getters.token

  return await axios({
    method: 'delete',
    url: '/lectures_calendar/',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    data: data,
    headers: {
      "Authorization": "Bearer " + token
    }
  })
    .then((response) => {
      if (response.status === 200) {
        return response.data
      }
    })
    .catch(function (error) {
      console.log('error: ', error)
    })
    .finally(function () {
      // always executed
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
    .then((response) => {
      if (response.status === 200) {
        return response.data
      }
    })
}

async function getOneByID(id) {
  const token = store.getters.token

  return await axios({
    method: 'get',
    url: `/lectures_calendar/${id}`,
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    headers: {
      "Authorization": "Bearer " + token
    },
  })
    .then((response) => {
      if (response.status === 200) {
        return response.data
      }
    })
}

export default {
  getAllList,
  postCreateOne,
  postDeleteByID,
  getManyByCourseID,
  getOneByID
}
