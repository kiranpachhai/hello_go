FROM scratch
ADD bin/hello_go /hello_go
CMD ["/hello_go"]
