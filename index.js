async function fieldValidation() {
  const url = "http://127.0.0.1:8034/login"

  const FirstName = document.querySelector('#FirstNameId')
  const LastName = document.querySelector('#LastNameId')
  const Email = document.querySelector('#emailId')
  const Password = document.querySelector('#passwordId')


  const data = {
    FirstName: FirstName.value,
    LastName: LastName.value,
    Email: Email.value,
    Password: Password.value,
  }

  const response = fetch(url, {
    method: 'POST', // *GET, POST, PUT, DELETE, etc.
    mode: 'cors', // no-cors, *cors, same-origin
    cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
    credentials: 'same-origin', // include, *same-origin, omit
    headers: {
      'Content-Type': 'application/json'
      // 'Content-Type': 'application/x-www-form-urlencoded',
    },
    redirect: 'follow', // manual, *follow, error
    referrerPolicy: 'no-referrer', // no-referrer, *client
    body: JSON.stringify(data) // body data type must match "Content-Type" header
  }).then((response) => {
    return response.json();
  })
      .then((data) => {

        function errorMessage (object){
          for (let i = 0; i < object.length; i++){
            console.log(i)
           console.log(object[i].FieldValue)
            console.log(object[i].ErrMassage)

            if(object[i].FieldValue === "FirstName"){
              document.querySelector('.firstNameErr').innerHTML = object[i].ErrMassage
            }

            if(object[i].FieldValue === "LastName"){
              document.querySelector('.lastNameErr').innerHTML = object[i].ErrMassage
            }

            if(object[i].FieldValue === "Password"){
              document.querySelector('.passwordErr').innerHTML = object[i].ErrMassage
            }
            if(object[i].FieldValue === "Email"){
              document.querySelector('.emailErr').innerHTML = object[i].ErrMassage
            }
          }

        }
        errorMessage (data)
      });

}


function ready() {
  const button = document.querySelector('#send-form');
  button.addEventListener('click', fieldValidation);
}

document.addEventListener('DOMContentLoaded', ready);
