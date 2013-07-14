#Author: sheppard(ysf1026@gmail.com) 2013-07-14
#       Desc:

project(claw-cgp)

cmake_minimum_required(VERSION 2.8)

include_directories(
     ${PROJECT_SOURCE_DIR}
)

set(BIN_NAME "client")
message(STATUS "${BIN_NAME}... @${PROJECT_SOURCE_DIR}")

file(GLOB_RECURSE SRCS *.cpp *.h *.inl)

set(Sources
     ${DEP_SRCS}
	 ${CGP_SRCS}
     ${SRCS}
)

add_executable(${BIN_NAME} ${Sources})
