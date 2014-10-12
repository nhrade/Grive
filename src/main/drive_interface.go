/*
*   drive_interface.go
*   author: K.A
*/

package main

import
(
    "code.google.com/p/google-api-go-client/drive/v2"
    "code.google.com/p/goauth2/oauth"
    "github.com/skratchdot/open-golang/open"
    "log"
    "os"
    "net/http"
    "fmt"
)

const
(
    CLIENT_ID = "726901253511-bneug4hg5tgh5pnsgmltuj1a5cqchuo4.apps.googleusercontent.com"
    CLIENT_SECRET = "bU0MjgUVXH4CtLZazOyuw5vp"
    AUTH_FNAME = "auth.txt"
)

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

/*
func writeToFile(fname string, con string) {
    byteCon := make([]byte, len(con))
    copy(byteCon[:], con)
    ioutil.WriteFile(fname, byteCon, 0644)
}


func getToken() string {
    if con, err := ioutil.ReadFile(TOK_FNAME); err == nil {
        return string(con)
    }
    return ""
}
*/

func allFiles(d *drive.Service) ([]*drive.File, error) {
  var fs []*drive.File
  tok := ""
  for {
    q := d.Files.List()
    if tok != "" {
      q = q.PageToken(tok)
    }
    r, err := q.Do()
    if err != nil {
      fmt.Printf("An error occurred: %v\n", err)
      return fs, err
    }
    fs = append(fs, r.Items...)
    tok = r.NextPageToken
    if tok == "" {
      break
    }
  }
  return fs, nil
}

func deleteFile(title string, svc *drive.Service){
    files, ferr := allFiles(svc)
    if ferr != nil {
        log.Fatalf("Error occured: %v\n", ferr)
    }
    for _, f := range files {
        if f.Title == title{
            if err := svc.Files.Delete(f.Id).Do(); err != nil {
                log.Fatalf("Error occuredL %v\n", err)
            }
            log.Printf("Deleted: ID=%v Title=%v", f.Id, f.Title)
            os.Exit(0)
        }
    }
}

func auth() *drive.Service {
    //I literally tore code in and out of auth several times
      var t *oauth.Transport


      authUrl := config.AuthCodeURL("state")
      open.Run(authUrl)
      t = &oauth.Transport{
          Config:    config,
          Transport: http.DefaultTransport,
      }
      var code string
      // Read the code
      log.Printf("Enter verification code: ")
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
