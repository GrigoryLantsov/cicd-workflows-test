FROM golang:1.20

COPY main .
RUN chmod +x ./main
CMD ["./main"]
