A look at Go 1.13 Errors
--------------------------------------------------

A talk presented on [August Golang meetup](https://www.meetup.com/golang-users-berlin/events/259188830/)
showing where we are in the proposal to enhance error handling for Go 2 and what was added already
in go 1.13. I also wrote a [blog post](https://medium.com/onefootball-locker-room/a-look-at-go-1-13-errors-9f6c9f6accb6) about it

The slides use go [present](https://godoc.org/golang.org/x/tools/present). To render the presentation:
1. have go present `golang.org/x/tools/present` (`go get golang.org/x/tools/cmd/present`)
2. go into the slides folder (`cd slides`)
3. `present` or `$GOPATH/bin/present`
4. Open your web browser on [http://127.0.0.1:3999](http://127.0.0.1:3999)
5. click on [slides.slide](http://127.0.0.1:3999/slides.slide)

Also there is a sample project demonstrating how to use the new error functionalities added in
go 1.13. To run it:

### Requirement
 - go 1.13
 
### Run

`go run main.go`