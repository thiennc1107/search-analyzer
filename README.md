# Search analyzer
An web application to analyze search results of various keywords.
## Core feature
- Dashboard the view analytic results of each keyword search results.
- File upload page to upload the list of keyword to analyze.
- Basic authentication
### Technologies used
#### Backend
- Go programming language: A highly scalable but also productive programming language.
- Echo: A high performant http library for go.
- Ent: A Entity like ORM for go which is supported by Meta
#### Frontend
- Htmx: high power tools to make HATEOAS application without compromising interactivity from SPA framworks like React, Vue,...
- Bulma CSS: free, open source framework that provides ready-to-use frontend components that you can easily combine to build responsive web interfaces.
- Alpine JS: Minimal and lightweight frontend framework.

### Getting started

## Setup
- After checking out the repository, from within the root, start the Docker containers for the database and cache by executing make up:
``` bash
git clone git@github.com:thiennc1107/search-analyzer.git
cd search-analyzer
make up
```

- After setup the environment, you can start the application by executing `make run`, by default, the application should be able to access in browser at `localhost:8000`. The database scheme will also be automatically migrated.
- If you want to reset all data, execute `make reset`
## Configuration
The config package provides a flexible, extensible way to store all configuration for the application. Configuration is added to the Container as a Service, making it accessible across most of the application.

Be sure to review and adjust all of the default configuration values provided in `config/config.yaml`.
