const Btn_Edit = document.getElementById("Edit_profil_btn")
const Overview = document.getElementById("Overview")
const Profile = document.getElementById("Manage_Profile")
const Profile_Close = document.getElementById("Close")
const Modify_username = document.getElementById("user_btn")
const Modify_email = document.getElementById("mail_btn")
const Modify_pwd = document.getElementById("password_btn")

// const ToChange = document.getElementById("Tochange")


const Profil_Mod = document.getElementById("Special_modif")
const field_name = document.getElementById("Name_Modif_1")
const filed_describ = document.getElementById("Name_Modif_2")
const first_label = document.getElementById("Label_Input_modif")
const Second_label = document.getElementById("Label_Input__curt_pwd")
const Exit = document.getElementById("Exit")
const ToChange = document.getElementById("ToChange")


Btn_Edit.addEventListener("click", ()=>{
    Overview.style.display = "none"
    Profile.style.display = "block"
})

Profile_Close.addEventListener("click", ()=>{
    Profile.style.display = "none"
    Overview.style.display = "block"
})

Modify_username.addEventListener("click", ()=>{
    Profile.style.display = "none"
    Profil_Mod.style.display = "block"
    field_name.innerHTML ="Username"
    filed_describ.innerHTML = "Write your new username and current password"
    first_label.innerHTML = "Username"
    first_label.setAttribute("name", "username")
    first_label.setAttribute("value", "Name")
    ToChange.setAttribute("value", "Name")
    Second_label.innerHTML = "Current Password"
})

Modify_email.addEventListener("click", ()=>{
    Profile.style.display = "none"
    Profil_Mod.style.display = "block"
    field_name.innerHTML ="Email"
    filed_describ.innerHTML = "Write your new email and current password"
    first_label.innerHTML = "Email"
    first_label.setAttribute("name", "email")
    first_label.setAttribute("value", "Email")
    ToChange.setAttribute("value", "Email")
    Second_label.innerHTML = "Current Password"
})

Modify_pwd.addEventListener("click", ()=>{
    Profile.style.display = "none"
    Profil_Mod.style.display = "block"
    field_name.innerHTML ="Password"
    filed_describ.innerHTML = "Write two times your new password"
    first_label.innerHTML = "New Password"
    first_label.setAttribute("name", "new_password")
    Second_label.innerHTML = "Verification New Password"
    ToChange.setAttribute("value","Password")
    document.getElementById("Input__info").setAttribute("type","password")
    document.getElementById("Input__curt_pwd").setAttribute("type","password")
})
Exit.addEventListener("click", ()=>{
    Profile.style.display = "block"
    Profil_Mod.style.display = "none"
    document.getElementById("Input__info").value = ""
    document.getElementById("Input__curt_pwd").value = ""
})

// let my_cookie = Select_Login_cookie()
// if (my_cookie != ""){
//     my_cookie = JSON.parse(Cookie_cooker(my_cookie))
//     console.log(my_cookie)

//     document.getElementById("Post_stat").innerHTML = `Post : <strong>${my_cookie.nb_posts}</strong>`
//     document.getElementById("Liked_stat").innerHTML = `Liked Post : <strong>${my_cookie.nb_likes}</strong>`

//     document.getElementById('Personal_user').innerHTML = `${my_cookie.user}`
//     document.getElementById('Personal_email').innerHTML = `${my_cookie.mail}`
//     document.getElementById("DeleteInput").setAttribute("name", `${my_cookie.mail}`)

// }

// function Select_Login_cookie (){
//     let my_cookie_login = ""

//     document.cookie.split("; ").forEach((elem)=>{
//         // console.log(elem.slice(0, 5))
//         // console.log(elem)
//         if (elem.slice(0,5) == "Login"){
//             my_cookie_login = elem.slice(7,-1)
//         }
//     })
//     // console.log(my_cookie_login)
//     return my_cookie_login
// }

// function Cookie_cooker (initial_cookie){
//     const new_hot_cookie = initial_cookie.split("")
//     new_hot_cookie.forEach((element, index) => {
//         if (element == "'"){
//             new_hot_cookie[index] = '"'
//         }
//     });
//     return new_hot_cookie.join("")
// }
