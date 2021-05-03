# dns-trainings
DNS Trainings

## Usage

### udpserver

```sh
$ go run udpserver.go
```

```sh
$ dig @127.0.0.1 example.com. a in -p 10053
```

### udpserver2

Add DNS "NAME" parsing to udpserver.go

```sh
$ go run udpserver2.go
```

```sh
$ dig @127.0.0.1 0.1.2.3.4.5.6.7.8.9.a.b.c.d.e.local. a in -p 10053
2021/05/03 15:52:25 Now listen
2021/05/03 15:52:28 p=  1, label length= 1, name=0.
2021/05/03 15:52:28 p=  3, label length= 1, name=0.1.
2021/05/03 15:52:28 p=  5, label length= 1, name=0.1.2.
2021/05/03 15:52:28 p=  7, label length= 1, name=0.1.2.3.
2021/05/03 15:52:28 p=  9, label length= 1, name=0.1.2.3.4.
2021/05/03 15:52:28 p= 11, label length= 1, name=0.1.2.3.4.5.
2021/05/03 15:52:28 p= 13, label length= 1, name=0.1.2.3.4.5.6.
2021/05/03 15:52:28 p= 15, label length= 1, name=0.1.2.3.4.5.6.7.
2021/05/03 15:52:28 p= 17, label length= 1, name=0.1.2.3.4.5.6.7.8.
2021/05/03 15:52:28 p= 19, label length= 1, name=0.1.2.3.4.5.6.7.8.9.
2021/05/03 15:52:28 p= 21, label length= 1, name=0.1.2.3.4.5.6.7.8.9.a.
2021/05/03 15:52:28 p= 23, label length= 1, name=0.1.2.3.4.5.6.7.8.9.a.b.
2021/05/03 15:52:28 p= 25, label length= 1, name=0.1.2.3.4.5.6.7.8.9.a.b.c.
2021/05/03 15:52:28 p= 27, label length= 1, name=0.1.2.3.4.5.6.7.8.9.a.b.c.d.
2021/05/03 15:52:28 p= 29, label length= 1, name=0.1.2.3.4.5.6.7.8.9.a.b.c.d.e.
2021/05/03 15:52:28 p= 31, label length= 5, name=0.1.2.3.4.5.6.7.8.9.a.b.c.d.e.local.
2021/05/03 15:52:28 From: 127.0.0.1:64191, Data: 5296012000010000000000 , 013001310132013301340135013601370138013901610162016301640165056c6f63616c00000100010000291000000000000000 , 0.1.2.3.4.5.6.7.8.9.a.b.c.d.e.local.
```


