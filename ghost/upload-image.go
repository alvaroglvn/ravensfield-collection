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
		log.Fatal(err)
	}
	//write image content to form file
	_, err = part.Write(imgContent)
	if err != nil {
		log.Fatal(err)
	}

	//set purpose field
	if err := writer.WriteField("purpose", "image"); err != nil {
		log.Fatal(err)
	}

	//close writer
	if err := writer.Close(); err != nil {
		log.Fatal(err)
	}

	return writer, buf
}

func UploadImage(config internal.ApiConfig, imgContent []byte, fileName string) (imageURL string, err error) {
	//turn data into multipart dataform
	writer, body := createMultipartFormData(imgContent, fileName)
	var uri = config.GhostURL + "/ghost/api/admin/images/upload/"

	//ghost authorization
	ghostKey := config.GhostKey
	ghostToken, err := CreateAdminToken(ghostKey)
	if err != nil {
		return "", fmt.Errorf("authorization token error: %s", err)
	}

	//send request
	request, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return "", fmt.Errorf("request error: %s", err)
	}
	//set headers
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", "Ghost "+ghostToken)

	//get response
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("response error: %s", err)
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("reading error: %s", err)
	}

	var images ImageResponse
	if err := json.Unmarshal(content[:], &images); err != nil {
		return "", fmt.Errorf("unmarshalling err: %s", err)
	}

	return images.Images[0].URL, nil
}
