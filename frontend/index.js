const ENDPOINT = "http://127.0.0.1:8000";

const username = localStorage.getItem('username');
let token = localStorage.getItem('token');
const logoutBtn = document.getElementById("logout-btn");
const todoText = document.getElementById("todo-text");

window.addEventListener('load', async () => {
    if (!(username && token)) {
        window.location = "http://127.0.0.1:5500/frontend/authenticate.html"
    }
    document.getElementById("hello-user").innerHTML = `Hello, ${username}`;
    const data = await fetchCurrentTodos();
    for (let i = 0; i < data.length; i++) {
        const todo = data[i];
        addTodoToDOM(todo.id, todo.task);
    }
});

logoutBtn.addEventListener('click', () => {
    localStorage.clear();
    window.location = "http://127.0.0.1:5500/frontend/authenticate.html"
});

todoText.addEventListener('keyup', async (e) => {
    console.log(e.key);
    if (e.key === "Enter") {
        await sendTodo();
    }
});

async function refreshTokens(callback) {
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
        callback();
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

    if (response.status === 201) {
        todoText.value = "";
        const result = await response.json();
        addTodoToDOM(result.data.id, result.data.task);
    } else if (response.status === 401) {
        await refreshTokens(sendTodo);
    }
}

async function fetchCurrentTodos() {
    const response = await fetch(`${ENDPOINT}/todos`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        }
    });
    if (response.status === 200) {
        const result = await response.json();
        return result.data;
    } else if (response.status === 401) {
        await refreshTokens(() => {
            window.location.href = "http://127.0.0.1:5500/frontend/index.html";
        });
    }
}

function addTodoToDOM(id, todo) {
    const div = document.createElement("div");
    div.setAttribute("class", "todo");

    const h5 = document.createElement("h5");
    h5.innerHTML = todo;

    const i = document.createElement("i");
    i.setAttribute("id", id);
    i.setAttribute("class", "material-icons");
    i.innerHTML = "delete";

    div.appendChild(h5);
    div.appendChild(i);

    document.getElementById("main-body").appendChild(div);

    setDeleteListener(id);
}

function setDeleteListener(id) {
    document.getElementById(id).addEventListener('click', async () => {
        const response = await fetch(`${ENDPOINT}/todo/${id}`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            }
        });
        if (response.status === 204) {
            removeTodoFromDOM(id);
        } else if (response.status === 401) {
            await refreshTokens(() => {
                window.location.href = "http://127.0.0.1:5500/frontend/index.html";
            });
        }
    });
}

function removeTodoFromDOM(id) {
    const parent = document.getElementById(id).parentNode;
    document.getElementById("main-body").removeChild(parent);
}