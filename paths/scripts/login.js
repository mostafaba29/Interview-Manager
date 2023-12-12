//const position = document.getElementById('position')
const loginBtn = document.getElementById('login');
const logoutBtn = document.getElementById('logout');

loginBtn.addEventListener('click',()=>{
    let username = document.getElementById('username').value;
    let password = document.getElementById('password').value;
    var credentials = {
        username: username,
        password: password
    };
    fetch('http://localhost:8000/auth/login',{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(credentials), 
    })
    .then(response => response.json())
    .then(data =>{
        console.log(data);
    })
    .catch(error => {
        console.log('error durin login',error);
    })
})

logoutBtn.addEventListener('click',()=>{
    fetch('http://localhost:8000/auth/logout',{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
    })
    .then(response => response.json())
    .then(data =>{
        console.log(data);
    })
    .catch(error => {
        console.log('error durin logout',error);
    })
})
