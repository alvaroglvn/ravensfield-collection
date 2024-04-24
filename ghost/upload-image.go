package ghost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	//"os"
	"strings"

	"github.com/alvaroglvn/ravensfield-collection/internal"
)

// https://res.cloudinary.com/alvaroglvn-image-cloud/image/upload/v1710866517/ravensfield-objects/img-odybhdmIDozdNVGIDKIIVTEY_mczgjd.png

func createMultipartFormData(imgUrl string) (*multipart.Writer, io.Reader) {

	const paramName = "file"

	file, err := http.Get(imgUrl)
	if err != nil {
		log.Println(err)
	}
	defer file.Body.Close()

	fileContents, err := io.ReadAll(file.Body)
	if err != nil {
		log.Println(err)
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	header := make(textproto.MIMEHeader)
	header.Set("Content-Type", "image/webp")
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
		escapeQuotes(paramName), escapeQuotes("object.webp")))
	part, err := writer.CreatePart(header)
	if err != nil {
		log.Println(err)
	}
	_, err = part.Write(fileContents)
	if err != nil {
		log.Println(err)
	}

	params := map[string]string{
		"purpose": "image",
		"ref":     imgUrl,
	}
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		log.Println(err)
	}

	return writer, body
}

func UploadImage(config internal.ApiConfig, imgUrl string) (imageURL string, err error) {
	writer, body := createMultipartFormData(imgUrl)
	var uri = config.GhostURL + "/ghost/api/admin/images/upload/"

	//ghost authorization
	ghostKey := config.GhostKey
	ghostToken, err := CreateAdminToken(ghostKey)
	if err != nil {
		return "", fmt.Errorf("authorization token error: %s", err)
	}

	request, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", "Ghost "+ghostToken)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			fmt.Printf("cannot close body reader: %v\n", err)
		}
	}()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBody := string(content[:])
	if response.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("wrong http code: %d (%q)",
			response.StatusCode, responseBody)
	}

	var images ImageResponse
	if err := json.Unmarshal(content[:], &images); err != nil {
		return "", err
	}

	return images.Images[0].URL, nil
}

type ImageResponse struct {
	Images []Image
}
type Image struct {
	URL string
}

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")
