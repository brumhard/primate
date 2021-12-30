///
//  Generated code. Do not modify.
//  source: dashboard/v1/dashboard.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dashboard.pbenum.dart';

export 'dashboard.pbenum.dart';

class ListPullRequestsRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ListPullRequestsRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'dashboard.v1'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  ListPullRequestsRequest._() : super();
  factory ListPullRequestsRequest() => create();
  factory ListPullRequestsRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ListPullRequestsRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ListPullRequestsRequest clone() => ListPullRequestsRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ListPullRequestsRequest copyWith(void Function(ListPullRequestsRequest) updates) => super.copyWith((message) => updates(message as ListPullRequestsRequest)) as ListPullRequestsRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ListPullRequestsRequest create() => ListPullRequestsRequest._();
  ListPullRequestsRequest createEmptyInstance() => create();
  static $pb.PbList<ListPullRequestsRequest> createRepeated() => $pb.PbList<ListPullRequestsRequest>();
  @$core.pragma('dart2js:noInline')
  static ListPullRequestsRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ListPullRequestsRequest>(create);
  static ListPullRequestsRequest? _defaultInstance;
}

class StreamPullRequestsRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'StreamPullRequestsRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'dashboard.v1'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  StreamPullRequestsRequest._() : super();
  factory StreamPullRequestsRequest() => create();
  factory StreamPullRequestsRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory StreamPullRequestsRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  StreamPullRequestsRequest clone() => StreamPullRequestsRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  StreamPullRequestsRequest copyWith(void Function(StreamPullRequestsRequest) updates) => super.copyWith((message) => updates(message as StreamPullRequestsRequest)) as StreamPullRequestsRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static StreamPullRequestsRequest create() => StreamPullRequestsRequest._();
  StreamPullRequestsRequest createEmptyInstance() => create();
  static $pb.PbList<StreamPullRequestsRequest> createRepeated() => $pb.PbList<StreamPullRequestsRequest>();
  @$core.pragma('dart2js:noInline')
  static StreamPullRequestsRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<StreamPullRequestsRequest>(create);
  static StreamPullRequestsRequest? _defaultInstance;
}

class ListPullRequestsResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ListPullRequestsResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'dashboard.v1'), createEmptyInstance: create)
    ..pc<Repository>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'items', $pb.PbFieldType.PM, subBuilder: Repository.create)
    ..hasRequiredFields = false
  ;

  ListPullRequestsResponse._() : super();
  factory ListPullRequestsResponse({
    $core.Iterable<Repository>? items,
  }) {
    final _result = create();
    if (items != null) {
      _result.items.addAll(items);
    }
    return _result;
  }
  factory ListPullRequestsResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ListPullRequestsResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ListPullRequestsResponse clone() => ListPullRequestsResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ListPullRequestsResponse copyWith(void Function(ListPullRequestsResponse) updates) => super.copyWith((message) => updates(message as ListPullRequestsResponse)) as ListPullRequestsResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ListPullRequestsResponse create() => ListPullRequestsResponse._();
  ListPullRequestsResponse createEmptyInstance() => create();
  static $pb.PbList<ListPullRequestsResponse> createRepeated() => $pb.PbList<ListPullRequestsResponse>();
  @$core.pragma('dart2js:noInline')
  static ListPullRequestsResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ListPullRequestsResponse>(create);
  static ListPullRequestsResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Repository> get items => $_getList(0);
}

class Repository extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Repository', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'dashboard.v1'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..pc<PullRequest>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pullrequests', $pb.PbFieldType.PM, subBuilder: PullRequest.create)
    ..hasRequiredFields = false
  ;

  Repository._() : super();
  factory Repository({
    $core.String? name,
    $core.Iterable<PullRequest>? pullrequests,
  }) {
    final _result = create();
    if (name != null) {
      _result.name = name;
    }
    if (pullrequests != null) {
      _result.pullrequests.addAll(pullrequests);
    }
    return _result;
  }
  factory Repository.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Repository.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Repository clone() => Repository()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Repository copyWith(void Function(Repository) updates) => super.copyWith((message) => updates(message as Repository)) as Repository; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Repository create() => Repository._();
  Repository createEmptyInstance() => create();
  static $pb.PbList<Repository> createRepeated() => $pb.PbList<Repository>();
  @$core.pragma('dart2js:noInline')
  static Repository getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Repository>(create);
  static Repository? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<PullRequest> get pullrequests => $_getList(1);
}

class PullRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'PullRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'dashboard.v1'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'title')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'url')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'user')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'sourceBranch')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'targetBranch')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'lastUpdatedAt')
    ..e<PullRequest_Status>(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'status', $pb.PbFieldType.OE, defaultOrMaker: PullRequest_Status.STATUS_UNSPECIFIED, valueOf: PullRequest_Status.valueOf, enumValues: PullRequest_Status.values)
    ..hasRequiredFields = false
  ;

  PullRequest._() : super();
  factory PullRequest({
    $core.String? title,
    $core.String? url,
    $core.String? user,
    $core.String? sourceBranch,
    $core.String? targetBranch,
    $core.String? lastUpdatedAt,
    PullRequest_Status? status,
  }) {
    final _result = create();
    if (title != null) {
      _result.title = title;
    }
    if (url != null) {
      _result.url = url;
    }
    if (user != null) {
      _result.user = user;
    }
    if (sourceBranch != null) {
      _result.sourceBranch = sourceBranch;
    }
    if (targetBranch != null) {
      _result.targetBranch = targetBranch;
    }
    if (lastUpdatedAt != null) {
      _result.lastUpdatedAt = lastUpdatedAt;
    }
    if (status != null) {
      _result.status = status;
    }
    return _result;
  }
  factory PullRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PullRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PullRequest clone() => PullRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PullRequest copyWith(void Function(PullRequest) updates) => super.copyWith((message) => updates(message as PullRequest)) as PullRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static PullRequest create() => PullRequest._();
  PullRequest createEmptyInstance() => create();
  static $pb.PbList<PullRequest> createRepeated() => $pb.PbList<PullRequest>();
  @$core.pragma('dart2js:noInline')
  static PullRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PullRequest>(create);
  static PullRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get title => $_getSZ(0);
  @$pb.TagNumber(1)
  set title($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTitle() => $_has(0);
  @$pb.TagNumber(1)
  void clearTitle() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get url => $_getSZ(1);
  @$pb.TagNumber(2)
  set url($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasUrl() => $_has(1);
  @$pb.TagNumber(2)
  void clearUrl() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get user => $_getSZ(2);
  @$pb.TagNumber(3)
  set user($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasUser() => $_has(2);
  @$pb.TagNumber(3)
  void clearUser() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get sourceBranch => $_getSZ(3);
  @$pb.TagNumber(4)
  set sourceBranch($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasSourceBranch() => $_has(3);
  @$pb.TagNumber(4)
  void clearSourceBranch() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get targetBranch => $_getSZ(4);
  @$pb.TagNumber(5)
  set targetBranch($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasTargetBranch() => $_has(4);
  @$pb.TagNumber(5)
  void clearTargetBranch() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get lastUpdatedAt => $_getSZ(5);
  @$pb.TagNumber(6)
  set lastUpdatedAt($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasLastUpdatedAt() => $_has(5);
  @$pb.TagNumber(6)
  void clearLastUpdatedAt() => clearField(6);

  @$pb.TagNumber(7)
  PullRequest_Status get status => $_getN(6);
  @$pb.TagNumber(7)
  set status(PullRequest_Status v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasStatus() => $_has(6);
  @$pb.TagNumber(7)
  void clearStatus() => clearField(7);
}

