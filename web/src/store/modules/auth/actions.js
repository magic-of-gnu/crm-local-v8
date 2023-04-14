import axios from 'axios'
import { ErrorCodes } from 'vue'

export default {
  async login (context, payload) {
    const urlLogin = "http://malcorp.test/api/server/login"

    console.log("payload: ", payload)

    const responseData = await axios({
      method: 'post',
      url: '/login',
      baseURL: 'http://malcorp.test/api/server',
      responseType: 'json',
      data: payload,
    }).then((response) => {
      if (response.status === 200) {
        return response.data
      }
      const error = new Error(response.data.message || 'Failed to authenticate.');
      throw error
    
    })

    localStorage.setItem('token', responseData.data.token)
    localStorage.setItem('userID', responseData.data.user.id)

    context.commit('setUser', {
      token: responseData.data.token,
      userId: responseData.data.user.id,
      tokenExpiration: responseData.data.expiresIn,
    })
  },
  tryLogin(context) {
    const token = localStorage.getItem('token');
    const userID = localStorage.getItem('userID')

    if (token) {
      context.commit('setUser', {
        token: token,
      })
    }

  }

}