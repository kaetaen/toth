package main

import (
	"log"
	"os"

	"github.com/gen2brain/beeep"
	"github.com/otiai10/gosseract/v2"
	"golang.design/x/clipboard"
)

func getImage() ([]byte, bool) {
	imgByte := clipboard.Read(clipboard.FmtImage)
	imgNotFound := false

	if len(imgByte) == 0 {
		imgNotFound = true
	}
	return imgByte, imgNotFound
}

func createTmpImg(img []byte) *os.File {
	FILE_PREFIX := "*.png"
	USER_DIR, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	tmpFile, err := os.CreateTemp(USER_DIR, FILE_PREFIX)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmpFile.Write(img); err != nil {
		log.Fatal(err)
	}
	tmpFile.Close()

	return tmpFile
}

func tesseractOCRProcessor(imgFile *os.File) string {
	tesseractClient := gosseract.NewClient()
	tesseractClient.SetImage(imgFile.Name())

	extractedText, err := tesseractClient.Text()
	if (err) != nil {
		panic(err)
	}

	tesseractClient.Close()
	return extractedText
}

func ExtractTextFromImgClipboard() {
	for {
		err := clipboard.Init()
		if err != nil {
			panic(err)
		}

		copiedImgByte, imgNotFound := getImage()
		if imgNotFound {
			log.Println("No image found on clipboard.")
			continue
		}

		tmpImg := createTmpImg(copiedImgByte)
		imgText := tesseractOCRProcessor(tmpImg)

		clipboard.Write(clipboard.FmtText, []byte(imgText))

		beeep.Notify("✅ Your text is ready", "📋 Just paste", tmpImg.Name())

		os.Remove(tmpImg.Name())
	}

}

func main() {
	ExtractTextFromImgClipboard()
}
