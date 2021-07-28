const ENDPOINT = "http://127.0.0.1:8000";

// forms
const loginForm = document.getElementById('login');
const registerForm = document.getElementById('register');

// switcher
const loginSwitcher = document.getElementById('login-switcher');
const registerSwitcher = document.getElementById('register-switcher');

// action button
const loginButton = document.getElementById('login-btn');
const registerButton = document.getElementById('register-btn');

loginSwitcher.addEventListener('click', () => {
    registerForm.style.display = 'none';
    loginForm.style.display = 'flex';
});

registerSwitcher.addEventListener('click', () => {
    registerForm.style.display = 'flex';
    loginForm.style.display = 'none';
});

loginButton.addEventListener('click', async (e) => {
    e.preventDefault();
    const username = loginForm.username.value;
    const password = loginForm.password.value;
    if (username.length < 1 || password.length < 1) {
        alert('Please fill in all fields!');
        return;
    }
    const response = await fetch(`${ENDPOINT}/login`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json",
        },
        body: JSON.stringify({
            username: username,
            password: username,
        }),
        mode: "cors",
    });

    if (response.status === 200) {
        const result = await response.json();
        console.log(result);
        localStorage.setItem('token', result.data.token);
        localStorage.setItem('refresh_token', result.data.refresh_token);
        localStorage.setItem('username', result.data.username);
    } else if (response.status === 403) {
        const error = await response.json();
        alert(error.message);
    }
    else {
        const error = await response.json();
        console.log(error);
        alert('Login failed');
    }
});

registerButton.addEventListener('click', async (e) => {
    e.preventDefault();
    const username = registerForm.username.value;
    const password = registerForm.password.value;
    const confirmPassword = registerForm.confirmPassword.value;
    if (username.length < 1 || password.length < 1) {
        alert('Please fill in all fields!');
        return;
    } else if (password !== confirmPassword) {
        alert('Password does not match!');
        return;
    }

    const response = await fetch(`${ENDPOINT}/register`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json",
        },
        body: JSON.stringify({
            username: username,
            password: password,
        }),
        mode: "cors",
    });

    if (response.status === 201) {
        const result = await response.json();
        console.log(result);
        localStorage.setItem('token', result.data.token);
        localStorage.setItem('refresh_token', result.data.refresh_token);
        localStorage.setItem('username', result.data.username);
    } else if (response.status === 403) {
        const error = await response.json();
        alert(error.message);
    } else {
        const error = await response.json();
        console.log(error);
        alert('Register failed');
    }
});