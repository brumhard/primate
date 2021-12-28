///
//  Generated code. Do not modify.
//  source: dashboard/v1/dashboard.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'dashboard.pb.dart' as $0;
import 'dashboard.pbjson.dart';

export 'dashboard.pb.dart';

abstract class DashboardServiceBase extends $pb.GeneratedService {
  $async.Future<$0.GetExampleResponse> getExample($pb.ServerContext ctx, $0.GetExampleRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'GetExample': return $0.GetExampleRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'GetExample': return this.getExample(ctx, request as $0.GetExampleRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => DashboardServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => DashboardServiceBase$messageJson;
}

