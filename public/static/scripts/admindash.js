axios({
    method: 'post',
    url: '/api/admin/list',
    headers: {
        "token": getCookie("token")
    }
}).then(function (response) {
    if (response.data === null) return
    response.data.map( ({File, Type}, index) => {
        $("#efs").append(`
                <div id="${index}">
                <th><p style="display: inline; color: #7a7a7a;">${Type}</p></th>
                <th><a style="color: #8a8a8a" href="https://i.waifu.pics/${File}">${File}</a></th>
                <th><a filename="${File}" id="${index}" style="color: var(--error-color);" class="dlfl">Delete</a></th>
                <th><a filename="${File}" id="${index}" style="color: var(--primary-color);" class="vfl">Verify</a></th>
                </div>
                `)
    })
})

function logout() {
    deleteCookie("token")
    window.location.reload()
}

// Delete files
$(document).on('click','.dlfl', function() {
    let id = $(this).attr('id')
    let file = $(this).attr('filename')
    // make delete request with id
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
        removeMsg = `This file has been deleted!`
        document.getElementById(id).innerHTML = removeMsg
    })
})

$(document).on('click','.vfl', function() {
    let id = $(this).attr('id')
    let file = $(this).attr('filename')
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
    }).then(function() {
        removeMsg = `This file has been verified!`
        document.getElementById(id).innerHTML = removeMsg
    })
})