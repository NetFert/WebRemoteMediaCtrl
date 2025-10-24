package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-vgo/robotgo"
)

func Media(c *gin.Context) {
	operation := c.Query("operation") //play pause stop previous next volume_up volume_down mute un_mute

	var err error
	switch operation {
	case "play":
		err = robotgo.KeyTap(robotgo.AudioPlay)
	case "pause":
		err = robotgo.KeyTap(robotgo.AudioPause)
	case "stop":
		err = robotgo.KeyTap(robotgo.AudioStop)
	case "previous":
		err = robotgo.KeyTap(robotgo.AudioPrev)
	case "next":
		err = robotgo.KeyTap(robotgo.AudioNext)
	case "volume_up":
		err = robotgo.KeyTap(robotgo.AudioVolUp)
	case "volume_down":
		err = robotgo.KeyTap(robotgo.AudioVolDown)
	case "mute":
		err = robotgo.KeyTap(robotgo.AudioMute)
	case "un_mute":
		err = robotgo.KeyTap(robotgo.AudioMute)
	default:
		err = errors.New("Unknown operation" + operation)
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  false,
			"msg": "Operation failed:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":  true,
		"msg": "Success",
	})
}
