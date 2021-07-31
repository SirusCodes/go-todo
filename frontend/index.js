const username = localStorage.getItem('username');
const token = localStorage.getItem('token');

window.addEventListener('load', () => {
    if (!(username && token)) {
        window.location = "http://127.0.0.1:5500/frontend/authenticate.html"
    }
    document.getElementById("hello-user").innerHTML = `Hello, ${username}`;
});

const logoutBtn = document.getElementById("logout-btn");

logoutBtn.addEventListener('click', () => {
    localStorage.clear();
    window.location = "http://127.0.0.1:5500/frontend/authenticate.html"
});