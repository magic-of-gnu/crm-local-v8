// import axios from 'axios'
// const axios = require('axios')
import axios from 'axios'

async function getRoomsList() {
  return await axios({
    method: 'get',
    url: '/rooms/list',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
  })
    .then((response) => {
      if (response.status === 200) {
        console.log('rooms then; ', response)
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

async function postRoomsCreateOne(data) {
  return await axios({
    method: 'post',
    url: '/rooms/create_one',
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
      console.log('create one finally')
      // always executed
    })
}

export default {
  getRoomsList,
  postRoomsCreateOne,
}
