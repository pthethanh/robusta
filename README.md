# robusta

Just a simple website to share technology knowledge. Live demo is available at: https://goway.herokuapp.com

This project follows [golang-standards/project-layout](https://github.com/golang-standards/project-layout) and hence can be used as a reference for new Go developers.

**Note** that this project now is no longer in active development, but feel free to make pull request.

## Prerequisites

Make sure you have the development environment matches with these notes below so we can mitigate any problems of version mismatch.

- Backend:
  - Go SDK: 1.13.
    Make sure to set `$GOROOT` and `$GOPATH` correctly.
    You can check those environment variable by typing: `go env`.
  - MongoDB: 4.1 (latest 4.1.8).
  - Go module

- Frontend:
  - NodeJS: 11.10 (latest 11.10.0).
  - Use `yarn` instead of `npm`.
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

#### 2. Start development environment with Docker

You can use Docker to start all services at once. This will support auto reload for both frontend and backend

```shell
$ cd /web && yarn install && cd ../
$ make compose_dev
```

After started, services will be available at `localhost` with ports as below:
```
MongoDB: 27017
Backend: 8080
Frontend: 8081
```
#### Testing Google Authentication locally

Google doesn't support OAuth with localhost or any domain with port, but our server is running on 8080, hence we need to create a local domain named `mylocalhost.com` and make sure all requests to that domain is forwarded to the server that is running at port 8080

**Following steps need to be done**:

**1. Update iptables**

```shell
sudo iptables -t nat -I OUTPUT -p tcp -d 127.0.0.1 --dport 80 -j REDIRECT --to-ports 8080
```
**2. Update /etc/hosts**

```shell
127.0.0.1 mylocalhost.com
```
**3. Update Google OAuth2 configurations**

Make sure `mylocalhost.com` is updated in Google Console
Add `http://mylocalhost.com/auth/google/callback` into the `Authorized redirect URIs` of [Google Console API](https://console.developers.google.com/apis/credentials/oauthclient/511091284450-9oes9tddrskgtlfcafblrke7pk28lthp.apps.googleusercontent.com?project=goway-1546095349277)

#### Checking for web performance using size-limit

```
make size
```
For more detail, see [size-limit on github](https://github.com/ai/size-limit)

#### Useful VS Code plugins
- Vetur (octref.vetur)
- Go (ms-vscode.go)
- vue-format (febean.vue-format)
- Code Spell Checker (streetsidesoftware.code-spell-checker)
- GitLens — Git supercharged (eamodio.gitlens)
- Docker (ms-azuretools.vscode-docker)

## Notes

- Make sure to run `go fmt`, `go vet`, `go test`, and `go build / go install` before pushing your code to Github.
  Or you can just run `make` before pushing.
- Never commit directly to `master` or `develop` branches (you don't have permission to do so, anyway). Instead, checkout from `develop` branch to a separated branch then work on that.  
  Whenever you finish your work, you can create a Pull Request (PR) / Merge Request (MR) to ask for code review and merging your branch back to `develop`.   
  `master` branch will be reserved when administrator decide to release a stable version of application.
- All documentations must be written in [Markdown](https://guides.github.com/features/mastering-markdown/) format, recommended to use [Typora](https://typora.io/) as the Markdown editor.

## Tech requirements:

Color: https://www.colourlovers.com/palette/173625/Browser_wars

## TODO
Need update
