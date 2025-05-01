## JWT Auth service
### Endpoints

#### /access/{guid}
Данный эндпоинт выдает пару Access, Refresh токенов.
Access токен - полноценный JWT токен. Refresh токен представлен в произвольной форме в формате base64.
Пара будет выдана если юзер с id=guid по текущему IP не заходил или заходил, но refresh токен просрочен
или актуальная версия токена изменилась. 

#### /refresh
Данный эндпоинт принимает "refresh" в теле и выдает пару Access, Refresh токенов с измененным Access токеном.
Пара будет выдана если токены-пара, если версия токена совпадает с актуальной и оба токена валидны, не просрочены.
Если юзер обновляет токены с другого IP, произойдет отправка уведомления на почту (в консоль) об этом.  

### Как запустить
в следующих файлах вписать необходимые параметры
```bash
    ./configs/.env
    ./configs/.docker_env
```
*при запуске создаются таблицы и первые три пользователя с id=1, id=2, id=3*

#### Через docker compose
Далее в корне проекта запустить
```bash
    docker compose --env-file=./configs/.env --env-file=./configs/.docker_env up -d
```
или
```bash
    make dockerup
```
сервис будет ожидать на localhost:${DOCKER_HOST_PORT}    
swagger - localhost:${DOCKER_HOST_PORT}/swagger
#### В терминале
Далее в корне проекта запустить
```bash
    go run ./cmd/gojwt/main.go
```
или
```bash
    make terminalup
```
сервис будет ожидать на localhost:${PORT}    
swagger - localhost:${PORT}/swagger

### Что нужно сделать
- добавить Redis, чтобы забрасывать туда access токены после операции /refresh.
