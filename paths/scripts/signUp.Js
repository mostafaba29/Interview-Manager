//const position = document.getElementById('position')
const signupBtn = document.getElementById('signup');

signupBtn.addEventListener('click',()=>{
    let username = document.getElementById('username').value;
    let password = document.getElementById('password').value;
    let position = document.getElementById('position').value;
    console.log(password);
    var credentials = {
        username: username,
        password: password,
        position: position
    };
    fetch('http://localhost:8000/auth/signup',{
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