import axios from 'axios'

export default {
  async login (context, payload) {
    const responseData = await axios({
      method: 'post',
      url: '/login',
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
  },
  logout(context) {
    localStorage.removeItem('token')
    localStorage.removeItem('userID')

    context.commit('setUser', {
      token: null,
      userId: null,
      tokenExpiration: null,
    })
  }


}