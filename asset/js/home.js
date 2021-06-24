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

PopulatPost = [...document.getElementsByClassName("theme")]
for (let step = 0; step < 5; step++) {
    if (PopulatPost[step].innerHTML.slice(3,PopulatPost[step].innerHTML.length-4) == "Fast_Food"){
        PopulatPost[step].innerHTML = "<p>Fast Food</p>"
    }
    if (PopulatPost[step].innerHTML.slice(3,PopulatPost[step].innerHTML.length-4) == "America_Latina"){
        PopulatPost[step].innerHTML = "<p>America Latina</p>"
    }
}

Tags = [...document.getElementsByClassName("tag")]
Tags.forEach(element => {
    console.log(element.innerHTML)
    if (element.innerText == "America_Latina"){
        element.innerHTML = "<p>America Latina</p>"
    }
    if (element.innerText == "Fast_Food"){
        element.innerHTML = "<p>Fast Food</p>"
    }
});