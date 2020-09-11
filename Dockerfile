FROM scratch

EXPOSE 9001

# Copy our static executable.
COPY example /go/bin/example

ENTRYPOINT ["/go/bin/example"]