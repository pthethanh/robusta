# robusta

[![Join the chat at https://gitter.im/pthethanh/robusta](https://badges.gitter.im/pthethanh/robusta.svg)](https://gitter.im/pthethanh/robusta?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Tool to track guests visit to my hotels over the time

## Prerequisites

Make sure you have the development environment matches with these notes below so we can mitigate any problems of version mismatch.

- Backend:
  - Go SDK: 1.10 (latest 1.10.8).
    Make sure to set `$GOROOT` and `$GOPATH` correctly.
    You can check those environment variable by typing: `go env`.
  - Install [govendor](https://github.com/kardianos/govendor) for dependency management.
  - MongoDB: 4.1 (latest 4.1.8).

- Frontend:
  - NodeJS: 11.10 (latest 11.10.0).
  - Use `yarn` instead of `npm` is possible.
  - Framework: [VueJS 2.x](https://vuejs.org) (latest 2.6.6).
  - UI components framework: [Element-UI](https://element.eleme.io) (latest 2.5.4).

- Commons:
  - OS: Should use Linux (latest Ubuntu or your choice of distro) if possible.
    Windows does not play well with Docker and some other techs we may use.
    If you still prefer to use Windows, so you may have to cope with problems by yourself later
    since we're assuming everything will be developed and run on Linux.
  - Install [Docker CE](https://docs.docker.com/install/) (latest 18.x) and [docker-compose](https://docs.docker.com/compose/install/).
  - Install [git](https://git-scm.com/) for manage source code.
  - IDE of your choice, recommended `Goland` or `VS Code`.

## Development

#### 1. Clone code to local

```shell
$ go get -u -v github.com/pthethanh/robusta
or
$ cd $GOPATH/src
$ git clone https://github.com/pthethanh/robusta.git
```
After this step, source code must be available at `$GOPATH/src/github.com/pthethanh/robusta`.

#### 2. Start development environment manually

- Start MongoDB service at localhost:27017. The easiest way is to run the Docker as below:

  ```shell
  $ docker run -p 27017:27017 -v /opt/data/mongo_home:/data/db --name mongo -d mongo:4.1.8
  ```

- Start backend API service (Go):

  ```shell
  $ go run main.go
  # Backend service will start on port :8080.
  ```

- Open another terminal and start frontend service (NodeJS):

  ```shell
  $ cd web
  $ yarn install   // Do this only once
  $ yarn serve
  # Frontend service will start at port :8081 and connect to backend service at :8080.
  # Open your browser then head to http://localhost:8081.
  ```

#### 3. Start development environment with Docker

Instead of manually start services like step 2. You can use Docker to start all services at once.

```shell
$ cd /web && yarn install && cd ../   # Do this only once to install frontend modules first
$ make compose
```

After started, services will be available at `localhost` with ports as below:
```
MongoDB: 27017
Backend: 8080 (hot-reload dashboard: 8082)
Frontend: 8081
```

## Notes

- Make sure to run `go fmt`, `go vet`, `go test`, and `go build / go install` before pushing your code to Github.
  Or you can just run `make` before pushing.
- Never commit directly to `master` or `develop` branches (you don't have permission to do so, anyway). Instead, checkout from `develop` branch to a separated branch then work on that.  
  Whenever you finish your work, you can create a Pull Request (PR) / Merge Request (MR) to ask for code review and merging your branch back to `develop`.   
  `master` branch will be reserved when administrator decide to release a stable version of application.
- All documentations must be written in [Markdown](https://guides.github.com/features/mastering-markdown/) format, recommended to use [Typora](https://typora.io/) as the Markdown editor.

## Tech requirements:

- Authentication:
  - LDAP (AD) (backend authentication)
  - JWT (frontend authentication)
- Database
  - MongoDB
- Frontend
  - Full frontend framework [Vue](https://vuejs.org)
  - Component toolkit: [Element-UI](https://element.eleme.io)

## TODO
Need update
