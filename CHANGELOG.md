# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.7.0, 2026-10-30

- Breaking: `CreateConversationAttributeListRequest.Options` is now `[]CreateConversationAttributeOptionRequest` instead of a pointer to the slice. The struct is no longer comparable.
- Breaking: `CreateConversationAttributeRelationshipRequest.Reference` is now a struct value, and its `ObjectTypeId` is a `string` instead of `*string`.
- Breaking: `Contacts` is now a pointer to the slice on `CreateTicketRequest`, `CreateTicketJSONBody`, `CreateTicketJSONRequestBody`, `EnqueueCreateTicketJSONBody` and `EnqueueCreateTicketJSONRequestBody`.
- Breaking: `CreateTicketTypeAttributeRequest.Description` is now `*string`.
- Breaking: `TicketStateDetailed.TicketTypes` data items are now an inline struct instead of `*TicketType`, with a typed `Category` of `TicketStateDetailedTicketTypesDataCategory`.
- Added `TicketStateDetailedTicketTypesDataCategory` with the constants `TicketStateDetailedTicketTypesDataCategoryBackOffice`, `TicketStateDetailedTicketTypesDataCategoryCustomer` and `TicketStateDetailedTicketTypesDataCategoryTracker`, plus a `Valid` method.
- Added `Ticket.CompanyId`, `TicketReply.AppPackageCode` and `TicketType.IsInternal`.
- Added typed error bodies and their `GetJSON` methods: `JSON400`, `JSON404` and `JSON500` on `CreateTicketResponse`, `JSON400` on `CreateTicketTypeAttributeResponse`, `CreateTicketTypeResponse` and `SearchTicketsResponse`, `JSON404` on `GetTicketResponse` and `GetTicketTypeResponse`, `JSON500` on `EnqueueCreateTicketResponse`, and `JSON422` on `DeleteTicketResponse`, `UpdateTicketResponse`, `DeleteConversationAttributeResponse` and `UpdateConversationAttributeResponse`.
- Documentation: the `AdminId` comment on the conversation attribute types now says it is `null` when there is no admin.

## v0.6.0, 2026-10-29

- Breaking: the Fin Voice call methods take a new params argument. `CollectFinVoiceCallById`, `CollectFinVoiceCallByExternalId`, `CollectFinVoiceCallByPhoneNumber` and `CollectFinVoiceCallsByConversationId` now take `*CollectFinVoiceCallByIdParams`, `*CollectFinVoiceCallByExternalIdParams`, `*CollectFinVoiceCallByPhoneNumberParams` and `*CollectFinVoiceCallsByConversationIdParams`. `RegisterFinVoiceCall` and `RegisterFinVoiceCallWithBody` take `*RegisterFinVoiceCallParams` before the body. This also changes the `ClientWithResponses` methods, `ClientInterface`, `ClientWithResponsesInterface` and the `NewCollectFinVoiceCall*Request` and `NewRegisterFinVoiceCallRequest*` builders.
- Breaking: `CursorPages.Next` is now `*string` instead of a struct with `Page` and `StartingAfter`. `CursorPages` also has a new `Prev` field.
- Breaking: `CreateOfficeHoursScheduleRequest.TimeZoneName`, `ListActivityLogsParams.CreatedAtAfter`, `SearchActivityLogsJSONBody.CreatedAtAfter` and `SearchActivityLogsJSONRequestBody.CreatedAtAfter` are now pointers.
- Breaking: the `JSON200` field and `GetJSON200` method on `ListCallsWithTranscriptsResponse` have a changed type. `ValidationError` is removed.
- Breaking: removed the constants `ActivityLogActivityTypeConversationDeletionScheduleDeletion`, `ActivityLogActivityTypeRulesetActivationTitlePreview`, `ActivityLogActivityTypeSearchBrowseEnabledChange` and `ActivityLogActivityTypeSearchBrowseRequiredChange`.
- Added about 270 `ActivityLogActivityType*` constants, for example `ActivityLogActivityTypeQaMonitorCreation`, `ActivityLogActivityTypeScimGroupCreation` and `ActivityLogActivityTypeRulesetActivation`.
- Added `ListActivityLogsParams.Page` and `PerPage`, `ActivityLog.Type`, `AdminWithApp.Role` and `App.Secure`.
- Added `Call.Duration`, `Call.HoldTime`, `Call.QueueTime` and `Call.TalkTime`, plus `AiCallResponse.FinCallDuration`, `Language` and `VendorInternalId`.
- Added `RegisterFinVoiceCallRequest.VendorInternalId` and the source constants `RegisterFinVoiceCallRequestSourceDialpad`, `RegisterFinVoiceCallRequestSourceGeneric` and `RegisterFinVoiceCallRequestSourceNiceCxone`.
- Added typed error bodies (`JSON400`, `JSON401`, `JSON404`, `JSON405`, `JSON422` and their `GetJSON` methods) on responses for office hours, `ListActivityLogs`, `ListCalls`, `IdentifyAdmin`, `SetAwayAdmin` and `ShowCallTranscript`.

