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
const Save = document.getElementById("Save")
const Exit = document.getElementById("Exit")
const ToChange = document.getElementById("ToChange")

const FirstInput = document.getElementById("Input__info")
const SecondInput = document.getElementById("Input__curt_pwd")

// Display the modify page
Btn_Edit.addEventListener("click", ()=>{
    Overview.style.display = "none"
    Profile.style.display = "block"
})
// Hide the modify page
Profile_Close.addEventListener("click", ()=>{
    Profile.style.display = "none"
    Overview.style.display = "block"
})

// // all listenent depend on where the modification is up
// username
// email
// password
// exit reset all field
Modify_username.addEventListener("click", ()=>{
    Profile.style.display = "none"
    Profil_Mod.style.display = "block"
    field_name.innerHTML ="Username"
    filed_describ.innerHTML = "Write your new username and current password"
    document.getElementById("Input__info").setAttribute("type","text")
    first_label.innerHTML = "Username"
    ToChange.setAttribute("value", "Name")
    Second_label.innerHTML = "Current Password"
})

Modify_email.addEventListener("click", ()=>{
    Profile.style.display = "none"
    Profil_Mod.style.display = "block"
    field_name.innerHTML ="Email"
    filed_describ.innerHTML = "Write your new email and current password"
    document.getElementById("Input__info").setAttribute("type","text")
    first_label.innerHTML = "Email"
    ToChange.setAttribute("value", "Email")
    Second_label.innerHTML = "Current Password"
})

