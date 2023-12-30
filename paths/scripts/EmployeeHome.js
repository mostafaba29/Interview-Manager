
const logoutBtn = document.getElementById('logoutBtn');
const subForm = document.getElementById('formContainer');
const subBtn = document.getElementById('Create-appointment-btn')
/*document.addEventListener("DOMContentLoaded",()=>{
    subForm.classList.toggle('hidden');
    subForm.classList.add('invisible');
    subBtn.addEventListener('click',()=>{
        subForm.classList.toggle('hidden');
        document.body.classList.togglte('blur-sm');
    });
    const myForm =document.getElementById('myForm');
    myForm.addEventListener('submit',function(event){
        event.preventDefault();
        subForm.classList.add('hidden');
        document.body.classList.toggle('blur-sm');
    })
})*/

subBtn.addEventListener('click',()=>{
    //document.body.classList.add('blur-sm');
    const formContainer = document.createElement('div');
    formContainer.id ='formContainer';
    formContainer.classList.add('flex','fixed','top-0','left-0','w-full','h-full','items-center','justify-center','z-50');
    const formContent = document.createElement('div');
    formContent.classList.add('bg-white','p-8','rounded','shadow-md');
    formContent.innerHTML=`
    <h2 class="text-2xl font-bold mb-4">Create Appointment</h2>
          <form id="createAppointmentForm">
            <!-- Time input -->
            <label for="timeInput" class="block mb-2">Time:</label>
            <input type="time" id="timeInput" name="time" class="w-full border p-2 mb-4">

            <!-- Name input -->
            <label for="nameInput" class="block mb-2">Name:</label>
            <input type="text" id="nameInput" name="name" class="w-full border p-2 mb-4">

            <!-- Reason input -->
            <label for="reasonInput" class="block mb-2">Reason:</label>
            <textarea id="reasonInput" name="reason" rows="4" class="w-full border p-2 mb-4"></textarea>
            <!-- Mangar Name -->
            <label for="managerName" class="block mb-2">Manager Name:</label>
            <input type="text" id ="managerName" name="managerName" class="w-full border p-2 mb-4">
            <!-- Submit button -->
            <div class="flex justify-end">
            <button type="submit" class="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600">Submit</button>
            <button type="button" class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600">cancel</button>
            </div>
          </form>
    `;
    formContainer.appendChild(formContent);
    document.body.appendChild(formContainer);
    //formContainer.classList.remove()
    const createAppointmentForm = document.getElementById("createAppointmentForm");
    const cancelBtn = formContainer.querySelector('button[type="button"]');
    cancelBtn.addEventListener('click', function() {
      formContainer.remove();
    })
        createAppointmentForm.addEventListener("submit", async function (event) {
          event.preventDefault();
          const formData = new FormData(createAppointmentForm);
          const time = formData.get("time");
          const name = formData.get("name");
          const reason = formData.get("reason");
          const mName =formData.get("managerName");
          let appointmentData = {
            Meeting_Time:time,
            client:name,
            Description:reason,
            Manager_Name:mName
          }
            fetch('http://localhost:5000/auth/create',{
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials:"include",
                body: JSON.stringify(appointmentData), 
            })
            .then(response => response.json())
            .then(data =>{
                console.log(data);
            })
            .catch(error => {
                console.log('error creating appointment',error);
            })
          formContainer.remove();

          try {
            const response = await fetch('http://localhost:5000/auth/allappointments');
            const data = await response.json();
            renderTable(data);
          } catch (error) {
            console.error('Error', error);
          }
        });
    })
    /*function renderTable(data){
        const tableBody = document.querySelector('#appointmentTable');
        tableBody.innerHTML = '';
        if (data && data.length > 0){
        data.slice(0, 20).forEach(row => {
          const newRow = document.createElement('tr');
          newRow.innerHTML = `
              <td>${row.meeting_time}</td>
              <td>${row.client}</td>
              <td>${row.description}</td>
              <td>${row.status}</td>
              <td>${row.manager_name}</td>
          `;
          tableBody.appendChild(newRow);
      });
    }
  }*/
/*CreateAppointment.addEventListener('click', ()=>{
    let appointmentTable = document.getElementById('appointmentTable');
    let nRow = document.createElement('tr');
    appointmentTable.appendChild(nRow);
    for(let i=0;i<4;i++){
        x = document.createElement('td');
        x.textContent = prompt('typesomeinput','defaultValue');
        nRow.appendChild(x)
    }
    let cells = nRow.getElementsByTagName('td');

    let data = {
        Meeting_Time:cells[0].textContent,
        Client:cells[1].textContent,
        Description:cells[2].textContent,
        Manager_Name:cells[3].textContent
    };

    fetch('http://localhost:3000/auth/create',{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data), 
    })
    .then(response => response.json())
    .then(data =>{
        console.log(data);
    })
    .catch(error => {
        console.log('error creating appointment',error);
    })

})*/

logoutBtn.addEventListener('click',()=>{
    fetch('http://localhost:5000/auth/logout',{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
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