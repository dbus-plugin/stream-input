# stream-input
a [dbus](https://github.com/funkygao/dbus) Input plugin that can integrate with external system.

### Usage

stream-input is an implementation of [dbus Input](https://github.com/funkygao/dbus/blob/master/engine/input.go#L17) interface. You only need to import the driver.

[input bootstrap](https://github.com/funkygao/dbus/blob/master/plugins/input/bootstrap.go)

```go
import _ "github.com/dbus-plugin/stream-input"
```

### Configurattion

- cmd

Example:

```
cmd: ["tail", "-F', "/var/log/messages"]
```
