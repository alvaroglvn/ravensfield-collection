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
	"os"
	"strings"

	"github.com/alvaroglvn/ravensfield-collection/internal"
)

// func UploadImgToGhost(imageData []byte, config internal.ApiConfig) error {

// 	//STEP 1 - MULTIPART DATAFORM

// 	var buffer bytes.Buffer
// 	multiwriter := multipart.NewWriter(&buffer)

// 	// //add form file
// 	// formFile, err := multiwriter.CreateFormFile("file", "image.webp")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// //write blob to form
// 	// if _, err := formFile.Write(imageData); err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	header := make(textproto.MIMEHeader)
// 	header.Set("Content-Type", "image/webp")
// 	header.Set("Content-Disposition", "form-data; name=\"file\"; filename=\"file.webp\"")
// 	part, err := multiwriter.CreatePart(header)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = part.Write(imageData)
// 	if err != nil {
// 		return err
// 	}

// 	//add purpose field
// 	if err := multiwriter.WriteField("purpose", "image"); err != nil {
// 		log.Fatal(err)
// 	}

// 	//get content type
// 	contentType := multiwriter.FormDataContentType()

// 	//close multiwriter
// 	if err := multiwriter.Close(); err != nil {
// 		log.Fatal(err)
// 	}

// 	//STEP 2 - UPLOAD REQUEST

// 	//ghost authorization
// 	ghostKey := config.GhostKey
// 	ghostToken, err := CreateAdminToken(ghostKey)
// 	if err != nil {
// 		return fmt.Errorf("authorization token error: %s", err)
// 	}

// 	//create http request
// 	endpoint := config.GhostURL + "/ghost/api/admin/images/upload/"
// 	req, err := http.NewRequest("POST", endpoint, &buffer)
// 	if err != nil {
// 		return fmt.Errorf("error creating http request: %s", err)
// 	}

// 	//set headers
// 	req.Header.Set("Content-Type", contentType)
// 	req.Header.Set("Authorization", "Ghost "+ghostToken)

// 	reqDump, err := httputil.DumpRequestOut(req, true)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("REQUEST:\n%s", string(reqDump))

// 	//send request
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return fmt.Errorf("failed to send request: %s", err)
// 	}
// 	defer resp.Body.Close()

// 	response, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println(string(response))

// 	return nil
// }

// func UploadImgToGhost(formData bytes.Buffer, contentType string, config internal.ApiConfig) error {
// 	//ghost authorization
// 	ghostKey := config.GhostKey
// 	ghostToken, err := CreateAdminToken(ghostKey)
// 	if err != nil {
// 		return fmt.Errorf("authorization token error: %s", err)
// 	}

// 	//create http request
// 	endpoint := config.GhostURL + "/ghost/api/admin/images/upload/"
// 	req, err := http.NewRequest("POST", endpoint, &formData)
// 	if err != nil {
// 		return fmt.Errorf("error creating http request: %s", err)
// 	}

// 	// //set headers
// 	req.Header.Set("Content-Type", contentType)
// 	req.Header.Set("Authorization", "Ghost "+ghostToken)

// 	// //send request
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return fmt.Errorf("failed to send request: %s", err)
// 	}
// 	defer resp.Body.Close()

// 	response, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println(string(response))

// 	return nil
// }

func createMultipartFormData(path string) (*multipart.Writer, io.Reader) {

	const paramName = "file"

	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	fileContents, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	fi, err := file.Stat()
	if err != nil {
		log.Println(err)
	}
	if err := file.Close(); err != nil {
		log.Println(err)
	}
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Type", "image/jpeg")
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
		escapeQuotes(paramName), escapeQuotes(fi.Name())))
	part, err := writer.CreatePart(h)
	if err != nil {
		log.Println(err)
	}
	_, err = part.Write(fileContents)
	if err != nil {
		log.Println(err)
	}

	params := map[string]string{
		"purpose": "image",
		"ref":     path,
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

func UploadImage(config internal.ApiConfig, path string) (imageURL string, err error) {
	writer, body := createMultipartFormData(path)
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
