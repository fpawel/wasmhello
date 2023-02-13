# Build
```
make run 
```
- Example:
```
APP_PORT=8001 GOCMD=/home/fpawel/sdk/go1.16.7/bin/go make run 
```

- Why specify the path to the **Golang** binary to build instead of use `go` command as usual?

Because of [this](https://github.com/maxence-charriere/go-app/issues/569)