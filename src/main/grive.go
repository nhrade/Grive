/*
*    grypt.go
*    author: K.A
*/

package main

import (
    "code.google.com/p/google-api-go-client/drive/v2"
    "os"
    "log"
)




func basicUsage(usage string){
    log.Fatalf("Wrong arguments, usage: grive %v", usage)
}

func cliArgs(svc *drive.Service){
    if len(os.Args) < 3{
        log.Fatal("Too few arguments, usage: grive [command] [args]")
    }
    cmd := os.Args[1]
    //Handle cli arguments
    switch cmd {
        case "upload":
            if len(os.Args) < 4{
                basicUsage("upload [title] [filename]")
            }
            file := File{os.Args[2], os.Args[3]}
            uploadFile(&file, svc)
        case "delete":
            if len(os.Args) != 3{
                basicUsage("delete [title]")
            }
            deleteFile(os.Args[2], svc)
        case "rename":
            if(len(os.Args) != 4){
                basicUsage("rename [targetTitle] [srcTitle]")
            }
            renameFile(os.Args[2], os.Args[3], svc)
        default:
            log.Fatal("Err: Unknown command")
    }
}

func main() {
    svc := auth()
    cliArgs(svc)

/*



  if err != nil {
    log.Fatalf("An error occurred uploading the document: %v\n", err)
  }
  log.Printf("Created: ID=%v, Title=%v\n", r.Id, r.Title)
  */
}
