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

	"github.com/alvaroglvn/ravensfield-collection/internal"
)

type ImageResponse struct {
	Images []Image
}
type Image struct {
	URL string
}

func createMultipartFormData(imgContent []byte, fileName string) (*multipart.Writer, io.Reader) {

	//create buffer
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	//set headers for form file field
	formFileHead := make(textproto.MIMEHeader)
	formFileHead.Set("Content-Type", "image/webp")
	formFileHead.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s.webp"`, fileName))
	part, err := writer.CreatePart(formFileHead)
	if err != nil {
		log.Println(err)
	}
	//write image content to form file
	_, err = part.Write(imgContent)
	if err != nil {
		log.Println(err)
	}

	//set purpose field
	if err := writer.WriteField("purpose", "image"); err != nil {
		log.Println(err)
	}

	//close writer
	if err := writer.Close(); err != nil {
		log.Println(err)
	}

	return writer, buf
}

func UploadImage(config internal.ApiConfig, imgContent []byte, fileName string) (imageURL string, err error) {
	writer, body := createMultipartFormData(imgContent, fileName)
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
