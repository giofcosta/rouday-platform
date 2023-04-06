FROM alpine:latest

RUN mkdir /app

COPY roudayApp /app

CMD [ "/app/roudayApp" ]