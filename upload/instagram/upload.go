// Contains all function related to uploading a Reel to Instagram
package instagram

import (
	"io"
	"os"

	"github.com/Dossified/Dossified-Shorts-Generator/config"
	"github.com/Dossified/Dossified-Shorts-Generator/logging"
	"github.com/Dossified/Dossified-Shorts-Generator/utils"
	"github.com/Dossified/Dossified-Shorts-Generator/video"

	"github.com/Davincible/goinsta/v3"
)

// Authenticates with Instagram & automatically uploads input file
func UploadToInstagram(filepath string, videoMode string) {
	logging.Info("Instagram upload initiated")
	logging.Info("Logging into Instagram")
    // Authentication
	instagramUsername := config.GetConfiguration().InstagramUsername
	instagramPassword := config.GetConfiguration().InstagramPassword
	insta := goinsta.New(instagramUsername, instagramPassword)
	err := insta.Login()
	utils.CheckError(err)

	defer insta.Export("./.goinsta")

    // Load file & upload
	logging.Info("Loading video file for Instagram upload")
	file := loadFile(filepath)
	logging.Info("Uploading to Instagram")
	upload(file, insta, videoMode)
	logging.Info("Upload to Instagram successful")
}

// Loads video file info memory
func loadFile(filepath string) io.Reader {
	file, err := os.Open(filepath)
	utils.CheckError(err)
	return file
}

// Uploads video to instagram
func upload(
	file io.Reader,
	insta *goinsta.Instagram,
	videoMode string,
) {
	videoTitle := video.GetVideoTitle(videoMode)
	_, err := insta.Upload(
		&goinsta.UploadOptions{
			File:    file,
			Caption: videoTitle,
			UserTags: &[]goinsta.UserTag{
				{
					User: &goinsta.User{
						ID:        insta.Account.ID,
						IsPrivate: true,
					},
				},
			},
		},
	)
	utils.CheckError(err)
}
