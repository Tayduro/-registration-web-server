async function signIn() {

    const url = window.location.origin + "/entrance"

    const Email = document.querySelector('#emailId')
    const Password = document.querySelector('#passwordId')
    //

    const data = {
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
        return response.json();
    })
        .then((data) => {

            localStorage.setItem('token', data)

            if(data.length > 0){
                window.location.href = 'profile.html'
            }
        });
}


function ready() {
    const buttonSignIn = document.querySelector('#sign-in');
    buttonSignIn.addEventListener('click', signIn);
}

document.addEventListener('DOMContentLoaded', ready);





