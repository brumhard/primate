///
//  Generated code. Do not modify.
//  source: dashboard/v1/dashboard.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use listPullRequestsRequestDescriptor instead')
const ListPullRequestsRequest$json = const {
  '1': 'ListPullRequestsRequest',
};

/// Descriptor for `ListPullRequestsRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List listPullRequestsRequestDescriptor = $convert.base64Decode('ChdMaXN0UHVsbFJlcXVlc3RzUmVxdWVzdA==');
@$core.Deprecated('Use streamPullRequestsRequestDescriptor instead')
const StreamPullRequestsRequest$json = const {
  '1': 'StreamPullRequestsRequest',
};

/// Descriptor for `StreamPullRequestsRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List streamPullRequestsRequestDescriptor = $convert.base64Decode('ChlTdHJlYW1QdWxsUmVxdWVzdHNSZXF1ZXN0');
@$core.Deprecated('Use listPullRequestsResponseDescriptor instead')
const ListPullRequestsResponse$json = const {
  '1': 'ListPullRequestsResponse',
  '2': const [
    const {'1': 'items', '3': 1, '4': 3, '5': 11, '6': '.dashboard.v1.Repository', '10': 'items'},
  ],
};

/// Descriptor for `ListPullRequestsResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List listPullRequestsResponseDescriptor = $convert.base64Decode('ChhMaXN0UHVsbFJlcXVlc3RzUmVzcG9uc2USLgoFaXRlbXMYASADKAsyGC5kYXNoYm9hcmQudjEuUmVwb3NpdG9yeVIFaXRlbXM=');
@$core.Deprecated('Use repositoryDescriptor instead')
const Repository$json = const {
  '1': 'Repository',
  '2': const [
    const {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'pullrequests', '3': 2, '4': 3, '5': 11, '6': '.dashboard.v1.PullRequest', '10': 'pullrequests'},
  ],
};

/// Descriptor for `Repository`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List repositoryDescriptor = $convert.base64Decode('CgpSZXBvc2l0b3J5EhIKBG5hbWUYASABKAlSBG5hbWUSPQoMcHVsbHJlcXVlc3RzGAIgAygLMhkuZGFzaGJvYXJkLnYxLlB1bGxSZXF1ZXN0UgxwdWxscmVxdWVzdHM=');
@$core.Deprecated('Use pullRequestDescriptor instead')
const PullRequest$json = const {
  '1': 'PullRequest',
  '2': const [
    const {'1': 'title', '3': 1, '4': 1, '5': 9, '10': 'title'},
    const {'1': 'url', '3': 2, '4': 1, '5': 9, '10': 'url'},
    const {'1': 'user', '3': 3, '4': 1, '5': 9, '10': 'user'},
  ],
};

/// Descriptor for `PullRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pullRequestDescriptor = $convert.base64Decode('CgtQdWxsUmVxdWVzdBIUCgV0aXRsZRgBIAEoCVIFdGl0bGUSEAoDdXJsGAIgASgJUgN1cmwSEgoEdXNlchgDIAEoCVIEdXNlcg==');
