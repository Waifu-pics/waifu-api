function gotoPage(param) {
    if (param === "sfw") {
        window.location.replace(`/`)
        return
    }
    window.location.replace(`/${param}`)
}