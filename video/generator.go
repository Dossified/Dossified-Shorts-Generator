package video

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Dominique-Roth/Dossified-Shorts-Generator/logging"
	"github.com/Dominique-Roth/Dossified-Shorts-Generator/utils"

	"github.com/u2takey/ffmpeg-go"
	"go.uber.org/zap"
)

const VIDEO_PATH = "output/videos/"
const IMAGE_PATH = "output/screenshots/"

func CreateVideo() {
    logging.Info("Creating video")

    createVideoDir()

    files, err := os.ReadDir(IMAGE_PATH)
    utils.CheckError(err)

    for _, file := range files {
        if !strings.Contains(file.Name(), ".png") {
            continue
        }
        logging.Debug("Preparing image", zap.String("File", file.Name()))
        prepareImage(file.Name())
    }
    //ffmpeg_go.Input(
    //    "output/screenshots/*.png",
    //    ffmpeg_go.KwArgs{
    //        "pattern_type": "glob",
    //        "framerate": "1/6",
    //    }).
	//	Output(VIDEO_PATH).
	//	OverWriteOutput().
	//	ErrorToStdOut().
	//	Run()
    logging.Info("Video created", zap.String("path", VIDEO_PATH))
}

func createVideoDir() string {
	path := filepath.Join(".", VIDEO_PATH)
	err := os.MkdirAll(path, os.ModePerm)
	utils.CheckError(err)
	logging.Debug("Screenshot path", zap.String("path", path))
	return path
}

func mergeVideos() {

}

func prepareImage(imageFileName string) {

    //videoFilePath := createVideoFromImage(imageFileName)
    createVideoFromImage(imageFileName)
    //addFadeIn(videoFilePath)
}

func createVideoFromImage(imageFileName string) string {
	executable_path, err := os.Getwd()
	utils.CheckError(err)
    outputFilePath := executable_path + "/" + VIDEO_PATH + strings.Replace(imageFileName, ".png", "", -1) + "RAW.mp4" 
    ffmpeg_go.Input(
            IMAGE_PATH + imageFileName,
            ffmpeg_go.KwArgs{
                "t": "6",
                "loop": "1",
                "framerate": "30",
        }).
		Output(
            outputFilePath,
            ffmpeg_go.KwArgs{
                "vf": "fade=in:0:30, fade=out:150:30",
        }).
		OverWriteOutput().
		ErrorToStdOut().
		Run()
    return outputFilePath
}

func addBackgroundMusic() {}

