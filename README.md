Grive
=====

A command line interface for google drive written in go

Made @ HackRU


To build go must be installed along with the dependencies using the go get tool
  - github.com/skratchdot/open-golang
  - code.google.com/p/google-api-go-client/drive/v2
  - code.google.com/p/goauth2/oauth
  
then run
  - `go build`

To use
make grive executable and then `./grive [command] [args]`

To upload a file
`./grive upload [title] [filename]`

To delete a file
`./grive delete [title]`

To rename file
`./grive rename [target] [source]`
