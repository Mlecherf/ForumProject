const arrowLeftTheme = document.querySelectorAll(".arrowLeftTheme")[0]
const arrowRightTheme = document.querySelectorAll(".arrowRightTheme")[0]

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

const arrowLeftPost = document.querySelectorAll(".arrowLeftPost")[0]
const arrowRightPost = document.querySelectorAll(".arrowRightPost")[0]

arrowLeftPost.addEventListener("click", () => {
    const post = document.getElementById("post")
    post.scrollTo({
        top:0,
        left:post.scrollLeft-500,
        behavior:"smooth"
    })
})

arrowRightPost.addEventListener("click", () => {
    const post = document.getElementById("post")
    post.scrollTo({
        top:0,
        left:post.scrollLeft+500,
        behavior:"smooth"
    })
})