cmake_minimum_required(VERSION 3.21)
project(opencv_example)
set(CMAKE_CXX_STANDARD 14)
find_package(OpenCV REQUIRED)
include_directories(${OpenCV_INCLUDE_DIRS} opencv_contrib/modules/aruco/samples)
add_executable(opencv_example main.cpp)
target_link_libraries(opencv_example ${OpenCV_LIBS})