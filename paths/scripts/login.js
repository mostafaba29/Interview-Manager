//const position = document.getElementById('position')
const loginBtn = document.getElementById('login');
//const logoutBtn = document.getElementById('logout');

loginBtn.addEventListener('click',()=>{
    let username = document.getElementById('username').value;
    let password = document.getElementById('password').value;
    var credentials = {
        username: username,
        password: password
    };
    if(username == ""){
        alert("Username can't be empty");
    }
    if(password == ""){
        alert("Password can't empty");
    }
    fetch('http://localhost:3000/auth/login',{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(credentials), 
    })
    .then(response => response.json())
    .then(data =>{
        if(data == 'logged in'){
            alert('logged in');
            //sessionStorage(token);
            window.location.href = 'EmployeeHome.html';
        }else{
            alert(data);
        }*/
    })
    .catch(error => {
        console.log('error during login',error);
    })
})

function sessionStorage(token){
    localStorage.setItem('token',token);
}

/*logoutBtn.addEventListener('click',()=>{
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
        console.log('error during logout',error);
    })
})*/
