# stream-input
a [dbus](https://github.com/funkygao/dbus) Input plugin that can integrate with external system.

### Usage

stream-input is an implementation of [dbus Input](https://github.com/funkygao/dbus/blob/master/engine/input.go#L17) interface. You only need to import the plugin.

```go
import _ "github.com/dbus-plugin/stream-input"
```

### Configurattion

- cmd

Example:

```
cmd: ["tail", "-F', "/var/log/messages"]
```

#### A complete example

```
{
    plugins: [
        {
            name: "StreamInput"
            cmd: ["cat", "/var/log/mail.log"]
        }

        {
            name:   "MockOutput"
            blackhole: false
            match:  ["StreamInput", ]
        }
    ]
}
```
