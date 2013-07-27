#!/bin/bash
#Author: sheppard(ysf1026@gmail.com) 2013-07-27
#    Desc: 

#cmake files redirect to bin/cmake_build
cmake -H. -Bbin/cmake_build
cd bin/cmake_build
make
