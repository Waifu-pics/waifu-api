if (localStorage.getItem("token") != null) {
    axios({
        method: 'post',
        url: '/api/admin/auth/verify',
        data: {
            'token': localStorage.getItem("token")
        }
    }).then(function () {
        window.location.replace("/admin/dash")
    }).catch(function () {
        localStorage.removeItem("token")
    })
}

function tryLogin() {
    axios({
        method: 'post',
        url: '/api/admin/auth/login',
        data: {
            'username': document.getElementById("usernameField").value,
            'password': document.getElementById("passwordField").value
        }
    }).then(function (data) {
        localStorage.setItem('token', data.data)
        window.location.replace("/admin/dash")
    }).catch(function (data) {
        let resBox = document.getElementById("responseBox")
        resBox.style.visibility = "visible"

        resBox.innerHTML = data.response.data
        
        setTimeout(function() { 
            resBox.style.visibility = "hidden"
        }, 5000)
    })
}