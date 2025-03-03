# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/service/signal.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from flyteidl.admin import signal_pb2 as flyteidl_dot_admin_dot_signal__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1d\x66lyteidl/service/signal.proto\x12\x10\x66lyteidl.service\x1a\x1cgoogle/api/annotations.proto\x1a\x1b\x66lyteidl/admin/signal.proto2\x9a\x03\n\rSignalService\x12W\n\x11GetOrCreateSignal\x12(.flyteidl.admin.SignalGetOrCreateRequest\x1a\x16.flyteidl.admin.Signal\"\x00\x12\xc1\x01\n\x0bListSignals\x12!.flyteidl.admin.SignalListRequest\x1a\x1a.flyteidl.admin.SignalList\"s\x82\xd3\xe4\x93\x02m\x12k/api/v1/signals/{workflow_execution_id.project}/{workflow_execution_id.domain}/{workflow_execution_id.name}\x12l\n\tSetSignal\x12 .flyteidl.admin.SignalSetRequest\x1a!.flyteidl.admin.SignalSetResponse\"\x1a\x82\xd3\xe4\x93\x02\x14:\x01*\"\x0f/api/v1/signalsB\xc3\x01\n\x14\x63om.flyteidl.serviceB\x0bSignalProtoP\x01Z=github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/service\xa2\x02\x03\x46SX\xaa\x02\x10\x46lyteidl.Service\xca\x02\x10\x46lyteidl\\Service\xe2\x02\x1c\x46lyteidl\\Service\\GPBMetadata\xea\x02\x11\x46lyteidl::Serviceb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'flyteidl.service.signal_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\024com.flyteidl.serviceB\013SignalProtoP\001Z=github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/service\242\002\003FSX\252\002\020Flyteidl.Service\312\002\020Flyteidl\\Service\342\002\034Flyteidl\\Service\\GPBMetadata\352\002\021Flyteidl::Service'
  _SIGNALSERVICE.methods_by_name['ListSignals']._options = None
  _SIGNALSERVICE.methods_by_name['ListSignals']._serialized_options = b'\202\323\344\223\002m\022k/api/v1/signals/{workflow_execution_id.project}/{workflow_execution_id.domain}/{workflow_execution_id.name}'
  _SIGNALSERVICE.methods_by_name['SetSignal']._options = None
  _SIGNALSERVICE.methods_by_name['SetSignal']._serialized_options = b'\202\323\344\223\002\024:\001*\"\017/api/v1/signals'
  _globals['_SIGNALSERVICE']._serialized_start=111
  _globals['_SIGNALSERVICE']._serialized_end=521
# @@protoc_insertion_point(module_scope)
