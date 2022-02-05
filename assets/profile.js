async function checkToken() {
    const url = window.location.origin + "/send-form"

    // const Email = document.querySelector('#emailId')
    // const Password = document.querySelector('#passwordId')
    //

    const token = localStorage.getItem('token');
    console.log(token)


    const data = {
        // Email: Email.value,
        // Password: Password.value,

    }

    fetch(url, {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: {
            'Authorization': token,
            'Content-Type': 'application/json'
        },
        redirect: 'follow',
        referrerPolicy: 'no-referrer',
        body: JSON.stringify(data)
    }).then((response) => {
        return response.json();
    })
        .then((data) => {

            console.log( data)


        });
}


function ready() {
    const buttonCheckToken = document.querySelector('#send-form');
    const buttonLogOut = document.querySelector('#log-out');
    buttonCheckToken.addEventListener('click', checkToken);
    buttonLogOut.addEventListener('click', logOut);
}

document.addEventListener('DOMContentLoaded', ready);


async function logOut () {
    const url = window.location.origin + "/log-out"

    const token = localStorage.getItem('token');

    const data = {

    }

    fetch(url, {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: {
            'Authorization': token,
            'Content-Type': 'application/json'
        },
        redirect: 'follow',
        referrerPolicy: 'no-referrer',
        body: JSON.stringify(data)
    }).then((response) => {
        return response.json();
    })
        .then((data) => {

            console.log( data)


        });


    // window.location.href = 'index.html'
}