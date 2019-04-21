# golang-autoclicker
Auto Clicker in Golang to Simulate Mouse Clicks

## Robotgo Package

Used [robotgo](https://github.com/go-vgo/robotgo) package for this

## Overview

I wanted a tool to simulate clicks for a game called [Plarium Vikings]() to auto click and apply tasks within the game.

## 1. Getting the Mouse Coordinates

Using `getpointer.go`: it will sleep for 2 seconds, then return the coordinates of your mouse's position. You will use those coordinates to plot where the mouse need to move.

Example:

```go
$ go run getpointer.go
pos: 966 621
pos: 849 356
```

## 2. Main App: Move and Click

I have 2 positions that I want the mouse to move to and click, then with a combination of a while loop, I just reiterate between clicking on two positions

After you start your mouse will move to the plotted locations and click as soon as it gets to the coordinate.

```go
$ go run autoclick.go
Starting iteration
Moving Mouse
Clicking
Delaying with 2 Seconds
Moving Mouse
Clicking
Delaying with 2 Seconds
finising iteration
...
finising iteration
finished all iterations
```
