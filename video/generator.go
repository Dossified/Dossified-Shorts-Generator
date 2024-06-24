package video

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Dossified/Dossified-Shorts-Generator/logging"
	"github.com/Dossified/Dossified-Shorts-Generator/utils"

	"github.com/u2takey/ffmpeg-go"
	"go.uber.org/zap"
)

const VIDEO_PATH = "output/videos/"
const IMAGE_PATH = "output/screenshots/"
const ASSET_PATH = "assets/"

func CreateVideo(videoMode string) string {
	logging.Info("Creating video")

	createVideoDir()

	videoFiles := make([]string, 0)
	switch videoMode {
	case "news":
		videoFiles = append(videoFiles, createChapterIntroVideo("news"))
		videoFiles = append(videoFiles, createVideoSnippets("news")...)
		break
	case "events":
		videoFiles = append(videoFiles, createChapterIntroVideo("events"))
		videoFiles = append(videoFiles, createVideoSnippets("events")...)
		break
	case "coins":
		break
	}

	mergeVideos(videoFiles)
	removeTemporaryVideos(videoFiles)

	currentPath, err := os.Getwd()
	utils.CheckError(err)
	videoPath := currentPath + "/" + VIDEO_PATH + "out.mp4"
	addBackgroundMusic(videoPath)
	videoPath = currentPath + "/" + VIDEO_PATH + "videoFinal.mp4"
	logging.Info("Video created", zap.String("path", videoPath))
	return videoPath
}

func createVideoSnippets(subFolder string) []string {
	files, err := os.ReadDir(IMAGE_PATH + "/" + subFolder + "/")
	utils.CheckError(err)

	videoFiles := make([]string, 0)
	for _, file := range files {
		if !strings.Contains(file.Name(), ".png") {
			continue
		}
		logging.Debug("Preparing image", zap.String("File", file.Name()))
		videoFilePath := createVideoFromImage(IMAGE_PATH+subFolder+"/", file.Name(), 3)
		videoFiles = append(videoFiles, videoFilePath)
	}
	return videoFiles
}

func createChapterIntroVideo(chapterName string) string {
	videoFilePath := createVideoFromImage(ASSET_PATH, chapterName+".png", 3)
	return videoFilePath
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

func createVideoFromImage(imagePath string, imageFileName string, length int) string {
	currentPath, err := os.Getwd()
	utils.CheckError(err)
	outputFilePath := currentPath + "/" + VIDEO_PATH + strings.Replace(
		imageFileName,
		".png",
		"",
		-1,
	) + ".mp4"
	frameRate := 30
	fadeOutFrame := (frameRate * length) - frameRate
	ffmpeg_go.Input(
		imagePath+imageFileName,
		ffmpeg_go.KwArgs{
			"t":         fmt.Sprint(length),
			"loop":      "1",
			"framerate": fmt.Sprint(frameRate),
		}).
		Output(
			outputFilePath,
			ffmpeg_go.KwArgs{
				"vf": "fade=in:0:30, fade=out:" + fmt.Sprint(fadeOutFrame) + ":30",
			}).
		OverWriteOutput().
		ErrorToStdOut().
		Run()
	removeOldScreenshots(imagePath, imageFileName)
	return outputFilePath
}

func removeOldScreenshots(imagePath string, imageFileName string) {
	if imagePath == ASSET_PATH {
		return
	}
	err := os.Remove(imagePath + imageFileName)
	utils.CheckError(err)
}

func addBackgroundMusic(videoPath string) {
	currentPath, err := os.Getwd()
	utils.CheckError(err)
	inputVideo := ffmpeg_go.Input(videoPath)
	inputAudio := ffmpeg_go.Input(currentPath + "/assets/bg_music.wav", ffmpeg_go.KwArgs{
        "stream_loop": "-1",
    })
	out := ffmpeg_go.Output(
		[]*ffmpeg_go.Stream{inputVideo, inputAudio},
		VIDEO_PATH+"videoFinal.mp4",
		ffmpeg_go.KwArgs{
			"map":      "1:a",
			"c:v":      "copy",
			"shortest": "",
		},
	)
	out.OverWriteOutput().ErrorToStdOut().Run()
}