## v0.5.1, 2026-10-29

- Documentation only. No exported identifiers changed.
- The doc comments on `BulkUpdateContacts` and `BulkDeleteContacts` (and their `WithBody` and `WithResponse` variants) now give the limit as 1,000 contacts per request instead of 100.
- Updated doc comments for the article, collection, content snippet, content import source, workflow export and Fin ask and escalate methods. They now describe the `404` and `422` response bodies, the `help_center_id` rules on collection updates, the IPv4 allowlist lockout `422`, the workflow ruleset ID versus `workflow.id`, and that `Title` on content snippet requests may be missing, null or empty.
- Updated field comments: `Name` on folders is returned HTML-escaped, `Page` above 5001 returns a `400`, `Query` should stay under 500 characters, and `TotalCount` on content import source lists includes deleted sources.
- Moved `ParsePreviewMergeContactResponse` and `ParseSearchContactsResponse` to a different generated file, and reordered some response types between files. Behavior is the same.

## v0.5.0, 2026-10-28

- Breaking: the note methods take a string ID instead of an int. This changes `(*Client).CreateNote`, `CreateNoteWithBody`, `DeleteNote` and `RetrieveNote`, the matching `ClientWithResponses` methods, `ClientInterface`, `ClientWithResponsesInterface`, and `NewCreateNoteRequest`, `NewCreateNoteRequestWithBody`, `NewDeleteNoteRequest` and `NewRetrieveNoteRequest`.
- Breaking: the reporting export and download endpoints no longer send `app_id`, `client_id` or `job_identifier` query parameters. `GetExportReportingDataJobIdentifierParams` loses `AppId`, `ClientId` and `JobIdentifier`, and `GetDownloadReportingDataJobIdentifierParams` loses `AppId` and `JobIdentifier`.
- Breaking: `SegmentList.Pages` is removed, `NoteList.Pages` is now `*PagesLink`, `ConvertVisitorRequest.User` is now a pointer, and `ConvertVisitorRequestVisitor2` and its union methods are removed.
- Breaking: `ListTagsResponse.JSON200` is now `*ListTags200JSONResponseBody`, and `LisDataEventsParams.Type` is now `LisDataEventsParamsType`.
- Breaking: `Metadata` on `DataEvent` and `CreateDataEventRequest`, and `CustomAttributes` on `Visitor`, `UpdateVisitorRequest` and `CustomObjectInstance`, are now maps of union types instead of `map[string]string`. `Error.Errors[].Code` is a string-or-integer union, read with `AsErrorErrorsCode0()`.
- Breaking: the inline structs on `Admin.Avatar`, `Visitor.Companies`, `Visitor.Segments`, `TagCompanyRequest.Companies`, `UntagCompanyRequest.Companies`, `TagMultipleUsersRequest.Users` and `CreateConversationRequest.From` changed shape, and `CreateConversationRequestFromTypeContact` is removed.
- Added `CreateOrUpdateTagRequest.IdsToAdd` and `TaggableType`, `CreateDataEventSummariesRequest` union methods, `ConvertVisitorRequestUser2`, and `JSON400`, `JSON401`, `JSON404`, `JSON422`, `JSON429` and `JSON500` fields on many response types.
- Fixed the spec patch tool: Intercom changed the `/tags` create body from `oneOf` to `anyOf`, which left `CreateTagJSONRequestBody` without JSON marshaling so it sent `{}`. The patch now hoists `anyOf` bodies too, and `CreateTagJSONRequestBody` is an alias of the new `ClientBodyCreateTag` union.

