const arrowLeftTheme = document.querySelectorAll(".arrowTopTheme")[0]
const arrowRightTheme = document.querySelectorAll(".arrowBottomTheme")[0]

arrowLeftTheme.addEventListener("click", () => {
    console.log("Aidr MWOA")
    const theme = document.getElementById("ici")
    theme.scrollTo({
        left:0,
        top:theme.scrollTop-500,
        behavior:"smooth"
    })
})

arrowRightTheme.addEventListener("click", () => {
    console.log("j'aime pas quand il ny pas derreur mais que cela ne marche PAS NIQUE !!!!!!!!!!!!!!!!!! tg")
    const theme = document.getElementById("ici")
    theme.scrollTo({
        left:0,
        top:theme.scrollTop+500,
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