#!/usr/bin/ruby
#Author: sheppard(ysf1026@gmail.com) 2013-07-15
#   Desc: 

def_path = "./protocol_definitions"
out_path = "./src/common/proto"
puts `protoc -I=#{def_path} --cpp_out=#{out_path} #{def_path}/*.proto`
puts "def_path is: #{def_path}"
puts "ok, out_path is: #{out_path}"
