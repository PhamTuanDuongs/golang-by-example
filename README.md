# golang-by-example
## To start using gRPC C++client
### I. Install tools and library
1. Install cmake and gRPC via link: https://grpc.io/docs/languages/cpp/quickstart/
2. Change version Cmake in CMakeLists.txt same with version that is installed on your pc
### II. Build the C++client
#### 1.Change to the example’s directory:
     $ cd golang-by-example/C++client/
#### 2.Change to the example’s directory: 
     $ mkdir -p cmake/build
     $ pushd cmake/build
     $ cmake -DCMAKE_PREFIX_PATH=$MY_INSTALL_DIR ../..
     $ make -j 4
#### Note 
if run **cmake-DCMAKE_PREFIX_PATH=$MY_INSTALL_DIR ../..** </br>
It throw:
CMake Error at CMakeLists.txt:9 (find_package):
  Could not find a package configuration file provided by "Protobuf" with any
  of the following names:

    ProtobufConfig.cmake
    protobuf-config.cmake

  Add the installation prefix of "Protobuf" to CMAKE_PREFIX_PATH or set
  "Protobuf_DIR" to a directory containing one of the above files.  If
  "Protobuf" provides a separate development package or SDK, be sure it has
  been installed.</br>
#### You need to follow with steps below to fix
**1. Remove the existing grpc build directory:** </br>

      cd /path/to/grpc  
      rm -rf cmake/build
      
</br>**2. Create a new build directory and navigate into it:** </br>

     mkdir -p cmake/build
     cd cmake/build
     
</br>**3. Configure the build with Cmake:** </br>

     cmake ../..
   
</br>**4. Build gRPC** </br>

     make -j 4
     
</br>**5. Install gRPC** </br>

     sudo make install

#### 3. Run file build:
    ./main
#### 3. Try it!
1. Before run C++Client you must run Go Server first in folder **chatBidirectionalStream**  



