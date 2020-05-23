function createCookie(key, value) {
    document.cookie = `${key}=${value}`
}

function deleteCookie(key) {
    document.cookie = `${key}=; Max-Age=-99999999;`
}

function getCookie(key) {
    const value = `; ${document.cookie}`
    const parts = value.split(`; ${key}=`)
    if (parts.length === 2) return parts.pop().split(';').shift()
}