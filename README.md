# Микросервис (gRPC) для работы с базой данных Postgres
Конфигурация находится в ```./config/prod.yaml```
```yaml
    env: "prod" #Прод конфиг
    grpc: #Конфигурация сервера
      port: 44040 #Порт который будет прослушивать gRPC сервер
      timeout: 5s #Таймаут на запрос
    database: #Конфигурация базы данных
      host: localhost #Хост
      user: postgres #Пользователь БД
      password: postgres #Пароль БД
      name: postgres #Имя базы данных
      port: 5432 #Порт базы данных
      migrations: true #Применять миграции при запуске
    paths: #Пути
      files: "./files" #Директория с файлами
      logDir: "./log" #Директория для логов
      logName: "databaseService.log" #Файл лога
```

## Запуск
Присутствует запуск с аргументами: ```databaseService --config=./config/prod.yaml```
# Make
В корне проекта присувствует **makefile**, содержит команды:
```
make build - сборка проекта, бинарник создается в корне директории
make run_local - запуск сервиса с локальной конфигурацией (--config=config/local.yaml)
make run_prod - запуск сервиса с продакшен конфигурацией (--config=config/prod.yaml)
make gen_proto - генерация файлов пакета из proto файлов (сразу скопирует файлы в pkg/service)
make clear - отчистка локальных папок
```

## Тесты
Тесты находятся в папке: ```tests```

## Docker
В корне присутствует **Dockerfile** и **docker-compose.yml**. По умолчанию будет скопирована конфигурация
из папки ```./config/prod.yaml``` и скопирована в контейнер как ```local.yaml```, сервис запустится с аргументом
```--config=local.yaml```

> [!IMPORTANT]
> Если не указать аргументы запуска, по умолчанию будет загружен конфиг по пути ```./config/local.yaml```, если
> конфигурация по данному пути не будет найдена, то сервис попытается загрузить конфигурацию по пути ```./local.
> yaml```, в случае если конфигурация так и не будет найдена, сервис завершится с ошибкой.

# Protocol Buffers
Proto файл находится в папке ```./protos/proto```
Для кодогенерации есть **Makefile** файл, открываем терминал в корне проекта и пишем ```make
gen_proto``` - для генерации сервера и клиента на языке **Golang**. Для генерации документации используем ```make
gen_doc```

> [!IMPORTANT]
> Для использования моего **makefile** обязательно необходима утилита **[make](https://www.make.com/en)** и установленные версии **protoc**
> под
> необходимый язык (В моем случае это **[grpcio, grpcio-tools](https://grpc.io/docs/languages/python/quickstart/)** для языка **Python** и **[protoc-gen-go,
> protoc-gen-go-grpc](https://grpc.io/docs/languages/go/quickstart/)** для языка
> **Golang**).
> Необходимые инструменты берем с официальной документации **[gRPC](https://grpc.io/)**

# Документация
С документацией по API можно ознакомиться [тут](https://github.com/CollapsorScripts/Todo-Database-Service/blob/main/docs/docs.md).