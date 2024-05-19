package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	godotenv.Load()
	api := slack.New(os.Getenv("BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}

	// fileArr := []string{"resume.pdf"}

	// for i:=0;i<len(fileArr);i++{
	// 	params := slack.FileUploadParameters{
	// 		Channels: channelArr,
	// 		File: fileArr[i],
	// 	}
	// 	file, err := api.UploadFile(params)

	// 	if err !=nil{
	// 		fmt.Println(err)
	// 	}

	// 	fmt.Printf("Name : %s, URL: %s\n", file.Name, file.URL)
	// }
    attachment := slack.Attachment{
        Pretext: "Super Bot Message",
        Text:    "some text",
        Color: "4af030",
        Fields: []slack.AttachmentField{
            {
                Title: "Date",
                Value: time.Now().String(),
            },
        },
    }
 
    _, timestamp, err := api.PostMessage(
        channelArr[0],
 
        slack.MsgOptionAttachments(attachment),
    )
 
    if err != nil {
        panic(err)
    }
    fmt.Printf("Message sent at %s", timestamp)
}