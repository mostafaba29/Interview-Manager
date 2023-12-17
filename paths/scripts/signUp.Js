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
    fetch('http://localhost:3000/auth/signup',{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(credentials), 
    })
    .then(response => response.json())
    .then(data =>{
        if(data == 'account created' && position == 'Manager'){
            alert('signed up as Manager');
            window.location.href = 'ManagerHome.html';
        }else if(data == 'account created' && position == 'Employee'){
            alert('signed up as Employee');
            window.location.href = 'EmployeeHome.html';
        }else{
            alert(data);
        }
    })
    .catch(error => {
        console.log('error during sign up',error);
    })
})