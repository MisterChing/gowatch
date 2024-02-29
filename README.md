# gowatch
![Go](https://github.com/silenceper/gowatch/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/silenceper/gowatch)](https://goreportcard.com/report/github.com/silenceper/gowatch)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/silenceper/gowatch)](https://pkg.go.dev/github.com/silenceper/gowatch)
[![codecov](https://codecov.io/gh/silenceper/gowatch/branch/master/graph/badge.svg)](https://codecov.io/gh/silenceper/gowatch)

[中文文档](./README_ZH_CN.md)

gowatch is a command line tool that builds and (re)starts your go project everytime you save a Go or template file.

![gowatch](./screenshot/gowatch.png)

## Installation
To install `gowatch` use the `go insall` command:

```go
go install github.com/silenceper/gowatch@latest
```

Then you can add `gowatch` binary to PATH environment variable in your ~/.bashrc or ~/.bash_profile file:

>If you already have `gowatch` installed, updating `gowatch` is simple:

```
go get -u github.com/silenceper/gowatch
```
## Usage
```
cd /path/to/myapp
```
Start gowatch:

```
gowatch
```


Gowatch will watch for file events, and every time you create/modify/delete a file it will build and restart the application,If `go build` returns an error, it will log it in stdout.

### Support Options

- -o : Not required, specify the target file path for the build
- -p : Not required, specify the package to be built (can also be a single file)
- -args: Not required, specify program runtime parameters, for example: -args = '-host =: 8080, -name = demo'
- -v: Not required, display gowatch version information
- -l: Not required, specify the log level, default is debug
- -h: Not required, show usage

example:

`gowatch -o ./bin/demo -p ./cmd/demo`

### Configuration file

In most cases, you don't need to specify the configuration. You can meet most of the requirements by directly executing the `gowatch` command.

You can initialize the gowatch.yml configuration file to the current directory with the default settings running the following command.
```
gowatch init
```
Create a `gowatch.yml` file in the execution directory:
```
# gowatch.yml configuration example

# The name of the executable file generated under the current directory execution. The default is the current directory name.
appname: "test"

# Specify the command to run after builds done
run_cmd: "./run.sh"

# Specify the directory where the compiled object files are stored
output: /bin/demo
# The file name suffix that needs to be monitored. By default, there is only a '.go' file.
watch_exts:
    - .yml

# The directory that needs to listen for file changes. By default, only the current directory.
watch_paths:
    - ../pk

# Additional parameters that need to be added when running the application
cmd_args:
    - arg1=val1

# Additional parameters that need to be added when building the application
build_args:
    - -race

# Need to increase environment variables, the current environment variables are loaded by default
envs:
    - a=b

# Whether to listen to file changes in the 'vendor' folder
vendor_watch: false

# Directory that do not need to listen for file changes
excluded_paths:
    - path

# main package path, can also be a single file, multiple files separated by commas
build_pkg: ""

# build tags
build_tags: ""

# Commands that can be executed before build the app
#prev_build_cmds:
#  - swag init

# Whether to prohibit automatic operation
disable_run: false

# use GOGC on build
build_go_gc: false

# log level, support debug, info, warn, error, fatal
log_level: "debug"
```

## Author
[@silenceper](http://silenceper.com)


>Inspired by [bee](https://github.com/beego/bee)
