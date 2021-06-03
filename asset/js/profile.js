const Btn_Edit = document.getElementById("Edit_profil_btn")
const Overview = document.getElementById("Overview")
const Profile = document.getElementById("Manage_Profile")
const Profile_Close = document.getElementById("Close")
const Modify_username = document.getElementById("user_btn")
const Modify_email = document.getElementById("mail_btn")
const Modify_pwd = document.getElementById("password_btn")


const Profil_Mod = document.getElementById("Special_modif")
const field_name = document.getElementById("Name_Modif_1")
const filed_describ = document.getElementById("Name_Modif_2")
const first_label = document.getElementById("Label_Input_modif")
const Second_label = document.getElementById("Label_Input__curt_pwd")
const Exit = document.getElementById("Exit")



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
    Second_label.innerHTML = "Current Password"
    Exit.addEventListener("click", ()=>{
        Profile.style.display = "block"
    Profil_Mod.style.display = "none"
    })
})

Modify_email.addEventListener("click", ()=>{
    Profile.style.display = "none"
    Profil_Mod.style.display = "block"
    field_name.innerHTML ="Email"
    filed_describ.innerHTML = "Write your new email and current password"
    first_label.innerHTML = "Email"
    Second_label.innerHTML = "Current Password"
    Exit.addEventListener("click", ()=>{
        Profile.style.display = "block"
    Profil_Mod.style.display = "none"
    })
})

Modify_pwd.addEventListener("click", ()=>{
    Profile.style.display = "none"
    Profil_Mod.style.display = "block"
    field_name.innerHTML ="Password"
    filed_describ.innerHTML = "Write two times your new password"
    first_label.innerHTML = "New Password"
    Second_label.innerHTML = "Verification New Password"
    Exit.addEventListener("click", ()=>{
        Profile.style.display = "block"
    Profil_Mod.style.display = "none"
    })
})