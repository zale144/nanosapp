// loginAccount sends requests to login with provided credentials
const loginAccount = (e) => {
    e.preventDefault();
    const errContainer = document.getElementById('l-error-container');
    errContainer.innerHTML = '<img src=/static/image/ajax-loader.gif />';
    let headers = new Headers();
    headers.set('Authorization', 'Basic ' + btoa(document.getElementById('l-username').value +
        ':' + document.getElementById('l-password').value));

    fetch('/login', {
        method: 'POST',
        headers: headers,
        credentials: 'same-origin'
    })
        .then(res => {
            if (!res.ok) {
                throw res.json();
            } else {
                return res.json()
            }
        })
        .then(
            (result) => {
                if (result.token) {
                    localStorage.setItem('access_token', result.token);
                    location.href = '/';
                }
            }).catch((err) => {
        err.then(data => {
            errContainer.innerHTML = data.message;
        })
    })
};

// registerAccount sends requests to register a new account
const registerAccount = (e) => {
    e.preventDefault();
    const errContainer = document.getElementById('r-error-container');
    errContainer.innerHTML = '<img src=/static/image/ajax-loader.gif />';

    fetch('/register', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            username: document.getElementById('r-username').value,
            password: document.getElementById('r-password').value,
            confirmPassword: document.getElementById('r-confirmPassword').value
        }),
        credentials: 'same-origin'
    })
        .then(res => {
            if (!res.ok) {
                throw res.json();
            } else {
                return res.json()
            }
        })
        .then(
            (result) => {
                if (result.token) {
                    localStorage.setItem('access_token', result.token);
                    location.href = '/';
                }
            }).catch((err) => {
        err.then(data => {
            errContainer.innerHTML = data.message;
        })
    })
};

// logout sends a request to logout the user
const logout = () => {
    fetch('/logout', {
        method: 'GET',
        credentials: 'same-origin'
    })
        .then(
            (result) => {
                location.reload();
            },
            (error) => {
                // TODO handle
            }
        )
};