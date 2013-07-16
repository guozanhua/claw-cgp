// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: cmd.proto

#ifndef PROTOBUF_cmd_2eproto__INCLUDED
#define PROTOBUF_cmd_2eproto__INCLUDED

#include <string>

#include <google/protobuf/stubs/common.h>

#if GOOGLE_PROTOBUF_VERSION < 2005000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please update
#error your headers.
#endif
#if 2005000 < GOOGLE_PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/unknown_field_set.h>
// @@protoc_insertion_point(includes)

namespace claw {
namespace cgp {
namespace proto {

// Internal implementation detail -- do not call these.
void  protobuf_AddDesc_cmd_2eproto();
void protobuf_AssignDesc_cmd_2eproto();
void protobuf_ShutdownFile_cmd_2eproto();

class Cmd;

// ===================================================================

class Cmd : public ::google::protobuf::Message {
 public:
  Cmd();
  virtual ~Cmd();

  Cmd(const Cmd& from);

  inline Cmd& operator=(const Cmd& from) {
    CopyFrom(from);
    return *this;
  }

  inline const ::google::protobuf::UnknownFieldSet& unknown_fields() const {
    return _unknown_fields_;
  }

  inline ::google::protobuf::UnknownFieldSet* mutable_unknown_fields() {
    return &_unknown_fields_;
  }

  static const ::google::protobuf::Descriptor* descriptor();
  static const Cmd& default_instance();

  void Swap(Cmd* other);

  // implements Message ----------------------------------------------

  Cmd* New() const;
  void CopyFrom(const ::google::protobuf::Message& from);
  void MergeFrom(const ::google::protobuf::Message& from);
  void CopyFrom(const Cmd& from);
  void MergeFrom(const Cmd& from);
  void Clear();
  bool IsInitialized() const;

  int ByteSize() const;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input);
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const;
  ::google::protobuf::uint8* SerializeWithCachedSizesToArray(::google::protobuf::uint8* output) const;
  int GetCachedSize() const { return _cached_size_; }
  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const;
  public:

  ::google::protobuf::Metadata GetMetadata() const;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // optional int32 id = 1;
  inline bool has_id() const;
  inline void clear_id();
  static const int kIdFieldNumber = 1;
  inline ::google::protobuf::int32 id() const;
  inline void set_id(::google::protobuf::int32 value);

  // optional string content = 2;
  inline bool has_content() const;
  inline void clear_content();
  static const int kContentFieldNumber = 2;
  inline const ::std::string& content() const;
  inline void set_content(const ::std::string& value);
  inline void set_content(const char* value);
  inline void set_content(const char* value, size_t size);
  inline ::std::string* mutable_content();
  inline ::std::string* release_content();
  inline void set_allocated_content(::std::string* content);

  // @@protoc_insertion_point(class_scope:claw.cgp.proto.Cmd)
 private:
  inline void set_has_id();
  inline void clear_has_id();
  inline void set_has_content();
  inline void clear_has_content();

  ::google::protobuf::UnknownFieldSet _unknown_fields_;

  ::std::string* content_;
  ::google::protobuf::int32 id_;

  mutable int _cached_size_;
  ::google::protobuf::uint32 _has_bits_[(2 + 31) / 32];

  friend void  protobuf_AddDesc_cmd_2eproto();
  friend void protobuf_AssignDesc_cmd_2eproto();
  friend void protobuf_ShutdownFile_cmd_2eproto();

  void InitAsDefaultInstance();
  static Cmd* default_instance_;
};
// ===================================================================


// ===================================================================

// Cmd

// optional int32 id = 1;
inline bool Cmd::has_id() const {
  return (_has_bits_[0] & 0x00000001u) != 0;
}
inline void Cmd::set_has_id() {
  _has_bits_[0] |= 0x00000001u;
}
inline void Cmd::clear_has_id() {
  _has_bits_[0] &= ~0x00000001u;
}
inline void Cmd::clear_id() {
  id_ = 0;
  clear_has_id();
}
inline ::google::protobuf::int32 Cmd::id() const {
  return id_;
}
inline void Cmd::set_id(::google::protobuf::int32 value) {
  set_has_id();
  id_ = value;
}

// optional string content = 2;
inline bool Cmd::has_content() const {
  return (_has_bits_[0] & 0x00000002u) != 0;
}
inline void Cmd::set_has_content() {
  _has_bits_[0] |= 0x00000002u;
}
inline void Cmd::clear_has_content() {
  _has_bits_[0] &= ~0x00000002u;
}
inline void Cmd::clear_content() {
  if (content_ != &::google::protobuf::internal::kEmptyString) {
    content_->clear();
  }
  clear_has_content();
}
inline const ::std::string& Cmd::content() const {
  return *content_;
}
inline void Cmd::set_content(const ::std::string& value) {
  set_has_content();
  if (content_ == &::google::protobuf::internal::kEmptyString) {
    content_ = new ::std::string;
  }
  content_->assign(value);
}
inline void Cmd::set_content(const char* value) {
  set_has_content();
  if (content_ == &::google::protobuf::internal::kEmptyString) {
    content_ = new ::std::string;
  }
  content_->assign(value);
}
inline void Cmd::set_content(const char* value, size_t size) {
  set_has_content();
  if (content_ == &::google::protobuf::internal::kEmptyString) {
    content_ = new ::std::string;
  }
  content_->assign(reinterpret_cast<const char*>(value), size);
}
inline ::std::string* Cmd::mutable_content() {
  set_has_content();
  if (content_ == &::google::protobuf::internal::kEmptyString) {
    content_ = new ::std::string;
  }
  return content_;
}
inline ::std::string* Cmd::release_content() {
  clear_has_content();
  if (content_ == &::google::protobuf::internal::kEmptyString) {
    return NULL;
  } else {
    ::std::string* temp = content_;
    content_ = const_cast< ::std::string*>(&::google::protobuf::internal::kEmptyString);
    return temp;
  }
}
inline void Cmd::set_allocated_content(::std::string* content) {
  if (content_ != &::google::protobuf::internal::kEmptyString) {
    delete content_;
  }
  if (content) {
    set_has_content();
    content_ = content;
  } else {
    clear_has_content();
    content_ = const_cast< ::std::string*>(&::google::protobuf::internal::kEmptyString);
  }
}


// @@protoc_insertion_point(namespace_scope)

}  // namespace proto
}  // namespace cgp
}  // namespace claw

#ifndef SWIG
namespace google {
namespace protobuf {


}  // namespace google
}  // namespace protobuf
#endif  // SWIG

// @@protoc_insertion_point(global_scope)

#endif  // PROTOBUF_cmd_2eproto__INCLUDED
