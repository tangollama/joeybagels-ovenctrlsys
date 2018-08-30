# New Relic Infrastructure Integration for ovenctrlsys

Reports status and metrics for ovenctrlsys service

## Requirements
These instructions assumes that one is deploying the code to a Linux-based VM (specifically Amazon Linux 2). If you've not deployed the New Relic Infrastructure agent to an EC2 instance, look to [this post](https://blog.newrelic.com/product-news/installing-infrastructure-agent-aws/) for simple instructions. It also assumes development on Mac OSX. Both of these "requirements" can easily be modified for your environment, but I'm writing the instructions for my setup. 

- Deploy the [New Relic Infrastructure agent](https://newrelic.com/products/infrastructure)
- Setup [Go](https://golang.org/) on your dev environment
- Setup make (I recommend using [homebrew](https://brew.sh/)) `brew install make`
- Run cmd `git clone https://github.com/tangollama/joeybagels-ovenctrlsys.git` in your $GOPATH/src directory
- cmd `cd joeybagels-ovenctrlsys`
- cmd `make compile`

## Installation

Assuming, you can execute `./bin/jb-overctrlsys` on your target environment (if you can't, you'll need to run `make compile` on the target host or - better yet - pass the target environment variables to the go build command i.e. `env GOOS=linux GOARCH=amd64 go build -o bin/jb-ovenctrlsys src/ovenctrlsys.go src/testdata.go`), then the following need to occur on your host that is already running the New Relic Infrastructure agent:

1. Copy the file bin/jb-overctrlsys into the executable path of the host (i.e. put it in a directory listed in `echo $PATH`). I recommend putting the executable in the same directory as newrelic-infra. (You can find this by running a `which newrelic-infra` on the host.) `sudo cp ./bin/jb-overctrlsys /usr/bin`
2. Copy jb-ovenctrlsys-config.yml into /etc/newrelic-infra/integrations.d `sudo cp jb-ovenctrlsys-config.yml /etc/newrelic-infra/integrations.d`
3. Copy jb-ovenctrlsys-definition.yml into /var/db/newrelic-infra/custom-integrations `sudo cp jb-ovenctrlsys-definition.yml /var/db/newrelic-infra/custom-integrations`
4. Restart the newrelic-infra agent `sudo systemctl restart newrelic-infra.service`

## Usage

`./jb-overctrlsys`

## Compatibility

* Supported OS: Amazon Linux 2
* ovenctrlsys versions: 0.1.0
* Edition: Example
