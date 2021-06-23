const LikePost = document.getElementById("Like-Button")
LikePost.value = "non"
const likeForm = document.getElementById("Likeform")
const PostModif = document.getElementById("modify-post")
const modif = document.getElementById("modify")
ADD_Like()
function ADD_Like(){
    LikePost.addEventListener("click",()=>{
        if(LikePost.value == "non"){
            console.log("false to true");
            LikePost.value = "oui";
            LikePost.style.color = "#FFCB77"
        }else if (LikePost.value == "oui"){
            console.log("true to false");
            LikePost.value = "non";
            LikePost.style.color = "#c29958"
        }
    })
}
likeForm.addEventListener("submit",function(){
    fetch('/like', {
        headers : {
            'Accept': 'application/json',
            'Content-Type':'application/json'
        },
        body:JSON.stringify({
            Like: LikePost.value
        })  
    })
    .then(function(response){
        return response.text
    })
    .catch(function(error){
        console.error(error)
    })
})
PostModif.addEventListener("click",function(){
    modif.style.display = "block"
})


