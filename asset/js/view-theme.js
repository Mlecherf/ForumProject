// On récupère nos boutons sur les côtés. 
const arrowTop = document.querySelectorAll(".arrowTop")[0]
const arrowBottom = document.querySelectorAll(".arrowBottom")[0]
// Et ceux pour le responsive
const arrowLeft = document.querySelectorAll(".arrowLeft")[0]
const arrowRight = document.querySelectorAll(".arrowRight")[0]

// Si on click sur le bouton, on décale theme de 500px sur un côté où l'autre, à la vertical
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

// De même pour le responsive, à l'horizontal
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


