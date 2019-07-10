## bark-helper-go

A small tool for syncing cliptext from Windows to iOS.

### Usage

```
Usage of bark-helper-go.exe:
  -s string
        service commands [install|uninstall|start|stop|run] (default "run")
  -t string
        set token
```

examples:

```
bark-helper-go.exe -t "token" // common way to start

bark-helper-go.exe -t "token" -c "install" // install as a service
bark-helper-go.exe -t "token" -c "start" // start the service
```