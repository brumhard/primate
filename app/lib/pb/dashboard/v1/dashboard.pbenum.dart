///
//  Generated code. Do not modify.
//  source: dashboard/v1/dashboard.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class PullRequest_Status extends $pb.ProtobufEnum {
  static const PullRequest_Status STATUS_UNSPECIFIED = PullRequest_Status._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_UNSPECIFIED');
  static const PullRequest_Status STATUS_DRAFT = PullRequest_Status._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_DRAFT');
  static const PullRequest_Status STATUS_ACTIVE = PullRequest_Status._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_ACTIVE');
  static const PullRequest_Status STATUS_CLOSED = PullRequest_Status._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_CLOSED');

  static const $core.List<PullRequest_Status> values = <PullRequest_Status> [
    STATUS_UNSPECIFIED,
    STATUS_DRAFT,
    STATUS_ACTIVE,
    STATUS_CLOSED,
  ];

  static final $core.Map<$core.int, PullRequest_Status> _byValue = $pb.ProtobufEnum.initByValue(values);
  static PullRequest_Status? valueOf($core.int value) => _byValue[value];

  const PullRequest_Status._($core.int v, $core.String n) : super(v, n);
}

