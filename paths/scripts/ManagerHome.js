// getting logout and show button for manager
const logoutBtn = document.getElementById('logoutBtn');

//event handling for logout button
logoutBtn.addEventListener('click',()=>{
    fetch('http://localhost:5000/auth/logout',{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        credentials:"include",
    })
    .then(response => response.json())
    .then(data =>{
        console.log(data);
        window.location.href = "login.html";
    })
    .catch(error => {
        console.log('error during logout',error);
    })
})