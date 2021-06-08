const Mail = document.getElementById("Insert_a_mail")
const Password = document.getElementById("Insert_a_password")
const Submit_btn = document.getElementsByClassName("submit")
const Print_err = document.getElementsByClassName("Print_error")
// ------------------------------------------------------------
Mail.addEventListener("input", ()=>{
    let mailValue = Mail.value
    let MailError = MailVerification(mailValue)
    if (mailValue == ""){
        Mail.style.borderColor = "red";
        Print_err[0].innerHTML = "Empty eMail."
    }else if(MailError.length != 0){
        Print_err[0].innerHTML = ""
        Mail.style.borderColor = "red";
        MailError.forEach(elem =>{
            Print_err[0].innerHTML += elem +"<br>"
        })
    }else{
        Mail.style.borderColor = "green";
        Print_err[0].innerHTML = ""
    }
    verif()
})
// ------------------------------------------------------------
function MailVerification(Mail){
    let MailAt = 0
    const input_err = []
    if(Mail.length<5){
        if (input_err.includes("Invalid Email.") == false){
            input_err.push("Invalid Email.")
        }
    }else{
    let Mail_verif =Mail.split("@").length
        if (Mail_verif == 2){
            let Domain_verif = Mail.split("@")[1].split(".")
            if (Domain_verif.length != 2){
                if (input_err.includes("Invalid Email.") == false){
                    input_err.push("Invalid Email.")
                }
            }else{
                if (Domain_verif[1] == ""){
                    if (input_err.includes("Invalid Email.") == false){
                        input_err.push("Invalid Email.")
                    }
                }
            }
        }else{
            if (input_err.includes("Invalid Email.") == false){
                input_err.push("Invalid Email.")
            }
        }
    Mail.split("").forEach((element, indexMail)=> {
        if (element == " "){
            if (input_err.includes("Invalid Email.") == false){
                input_err.push("Invalid Email.")
            }
        }
        if(['à', 'ç', 'é', 'è','ê','ù'].includes(element)){
            if (input_err.includes("Invalid Email.") == false){
                input_err.push("Invalid Email.")
            }
        }
        if (element == "@"){
            MailAt ++
        }
        if (indexMail > 0){

             if (element == "." && Mail[indexMail -1] == "."){
                if (input_err.includes("Invalid Email.") == false){
                    input_err.push("Invalid Email.")
                }
            }else if(element == "@" && Mail[indexMail -1] == "."){
                if (input_err.includes("Invalid Email.") == false){
                    input_err.push("Invalid Email.")
                }
            }else if (element == "." && Mail[indexMail -1] == "@"){
                if (input_err.includes("Invalid Email.") == false){
                    input_err.push("Invalid Email.")
                    
                }
            }
        }
    })
    if (MailAt > 1 ){
        if (input_err.includes("Invalid Email.") == false){
            input_err.push("Invalid Email.")
        }
    }else if (MailAt < 1 ){
        if (input_err.includes("Invalid Email.") == false){
            input_err.push("Invalid Email.")
        }
    }
}
    return input_err
}
// ------------------------------------------------------------
Password.addEventListener("input", ()=>{
    let passwordValue = Password.value
    let PasswordError = PasswordVerification(passwordValue)
    if (PasswordError.length != 0){
        Print_err[1].innerHTML = ""
        Password.style.borderColor = "red";
        PasswordError.forEach(elem =>{
            Print_err[1].innerHTML += elem +"<br>"
        })
    }else{
        Password.style.borderColor = "green";
        Print_err[1].innerHTML = ""
    }
    verif()
})
// ------------------------------------------------------------
function PasswordVerification (user){
    const input_err = []
    let Lower = 0
    let Upper = 0
    let Nb = 0
    if(user.length < 6){
        if (input_err.includes("Invalid password.") == false){
            input_err.push("Invalid password.")
        }
    }
    user.split("").forEach((element, index)=> {
        let ascii = element.charCodeAt(0)
        if (element == " "){
            if (input_err.includes("Invalid password.") == false){
                input_err.push("Invalid password.")
            }
        }
        if (ascii >= 65 && ascii <= 90){
            Upper ++
        }else if (ascii >= 97 && ascii <= 122){
            Lower ++
        }else if (ascii >= 48 && ascii <= 57){
            Nb ++
        }
        })
    if (Lower < 1|| Upper < 1||Nb < 1){
        if (input_err.includes("Invalid password.") == false){
            input_err.push("Invalid password.")
        }
    }
    return input_err
}
// ------------------------------------------------------------
function verif (){
    if (Mail.style.borderColor == "green" && Password.style.borderColor == "green" ){
        Submit_btn[0].removeAttribute('disabled')
    }else{
        Submit_btn[0].setAttribute('disabled' , true)
    }
}
verif()