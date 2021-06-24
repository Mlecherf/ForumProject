// On récupère nos boutons sur les côtés. 
const arrowLeftTheme = document.querySelectorAll(".arrowLeftTheme")[0]
const arrowRightTheme = document.querySelectorAll(".arrowRightTheme")[0]

// Si on click sur le bouton, on décale theme de 500px sur un côté où l'autre, à l'horizontal
arrowLeftTheme.addEventListener("click", () => {
    const theme = document.getElementById("theme")
    theme.scrollTo({
        top:0,
        left:theme.scrollLeft-500,
        behavior:"smooth"
    })
})
arrowRightTheme.addEventListener("click", () => {
    const theme = document.getElementById("theme")
    theme.scrollTo({
        top:0,
        left:theme.scrollLeft+500,
        behavior:"smooth"
    })
})
