// forms
const loginForm = document.getElementById('login');
const registerForm = document.getElementById('register');

// switcher
const loginSwitcher = document.getElementById('login-switcher');
const registerSwitcher = document.getElementById('register-switcher');

loginSwitcher.addEventListener('click', () => {
    registerForm.style.display = 'none';
    loginForm.style.display = 'flex';
});

registerSwitcher.addEventListener('click', () => {
    registerForm.style.display = 'flex';
    loginForm.style.display = 'none';
});

