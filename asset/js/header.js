// Get the input field
var input = document.getElementById("Search__Text");

// Execute a function when the user releases a key on the keyboard
input.addEventListener("keyup", function(event) {
  // Number 13 is the "Enter" key on the keyboard
  if (event.keyCode === 13) {
    // Cancel the default action, if needed
    event.preventDefault();
    // Trigger the button element with a click
    console.log("recherche depuis le header")
    setTimeout(() => {
      document.getElementById("Search__BTN").removeAttribute("disabled");
      document.getElementById("Search__BTN").click();
    }, 2000);
  }
});

document.getElementById("close_post").addEventListener("click", ()=>{
    document.getElementById('pop_post_add').style.display = "none"
})
document.getElementById("Post_add").addEventListener("click", ()=>{
    document.getElementById('pop_post_add').style.display = "block"
})

document.getElementsByClassName("Login")[0].addEventListener("click", (event)=>{
    document.getElementById("Pop_up").style.display = "block"
    event.stopPropagation()
})
document.addEventListener("click", ()=>{
    document.getElementById("Pop_up").style.display = "none"
})
// document.cookie = "Login ='{'user':'Clem','mail':'mail@cookie.com','nb_posts':'10','nb_likes':'15'}'"
// document.cookie = "Connect = true"


let my_cookie_header = Select_Login_cookie()
if (my_cookie_header == "true"){
    document.getElementsByClassName("Login")[0].setAttribute("src", "https://img.icons8.com/fluent-systems-regular/45/000000/user-male-circle.png")
    document.getElementById("Post_add").style.display = "flex"
    document.getElementById("First_pop").innerHTML = "Profile"
    document.getElementById("First_pop").setAttribute("href","/profile")
    document.getElementById("Second_pop").innerHTML = "Logout"
    document.getElementById("Second_pop").setAttribute("href","/")
}else{
    document.getElementsByClassName("Login")[0].setAttribute("src", "https://img.icons8.com/windows/50/000000/user-ninja.png")
    document.getElementById("First_pop").innerHTML = "Register"
    document.getElementById("First_pop").setAttribute("href","/register")
    document.getElementById("Second_pop").innerHTML = "Login"
    document.getElementById("Second_pop").setAttribute("href","/login")
}

function Select_Login_cookie (){
    let my_cookie_login = ""

    document.cookie.split("; ").forEach((elem)=>{
        if (elem.slice(0,7) == "Connect"){
            my_cookie_login = elem.slice(8,elem.length)
        }
    })
    return my_cookie_login
}

const Post_tag = [...document.getElementsByClassName('tag__post')]
const restes_tag = document.getElementById('reset__tag')
const nb_tag = document.getElementById('nb_tag')
let nb_tag_up = 0;

Post_tag.forEach((elem, index)=>{
    elem.addEventListener('click', ()=>{
        Up_tag(index)
    })
})

restes_tag.addEventListener('click', ()=>{
    Post_tag.forEach((elem)=>{
        elem.value = "down"
    })
    nb_tag.innerHTML = `Tag : 0/4`
})

function Up_tag (index_tag){
    Post_tag[index_tag].value = "up"
    nb_tag_up ++
    nb_tag.innerHTML = `Tag : ${nb_tag_up}/4`
}
