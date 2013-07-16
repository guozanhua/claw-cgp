#!/usr/bin/ruby
#Author: sheppard(ysf1026@gmail.com) 2013-07-12
#   Desc: 

require 'rbconfig'

def win_check_pause yes
  if yes
    system "pause"
  end
end

def os
  case RbConfig::CONFIG['host_os']
    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/i
      :windows
    when /linux/i
      :linux
    when /solaris|bsd/i
      :unix
    when /darwin|mac os/
      :mac
    else
      :unknow
    end
end
