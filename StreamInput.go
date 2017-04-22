// Package streaminput provides a dbus Input plugin that can integrate with external system
package streaminput

import (
	"bufio"

	"github.com/funkygao/dbus/engine"
	"github.com/funkygao/dbus/pkg/model"
	"github.com/funkygao/golib/pipestream"
	conf "github.com/funkygao/jsconf"
)

var (
	_ engine.Input = &StreamInput{}
)

// StreamInput is a dbus Input plugin that can integrate with external system.
type StreamInput struct {
	cmdAndArgs []string
}

func (this *StreamInput) Init(config *conf.Conf) {
	this.cmdAndArgs = config.StringList("cmd", nil)
	if len(this.cmdAndArgs) == 0 {
		panic("empty cmd in config")
	}
}

func (this *StreamInput) SampleConfig() string {
	return `
	cmd: ["tail", "-F", "/var/log/messages"]
	`
}

func (this *StreamInput) Ack(pack *engine.Packet) error {
	return nil
}

func (this *StreamInput) End(r engine.InputRunner) {}

func (this *StreamInput) Run(r engine.InputRunner, h engine.PluginHelper) error {
	cmd := pipestream.New(this.cmdAndArgs[0], this.cmdAndArgs[1:]...)
	if err := cmd.Open(); err != nil {
		return err
	}
	defer cmd.Close()

	scanner := bufio.NewScanner(cmd.Reader())
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Bytes()

		select {
		case <-r.Stopper():
			return nil
		default:
		}

		pack := <-r.Exchange().InChan()
		pack.Payload = model.Bytes(line)
		r.Exchange().Emit(pack)
	}

	return nil
}

func init() {
	engine.RegisterPlugin("StreamInput", func() engine.Plugin {
		return new(StreamInput)
	})
}
