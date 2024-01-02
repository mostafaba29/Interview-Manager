//getting logout,form creation and submission elements
const logoutBtn = document.getElementById('logoutBtn');
const subForm = document.getElementById('formContainer');
const subBtn = document.getElementById('Create-appointment-btn')


//event handling for creating a new appointment
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
            <!-- Submit and Cancel buttons -->
            <button type="submit" class="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600">Submit</button>
            <button type="button" class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600">Cancel</button>
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
        window.location.href = "login.html";
    })
    .catch(error => {
        console.log('error during logout',error);
    })
})