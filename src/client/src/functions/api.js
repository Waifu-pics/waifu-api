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
  checkLoggedIn: async () => {
    return Axios.post(`/api/admin/login`).then(res => {
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
  generateImage: (endpoint, nsfw, text, callback) => {
    Axios({
      method: "post",
      url: `/api/gen`,
      responseType: 'arraybuffer',
      data: {
        endpoint: {
          nsfw: nsfw,
          type: endpoint,
        },
        text: {
          top: text.top,
          bottom: text.bottom,
        },
      },
    }).then((response) => {
      callback(response.data, null)
    }).catch((err) => {
      if (err.response.data) {
        // Response on success sends arraybuffer, on fail sends text, convert buf to text
        let errormsg = String.fromCharCode.apply(null, new Uint8Array(err.response.data))

        try {
          callback(null, JSON.parse(errormsg).message)
        } catch(e) {
          callback(null, "You are being rate limited")
        }
        
        return
      }
      callback(null, "Image could not be generated")
    })
  },
}
