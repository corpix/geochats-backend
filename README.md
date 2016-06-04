go-boilerplate
--------------
This is a golang headless project boilerplate.

# Usecases
* Quick way to bootstrap a golang tool or server project

# Install
Make sure your [GOPATH](https://golang.org/doc/code.html#GOPATH) is ok.

You could install boilerplate as usual golang module:
```shell
go get -u github.com/corpix/go-boilerplate/...
```

But this is not so useful, **more often you want the project based on this boilerplate**.

Create a package:

``` shellsession
cd "${GOPATH}/src"
mkdir -p github.com/$USER/my-awesome-project
cd github.com/$USER/my-awesome-project
```

On Linux/OS X:

``` shell
curl -Ls https://raw.githubusercontent.com/corpix/go-boilerplate/0247ae4d8164260aa6e6ed02bef17c4d39cc2c67/init.sh | bash
```

You are done!

# Usage example

``` shell
go-boilerplate -h
NAME:
   go-boilerplate - APP_USAGE

USAGE:
   go-boilerplate [global options] command [command options] [arguments...]

AUTHOR(S):
   AUTHOR_NAME <AUTHOR_EMAIL>

COMMANDS:
     greet, g  greets the user

GLOBAL OPTIONS:
   --debug                      debug mode [$DEBUG]
   --log-level value, -l value  log level(debug, info, warn, error, fatal, panic) (default: "info")
   --log-file value             log file to log entries to
   --log-file-mode value        octal mode for log file (default: 384)
   --log-formatter value        log formatter to use(available: json or none)
   --config value               path to configuration file (default: "config.xml")
   --help, -h                   show help
   --version, -v                print the version
```

# Hack on it

* Fork this project
* Clone it
* Make a branch
* Hack on it!
* Test it
* Push your fork
* Pull request your awesome feature or fix!

``` shellsession
mkdir -p $GOPATH/src/github.com/$USER
git clone git@github.com:$USER/go-boilerplate.git $GOPATH/src/github.com/$USER
cd $_
git checkout -b awesome-feature
# Hacking in progress ...
make test
git add .
git commit -m 'Added awesome feature for ...'
git push
# Create a pull request!
```
