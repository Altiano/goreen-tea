# Quick Start
```
git clone git@gitlab.com:altiano/goreen-tea.git --depth 1
cd goreen-tea
bash ./init.sh <project-base-url> <porject-name>
```

example:
```
bash ./init.sh gitlab.com/altiano awesomeness-api      
```


# Features
- Rest API with [Iris](https://github.com/kataras/iris)
- MongoDB as the database
- Unit test with [go-mock](https://github.com/golang/mock)
- [OpenTelemetry](https://github.com/open-telemetry/opentelemetry-go) with [Jaeger](https://www.jaegertracing.io) tracing
- DI with google [wire](https://github.com/google/wire) 
- Docker & Kubernetes setup
- Gitlab CI
- Domain-centric architecture
- Golang validation with [Validator v9](https://github.com/go-playground/validator)
- VS Code tasks & snippets


# Intro 

I wrote an introduction article with the restaurant order sample at [Goreen tea intro](https://blog.altiano.dev/goreen-tea-introduction)


The core architecture is inpired by Clean architecture, 
These how I would describe the archicture with my own understanding:
 
 - Frameworks/Drivers : things that do the actual IO (network, disk, RAM)
 - Adapters : the thing that maps the raw input/output to/from your application logic (use cases)
 - Use cases : high-level orchestration of your logic modules with the Frameworks/Drivers
 - Entities : A logic module

For more detail: [The clean architeture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

Disclaimer, I do not say that this is the correct way to implement that architecture but rather,
I take inspiration from that and adjust it in a way that easier to develop for me.

# Core architecture
![goreen tea architecture](https://blog.altiano.dev/goreen-tea-introduction/architecture.svg "goreen tea architecture")

# Structure
These models will serve as starting point, but you should adjust to meet your needs.
For example, `src/shared/utils.go` may be refactored to a different folder inside `src/shared`.

## `/src`

This is the most important because it is contains all the applicaiton code.
If we compare these with clean architecture terminology, then the intention are as follow:

- `/app` : is about orchestrating the higher level view of your application use cases. It mainly offload the actual work to the domain and just have simple conditional controls. Should not do any calculation on its own.

- `domain` : then will carry out a particular / modular goal that should have clearly defined scoped that is only contain business language/logic.
	- `repo` : the adapaters for `database` framework.
	- `models` : shared object models, errors or constants.
	- `mocks` : mocking objects for unit tests.

- `/frameworks` : these are the layer that responsible for the implementation detail of program IO.
	- Input : To parse HTTP request (`rest/iris.go`), to listen for pub/sub events, to receive filesystem events etc.
	- Output : To store data to the memory for temporary storage (e.g. `memcache/memcache.go`) to the disk for persistent (e.g. `database/mongodb.go`) or to HTTP response (`rest/iris.go)` etc

- `/shared` : shared models, utils, constants and config.go
	- `config.go` loads configuraiton from .env file or environment variables
	- `consts.go` application-wide constants
	- `errors.go` application-wide error constants
	- `utils.go` simple appliation-wide utilities

### Framework & Adapters mapping
Below the mapping between Framework and its adapater:

 - *Controller* is the Adapter to *rest* Framework (Iris, Echo, GraphQL etc)
 - *Resolvers* is the Adapter to *graphql* Framework (standard GraphQL, Apollo, etc)
 - *Handlers* is the Adapter to *websocket* Framework (...)
 - *Procedures* is the Adapter to *rpc* Framework (GRPC, 🤷‍♂️ etc)
 - *Listeners* is the Adapter to *messaging* Framework (Kafka, RabitMQ, Nats, etc)
 - *Runner* is the Adapter to *jobs* Framework (custom one etc)
 - *Repo* is the Adapter to *database* Framework (MongoDB, MySQL, etc)

These are also adapters but so simple that each only implemented inside the framework package.

 - *memcache.go* is the Adapter to *memcache* Framework (Redis, BuntDB, etc)
 - *opentelemetry.go* is the Adapter to *tracing* Framework (OpenTelemetry, Jaeger, Newrelic, etc)
 - *-* is the Adapter to *external http calls* Framework [TODO]
 - *dummyEmail.go* is the Adapter to *email* Framework [TODO]

## `/main.go` and `/di` folder
Entry point to bootstrap the application.


## `/docker`
 - prod.dockerfile : using the multi-stage compilation to produce one executable including all golang runtime deps 
 - debug.dockerfile : useful for debugging the container where all your source is there. 

## `/k8s`
Here you can define how your application final k8s manifest should look like.
I use kustomize to help building the manifest.
```
cd k8s/base
kustomize build . > myapp.yaml
```
If you want to override either for development purpose, you can use the dev folder, where it derived from base.
In production you could do s similar thing after the base manifest is compiled.
Learn more kustomize at : https://kustomize.io

PS: I use k3s for my kubernetes cluster https://k3s.io


## etc
 - .gitlab-ci.yml : Gitlab CI
 - .misc/http/*.http : VS Code RestClient directories
 - .vscode : VS Code Golang debugger
 - .env : used by config.go
 - Init.sh : script to initialize your project for the first time.


## TODO
- [ ] Data store
    - [ ] MongoDB seeding
    - [ ] MySQL support
    - [ ] Redis
- [ ] Servers
    - [ ] GRPC
    - [ ] GraphQL
- [ ] Messaging
    - [ ] Kafka
    - [ ] NATS

# Project Icon
<a href="https://iconscout.com/icons/green-tea" target="_blank">Green Tea Icon</a> by <a href="https://iconscout.com/contributors/pro-symbols">Pro Symbols</a> on <a href="https://iconscout.com">Iconscout</a>