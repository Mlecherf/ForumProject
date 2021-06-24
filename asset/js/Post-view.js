const LikePost = document.getElementById("Like-Button")
LikePost.value = "non"
const PostModif = document.getElementById("modify-post")
const modif = document.getElementById("pop_post_modify")
// ADD_Like()
// function ADD_Like(){
//     LikePost.addEventListener("click",()=>{
//         if(LikePost.value == "non"){
//             console.log("false to true");
//             LikePost.value = "oui";
//             LikePost.style.color = "#FFCB77"
//         }else if (LikePost.value == "oui"){
//             console.log("true to false");
//             LikePost.value = "non";
//             LikePost.style.color = "#c29958"
//         }
//         fetch('/like', {
//             headers : {
//                 'Accept': 'application/json',
//                 'Content-Type':'application/json'
//             },
//             body:JSON.stringify({
//                 Like: LikePost.value
//             })  
//         })
//         .then(function(response){
//             return response.text
//         })
//         .catch(function(error){
//             console.error(error)
//         })
//     })
// }
const NameMod = document.getElementById("post_name_mod");
const ContentMod = document.getElementById("post_content_mod");



PostModif.addEventListener("click",function(){
    modif.style.display = "block"
    NameMod.value = document.getElementById("User-Post-Name").innerHTML
    ContentMod.value = document.getElementById("text").innerHTML
    document.getElementById("label_name_mod").innerHTML = `Name : character ${NameMod.value.length}/25`
    document.getElementById("label_content_mod").innerHTML = `Description : character ${ContentMod.value.length}/2000`
    document.getElementById("post_submit_mod").innerHTML = "Change"
})
document.getElementsByClassName("closemod")[0].addEventListener("click",()=>{
    modif.style.display = "none"
    document.getElementById("label_name_mod").innerHTML = "Name : character 0/25"
    NameMod.value = ""
    NameMod.style.borderColor = "#FFCB77"
    document.getElementById("label_content_mod").innerHTML = `Description : character 0/2000`
    ContentMod.value = ""
    ContentMod.style.borderColor = "#FFCB77"
    document.getElementById("post_submit_mod").setAttribute("disabled",true)
})

// --name input management--
NameMod.addEventListener("input", ()=>{
    console.log("name")
    console.log(document.getElementById("label_name_mod").innerHTML)
    if (NameMod.value.length < 4){
        document.getElementById("label_name_mod").innerHTML = `Name : not enough character ${NameMod.value.length}/25`
        NameMod.style.borderColor = "red"
    }else if(NameMod.value.length > 25){
        document.getElementById("label_name_mod").innerHTML = `Name : too many character ${NameMod.value.length}/25`
        NameMod.style.borderColor = "red"
    }else{
        document.getElementById("label_name_mod").innerHTML = `Name : character ${NameMod.value.length}/25`
        NameMod.style.borderColor = "green"
    }
    ADD_enabled()
})
// content post management
ContentMod.addEventListener("input", ()=>{
    console.log("content")
    console.log(document.getElementById("label_content_mod").innerHTML)
    if (ContentMod.value.length < 4){
       document.getElementById("label_content_mod").innerHTML = `Description : not enough character ${ContentMod.value.length}/2000`
       ContentMod.style.borderColor = "red"
    }else if(ContentMod.value.length > 2000){
       document.getElementById("label_content_mod").innerHTML = `Description : too many character ${ContentMod.value.length}/2000`
       ContentMod.style.borderColor = "red"
    }else{
       document.getElementById("label_content_mod").innerHTML = `Description : character ${ContentMod.value.length}/2000`
       ContentMod.style.borderColor = "green"
    }
    ADD_enabled()
})
function ADD_enabled (){
    if ((NameMod.style.borderColor == "green")&&(ContentMod.style.borderColor == "green")){
        document.getElementById("post_submit_mod").removeAttribute("disabled")
    }else{
        document.getElementById("post_submit_mod").setAttribute("disabled",true)
    }
} 
document.getElementById("myFormMod").addEventListener('submit',function (e){
    modif.style.display = "none"
    e.preventDefault()
    fetch('/modifpost', {
        method: 'post',
        headers : {
            'Accept': 'application/json',
            'Content-Type':'application/json'
        },
        body:JSON.stringify({
            Name: NameMod.value,
            Content: ContentMod.value,
            Id_post: document.getElementById("Post_id_modif").value
        })
    })
    .then(function(response){
        return response.text
    })
    .catch(function(error){
        console.error(error)
    })
    document.getElementById("label_name_mod").innerHTML = "Name : character 0/25"
    NameMod.value = ""
    NameMod.style.borderColor = "#FFCB77"
    document.getElementById("label_content_mod").innerHTML = `Description : character 0/2000`
    ContentMod.value = ""
    ContentMod.style.borderColor = "#FFCB77"
    document.getElementById("post_submit_mod").setAttribute("disabled",true)
})
<<<<<<< HEAD
const scrollbar = document.getElementById("text")

scrollbar.addEventListener("click",()=>{
    if(NameMod.value.length<280){
        scrollbar.style.overflowY = "hidden"
    }else{
        scrollbar.style.overflowY = "scroll"
    }
})


Tags = [...document.getElementsByClassName("Tag")]
Tags.forEach(element => {
    console.log(element.innerText)
    console.log(element.innerText === "America_Latina")
    if (element.innerText == "America_Latina"){
        element.innerText = "America Latina"
    }
    if (element.innerText == "Fast_Food"){
        element.innerText = "Fast Food"
    }
});
=======
>>>>>>> e77d9cc647cdccadc0d7a2dbfc7af5f92f8359a3
