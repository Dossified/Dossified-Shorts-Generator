// Contains all function related to video editing & rendering
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

// Generates the video from the previously taken screenshots
func CreateVideo(videoMode string) (string, string) {
	logging.Info("Creating video")

    // Make sure output directory exists
	createVideoDir()

    // First we create a list of all images we want to concatenate
    // This include intro & outro images
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

    // Here we merge the images into videos
    // We create 2 videos bc. YouTube has a short video limit of 60 seconds
    // Therefore we create a second video which gets cut after the 60s mark
	mergeVideos(videoFiles, false)
	mergeVideos(videoFiles, true)
	removeTemporaryVideos(videoFiles)

	currentPath, err := os.Getwd()
	utils.CheckError(err)

    // Add background music to the video
	videoPath := currentPath + "/" + VIDEO_PATH + "out.mp4"
	addBackgroundMusic(videoPath, false)

    // Add background music to the video
	videoPathYT := currentPath + "/" + VIDEO_PATH + "outYT.mp4"
	addBackgroundMusic(videoPathYT, true)
	videoPath = currentPath + "/" + VIDEO_PATH + "videoFinal.mp4"
	logging.Info("Video created", zap.String("path", videoPath))
	return videoPath, videoPathYT
}

// Adds all image file paths of our screenshots to a list of strings
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

// Creates a video from the intro image
func createChapterIntroVideo(chapterName string) string {
	videoFilePath := createVideoFromImage(ASSET_PATH, chapterName+".png", 3)
	return videoFilePath
}

// Creates video output directory if it does not exist yet
func createVideoDir() string {
	path := filepath.Join(".", VIDEO_PATH)
	err := os.MkdirAll(path, os.ModePerm)
	utils.CheckError(err)
	logging.Debug("Video path", zap.String("path", path))
	return path
}

// Deletes temporary video files
func removeTemporaryVideos(videoFiles []string) {
	for _, file := range videoFiles {
		err := os.Remove(file)
		utils.CheckError(err)
	}
}

// Merges out seperate videos for intro & content into a single video
func mergeVideos(videoFiles []string, youtubeMode bool) {
	videoPathsTextFile, err := os.Create(VIDEO_PATH + "videos")
	utils.CheckError(err)

	for _, file := range videoFiles {
		videoPathsTextFile.WriteString("file '" + file + "'\n")
	}
	kwargs := ffmpeg_go.KwArgs{
		"f":    "concat",
		"safe": "0",
	}
	fileName := "out.mp4"
	if youtubeMode {
		kwargs["t"] = "60"
		fileName = "outYT.mp4"
	}
	ffmpeg_go.Input(
		VIDEO_PATH+"videos",
		kwargs,
	).Output(VIDEO_PATH + fileName).OverWriteOutput().ErrorToStdOut().Run()
	os.Remove(VIDEO_PATH + "videos")
}

// Generates a video with fade from the passed image path
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

    // Delete the screenshot after we're done
	removeOldScreenshots(imagePath, imageFileName)
	return outputFilePath
}

// Deletes input image by path
func removeOldScreenshots(imagePath string, imageFileName string) {
	if imagePath == ASSET_PATH {
		return
	}
	err := os.Remove(imagePath + imageFileName)
	utils.CheckError(err)
}

// Adds background music to the input video
func addBackgroundMusic(videoPath string, youtubeMode bool) {
	currentPath, err := os.Getwd()
	utils.CheckError(err)
	inputVideo := ffmpeg_go.Input(videoPath)
	inputAudio := ffmpeg_go.Input(currentPath+"/assets/bg_music.wav", ffmpeg_go.KwArgs{
		"stream_loop": "-1",
	})
	outFileName := "videoFinal.mp4"
	if youtubeMode {
		outFileName = "videoFinalYT.mp4"
	}
	out := ffmpeg_go.Output(
		[]*ffmpeg_go.Stream{inputVideo, inputAudio},
		VIDEO_PATH+outFileName,
		ffmpeg_go.KwArgs{
			"map":      "1:a",
			"c:v":      "copy",
			"shortest": "",
		},
	)
	out.OverWriteOutput().ErrorToStdOut().Run()
}
