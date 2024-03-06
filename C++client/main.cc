#include <iostream>
#include <memory>
#include <string>
#include <grpcpp/grpcpp.h>
#include "chatproto.grpc.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::ClientReaderWriter;
using grpc::Status;
using Chat::Request;
using Chat::Response;
using Chat::ServiceChat;

class ChatCClient {
public:
    explicit ChatCClient(std::shared_ptr<Channel> channel) : stub_(ServiceChat::NewStub(channel)) {}

    void ChatC(const std::string& message, const std::string& tag) {
        // Set up the client context and stream for bidirectional communication.
        ClientContext context;
        std::unique_ptr<ClientReaderWriter<Request, Response>> stream(stub_->ChatC(&context));

        if (!stream) {
            std::cerr << "Failed to create stream." << std::endl;
            return;
        }

        // Send a message to the server.
        Request request;
        request.set_mess(message);
        request.set_tag(tag);
        if (!stream->Write(request)) {
            std::cerr << "Failed to write to stream." << std::endl;
            return;
        }

        // Read responses from the server.
        Response response;
        while (stream->Read(&response)) {
            std::cout << "Received from server: " << response.mess() << std::endl;
        }

        // Close the stream.
        Status status = stream->Finish();
        if (status.ok()) {
            std::cout << "ChatC completed successfully." << std::endl;
        } else {
            std::cerr << "ChatC RPC failed: " << status.error_code() << ": " << status.error_message() << std::endl;
        }
    }

private:
    std::unique_ptr<ServiceChat::Stub> stub_;
};

int main() {
    // Set up a gRPC channel to the server.
    std::string server_address("localhost:9999");
    ChatCClient client(grpc::CreateChannel(server_address, grpc::InsecureChannelCredentials()));

    // Continuous chat loop.
    while (true) {
        // Get user input for message and tag.
        std::cout << "Enter message: ";
        std::string message;
        std::getline(std::cin, message);

        std::cout << "Enter tag: ";
        std::string tag;
        std::getline(std::cin, tag);

        // ChatC with the server using user input.
        client.ChatC(message, tag);

        // Ask the user if they want to continue chatting.
        std::cout << "Do you want to continue chatting? (y/n): ";
        char response;
        std::cin >> response;
        if (response != 'y') {
            break; // Exit the loop if the user does not want to continue.
        }

        // Clear the newline character from the input buffer.
        std::cin.ignore(std::numeric_limits<std::streamsize>::max(), '\n');
    }

    return 0;
}
