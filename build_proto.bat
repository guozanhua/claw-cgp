title build_proto
tools\win\protoc -I"protocol_definitions" --cpp_out="common/proto" protocol_definitions\*.proto
pause