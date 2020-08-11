const { default: Axios } = require("axios")

export const api = {
  getEndpoints: function (url) {
    let endpoints = []

    Axios({
      method: "get",
      url: `${process.env.VUE_APP_APIROOT}/api/endpoints`,
    }).then((response) => {
      response.data.forEach(elem => {
        if (url) {
          endpoints.push(`${process.env.VUE_APP_APIROOT}/api/${elem}`)
        } else {
          endpoints.push(elem)
        }
      })
    })

    return endpoints
  },
}
