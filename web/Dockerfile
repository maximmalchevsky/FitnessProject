FROM nginx:alpine

# Копируем файлы в директорию, откуда Nginx будет их обслуживать
COPY . /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf
# Открываем порт 3000
EXPOSE 3000

# Запускаем Nginx
CMD ["nginx", "-g", "daemon off;"]