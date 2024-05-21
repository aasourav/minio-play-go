package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	endpoint := "172.19.255.201:9000"
	accessKeyID := "puNPwHUvuNbzap8JWW7u"
	secretAccessKey := "llSVWiTyBK3Cv7p4R6QJvbYDZ44HrW78237iZ5bN"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// log.Printf("%#v\n", minioClient) // minioClient is now setup
	// err = minioClient.MakeBucket(context.Background(), "new-bucket", minio.MakeBucketOptions{Region: "api-south"})

	// if err != nil {
	// 	log.Fatalln(err)
	// } else {
	// 	log.Printf("success")
	// }

	// err = minioClient.FGetObject(context.Background(), "aas", "bleachbit_4.6.0-0_all_ubuntu2204.deb", "downloads.deb", minio.GetObjectOptions{})

	file, err := os.Open("./cookies.txt")
	defer file.Close()

	fileStat, err := file.Stat()
	minioClient.PutObject(context.Background(), "aas", "test-obj.txt", file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})

	// reqParams := make(url.Values)
	// reqParams.Set("response-content-disposition", "attachment; filename=\""+"20240521_08.sql"+"\"")

	// for k, v := range reqParams {
	// 	fmt.Printf("key[%s] value[%s]\n", k, v)
	// }
	// url, _ := minioClient.PresignedGetObject(context.Background(), "mysql", "20240521_08.sql", time.Hour*24*7, reqParams)

	objectCh := minioClient.ListObjects(context.Background(), "mysql", minio.ListObjectsOptions{
		Recursive: true,
	})

	for object := range objectCh {
		if object.Err != nil {
			log.Fatalln(object.Err)
		}

		// Generate a presigned URL for each object
		presignedURL, err := minioClient.PresignedGetObject(context.Background(), "mysql", object.Key, time.Hour*24*7, nil)
		if err != nil {
			log.Fatalln(err)
		}

		// Print the presigned URL
		fmt.Printf("Presigned URL for object %s: %s\n", object.Key, presignedURL)
	}

}
