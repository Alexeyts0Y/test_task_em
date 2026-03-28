# Тестовое задание Effective mobile
## Описание проекта
### Использованные технологии:
<img width="110" height="110" title="Golang" alt="golang" src="https://github.com/user-attachments/assets/9218acc9-6b90-4dbd-9ef4-4403f9664312"/>
<img width="100" height="100" title="Gin" alt="gin" src="https://github.com/user-attachments/assets/907271ec-aa81-4202-b9c7-5c945e1c4abb" />
<img width="110" height="110" title="PostgreSQL" alt="postgresql" src="https://github.com/user-attachments/assets/82b28ccc-f900-4877-a3d4-c19b439e180d"/>
<img width="110" height="110" title="Docker" alt="docker" src="https://github.com/user-attachments/assets/3b3e98c7-2d7f-4013-b603-f5694bd174e9" />

### Реализованный функционал:
#### Управление подписками (CRUDL)
- Создание подписки с привязкой к user_id, указанием цены и даты начала.
- Получение информации о конкретной подписке по её уникальному ID.
- Обновление данных существующей подписки (изменение цены, сервиса или дат).
- Удаление записи о подписке из базы данных.
- Получение списка подписок с фильтрацией по user_id и названию сервиса (service_name).

#### Бизнес-логика и расчеты
- Расчет суммарной стоимости подписок для пользователя за произвольный период.
- Определение пересечения периодов (алгоритм Overlap) между сроком действия подписки и запрашиваемым интервалом.
- Поддержка бессрочных подписок (автоматический расчет стоимости для активных услуг без даты окончания).
- Конвертация строковых дат формата MM-YYYY в числовые эквиваленты для точных математических вычислений.
#### Инфраструктура и надежность
- Автоматическое развертывание БД (миграции) при старте приложения через docker-compose.
- Валидация входящих запросов (проверка UUID, обязательных полей и форматов дат).
- Структурированное логирование всех ключевых операций и ошибок в формате JSON.
- Контроль состояния базы данных (Healthcheck) для предотвращения ошибок запуска приложения.

### Установка и запуск (Без использования Docker)
1. Склонируйте репозиторий
```
git clone https://github.com/Alexeyts0Y/test_task_em
cd test_task_em
```

2. Создайте файл .env по шаблону .env.template
```
# .env

# Имя вашей базы данных
DB_NAME=your_db_name

# Хост вашей базы данных
DB_HOST=your_db_host

# Пароль от вашей базы данных
DB_PASSWORD=your_db_password

# Порт вашей базы данных
DB_PORT=your_db_port

# Имя пользователя базы данных
DB_USER=your_db_user
```

3. Находясь в корне проекта, запустите приложение, введя:
```
go run cmd/app/main.go
```

### Установка и запуск (С использованием Docker)
1. Склонируйте репозиторий
```
git clone https://github.com/Alexeyts0Y/test_task_em
cd test_task_em
```

2. Создайте файл .env по шаблону .env.template
```
# .env

# Имя вашей базы данных
DB_NAME=your_db_name

# Хост вашей базы данных
DB_HOST=your_db_host

# Пароль от вашей базы данных
DB_PASSWORD=your_db_password

# Порт вашей базы данных
DB_PORT=your_db_port

# Имя пользователя базы данных
DB_USER=your_db_user
```

3. Запустите docker-compose
```
docker-compose up --build
```

### Структура проекта
```
test_task_em/
├── cmd/
│   └── app/
│       └── main.go              # Точка входа в приложение
├── internal/
│   ├── config/                  # Конфиги
│   ├── handlers/                # Хендлеры
│   ├── models/                  # Модели данных
│   ├── repository/              # Репозитории
│   └── service/                 # Бизнес логика
├── migrations/                  # SQL миграции (embed)
├── .env.template                # Шаблон для переменных окружения
├── docker-compose.yml           # Оркестрация контейнеров
├── Dockerfile                   # Инструкция сборки образа
├── go.mod                       # Зависимости проекта
└── go.sum                       # Хеши зависимостей
```
