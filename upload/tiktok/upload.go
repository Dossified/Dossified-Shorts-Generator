package tiktok

import (
    "log"
    "bytes"
    "io"
    "encoding/json"
    "net/http"
)

func UploadToTikTok(filepath string) {

}

func authenticateTikTok() {

}

func queryCreatorInfo() {
    creatorInfoEndpoint := "https://open.tiktokapis.com/v2/post/publish/creator_info/query/"

}

func uploadVideoRequest() {
    //Encode the data
    postBody, _ := json.Marshal(map[string]map[string]string {
        "source_info": {
            "source": "FILE_UPLOAD",
            "video_size": "0", // ToDo
            "chunk_size": "0", // ToDo
            "total_chunk_count": "0", // ToDo
        },
    })
    responseBody := bytes.NewBuffer(postBody)
    //Leverage Go's HTTP Post function to make request
    resp, err := http.Post(
        "https://open.tiktokapis.com/v2/post/publish/inbox/video/init/",
        "application/json",
        responseBody,
    )
    //Handle Error
    if err != nil {
        log.Fatalf("An Error Occured %v", err)
    }
    defer resp.Body.Close()
    //Read the response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    sb := string(body)
    log.Printf(sb)
}

func uploadVideo() {

}
