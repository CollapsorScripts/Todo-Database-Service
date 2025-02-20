.PHONY: all

ENTRY_POINT := "./cmd/entrypoint"
EXE_NAME := "./databaseService"
LOCAL := "--config=./config/local.yaml"
PROD := "--config=./config/prod.yaml"
FILES_DIR := "./files"
PROTOS_DIR := "protos"

# ======================================================
# build - сборка
# run_local - запуск сервиса с локальной конфигурацией
# run_prod - запуск сервиса с продакшен конфигурацией
# gen_proto - генерация файлов пакета из proto файлов
# clear - отчистка локальных папок
# ======================================================

all: build run_local

build:
	@echo Compile and build...
	@go build -o $(EXE_NAME) $(ENTRY_POINT)

run_local:
	@echo Run local app: $(EXE_NAME) $(LOCAL)
	@$(EXE_NAME) $(LOCAL)

run_prod:
	@echo Run prod app: $(EXE_NAME) $(PROD)
	@$(EXE_NAME) $(PROD)

gen_proto:
	@make -C $(PROTOS_DIR)

gen_doc:
	@protoc --doc_out=./docs --doc_opt=markdown,docs.md protos/proto/Service/*.proto

clear:
	@echo Cleaning files...
	@rm -rf $(FILES_DIR) $(EXE_NAME)