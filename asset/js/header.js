// Search bar
// 
// Get the input field
const input = document.getElementById("Search__Text");
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


// Popup on click account icon
// 
document.getElementsByClassName("Login")[0].addEventListener("click", (event)=>{
    document.getElementById("Pop_up").style.display = "block"
    event.stopPropagation()
})
document.addEventListener("click", ()=>{
    document.getElementById("Pop_up").style.display = "none"
})



// Cookie gestion 
// depending on cookie value, change occur in the header 
// if connect = true |-> account image = user-male / post btn = display / gestion = Profile & Logout 
// if connect = false |-> account image = ninja / post btn = hidden / gestion = Register & Login
// 
let my_cookie_header = Select_Login_cookie()
if (my_cookie_header == "true"){
    document.getElementsByClassName("Login")[0].setAttribute("src", "https://img.icons8.com/fluent-systems-regular/45/000000/user-male-circle.png")
    document.getElementById("Post_add").style.display = "flex"
    document.getElementById("First_pop").innerHTML = "Profile"
    document.getElementById("First_pop").setAttribute("href","/profile")
    document.getElementById("Second_pop").innerHTML = "Logout"
    document.getElementById("Second_pop").setAttribute("href","/logout")
}else{
    document.getElementsByClassName("Login")[0].setAttribute("src", "https://img.icons8.com/windows/50/000000/user-ninja.png")
    document.getElementById("First_pop").innerHTML = "Register"
    document.getElementById("First_pop").setAttribute("href","/register")
    document.getElementById("Second_pop").innerHTML = "Login"
    document.getElementById("Second_pop").setAttribute("href","/login")
}
// Select Connect cookie 
function Select_Login_cookie (){
    let my_cookie_login = ""

    document.cookie.split("; ").forEach((elem)=>{
        if (elem.slice(0,7) == "Connect"){
            my_cookie_login = elem.slice(8,elem.length)
        }
    })
    return my_cookie_login
}


// Pop-up Tag
// Gestion of Post Add
// Name 4-25
// Tag 1-4
// Content 4-2000
const NameAdd = document.getElementById("post_name_add");
const ContentAdd = document.getElementById("post_content_add");
const Post_tag = [...document.getElementsByClassName('tag__post')]
const restes_tag = document.getElementById('reset__tag')
const nb_tag = document.getElementById('nb_tag')
let nb_tag_up = 0;
let Check_Tag = ""
const myForm = document.getElementById("myForm")

// Close the popup
// reinitialise all js var / html innerHTML / input border color & input value
document.getElementById("close_post").addEventListener("click", ()=>{
    document.getElementById('pop_post_add').style.display = "none"
    document.getElementById("label_name_post").innerHTML = "Name : character 0/25"
    NameAdd.value = ""
    NameAdd.style.borderColor = "#FFCB77"
    Post_tag.forEach((elem)=>{
        elem.checked = false;
        elem.value = "down"
        elem.removeAttribute("disabled")
    })
    nb_tag.innerHTML = `Tag : 0/4`
    nb_tag_up = 0
    document.getElementById("label_content_post").innerHTML = `Description : character 0/2000`
    ContentAdd.value = ""
    ContentAdd.style.borderColor = "#FFCB77"
    document.getElementById("post_submit").setAttribute("disabled",true)
})
// Open the popup
document.getElementById("Post_add").addEventListener("click", ()=>{
    document.getElementById('pop_post_add').style.display = "block"
})


// --name input management--
// looking name-input length
NameAdd.addEventListener("input", ()=>{
    if (NameAdd.value.length < 4){
        document.getElementById("label_name_post").innerHTML = `Name : not enough character ${NameAdd.value.length}/25`
        NameAdd.style.borderColor = "red"
    }else if(NameAdd.value.length > 25){
        document.getElementById("label_name_post").innerHTML = `Name : too many character ${NameAdd.value.length}/25`
        NameAdd.style.borderColor = "red"
    }else{
        document.getElementById("label_name_post").innerHTML = `Name : character ${NameAdd.value.length}/25`
        NameAdd.style.borderColor = "green"
    }
    ADD_enabled()
})
// content post management
// looking content-input length
ContentAdd.addEventListener("input", ()=>{
    if (ContentAdd.value.length < 4){
       document.getElementById("label_content_post").innerHTML = `Description : not enough character ${ContentAdd.value.length}/2000`
       ContentAdd.style.borderColor = "red"
    }else if(ContentAdd.value.length > 2000){
       document.getElementById("label_content_post").innerHTML = `Description : too many character ${ContentAdd.value.length}/2000`
       ContentAdd.style.borderColor = "red"
    }else{
       document.getElementById("label_content_post").innerHTML = `Description : character ${ContentAdd.value.length}/2000`
       ContentAdd.style.borderColor = "green"
    }
    ADD_enabled()
})
// Tag management
// Check all up tag 
// Max 4 up | Min 1 up
// If 4 up other down and block
// Reset tag btn -> all tag down
Post_tag.forEach((elem)=>{
    elem.value = "down"
})
Post_tag.forEach((elem, index)=>{
    elem.addEventListener('click', ()=>{
        if (elem.value == "down"){
            Check_Tag += `$${elem.name}`
            nb_tag_up ++
            nb_tag.innerHTML = `Tag : ${nb_tag_up}/4`
            elem.value = "up"
        }
        if (nb_tag_up == 4){
            Post_tag.forEach((elem)=>{
                if (elem.value == "down"){
                    elem.setAttribute("disabled", true)
                }
            })
        }
        ADD_enabled()
    })
})
restes_tag.addEventListener('click', ()=>{
    Post_tag.forEach((elem)=>{
        elem.checked = false;
        elem.value = "down"
        elem.removeAttribute("disabled")
    })
    nb_tag.innerHTML = `Tag : 0/4`
    nb_tag_up = 0
    Check_Tag = ""
})

// Looking to nb tag up and coolor of name & content border
// Disable or not the add btn
function ADD_enabled (){
    if ((nb_tag_up >0 && nb_tag_up <5)&&(NameAdd.style.borderColor == "green")&&(ContentAdd.style.borderColor == "green")){
        document.getElementById("post_submit").removeAttribute("disabled")
    }else{
        document.getElementById("post_submit").setAttribute("disabled",true)
    }
}
// On submit post
// Send all data to Golang server through /recup path
myForm.addEventListener('submit',function (e){
    document.getElementById("pop_post_add").style.display="none";
    e.preventDefault()
    fetch('/recup', {
        method: 'post',
        headers : {
            'Accept': 'application/json',
            'Content-Type':'application/json'
        },
        body:JSON.stringify({
            Name: NameAdd.value,
            Content: ContentAdd.value,
            Tags: Check_Tag
        })
    })
    .then(function(response){
        return response.text
    })
    .catch(function(error){
        console.error(error)
    })
    document.getElementById("label_name_post").innerHTML = "Name : character 0/25"
    NameAdd.value = ""
    NameAdd.style.borderColor = "#FFCB77"
    nb_tag.innerHTML = `Tag : 0/4`
    nb_tag_up = 0
    document.getElementById("label_content_post").innerHTML = `Description : character 0/2000`
    ContentAdd.value = ""
    ContentAdd.style.borderColor = "#FFCB77"

})