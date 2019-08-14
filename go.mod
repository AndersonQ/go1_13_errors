module github.com/AndersonQ/go1_13_errors

go 1.13

require (
	github.com/pkg/errors v0.8.1
	github.com/rs/zerolog v1.14.3
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80 // indirect
	golang.org/x/tools v0.0.0-20190808195139-e713427fea3f // indirect
)

replace github.com/pkg/errors => ../errors
