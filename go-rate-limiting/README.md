# GOLANG X REDIS ... RATE LIMITING

## Things to Learn

- decorator pattern in golang.

## What is **RATE LIMITING**?

Rate limiting is a way to limit the number of request a user can send in a given time.

## Possible types of **RATE LIMITING**?

- IP Based
- Token Based

### Let's try out the IP Based Rate Limiting

> Before that ...

I want to test the server against different IP addresses. So I have to build a client that can send the request to the server with different IP address. So let's see what I have to do to build something like that.

One idea is **IP Scopping**.
