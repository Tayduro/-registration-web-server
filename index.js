function fieldValidation() {
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
  const button = document.querySelector('#send-form');
  button.addEventListener('click', fieldValidation);
}

document.addEventListener('DOMContentLoaded', ready);
