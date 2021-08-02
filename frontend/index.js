const ENDPOINT = "http://127.0.0.1:8000";

const username = localStorage.getItem('username');
const token = localStorage.getItem('token');
const logoutBtn = document.getElementById("logout-btn");
const todoText = document.getElementById("todo-text");

window.addEventListener('load', () => {
    if (!(username && token)) {
        window.location = "http://127.0.0.1:5500/frontend/authenticate.html"
    }
    document.getElementById("hello-user").innerHTML = `Hello, ${username}`;
});

logoutBtn.addEventListener('click', () => {
    localStorage.clear();
    window.location = "http://127.0.0.1:5500/frontend/authenticate.html"
});

todoText.addEventListener('keyup', async (e) => {
    if (e.key === "Enter") {
        await sendTodo();
    }
});

async function refreshTokens() {
    const response = await fetch(`${ENDPOINT}/refresh`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify({
            username: username,
            refresh_token: localStorage.getItem('refresh_token'),
        })
    });

    if (response.status === 201) {
        const result = await response.json();
        localStorage.setItem('token', result.data.token);
        token = result.data.token;
        sendTodo();
    }
}

async function sendTodo() {
    const response = await fetch(`${ENDPOINT}/todo`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify({
            task: todoText.value
        })
    });

    if (response.status === 200) {
        todoText.value = "";
    } else if (response.status === 401) {
        refreshTokens();
    }

}