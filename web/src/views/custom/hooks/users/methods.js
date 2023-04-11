import axios from 'axios'


async function getAllList() {
  return await axios({
    method: 'get',
    url: '/users/list',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
  })
    .then((response) => {
      if (response.status === 200) {
        return response.data
      }
    })
    .catch(function (error) {
    })
    .finally(function () {
      // always executed
    })
}

async function postCreateOne(data) {
  return await axios({
    method: 'post',
    url: '/users/create_one',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    data: data,
  })
    .then((response) => {
      if (response.status === 200) {
        return response.data
      }
    })
    .catch(function (error) {
    })
    .finally(function () {
      // always executed
    })
}

async function postDeleteByID(data) {
  return await axios({
    method: 'delete',
    url: '/users',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    data: data,
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


export default {
  getAllList,
  postCreateOne,
  postDeleteByID,
}
