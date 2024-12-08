document.addEventListener("DOMContentLoaded", function () {
    const loginForm = document.querySelector('form[action="/login"]');
    if (loginForm) {
        loginForm.addEventListener("submit", function (e) {
            e.preventDefault();

            const username = document.getElementById("username").value;
            const password = document.getElementById("password").value;

            if (!username || !password) {
                alert("Пожалуйста, заполните все поля!");
                return;
            }

            fetch('http://localhost:8082/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ username, password }),
            })
                .then(response => response.json())
                .then(data => {
                    if (data.success) {
                        window.location.href = '/notes';
                    } else {
                        alert(data.message || 'Ошибка входа');
                    }
                })
                .catch(error => {
                    console.error('Ошибка:', error);
                    alert('Произошла ошибка. Попробуйте снова.');
                });
        });
    }

    const registerForm = document.querySelector('form[action="/register"]');
    if (registerForm) {
        registerForm.addEventListener("submit", function (e) {
            e.preventDefault();

            const username = document.getElementById("username").value;
            const password = document.getElementById("password").value;

            if (!username || !password) {
                alert("Пожалуйста, заполните все поля!");
                return;
            }

            fetch('http://localhost:8082/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ username, password }),
            })
                .then(response => response.json())
                .then(data => {
                    if (data.success) {
                        window.location.href = '/login';
                    } else {
                        alert(data.message || 'Ошибка регистрации');
                    }
                })
                .catch(error => {
                    console.error('Ошибка:', error);
                    alert('Произошла ошибка. Попробуйте снова.');
                });
        });
    }

    const deleteButtons = document.querySelectorAll('.btn-danger');
    deleteButtons.forEach(button => {
        button.addEventListener("click", function (e) {
            const noteId = this.getAttribute("data-id");
            const confirmDelete = confirm("Вы уверены, что хотите удалить эту заметку?");
            if (confirmDelete) {
                fetch(`http://localhost:8082/delete-note/${noteId}`, {
                    method: 'POST',
                })
                    .then(response => response.json())
                    .then(data => {
                        if (data.success) {
                            this.closest('.card').remove();
                        } else {
                            alert('Ошибка при удалении заметки');
                        }
                    })
                    .catch(error => {
                        console.error('Ошибка:', error);
                        alert('Произошла ошибка. Попробуйте снова.');
                    });
            }
        });
    });
});
