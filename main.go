package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6072266060418-6115048569936-9UH3QR47X6CG8hEHloE5ut5u")
	os.Setenv("CHANNEL_ID", "C0626M16TFE")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv(("CHANNEL_ID"))}
	fileArr := []string{"ansys.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name:%s, URL:%s\n", file.Name, file.URL)
	}

}
