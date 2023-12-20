//const position = document.getElementById('position')
const loginBtn = document.getElementById('login');
const userDiv = document.querySelector('#loginbox');
const passDiv = document.querySelector('#passbox');
//const logoutBtn = document.getElementById('logout');

loginBtn.addEventListener('click',()=>{
    let username = document.getElementById('username').value;
    let password = document.getElementById('password').value;
    var credentials = {
        username: username,
        password: password
    };
    if(username == ""){
        let errorMessage = document.createElement("p");
        errorMessage.textContent = "Username can't be empty";
        errorMessage.classList.add("co-red-600","font-bold");
        userDiv.insertBefore(errorMessage,userDiv.firstChild.nextSibling);
        //alert("Username can't be empty");
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
        const sessionToken = data.session;
        document.cookie = `session_id=${sessionToken}; path=/; HttpOnly=true; SameSite=lax`;
        //localStorage.setItem('session token',data.session);
        //sessionStorage.setItem('session token',data.session);

        if(data.user == 'logged in as Employee'){
            window.location.href = 'EmployeeHome.html';
        }else if (data.user == 'logged in as Manager'){
            window.location.href = 'ManagerHome.html';
        }
    })
    .catch(error => {
        console.log('error during login',error);
    })
})


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

