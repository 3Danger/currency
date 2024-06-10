<span style="font-size:2em;">README для проекта [currency](https://github.com/3Danger/currency.git)</span>

Склонируйте репозиторий:
``` shell
git clone https://github.com/3Danger/currency.git
```

Перейдите в директорию проекта:
``` shell
cd currency
```

Дополните файл .env вашим токеном:
```dotenv
HTTP_TOKEN={{YOUR_TOKEN}}
```

Установите зависимости:
``` shell
go mod tidy
```

Для запуска Postgres в контейнере используйте команду:
``` makefile
make postgres-up
```

Примените миграции:
``` makefile
make migrate-postgres-up
```

Для запуска воркеров используйте команду:
``` makefile
make run-workers
```

Для запуска REST API используйте команду:
``` makefile
make run-rest
```

<span style="font-size:1.5em;">Помощь</span>

Если у вас возникли вопросы или проблемы, пожалуйста, создайте issue в репозитории.
