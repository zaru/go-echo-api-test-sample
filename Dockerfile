FROM golang:1.9.2-stretch

WORKDIR /go/src/github.com/zaru/go-echo-api-test-sample

# Golang package
RUN go get -u \
    github.com/golang/dep/cmd/dep \
    github.com/tockins/realize \
    github.com/pressly/goose/cmd/goose
COPY . .
RUN dep ensure

# bashrc
# current path が長くなるので、行末に改行をいれる
RUN echo 'PS1="${debian_chroot:+($debian_chroot)}\e[35m\u@\h:\w\e[0m\\n$ "' >> ~/.bashrc; \
    echo "alias ll='ls -la'" >> ~/.bashrc


EXPOSE 1324
