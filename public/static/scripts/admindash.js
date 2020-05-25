axios({
    method: 'post',
    url: '/api/admin/list',
    headers: {
        "token": getCookie("token")
    }
}).then(function (response) {
    if (response.data === null) return
    response.data.map( ({File, Type}, index) => {
        let listElem = document.createElement("div")
        listElem.id = index.toString()
        listElem.innerHTML = `
            <th><p style="display: inline; color: #7a7a7a;">${Type}</p></th>
            <th><a style="color: #8a8a8a" href="${Root}${File}">${File}</a></th>
            <th><a filename="${File}" id="${index}" style="color: var(--error-color);" class="deleteFile">Delete</a></th>
            <th><a filename="${File}" id="${index}" style="color: var(--primary-color);" class="verifyFile">Verify</a></th>
        `

        document.getElementById("pendingList").appendChild(listElem)

    })
    setClickListener()
})

function logout() {
    deleteCookie("token")
    window.location.reload()
}

function setClickListener() {
    // Delete file
    let deleteButtons = document.querySelectorAll(".deleteFile")
    for(let i = 0, x = deleteButtons.length; i < x; i++) {
        deleteButtons[i].onclick = () => {
            let id = deleteButtons[i].getAttribute("id")
            let file = deleteButtons[i].getAttribute("filename")
            axios({
                method: 'post',
                url: '/api/admin/verify',
                data: {
                    'file': file,
                    'isVer': false
                },
                headers: {
                    'token': getCookie("token")
                }
            }).then(function () {
                document.getElementById(id).innerHTML = "This file has been deleted!"
            })
        }
    }

    // Verify file
    let verifyButtons = document.querySelectorAll(".verifyFile")
    for(let i = 0, x = verifyButtons.length; i < x; i++) {
        verifyButtons[i].onclick = () => {
            let id = verifyButtons[i].getAttribute("id")
            let file = verifyButtons[i].getAttribute("filename")
            axios({
                method: 'post',
                url: '/api/admin/verify',
                data: {
                    'file': file,
                    'isVer': true
                },
                headers: {
                    'token': getCookie("token")
                }
            }).then(function () {
                document.getElementById(id).innerHTML = "This file has been verified!"
            })
        }
    }
}