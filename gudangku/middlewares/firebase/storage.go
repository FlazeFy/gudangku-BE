package firebase

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

type Firebase struct {
	StorageClient *storage.Client
}

var firebaseInstance *Firebase

func initFirebase() (*Firebase, error) {
	if firebaseInstance != nil {
		return firebaseInstance, nil
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("secret/gudangku-94edc-firebase-adminsdk-we9nr-31d47a729d.json"))
	if err != nil {
		return nil, fmt.Errorf("error initializing Google Cloud Storage client: %w", err)
	}

	firebaseInstance = &Firebase{StorageClient: client}
	return firebaseInstance, nil
}

func UploadFile(ctx, user_id, username string, file *multipart.FileHeader, fileExt string) (string, error) {
	firebase, err := initFirebase()
	if err != nil {
		return "", fmt.Errorf("failed to initialize Firebase: %w", err)
	}
	bucket := firebase.StorageClient.Bucket("gudangku-94edc.appspot.com")
	fileReader, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer fileReader.Close()

	id := uuid.New().String()
	objectName := fmt.Sprintf("%s/%s/%s", ctx, user_id+"_"+username, id+"."+fileExt)

	writer := bucket.Object(objectName).NewWriter(context.Background())
	writer.ContentType = MimeType(fileExt)
	writer.ACL = []storage.ACLRule{
		{Entity: storage.AllUsers, Role: storage.RoleReader},
	}

	if _, err := io.Copy(writer, fileReader); err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}
	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %w", err)
	}

	attrs, err := bucket.Object(objectName).Attrs(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get object attributes: %w", err)
	}

	return attrs.MediaLink, nil
}
