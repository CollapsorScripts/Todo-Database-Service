FROM --platform=linux/arm64 arm64/golang:latest as builder

SHELL ["/bin/bash", "-c"]

# Устанавливаем переменные
ENV GOARCH=arm64
ENV TARGETOS=linux
ENV ENTRY_POINT=./cmd/entrypoint
ENV PROGRAM=./service
ENV WKDIR=/build

# Рабочая директория
WORKDIR ${WKDIR}
COPY . ${WKDIR}

# Компилируем
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${GOARCH} go build -o ${PROGRAM} ${ENTRY_POINT}

# Создаем финальный образ
FROM arm64v8/alpine:latest

# Устанавливаем переменные
ENV PROGRAM=service
ENV WKDIR=/app
ENV BUILDIR=/build
#Порт для прослушки
ENV PORT=44040

# Рабочая директория
WORKDIR ${WKDIR}

# Копируем исполняемый файл из предыдущего образа
COPY --from=builder ${BUILDIR}/${PROGRAM} ./${PROGRAM}

# Добавляем сертификаты
RUN apk add --upgrade --no-cache ca-certificates && update-ca-certificates

# Устанавливаем время
RUN apk add tzdata && echo "Europe/Moscow" > /etc/timezone && ln -s /usr/share/zoneinfo/Europe/Moscow /etc/localtime

# Копируем файл конфигурации и сертификаты/ключи в контейнер
COPY config/prod.yaml local.yaml

# Открываем порты
EXPOSE ${PORT}