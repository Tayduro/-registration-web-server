async function signIn() {
<<<<<<< HEAD
    const url = window.location.origin + "/entrance"
=======
    const url = window.location.origin + "/login"
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3

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
<<<<<<< HEAD
=======
            console.log( data)
            console.log( data.length)
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3
            if(data.length > 0){
                window.location.href = 'profile.html'
            }
        });
}


function ready() {
<<<<<<< HEAD
    const buttonSignIn = document.querySelector('#sign-in');
=======
    // const buttonSignUp = document.querySelector('#send-form');
    const buttonSignIn = document.querySelector('#sign-in');
    // buttonSignUp.addEventListener('click', signUp);
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3
    buttonSignIn.addEventListener('click', signIn);
}

document.addEventListener('DOMContentLoaded', ready);

<<<<<<< HEAD
=======

// async function signUp () {
//     window.location.href = 'index.html'
// }
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3
