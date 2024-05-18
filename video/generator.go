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

func CreateVideo() string {
	logging.Info("Creating video")

	createVideoDir()

	files, err := os.ReadDir(IMAGE_PATH)
	utils.CheckError(err)

	videoFiles := make([]string, 0)
	for _, file := range files {
		if !strings.Contains(file.Name(), ".png") {
			continue
		}
		logging.Debug("Preparing image", zap.String("File", file.Name()))
		videoFilePath := createVideoFromImage(file.Name())
		videoFiles = append(videoFiles, videoFilePath)
	}
	mergeVideos(videoFiles)
	removeTemporaryVideos(videoFiles)

	currentPath, err := os.Getwd()
	utils.CheckError(err)
	videoPath := currentPath + "/" + VIDEO_PATH + "out.mp4"
	logging.Info("Video created", zap.String("path", videoPath))
	return videoPath
}

func createVideoDir() string {
	path := filepath.Join(".", VIDEO_PATH)
	err := os.MkdirAll(path, os.ModePerm)
	utils.CheckError(err)
	logging.Debug("Video path", zap.String("path", path))
	return path
}

func removeTemporaryVideos(videoFiles []string) {
	for _, file := range videoFiles {
		err := os.Remove(file)
		utils.CheckError(err)
	}
}

func mergeVideos(videoFiles []string) {

	//executable_path, err := os.Getwd()
	//utils.CheckError(err)
	//outputFilePath := executable_path + "/" + VIDEO_PATH + "*.mp4"

	videoPathsTextFile, err := os.Create(VIDEO_PATH + "videos")
	utils.CheckError(err)

	for _, file := range videoFiles {
		videoPathsTextFile.WriteString("file '" + file + "'\n")
	}
	ffmpeg_go.Input(
		VIDEO_PATH+"videos",
		ffmpeg_go.KwArgs{
			"f":    "concat",
			"safe": "0",
		}).Output(VIDEO_PATH + "out.mp4").OverWriteOutput().ErrorToStdOut().Run()
	os.Remove(VIDEO_PATH + "videos")
}

func createVideoFromImage(imageFileName string) string {
	currentPath, err := os.Getwd()
	utils.CheckError(err)
	outputFilePath := currentPath + "/" + VIDEO_PATH + strings.Replace(
		imageFileName,
		".png",
		"",
		-1,
	) + ".mp4"
	ffmpeg_go.Input(
		IMAGE_PATH+imageFileName,
		ffmpeg_go.KwArgs{
			"t":         "6",
			"loop":      "1",
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
