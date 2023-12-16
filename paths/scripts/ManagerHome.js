const showBtn = document.getElementById('showBtn');

showBtn.addEventListener('click',()=>{
    fetch('http://localhost:3000/allappointments')
        .then(response => {
            // Check if the response status is OK (status code 200-299)
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            
            // Parse the response body as JSON
            return response.json();
        })
        .then(data => {
            let appointmentTable = document.getElementById('appointmentTable');
            let nRow = document.createElement('tr');
            appointmentTable.appendChild(nRow);
            let t = document.createElement('td');
            t=data.field1;
            nRow.appendChild(t);
            let n = document.createElement('td');
            n=data.field2;
            nRow.appendChild(n);
            let r = document.createElement('td');
            r=data.field3;
            nRow.appendChild(r);
            
        })
        .catch(error => {
            // Log and handle errors
            console.error('Error:', error);
        });
})