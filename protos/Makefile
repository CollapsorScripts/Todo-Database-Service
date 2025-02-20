.PHONY: gen

PKG_FOLDER := "../pkg/service"

gen:
		@echo Step 1: Generate proto files...
		@protoc -I proto proto/Service/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
		@echo Files created...
		@echo Step 2: Copy files to ${PKG_FOLDER}
		@cp ./gen/go/Service/*.go ${PKG_FOLDER}
		@echo Successfully!!!