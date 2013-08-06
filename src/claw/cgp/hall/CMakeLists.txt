#Author: sheppard(ysf1026@gmail.com) 2013-07-15
#       Desc:

project(claw-cgp)

cmake_minimum_required(VERSION 2.8)

include_directories(
    ${PROJECT_SOURCE_DIR}
)

set(BIN_NAME "hall")
message(STATUS "${BIN_NAME}... @${PROJECT_SOURCE_DIR}")

file(GLOB_RECURSE SRCS *.cpp *.h *.inl)

set(Sources
    ${COMMON_SRCS}
    ${SRCS}
)

add_executable(${BIN_NAME} ${Sources})
target_link_libraries(${BIN_NAME} pthread zmq protobuf glog claw-gse jsoncpp)
