# Minimum CMake required
cmake_minimum_required(VERSION  3.19.6)

# Project
project(c++client)

# Protobuf
set(protobuf_MODULE_COMPATIBLE TRUE)
find_package(Protobuf CONFIG REQUIRED)
message(STATUS "Using protobuf ${protobuf_VERSION}")

# Protobuf-compiler
set(_PROTOBUF_PROTOC $<TARGET_FILE:protobuf::protoc>)

# gRPC
find_package(gRPC CONFIG REQUIRED)
message(STATUS "Using gRPC ${gRPC_VERSION}")
set(_GRPC_GRPCPP gRPC::grpc++)
set(_GRPC_CPP_PLUGIN_EXECUTABLE $<TARGET_FILE:gRPC::grpc_cpp_plugin>)

# Proto file
get_filename_component(hw_proto "chatproto.proto" ABSOLUTE)
get_filename_component(hw_proto_path "${hw_proto}" PATH)

set(CODEGEN_DIR "${CMAKE_CURRENT_SOURCE_DIR}/proto")
# Generated sources
set(hw_proto_srcs "${CODEGEN_DIR}/chatproto.pb.cc")
set(hw_proto_hdrs "${CODEGEN_DIR}/chatproto.pb.h")
set(hw_grpc_srcs "${CODEGEN_DIR}/chatproto.grpc.pb.cc")
set(hw_grpc_hdrs "${CODEGEN_DIR}/chatproto.grpc.pb.h")
add_custom_command(
      OUTPUT "${hw_proto_srcs}" "${hw_proto_hdrs}" "${hw_grpc_srcs}" "${hw_grpc_hdrs}"
      COMMAND ${_PROTOBUF_PROTOC}
      ARGS --grpc_out "${CODEGEN_DIR}"
        --cpp_out "${CODEGEN_DIR}"
        -I "${hw_proto_path}"
        --plugin=protoc-gen-grpc="${_GRPC_CPP_PLUGIN_EXECUTABLE}"
        "${hw_proto}"
      DEPENDS "${hw_proto}")

# Include generated *.pb.h files
include_directories("${CODEGEN_DIR}")

# Targets (client|server)
foreach(_target 
    main)
  add_executable(${_target} "${_target}.cc"
    ${hw_proto_srcs}
    ${hw_grpc_srcs})
  target_link_libraries(${_target}
    ${_REFLECTION}
    ${_GRPC_GRPCPP}
    ${_PROTOBUF_LIBPROTOBUF})
endforeach()