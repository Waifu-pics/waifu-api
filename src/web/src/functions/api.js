const { default: Axios } = require("axios")

export const api = {
  getEndpoints: (callback) => {
    Axios({
      method: "get",
      url: `/api/endpoints`,
    }).then((response) => {
      callback(response.data)
    })
  },
  checkLoggedIn: () => {
    return Axios.post(`/api/login`).then(res => {
      return res.status
    })
  },
  getOne: (endpoint, nsfw, callback) => {
    Axios({
      method: "get",
      url: `/api/${nsfw ? "nsfw" : "sfw"}/${endpoint}`,
    }).then((response) => {
      callback(response.data)
    })
  },
}
