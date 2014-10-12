/*
*   drive_interface.go
*   author: K.A
*/

package main

import
(
    "code.google.com/p/google-api-go-client/drive/v2"
    "code.google.com/p/goauth2/oauth"
    "log"
    "os"
    "net/http"
    "fmt"
)

const CLIENT_ID = "726901253511-bneug4hg5tgh5pnsgmltuj1a5cqchuo4.apps.googleusercontent.com"
const CLIENT_SECRET = "bU0MjgUVXH4CtLZazOyuw5vp"

type File struct {
    title, filename string
}

// Settings for authorization.
var config = &oauth.Config{
  ClientId:     CLIENT_ID,
  ClientSecret: CLIENT_SECRET,
  Scope:        "https://www.googleapis.com/auth/drive",
  RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
  AuthURL:      "https://accounts.google.com/o/oauth2/auth",
  TokenURL:     "https://accounts.google.com/o/oauth2/token",
}



func uploadFile(file *File, svc *drive.Service){

    f := &drive.File{
         Title: file.title,
         Description: "",
    }
    m, err := os.Open(file.filename)
    if err != nil {
        log.Fatalf("Err: Could not open file. %v\n", err)
    }
    r, err := svc.Files.Insert(f).Media(m).Do()
    if err != nil{
        log.Fatalf("Err: Could not upload file %v\n", err)
    }
    log.Printf("Created: ID=%v Title=%v\n", r.Id, r.Title)
}


func auth() *drive.Service {

      // Generate a URL to visit for authorization.

      authUrl := config.AuthCodeURL("state")
      log.Printf("Go to the following link in your browser: %v\n", authUrl)

      t := &oauth.Transport{
        Config:    config,
        Transport: http.DefaultTransport,
      }

      // Read the code, and exchange it for a token.
      log.Printf("Enter verification code: ")
      var code string
      fmt.Scanln(&code)
      _, err := t.Exchange(code)
      if err != nil {
        log.Fatalf("An error occurred exchanging the code: %v\n", err)
      }

      // Create a new authorized Drive client.
      svc, err := drive.New(t.Client())
      if err != nil {
        log.Fatalf("An error occurred creating Drive client: %v\n", err)
      }
      return svc
}

func getFileTitle(d *drive.Service, id string) (title string, e error) {
    f, err := d.Files.Get(id).Do()
    if err != nil {
        fmt.Printf("Boom goes the dynamite. %v\n", err)
        return "", err
    }
    return f.Title, nil
}
