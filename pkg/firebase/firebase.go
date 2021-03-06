package firebase

import (
	"context"
	"path/filepath"

	firebaseAuth "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/labstack/gommon/log"
	"google.golang.org/api/option"
)

func SetupFirebase() *auth.Client {

	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	//Firebase admin SDK initialization
	app, err := firebaseAuth.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}

	//Firebase Auth
	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Error("Firebase Auth Error: ", err)
		panic("Firebase load error")
	}

	return auth
}