## v0.4.3, 2026-10-28

- Updated Semgrep in the development image to 1.182.1 and refreshed its locked dependencies. certifi 2026.10.23 drops the emSign root CAs whose intermediate keys were compromised (GHSA-6q48-6c7w-9gw2).
- Dependabot no longer opens pull requests for `mcp` or the OpenTelemetry packages in `requirements-semgrep.txt`. Semgrep pins their versions itself, so bumping one of them on its own cannot install.
- Updated kin-openapi, which the code generator uses, from 0.142.0 to 0.149.0. The newer version fixes several panics on malformed specs reported in GitHub security advisories.
- The security scan merges its two gitleaks runs into one SARIF run. GitHub code scanning rejects a report with two runs from the same tool, so no scan results had been uploaded.

## v0.4.2, 2026-10-27

- The development image is based on Debian (`golang:1.26.10-trixie`) instead of Alpine. Alpine keeps only the latest build of each package, so the exact pins broke the image build whenever a package was updated. System packages now come from the snapshot.debian.org date the base image was built from, which keeps them fixed for a given base image digest. jq 1.8.2, shellcheck 0.11.0 and shfmt 3.13.1 keep their versions, and Semgrep runs in its own virtual environment.

## v0.4.1, 2026-10-27

- Documentation only. The doc comment on `CreateDataConnectorRequest.Headers` and `UpdateDataConnectorRequest.Headers` now warns against putting credentials in header values, because they appear unmasked in execution logs except `Authorization`, and points to `token_ids` instead.

## v0.4.0, 2026-10-24

- Breaking: `DeleteCustomObjectInstancesByExternalId` and `DeleteCustomObjectInstancesById` change arguments on `Client`, `ClientWithResponses` and both interfaces. `DeleteCustomObjectInstancesByExternalId` drops one string argument, `DeleteCustomObjectInstancesById` gains one, and `NewDeleteCustomObjectInstancesByExternalIdRequest` and `NewDeleteCustomObjectInstancesByIdRequest` change the same way. `DeleteCustomObjectInstancesByIdParams.ExternalId` moves to `DeleteCustomObjectInstancesByExternalIdParams.ExternalId`.
- Breaking: `UpdateNewsItem` and `UpdateNewsItemWithResponse` now take `UpdateNewsItemRequest` as the body. `UpdateNewsItemJSONRequestBody` changes from `NewsItemRequest` to it. Added `UpdateNewsItemRequestState` with `Draft` and `Live` values.
- Breaking: `CustomObjectInstancesPaginatedList` loses `Data`, `Pages`, `TotalCount` and `Type`, and is no longer comparable. `ListCustomObjectInstances200JSONResponseBody` is a new union with `AsCustomObjectInstance`, `AsCustomObjectInstanceNotFound`, `AsCustomObjectInstancesPaginatedList` and matching `From`, `Merge` and JSON methods.
- Breaking: `CreateOrUpdateCustomObjectInstanceRequest.CustomAttributes` is now a map of `CreateOrUpdateCustomObjectInstanceRequest_CustomAttributes_AdditionalProperties` instead of `map[string]string`.
- Breaking: `NewsItem.Labels`, `NewsItem.Reactions` and `NewsItemRequest.Reactions` are now `*[]string` instead of `*[]*string`. `Predicate.Value` is now `interface{}`. `PaginatedResponse.Pages` is now `*PagesLink`. Removed `PaginatedResponseTypeConversationList`.
- Breaking: removed `BulkContentActionsResponse.JSON403` and `GetJSON403`.
- Added `NewsItemRequest.CoverImageUrl`, `Predicate.Predicates`, `InternalArticleListItem.Tags`, `Locale` and `JsonBlocks` on the internal article create and update requests, and `AdminId` on `DetachTagFromContentSnippetParams` and `DetachTagFromInternalArticleParams`.
- Added `Page` and `PerPage` on `ListInternalArticlesParams`, `SearchInternalArticlesParams`, `ListLiveNewsfeedItemsParams`, `ListNewsItemsParams` and `ListNewsfeedsParams`.
- Added `JSON400`, `JSON401`, `JSON404`, `JSON422` and `JSON403` response fields with `GetJSON` methods on many content snippet, internal article, news item, newsfeed, folder, audience, conversation, dynamic content, custom object and search responses.

