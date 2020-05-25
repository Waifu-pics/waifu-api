scrollTop()

document.getElementById("relBtn").onclick = () => {
    document.getElementById("photos").innerHTML = ""
    getMore()
    scrollTop()
}

function scrollTop() {
    document.body.scrollTop = 0; // For Safari
    document.documentElement.scrollTop = 0; // For Chrome, Firefox, IE and Opera
}

let exclude = [] // Exclude array for pagination

// Original request on load
axios({
    method: 'post',
    url: '/api/many/' + Endpoint
}).then((response) => {
    response.data.data.map(function(file) {
        // Create the image element
        let image = document.createElement("img")
        image.src = Root + file

        document.getElementById("photos").appendChild(image)
        exclude.push(file)
    })
})

// Request for nextpage
function getMore() {
    axios({
        method: 'post',
        url: '/api/many/' + Endpoint,
        data: {
            exclude: exclude
        }
    }).then((response) => {
        response.data.data.map(function(file) {
            // Create the image element
            let image = document.createElement("img")
            image.src = Root + file

            document.getElementById("photos").appendChild(image)
            exclude.push(file)
        })
        if (response.data.data.length === 0) {
            let errorMsg = document.getElementById("endMsg")
            errorMsg.innerHTML = "You have reached the end!"
            errorMsg.classList.add("centered")
            document.getElementById("relBtn").remove()
        }
    })
}