async function fieldValidation() {
  const url = window.location.origin + "/login"

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

 fetch(url, {
    method: 'POST',
    mode: 'cors',
    cache: 'no-cache',
    credentials: 'same-origin',
    headers: {
      'Content-Type': 'application/json'
    },
    redirect: 'follow',
    referrerPolicy: 'no-referrer',
    body: JSON.stringify(data)
  }).then((response) => {
     console.log(data)
    return response.json();
  })
      .then((data) => {

        function errorMessage (object){
          for (let i = 0; i < object.length; i++){
              if(!object[i].FieldValue){
                  document.querySelector('.firstNameErr').innerHTML = "";
                  document.querySelector('.lastNameErr').innerHTML = "";
                  document.querySelector('.passwordErr').innerHTML = "";
                  document.querySelector('.emailErr').innerHTML = "";
                  window.location.href = 'confirmedRegistration.html'
                  return
              }

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
