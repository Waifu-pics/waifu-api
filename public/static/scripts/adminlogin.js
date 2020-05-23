function tryLogin() {
    axios({
        method: 'post',
        url: '/api/admin/login',
        data: {
            'username': document.getElementById("usernameField").value,
            'password': document.getElementById("passwordField").value
        }
    }).then(function (data) {
        createCookie("token", data.data)
        window.location.reload()
    }).catch(function (data) {
        let resBox = document.getElementById("responseBox")
        resBox.style.visibility = "visible"

        resBox.innerHTML = data.response.data
        
        setTimeout(function() { 
            resBox.style.visibility = "hidden"
        }, 5000)
    })
}