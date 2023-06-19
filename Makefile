# Nombre del Servicio
SERVICE_NAME := cleara
# Llave para JWT
APP_KEY=this_is_my_app_key
# Modo: development | production
MODE=development

# Database Connection
DB_URI=mysql://appdb:123456@localhost:2727
DB_USER=root
DB_PASSWORD=
DB=localdb
DB_SERVER=localhost

# Mailtrap
MAIL_MAILER=smtp
MAIL_HOST=smtp.mailtrap.io
MAIL_PORT=2525
MAIL_USERNAME=
MAIL_PASSWORD=
MAIL_ENCRYPTION=tls

build:
	env CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -ldflags '-w -s' -a -installsuffix cgo -o $(SERVICE_NAME)

#build_docker: # todo debo actualizarlo
#	docker run --rm -e CGO_ENABLED=0 -e GOOS=linux -e GO111MODULE=on -v $(shell pwd):/go/src/$(SERVICE_NAME) -w "/go/src/$(SERVICE_NAME)" golang:1.15.2-alpine3.12 sh -c "apk update && apk upgrade && apk add --no-cache bash git openssh && go get && go build  -ldflags '-w -s' -a -installsuffix cgo -o $(SERVICE_NAME)"

build_doc:
	env GO111MODULE=on swag init

run_local:
	env MODE=$(MODE) DB_USER=$(DB_USER) DB_SERVER=$(DB_SERVER) DB_PASSWORD=$(DB_PASSWORD) DB=$(DB) SecretKey=$(APP_KEY) MAIL_HOST=$(MAIL_HOST) MAIL_PORT=$(MAIL_PORT) MAIL_USERNAME=$(MAIL_USERNAME) MAIL_PASSWORD=$(MAIL_PASSWORD) ./$(SERVICE_NAME)

test:
	env MODE=$(MODE) DB_USER=$(DB_USER) DB_SERVER=$(DB_SERVER) DB_PASSWORD=$(DB_PASSWORD) DB=$(DB) SecretKey=$(APP_KEY) MAIL_HOST=$(MAIL_HOST) MAIL_PORT=$(MAIL_PORT) MAIL_USERNAME=$(MAIL_USERNAME) MAIL_PASSWORD=$(MAIL_PASSWORD) go test -v