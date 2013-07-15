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
set(LIBRARY_OUTPUT_PATH ${CMAKE_CURRENT_SOURCE_DIR}/dep/lib)

if(UNIX)
include_directories(
    ${CMAKE_CURRENT_SOURCE_DIR}/src #add this path for tests
    /usr/local/include
)

link_directories(
    ${CMAKE_CURRENT_SOURCE_DIR}/lib
    /usr/local/lib
)
elseif(MSVC)
include_directories(
     ${CMAKE_CURRENT_SOURCE_DIR}/src #add this path for tests
     ${CMAKE_CURRENT_SOURCE_DIR}/dep/include
     ${CMAKE_CURRENT_SOURCE_DIR}/dep/win/include
)

link_directories(
     ${CMAKE_CURRENT_SOURCE_DIR}/dep/win/lib/debug
)
endif()

if(MSVC)
     file(GLOB_RECURSE DEP_SRCS dep/win/include/*.h dep/win/lib/boost/*.cpp)
endif()
file(GLOB_RECURSE CGP_SRCS src/claw/gse/*.h src/claw/gse/*.cpp)

add_subdirectory(test)
