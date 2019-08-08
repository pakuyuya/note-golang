package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/service/s3"
    "fmt"
)

// @see https://rooter.jp/aws/aws-sdk-go-s3-four-tips/
// @see https://hori-ryota.com/blog/aws-s3-commonprefixes/
func main() {
	fmt.Println("try connection")

	regionName := "ap-northeast-1"
	profileName := "default"
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(regionName),
		Credentials: credentials.NewSharedCredentials("", profileName),
	}))
	fmt.Println("new s3 object")
	svc := s3.New(sess)

	bucketName := "your-bucket"
	prefix := ""

	fmt.Println("get list objects")
	//prefixを指定して検索
	resp, err := svc.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(prefix),
		Delimiter: aws.String("/"),
	})
	if err != nil {
		panic(err)
	}

	// ファイル一覧
	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}

	// フォルダ（prefix）一覧
	for _, item := range resp.CommonPrefixes {
		fmt.Println("Name:         ", *item.Prefix)
	}
}