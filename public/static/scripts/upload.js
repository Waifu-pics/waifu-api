if (getCookie("token") != null) {
    axios({
        method: 'post',
        url: '/api/admin/token',
        headers: {
            'token': getCookie("token")
        }
    }).then(function () {
        zone.options.headers.token = getCookie("token")
        document.getElementById("adminBox").style.visibility = "visible"
        document.getElementById("responseBox").style.marginTop = "20px"
    }).catch(function () {
        deleteCookie("token")
    })
}

function reDefOpt() {
    zone.options.headers.type = $("#select option:selected").text().toLowerCase()
}

var zone = new Dropzone("#zone", {
    url: "/api/upload",
    paramName: "uploadFile",

    previewTemplate: `<div id="tpl"><div class="dz-preview dz-processing dz-image-preview dz-complete"></div></div>`,
    timeout: 0,

    headers: {
        'type': $("#select option:selected").text().toLowerCase(),
        'token': null
    },

    init: function () {
        this.on("success", function () {
                unHide("success")
                document.getElementById('responseBox').innerHTML = `File was successfully uploaded`
            }),

            this.on("error", function (data) {
                unHide("error")
                document.getElementById('responseBox').innerHTML = `${data.xhr.response}`
            })
    }

})

function unHide(resType) {
    let resBox = document.getElementById("responseBox")

    if (resType == "success") {
        resBox.classList.add("terminal-alert-primary")
        resBox.style.visibility = "visible"
        deleteBox("success")
    } else {
        resBox.classList.add("terminal-alert-error")
        resBox.style.visibility = "visible"
        deleteBox("error")
    }

    function deleteBox(resType) {
        setTimeout(function () {
            resBox.style.visibility = "hidden"
            if (resType == "success") {
                resBox.classList.remove("terminal-alert-primary")
            } else {
                resBox.classList.remove("terminal-alert-error")
            }
        }, 5000)
    }

}

let modal = document.getElementById("guidelineModal")
let btn = document.getElementById("guideBTN")
let span = document.getElementsByClassName("close")[0]

btn.onclick = function () {
    modal.style.display = "block";
}

// When the user clicks on <span> (x), close the modal
span.onclick = function () {
    modal.style.display = "none";
}

// When the user clicks anywhere outside of the modal, close it
window.onclick = function (event) {
    if (event.target == modal) {
        modal.style.display = "none";
    }
}