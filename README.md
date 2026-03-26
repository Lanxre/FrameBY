# POSTGRES (База данных)

- [Скачать PostgreSQL](https://www.postgresql.org/download/)
- *После установки*, запусти сервер PostgreSQL и создай базу данных с название `frameby`
- [Скачать pgAdmin](https://www.pgadmin.org/download/), опцианально, если нужен удобный интерфейс для работы, а не через командную строку. (возможно нужен будет впн)

# Настройка переменных окружения

Если менял пользователя для бд не `postgres` и поставил пароль не `root`, то в `.env` файле нужно указать эти данные.
Если оставил дефолтные значения, то можно пропустить пункт `Как настроить`, конфиг сам создаст дефолтные значения.

## Как настроить

```
$ cd backend
```

Заходишь в бекенд проекта и переименовываешь файл `.env.example` в `.env` и заполняешь его данными

Должно быть так:

```
# SERVER
SERVER_ADDRESS=:8080
SERVER_WRITE_TIMEOUT=15s
SERVER_READ_TIMEOUT=15s

#DB
DB_PASSWORD=YOU_DB_PASSWORD - default 'root'
DB_HOST=localhost
DB_PORT=5432
DB_NAME=frameby
DB_SSLMODE=disable

DATABASE_URL=postgres://postgres:agentdb@localhost:5432/frameby?sslmode=disable

#STATE
STATE=DsGqMftyi34NG7725PN67V

#JWT
JWT_SECRET=4jAJLUqf0i9+hPAhRwhYN4zB58Wdxw1d+aDze2gXf785wrgEr1w5Ewnau/VQ1zNX

#FRONTEND
FRONTEND_URL=http://localhost:3000

#BACKEND
BACKEND_URL=http://localhost:8080
```

# Установка Bun

[Скачать Bun](https://bun.com/)

# Установка зависимостей

```
$ bun run install:all
```

После того как установил все зависимости, запуск проекта.

```
$ bun run dev
```