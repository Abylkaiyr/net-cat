# NET-CAT

Project creates a simple TCP chat

The max size of the room is 10 clients. 

Project uses GOROUTINES, Channels! 

USAGE:

1. Clone it
2. Move to the root directory of the project


To run the server:

```bash
make server || go run cmd/server/main.go
```

To add a new client to the server click to the ADD NEW SPLIT TERMINAL


```bash
make client || go run cmd/client/main.go || nc localhost 8989
```




### Example
- `1st` Terminal
```bash
$ go run cmd/server/main.go || make server
Listening on the port :8989
2022/08/30 19:16:23 Listening the connections on the server 127.0.0.1:8989
$ 
```
- `2nd` Terminal
```bash
$ go run cmd/client/main.go || nc localhost 8989 || make client
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]: Abylkaiyr
[2021-10-27 19:14:10][Abylkaiyr]:Hello 
Dias has joined our chat...
[2021-10-27 19:14:43][Dias]:Hello
[2021-10-27 19:14:43][Abylkaiyr]:How Are you?
[2021-10-27 19:15:31][Dias]:I am fine, and you?
[2021-10-27 19:15:51][Abylkaiyr]:Good
Dias has left our chat...
[2021-10-27 19:16:10][Abylkaiyr]:^C
$ 
```
- `3rd` Terminal
```bash
$go run cmd/client/main.go || nc localhost 8989 || make client
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]: Dias
[2021-10-27 19:14:20][Abylkaiyr]:Hello
[2021-10-27 19:14:32][Dias]:Hello
[2021-10-27 19:15:04][Abylkaiyr]:How Are you?
[2021-10-27 19:15:04][Dias]:I am fine, and you?
[2021-10-27 19:15:31][Abylkaiyr]:Good
[2021-10-27 19:16:03][Dias]:^C
$ 
```




# Made by:

      Abylkaiyr & Dias12