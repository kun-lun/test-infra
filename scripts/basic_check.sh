set -x
go get -u golang.org/x/lint/golint
go get -u github.com/onsi/ginkgo/ginkgo
go get -u github.com/onsi/gomega/...
export PATH="$PATH:$GOPATH/bin"
golint `go list ./... | grep -v ./vendor/`
ginkgo -r