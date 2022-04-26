# bloXroute Challenge

This is for the bloXroute Challenge to set up a go server subscribed to an SQS queue and a client that from the CLI can issue commands to the server.

## How To Run

Requirements for this are go with modules support.  I am using GO 1.17 because that's what I had currently installed.

Clone the repository

### Server

From the root folder

``` bin/bash
$ cd server
$ go get
$ go run main.go
```


### Client

From the root folder

``` bin/bash
$ cd client
$ go get
$ go run main.go [GetItems, GetItem, AddItem, DeleteItem]
```

## Example Usage:

![image](https://user-images.githubusercontent.com/8403969/165197596-f9369724-0b39-483f-9656-f4abc1593550.png)
