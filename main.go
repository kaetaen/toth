package main

import (
	"log"
	"os"

	"github.com/gen2brain/beeep"
	"github.com/otiai10/gosseract/v2"
	"golang.design/x/clipboard"
)

const FILE_PREFIX = "*.png"

func main() {
	USER_DIR, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	for {
		err := clipboard.Init()
		if err != nil {
			panic(err)
		}

		copiedImgByte := clipboard.Read(clipboard.FmtImage)
		if len(copiedImgByte) == 0 {
			log.Println("No image found on clipboard.")
			continue
		}

		tmpFile, err := os.CreateTemp(USER_DIR, FILE_PREFIX)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := tmpFile.Write(copiedImgByte); err != nil {
			log.Fatal(err)
		}
		tmpFile.Close()

		tesseractClient := gosseract.NewClient()

		tesseractClient.SetImage(tmpFile.Name())
		extractedText, err := tesseractClient.Text()
		if (err) != nil {
			panic(err)
		}

		clipboard.Write(clipboard.FmtText, []byte(extractedText))

		beeep.Notify("✅ Your text is ready", "📋 Just paste", tmpFile.Name())

		tesseractClient.Close()
		os.Remove(tmpFile.Name())
	}

}
