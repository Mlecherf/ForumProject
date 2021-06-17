const arrowTop = document.querySelectorAll(".arrowTop")[0]
const arrowBottom = document.querySelectorAll(".arrowBottom")[0]

arrowTop.addEventListener("click", () => {
    const theme = document.getElementById("test")
    theme.scrollTo({
        top:theme.scrollTop-500,
        left:0,
        behavior:"smooth"
    })
})

arrowBottom.addEventListener("click", () => {
    const theme = document.getElementById("test")
    theme.scrollTo({
        top:theme.scrollTop+500,
        left:0,
        behavior:"smooth"
    })
})
