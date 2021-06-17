const arrowTop = document.querySelectorAll(".arrowTop")[0]
const arrowBottom = document.querySelectorAll(".arrowBottom")[0]

arrowTop.addEventListener("click", () => {
    const theme = document.getElementById("test")
    theme.scrollTo({
        top:theme.scrollTop-400,
        left:0,
        behavior:"smooth"
    })
})

arrowBottom.addEventListener("click", () => {
    const theme = document.getElementById("test")
    theme.scrollTo({
        top:theme.scrollTop+400,
        left:0,
        behavior:"smooth"
    })
})
 



// const top__ = document.getElementById("go_top")
// const bottom = document.getElementById('go_bottom')

// bottom.addEventListener('click', (e)=>{
//     const theme = document.getElementById("theme")
//     theme.scrollTo({
//         left:0,
//         top:theme.scrollLeft+500,
//         behavior:"smooth"
//     })

// })