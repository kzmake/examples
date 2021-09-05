# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: wallet/v1/wallet.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='wallet/v1/wallet.proto',
  package='basic.wallet.v1',
  syntax='proto3',
  serialized_options=b'ZHgithub.com/kzmake/examples/basic-service-invocation/api/wallet/v1;wallet',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x16wallet/v1/wallet.proto\x12\x0f\x62\x61sic.wallet.v1\x1a\x1cgoogle/api/annotations.proto\"\x0f\n\rCreateRequest\"A\n\x0e\x43reateResponse\x12/\n\x06wallet\x18\x01 \x01(\x0b\x32\x17.basic.wallet.v1.WalletR\x06wallet\"p\n\x06Wallet\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x37\n\x07\x62\x61lance\x18\x02 \x01(\x0b\x32\x1d.basic.wallet.v1.Wallet.MoneyR\x07\x62\x61lance\x1a\x1d\n\x05Money\x12\x14\n\x05value\x18\x01 \x01(\x03R\x05value2Z\n\rWalletService\x12I\n\x06\x43reate\x12\x1e.basic.wallet.v1.CreateRequest\x1a\x1f.basic.wallet.v1.CreateResponseBJZHgithub.com/kzmake/examples/basic-service-invocation/api/wallet/v1;walletb\x06proto3'
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,])




_CREATEREQUEST = _descriptor.Descriptor(
  name='CreateRequest',
  full_name='basic.wallet.v1.CreateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
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
  serialized_start=73,
  serialized_end=88,
)


_CREATERESPONSE = _descriptor.Descriptor(
  name='CreateResponse',
  full_name='basic.wallet.v1.CreateResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='wallet', full_name='basic.wallet.v1.CreateResponse.wallet', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='wallet', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=90,
  serialized_end=155,
)


_WALLET_MONEY = _descriptor.Descriptor(
  name='Money',
  full_name='basic.wallet.v1.Wallet.Money',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='basic.wallet.v1.Wallet.Money.value', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='value', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=240,
  serialized_end=269,
)

_WALLET = _descriptor.Descriptor(
  name='Wallet',
  full_name='basic.wallet.v1.Wallet',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='basic.wallet.v1.Wallet.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='id', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='balance', full_name='basic.wallet.v1.Wallet.balance', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='balance', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_WALLET_MONEY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=157,
  serialized_end=269,
)

_CREATERESPONSE.fields_by_name['wallet'].message_type = _WALLET
_WALLET_MONEY.containing_type = _WALLET
_WALLET.fields_by_name['balance'].message_type = _WALLET_MONEY
DESCRIPTOR.message_types_by_name['CreateRequest'] = _CREATEREQUEST
DESCRIPTOR.message_types_by_name['CreateResponse'] = _CREATERESPONSE
DESCRIPTOR.message_types_by_name['Wallet'] = _WALLET
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CreateRequest = _reflection.GeneratedProtocolMessageType('CreateRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEREQUEST,
  '__module__' : 'wallet.v1.wallet_pb2'
  # @@protoc_insertion_point(class_scope:basic.wallet.v1.CreateRequest)
  })
_sym_db.RegisterMessage(CreateRequest)

CreateResponse = _reflection.GeneratedProtocolMessageType('CreateResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATERESPONSE,
  '__module__' : 'wallet.v1.wallet_pb2'
  # @@protoc_insertion_point(class_scope:basic.wallet.v1.CreateResponse)
  })
_sym_db.RegisterMessage(CreateResponse)

Wallet = _reflection.GeneratedProtocolMessageType('Wallet', (_message.Message,), {

  'Money' : _reflection.GeneratedProtocolMessageType('Money', (_message.Message,), {
    'DESCRIPTOR' : _WALLET_MONEY,
    '__module__' : 'wallet.v1.wallet_pb2'
    # @@protoc_insertion_point(class_scope:basic.wallet.v1.Wallet.Money)
    })
  ,
  'DESCRIPTOR' : _WALLET,
  '__module__' : 'wallet.v1.wallet_pb2'
  # @@protoc_insertion_point(class_scope:basic.wallet.v1.Wallet)
  })
_sym_db.RegisterMessage(Wallet)
_sym_db.RegisterMessage(Wallet.Money)


DESCRIPTOR._options = None

_WALLETSERVICE = _descriptor.ServiceDescriptor(
  name='WalletService',
  full_name='basic.wallet.v1.WalletService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=271,
  serialized_end=361,
  methods=[
  _descriptor.MethodDescriptor(
    name='Create',
    full_name='basic.wallet.v1.WalletService.Create',
    index=0,
    containing_service=None,
    input_type=_CREATEREQUEST,
    output_type=_CREATERESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_WALLETSERVICE)

DESCRIPTOR.services_by_name['WalletService'] = _WALLETSERVICE

# @@protoc_insertion_point(module_scope)
