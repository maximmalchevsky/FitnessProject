const personal6Button = document.getElementById('personal6');
if (personal6Button) {
    personal6Button.addEventListener('click', () => {
        sendSubscriptionRequest('Персональный', 6);
    });
}

const personal12Button = document.getElementById('personal12');
if (personal12Button) {
    personal12Button.addEventListener('click', () => {
        sendSubscriptionRequest('Персональный', 12);
    });
}

const corporate12Button = document.getElementById('corporate12');
if (corporate12Button) {
    corporate12Button.addEventListener('click', () => {
        sendSubscriptionRequest('Корпоративный', 12);
    });
}

function sendSubscriptionRequest(type, duration) {
    const name = document.getElementById('name').value;
    const phone = document.getElementById('phone').value;

    if (!name || !phone) {
        alert('Пожалуйста, заполните имя и телефон.');
        return;
    }

    const data = {
        duration,
        name,
        phone,
        type
    };

    fetch('http://localhost:8080/pass', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json'
        },
        body: JSON.stringify(data)
    })
        .then(result => {
            alert("Форма успешно отправлена!")
            console.log('Успешный ответ:', result);
            document.getElementById('response').textContent = JSON.stringify(result, null, 2);
        })
        .catch(error => {
            console.error('Произошла ошибка:', error);
            document.getElementById('response').textContent = `Ошибка: ${error.message}`;
        });
}