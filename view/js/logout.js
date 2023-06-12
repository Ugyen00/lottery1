function logout(){
    fetch("/logout")
    .then(resp =>{
        if (resp.ok){
            window.location.href = "index.html"
        }else{
            throw new Error(resp.statusText)
        }
    })
    .catch(e =>{
        alert(e)
    })
}