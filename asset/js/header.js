// Get the input field
var input = document.getElementById("Search__Text");

// Execute a function when the user releases a key on the keyboard
input.addEventListener("keyup", function(event) {
  // Number 13 is the "Enter" key on the keyboard
  if (event.keyCode === 13) {
    // Cancel the default action, if needed
    event.preventDefault();
    // Trigger the button element with a click
    console.log("blab")
    setTimeout(() => {
      document.getElementById("Search__BTN").removeAttribute("disabled");
      document.getElementById("Search__BTN").click();
    }, 2000);
  }
});

// document.cookie = "Login ='{'user':'Clem','mail':'mail@cookie.com','nb_posts':'10','nb_likes':'15'}'"
let my_cookie_header = Select_Login_cookie()
if (my_cookie_header != ""){
    my_cookie_header = JSON.parse(Cookie_cooker(my_cookie_header))
    console.log(my_cookie_header)
    if (my_cookie_header.user != ""){
        document.getElementsByClassName("Login")[0].setAttribute("src", "https://img.icons8.com/fluent-systems-regular/45/000000/user-male-circle.png")
        document.getElementById("Post_add").style.display = "block"
    }
}else{
    document.getElementsByClassName("Login")[0].setAttribute("src", "https://img.icons8.com/windows/50/000000/user-ninja.png")
    document.getElementById("Post_add").style.display = "none"
}

function Select_Login_cookie (){
    let my_cookie_login = ""

    document.cookie.split("; ").forEach((elem)=>{
        // console.log(elem.slice(0, 5))
        // console.log(elem)
        if (elem.slice(0,5) == "Login"){
            my_cookie_login = elem.slice(7,-1)
        }
    })
    // console.log(my_cookie_login)
    return my_cookie_login
}

function Cookie_cooker (initial_cookie){
    const new_hot_cookie = initial_cookie.split("")
    new_hot_cookie.forEach((element, index) => {
        if (element == "'"){
            new_hot_cookie[index] = '"'
        }
    });
    return new_hot_cookie.join("")
}
