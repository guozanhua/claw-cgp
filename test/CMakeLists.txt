#Author: sheppard(ysf1026@gmail.com) 2013-07-14
#		Desc: tests

project(claw-cgp)

cmake_minimum_required(VERSION 2.8)

include_directories(
     ${PROJECT_SOURCE_DIR}
)

if(MSVC)
include_directories(
     ${CMAKE_CURRENT_SOURCE_DIR}/dep/include
)
link_directories(
     ${CMAKE_CURRENT_SOURCE_DIR}/dep/lib
)
endif()

set(BIN_NAME "test")
message(STATUS "${BIN_NAME}... @${PROJECT_SOURCE_DIR}")

file(GLOB_RECURSE SRCS *.cpp *.h *.inl)

set(Sources
     ${DEP_SRCS}
	 ${CGP_SRCS}
     ${SRCS}
)

add_executable(${BIN_NAME} ${Sources})
if(UNIX)
    target_link_libraries(${BIN_NAME} gtest)
elseif(MSVC)
    target_link_libraries(${BIN_NAME} gtestd)
endif()
