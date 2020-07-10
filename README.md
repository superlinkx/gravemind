# Project Gravemind
## A web-based dashboard for keystroke

### Getting Started

#### Build project for development
`go build`

#### Running project in development
This project is a bit old and was written in a rush. There are some things hardcoded that should not have been, but you can override configuration with the `-config` flag:
`./gravemind -config ./configs/dev.json`

Note: The dev json configuration is suitable for running and testing within this directory. For deployments, the `configs/gravemind.json` would be installed in `/etc/gravemind/gravemind.json` and the gravemind binary would be located in `/usr/local/bin/gravemind`

#### Build project for RasPi3
`GOARM=7 GOARCH=arm go build -o ./.build/usr/local/bin/gravemind`

### Flags and Configuration
The `gravemind` binary is built when running a go build. There are two flags it understands:  
`-h \ --help`: Bring up the help information for the command with documented flags  
`-config`: Provide a different configuration file (Default is `/etc/gravemind/gravemind.json`)

Paths to resource files (CSVs from Keystroke and dashboard.html file), port for the Gravemind service, and business name can be found and configured in `configs/gravemind.json`.