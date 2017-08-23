package main

import (
	"io"
	"log"
	"os/exec"

	"github.com/gin-gonic/gin"
)

// type CameraFrameProvider interface {
// 	Listen(chan<- []byte)
// 	Stop(chan<- []byte)
// }

// type DesktopCameraFrameProvider struct {
// 	bcast broadcast.Broadcaster
// }

// func videoProcess(bcast broadcast.Broadcaster) {
// 	// raspivid := exec.Command("raspivid", "-t", "0", "-w", "640", "-h", "480", "-fps", "30", "-o", "-", "-pf", "baseline")
// 	raspivid := exec.Command("bash", "-c", "while :; do dd if=sample.h264; sleep 20; done")
// 	raspividOut, err := raspivid.StdoutPipe()

// 	if err != nil {
// 		log.Panicln(err)
// 	}

// 	err = raspivid.Start()
// 	if err != nil {
// 		log.Panicln(err)
// 	}

// 	// Kill the raspivid process on client exit
// 	defer raspivid.Process.Kill()

// 	buf := make([]byte, 10*1024*1024)
// 	for {
// 		n, _ := raspividOut.Read(buf)
// 		if n > 0 {
// 			bcast.Submit(buf[:n])
// 		}
// 		time.Sleep(20 * time.Millisecond)
// 	}
// }

func apiCameraMount(app *App) {
	app.Router.GET("/camera", func(c *gin.Context) {
		raspivid := exec.Command("raspivid", "-t", "0", "-w", "640", "-h", "480", "-fps", "20", "-b", "500000", "-g", "10", "-o", "-", "-pf", "baseline")
		// raspivid := exec.Command("dd", "if=/dev/urandom", "count=100000")
		raspividOut, err := raspivid.StdoutPipe()

		if err != nil {
			log.Panicln(err)
		}

		err = raspivid.Start()
		if err != nil {
			log.Panicln(err)
		}

		// Kill the raspivid process on client exit
		defer raspivid.Process.Kill()

		buf := make([]byte, 10*1024*1024)
		c.Stream(func(w io.Writer) bool {
			n, _ := raspividOut.Read(buf)
			// Check the buffer doesn't grow out of control while we were writing to
			// the client.
			if n > 0 && n < 50*1024 {
				w.Write(buf[:n])
			}
			return true
		})
	})
}

// func apiCameraMount2(app *App) {
// 	bcast := broadcast.NewBroadcaster(1)
// 	go videoProcess(bcast)

// 	app.Router.GET("/camera", func(c *gin.Context) {
// 		ch := make(chan interface{})
// 		bcast.Register(ch)
// 		defer bcast.Unregister(ch)

// 		c.Header("Content-Type", "video/H264")
// 		c.Stream(func(w io.Writer) bool {
// 			select {
// 			case frame := <-ch:
// 				w.Write(frame.([]byte))
// 			}
// 			return true
// 		})
// 	})
// }
