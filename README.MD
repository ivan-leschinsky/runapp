# runapp

Simple app to ensure some app is running on macOS/linux

![Go version](https://img.shields.io/badge/go-1.18-blue)

---

### Install from script

```sh
curl -s https://raw.githubusercontent.com/ivan-leschinsky/runapp/master/install.sh | bash
````

App is intended to run periodically with cron

### Run in development:

```sh
go mod download
go run .
```

### Install from development
```sh
go build .
go install
```
