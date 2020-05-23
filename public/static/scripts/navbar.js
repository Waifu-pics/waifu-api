const navTabs = document.querySelectorAll(".NAVTAB")
let path = window.location.pathname

navTabs.forEach(navTab => {
    if (path === "/pages") {
        document.getElementById("PAGESTAB").classList.add("active")
    } else if (navTab.getAttribute("href") === path) {
        document.getElementById(navTab.innerHTML + "TAB").classList.add("active")
    }
})