# go-string-non-thread-safe
An example to show string is not thread safe.

### output
````shell
b
05
b
b
b
b
b
b
b
b
````


### assembly
````
go tool compile -S -S main.go 
...
        rel 5+4 t=17 TLS+0
        rel 49+4 t=8 time.Sleep+0
        rel 75+4 t=16 "".a+4
        rel 85+4 t=16 runtime.writeBarrier+-1
        rel 95+4 t=16 go.string."0"+0
        rel 102+4 t=16 "".a+0
        rel 111+4 t=16 "".a+0
        rel 118+4 t=16 go.string."0"+0
        rel 123+4 t=8 runtime.gcWriteBarrierCX+0
        rel 133+4 t=16 "".a+4
        rel 143+4 t=16 runtime.writeBarrier+-1
        rel 153+4 t=16 go.string."bb"+0
        rel 160+4 t=16 "".a+0
        rel 169+4 t=16 "".a+0
        rel 176+4 t=16 go.string."bb"+0
        rel 181+4 t=8 runtime.gcWriteBarrierCX+0
        rel 193+4 t=8 runtime.morestack_noctxt+0
...
````