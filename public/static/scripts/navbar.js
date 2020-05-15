if (window.location.href == `${window.location.origin}/`) {    
    document.getElementById("HOMETAB").classList.add("active")
} else if (window.location.href == `${window.location.origin}/nsfw`) {
    document.getElementById("NSFWTAB").classList.add("active")
} else if (window.location.href == `${window.location.origin}/upload`) {
    document.getElementById("UPLOADTAB").classList.add("active")
} else if (window.location.href == `${window.location.origin}/docs`) {
    document.getElementById("APITAB").classList.add("active")
}