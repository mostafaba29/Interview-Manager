//constants of the buttons and the parent div
const sendBtn = document.getElementById('sendBtn');
const reciveBtn = document.getElementById('recieveBtn');
const mContainer = document.getElementById('basicFunc');

//code to enable the send button to collect the data form the input and send it via the fethc API
sendBtn.addEventListener('click',()=>{
    let inputTextbox = document.getElementById('inputMessage');
    // The data is collected in a variable and converted to JSON format
    let data = {
        textValue:inputTextbox
    };
    var jsonData = JSON.stringify(data);
    fetch('server',{
        method: 'POST',
        headers: {
            'Content-Type':'applicaton/json'
        },
        body: jsonData,
    })
    .then(response => response.json())
    .then(data => console.log('Server response:',data))
    .catch(error => console.log('Error:',error));

})
// code to recive the data and create a label to show it in the form 
reciveBtn.addEventListener('click',()=>{
    fetch('server')
    .then(response =>response.json())
    .then(data =>{
        var outputText = document.createElement("label");
        outputText.className="m-1";
        outputText.id="message";
        mContainer.insertBefore(outputText,mContainer.firstChild)
        //the recived data should be in JSON format and thus it's conveted here
        var jsonString = JSON.stringify(data);
        outputText.textContent=jsonString;
    })
    .catch(error=>console.log('error:',error));
})