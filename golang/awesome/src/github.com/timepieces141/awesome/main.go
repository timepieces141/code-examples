package main

import (
    "flag"
    "github.com/timepieces141/awesome/configuration"
    "github.com/timepieces141/awesome/heartbeat"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"
)

var status bool = false

func main() {
    log.Println("--------------------------------")
    log.Println("Awesome! starting up...  ")
    log.Println("--------------------------------")

    // check for a command line flag that tells us where the config file is,
    // else assume the default location
    var configFile string
    flag.StringVar(&configFile, "config-file", configuration.DEFAULT_CONFIG_FILE, "Location of the configuration file");
    flag.Parse()

    // channel for stopping loops inside services
    stop := make(chan []byte, 0)

    // setup status for heartbeat service
    statusFunction := func() bool {
        return status
    }
    go heartbeat.SetupHeartbeat(statusFunction, &stop)

    // more services could be established here, and if they have "alive" loops,
    // then they can also have access to "stop"

    // we're ready
    log.Println("Awesome is ready.")
    status = true

    // listen for the terminate, interrupt, kill signals and block for it
    sigChannel := make(chan os.Signal, 1)
    signal.Notify(sigChannel, os.Interrupt, os.Kill, syscall.SIGTERM)
    s := <- sigChannel
    log.Println("Signal Received:", s)

    // break all the loops
    status = false
    close(stop)

    time.Sleep(1 * time.Second)
    log.Println("--------------------------------")
    log.Println("Awesome shutting down...")
    log.Println("--------------------------------")
}