## v0.3.0, 2026-10-23

- Breaking: `AskFinResponse.JSON200`, `ReplyToFinResponse.JSON200`, `EscalateFinConversationResponse.JSON200` and `ListFinCapabilitiesResponse.JSON200` have new body types, and `SseSubscriptionUrl` is removed from the Ask Fin body. `CreateContactResponse.JSON200` and `MergeContactResponse.JSON200` also have new body types, as do their `GetJSON200` methods.
- Breaking: `AttachSubscriptionTypeToContactJSONBody.ConsentType` is now `*string`. `FinAgentMessage.Timestamp` is now `*time.Time`. `UpdateExternalPageRequest.ExternalId` is now `string` and `Url` is now `*string`. `ListAllCompaniesParams.Order` is now `*ListAllCompaniesParamsOrder`.
- Breaking: `Contact.Avatar` is now `*string`. `Company.Notes` is now a struct with `Data`, `HasMore` and a total count. `CompanyList.Pages` is now `*PagesLink` and `CompanyScroll.Pages` is now `*map[string]interface{}`.
- Breaking: `CreateDataConnectorRequest` and `UpdateDataConnectorRequest` `DataInputs` and `Headers` have new item types, and header `Name` and `Value` are now plain strings. `RunFinProcedureJSONBody.Message` is now an inline struct with a `RunFinProcedureJSONBodyMessageAuthor` author.
- Breaking: removed `CompanyNote.Author`, `Company` and `CreatedAt`, `ContactAttachedCompanies.Companies` (use `Data`), `UpdateExternalPageRequest.FinAvailability`, `ContactsBulkJobStatePending`, and `ExportWorkflowResponse.JSON403` with `GetJSON403`.
- Breaking: removed `CreateExternalPageRequestLocale`, `ExternalPageLocale`, `UpdateExternalPageRequestLocale` and their `En` constants and `Valid` methods.
- Added `ResponseFormat` with `Html` and `Markdown` values on the Fin start, reply and ask request bodies.
- Added `CompanyUserList`, `CompanySearchResponse`, `ContactSubscriptionTypes`, and the `AsCompanyUserList`, `FromCompanyUserList` and `MergeCompanyUserList` union methods on the company retrieve body. `RetrieveCompanyParams` and `RetrieveACompanyByIdParams` gain `Type`, and the latter also `Page` and `PerPage`.
- Added contact fields `SmsConsent`, `UnsubscribedFromSms`, `LanguageOverride`, `Referrer`, `Utm*` fields, `OptedInSubscriptionTypes`, `OptedOutSubscriptionTypes`, `EnabledPushMessaging`, and `ContinentCode` and `CountryCode` on `ContactLocation`.
- Added `PerPage`, `StartingAfter`, `Page`, `Order` and `EscapeIdentity` params on the contact, company, external page and content import source list calls, and `IncludeMergeHistory` on `ListAttachedContactsParams`.
- Added `EnforceAuth` and `InternalDescription` on data connector types, `TransformCodeBlock` on `DataConnectorDetail`, new data connector error type and source constants, `ExternalPage.Tags`, `AiAgentAvailability`, `AiCopilotAvailability` and `AiSalesAgentAvailability` on external page requests, `CreateOrUpdateCompanyRequest.Id`, `UpdateCompanyRequest.UpdateLastRequestAt` and `CompanyNote.Url`.
- Added `JSON401`, `JSON403`, `JSON404` and `JSON422` response fields on several contact, content import source, external page and Fin responses.

## v0.2.0, 2026-10-22

