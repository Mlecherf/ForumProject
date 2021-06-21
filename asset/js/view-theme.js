const arrowTop = document.querySelectorAll(".arrowTop")[0]
const arrowBottom = document.querySelectorAll(".arrowBottom")[0]
const arrowLeft = document.querySelectorAll(".arrowLeft")[0]
const arrowRight = document.querySelectorAll(".arrowRight")[0]

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

arrowLeft.addEventListener("click", () => {
    const theme = document.getElementById("test")
    theme.scrollTo({
        top:0,
        left:theme.scrollLeft-200,
        behavior:"smooth"
    })
})

arrowRight.addEventListener("click", () => {
    const theme = document.getElementById("test")
    theme.scrollTo({
        top:0,
        left:theme.scrollLeft+200,
        behavior:"smooth"
    })
})


