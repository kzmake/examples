# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: blog/v1/blog.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='blog/v1/blog.proto',
  package='basic.blog.v1',
  syntax='proto3',
  serialized_options=b'ZDgithub.com/kzmake/examples/basic-service-invocation/api/blog/v1;blog',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x12\x62log/v1/blog.proto\x12\rbasic.blog.v1\x1a\x1cgoogle/api/annotations.proto\"+\n\rCreateRequest\x12\x1a\n\x08username\x18\x01 \x01(\tR\x08username\"9\n\x0e\x43reateResponse\x12\'\n\x04\x62log\x18\x01 \x01(\x0b\x32\x13.basic.blog.v1.BlogR\x04\x62log\"\x18\n\x04\x42log\x12\x10\n\x03url\x18\x01 \x01(\tR\x03url2T\n\x0b\x42logService\x12\x45\n\x06\x43reate\x12\x1c.basic.blog.v1.CreateRequest\x1a\x1d.basic.blog.v1.CreateResponseBFZDgithub.com/kzmake/examples/basic-service-invocation/api/blog/v1;blogb\x06proto3'
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,])




_CREATEREQUEST = _descriptor.Descriptor(
  name='CreateRequest',
  full_name='basic.blog.v1.CreateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='username', full_name='basic.blog.v1.CreateRequest.username', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='username', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=67,
  serialized_end=110,
)


_CREATERESPONSE = _descriptor.Descriptor(
  name='CreateResponse',
  full_name='basic.blog.v1.CreateResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='blog', full_name='basic.blog.v1.CreateResponse.blog', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='blog', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=112,
  serialized_end=169,
)


_BLOG = _descriptor.Descriptor(
  name='Blog',
  full_name='basic.blog.v1.Blog',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='url', full_name='basic.blog.v1.Blog.url', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='url', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=171,
  serialized_end=195,
)

_CREATERESPONSE.fields_by_name['blog'].message_type = _BLOG
DESCRIPTOR.message_types_by_name['CreateRequest'] = _CREATEREQUEST
DESCRIPTOR.message_types_by_name['CreateResponse'] = _CREATERESPONSE
DESCRIPTOR.message_types_by_name['Blog'] = _BLOG
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CreateRequest = _reflection.GeneratedProtocolMessageType('CreateRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEREQUEST,
  '__module__' : 'blog.v1.blog_pb2'
  # @@protoc_insertion_point(class_scope:basic.blog.v1.CreateRequest)
  })
_sym_db.RegisterMessage(CreateRequest)

CreateResponse = _reflection.GeneratedProtocolMessageType('CreateResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATERESPONSE,
  '__module__' : 'blog.v1.blog_pb2'
  # @@protoc_insertion_point(class_scope:basic.blog.v1.CreateResponse)
  })
_sym_db.RegisterMessage(CreateResponse)

Blog = _reflection.GeneratedProtocolMessageType('Blog', (_message.Message,), {
  'DESCRIPTOR' : _BLOG,
  '__module__' : 'blog.v1.blog_pb2'
  # @@protoc_insertion_point(class_scope:basic.blog.v1.Blog)
  })
_sym_db.RegisterMessage(Blog)


DESCRIPTOR._options = None

_BLOGSERVICE = _descriptor.ServiceDescriptor(
  name='BlogService',
  full_name='basic.blog.v1.BlogService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=197,
  serialized_end=281,
  methods=[
  _descriptor.MethodDescriptor(
    name='Create',
    full_name='basic.blog.v1.BlogService.Create',
    index=0,
    containing_service=None,
    input_type=_CREATEREQUEST,
    output_type=_CREATERESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_BLOGSERVICE)

DESCRIPTOR.services_by_name['BlogService'] = _BLOGSERVICE

# @@protoc_insertion_point(module_scope)