- Added `GetSubscriptionType` on `Client` and `ClientWithResponses`, with `GetSubscriptionTypeParams`, `GetSubscriptionTypeResponse`, `NewGetSubscriptionTypeRequest` and `ParseGetSubscriptionTypeResponse`.
- Added `DataEvent`, `DataEventList` and `LisDataEvents200JSONResponseBody`, a union with `AsDataEventList`, `AsDataEventSummary` and matching `From`, `Merge` and JSON methods.
- Added `CallTranscriptTurn`, and `ShowCallTranscriptResponse` now has `JSON200`. `Headers200` and `ShowCallTranscriptResponse200Headers` are removed.
- Added article fields `AudienceIds`, `ContentId`, `Locale` and `ParentIds`, `Page` and `PerPage` on the list and search params, `SearchArticlesParamsState`, and `JSON400` and `JSON422` response fields on several article and collection responses.
- Breaking: `Pages` on article, collection, internal article and help center redirect lists is now `*PagesLink` instead of `*CursorPages`. `CursorPages.Next` is now an inline struct instead of `*StartingAfterPaging`.
- Breaking: `CustomAttributes` is now a struct with `AITopicAndSubtopic`, `CXScoreReasons` and `AdditionalProperties`. `AiTopicSubtopicList`, `CxScoreReasonList` and their `As`, `From` and `Merge` methods are removed. Added `AiTopicSubtopic` and `CxScoreReason`.
- Breaking: `CustomAttributes` on company requests is now a map of union types instead of `map[string]string`. This affects `Company`, `CreateOrUpdateCompanyRequest` and `UpdateCompanyRequest`.
- Breaking: `CreateArticleRequest.AuthorId` and `Title` and `CreateCollectionRequest.Name` are now pointers. `CreateArticleRequest.ParentType` and `UpdateArticleRequest.ParentType` are removed.
- Breaking: `CreateDataEventSummariesRequest.EventSummaries` is now a slice of structs, and the type is no longer comparable. `DataEventSummary` loses `Email`, `Events`, `IntercomUserId`, `Type` and `UserId`.
- Breaking: `GetCustomObjectInstancesByIdResponse.JSON200` and `RetrieveCompanyResponse.JSON200` have new body types.
- Breaking: removed `ArticleStatistics.Conversions` (use `Conversations`), `AwayStatusReason.CreatedAt` and `UpdatedAt`, `ConversationPartMetadata.QuickReplyUuid`, `ConversationTeammates.Teammates` (use `Admins`), `PhoneSwitch.Phone`, `Visitor.LasRequestAt` (use `LastRequestAt`) and `TicketTypeAttributeList.TicketTypeAttributes`.
- Breaking: `InternalArticleListItem.Id` is now `*int` and its `Type` is removed. `CreateTicketTypeRequest.Category` is no longer a pointer. `Error.Errors` items have `Data` instead of `Field`.
- Breaking: renamed `SlaAppliedSlaStatusCancelled` to `SlaAppliedSlaStatusCanceled` and `UpdateContentImportSourceRequestSyncBehaviorAutomated` to `UpdateContentImportSourceRequestSyncBehaviorAutomatic`.
- Raised the Go requirement in `go.mod` from 1.26.9 to 1.26.10 to pick up the `crypto/x509` fix for GO-2026-6809. The development image and the setup reference match.

## v0.1.1, 2026-10-17

- Documentation only. The agent skill explains how to look up every endpoint, parameter and model with `go doc` or on pkg.go.dev, and declares its license.
- The skill no longer lists the API credentials under OpenClaw's `requires.env`, which kept OpenClaw from loading it until they were set. It requires only `go`; the credentials are still documented as what your program reads.
- The README points to the full API reference, and the examples in the README and the skill were checked against the current code.

## v0.1.0, 2026-10-17

- Added `intercom`, a REST client generated with oapi-codegen from Intercom's OpenAPI document (Intercom API, spec version 2.17, OpenAPI 3.0.1). It covers all 258 operations across 182 paths, including admins, contacts, companies, conversations, messages, articles, tickets and exports.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- The spec patch adds missing reporting job path parameters, resolves two inline enum Go name collisions, collapses single-branch contact body unions, and names four other inline union schemas so typed requests retain their JSON methods and serialize their fields.
- `make generate` selects the newest numbered stable spec from Intercom's public repository and regenerates the client. Initial source: https://raw.githubusercontent.com/intercom/Intercom-OpenAPI/main/descriptions/2.17/api.intercom.io.yaml. The spec is not committed.
- Added the `go-intercom` agent skill and Claude Code and Codex plugin manifests.
