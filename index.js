function errorFirstName() {
  const input = document.querySelector('#FirstNameId');
  const errorDiv = input.nextSibling;
  if (input.value === '') {
    errorDiv.innerHTML = '';
    return false;
  }
  if (input.value.length > 255) {
    errorDiv.innerHTML = 'max 255';
    return false;
  }
  errorDiv.innerHTML = '';
  return true;
}

function errorLastName() {
  const input = document.querySelector('#LastNameId');
  const errorDiv = input.nextSibling;
  if (input.value === '') {
    errorDiv.innerHTML = '';
    return false;
  }
  if (input.value.length > 255) {
    errorDiv.innerHTML = 'max 255';
    return false;
  }
  errorDiv.innerHTML = '';
  return true;
}

function passwordError() {
  // const input = event.target;
  const input = document.querySelector('#passwordId');
  const errorDiv = input.nextSibling;
  if (input.value === '') {
    errorDiv.innerHTML = '';
    return false;
  }
  if (input.value.length < 8 || input.value.length > 64) {
    errorDiv.innerHTML = 'error';
    return false;
  }
  errorDiv.innerHTML = '';
  return true;
}

function errorEmail() {
  const input = document.querySelector('#emailId');
  const errorDiv = input.nextSibling;
  if (input.value === '') {
    errorDiv.innerHTML = '';
    return false;
  }
  if (input.validity.typeMismatch === true) {
    errorDiv.innerHTML = 'email is invalid';
    return false;
  }
  errorDiv.innerHTML = '';
  return true;
}

function fieldValidation() {
  if (errorEmail() && passwordError() && errorFirstName() && errorLastName()) {
    window.location.href = 'confirmedRegistration.html';
  }
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

  const button = document.querySelector('#button');
  button.addEventListener('click', fieldValidation);
}

document.addEventListener('DOMContentLoaded', ready);
