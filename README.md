# <p style="text-align: center;">Go Realtime Chat Server</p>

## Content
- [Architecture](#architecture)
- [Features](#features)
- [References](#references)

## Architecture
### *1. This project is build in Go, conforming **Clean Architecture***

![Clean Architecture](./assets/clean_architecture.png)

### *2. Chat feature using Websocket and **Hub Architecture***
    - Using go routines and channels to handle rooms and messages
    - Room will be saved in main memory (deleted after reseting server)

![Hub Architecture](./assets/hub_architecture.jpg)


## Features: 
    - Log in, Log out, Sign up
    - Create Room, Join Room, Send Message To Room

