PROJECT_NAME=robusta
BUILD_VERSION=$(shell cat VERSION)
DOCKER_IMAGE=$(PROJECT_NAME):$(BUILD_VERSION)
GO_BUILD_ENV=CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on
GO_FILES=$(shell go list ./... | grep -v /vendor/)
REALIZE_VERSION=2.0.2
REALIZE_IMAGE=realize:$(REALIZE_VERSION)

.SILENT:

all: mod_tidy fmt vet test install 

build:
	$(GO_BUILD_ENV) go build -v -o $(PROJECT_NAME)-$(BUILD_VERSION).bin .

install:
	$(GO_BUILD_ENV) go install

vet:
	$(GO_BUILD_ENV) go vet $(GO_FILES)

fmt:
	$(GO_BUILD_ENV) go fmt $(GO_FILES)

test:
	$(GO_BUILD_ENV) go test $(GO_FILES) -cover -v

integration_test:
	$(GO_BUILD_ENV) go test -tags=integration $(GO_FILES) -cover -v

mod_tidy:
	$(GO_BUILD_ENV) go mod tidy

compose_dev: realize
	cd deployment/dev && REALIZE_VERSION=$(REALIZE_VERSION) docker-compose up

compose_prod: docker
	cd deployment/docker && BUILD_VERSION=$(BUILD_VERSION) docker-compose up

docker_prebuild: vet test web_build build
	mkdir -p deployment/docker/web/dist
	mkdir -p deployment/docker/configs
	mkdir -p deployment/docker/templates
	mv $(PROJECT_NAME)-$(BUILD_VERSION).bin deployment/docker/$(PROJECT_NAME).bin; \
	cp -R web/dist deployment/docker/web/; \
	cp -R templates deployment/docker/; \
	cp -R configs deployment/docker/;

docker_build:
	cd deployment/docker; \
	docker build -t $(DOCKER_IMAGE) .;

docker_postbuild:
	cd deployment/docker; \
	rm -rf $(PROJECT_NAME).bin 2> /dev/null;\
	rm -rf web 2> /dev/null; \
	rm -rf templates 2> /dev/null; \
	rm -rf configs 2> /dev/null;

docker: docker_prebuild docker_build docker_postbuild

docker_run:
	docker run -p 8080:8080 $(DOCKER_IMAGE)

web_build:
	cd web; \
	yarn install; \
	yarn build

heroku_config:
	heroku login
	cd scripts && ./env-heroku.sh

heroku_predeploy:
	cd deployment/docker; \
	heroku container:login; \
	heroku container:push web --app goway; \
	heroku container:release web --app goway; \
	heroku open --app goway

deploy_heroku: docker_prebuild heroku_predeploy docker_postbuild

web_lint:
	cd web; \
	yarn lint
size:
	cd web; \
	yarn size
size_why:
	cd web; \
	yarn size-why

web_plugins:
	cd web; \
	rm -rf ./node_modules/code-example
	cd web/src/components/Editor/plugins/code-example; \
	npm run build
	cd web; \
	mkdir ./node_modules/code-example; \
	cp -R ./src/components/Editor/plugins/code-example/ ./node_modules/

oauth_local:
	sudo iptables -t nat -I OUTPUT -p tcp -d 127.0.0.1 --dport 80 -j REDIRECT --to-ports 8080
	sudo iptables -t nat -I OUTPUT -p tcp -d 127.0.0.1 --dport 443 -j REDIRECT --to-ports 443
	# sudo echo "127.0.0.1 mylocalhost.com" >> /etc/hosts

# https://github.com/wagoodman/dive for more detail
dive:
	dive $(DOCKER_IMAGE)

# https://github.com/oxequa/realize  
realize:
	cd deployment/dev; \
	docker build -t $(REALIZE_IMAGE) .;

# https://github.com/FiloSottile/mkcert
gen_certs:
	mkcert -install -cert-file deployment/certs/server.pem -key-file deployment/certs/server-key.pem mylocalhost.com localhost 127.0.0.1 ::1