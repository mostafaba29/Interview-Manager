//getting input and signup button
const signupBtn = document.getElementById('signup');
const userDiv = document.getElementById('userbox');
const passDiv = document.getElementById('passbox')
//event handling for signup
signupBtn.addEventListener('click',()=>{
    let username = document.getElementById('username').value;
    let password = document.getElementById('password').value;
    let position = document.getElementById('position').value;
    localStorage.setItem('user',username);
    localStorage.setItem('password',password);
    if(username == ""){
        let errorMessage = document.createElement("p");
        errorMessage.textContent = "Username can't be empty";
        errorMessage.classList.add("text-red-600","font-bold");
        userDiv.insertBefore(errorMessage,userDiv.firstChild.nextSibling);
        //alert("Username can't be empty");
    }
    if(password == ""){
        let passErrorMessage = document.createElement("p");
        passErrorMessage.textContent = "Password can't be empty";
        passErrorMessage.classList.add("text-red-600","font-bold");
        passDiv.insertBefore(passErrorMessage,passDiv.firstChild.nextSibling);
    }
    //console.log(password);
    var credentials = {
        username: username,
        password: password,
        position: position
    };
    fetch('http://localhost:5000/signup',{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        credentials:"include",
        body: JSON.stringify(credentials), 
    })
    .then(response => response.json())
    .then(data =>{
        if((data.message == 'account created' && data.position == 'Manager')||(data.message == 'account created' && data.position == 'Employee')){
            //alert('signed up as Manager');
            window.location.href = 'login.html';
        }else{
            console.log(data);
        }
    })
    .catch(error => {
        console.log('error during sign up',error);
    })
})