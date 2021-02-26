Broken mTLS in gRPC
=====================

This repo demonstrates that Node version after 14.15.3 cannot connect to a Go-based gRPC server using mutual-TLS.

### Instructions

1. Compile server with `go build .` and run `./mtls-broken` from root of Git repo
2. In `client/` directory run `npm install`
3. Run client script `node ./helloworld-client.js` using different version of NodeJS


## Node v15.10.0
```
bash-5.1$ nvm use 15.10.0
Now using node v15.10.0 (npm v7.5.3)
bash-5.1$ node ./helloworld-client.js 
Error: 14 UNAVAILABLE: No connection established
    at Object.callErrorFromStatus (/Users/moloch/git/mtls-broken/client/node_modules/@grpc/grpc-js/build/src/call.js:31:26)
    at Object.onReceiveStatus (/Users/moloch/git/mtls-broken/client/node_modules/@grpc/grpc-js/build/src/client.js:176:52)
    at Object.onReceiveStatus (/Users/moloch/git/mtls-broken/client/node_modules/@grpc/grpc-js/build/src/client-interceptors.js:336:141)
    at Object.onReceiveStatus (/Users/moloch/git/mtls-broken/client/node_modules/@grpc/grpc-js/build/src/client-interceptors.js:299:181)
    at /Users/moloch/git/mtls-broken/client/node_modules/@grpc/grpc-js/build/src/call-stream.js:130:78
    at processTicksAndRejections (node:internal/process/task_queues:76:11) {
  code: 14,
  details: 'No connection established',
  metadata: Metadata { internalRepr: Map(0) {}, options: {} }
}
/Users/moloch/git/mtls-broken/client/helloworld-client.js:25
    console.log('Greeting:', response.getMessage());
                                      ^

TypeError: Cannot read property 'getMessage' of undefined
    at Object.callback (/Users/moloch/git/mtls-broken/client/helloworld-client.js:25:39)
    at Object.onReceiveStatus (/Users/moloch/git/mtls-broken/client/node_modules/@grpc/grpc-js/build/src/client.js:176:36)
    at Object.onReceiveStatus (/Users/moloch/git/mtls-broken/client/node_modules/@grpc/grpc-js/build/src/client-interceptors.js:336:141)
    at Object.onReceiveStatus (/Users/moloch/git/mtls-broken/client/node_modules/@grpc/grpc-js/build/src/client-interceptors.js:299:181)
    at /Users/moloch/git/mtls-broken/client/node_modules/@grpc/grpc-js/build/src/call-stream.js:130:78
    at processTicksAndRejections (node:internal/process/task_queues:76:11)
```

## Node v14.15.4

```
bash-5.1$ node ./helloworld-client.js 
Error: 14 UNAVAILABLE: No connection established
    at Object.callErrorFromStatus (/Users/moloch/git/mtls-broken/client/node_modules/@grpc/grpc-js/build/src/call.js:31:26)
    at Object.onReceiveStatus (/Users/moloch/git/mtls-broken/client/node_modules/@grpc/grpc-js/build/src/client.js:176:52)
    at Object.onReceiveStatus (/Users/moloch/git/mtls-broken/client/node_modules/@grpc/grpc-js/build/src/client-interceptors.js:336:141)
    at Object.onReceiveStatus (/Users/moloch/git/mtls-broken/client/node_modules/@grpc/grpc-js/build/src/client-interceptors.js:299:181)
    at /Users/moloch/git/mtls-broken/client/node_modules/@grpc/grpc-js/build/src/call-stream.js:130:78
    at processTicksAndRejections (internal/process/task_queues.js:75:11) {
  code: 14,
  details: 'No connection established',
  metadata: Metadata { internalRepr: Map(0) {}, options: {} }
}
/Users/moloch/git/mtls-broken/client/helloworld-client.js:25
    console.log('Greeting:', response.getMessage());
                                      ^

TypeError: Cannot read property 'getMessage' of undefined
    at Object.callback (/Users/moloch/git/mtls-broken/client/helloworld-client.js:25:39)
    at Object.onReceiveStatus (/Users/moloch/git/mtls-broken/client/node_modules/@grpc/grpc-js/build/src/client.js:176:36)
    at Object.onReceiveStatus (/Users/moloch/git/mtls-broken/client/node_modules/@grpc/grpc-js/build/src/client-interceptors.js:336:141)
    at Object.onReceiveStatus (/Users/moloch/git/mtls-broken/client/node_modules/@grpc/grpc-js/build/src/client-interceptors.js:299:181)
    at /Users/moloch/git/mtls-broken/client/node_modules/@grpc/grpc-js/build/src/call-stream.js:130:78
    at processTicksAndRejections (internal/process/task_queues.js:75:11)
```


## Node v14.15.3
```
bash-5.1$ nvm use 14.15.3
Now using node v14.15.3 (npm v6.14.9)
bash-5.1$ node ./helloworld-client.js 
Greeting: Hello john
```