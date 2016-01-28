package status

func SetupHeartbeat(statusFunction func() bool, stopStatus *chan []byte) {
    var running bool = true

    // setup the timer
    timeout := make(chan bool, 1)
    go func() {
        for ; running ; {
            time.Sleep(5 * time.Second)
            timeout <- true
        }
    }()

    for ; running ; {
        select{
            // send a heartbeat (stdout and wait for this example)
            case data, _ := <- timeout
                log.Println("I'm still here!!")

            // listen for the stop
            case _, ok := <-*stop:
                if ok || !ok {
                    running = false
                }
        }
    }
}
