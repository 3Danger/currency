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

Обновить зависимости, запустить контейнер postgres, применить миграции
``` makefile
make setup
```

Для запуска синхронизации валют:
``` makefile
make run-workers
```

Для запуска REST API используйте команду:
``` makefile
make run-rest
```

Открыть swagger документацию можно по пути:
- http://localhost:8080/api/swagger/index.html


<span style="font-size:1.5em;">Помощь</span>

Если у вас возникли вопросы или проблемы, пожалуйста, создайте issue в репозитории.
