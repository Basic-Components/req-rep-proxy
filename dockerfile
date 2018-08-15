FROM hsz1273327/req-rep-proxy:Base-v0
ADD . /app/src/github.com/Basic-Components/req-rep-proxy
ENV GOPATH="/app"
WORKDIR /app/src/github.com/Basic-Components/req-rep-proxy
RUN go build