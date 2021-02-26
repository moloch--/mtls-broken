
var jspb = require('google-protobuf');
var messages = require('./helloworld_pb');
var services = require('./helloworld_grpc_pb');
var fs = require('fs');
var path = require('path');
var grpc = require('@grpc/grpc-js');


var client = new services.GreeterClient("localhost:50051", grpc.credentials.createSsl(
    Buffer.from("asdf"), // fs.readFileSync(path.join(process.cwd(), "..", "ca-cert.pem")), // CA
    fs.readFileSync(path.join(process.cwd(), "client-key.pem")), // Private Key
    fs.readFileSync(path.join(process.cwd(), "client-cert.pem")), // Cert
    {
        checkServerIdentity: () => undefined
    }
));

var request = new messages.HelloRequest();
request.setName("john");
client.sayHello(request, function(err, response) {
    if (err) {
        console.error(err);
    }
    console.log('Greeting:', response.getMessage());
});

