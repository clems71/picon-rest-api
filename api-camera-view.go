package main

import (
	"io"
	"log"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
)

func apiCameraMount(app *App) {
	app.Router.GET("/camera", func(c *gin.Context) {
		raspivid := exec.Command("raspivid", "-t", "0", "-w", "640", "-h", "480", "-fps", "30", "-o", "-", "-pf", "baseline")
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
			if n > 0 {
				w.Write(buf[:n])
			}
			time.Sleep(20 * time.Millisecond)
			return true
		})
	})
}
