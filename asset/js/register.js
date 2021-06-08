const Email = document.getElementById("Email__input")
const Username = document.getElementById("Username__input")
const Password = document.getElementById("Password__input")
const Password_Verif = document.getElementById("Password_Verif__input")
const Submit_btn = document.getElementsByClassName("submit")
const Disclaimer_Error = document.getElementsByClassName("Disclaimer_error")

Email.addEventListener("input", (event)=>{
    let email_value = Email.value
    let error_encounter = Email_Verification(email_value)
    if (email_value == ""){
        Email.style.borderColor = "red";
        Disclaimer_Error[0].innerHTML = "Invalide Email"
    }else if(error_encounter.length != 0){
        Disclaimer_Error[0].innerHTML = ""
        Email.style.borderColor = "red";
        error_encounter.forEach(elem =>{
            Disclaimer_Error[0].innerHTML += elem +"<br>"
        })
    }else{
        Email.style.borderColor = "green";
        Disclaimer_Error[0].innerHTML = ""
    }
    verif()
})

function Email_Verification (email){
    let double_mail = 0
    const input_err = []

    if (email.length < 5){
        if (input_err.includes("Too short") == false){
            input_err.push("Too short")
        }
    }else{
        let verif_mail = email.split("@").length
        if (verif_mail == 2){
            let verif_point = email.split("@")[1].split(".")
            if (verif_point.length != 2){
                if (input_err.includes("End of mail invalide") == false){
                    input_err.push("End of mail invalide")
                }
            }else{
                if (verif_point[1] == ""){
                    if (input_err.includes("End of mail invalide") == false){
                        input_err.push("End of mail invalide")
                    }
                }
            }
        }else{
            if (input_err.includes("End of mail invalide") == false){
                input_err.push("End of mail invalide")
            }
        }
    }
    email.split("").forEach((element, index)=> {
        let ascii = element.charCodeAt(0)
        if (ascii < 32 || ascii > 126){
            console.log(element)
            if (input_err.includes("Invalide character") == false){
                input_err.push("Invalide character")
            }
        }
        if (element == " "){
            if (input_err.includes("Space character not allowed") == false){
                input_err.push("Space character not allowed")
            }
        }
        if (element == "@"){
            double_mail ++
        }
        if (index > 0){
            if (element == "." && email[index -1] == "."){
                if (input_err.includes(". after an other .") == false){
                    input_err.push(". after an other .")
                }
            }
            if (element == "." && email[index -1] == "@"){
                if (input_err.includes(". after an @") == false){
                    input_err.push(". after an @")
                }
            }
            if (element == "@" && email[index -1] == "."){
                if (input_err.includes(". before an @") == false){
                    input_err.push(". before an @")
                }
            }
        }
    })
    if (double_mail > 1 ){
        if (input_err.includes("Too many @") == false){
            input_err.push("Too many @")
        }
    }else if (double_mail < 1 ){
        if (input_err.includes("Missing @") == false){
            input_err.push("Missing @")
        }
    }
    return input_err
}

Username.addEventListener("input", (event)=>{
    let username_value = Username.value
    let error_encounter = Username_Verification(username_value)
    if (error_encounter.length != 0){
        Disclaimer_Error[1].innerHTML = ""
        Username.style.borderColor = "red";
        error_encounter.forEach(elem =>{
            Disclaimer_Error[1].innerHTML += elem +"<br>"
        })
    }else{
        Username.style.borderColor = "green";
        Disclaimer_Error[1].innerHTML = ""
    }
    verif()
})

function Username_Verification (user){
    const input_err = []
    if(user.length < 4){
        if (input_err.includes("Not enough characters") == false){
            input_err.push("Not enough characters")
        }
    }else if (user.length > 12){
        if (input_err.includes("Too many characters") == false){
            input_err.push("Too many characters")
        }
    }
    user.split("").forEach((element, index)=> {
        if (element == " "){
            if (input_err.includes("Space character not allowed") == false){
                input_err.push("Space character not allowed")
            }
        }
    })
    return input_err
}

Password.addEventListener("input", (event)=>{
    let pwd_value = Password.value
    let pwd_verif_value = Password_Verif.value
    let error_encounter = Password_Verification(pwd_value)
    if (error_encounter.length != 0){
        Disclaimer_Error[2].innerHTML = ""
        Password.style.borderColor = "red";
        error_encounter.forEach(elem =>{
            Disclaimer_Error[2].innerHTML += elem +"<br>"
        })
    }else{
        Password.style.borderColor = "green";
    }
    if(pwd_verif_value == pwd_value ){
        Password_Verif.style.borderColor = "green";
        Disclaimer_Error[3].innerHTML = ""
    }
    verif()
})

function Password_Verification (user){
    Disclaimer_Error[2].innerHTML = ""
    const input_err = []
    let Lower = 0
    let Upper = 0
    let Nb = 0
    if(user.length < 6){
        if (input_err.includes("Not enough characters") == false){
            input_err.push("Not enough characters")
        }
    }
    user.split("").forEach((element, index)=> {
        let ascii = element.charCodeAt(0)
        if (element == " "){
            if (input_err.includes("Space character not allowed") == false){
                input_err.push("Space character not allowed")
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
    if (Lower < 1){
        if (input_err.includes("Not enough lowercase") == false){
            input_err.push("Not enough lowercase")
        }
    }
    if (Upper < 1){
        if (input_err.includes("Not enough uppsercase") == false){
            input_err.push("Not enough uppsercase")
        }
    }
    if (Nb < 1){
        if (input_err.includes("Not enough number") == false){
            input_err.push("Not enough number")
        }
    }
    return input_err
}

Password_Verif.addEventListener("input", (event)=>{
    let pwd_verif_value = Password_Verif.value
    let pwd_value = Password.value

    if(pwd_verif_value != pwd_value ){
        Password_Verif.style.borderColor = "red";
        Disclaimer_Error[3].innerHTML = "Invalide Password"
    }else{
        Password_Verif.style.borderColor = "green";
        Disclaimer_Error[3].innerHTML = ""
    }
    verif()
})

function verif (){
    if (Email.style.borderColor == "green" && Username.style.borderColor == "green" && Password.style.borderColor == "green"  && Password_Verif.style.borderColor == "green"){
        Submit_btn[0].removeAttribute('disabled')
    }
}


