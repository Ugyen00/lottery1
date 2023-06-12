if (document.cookie == ""){
    alert("User not logged in!!")
    window.open("index.html","_self")
}else{
    console.log(document.cookie)
    console.log("cookie set")
}