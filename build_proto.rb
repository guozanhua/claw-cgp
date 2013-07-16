#!/usr/bin/ruby
#Author: sheppard(ysf1026@gmail.com) 2013-07-15
#   Desc: 

$LOAD_PATH.unshift File.join(File.dirname(__FILE__), "./tool/dep")
require "base"

def_path = "./protocol_definition"
out_path = "./src/claw/cgp/common/proto"
protoc_path = ""
if os == :windows
  protoc_path = "./tool/win/"
end

puts `#{protoc_path}protoc -I=#{def_path} --cpp_out=#{out_path} #{def_path}/*.proto`
puts "def_path is: #{def_path}"
puts "ok, out_path is: #{out_path}"

if os == :windows
  system "pause"
end
