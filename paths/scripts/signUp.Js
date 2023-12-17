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
    fetch('http://localhost:3000/signup',{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(credentials), 
    })
    .then(response => response.json())
    .then(data =>{
        if(data == 'signed up' && position == 'Manager'){
            alert('signed up as manager');
            window.location.href = 'ManagerHome.html';
        }else if(data == 'signed up' && position == 'Employee'){
            alert('signed up as employee');
            window.location.href = 'EmployeeHome.html';
        }else{
            alert(data);
        }
    })
    .catch(error => {
        console.log('error during sign up',error);
    })
})