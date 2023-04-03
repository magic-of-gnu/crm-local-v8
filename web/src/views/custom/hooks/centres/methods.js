// import axios from 'axios'
// const axios = require('axios')
import axios from 'axios'

async function getCentresList() {
  return await axios({
    method: 'get',
    url: '/centres/list',
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

async function postCentresCreateOne(data) {
  return await axios({
    method: 'post',
    url: '/rooms/create_one',
    baseURL: 'http://malcorp.test/api/server',
    crossOrigin: true,
    responseType: 'json',
    data: {
      centre_id: data.centre_id,
      description: data.description,
    },
  })
    .then((response) => {
      console.log('create_one then')
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
  getCentresList,
  postCentresCreateOne,
}
