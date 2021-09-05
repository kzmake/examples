# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: greeter/v1/greeter.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='greeter/v1/greeter.proto',
  package='basic.greeter.v1',
  syntax='proto3',
  serialized_options=b'ZMgithub.com/kzmake/examples/basic-service-invocation-py/api/greeter/v1;greeter',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x18greeter/v1/greeter.proto\x12\x10\x62\x61sic.greeter.v1\x1a\x1cgoogle/api/annotations.proto\"\"\n\x0cHelloRequest\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\"!\n\rHelloResponse\x12\x10\n\x03msg\x18\x01 \x01(\tR\x03msg2q\n\x07Greeter\x12\x66\n\x05Hello\x12\x1e.basic.greeter.v1.HelloRequest\x1a\x1f.basic.greeter.v1.HelloResponse\"\x1c\x82\xd3\xe4\x93\x02\x16\"\x11/greeter/v1/hello:\x01*BOZMgithub.com/kzmake/examples/basic-service-invocation-py/api/greeter/v1;greeterb\x06proto3'
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,])




_HELLOREQUEST = _descriptor.Descriptor(
  name='HelloRequest',
  full_name='basic.greeter.v1.HelloRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='basic.greeter.v1.HelloRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='name', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=76,
  serialized_end=110,
)


_HELLORESPONSE = _descriptor.Descriptor(
  name='HelloResponse',
  full_name='basic.greeter.v1.HelloResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='msg', full_name='basic.greeter.v1.HelloResponse.msg', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='msg', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_end=145,
)

DESCRIPTOR.message_types_by_name['HelloRequest'] = _HELLOREQUEST
DESCRIPTOR.message_types_by_name['HelloResponse'] = _HELLORESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

HelloRequest = _reflection.GeneratedProtocolMessageType('HelloRequest', (_message.Message,), {
  'DESCRIPTOR' : _HELLOREQUEST,
  '__module__' : 'greeter.v1.greeter_pb2'
  # @@protoc_insertion_point(class_scope:basic.greeter.v1.HelloRequest)
  })
_sym_db.RegisterMessage(HelloRequest)

HelloResponse = _reflection.GeneratedProtocolMessageType('HelloResponse', (_message.Message,), {
  'DESCRIPTOR' : _HELLORESPONSE,
  '__module__' : 'greeter.v1.greeter_pb2'
  # @@protoc_insertion_point(class_scope:basic.greeter.v1.HelloResponse)
  })
_sym_db.RegisterMessage(HelloResponse)


DESCRIPTOR._options = None

_GREETER = _descriptor.ServiceDescriptor(
  name='Greeter',
  full_name='basic.greeter.v1.Greeter',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=147,
  serialized_end=260,
  methods=[
  _descriptor.MethodDescriptor(
    name='Hello',
    full_name='basic.greeter.v1.Greeter.Hello',
    index=0,
    containing_service=None,
    input_type=_HELLOREQUEST,
    output_type=_HELLORESPONSE,
    serialized_options=b'\202\323\344\223\002\026\"\021/greeter/v1/hello:\001*',
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_GREETER)

DESCRIPTOR.services_by_name['Greeter'] = _GREETER

# @@protoc_insertion_point(module_scope)
