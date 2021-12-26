// function errorFirstName() {
//   const input = document.querySelector('#FirstNameId');
//   const errorDiv = input.nextSibling.nextSibling;
//   if (input.value === '') {
//     errorDiv.innerHTML = '';
//     return false;
//   }
//   if (input.value.length > 255) {
//     errorDiv.innerHTML = 'max 255';
//     return false;
//   }
//   errorDiv.innerHTML = '';
//   return true;
// }
//
// function errorLastName() {
//   const input = document.querySelector('#LastNameId');
//   const errorDiv = input.nextSibling.nextSibling;
//   if (input.value === '') {
//     errorDiv.innerHTML = '';
//     return false;
//   }
//   if (input.value.length > 255) {
//     errorDiv.innerHTML = 'max 255';
//     return false;
//   }
//   errorDiv.innerHTML = '';
//   return true;
// }
//
// function passwordError() {
//   const input = document.querySelector('#passwordId');
//   const errorDiv = input.nextSibling.nextSibling;
//   console.log(errorDiv)
//   if (input.value === '') {
//     errorDiv.innerHTML = '';
//     return false;
//   }
//   if (input.value.length < 8 || input.value.length > 64) {
//     errorDiv.innerHTML = 'error';
//     return false;
//   }
//   errorDiv.innerHTML = '';
//   return true;
// }
// function errorEmail() {
//   const input = document.querySelector('#emailId');
//   const errorDiv = input.nextSibling.nextSibling;
//   if (input.value === '') {
//     errorDiv.innerHTML = '';
//     return false;
//   }
//   if (input.validity.typeMismatch === true) {
//     errorDiv.innerHTML = 'email is invalid';
//     return false;
//   }
//   errorDiv.innerHTML = '';
//   return true;
// }


function fieldValidation() {
  // if (errorEmail() && passwordError() && errorFirstName() && errorLastName()) {
  //   window.location.href = 'confirmedRegistration.html';
  // }
  const url = "http://127.0.0.1:8033/login"

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
  });
}


function ready() {

  const input1 = document.querySelector('#FirstNameId');
  input1.addEventListener('input', errorFirstName);

  const input2 = document.querySelector('#LastNameId');
  input2.addEventListener('input', errorLastName);

  const input3 = document.querySelector('#emailId');
  input3.addEventListener('input', errorEmail);

  const input4 =  document.querySelector('#passwordId');
  input4.addEventListener('input', passwordError);

  const button = document.querySelector('#send-form');
  button.addEventListener('click', fieldValidation);
}

document.addEventListener('DOMContentLoaded', ready);
