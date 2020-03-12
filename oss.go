package ossservice

import (
	"bytes"
	"fmt"
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	ossClient *oss.Client
	ossBucket *oss.Bucket
)

//ossService - oss service
type ossService struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	Bucket          string
	Folder          string
}

//connectToOss - establishes a new connection to the oss with the Endpoint,AccessKeyID and the SecretAccessKey provided
func (o *ossService) connectToOss() *oss.Client {

	client, err := oss.New(o.Endpoint, o.AccessKeyID, o.SecretAccessKey)
	if err != nil {
		fmt.Println("error connecting to oss")
	}
	return client

}

//GetOssBucket establishes a connection with oss bucket  with the provided bucket name
func (o *ossService) GetOssBucket() *oss.Bucket {
	ossClient = o.connectToOss()
	bucket, err := ossClient.Bucket(o.Bucket)
	if err != nil {
		fmt.Println("error connecting to oss bucket")
	}
	return bucket
}

//DownloadObject downloads the object from the oss to the path specified and return the path
func (o *ossService) DownloadObject(object string, path string, size int64) (string, error) {
	ossBucket = o.GetOssBucket()
	err := ossBucket.DownloadFile(object, path, size)
	if err != nil {
		return "", err
	}
	return path, err
}

//GetByteObject get the object from the oss as object buffer (bytes) and return the path
func (o *ossService) GetByteObject(object string) (*bytes.Buffer, error) {
	ossBucket = o.GetOssBucket()
	buf := new(bytes.Buffer)
	body, err := ossBucket.GetObject(object)
	if err != nil {
		return buf, err
	}
	io.Copy(buf, body)
	body.Close()
	return buf, err
}