Modify_pwd.addEventListener("click", ()=>{
    Profile.style.display = "none"
    Profil_Mod.style.display = "block"
    field_name.innerHTML ="Password"
    filed_describ.innerHTML = "Write two times your new password"
    document.getElementById("Input__info").setAttribute("type","password")
    first_label.innerHTML = "New Password"
    Second_label.innerHTML = "Verification New Password"
    ToChange.setAttribute("value","Password")
})
Exit.addEventListener("click", ()=>{
    Profile.style.display = "block"
    Profil_Mod.style.display = "none"
    document.getElementById("Input__info").value = ""
    document.getElementById("Input__curt_pwd").value = ""
    FirstInput.style.borderColor = "#FFCB77"
    SecondInput.style.borderColor = "#FFCB77"
})
// Save.removeAttribute("disabled")
FirstInput.addEventListener('input', ()=>{
    if (ToChange.value == "Name"){
        let username_value = FirstInput.value
        // console.log(FirstInput.value)
        Username_Verification(username_value)
    }else if  (ToChange.value == "Email"){
        let email_value = FirstInput.value
        // console.log(FirstInput.value)
        Email_Verification(email_value)
    }else if (ToChange.value == "Password"){
        let pwd_value = FirstInput.value
        // console.log(FirstInput.value)
        Pwd_Verification(pwd_value)
    }
    if (SecondInput.style.borderColor == "green" && FirstInput.style.borderColor == "green"){
        Save.removeAttribute("disabled")
    }

})
SecondInput.addEventListener('input', ()=>{
    let pwd_value = SecondInput.value
    Pwd_Verification_verification(pwd_value)
        if (SecondInput.style.borderColor == "green" && FirstInput.style.borderColor == "green"){
            Save.removeAttribute("disabled")
        }
})
// Same verification than register and Login
function Username_Verification(user){
    // console.log(user, user.length)
    if(user.length < 4){
        first_label.innerHTML = "Username <strong>Not enough characters</strong>"
        FirstInput.style.borderColor = "red"
    }else if (user.length > 12){
        first_label.innerHTML = "Username <strong>Too many characters</strong>"
        FirstInput.style.borderColor = "red"
    }else{
        first_label.innerHTML = "Username"
        FirstInput.style.borderColor = "green"
    }
    user.split("").forEach((element, index)=> {
        if (element == " "){
            first_label.innerHTML = "Username <strong>Space character not allowed</strong>"
            FirstInput.style.borderColor = "red"
        }
    })
}
function Email_Verification(email){
    if(email.length < 5){
        first_label.innerHTML = "Email <strong>Not enough characters</strong>"
        FirstInput.style.borderColor = "red"
    }else{
        let verif_mail = email.split("@").length
        if (verif_mail == 2){
            let verif_point = email.split("@")[1].split(".")
            if (verif_point.length != 2){
                first_label.innerHTML = "Email <strong>End of mail invalide</strong>"
                FirstInput.style.borderColor = "red"
            }else{
                if (verif_point[1] == ""){
                    first_label.innerHTML = "Email <strong>End of mail invalide</strong>"
                    FirstInput.style.borderColor = "red"
                }else{
                    first_label.innerHTML = "Email"
                    FirstInput.style.borderColor = "green"
                }
            }
        }else{
            first_label.innerHTML = "Email <strong>Invalide number of @</strong>"
            FirstInput.style.borderColor = "red"
        }
    }
    email.split("").forEach((element, index)=> {
        let ascii = element.charCodeAt(0)
        if (ascii < 32 || ascii > 126){
            first_label.innerHTML = "Email <strong>Invalide character</strong>"
            FirstInput.style.borderColor = "red"
        }
        if (element == " "){
            first_label.innerHTML = "Email <strong>Space character not allowed</strong>"
            FirstInput.style.borderColor = "red"
        }
        if (index > 0){
            if (element == "." && email[index -1] == "."){
                first_label.innerHTML = "Email <strong>. after an other .</strong>"
                FirstInput.style.borderColor = "red"
            }
            if (element == "." && email[index -1] == "@"){
                first_label.innerHTML = "Email <strong>. after an @</strong>"
                FirstInput.style.borderColor = "red"
            }
            if (element == "@" && email[index -1] == "."){
                first_label.innerHTML = "Email <strong>. before an @</strong>"
                FirstInput.style.borderColor = "red"
            }
        }
    })
}
function Pwd_Verification(pwd){
    let to_verif = SecondInput.value
    let Lower = 0
    let Upper = 0
    let Nb = 0
    let Space_count = 0

    pwd.split("").forEach((element, index)=> {
        let ascii = element.charCodeAt(0)
        if (ascii >= 65 && ascii <= 90){
            Upper ++
        }else if (ascii >= 97 && ascii <= 122){
            Lower ++
        }else if (ascii >= 48 && ascii <= 57){
            Nb ++
        }else if(element == " "){
            Space_count ++
        }
    })

    if(pwd.length < 6){
        first_label.innerHTML = "Password <strong>Not enough characters</strong>"
        FirstInput.style.borderColor = "red"
    }else if (Space_count >0 ){
        first_label.innerHTML = "Password <strong>Space character not allowed</strong>"
        FirstInput.style.borderColor = "red"
    }else if (Lower < 1){
        first_label.innerHTML = "Password <strong>Not enough lowercase</strong>"
        FirstInput.style.borderColor = "red"
    }else if (Upper < 1){
        first_label.innerHTML = "Password <strong>Not enough uppsercase</strong>"
        FirstInput.style.borderColor = "red"
    }else if (Nb < 1){
        first_label.innerHTML = "Password <strong>Not enough number</strong>"
        FirstInput.style.borderColor = "red"
    }else if(pwd != to_verif){
        first_label.innerHTML = "Password"
        FirstInput.style.borderColor = "green"
        Second_label.innerHTML = "Password Verification <strong>not the same</strong>"
        SecondInput.style.borderColor = "red"
    }else if(pwd == to_verif){
        first_label.innerHTML = "Password"
        FirstInput.style.borderColor = "green"
        Second_label.innerHTML = "Password Verification"
        SecondInput.style.borderColor = "green"
    }
    else{
        first_label.innerHTML = "Password"
        FirstInput.style.borderColor = "green"
    }

}
function Pwd_Verification_verification(pwd){
    let to_verif = FirstInput.value
    let Lower = 0
    let Upper = 0
    let Nb = 0
    let Space_count = 0
    if(pwd.length < 6){
        Second_label.innerHTML = "Password verification <strong>Not enough characters</strong>"
        SecondInput.style.borderColor = "red"
    }
    pwd.split("").forEach((element, index)=> {
        let ascii = element.charCodeAt(0)
        if (ascii >= 65 && ascii <= 90){
            Upper ++
        }else if (ascii >= 97 && ascii <= 122){
            Lower ++
        }else if (ascii >= 48 && ascii <= 57){
            Nb ++
        }else if(element == " "){
            Space_count ++
        }
    })
    if (Space_count >0 ){
        Second_label.innerHTML = "Password verification <strong>Space character not allowed</strong>"
        SecondInput.style.borderColor = "red"
    }else if (Lower < 1){
        Second_label.innerHTML = "Password verification <strong>Not enough lowercase</strong>"
        SecondInput.style.borderColor = "red"
    }else if (Upper < 1){
        Second_label.innerHTML = "Password verification <strong>Not enough uppsercase</strong>"
        SecondInput.style.borderColor = "red"
    }else if (Nb < 1){
        Second_label.innerHTML = "Password verification <strong>Not enough number</strong>"
        SecondInput.style.borderColor = "red"
    }else if(pwd != to_verif && ToChange.value == "Password"){
        Second_label.innerHTML = "Password Verification <strong>not the same</strong>"
        SecondInput.style.borderColor = "red"
    }
    else if(pwd == to_verif  && ToChange.value == "Password"){
        first_label.innerHTML = "Password"
        FirstInput.style.borderColor = "green"
        Second_label.innerHTML = "Password Verification"
        SecondInput.style.borderColor = "green"
    }
    else{
        Second_label.innerHTML = "Password verification"
        SecondInput.style.borderColor = "green"
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