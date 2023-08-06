# RapidDocker
RapidDocker is a tool to quickly create a docker container for your project. Currently used to start database containers for local development.


## Setup
First you'll need to have Go installed. You can download it [here](https://golang.org/dl/).
Along with Docker, which you can download [here](https://docs.docker.com/get-docker/).

Then you can compile RapidDocker by running:
```bash
sudo go build -o /usr/local/bin/rapid-docker
```

## Usage
To setup and run a database container for your project, run:
```bash
rapid-docker
```