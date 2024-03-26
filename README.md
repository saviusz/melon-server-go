# Melon Server (Go)
This projects serves as backend for songbook system build by me.

## Development
To start developing you must have installed:
* go runtime
* [Taskfile](https://taskfile.dev/installation/) buildrunner


### Installing

First you need to clone the repo:
```sh
git clone https://github.com/saviusz/melon-server-go.git
```

and then install deps
```sh
task install
```

>[!NOTE]
> If you have problems with running `task watch` directly after installation, try restarting the terminal

### Running

To run project you can use:
```sh
task run
```

But, you'd probably prefer to use watchdog:
```sh
task watch
```

### Testing
To run test cases use:
```sh
task test
```

### Building
To build project use:
```sh
task build
```

Built binary would be located in `/dist/` folder