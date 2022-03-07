FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN cd /build && git clone https://github.com/Tahmid79/gin-demo.git

RUN cd /build/gin-demo && go build -o app
RUN cd /build/gin-demo && ./app

EXPOSE 5000

ENTRYPOINT [ "/build/gin-demo/app" ]

