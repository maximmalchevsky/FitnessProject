// Добавляем обработчик события на кнопку
document.getElementById('createTraining').addEventListener('click', sendTrainingData);

// Функция для отправки данных на сервер
function sendTrainingData() {
    // Получаем значения из полей ввода
    const phone = document.getElementById('phoneInput').value.trim();
    const name = document.getElementById('nameInput').value.trim();

    // Проверяем, что оба поля заполнены
    if (!phone || !name) {
        alert("Пожалуйста, заполните оба поля!");
        return;
    }

    const url = 'http://localhost:8080/training';

    // Формируем данные для отправки
    const data = {
        name: name,
        phone: phone
    };

    // Выполняем POST-запрос
    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json'
        },
        body: JSON.stringify(data)
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`Ошибка: ${response.status}`);
            }
            return response.json(); // Обрабатываем ответ как JSON
        })
        .then(result => {
            console.log('Успешный ответ:', result);
            document.getElementById('response').textContent = JSON.stringify(result, null, 2);
        })
        .catch(error => {
            console.error('Произошла ошибка:', error);
            document.getElementById('response').textContent = `Ошибка: ${error.message}`;
        });
}

