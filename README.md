# <p style="text-align: center;">Go Realtime Chat Server</p>

## Content
- [Quick start](#quick-start)
- [Architecture](#architecture)
- [Features](#features)
- [References](#references)
## Quick start

***// TODO***

## Architecture
### *1. This project is build in Go, conforming **Clean Architecture***

![Clean Architecture](./assets/clean_architecture.png)

### *2. Chat feature using Websocket and **Hub Architecture***
    - Using go routines and channels to handle rooms and messages
    - Room will be saved in main memory (deleted after reseting server)

![Hub Architecture](./assets/hub_architecture.jpg)


## Features: 
    - Log in, Log out, Sign up
    - Get All Rooms, Get Room By ID
    - Create Room, Join Room, Send Message To Room



## References: 
[1] [Go Next TS Chat Project](https://github.com/dhij/go-next-ts_chat)