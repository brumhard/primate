///
//  Generated code. Do not modify.
//  source: dashboard/v1/dashboard.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class GetExampleRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetExampleRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'awesomecli.v1'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  GetExampleRequest._() : super();
  factory GetExampleRequest() => create();
  factory GetExampleRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetExampleRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetExampleRequest clone() => GetExampleRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetExampleRequest copyWith(void Function(GetExampleRequest) updates) => super.copyWith((message) => updates(message as GetExampleRequest)) as GetExampleRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetExampleRequest create() => GetExampleRequest._();
  GetExampleRequest createEmptyInstance() => create();
  static $pb.PbList<GetExampleRequest> createRepeated() => $pb.PbList<GetExampleRequest>();
  @$core.pragma('dart2js:noInline')
  static GetExampleRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetExampleRequest>(create);
  static GetExampleRequest? _defaultInstance;
}

class GetExampleResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetExampleResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'awesomecli.v1'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  GetExampleResponse._() : super();
  factory GetExampleResponse() => create();
  factory GetExampleResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetExampleResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetExampleResponse clone() => GetExampleResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetExampleResponse copyWith(void Function(GetExampleResponse) updates) => super.copyWith((message) => updates(message as GetExampleResponse)) as GetExampleResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetExampleResponse create() => GetExampleResponse._();
  GetExampleResponse createEmptyInstance() => create();
  static $pb.PbList<GetExampleResponse> createRepeated() => $pb.PbList<GetExampleResponse>();
  @$core.pragma('dart2js:noInline')
  static GetExampleResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetExampleResponse>(create);
  static GetExampleResponse? _defaultInstance;
}

class DashboardServiceApi {
  $pb.RpcClient _client;
  DashboardServiceApi(this._client);

  $async.Future<GetExampleResponse> getExample($pb.ClientContext? ctx, GetExampleRequest request) {
    var emptyResponse = GetExampleResponse();
    return _client.invoke<GetExampleResponse>(ctx, 'DashboardService', 'GetExample', request, emptyResponse);
  }
}

