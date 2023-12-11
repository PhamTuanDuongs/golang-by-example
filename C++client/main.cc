#include <iostream>
#include <memory>
#include <grpcpp/grpcpp.h>
#include <grpcpp/impl/codegen/async_stream.h>
#include <thread>
#include "proto/chatproto.grpc.pb.h"

using Chat::Request;
using Chat::Response;
using Chat::ServiceChat;
using grpc::Channel;
using grpc::ClientContext;
using grpc::ClientReader;
using grpc::ClientReaderWriter;
using grpc::Status;

class ChatCClient
{
public:
    ChatCClient(std::shared_ptr<Channel> channel) : stub_(ServiceChat::NewStub(channel)) {}
    void ChatC()
    {
        // Container for the data we expect from the server.
        auto reply = Response();
        ClientContext context;

        std::unique_ptr<ClientReaderWriter<Request, Response>> stream(stub_->ChatC(&context));

        while (stream->Read(&reply))
        {
            std::cout << reply.mess().c_str() << std::endl;
        }

        Status status = stream->Finish();
        // Act upon its status.
        if (status.ok())
        {
            std::cout << "Read information from Server C successfully" << std::endl;
        }
        else
        {
            std::cout << status.error_code() << ": " << status.error_message() << std::endl;
        }
    }

private:
    std::unique_ptr<ServiceChat::Stub> stub_;
};

int main()
{
    // auto client = ChatCClient(grpc::CreateChannel(server_address, grpc::InsecureChannelCredentials()));
    ChatCClient chatclient(grpc::CreateChannel("localhost:9999", grpc::InsecureChannelCredentials()));
    // // Call the ChatC method
    chatclient.ChatC();
    std::getchar();
    return 0;
}
