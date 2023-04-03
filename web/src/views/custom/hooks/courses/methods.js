import axios from 'axios'

async function getCoursesList() {
  return await axios({
    method: 'get',
    url: '/courses/list',
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
      console.log('error: ', error)
    })
    .finally(function () {
      console.log('finally')
      // always executed
    })
}

async function postCreateOne(data) {
  return await axios({
    method: 'post',
    url: '/courses/create_one',
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
  getCoursesList,
  postCreateOne,
}
