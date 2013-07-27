#Author: sheppard(ysf1026@gmail.com) 2013-07-14
#       Desc: project root

project(claw-cgp)

cmake_minimum_required(VERSION 2.8)

message(STATUS "Project: claw-cgp")
message(STATUS "Project Directory: ${CMAKE_CURRENT_SOURCE_DIR}")

set(CMAKE_BUILD_TYPE DEBUG)
if(UNIX)
    add_definitions(-D_LINUX_)
elseif(MSVC)
    add_definitions(-D_WINDOWS_)
endif()
if(UNIX)
    set(CMAKE_CXX_FLAGS_DEBUG "-g -Wall")
else()
    set(CMAKE_CXX_FLAGS_DEBUG "-DEBUG -W3 -MDd -EHa -ZI -Od")
endif()

set(EXECUTABLE_OUTPUT_PATH ${CMAKE_CURRENT_SOURCE_DIR}/bin)

include_directories(
    ${CMAKE_CURRENT_SOURCE_DIR}/src
)

file(GLOB_RECURSE CGP_SRCS src/claw/cgp/*.h src/claw/cgp/*.cpp) #for test
list(REMOVE_ITEM CGP_SRCS "${CMAKE_CURRENT_SOURCE_DIR}/src/claw/cgp/client/client.cpp")
list(REMOVE_ITEM CGP_SRCS "${CMAKE_CURRENT_SOURCE_DIR}/src/claw/cgp/hall/hall.cpp")
list(REMOVE_ITEM CGP_SRCS "${CMAKE_CURRENT_SOURCE_DIR}/src/claw/cgp/login/login.cpp")

file(GLOB_RECURSE COMMON_SRCS src/claw/cgp/common/*.h src/claw/cgp/common/*.cpp src/claw/cgp/common/*.cc)

add_subdirectory(src)
add_subdirectory(test)
