///
//  Generated code. Do not modify.
//  source: dashboard/v1/dashboard.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use getExampleRequestDescriptor instead')
const GetExampleRequest$json = const {
  '1': 'GetExampleRequest',
};

/// Descriptor for `GetExampleRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getExampleRequestDescriptor = $convert.base64Decode('ChFHZXRFeGFtcGxlUmVxdWVzdA==');
@$core.Deprecated('Use getExampleResponseDescriptor instead')
const GetExampleResponse$json = const {
  '1': 'GetExampleResponse',
};

/// Descriptor for `GetExampleResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getExampleResponseDescriptor = $convert.base64Decode('ChJHZXRFeGFtcGxlUmVzcG9uc2U=');
const $core.Map<$core.String, $core.dynamic> DashboardServiceBase$json = const {
  '1': 'DashboardService',
  '2': const [
    const {'1': 'GetExample', '2': '.awesomecli.v1.GetExampleRequest', '3': '.awesomecli.v1.GetExampleResponse', '4': const {}},
  ],
};

@$core.Deprecated('Use dashboardServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> DashboardServiceBase$messageJson = const {
  '.awesomecli.v1.GetExampleRequest': GetExampleRequest$json,
  '.awesomecli.v1.GetExampleResponse': GetExampleResponse$json,
};

/// Descriptor for `DashboardService`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List dashboardServiceDescriptor = $convert.base64Decode('ChBEYXNoYm9hcmRTZXJ2aWNlElMKCkdldEV4YW1wbGUSIC5hd2Vzb21lY2xpLnYxLkdldEV4YW1wbGVSZXF1ZXN0GiEuYXdlc29tZWNsaS52MS5HZXRFeGFtcGxlUmVzcG9uc2UiAA==');
