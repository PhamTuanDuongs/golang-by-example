#include <iostream>
#include <memory>
#include <grpcpp/grpcpp.h>
#include <grpcpp/impl/codegen/async_stream.h>
#include <thread>
#include "proto/chatproto.grpc.pb.h"

using Chat::ChatService;
using Chat::Request;
using Chat::Response;
using grpc::Channel;
using grpc::ClientContext;
using grpc::ClientReaderWriter;
using grpc::Status;

class ChatCClient
{
public:
    ChatCClient(std::shared_ptr<Channel> channel) : stub_(ChatService::NewStub(channel)) {}

     void ChatC()
{
    ClientContext context;
    std::unique_ptr<ClientReaderWriter<Request, Response>> stream(stub_->ChatC(&context));

    // Check if the connection is successful
    if (!stream) {
        std::cerr << "Error connecting to the server" << std::endl;
        return;
    }

    // Start a separate thread to read responses asynchronously
    std::thread response_reader([this, &stream]() {
        Response response;
        while (stream->Read(&response)) {
            std::cout << "Received response: " << response.DebugString() << std::endl;
        }
    });

    // Wait for the response reader to finish
    response_reader.join();
}



private:
    std::unique_ptr<ChatService::Stub> stub_;
};

int main()
{
    // Create a channel to connect to the server
    std::string server_address("localhost:9999"); // Replace with your server address
    ChatCClient client(grpc::CreateChannel(server_address, grpc::InsecureChannelCredentials()));

    // Call the ChatC method
    client.ChatC();

    return 0;
}
