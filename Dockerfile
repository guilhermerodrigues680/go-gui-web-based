FROM scratch

EXPOSE 9001

COPY go-gui-web-based /go/bin/go-gui-web-based

ENTRYPOINT ["/go/bin/go-gui-web-based"]
