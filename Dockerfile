FROM golang:latest

ENV HOME /go/src

WORKDIR ${HOME}

COPY . ${HOME}

CMD go version\
&& go env -w GOPROXY=https://goproxy.cn\
&& go build -o out -buildmode=plugin ${HOME}/plugins/simple/*.go\
&& go build -o out -buildmode=plugin ${HOME}/plugins/hello-world/*.go\
&& echo "build sucessful!\n"\
&& go run i.go
