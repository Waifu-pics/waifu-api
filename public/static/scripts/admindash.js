if (localStorage.getItem("token") != null) {
    axios({
        method: 'post',
        url: '/api/admin/verifytoken',
        data: {
            'token': localStorage.getItem("token")
        }
    }).then(function () {
        axios({
            method: 'post',
            url: '/api/admin/listfile',
            data: {
                'token': localStorage.getItem("token"),
            }
        }).then(function (response) {
            response.data.map( ({file, type}, index) => {
                $("#efs").append(`
                <div id="${index}">
                <th><p style="display: inline; color: #7a7a7a;">${type}</p></th>
                <th><a style="color: #8a8a8a" href="https://i.waifu.pics/${file}">${file}</a></th>
                <th><a filename="${file}" id="${index}" style="color: var(--error-color);" class="dlfl">Delete</a></th>
                <th><a filename="${file}" id="${index}" style="color: var(--primary-color);" class="vfl">Verify</a></th>
                </div>
                `)
            })
        })
    }).catch(function () {
        localStorage.removeItem("token")
        window.location.replace("/admin/login")
    })
} else {
    window.location.replace("/admin/login")
}

function logout() {
    localStorage.removeItem("token")
    window.location.replace("/admin/login")
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
        'token': localStorage.getItem("token"),
        'image': file,
        'isVer': false
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
        'token': localStorage.getItem("token"),
        'image': file,
        'isVer': true
    }
}).then(function() {
    removeMsg = `This file has been verified!`
    document.getElementById(id).innerHTML = removeMsg
})
})