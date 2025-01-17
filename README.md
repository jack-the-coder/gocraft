# GoCraft

A Minecraft-like game written in Go, just for fun!

![Screenshot](https://i.imgur.com/vrGRDg1.png)

## Features

- Basic terrain generation
- Add and remove blocks.
- Move and fly.
- Multiplayer support

## Dependencies

### For Go

- Go 1.10+

### For GLFW

- On macOS, you need Xcode or Command Line Tools for Xcode (`xcode-select --install`) for required headers and libraries.
- On Ubuntu/Debian-like Linux distributions, you need `libgl1-mesa-dev` and `xorg-dev` packages.
- On CentOS/Fedora-like Linux distributions, you need `libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel mesa-libGL-devel libXi-devel` packages.


## Install

`go get github.com/icexin/gocraft`

## Run

Suppose `$GOPATH/bin` is in your `PATH`, use command below to run. 

`cd $GOPATH/src/github.com/icexin/gocraft && gocraft`

## How to play

- WASD to move around.
- Tab to toggle flying mode.
- Space to jump.
- Left and right click to add/remove block.
- E and R to cycle through the blocks.

## Multiplayer

Multiplayer is supported now!

The server code is at https://github.com/icexin/gocraft-server .

You can use `gocraft -s gocraft.icexin.com` to connect the public server.

Since the player on public server is anonymous, be carefull for your work!

If any network error occurs, the game will end with a panic, may changed in the future.

Local cache is saved as `cache_$server.db`, you can use `gocraft -db xxx.db` to offline use.

## Roadmap

- [x] Persistent changed blocks
- [x] Multiplayer support
- [ ] Seedable terrain generation
- [ ] Ambient Occlusion support

## Implementation Details

Many implementations is inspired by https://github.com/fogleman/Craft, thanks for Fogleman's good work!

Multiplayer is implementated used a duplex RPC call: client can call server to update blocks or fetch chunks, server can also push changes to clients. 
