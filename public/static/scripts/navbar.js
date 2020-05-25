const navTabs = document.querySelectorAll(".NAVTAB")
let path = window.location.pathname

navTabs.forEach(navTab => {
    if (path === "/pages") {
        document.getElementById("PAGESTAB").classList.add("active")
    } if (path === "/") {
        document.getElementById("HOMETAB").classList.add("active")
    } else if (navTab.getAttribute("href") === path) {
        document.getElementById(navTab.innerHTML + "TAB").classList.add("active")
    }
})

function menuToggle() {
    const x = document.getElementById("hideLinks")
    if (!x.classList.contains("hidden")) {
        x.classList.add("hidden")
    } else {
        x.classList.remove("hidden")
    }
}