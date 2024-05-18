package video

import (
	"github.com/u2takey/ffmpeg-go"
)

func createVideo() {

	ffmpeg_go.Input("*.png", ffmpeg_go.KwArgs{"pattern_type": "glob"}).
		Output("out.mp4").
		OverWriteOutput().
		ErrorToStdOut().
		Run()
}
