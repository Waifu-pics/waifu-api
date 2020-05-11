// Scroll to top on load
$(this).scrollTop(0)

$(document).on('click','#relBtn', function() {
    $("#photos").empty()
    getMore()
    $(this).scrollTop(0)
})

let exclude = []
axios({
    method: 'post',
    url: '/api/' + window.ENDPOINT
}).then((response) => {
    response.data.data.map(function(file) {
        $("#photos").append(`<img src="${window.IMGSERVE + file}" alt="">`)
        exclude.push(file)
    })
})

const getMore = () => {
    axios({
        method: 'post',
        url: '/api/' + window.ENDPOINT,
        data: {
            exclude: exclude
        }
    }).then((response) => {
        response.data.data.map(function(file) {
            $("#photos").append(`<img src="${window.IMGSERVE + file}" alt="">`)
            exclude.push(file)
        })
        if (response.data.data.length == 0) {
            console.log("end")
            $("#endMsg").text("You have reached the end!")
            $("#endMsg").addClass("centered")
            $("#relBtn").remove()
        }
    })
}