const CreateAppointment = document.getElementsByClassName('Create-appointment-btn')[0];

CreateAppointment.addEventListener('click', ()=>{
    let appointmentTable = document.getElementById('appointmentTable');
    let nRow = document.createElement('tr');
    appointmentTable.appendChild(nRow);
    for(let i=0;i<5;i++){
        x = document.createElement('td');
        x.textContent = prompt('typesomeinput','defaultValue');
        nRow.appendChild(x)
    }
    fetch('http://localhost:8000/create',{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(nRow), 
    })
    .then(response => response.json())
    .then(data =>{
        console.log(data);
    })
    .catch(error => {
        console.log('error creating appointment',error);
    })

})