package cli

import (
  "os"
  "os/signal"
  "syscall"

  "github.com/mitchellh/cli"

  "github.com/tyrchen/podgen/cli/command"
)

// Commands returns the mapping of CLI commands. The meta
// parameter lets you set meta options for all commands.
func Commands(metaPtr *command.Meta) map[string]cli.CommandFactory {
  if metaPtr == nil {
    metaPtr = new(command.Meta)
  }

  meta := *metaPtr
  if meta.Ui == nil {
    meta.Ui = &cli.BasicUi{
      Writer:      os.Stdout,
      ErrorWriter: os.Stderr,
    }
  }

  return map[string]cli.CommandFactory{
    "init": func() (cli.Command, error) {
      return &command.InitCommand{
        Meta: meta,
      }, nil
    },
  }
}

// makeShutdownCh returns a channel that can be used for shutdown
// notifications for commands. This channel will send a message for every
// interrupt or SIGTERM received.
func makeShutdownCh() <-chan struct{} {
  resultCh := make(chan struct{})

  signalCh := make(chan os.Signal, 4)
  signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
  go func() {
    for {
      <-signalCh
      resultCh <- struct{}{}
    }
  }()
  return resultCh
}
