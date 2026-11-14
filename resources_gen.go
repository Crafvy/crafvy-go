// Code generated from sdks/openapi.json by `pnpm sdk:generate`. DO NOT EDIT.
// Generated from sdks/openapi.json by `pnpm sdk:generate` — do not edit.

package crafvy

import "context"

// resources is every top-level resource; Client embeds it, so they read client.Projects, client.CRM, …
type resources struct {
	AbTests           *AbTestsService
	Ads               *AdsService
	Agents            *AgentsService
	AI                *AiService
	Analytics         *AnalyticsService
	Announcements     *AnnouncementsService
	Approvals         *ApprovalsService
	Apps              *AppsService
	AssetFolders      *AssetFoldersService
	Assets            *AssetsService
	AssetShares       *AssetSharesService
	AuditLogs         *AuditLogsService
	Banking           *BankingService
	Billing           *BillingService
	BuilderBlocks     *BuilderBlocksService
	Calendar          *CalendarService
	Code              *CodeService
	Community         *CommunityService
	Connections       *ConnectionsService
	Countries         *CountriesService
	Creators          *CreatorsService
	CRM               *CrmService
	Currencies        *CurrenciesService
	Dashboards        *DashboardsService
	Data              *DataService
	Databases         *DatabasesService
	Documents         *DocumentsService
	Endpoints         *EndpointsService
	FeatureGates      *FeatureGatesService
	FeatureRequests   *FeatureRequestsService
	Help              *HelpService
	Invoicing         *InvoicingService
	ItemFolders       *ItemFoldersService
	Languages         *LanguagesService
	Mail              *MailService
	Marketplace       *MarketplaceService
	Messenger         *MessengerService
	Milestones        *MilestonesService
	Notifications     *NotificationsService
	OAuth             *OauthService
	Organizations     *OrganizationsService
	Pages             *PagesService
	PDF               *PdfService
	Permissions       *PermissionsService
	PrivacyRequests   *PrivacyRequestsService
	Products          *ProductsService
	Projects          *ProjectsService
	Push              *PushService
	Referrals         *ReferralsService
	RegionalSettings  *RegionalSettingsService
	Reports           *ReportsService
	RetentionPolicies *RetentionPoliciesService
	Roles             *RolesService
	Share             *ShareService
	Social            *SocialService
	Submissions       *SubmissionsService
	Subscriptions     *SubscriptionsService
	Support           *SupportService
	Tasks             *TasksService
	TaxRates          *TaxRatesService
	Teams             *TeamsService
	Telephony         *TelephonyService
	Templates         *TemplatesService
	Trash             *TrashService
	User              *UserService
	Users             *UsersService
	Variables         *VariablesService
	Webhooks          *WebhooksService
	Workflows         *WorkflowsService
	Workspaces        *WorkspacesService
}

func (r *resources) init(c *Client) {
	r.AbTests = newAbTestsService(c)
	r.Ads = newAdsService(c)
	r.Agents = newAgentsService(c)
	r.AI = newAiService(c)
	r.Analytics = newAnalyticsService(c)
	r.Announcements = newAnnouncementsService(c)
	r.Approvals = newApprovalsService(c)
	r.Apps = newAppsService(c)
	r.AssetFolders = newAssetFoldersService(c)
	r.Assets = newAssetsService(c)
	r.AssetShares = newAssetSharesService(c)
	r.AuditLogs = newAuditLogsService(c)
	r.Banking = newBankingService(c)
	r.Billing = newBillingService(c)
	r.BuilderBlocks = newBuilderBlocksService(c)
	r.Calendar = newCalendarService(c)
	r.Code = newCodeService(c)
	r.Community = newCommunityService(c)
	r.Connections = newConnectionsService(c)
	r.Countries = newCountriesService(c)
	r.Creators = newCreatorsService(c)
	r.CRM = newCrmService(c)
	r.Currencies = newCurrenciesService(c)
	r.Dashboards = newDashboardsService(c)
	r.Data = newDataService(c)
	r.Databases = newDatabasesService(c)
	r.Documents = newDocumentsService(c)
	r.Endpoints = newEndpointsService(c)
	r.FeatureGates = newFeatureGatesService(c)
	r.FeatureRequests = newFeatureRequestsService(c)
	r.Help = newHelpService(c)
	r.Invoicing = newInvoicingService(c)
	r.ItemFolders = newItemFoldersService(c)
	r.Languages = newLanguagesService(c)
	r.Mail = newMailService(c)
	r.Marketplace = newMarketplaceService(c)
	r.Messenger = newMessengerService(c)
	r.Milestones = newMilestonesService(c)
	r.Notifications = newNotificationsService(c)
	r.OAuth = newOauthService(c)
	r.Organizations = newOrganizationsService(c)
	r.Pages = newPagesService(c)
	r.PDF = newPdfService(c)
	r.Permissions = newPermissionsService(c)
	r.PrivacyRequests = newPrivacyRequestsService(c)
	r.Products = newProductsService(c)
	r.Projects = newProjectsService(c)
	r.Push = newPushService(c)
	r.Referrals = newReferralsService(c)
	r.RegionalSettings = newRegionalSettingsService(c)
	r.Reports = newReportsService(c)
	r.RetentionPolicies = newRetentionPoliciesService(c)
	r.Roles = newRolesService(c)
	r.Share = newShareService(c)
	r.Social = newSocialService(c)
	r.Submissions = newSubmissionsService(c)
	r.Subscriptions = newSubscriptionsService(c)
	r.Support = newSupportService(c)
	r.Tasks = newTasksService(c)
	r.TaxRates = newTaxRatesService(c)
	r.Teams = newTeamsService(c)
	r.Telephony = newTelephonyService(c)
	r.Templates = newTemplatesService(c)
	r.Trash = newTrashService(c)
	r.User = newUserService(c)
	r.Users = newUsersService(c)
	r.Variables = newVariablesService(c)
	r.Webhooks = newWebhooksService(c)
	r.Workflows = newWorkflowsService(c)
	r.Workspaces = newWorkspacesService(c)
}

// AbTestsService — A/B tests: The stored A/B test rows of builder pages. A test reaches visitors only through the page's A/B actions (`pages.abTestAction`), which take variant B live through the publish gates; these operations create and edit drafts, read results and apply a winner to the page's draft.
type AbTestsService struct {
	c *Client
}

func newAbTestsService(c *Client) *AbTestsService {
	return &AbTestsService{c: c}
}

// Create — Creates an A/B test on a page, as a draft.
//
// Needs `page:update` in the page's workspace and a plan with A/B testing. The test is stored as a draft; it reaches visitors only when started from the page's A/B actions (`pages.abTestAction`).
//
// POST /api/ab-tests
func (r *AbTestsService) Create(ctx context.Context, body *CreateAbTestInput, opts ...CallOption) (*AbTest, error) {
	var out *AbTest
	err := r.c.call(ctx, "POST", "/api/ab-tests", nil, body, false, "data", &out, opts)
	return out, err
}

// Delete — Deletes an A/B test with its recorded events and visitor sessions.
//
// Needs `page:update`.
//
// DELETE /api/ab-tests/{testId}
//
// testID: The A/B test's id.
func (r *AbTestsService) Delete(ctx context.Context, testID string, opts ...CallOption) (*AbTestsDeleteResponse, error) {
	var out *AbTestsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/ab-tests/"+pathSeg(testID), nil, nil, false, "", &out, opts)
	return out, err
}

// Get — Returns one A/B test.
//
// GET /api/ab-tests/{testId}
//
// testID: The A/B test's id.
func (r *AbTestsService) Get(ctx context.Context, testID string, opts ...CallOption) (*AbTest, error) {
	var out *AbTest
	err := r.c.call(ctx, "GET", "/api/ab-tests/"+pathSeg(testID), nil, nil, false, "data", &out, opts)
	return out, err
}

// List — Lists the A/B tests of one page, or without a page the tests the caller created.
//
// GET /api/ab-tests
func (r *AbTestsService) List(ctx context.Context, query *AbTestsListQuery, opts ...CallOption) ([]AbTest, error) {
	var out []AbTest
	err := r.c.call(ctx, "GET", "/api/ab-tests", query, nil, false, "data", &out, opts)
	return out, err
}

// Promote — Applies the winning variant to the page's draft and ends the test.
//
// Needs `page:update`. Variant B's live snapshot becomes the page's draft (unpublished work it replaces is filed as a version first); for A the live page already is the winner and the draft is left alone. Nothing goes live: publishing the page is a separate step.
//
// POST /api/ab-tests/{testId}/promote
//
// testID: The A/B test's id.
func (r *AbTestsService) Promote(ctx context.Context, testID string, body *AbTestsPromoteBody, opts ...CallOption) (*AbTestsPromoteResponseData, error) {
	var out *AbTestsPromoteResponseData
	err := r.c.call(ctx, "POST", "/api/ab-tests/"+pathSeg(testID)+"/promote", nil, body, false, "data", &out, opts)
	return out, err
}

// Update — Changes a draft A/B test's variants, split, goal or dates.
//
// Needs `page:update`. A test that runs or ran on the live site, or a change to `running`, answers 409 `USE_PAGE_AB_TEST`: those go through the page's A/B actions.
//
// PUT /api/ab-tests/{testId}
//
// testID: The A/B test's id.
func (r *AbTestsService) Update(ctx context.Context, testID string, body *UpdateAbTestInput, opts ...CallOption) (*AbTest, error) {
	var out *AbTest
	err := r.c.call(ctx, "PUT", "/api/ab-tests/"+pathSeg(testID), nil, body, false, "data", &out, opts)
	return out, err
}

// AdsService — Ad connections: The ad accounts an organization reports website conversions to (Visitor Journeys): Meta, TikTok, Google Ads and LinkedIn. Every operation needs `ads:manage_connections` in the organization (its administrators by default), and the network must be on the plan. Google Ads and LinkedIn — and Meta or TikTok by sign-in — are connected in the dashboard, where the person signs in to the network.
type AdsService struct {
	c *Client

	Connections *AdsConnectionsService
}

func newAdsService(c *Client) *AdsService {
	return &AdsService{c: c, Connections: newAdsConnectionsService(c)}
}

// AdsConnectionsService — ads.connections
type AdsConnectionsService struct {
	c *Client
}

func newAdsConnectionsService(c *Client) *AdsConnectionsService {
	return &AdsConnectionsService{c: c}
}

// Create — Connects a Meta dataset or a TikTok pixel with a pasted access token.
//
// Meta needs `account.datasetId` (digits) and `credentials.accessToken`, optionally `account.testEventCode`; TikTok needs `account.pixelCode` and `credentials.accessToken`, optionally `account.offlineEventSetId`. Google Ads and LinkedIn are refused (`field: "network"`): they are connected by signing in, in the dashboard. A field that does not fit answers 400 `adConnectionInvalid` with the `field`; 403 `journeysNetworkNotOnPlan` when the network is not on the plan. The token is stored encrypted and never shown again.
//
// POST /api/ads/connections
func (r *AdsConnectionsService) Create(ctx context.Context, body *AdsConnectionsCreateBody, opts ...CallOption) (*AdConnection, error) {
	var out *AdConnection
	err := r.c.call(ctx, "POST", "/api/ads/connections", nil, body, false, "data", &out, opts)
	return out, err
}

// Delete — Disconnects an ad account and deletes its stored credentials and delivery log.
//
// A Google refresh token is also revoked at Google. A project that named the connection keeps the id; nothing is sent to a connection that no longer exists.
//
// DELETE /api/ads/connections/{connectionId}
//
// connectionID: The ad connection's id.
func (r *AdsConnectionsService) Delete(ctx context.Context, connectionID string, opts ...CallOption) (*AdsConnectionsDeleteResponse, error) {
	var out *AdsConnectionsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/ads/connections/"+pathSeg(connectionID), nil, nil, false, "", &out, opts)
	return out, err
}

// List — Lists an organization's ad connections, without their credentials.
//
// GET /api/ads/connections
func (r *AdsConnectionsService) List(ctx context.Context, query *AdsConnectionsListQuery, opts ...CallOption) ([]AdConnection, error) {
	var out []AdConnection
	err := r.c.call(ctx, "GET", "/api/ads/connections", query, nil, false, "data", &out, opts)
	return out, err
}

// Test — Sends one test event through an ad connection and returns the network's answer.
//
// Meta and TikTok receive an event under the connection's test event code (Meta keeps it in the dataset like any other); Google Ads validates the request without taking it in; LinkedIn is only read. No data about any person is sent. A refused token marks the connection `error`.
//
// POST /api/ads/connections/{connectionId}/test
//
// connectionID: The ad connection's id.
func (r *AdsConnectionsService) Test(ctx context.Context, connectionID string, opts ...CallOption) (*AdsConnectionsTestResponseData, error) {
	var out *AdsConnectionsTestResponseData
	err := r.c.call(ctx, "POST", "/api/ads/connections/"+pathSeg(connectionID)+"/test", nil, nil, false, "data", &out, opts)
	return out, err
}

// Update — Renames an ad connection, changes its account ids, replaces its access token or switches it off and on.
//
// `account` is merged into the stored ids (an empty string removes one). A new `credentials.accessToken` replaces the stored one and clears an error; a connection made by signing in takes no pasted token. 400 `adConnectionInvalid` with the `field` that does not fit.
//
// PATCH /api/ads/connections/{connectionId}
//
// connectionID: The ad connection's id.
func (r *AdsConnectionsService) Update(ctx context.Context, connectionID string, body *AdsConnectionsUpdateBody, opts ...CallOption) (*AdConnection, error) {
	var out *AdConnection
	err := r.c.call(ctx, "PATCH", "/api/ads/connections/"+pathSeg(connectionID), nil, body, false, "data", &out, opts)
	return out, err
}

// AgentsService — Agents: The Agent Hub: agents that work towards a goal in the background, acting as their owner — their groups, runs, questions, schedules and shared memory.
type AgentsService struct {
	c *Client

	Groups    *AgentsGroupsService
	Memory    *AgentsMemoryService
	Questions *AgentsQuestionsService
	Runs      *AgentsRunsService
	Schedules *AgentsSchedulesService
}

func newAgentsService(c *Client) *AgentsService {
	return &AgentsService{c: c, Groups: newAgentsGroupsService(c), Memory: newAgentsMemoryService(c), Questions: newAgentsQuestionsService(c), Runs: newAgentsRunsService(c), Schedules: newAgentsSchedulesService(c)}
}

// Activity — What an agent has done: its 50 latest tasks with their results, and today's spend.
//
// Needs `agent:read`.
//
// GET /api/agents/{agentId}/activity
//
// agentID: The agent.
func (r *AgentsService) Activity(ctx context.Context, agentID string, opts ...CallOption) (*AgentsActivityResponse, error) {
	var out *AgentsActivityResponse
	err := r.c.call(ctx, "GET", "/api/agents/"+pathSeg(agentID)+"/activity", nil, nil, false, "", &out, opts)
	return out, err
}

// Create — Creates an agent in a group.
//
// Needs the `agent:create` permission and the Agent Hub on the plan. The agent acts as its creator, with their access. Counts against the plan's agent limit when an `organizationId` is given.
//
// POST /api/agents
func (r *AgentsService) Create(ctx context.Context, body *AgentsCreateBody, opts ...CallOption) (*Agent, error) {
	var out *Agent
	err := r.c.call(ctx, "POST", "/api/agents", nil, body, false, "agent", &out, opts)
	return out, err
}

// Delete — Moves an agent to the trash.
//
// Needs `agent:delete`. Its past runs keep their history. Restored from the trash, it comes back stopped.
//
// DELETE /api/agents/{agentId}
//
// agentID: The agent.
func (r *AgentsService) Delete(ctx context.Context, agentID string, opts ...CallOption) (*AgentsDeleteResponse, error) {
	var out *AgentsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/agents/"+pathSeg(agentID), nil, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the agents the caller may use: their own, and those shared with the workspace or organisation named.
//
// GET /api/agents
func (r *AgentsService) List(ctx context.Context, query *AgentsListQuery, opts ...CallOption) ([]Agent, error) {
	var out []Agent
	err := r.c.call(ctx, "GET", "/api/agents", query, nil, false, "agents", &out, opts)
	return out, err
}

// Queue — Everything queued, running, waiting or stuck across the runs the caller may see, in execution order.
//
// Needs `agent:read`.
//
// GET /api/agents/queue
func (r *AgentsService) Queue(ctx context.Context, query *AgentsQueueQuery, opts ...CallOption) ([]AgentsQueueResponseTasksItem, error) {
	var out []AgentsQueueResponseTasksItem
	err := r.c.call(ctx, "GET", "/api/agents/queue", query, nil, false, "tasks", &out, opts)
	return out, err
}

// Tools — What each agent grant opens: the platform tools behind it, and whether the caller could run each of them here.
//
// An agent acts as its owner, so a tool its owner cannot use (permissions, the plan, the administrator's switches) is one their agent cannot use either.
//
// GET /api/agents/tools
func (r *AgentsService) Tools(ctx context.Context, query *AgentsToolsQuery, opts ...CallOption) ([]AgentsToolsResponseDataGrantsItem, error) {
	var out []AgentsToolsResponseDataGrantsItem
	err := r.c.call(ctx, "GET", "/api/agents/tools", query, nil, false, "data.grants", &out, opts)
	return out, err
}

// Update — Changes an agent, or starts and stops it with `isActive`.
//
// Needs `agent:create`. Only the fields sent change. Stopping an agent lets a task already running finish; the next one is refused.
//
// PATCH /api/agents/{agentId}
//
// agentID: The agent.
func (r *AgentsService) Update(ctx context.Context, agentID string, body *AgentsUpdateBody, opts ...CallOption) (*Agent, error) {
	var out *Agent
	err := r.c.call(ctx, "PATCH", "/api/agents/"+pathSeg(agentID), nil, body, false, "agent", &out, opts)
	return out, err
}

// AgentsGroupsService — agents.groups
type AgentsGroupsService struct {
	c *Client
}

func newAgentsGroupsService(c *Client) *AgentsGroupsService {
	return &AgentsGroupsService{c: c}
}

// Create — Creates an agent group.
//
// Needs `agent:create` and the Agent Hub on the plan; counts against the plan's group limit when an `organizationId` is given.
//
// POST /api/agents/groups
func (r *AgentsGroupsService) Create(ctx context.Context, body *AgentsGroupsCreateBody, opts ...CallOption) (*AgentGroup, error) {
	var out *AgentGroup
	err := r.c.call(ctx, "POST", "/api/agents/groups", nil, body, false, "group", &out, opts)
	return out, err
}

// Delete — Deletes an agent group and moves its agents to the trash.
//
// Needs `agent:delete`. Refused (409) while any of its agents is started — stop them first.
//
// DELETE /api/agents/groups/{groupId}
//
// groupID: The group.
func (r *AgentsGroupsService) Delete(ctx context.Context, groupID string, opts ...CallOption) (*AgentsGroupsDeleteResponse, error) {
	var out *AgentsGroupsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/agents/groups/"+pathSeg(groupID), nil, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the agent groups the caller may see, with their member counts.
//
// GET /api/agents/groups
func (r *AgentsGroupsService) List(ctx context.Context, query *AgentsGroupsListQuery, opts ...CallOption) ([]AgentGroup, error) {
	var out []AgentGroup
	err := r.c.call(ctx, "GET", "/api/agents/groups", query, nil, false, "groups", &out, opts)
	return out, err
}

// Update — Renames, recolours or changes an agent group.
//
// Needs `agent:create`. Only the fields sent change.
//
// PATCH /api/agents/groups/{groupId}
//
// groupID: The group.
func (r *AgentsGroupsService) Update(ctx context.Context, groupID string, body *AgentsGroupsUpdateBody, opts ...CallOption) (*AgentGroup, error) {
	var out *AgentGroup
	err := r.c.call(ctx, "PATCH", "/api/agents/groups/"+pathSeg(groupID), nil, body, false, "group", &out, opts)
	return out, err
}

// AgentsMemoryService — agents.memory
type AgentsMemoryService struct {
	c *Client
}

func newAgentsMemoryService(c *Client) *AgentsMemoryService {
	return &AgentsMemoryService{c: c}
}

// Add — Adds a fact to a workspace's shared memory, which every agent and the assistant read.
//
// Needs `ai_persona:manage_workspace` in that workspace. When the platform has memory switched off nothing is stored, and the answer says `stored: false`.
//
// POST /api/agents/brain
func (r *AgentsMemoryService) Add(ctx context.Context, body *AgentsMemoryAddBody, opts ...CallOption) (*AgentsMemoryAddResponse, error) {
	var out *AgentsMemoryAddResponse
	err := r.c.call(ctx, "POST", "/api/agents/brain", nil, body, false, "", &out, opts)
	return out, err
}

// Forget — Forgets memory entries.
//
// A shared entry needs `ai_persona:manage_workspace` in its workspace; an agent's private one, the right to manage that agent. Entries the caller may not forget are left alone.
//
// DELETE /api/agents/brain
func (r *AgentsMemoryService) Forget(ctx context.Context, query *AgentsMemoryForgetQuery, opts ...CallOption) (float64, error) {
	var out float64
	err := r.c.call(ctx, "DELETE", "/api/agents/brain", query, nil, false, "deleted", &out, opts)
	return out, err
}

// List — Lists what the agents know: the workspace memory of the caller's workspaces, at most 500 entries, latest first.
//
// An agent's private notes are listed only to who may manage that agent. Personal notes are not here.
//
// GET /api/agents/brain
func (r *AgentsMemoryService) List(ctx context.Context, query *AgentsMemoryListQuery, opts ...CallOption) ([]AgentMemoryEntry, error) {
	var out []AgentMemoryEntry
	err := r.c.call(ctx, "GET", "/api/agents/brain", query, nil, false, "entries", &out, opts)
	return out, err
}

// AgentsQuestionsService — agents.questions
type AgentsQuestionsService struct {
	c *Client
}

func newAgentsQuestionsService(c *Client) *AgentsQuestionsService {
	return &AgentsQuestionsService{c: c}
}

// Answer — Answers an agent's question, and the work goes on.
//
// Needs `agent:execute`; only the run's owner may answer. The answer's type follows the question's `kind`: text for `text`, one option value for `single`, a list for `multi`, true or false for `confirm` (a yes runs the held-back step, with the owner's access checked again). 409 when it was answered already.
//
// POST /api/agents/questions/{questionId}
//
// questionID: The question.
// The answer is one of string, []string, bool, decoded as generic JSON (map[string]any, []any, …).
func (r *AgentsQuestionsService) Answer(ctx context.Context, questionID string, body *AgentsQuestionsAnswerBody, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "POST", "/api/agents/questions/"+pathSeg(questionID), nil, body, false, "answer", &out, opts)
	return out, err
}

// Decline — Declines to answer, which stops the run.
//
// Needs `agent:execute`. Every open question of the run is closed and its queued and waiting tasks are cancelled.
//
// DELETE /api/agents/questions/{questionId}
//
// questionID: The question.
func (r *AgentsQuestionsService) Decline(ctx context.Context, questionID string, opts ...CallOption) (*AgentsQuestionsDeclineResponse, error) {
	var out *AgentsQuestionsDeclineResponse
	err := r.c.call(ctx, "DELETE", "/api/agents/questions/"+pathSeg(questionID), nil, nil, false, "", &out, opts)
	return out, err
}

// List — Lists what the agents need the caller to decide, oldest first.
//
// Needs `agent:read`. A question about a resource (a database, a project…) carries as options what the caller may choose from right now.
//
// GET /api/agents/questions
func (r *AgentsQuestionsService) List(ctx context.Context, query *AgentsQuestionsListQuery, opts ...CallOption) ([]AgentQuestion, error) {
	var out []AgentQuestion
	err := r.c.call(ctx, "GET", "/api/agents/questions", query, nil, false, "questions", &out, opts)
	return out, err
}

// AgentsRunsService — agents.runs
type AgentsRunsService struct {
	c *Client
}

func newAgentsRunsService(c *Client) *AgentsRunsService {
	return &AgentsRunsService{c: c}
}

// Cancel — Cancels a run that is still going.
//
// Needs `agent:execute`. Queued and waiting tasks are cancelled; a task already running finishes. 409 when the run has finished.
//
// PATCH /api/agents/runs/{runId}
//
// runID: The run.
func (r *AgentsRunsService) Cancel(ctx context.Context, runID string, opts ...CallOption) (*AgentsRunsCancelResponse, error) {
	var out *AgentsRunsCancelResponse
	err := r.c.call(ctx, "PATCH", "/api/agents/runs/"+pathSeg(runID), nil, nil, false, "", &out, opts)
	return out, err
}

// Delete — Deletes a finished run with its tasks, timeline and questions.
//
// Needs `agent:delete`. 409 while the run is still going — cancel it first.
//
// DELETE /api/agents/runs/{runId}
//
// runID: The run.
func (r *AgentsRunsService) Delete(ctx context.Context, runID string, opts ...CallOption) (*AgentsRunsDeleteResponse, error) {
	var out *AgentsRunsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/agents/runs/"+pathSeg(runID), nil, nil, false, "", &out, opts)
	return out, err
}

// Get — One run with its task tree, its timeline, its questions and the agents involved.
//
// GET /api/agents/runs/{runId}
//
// runID: The run.
func (r *AgentsRunsService) Get(ctx context.Context, runID string, opts ...CallOption) (*AgentsRunsGetResponse, error) {
	var out *AgentsRunsGetResponse
	err := r.c.call(ctx, "GET", "/api/agents/runs/"+pathSeg(runID), nil, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the 50 latest runs the caller may see.
//
// GET /api/agents/runs
func (r *AgentsRunsService) List(ctx context.Context, query *AgentsRunsListQuery, opts ...CallOption) ([]AgentRun, error) {
	var out []AgentRun
	err := r.c.call(ctx, "GET", "/api/agents/runs", query, nil, false, "runs", &out, opts)
	return out, err
}

// Retry — Tries a finished run again: the tasks that did not succeed run again, the ones that did are kept.
//
// Needs `agent:execute`, and only the run's owner may retry it. A run that never planned anything is planned again. 409 while it is still going.
//
// POST /api/agents/runs/{runId}/retry
//
// runID: The run.
func (r *AgentsRunsService) Retry(ctx context.Context, runID string, opts ...CallOption) (*AgentRun, error) {
	var out *AgentRun
	err := r.c.call(ctx, "POST", "/api/agents/runs/"+pathSeg(runID)+"/retry", nil, nil, false, "run", &out, opts)
	return out, err
}

// Start — Starts a run: agents work towards a goal in the background.
//
// Needs `agent:execute` and the Agent Hub on the plan. A supervised run (the default) is planned by a supervisor that hands out tasks; a collaborative run (its own plan feature) lets two or more agents of one group talk in rounds. The run spends the caller's AI allowance, within `tokenBudget`.
//
// POST /api/agents/runs
func (r *AgentsRunsService) Start(ctx context.Context, body *AgentsRunsStartBody, opts ...CallOption) (*AgentRun, error) {
	var out *AgentRun
	err := r.c.call(ctx, "POST", "/api/agents/runs", nil, body, false, "run", &out, opts)
	return out, err
}

// AgentsSchedulesService — agents.schedules
type AgentsSchedulesService struct {
	c *Client
}

func newAgentsSchedulesService(c *Client) *AgentsSchedulesService {
	return &AgentsSchedulesService{c: c}
}

// Create — Creates a schedule that starts a run every so many minutes.
//
// Needs `agent:execute` and the Agent Hub on the plan; counts against the plan's schedule limit when an `organizationId` is given. Its first run starts on the next tick.
//
// POST /api/agents/schedules
func (r *AgentsSchedulesService) Create(ctx context.Context, body *AgentsSchedulesCreateBody, opts ...CallOption) (*AgentSchedule, error) {
	var out *AgentSchedule
	err := r.c.call(ctx, "POST", "/api/agents/schedules", nil, body, false, "schedule", &out, opts)
	return out, err
}

// Delete — Deletes one of the caller's schedules; runs it already started stay.
//
// DELETE /api/agents/schedules/{scheduleId}
//
// scheduleID: The schedule.
func (r *AgentsSchedulesService) Delete(ctx context.Context, scheduleID string, opts ...CallOption) (*AgentsSchedulesDeleteResponse, error) {
	var out *AgentsSchedulesDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/agents/schedules/"+pathSeg(scheduleID), nil, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the runs that start themselves: the caller's schedules, and those of the workspace named.
//
// GET /api/agents/schedules
func (r *AgentsSchedulesService) List(ctx context.Context, query *AgentsSchedulesListQuery, opts ...CallOption) ([]AgentsSchedulesListResponseSchedulesItem, error) {
	var out []AgentsSchedulesListResponseSchedulesItem
	err := r.c.call(ctx, "GET", "/api/agents/schedules", query, nil, false, "schedules", &out, opts)
	return out, err
}

// Update — Pauses, resumes or changes one of the caller's schedules.
//
// Needs `agent:execute`. Resuming counts against the plan's schedule limit. A changed interval counts from now.
//
// PATCH /api/agents/schedules/{scheduleId}
//
// scheduleID: The schedule.
func (r *AgentsSchedulesService) Update(ctx context.Context, scheduleID string, body *AgentsSchedulesUpdateBody, opts ...CallOption) (*AgentSchedule, error) {
	var out *AgentSchedule
	err := r.c.call(ctx, "PATCH", "/api/agents/schedules/"+pathSeg(scheduleID), nil, body, false, "schedule", &out, opts)
	return out, err
}

// AiService — AI: The platform's AI: generating and checking content, the assistant's chats, personas and skills, the receipts of what was generated, and the organization's AI-labelling policy and activity. Every generation counts against the organization's AI quota on its plan.
type AiService struct {
	c *Client

	Chats    *AiChatsService
	Marking  *AiMarkingService
	Personas *AiPersonasService
	Receipts *AiReceiptsService
	Skills   *AiSkillsService
	Video    *AiVideoService
}

func newAiService(c *Client) *AiService {
	return &AiService{c: c, Chats: newAiChatsService(c), Marking: newAiMarkingService(c), Personas: newAiPersonasService(c), Receipts: newAiReceiptsService(c), Skills: newAiSkillsService(c), Video: newAiVideoService(c)}
}

// AnalyzeImage — Describes an image and derives a colour palette, light and dark website themes, fonts and design suggestions from it.
//
// Needs image AI on the organization's plan. Answers 503 `aiNotConfigured` when the platform has no AI configured, 403 when the person opted out of AI processing (Settings → Privacy), and 429 when the organization's AI quota is used up. The image is fetched from `imageUrl`; the same route also takes a multipart form with `file` (JPEG, PNG, WebP or GIF, 10 MB at most), `organizationId` and `focus`. The analysis is the AI's answer: its fields (`description`, `mood`, `colors`, `palette`, `darkPalette`, `typography`, `designSuggestions`) are asked for but not guaranteed.
//
// POST /api/ai/analyze-image
func (r *AiService) AnalyzeImage(ctx context.Context, body *AiAnalyzeImageBody, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/ai/analyze-image", nil, body, false, "analysis", &out, opts)
	return out, err
}

// AnalyzeURL — Reads a public web page and describes its sections, colours, fonts and content, to recreate or take inspiration from it.
//
// Answers 503 `aiNotConfigured` when the platform has no AI configured, 403 when the person opted out of AI processing (Settings → Privacy), and 429 when the organization's AI quota is used up. Only public pages are fetched (3 MB at most). The analysis is the AI's answer: its fields (`title`, `description`, `pageType`, `layout`, `sections`, `colors`, `darkColors`, `fonts`, `content`, `designNotes`) are asked for but not guaranteed.
//
// POST /api/ai/analyze-url
func (r *AiService) AnalyzeURL(ctx context.Context, body *AiAnalyzeUrlBody, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/ai/analyze-url", nil, body, false, "analysis", &out, opts)
	return out, err
}

// BulkOperation — Rewrites or translates every text of a page's component tree at once and answers the changed tree.
//
// Needs `project.create`. Answers 503 `aiNotConfigured` when the platform has no AI configured, 403 when the person opted out of AI processing (Settings → Privacy), and 429 when the organization's AI quota is used up. Off when the platform administrator switched bulk operations off (403). Nothing is saved: the tree comes back for the caller to store.
//
// POST /api/ai/bulk-operations
func (r *AiService) BulkOperation(ctx context.Context, body *AiBulkOperationBody, opts ...CallOption) (*AiBulkOperationResponseData, error) {
	var out *AiBulkOperationResponseData
	err := r.c.call(ctx, "POST", "/api/ai/bulk-operations", nil, body, false, "data", &out, opts)
	return out, err
}

// CheckAccessibility — Checks a page's component tree for accessibility problems and answers the issues with WCAG criteria and fixes.
//
// Answers 503 `aiNotConfigured` when the platform has no AI configured, 403 when the person opted out of AI processing (Settings → Privacy), and 429 when the organization's AI quota is used up. Off when the platform administrator switched the checker off (403).
//
// POST /api/ai/accessibility-check
func (r *AiService) CheckAccessibility(ctx context.Context, body *AiCheckAccessibilityBody, opts ...CallOption) (*AIAccessibilityReport, error) {
	var out *AIAccessibilityReport
	err := r.c.call(ctx, "POST", "/api/ai/accessibility-check", nil, body, false, "data", &out, opts)
	return out, err
}

// ComponentAction — Rewrites, generates or improves the content of one builder component, or suggests a layout for it.
//
// Answers only the props that changed, to merge into the component. Answers 503 `aiNotConfigured` when the platform has no AI configured, 403 when the person opted out of AI processing (Settings → Privacy), and 429 when the organization's AI quota is used up. Off when the platform administrator switched component actions off (403). With a `projectId` the caller may read, the project's facts and tone go into the copy operations.
//
// POST /api/ai/component
func (r *AiService) ComponentAction(ctx context.Context, body *AiComponentActionBody, opts ...CallOption) (*AiComponentActionResponse, error) {
	var out *AiComponentActionResponse
	err := r.c.call(ctx, "POST", "/api/ai/component", nil, body, false, "", &out, opts)
	return out, err
}

// GenerateContent — Writes text from a prompt — a headline, a description, a post — in the requested format, tone and length.
//
// Answers 503 `aiNotConfigured` when the platform has no AI configured, 403 when the person opted out of AI processing (Settings → Privacy), and 429 when the organization's AI quota is used up. Off when the platform administrator switched content generation off (403). With a `projectId` the caller may read, the project's facts, avoided words and tone go into the prompt. The organization's AI-labelling policy decides whether a visible "generated with AI" line is added (`marked`); `markAsAi` chooses where the policy lets the person choose.
//
// POST /api/ai/generate-content
func (r *AiService) GenerateContent(ctx context.Context, body *AiGenerateContentBody, opts ...CallOption) (*AiGenerateContentResponseData, error) {
	var out *AiGenerateContentResponseData
	err := r.c.call(ctx, "POST", "/api/ai/generate-content", nil, body, false, "data", &out, opts)
	return out, err
}

// GenerateImage — Generates an image from a prompt and keeps it as an asset of the organization.
//
// Needs `asset.ai_generate` and `asset.create` in the organization, image generation on its plan and room in its storage; counts against its image quota. The image lands in the asset library with the organization's default visibility; `url` is the asset's own.
//
// POST /api/ai/generate-image
func (r *AiService) GenerateImage(ctx context.Context, body *AiGenerateImageBody, opts ...CallOption) (*AiGenerateImageResponseData, error) {
	var out *AiGenerateImageResponseData
	err := r.c.call(ctx, "POST", "/api/ai/generate-image", nil, body, false, "data", &out, opts)
	return out, err
}

// GenerateProject — Builds a whole project from a description — its pages, and the databases and endpoints it needs.
//
// Needs `project.create`; the workspace must belong to the organization. Answers 503 `aiNotConfigured` when the platform has no AI configured, 403 when the person opted out of AI processing (Settings → Privacy), and 429 when the organization's AI quota is used up. Plans the project, then generates page by page, and answers when everything is made — this can take minutes. Pages that could not be made are listed in `errors`; the rest of the project stays.
//
// POST /api/ai/generate-project
func (r *AiService) GenerateProject(ctx context.Context, body *AiGenerateProjectBody, opts ...CallOption) (*AiGenerateProjectResponseData, error) {
	var out *AiGenerateProjectResponseData
	err := r.c.call(ctx, "POST", "/api/ai/generate-project", nil, body, false, "data", &out, opts)
	return out, err
}

// GenerateSEO — Suggests a meta title, a meta description and keywords for a page.
//
// Answers 503 `aiNotConfigured` when the platform has no AI configured, 403 when the person opted out of AI processing (Settings → Privacy), and 429 when the organization's AI quota is used up. Off when the platform administrator switched SEO generation off (403).
//
// POST /api/ai/generate-seo
func (r *AiService) GenerateSEO(ctx context.Context, body *AiGenerateSeoBody, opts ...CallOption) (*AiGenerateSeoResponseData, error) {
	var out *AiGenerateSeoResponseData
	err := r.c.call(ctx, "POST", "/api/ai/generate-seo", nil, body, false, "data", &out, opts)
	return out, err
}

// GetTimeSaved — Estimates the time the caller's own AI actions saved, per surface and per month.
//
// GET /api/ai/time-saved
func (r *AiService) GetTimeSaved(ctx context.Context, query *AiGetTimeSavedQuery, opts ...CallOption) (*AITimeSavedReport, error) {
	var out *AITimeSavedReport
	err := r.c.call(ctx, "GET", "/api/ai/time-saved", query, nil, false, "", &out, opts)
	return out, err
}

// GetUsage — Returns AI usage per day and per surface, for the caller or for an organization.
//
// Without `organizationId` (or with `userId`) it is the caller's own usage. `organizationId` alone is the whole organization's and needs `ai_usage:view` there; with `userId` as well, that member's usage in that organization (same permission).
//
// GET /api/ai/usage
func (r *AiService) GetUsage(ctx context.Context, query *AiGetUsageQuery, opts ...CallOption) (*AIUsageRollup, error) {
	var out *AIUsageRollup
	err := r.c.call(ctx, "GET", "/api/ai/usage", query, nil, false, "data", &out, opts)
	return out, err
}

// QueryDatabases — Answers a question in plain language over the organization's databases the caller may read, as figures ready for a chart.
//
// Answers 503 `aiNotConfigured` when the platform has no AI configured, 403 when the person opted out of AI processing (Settings → Privacy), and 429 when the organization's AI quota is used up. The AI picks one database and a filter, grouping and aggregation; the platform runs it with the caller's row permissions. `truncated` says only the first `cap` records were counted.
//
// POST /api/ai/db-query
func (r *AiService) QueryDatabases(ctx context.Context, body *AiQueryDatabasesBody, opts ...CallOption) (*AiQueryDatabasesResponseData, error) {
	var out *AiQueryDatabasesResponseData
	err := r.c.call(ctx, "POST", "/api/ai/db-query", nil, body, false, "data", &out, opts)
	return out, err
}

// RewriteDocumentText — Improves, shortens, expands, corrects, continues, summarizes or translates a selection from a document page.
//
// Needs `document.update` and edit access to the document, AI switched on for the document and for the workspace. Refused on vault and restricted pages and for a selection containing text marked sensitive. 20 calls a minute per person. Nothing is saved: the text comes back for the caller to place.
//
// POST /api/ai/documents/rewrite
func (r *AiService) RewriteDocumentText(ctx context.Context, body *AiRewriteDocumentTextBody, opts ...CallOption) (string, error) {
	var out string
	err := r.c.call(ctx, "POST", "/api/ai/documents/rewrite", nil, body, false, "text", &out, opts)
	return out, err
}

// Subtitles — Polishes or translates subtitle cues, keeping their timings, or finds chapter markers in them.
//
// Needs `asset.record_video`, video recording on the organization's plan and the `video:ai_subtitles` feature. `polish` and `translate` answer the cues; `chapters` answers 3 to 8 `{ time, title }` markers, the first at 0.
//
// POST /api/ai/subtitles
func (r *AiService) Subtitles(ctx context.Context, body *AiSubtitlesBody, opts ...CallOption) ([]any, error) {
	var out []any
	err := r.c.call(ctx, "POST", "/api/ai/subtitles", nil, body, false, "result", &out, opts)
	return out, err
}

// AiChatsService — ai.chats
type AiChatsService struct {
	c *Client
}

func newAiChatsService(c *Client) *AiChatsService {
	return &AiChatsService{c: c}
}

// Create — Creates an empty assistant chat in a workspace.
//
// POST /api/ai/chats
func (r *AiChatsService) Create(ctx context.Context, body *AiChatsCreateBody, opts ...CallOption) (*AIChat, error) {
	var out *AIChat
	err := r.c.call(ctx, "POST", "/api/ai/chats", nil, body, false, "chat", &out, opts)
	return out, err
}

// Delete — Deletes one of the caller's assistant chats and the files attached in it.
//
// DELETE /api/ai/chats/{chatId}
//
// chatID: The chat's id.
func (r *AiChatsService) Delete(ctx context.Context, chatID string, opts ...CallOption) (*AiChatsDeleteResponse, error) {
	var out *AiChatsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/ai/chats/"+pathSeg(chatID), nil, nil, false, "", &out, opts)
	return out, err
}

// Get — Returns one of the caller's assistant chats with its messages, oldest first.
//
// GET /api/ai/chats/{chatId}
//
// chatID: The chat's id.
func (r *AiChatsService) Get(ctx context.Context, chatID string, opts ...CallOption) (*AiChatsGetResponse, error) {
	var out *AiChatsGetResponse
	err := r.c.call(ctx, "GET", "/api/ai/chats/"+pathSeg(chatID), nil, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the caller's 50 most recent assistant chats in a workspace.
//
// A chat is its owner's alone. Chats without a message are left out.
//
// GET /api/ai/chats
func (r *AiChatsService) List(ctx context.Context, query *AiChatsListQuery, opts ...CallOption) ([]AIChatSummary, error) {
	var out []AIChatSummary
	err := r.c.call(ctx, "GET", "/api/ai/chats", query, nil, false, "chats", &out, opts)
	return out, err
}

// Rename — Renames one of the caller's assistant chats.
//
// PATCH /api/ai/chats/{chatId}
//
// chatID: The chat's id.
func (r *AiChatsService) Rename(ctx context.Context, chatID string, body *AiChatsRenameBody, opts ...CallOption) (*AIChatSummary, error) {
	var out *AIChatSummary
	err := r.c.call(ctx, "PATCH", "/api/ai/chats/"+pathSeg(chatID), nil, body, false, "chat", &out, opts)
	return out, err
}

// AiMarkingService — ai.marking
type AiMarkingService struct {
	c *Client
}

func newAiMarkingService(c *Client) *AiMarkingService {
	return &AiMarkingService{c: c}
}

// ExportActivity — Downloads the AI generations of the organizations the caller administers as a CSV or JSON file, for an audit.
//
// Same access as the activity list; at most the newest 10,000 events. Answers the file itself (`text/csv` or `application/json`). No receipt's share link, and a prompt preview only on the caller's own events.
//
// GET /api/ai-marking/activity/export
func (r *AiMarkingService) ExportActivity(ctx context.Context, query *AiMarkingExportActivityQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/ai-marking/activity/export", query, nil, false, "", &out, opts)
	return out, err
}

// GetPolicy — Returns an organization's policy for labelling AI-generated content.
//
// `optional`: unlabelled unless the person asks. `recommended`: labelled unless the person declines. `required`: always labelled. Any member may read it.
//
// GET /api/ai-marking/policy
func (r *AiMarkingService) GetPolicy(ctx context.Context, query *AiMarkingGetPolicyQuery, opts ...CallOption) (string, error) {
	var out string
	err := r.c.call(ctx, "GET", "/api/ai-marking/policy", query, nil, false, "data.policy", &out, opts)
	return out, err
}

// ListActivity — Lists the AI generations of the organizations the caller administers, newest first — what was made, by whom, and whether it was labelled.
//
// Needs the admin or owner role in the organization (without `organizationId`, every organization the caller administers). `promptPreview` is filled only on the caller's own events. Page with `before`: the `createdAt` of the last event seen.
//
// GET /api/ai-marking/activity
func (r *AiMarkingService) ListActivity(ctx context.Context, query *AiMarkingListActivityQuery, opts ...CallOption) ([]AIGenerationEvent, error) {
	var out []AIGenerationEvent
	err := r.c.call(ctx, "GET", "/api/ai-marking/activity", query, nil, false, "data.events", &out, opts)
	return out, err
}

// SetPolicy — Sets an organization's policy for labelling AI-generated content.
//
// Needs the admin or owner role in the organization.
//
// PUT /api/ai-marking/policy
func (r *AiMarkingService) SetPolicy(ctx context.Context, body *AiMarkingSetPolicyBody, opts ...CallOption) (string, error) {
	var out string
	err := r.c.call(ctx, "PUT", "/api/ai-marking/policy", nil, body, false, "data.policy", &out, opts)
	return out, err
}

// AiPersonasService — ai.personas
type AiPersonasService struct {
	c *Client
}

func newAiPersonasService(c *Client) *AiPersonasService {
	return &AiPersonasService{c: c}
}

// Clear — Removes the caller's own assistant persona.
//
// DELETE /api/ai/personas
func (r *AiPersonasService) Clear(ctx context.Context, opts ...CallOption) (*AiPersonasClearResponse, error) {
	var out *AiPersonasClearResponse
	err := r.c.call(ctx, "DELETE", "/api/ai/personas", nil, nil, false, "", &out, opts)
	return out, err
}

// ClearOrganization — Removes an organization's assistant persona.
//
// Needs `ai_persona.manage_organization` in the organization.
//
// DELETE /api/ai/personas/organization/{organizationId}
//
// organizationID: The organization's id.
func (r *AiPersonasService) ClearOrganization(ctx context.Context, organizationID string, opts ...CallOption) (*AiPersonasClearOrganizationResponse, error) {
	var out *AiPersonasClearOrganizationResponse
	err := r.c.call(ctx, "DELETE", "/api/ai/personas/organization/"+pathSeg(organizationID), nil, nil, false, "", &out, opts)
	return out, err
}

// ClearWorkspace — Removes a workspace's assistant persona.
//
// Needs `ai_persona.manage_workspace` in the workspace.
//
// DELETE /api/ai/personas/workspace/{workspaceId}
//
// workspaceID: The workspace's id.
func (r *AiPersonasService) ClearWorkspace(ctx context.Context, workspaceID string, opts ...CallOption) (*AiPersonasClearWorkspaceResponse, error) {
	var out *AiPersonasClearWorkspaceResponse
	err := r.c.call(ctx, "DELETE", "/api/ai/personas/workspace/"+pathSeg(workspaceID), nil, nil, false, "", &out, opts)
	return out, err
}

// DeriveFromSamples — Derives a brand voice — persona fields — from 3 to 15 samples of the organization's writing.
//
// Needs the `brand_voice_trainer` feature and `ai_persona.manage_organization` in the organization. Answers 503 `aiNotConfigured` when the platform has no AI configured, 403 when the person opted out of AI processing (Settings → Privacy), and 429 when the organization's AI quota is used up. Saves nothing: apply the result with PUT /api/ai/personas/organization/{organizationId}. Samples shorter than 20 characters are ignored; each is read up to 4,000.
//
// POST /api/ai/persona/from-samples
func (r *AiPersonasService) DeriveFromSamples(ctx context.Context, body *AiPersonasDeriveFromSamplesBody, opts ...CallOption) (*AIDerivedPersona, error) {
	var out *AIDerivedPersona
	err := r.c.call(ctx, "POST", "/api/ai/persona/from-samples", nil, body, false, "data.persona", &out, opts)
	return out, err
}

// Get — Returns the caller's own assistant persona, or null.
//
// GET /api/ai/personas
func (r *AiPersonasService) Get(ctx context.Context, opts ...CallOption) (*AIPersona, error) {
	var out *AIPersona
	err := r.c.call(ctx, "GET", "/api/ai/personas", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetOrganization — Returns an organization's assistant persona, or null.
//
// Needs `ai_persona.manage_organization` in the organization.
//
// GET /api/ai/personas/organization/{organizationId}
//
// organizationID: The organization's id.
func (r *AiPersonasService) GetOrganization(ctx context.Context, organizationID string, opts ...CallOption) (*AIPersona, error) {
	var out *AIPersona
	err := r.c.call(ctx, "GET", "/api/ai/personas/organization/"+pathSeg(organizationID), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetWorkspace — Returns a workspace's assistant persona, or null.
//
// Needs `ai_persona.manage_workspace` in the workspace.
//
// GET /api/ai/personas/workspace/{workspaceId}
//
// workspaceID: The workspace's id.
func (r *AiPersonasService) GetWorkspace(ctx context.Context, workspaceID string, opts ...CallOption) (*AIPersona, error) {
	var out *AIPersona
	err := r.c.call(ctx, "GET", "/api/ai/personas/workspace/"+pathSeg(workspaceID), nil, nil, false, "data", &out, opts)
	return out, err
}

// Preview — Answers a prompt in the voice of a persona that is not saved yet, to try it before rolling it out.
//
// Needs the `brand_voice_trainer` feature and `ai_persona.manage_organization` in the organization. Answers 503 `aiNotConfigured` when the platform has no AI configured, 403 when the person opted out of AI processing (Settings → Privacy), and 429 when the organization's AI quota is used up.
//
// POST /api/ai/persona/preview
func (r *AiPersonasService) Preview(ctx context.Context, body *AiPersonasPreviewBody, opts ...CallOption) (string, error) {
	var out string
	err := r.c.call(ctx, "POST", "/api/ai/persona/preview", nil, body, false, "data.reply", &out, opts)
	return out, err
}

// Set — Saves the caller's own assistant persona — who they are, their work, tone, reply length and format.
//
// PUT /api/ai/personas
func (r *AiPersonasService) Set(ctx context.Context, body *AIPersonaInput, opts ...CallOption) (*AIPersona, error) {
	var out *AIPersona
	err := r.c.call(ctx, "PUT", "/api/ai/personas", nil, body, false, "data", &out, opts)
	return out, err
}

// SetOrganization — Saves an organization's assistant persona.
//
// Needs `ai_persona.manage_organization` in the organization.
//
// PUT /api/ai/personas/organization/{organizationId}
//
// organizationID: The organization's id.
func (r *AiPersonasService) SetOrganization(ctx context.Context, organizationID string, body *AISharedPersonaInput, opts ...CallOption) (*AIPersona, error) {
	var out *AIPersona
	err := r.c.call(ctx, "PUT", "/api/ai/personas/organization/"+pathSeg(organizationID), nil, body, false, "data", &out, opts)
	return out, err
}

// SetWorkspace — Saves a workspace's assistant persona.
//
// Needs `ai_persona.manage_workspace` in the workspace.
//
// PUT /api/ai/personas/workspace/{workspaceId}
//
// workspaceID: The workspace's id.
func (r *AiPersonasService) SetWorkspace(ctx context.Context, workspaceID string, body *AISharedPersonaInput, opts ...CallOption) (*AIPersona, error) {
	var out *AIPersona
	err := r.c.call(ctx, "PUT", "/api/ai/personas/workspace/"+pathSeg(workspaceID), nil, body, false, "data", &out, opts)
	return out, err
}

// AiReceiptsService — ai.receipts
type AiReceiptsService struct {
	c *Client
}

func newAiReceiptsService(c *Client) *AiReceiptsService {
	return &AiReceiptsService{c: c}
}

// Get — Returns one receipt of an AI generation.
//
// Needs the `ai_receipts` feature. The caller's own, or any of the organization's with `ai_usage.view_all` there; anything else answers 404.
//
// GET /api/ai/receipts/{receiptId}
//
// receiptID: The receipt's id.
func (r *AiReceiptsService) Get(ctx context.Context, receiptID string, query *AiReceiptsGetQuery, opts ...CallOption) (*AIReceipt, error) {
	var out *AIReceipt
	err := r.c.call(ctx, "GET", "/api/ai/receipts/"+pathSeg(receiptID), query, nil, false, "receipt", &out, opts)
	return out, err
}

// List — Lists the receipts of AI generations — the caller's own, or a whole organization's.
//
// Needs the `ai_receipts` feature. `scope=org` needs `organizationId` and `ai_usage.view_all` there.
//
// GET /api/ai/receipts
func (r *AiReceiptsService) List(ctx context.Context, query *AiReceiptsListQuery, opts ...CallOption) ([]AIReceipt, error) {
	var out []AIReceipt
	err := r.c.call(ctx, "GET", "/api/ai/receipts", query, nil, false, "receipts", &out, opts)
	return out, err
}

// Share — Makes a public link token for a receipt, or returns the one it has.
//
// Same access as reading the receipt. Anyone with the token can read the receipt (without the token itself) until it is revoked.
//
// POST /api/ai/receipts/{receiptId}/share
//
// receiptID: The receipt's id.
func (r *AiReceiptsService) Share(ctx context.Context, receiptID string, query *AiReceiptsShareQuery, opts ...CallOption) (*AiReceiptsShareResponse, error) {
	var out *AiReceiptsShareResponse
	err := r.c.call(ctx, "POST", "/api/ai/receipts/"+pathSeg(receiptID)+"/share", query, nil, false, "", &out, opts)
	return out, err
}

// Unshare — Revokes a receipt's public link.
//
// DELETE /api/ai/receipts/{receiptId}/share
//
// receiptID: The receipt's id.
func (r *AiReceiptsService) Unshare(ctx context.Context, receiptID string, query *AiReceiptsUnshareQuery, opts ...CallOption) (*AiReceiptsUnshareResponse, error) {
	var out *AiReceiptsUnshareResponse
	err := r.c.call(ctx, "DELETE", "/api/ai/receipts/"+pathSeg(receiptID)+"/share", query, nil, false, "", &out, opts)
	return out, err
}

// AiSkillsService — ai.skills
type AiSkillsService struct {
	c *Client
}

func newAiSkillsService(c *Client) *AiSkillsService {
	return &AiSkillsService{c: c}
}

// AddTemplate — Adds a built-in template as a skill, in the caller's language.
//
// Without a level it becomes the caller's own skill, switched on. At a workspace or organization it needs that level's manage permission, and starts approved only for somebody who may approve there. Once per template and level: a second add answers 409 `alreadyAdded` with the existing skill's id.
//
// POST /api/ai/skills/templates
func (r *AiSkillsService) AddTemplate(ctx context.Context, body *AiSkillsAddTemplateBody, opts ...CallOption) (*AISkill, error) {
	var out *AISkill
	err := r.c.call(ctx, "POST", "/api/ai/skills/templates", nil, body, false, "data", &out, opts)
	return out, err
}

// Approve — Approves and switches on, or rejects and switches off, a workspace's or organization's skill.
//
// Needs `ai_skill.approve` in the skill's own workspace or organization. A person's own skill needs no approval (400).
//
// POST /api/ai/skills/{skillId}/approve
//
// skillID: The skill's id.
func (r *AiSkillsService) Approve(ctx context.Context, skillID string, body *AiSkillsApproveBody, opts ...CallOption) (*AISkill, error) {
	var out *AISkill
	err := r.c.call(ctx, "POST", "/api/ai/skills/"+pathSeg(skillID)+"/approve", nil, body, false, "data", &out, opts)
	return out, err
}

// Create — Creates an assistant skill of the caller's own.
//
// POST /api/ai/skills
func (r *AiSkillsService) Create(ctx context.Context, body *AICreateSkillInput, opts ...CallOption) (*AISkill, error) {
	var out *AISkill
	err := r.c.call(ctx, "POST", "/api/ai/skills", nil, body, false, "data", &out, opts)
	return out, err
}

// CreateOrganization — Creates an assistant skill shared with an organization.
//
// Needs `ai_skill.manage_organization` in the organization. It starts approved and on for somebody with `ai_skill.approve` there, pending and off otherwise.
//
// POST /api/ai/skills/organization/{organizationId}
//
// organizationID: The organization's id.
func (r *AiSkillsService) CreateOrganization(ctx context.Context, organizationID string, body *AICreateSharedSkillInput, opts ...CallOption) (*AISkill, error) {
	var out *AISkill
	err := r.c.call(ctx, "POST", "/api/ai/skills/organization/"+pathSeg(organizationID), nil, body, false, "data", &out, opts)
	return out, err
}

// CreateWorkspace — Creates an assistant skill shared with a workspace.
//
// Needs `ai_skill.manage_workspace` in the workspace. It starts approved and on for somebody with `ai_skill.approve` there, pending and off otherwise.
//
// POST /api/ai/skills/workspace/{workspaceId}
//
// workspaceID: The workspace's id.
func (r *AiSkillsService) CreateWorkspace(ctx context.Context, workspaceID string, body *AICreateSharedSkillInput, opts ...CallOption) (*AISkill, error) {
	var out *AISkill
	err := r.c.call(ctx, "POST", "/api/ai/skills/workspace/"+pathSeg(workspaceID), nil, body, false, "data", &out, opts)
	return out, err
}

// Delete — Deletes an assistant skill.
//
// DELETE /api/ai/skills/{skillId}
//
// skillID: The skill's id.
func (r *AiSkillsService) Delete(ctx context.Context, skillID string, opts ...CallOption) (*AiSkillsDeleteResponse, error) {
	var out *AiSkillsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/ai/skills/"+pathSeg(skillID), nil, nil, false, "", &out, opts)
	return out, err
}

// Get — Returns one assistant skill.
//
// The caller's own skill, or a workspace's or organization's skill where the caller manages that level's skills.
//
// GET /api/ai/skills/{skillId}
//
// skillID: The skill's id.
func (r *AiSkillsService) Get(ctx context.Context, skillID string, opts ...CallOption) (*AISkill, error) {
	var out *AISkill
	err := r.c.call(ctx, "GET", "/api/ai/skills/"+pathSeg(skillID), nil, nil, false, "data", &out, opts)
	return out, err
}

// List — Lists the caller's own assistant skills, most recently changed first.
//
// GET /api/ai/skills
func (r *AiSkillsService) List(ctx context.Context, opts ...CallOption) ([]AISkill, error) {
	var out []AISkill
	err := r.c.call(ctx, "GET", "/api/ai/skills", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListOrganization — Lists an organization's shared assistant skills.
//
// Needs `ai_skill.manage_organization` in the organization.
//
// GET /api/ai/skills/organization/{organizationId}
//
// organizationID: The organization's id.
func (r *AiSkillsService) ListOrganization(ctx context.Context, organizationID string, opts ...CallOption) ([]AISkill, error) {
	var out []AISkill
	err := r.c.call(ctx, "GET", "/api/ai/skills/organization/"+pathSeg(organizationID), nil, nil, false, "data", &out, opts)
	return out, err
}

// ListTemplates — Lists the built-in skill templates and which of them were already added.
//
// Without a level, as the caller's own skills; with `workspaceId` or `organizationId` (one of them), at that level, which needs its manage permission there.
//
// GET /api/ai/skills/templates
func (r *AiSkillsService) ListTemplates(ctx context.Context, query *AiSkillsListTemplatesQuery, opts ...CallOption) ([]AiSkillsListTemplatesResponseDataItem, error) {
	var out []AiSkillsListTemplatesResponseDataItem
	err := r.c.call(ctx, "GET", "/api/ai/skills/templates", query, nil, false, "data", &out, opts)
	return out, err
}

// ListTools — Lists, by hub, the assistant tools the caller may use in a workspace — what a skill may suggest.
//
// GET /api/ai/skills/tools
func (r *AiSkillsService) ListTools(ctx context.Context, query *AiSkillsListToolsQuery, opts ...CallOption) ([]AiSkillsListToolsResponseDataHubsItem, error) {
	var out []AiSkillsListToolsResponseDataHubsItem
	err := r.c.call(ctx, "GET", "/api/ai/skills/tools", query, nil, false, "data.hubs", &out, opts)
	return out, err
}

// ListWorkspace — Lists a workspace's shared assistant skills.
//
// Needs `ai_skill.manage_workspace` in the workspace.
//
// GET /api/ai/skills/workspace/{workspaceId}
//
// workspaceID: The workspace's id.
func (r *AiSkillsService) ListWorkspace(ctx context.Context, workspaceID string, opts ...CallOption) ([]AISkill, error) {
	var out []AISkill
	err := r.c.call(ctx, "GET", "/api/ai/skills/workspace/"+pathSeg(workspaceID), nil, nil, false, "data", &out, opts)
	return out, err
}

// SetOverride — Switches an inherited workspace or organization skill off, or back on, for the caller alone.
//
// Refused (403) for a skill its level enforces.
//
// PUT /api/ai/skills/{skillId}/override
//
// skillID: The skill's id.
func (r *AiSkillsService) SetOverride(ctx context.Context, skillID string, body *AiSkillsSetOverrideBody, opts ...CallOption) (*AiSkillsSetOverrideResponseData, error) {
	var out *AiSkillsSetOverrideResponseData
	err := r.c.call(ctx, "PUT", "/api/ai/skills/"+pathSeg(skillID)+"/override", nil, body, false, "data", &out, opts)
	return out, err
}

// Update — Changes an assistant skill.
//
// The caller's own skill, or a shared one where the caller manages that level's skills. An edit of a shared skill by somebody who may not approve there sends it back for approval; switching a shared skill on or off is POST /api/ai/skills/{skillId}/approve.
//
// PATCH /api/ai/skills/{skillId}
//
// skillID: The skill's id.
func (r *AiSkillsService) Update(ctx context.Context, skillID string, body *AIUpdateSkillInput, opts ...CallOption) (*AISkill, error) {
	var out *AISkill
	err := r.c.call(ctx, "PATCH", "/api/ai/skills/"+pathSeg(skillID), nil, body, false, "data", &out, opts)
	return out, err
}

// AiVideoService — ai.video
type AiVideoService struct {
	c *Client
}

func newAiVideoService(c *Client) *AiVideoService {
	return &AiVideoService{c: c}
}

// Create — Starts a video or GIF generation job and answers its id.
//
// The job runs in the background; poll GET /api/ai/video/{jobId}. The duration is capped at the platform's maximum. `screenshot-walkthrough` needs a public `url` and captures it step by step. Plan limits and the platform's video settings apply.
//
// POST /api/ai/video
func (r *AiVideoService) Create(ctx context.Context, body *AiVideoCreateBody, opts ...CallOption) (*AiVideoCreateResponse, error) {
	var out *AiVideoCreateResponse
	err := r.c.call(ctx, "POST", "/api/ai/video", nil, body, false, "", &out, opts)
	return out, err
}

// Get — Returns a video generation job's status and, once complete, its result.
//
// The caller's own job, or one made for an organization they belong to; any other id answers 404.
//
// GET /api/ai/video/{jobId}
//
// jobID: The job's id.
func (r *AiVideoService) Get(ctx context.Context, jobID string, opts ...CallOption) (*AIVideoJob, error) {
	var out *AIVideoJob
	err := r.c.call(ctx, "GET", "/api/ai/video/"+pathSeg(jobID), nil, nil, false, "", &out, opts)
	return out, err
}

// AnalyticsService — Analytics: Visitor events of builder pages: recording them from a page, the click and scroll heatmaps they make, and refreshing the daily summaries the dashboards read.
type AnalyticsService struct {
	c *Client
}

func newAnalyticsService(c *Client) *AnalyticsService {
	return &AnalyticsService{c: c}
}

// Aggregate — Recomputes a project's daily page and project summaries from its recorded events.
//
// Without `date` or `days`, today. Without `projectId` only a platform administrator may run it. The answer counts the summaries written.
//
// POST /api/analytics/aggregate
func (r *AnalyticsService) Aggregate(ctx context.Context, query *AnalyticsAggregateQuery, opts ...CallOption) (*AnalyticsAggregateResponse, error) {
	var out *AnalyticsAggregateResponse
	err := r.c.call(ctx, "POST", "/api/analytics/aggregate", query, nil, false, "", &out, opts)
	return out, err
}

// Heatmap — Returns a page's click or scroll heatmap for a period.
//
// Clicks come as normalised `[x, y, count]` points on a 200×200 grid plus clicks per component; scrolls as ten depth bands with the visitors who reached each.
//
// GET /api/analytics/heatmap
//
// The answer is one of *AnalyticsHeatmapResponseData, *AnalyticsHeatmapResponseData2, decoded as generic JSON (map[string]any, []any, …).
func (r *AnalyticsService) Heatmap(ctx context.Context, query *AnalyticsHeatmapQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/analytics/heatmap", query, nil, false, "data", &out, opts)
	return out, err
}

// Track — Records a visitor event on a page.
//
// Needs no token: published pages call it. Limited to 120 events a minute per address (beyond that it answers 204 and drops the event); while the platform's analytics are off it answers success and records nothing.
//
// POST /api/analytics/track — No token needed.
func (r *AnalyticsService) Track(ctx context.Context, body *AnalyticsEvent, opts ...CallOption) (*AnalyticsTrackResponse, error) {
	var out *AnalyticsTrackResponse
	err := r.c.call(ctx, "POST", "/api/analytics/track", nil, body, false, "", &out, opts)
	return out, err
}

// AnnouncementsService — Announcements
type AnnouncementsService struct {
	c *Client
}

func newAnnouncementsService(c *Client) *AnnouncementsService {
	return &AnnouncementsService{c: c}
}

// Dismiss — Dismisses an announcement for the caller.
//
// Dismissing it twice is harmless (`alreadyDismissed: true`).
//
// POST /api/announcements/{id}/dismiss
//
// id: The announcement's id.
func (r *AnnouncementsService) Dismiss(ctx context.Context, id string, opts ...CallOption) (*AnnouncementsDismissResponseData, error) {
	var out *AnnouncementsDismissResponseData
	err := r.c.call(ctx, "POST", "/api/announcements/"+pathSeg(id)+"/dismiss", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListActive — Lists the platform's announcements that are live now and addressed to the caller.
//
// Only announcements within their time window and whose audience includes the caller, dismissed ones too (`dismissed: true`). Pinned first, then by priority, then newest.
//
// GET /api/announcements/active
func (r *AnnouncementsService) ListActive(ctx context.Context, opts ...CallOption) ([]Announcement, error) {
	var out []Announcement
	err := r.c.call(ctx, "GET", "/api/announcements/active", nil, nil, false, "data", &out, opts)
	return out, err
}

// ApprovalsService — Approvals
type ApprovalsService struct {
	c *Client
}

func newApprovalsService(c *Client) *ApprovalsService {
	return &ApprovalsService{c: c}
}

// Approve — Approves the current step of an approval as the caller.
//
// When the last step is approved the approval completes, and what it was for takes effect (a page publish, for instance).
//
// POST /api/approvals/{approvalId}/approve
//
// approvalID: The approval's id.
func (r *ApprovalsService) Approve(ctx context.Context, approvalID string, body *ApprovalsApproveBody, opts ...CallOption) (*Approval, error) {
	var out *Approval
	err := r.c.call(ctx, "POST", "/api/approvals/"+pathSeg(approvalID)+"/approve", nil, body, false, "data.workflow", &out, opts)
	return out, err
}

// Bulk — Approves, rejects or cancels up to 50 approvals at once.
//
// Each approval is decided on its own; `results` says which worked.
//
// POST /api/approvals/bulk
func (r *ApprovalsService) Bulk(ctx context.Context, body *ApprovalsBulkBody, opts ...CallOption) ([]ApprovalsBulkResponseDataResultsItem, error) {
	var out []ApprovalsBulkResponseDataResultsItem
	err := r.c.call(ctx, "POST", "/api/approvals/bulk", nil, body, false, "data.results", &out, opts)
	return out, err
}

// Cancel — Cancels a pending approval.
//
// Only its requester or an administrator can cancel it.
//
// DELETE /api/approvals/{approvalId}
//
// approvalID: The approval's id.
func (r *ApprovalsService) Cancel(ctx context.Context, approvalID string, opts ...CallOption) (*ApprovalsCancelResponse, error) {
	var out *ApprovalsCancelResponse
	err := r.c.call(ctx, "DELETE", "/api/approvals/"+pathSeg(approvalID), nil, nil, false, "", &out, opts)
	return out, err
}

// Create — Requests an approval of something — a page, project, asset, workflow, row — from the given approvers.
//
// The approvers of the first step are notified.
//
// POST /api/approvals
func (r *ApprovalsService) Create(ctx context.Context, body *CreateApprovalInput, opts ...CallOption) (*Approval, error) {
	var out *Approval
	err := r.c.call(ctx, "POST", "/api/approvals", nil, body, false, "data", &out, opts)
	return out, err
}

// CreateDelegation — Lets another member decide the caller's approvals of an organization, for a period.
//
// POST /api/approvals/delegations
func (r *ApprovalsService) CreateDelegation(ctx context.Context, body *ApprovalsCreateDelegationBody, opts ...CallOption) (*ApprovalDelegation, error) {
	var out *ApprovalDelegation
	err := r.c.call(ctx, "POST", "/api/approvals/delegations", nil, body, false, "data.delegation", &out, opts)
	return out, err
}

// DeleteDelegation — Ends one of the caller's approval delegations.
//
// DELETE /api/approvals/delegations
func (r *ApprovalsService) DeleteDelegation(ctx context.Context, query *ApprovalsDeleteDelegationQuery, opts ...CallOption) (*ApprovalsDeleteDelegationResponse, error) {
	var out *ApprovalsDeleteDelegationResponse
	err := r.c.call(ctx, "DELETE", "/api/approvals/delegations", query, nil, false, "", &out, opts)
	return out, err
}

// Get — Returns one approval with its steps and approvers.
//
// GET /api/approvals/{approvalId}
//
// approvalID: The approval's id.
func (r *ApprovalsService) Get(ctx context.Context, approvalID string, opts ...CallOption) (*Approval, error) {
	var out *Approval
	err := r.c.call(ctx, "GET", "/api/approvals/"+pathSeg(approvalID), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetSettings — Returns an organization's approval settings: whether publishing needs approval and who approves.
//
// GET /api/approvals/settings
func (r *ApprovalsService) GetSettings(ctx context.Context, query *ApprovalsGetSettingsQuery, opts ...CallOption) (*ApprovalsGetSettingsResponseData, error) {
	var out *ApprovalsGetSettingsResponseData
	err := r.c.call(ctx, "GET", "/api/approvals/settings", query, nil, false, "data", &out, opts)
	return out, err
}

// History — Lists what happened to an approval — requested, approved, rejected, cancelled — oldest first.
//
// GET /api/approvals/{approvalId}/history
//
// approvalID: The approval's id.
func (r *ApprovalsService) History(ctx context.Context, approvalID string, opts ...CallOption) ([]ApprovalHistoryEntry, error) {
	var out []ApprovalHistoryEntry
	err := r.c.call(ctx, "GET", "/api/approvals/"+pathSeg(approvalID)+"/history", nil, nil, false, "data", &out, opts)
	return out, err
}

// List — Lists the approvals the caller may see.
//
// An approval with strict visibility is listed only to its requester and its approvers.
//
// GET /api/approvals
func (r *ApprovalsService) List(ctx context.Context, query *ApprovalsListQuery, opts ...CallOption) ([]Approval, error) {
	var out []Approval
	err := r.c.call(ctx, "GET", "/api/approvals", query, nil, false, "workflows", &out, opts)
	return out, err
}

// ListDelegations — Lists the approval delegations from or to the caller.
//
// GET /api/approvals/delegations
func (r *ApprovalsService) ListDelegations(ctx context.Context, opts ...CallOption) ([]ApprovalDelegation, error) {
	var out []ApprovalDelegation
	err := r.c.call(ctx, "GET", "/api/approvals/delegations", nil, nil, false, "data.delegations", &out, opts)
	return out, err
}

// Reject — Rejects an approval as the caller, with a reason.
//
// POST /api/approvals/{approvalId}/reject
//
// approvalID: The approval's id.
// The answer is one of *Approval, *ApprovalsRejectResponseData2, decoded as generic JSON (map[string]any, []any, …).
func (r *ApprovalsService) Reject(ctx context.Context, approvalID string, body *ApprovalsRejectBody, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "POST", "/api/approvals/"+pathSeg(approvalID)+"/reject", nil, body, false, "data", &out, opts)
	return out, err
}

// Update — Changes an approval's title, description, priority or due date.
//
// The status changes only through approve, reject or cancel; sending `status` is refused.
//
// PATCH /api/approvals/{approvalId}
//
// approvalID: The approval's id.
func (r *ApprovalsService) Update(ctx context.Context, approvalID string, body *ApprovalsUpdateBody, opts ...CallOption) (*Approval, error) {
	var out *Approval
	err := r.c.call(ctx, "PATCH", "/api/approvals/"+pathSeg(approvalID), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateSettings — Changes an organization's approval settings.
//
// Organization administrators only. Name approvers by people or by roles, not both.
//
// PUT /api/approvals/settings
func (r *ApprovalsService) UpdateSettings(ctx context.Context, body *ApprovalsUpdateSettingsBody, opts ...CallOption) (*ApprovalsUpdateSettingsResponse, error) {
	var out *ApprovalsUpdateSettingsResponse
	err := r.c.call(ctx, "PUT", "/api/approvals/settings", nil, body, false, "", &out, opts)
	return out, err
}

// AppsService — Apps: A project's or a hub's own installable app (PWA) and its native desktop and mobile builds.
type AppsService struct {
	c *Client

	Builds             *AppsBuildsService
	Publications       *AppsPublicationsService
	SigningCredentials *AppsSigningCredentialsService
}

func newAppsService(c *Client) *AppsService {
	return &AppsService{c: c, Builds: newAppsBuildsService(c), Publications: newAppsPublicationsService(c), SigningCredentials: newAppsSigningCredentialsService(c)}
}

// AppsBuildsService — apps.builds
type AppsBuildsService struct {
	c *Client
}

func newAppsBuildsService(c *Client) *AppsBuildsService {
	return &AppsBuildsService{c: c}
}

// Create — Queues a native build of a publication for one platform.
//
// Needs `project:publish_native`, the plan's native builds and a verified email; the plan's builds per month and artifact storage apply. Poll `apps.builds.list` for its status.
//
// POST /api/apps/publications/{publicationId}/builds
//
// publicationID: The publication's id.
func (r *AppsBuildsService) Create(ctx context.Context, publicationID string, body *AppsBuildsCreateBody, opts ...CallOption) (*AppBuild, error) {
	var out *AppBuild
	err := r.c.call(ctx, "POST", "/api/apps/publications/"+pathSeg(publicationID)+"/builds", nil, body, false, "data.build", &out, opts)
	return out, err
}

// Delete — Cancels a queued or running build, or deletes a finished one with its file.
//
// DELETE /api/apps/builds/{buildId}
//
// buildID: The build's id.
func (r *AppsBuildsService) Delete(ctx context.Context, buildID string, query *AppsBuildsDeleteQuery, opts ...CallOption) (*AppsBuildsDeleteResponseData, error) {
	var out *AppsBuildsDeleteResponseData
	err := r.c.call(ctx, "DELETE", "/api/apps/builds/"+pathSeg(buildID), query, nil, false, "data", &out, opts)
	return out, err
}

// Download — Downloads a completed build's installer or package.
//
// Answers the file itself (`.dmg`, `.msi`, `.apk`, …), not JSON; the header `X-Content-Sha256` carries its hash. 400 while the build is not completed, 410 once its file expired.
//
// GET /api/apps/builds/{buildId}/download
//
// buildID: The build's id.
func (r *AppsBuildsService) Download(ctx context.Context, buildID string, query *AppsBuildsDownloadQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/apps/builds/"+pathSeg(buildID)+"/download", query, nil, false, "", &out, opts)
	return out, err
}

// List — Lists a publication's builds, newest first.
//
// At most 50.
//
// GET /api/apps/publications/{publicationId}/builds
//
// publicationID: The publication's id.
func (r *AppsBuildsService) List(ctx context.Context, publicationID string, query *AppsBuildsListQuery, opts ...CallOption) ([]AppBuild, error) {
	var out []AppBuild
	err := r.c.call(ctx, "GET", "/api/apps/publications/"+pathSeg(publicationID)+"/builds", query, nil, false, "data.builds", &out, opts)
	return out, err
}

// Update — Sets a build's release notes.
//
// PATCH /api/apps/builds/{buildId}
//
// buildID: The build's id.
func (r *AppsBuildsService) Update(ctx context.Context, buildID string, body *AppsBuildsUpdateBody, opts ...CallOption) (*AppsBuildsUpdateResponse, error) {
	var out *AppsBuildsUpdateResponse
	err := r.c.call(ctx, "PATCH", "/api/apps/builds/"+pathSeg(buildID), nil, body, false, "", &out, opts)
	return out, err
}

// AppsPublicationsService — apps.publications
type AppsPublicationsService struct {
	c *Client
}

func newAppsPublicationsService(c *Client) *AppsPublicationsService {
	return &AppsPublicationsService{c: c}
}

// Create — Creates the app of a project or of a hub.
//
// Give exactly one of `projectId` (needs `project:publish_pwa`) and `hubKey` (needs `organization:manage_hub_apps`); a project or a hub has at most one publication (409). Needs a verified email; counts against the plan's `max_app_publications`.
//
// POST /api/apps/publications
func (r *AppsPublicationsService) Create(ctx context.Context, body *AppsPublicationsCreateBody, opts ...CallOption) (*AppPublication, error) {
	var out *AppPublication
	err := r.c.call(ctx, "POST", "/api/apps/publications", nil, body, false, "data.publication", &out, opts)
	return out, err
}

// Delete — Deletes an app publication and its builds.
//
// DELETE /api/apps/publications/{publicationId}
//
// publicationID: The publication's id.
func (r *AppsPublicationsService) Delete(ctx context.Context, publicationID string, query *AppsPublicationsDeleteQuery, opts ...CallOption) (*AppsPublicationsDeleteResponse, error) {
	var out *AppsPublicationsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/apps/publications/"+pathSeg(publicationID), query, nil, false, "", &out, opts)
	return out, err
}

// Get — Returns one app publication.
//
// GET /api/apps/publications/{publicationId}
//
// publicationID: The publication's id.
func (r *AppsPublicationsService) Get(ctx context.Context, publicationID string, query *AppsPublicationsGetQuery, opts ...CallOption) (*AppPublication, error) {
	var out *AppPublication
	err := r.c.call(ctx, "GET", "/api/apps/publications/"+pathSeg(publicationID), query, nil, false, "data.publication", &out, opts)
	return out, err
}

// List — Lists an organization's app publications, most recently changed first.
//
// At most 50.
//
// GET /api/apps/publications
func (r *AppsPublicationsService) List(ctx context.Context, query *AppsPublicationsListQuery, opts ...CallOption) ([]AppPublication, error) {
	var out []AppPublication
	err := r.c.call(ctx, "GET", "/api/apps/publications", query, nil, false, "data.publications", &out, opts)
	return out, err
}

// Update — Changes an app publication's name, look, version or channels.
//
// Only the fields sent change. Turning on `desktopEnabled` or `mobileEnabled` needs the plan's native builds.
//
// PATCH /api/apps/publications/{publicationId}
//
// publicationID: The publication's id.
func (r *AppsPublicationsService) Update(ctx context.Context, publicationID string, body *AppsPublicationsUpdateBody, opts ...CallOption) (*AppPublication, error) {
	var out *AppPublication
	err := r.c.call(ctx, "PATCH", "/api/apps/publications/"+pathSeg(publicationID), nil, body, false, "data.publication", &out, opts)
	return out, err
}

// AppsSigningCredentialsService — apps.signingCredentials
type AppsSigningCredentialsService struct {
	c *Client
}

func newAppsSigningCredentialsService(c *Client) *AppsSigningCredentialsService {
	return &AppsSigningCredentialsService{c: c}
}

// Create — Stores signing material for native builds, encrypted.
//
// Needs `organization:manage_signing_credentials` and the plan's store submission. `blob` is the material, laid out per platform as the build runner reads it; binary files (.p12, .jks, .pfx, .mobileprovision) base64-encoded:
//
// - `apple`: `teamId`, `appStoreConnectKeyId`, `appStoreConnectIssuerId`, `apiKeyP8`, `distributionCertP12`, `distributionCertPassword`, `provisioningProfile`, `autoSubmitAfterBuild`
// - `google`: `serviceAccountJson`, `keystoreJks`, `keystorePassword`, `keyAlias`, `keyPassword`, `packageName`, `appSigningSha256`, `submitTrack`, `autoSubmitAfterBuild`
// - `macos-developer-id`: `developerIdCertP12`, `developerIdCertPassword`, `signingIdentity`, `teamId`, `appleId`, `appSpecificPassword`
// - `windows-authenticode`: `codeSigningPfx`, `codeSigningPfxPassword`, `timestampUrl`
//
// POST /api/apps/signing-credentials
func (r *AppsSigningCredentialsService) Create(ctx context.Context, body *AppsSigningCredentialsCreateBody, opts ...CallOption) (*AppsSigningCredentialsCreateResponseDataCredential, error) {
	var out *AppsSigningCredentialsCreateResponseDataCredential
	err := r.c.call(ctx, "POST", "/api/apps/signing-credentials", nil, body, false, "data.credential", &out, opts)
	return out, err
}

// Delete — Deletes a signing credential.
//
// Needs `organization:manage_signing_credentials`.
//
// DELETE /api/apps/signing-credentials/{credentialId}
//
// credentialID: The credential's id.
func (r *AppsSigningCredentialsService) Delete(ctx context.Context, credentialID string, query *AppsSigningCredentialsDeleteQuery, opts ...CallOption) (*AppsSigningCredentialsDeleteResponse, error) {
	var out *AppsSigningCredentialsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/apps/signing-credentials/"+pathSeg(credentialID), query, nil, false, "", &out, opts)
	return out, err
}

// List — Lists an organization's signing credentials, without their secrets.
//
// GET /api/apps/signing-credentials
func (r *AppsSigningCredentialsService) List(ctx context.Context, query *AppsSigningCredentialsListQuery, opts ...CallOption) ([]AppSigningCredential, error) {
	var out []AppSigningCredential
	err := r.c.call(ctx, "GET", "/api/apps/signing-credentials", query, nil, false, "data.credentials", &out, opts)
	return out, err
}

// AssetFoldersService — File folders
type AssetFoldersService struct {
	c *Client
}

func newAssetFoldersService(c *Client) *AssetFoldersService {
	return &AssetFoldersService{c: c}
}

// Create — Creates a file folder.
//
// POST /api/asset-folders — Token permission: write:assets.
func (r *AssetFoldersService) Create(ctx context.Context, body *CreateAssetFolderInput, opts ...CallOption) (*AssetFolder, error) {
	var out *AssetFolder
	err := r.c.call(ctx, "POST", "/api/asset-folders", nil, body, false, "folder", &out, opts)
	return out, err
}

// Delete — Deletes a file folder and moves its files to the trash.
//
// Refused while the folder has subfolders, and when the caller may not delete every file in it. The folder's shares are removed.
//
// DELETE /api/asset-folders/{folderId} — Token permission: write:assets.
//
// folderID: The folder's id.
func (r *AssetFoldersService) Delete(ctx context.Context, folderID string, opts ...CallOption) (*AssetFoldersDeleteResponse, error) {
	var out *AssetFoldersDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/asset-folders/"+pathSeg(folderID), nil, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the file folders of an organization the caller may see.
//
// GET /api/asset-folders — Token permission: read:assets.
func (r *AssetFoldersService) List(ctx context.Context, query *AssetFoldersListQuery, opts ...CallOption) ([]AssetFolder, error) {
	var out []AssetFolder
	err := r.c.call(ctx, "GET", "/api/asset-folders", query, nil, false, "folders", &out, opts)
	return out, err
}

// Update — Renames a file folder, moves it, or changes its visibility.
//
// PATCH /api/asset-folders/{folderId} — Token permission: write:assets.
//
// folderID: The folder's id.
func (r *AssetFoldersService) Update(ctx context.Context, folderID string, body *UpdateAssetFolderInput, opts ...CallOption) (*AssetFolder, error) {
	var out *AssetFolder
	err := r.c.call(ctx, "PATCH", "/api/asset-folders/"+pathSeg(folderID), nil, body, false, "folder", &out, opts)
	return out, err
}

// AssetsService — Files
type AssetsService struct {
	c *Client

	Versions *AssetsVersionsService
}

func newAssetsService(c *Client) *AssetsService {
	return &AssetsService{c: c, Versions: newAssetsVersionsService(c)}
}

// Bulk — Moves, trashes or changes the visibility of up to 500 files and 500 folders at once.
//
// Each item is handled on its own; `breakdown` and `errors` say what did not work.
//
// POST /api/assets/bulk — Token permission: write:assets.
func (r *AssetsService) Bulk(ctx context.Context, body *AssetsBulkBody, opts ...CallOption) (*AssetsBulkResponse, error) {
	var out *AssetsBulkResponse
	err := r.c.call(ctx, "POST", "/api/assets/bulk", nil, body, false, "", &out, opts)
	return out, err
}

// CreateText — Creates a text file — Markdown, plain text, JSON, CSV, YAML or HTML — from its content.
//
// At most 5 MB; counts against the plan's storage. JSON must parse.
//
// POST /api/assets/create-text — Token permission: write:assets.
func (r *AssetsService) CreateText(ctx context.Context, body *AssetsCreateTextBody, opts ...CallOption) (*Asset, error) {
	var out *Asset
	err := r.c.call(ctx, "POST", "/api/assets/create-text", nil, body, false, "asset", &out, opts)
	return out, err
}

// Delete — Moves a file to the trash.
//
// DELETE /api/assets/{assetId} — Token permission: write:assets.
//
// assetID: The file's id.
func (r *AssetsService) Delete(ctx context.Context, assetID string, opts ...CallOption) (*AssetsDeleteResponse, error) {
	var out *AssetsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/assets/"+pathSeg(assetID), nil, nil, false, "", &out, opts)
	return out, err
}

// EmptyTrash — Deletes for good every file in an organization's trash that the caller may delete.
//
// DELETE /api/assets/trash — Token permission: write:assets.
func (r *AssetsService) EmptyTrash(ctx context.Context, query *AssetsEmptyTrashQuery, opts ...CallOption) (*AssetsEmptyTrashResponse, error) {
	var out *AssetsEmptyTrashResponse
	err := r.c.call(ctx, "DELETE", "/api/assets/trash", query, nil, false, "", &out, opts)
	return out, err
}

// GenerateImage — Generates an image with AI from a prompt and saves it to the media library.
//
// Needs `asset.ai_generate`; counts against the organization's AI budget and storage.
//
// POST /api/assets/ai-generate-image — Token permission: write:assets.
func (r *AssetsService) GenerateImage(ctx context.Context, body *AssetsGenerateImageBody, opts ...CallOption) (*Asset, error) {
	var out *Asset
	err := r.c.call(ctx, "POST", "/api/assets/ai-generate-image", nil, body, false, "asset", &out, opts)
	return out, err
}

// Get — Returns one file's details.
//
// GET /api/assets/{assetId} — Token permission: read:assets.
//
// assetID: The file's id.
func (r *AssetsService) Get(ctx context.Context, assetID string, opts ...CallOption) (*Asset, error) {
	var out *Asset
	err := r.c.call(ctx, "GET", "/api/assets/"+pathSeg(assetID), nil, nil, false, "asset", &out, opts)
	return out, err
}

// Intelligence — Reports an organization's storage: totals, the biggest folders and uploaders, stale files and duplicates.
//
// Needs `asset.intelligence`.
//
// GET /api/assets/intelligence — Token permission: read:assets.
func (r *AssetsService) Intelligence(ctx context.Context, query *AssetsIntelligenceQuery, opts ...CallOption) (*AssetsIntelligenceResponse, error) {
	var out *AssetsIntelligenceResponse
	err := r.c.call(ctx, "GET", "/api/assets/intelligence", query, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the files of an organization the caller may see, newest first.
//
// GET /api/assets — Token permission: read:assets.
func (r *AssetsService) List(ctx context.Context, query *AssetsListQuery, opts ...CallOption) ([]Asset, error) {
	var out []Asset
	err := r.c.call(ctx, "GET", "/api/assets", query, nil, false, "assets", &out, opts)
	return out, err
}

// ListTrash — Lists the files in an organization's trash that the caller may restore, with when each is deleted for good.
//
// GET /api/assets/trash — Token permission: read:assets.
func (r *AssetsService) ListTrash(ctx context.Context, query *AssetsListTrashQuery, opts ...CallOption) ([]AssetsListTrashResponseAssetsItem, error) {
	var out []AssetsListTrashResponseAssetsItem
	err := r.c.call(ctx, "GET", "/api/assets/trash", query, nil, false, "assets", &out, opts)
	return out, err
}

// StorageUsage — Returns how much of its storage an organization uses.
//
// GET /api/assets/storage-usage — Token permission: read:assets.
func (r *AssetsService) StorageUsage(ctx context.Context, query *AssetsStorageUsageQuery, opts ...CallOption) (*AssetsStorageUsageResponseData, error) {
	var out *AssetsStorageUsageResponseData
	err := r.c.call(ctx, "GET", "/api/assets/storage-usage", query, nil, false, "data", &out, opts)
	return out, err
}

// Sync — The files and folders of an organization changed since a moment, oldest change first, for keeping a copy in step.
//
// Needs desktop sync on the plan and the role's `integration:use_desktop` permission (403 `desktopSyncNotPermitted`). Only what the caller may see; the platform's own folders and the files in them are left out. 500 items a page: while `nextCursor` is set, ask again with it and the same `since`. With `since`, `deletedAssetIds` names the files moved to the trash since then (on the first page only). Keep `serverTime` as the next `since`.
//
// GET /api/sync
func (r *AssetsService) Sync(ctx context.Context, query *AssetsSyncQuery, opts ...CallOption) (*AssetsSyncResponse, error) {
	var out *AssetsSyncResponse
	err := r.c.call(ctx, "GET", "/api/sync", query, nil, false, "", &out, opts)
	return out, err
}

// Transcribe — Transcribes an audio or video file with AI and returns the text, without saving it.
//
// Needs `asset.transcribe`; counts against the organization's AI budget. Save the text with `assets.createText`.
//
// POST /api/assets/transcribe — Token permission: write:assets.
func (r *AssetsService) Transcribe(ctx context.Context, body *AssetsTranscribeBody, opts ...CallOption) (*AssetsTranscribeResponse, error) {
	var out *AssetsTranscribeResponse
	err := r.c.call(ctx, "POST", "/api/assets/transcribe", nil, body, false, "", &out, opts)
	return out, err
}

// TrashAction — Restores files from the trash, or deletes them for good.
//
// Files the caller may not manage are reported in `errors` and left alone.
//
// POST /api/assets/trash — Token permission: write:assets.
func (r *AssetsService) TrashAction(ctx context.Context, body *AssetsTrashActionBody, opts ...CallOption) (*AssetsTrashActionResponse, error) {
	var out *AssetsTrashActionResponse
	err := r.c.call(ctx, "POST", "/api/assets/trash", nil, body, false, "", &out, opts)
	return out, err
}

// Update — Renames a file, moves it to another folder or workspace, or changes its tags or visibility.
//
// PATCH /api/assets/{assetId} — Token permission: write:assets.
//
// assetID: The file's id.
func (r *AssetsService) Update(ctx context.Context, assetID string, body *UpdateAssetInput, opts ...CallOption) (*Asset, error) {
	var out *Asset
	err := r.c.call(ctx, "PATCH", "/api/assets/"+pathSeg(assetID), nil, body, false, "asset", &out, opts)
	return out, err
}

// Upload — Uploads a file to the media library.
//
// Needs `asset.create`; counts against the plan's storage and file-size limits. With `overwriteAssetId` the file replaces that file's content as a new version, and the answer is `{ success, data }` instead of `{ asset }`.
//
// POST /api/assets — Token permission: write:assets.
func (r *AssetsService) Upload(ctx context.Context, body *AssetsUploadBody, opts ...CallOption) (*AssetsUploadResponse, error) {
	var out *AssetsUploadResponse
	err := r.c.call(ctx, "POST", "/api/assets", nil, body, true, "", &out, opts)
	return out, err
}

// Usage — Lists the pages that show a file, so you know whether it is safe to delete.
//
// Needs `asset.intelligence`.
//
// GET /api/assets/{assetId}/usage — Token permission: read:assets.
//
// assetID: The file's id.
func (r *AssetsService) Usage(ctx context.Context, assetID string, query *AssetsUsageQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/assets/"+pathSeg(assetID)+"/usage", query, nil, false, "references", &out, opts)
	return out, err
}

// Zip — Downloads files and folders as one ZIP archive, or with `check: true` says what it would hold.
//
// Without `check` the answer is the archive itself (`application/zip`). Only files the caller may view are included; folders are kept as paths.
//
// POST /api/assets/zip — Token permission: write:assets.
func (r *AssetsService) Zip(ctx context.Context, body *AssetsZipBody, opts ...CallOption) (*AssetsZipResponse, error) {
	var out *AssetsZipResponse
	err := r.c.call(ctx, "POST", "/api/assets/zip", nil, body, false, "", &out, opts)
	return out, err
}

// AssetsVersionsService — assets.versions
type AssetsVersionsService struct {
	c *Client
}

func newAssetsVersionsService(c *Client) *AssetsVersionsService {
	return &AssetsVersionsService{c: c}
}

// Delete — Deletes a version of a file.
//
// DELETE /api/assets/{assetId}/versions/{versionId} — Token permission: write:assets.
//
// assetID: The file's id.
// versionID: The version's id.
func (r *AssetsVersionsService) Delete(ctx context.Context, assetID string, versionID string, query *AssetsVersionsDeleteQuery, opts ...CallOption) (*AssetsVersionsDeleteResponse, error) {
	var out *AssetsVersionsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/assets/"+pathSeg(assetID)+"/versions/"+pathSeg(versionID), query, nil, false, "", &out, opts)
	return out, err
}

// Diff — Compares the text of two versions of a file.
//
// POST /api/assets/{assetId}/versions/diff — Token permission: write:assets.
//
// assetID: The file's id.
func (r *AssetsVersionsService) Diff(ctx context.Context, assetID string, body *AssetsVersionsDiffBody, opts ...CallOption) (*AssetVersionDiff, error) {
	var out *AssetVersionDiff
	err := r.c.call(ctx, "POST", "/api/assets/"+pathSeg(assetID)+"/versions/diff", nil, body, false, "data", &out, opts)
	return out, err
}

// Download — Downloads the content of a version of a file.
//
// Answers the file itself, not JSON.
//
// GET /api/assets/{assetId}/versions/{versionId}/download — Token permission: read:assets.
//
// assetID: The file's id.
// versionID: The version's id.
func (r *AssetsVersionsService) Download(ctx context.Context, assetID string, versionID string, query *AssetsVersionsDownloadQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/assets/"+pathSeg(assetID)+"/versions/"+pathSeg(versionID)+"/download", query, nil, false, "", &out, opts)
	return out, err
}

// Get — Returns one version of a file, with its extracted text.
//
// GET /api/assets/{assetId}/versions/{versionId} — Token permission: read:assets.
//
// assetID: The file's id.
// versionID: The version's id.
func (r *AssetsVersionsService) Get(ctx context.Context, assetID string, versionID string, query *AssetsVersionsGetQuery, opts ...CallOption) (*AssetsVersionsGetResponseData, error) {
	var out *AssetsVersionsGetResponseData
	err := r.c.call(ctx, "GET", "/api/assets/"+pathSeg(assetID)+"/versions/"+pathSeg(versionID), query, nil, false, "data", &out, opts)
	return out, err
}

// List — Lists a file's versions, newest first, and with `q` the versions whose text matches.
//
// GET /api/assets/{assetId}/versions — Token permission: read:assets.
//
// assetID: The file's id.
func (r *AssetsVersionsService) List(ctx context.Context, assetID string, query *AssetsVersionsListQuery, opts ...CallOption) ([]AssetVersion, error) {
	var out []AssetVersion
	err := r.c.call(ctx, "GET", "/api/assets/"+pathSeg(assetID)+"/versions", query, nil, false, "data.versions", &out, opts)
	return out, err
}

// Restore — Makes an earlier version a file's current content, keeping the content it replaces as a version.
//
// POST /api/assets/{assetId}/versions/{versionId}/restore — Token permission: write:assets.
//
// assetID: The file's id.
// versionID: The version's id.
func (r *AssetsVersionsService) Restore(ctx context.Context, assetID string, versionID string, body *AssetsVersionsRestoreBody, opts ...CallOption) (*AssetsVersionsRestoreResponseData, error) {
	var out *AssetsVersionsRestoreResponseData
	err := r.c.call(ctx, "POST", "/api/assets/"+pathSeg(assetID)+"/versions/"+pathSeg(versionID)+"/restore", nil, body, false, "data", &out, opts)
	return out, err
}

// Update — Labels a version of a file, or marks it as a milestone.
//
// PATCH /api/assets/{assetId}/versions/{versionId} — Token permission: write:assets.
//
// assetID: The file's id.
// versionID: The version's id.
func (r *AssetsVersionsService) Update(ctx context.Context, assetID string, versionID string, body *AssetsVersionsUpdateBody, opts ...CallOption) (*AssetVersion, error) {
	var out *AssetVersion
	err := r.c.call(ctx, "PATCH", "/api/assets/"+pathSeg(assetID)+"/versions/"+pathSeg(versionID), nil, body, false, "data.version", &out, opts)
	return out, err
}

// AssetSharesService — File shares
type AssetSharesService struct {
	c *Client
}

func newAssetSharesService(c *Client) *AssetSharesService {
	return &AssetSharesService{c: c}
}

// Create — Shares a file with a member, or makes a share link for it.
//
// A link's expiry is capped at the organization's maximum, and a public link is refused when the organization does not allow them.
//
// POST /api/assets/{assetId}/shares — Token permission: write:assets.
//
// assetID: The file's id.
func (r *AssetSharesService) Create(ctx context.Context, assetID string, body *CreateAssetShareInput, opts ...CallOption) (*AssetShare, error) {
	var out *AssetShare
	err := r.c.call(ctx, "POST", "/api/assets/"+pathSeg(assetID)+"/shares", nil, body, false, "share", &out, opts)
	return out, err
}

// CreateFolder — Shares a folder with a member, or makes a share link for it.
//
// POST /api/asset-folders/{folderId}/shares — Token permission: write:assets.
//
// folderID: The folder's id.
func (r *AssetSharesService) CreateFolder(ctx context.Context, folderID string, body *CreateAssetShareInput, opts ...CallOption) (*AssetShare, error) {
	var out *AssetShare
	err := r.c.call(ctx, "POST", "/api/asset-folders/"+pathSeg(folderID)+"/shares", nil, body, false, "share", &out, opts)
	return out, err
}

// Delete — Removes a file share.
//
// DELETE /api/assets/{assetId}/shares/{shareId} — Token permission: write:assets.
//
// assetID: The file's id.
// shareID: The share's id.
func (r *AssetSharesService) Delete(ctx context.Context, assetID string, shareID string, opts ...CallOption) (*AssetSharesDeleteResponse, error) {
	var out *AssetSharesDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/assets/"+pathSeg(assetID)+"/shares/"+pathSeg(shareID), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteByID — Removes any share — of a file or a folder — by its id.
//
// DELETE /api/asset-shares/{shareId}
//
// shareID: The share's id.
func (r *AssetSharesService) DeleteByID(ctx context.Context, shareID string, opts ...CallOption) (*AssetSharesDeleteByIdResponse, error) {
	var out *AssetSharesDeleteByIdResponse
	err := r.c.call(ctx, "DELETE", "/api/asset-shares/"+pathSeg(shareID), nil, nil, false, "", &out, opts)
	return out, err
}

// List — Lists a file's live shares — with members and as links.
//
// GET /api/assets/{assetId}/shares — Token permission: read:assets.
//
// assetID: The file's id.
func (r *AssetSharesService) List(ctx context.Context, assetID string, opts ...CallOption) ([]AssetShare, error) {
	var out []AssetShare
	err := r.c.call(ctx, "GET", "/api/assets/"+pathSeg(assetID)+"/shares", nil, nil, false, "shares", &out, opts)
	return out, err
}

// ListFolder — Lists a folder's live shares.
//
// GET /api/asset-folders/{folderId}/shares — Token permission: read:assets.
//
// folderID: The folder's id.
func (r *AssetSharesService) ListFolder(ctx context.Context, folderID string, opts ...CallOption) ([]AssetShare, error) {
	var out []AssetShare
	err := r.c.call(ctx, "GET", "/api/asset-folders/"+pathSeg(folderID)+"/shares", nil, nil, false, "shares", &out, opts)
	return out, err
}

// Update — Changes a file share's permission, expiry or password, or turns it off.
//
// PATCH /api/assets/{assetId}/shares/{shareId} — Token permission: write:assets.
//
// assetID: The file's id.
// shareID: The share's id.
func (r *AssetSharesService) Update(ctx context.Context, assetID string, shareID string, body *UpdateAssetShareInput, opts ...CallOption) (*AssetShare, error) {
	var out *AssetShare
	err := r.c.call(ctx, "PATCH", "/api/assets/"+pathSeg(assetID)+"/shares/"+pathSeg(shareID), nil, body, false, "share", &out, opts)
	return out, err
}

// UpdateByID — Changes any share — of a file or a folder — by its id.
//
// PATCH /api/asset-shares/{shareId}
//
// shareID: The share's id.
func (r *AssetSharesService) UpdateByID(ctx context.Context, shareID string, body *UpdateAssetShareInput, opts ...CallOption) (*AssetShare, error) {
	var out *AssetShare
	err := r.c.call(ctx, "PATCH", "/api/asset-shares/"+pathSeg(shareID), nil, body, false, "share", &out, opts)
	return out, err
}

// AuditLogsService — Audit log: An organization's audit trail: every sign-in, change of a member, role or setting and every write the platform records for it, for its owners and admins.
type AuditLogsService struct {
	c *Client
}

func newAuditLogsService(c *Client) *AuditLogsService {
	return &AuditLogsService{c: c}
}

// List — Lists an organization's audit log, newest first.
//
// Needs an owner or admin of the organization (403 `adminRequired`), and the audit log on the plan (402 `featureNotOnPlan`). `format=csv` downloads the same lines as a CSV file.
//
// GET /api/audit-logs
func (r *AuditLogsService) List(ctx context.Context, query *AuditLogsListQuery, opts ...CallOption) ([]AuditLogEvent, error) {
	var out []AuditLogEvent
	err := r.c.call(ctx, "GET", "/api/audit-logs", query, nil, false, "logs", &out, opts)
	return out, err
}

// BankingService — Banking: The organisation's bank accounts and their lines: statement imports, matching to invoices and purchases, rules, receipts, expected payments, payment orders, the liquidity forecast. Every operation works in one organisation (`organizationId`) and needs the role's `banking:*` permissions held at organisation level, plus the right on each account (a grant, its visibility). Connecting a bank goes through finAPI's web form in the dashboard.
type BankingService struct {
	c *Client

	Accounts     *BankingAccountsService
	Anomalies    *BankingAnomaliesService
	Expected     *BankingExpectedService
	Matches      *BankingMatchesService
	Matching     *BankingMatchingService
	Payments     *BankingPaymentsService
	Receipts     *BankingReceiptsService
	Rules        *BankingRulesService
	Settings     *BankingSettingsService
	Transactions *BankingTransactionsService
}

func newBankingService(c *Client) *BankingService {
	return &BankingService{c: c, Accounts: newBankingAccountsService(c), Anomalies: newBankingAnomaliesService(c), Expected: newBankingExpectedService(c), Matches: newBankingMatchesService(c), Matching: newBankingMatchingService(c), Payments: newBankingPaymentsService(c), Receipts: newBankingReceiptsService(c), Rules: newBankingRulesService(c), Settings: newBankingSettingsService(c), Transactions: newBankingTransactionsService(c)}
}

// Forecast — The liquidity forecast: today's balances moved by open invoices, recurring invoices, open purchases and recurring payments — a row per day.
//
// Only accounts the caller holds the `balance` right on; `available: false` when that is none.
//
// GET /api/banking/forecast — Token permission: read:banking.
func (r *BankingService) Forecast(ctx context.Context, query *BankingForecastQuery, opts ...CallOption) (*BankingForecastResponse, error) {
	var out *BankingForecastResponse
	err := r.c.call(ctx, "GET", "/api/banking/forecast", query, nil, false, "", &out, opts)
	return out, err
}

// ImportPreview — Tells what a statement file is before importing it: its format and, for a CSV, its header, a few rows and a recognised bank's mapping.
//
// Nothing is stored.
//
// POST /api/banking/import-preview — Token permission: write:banking.
func (r *BankingService) ImportPreview(ctx context.Context, body *BankingImportPreviewBody, query *BankingImportPreviewQuery, opts ...CallOption) (*BankingImportPreviewResponse, error) {
	var out *BankingImportPreviewResponse
	err := r.c.call(ctx, "POST", "/api/banking/import-preview", query, body, true, "", &out, opts)
	return out, err
}

// Reports — The Banking hub's numbers as rows: balances, balance history, in and out by month, by category, top counterparts, the forecast.
//
// GET /api/banking/reports — Token permission: read:banking.
func (r *BankingService) Reports(ctx context.Context, query *BankingReportsQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/banking/reports", query, nil, false, "rows", &out, opts)
	return out, err
}

// BankingAccountsService — banking.accounts
type BankingAccountsService struct {
	c *Client

	Grants  *BankingAccountsGrantsService
	Imports *BankingAccountsImportsService
	Ledger  *BankingAccountsLedgerService
}

func newBankingAccountsService(c *Client) *BankingAccountsService {
	return &BankingAccountsService{c: c, Grants: newBankingAccountsGrantsService(c), Imports: newBankingAccountsImportsService(c), Ledger: newBankingAccountsLedgerService(c)}
}

// Create — Creates an account whose statements are imported as files.
//
// Needs `banking:connect`; counts against the plan's bank account limit. Visibility: the request's, else Settings → Banking's default.
//
// POST /api/banking/accounts — Token permission: write:banking.
func (r *BankingAccountsService) Create(ctx context.Context, body *BankingAccountsCreateBody, query *BankingAccountsCreateQuery, opts ...CallOption) (*BankAccount, error) {
	var out *BankAccount
	err := r.c.call(ctx, "POST", "/api/banking/accounts", query, body, false, "account", &out, opts)
	return out, err
}

// Delete — Deletes a bank account with nothing matched or booked.
//
// 409 `bankingAccountHasBooked` otherwise — archive it instead.
//
// DELETE /api/banking/accounts/{accountId} — Token permission: write:banking.
//
// accountID: The bank account.
func (r *BankingAccountsService) Delete(ctx context.Context, accountID string, query *BankingAccountsDeleteQuery, opts ...CallOption) (*BankingAccountsDeleteResponse, error) {
	var out *BankingAccountsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/banking/accounts/"+pathSeg(accountID), query, nil, false, "", &out, opts)
	return out, err
}

// Get — One bank account as the caller may see it.
//
// GET /api/banking/accounts/{accountId} — Token permission: read:banking.
//
// accountID: The bank account.
func (r *BankingAccountsService) Get(ctx context.Context, accountID string, query *BankingAccountsGetQuery, opts ...CallOption) (*BankAccount, error) {
	var out *BankAccount
	err := r.c.call(ctx, "GET", "/api/banking/accounts/"+pathSeg(accountID), query, nil, false, "account", &out, opts)
	return out, err
}

// List — Lists the bank accounts the caller sees, each with the caller's rights on it.
//
// Needs `banking:read`. The balance is null without the account's `balance` right.
//
// GET /api/banking/accounts — Token permission: read:banking.
func (r *BankingAccountsService) List(ctx context.Context, query *BankingAccountsListQuery, opts ...CallOption) ([]BankAccount, error) {
	var out []BankAccount
	err := r.c.call(ctx, "GET", "/api/banking/accounts", query, nil, false, "accounts", &out, opts)
	return out, err
}

// Reconciliation — One month of a bank account against the books: the bank's opening, in, out and closing, each mapped company's ledger balance and the difference, and the lines nothing explains yet.
//
// Needs the account's `transactions` right. The books' side only for a holder of `invoicing:export` or `invoicing:bookkeeping`: otherwise `books` is empty and `booksHidden` true.
//
// GET /api/banking/accounts/{accountId}/reconciliation — Token permission: read:banking.
//
// accountID: The bank account.
func (r *BankingAccountsService) Reconciliation(ctx context.Context, accountID string, query *BankingAccountsReconciliationQuery, opts ...CallOption) (*BankingAccountsReconciliationResponse, error) {
	var out *BankingAccountsReconciliationResponse
	err := r.c.call(ctx, "GET", "/api/banking/accounts/"+pathSeg(accountID)+"/reconciliation", query, nil, false, "", &out, opts)
	return out, err
}

// Update — Changes a bank account's name, type, holder, visibility, CSV mapping or low-balance warning, or archives it.
//
// Needs the account's `manage` right (its creator, a grant, or `banking:settings`).
//
// PATCH /api/banking/accounts/{accountId} — Token permission: write:banking.
//
// accountID: The bank account.
func (r *BankingAccountsService) Update(ctx context.Context, accountID string, body *BankingAccountsUpdateBody, query *BankingAccountsUpdateQuery, opts ...CallOption) (*BankAccount, error) {
	var out *BankAccount
	err := r.c.call(ctx, "PATCH", "/api/banking/accounts/"+pathSeg(accountID), query, body, false, "account", &out, opts)
	return out, err
}

// BankingAccountsGrantsService — banking.accounts.grants
type BankingAccountsGrantsService struct {
	c *Client
}

func newBankingAccountsGrantsService(c *Client) *BankingAccountsGrantsService {
	return &BankingAccountsGrantsService{c: c}
}

// List — Who holds which right on a bank account.
//
// Needs the account's `manage` right.
//
// GET /api/banking/accounts/{accountId}/grants — Token permission: read:banking.
//
// accountID: The bank account.
func (r *BankingAccountsGrantsService) List(ctx context.Context, accountID string, query *BankingAccountsGrantsListQuery, opts ...CallOption) ([]BankGrant, error) {
	var out []BankGrant
	err := r.c.call(ctx, "GET", "/api/banking/accounts/"+pathSeg(accountID)+"/grants", query, nil, false, "grants", &out, opts)
	return out, err
}

// Remove — Takes a grantee's rights on a bank account away.
//
// DELETE /api/banking/accounts/{accountId}/grants/{grantId} — Token permission: write:banking.
//
// accountID: The bank account.
// grantID: The grant.
func (r *BankingAccountsGrantsService) Remove(ctx context.Context, accountID string, grantID string, query *BankingAccountsGrantsRemoveQuery, opts ...CallOption) (*BankingAccountsGrantsRemoveResponse, error) {
	var out *BankingAccountsGrantsRemoveResponse
	err := r.c.call(ctx, "DELETE", "/api/banking/accounts/"+pathSeg(accountID)+"/grants/"+pathSeg(grantID), query, nil, false, "", &out, opts)
	return out, err
}

// Set — Gives a person or a role rights on a bank account, or changes theirs.
//
// Needs the account's `manage` right. One of `userId` and `roleId`. `pay` may carry a limit per order.
//
// POST /api/banking/accounts/{accountId}/grants — Token permission: write:banking.
//
// accountID: The bank account.
func (r *BankingAccountsGrantsService) Set(ctx context.Context, accountID string, body *BankingAccountsGrantsSetBody, query *BankingAccountsGrantsSetQuery, opts ...CallOption) (*BankingAccountsGrantsSetResponse, error) {
	var out *BankingAccountsGrantsSetResponse
	err := r.c.call(ctx, "POST", "/api/banking/accounts/"+pathSeg(accountID)+"/grants", query, body, false, "", &out, opts)
	return out, err
}

// BankingAccountsImportsService — banking.accounts.imports
type BankingAccountsImportsService struct {
	c *Client
}

func newBankingAccountsImportsService(c *Client) *BankingAccountsImportsService {
	return &BankingAccountsImportsService{c: c}
}

// List — The statement files imported into a bank account.
//
// Needs the account's `transactions` right.
//
// GET /api/banking/accounts/{accountId}/imports — Token permission: read:banking.
//
// accountID: The bank account.
func (r *BankingAccountsImportsService) List(ctx context.Context, accountID string, query *BankingAccountsImportsListQuery, opts ...CallOption) ([]BankImport, error) {
	var out []BankImport
	err := r.c.call(ctx, "GET", "/api/banking/accounts/"+pathSeg(accountID)+"/imports", query, nil, false, "imports", &out, opts)
	return out, err
}

// Upload — Imports a CAMT.053, MT940 or CSV statement file into a bank account.
//
// Needs the account's `manage` right. Read by its bytes, 20 MB at most; the original is kept unchangeable with its SHA-256, and the same lines are never imported twice. A CSV needs a mapping: `csvMapping` (JSON), or the one saved on the account.
//
// POST /api/banking/accounts/{accountId}/imports — Token permission: write:banking.
//
// accountID: The bank account.
func (r *BankingAccountsImportsService) Upload(ctx context.Context, accountID string, body *BankingAccountsImportsUploadBody, query *BankingAccountsImportsUploadQuery, opts ...CallOption) (*BankingAccountsImportsUploadResponse, error) {
	var out *BankingAccountsImportsUploadResponse
	err := r.c.call(ctx, "POST", "/api/banking/accounts/"+pathSeg(accountID)+"/imports", query, body, true, "", &out, opts)
	return out, err
}

// BankingAccountsLedgerService — banking.accounts.ledger
type BankingAccountsLedgerService struct {
	c *Client
}

func newBankingAccountsLedgerService(c *Client) *BankingAccountsLedgerService {
	return &BankingAccountsLedgerService{c: c}
}

// Get — Which account of each company's books a bank account is posted to.
//
// GET /api/banking/accounts/{accountId}/ledger — Token permission: read:banking.
//
// accountID: The bank account.
func (r *BankingAccountsLedgerService) Get(ctx context.Context, accountID string, query *BankingAccountsLedgerGetQuery, opts ...CallOption) ([]BankingAccountsLedgerGetResponseMappingsItem, error) {
	var out []BankingAccountsLedgerGetResponseMappingsItem
	err := r.c.call(ctx, "GET", "/api/banking/accounts/"+pathSeg(accountID)+"/ledger", query, nil, false, "mappings", &out, opts)
	return out, err
}

// Remove — Removes a bank account's mapping to a company's books; the chart account stays.
//
// DELETE /api/banking/accounts/{accountId}/ledger — Token permission: write:banking.
//
// accountID: The bank account.
func (r *BankingAccountsLedgerService) Remove(ctx context.Context, accountID string, query *BankingAccountsLedgerRemoveQuery, opts ...CallOption) (*BankingAccountsLedgerRemoveResponse, error) {
	var out *BankingAccountsLedgerRemoveResponse
	err := r.c.call(ctx, "DELETE", "/api/banking/accounts/"+pathSeg(accountID)+"/ledger", query, nil, false, "", &out, opts)
	return out, err
}

// Set — Maps a bank account to a company's books: to an existing bank account of the chart, or to a new one made for it.
//
// Needs the account's `manage` right and `invoicing:bookkeeping`. Without `ledgerAccount` its own chart account is created (SKR03 1201 ff., SKR04 1801 ff.).
//
// PUT /api/banking/accounts/{accountId}/ledger — Token permission: write:banking.
//
// accountID: The bank account.
func (r *BankingAccountsLedgerService) Set(ctx context.Context, accountID string, body *BankingAccountsLedgerSetBody, query *BankingAccountsLedgerSetQuery, opts ...CallOption) (*BankingAccountsLedgerSetResponse, error) {
	var out *BankingAccountsLedgerSetResponse
	err := r.c.call(ctx, "PUT", "/api/banking/accounts/"+pathSeg(accountID)+"/ledger", query, body, false, "", &out, opts)
	return out, err
}

// BankingAnomaliesService — banking.anomalies
type BankingAnomaliesService struct {
	c *Client
}

func newBankingAnomaliesService(c *Client) *BankingAnomaliesService {
	return &BankingAnomaliesService{c: c}
}

// Review — Dismisses an anomaly of a bank line (a double debit, an unusual amount, a new payee), or restores it.
//
// Needs the `match` right on the line's account. A dismissed anomaly is not raised again for the same line and kind.
//
// POST /api/banking/anomalies/{anomalyId} — Token permission: write:banking.
//
// anomalyID: The anomaly.
// body: one of *BankingAnomaliesReviewBody, *BankingAnomaliesReviewBody2.
func (r *BankingAnomaliesService) Review(ctx context.Context, anomalyID string, body any, query *BankingAnomaliesReviewQuery, opts ...CallOption) (*BankAnomaly, error) {
	var out *BankAnomaly
	err := r.c.call(ctx, "POST", "/api/banking/anomalies/"+pathSeg(anomalyID), query, body, false, "anomaly", &out, opts)
	return out, err
}

// BankingExpectedService — banking.expected
type BankingExpectedService struct {
	c *Client

	Plans *BankingExpectedPlansService
}

func newBankingExpectedService(c *Client) *BankingExpectedService {
	return &BankingExpectedService{c: c, Plans: newBankingExpectedPlansService(c)}
}

// List — Expected payments: what the payment plans expect, and whether it came, with totals per currency.
//
// Needs expected payments on the plan. Only rows whose plan's account lines the caller reads and whose record they may see.
//
// GET /api/banking/expected — Token permission: read:banking.
func (r *BankingExpectedService) List(ctx context.Context, query *BankingExpectedListQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/banking/expected", query, nil, false, "items", &out, opts)
	return out, err
}

// Record — A database record's payment account: every expected payment of the plans the caller may see, the bank lines that paid them, and totals per currency.
//
// Needs expected payments on the plan. 404 when the caller may not see the record. `account` is null when the record's database has no payment plan the caller may see (with `active=1`, none switched on).
//
// GET /api/banking/expected/record — Token permission: read:banking.
func (r *BankingExpectedService) Record(ctx context.Context, query *BankingExpectedRecordQuery, opts ...CallOption) (*BankRecordPaymentAccount, error) {
	var out *BankRecordPaymentAccount
	err := r.c.call(ctx, "GET", "/api/banking/expected/record", query, nil, false, "account", &out, opts)
	return out, err
}

// Review — Waives an expected payment (a reduction, a write-off), or takes the waiver back.
//
// Needs `banking:match`, the `match` right on one of the plan's accounts and the record visible.
//
// POST /api/banking/expected/{expectedId} — Token permission: write:banking.
//
// expectedID: The expected payment.
// body: one of *BankingExpectedReviewBody, *BankingExpectedReviewBody2.
func (r *BankingExpectedService) Review(ctx context.Context, expectedID string, body any, query *BankingExpectedReviewQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/banking/expected/"+pathSeg(expectedID), query, body, false, "expected", &out, opts)
	return out, err
}

// BankingExpectedPlansService — banking.expected.plans
type BankingExpectedPlansService struct {
	c *Client
}

func newBankingExpectedPlansService(c *Client) *BankingExpectedPlansService {
	return &BankingExpectedPlansService{c: c}
}

// Create — Creates a payment plan on a database; its expected payments are generated at once.
//
// Needs `banking:match`, edit access to the database, the `match` right on every account it names, and expected payments on the plan.
//
// POST /api/banking/expected/plans — Token permission: write:banking.
func (r *BankingExpectedPlansService) Create(ctx context.Context, body *BankPaymentPlanInput, query *BankingExpectedPlansCreateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/banking/expected/plans", query, body, false, "plan", &out, opts)
	return out, err
}

// Delete — Deletes a payment plan, or archives it when anything was ever paid or matched.
//
// DELETE /api/banking/expected/plans/{planId} — Token permission: write:banking.
//
// planID: The payment plan.
func (r *BankingExpectedPlansService) Delete(ctx context.Context, planID string, query *BankingExpectedPlansDeleteQuery, opts ...CallOption) (*BankingExpectedPlansDeleteResponse, error) {
	var out *BankingExpectedPlansDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/banking/expected/plans/"+pathSeg(planID), query, nil, false, "", &out, opts)
	return out, err
}

// Generate — Generates a payment plan's expected payments now rather than at the next hourly run.
//
// Needs the right to change the plan.
//
// POST /api/banking/expected/plans/{planId}/generate — Token permission: write:banking.
//
// planID: The payment plan.
func (r *BankingExpectedPlansService) Generate(ctx context.Context, planID string, query *BankingExpectedPlansGenerateQuery, opts ...CallOption) (*BankingExpectedPlansGenerateResponseResult, error) {
	var out *BankingExpectedPlansGenerateResponseResult
	err := r.c.call(ctx, "POST", "/api/banking/expected/plans/"+pathSeg(planID)+"/generate", query, nil, false, "result", &out, opts)
	return out, err
}

// Get — One payment plan.
//
// GET /api/banking/expected/plans/{planId} — Token permission: read:banking.
//
// planID: The payment plan.
func (r *BankingExpectedPlansService) Get(ctx context.Context, planID string, query *BankingExpectedPlansGetQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/banking/expected/plans/"+pathSeg(planID), query, nil, false, "plan", &out, opts)
	return out, err
}

// List — Lists the payment plans the caller may see, newest first.
//
// GET /api/banking/expected/plans — Token permission: read:banking.
func (r *BankingExpectedPlansService) List(ctx context.Context, query *BankingExpectedPlansListQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/banking/expected/plans", query, nil, false, "plans", &out, opts)
	return out, err
}

// Update — Changes a payment plan; the rows nothing touched yet are planned again.
//
// PATCH /api/banking/expected/plans/{planId} — Token permission: write:banking.
//
// planID: The payment plan.
func (r *BankingExpectedPlansService) Update(ctx context.Context, planID string, body *BankingExpectedPlansUpdateBody, query *BankingExpectedPlansUpdateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "PATCH", "/api/banking/expected/plans/"+pathSeg(planID), query, body, false, "plan", &out, opts)
	return out, err
}

// BankingMatchesService — banking.matches
type BankingMatchesService struct {
	c *Client
}

func newBankingMatchesService(c *Client) *BankingMatchesService {
	return &BankingMatchesService{c: c}
}

// Undo — Undoes a confirmed match: the payment it recorded is removed (its entry reversed in the books), a booking is reversed.
//
// Needs `banking:match` and the account's `match` right; a match to a document needs `invoicing:payments` and the document visible, a booking `invoicing:bookkeeping`. A payment the match only took over (a confirmed payment run's) stays.
//
// DELETE /api/banking/matches/{matchId} — Token permission: write:banking.
//
// matchID: The match.
func (r *BankingMatchesService) Undo(ctx context.Context, matchID string, query *BankingMatchesUndoQuery, opts ...CallOption) (*BankingMatchesUndoResponseTransaction, error) {
	var out *BankingMatchesUndoResponseTransaction
	err := r.c.call(ctx, "DELETE", "/api/banking/matches/"+pathSeg(matchID), query, nil, false, "transaction", &out, opts)
	return out, err
}

// BankingMatchingService — banking.matching
type BankingMatchingService struct {
	c *Client
}

func newBankingMatchingService(c *Client) *BankingMatchingService {
	return &BankingMatchingService{c: c}
}

// Count — How many lines wait for a look in the accounts whose lines the caller may read.
//
// GET /api/banking/matching/count — Token permission: read:banking.
func (r *BankingMatchingService) Count(ctx context.Context, query *BankingMatchingCountQuery, opts ...CallOption) (*BankingMatchingCountResponse, error) {
	var out *BankingMatchingCountResponse
	err := r.c.call(ctx, "GET", "/api/banking/matching/count", query, nil, false, "", &out, opts)
	return out, err
}

// Queue — The review queue: open lines with what the engine suggested, suggestions first.
//
// Documents the caller may not see are left out.
//
// GET /api/banking/matching/queue — Token permission: read:banking.
func (r *BankingMatchingService) Queue(ctx context.Context, query *BankingMatchingQueueQuery, opts ...CallOption) (*BankingMatchingQueueResponse, error) {
	var out *BankingMatchingQueueResponse
	err := r.c.call(ctx, "GET", "/api/banking/matching/queue", query, nil, false, "", &out, opts)
	return out, err
}

// Run — Runs matching again over the open lines of one account, or of every account the caller may match.
//
// Needs `banking:match`. For after new invoices were issued or a rule was added; auto-matching follows Settings → Banking.
//
// POST /api/banking/matching/run — Token permission: write:banking.
func (r *BankingMatchingService) Run(ctx context.Context, body *BankingMatchingRunBody, query *BankingMatchingRunQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/banking/matching/run", query, body, false, "result", &out, opts)
	return out, err
}

// BankingPaymentsService — banking.payments
type BankingPaymentsService struct {
	c *Client

	Settings *BankingPaymentsSettingsService
}

func newBankingPaymentsService(c *Client) *BankingPaymentsService {
	return &BankingPaymentsService{c: c, Settings: newBankingPaymentsSettingsService(c)}
}

// Act — Submits, approves, rejects, sends or cancels a payment order.
//
// Needs `banking:pay`.
//
// - `submit`: to approval (approved at once when the rule needs none).
// - `approve` / `reject`: four eyes — never the order's creator; `pay` on the account with a limit at least the total. A rejection needs a note.
// - `send`: to the bank; answers `webFormUrl`, which a person must open to confirm with their bank. The order must equal what was approved (409 `bankingPaymentChanged`).
// - `cancel`: before it was sent, or when the bank's outcome stayed unclear.
//
// POST /api/banking/payments/{orderId} — Token permission: write:banking.
//
// orderID: The payment order.
// body: one of *BankingPaymentsActBody, *BankingPaymentsActBody2, *BankingPaymentsActBody3, *BankingPaymentsActBody4, *BankingPaymentsActBody5.
func (r *BankingPaymentsService) Act(ctx context.Context, orderID string, body any, query *BankingPaymentsActQuery, opts ...CallOption) (*BankingPaymentsActResponse, error) {
	var out *BankingPaymentsActResponse
	err := r.c.call(ctx, "POST", "/api/banking/payments/"+pathSeg(orderID), query, body, false, "", &out, opts)
	return out, err
}

// Create — Creates a payment order — one transfer or several — from a connected account, and sends it to approval.
//
// Needs `banking:pay` and the account's `pay` right with a limit at least the total (403 `bankingPaymentOverLimit`), and payment orders on the plan. 422 `bankingPaymentInvalid` with `problems`. With `submit` (the default) it goes to approval at once, or is approved when the organisation's rule asks for none. Sending it to the bank needs a person (see the `send` action).
//
// POST /api/banking/payments — Token permission: write:banking.
func (r *BankingPaymentsService) Create(ctx context.Context, body *BankingPaymentsCreateBody, query *BankingPaymentsCreateQuery, opts ...CallOption) (*BankPaymentOrder, error) {
	var out *BankPaymentOrder
	err := r.c.call(ctx, "POST", "/api/banking/payments", query, body, false, "order", &out, opts)
	return out, err
}

// FromPurchases — Creates a payment order paying purchase invoices: an open payment run, or ticked invoices.
//
// Needs `banking:pay` with the account's `pay` right and limit, and `invoicing:read` with `invoicing:payments`. Send `runId`, or `issuerId` with `purchaseIds`.
//
// POST /api/banking/payments/from-purchases — Token permission: write:banking.
func (r *BankingPaymentsService) FromPurchases(ctx context.Context, body *BankingPaymentsFromPurchasesBody, query *BankingPaymentsFromPurchasesQuery, opts ...CallOption) (*BankPaymentOrder, error) {
	var out *BankPaymentOrder
	err := r.c.call(ctx, "POST", "/api/banking/payments/from-purchases", query, body, false, "order", &out, opts)
	return out, err
}

// Get — One payment order with its items, approvals and history.
//
// A submitted order is read back from the bank first.
//
// GET /api/banking/payments/{orderId} — Token permission: read:banking.
//
// orderID: The payment order.
func (r *BankingPaymentsService) Get(ctx context.Context, orderID string, query *BankingPaymentsGetQuery, opts ...CallOption) (*BankPaymentOrder, error) {
	var out *BankPaymentOrder
	err := r.c.call(ctx, "GET", "/api/banking/payments/"+pathSeg(orderID), query, nil, false, "order", &out, opts)
	return out, err
}

// List — Lists payment orders: the caller's own and those of accounts where they hold `pay` or `transactions`.
//
// Needs payment orders on the plan.
//
// GET /api/banking/payments — Token permission: read:banking.
func (r *BankingPaymentsService) List(ctx context.Context, query *BankingPaymentsListQuery, opts ...CallOption) ([]BankPaymentOrder, error) {
	var out []BankPaymentOrder
	err := r.c.call(ctx, "GET", "/api/banking/payments", query, nil, false, "orders", &out, opts)
	return out, err
}

// Options — What the caller may pay from: the connected accounts where they hold `pay` with their limit, the four-eyes rule, and with `accountId` what that account can do at the bank.
//
// Needs payment orders on the plan. `accounts` is empty without `banking:pay`. `capabilities` (several transfers in one order, instant) is null when the bank provider does not answer. `paused`: the platform has paused calls to the bank provider.
//
// GET /api/banking/payments/options — Token permission: read:banking.
func (r *BankingPaymentsService) Options(ctx context.Context, query *BankingPaymentsOptionsQuery, opts ...CallOption) (*BankingPaymentsOptionsResponse, error) {
	var out *BankingPaymentsOptionsResponse
	err := r.c.call(ctx, "GET", "/api/banking/payments/options", query, nil, false, "", &out, opts)
	return out, err
}

// Update — Changes a payment order that is not sent yet; a waiting or approved one goes back to approval.
//
// Needs `banking:pay`: its creator, or a `pay` holder within their limit.
//
// PATCH /api/banking/payments/{orderId} — Token permission: write:banking.
//
// orderID: The payment order.
func (r *BankingPaymentsService) Update(ctx context.Context, orderID string, body *BankingPaymentsUpdateBody, query *BankingPaymentsUpdateQuery, opts ...CallOption) (*BankPaymentOrder, error) {
	var out *BankPaymentOrder
	err := r.c.call(ctx, "PATCH", "/api/banking/payments/"+pathSeg(orderID), query, body, false, "order", &out, opts)
	return out, err
}

// BankingPaymentsSettingsService — banking.payments.settings
type BankingPaymentsSettingsService struct {
	c *Client
}

func newBankingPaymentsSettingsService(c *Client) *BankingPaymentsSettingsService {
	return &BankingPaymentsSettingsService{c: c}
}

// Get — Settings → Banking → Payments: the organisation's four-eyes rule.
//
// GET /api/banking/payments/settings — Token permission: read:banking.
func (r *BankingPaymentsSettingsService) Get(ctx context.Context, query *BankingPaymentsSettingsGetQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/banking/payments/settings", query, nil, false, "settings", &out, opts)
	return out, err
}

// Update — Changes the four-eyes rule for payment orders.
//
// Needs `banking:settings`. Applies to orders sent to approval from then on.
//
// PUT /api/banking/payments/settings — Token permission: write:banking.
func (r *BankingPaymentsSettingsService) Update(ctx context.Context, body *BankingPaymentsSettingsUpdateBody, query *BankingPaymentsSettingsUpdateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "PUT", "/api/banking/payments/settings", query, body, false, "settings", &out, opts)
	return out, err
}

// BankingReceiptsService — banking.receipts
type BankingReceiptsService struct {
	c *Client

	Settings *BankingReceiptsSettingsService
}

func newBankingReceiptsService(c *Client) *BankingReceiptsService {
	return &BankingReceiptsService{c: c, Settings: newBankingReceiptsSettingsService(c)}
}

// Book — Books with the receipt: creates a purchase invoice draft from it, prefilled with the bank line.
//
// Needs the account's `match` right, and in Invoicing `invoicing:purchases` and `invoicing:create` with the books on (409 `bankingReceiptBooksOff`). `aiAvailable` tells whether reading the draft with AI is offered.
//
// POST /api/banking/receipts/{receiptId}/book — Token permission: write:banking.
//
// receiptID: The receipt link.
func (r *BankingReceiptsService) Book(ctx context.Context, receiptID string, query *BankingReceiptsBookQuery, opts ...CallOption) (*BankingReceiptsBookResponse, error) {
	var out *BankingReceiptsBookResponse
	err := r.c.call(ctx, "POST", "/api/banking/receipts/"+pathSeg(receiptID)+"/book", query, nil, false, "", &out, opts)
	return out, err
}

// Decide — Accepts or rejects a suggested receipt.
//
// Needs `banking:match` and the `match` right on the line's account.
//
// PATCH /api/banking/receipts/{receiptId} — Token permission: write:banking.
//
// receiptID: The receipt link.
func (r *BankingReceiptsService) Decide(ctx context.Context, receiptID string, body *BankingReceiptsDecideBody, query *BankingReceiptsDecideQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "PATCH", "/api/banking/receipts/"+pathSeg(receiptID), query, body, false, "receipt", &out, opts)
	return out, err
}

// File — Downloads a receipt's file: the uploaded file, or the linked document's original.
//
// Needs the account's `transactions` right; a linked document only when the caller may see it in Invoicing (404 otherwise). Inline for PDF, JPEG and PNG, an attachment otherwise.
//
// GET /api/banking/receipts/{receiptId}/file — Token permission: read:banking.
//
// receiptID: The receipt link.
func (r *BankingReceiptsService) File(ctx context.Context, receiptID string, query *BankingReceiptsFileQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/banking/receipts/"+pathSeg(receiptID)+"/file", query, nil, false, "", &out, opts)
	return out, err
}

// Missing — The outgoing lines that miss their receipt, newest first, with suggestions.
//
// Which lines ask for a receipt is Settings → Banking → Receipts.
//
// GET /api/banking/receipts — Token permission: read:banking.
func (r *BankingReceiptsService) Missing(ctx context.Context, query *BankingReceiptsMissingQuery, opts ...CallOption) (*BankingReceiptsMissingResponse, error) {
	var out *BankingReceiptsMissingResponse
	err := r.c.call(ctx, "GET", "/api/banking/receipts", query, nil, false, "", &out, opts)
	return out, err
}

// Run — Finds receipts: scores every line that misses one against the received documents and purchase invoices again — exact and unique is linked, the rest suggested.
//
// Needs `banking:match`.
//
// POST /api/banking/receipts/run — Token permission: write:banking.
func (r *BankingReceiptsService) Run(ctx context.Context, query *BankingReceiptsRunQuery, opts ...CallOption) (*BankingReceiptsRunResponseResult, error) {
	var out *BankingReceiptsRunResponseResult
	err := r.c.call(ctx, "POST", "/api/banking/receipts/run", query, nil, false, "result", &out, opts)
	return out, err
}

// Unlink — Removes a linked receipt; an uploaded file stays stored.
//
// DELETE /api/banking/receipts/{receiptId} — Token permission: write:banking.
//
// receiptID: The receipt link.
func (r *BankingReceiptsService) Unlink(ctx context.Context, receiptID string, query *BankingReceiptsUnlinkQuery, opts ...CallOption) (*BankingReceiptsUnlinkResponse, error) {
	var out *BankingReceiptsUnlinkResponse
	err := r.c.call(ctx, "DELETE", "/api/banking/receipts/"+pathSeg(receiptID), query, nil, false, "", &out, opts)
	return out, err
}

// BankingReceiptsSettingsService — banking.receipts.settings
type BankingReceiptsSettingsService struct {
	c *Client
}

func newBankingReceiptsSettingsService(c *Client) *BankingReceiptsSettingsService {
	return &BankingReceiptsSettingsService{c: c}
}

// Get — Settings → Banking → Receipts: which outgoing lines ask for a receipt.
//
// GET /api/banking/receipts/settings — Token permission: read:banking.
func (r *BankingReceiptsSettingsService) Get(ctx context.Context, query *BankingReceiptsSettingsGetQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/banking/receipts/settings", query, nil, false, "settings", &out, opts)
	return out, err
}

// Update — Changes which outgoing lines ask for a receipt.
//
// Needs `banking:settings`.
//
// PUT /api/banking/receipts/settings — Token permission: write:banking.
func (r *BankingReceiptsSettingsService) Update(ctx context.Context, body *BankingReceiptsSettingsUpdateBody, query *BankingReceiptsSettingsUpdateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "PUT", "/api/banking/receipts/settings", query, body, false, "settings", &out, opts)
	return out, err
}

// BankingRulesService — banking.rules
type BankingRulesService struct {
	c *Client
}

func newBankingRulesService(c *Client) *BankingRulesService {
	return &BankingRulesService{c: c}
}

// Create — Creates a rule: lines that meet its conditions are ignored, categorised, booked or tied to a customer or supplier.
//
// Needs `banking:match`; a rule naming an account needs the `match` right on it. A `ledger` action names a company and an account of its chart that is not money, personal or inactive.
//
// POST /api/banking/rules — Token permission: write:banking.
func (r *BankingRulesService) Create(ctx context.Context, body *BankingRulesCreateBody, query *BankingRulesCreateQuery, opts ...CallOption) (*BankRule, error) {
	var out *BankRule
	err := r.c.call(ctx, "POST", "/api/banking/rules", query, body, false, "rule", &out, opts)
	return out, err
}

// Delete — Deletes a banking rule.
//
// DELETE /api/banking/rules/{ruleId} — Token permission: write:banking.
//
// ruleID: The rule.
func (r *BankingRulesService) Delete(ctx context.Context, ruleID string, query *BankingRulesDeleteQuery, opts ...CallOption) (*BankingRulesDeleteResponse, error) {
	var out *BankingRulesDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/banking/rules/"+pathSeg(ruleID), query, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the organisation's banking rules, manual and learned, by priority.
//
// GET /api/banking/rules — Token permission: read:banking.
func (r *BankingRulesService) List(ctx context.Context, query *BankingRulesListQuery, opts ...CallOption) ([]BankRule, error) {
	var out []BankRule
	err := r.c.call(ctx, "GET", "/api/banking/rules", query, nil, false, "rules", &out, opts)
	return out, err
}

// Update — Changes a banking rule; changing a learned rule's conditions or action makes it a manual one.
//
// PATCH /api/banking/rules/{ruleId} — Token permission: write:banking.
//
// ruleID: The rule.
func (r *BankingRulesService) Update(ctx context.Context, ruleID string, body *BankingRulesUpdateBody, query *BankingRulesUpdateQuery, opts ...CallOption) (*BankRule, error) {
	var out *BankRule
	err := r.c.call(ctx, "PATCH", "/api/banking/rules/"+pathSeg(ruleID), query, body, false, "rule", &out, opts)
	return out, err
}

// BankingSettingsService — banking.settings
type BankingSettingsService struct {
	c *Client
}

func newBankingSettingsService(c *Client) *BankingSettingsService {
	return &BankingSettingsService{c: c}
}

// Get — Settings → Banking.
//
// GET /api/banking/settings — Token permission: read:banking.
func (r *BankingSettingsService) Get(ctx context.Context, query *BankingSettingsGetQuery, opts ...CallOption) (*BankSettings, error) {
	var out *BankSettings
	err := r.c.call(ctx, "GET", "/api/banking/settings", query, nil, false, "settings", &out, opts)
	return out, err
}

// Update — Changes Settings → Banking: auto-matching, the suggestion threshold, new accounts' visibility, consent reminders, AI on bank data.
//
// Needs `banking:settings`.
//
// PUT /api/banking/settings — Token permission: write:banking.
func (r *BankingSettingsService) Update(ctx context.Context, body *BankingSettingsUpdateBody, query *BankingSettingsUpdateQuery, opts ...CallOption) (*BankSettings, error) {
	var out *BankSettings
	err := r.c.call(ctx, "PUT", "/api/banking/settings", query, body, false, "settings", &out, opts)
	return out, err
}

// BankingTransactionsService — banking.transactions
type BankingTransactionsService struct {
	c *Client
}

func newBankingTransactionsService(c *Client) *BankingTransactionsService {
	return &BankingTransactionsService{c: c}
}

// AddReceipt — Links a received document or purchase invoice to a bank line as its receipt, or says it needs none.
//
// Needs the account's `match` right. A file can be uploaded instead as multipart form data in a `file` field (PDF, JPEG, PNG or HEIC, 20 MB at most).
//
// POST /api/banking/transactions/{transactionId}/receipts — Token permission: write:banking.
//
// transactionID: The bank line.
// body: one of *BankingTransactionsAddReceiptBody, *BankingTransactionsAddReceiptBody2.
func (r *BankingTransactionsService) AddReceipt(ctx context.Context, transactionID string, body any, query *BankingTransactionsAddReceiptQuery, opts ...CallOption) (*BankingTransactionsAddReceiptResponse, error) {
	var out *BankingTransactionsAddReceiptResponse
	err := r.c.call(ctx, "POST", "/api/banking/transactions/"+pathSeg(transactionID)+"/receipts", query, body, false, "", &out, opts)
	return out, err
}

// Annotate — Sets a bank line's notes and category; the bank's own data never changes.
//
// Needs the account's `match` right.
//
// PATCH /api/banking/transactions/{transactionId} — Token permission: write:banking.
//
// transactionID: The bank line.
func (r *BankingTransactionsService) Annotate(ctx context.Context, transactionID string, body *BankingTransactionsAnnotateBody, query *BankingTransactionsAnnotateQuery, opts ...CallOption) (*BankTransaction, error) {
	var out *BankTransaction
	err := r.c.call(ctx, "PATCH", "/api/banking/transactions/"+pathSeg(transactionID), query, body, false, "transaction", &out, opts)
	return out, err
}

// Export — Downloads the lines the list shows for the same filters as a CSV or Excel file.
//
// Needs `banking:export`. Newest first, up to a maximum; the header `X-Export-Truncated: 1` says when the filter held more.
//
// GET /api/banking/transactions/export — Token permission: read:banking.
func (r *BankingTransactionsService) Export(ctx context.Context, query *BankingTransactionsExportQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/banking/transactions/export", query, nil, false, "", &out, opts)
	return out, err
}

// Get — One bank line, with its category suggestion and its anomalies.
//
// GET /api/banking/transactions/{transactionId} — Token permission: read:banking.
//
// transactionID: The bank line.
func (r *BankingTransactionsService) Get(ctx context.Context, transactionID string, query *BankingTransactionsGetQuery, opts ...CallOption) (*BankTransaction, error) {
	var out *BankTransaction
	err := r.c.call(ctx, "GET", "/api/banking/transactions/"+pathSeg(transactionID), query, nil, false, "transaction", &out, opts)
	return out, err
}

// List — Lists one page of the lines of the accounts whose lines the caller may read, with the sums in and out.
//
// Needs `banking:read`; only accounts the caller holds the `transactions` right on (naming one without it is 403 `bankingAccountRightMissing`).
//
// GET /api/banking/transactions — Token permission: read:banking.
func (r *BankingTransactionsService) List(ctx context.Context, query *BankingTransactionsListQuery, opts ...CallOption) ([]BankTransaction, error) {
	var out []BankTransaction
	err := r.c.call(ctx, "GET", "/api/banking/transactions", query, nil, false, "items", &out, opts)
	return out, err
}

// Match — Matches a bank line to invoices, purchases, expected payments or bookings — or accepts, rejects, ignores or rechecks its suggestions.
//
// Needs `banking:match` and the account's `match` right; matching documents also needs `invoicing:read` and `invoicing:payments`, bookings `invoicing:bookkeeping`.
//
// - `apply`: several targets split the line; amounts positive, together at most what is left.
// - `accept`: confirm the open suggestions as they stand.
// - `reject`: "not this" — the open suggestions go.
// - `ignore`: the line needs no match (or does again).
// - `recheck`: run the engine on this line again.
//
// POST /api/banking/transactions/{transactionId}/matching — Token permission: write:banking.
//
// transactionID: The bank line.
// body: one of *BankingTransactionsMatchBody, *BankingTransactionsMatchBody2, *BankingTransactionsMatchBody3, *BankingTransactionsMatchBody4, *BankingTransactionsMatchBody5.
func (r *BankingTransactionsService) Match(ctx context.Context, transactionID string, body any, query *BankingTransactionsMatchQuery, opts ...CallOption) (*BankingTransactionsMatchResponse, error) {
	var out *BankingTransactionsMatchResponse
	err := r.c.call(ctx, "POST", "/api/banking/transactions/"+pathSeg(transactionID)+"/matching", query, body, false, "", &out, opts)
	return out, err
}

// MatchCandidates — What a bank line could be matched to — open invoices for money in, purchase invoices still owed for money out, expected payments — in the line's currency, ranked for this line.
//
// Needs `banking:match` and the account's `match` right; documents also need `invoicing:read` and `invoicing:payments`, without which only expected payments are offered. Only what the caller may see.
//
// GET /api/banking/transactions/{transactionId}/matching/candidates — Token permission: read:banking.
//
// transactionID: The bank line.
func (r *BankingTransactionsService) MatchCandidates(ctx context.Context, transactionID string, query *BankingTransactionsMatchCandidatesQuery, opts ...CallOption) ([]BankMatchCandidate, error) {
	var out []BankMatchCandidate
	err := r.c.call(ctx, "GET", "/api/banking/transactions/"+pathSeg(transactionID)+"/matching/candidates", query, nil, false, "candidates", &out, opts)
	return out, err
}

// Matches — A bank line's matches — confirmed, suggested and undone — what is left of it, and whether the caller may match it.
//
// `hidden` counts the matches to documents the caller may not see.
//
// GET /api/banking/transactions/{transactionId}/matching — Token permission: read:banking.
//
// transactionID: The bank line.
func (r *BankingTransactionsService) Matches(ctx context.Context, transactionID string, query *BankingTransactionsMatchesQuery, opts ...CallOption) (*BankingTransactionsMatchesResponse, error) {
	var out *BankingTransactionsMatchesResponse
	err := r.c.call(ctx, "GET", "/api/banking/transactions/"+pathSeg(transactionID)+"/matching", query, nil, false, "", &out, opts)
	return out, err
}

// ReceiptCandidates — Received documents and purchase invoices that could be a bank line's receipt, best first.
//
// Needs the account's `match` right. Without `q`, the documents the receipt rules score for this line; `q` narrows by number or supplier and also finds documents outside the date window. At most 30, only those the caller may see.
//
// GET /api/banking/transactions/{transactionId}/receipts/candidates — Token permission: read:banking.
//
// transactionID: The bank line.
func (r *BankingTransactionsService) ReceiptCandidates(ctx context.Context, transactionID string, query *BankingTransactionsReceiptCandidatesQuery, opts ...CallOption) ([]BankReceiptCandidate, error) {
	var out []BankReceiptCandidate
	err := r.c.call(ctx, "GET", "/api/banking/transactions/"+pathSeg(transactionID)+"/receipts/candidates", query, nil, false, "candidates", &out, opts)
	return out, err
}

// Receipts — A bank line's receipts, linked and suggested, and whether it still misses one.
//
// GET /api/banking/transactions/{transactionId}/receipts — Token permission: read:banking.
//
// transactionID: The bank line.
func (r *BankingTransactionsService) Receipts(ctx context.Context, transactionID string, query *BankingTransactionsReceiptsQuery, opts ...CallOption) (*BankingTransactionsReceiptsResponse, error) {
	var out *BankingTransactionsReceiptsResponse
	err := r.c.call(ctx, "GET", "/api/banking/transactions/"+pathSeg(transactionID)+"/receipts", query, nil, false, "", &out, opts)
	return out, err
}

// Suggest — Suggests a category and ledger account for a bank line, from the counterpart's history and then the platform AI.
//
// Needs `banking:match` and the account's `match` right. The suggestion is kept beside the line and applied by a person, never booked by itself. The AI half needs AI on bank data on the plan and allowed in Settings → Banking. 10 a minute.
//
// POST /api/banking/transactions/{transactionId}/suggest — Token permission: write:banking.
//
// transactionID: The bank line.
func (r *BankingTransactionsService) Suggest(ctx context.Context, transactionID string, query *BankingTransactionsSuggestQuery, opts ...CallOption) (*BankingTransactionsSuggestResponse, error) {
	var out *BankingTransactionsSuggestResponse
	err := r.c.call(ctx, "POST", "/api/banking/transactions/"+pathSeg(transactionID)+"/suggest", query, nil, false, "", &out, opts)
	return out, err
}

// BillingService — Billing: An organization's own plan with the platform: its status, billing address, payment history and the invoices the platform issued to it. Changing or paying for a plan happens in the dashboard, where the payment provider's form can ask the person to confirm.
type BillingService struct {
	c *Client
}

func newBillingService(c *Client) *BillingService {
	return &BillingService{c: c}
}

// DocumentPDF — Downloads the PDF of a document the platform issued to the caller, as archived.
//
// An organization's document also needs `invoices:download` there. 404 for anything else. The header `X-Content-SHA256` carries the stored hash; every download is written to the audit log.
//
// GET /api/billing/documents/{documentId}/pdf
//
// documentID: The document's id, from GET /api/billing/documents.
func (r *BillingService) DocumentPDF(ctx context.Context, documentID string, query *BillingDocumentPdfQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/billing/documents/"+pathSeg(documentID)+"/pdf", query, nil, false, "", &out, opts)
	return out, err
}

// DocumentXml — Downloads the e-invoice XML of a document the platform issued to the caller, as archived.
//
// Same access as the PDF. 404 also when the document was not issued as an e-invoice; 502 `invoicingEinvoiceXmlMissing` when the archive cannot be read.
//
// GET /api/billing/documents/{documentId}/xml
//
// documentID: The document's id, from GET /api/billing/documents.
func (r *BillingService) DocumentXml(ctx context.Context, documentID string, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/billing/documents/"+pathSeg(documentID)+"/xml", nil, nil, false, "", &out, opts)
	return out, err
}

// ExplainDocument — Explains a document the platform issued to the caller in plain language, with the platform AI.
//
// Counts against the buyer organization's AI tokens (a personal purchase: the caller's first organization). 503 `aiNotConfigured` without platform AI, 403 `aiOptedOut` when the person opted out of AI, 429 over the quota. Cached per document and language for 24 hours.
//
// POST /api/billing/documents/{documentId}/explain
//
// documentID: The document's id, from GET /api/billing/documents.
func (r *BillingService) ExplainDocument(ctx context.Context, documentID string, opts ...CallOption) (*BillingExplainDocumentResponse, error) {
	var out *BillingExplainDocumentResponse
	err := r.c.call(ctx, "POST", "/api/billing/documents/"+pathSeg(documentID)+"/explain", nil, nil, false, "", &out, opts)
	return out, err
}

// GetAddress — Returns an organization's billing address and VAT number.
//
// Needs `billing:read` and membership of the organization.
//
// GET /api/billing/address
func (r *BillingService) GetAddress(ctx context.Context, query *BillingGetAddressQuery, opts ...CallOption) (*BillingAddress, error) {
	var out *BillingAddress
	err := r.c.call(ctx, "GET", "/api/billing/address", query, nil, false, "", &out, opts)
	return out, err
}

// GetStatus — Returns the status of an organization's plan.
//
// Needs `billing:read` and membership of the organization. A paid plan is preferred over a free one when both are assigned.
//
// GET /api/billing/status
func (r *BillingService) GetStatus(ctx context.Context, query *BillingGetStatusQuery, opts ...CallOption) (*BillingStatus, error) {
	var out *BillingStatus
	err := r.c.call(ctx, "GET", "/api/billing/status", query, nil, false, "", &out, opts)
	return out, err
}

// ListDocuments — Lists the invoices, credit notes and other documents the platform issued to the caller or the caller's organizations.
//
// An organization's documents are listed where the caller holds `invoices:read` at organization level; personal purchases always. Drafts are never listed.
//
// GET /api/billing/documents
func (r *BillingService) ListDocuments(ctx context.Context, query *BillingListDocumentsQuery, opts ...CallOption) (*BillingListDocumentsResponseData, error) {
	var out *BillingListDocumentsResponseData
	err := r.c.call(ctx, "GET", "/api/billing/documents", query, nil, false, "data", &out, opts)
	return out, err
}

// ListHistory — Lists an organization's billing events, newest first.
//
// Needs `billing:read` and membership of the organization.
//
// GET /api/billing/history
func (r *BillingService) ListHistory(ctx context.Context, query *BillingListHistoryQuery, opts ...CallOption) ([]BillingEvent, error) {
	var out []BillingEvent
	err := r.c.call(ctx, "GET", "/api/billing/history", query, nil, false, "events", &out, opts)
	return out, err
}

// ListMarketplaceInvoices — Lists the invoices marketplace sellers issued to the caller or the caller's organizations.
//
// An organization's invoices are listed where the caller holds `invoices:read` at organization level. 403 `featureDisabledByAdmin` while the marketplace is switched off.
//
// GET /api/billing/documents/marketplace
func (r *BillingService) ListMarketplaceInvoices(ctx context.Context, query *BillingListMarketplaceInvoicesQuery, opts ...CallOption) ([]MarketplaceInvoice, error) {
	var out []MarketplaceInvoice
	err := r.c.call(ctx, "GET", "/api/billing/documents/marketplace", query, nil, false, "data.rows", &out, opts)
	return out, err
}

// MarketplaceInvoicePDF — Downloads the PDF of a marketplace seller's invoice to the caller or the caller's organization.
//
// An organization's invoice also needs `invoices:download` there; 404 otherwise. Every download is written to the audit log.
//
// GET /api/billing/documents/marketplace/{invoiceId}/pdf
//
// invoiceID: The invoice's id, from GET /api/billing/documents/marketplace.
func (r *BillingService) MarketplaceInvoicePDF(ctx context.Context, invoiceID string, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/billing/documents/marketplace/"+pathSeg(invoiceID)+"/pdf", nil, nil, false, "", &out, opts)
	return out, err
}

// UpdateAddress — Saves an organization's billing address and VAT number.
//
// Needs `billing:manage` in that organization. A field left out keeps its value; an empty string clears it. An EU VAT number is checked (400 when malformed), and the details are passed on to the payment provider when the organization is already a customer there (`providerSync`).
//
// PATCH /api/billing/address
func (r *BillingService) UpdateAddress(ctx context.Context, body *UpdateBillingAddressInput, opts ...CallOption) (*BillingUpdateAddressResponse, error) {
	var out *BillingUpdateAddressResponse
	err := r.c.call(ctx, "PATCH", "/api/billing/address", nil, body, false, "", &out, opts)
	return out, err
}

// Usage — Returns an organization's plan with its limits and what is used of each, AI and video usage this month, and whether the caller may change the plan.
//
// Needs membership of the organization. `limits` and `usage` are keyed alike (`max_projects`, `max_storage_mb` …); a limit of `-1` means no limit. Mailboxes, mail storage, calendars, booking pages and organizations are the caller's own, against the best plan among their organizations; the organization's shared mailboxes are `sharedMail`.
//
// GET /api/usage
func (r *BillingService) Usage(ctx context.Context, query *BillingUsageQuery, opts ...CallOption) (*BillingUsageResponse, error) {
	var out *BillingUsageResponse
	err := r.c.call(ctx, "GET", "/api/usage", query, nil, false, "", &out, opts)
	return out, err
}

// BuilderBlocksService — Builder blocks: An organisation's own blocks for the website builder — each with versions of its props schema, default props and styles and an HTML template — beside the platform's blocks, which are read-only. The organisation acted on is the one the request names, else the one the token is bound to; the caller must belong to it.
type BuilderBlocksService struct {
	c *Client

	Versions *BuilderBlocksVersionsService
}

func newBuilderBlocksService(c *Client) *BuilderBlocksService {
	return &BuilderBlocksService{c: c, Versions: newBuilderBlocksVersionsService(c)}
}

// Catalog — Lists every block the editor offers the caller — the platform's and the organisation's own — with its current schema, defaults and whether the plan includes it.
//
// Hidden blocks are left out; a block the plan does not include is listed with `policy.accessible: false`. Answers an `ETag`; with a matching `If-None-Match` it answers 304.
//
// GET /api/builder-blocks/catalog
func (r *BuilderBlocksService) Catalog(ctx context.Context, query *BuilderBlocksCatalogQuery, opts ...CallOption) ([]BuilderBlockCatalogEntry, error) {
	var out []BuilderBlockCatalogEntry
	err := r.c.call(ctx, "GET", "/api/builder-blocks/catalog", query, nil, false, "data.blocks", &out, opts)
	return out, err
}

// Create — Creates a block of the organisation, with its first version.
//
// Needs `component:manage_org_blocks` held in the organisation acted on. The block starts as `beta`. 409 when the organisation already has a block with that slug.
//
// POST /api/builder-blocks
func (r *BuilderBlocksService) Create(ctx context.Context, body *BuilderBlocksCreateBody, query *BuilderBlocksCreateQuery, opts ...CallOption) (string, error) {
	var out string
	err := r.c.call(ctx, "POST", "/api/builder-blocks", query, body, false, "data.id", &out, opts)
	return out, err
}

// Delete — Moves a block of the organisation to the trash.
//
// Needs `component:manage_org_blocks` held in the organisation acted on. It can be restored from the trash until it is deleted for good.
//
// DELETE /api/builder-blocks/{blockId}
//
// blockID: The block's id, e.g. `user-1a2b3c4d-pricing-card`.
func (r *BuilderBlocksService) Delete(ctx context.Context, blockID string, query *BuilderBlocksDeleteQuery, opts ...CallOption) (*BuilderBlocksDeleteResponseData, error) {
	var out *BuilderBlocksDeleteResponseData
	err := r.c.call(ctx, "DELETE", "/api/builder-blocks/"+pathSeg(blockID), query, nil, false, "data", &out, opts)
	return out, err
}

// List — Lists one page of the organisation's own blocks, the most recently changed first.
//
// Needs `component:manage_org_blocks` held in the organisation acted on. Blocks in the trash are left out.
//
// GET /api/builder-blocks
func (r *BuilderBlocksService) List(ctx context.Context, query *BuilderBlocksListQuery, opts ...CallOption) (*BuilderBlocksListResponseData, error) {
	var out *BuilderBlocksListResponseData
	err := r.c.call(ctx, "GET", "/api/builder-blocks", query, nil, false, "data", &out, opts)
	return out, err
}

// OfferInMarketplace — Prepares a product in the Products hub that offers the block in the marketplace, or finds the one already there.
//
// Needs the marketplace switched on, `component:manage_org_blocks`, `products:create` in the block's organisation, and selling in the marketplace on its plan. Nothing is listed yet: the product is finished and submitted in the Products hub, and a purchase delivers the block's current version.
//
// POST /api/builder-blocks/{blockId}/publish
//
// blockID: The block's id, e.g. `user-1a2b3c4d-pricing-card`.
func (r *BuilderBlocksService) OfferInMarketplace(ctx context.Context, blockID string, body *BuilderBlocksOfferInMarketplaceBody, query *BuilderBlocksOfferInMarketplaceQuery, opts ...CallOption) (*BuilderBlocksOfferInMarketplaceResponseData, error) {
	var out *BuilderBlocksOfferInMarketplaceResponseData
	err := r.c.call(ctx, "POST", "/api/builder-blocks/"+pathSeg(blockID)+"/publish", query, body, false, "data", &out, opts)
	return out, err
}

// Update — Changes a block's name, description, category, icon, tags, notes or whether it is offered.
//
// Needs `component:manage_org_blocks` held in the organisation acted on. Only the fields sent change; the schema and template change with a new version (builderBlocks.versions.create).
//
// PATCH /api/builder-blocks/{blockId}
//
// blockID: The block's id, e.g. `user-1a2b3c4d-pricing-card`.
func (r *BuilderBlocksService) Update(ctx context.Context, blockID string, body *BuilderBlocksUpdateBody, query *BuilderBlocksUpdateQuery, opts ...CallOption) (*BuilderBlocksUpdateResponseData, error) {
	var out *BuilderBlocksUpdateResponseData
	err := r.c.call(ctx, "PATCH", "/api/builder-blocks/"+pathSeg(blockID), query, body, false, "data", &out, opts)
	return out, err
}

// Usage — Tells how often a block of the organisation is used, and on which pages.
//
// Needs `component:manage_org_blocks` held in the organisation acted on. `pages` names at most 25.
//
// GET /api/builder-blocks/{blockId}/usage
//
// blockID: The block's id, e.g. `user-1a2b3c4d-pricing-card`.
func (r *BuilderBlocksService) Usage(ctx context.Context, blockID string, query *BuilderBlocksUsageQuery, opts ...CallOption) (*BuilderBlocksUsageResponseData, error) {
	var out *BuilderBlocksUsageResponseData
	err := r.c.call(ctx, "GET", "/api/builder-blocks/"+pathSeg(blockID)+"/usage", query, nil, false, "data", &out, opts)
	return out, err
}

// BuilderBlocksVersionsService — builderBlocks.versions
type BuilderBlocksVersionsService struct {
	c *Client
}

func newBuilderBlocksVersionsService(c *Client) *BuilderBlocksVersionsService {
	return &BuilderBlocksVersionsService{c: c}
}

// Create — Publishes a new version of a block's schema, defaults and template, which becomes the current one.
//
// Needs `component:manage_org_blocks` held in the organisation acted on.
//
// POST /api/builder-blocks/{blockId}/versions
//
// blockID: The block's id, e.g. `user-1a2b3c4d-pricing-card`.
func (r *BuilderBlocksVersionsService) Create(ctx context.Context, blockID string, body *BuilderBlocksVersionsCreateBody, query *BuilderBlocksVersionsCreateQuery, opts ...CallOption) (float64, error) {
	var out float64
	err := r.c.call(ctx, "POST", "/api/builder-blocks/"+pathSeg(blockID)+"/versions", query, body, false, "data.version", &out, opts)
	return out, err
}

// List — Lists every version of a block, the newest first.
//
// Needs `component:manage_org_blocks` held in the organisation acted on.
//
// GET /api/builder-blocks/{blockId}/versions
//
// blockID: The block's id, e.g. `user-1a2b3c4d-pricing-card`.
func (r *BuilderBlocksVersionsService) List(ctx context.Context, blockID string, query *BuilderBlocksVersionsListQuery, opts ...CallOption) ([]BuilderBlockVersion, error) {
	var out []BuilderBlockVersion
	err := r.c.call(ctx, "GET", "/api/builder-blocks/"+pathSeg(blockID)+"/versions", query, nil, false, "data.versions", &out, opts)
	return out, err
}

// CalendarService — Calendar: Calendars and their events, sharing, ICS subscriptions and CalDAV sync; booking pages with their bookings and waiting lists; Google and Microsoft accounts; meeting polls; the scheduling helpers (conflicts, holidays, AI). Everything needs Calendar on the caller's plan (booking pages: booking pages on the plan).
type CalendarService struct {
	c *Client
}

func newCalendarService(c *Client) *CalendarService {
	return &CalendarService{c: c}
}

// AddAttendee — Adds a guest to an event and invites them by mail.
//
// Events of calendars the caller owns. Answers 201.
//
// POST /api/calendar/events/{eventId}/attendees
//
// eventID: The event.
func (r *CalendarService) AddAttendee(ctx context.Context, eventID string, body *CalendarAddAttendeeBody, opts ...CallOption) (*CalendarEventAttendee, error) {
	var out *CalendarEventAttendee
	err := r.c.call(ctx, "POST", "/api/calendar/events/"+pathSeg(eventID)+"/attendees", nil, body, false, "data", &out, opts)
	return out, err
}

// AICommand — Reads an event out of a sentence (`parse_event`), or suggests free meeting times (`suggest_slots`).
//
// parse_event needs AI scheduling on the plan and counts against the organisation's AI allowance. suggest_slots needs no AI: it ranks 3–5 free times within the caller's working hours and own events.
//
// POST /api/calendar/ai
//
// The answer is one of *CalendarAiCommandResponseData, *CalendarAiCommandResponseData2, decoded as generic JSON (map[string]any, []any, …).
func (r *CalendarService) AICommand(ctx context.Context, body *CalendarAiCommandBody, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "POST", "/api/calendar/ai", nil, body, false, "data", &out, opts)
	return out, err
}

// CancelBooking — Cancels a booking: as its host, or as the guest with the booking's cancel token.
//
// The guest is told by mail and the slot is freed. 400 when already cancelled.
//
// POST /api/calendar/bookings/{id}/cancel
//
// id: The booking.
func (r *CalendarService) CancelBooking(ctx context.Context, id string, body *CalendarCancelBookingBody, opts ...CallOption) (*Booking, error) {
	var out *Booking
	err := r.c.call(ctx, "POST", "/api/calendar/bookings/"+pathSeg(id)+"/cancel", nil, body, false, "data", &out, opts)
	return out, err
}

// CheckConflicts — Lists the caller's events that overlap a proposed time, recurring ones and travel buffers included.
//
// Across the calendars the caller owns; at most 25.
//
// POST /api/calendar/conflicts
func (r *CalendarService) CheckConflicts(ctx context.Context, body *CalendarCheckConflictsBody, opts ...CallOption) ([]CalendarCheckConflictsResponseDataConflictsItem, error) {
	var out []CalendarCheckConflictsResponseDataConflictsItem
	err := r.c.call(ctx, "POST", "/api/calendar/conflicts", nil, body, false, "data.conflicts", &out, opts)
	return out, err
}

// ConfirmBooking — Confirms a pending booking of a page the caller owns.
//
// A booking held back as suspected spam is released at the same time.
//
// POST /api/calendar/bookings/{id}/confirm
//
// id: The booking.
func (r *CalendarService) ConfirmBooking(ctx context.Context, id string, opts ...CallOption) (*Booking, error) {
	var out *Booking
	err := r.c.call(ctx, "POST", "/api/calendar/bookings/"+pathSeg(id)+"/confirm", nil, nil, false, "data", &out, opts)
	return out, err
}

// CreateBookingPage — Creates a booking page on one of the caller's calendars.
//
// Counts against the plan's booking page limit. The slug is the page's public address and must be free.
//
// POST /api/calendar/booking-pages
func (r *CalendarService) CreateBookingPage(ctx context.Context, body *CalendarCreateBookingPageBody, opts ...CallOption) (*BookingPage, error) {
	var out *BookingPage
	err := r.c.call(ctx, "POST", "/api/calendar/booking-pages", nil, body, false, "data", &out, opts)
	return out, err
}

// CreateCalendar — Creates a calendar.
//
// Counts against the plan's calendar limit. A workspace or organisation calendar takes `workspaceId` / `organizationId`, else the `X-Workspace-Id` / `X-Organization-Id` header.
//
// POST /api/calendar/calendars
func (r *CalendarService) CreateCalendar(ctx context.Context, body *CalendarCreateCalendarBody, opts ...CallOption) (*Calendar, error) {
	var out *Calendar
	err := r.c.call(ctx, "POST", "/api/calendar/calendars", nil, body, false, "data", &out, opts)
	return out, err
}

// CreateEvent — Creates an event in a calendar; its guests are invited by mail.
//
// Needs edit access to the calendar. `createConference` adds a Google Meet or Teams link when the calendar mirrors such an account.
//
// POST /api/calendar/calendars/{calendarId}/events
//
// calendarID: The calendar.
func (r *CalendarService) CreateEvent(ctx context.Context, calendarID string, body *CalendarCreateEventBody, opts ...CallOption) (*CalendarEvent, error) {
	var out *CalendarEvent
	err := r.c.call(ctx, "POST", "/api/calendar/calendars/"+pathSeg(calendarID)+"/events", nil, body, false, "data", &out, opts)
	return out, err
}

// CreateEventConference — Creates a Google Meet or Teams link for an event of a calendar linked to Google or Microsoft.
//
// Events of calendars the caller owns. 409 when the calendar is not linked, 422 when the account cannot make online meetings. Google may still be making it (`status: pending`); the next sync brings it.
//
// POST /api/calendar/events/{eventId}/conference
//
// eventID: The event.
func (r *CalendarService) CreateEventConference(ctx context.Context, eventID string, opts ...CallOption) (*CalendarCreateEventConferenceResponseData, error) {
	var out *CalendarCreateEventConferenceResponseData
	err := r.c.call(ctx, "POST", "/api/calendar/events/"+pathSeg(eventID)+"/conference", nil, nil, false, "data", &out, opts)
	return out, err
}

// CreateProposal — Creates a meeting poll: 2–8 candidate times the guests vote on without an account.
//
// The guests vote at /schedule/<token>; finalize it to make the winning time an event.
//
// POST /api/calendar/proposals
func (r *CalendarService) CreateProposal(ctx context.Context, body *CalendarCreateProposalBody, opts ...CallOption) (*CalendarCreateProposalResponseData, error) {
	var out *CalendarCreateProposalResponseData
	err := r.c.call(ctx, "POST", "/api/calendar/proposals", nil, body, false, "data", &out, opts)
	return out, err
}

// CreateShare — Shares a calendar with a person by e-mail address.
//
// The calendar's owner only. Needs `calendar:share`. An address with no account yet is kept and links when they sign up. 409 when already shared. With an organisation-bound token, only members of that organisation are linked.
//
// POST /api/calendar/calendars/{calendarId}/shares
//
// calendarID: The calendar.
func (r *CalendarService) CreateShare(ctx context.Context, calendarID string, body *CalendarCreateShareBody, opts ...CallOption) (*CalendarShare, error) {
	var out *CalendarShare
	err := r.c.call(ctx, "POST", "/api/calendar/calendars/"+pathSeg(calendarID)+"/shares", nil, body, false, "data", &out, opts)
	return out, err
}

// DecideSpamBooking — Decides on a booking held back as suspected spam: let it through, or discard it without mailing anybody.
//
// The booking page's owner; only a pending booking marked as suspected.
//
// POST /api/calendar/bookings/{id}/spam
//
// id: The booking.
func (r *CalendarService) DecideSpamBooking(ctx context.Context, id string, body *CalendarDecideSpamBookingBody, opts ...CallOption) (*Booking, error) {
	var out *Booking
	err := r.c.call(ctx, "POST", "/api/calendar/bookings/"+pathSeg(id)+"/spam", nil, body, false, "data", &out, opts)
	return out, err
}

// DeleteBooking — Cancels a booking as its host; the guest is told by mail.
//
// Never deleted: it stays as cancelled. One cancelled before answers `alreadyCancelled`.
//
// DELETE /api/calendar/bookings/{id}
//
// id: The booking.
func (r *CalendarService) DeleteBooking(ctx context.Context, id string, opts ...CallOption) (*CalendarDeleteBookingResponse, error) {
	var out *CalendarDeleteBookingResponse
	err := r.c.call(ctx, "DELETE", "/api/calendar/bookings/"+pathSeg(id), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteBookingPage — Moves one of the caller's booking pages to the trash.
//
// DELETE /api/calendar/booking-pages/{id}
//
// id: The booking page.
func (r *CalendarService) DeleteBookingPage(ctx context.Context, id string, opts ...CallOption) (*CalendarDeleteBookingPageResponse, error) {
	var out *CalendarDeleteBookingPageResponse
	err := r.c.call(ctx, "DELETE", "/api/calendar/booking-pages/"+pathSeg(id), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteCalendar — Moves a calendar to the trash, or deletes it for good with `force`.
//
// The calendar's owner only.
//
// DELETE /api/calendar/calendars/{calendarId}
//
// calendarID: The calendar.
func (r *CalendarService) DeleteCalendar(ctx context.Context, calendarID string, query *CalendarDeleteCalendarQuery, opts ...CallOption) (*CalendarDeleteCalendarResponse, error) {
	var out *CalendarDeleteCalendarResponse
	err := r.c.call(ctx, "DELETE", "/api/calendar/calendars/"+pathSeg(calendarID), query, nil, false, "", &out, opts)
	return out, err
}

// DeleteEvent — Moves an event of a calendar the caller owns to the trash.
//
// DELETE /api/calendar/events/{eventId}
//
// eventID: The event.
func (r *CalendarService) DeleteEvent(ctx context.Context, eventID string, opts ...CallOption) (*CalendarDeleteEventResponse, error) {
	var out *CalendarDeleteEventResponse
	err := r.c.call(ctx, "DELETE", "/api/calendar/events/"+pathSeg(eventID), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteShare — Stops sharing a calendar with a person.
//
// The calendar's owner only. Needs `calendar:share`.
//
// DELETE /api/calendar/calendars/{calendarId}/shares/{shareId}
//
// calendarID: The calendar.
// shareID: The share.
func (r *CalendarService) DeleteShare(ctx context.Context, calendarID string, shareID string, opts ...CallOption) (*CalendarDeleteShareResponse, error) {
	var out *CalendarDeleteShareResponse
	err := r.c.call(ctx, "DELETE", "/api/calendar/calendars/"+pathSeg(calendarID)+"/shares/"+pathSeg(shareID), nil, nil, false, "", &out, opts)
	return out, err
}

// DisableCaldavSync — Turns off CalDAV sync for a calendar and removes its CalDAV copy.
//
// The calendar's owner only. Needs `calendar:caldav_enable`. The events here stay.
//
// DELETE /api/calendar/calendars/{calendarId}/caldav-sync
//
// calendarID: The calendar.
func (r *CalendarService) DisableCaldavSync(ctx context.Context, calendarID string, opts ...CallOption) (*CalendarDisableCaldavSyncResponse, error) {
	var out *CalendarDisableCaldavSyncResponse
	err := r.c.call(ctx, "DELETE", "/api/calendar/calendars/"+pathSeg(calendarID)+"/caldav-sync", nil, nil, false, "", &out, opts)
	return out, err
}

// DisconnectIntegration — Disconnects a Google or Microsoft account; its mirrored calendars go to the trash.
//
// DELETE /api/calendar/integrations/{id}
//
// id: The connected Google or Microsoft account.
func (r *CalendarService) DisconnectIntegration(ctx context.Context, id string, opts ...CallOption) (*CalendarDisconnectIntegrationResponse, error) {
	var out *CalendarDisconnectIntegrationResponse
	err := r.c.call(ctx, "DELETE", "/api/calendar/integrations/"+pathSeg(id), nil, nil, false, "", &out, opts)
	return out, err
}

// EnableCaldavSync — Turns on CalDAV sync for a calendar and copies its events there.
//
// The calendar's owner only. Needs `calendar:caldav_enable`. More than 50 events are copied in the background (`backfillQueued`).
//
// POST /api/calendar/calendars/{calendarId}/caldav-sync
//
// calendarID: The calendar.
func (r *CalendarService) EnableCaldavSync(ctx context.Context, calendarID string, opts ...CallOption) (*CalendarEnableCaldavSyncResponseData, error) {
	var out *CalendarEnableCaldavSyncResponseData
	err := r.c.call(ctx, "POST", "/api/calendar/calendars/"+pathSeg(calendarID)+"/caldav-sync", nil, nil, false, "data", &out, opts)
	return out, err
}

// EventBriefing — Writes an AI briefing for a meeting: who is coming, recent mail with them, earlier meetings.
//
// Needs view access to the event's content. Needs AI scheduling on the plan and counts against the organisation's AI allowance.
//
// POST /api/calendar/events/{eventId}/briefing
//
// eventID: The event.
func (r *CalendarService) EventBriefing(ctx context.Context, eventID string, opts ...CallOption) (*CalendarEventBriefingResponseData, error) {
	var out *CalendarEventBriefingResponseData
	err := r.c.call(ctx, "POST", "/api/calendar/events/"+pathSeg(eventID)+"/briefing", nil, nil, false, "data", &out, opts)
	return out, err
}

// EventMeetingNeed — Asks the AI whether an event needs to be a meeting, with a reply to send instead when it does not.
//
// Needs view access to the event's content. Needs AI scheduling on the plan and counts against the organisation's AI allowance.
//
// POST /api/calendar/events/{eventId}/meeting-need
//
// eventID: The event.
func (r *CalendarService) EventMeetingNeed(ctx context.Context, eventID string, opts ...CallOption) (*CalendarEventMeetingNeedResponseData, error) {
	var out *CalendarEventMeetingNeedResponseData
	err := r.c.call(ctx, "POST", "/api/calendar/events/"+pathSeg(eventID)+"/meeting-need", nil, nil, false, "data", &out, opts)
	return out, err
}

// EventOutcome — Drafts a meeting's outcome with AI (`generate`), or saves one (`save`).
//
// generate: view access; needs AI scheduling on the plan and counts against the organisation's AI allowance. Nothing is stored. save: edit access.
//
// POST /api/calendar/events/{eventId}/outcome
//
// eventID: The event.
// body: one of *CalendarEventOutcomeBody, *CalendarEventOutcomeBody2.
// The answer is one of *CalendarEventOutcomeResponseData, *CalendarEventOutcomeResponseData2, decoded as generic JSON (map[string]any, []any, …).
func (r *CalendarService) EventOutcome(ctx context.Context, eventID string, body any, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "POST", "/api/calendar/events/"+pathSeg(eventID)+"/outcome", nil, body, false, "data", &out, opts)
	return out, err
}

// EventRecurrence — Keeps a recurring event going (`renew`) or ends its series today (`kill`).
//
// Needs edit access. Ending keeps the past occurrences. 400 for an event that does not repeat.
//
// POST /api/calendar/events/{eventId}/recurrence
//
// eventID: The event.
// The answer is one of *CalendarEventRecurrenceResponseData, *CalendarEventRecurrenceResponseData2, decoded as generic JSON (map[string]any, []any, …).
func (r *CalendarService) EventRecurrence(ctx context.Context, eventID string, body *CalendarEventRecurrenceBody, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "POST", "/api/calendar/events/"+pathSeg(eventID)+"/recurrence", nil, body, false, "data", &out, opts)
	return out, err
}

// FinalizeProposal — Makes one time of a meeting poll a calendar event and closes the poll.
//
// The poll's organiser only; needs edit access to the calendar. 409 when the time collides with an event there, unless `force=true`.
//
// POST /api/calendar/proposals/{token}/finalize
//
// token: The poll's token.
func (r *CalendarService) FinalizeProposal(ctx context.Context, token string, body *CalendarFinalizeProposalBody, query *CalendarFinalizeProposalQuery, opts ...CallOption) (*CalendarFinalizeProposalResponseData, error) {
	var out *CalendarFinalizeProposalResponseData
	err := r.c.call(ctx, "POST", "/api/calendar/proposals/"+pathSeg(token)+"/finalize", query, body, false, "data", &out, opts)
	return out, err
}

// GetBooking — One booking of a page the caller hosts, with its page.
//
// GET /api/calendar/bookings/{id}
//
// id: The booking.
func (r *CalendarService) GetBooking(ctx context.Context, id string, opts ...CallOption) (*CalendarGetBookingResponseData, error) {
	var out *CalendarGetBookingResponseData
	err := r.c.call(ctx, "GET", "/api/calendar/bookings/"+pathSeg(id), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetBookingPage — One of the caller's booking pages.
//
// GET /api/calendar/booking-pages/{id}
//
// id: The booking page.
func (r *CalendarService) GetBookingPage(ctx context.Context, id string, opts ...CallOption) (*BookingPage, error) {
	var out *BookingPage
	err := r.c.call(ctx, "GET", "/api/calendar/booking-pages/"+pathSeg(id), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetCaldavSync — Whether a calendar syncs over CalDAV, and the server address and login to add it to Apple Calendar, Outlook or DAVx5.
//
// Anybody who can view the calendar (a public one: its owner or a `manage` share). Needs `calendar:caldav_enable`. `connection` is set while sync is on.
//
// GET /api/calendar/calendars/{calendarId}/caldav-sync
//
// calendarID: The calendar.
func (r *CalendarService) GetCaldavSync(ctx context.Context, calendarID string, opts ...CallOption) (*CalendarGetCaldavSyncResponseData, error) {
	var out *CalendarGetCaldavSyncResponseData
	err := r.c.call(ctx, "GET", "/api/calendar/calendars/"+pathSeg(calendarID)+"/caldav-sync", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetCalendar — One calendar the caller owns or that is shared with them.
//
// GET /api/calendar/calendars/{calendarId}
//
// calendarID: The calendar.
func (r *CalendarService) GetCalendar(ctx context.Context, calendarID string, opts ...CallOption) (*Calendar, error) {
	var out *Calendar
	err := r.c.call(ctx, "GET", "/api/calendar/calendars/"+pathSeg(calendarID), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetCalendarFeedToken — A calendar's ICS subscription URL, for Google, Apple or Outlook calendars.
//
// The calendar's owner only. The URL carries a secret: whoever has it reads the calendar.
//
// GET /api/calendar/calendars/{calendarId}/feed
//
// calendarID: The calendar.
func (r *CalendarService) GetCalendarFeedToken(ctx context.Context, calendarID string, opts ...CallOption) (*CalendarGetCalendarFeedTokenResponseData, error) {
	var out *CalendarGetCalendarFeedTokenResponseData
	err := r.c.call(ctx, "GET", "/api/calendar/calendars/"+pathSeg(calendarID)+"/feed", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetEvent — One event of a calendar the caller owns.
//
// GET /api/calendar/events/{eventId}
//
// eventID: The event.
func (r *CalendarService) GetEvent(ctx context.Context, eventID string, opts ...CallOption) (*CalendarEvent, error) {
	var out *CalendarEvent
	err := r.c.call(ctx, "GET", "/api/calendar/events/"+pathSeg(eventID), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetEventIcs — Downloads an event as an .ics file.
//
// Events of calendars the caller owns.
//
// GET /api/calendar/events/{eventId}/ics
//
// eventID: The event.
func (r *CalendarService) GetEventIcs(ctx context.Context, eventID string, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/calendar/events/"+pathSeg(eventID)+"/ics", nil, nil, false, "", &out, opts)
	return out, err
}

// GetEventOutcome — A meeting's saved outcome summary.
//
// Needs view access to the event's content.
//
// GET /api/calendar/events/{eventId}/outcome
//
// eventID: The event.
func (r *CalendarService) GetEventOutcome(ctx context.Context, eventID string, opts ...CallOption) (*CalendarGetEventOutcomeResponseData, error) {
	var out *CalendarGetEventOutcomeResponseData
	err := r.c.call(ctx, "GET", "/api/calendar/events/"+pathSeg(eventID)+"/outcome", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetSettings — The caller's calendar preferences, defaults filled in.
//
// GET /api/calendar/settings
func (r *CalendarService) GetSettings(ctx context.Context, opts ...CallOption) (*CalendarSettings, error) {
	var out *CalendarSettings
	err := r.c.call(ctx, "GET", "/api/calendar/settings", nil, nil, false, "data.settings", &out, opts)
	return out, err
}

// Holidays — The public holidays of the caller's holiday country and region for a year.
//
// Empty while no country is set in the calendar settings.
//
// GET /api/calendar/holidays
func (r *CalendarService) Holidays(ctx context.Context, query *CalendarHolidaysQuery, opts ...CallOption) ([]CalendarHolidaysResponseDataHolidaysItem, error) {
	var out []CalendarHolidaysResponseDataHolidaysItem
	err := r.c.call(ctx, "GET", "/api/calendar/holidays", query, nil, false, "data.holidays", &out, opts)
	return out, err
}

// ListAttendees — Lists an event's guests with their answers.
//
// Events of calendars the caller owns.
//
// GET /api/calendar/events/{eventId}/attendees
//
// eventID: The event.
func (r *CalendarService) ListAttendees(ctx context.Context, eventID string, opts ...CallOption) ([]CalendarEventAttendee, error) {
	var out []CalendarEventAttendee
	err := r.c.call(ctx, "GET", "/api/calendar/events/"+pathSeg(eventID)+"/attendees", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListBookingPages — Lists the caller's booking pages.
//
// GET /api/calendar/booking-pages
func (r *CalendarService) ListBookingPages(ctx context.Context, opts ...CallOption) ([]BookingPage, error) {
	var out []BookingPage
	err := r.c.call(ctx, "GET", "/api/calendar/booking-pages", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListBookings — Lists one page of the bookings of the pages the caller hosts, latest start first.
//
// GET /api/calendar/bookings
func (r *CalendarService) ListBookings(ctx context.Context, query *CalendarListBookingsQuery, opts ...CallOption) (*CalendarListBookingsResponseData, error) {
	var out *CalendarListBookingsResponseData
	err := r.c.call(ctx, "GET", "/api/calendar/bookings", query, nil, false, "data", &out, opts)
	return out, err
}

// ListCalendars — Lists the calendars the caller can see: their own, those shared with them, and their organisations' and workspaces'.
//
// The headers `X-Organization-Id` / `X-Workspace-Id` narrow the organisation and workspace calendars to one. A shared calendar carries `_sharePermission`; `feedToken` is the owner's only.
//
// GET /api/calendar/calendars
func (r *CalendarService) ListCalendars(ctx context.Context, opts ...CallOption) ([]CalendarListCalendarsResponseDataItem, error) {
	var out []CalendarListCalendarsResponseDataItem
	err := r.c.call(ctx, "GET", "/api/calendar/calendars", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListEvents — Lists a calendar's events in a window, recurring ones expanded into their occurrences.
//
// With `start` and `end`: the events touching that window, each occurrence of a recurring event on its own (at most 1,000 rows). Without: the events around `from` (default now) unexpanded, at most 2,000. A private event on a calendar the caller may only view comes back as a busy block (`masked`).
//
// GET /api/calendar/calendars/{calendarId}/events
//
// calendarID: The calendar.
func (r *CalendarService) ListEvents(ctx context.Context, calendarID string, query *CalendarListEventsQuery, opts ...CallOption) ([]CalendarEvent, error) {
	var out []CalendarEvent
	err := r.c.call(ctx, "GET", "/api/calendar/calendars/"+pathSeg(calendarID)+"/events", query, nil, false, "data", &out, opts)
	return out, err
}

// ListIntegrationCalendars — Lists a connected account's calendars and which of them are mirrored here.
//
// Asks the provider (502 when it does not answer). 400 when the account is no longer connected.
//
// GET /api/calendar/integrations/{id}/calendars
//
// id: The connected Google or Microsoft account.
func (r *CalendarService) ListIntegrationCalendars(ctx context.Context, id string, opts ...CallOption) (*CalendarListIntegrationCalendarsResponseData, error) {
	var out *CalendarListIntegrationCalendarsResponseData
	err := r.c.call(ctx, "GET", "/api/calendar/integrations/"+pathSeg(id)+"/calendars", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListIntegrations — Lists the caller's connected Google and Microsoft accounts.
//
// Connecting one happens in a browser; the tokens are never answered.
//
// GET /api/calendar/integrations
func (r *CalendarService) ListIntegrations(ctx context.Context, opts ...CallOption) ([]CalendarIntegration, error) {
	var out []CalendarIntegration
	err := r.c.call(ctx, "GET", "/api/calendar/integrations", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListProposals — Lists the caller's meeting polls, newest first, with their vote counts.
//
// The newest 100.
//
// GET /api/calendar/proposals
func (r *CalendarService) ListProposals(ctx context.Context, opts ...CallOption) ([]CalendarListProposalsResponseDataPollsItem, error) {
	var out []CalendarListProposalsResponseDataPollsItem
	err := r.c.call(ctx, "GET", "/api/calendar/proposals", nil, nil, false, "data.polls", &out, opts)
	return out, err
}

// ListShares — Lists whom a calendar is shared with.
//
// The calendar's owner only. Needs `calendar:share`. A person is named only when they share an organisation with the owner.
//
// GET /api/calendar/calendars/{calendarId}/shares
//
// calendarID: The calendar.
func (r *CalendarService) ListShares(ctx context.Context, calendarID string, opts ...CallOption) ([]CalendarShare, error) {
	var out []CalendarShare
	err := r.c.call(ctx, "GET", "/api/calendar/calendars/"+pathSeg(calendarID)+"/shares", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListWaitlist — Lists the waiting lists of the caller's booking pages, in line order.
//
// GET /api/calendar/booking-waitlist
func (r *CalendarService) ListWaitlist(ctx context.Context, query *CalendarListWaitlistQuery, opts ...CallOption) ([]BookingWaitlistEntry, error) {
	var out []BookingWaitlistEntry
	err := r.c.call(ctx, "GET", "/api/calendar/booking-waitlist", query, nil, false, "data.entries", &out, opts)
	return out, err
}

// PullCaldavSync — Brings in now what changed in a calendar's CalDAV copy.
//
// The calendar's owner only. Needs `calendar:caldav_enable`; 409 when sync is off. `skipped` when nothing changed.
//
// POST /api/calendar/calendars/{calendarId}/caldav-sync/refresh
//
// calendarID: The calendar.
func (r *CalendarService) PullCaldavSync(ctx context.Context, calendarID string, opts ...CallOption) (*CalendarPullCaldavSyncResponseData, error) {
	var out *CalendarPullCaldavSyncResponseData
	err := r.c.call(ctx, "POST", "/api/calendar/calendars/"+pathSeg(calendarID)+"/caldav-sync/refresh", nil, nil, false, "data", &out, opts)
	return out, err
}

// Refresh — Brings in now what changed in the caller's CalDAV calendars and connected accounts.
//
// At most once every 30 seconds; a call inside that answers `skipped: true`.
//
// POST /api/calendar/refresh
func (r *CalendarService) Refresh(ctx context.Context, opts ...CallOption) (*CalendarRefreshResponseData, error) {
	var out *CalendarRefreshResponseData
	err := r.c.call(ctx, "POST", "/api/calendar/refresh", nil, nil, false, "data", &out, opts)
	return out, err
}

// RemoveAttendee — Removes a guest from an event and tells them it is off.
//
// Events of calendars the caller owns.
//
// DELETE /api/calendar/events/{eventId}/attendees
//
// eventID: The event.
func (r *CalendarService) RemoveAttendee(ctx context.Context, eventID string, query *CalendarRemoveAttendeeQuery, opts ...CallOption) (*CalendarRemoveAttendeeResponse, error) {
	var out *CalendarRemoveAttendeeResponse
	err := r.c.call(ctx, "DELETE", "/api/calendar/events/"+pathSeg(eventID)+"/attendees", query, nil, false, "", &out, opts)
	return out, err
}

// RemoveWaitlistEntry — Takes somebody off a booking page's waiting list; a slot held for them goes to the next in line.
//
// DELETE /api/calendar/booking-waitlist/{entryId}
//
// entryID: The waiting-list entry.
func (r *CalendarService) RemoveWaitlistEntry(ctx context.Context, entryID string, opts ...CallOption) (*CalendarRemoveWaitlistEntryResponse, error) {
	var out *CalendarRemoveWaitlistEntryResponse
	err := r.c.call(ctx, "DELETE", "/api/calendar/booking-waitlist/"+pathSeg(entryID), nil, nil, false, "", &out, opts)
	return out, err
}

// RescheduleBooking — Moves a booking to another time: as the host of its page, or as the guest with the reschedule token.
//
// The new time must be bookable on the page (409 otherwise); a guest also within its notice and advance limits. A new reschedule token is made.
//
// POST /api/calendar/bookings/{id}/reschedule
//
// id: The booking.
func (r *CalendarService) RescheduleBooking(ctx context.Context, id string, body *CalendarRescheduleBookingBody, opts ...CallOption) (*CalendarRescheduleBookingResponseData, error) {
	var out *CalendarRescheduleBookingResponseData
	err := r.c.call(ctx, "POST", "/api/calendar/bookings/"+pathSeg(id)+"/reschedule", nil, body, false, "data", &out, opts)
	return out, err
}

// RevokeCalendarFeed — Switches a calendar's ICS subscription off until it is rotated.
//
// The calendar's owner only.
//
// DELETE /api/calendar/calendars/{calendarId}/feed
//
// calendarID: The calendar.
func (r *CalendarService) RevokeCalendarFeed(ctx context.Context, calendarID string, opts ...CallOption) (*CalendarRevokeCalendarFeedResponse, error) {
	var out *CalendarRevokeCalendarFeedResponse
	err := r.c.call(ctx, "DELETE", "/api/calendar/calendars/"+pathSeg(calendarID)+"/feed", nil, nil, false, "", &out, opts)
	return out, err
}

// RotateCalendarFeed — Makes a new ICS subscription URL; every old one stops working.
//
// The calendar's owner only. Also switches a revoked feed back on.
//
// POST /api/calendar/calendars/{calendarId}/feed
//
// calendarID: The calendar.
func (r *CalendarService) RotateCalendarFeed(ctx context.Context, calendarID string, opts ...CallOption) (*CalendarRotateCalendarFeedResponseData, error) {
	var out *CalendarRotateCalendarFeedResponseData
	err := r.c.call(ctx, "POST", "/api/calendar/calendars/"+pathSeg(calendarID)+"/feed", nil, nil, false, "data", &out, opts)
	return out, err
}

// Search — Searches the events of the caller's own and shared calendars by title, description, place or guest.
//
// From 90 days back to a year ahead, newest first. Private events of calendars the caller may only view never match. 30 a minute; fewer than 2 characters answer nothing.
//
// GET /api/calendar/search
func (r *CalendarService) Search(ctx context.Context, query *CalendarSearchQuery, opts ...CallOption) ([]CalendarSearchResponseDataResultsItem, error) {
	var out []CalendarSearchResponseDataResultsItem
	err := r.c.call(ctx, "GET", "/api/calendar/search", query, nil, false, "data.results", &out, opts)
	return out, err
}

// SetIntegrationCalendars — Chooses which of a connected account's calendars are mirrored here, then syncs.
//
// Each chosen one gets a calendar of its own; one no longer chosen goes to the trash here, its events stay at the provider.
//
// PUT /api/calendar/integrations/{id}/calendars
//
// id: The connected Google or Microsoft account.
func (r *CalendarService) SetIntegrationCalendars(ctx context.Context, id string, body *CalendarSetIntegrationCalendarsBody, opts ...CallOption) (*CalendarSetIntegrationCalendarsResponseData, error) {
	var out *CalendarSetIntegrationCalendarsResponseData
	err := r.c.call(ctx, "PUT", "/api/calendar/integrations/"+pathSeg(id)+"/calendars", nil, body, false, "data", &out, opts)
	return out, err
}

// SyncIntegration — Syncs a connected account now, both ways.
//
// POST /api/calendar/integrations/{id}/sync
//
// id: The connected Google or Microsoft account.
func (r *CalendarService) SyncIntegration(ctx context.Context, id string, opts ...CallOption) (*CalendarSyncCounts, error) {
	var out *CalendarSyncCounts
	err := r.c.call(ctx, "POST", "/api/calendar/integrations/"+pathSeg(id)+"/sync", nil, nil, false, "data", &out, opts)
	return out, err
}

// TravelEstimate — Estimates the travel time in minutes between two places.
//
// Needs the travel buffer feature. `minutes` is null when the platform has no maps provider. 30 a minute.
//
// POST /api/calendar/travel-estimate
func (r *CalendarService) TravelEstimate(ctx context.Context, body *CalendarTravelEstimateBody, opts ...CallOption) (*CalendarTravelEstimateResponseData, error) {
	var out *CalendarTravelEstimateResponseData
	err := r.c.call(ctx, "POST", "/api/calendar/travel-estimate", nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateBookingPage — Changes one of the caller's booking pages.
//
// Only the fields sent change.
//
// PUT /api/calendar/booking-pages/{id}
//
// id: The booking page.
func (r *CalendarService) UpdateBookingPage(ctx context.Context, id string, body *CalendarUpdateBookingPageBody, opts ...CallOption) (*BookingPage, error) {
	var out *BookingPage
	err := r.c.call(ctx, "PUT", "/api/calendar/booking-pages/"+pathSeg(id), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateCalendar — Changes a calendar's name, description, colour, zone, visibility or settings.
//
// The calendar's owner only. Only the fields sent change.
//
// PUT /api/calendar/calendars/{calendarId}
//
// calendarID: The calendar.
func (r *CalendarService) UpdateCalendar(ctx context.Context, calendarID string, body *CalendarUpdateCalendarBody, opts ...CallOption) (*Calendar, error) {
	var out *Calendar
	err := r.c.call(ctx, "PUT", "/api/calendar/calendars/"+pathSeg(calendarID), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateEvent — Changes an event of a calendar the caller owns; guests are told what changed.
//
// Only the fields sent change.
//
// PUT /api/calendar/events/{eventId}
//
// eventID: The event.
func (r *CalendarService) UpdateEvent(ctx context.Context, eventID string, body *CalendarUpdateEventBody, opts ...CallOption) (*CalendarEvent, error) {
	var out *CalendarEvent
	err := r.c.call(ctx, "PUT", "/api/calendar/events/"+pathSeg(eventID), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateSettings — Changes the caller's calendar preferences.
//
// Only the fields sent change; an unknown field is refused.
//
// PUT /api/calendar/settings
func (r *CalendarService) UpdateSettings(ctx context.Context, body *CalendarUpdateSettingsBody, opts ...CallOption) (*CalendarSettings, error) {
	var out *CalendarSettings
	err := r.c.call(ctx, "PUT", "/api/calendar/settings", nil, body, false, "data.settings", &out, opts)
	return out, err
}

// UpdateShare — Changes what a person a calendar is shared with may do.
//
// The calendar's owner only. Needs `calendar:share`.
//
// PATCH /api/calendar/calendars/{calendarId}/shares/{shareId}
//
// calendarID: The calendar.
// shareID: The share.
func (r *CalendarService) UpdateShare(ctx context.Context, calendarID string, shareID string, body *CalendarUpdateShareBody, opts ...CallOption) (*CalendarShare, error) {
	var out *CalendarShare
	err := r.c.call(ctx, "PATCH", "/api/calendar/calendars/"+pathSeg(calendarID)+"/shares/"+pathSeg(shareID), nil, body, false, "data", &out, opts)
	return out, err
}

// CodeService — Coding hub: Code projects: their files, environment variables, git sync, the databases their code may reach, and the sites they are published as.
type CodeService struct {
	c *Client

	Assets      *CodeAssetsService
	Data        *CodeDataService
	Deployments *CodeDeploymentsService
	Env         *CodeEnvService
	Files       *CodeFilesService
	Git         *CodeGitService
	Projects    *CodeProjectsService
}

func newCodeService(c *Client) *CodeService {
	return &CodeService{c: c, Assets: newCodeAssetsService(c), Data: newCodeDataService(c), Deployments: newCodeDeploymentsService(c), Env: newCodeEnvService(c), Files: newCodeFilesService(c), Git: newCodeGitService(c), Projects: newCodeProjectsService(c)}
}

// Storage — Returns how much code storage each project the caller can open uses.
//
// Uploaded files count in the Assets hub, not here.
//
// GET /api/code/storage
func (r *CodeService) Storage(ctx context.Context, query *CodeStorageQuery, opts ...CallOption) (*CodeStorageResponse, error) {
	var out *CodeStorageResponse
	err := r.c.call(ctx, "GET", "/api/code/storage", query, nil, false, "", &out, opts)
	return out, err
}

// CodeAssetsService — code.assets
type CodeAssetsService struct {
	c *Client
}

func newCodeAssetsService(c *Client) *CodeAssetsService {
	return &CodeAssetsService{c: c}
}

// GetText — Returns the text of a file in the Assets hub, as the code editor opens it.
//
// For a text file up to 2 MB the caller can see. Otherwise 415 with `reason` (`excluded`, `not_text`, `too_large`), or 502 `unreadable`.
//
// GET /api/code/assets/{assetId}
//
// assetID: The asset's id.
func (r *CodeAssetsService) GetText(ctx context.Context, assetID string, opts ...CallOption) (*CodeAssetsGetTextResponse, error) {
	var out *CodeAssetsGetTextResponse
	err := r.c.call(ctx, "GET", "/api/code/assets/"+pathSeg(assetID), nil, nil, false, "", &out, opts)
	return out, err
}

// SaveText — Replaces the text of a file in the Assets hub, keeping the previous content as a version.
//
// The asset keeps its id and URL, so what links to it shows the new text. Needs the right to manage the asset; at most 2 MB (413).
//
// PUT /api/code/assets/{assetId}
//
// assetID: The asset's id.
func (r *CodeAssetsService) SaveText(ctx context.Context, assetID string, body *CodeAssetsSaveTextBody, opts ...CallOption) (*CodeAssetsSaveTextResponse, error) {
	var out *CodeAssetsSaveTextResponse
	err := r.c.call(ctx, "PUT", "/api/code/assets/"+pathSeg(assetID), nil, body, false, "", &out, opts)
	return out, err
}

// CodeDataService — code.data
type CodeDataService struct {
	c *Client
}

func newCodeDataService(c *Client) *CodeDataService {
	return &CodeDataService{c: c}
}

// Bind — Lets a code project's code reach a database under an alias.
//
// Needs `code:update` and, on the database, read access (`access: read`) or edit access (`write`); the database must be the project's organization's. A refusal answers with `problem`: `not_found`, `forbidden`, `other_tenant`, `too_many`, `alias_taken` or the alias's shape.
//
// POST /api/code/projects/{projectId}/data
//
// projectID: The code project's id.
func (r *CodeDataService) Bind(ctx context.Context, projectID string, body *CodeDataBindBody, opts ...CallOption) (*CodeDataBindResponseBinding, error) {
	var out *CodeDataBindResponseBinding
	err := r.c.call(ctx, "POST", "/api/code/projects/"+pathSeg(projectID)+"/data", nil, body, false, "binding", &out, opts)
	return out, err
}

// List — Lists the databases a code project's code may reach, and those the caller could add.
//
// GET /api/code/projects/{projectId}/data
//
// projectID: The code project's id.
func (r *CodeDataService) List(ctx context.Context, projectID string, opts ...CallOption) ([]CodeDataBinding, error) {
	var out []CodeDataBinding
	err := r.c.call(ctx, "GET", "/api/code/projects/"+pathSeg(projectID)+"/data", nil, nil, false, "bindings", &out, opts)
	return out, err
}

// Unbind — Stops a code project's code from reaching a database.
//
// Needs `code:update`.
//
// DELETE /api/code/projects/{projectId}/data
//
// projectID: The code project's id.
func (r *CodeDataService) Unbind(ctx context.Context, projectID string, query *CodeDataUnbindQuery, opts ...CallOption) (*CodeDataUnbindResponse, error) {
	var out *CodeDataUnbindResponse
	err := r.c.call(ctx, "DELETE", "/api/code/projects/"+pathSeg(projectID)+"/data", query, nil, false, "", &out, opts)
	return out, err
}

// Update — Changes whether a code project may only read or also write a bound database.
//
// Needs `code:update`; `write` needs edit access to the database.
//
// PATCH /api/code/projects/{projectId}/data
//
// projectID: The code project's id.
func (r *CodeDataService) Update(ctx context.Context, projectID string, body *CodeDataUpdateBody, opts ...CallOption) (*CodeDataUpdateResponseBinding, error) {
	var out *CodeDataUpdateResponseBinding
	err := r.c.call(ctx, "PATCH", "/api/code/projects/"+pathSeg(projectID)+"/data", nil, body, false, "binding", &out, opts)
	return out, err
}

// CodeDeploymentsService — code.deployments
type CodeDeploymentsService struct {
	c *Client
}

func newCodeDeploymentsService(c *Client) *CodeDeploymentsService {
	return &CodeDeploymentsService{c: c}
}

// Delete — Takes a published site down and frees its name.
//
// Needs `code:update`. `alreadyGone: true` when nothing was published.
//
// DELETE /api/code/projects/{projectId}/deployments
//
// projectID: The code project's id.
func (r *CodeDeploymentsService) Delete(ctx context.Context, projectID string, query *CodeDeploymentsDeleteQuery, opts ...CallOption) (*CodeDeploymentsDeleteResponse, error) {
	var out *CodeDeploymentsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/code/projects/"+pathSeg(projectID)+"/deployments", query, nil, false, "", &out, opts)
	return out, err
}

// List — Lists a code project's published sites, with their containers and this month's compute use.
//
// GET /api/code/projects/{projectId}/deployments
//
// projectID: The code project's id.
func (r *CodeDeploymentsService) List(ctx context.Context, projectID string, opts ...CallOption) ([]CodeDeploymentsListResponseDeploymentsItem, error) {
	var out []CodeDeploymentsListResponseDeploymentsItem
	err := r.c.call(ctx, "GET", "/api/code/projects/"+pathSeg(projectID)+"/deployments", nil, nil, false, "deployments", &out, opts)
	return out, err
}

// Logs — Returns the last 200 lines a published site's container wrote.
//
// `live: false` is what was captured when the container stopped (`reason`: `stopped` or `unreachable`). 404 with `reason: not_published` when the environment is not published.
//
// GET /api/code/projects/{projectId}/deployments/logs
//
// projectID: The code project's id.
func (r *CodeDeploymentsService) Logs(ctx context.Context, projectID string, query *CodeDeploymentsLogsQuery, opts ...CallOption) (*CodeDeploymentsLogsResponse, error) {
	var out *CodeDeploymentsLogsResponse
	err := r.c.call(ctx, "GET", "/api/code/projects/"+pathSeg(projectID)+"/deployments/logs", query, nil, false, "", &out, opts)
	return out, err
}

// Publish — Publishes a code project as a site, or publishes it again.
//
// Needs `code:update`. Builds the current files in a container of their own and runs the project's start script; the answer comes when it is up. The first publish of an environment takes `slug`; the plan's environments and deployments apply. A refusal answers with `reason`: `no_domain` or `not_enabled` (403), `slug` with `problem` (`too_short`, `too_long`, `shape`, `taken`), `no_start_script` (400), `limit` (403).
//
// POST /api/code/projects/{projectId}/deployments
//
// projectID: The code project's id.
func (r *CodeDeploymentsService) Publish(ctx context.Context, projectID string, body *CodeDeploymentsPublishBody, opts ...CallOption) (*CodeDeployment, error) {
	var out *CodeDeployment
	err := r.c.call(ctx, "POST", "/api/code/projects/"+pathSeg(projectID)+"/deployments", nil, body, false, "deployment", &out, opts)
	return out, err
}

// RemoveDomain — Removes a published site's custom domain.
//
// Needs `code:update`.
//
// DELETE /api/code/projects/{projectId}/deployments/domain
//
// projectID: The code project's id.
func (r *CodeDeploymentsService) RemoveDomain(ctx context.Context, projectID string, query *CodeDeploymentsRemoveDomainQuery, opts ...CallOption) (*CodeDeploymentsRemoveDomainResponse, error) {
	var out *CodeDeploymentsRemoveDomainResponse
	err := r.c.call(ctx, "DELETE", "/api/code/projects/"+pathSeg(projectID)+"/deployments/domain", query, nil, false, "", &out, opts)
	return out, err
}

// SetDomain — Claims a custom domain for a published site.
//
// Needs `code:update`, custom domains on the plan and a published environment (409 otherwise). The domain is pending until `code.deployments.verifyDomain` finds the TXT `record` the answer names.
//
// PUT /api/code/projects/{projectId}/deployments/domain
//
// projectID: The code project's id.
func (r *CodeDeploymentsService) SetDomain(ctx context.Context, projectID string, body *CodeDeploymentsSetDomainBody, query *CodeDeploymentsSetDomainQuery, opts ...CallOption) (*CodeDeploymentsSetDomainResponse, error) {
	var out *CodeDeploymentsSetDomainResponse
	err := r.c.call(ctx, "PUT", "/api/code/projects/"+pathSeg(projectID)+"/deployments/domain", query, body, false, "", &out, opts)
	return out, err
}

// Update — Turns a required sign-in on or off for a published site, without publishing again.
//
// Needs `code:update`. Visitors sign in as the workspace's site users, so turning it on needs a project with a workspace (400 `problem: no_workspace`).
//
// PATCH /api/code/projects/{projectId}/deployments
//
// projectID: The code project's id.
func (r *CodeDeploymentsService) Update(ctx context.Context, projectID string, body *CodeDeploymentsUpdateBody, opts ...CallOption) (*CodeDeploymentsUpdateResponseDeployment, error) {
	var out *CodeDeploymentsUpdateResponseDeployment
	err := r.c.call(ctx, "PATCH", "/api/code/projects/"+pathSeg(projectID)+"/deployments", nil, body, false, "deployment", &out, opts)
	return out, err
}

// VerifyDomain — Checks a site's custom domain for its TXT record and marks it verified when it is there.
//
// Needs `code:update`.
//
// POST /api/code/projects/{projectId}/deployments/domain
//
// projectID: The code project's id.
func (r *CodeDeploymentsService) VerifyDomain(ctx context.Context, projectID string, query *CodeDeploymentsVerifyDomainQuery, opts ...CallOption) (*CodeDeploymentsVerifyDomainResponse, error) {
	var out *CodeDeploymentsVerifyDomainResponse
	err := r.c.call(ctx, "POST", "/api/code/projects/"+pathSeg(projectID)+"/deployments/domain", query, nil, false, "", &out, opts)
	return out, err
}

// CodeEnvService — code.env
type CodeEnvService struct {
	c *Client
}

func newCodeEnvService(c *Client) *CodeEnvService {
	return &CodeEnvService{c: c}
}

// Delete — Removes an environment variable of a code project.
//
// Needs `code:update`. Applies on the container's next start.
//
// DELETE /api/code/projects/{projectId}/env
//
// projectID: The code project's id.
func (r *CodeEnvService) Delete(ctx context.Context, projectID string, query *CodeEnvDeleteQuery, opts ...CallOption) (*CodeEnvDeleteResponse, error) {
	var out *CodeEnvDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/code/projects/"+pathSeg(projectID)+"/env", query, nil, false, "", &out, opts)
	return out, err
}

// List — Lists a code project's environment variables, without their values.
//
// GET /api/code/projects/{projectId}/env
//
// projectID: The code project's id.
func (r *CodeEnvService) List(ctx context.Context, projectID string, opts ...CallOption) ([]CodeEnvVariable, error) {
	var out []CodeEnvVariable
	err := r.c.call(ctx, "GET", "/api/code/projects/"+pathSeg(projectID)+"/env", nil, nil, false, "variables", &out, opts)
	return out, err
}

// Set — Sets an environment variable of a code project.
//
// Needs `code:update`. Stored encrypted; the value never comes back. A running container keeps the environment it started with — the change applies on its next start. A refusal answers 400 with `problem`: `shape`, `too_long`, `reserved`, `value_too_long` or `too_many`.
//
// PUT /api/code/projects/{projectId}/env
//
// projectID: The code project's id.
func (r *CodeEnvService) Set(ctx context.Context, projectID string, body *CodeEnvSetBody, opts ...CallOption) (*CodeEnvSetResponse, error) {
	var out *CodeEnvSetResponse
	err := r.c.call(ctx, "PUT", "/api/code/projects/"+pathSeg(projectID)+"/env", nil, body, false, "", &out, opts)
	return out, err
}

// CodeFilesService — code.files
type CodeFilesService struct {
	c *Client
}

func newCodeFilesService(c *Client) *CodeFilesService {
	return &CodeFilesService{c: c}
}

// Create — Creates a file in a code project.
//
// Needs `code:update`. Text up to 1 MB; a non-text file is uploaded to the Assets hub and added here by `assetId`. 409 when the path exists; the project's file count and storage limits apply.
//
// POST /api/code/projects/{projectId}/files
//
// projectID: The code project's id.
func (r *CodeFilesService) Create(ctx context.Context, projectID string, body *CodeFilesCreateBody, opts ...CallOption) (*CodeFile, error) {
	var out *CodeFile
	err := r.c.call(ctx, "POST", "/api/code/projects/"+pathSeg(projectID)+"/files", nil, body, false, "file", &out, opts)
	return out, err
}

// Delete — Deletes a file, or a folder with everything in it.
//
// Needs `code:update`.
//
// DELETE /api/code/projects/{projectId}/files
//
// projectID: The code project's id.
func (r *CodeFilesService) Delete(ctx context.Context, projectID string, query *CodeFilesDeleteQuery, opts ...CallOption) (*CodeFilesDeleteResponse, error) {
	var out *CodeFilesDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/code/projects/"+pathSeg(projectID)+"/files", query, nil, false, "", &out, opts)
	return out, err
}

// Get — Returns one file of a code project, with its content.
//
// While the project's container runs, the content is the container's; `stale: true` means the container did not answer in time and the last saved copy was served.
//
// GET /api/code/projects/{projectId}/files
//
// projectID: The code project's id.
func (r *CodeFilesService) Get(ctx context.Context, projectID string, query *CodeFilesGetQuery, opts ...CallOption) (*CodeFile, error) {
	var out *CodeFile
	err := r.c.call(ctx, "GET", "/api/code/projects/"+pathSeg(projectID)+"/files", query, nil, false, "file", &out, opts)
	return out, err
}

// Update — Saves a file's content, or renames or moves a file or a folder.
//
// Needs `code:update`. With `newPath` the file — or the folder and everything in it — moves there and the answer carries `moved`; otherwise `content` is saved and the answer carries the `file`, or `unchanged: true` when nothing changed.
//
// PATCH /api/code/projects/{projectId}/files
//
// projectID: The code project's id.
func (r *CodeFilesService) Update(ctx context.Context, projectID string, body *CodeFilesUpdateBody, opts ...CallOption) (*CodeFilesUpdateResponse, error) {
	var out *CodeFilesUpdateResponse
	err := r.c.call(ctx, "PATCH", "/api/code/projects/"+pathSeg(projectID)+"/files", nil, body, false, "", &out, opts)
	return out, err
}

// CodeGitService — code.git
type CodeGitService struct {
	c *Client

	Credentials *CodeGitCredentialsService
}

func newCodeGitService(c *Client) *CodeGitService {
	return &CodeGitService{c: c, Credentials: newCodeGitCredentialsService(c)}
}

// Connect — Connects a code project to a repository, or disconnects it.
//
// Needs `code:update`. Without `remoteUrl` the project is disconnected. Give a credential the caller may use (`code.git.credentials.list`): `credentialId` for a token, `connectionId` for a signed-in GitHub account.
//
// PUT /api/code/projects/{projectId}/git
//
// projectID: The code project's id.
func (r *CodeGitService) Connect(ctx context.Context, projectID string, body *CodeGitConnectBody, opts ...CallOption) (*CodeGitConnectResponse, error) {
	var out *CodeGitConnectResponse
	err := r.c.call(ctx, "PUT", "/api/code/projects/"+pathSeg(projectID)+"/git", nil, body, false, "", &out, opts)
	return out, err
}

// History — Lists the commits of a code project's branch, its branches, one commit's changes, or one file as last synced.
//
// One of: nothing — the branch's latest 30 commits (`commits`); `branches=1` — the repository's branches; `sha` — that commit's `message` and `files` with their diffs; `path` — that file's `content` at the commit the project last synced to (null when there is none).
//
// GET /api/code/projects/{projectId}/git/history
//
// projectID: The code project's id.
func (r *CodeGitService) History(ctx context.Context, projectID string, query *CodeGitHistoryQuery, opts ...CallOption) (*CodeGitHistoryResponse, error) {
	var out *CodeGitHistoryResponse
	err := r.c.call(ctx, "GET", "/api/code/projects/"+pathSeg(projectID)+"/git/history", query, nil, false, "", &out, opts)
	return out, err
}

// Repos — Lists the repositories a git credential can push to, or one repository's branches.
//
// Give `credentialId` or `connectionId` (with its `workspaceId`). With `repo`, the answer is that repository's `branches` instead of `repos`.
//
// GET /api/code/git/repos
func (r *CodeGitService) Repos(ctx context.Context, query *CodeGitReposQuery, opts ...CallOption) (*CodeGitReposResponse, error) {
	var out *CodeGitReposResponse
	err := r.c.call(ctx, "GET", "/api/code/git/repos", query, nil, false, "", &out, opts)
	return out, err
}

// Status — Returns a code project's repository connection and what differs from it.
//
// Needs Git sync on the account. `connected: false` with a `reason` when there is no repository or no usable credential.
//
// GET /api/code/projects/{projectId}/git
//
// projectID: The code project's id.
func (r *CodeGitService) Status(ctx context.Context, projectID string, opts ...CallOption) (*CodeGitStatusResponse, error) {
	var out *CodeGitStatusResponse
	err := r.c.call(ctx, "GET", "/api/code/projects/"+pathSeg(projectID)+"/git", nil, nil, false, "", &out, opts)
	return out, err
}

// Sync — Pulls from, pushes to, or switches the branch of a code project's repository.
//
// Needs `code:update`. `pull` refuses to overwrite local changes unless `force`; `push` makes one commit of what changed and refuses when the branch moved since the last sync; `switch-branch` refuses with local changes (409) unless `force`, which discards them.
//
// POST /api/code/projects/{projectId}/git
//
// projectID: The code project's id.
func (r *CodeGitService) Sync(ctx context.Context, projectID string, body *CodeGitSyncBody, opts ...CallOption) (*CodeGitSyncResponse, error) {
	var out *CodeGitSyncResponse
	err := r.c.call(ctx, "POST", "/api/code/projects/"+pathSeg(projectID)+"/git", nil, body, false, "", &out, opts)
	return out, err
}

// CodeGitCredentialsService — code.git.credentials
type CodeGitCredentialsService struct {
	c *Client
}

func newCodeGitCredentialsService(c *Client) *CodeGitCredentialsService {
	return &CodeGitCredentialsService{c: c}
}

// Create — Stores a git host access token, after the host accepted it.
//
// Needs `code:update`; a credential shared with an organization (`scope: organization`) also needs `organization:update`. The token is stored encrypted and never comes back.
//
// POST /api/code/git/credentials
func (r *CodeGitCredentialsService) Create(ctx context.Context, body *CodeGitCredentialsCreateBody, opts ...CallOption) (*CodeGitCredential, error) {
	var out *CodeGitCredential
	err := r.c.call(ctx, "POST", "/api/code/git/credentials", nil, body, false, "credential", &out, opts)
	return out, err
}

// Delete — Deletes a git credential or a signed-in GitHub account.
//
// DELETE /api/code/git/credentials
func (r *CodeGitCredentialsService) Delete(ctx context.Context, query *CodeGitCredentialsDeleteQuery, opts ...CallOption) (*CodeGitCredentialsDeleteResponse, error) {
	var out *CodeGitCredentialsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/code/git/credentials", query, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the git credentials the caller may use, without their tokens.
//
// The caller's own tokens and their organizations' shared ones; with `workspaceId`, also the GitHub accounts signed in to in that workspace.
//
// GET /api/code/git/credentials
func (r *CodeGitCredentialsService) List(ctx context.Context, query *CodeGitCredentialsListQuery, opts ...CallOption) ([]CodeGitCredential, error) {
	var out []CodeGitCredential
	err := r.c.call(ctx, "GET", "/api/code/git/credentials", query, nil, false, "credentials", &out, opts)
	return out, err
}

// CodeProjectsService — code.projects
type CodeProjectsService struct {
	c *Client
}

func newCodeProjectsService(c *Client) *CodeProjectsService {
	return &CodeProjectsService{c: c}
}

// Create — Creates a code project, from a template or with a starter file.
//
// Needs `code:create` and the Coding hub on the plan; counts against the organization's `max_code_projects`. Without `template` the project starts with the language's starter file, unless `empty` is true.
//
// POST /api/code/projects
func (r *CodeProjectsService) Create(ctx context.Context, body *CodeProjectsCreateBody, opts ...CallOption) (*CodeProject, error) {
	var out *CodeProject
	err := r.c.call(ctx, "POST", "/api/code/projects", nil, body, false, "project", &out, opts)
	return out, err
}

// Delete — Moves a code project to the trash.
//
// Needs `code:delete`, and only the owner may. What runs stops and its sites go offline; the files wait in the trash until it is deleted for good.
//
// DELETE /api/code/projects/{projectId}
//
// projectID: The code project's id.
func (r *CodeProjectsService) Delete(ctx context.Context, projectID string, opts ...CallOption) (*CodeProjectsDeleteResponse, error) {
	var out *CodeProjectsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/code/projects/"+pathSeg(projectID), nil, nil, false, "", &out, opts)
	return out, err
}

// Download — Downloads a code project's files as a ZIP file.
//
// Answers the file itself (`application/zip`), not JSON. Uploaded files live in the Assets hub and are only listed, in `UPLOADS.txt`.
//
// GET /api/code/projects/{projectId}/download
//
// projectID: The code project's id.
func (r *CodeProjectsService) Download(ctx context.Context, projectID string, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/code/projects/"+pathSeg(projectID)+"/download", nil, nil, false, "", &out, opts)
	return out, err
}

// Get — Returns a code project with its file tree, without the files' contents.
//
// Read a file's content with `code.files.get`.
//
// GET /api/code/projects/{projectId}
//
// projectID: The code project's id.
func (r *CodeProjectsService) Get(ctx context.Context, projectID string, opts ...CallOption) (*CodeProjectsGetResponse, error) {
	var out *CodeProjectsGetResponse
	err := r.c.call(ctx, "GET", "/api/code/projects/"+pathSeg(projectID), nil, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the caller's code projects, and the shared ones of a workspace or organization.
//
// Always the caller's own. With `workspaceId`, also that workspace's projects shared with it; with `organizationId`, also the organization's projects shared with it — each only where the caller belongs. Most recently opened first; projects in the trash are left out.
//
// GET /api/code/projects
func (r *CodeProjectsService) List(ctx context.Context, query *CodeProjectsListQuery, opts ...CallOption) ([]CodeProjectsListResponseProjectsItem, error) {
	var out []CodeProjectsListResponseProjectsItem
	err := r.c.call(ctx, "GET", "/api/code/projects", query, nil, false, "projects", &out, opts)
	return out, err
}

// Update — Renames, describes, archives or re-scopes a code project.
//
// Needs `code:update`. Only the owner may change `visibility`; people who lose write access are disconnected from live editing at once.
//
// PATCH /api/code/projects/{projectId}
//
// projectID: The code project's id.
func (r *CodeProjectsService) Update(ctx context.Context, projectID string, body *CodeProjectsUpdateBody, opts ...CallOption) (*CodeProject, error) {
	var out *CodeProject
	err := r.c.call(ctx, "PATCH", "/api/code/projects/"+pathSeg(projectID), nil, body, false, "project", &out, opts)
	return out, err
}

// CommunityService — Community: Community hubs — forums with members, posts, replies, polls, AMAs, tags, invitations, join requests, reputation and moderation. A hub is named by its slug. Creating and changing things needs Community on the caller's plan; what a caller may do in a hub follows their role there (owner, moderator, member).
type CommunityService struct {
	c *Client
}

func newCommunityService(c *Client) *CommunityService {
	return &CommunityService{c: c}
}

// AcceptAnswer — Marks a reply as a question's accepted answer.
//
// The question's author or a moderator. Its author earns reputation.
//
// POST /api/community/hubs/{slug}/posts/{postId}/accept-answer
//
// slug: The hub's slug.
// postID: The post.
func (r *CommunityService) AcceptAnswer(ctx context.Context, slug string, postID string, body *CommunityAcceptAnswerBody, opts ...CallOption) (*CommunityAcceptAnswerResponseData, error) {
	var out *CommunityAcceptAnswerResponseData
	err := r.c.call(ctx, "POST", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID)+"/accept-answer", nil, body, false, "data", &out, opts)
	return out, err
}

// AcceptInvite — Joins a hub with an invitation.
//
// 409 when already a member, 404 when the invitation is expired or used up.
//
// POST /api/community/invites/{token}
//
// token: The invitation's token.
func (r *CommunityService) AcceptInvite(ctx context.Context, token string, opts ...CallOption) (*CommunityAcceptInviteResponseData, error) {
	var out *CommunityAcceptInviteResponseData
	err := r.c.call(ctx, "POST", "/api/community/invites/"+pathSeg(token), nil, nil, false, "data", &out, opts)
	return out, err
}

// AISummarize — An AI summary of a post's discussion.
//
// Null with `reason: not_enough_replies` for a short one. Cached until ten more replies arrive. Not for a person who opted out of AI.
//
// GET /api/community/hubs/{slug}/posts/{postId}/summary
//
// slug: The hub's slug.
// postID: The post.
func (r *CommunityService) AISummarize(ctx context.Context, slug string, postID string, opts ...CallOption) (*CommunityAiSummarizeResponseData, error) {
	var out *CommunityAiSummarizeResponseData
	err := r.c.call(ctx, "GET", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID)+"/summary", nil, nil, false, "data", &out, opts)
	return out, err
}

// AmaRecap — Writes the recap post of an AMA.
//
// The AMA's host or a moderator; needs the AMA feature.
//
// POST /api/community/hubs/{slug}/posts/{postId}/recap
//
// slug: The hub's slug.
// postID: The post.
func (r *CommunityService) AmaRecap(ctx context.Context, slug string, postID string, opts ...CallOption) (*CommunityAmaRecapResponseData, error) {
	var out *CommunityAmaRecapResponseData
	err := r.c.call(ctx, "POST", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID)+"/recap", nil, nil, false, "data", &out, opts)
	return out, err
}

// ApprovePost — Publishes a post waiting for review, or rejects it.
//
// The hub's owner and moderators only.
//
// POST /api/community/hubs/{slug}/posts/{postId}/approve
//
// slug: The hub's slug.
// postID: The post.
func (r *CommunityService) ApprovePost(ctx context.Context, slug string, postID string, body *CommunityApprovePostBody, opts ...CallOption) (*CommunityApprovePostResponseData, error) {
	var out *CommunityApprovePostResponseData
	err := r.c.call(ctx, "POST", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID)+"/approve", nil, body, false, "data", &out, opts)
	return out, err
}

// ApproveReply — Publishes a reply waiting for review, or rejects it.
//
// The hub's owner and moderators only.
//
// POST /api/community/hubs/{slug}/replies/{replyId}/approve
//
// slug: The hub's slug.
// replyID: The reply.
func (r *CommunityService) ApproveReply(ctx context.Context, slug string, replyID string, body *CommunityApproveReplyBody, opts ...CallOption) (*CommunityApproveReplyResponseData, error) {
	var out *CommunityApproveReplyResponseData
	err := r.c.call(ctx, "POST", "/api/community/hubs/"+pathSeg(slug)+"/replies/"+pathSeg(replyID)+"/approve", nil, body, false, "data", &out, opts)
	return out, err
}

// BecomeMember — Joins a hub and answers the membership, or the join request when it needs approval.
//
// Like POST …/join, with the row in the answer. 409 when already a member or already asked.
//
// POST /api/community/hubs/{slug}/members
//
// slug: The hub's slug.
// The answer is one of *CommunityBecomeMemberResponseData, *CommunityBecomeMemberResponseData2, decoded as generic JSON (map[string]any, []any, …).
func (r *CommunityService) BecomeMember(ctx context.Context, slug string, body *CommunityBecomeMemberBody, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "POST", "/api/community/hubs/"+pathSeg(slug)+"/members", nil, body, false, "data", &out, opts)
	return out, err
}

// BookmarkPost — Saves a post to the caller's bookmarks.
//
// POST /api/community/hubs/{slug}/posts/{postId}/bookmark
//
// slug: The hub's slug.
// postID: The post.
func (r *CommunityService) BookmarkPost(ctx context.Context, slug string, postID string, opts ...CallOption) (*CommunityBookmarkPostResponse, error) {
	var out *CommunityBookmarkPostResponse
	err := r.c.call(ctx, "POST", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID)+"/bookmark", nil, nil, false, "", &out, opts)
	return out, err
}

// ClearAcceptedAnswer — Takes back a question's accepted answer.
//
// The question's author or a moderator.
//
// DELETE /api/community/hubs/{slug}/posts/{postId}/accept-answer
//
// slug: The hub's slug.
// postID: The post.
func (r *CommunityService) ClearAcceptedAnswer(ctx context.Context, slug string, postID string, opts ...CallOption) (*CommunityClearAcceptedAnswerResponse, error) {
	var out *CommunityClearAcceptedAnswerResponse
	err := r.c.call(ctx, "DELETE", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID)+"/accept-answer", nil, nil, false, "", &out, opts)
	return out, err
}

// CreateHub — Creates a hub with the caller as its owner; the slug is made from the name.
//
// Needs `community:create_hub`. An organisation's hub counts against its plan's hub limit. Five an hour.
//
// POST /api/community/hubs
func (r *CommunityService) CreateHub(ctx context.Context, body *CommunityCreateHubBody, opts ...CallOption) (*CommunityHub, error) {
	var out *CommunityHub
	err := r.c.call(ctx, "POST", "/api/community/hubs", nil, body, false, "data", &out, opts)
	return out, err
}

// CreateInvite — Creates an invitation link to a hub, mailed when an address is given.
//
// The hub's owner and moderators only. Only the owner invites moderators. 40 an hour.
//
// POST /api/community/hubs/{slug}/invites
//
// slug: The hub's slug.
func (r *CommunityService) CreateInvite(ctx context.Context, slug string, body *CommunityCreateInviteBody, opts ...CallOption) (*CommunityInvite, error) {
	var out *CommunityInvite
	err := r.c.call(ctx, "POST", "/api/community/hubs/"+pathSeg(slug)+"/invites", nil, body, false, "data", &out, opts)
	return out, err
}

// CreatePost — Posts in a hub.
//
// Active members only. With post approval or AI moderation on, it may wait for a moderator (`status: pending`). An AMA needs a moderator, the AMA feature and its start and end. Rate-limited.
//
// POST /api/community/hubs/{slug}/posts
//
// slug: The hub's slug.
func (r *CommunityService) CreatePost(ctx context.Context, slug string, body *CommunityCreatePostBody, opts ...CallOption) (*CommunityPost, error) {
	var out *CommunityPost
	err := r.c.call(ctx, "POST", "/api/community/hubs/"+pathSeg(slug)+"/posts", nil, body, false, "data", &out, opts)
	return out, err
}

// CreateReply — Replies to a post, or to a reply on it.
//
// Active members; a locked post only takes moderators' replies. Two levels deep at most. With moderation on it may wait for review (`pending`). The author and those mentioned are told.
//
// POST /api/community/hubs/{slug}/posts/{postId}/replies
//
// slug: The hub's slug.
// postID: The post.
func (r *CommunityService) CreateReply(ctx context.Context, slug string, postID string, body *CommunityCreateReplyBody, opts ...CallOption) (*CommunityCreateReplyResponseData, error) {
	var out *CommunityCreateReplyResponseData
	err := r.c.call(ctx, "POST", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID)+"/replies", nil, body, false, "data", &out, opts)
	return out, err
}

// CreateTag — Adds a tag to a hub.
//
// The hub's owner and moderators only. 409 when a tag of that name exists.
//
// POST /api/community/hubs/{slug}/tags
//
// slug: The hub's slug.
func (r *CommunityService) CreateTag(ctx context.Context, slug string, body *CommunityCreateTagBody, opts ...CallOption) (*CommunityTag, error) {
	var out *CommunityTag
	err := r.c.call(ctx, "POST", "/api/community/hubs/"+pathSeg(slug)+"/tags", nil, body, false, "data", &out, opts)
	return out, err
}

// DeleteDraft — Discards the caller's unfinished post in a hub.
//
// DELETE /api/community/hubs/{slug}/drafts
//
// slug: The hub's slug.
func (r *CommunityService) DeleteDraft(ctx context.Context, slug string, opts ...CallOption) (*CommunityDeleteDraftResponse, error) {
	var out *CommunityDeleteDraftResponse
	err := r.c.call(ctx, "DELETE", "/api/community/hubs/"+pathSeg(slug)+"/drafts", nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteHub — Deletes a hub with everything in it, for good.
//
// Its owner only.
//
// DELETE /api/community/hubs/{slug}
//
// slug: The hub's slug.
func (r *CommunityService) DeleteHub(ctx context.Context, slug string, opts ...CallOption) (*CommunityDeleteHubResponse, error) {
	var out *CommunityDeleteHubResponse
	err := r.c.call(ctx, "DELETE", "/api/community/hubs/"+pathSeg(slug), nil, nil, false, "", &out, opts)
	return out, err
}

// DeletePost — Deletes a post.
//
// Its author or a moderator.
//
// DELETE /api/community/hubs/{slug}/posts/{postId}
//
// slug: The hub's slug.
// postID: The post.
func (r *CommunityService) DeletePost(ctx context.Context, slug string, postID string, opts ...CallOption) (*CommunityDeletePostResponse, error) {
	var out *CommunityDeletePostResponse
	err := r.c.call(ctx, "DELETE", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteReply — Deletes a reply.
//
// Its author or a moderator.
//
// DELETE /api/community/hubs/{slug}/posts/{postId}/replies/{replyId}
//
// slug: The hub's slug.
// postID: The post.
// replyID: The reply.
func (r *CommunityService) DeleteReply(ctx context.Context, slug string, postID string, replyID string, opts ...CallOption) (*CommunityDeleteReplyResponse, error) {
	var out *CommunityDeleteReplyResponse
	err := r.c.call(ctx, "DELETE", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID)+"/replies/"+pathSeg(replyID), nil, nil, false, "", &out, opts)
	return out, err
}

// Feed — The newest posts of the hubs the caller is an active member of.
//
// GET /api/community/feed
func (r *CommunityService) Feed(ctx context.Context, query *CommunityFeedQuery, opts ...CallOption) ([]CommunityFeedResponseDataPostsItem, error) {
	var out []CommunityFeedResponseDataPostsItem
	err := r.c.call(ctx, "GET", "/api/community/feed", query, nil, false, "data.posts", &out, opts)
	return out, err
}

// GetDraft — The caller's unfinished post in a hub, or null.
//
// One draft per person and hub; active members only.
//
// GET /api/community/hubs/{slug}/drafts
//
// slug: The hub's slug.
func (r *CommunityService) GetDraft(ctx context.Context, slug string, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/community/hubs/"+pathSeg(slug)+"/drafts", nil, nil, false, "data.draft", &out, opts)
	return out, err
}

// GetHub — One hub, with the caller's role and status in it.
//
// A hub that is not public or unlisted only for its members (and, for an organisation or workspace hub, that organisation's or workspace's members).
//
// GET /api/community/hubs/{slug}
//
// slug: The hub's slug.
func (r *CommunityService) GetHub(ctx context.Context, slug string, opts ...CallOption) (*CommunityGetHubResponseData, error) {
	var out *CommunityGetHubResponseData
	err := r.c.call(ctx, "GET", "/api/community/hubs/"+pathSeg(slug), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetInvite — What an invitation leads to: the hub, who invited, the role and until when.
//
// GET /api/community/invites/{token}
//
// token: The invitation's token.
func (r *CommunityService) GetInvite(ctx context.Context, token string, opts ...CallOption) (*CommunityGetInviteResponseData, error) {
	var out *CommunityGetInviteResponseData
	err := r.c.call(ctx, "GET", "/api/community/invites/"+pathSeg(token), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetMemberBadges — A member's badges and reputation in a hub.
//
// Needs community gamification switched on. A hub that is not public or unlisted only for its active members.
//
// GET /api/community/hubs/{slug}/members/{memberId}/badges
//
// slug: The hub's slug.
// memberID: The member's user id.
func (r *CommunityService) GetMemberBadges(ctx context.Context, slug string, memberID string, opts ...CallOption) (*CommunityGetMemberBadgesResponseData, error) {
	var out *CommunityGetMemberBadgesResponseData
	err := r.c.call(ctx, "GET", "/api/community/hubs/"+pathSeg(slug)+"/members/"+pathSeg(memberID)+"/badges", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetMemberProfile — A member's profile in a hub: who they are, their role, badges and latest posts and replies.
//
// E-mail and the extended profile only where the member allows it (always for themselves and moderators).
//
// GET /api/community/hubs/{slug}/members/{memberId}/profile
//
// slug: The hub's slug.
// memberID: The member's user id.
func (r *CommunityService) GetMemberProfile(ctx context.Context, slug string, memberID string, opts ...CallOption) (*CommunityGetMemberProfileResponseData, error) {
	var out *CommunityGetMemberProfileResponseData
	err := r.c.call(ctx, "GET", "/api/community/hubs/"+pathSeg(slug)+"/members/"+pathSeg(memberID)+"/profile", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetPost — One post with its author, hub, reaction counts and what the caller did on it.
//
// Counts one view per person a day.
//
// GET /api/community/hubs/{slug}/posts/{postId}
//
// slug: The hub's slug.
// postID: The post.
func (r *CommunityService) GetPost(ctx context.Context, slug string, postID string, opts ...CallOption) (*CommunityGetPostResponseData, error) {
	var out *CommunityGetPostResponseData
	err := r.c.call(ctx, "GET", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID), nil, nil, false, "data", &out, opts)
	return out, err
}

// Join — Joins a hub, or asks to when it needs approval.
//
// A private or invitation-only hub is joined through an invitation (403).
//
// POST /api/community/hubs/{slug}/join
//
// slug: The hub's slug.
func (r *CommunityService) Join(ctx context.Context, slug string, opts ...CallOption) (*CommunityJoinResponseData, error) {
	var out *CommunityJoinResponseData
	err := r.c.call(ctx, "POST", "/api/community/hubs/"+pathSeg(slug)+"/join", nil, nil, false, "data", &out, opts)
	return out, err
}

// Leaderboard — A hub's members with the most reputation, with their top badges.
//
// Needs community gamification switched on. A hub that is not public or unlisted only for its active members.
//
// GET /api/community/hubs/{slug}/leaderboard
//
// slug: The hub's slug.
func (r *CommunityService) Leaderboard(ctx context.Context, slug string, query *CommunityLeaderboardQuery, opts ...CallOption) ([]CommunityLeaderboardResponseDataLeaderboardItem, error) {
	var out []CommunityLeaderboardResponseDataLeaderboardItem
	err := r.c.call(ctx, "GET", "/api/community/hubs/"+pathSeg(slug)+"/leaderboard", query, nil, false, "data.leaderboard", &out, opts)
	return out, err
}

// Leave — Leaves a hub.
//
// Its owner cannot leave: transfer it or delete it first.
//
// DELETE /api/community/hubs/{slug}/join
//
// slug: The hub's slug.
func (r *CommunityService) Leave(ctx context.Context, slug string, opts ...CallOption) (*CommunityLeaveResponse, error) {
	var out *CommunityLeaveResponse
	err := r.c.call(ctx, "DELETE", "/api/community/hubs/"+pathSeg(slug)+"/join", nil, nil, false, "", &out, opts)
	return out, err
}

// ListBookmarks — Lists the caller's bookmarked posts across hubs, newest first.
//
// At most 100.
//
// GET /api/community/bookmarks
func (r *CommunityService) ListBookmarks(ctx context.Context, opts ...CallOption) ([]CommunityListBookmarksResponseDataBookmarksItem, error) {
	var out []CommunityListBookmarksResponseDataBookmarksItem
	err := r.c.call(ctx, "GET", "/api/community/bookmarks", nil, nil, false, "data.bookmarks", &out, opts)
	return out, err
}

// ListHubs — Lists one page of the hubs the caller can see, biggest first: public ones and those they belong to.
//
// 20 a page. Archived hubs are left out.
//
// GET /api/community/hubs
func (r *CommunityService) ListHubs(ctx context.Context, query *CommunityListHubsQuery, opts ...CallOption) (*CommunityListHubsResponseData, error) {
	var out *CommunityListHubsResponseData
	err := r.c.call(ctx, "GET", "/api/community/hubs", query, nil, false, "data", &out, opts)
	return out, err
}

// ListInvites — Lists a hub's invitations.
//
// The hub's owner and moderators only.
//
// GET /api/community/hubs/{slug}/invites
//
// slug: The hub's slug.
func (r *CommunityService) ListInvites(ctx context.Context, slug string, opts ...CallOption) ([]CommunityInvite, error) {
	var out []CommunityInvite
	err := r.c.call(ctx, "GET", "/api/community/hubs/"+pathSeg(slug)+"/invites", nil, nil, false, "data.invites", &out, opts)
	return out, err
}

// ListJoinRequests — Lists a hub's open join requests, newest first.
//
// The hub's owner and moderators only.
//
// GET /api/community/hubs/{slug}/join-requests
//
// slug: The hub's slug.
func (r *CommunityService) ListJoinRequests(ctx context.Context, slug string, opts ...CallOption) ([]CommunityListJoinRequestsResponseDataRequestsItem, error) {
	var out []CommunityListJoinRequestsResponseDataRequestsItem
	err := r.c.call(ctx, "GET", "/api/community/hubs/"+pathSeg(slug)+"/join-requests", nil, nil, false, "data.requests", &out, opts)
	return out, err
}

// ListMembers — Lists one page of a hub's members, newest first, with their post counts and badges.
//
// 20 a page. Moderators may list other statuses and see e-mail addresses; everybody else sees active members.
//
// GET /api/community/hubs/{slug}/members
//
// slug: The hub's slug.
func (r *CommunityService) ListMembers(ctx context.Context, slug string, query *CommunityListMembersQuery, opts ...CallOption) (*CommunityListMembersResponseData, error) {
	var out *CommunityListMembersResponseData
	err := r.c.call(ctx, "GET", "/api/community/hubs/"+pathSeg(slug)+"/members", query, nil, false, "data", &out, opts)
	return out, err
}

// ListPosts — Lists one page of a hub's posts, pinned first.
//
// 20 a page. Poll options come with their counts and the caller's own vote (`userVotedIndex`, -1 for none). `status=pending` lists the review queue, for moderators.
//
// GET /api/community/hubs/{slug}/posts
//
// slug: The hub's slug.
func (r *CommunityService) ListPosts(ctx context.Context, slug string, query *CommunityListPostsQuery, opts ...CallOption) (*CommunityListPostsResponseData, error) {
	var out *CommunityListPostsResponseData
	err := r.c.call(ctx, "GET", "/api/community/hubs/"+pathSeg(slug)+"/posts", query, nil, false, "data", &out, opts)
	return out, err
}

// ListReplies — Lists one page of a post's replies, the accepted answer first, each with the replies under it.
//
// 20 top-level replies a page. A deleted reply that was answered stays as a placeholder.
//
// GET /api/community/hubs/{slug}/posts/{postId}/replies
//
// slug: The hub's slug.
// postID: The post.
func (r *CommunityService) ListReplies(ctx context.Context, slug string, postID string, query *CommunityListRepliesQuery, opts ...CallOption) (*CommunityListRepliesResponseData, error) {
	var out *CommunityListRepliesResponseData
	err := r.c.call(ctx, "GET", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID)+"/replies", query, nil, false, "data", &out, opts)
	return out, err
}

// ListReputationEvents — How a member's reputation in a hub came about, newest first.
//
// Needs community gamification switched on. The member themselves or a moderator.
//
// GET /api/community/hubs/{slug}/members/{memberId}/reputation-events
//
// slug: The hub's slug.
// memberID: The member's user id.
func (r *CommunityService) ListReputationEvents(ctx context.Context, slug string, memberID string, query *CommunityListReputationEventsQuery, opts ...CallOption) ([]CommunityListReputationEventsResponseDataEventsItem, error) {
	var out []CommunityListReputationEventsResponseDataEventsItem
	err := r.c.call(ctx, "GET", "/api/community/hubs/"+pathSeg(slug)+"/members/"+pathSeg(memberID)+"/reputation-events", query, nil, false, "data.events", &out, opts)
	return out, err
}

// ListTags — Lists a hub's tags.
//
// GET /api/community/hubs/{slug}/tags
//
// slug: The hub's slug.
func (r *CommunityService) ListTags(ctx context.Context, slug string, opts ...CallOption) ([]CommunityTag, error) {
	var out []CommunityTag
	err := r.c.call(ctx, "GET", "/api/community/hubs/"+pathSeg(slug)+"/tags", nil, nil, false, "data.tags", &out, opts)
	return out, err
}

// LockPost — Locks a post against new replies, or unlocks it.
//
// The hub's owner and moderators only. Each call flips it.
//
// POST /api/community/hubs/{slug}/posts/{postId}/lock
//
// slug: The hub's slug.
// postID: The post.
func (r *CommunityService) LockPost(ctx context.Context, slug string, postID string, opts ...CallOption) (*CommunityLockPostResponseData, error) {
	var out *CommunityLockPostResponseData
	err := r.c.call(ctx, "POST", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID)+"/lock", nil, nil, false, "data", &out, opts)
	return out, err
}

// Moderation — The caller's moderation work across the hubs they moderate: the review queue, open reports and recent activity.
//
// The tab named in `tab` comes back as the page asked for; the other two as a preview of five.
//
// GET /api/community/moderation
func (r *CommunityService) Moderation(ctx context.Context, query *CommunityModerationQuery, opts ...CallOption) (*CommunityModerationResponseData, error) {
	var out *CommunityModerationResponseData
	err := r.c.call(ctx, "GET", "/api/community/moderation", query, nil, false, "data", &out, opts)
	return out, err
}

// PinPost — Pins a post to the top of its hub, or unpins it.
//
// The hub's owner and moderators only. Each call flips it.
//
// POST /api/community/hubs/{slug}/posts/{postId}/pin
//
// slug: The hub's slug.
// postID: The post.
func (r *CommunityService) PinPost(ctx context.Context, slug string, postID string, opts ...CallOption) (*CommunityPinPostResponseData, error) {
	var out *CommunityPinPostResponseData
	err := r.c.call(ctx, "POST", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID)+"/pin", nil, nil, false, "data", &out, opts)
	return out, err
}

// ReactPost — Adds an emoji reaction to a post, or takes it back.
//
// Active members only. Each call with the same emoji flips it.
//
// POST /api/community/hubs/{slug}/posts/{postId}/react
//
// slug: The hub's slug.
// postID: The post.
func (r *CommunityService) ReactPost(ctx context.Context, slug string, postID string, body *CommunityReactPostBody, opts ...CallOption) (*CommunityReactPostResponseData, error) {
	var out *CommunityReactPostResponseData
	err := r.c.call(ctx, "POST", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID)+"/react", nil, body, false, "data", &out, opts)
	return out, err
}

// ReactReply — Adds an emoji reaction to a reply, or takes it back.
//
// Members only. Each call with the same emoji flips it.
//
// POST /api/community/hubs/{slug}/posts/{postId}/replies/{replyId}/react
//
// slug: The hub's slug.
// postID: The post.
// replyID: The reply.
func (r *CommunityService) ReactReply(ctx context.Context, slug string, postID string, replyID string, body *CommunityReactReplyBody, opts ...CallOption) (*CommunityReactReplyResponseData, error) {
	var out *CommunityReactReplyResponseData
	err := r.c.call(ctx, "POST", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID)+"/replies/"+pathSeg(replyID)+"/react", nil, body, false, "data", &out, opts)
	return out, err
}

// RemoveMember — Removes a member from a hub.
//
// A moderator, or the member themselves. The owner cannot be removed by anybody else.
//
// DELETE /api/community/hubs/{slug}/members/{memberId}
//
// slug: The hub's slug.
// memberID: The member's user id.
func (r *CommunityService) RemoveMember(ctx context.Context, slug string, memberID string, opts ...CallOption) (*CommunityRemoveMemberResponse, error) {
	var out *CommunityRemoveMemberResponse
	err := r.c.call(ctx, "DELETE", "/api/community/hubs/"+pathSeg(slug)+"/members/"+pathSeg(memberID), nil, nil, false, "", &out, opts)
	return out, err
}

// Report — Reports a post or reply to the hub's moderators.
//
// Members of the hub; not one's own content. Reporting the same thing again answers the open report (`duplicate: true`). 30 an hour.
//
// POST /api/community/reports
func (r *CommunityService) Report(ctx context.Context, body *CommunityReportBody, opts ...CallOption) (*CommunityReportResponseData, error) {
	var out *CommunityReportResponseData
	err := r.c.call(ctx, "POST", "/api/community/reports", nil, body, false, "data", &out, opts)
	return out, err
}

// ResolveJoinRequest — Lets somebody into a hub or turns them away; they are told either way.
//
// The hub's owner and moderators only. 409 when already decided.
//
// PUT /api/community/hubs/{slug}/join-requests/{requestId}
//
// slug: The hub's slug.
// requestID: The join request.
func (r *CommunityService) ResolveJoinRequest(ctx context.Context, slug string, requestID string, body *CommunityResolveJoinRequestBody, opts ...CallOption) (*CommunityResolveJoinRequestResponseData, error) {
	var out *CommunityResolveJoinRequestResponseData
	err := r.c.call(ctx, "PUT", "/api/community/hubs/"+pathSeg(slug)+"/join-requests/"+pathSeg(requestID), nil, body, false, "data", &out, opts)
	return out, err
}

// ResolveReport — Settles a report: dismiss it, remove the content, or ban its author from the hub.
//
// The hub's moderators. 409 when already settled.
//
// PATCH /api/community/moderation/reports/{id}
//
// id: The report.
func (r *CommunityService) ResolveReport(ctx context.Context, id string, body *CommunityResolveReportBody, opts ...CallOption) (*CommunityResolveReportResponseData, error) {
	var out *CommunityResolveReportResponseData
	err := r.c.call(ctx, "PATCH", "/api/community/moderation/reports/"+pathSeg(id), nil, body, false, "data", &out, opts)
	return out, err
}

// SaveDraft — Saves the caller's unfinished post in a hub, replacing the one before.
//
// PUT /api/community/hubs/{slug}/drafts
//
// slug: The hub's slug.
func (r *CommunityService) SaveDraft(ctx context.Context, slug string, body *CommunitySaveDraftBody, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "PUT", "/api/community/hubs/"+pathSeg(slug)+"/drafts", nil, body, false, "data.draft", &out, opts)
	return out, err
}

// Search — Searches the posts of public hubs and of the hubs the caller belongs to, newest first.
//
// At most 30. An unlisted hub is searched only when named in `hub`. Fewer than 2 characters answer nothing.
//
// GET /api/community/search
func (r *CommunityService) Search(ctx context.Context, query *CommunitySearchQuery, opts ...CallOption) ([]CommunitySearchResponseDataPostsItem, error) {
	var out []CommunitySearchResponseDataPostsItem
	err := r.c.call(ctx, "GET", "/api/community/search", query, nil, false, "data.posts", &out, opts)
	return out, err
}

// TransferHub — Hands a hub to another of its active members; the old owner becomes a moderator.
//
// Its owner only.
//
// POST /api/community/hubs/{slug}/transfer
//
// slug: The hub's slug.
func (r *CommunityService) TransferHub(ctx context.Context, slug string, body *CommunityTransferHubBody, opts ...CallOption) (*CommunityTransferHubResponse, error) {
	var out *CommunityTransferHubResponse
	err := r.c.call(ctx, "POST", "/api/community/hubs/"+pathSeg(slug)+"/transfer", nil, body, false, "", &out, opts)
	return out, err
}

// UnbookmarkPost — Removes a post from the caller's bookmarks.
//
// DELETE /api/community/hubs/{slug}/posts/{postId}/bookmark
//
// slug: The hub's slug.
// postID: The post.
func (r *CommunityService) UnbookmarkPost(ctx context.Context, slug string, postID string, opts ...CallOption) (*CommunityUnbookmarkPostResponse, error) {
	var out *CommunityUnbookmarkPostResponse
	err := r.c.call(ctx, "DELETE", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID)+"/bookmark", nil, nil, false, "", &out, opts)
	return out, err
}

// UpdateHub — Changes a hub: its texts, look, visibility, joining, moderation settings, slug or archive state.
//
// The hub's owner and moderators only. Only the fields sent change; `settings` is merged into what is stored. A new slug must be free (409).
//
// PUT /api/community/hubs/{slug}
//
// slug: The hub's slug.
func (r *CommunityService) UpdateHub(ctx context.Context, slug string, body *CommunityUpdateHubBody, opts ...CallOption) (*CommunityHub, error) {
	var out *CommunityHub
	err := r.c.call(ctx, "PUT", "/api/community/hubs/"+pathSeg(slug), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateMember — Changes a member's role, bans or unbans them, or mutes them for a while.
//
// The hub's owner and moderators only. Only the owner can touch the owner or make somebody owner.
//
// PUT /api/community/hubs/{slug}/members/{memberId}
//
// slug: The hub's slug.
// memberID: The member's user id.
func (r *CommunityService) UpdateMember(ctx context.Context, slug string, memberID string, body *CommunityUpdateMemberBody, opts ...CallOption) (*CommunityMember, error) {
	var out *CommunityMember
	err := r.c.call(ctx, "PUT", "/api/community/hubs/"+pathSeg(slug)+"/members/"+pathSeg(memberID), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdatePost — Edits a post.
//
// Its author within 30 minutes of posting, or a moderator at any time.
//
// PUT /api/community/hubs/{slug}/posts/{postId}
//
// slug: The hub's slug.
// postID: The post.
func (r *CommunityService) UpdatePost(ctx context.Context, slug string, postID string, body *CommunityUpdatePostBody, opts ...CallOption) (*CommunityPost, error) {
	var out *CommunityPost
	err := r.c.call(ctx, "PUT", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateReply — Edits a reply.
//
// Its author only.
//
// PUT /api/community/hubs/{slug}/posts/{postId}/replies/{replyId}
//
// slug: The hub's slug.
// postID: The post.
// replyID: The reply.
func (r *CommunityService) UpdateReply(ctx context.Context, slug string, postID string, replyID string, body *CommunityUpdateReplyBody, opts ...CallOption) (*CommunityReply, error) {
	var out *CommunityReply
	err := r.c.call(ctx, "PUT", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID)+"/replies/"+pathSeg(replyID), nil, body, false, "data", &out, opts)
	return out, err
}

// UpvotePost — Upvotes a post, or takes the upvote back.
//
// Active members only. Each call flips it.
//
// POST /api/community/hubs/{slug}/posts/{postId}/upvote
//
// slug: The hub's slug.
// postID: The post.
func (r *CommunityService) UpvotePost(ctx context.Context, slug string, postID string, opts ...CallOption) (*CommunityUpvotePostResponseData, error) {
	var out *CommunityUpvotePostResponseData
	err := r.c.call(ctx, "POST", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID)+"/upvote", nil, nil, false, "data", &out, opts)
	return out, err
}

// UpvoteReply — Upvotes a reply, or takes the upvote back.
//
// Active members only. Each call flips it.
//
// POST /api/community/hubs/{slug}/posts/{postId}/replies/{replyId}/upvote
//
// slug: The hub's slug.
// postID: The post.
// replyID: The reply.
func (r *CommunityService) UpvoteReply(ctx context.Context, slug string, postID string, replyID string, opts ...CallOption) (*CommunityUpvoteReplyResponseData, error) {
	var out *CommunityUpvoteReplyResponseData
	err := r.c.call(ctx, "POST", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID)+"/replies/"+pathSeg(replyID)+"/upvote", nil, nil, false, "data", &out, opts)
	return out, err
}

// VotePoll — Votes in a poll; voting for the same option again takes the vote back.
//
// Members; in a public, open hub voting joins it. Someone banned from the hub is refused (403).
//
// POST /api/community/hubs/{slug}/posts/{postId}/vote
//
// slug: The hub's slug.
// postID: The post.
func (r *CommunityService) VotePoll(ctx context.Context, slug string, postID string, body *CommunityVotePollBody, opts ...CallOption) (*CommunityVotePollResponseData, error) {
	var out *CommunityVotePollResponseData
	err := r.c.call(ctx, "POST", "/api/community/hubs/"+pathSeg(slug)+"/posts/"+pathSeg(postID)+"/vote", nil, body, false, "data", &out, opts)
	return out, err
}

// ConnectionsService — Connections: Credentials for the services workflows talk to.
type ConnectionsService struct {
	c *Client
}

func newConnectionsService(c *Client) *ConnectionsService {
	return &ConnectionsService{c: c}
}

// Create — Stores credentials for an integration as a connection.
//
// POST /api/connections
func (r *ConnectionsService) Create(ctx context.Context, body *CreateConnectionInput, opts ...CallOption) (*Connection, error) {
	var out *Connection
	err := r.c.call(ctx, "POST", "/api/connections", nil, body, false, "", &out, opts)
	return out, err
}

// Delete — Deletes a connection; workflows using it stop working (see `connections.dependents`).
//
// DELETE /api/connections/{connectionId}
//
// connectionID: The connection's id.
func (r *ConnectionsService) Delete(ctx context.Context, connectionID string, opts ...CallOption) (*ConnectionsDeleteResponse, error) {
	var out *ConnectionsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/connections/"+pathSeg(connectionID), nil, nil, false, "", &out, opts)
	return out, err
}

// Dependents — Lists the workflows that use a connection.
//
// GET /api/connections/{connectionId}/dependents
//
// connectionID: The connection's id.
func (r *ConnectionsService) Dependents(ctx context.Context, connectionID string, opts ...CallOption) ([]ConnectionsDependentsResponseItem, error) {
	var out []ConnectionsDependentsResponseItem
	err := r.c.call(ctx, "GET", "/api/connections/"+pathSeg(connectionID)+"/dependents", nil, nil, false, "", &out, opts)
	return out, err
}

// Get — Returns one connection, with secrets masked.
//
// GET /api/connections/{connectionId}
//
// connectionID: The connection's id.
func (r *ConnectionsService) Get(ctx context.Context, connectionID string, opts ...CallOption) (*Connection, error) {
	var out *Connection
	err := r.c.call(ctx, "GET", "/api/connections/"+pathSeg(connectionID), nil, nil, false, "", &out, opts)
	return out, err
}

// GetOAuthURL — Returns the address where a person authorizes an OAuth connection with the provider.
//
// GET /api/connections/{connectionId}/oauth-url
//
// connectionID: The connection's id.
func (r *ConnectionsService) GetOAuthURL(ctx context.Context, connectionID string, query *ConnectionsGetOAuthUrlQuery, opts ...CallOption) (string, error) {
	var out string
	err := r.c.call(ctx, "GET", "/api/connections/"+pathSeg(connectionID)+"/oauth-url", query, nil, false, "url", &out, opts)
	return out, err
}

// List — Lists a workspace's connections, with secrets masked.
//
// The caller's personal connections are included; other people's are not.
//
// GET /api/connections
func (r *ConnectionsService) List(ctx context.Context, query *ConnectionsListQuery, opts ...CallOption) ([]Connection, error) {
	var out []Connection
	err := r.c.call(ctx, "GET", "/api/connections", query, nil, false, "", &out, opts)
	return out, err
}

// ListRequirements — Lists what each integration needs to connect: credential types and fields.
//
// GET /api/connections/requirements
func (r *ConnectionsService) ListRequirements(ctx context.Context, opts ...CallOption) ([]ConnectionsListRequirementsResponseItem, error) {
	var out []ConnectionsListRequirementsResponseItem
	err := r.c.call(ctx, "GET", "/api/connections/requirements", nil, nil, false, "", &out, opts)
	return out, err
}

// Refresh — Refreshes an OAuth connection's access token when it is about to expire.
//
// POST /api/connections/{connectionId}/refresh
//
// connectionID: The connection's id.
func (r *ConnectionsService) Refresh(ctx context.Context, connectionID string, opts ...CallOption) (*ConnectionsRefreshResponse, error) {
	var out *ConnectionsRefreshResponse
	err := r.c.call(ctx, "POST", "/api/connections/"+pathSeg(connectionID)+"/refresh", nil, nil, false, "", &out, opts)
	return out, err
}

// Test — Checks that a connection's credentials work.
//
// POST /api/connections/{connectionId}/test
//
// connectionID: The connection's id.
func (r *ConnectionsService) Test(ctx context.Context, connectionID string, opts ...CallOption) (*ConnectionsTestResponse, error) {
	var out *ConnectionsTestResponse
	err := r.c.call(ctx, "POST", "/api/connections/"+pathSeg(connectionID)+"/test", nil, nil, false, "", &out, opts)
	return out, err
}

// Update — Renames a connection or changes its credentials.
//
// PATCH /api/connections/{connectionId}
//
// connectionID: The connection's id.
func (r *ConnectionsService) Update(ctx context.Context, connectionID string, body *UpdateConnectionInput, opts ...CallOption) (*Connection, error) {
	var out *Connection
	err := r.c.call(ctx, "PATCH", "/api/connections/"+pathSeg(connectionID), nil, body, false, "", &out, opts)
	return out, err
}

// CountriesService — Countries
type CountriesService struct {
	c *Client
}

func newCountriesService(c *Client) *CountriesService {
	return &CountriesService{c: c}
}

// List — Lists the countries the platform offers, with each one's regional defaults and tax regime.
//
// Answers without a token. Only enabled countries, in the administrator's order; cached for five minutes.
//
// GET /api/countries — No token needed.
func (r *CountriesService) List(ctx context.Context, opts ...CallOption) ([]Country, error) {
	var out []Country
	err := r.c.call(ctx, "GET", "/api/countries", nil, nil, false, "", &out, opts)
	return out, err
}

// CreatorsService — Creator profiles: A person's creator profile — what they chose to show, their marketplace listings, the hubs and templates they share with the caller, and their community reputation. Not anonymous: every call needs a signed-in person or a token; never an address, a phone number or an organisation membership.
type CreatorsService struct {
	c *Client
}

func newCreatorsService(c *Client) *CreatorsService {
	return &CreatorsService{c: c}
}

// Faces — Lists where else a person shows up that the caller can see: their shop, shared hubs, templates and reports.
//
// A profile is seen by its owner; by anybody signed in when public; by people of the same organisation when shared with the organisation. Otherwise 404, as if there were none.
//
// GET /api/public/creator/{handle}/faces
//
// handle: The person's handle (1–32 letters, digits, `-` or `_`); case does not matter.
func (r *CreatorsService) Faces(ctx context.Context, handle string, opts ...CallOption) (*CreatorFaces, error) {
	var out *CreatorFaces
	err := r.c.call(ctx, "GET", "/api/public/creator/"+pathSeg(handle)+"/faces", nil, nil, false, "data", &out, opts)
	return out, err
}

// Get — Returns a person's creator profile by handle.
//
// A profile is seen by its owner; by anybody signed in when public; by people of the same organisation when shared with the organisation. Otherwise 404, as if there were none.
//
// GET /api/public/creator/{handle}
//
// handle: The person's handle (1–32 letters, digits, `-` or `_`); case does not matter.
func (r *CreatorsService) Get(ctx context.Context, handle string, opts ...CallOption) (*CreatorProfile, error) {
	var out *CreatorProfile
	err := r.c.call(ctx, "GET", "/api/public/creator/"+pathSeg(handle), nil, nil, false, "", &out, opts)
	return out, err
}

// RegenerateAbout — Writes the short introduction on the caller's own profile again with the platform's AI, from their listings, badges and bio.
//
// The profile's owner only, and only with an organisation (403 `no_organization` without one). Needs the AI introduction on their organisation's plan and its AI allowance (402 when used up); 4 an hour. 400 `creatorAboutNotEnoughSignal` with no listing, badge or bio to write from. The text is shown again (if it was hidden) and marked as written by AI.
//
// POST /api/public/creator/{handle}/about
//
// handle: The person's handle (1–32 letters, digits, `-` or `_`); case does not matter.
func (r *CreatorsService) RegenerateAbout(ctx context.Context, handle string, opts ...CallOption) (*CreatorsRegenerateAboutResponseData, error) {
	var out *CreatorsRegenerateAboutResponseData
	err := r.c.call(ctx, "POST", "/api/public/creator/"+pathSeg(handle)+"/about", nil, nil, false, "data", &out, opts)
	return out, err
}

// Reputation — Returns a person's reputation per community hub and their latest reputation events, in the hubs the caller also belongs to.
//
// A profile is seen by its owner; by anybody signed in when public; by people of the same organisation when shared with the organisation. Otherwise 404, as if there were none. The profile's owner sees every hub of theirs.
//
// GET /api/public/creator/{handle}/reputation-events
//
// handle: The person's handle (1–32 letters, digits, `-` or `_`); case does not matter.
func (r *CreatorsService) Reputation(ctx context.Context, handle string, opts ...CallOption) (*CreatorReputation, error) {
	var out *CreatorReputation
	err := r.c.call(ctx, "GET", "/api/public/creator/"+pathSeg(handle)+"/reputation-events", nil, nil, false, "data", &out, opts)
	return out, err
}

// SetAboutHidden — Hides the AI-written introduction on the caller's own profile, or shows it again.
//
// The profile's owner only.
//
// PATCH /api/public/creator/{handle}/about
//
// handle: The person's handle (1–32 letters, digits, `-` or `_`); case does not matter.
func (r *CreatorsService) SetAboutHidden(ctx context.Context, handle string, body *CreatorsSetAboutHiddenBody, opts ...CallOption) (*CreatorsSetAboutHiddenResponse, error) {
	var out *CreatorsSetAboutHiddenResponse
	err := r.c.call(ctx, "PATCH", "/api/public/creator/"+pathSeg(handle)+"/about", nil, body, false, "", &out, opts)
	return out, err
}

// CrmService — CRM
type CrmService struct {
	c *Client

	Actions            *CrmActionsService
	Activities         *CrmActivitiesService
	Brief              *CrmBriefService
	CaptureRules       *CrmCaptureRulesService
	Consents           *CrmConsentsService
	DealRoom           *CrmDealRoomService
	Duplicates         *CrmDuplicatesService
	Enrichment         *CrmEnrichmentService
	Groups             *CrmGroupsService
	Journey            *CrmJourneyService
	Matches            *CrmMatchesService
	Matching           *CrmMatchingService
	Meetings           *CrmMeetingsService
	OfflineConversions *CrmOfflineConversionsService
	Pipeline           *CrmPipelineService
	Playbooks          *CrmPlaybooksService
	Portal             *CrmPortalService
	Records            *CrmRecordsService
	Sales              *CrmSalesService
	Spaces             *CrmSpacesService
	Tasks              *CrmTasksService
	Website            *CrmWebsiteService
}

func newCrmService(c *Client) *CrmService {
	return &CrmService{c: c, Actions: newCrmActionsService(c), Activities: newCrmActivitiesService(c), Brief: newCrmBriefService(c), CaptureRules: newCrmCaptureRulesService(c), Consents: newCrmConsentsService(c), DealRoom: newCrmDealRoomService(c), Duplicates: newCrmDuplicatesService(c), Enrichment: newCrmEnrichmentService(c), Groups: newCrmGroupsService(c), Journey: newCrmJourneyService(c), Matches: newCrmMatchesService(c), Matching: newCrmMatchingService(c), Meetings: newCrmMeetingsService(c), OfflineConversions: newCrmOfflineConversionsService(c), Pipeline: newCrmPipelineService(c), Playbooks: newCrmPlaybooksService(c), Portal: newCrmPortalService(c), Records: newCrmRecordsService(c), Sales: newCrmSalesService(c), Spaces: newCrmSpacesService(c), Tasks: newCrmTasksService(c), Website: newCrmWebsiteService(c)}
}

// Connected — What the other hubs know about this person — mail, form submissions, bookings, meetings, signatures — newest first.
//
// Read live for the caller: per hub its own rules apply, and nothing they may not see there is shown or counted. `failedSources` names a hub that could not be read.
//
// GET /api/crm/spaces/{spaceId}/records/{recordId}/connected — Token permission: read:crm.
//
// spaceID: The space.
// recordID: The record.
func (r *CrmService) Connected(ctx context.Context, spaceID string, recordID string, query *CrmConnectedQuery, opts ...CallOption) (*CrmConnectedResponse, error) {
	var out *CrmConnectedResponse
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/connected", query, nil, false, "", &out, opts)
	return out, err
}

// Mailboxes — The caller's mailboxes of the workspace's organisation, each with whether it shows in the CRM.
//
// Only a mailbox's owner switches it, in the Mail API.
//
// GET /api/crm/mailboxes — Token permission: read:crm.
func (r *CrmService) Mailboxes(ctx context.Context, query *CrmMailboxesQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/crm/mailboxes", query, nil, false, "mailboxes", &out, opts)
	return out, err
}

// CrmActionsService — crm.actions
type CrmActionsService struct {
	c *Client
}

func newCrmActionsService(c *Client) *CrmActionsService {
	return &CrmActionsService{c: c}
}

// BookingLink — Shares a booking link with the person: prefilled in the composer, as a link to copy, or as a mail draft in the caller's mailbox.
//
// Needs edit level. Never sends mail by itself. The record's timeline notes it.
//
// POST /api/crm/spaces/{spaceId}/records/{recordId}/actions/booking-link — Token permission: write:crm.
//
// spaceID: The space.
// recordID: The record.
func (r *CrmActionsService) BookingLink(ctx context.Context, spaceID string, recordID string, body *CrmActionsBookingLinkBody, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/actions/booking-link", nil, body, false, "result", &out, opts)
	return out, err
}

// Document — Creates a document from a template with the record's values filled in, private to the caller.
//
// Needs edit level and `document:create`; counts against the plan's document limit.
//
// POST /api/crm/spaces/{spaceId}/records/{recordId}/actions/document — Token permission: write:crm.
//
// spaceID: The space.
// recordID: The record.
func (r *CrmActionsService) Document(ctx context.Context, spaceID string, recordID string, body *CrmActionsDocumentBody, opts ...CallOption) (*CrmActionsDocumentResponseDocument, error) {
	var out *CrmActionsDocumentResponseDocument
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/actions/document", nil, body, false, "document", &out, opts)
	return out, err
}

// Get — What a record's actions may offer now: the person they reach, the booking pages to share, the templates a document may start from, whether signing is possible.
//
// Needs edit level on the record (403 otherwise).
//
// GET /api/crm/spaces/{spaceId}/records/{recordId}/actions — Token permission: read:crm.
//
// spaceID: The space.
// recordID: The record.
func (r *CrmActionsService) Get(ctx context.Context, spaceID string, recordID string, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/actions", nil, nil, false, "actions", &out, opts)
	return out, err
}

// RequestSignature — Sends a document for signing — to the record's person, or the signers named.
//
// Needs edit level and e-signature on the plan. The document is rendered as a PDF and sent through the e-signature hub; the record's timeline notes it.
//
// POST /api/crm/spaces/{spaceId}/records/{recordId}/actions/signature — Token permission: write:crm.
//
// spaceID: The space.
// recordID: The record.
func (r *CrmActionsService) RequestSignature(ctx context.Context, spaceID string, recordID string, body *CrmActionsRequestSignatureBody, opts ...CallOption) (string, error) {
	var out string
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/actions/signature", nil, body, false, "requestId", &out, opts)
	return out, err
}

// SignatureOptions — The documents the caller may send for signing to the record's person.
//
// GET /api/crm/spaces/{spaceId}/records/{recordId}/actions/signature — Token permission: read:crm.
//
// spaceID: The space.
// recordID: The record.
func (r *CrmActionsService) SignatureOptions(ctx context.Context, spaceID string, recordID string, query *CrmActionsSignatureOptionsQuery, opts ...CallOption) (*CrmActionsSignatureOptionsResponse, error) {
	var out *CrmActionsSignatureOptionsResponse
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/actions/signature", query, nil, false, "", &out, opts)
	return out, err
}

// CrmActivitiesService — crm.activities
type CrmActivitiesService struct {
	c *Client
}

func newCrmActivitiesService(c *Client) *CrmActivitiesService {
	return &CrmActivitiesService{c: c}
}

// Delete — Removes an activity; entries the system recorded cannot be removed.
//
// DELETE /api/crm/spaces/{spaceId}/records/{recordId}/activities/{activityId} — Token permission: write:crm.
func (r *CrmActivitiesService) Delete(ctx context.Context, spaceID string, recordID string, activityID string, opts ...CallOption) (*CrmActivitiesDeleteResponse, error) {
	var out *CrmActivitiesDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/activities/"+pathSeg(activityID), nil, nil, false, "", &out, opts)
	return out, err
}

// Log — Log an activity — a call, meeting, email, note or one of the space's own types — with the answers to its question flow.
//
// POST /api/crm/spaces/{spaceId}/records/{recordId}/activities — Token permission: write:crm.
func (r *CrmActivitiesService) Log(ctx context.Context, spaceID string, recordID string, body *CrmLogActivityInput, opts ...CallOption) (*CrmActivitiesLogResponse, error) {
	var out *CrmActivitiesLogResponse
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/activities", nil, body, false, "", &out, opts)
	return out, err
}

// LogCall — Log a call — direction, outcome, duration, a note, the call flow's answers, and optionally a follow-up task.
//
// POST /api/crm/spaces/{spaceId}/records/{recordId}/calls — Token permission: write:crm.
func (r *CrmActivitiesService) LogCall(ctx context.Context, spaceID string, recordID string, body *CrmLogCallInput, opts ...CallOption) (*CrmActivitiesLogCallResponse, error) {
	var out *CrmActivitiesLogCallResponse
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/calls", nil, body, false, "", &out, opts)
	return out, err
}

// Timeline — One page of a record's timeline, newest first.
//
// GET /api/crm/spaces/{spaceId}/records/{recordId}/timeline — Token permission: read:crm.
func (r *CrmActivitiesService) Timeline(ctx context.Context, spaceID string, recordID string, query *CrmActivitiesTimelineQuery, opts ...CallOption) ([]CrmTimelineEntry, error) {
	var out []CrmTimelineEntry
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/timeline", query, nil, false, "entries", &out, opts)
	return out, err
}

// Update — Change an activity (its author, or a manager of the space).
//
// PATCH /api/crm/spaces/{spaceId}/records/{recordId}/activities/{activityId} — Token permission: write:crm.
func (r *CrmActivitiesService) Update(ctx context.Context, spaceID string, recordID string, activityID string, body *CrmActivitiesUpdateBody, opts ...CallOption) (*CrmActivitiesUpdateResponse, error) {
	var out *CrmActivitiesUpdateResponse
	err := r.c.call(ctx, "PATCH", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/activities/"+pathSeg(activityID), nil, body, false, "", &out, opts)
	return out, err
}

// CrmBriefService — crm.brief
type CrmBriefService struct {
	c *Client
}

func newCrmBriefService(c *Client) *CrmBriefService {
	return &CrmBriefService{c: c}
}

// Get — "Brief me": a short briefing before a call or meeting — who, last contact, open deals and tasks, what to mention, risks — written by the platform's AI.
//
// Nothing is stored. Written only from what the caller may see; the AI allowance applies (429).
//
// POST /api/crm/spaces/{spaceId}/records/{recordId}/brief — Token permission: write:crm.
//
// spaceID: The space.
// recordID: The record.
func (r *CrmBriefService) Get(ctx context.Context, spaceID string, recordID string, body *CrmBriefGetBody, opts ...CallOption) (*CrmBriefing, error) {
	var out *CrmBriefing
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/brief", nil, body, false, "briefing", &out, opts)
	return out, err
}

// SaveNote — Saves a briefing as a note on the record, marked AI-assisted.
//
// Needs edit level on the space.
//
// POST /api/crm/spaces/{spaceId}/records/{recordId}/brief/note — Token permission: write:crm.
//
// spaceID: The space.
// recordID: The record.
func (r *CrmBriefService) SaveNote(ctx context.Context, spaceID string, recordID string, body *CrmBriefSaveNoteBody, opts ...CallOption) (*CrmTimelineEntry, error) {
	var out *CrmTimelineEntry
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/brief/note", nil, body, false, "entry", &out, opts)
	return out, err
}

// CrmCaptureRulesService — crm.captureRules
type CrmCaptureRulesService struct {
	c *Client
}

func newCrmCaptureRulesService(c *Client) *CrmCaptureRulesService {
	return &CrmCaptureRulesService{c: c}
}

// Create — Creates an automatic-contact rule: what arrives from a source (a form, a waitlist…) becomes people — and optionally companies, deals and tasks — in the CRM.
//
// Needs `crm:manage` in the workspace, edit level on every space it writes to, the source readable by the caller, and automatic contacts on the plan (counts against its rule limit). The rule runs as the caller.
//
// POST /api/crm/capture-rules — Token permission: write:crm.
func (r *CrmCaptureRulesService) Create(ctx context.Context, body *CrmCaptureRulesCreateBody, opts ...CallOption) (*CrmCaptureRule, error) {
	var out *CrmCaptureRule
	err := r.c.call(ctx, "POST", "/api/crm/capture-rules", nil, body, false, "rule", &out, opts)
	return out, err
}

// Delete — Deletes an automatic-contact rule.
//
// DELETE /api/crm/capture-rules/{ruleId} — Token permission: write:crm.
func (r *CrmCaptureRulesService) Delete(ctx context.Context, ruleID string, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "DELETE", "/api/crm/capture-rules/"+pathSeg(ruleID), nil, nil, false, "", &out, opts)
	return out, err
}

// List — The workspace's automatic-contact rules, each with its source, the spaces it writes to, its last run and its counts.
//
// GET /api/crm/capture-rules — Token permission: read:crm.
func (r *CrmCaptureRulesService) List(ctx context.Context, query *CrmCaptureRulesListQuery, opts ...CallOption) (*CrmCaptureRulesListResponse, error) {
	var out *CrmCaptureRulesListResponse
	err := r.c.call(ctx, "GET", "/api/crm/capture-rules", query, nil, false, "", &out, opts)
	return out, err
}

// Sources — What a rule may point at — sources, their items, the spaces.
//
// GET /api/crm/capture-rules/sources — Token permission: read:crm.
func (r *CrmCaptureRulesService) Sources(ctx context.Context, query *CrmCaptureRulesSourcesQuery, opts ...CallOption) (*CrmCaptureChoices, error) {
	var out *CrmCaptureChoices
	err := r.c.call(ctx, "GET", "/api/crm/capture-rules/sources", query, nil, false, "choices", &out, opts)
	return out, err
}

// Update — Change a rule — `{ enabled: false }` alone switches it off.
//
// PATCH /api/crm/capture-rules/{ruleId} — Token permission: write:crm.
func (r *CrmCaptureRulesService) Update(ctx context.Context, ruleID string, body *CrmCaptureRulesUpdateBody, opts ...CallOption) (*CrmCaptureRule, error) {
	var out *CrmCaptureRule
	err := r.c.call(ctx, "PATCH", "/api/crm/capture-rules/"+pathSeg(ruleID), nil, body, false, "rule", &out, opts)
	return out, err
}

// CrmConsentsService — crm.consents
type CrmConsentsService struct {
	c *Client
}

func newCrmConsentsService(c *Client) *CrmConsentsService {
	return &CrmConsentsService{c: c}
}

// ForSpace — Consent status per record and channel, and the counts.
//
// GET /api/crm/spaces/{spaceId}/consents — Token permission: read:crm.
func (r *CrmConsentsService) ForSpace(ctx context.Context, spaceID string, opts ...CallOption) (*CrmConsentsForSpaceResponse, error) {
	var out *CrmConsentsForSpaceResponse
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/consents", nil, nil, false, "", &out, opts)
	return out, err
}

// Get — A person's consent ledger.
//
// GET /api/crm/spaces/{spaceId}/records/{recordId}/consents — Token permission: read:crm.
func (r *CrmConsentsService) Get(ctx context.Context, spaceID string, recordID string, opts ...CallOption) (*CrmConsentsGetResponse, error) {
	var out *CrmConsentsGetResponse
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/consents", nil, nil, false, "", &out, opts)
	return out, err
}

// Grant — Records a consent the person gave (`action: grant`), or that they withdrew one (`action: revoke`).
//
// Needs edit level on the space. A consent recorded here is marked manual, with the caller as the actor.
//
// POST /api/crm/spaces/{spaceId}/records/{recordId}/consents — Token permission: write:crm.
//
// body: one of *CrmConsentsGrantBody, *CrmConsentsGrantBody2.
func (r *CrmConsentsService) Grant(ctx context.Context, spaceID string, recordID string, body any, opts ...CallOption) (string, error) {
	var out string
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/consents", nil, body, false, "consentId", &out, opts)
	return out, err
}

// Reconsent — Asks the person to confirm their consent by double opt-in: a mail with a one-time link in their language.
//
// Needs edit level. E-mail channels only; refused for a person whose latest decision is a withdrawal they made themselves. 30 an hour.
//
// POST /api/crm/spaces/{spaceId}/records/{recordId}/consents/reconsent — Token permission: write:crm.
//
// spaceID: The space.
// recordID: The record.
func (r *CrmConsentsService) Reconsent(ctx context.Context, spaceID string, recordID string, body *CrmConsentsReconsentBody, opts ...CallOption) (*CrmConsentsReconsentResponse, error) {
	var out *CrmConsentsReconsentResponse
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/consents/reconsent", nil, body, false, "", &out, opts)
	return out, err
}

// CrmDealRoomService — crm.dealRoom
type CrmDealRoomService struct {
	c *Client
}

func newCrmDealRoomService(c *Client) *CrmDealRoomService {
	return &CrmDealRoomService{c: c}
}

// AddStep — Add a step to the mutual action plan.
//
// POST /api/crm/spaces/{spaceId}/records/{dealId}/deal-room/steps — Token permission: write:crm.
func (r *CrmDealRoomService) AddStep(ctx context.Context, spaceID string, dealID string, body *CrmDealRoomStepInput, opts ...CallOption) (*CrmDealRoomStep, error) {
	var out *CrmDealRoomStep
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(dealID)+"/deal-room/steps", nil, body, false, "step", &out, opts)
	return out, err
}

// DeleteStep — Remove a step.
//
// DELETE /api/crm/spaces/{spaceId}/records/{dealId}/deal-room/steps/{stepId} — Token permission: write:crm.
func (r *CrmDealRoomService) DeleteStep(ctx context.Context, spaceID string, dealID string, stepID string, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "DELETE", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(dealID)+"/deal-room/steps/"+pathSeg(stepID), nil, nil, false, "", &out, opts)
	return out, err
}

// Get — The room: plan, shared items, thread.
//
// GET /api/crm/spaces/{spaceId}/records/{dealId}/deal-room — Token permission: read:crm.
func (r *CrmDealRoomService) Get(ctx context.Context, spaceID string, dealID string, opts ...CallOption) (*CrmDealRoomState, error) {
	var out *CrmDealRoomState
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(dealID)+"/deal-room", nil, nil, false, "room", &out, opts)
	return out, err
}

// Message — Writes a message to the customer, or a comment on a step, in the deal room; plain text.
//
// POST /api/crm/spaces/{spaceId}/records/{dealId}/deal-room/messages — Token permission: write:crm.
//
// body: one of *CrmDealRoomMessageBody, *CrmDealRoomMessageBody2.
func (r *CrmDealRoomService) Message(ctx context.Context, spaceID string, dealID string, body any, opts ...CallOption) (*CrmDealRoomMessage, error) {
	var out *CrmDealRoomMessage
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(dealID)+"/deal-room/messages", nil, body, false, "message", &out, opts)
	return out, err
}

// Options — What could be shared, of one kind.
//
// GET /api/crm/spaces/{spaceId}/records/{dealId}/deal-room/options — Token permission: read:crm.
func (r *CrmDealRoomService) Options(ctx context.Context, spaceID string, dealID string, query *CrmDealRoomOptionsQuery, opts ...CallOption) ([]CrmDealRoomOption, error) {
	var out []CrmDealRoomOption
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(dealID)+"/deal-room/options", query, nil, false, "options", &out, opts)
	return out, err
}

// Share — Share a document, file, signature request, booking page or match into the room.
//
// POST /api/crm/spaces/{spaceId}/records/{dealId}/deal-room/items — Token permission: write:crm.
func (r *CrmDealRoomService) Share(ctx context.Context, spaceID string, dealID string, body *CrmDealRoomShareBody, opts ...CallOption) (*CrmDealRoomItem, error) {
	var out *CrmDealRoomItem
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(dealID)+"/deal-room/items", nil, body, false, "item", &out, opts)
	return out, err
}

// Unshare — Take a shared item back out.
//
// DELETE /api/crm/spaces/{spaceId}/records/{dealId}/deal-room/items/{itemId} — Token permission: write:crm.
func (r *CrmDealRoomService) Unshare(ctx context.Context, spaceID string, dealID string, itemID string, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "DELETE", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(dealID)+"/deal-room/items/"+pathSeg(itemID), nil, nil, false, "", &out, opts)
	return out, err
}

// UpdateStep — Change, tick off or reorder a step.
//
// PATCH /api/crm/spaces/{spaceId}/records/{dealId}/deal-room/steps/{stepId} — Token permission: write:crm.
func (r *CrmDealRoomService) UpdateStep(ctx context.Context, spaceID string, dealID string, stepID string, body *CrmDealRoomStepPatch, opts ...CallOption) (*CrmDealRoomStep, error) {
	var out *CrmDealRoomStep
	err := r.c.call(ctx, "PATCH", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(dealID)+"/deal-room/steps/"+pathSeg(stepID), nil, body, false, "step", &out, opts)
	return out, err
}

// CrmDuplicatesService — crm.duplicates
type CrmDuplicatesService struct {
	c *Client
}

func newCrmDuplicatesService(c *Client) *CrmDuplicatesService {
	return &CrmDuplicatesService{c: c}
}

// Check — Records the given cells may be the same as — at most five.
//
// POST /api/crm/spaces/{spaceId}/duplicates/check — Token permission: write:crm.
//
// body: one of *CrmDuplicatesCheckBody, *CrmDuplicatesCheckBody2.
func (r *CrmDuplicatesService) Check(ctx context.Context, spaceID string, body any, opts ...CallOption) ([]CrmDuplicateCandidate, error) {
	var out []CrmDuplicateCandidate
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/duplicates/check", nil, body, false, "candidates", &out, opts)
	return out, err
}

// List — Groups of records that look like the same person or company.
//
// GET /api/crm/spaces/{spaceId}/duplicates — Token permission: read:crm.
func (r *CrmDuplicatesService) List(ctx context.Context, spaceID string, opts ...CallOption) (*CrmDuplicatesListResponse, error) {
	var out *CrmDuplicatesListResponse
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/duplicates", nil, nil, false, "", &out, opts)
	return out, err
}

// Merge — Merges records into one; the others are deleted.
//
// Needs manage level on the space.
//
// POST /api/crm/spaces/{spaceId}/merge — Token permission: write:crm.
func (r *CrmDuplicatesService) Merge(ctx context.Context, spaceID string, body *CrmMergeRequest, opts ...CallOption) (*CrmDuplicatesMergeResponse, error) {
	var out *CrmDuplicatesMergeResponse
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/merge", nil, body, false, "", &out, opts)
	return out, err
}

// PreviewMerge — What merging would do; nothing is written.
//
// POST /api/crm/spaces/{spaceId}/merge/preview — Token permission: write:crm.
func (r *CrmDuplicatesService) PreviewMerge(ctx context.Context, spaceID string, body *CrmMergeRequest, opts ...CallOption) (*CrmMergePreview, error) {
	var out *CrmMergePreview
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/merge/preview", nil, body, false, "preview", &out, opts)
	return out, err
}

// CrmEnrichmentService — crm.enrichment
type CrmEnrichmentService struct {
	c *Client
}

func newCrmEnrichmentService(c *Client) *CrmEnrichmentService {
	return &CrmEnrichmentService{c: c}
}

// Bulk — Enriches the selected records: a run each, for those with a website or company address.
//
// Needs edit level and company enrichment on the plan. The month's allowance is checked for all before any is queued (403 limit).
//
// POST /api/crm/spaces/{spaceId}/enrichment/bulk — Token permission: write:crm.
//
// spaceID: The space.
func (r *CrmEnrichmentService) Bulk(ctx context.Context, spaceID string, body *CrmEnrichmentBulkBody, opts ...CallOption) (*CrmEnrichmentBulkResponse, error) {
	var out *CrmEnrichmentBulkResponse
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/enrichment/bulk", nil, body, false, "", &out, opts)
	return out, err
}

// Configure — Replaces a space's company-enrichment settings.
//
// Needs manage level and `crm:manage`. Each mapped column must exist, fit its field and be mapped once (400 `crmEnrichmentInvalid` with `problems`). Automatic enrichment and refreshing need it on the plan.
//
// PUT /api/crm/spaces/{spaceId}/enrichment — Token permission: write:crm.
//
// spaceID: The space.
func (r *CrmEnrichmentService) Configure(ctx context.Context, spaceID string, body *CrmEnrichmentConfigureBody, opts ...CallOption) (*CrmEnrichmentConfigureResponse, error) {
	var out *CrmEnrichmentConfigureResponse
	err := r.c.call(ctx, "PUT", "/api/crm/spaces/"+pathSeg(spaceID)+"/enrichment", nil, body, false, "", &out, opts)
	return out, err
}

// Decide — Accepts or rejects enrichment proposals; accepted values are written to the record.
//
// Needs edit level. A cell changed since the proposal answers 409 `recordFieldChanged` unless it is in `force`. `acceptEmpty` accepts every open proposal whose cell is empty now.
//
// POST /api/crm/spaces/{spaceId}/records/{recordId}/enrichment/decide — Token permission: write:crm.
//
// spaceID: The space.
// recordID: The record.
func (r *CrmEnrichmentService) Decide(ctx context.Context, spaceID string, recordID string, body *CrmEnrichmentDecideBody, opts ...CallOption) (*CrmEnrichmentDecideResponse, error) {
	var out *CrmEnrichmentDecideResponse
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/enrichment/decide", nil, body, false, "", &out, opts)
	return out, err
}

// Get — A record's company enrichment: its last run, its open proposals with what each cell holds now, and the month's allowance.
//
// GET /api/crm/spaces/{spaceId}/records/{recordId}/enrichment — Token permission: read:crm.
//
// spaceID: The space.
// recordID: The record.
func (r *CrmEnrichmentService) Get(ctx context.Context, spaceID string, recordID string, opts ...CallOption) (*CrmEnrichmentGetResponse, error) {
	var out *CrmEnrichmentGetResponse
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/enrichment", nil, nil, false, "", &out, opts)
	return out, err
}

// Run — Queues an enrichment run for a record: its company's public details become proposals.
//
// Needs edit level. One run per record at a time (409 `crmEnrichmentRunning`); the month's allowance applies (403 limit).
//
// POST /api/crm/spaces/{spaceId}/records/{recordId}/enrichment — Token permission: write:crm.
//
// spaceID: The space.
// recordID: The record.
func (r *CrmEnrichmentService) Run(ctx context.Context, spaceID string, recordID string, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/enrichment", nil, nil, false, "run", &out, opts)
	return out, err
}

// Settings — A space's company-enrichment settings, the mapping its columns suggest, and the organisation's allowance this month.
//
// Needs company enrichment on the plan.
//
// GET /api/crm/spaces/{spaceId}/enrichment — Token permission: read:crm.
//
// spaceID: The space.
func (r *CrmEnrichmentService) Settings(ctx context.Context, spaceID string, opts ...CallOption) (*CrmEnrichmentSettingsResponse, error) {
	var out *CrmEnrichmentSettingsResponse
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/enrichment", nil, nil, false, "", &out, opts)
	return out, err
}

// CrmGroupsService — crm.groups
type CrmGroupsService struct {
	c *Client
}

func newCrmGroupsService(c *Client) *CrmGroupsService {
	return &CrmGroupsService{c: c}
}

// Delete — Deletes a group with every space in it, their records, activities and reminders.
//
// Needs `crm:manage`, `database:delete` and delete level on every space of the group: all or nothing, checked before the first goes.
//
// DELETE /api/crm/groups/{groupId} — Token permission: write:crm.
//
// groupID: The group.
func (r *CrmGroupsService) Delete(ctx context.Context, groupID string, opts ...CallOption) (*CrmGroupsDeleteResponse, error) {
	var out *CrmGroupsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/crm/groups/"+pathSeg(groupID), nil, nil, false, "", &out, opts)
	return out, err
}

// Rename — Renames a group of spaces, on every space in it.
//
// Needs `crm:manage` and manage level on every space of the group — renamed everywhere or nowhere (403 `crmGroupNotEditable`).
//
// PATCH /api/crm/groups/{groupId} — Token permission: write:crm.
//
// groupID: The group.
func (r *CrmGroupsService) Rename(ctx context.Context, groupID string, body *CrmGroupsRenameBody, opts ...CallOption) (*CrmGroupsRenameResponse, error) {
	var out *CrmGroupsRenameResponse
	err := r.c.call(ctx, "PATCH", "/api/crm/groups/"+pathSeg(groupID), nil, body, false, "", &out, opts)
	return out, err
}

// CrmJourneyService — crm.journey
type CrmJourneyService struct {
	c *Client
}

func newCrmJourneyService(c *Client) *CrmJourneyService {
	return &CrmJourneyService{c: c}
}

// Get — A record's visitor journeys: the consented visitors of the organisation's sites linked to it, their touches and events.
//
// Needs `journeys:read` and visitor journeys on the plan.
//
// GET /api/crm/spaces/{spaceId}/records/{recordId}/journey — Token permission: read:crm.
//
// spaceID: The space.
// recordID: The record.
func (r *CrmJourneyService) Get(ctx context.Context, spaceID string, recordID string, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/journey", nil, nil, false, "data", &out, opts)
	return out, err
}

// Summarise — Summarises with the platform's AI how this person found and used the organisation's sites; nothing is stored.
//
// Needs `journeys:read`; the AI allowance applies (429). The AI reads touches, pages and timestamps only — never an e-mail or a name.
//
// POST /api/crm/spaces/{spaceId}/records/{recordId}/journey/summary — Token permission: write:crm.
//
// spaceID: The space.
// recordID: The record.
func (r *CrmJourneyService) Summarise(ctx context.Context, spaceID string, recordID string, body *CrmJourneySummariseBody, opts ...CallOption) (*CrmJourneySummariseResponse, error) {
	var out *CrmJourneySummariseResponse
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/journey/summary", nil, body, false, "", &out, opts)
	return out, err
}

// CrmMatchesService — crm.matches
type CrmMatchesService struct {
	c *Client
}

func newCrmMatchesService(c *Client) *CrmMatchesService {
	return &CrmMatchesService{c: c}
}

// ForRecord — The records matching this one, ranked, with reasons.
//
// GET /api/crm/spaces/{spaceId}/records/{recordId}/matches — Token permission: read:crm.
func (r *CrmMatchesService) ForRecord(ctx context.Context, spaceID string, recordID string, opts ...CallOption) (*CrmMatchesForRecordResponse, error) {
	var out *CrmMatchesForRecordResponse
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/matches", nil, nil, false, "", &out, opts)
	return out, err
}

// ForSpace — The space's best open matches.
//
// GET /api/crm/spaces/{spaceId}/matches — Token permission: read:crm.
func (r *CrmMatchesService) ForSpace(ctx context.Context, spaceID string, opts ...CallOption) (*CrmMatchesForSpaceResponse, error) {
	var out *CrmMatchesForSpaceResponse
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/matches", nil, nil, false, "", &out, opts)
	return out, err
}

// Propose — Logs a match as proposed: a "Match proposed" activity on each record the caller may change, and a link in the proposed column.
//
// Needs edit level on at least one of the two spaces.
//
// POST /api/crm/spaces/{spaceId}/records/{recordId}/matches/propose — Token permission: write:crm.
//
// spaceID: The space.
// recordID: This record.
func (r *CrmMatchesService) Propose(ctx context.Context, spaceID string, recordID string, body *CrmMatchesProposeBody, opts ...CallOption) (*CrmMatchesProposeResponse, error) {
	var out *CrmMatchesProposeResponse
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/matches/propose", nil, body, false, "", &out, opts)
	return out, err
}

// CrmMatchingService — crm.matching
type CrmMatchingService struct {
	c *Client
}

func newCrmMatchingService(c *Client) *CrmMatchingService {
	return &CrmMatchingService{c: c}
}

// Get — A space's matching definitions (properties and search profiles), those of other spaces looking at it, and what no longer fits.
//
// GET /api/crm/spaces/{spaceId}/matching — Token permission: read:crm.
//
// spaceID: The space.
func (r *CrmMatchingService) Get(ctx context.Context, spaceID string, query *CrmMatchingGetQuery, opts ...CallOption) (*CrmMatchingGetResponse, error) {
	var out *CrmMatchingGetResponse
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/matching", query, nil, false, "", &out, opts)
	return out, err
}

// Put — Replaces a space's matching definitions.
//
// Needs manage level and `crm:manage`. Each definition names another space of the workspace the caller can open and fits both spaces' columns (400 `crmMatchingInvalid` with `problems`).
//
// PUT /api/crm/spaces/{spaceId}/matching — Token permission: write:crm.
//
// spaceID: The space.
func (r *CrmMatchingService) Put(ctx context.Context, spaceID string, body *CrmMatchingPutBody, opts ...CallOption) (*CrmMatchingPutResponse, error) {
	var out *CrmMatchingPutResponse
	err := r.c.call(ctx, "PUT", "/api/crm/spaces/"+pathSeg(spaceID)+"/matching", nil, body, false, "", &out, opts)
	return out, err
}

// CrmMeetingsService — crm.meetings
type CrmMeetingsService struct {
	c *Client
}

func newCrmMeetingsService(c *Client) *CrmMeetingsService {
	return &CrmMeetingsService{c: c}
}

// Apply — Writes a meeting's outcome to the CRM: a Meeting activity per person, a task per next step — each once per event.
//
// An item applied before is skipped and reported `alreadyApplied`.
//
// POST /api/crm/meetings/{eventId} — Token permission: write:crm.
//
// eventID: The calendar event.
func (r *CrmMeetingsService) Apply(ctx context.Context, eventID string, body *CrmMeetingsApplyBody, opts ...CallOption) (*CrmMeetingsApplyResponse, error) {
	var out *CrmMeetingsApplyResponse
	err := r.c.call(ctx, "POST", "/api/crm/meetings/"+pathSeg(eventID), nil, body, false, "", &out, opts)
	return out, err
}

// Proposal — What a calendar meeting's outcome would write to the CRM: the attendees' records, their open deals, the next steps as tasks.
//
// Writes nothing. Only records the caller may see.
//
// GET /api/crm/meetings/{eventId} — Token permission: read:crm.
//
// eventID: The calendar event.
func (r *CrmMeetingsService) Proposal(ctx context.Context, eventID string, query *CrmMeetingsProposalQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/crm/meetings/"+pathSeg(eventID), query, nil, false, "proposal", &out, opts)
	return out, err
}

// CrmOfflineConversionsService — crm.offlineConversions
type CrmOfflineConversionsService struct {
	c *Client
}

func newCrmOfflineConversionsService(c *Client) *CrmOfflineConversionsService {
	return &CrmOfflineConversionsService{c: c}
}

// Get — A deals space's offline-conversion settings — which stage reports which conversion to which ad network — and its last deliveries.
//
// GET /api/crm/spaces/{spaceId}/offline-conversions — Token permission: read:crm.
//
// spaceID: The deals space.
func (r *CrmOfflineConversionsService) Get(ctx context.Context, spaceID string, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/offline-conversions", nil, nil, false, "data", &out, opts)
	return out, err
}

// Put — Replaces a deals space's offline-conversion settings.
//
// Needs manage level, `crm:manage` and offline conversions on the plan. 400 `crmOfflineInvalid` with `problems`.
//
// PUT /api/crm/spaces/{spaceId}/offline-conversions — Token permission: write:crm.
//
// spaceID: The deals space.
func (r *CrmOfflineConversionsService) Put(ctx context.Context, spaceID string, body *CrmOfflineConversionsPutBody, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "PUT", "/api/crm/spaces/"+pathSeg(spaceID)+"/offline-conversions", nil, body, false, "data", &out, opts)
	return out, err
}

// CrmPipelineService — crm.pipeline
type CrmPipelineService struct {
	c *Client
}

func newCrmPipelineService(c *Client) *CrmPipelineService {
	return &CrmPipelineService{c: c}
}

// Add — Adds another pipeline beside a deals space: a new deals space in the same group, its stages copied from this one or the template's.
//
// Needs `crm:manage` and `database:create`; counts against the plan's CRM space limit.
//
// POST /api/crm/spaces/{spaceId}/pipeline — Token permission: write:crm.
//
// spaceID: The deals space.
func (r *CrmPipelineService) Add(ctx context.Context, spaceID string, body *CrmPipelineAddBody, opts ...CallOption) (*CrmPipelineAddResponse, error) {
	var out *CrmPipelineAddResponse
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/pipeline", nil, body, false, "", &out, opts)
	return out, err
}

// Update — Changes a deals space's pipeline as one list: add, rename, recolour, reorder and remove stages, mark them won or lost.
//
// Needs edit level on the table. Cells and views holding a renamed stage follow; a removed stage's deals move by `moves` (or lose their stage). 400 `crmPipelineInvalid` with `problems`.
//
// PATCH /api/crm/spaces/{spaceId}/pipeline — Token permission: write:crm.
//
// spaceID: The deals space.
func (r *CrmPipelineService) Update(ctx context.Context, spaceID string, body *CrmPipelineUpdateBody, opts ...CallOption) (*CrmPipelineUpdateResponse, error) {
	var out *CrmPipelineUpdateResponse
	err := r.c.call(ctx, "PATCH", "/api/crm/spaces/"+pathSeg(spaceID)+"/pipeline", nil, body, false, "", &out, opts)
	return out, err
}

// CrmPlaybooksService — crm.playbooks
type CrmPlaybooksService struct {
	c *Client
}

func newCrmPlaybooksService(c *Client) *CrmPlaybooksService {
	return &CrmPlaybooksService{c: c}
}

// Get — The space's playbooks and what the editor offers.
//
// GET /api/crm/spaces/{spaceId}/playbooks — Token permission: read:crm.
func (r *CrmPlaybooksService) Get(ctx context.Context, spaceID string, opts ...CallOption) (*CrmPlaybooksGetResponse, error) {
	var out *CrmPlaybooksGetResponse
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/playbooks", nil, nil, false, "", &out, opts)
	return out, err
}

// Put — Replaces the space's playbooks.
//
// 400 `crmPlaybooksInvalid` with `problems`.
//
// PUT /api/crm/spaces/{spaceId}/playbooks — Token permission: write:crm.
func (r *CrmPlaybooksService) Put(ctx context.Context, spaceID string, body *CrmPlaybooksPutBody, opts ...CallOption) (*CrmPlaybooksPutResponse, error) {
	var out *CrmPlaybooksPutResponse
	err := r.c.call(ctx, "PUT", "/api/crm/spaces/"+pathSeg(spaceID)+"/playbooks", nil, body, false, "", &out, opts)
	return out, err
}

// Retry — Run a run's skipped or failed steps again, as the caller.
//
// POST /api/crm/spaces/{spaceId}/records/{dealId}/playbooks/{runId}/retry — Token permission: write:crm.
func (r *CrmPlaybooksService) Retry(ctx context.Context, spaceID string, dealID string, runID string, opts ...CallOption) ([]CrmPlaybookRun, error) {
	var out []CrmPlaybookRun
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(dealID)+"/playbooks/"+pathSeg(runID)+"/retry", nil, nil, false, "runs", &out, opts)
	return out, err
}

// Runs — A deal's playbook runs.
//
// GET /api/crm/spaces/{spaceId}/records/{dealId}/playbooks — Token permission: read:crm.
func (r *CrmPlaybooksService) Runs(ctx context.Context, spaceID string, dealID string, opts ...CallOption) (*CrmPlaybooksRunsResponse, error) {
	var out *CrmPlaybooksRunsResponse
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(dealID)+"/playbooks", nil, nil, false, "", &out, opts)
	return out, err
}

// CrmPortalService — crm.portal
type CrmPortalService struct {
	c *Client
}

func newCrmPortalService(c *Client) *CrmPortalService {
	return &CrmPortalService{c: c}
}

// Invitations — The person's invitations, newest first.
//
// GET /api/crm/spaces/{spaceId}/records/{recordId}/portal-invitations — Token permission: read:crm.
func (r *CrmPortalService) Invitations(ctx context.Context, spaceID string, recordID string, opts ...CallOption) (*CrmPortalInvitationsResponse, error) {
	var out *CrmPortalInvitationsResponse
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/portal-invitations", nil, nil, false, "", &out, opts)
	return out, err
}

// Invite — Invite the person: a single-use sign-in link by mail.
//
// POST /api/crm/spaces/{spaceId}/records/{recordId}/portal-invite — Token permission: write:crm.
func (r *CrmPortalService) Invite(ctx context.Context, spaceID string, recordID string, body *CrmPortalInviteBody, opts ...CallOption) (*CrmPortalInviteResult, error) {
	var out *CrmPortalInviteResult
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/portal-invite", nil, body, false, "result", &out, opts)
	return out, err
}

// InviteOptions — What an invitation would offer — address, websites, the mail as sent.
//
// GET /api/crm/spaces/{spaceId}/records/{recordId}/portal-invite — Token permission: read:crm.
func (r *CrmPortalService) InviteOptions(ctx context.Context, spaceID string, recordID string, query *CrmPortalInviteOptionsQuery, opts ...CallOption) (*CrmPortalInviteOptions, error) {
	var out *CrmPortalInviteOptions
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/portal-invite", query, nil, false, "invite", &out, opts)
	return out, err
}

// Record — Whether a record is shown in the customer portal and with which fields, and how many website accounts see it.
//
// GET /api/crm/spaces/{spaceId}/records/{recordId}/portal — Token permission: read:crm.
//
// spaceID: The space.
// recordID: The record.
func (r *CrmPortalService) Record(ctx context.Context, spaceID string, recordID string, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/portal", nil, nil, false, "portal", &out, opts)
	return out, err
}

// Resend — Send a new invitation in place of this one.
//
// POST /api/crm/portal/invitations/{invitationId}/resend — Token permission: write:crm.
func (r *CrmPortalService) Resend(ctx context.Context, invitationID string, opts ...CallOption) (*CrmPortalInviteResult, error) {
	var out *CrmPortalInviteResult
	err := r.c.call(ctx, "POST", "/api/crm/portal/invitations/"+pathSeg(invitationID)+"/resend", nil, nil, false, "result", &out, opts)
	return out, err
}

// SetSitePage — Chooses a website's customer portal page — a live page holding a customer portal block — or clears the choice.
//
// Needs `crm:manage` and `project:update`. 400 `crmPortalPageNotPortal`.
//
// PUT /api/crm/portal/sites/{projectId} — Token permission: write:crm.
//
// projectID: The website.
func (r *CrmPortalService) SetSitePage(ctx context.Context, projectID string, body *CrmPortalSetSitePageBody, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "PUT", "/api/crm/portal/sites/"+pathSeg(projectID), nil, body, false, "site", &out, opts)
	return out, err
}

// SiteInvitations — A website's invitations.
//
// GET /api/crm/portal/sites/{projectId}/invitations — Token permission: read:crm.
func (r *CrmPortalService) SiteInvitations(ctx context.Context, projectID string, opts ...CallOption) (*CrmPortalSiteInvitationsResponse, error) {
	var out *CrmPortalSiteInvitationsResponse
	err := r.c.call(ctx, "GET", "/api/crm/portal/sites/"+pathSeg(projectID)+"/invitations", nil, nil, false, "", &out, opts)
	return out, err
}

// Sites — The workspace's websites with their customer portal pages, the chosen one, and what stands in the way.
//
// GET /api/crm/portal/sites — Token permission: read:crm.
func (r *CrmPortalService) Sites(ctx context.Context, query *CrmPortalSitesQuery, opts ...CallOption) (*CrmPortalSitesResponse, error) {
	var out *CrmPortalSitesResponse
	err := r.c.call(ctx, "GET", "/api/crm/portal/sites", query, nil, false, "", &out, opts)
	return out, err
}

// Spaces — The CRM spaces a customer portal block on a website may show, each with the fields it shares.
//
// GET /api/crm/portal/spaces — Token permission: read:crm.
func (r *CrmPortalService) Spaces(ctx context.Context, query *CrmPortalSpacesQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/crm/portal/spaces", query, nil, false, "spaces", &out, opts)
	return out, err
}

// Withdraw — Withdraw an open invitation.
//
// POST /api/crm/portal/invitations/{invitationId}/withdraw — Token permission: write:crm.
func (r *CrmPortalService) Withdraw(ctx context.Context, invitationID string, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "POST", "/api/crm/portal/invitations/"+pathSeg(invitationID)+"/withdraw", nil, nil, false, "", &out, opts)
	return out, err
}

// CrmRecordsService — crm.records
type CrmRecordsService struct {
	c *Client
}

func newCrmRecordsService(c *Client) *CrmRecordsService {
	return &CrmRecordsService{c: c}
}

// Create — Creates a record (cells by column id), or several at once with `records`.
//
// Counts against the plan's CRM record limit. Relations and person cells are checked; a database with approvers files the new record for approval.
//
// POST /api/databases/{spaceId}/records
//
// spaceID: The space.
// body: one of *CrmRecordsCreateBody, *CrmRecordsCreateBody2.
func (r *CrmRecordsService) Create(ctx context.Context, spaceID string, body any, opts ...CallOption) (*CrmRecordsCreateResponse, error) {
	var out *CrmRecordsCreateResponse
	err := r.c.call(ctx, "POST", "/api/databases/"+pathSeg(spaceID)+"/records", nil, body, false, "", &out, opts)
	return out, err
}

// Delete — Moves a record to the trash.
//
// 409 `recordPendingApproval` while a change waits for approval.
//
// DELETE /api/databases/{spaceId}/records/{recordId}
//
// spaceID: The space.
// recordID: The record.
func (r *CrmRecordsService) Delete(ctx context.Context, spaceID string, recordID string, opts ...CallOption) (*CrmRecordsDeleteResponse, error) {
	var out *CrmRecordsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/databases/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID), nil, nil, false, "", &out, opts)
	return out, err
}

// Get — One record.
//
// GET /api/databases/{spaceId}/records/{recordId}
//
// spaceID: The space.
// recordID: The record.
func (r *CrmRecordsService) Get(ctx context.Context, spaceID string, recordID string, opts ...CallOption) (*CrmRecord, error) {
	var out *CrmRecord
	err := r.c.call(ctx, "GET", "/api/databases/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID), nil, nil, false, "record", &out, opts)
	return out, err
}

// List — One page of a space's records, with filters, sort and offset paging.
//
// Only the records the caller may see (row security).
//
// GET /api/databases/{spaceId}/records
//
// spaceID: The space.
func (r *CrmRecordsService) List(ctx context.Context, spaceID string, query *CrmRecordsListQuery, opts ...CallOption) (*CrmRecordsListResponse, error) {
	var out *CrmRecordsListResponse
	err := r.c.call(ctx, "GET", "/api/databases/"+pathSeg(spaceID)+"/records", query, nil, false, "", &out, opts)
	return out, err
}

// Update — Changes the cells of a record.
//
// With `merge: true` only the cells in `data` change; without it `data` is the whole row and a cell left out is removed. With `base` (each changed cell as the edit began), a cell somebody changed since is refused (409 `recordFieldChanged`) instead of overwritten. 409 `recordPendingApproval` while an earlier change waits for approval.
//
// PATCH /api/databases/{spaceId}/records/{recordId}
//
// spaceID: The space.
// recordID: The record.
func (r *CrmRecordsService) Update(ctx context.Context, spaceID string, recordID string, body *CrmRecordsUpdateBody, opts ...CallOption) (*CrmRecord, error) {
	var out *CrmRecord
	err := r.c.call(ctx, "PATCH", "/api/databases/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID), nil, body, false, "record", &out, opts)
	return out, err
}

// CrmSalesService — crm.sales
type CrmSalesService struct {
	c *Client
}

func newCrmSalesService(c *Client) *CrmSalesService {
	return &CrmSalesService{c: c}
}

// DealSignals — Stalled deals and their next steps.
//
// GET /api/crm/spaces/{spaceId}/deal-signals — Token permission: read:crm.
func (r *CrmSalesService) DealSignals(ctx context.Context, spaceID string, opts ...CallOption) (*CrmSalesDealSignalsResponse, error) {
	var out *CrmSalesDealSignalsResponse
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/deal-signals", nil, nil, false, "", &out, opts)
	return out, err
}

// FlowFunnel — The question flow funnel of an activity type: how the answers of logged calls and meetings went.
//
// For one space or a workspace's spaces; one of `spaceId` and `workspaceId` is required.
//
// GET /api/crm/sales/flows — Token permission: read:crm.
func (r *CrmSalesService) FlowFunnel(ctx context.Context, query *CrmSalesFlowFunnelQuery, opts ...CallOption) (*CrmFlowFunnelReport, error) {
	var out *CrmFlowFunnelReport
	err := r.c.call(ctx, "GET", "/api/crm/sales/flows", query, nil, false, "report", &out, opts)
	return out, err
}

// Report — The sales report: pipeline, win rate, sales cycle, forecast, activities, lost reasons, stalled deals.
//
// For one deals space (`spaceId`) or every deals space of a workspace the caller may open (`workspaceId`); one of the two is required. Only the deals the caller may see; periods on the caller's clock. 60 a minute.
//
// GET /api/crm/sales — Token permission: read:crm.
func (r *CrmSalesService) Report(ctx context.Context, query *CrmSalesReportQuery, opts ...CallOption) (*CrmSalesReport, error) {
	var out *CrmSalesReport
	err := r.c.call(ctx, "GET", "/api/crm/sales", query, nil, false, "report", &out, opts)
	return out, err
}

// CrmSpacesService — crm.spaces
type CrmSpacesService struct {
	c *Client
}

func newCrmSpacesService(c *Client) *CrmSpacesService {
	return &CrmSpacesService{c: c}
}

// Convert — Turns an existing table into a CRM space; its rows, columns, views, shares and relations stay as they are.
//
// Needs manage level on the table and `crm:manage`. The table and its rows move to the CRM's own plan limits (403 limit). 400 `crmConfigInvalid` with `problems`; 409 `crmAlreadySpace`.
//
// POST /api/crm/spaces/convert — Token permission: write:crm.
func (r *CrmSpacesService) Convert(ctx context.Context, body *CrmSpacesConvertBody, opts ...CallOption) (*CrmSpaceSummary, error) {
	var out *CrmSpaceSummary
	err := r.c.call(ctx, "POST", "/api/crm/spaces/convert", nil, body, false, "space", &out, opts)
	return out, err
}

// Convertible — The plain tables of a workspace the caller could turn into a CRM space.
//
// Needs `crm:manage` in the workspace. Names, icons, row counts and columns only.
//
// GET /api/crm/spaces/convert — Token permission: read:crm.
func (r *CrmSpacesService) Convertible(ctx context.Context, query *CrmSpacesConvertibleQuery, opts ...CallOption) (*CrmSpacesConvertibleResponse, error) {
	var out *CrmSpacesConvertibleResponse
	err := r.c.call(ctx, "GET", "/api/crm/spaces/convert", query, nil, false, "", &out, opts)
	return out, err
}

// Create — Creates a custom space with one Name column and the default activity types.
//
// Needs `crm:manage` and `database:create` in the workspace; counts against the plan's CRM space limit.
//
// POST /api/crm/spaces — Token permission: write:crm.
func (r *CrmSpacesService) Create(ctx context.Context, body *CrmSpacesCreateBody, opts ...CallOption) (*CrmSpaceSummary, error) {
	var out *CrmSpaceSummary
	err := r.c.call(ctx, "POST", "/api/crm/spaces", nil, body, false, "space", &out, opts)
	return out, err
}

// Design — Designs CRM spaces from a description with the platform's AI — answered as a proposal, never created here.
//
// Needs `crm:manage`, `database:create` and CRM AI on the plan; the AI allowance applies (429). Create the (edited) proposal with `crm.spaces.setup` and `design`.
//
// POST /api/crm/design — Token permission: write:crm.
func (r *CrmSpacesService) Design(ctx context.Context, body *CrmSpacesDesignBody, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/crm/design", nil, body, false, "proposal", &out, opts)
	return out, err
}

// Get — One space, its config and what the caller may do with it.
//
// The table itself — columns, views, records — is the Databases API.
//
// GET /api/crm/spaces/{spaceId} — Token permission: read:crm.
//
// spaceID: The space (a database of kind CRM space).
func (r *CrmSpacesService) Get(ctx context.Context, spaceID string, opts ...CallOption) (*CrmSpaceSummary, error) {
	var out *CrmSpaceSummary
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID), nil, nil, false, "space", &out, opts)
	return out, err
}

// List — The workspace's CRM spaces the caller can open, each with its config, record count and what no longer fits.
//
// Needs `crm:read` in the workspace.
//
// GET /api/crm/spaces — Token permission: read:crm.
func (r *CrmSpacesService) List(ctx context.Context, query *CrmSpacesListQuery, opts ...CallOption) ([]CrmSpaceSummary, error) {
	var out []CrmSpaceSummary
	err := r.c.call(ctx, "GET", "/api/crm/spaces", query, nil, false, "spaces", &out, opts)
	return out, err
}

// SetGroup — Moves a space into a group, out of any group, or into a new group of its own.
//
// Needs manage level on the table and `crm:manage`.
//
// PUT /api/crm/spaces/{spaceId}/group — Token permission: write:crm.
//
// spaceID: The space.
// body: one of *CrmSpacesSetGroupBody, *CrmSpacesSetGroupBody2.
func (r *CrmSpacesService) SetGroup(ctx context.Context, spaceID string, body any, opts ...CallOption) (*CrmSpacesSetGroupResponse, error) {
	var out *CrmSpacesSetGroupResponse
	err := r.c.call(ctx, "PUT", "/api/crm/spaces/"+pathSeg(spaceID)+"/group", nil, body, false, "", &out, opts)
	return out, err
}

// Setup — Creates spaces from a template of the gallery ("sales", "realEstate", …), or the classic People / Companies / Deals / Tasks set.
//
// A template's spaces form one new group with their relations both ways; a template that wants Tasks reuses the workspace's Tasks space by default. The classic set refuses a type that exists already (409 `crmSpaceTypeExists`) unless `force`.
//
// POST /api/crm/setup — Token permission: write:crm.
//
// body: CrmSetupInput.
func (r *CrmSpacesService) Setup(ctx context.Context, body any, opts ...CallOption) (*CrmSpacesSetupResponse, error) {
	var out *CrmSpacesSetupResponse
	err := r.c.call(ctx, "POST", "/api/crm/setup", nil, body, false, "", &out, opts)
	return out, err
}

// Summary — What a space's overview card shows: the pipeline per stage with count and amount, open / won / lost, open and done tasks, the records changed last.
//
// Counted only from the rows the caller may see, up to a read cap (`truncated`).
//
// GET /api/crm/spaces/{spaceId}/summary — Token permission: read:crm.
//
// spaceID: The space.
func (r *CrmSpacesService) Summary(ctx context.Context, spaceID string, opts ...CallOption) (*CrmSpacesSummaryResponse, error) {
	var out *CrmSpacesSummaryResponse
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/summary", nil, nil, false, "", &out, opts)
	return out, err
}

// UpdateConfig — Changes a space's CRM config: roles, pipeline settings, activity types, connected hubs, portal, website.
//
// Needs manage level on the table and `crm:manage`. A key left out stays; null clears a role. Checked against the columns as they are now: 400 `crmConfigInvalid` with `problems`. Switching on the portal, deal room or website needs it on the plan. Built-in activity types can be renamed or hidden, never removed.
//
// PATCH /api/crm/spaces/{spaceId} — Token permission: write:crm.
//
// spaceID: The space.
func (r *CrmSpacesService) UpdateConfig(ctx context.Context, spaceID string, body *CrmSpacesUpdateConfigBody, opts ...CallOption) (*CrmSpaceSummary, error) {
	var out *CrmSpaceSummary
	err := r.c.call(ctx, "PATCH", "/api/crm/spaces/"+pathSeg(spaceID), nil, body, false, "space", &out, opts)
	return out, err
}

// CrmTasksService — crm.tasks
type CrmTasksService struct {
	c *Client
}

func newCrmTasksService(c *Client) *CrmTasksService {
	return &CrmTasksService{c: c}
}

// SetDone — Ticks a CRM task done (the space's first done value), or opens it again.
//
// Needs edit level on the space.
//
// POST /api/crm/spaces/{spaceId}/records/{recordId}/done — Token permission: write:crm.
func (r *CrmTasksService) SetDone(ctx context.Context, spaceID string, recordID string, body *CrmTasksSetDoneBody, opts ...CallOption) (*CrmTasksSetDoneResponse, error) {
	var out *CrmTasksSetDoneResponse
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/done", nil, body, false, "", &out, opts)
	return out, err
}

// CrmWebsiteService — crm.website
type CrmWebsiteService struct {
	c *Client
}

func newCrmWebsiteService(c *Client) *CrmWebsiteService {
	return &CrmWebsiteService{c: c}
}

// CreatePages — Creates listing pages for a space on a website: a list page and a detail page per language, as drafts.
//
// Needs manage level, `crm:manage` and `page:create`. Nothing goes live: each page is published from the builder. 409 `crmWebsiteOff`, `crmWebsitePagesExist`, `pathTaken`.
//
// POST /api/crm/spaces/{spaceId}/website/pages — Token permission: write:crm.
//
// spaceID: The space.
func (r *CrmWebsiteService) CreatePages(ctx context.Context, spaceID string, body *CrmWebsiteCreatePagesBody, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/website/pages", nil, body, false, "pages", &out, opts)
	return out, err
}

// Get — A space on websites: whether it is on, how many records are public, the public columns and the listing pages.
//
// GET /api/crm/spaces/{spaceId}/website — Token permission: read:crm.
//
// spaceID: The space.
func (r *CrmWebsiteService) Get(ctx context.Context, spaceID string, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/website", nil, nil, false, "website", &out, opts)
	return out, err
}

// Preview — What visitors would see under a website configuration not saved yet: how many records would be public, and three of them.
//
// Needs manage level and `crm:manage`. Writes nothing.
//
// POST /api/crm/spaces/{spaceId}/website/preview — Token permission: write:crm.
//
// spaceID: The space.
func (r *CrmWebsiteService) Preview(ctx context.Context, spaceID string, body *CrmWebsitePreviewBody, opts ...CallOption) (*CrmWebsitePreviewResponse, error) {
	var out *CrmWebsitePreviewResponse
	err := r.c.call(ctx, "POST", "/api/crm/spaces/"+pathSeg(spaceID)+"/website/preview", nil, body, false, "", &out, opts)
	return out, err
}

// Record — Whether a record is public on the website now, with its live addresses and public columns.
//
// GET /api/crm/spaces/{spaceId}/records/{recordId}/website — Token permission: read:crm.
//
// spaceID: The space.
// recordID: The record.
func (r *CrmWebsiteService) Record(ctx context.Context, spaceID string, recordID string, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/crm/spaces/"+pathSeg(spaceID)+"/records/"+pathSeg(recordID)+"/website", nil, nil, false, "website", &out, opts)
	return out, err
}

// CurrenciesService — Currencies
type CurrenciesService struct {
	c *Client
}

func newCurrenciesService(c *Client) *CurrenciesService {
	return &CurrenciesService{c: c}
}

// List — Lists the currencies the platform offers, with how each one is written.
//
// Answers without a token. Only enabled currencies, in the administrator's order; cached for five minutes.
//
// GET /api/currencies — No token needed.
func (r *CurrenciesService) List(ctx context.Context, opts ...CallOption) ([]Currency, error) {
	var out []Currency
	err := r.c.call(ctx, "GET", "/api/currencies", nil, nil, false, "", &out, opts)
	return out, err
}

// Rates — Returns the latest stored exchange rates against one base currency.
//
// The base is the currency most rates are stored against. Before the first fetch the answer is `{ base: null, rates: {}, fetchedAt: null }`.
//
// GET /api/currency-rates
func (r *CurrenciesService) Rates(ctx context.Context, opts ...CallOption) (*CurrencyRates, error) {
	var out *CurrencyRates
	err := r.c.call(ctx, "GET", "/api/currency-rates", nil, nil, false, "", &out, opts)
	return out, err
}

// DashboardsService — Dashboards: An organisation's custom dashboards and their widgets. A dashboard is private to its creator, or visible to a workspace or the whole organisation; only its creator changes or deletes it.
type DashboardsService struct {
	c *Client
}

func newDashboardsService(c *Client) *DashboardsService {
	return &DashboardsService{c: c}
}

// AddWidget — Adds a widget to a dashboard.
//
// POST /api/dashboards/{dashboardId}/widgets
//
// dashboardID: The dashboard's id.
func (r *DashboardsService) AddWidget(ctx context.Context, dashboardID string, body *DashboardsAddWidgetBody, opts ...CallOption) (*DashboardWidget, error) {
	var out *DashboardWidget
	err := r.c.call(ctx, "POST", "/api/dashboards/"+pathSeg(dashboardID)+"/widgets", nil, body, false, "widget", &out, opts)
	return out, err
}

// Create — Creates a dashboard.
//
// Needs `dashboard:create` and membership of the organisation (and access to the workspace, if given).
//
// POST /api/dashboards
func (r *DashboardsService) Create(ctx context.Context, body *DashboardsCreateBody, opts ...CallOption) (*Dashboard, error) {
	var out *Dashboard
	err := r.c.call(ctx, "POST", "/api/dashboards", nil, body, false, "dashboard", &out, opts)
	return out, err
}

// Delete — Deletes a dashboard with its widgets.
//
// Needs `dashboard:delete`; only the dashboard's creator may.
//
// DELETE /api/dashboards/{id}
//
// id: The dashboard's id.
func (r *DashboardsService) Delete(ctx context.Context, id string, opts ...CallOption) (*DashboardsDeleteResponse, error) {
	var out *DashboardsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/dashboards/"+pathSeg(id), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteWidget — Removes a widget from a dashboard.
//
// DELETE /api/dashboards/{dashboardId}/widgets/{widgetId}
//
// dashboardID: The dashboard's id.
// widgetID: The widget's id.
func (r *DashboardsService) DeleteWidget(ctx context.Context, dashboardID string, widgetID string, opts ...CallOption) (*DashboardsDeleteWidgetResponse, error) {
	var out *DashboardsDeleteWidgetResponse
	err := r.c.call(ctx, "DELETE", "/api/dashboards/"+pathSeg(dashboardID)+"/widgets/"+pathSeg(widgetID), nil, nil, false, "", &out, opts)
	return out, err
}

// Get — Returns one dashboard with its widgets.
//
// Needs `dashboard:read`.
//
// GET /api/dashboards/{id}
//
// id: The dashboard's id.
func (r *DashboardsService) Get(ctx context.Context, id string, opts ...CallOption) (*DashboardsGetResponseDashboard, error) {
	var out *DashboardsGetResponseDashboard
	err := r.c.call(ctx, "GET", "/api/dashboards/"+pathSeg(id), nil, nil, false, "dashboard", &out, opts)
	return out, err
}

// List — Lists the dashboards of an organisation the caller can see.
//
// Needs `dashboard:read`. Without `organizationId` the list is empty. Workspace dashboards are listed only with their `workspaceId`.
//
// GET /api/dashboards
func (r *DashboardsService) List(ctx context.Context, query *DashboardsListQuery, opts ...CallOption) ([]Dashboard, error) {
	var out []Dashboard
	err := r.c.call(ctx, "GET", "/api/dashboards", query, nil, false, "dashboards", &out, opts)
	return out, err
}

// ListWidgets — Lists a dashboard's widgets, top to bottom.
//
// GET /api/dashboards/{dashboardId}/widgets
//
// dashboardID: The dashboard's id.
func (r *DashboardsService) ListWidgets(ctx context.Context, dashboardID string, opts ...CallOption) ([]DashboardWidget, error) {
	var out []DashboardWidget
	err := r.c.call(ctx, "GET", "/api/dashboards/"+pathSeg(dashboardID)+"/widgets", nil, nil, false, "widgets", &out, opts)
	return out, err
}

// Update — Changes a dashboard's name, description, visibility, order or default flag.
//
// Needs `dashboard:update`; only the dashboard's creator may.
//
// PATCH /api/dashboards/{id}
//
// id: The dashboard's id.
func (r *DashboardsService) Update(ctx context.Context, id string, body *DashboardsUpdateBody, opts ...CallOption) (*Dashboard, error) {
	var out *Dashboard
	err := r.c.call(ctx, "PATCH", "/api/dashboards/"+pathSeg(id), nil, body, false, "dashboard", &out, opts)
	return out, err
}

// UpdateWidget — Changes a widget's title, type, configuration or place on the grid.
//
// PATCH /api/dashboards/{dashboardId}/widgets/{widgetId}
//
// dashboardID: The dashboard's id.
// widgetID: The widget's id.
func (r *DashboardsService) UpdateWidget(ctx context.Context, dashboardID string, widgetID string, body *DashboardsUpdateWidgetBody, opts ...CallOption) (*DashboardWidget, error) {
	var out *DashboardWidget
	err := r.c.call(ctx, "PATCH", "/api/dashboards/"+pathSeg(dashboardID)+"/widgets/"+pathSeg(widgetID), nil, body, false, "widget", &out, opts)
	return out, err
}

// DataService — Data export: Everything of the caller's organizations they may see, as one JSON, Excel or CSV file.
type DataService struct {
	c *Client

	Exports *DataExportsService
}

func newDataService(c *Client) *DataService {
	return &DataService{c: c, Exports: newDataExportsService(c)}
}

// Export — Exports the caller's data at once, as a file.
//
// Needs `data:export`. Answers the file itself (JSON, .xlsx or CSV), not JSON in an envelope. For a large export, start a background one with `data.exports.start`.
//
// GET /api/data/export
func (r *DataService) Export(ctx context.Context, query *DataExportQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/data/export", query, nil, false, "", &out, opts)
	return out, err
}

// DataExportsService — data.exports
type DataExportsService struct {
	c *Client
}

func newDataExportsService(c *Client) *DataExportsService {
	return &DataExportsService{c: c}
}

// Download — Downloads the file of a finished background export.
//
// Needs `data:export`. Answers the file itself, not JSON. Only the person who started it; 404 until it is completed, 410 once it expired.
//
// GET /api/data/export/download
func (r *DataExportsService) Download(ctx context.Context, query *DataExportsDownloadQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/data/export/download", query, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the caller's last ten background exports, or returns one of them.
//
// Needs `data:export`. With `id` the answer is that export alone (an object, not a list); another person's export answers 404.
//
// GET /api/data/export/jobs
//
// The answer is one of *DataExportsListResponse, *DataExportJob, decoded as generic JSON (map[string]any, []any, …).
func (r *DataExportsService) List(ctx context.Context, query *DataExportsListQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/data/export/jobs", query, nil, false, "", &out, opts)
	return out, err
}

// Start — Starts a background export of the caller's data.
//
// Needs `data:export`. One at a time per person: while one runs the answer is 409 with that one's `jobId`. Poll `data.exports.list` with the id, then download it with `data.exports.download`.
//
// POST /api/data/export/jobs
func (r *DataExportsService) Start(ctx context.Context, body *DataExportsStartBody, opts ...CallOption) (*DataExportsStartResponse, error) {
	var out *DataExportsStartResponse
	err := r.c.call(ctx, "POST", "/api/data/export/jobs", nil, body, false, "", &out, opts)
	return out, err
}

// DatabasesService — Databases
type DatabasesService struct {
	c *Client

	Columns *DatabasesColumnsService
	Records *DatabasesRecordsService
	Shares  *DatabasesSharesService
}

func newDatabasesService(c *Client) *DatabasesService {
	return &DatabasesService{c: c, Columns: newDatabasesColumnsService(c), Records: newDatabasesRecordsService(c), Shares: newDatabasesSharesService(c)}
}

// Activity — Lists what happened in a table — records and columns added, changed, deleted — newest first.
//
// GET /api/databases/{databaseId}/activity
//
// databaseID: The table's id.
func (r *DatabasesService) Activity(ctx context.Context, databaseID string, query *DatabasesActivityQuery, opts ...CallOption) ([]DatabasesActivityResponseEntriesItem, error) {
	var out []DatabasesActivityResponseEntriesItem
	err := r.c.call(ctx, "GET", "/api/databases/"+pathSeg(databaseID)+"/activity", query, nil, false, "entries", &out, opts)
	return out, err
}

// Create — Creates a table, optionally with its columns.
//
// Needs `database.create`.
//
// POST /api/databases
func (r *DatabasesService) Create(ctx context.Context, body *DatabasesCreateBody, opts ...CallOption) (*Database, error) {
	var out *Database
	err := r.c.call(ctx, "POST", "/api/databases", nil, body, false, "database", &out, opts)
	return out, err
}

// Delete — Moves a table and its records to the trash.
//
// DELETE /api/databases/{databaseId}
//
// databaseID: The table's id.
func (r *DatabasesService) Delete(ctx context.Context, databaseID string, opts ...CallOption) (*DatabasesDeleteResponse, error) {
	var out *DatabasesDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/databases/"+pathSeg(databaseID), nil, nil, false, "", &out, opts)
	return out, err
}

// Get — Returns a table with its columns and views, and what the caller may do with it.
//
// GET /api/databases/{databaseId}
//
// databaseID: The table's id.
func (r *DatabasesService) Get(ctx context.Context, databaseID string, query *DatabasesGetQuery, opts ...CallOption) (*Database, error) {
	var out *Database
	err := r.c.call(ctx, "GET", "/api/databases/"+pathSeg(databaseID), query, nil, false, "database", &out, opts)
	return out, err
}

// Import — Imports a CSV file into a table, or with `dryRun` says what an import would do.
//
// Needs edit access. `options` maps each file column to an existing column, a new one, or nothing, and says how numbers, dates, invalid values and duplicates are handled. The plan's record limit applies to the whole file. An .xlsx is refused — save it as CSV.
//
// POST /api/databases/{databaseId}/import
//
// databaseID: The table's id.
func (r *DatabasesService) Import(ctx context.Context, databaseID string, body *DatabasesImportBody, opts ...CallOption) (*DatabasesImportResponse, error) {
	var out *DatabasesImportResponse
	err := r.c.call(ctx, "POST", "/api/databases/"+pathSeg(databaseID)+"/import", nil, body, true, "", &out, opts)
	return out, err
}

// List — Lists the tables of an organization or workspace the caller can open.
//
// GET /api/databases
func (r *DatabasesService) List(ctx context.Context, query *DatabasesListQuery, opts ...CallOption) ([]Database, error) {
	var out []Database
	err := r.c.call(ctx, "GET", "/api/databases", query, nil, false, "databases", &out, opts)
	return out, err
}

// People — Lists the people a table's person cells can name, and those they already name.
//
// GET /api/databases/{databaseId}/people
//
// databaseID: The table's id.
func (r *DatabasesService) People(ctx context.Context, databaseID string, opts ...CallOption) ([]DatabasesPeopleResponsePeopleItem, error) {
	var out []DatabasesPeopleResponsePeopleItem
	err := r.c.call(ctx, "GET", "/api/databases/"+pathSeg(databaseID)+"/people", nil, nil, false, "people", &out, opts)
	return out, err
}

// Update — Changes a table's name, look, visibility, columns, views or settings.
//
// `columns`, when sent, is the whole column list. A relation column whose pair in the linked table changes too is reported in `inverse`.
//
// PATCH /api/databases/{databaseId}
//
// databaseID: The table's id.
func (r *DatabasesService) Update(ctx context.Context, databaseID string, body *DatabasesUpdateBody, opts ...CallOption) (*Database, error) {
	var out *Database
	err := r.c.call(ctx, "PATCH", "/api/databases/"+pathSeg(databaseID), nil, body, false, "database", &out, opts)
	return out, err
}

// DatabasesColumnsService — databases.columns
type DatabasesColumnsService struct {
	c *Client
}

func newDatabasesColumnsService(c *Client) *DatabasesColumnsService {
	return &DatabasesColumnsService{c: c}
}

// Delete — Removes a column and its values from a table.
//
// DELETE /api/databases/{databaseId}/columns/{columnId}
//
// databaseID: The table's id.
// columnID: The column's id.
func (r *DatabasesColumnsService) Delete(ctx context.Context, databaseID string, columnID string, opts ...CallOption) (*DatabasesColumnsDeleteResponse, error) {
	var out *DatabasesColumnsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/databases/"+pathSeg(databaseID)+"/columns/"+pathSeg(columnID), nil, nil, false, "", &out, opts)
	return out, err
}

// DisableCalendarFeed — Stops a column's calendar feed; its URL stops working.
//
// DELETE /api/databases/{databaseId}/columns/{columnId}/calendar-feed
//
// databaseID: The table's id.
// columnID: The column's id.
func (r *DatabasesColumnsService) DisableCalendarFeed(ctx context.Context, databaseID string, columnID string, opts ...CallOption) (*DatabasesColumnsDisableCalendarFeedResponse, error) {
	var out *DatabasesColumnsDisableCalendarFeedResponse
	err := r.c.call(ctx, "DELETE", "/api/databases/"+pathSeg(databaseID)+"/columns/"+pathSeg(columnID)+"/calendar-feed", nil, nil, false, "", &out, opts)
	return out, err
}

// EnableCalendarFeed — Publishes a date column as a calendar feed and returns its subscription URL.
//
// POST /api/databases/{databaseId}/columns/{columnId}/calendar-feed
//
// databaseID: The table's id.
// columnID: The column's id.
func (r *DatabasesColumnsService) EnableCalendarFeed(ctx context.Context, databaseID string, columnID string, body *DatabasesColumnsEnableCalendarFeedBody, opts ...CallOption) (*DatabasesColumnsEnableCalendarFeedResponse, error) {
	var out *DatabasesColumnsEnableCalendarFeedResponse
	err := r.c.call(ctx, "POST", "/api/databases/"+pathSeg(databaseID)+"/columns/"+pathSeg(columnID)+"/calendar-feed", nil, body, false, "", &out, opts)
	return out, err
}

// Get — Returns one column of a table.
//
// GET /api/databases/{databaseId}/columns/{columnId}
//
// databaseID: The table's id.
// columnID: The column's id.
func (r *DatabasesColumnsService) Get(ctx context.Context, databaseID string, columnID string, opts ...CallOption) (*DatabaseColumn, error) {
	var out *DatabaseColumn
	err := r.c.call(ctx, "GET", "/api/databases/"+pathSeg(databaseID)+"/columns/"+pathSeg(columnID), nil, nil, false, "column", &out, opts)
	return out, err
}

// Update — Changes one column of a table — its name, type or options.
//
// `clearData: true` empties the column in every record — for a type change whose old values mean nothing in the new type. Pointing a relation at another table always empties it.
//
// PUT /api/databases/{databaseId}/columns/{columnId}
//
// databaseID: The table's id.
// columnID: The column's id.
func (r *DatabasesColumnsService) Update(ctx context.Context, databaseID string, columnID string, body *DatabasesColumnsUpdateBody, opts ...CallOption) (*DatabaseColumn, error) {
	var out *DatabaseColumn
	err := r.c.call(ctx, "PUT", "/api/databases/"+pathSeg(databaseID)+"/columns/"+pathSeg(columnID), nil, body, false, "column", &out, opts)
	return out, err
}

// DatabasesRecordsService — databases.records
type DatabasesRecordsService struct {
	c *Client
}

func newDatabasesRecordsService(c *Client) *DatabasesRecordsService {
	return &DatabasesRecordsService{c: c}
}

// Activity — Lists a record's history and logged activity, newest first.
//
// GET /api/databases/{databaseId}/records/{recordId}/activity
//
// databaseID: The table's id.
// recordID: The record's id.
func (r *DatabasesRecordsService) Activity(ctx context.Context, databaseID string, recordID string, query *DatabasesRecordsActivityQuery, opts ...CallOption) ([]DatabasesRecordsActivityResponseEntriesItem, error) {
	var out []DatabasesRecordsActivityResponseEntriesItem
	err := r.c.call(ctx, "GET", "/api/databases/"+pathSeg(databaseID)+"/records/"+pathSeg(recordID)+"/activity", query, nil, false, "entries", &out, opts)
	return out, err
}

// AIFormula — Computes an AI formula cell — `AI()`, `AI_CLASSIFY()`, `AI_SUMMARIZE()`, `AI_EXTRACT()` — for one record.
//
// Counts against the organization's AI budget.
//
// POST /api/databases/{databaseId}/records/{recordId}/ai-formula
//
// databaseID: The table's id.
// recordID: The record's id.
func (r *DatabasesRecordsService) AIFormula(ctx context.Context, databaseID string, recordID string, body *DatabasesRecordsAiFormulaBody, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "POST", "/api/databases/"+pathSeg(databaseID)+"/records/"+pathSeg(recordID)+"/ai-formula", nil, body, false, "value", &out, opts)
	return out, err
}

// Bulk — Deletes or duplicates up to 500 records at once.
//
// POST /api/databases/{databaseId}/records/bulk
//
// databaseID: The table's id.
func (r *DatabasesRecordsService) Bulk(ctx context.Context, databaseID string, body *DatabasesRecordsBulkBody, opts ...CallOption) (*DatabasesRecordsBulkResponse, error) {
	var out *DatabasesRecordsBulkResponse
	err := r.c.call(ctx, "POST", "/api/databases/"+pathSeg(databaseID)+"/records/bulk", nil, body, false, "", &out, opts)
	return out, err
}

// Create — Adds one record, or several with `records`.
//
// Counts against the plan's record limit. Linked records must exist and be readable by the caller.
//
// POST /api/databases/{databaseId}/records
//
// databaseID: The table's id.
func (r *DatabasesRecordsService) Create(ctx context.Context, databaseID string, body *CreateRecordInput, opts ...CallOption) (*DatabasesRecordsCreateResponse, error) {
	var out *DatabasesRecordsCreateResponse
	err := r.c.call(ctx, "POST", "/api/databases/"+pathSeg(databaseID)+"/records", nil, body, false, "", &out, opts)
	return out, err
}

// Delete — Moves a record to the trash.
//
// DELETE /api/databases/{databaseId}/records/{recordId}
//
// databaseID: The table's id.
// recordID: The record's id.
func (r *DatabasesRecordsService) Delete(ctx context.Context, databaseID string, recordID string, opts ...CallOption) (*DatabasesRecordsDeleteResponse, error) {
	var out *DatabasesRecordsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/databases/"+pathSeg(databaseID)+"/records/"+pathSeg(recordID), nil, nil, false, "", &out, opts)
	return out, err
}

// Get — Returns one record.
//
// GET /api/databases/{databaseId}/records/{recordId}
//
// databaseID: The table's id.
// recordID: The record's id.
func (r *DatabasesRecordsService) Get(ctx context.Context, databaseID string, recordID string, opts ...CallOption) (*DatabaseRecord, error) {
	var out *DatabaseRecord
	err := r.c.call(ctx, "GET", "/api/databases/"+pathSeg(databaseID)+"/records/"+pathSeg(recordID), nil, nil, false, "record", &out, opts)
	return out, err
}

// Links — Lists the records of other tables that link to a record and the caller may see.
//
// GET /api/databases/{databaseId}/records/{recordId}/links
//
// databaseID: The table's id.
// recordID: The record's id.
func (r *DatabasesRecordsService) Links(ctx context.Context, databaseID string, recordID string, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/databases/"+pathSeg(databaseID)+"/records/"+pathSeg(recordID)+"/links", nil, nil, false, "links", &out, opts)
	return out, err
}

// List — Lists a table's records, filtered and sorted.
//
// GET /api/databases/{databaseId}/records
//
// databaseID: The table's id.
func (r *DatabasesRecordsService) List(ctx context.Context, databaseID string, query *DatabasesRecordsListQuery, opts ...CallOption) ([]DatabaseRecord, error) {
	var out []DatabaseRecord
	err := r.c.call(ctx, "GET", "/api/databases/"+pathSeg(databaseID)+"/records", query, nil, false, "records", &out, opts)
	return out, err
}

// Update — Changes a record's cells.
//
// PATCH /api/databases/{databaseId}/records/{recordId}
//
// databaseID: The table's id.
// recordID: The record's id.
func (r *DatabasesRecordsService) Update(ctx context.Context, databaseID string, recordID string, body *UpdateRecordInput, opts ...CallOption) (*DatabaseRecord, error) {
	var out *DatabaseRecord
	err := r.c.call(ctx, "PATCH", "/api/databases/"+pathSeg(databaseID)+"/records/"+pathSeg(recordID), nil, body, false, "record", &out, opts)
	return out, err
}

// DatabasesSharesService — databases.shares
type DatabasesSharesService struct {
	c *Client
}

func newDatabasesSharesService(c *Client) *DatabasesSharesService {
	return &DatabasesSharesService{c: c}
}

// List — Lists who a table is shared with, and at what level.
//
// GET /api/databases/{databaseId}/shares
//
// databaseID: The table's id.
func (r *DatabasesSharesService) List(ctx context.Context, databaseID string, query *DatabasesSharesListQuery, opts ...CallOption) ([]DatabasesSharesListResponseSharesItem, error) {
	var out []DatabasesSharesListResponseSharesItem
	err := r.c.call(ctx, "GET", "/api/databases/"+pathSeg(databaseID)+"/shares", query, nil, false, "shares", &out, opts)
	return out, err
}

// Remove — Stops sharing a table with a person.
//
// DELETE /api/databases/{databaseId}/shares/{userId}
//
// databaseID: The table's id.
// userID: The person's user id.
func (r *DatabasesSharesService) Remove(ctx context.Context, databaseID string, userID string, opts ...CallOption) (*DatabasesSharesRemoveResponse, error) {
	var out *DatabasesSharesRemoveResponse
	err := r.c.call(ctx, "DELETE", "/api/databases/"+pathSeg(databaseID)+"/shares/"+pathSeg(userID), nil, nil, false, "", &out, opts)
	return out, err
}

// Set — Shares a table with a person, or changes their level.
//
// PUT /api/databases/{databaseId}/shares
//
// databaseID: The table's id.
func (r *DatabasesSharesService) Set(ctx context.Context, databaseID string, body *DatabasesSharesSetBody, opts ...CallOption) (*DatabasesSharesSetResponseShare, error) {
	var out *DatabasesSharesSetResponseShare
	err := r.c.call(ctx, "PUT", "/api/databases/"+pathSeg(databaseID)+"/shares", nil, body, false, "share", &out, opts)
	return out, err
}

// DocumentsService — Documents: The Documents hub: documents made of a tree of pages, with comments, versions, share links and mail merge.
type DocumentsService struct {
	c *Client

	Comments  *DocumentsCommentsService
	MailMerge *DocumentsMailMergeService
	Pages     *DocumentsPagesService
	Shares    *DocumentsSharesService
	Versions  *DocumentsVersionsService
}

func newDocumentsService(c *Client) *DocumentsService {
	return &DocumentsService{c: c, Comments: newDocumentsCommentsService(c), MailMerge: newDocumentsMailMergeService(c), Pages: newDocumentsPagesService(c), Shares: newDocumentsSharesService(c), Versions: newDocumentsVersionsService(c)}
}

// Bulk — Deletes, duplicates or changes the visibility of up to 500 documents at once.
//
// Needs `document:update`; each document also needs edit access. A document that fails does not fail the batch: it is listed in `failed` with a reason (`not_found`, `forbidden`, `limit`, …).
//
// POST /api/documents/bulk
func (r *DocumentsService) Bulk(ctx context.Context, body *DocumentsBulkBody, opts ...CallOption) (*DocumentsBulkResponse, error) {
	var out *DocumentsBulkResponse
	err := r.c.call(ctx, "POST", "/api/documents/bulk", nil, body, false, "", &out, opts)
	return out, err
}

// Compare — The paragraphs of two documents, ready to be compared.
//
// Each side holds only the pages the caller may read; a vault page's text is never included. The comparison itself is left to the caller.
//
// GET /api/documents/{documentId}/compare
//
// documentID: The document.
func (r *DocumentsService) Compare(ctx context.Context, documentID string, query *DocumentsCompareQuery, opts ...CallOption) (*DocumentsCompareResponse, error) {
	var out *DocumentsCompareResponse
	err := r.c.call(ctx, "GET", "/api/documents/"+pathSeg(documentID)+"/compare", query, nil, false, "", &out, opts)
	return out, err
}

// Create — Creates a document, empty or from a template.
//
// Needs `document:create` and the Documents hub on the plan; counts against the plan's document limit. The caller becomes its admin.
//
// POST /api/documents
func (r *DocumentsService) Create(ctx context.Context, body *DocumentsCreateBody, opts ...CallOption) (*Document, error) {
	var out *Document
	err := r.c.call(ctx, "POST", "/api/documents", nil, body, false, "document", &out, opts)
	return out, err
}

// Delete — Moves a document to the trash, or deletes a trashed one for good with `purge=1`.
//
// Needs `document:delete` and edit access. Purging refuses (409) a document that is not in the trash, so no single request takes a live document to nothing.
//
// DELETE /api/documents/{documentId}
//
// documentID: The document.
func (r *DocumentsService) Delete(ctx context.Context, documentID string, query *DocumentsDeleteQuery, opts ...CallOption) (*DocumentsDeleteResponse, error) {
	var out *DocumentsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/documents/"+pathSeg(documentID), query, nil, false, "", &out, opts)
	return out, err
}

// Duplicate — Copies a document with the pages the caller may read.
//
// Needs `document:create` and read access; counts against the plan's document limit.
//
// POST /api/documents/{documentId}/duplicate
//
// documentID: The document.
func (r *DocumentsService) Duplicate(ctx context.Context, documentID string, body *DocumentsDuplicateBody, opts ...CallOption) (*Document, error) {
	var out *Document
	err := r.c.call(ctx, "POST", "/api/documents/"+pathSeg(documentID)+"/duplicate", nil, body, false, "document", &out, opts)
	return out, err
}

// Export — Downloads a document as a Word, PDF, Markdown or HTML file.
//
// Answers the file itself. Pages the caller may not read are left out; a vault page is listed as a note, never with its text.
//
// GET /api/documents/{documentId}/export
//
// documentID: The document.
func (r *DocumentsService) Export(ctx context.Context, documentID string, query *DocumentsExportQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/documents/"+pathSeg(documentID)+"/export", query, nil, false, "", &out, opts)
	return out, err
}

// Get — One document, with what the caller may do with it.
//
// GET /api/documents/{documentId}
//
// documentID: The document.
func (r *DocumentsService) Get(ctx context.Context, documentID string, opts ...CallOption) (*Document, error) {
	var out *Document
	err := r.c.call(ctx, "GET", "/api/documents/"+pathSeg(documentID), nil, nil, false, "document", &out, opts)
	return out, err
}

// Google — The Google accounts the caller can import from and, with `connectionId`, their Google documents.
//
// A Google account is connected in the dashboard (Migration Center), which needs `migration:create` — reported as `canConnect`.
//
// GET /api/documents/google
func (r *DocumentsService) Google(ctx context.Context, query *DocumentsGoogleQuery, opts ...CallOption) (*DocumentsGoogleResponse, error) {
	var out *DocumentsGoogleResponse
	err := r.c.call(ctx, "GET", "/api/documents/google", query, nil, false, "", &out, opts)
	return out, err
}

// Import — Imports Markdown, HTML or a Google document into a document as new pages.
//
// Needs `document:update` and edit access; counts against the plan's page limit. Top-level headings become pages and the next level their subpages. Send one of `markdown`, `html`, or `googleConnectionId` with `googleFileId` (GET /api/documents/google lists them). A .docx, .md or .html file can be sent instead as multipart form data in a `file` field (20 MB at most).
//
// POST /api/documents/{documentId}/import
//
// documentID: The document.
func (r *DocumentsService) Import(ctx context.Context, documentID string, body *DocumentsImportBody, opts ...CallOption) (*DocumentsImportResponse, error) {
	var out *DocumentsImportResponse
	err := r.c.call(ctx, "POST", "/api/documents/"+pathSeg(documentID)+"/import", nil, body, false, "", &out, opts)
	return out, err
}

// List — Lists one page of the documents the caller may open, with counts.
//
// Needs `document:read` and the Documents hub on the plan. One of `organizationId` and `workspaceId` is required. With `trash=1` (and `organizationId`) it lists the organisation's trashed documents instead, all at once.
//
// GET /api/documents
func (r *DocumentsService) List(ctx context.Context, query *DocumentsListQuery, opts ...CallOption) ([]Document, error) {
	var out []Document
	err := r.c.call(ctx, "GET", "/api/documents", query, nil, false, "documents", &out, opts)
	return out, err
}

// Move — Moves a document to another workspace of its organisation, or to none.
//
// Needs `document:update`, edit access, and membership of the destination workspace.
//
// PUT /api/documents/{documentId}/move
//
// documentID: The document.
func (r *DocumentsService) Move(ctx context.Context, documentID string, body *DocumentsMoveBody, opts ...CallOption) (*DocumentsMoveResponse, error) {
	var out *DocumentsMoveResponse
	err := r.c.call(ctx, "PUT", "/api/documents/"+pathSeg(documentID)+"/move", nil, body, false, "", &out, opts)
	return out, err
}

// Restore — Takes a document back out of the trash.
//
// Needs `document:delete` and edit access. `restored: false` when it was not in the trash.
//
// POST /api/documents/{documentId}/restore
//
// documentID: The document.
func (r *DocumentsService) Restore(ctx context.Context, documentID string, opts ...CallOption) (bool, error) {
	var out bool
	err := r.c.call(ctx, "POST", "/api/documents/"+pathSeg(documentID)+"/restore", nil, nil, false, "restored", &out, opts)
	return out, err
}

// Search — Searches the names, page titles and text of the documents the caller may open.
//
// Restricted and vault pages are never searched. Fewer than two characters find nothing.
//
// GET /api/documents/search
func (r *DocumentsService) Search(ctx context.Context, query *DocumentsSearchQuery, opts ...CallOption) ([]DocumentsSearchResponseHitsItem, error) {
	var out []DocumentsSearchResponseHitsItem
	err := r.c.call(ctx, "GET", "/api/documents/search", query, nil, false, "hits", &out, opts)
	return out, err
}

// SetFavorite — Stars a document for the caller, or unstars it.
//
// PUT /api/documents/{documentId}/favorite
//
// documentID: The document.
func (r *DocumentsService) SetFavorite(ctx context.Context, documentID string, body *DocumentsSetFavoriteBody, opts ...CallOption) (bool, error) {
	var out bool
	err := r.c.call(ctx, "PUT", "/api/documents/"+pathSeg(documentID)+"/favorite", nil, body, false, "favorite", &out, opts)
	return out, err
}

// Templates — Lists the document templates: the platform's and the organisation's own, most used first.
//
// GET /api/documents/templates
func (r *DocumentsService) Templates(ctx context.Context, query *DocumentsTemplatesQuery, opts ...CallOption) ([]DocumentTemplate, error) {
	var out []DocumentTemplate
	err := r.c.call(ctx, "GET", "/api/documents/templates", query, nil, false, "templates", &out, opts)
	return out, err
}

// Update — Changes a document's name, description, icon, cover, visibility, access list, settings, AI indexing or tags.
//
// Needs `document:update` and edit access. Only the fields sent change. Changing `permissions` needs the document's creator or one of its admins; narrowing access closes the live editing sessions of who lost it.
//
// PATCH /api/documents/{documentId}
//
// documentID: The document.
func (r *DocumentsService) Update(ctx context.Context, documentID string, body *DocumentsUpdateBody, opts ...CallOption) (*Document, error) {
	var out *Document
	err := r.c.call(ctx, "PATCH", "/api/documents/"+pathSeg(documentID), nil, body, false, "document", &out, opts)
	return out, err
}

// DocumentsCommentsService — documents.comments
type DocumentsCommentsService struct {
	c *Client
}

func newDocumentsCommentsService(c *Client) *DocumentsCommentsService {
	return &DocumentsCommentsService{c: c}
}

// Create — Comments on a page, or answers a comment.
//
// Needs `document:resolve` and comment access, and comments allowed on the document. The people named and the thread's participants who can open the page are notified. Threads nest five levels at most.
//
// POST /api/documents/{documentId}/pages/{pageId}/comments
//
// documentID: The document.
// pageID: A page of it.
func (r *DocumentsCommentsService) Create(ctx context.Context, documentID string, pageID string, body *DocumentsCommentsCreateBody, opts ...CallOption) (*DocumentComment, error) {
	var out *DocumentComment
	err := r.c.call(ctx, "POST", "/api/documents/"+pathSeg(documentID)+"/pages/"+pathSeg(pageID)+"/comments", nil, body, false, "comment", &out, opts)
	return out, err
}

// Delete — Deletes a comment (its author only).
//
// DELETE /api/documents/{documentId}/pages/{pageId}/comments/{commentId}
//
// documentID: The document.
// pageID: A page of it.
// commentID: A comment on the page.
func (r *DocumentsCommentsService) Delete(ctx context.Context, documentID string, pageID string, commentID string, opts ...CallOption) (*DocumentsCommentsDeleteResponse, error) {
	var out *DocumentsCommentsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/documents/"+pathSeg(documentID)+"/pages/"+pathSeg(pageID)+"/comments/"+pathSeg(commentID), nil, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the comments on a page, threaded by `parentId`, with their reactions.
//
// GET /api/documents/{documentId}/pages/{pageId}/comments
//
// documentID: The document.
// pageID: A page of it.
func (r *DocumentsCommentsService) List(ctx context.Context, documentID string, pageID string, opts ...CallOption) ([]DocumentComment, error) {
	var out []DocumentComment
	err := r.c.call(ctx, "GET", "/api/documents/"+pathSeg(documentID)+"/pages/"+pathSeg(pageID)+"/comments", nil, nil, false, "comments", &out, opts)
	return out, err
}

// React — Adds the caller's reaction to a comment, or takes it away again.
//
// PUT /api/documents/{documentId}/pages/{pageId}/comments/{commentId}/react
//
// documentID: The document.
// pageID: A page of it.
// commentID: A comment on the page.
func (r *DocumentsCommentsService) React(ctx context.Context, documentID string, pageID string, commentID string, body *DocumentsCommentsReactBody, opts ...CallOption) (bool, error) {
	var out bool
	err := r.c.call(ctx, "PUT", "/api/documents/"+pathSeg(documentID)+"/pages/"+pathSeg(pageID)+"/comments/"+pathSeg(commentID)+"/react", nil, body, false, "reacted", &out, opts)
	return out, err
}

// Update — Edits a comment's text (its author only), or resolves and reopens its thread.
//
// PATCH /api/documents/{documentId}/pages/{pageId}/comments/{commentId}
//
// documentID: The document.
// pageID: A page of it.
// commentID: A comment on the page.
func (r *DocumentsCommentsService) Update(ctx context.Context, documentID string, pageID string, commentID string, body *DocumentsCommentsUpdateBody, opts ...CallOption) (*DocumentComment, error) {
	var out *DocumentComment
	err := r.c.call(ctx, "PATCH", "/api/documents/"+pathSeg(documentID)+"/pages/"+pathSeg(pageID)+"/comments/"+pathSeg(commentID), nil, body, false, "comment", &out, opts)
	return out, err
}

// DocumentsMailMergeService — documents.mailMerge
type DocumentsMailMergeService struct {
	c *Client
}

func newDocumentsMailMergeService(c *Client) *DocumentsMailMergeService {
	return &DocumentsMailMergeService{c: c}
}

// Clear — Clears a document's mail merge source; the fields stay in the text.
//
// DELETE /api/documents/{documentId}/mail-merge
//
// documentID: The document.
func (r *DocumentsMailMergeService) Clear(ctx context.Context, documentID string, opts ...CallOption) (*DocumentsMailMergeClearResponse, error) {
	var out *DocumentsMailMergeClearResponse
	err := r.c.call(ctx, "DELETE", "/api/documents/"+pathSeg(documentID)+"/mail-merge", nil, nil, false, "", &out, opts)
	return out, err
}

// Get — A document's mail merge: its source database and columns, the views to choose from, how many rows it covers, and optionally one row's values.
//
// The rows are always read as the caller. `source` is null when none is set; `refused` says why a set source cannot be read by the caller. `blocked` is the reason an export would refuse the merge.
//
// GET /api/documents/{documentId}/mail-merge
//
// documentID: The document.
func (r *DocumentsMailMergeService) Get(ctx context.Context, documentID string, query *DocumentsMailMergeGetQuery, opts ...CallOption) (*DocumentsMailMergeGetResponse, error) {
	var out *DocumentsMailMergeGetResponse
	err := r.c.call(ctx, "GET", "/api/documents/"+pathSeg(documentID)+"/mail-merge", query, nil, false, "", &out, opts)
	return out, err
}

// Update — Sets or changes a document's mail merge source, view and sort.
//
// Needs `document:update` and edit access. A partial update: `databaseId` is needed only for the first source, and a different one starts over. Only a database of the document's organisation the caller may read, a usable view of it, a column of it.
//
// PUT /api/documents/{documentId}/mail-merge
//
// documentID: The document.
func (r *DocumentsMailMergeService) Update(ctx context.Context, documentID string, body *DocumentsMailMergeUpdateBody, opts ...CallOption) (*DocumentsMailMergeUpdateResponse, error) {
	var out *DocumentsMailMergeUpdateResponse
	err := r.c.call(ctx, "PUT", "/api/documents/"+pathSeg(documentID)+"/mail-merge", nil, body, false, "", &out, opts)
	return out, err
}

// DocumentsPagesService — documents.pages
type DocumentsPagesService struct {
	c *Client
}

func newDocumentsPagesService(c *Client) *DocumentsPagesService {
	return &DocumentsPagesService{c: c}
}

// Backlinks — The pages of the document that link to this page.
//
// GET /api/documents/{documentId}/backlinks/{pageId}
//
// documentID: The document.
// pageID: A page of it.
func (r *DocumentsPagesService) Backlinks(ctx context.Context, documentID string, pageID string, opts ...CallOption) ([]DocumentsPagesBacklinksResponseBacklinksItem, error) {
	var out []DocumentsPagesBacklinksResponseBacklinksItem
	err := r.c.call(ctx, "GET", "/api/documents/"+pathSeg(documentID)+"/backlinks/"+pathSeg(pageID), nil, nil, false, "backlinks", &out, opts)
	return out, err
}

// Create — Adds an empty page to a document, at the end of its parent's children.
//
// Needs `document:update` and edit access; counts against the plan's page limit. Pages nest six levels deep at most (409 `depth`). A sensitivity other than `none` needs sensitive blocks on the plan, else the page is plain.
//
// POST /api/documents/{documentId}/pages
//
// documentID: The document.
func (r *DocumentsPagesService) Create(ctx context.Context, documentID string, body *DocumentsPagesCreateBody, opts ...CallOption) (*DocumentPage, error) {
	var out *DocumentPage
	err := r.c.call(ctx, "POST", "/api/documents/"+pathSeg(documentID)+"/pages", nil, body, false, "page", &out, opts)
	return out, err
}

// Delete — Deletes a page and its subpages.
//
// Refused (403 `pageHasRestrictedSubpages`) when a subpage is one the caller may not edit.
//
// DELETE /api/documents/{documentId}/pages/{pageId}
//
// documentID: The document.
// pageID: A page of it.
func (r *DocumentsPagesService) Delete(ctx context.Context, documentID string, pageID string, opts ...CallOption) (*DocumentsPagesDeleteResponse, error) {
	var out *DocumentsPagesDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/documents/"+pathSeg(documentID)+"/pages/"+pathSeg(pageID), nil, nil, false, "", &out, opts)
	return out, err
}

// Get — One page with its content.
//
// GET /api/documents/{documentId}/pages/{pageId}
//
// documentID: The document.
// pageID: A page of it.
func (r *DocumentsPagesService) Get(ctx context.Context, documentID string, pageID string, opts ...CallOption) (*DocumentPage, error) {
	var out *DocumentPage
	err := r.c.call(ctx, "GET", "/api/documents/"+pathSeg(documentID)+"/pages/"+pathSeg(pageID), nil, nil, false, "page", &out, opts)
	return out, err
}

// List — Lists the pages of a document the caller may read, in tree order, without their content.
//
// GET /api/documents/{documentId}/pages
//
// documentID: The document.
func (r *DocumentsPagesService) List(ctx context.Context, documentID string, opts ...CallOption) ([]DocumentPage, error) {
	var out []DocumentPage
	err := r.c.call(ctx, "GET", "/api/documents/"+pathSeg(documentID)+"/pages", nil, nil, false, "pages", &out, opts)
	return out, err
}

// Update — Changes a page's details (`mode: "meta"`) or replaces its content (`mode: "content"`).
//
// Needs `document:update` and edit access to the page.
//
// - **meta**: title, icon, place in the tree, status, sensitivity, visibility, header. Publishing a page of a document that requires approval files an approval instead (`approvalWorkflowId`).
// - **content**: `contentJson` (a ProseMirror document) and `contentHtml`. With `baseUpdatedAt`, a page changed since answers 409 `conflictReload` with what it holds now. With `locale`, the content is that language's translation. A version is kept at most once an hour, or when `snapshotVersion` is set.
//
// PATCH /api/documents/{documentId}/pages/{pageId}
//
// documentID: The document.
// pageID: A page of it.
func (r *DocumentsPagesService) Update(ctx context.Context, documentID string, pageID string, body *DocumentsPagesUpdateBody, opts ...CallOption) (*DocumentPage, error) {
	var out *DocumentPage
	err := r.c.call(ctx, "PATCH", "/api/documents/"+pathSeg(documentID)+"/pages/"+pathSeg(pageID), nil, body, false, "page", &out, opts)
	return out, err
}

// DocumentsSharesService — documents.shares
type DocumentsSharesService struct {
	c *Client
}

func newDocumentsSharesService(c *Client) *DocumentsSharesService {
	return &DocumentsSharesService{c: c}
}

// Create — Creates a public share link to a document or one of its pages.
//
// Needs `document:share`, edit access, public pages on the plan; counts against the plan's limit of public documents. Anyone with the link can read what it shares.
//
// POST /api/documents/{documentId}/shares
//
// documentID: The document.
func (r *DocumentsSharesService) Create(ctx context.Context, documentID string, body *DocumentsSharesCreateBody, opts ...CallOption) (*DocumentShare, error) {
	var out *DocumentShare
	err := r.c.call(ctx, "POST", "/api/documents/"+pathSeg(documentID)+"/shares", nil, body, false, "share", &out, opts)
	return out, err
}

// List — Lists a document's live share links.
//
// The link's secret `token` only for a caller who may edit the document.
//
// GET /api/documents/{documentId}/shares
//
// documentID: The document.
func (r *DocumentsSharesService) List(ctx context.Context, documentID string, opts ...CallOption) ([]DocumentShare, error) {
	var out []DocumentShare
	err := r.c.call(ctx, "GET", "/api/documents/"+pathSeg(documentID)+"/shares", nil, nil, false, "shares", &out, opts)
	return out, err
}

// Revoke — Revokes a share link; it stops working at once.
//
// DELETE /api/documents/{documentId}/shares/{shareId}
//
// documentID: The document.
// shareID: The share link.
func (r *DocumentsSharesService) Revoke(ctx context.Context, documentID string, shareID string, opts ...CallOption) (*DocumentsSharesRevokeResponse, error) {
	var out *DocumentsSharesRevokeResponse
	err := r.c.call(ctx, "DELETE", "/api/documents/"+pathSeg(documentID)+"/shares/"+pathSeg(shareID), nil, nil, false, "", &out, opts)
	return out, err
}

// DocumentsVersionsService — documents.versions
type DocumentsVersionsService struct {
	c *Client
}

func newDocumentsVersionsService(c *Client) *DocumentsVersionsService {
	return &DocumentsVersionsService{c: c}
}

// Get — One saved version of a page, with its content.
//
// GET /api/documents/{documentId}/pages/{pageId}/versions/{versionId}
//
// documentID: The document.
// pageID: A page of it.
// versionID: The version.
func (r *DocumentsVersionsService) Get(ctx context.Context, documentID string, pageID string, versionID string, opts ...CallOption) (*DocumentsVersionsGetResponseVersion, error) {
	var out *DocumentsVersionsGetResponseVersion
	err := r.c.call(ctx, "GET", "/api/documents/"+pathSeg(documentID)+"/pages/"+pathSeg(pageID)+"/versions/"+pathSeg(versionID), nil, nil, false, "version", &out, opts)
	return out, err
}

// List — Lists a page's saved versions, newest first, without their content.
//
// Without longer version history on the plan, only the 30 newest (`trimmed`).
//
// GET /api/documents/{documentId}/pages/{pageId}/versions
//
// documentID: The document.
// pageID: A page of it.
func (r *DocumentsVersionsService) List(ctx context.Context, documentID string, pageID string, opts ...CallOption) ([]DocumentVersionSummary, error) {
	var out []DocumentVersionSummary
	err := r.c.call(ctx, "GET", "/api/documents/"+pathSeg(documentID)+"/pages/"+pathSeg(pageID)+"/versions", nil, nil, false, "versions", &out, opts)
	return out, err
}

// Restore — Puts a saved version back as the page's content.
//
// POST /api/documents/{documentId}/pages/{pageId}/versions/{versionId}
//
// documentID: The document.
// pageID: A page of it.
// versionID: The version.
func (r *DocumentsVersionsService) Restore(ctx context.Context, documentID string, pageID string, versionID string, opts ...CallOption) (*DocumentPage, error) {
	var out *DocumentPage
	err := r.c.call(ctx, "POST", "/api/documents/"+pathSeg(documentID)+"/pages/"+pathSeg(pageID)+"/versions/"+pathSeg(versionID), nil, nil, false, "page", &out, opts)
	return out, err
}

// EndpointsService — API endpoints: APIs you build on the platform and publish to your own callers.
type EndpointsService struct {
	c *Client
}

func newEndpointsService(c *Client) *EndpointsService {
	return &EndpointsService{c: c}
}

// Consumers — Lists who calls an endpoint: each API key with the routes it called and how often.
//
// GET /api/endpoints/{endpointId}/consumers
//
// endpointID: The endpoint's id.
func (r *EndpointsService) Consumers(ctx context.Context, endpointID string, opts ...CallOption) ([]EndpointsConsumersResponseDataItemsItem, error) {
	var out []EndpointsConsumersResponseDataItemsItem
	err := r.c.call(ctx, "GET", "/api/endpoints/"+pathSeg(endpointID)+"/consumers", nil, nil, false, "data.items", &out, opts)
	return out, err
}

// Create — Creates an API endpoint.
//
// POST /api/endpoints
func (r *EndpointsService) Create(ctx context.Context, body *CreateEndpointInput, opts ...CallOption) (*Endpoint, error) {
	var out *Endpoint
	err := r.c.call(ctx, "POST", "/api/endpoints", nil, body, false, "endpoint", &out, opts)
	return out, err
}

// CreateAPIKey — Creates an API key for an endpoint's callers; the answer carries the key, once.
//
// POST /api/endpoints/{endpointId}/api-keys
//
// endpointID: The endpoint's id.
func (r *EndpointsService) CreateAPIKey(ctx context.Context, endpointID string, body *CreateEndpointApiKeyInput, opts ...CallOption) (*EndpointApiKey, error) {
	var out *EndpointApiKey
	err := r.c.call(ctx, "POST", "/api/endpoints/"+pathSeg(endpointID)+"/api-keys", nil, body, false, "", &out, opts)
	return out, err
}

// CreateEnvVar — Adds an environment variable to an endpoint.
//
// POST /api/endpoints/{endpointId}/env
//
// endpointID: The endpoint's id.
func (r *EndpointsService) CreateEnvVar(ctx context.Context, endpointID string, body *CreateEndpointEnvVarInput, opts ...CallOption) (*EndpointEnvVar, error) {
	var out *EndpointEnvVar
	err := r.c.call(ctx, "POST", "/api/endpoints/"+pathSeg(endpointID)+"/env", nil, body, false, "envVar", &out, opts)
	return out, err
}

// CreateWebhook — Adds a webhook to an endpoint; the answer carries its signing secret, once.
//
// POST /api/endpoints/{endpointId}/webhooks
//
// endpointID: The endpoint's id.
func (r *EndpointsService) CreateWebhook(ctx context.Context, endpointID string, body *CreateEndpointWebhookInput, opts ...CallOption) (*EndpointsCreateWebhookResponse, error) {
	var out *EndpointsCreateWebhookResponse
	err := r.c.call(ctx, "POST", "/api/endpoints/"+pathSeg(endpointID)+"/webhooks", nil, body, false, "", &out, opts)
	return out, err
}

// Delete — Moves an endpoint to the trash; its URL stops answering.
//
// DELETE /api/endpoints/{endpointId}
//
// endpointID: The endpoint's id.
func (r *EndpointsService) Delete(ctx context.Context, endpointID string, opts ...CallOption) (*EndpointsDeleteResponse, error) {
	var out *EndpointsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/endpoints/"+pathSeg(endpointID), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteAPIKey — Revokes an endpoint API key.
//
// DELETE /api/endpoints/{endpointId}/api-keys/{keyId}
//
// endpointID: The endpoint's id.
// keyID: The API key's id.
func (r *EndpointsService) DeleteAPIKey(ctx context.Context, endpointID string, keyID string, opts ...CallOption) (*EndpointsDeleteApiKeyResponse, error) {
	var out *EndpointsDeleteApiKeyResponse
	err := r.c.call(ctx, "DELETE", "/api/endpoints/"+pathSeg(endpointID)+"/api-keys/"+pathSeg(keyID), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteEnvVar — Removes an environment variable from an endpoint.
//
// DELETE /api/endpoints/{endpointId}/env/{varId}
//
// endpointID: The endpoint's id.
// varID: The variable's id.
func (r *EndpointsService) DeleteEnvVar(ctx context.Context, endpointID string, varID string, opts ...CallOption) (*EndpointsDeleteEnvVarResponse, error) {
	var out *EndpointsDeleteEnvVarResponse
	err := r.c.call(ctx, "DELETE", "/api/endpoints/"+pathSeg(endpointID)+"/env/"+pathSeg(varID), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteWebhook — Removes a webhook from an endpoint.
//
// DELETE /api/endpoints/{endpointId}/webhooks/{webhookId}
//
// endpointID: The endpoint's id.
// webhookID: The webhook's id.
func (r *EndpointsService) DeleteWebhook(ctx context.Context, endpointID string, webhookID string, opts ...CallOption) (*EndpointsDeleteWebhookResponse, error) {
	var out *EndpointsDeleteWebhookResponse
	err := r.c.call(ctx, "DELETE", "/api/endpoints/"+pathSeg(endpointID)+"/webhooks/"+pathSeg(webhookID), nil, nil, false, "", &out, opts)
	return out, err
}

// Dependencies — Lists what an endpoint's routes use: workflows, tables and environment variables.
//
// GET /api/endpoints/{endpointId}/dependencies
//
// endpointID: The endpoint's id.
func (r *EndpointsService) Dependencies(ctx context.Context, endpointID string, opts ...CallOption) (*EndpointsDependenciesResponseData, error) {
	var out *EndpointsDependenciesResponseData
	err := r.c.call(ctx, "GET", "/api/endpoints/"+pathSeg(endpointID)+"/dependencies", nil, nil, false, "data", &out, opts)
	return out, err
}

// ExplainCapture — Explains with AI why a recorded request failed and how to fix it.
//
// Counts against the organization's AI budget.
//
// POST /api/endpoints/{endpointId}/explain
//
// endpointID: The endpoint's id.
func (r *EndpointsService) ExplainCapture(ctx context.Context, endpointID string, body *EndpointsExplainCaptureBody, opts ...CallOption) (*EndpointsExplainCaptureResponseData, error) {
	var out *EndpointsExplainCaptureResponseData
	err := r.c.call(ctx, "POST", "/api/endpoints/"+pathSeg(endpointID)+"/explain", nil, body, false, "data", &out, opts)
	return out, err
}

// Get — Returns an endpoint with its versions, routes and the URL callers reach it at.
//
// GET /api/endpoints/{endpointId}
//
// endpointID: The endpoint's id.
func (r *EndpointsService) Get(ctx context.Context, endpointID string, opts ...CallOption) (*Endpoint, error) {
	var out *Endpoint
	err := r.c.call(ctx, "GET", "/api/endpoints/"+pathSeg(endpointID), nil, nil, false, "", &out, opts)
	return out, err
}

// GetActivity — Lists what happened to an endpoint — deploys, changes, errors — newest first.
//
// GET /api/endpoints/{endpointId}/activity
//
// endpointID: The endpoint's id.
func (r *EndpointsService) GetActivity(ctx context.Context, endpointID string, query *EndpointsGetActivityQuery, opts ...CallOption) ([]EndpointsGetActivityResponseActivitiesItem, error) {
	var out []EndpointsGetActivityResponseActivitiesItem
	err := r.c.call(ctx, "GET", "/api/endpoints/"+pathSeg(endpointID)+"/activity", query, nil, false, "activities", &out, opts)
	return out, err
}

// GetAnalytics — Returns an endpoint's requests, errors and response times, overall and per version.
//
// GET /api/endpoints/{endpointId}/analytics
//
// endpointID: The endpoint's id.
func (r *EndpointsService) GetAnalytics(ctx context.Context, endpointID string, query *EndpointsGetAnalyticsQuery, opts ...CallOption) (*EndpointsGetAnalyticsResponse, error) {
	var out *EndpointsGetAnalyticsResponse
	err := r.c.call(ctx, "GET", "/api/endpoints/"+pathSeg(endpointID)+"/analytics", query, nil, false, "", &out, opts)
	return out, err
}

// GetAuthConfig — Returns how an endpoint checks a JWT or OAuth2 bearer token; never a secret, only whether one is stored.
//
// GET /api/endpoints/{endpointId}/auth-config
//
// endpointID: The endpoint's id.
func (r *EndpointsService) GetAuthConfig(ctx context.Context, endpointID string, opts ...CallOption) (*EndpointAuthConfig, error) {
	var out *EndpointAuthConfig
	err := r.c.call(ctx, "GET", "/api/endpoints/"+pathSeg(endpointID)+"/auth-config", nil, nil, false, "", &out, opts)
	return out, err
}

// GetCapture — Returns one recorded request of an endpoint, with its request and response.
//
// GET /api/endpoints/{endpointId}/captures/{captureId}
//
// endpointID: The endpoint's id.
// captureID: The capture's id.
func (r *EndpointsService) GetCapture(ctx context.Context, endpointID string, captureID string, opts ...CallOption) (*EndpointsGetCaptureResponseData, error) {
	var out *EndpointsGetCaptureResponseData
	err := r.c.call(ctx, "GET", "/api/endpoints/"+pathSeg(endpointID)+"/captures/"+pathSeg(captureID), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetDocs — Returns an OpenAPI description of an endpoint for its callers.
//
// With `format=yaml` the answer is YAML text.
//
// GET /api/endpoints/{endpointId}/docs
//
// endpointID: The endpoint's id.
func (r *EndpointsService) GetDocs(ctx context.Context, endpointID string, query *EndpointsGetDocsQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/endpoints/"+pathSeg(endpointID)+"/docs", query, nil, false, "", &out, opts)
	return out, err
}

// GetTimings — Returns where the time of one request went, step by step.
//
// GET /api/endpoints/{endpointId}/waterfall/{metricId}
//
// endpointID: The endpoint's id.
// metricID: The request's metric id.
func (r *EndpointsService) GetTimings(ctx context.Context, endpointID string, metricID string, opts ...CallOption) (*EndpointsGetTimingsResponseData, error) {
	var out *EndpointsGetTimingsResponseData
	err := r.c.call(ctx, "GET", "/api/endpoints/"+pathSeg(endpointID)+"/waterfall/"+pathSeg(metricID), nil, nil, false, "data", &out, opts)
	return out, err
}

// List — Lists the API endpoints of an organization the caller may open.
//
// GET /api/endpoints
func (r *EndpointsService) List(ctx context.Context, query *EndpointsListQuery, opts ...CallOption) ([]Endpoint, error) {
	var out []Endpoint
	err := r.c.call(ctx, "GET", "/api/endpoints", query, nil, false, "endpoints", &out, opts)
	return out, err
}

// ListAPIKeys — Lists the API keys callers use for an endpoint, without the keys themselves.
//
// GET /api/endpoints/{endpointId}/api-keys
//
// endpointID: The endpoint's id.
func (r *EndpointsService) ListAPIKeys(ctx context.Context, endpointID string, opts ...CallOption) ([]EndpointApiKey, error) {
	var out []EndpointApiKey
	err := r.c.call(ctx, "GET", "/api/endpoints/"+pathSeg(endpointID)+"/api-keys", nil, nil, false, "", &out, opts)
	return out, err
}

// ListCaptures — Lists the recorded requests of an endpoint, newest first.
//
// GET /api/endpoints/{endpointId}/captures
//
// endpointID: The endpoint's id.
func (r *EndpointsService) ListCaptures(ctx context.Context, endpointID string, query *EndpointsListCapturesQuery, opts ...CallOption) ([]EndpointsListCapturesResponseDataItemsItem, error) {
	var out []EndpointsListCapturesResponseDataItemsItem
	err := r.c.call(ctx, "GET", "/api/endpoints/"+pathSeg(endpointID)+"/captures", query, nil, false, "data.items", &out, opts)
	return out, err
}

// ListEnvVars — Lists an endpoint's environment variables; secret values are masked.
//
// GET /api/endpoints/{endpointId}/env
//
// endpointID: The endpoint's id.
func (r *EndpointsService) ListEnvVars(ctx context.Context, endpointID string, opts ...CallOption) ([]EndpointEnvVar, error) {
	var out []EndpointEnvVar
	err := r.c.call(ctx, "GET", "/api/endpoints/"+pathSeg(endpointID)+"/env", nil, nil, false, "", &out, opts)
	return out, err
}

// ListWebhooks — Lists the webhooks an endpoint calls on its events, without their secrets.
//
// GET /api/endpoints/{endpointId}/webhooks
//
// endpointID: The endpoint's id.
func (r *EndpointsService) ListWebhooks(ctx context.Context, endpointID string, opts ...CallOption) ([]EndpointWebhook, error) {
	var out []EndpointWebhook
	err := r.c.call(ctx, "GET", "/api/endpoints/"+pathSeg(endpointID)+"/webhooks", nil, nil, false, "", &out, opts)
	return out, err
}

// LogActivity — Writes a note into an endpoint's activity — a deploy, a rollback, a config change.
//
// POST /api/endpoints/{endpointId}/activity
//
// endpointID: The endpoint's id.
func (r *EndpointsService) LogActivity(ctx context.Context, endpointID string, body *EndpointsLogActivityBody, opts ...CallOption) (*EndpointsLogActivityResponse, error) {
	var out *EndpointsLogActivityResponse
	err := r.c.call(ctx, "POST", "/api/endpoints/"+pathSeg(endpointID)+"/activity", nil, body, false, "", &out, opts)
	return out, err
}

// Replace — Changes an endpoint — its routes, versions, auth or settings — and records the change in its history.
//
// PUT /api/endpoints/{endpointId}
//
// endpointID: The endpoint's id.
func (r *EndpointsService) Replace(ctx context.Context, endpointID string, body *UpdateEndpointInput, opts ...CallOption) (*Endpoint, error) {
	var out *Endpoint
	err := r.c.call(ctx, "PUT", "/api/endpoints/"+pathSeg(endpointID), nil, body, false, "", &out, opts)
	return out, err
}

// ShadowResults — Compares a shadow version's answers with the live version's, request by request.
//
// GET /api/endpoints/{endpointId}/shadow-results
//
// endpointID: The endpoint's id.
func (r *EndpointsService) ShadowResults(ctx context.Context, endpointID string, query *EndpointsShadowResultsQuery, opts ...CallOption) (*EndpointsShadowResultsResponseData, error) {
	var out *EndpointsShadowResultsResponseData
	err := r.c.call(ctx, "GET", "/api/endpoints/"+pathSeg(endpointID)+"/shadow-results", query, nil, false, "data", &out, opts)
	return out, err
}

// SuspendAPIKey — Suspends an endpoint API key, or resumes it.
//
// POST /api/endpoints/{endpointId}/api-keys/{keyId}/suspend
//
// endpointID: The endpoint's id.
// keyID: The API key's id.
func (r *EndpointsService) SuspendAPIKey(ctx context.Context, endpointID string, keyID string, body *EndpointsSuspendApiKeyBody, opts ...CallOption) (*EndpointsSuspendApiKeyResponseData, error) {
	var out *EndpointsSuspendApiKeyResponseData
	err := r.c.call(ctx, "POST", "/api/endpoints/"+pathSeg(endpointID)+"/api-keys/"+pathSeg(keyID)+"/suspend", nil, body, false, "data", &out, opts)
	return out, err
}

// TestAuth — Checks whether the endpoint's saved token setup accepts a token, and if not, why.
//
// The token is used for this check only; it is neither logged nor stored.
//
// POST /api/endpoints/{endpointId}/auth-test
//
// endpointID: The endpoint's id.
func (r *EndpointsService) TestAuth(ctx context.Context, endpointID string, body *EndpointsTestAuthBody, opts ...CallOption) (*EndpointsTestAuthResponse, error) {
	var out *EndpointsTestAuthResponse
	err := r.c.call(ctx, "POST", "/api/endpoints/"+pathSeg(endpointID)+"/auth-test", nil, body, false, "", &out, opts)
	return out, err
}

// Update — Changes an endpoint without a new history entry — how a restore from its history is saved.
//
// PATCH /api/endpoints/{endpointId}
//
// endpointID: The endpoint's id.
func (r *EndpointsService) Update(ctx context.Context, endpointID string, body *UpdateEndpointInput, opts ...CallOption) (*Endpoint, error) {
	var out *Endpoint
	err := r.c.call(ctx, "PATCH", "/api/endpoints/"+pathSeg(endpointID), nil, body, false, "endpoint", &out, opts)
	return out, err
}

// UpdateAuthConfig — Sets how an endpoint checks a JWT or OAuth2 bearer token.
//
// A blank secret keeps the stored one. An HS256 secret has at least 32 bytes.
//
// PUT /api/endpoints/{endpointId}/auth-config
//
// endpointID: The endpoint's id.
func (r *EndpointsService) UpdateAuthConfig(ctx context.Context, endpointID string, body *EndpointsUpdateAuthConfigBody, opts ...CallOption) (*EndpointAuthConfig, error) {
	var out *EndpointAuthConfig
	err := r.c.call(ctx, "PUT", "/api/endpoints/"+pathSeg(endpointID)+"/auth-config", nil, body, false, "", &out, opts)
	return out, err
}

// UpdateEnvVar — Changes an endpoint's environment variable; only the fields sent change.
//
// PATCH /api/endpoints/{endpointId}/env/{varId}
//
// endpointID: The endpoint's id.
// varID: The variable's id.
func (r *EndpointsService) UpdateEnvVar(ctx context.Context, endpointID string, varID string, body *EndpointsUpdateEnvVarBody, opts ...CallOption) (*EndpointEnvVar, error) {
	var out *EndpointEnvVar
	err := r.c.call(ctx, "PATCH", "/api/endpoints/"+pathSeg(endpointID)+"/env/"+pathSeg(varID), nil, body, false, "envVar", &out, opts)
	return out, err
}

// UpdateWebhook — Changes an endpoint webhook's URL, events or state.
//
// PATCH /api/endpoints/{endpointId}/webhooks/{webhookId}
//
// endpointID: The endpoint's id.
// webhookID: The webhook's id.
func (r *EndpointsService) UpdateWebhook(ctx context.Context, endpointID string, webhookID string, body *EndpointsUpdateWebhookBody, opts ...CallOption) (*EndpointWebhook, error) {
	var out *EndpointWebhook
	err := r.c.call(ctx, "PATCH", "/api/endpoints/"+pathSeg(endpointID)+"/webhooks/"+pathSeg(webhookID), nil, body, false, "", &out, opts)
	return out, err
}

// FeatureGatesService — Feature gates
type FeatureGatesService struct {
	c *Client
}

func newFeatureGatesService(c *Client) *FeatureGatesService {
	return &FeatureGatesService{c: c}
}

// Get — Returns which features are on for the caller, by gate key.
//
// Each gate is evaluated for the caller, their roles and, with `organizationId`, that organisation's plan. `platform:marketplace` says whether the marketplace is switched on; while it is off every `marketplace*` gate is `false`. Cached privately for 60 seconds.
//
// GET /api/feature-gates
func (r *FeatureGatesService) Get(ctx context.Context, query *FeatureGatesGetQuery, opts ...CallOption) (FeatureGates, error) {
	var out FeatureGates
	err := r.c.call(ctx, "GET", "/api/feature-gates", query, nil, false, "data", &out, opts)
	return out, err
}

// FeatureRequestsService — Feature requests: The platform's public roadmap: ideas every signed-in person can submit, vote on and discuss.
type FeatureRequestsService struct {
	c *Client
}

func newFeatureRequestsService(c *Client) *FeatureRequestsService {
	return &FeatureRequestsService{c: c}
}

// AddComment — Comments on a feature request.
//
// Its submitter is notified.
//
// POST /api/feature-requests/{id}/comments
//
// id: The feature request's id.
func (r *FeatureRequestsService) AddComment(ctx context.Context, id string, body *FeatureRequestsAddCommentBody, opts ...CallOption) (*FeatureRequestComment, error) {
	var out *FeatureRequestComment
	err := r.c.call(ctx, "POST", "/api/feature-requests/"+pathSeg(id)+"/comments", nil, body, false, "data", &out, opts)
	return out, err
}

// Create — Submits a feature request, with the caller's vote on it.
//
// The platform's administrators are notified.
//
// POST /api/feature-requests
func (r *FeatureRequestsService) Create(ctx context.Context, body *FeatureRequestsCreateBody, opts ...CallOption) (*FeatureRequest, error) {
	var out *FeatureRequest
	err := r.c.call(ctx, "POST", "/api/feature-requests", nil, body, false, "data", &out, opts)
	return out, err
}

// Get — Returns one feature request with its comments and whether the caller voted.
//
// GET /api/feature-requests/{id}
//
// id: The feature request's id.
func (r *FeatureRequestsService) Get(ctx context.Context, id string, opts ...CallOption) (*FeatureRequestsGetResponseData, error) {
	var out *FeatureRequestsGetResponseData
	err := r.c.call(ctx, "GET", "/api/feature-requests/"+pathSeg(id), nil, nil, false, "data", &out, opts)
	return out, err
}

// List — Lists up to 500 feature requests, most voted or newest first, with whether the caller voted.
//
// GET /api/feature-requests
func (r *FeatureRequestsService) List(ctx context.Context, query *FeatureRequestsListQuery, opts ...CallOption) ([]FeatureRequestsListResponseDataItem, error) {
	var out []FeatureRequestsListResponseDataItem
	err := r.c.call(ctx, "GET", "/api/feature-requests", query, nil, false, "data", &out, opts)
	return out, err
}

// ListComments — Lists a feature request's comments, oldest first.
//
// GET /api/feature-requests/{id}/comments
//
// id: The feature request's id.
func (r *FeatureRequestsService) ListComments(ctx context.Context, id string, opts ...CallOption) ([]FeatureRequestsListCommentsResponseDataItem, error) {
	var out []FeatureRequestsListCommentsResponseDataItem
	err := r.c.call(ctx, "GET", "/api/feature-requests/"+pathSeg(id)+"/comments", nil, nil, false, "data", &out, opts)
	return out, err
}

// Unvote — Takes back the caller's vote on a feature request.
//
// Answers 404 when the caller had not voted.
//
// DELETE /api/feature-requests/{id}/vote
//
// id: The feature request's id.
func (r *FeatureRequestsService) Unvote(ctx context.Context, id string, opts ...CallOption) (*FeatureRequestsUnvoteResponseData, error) {
	var out *FeatureRequestsUnvoteResponseData
	err := r.c.call(ctx, "DELETE", "/api/feature-requests/"+pathSeg(id)+"/vote", nil, nil, false, "data", &out, opts)
	return out, err
}

// Vote — Adds the caller's vote to a feature request.
//
// A second vote answers 409.
//
// POST /api/feature-requests/{id}/vote
//
// id: The feature request's id.
func (r *FeatureRequestsService) Vote(ctx context.Context, id string, opts ...CallOption) (*FeatureRequestsVoteResponseData, error) {
	var out *FeatureRequestsVoteResponseData
	err := r.c.call(ctx, "POST", "/api/feature-requests/"+pathSeg(id)+"/vote", nil, nil, false, "data", &out, opts)
	return out, err
}

// HelpService — Help
type HelpService struct {
	c *Client
}

func newHelpService(c *Client) *HelpService {
	return &HelpService{c: c}
}

// GetArticle — Returns one published help article with its text.
//
// By id only; to find an article by its slug, list with `slug`.
//
// GET /api/help/articles/{id}
//
// id: The article's id.
func (r *HelpService) GetArticle(ctx context.Context, id string, opts ...CallOption) (*HelpArticle, error) {
	var out *HelpArticle
	err := r.c.call(ctx, "GET", "/api/help/articles/"+pathSeg(id), nil, nil, false, "data", &out, opts)
	return out, err
}

// ListArticles — Lists the published help articles of one language, with the help categories.
//
// The list leaves out each article's text; `help.getArticle` returns it. With `search`, the articles are ranked by relevance (full text over title, excerpt, tags and text, at most 20); otherwise by their sort order.
//
// GET /api/help/articles
func (r *HelpService) ListArticles(ctx context.Context, query *HelpListArticlesQuery, opts ...CallOption) (*HelpListArticlesResponseData, error) {
	var out *HelpListArticlesResponseData
	err := r.c.call(ctx, "GET", "/api/help/articles", query, nil, false, "data", &out, opts)
	return out, err
}

// InvoicingService — Invoicing: Quotes, order confirmations, invoices and credit notes with e-invoices (XRechnung, ZUGFeRD); customers, issuing companies, number ranges and PDF templates; payments and dunning; recurring invoices; received e-invoices, purchase invoices, suppliers and SEPA payment runs; the books (EÜR or double entry) with DATEV export. Every operation works in one organisation (`organizationId`) and needs the role's `invoicing:*` permissions held at organisation level. A finalized document is never changed or deleted: it is corrected with a credit note or a cancellation.
type InvoicingService struct {
	c *Client

	ApprovalRules *InvoicingApprovalRulesService
	Customers     *InvoicingCustomersService
	Documents     *InvoicingDocumentsService
	Dunning       *InvoicingDunningService
	Inbound       *InvoicingInboundService
	Issuers       *InvoicingIssuersService
	Ledger        *InvoicingLedgerService
	NumberRanges  *InvoicingNumberRangesService
	PaymentRuns   *InvoicingPaymentRunsService
	Payments      *InvoicingPaymentsService
	Purchases     *InvoicingPurchasesService
	Recurring     *InvoicingRecurringService
	Reports       *InvoicingReportsService
	Suppliers     *InvoicingSuppliersService
	Templates     *InvoicingTemplatesService
}

func newInvoicingService(c *Client) *InvoicingService {
	return &InvoicingService{c: c, ApprovalRules: newInvoicingApprovalRulesService(c), Customers: newInvoicingCustomersService(c), Documents: newInvoicingDocumentsService(c), Dunning: newInvoicingDunningService(c), Inbound: newInvoicingInboundService(c), Issuers: newInvoicingIssuersService(c), Ledger: newInvoicingLedgerService(c), NumberRanges: newInvoicingNumberRangesService(c), PaymentRuns: newInvoicingPaymentRunsService(c), Payments: newInvoicingPaymentsService(c), Purchases: newInvoicingPurchasesService(c), Recurring: newInvoicingRecurringService(c), Reports: newInvoicingReportsService(c), Suppliers: newInvoicingSuppliersService(c), Templates: newInvoicingTemplatesService(c)}
}

// InvoicingApprovalRulesService — invoicing.approvalRules
type InvoicingApprovalRulesService struct {
	c *Client
}

func newInvoicingApprovalRulesService(c *Client) *InvoicingApprovalRulesService {
	return &InvoicingApprovalRulesService{c: c}
}

// Create — Creates a purchase approval rule: the first active rule whose conditions all hold decides who approves, in one or two steps.
//
// Needs `invoicing:settings` and purchase approvals on the plan.
//
// POST /api/invoicing/purchase-approval-rules — Token permission: write:invoicing.
func (r *InvoicingApprovalRulesService) Create(ctx context.Context, body *InvoicingApprovalRulesCreateBody, query *InvoicingApprovalRulesCreateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/purchase-approval-rules", query, body, false, "rule", &out, opts)
	return out, err
}

// Delete — Deletes a purchase approval rule.
//
// DELETE /api/invoicing/purchase-approval-rules/{ruleId} — Token permission: write:invoicing.
//
// ruleID: The rule.
func (r *InvoicingApprovalRulesService) Delete(ctx context.Context, ruleID string, query *InvoicingApprovalRulesDeleteQuery, opts ...CallOption) (*InvoicingApprovalRulesDeleteResponse, error) {
	var out *InvoicingApprovalRulesDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/invoicing/purchase-approval-rules/"+pathSeg(ruleID), query, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the purchase approval rules in their order, and the roles and members a step can name.
//
// Needs `invoicing:settings` and purchase approvals on the plan.
//
// GET /api/invoicing/purchase-approval-rules — Token permission: read:invoicing.
func (r *InvoicingApprovalRulesService) List(ctx context.Context, query *InvoicingApprovalRulesListQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/purchase-approval-rules", query, nil, false, "rules", &out, opts)
	return out, err
}

// Update — Changes a purchase approval rule; invoices already waiting keep their steps.
//
// PATCH /api/invoicing/purchase-approval-rules/{ruleId} — Token permission: write:invoicing.
//
// ruleID: The rule.
func (r *InvoicingApprovalRulesService) Update(ctx context.Context, ruleID string, body *InvoicingApprovalRulesUpdateBody, query *InvoicingApprovalRulesUpdateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "PATCH", "/api/invoicing/purchase-approval-rules/"+pathSeg(ruleID), query, body, false, "rule", &out, opts)
	return out, err
}

// InvoicingCustomersService — invoicing.customers
type InvoicingCustomersService struct {
	c *Client
}

func newInvoicingCustomersService(c *Client) *InvoicingCustomersService {
	return &InvoicingCustomersService{c: c}
}

// CheckVAT — Checks a customer's stored VAT ID again with VIES / HMRC and stores the answer.
//
// POST /api/invoicing/customers/{customerId}/vat-check — Token permission: write:invoicing.
//
// customerID: The customer.
func (r *InvoicingCustomersService) CheckVAT(ctx context.Context, customerID string, query *InvoicingCustomersCheckVatQuery, opts ...CallOption) (*InvoicingCustomersCheckVatResponse, error) {
	var out *InvoicingCustomersCheckVatResponse
	err := r.c.call(ctx, "POST", "/api/invoicing/customers/"+pathSeg(customerID)+"/vat-check", query, nil, false, "", &out, opts)
	return out, err
}

// Create — Creates a customer; their VAT ID is checked with VIES / HMRC.
//
// Needs `invoicing:create`. Without a number of its own one is drawn from the customer range. Reverse charge applies only to a validated VAT ID.
//
// POST /api/invoicing/customers — Token permission: write:invoicing.
func (r *InvoicingCustomersService) Create(ctx context.Context, body *InvoicingCustomerInput, query *InvoicingCustomersCreateQuery, opts ...CallOption) (*InvoicingCustomer, error) {
	var out *InvoicingCustomer
	err := r.c.call(ctx, "POST", "/api/invoicing/customers", query, body, false, "customer", &out, opts)
	return out, err
}

// CreateFromCRM — Creates a customer from a CRM record, linked to it.
//
// Needs `invoicing:create`. `overrides` win over what the record maps to. 409 `invoicingCrmRecordTaken` when the record is a customer already.
//
// POST /api/invoicing/customers/from-crm — Token permission: write:invoicing.
func (r *InvoicingCustomersService) CreateFromCRM(ctx context.Context, body *InvoicingCustomersCreateFromCrmBody, query *InvoicingCustomersCreateFromCrmQuery, opts ...CallOption) (*InvoicingCustomer, error) {
	var out *InvoicingCustomer
	err := r.c.call(ctx, "POST", "/api/invoicing/customers/from-crm", query, body, false, "customer", &out, opts)
	return out, err
}

// Delete — Deletes a customer no document names.
//
// 409 `invoicingHasDocuments` otherwise — archive them instead.
//
// DELETE /api/invoicing/customers/{customerId} — Token permission: write:invoicing.
//
// customerID: The customer.
func (r *InvoicingCustomersService) Delete(ctx context.Context, customerID string, query *InvoicingCustomersDeleteQuery, opts ...CallOption) (*InvoicingCustomersDeleteResponse, error) {
	var out *InvoicingCustomersDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/invoicing/customers/"+pathSeg(customerID), query, nil, false, "", &out, opts)
	return out, err
}

// Get — One customer, their latest 20 documents and what they owe per currency.
//
// GET /api/invoicing/customers/{customerId} — Token permission: read:invoicing.
//
// customerID: The customer.
func (r *InvoicingCustomersService) Get(ctx context.Context, customerID string, query *InvoicingCustomersGetQuery, opts ...CallOption) (*InvoicingCustomer, error) {
	var out *InvoicingCustomer
	err := r.c.call(ctx, "GET", "/api/invoicing/customers/"+pathSeg(customerID), query, nil, false, "customer", &out, opts)
	return out, err
}

// Import — Imports customers from a CSV file — a dry run first, then the write.
//
// Needs `invoicing:create`. `options` (JSON): encoding, delimiter, hasHeader, `mapping` (a customer field or null per file column), numberFormat (comma | dot), `duplicate` (match by customerNumber or email, then skip or update; or null), `dryRun`. VAT IDs are stored unchecked. An .xlsx is refused: save the sheet as CSV.
//
// POST /api/invoicing/customers/import — Token permission: write:invoicing.
func (r *InvoicingCustomersService) Import(ctx context.Context, body *InvoicingCustomersImportBody, query *InvoicingCustomersImportQuery, opts ...CallOption) (*InvoicingCustomersImportResponse, error) {
	var out *InvoicingCustomersImportResponse
	err := r.c.call(ctx, "POST", "/api/invoicing/customers/import", query, body, true, "", &out, opts)
	return out, err
}

// List — Lists the customers the caller may see, by name.
//
// GET /api/invoicing/customers — Token permission: read:invoicing.
func (r *InvoicingCustomersService) List(ctx context.Context, query *InvoicingCustomersListQuery, opts ...CallOption) ([]InvoicingCustomer, error) {
	var out []InvoicingCustomer
	err := r.c.call(ctx, "GET", "/api/invoicing/customers", query, nil, false, "customers", &out, opts)
	return out, err
}

// PreviewFromCRM — What a CRM record would become as a customer, and which column fed which field.
//
// The record is read under the CRM's own rules. `existing` is the customer already linked to it.
//
// GET /api/invoicing/customers/from-crm — Token permission: read:invoicing.
func (r *InvoicingCustomersService) PreviewFromCRM(ctx context.Context, query *InvoicingCustomersPreviewFromCrmQuery, opts ...CallOption) (*InvoicingCustomersPreviewFromCrmResponse, error) {
	var out *InvoicingCustomersPreviewFromCrmResponse
	err := r.c.call(ctx, "GET", "/api/invoicing/customers/from-crm", query, nil, false, "", &out, opts)
	return out, err
}

// Update — Changes a customer, or archives and restores them.
//
// Needs `invoicing:create`. A changed VAT ID is checked again. Issued documents keep the customer as they were. Who sees them is changed by their creator or `invoicing:settings` only.
//
// PATCH /api/invoicing/customers/{customerId} — Token permission: write:invoicing.
//
// customerID: The customer.
func (r *InvoicingCustomersService) Update(ctx context.Context, customerID string, body *InvoicingCustomersUpdateBody, query *InvoicingCustomersUpdateQuery, opts ...CallOption) (*InvoicingCustomer, error) {
	var out *InvoicingCustomer
	err := r.c.call(ctx, "PATCH", "/api/invoicing/customers/"+pathSeg(customerID), query, body, false, "customer", &out, opts)
	return out, err
}

// InvoicingDocumentsService — invoicing.documents
type InvoicingDocumentsService struct {
	c *Client
}

func newInvoicingDocumentsService(c *Client) *InvoicingDocumentsService {
	return &InvoicingDocumentsService{c: c}
}

// AnswerQuote — Records the customer's answer to an issued quote: accepted or declined.
//
// Needs `invoicing:finalize`. Only an issued quote that is still open.
//
// POST /api/invoicing/documents/{documentId}/quote — Token permission: write:invoicing.
//
// documentID: The document (quote, invoice, credit note…).
func (r *InvoicingDocumentsService) AnswerQuote(ctx context.Context, documentID string, body *InvoicingDocumentsAnswerQuoteBody, query *InvoicingDocumentsAnswerQuoteQuery, opts ...CallOption) (*InvoicingDocument, error) {
	var out *InvoicingDocument
	err := r.c.call(ctx, "POST", "/api/invoicing/documents/"+pathSeg(documentID)+"/quote", query, body, false, "document", &out, opts)
	return out, err
}

// BankLines — The bank lines an invoice was paid with, as the Banking hub matched them.
//
// Date and amount always; account, counterpart and purpose only where the caller may read that account's lines.
//
// GET /api/invoicing/documents/{documentId}/bank — Token permission: read:invoicing.
//
// documentID: The document (quote, invoice, credit note…).
func (r *InvoicingDocumentsService) BankLines(ctx context.Context, documentID string, query *InvoicingDocumentsBankLinesQuery, opts ...CallOption) (*InvoicingDocumentsBankLinesResponse, error) {
	var out *InvoicingDocumentsBankLinesResponse
	err := r.c.call(ctx, "GET", "/api/invoicing/documents/"+pathSeg(documentID)+"/bank", query, nil, false, "", &out, opts)
	return out, err
}

// CheckEinvoice — Checks a draft's e-invoice with the KoSIT validator, written with the number it would get — nothing is drawn.
//
// Needs `invoicing:create` and e-invoicing on the plan. The verdict is written to the document's history. `report` is null when the draft is not an e-invoice.
//
// POST /api/invoicing/documents/{documentId}/einvoice-check — Token permission: write:invoicing.
//
// documentID: The document (quote, invoice, credit note…).
func (r *InvoicingDocumentsService) CheckEinvoice(ctx context.Context, documentID string, query *InvoicingDocumentsCheckEinvoiceQuery, opts ...CallOption) (*InvoicingDocumentsCheckEinvoiceResponse, error) {
	var out *InvoicingDocumentsCheckEinvoiceResponse
	err := r.c.call(ctx, "POST", "/api/invoicing/documents/"+pathSeg(documentID)+"/einvoice-check", query, nil, false, "", &out, opts)
	return out, err
}

// Convert — Makes a new draft from an issued document: an order confirmation or invoice from a quote, a partial or final invoice, a credit note or cancellation of an invoice.
//
// Needs `invoicing:create` (a credit note or cancellation: `invoicing:cancel`).
//
// - quote → order_confirmation | invoice
// - order_confirmation → invoice | delivery_note
// - quote / order confirmation → invoice with `invoiceKind`: `partial` (`percent` of the gross or `amountCents`) or `final` (every line, the issued partial invoices deducted)
// - invoice → credit_note (all lines, chosen `lines` with smaller quantities, or a gross `amountCents`) | cancellation
//
// 409 `invoicingIsDraft` (finalize the source first), 409 `invoicingInvalidTransition`, 422 `invoicingOverCredit` / `invoicingPartialExceeds`.
//
// POST /api/invoicing/documents/{documentId}/convert — Token permission: write:invoicing.
//
// documentID: The document (quote, invoice, credit note…).
func (r *InvoicingDocumentsService) Convert(ctx context.Context, documentID string, body *InvoicingDocumentsConvertBody, query *InvoicingDocumentsConvertQuery, opts ...CallOption) (*InvoicingDocument, error) {
	var out *InvoicingDocument
	err := r.c.call(ctx, "POST", "/api/invoicing/documents/"+pathSeg(documentID)+"/convert", query, body, false, "document", &out, opts)
	return out, err
}

// Create — Creates a draft — for a customer, from catalogue products and lines.
//
// Needs `invoicing:create`. The customer gives the language, currency, terms and discount; the template the texts. An invoice without a service date gets the document date. Credit notes and cancellations are made from their invoice (convert).
//
// POST /api/invoicing/documents — Token permission: write:invoicing.
func (r *InvoicingDocumentsService) Create(ctx context.Context, body *InvoicingDraftInput, query *InvoicingDocumentsCreateQuery, opts ...CallOption) (*InvoicingDocument, error) {
	var out *InvoicingDocument
	err := r.c.call(ctx, "POST", "/api/invoicing/documents", query, body, false, "document", &out, opts)
	return out, err
}

// Delete — Moves a draft to the trash.
//
// Needs `invoicing:create`. A finalized document is never deleted (409 `invoicingFinalized`, §147 AO).
//
// DELETE /api/invoicing/documents/{documentId} — Token permission: write:invoicing.
//
// documentID: The document (quote, invoice, credit note…).
func (r *InvoicingDocumentsService) Delete(ctx context.Context, documentID string, query *InvoicingDocumentsDeleteQuery, opts ...CallOption) (*InvoicingDocumentsDeleteResponse, error) {
	var out *InvoicingDocumentsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/invoicing/documents/"+pathSeg(documentID), query, nil, false, "", &out, opts)
	return out, err
}

// Duplicate — Makes a new draft of the same type with the same customer, lines and texts.
//
// Needs `invoicing:create`. No number, payments or links are copied. A credit note or cancellation is refused: it is made from its invoice.
//
// POST /api/invoicing/documents/{documentId}/duplicate — Token permission: write:invoicing.
//
// documentID: The document (quote, invoice, credit note…).
func (r *InvoicingDocumentsService) Duplicate(ctx context.Context, documentID string, query *InvoicingDocumentsDuplicateQuery, opts ...CallOption) (*InvoicingDocument, error) {
	var out *InvoicingDocument
	err := r.c.call(ctx, "POST", "/api/invoicing/documents/"+pathSeg(documentID)+"/duplicate", query, nil, false, "document", &out, opts)
	return out, err
}

// Export — Downloads documents as a CSV file, filtered like the list — or, with `kind=dunning`, the dunning notices sent.
//
// Needs `invoicing:export`. Every document of the organisation, whoever may see it in the hub: the books are always complete. At most 10,000 rows.
//
// GET /api/invoicing/documents/export — Token permission: read:invoicing.
func (r *InvoicingDocumentsService) Export(ctx context.Context, query *InvoicingDocumentsExportQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/invoicing/documents/export", query, nil, false, "", &out, opts)
	return out, err
}

// Finalize — Issues a draft: its number is drawn, it is frozen and its PDF archived.
//
// Needs `invoicing:finalize` (a credit note or cancellation: `invoicing:cancel`). From then on the document can never be changed or deleted. Refusals: 422 `invoicingIncomplete` with `issues` (what §14 UStG still misses), 422 `invoicingOverCredit`, 409 `invoicingNotDraft` / `invoicingAlreadyCancelled` / `invoicingNumberTaken`. An e-invoice needs e-invoicing on the plan and passes the KoSIT validator (422 `invoicingEinvoiceInvalid` with `report`; the draft keeps no number). The plan's monthly document limit counts every type but credit notes and cancellations.
//
// POST /api/invoicing/documents/{documentId}/finalize — Token permission: write:invoicing.
//
// documentID: The document (quote, invoice, credit note…).
func (r *InvoicingDocumentsService) Finalize(ctx context.Context, documentID string, query *InvoicingDocumentsFinalizeQuery, opts ...CallOption) (*InvoicingDocument, error) {
	var out *InvoicingDocument
	err := r.c.call(ctx, "POST", "/api/invoicing/documents/"+pathSeg(documentID)+"/finalize", query, nil, false, "document", &out, opts)
	return out, err
}

// Get — One document with its lines, history, payments, the documents of its deal and its e-invoice.
//
// GET /api/invoicing/documents/{documentId} — Token permission: read:invoicing.
//
// documentID: The document (quote, invoice, credit note…).
func (r *InvoicingDocumentsService) Get(ctx context.Context, documentID string, query *InvoicingDocumentsGetQuery, opts ...CallOption) (*InvoicingDocumentsGetResponse, error) {
	var out *InvoicingDocumentsGetResponse
	err := r.c.call(ctx, "GET", "/api/invoicing/documents/"+pathSeg(documentID), query, nil, false, "", &out, opts)
	return out, err
}

// List — Lists one page of the documents the caller may see, newest first with drafts on top, with totals per currency.
//
// Needs `invoicing:read`.
//
// GET /api/invoicing/documents — Token permission: read:invoicing.
func (r *InvoicingDocumentsService) List(ctx context.Context, query *InvoicingDocumentsListQuery, opts ...CallOption) ([]InvoicingDocument, error) {
	var out []InvoicingDocument
	err := r.c.call(ctx, "GET", "/api/invoicing/documents", query, nil, false, "documents", &out, opts)
	return out, err
}

// MarkSent — Records that an issued document went out another way: handed over, printed, mailed from elsewhere.
//
// Needs `invoicing:finalize`.
//
// POST /api/invoicing/documents/{documentId}/sent — Token permission: write:invoicing.
//
// documentID: The document (quote, invoice, credit note…).
func (r *InvoicingDocumentsService) MarkSent(ctx context.Context, documentID string, body *InvoicingDocumentsMarkSentBody, query *InvoicingDocumentsMarkSentQuery, opts ...CallOption) (*InvoicingDocument, error) {
	var out *InvoicingDocument
	err := r.c.call(ctx, "POST", "/api/invoicing/documents/"+pathSeg(documentID)+"/sent", query, body, false, "document", &out, opts)
	return out, err
}

// PDF — Downloads a document's PDF: an issued one as archived, a draft rendered now and stamped as a draft.
//
// The header `X-Content-SHA256` carries an issued document's stored hash. 502 `invoicingPdfFailed` when the renderer does not answer.
//
// GET /api/invoicing/documents/{documentId}/pdf — Token permission: read:invoicing.
//
// documentID: The document (quote, invoice, credit note…).
func (r *InvoicingDocumentsService) PDF(ctx context.Context, documentID string, query *InvoicingDocumentsPdfQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/invoicing/documents/"+pathSeg(documentID)+"/pdf", query, nil, false, "", &out, opts)
	return out, err
}

// Send — Sends an issued document by e-mail with its PDF (and an XRechnung's XML).
//
// Needs `invoicing:finalize`. From the Mail hub mailbox `accountId` (the Mail hub's rules apply), else from the platform's sender with Reply-To the issuer. Records the sending. 409 `invoicingIsDraft`, 422 `invoicingNoRecipient`, 502 `invoicingPdfFailed` / `invoicingSendFailed`.
//
// POST /api/invoicing/documents/{documentId}/send — Token permission: write:invoicing.
//
// documentID: The document (quote, invoice, credit note…).
func (r *InvoicingDocumentsService) Send(ctx context.Context, documentID string, body *InvoicingDocumentsSendBody, query *InvoicingDocumentsSendQuery, opts ...CallOption) (*InvoicingDocumentsSendResponse, error) {
	var out *InvoicingDocumentsSendResponse
	err := r.c.call(ctx, "POST", "/api/invoicing/documents/"+pathSeg(documentID)+"/send", query, body, false, "", &out, opts)
	return out, err
}

// SendDefaults — What sending a document by e-mail starts from: the recipients, the subject, and the mailboxes the caller may send from.
//
// Needs `invoicing:finalize`. No mailboxes: the platform's sender is used.
//
// GET /api/invoicing/documents/{documentId}/send — Token permission: read:invoicing.
//
// documentID: The document (quote, invoice, credit note…).
func (r *InvoicingDocumentsService) SendDefaults(ctx context.Context, documentID string, query *InvoicingDocumentsSendDefaultsQuery, opts ...CallOption) (*InvoicingDocumentsSendDefaultsResponse, error) {
	var out *InvoicingDocumentsSendDefaultsResponse
	err := r.c.call(ctx, "GET", "/api/invoicing/documents/"+pathSeg(documentID)+"/send", query, nil, false, "", &out, opts)
	return out, err
}

// Update — Changes a draft; totals, tax and legal notes are computed again.
//
// Needs `invoicing:create`. A finalized document is refused (409 `invoicingNotDraft`) — except `visibility` / `workspaceId` alone, which change at any status (its creator or `invoicing:settings`).
//
// PATCH /api/invoicing/documents/{documentId} — Token permission: write:invoicing.
//
// documentID: The document (quote, invoice, credit note…).
func (r *InvoicingDocumentsService) Update(ctx context.Context, documentID string, body *InvoicingDraftPatch, query *InvoicingDocumentsUpdateQuery, opts ...CallOption) (*InvoicingDocument, error) {
	var out *InvoicingDocument
	err := r.c.call(ctx, "PATCH", "/api/invoicing/documents/"+pathSeg(documentID), query, body, false, "document", &out, opts)
	return out, err
}

// Xml — Downloads an issued e-invoice's XML exactly as archived at finalize.
//
// XRechnung UBL / CII, or the factur-x.xml of a ZUGFeRD PDF; checked against its stored hash (`X-Content-SHA256`). 404 `invoicingEinvoiceNotIssued` for a draft or a document without an e-invoice.
//
// GET /api/invoicing/documents/{documentId}/xml — Token permission: read:invoicing.
//
// documentID: The document (quote, invoice, credit note…).
func (r *InvoicingDocumentsService) Xml(ctx context.Context, documentID string, query *InvoicingDocumentsXmlQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/invoicing/documents/"+pathSeg(documentID)+"/xml", query, nil, false, "", &out, opts)
	return out, err
}

// InvoicingDunningService — invoicing.dunning
type InvoicingDunningService struct {
	c *Client
}

func newInvoicingDunningService(c *Client) *InvoicingDunningService {
	return &InvoicingDunningService{c: c}
}

// GetNotice — One dunning notice.
//
// GET /api/invoicing/dunning/notices/{noticeId} — Token permission: read:invoicing.
//
// noticeID: The dunning notice.
func (r *InvoicingDunningService) GetNotice(ctx context.Context, noticeID string, query *InvoicingDunningGetNoticeQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/dunning/notices/"+pathSeg(noticeID), query, nil, false, "notice", &out, opts)
	return out, err
}

// List — The proposed dunning notices waiting for review, oldest first, with their invoices.
//
// Needs the plan's dunning feature.
//
// GET /api/invoicing/dunning — Token permission: read:invoicing.
func (r *InvoicingDunningService) List(ctx context.Context, query *InvoicingDunningListQuery, opts ...CallOption) ([]InvoicingDunningListResponseItemsItem, error) {
	var out []InvoicingDunningListResponseItemsItem
	err := r.c.call(ctx, "GET", "/api/invoicing/dunning", query, nil, false, "items", &out, opts)
	return out, err
}

// NoticePDF — Downloads a dunning notice as PDF: a sent one as archived, a proposed one as a preview.
//
// GET /api/invoicing/dunning/notices/{noticeId}/pdf — Token permission: read:invoicing.
//
// noticeID: The dunning notice.
func (r *InvoicingDunningService) NoticePDF(ctx context.Context, noticeID string, query *InvoicingDunningNoticePdfQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/invoicing/dunning/notices/"+pathSeg(noticeID)+"/pdf", query, nil, false, "", &out, opts)
	return out, err
}

// PauseCustomer — Excludes a customer from dunning, or includes them again.
//
// Needs `invoicing:payments` and the plan's dunning feature. While paused no notice is proposed for their invoices and waiting proposals are dropped.
//
// PATCH /api/invoicing/customers/{customerId}/dunning — Token permission: write:invoicing.
//
// customerID: The customer.
func (r *InvoicingDunningService) PauseCustomer(ctx context.Context, customerID string, body *InvoicingDunningPauseCustomerBody, query *InvoicingDunningPauseCustomerQuery, opts ...CallOption) (bool, error) {
	var out bool
	err := r.c.call(ctx, "PATCH", "/api/invoicing/customers/"+pathSeg(customerID)+"/dunning", query, body, false, "paused", &out, opts)
	return out, err
}

// PauseDocument — Pauses or resumes the dunning of an invoice; pausing drops a waiting proposal.
//
// Needs `invoicing:payments` and the plan's dunning feature.
//
// PATCH /api/invoicing/documents/{documentId}/dunning — Token permission: write:invoicing.
//
// documentID: The document (quote, invoice, credit note…).
func (r *InvoicingDunningService) PauseDocument(ctx context.Context, documentID string, body *InvoicingDunningPauseDocumentBody, query *InvoicingDunningPauseDocumentQuery, opts ...CallOption) (*InvoicingDocument, error) {
	var out *InvoicingDocument
	err := r.c.call(ctx, "PATCH", "/api/invoicing/documents/"+pathSeg(documentID)+"/dunning", query, body, false, "document", &out, opts)
	return out, err
}

// Propose — Proposes a dunning notice for an invoice now: the next level, or a chosen one above every level sent.
//
// Needs `invoicing:payments` and the plan's dunning feature. 409 `invoicingDunningNotDunnable` with `reason` (paid, cancelled, not due, paused, last level sent); 409 `invoicingDunningWaiting` when one waits for review already.
//
// POST /api/invoicing/documents/{documentId}/dunning — Token permission: write:invoicing.
//
// documentID: The document (quote, invoice, credit note…).
func (r *InvoicingDunningService) Propose(ctx context.Context, documentID string, body *InvoicingDunningProposeBody, query *InvoicingDunningProposeQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/documents/"+pathSeg(documentID)+"/dunning", query, body, false, "notice", &out, opts)
	return out, err
}

// RemoveNotice — Removes a proposed dunning notice; a sent one is a record and stays.
//
// Needs `invoicing:payments` and the plan's dunning feature. 409 `invoicingDunningNoticeSent`.
//
// DELETE /api/invoicing/dunning/notices/{noticeId} — Token permission: write:invoicing.
//
// noticeID: The dunning notice.
func (r *InvoicingDunningService) RemoveNotice(ctx context.Context, noticeID string, query *InvoicingDunningRemoveNoticeQuery, opts ...CallOption) (*InvoicingDunningRemoveNoticeResponse, error) {
	var out *InvoicingDunningRemoveNoticeResponse
	err := r.c.call(ctx, "DELETE", "/api/invoicing/dunning/notices/"+pathSeg(noticeID), query, nil, false, "", &out, opts)
	return out, err
}

// SendDefaults — What sending a dunning notice starts from: the recipients, the subject, the mailboxes the caller may send from.
//
// GET /api/invoicing/dunning/notices/{noticeId}/send — Token permission: read:invoicing.
//
// noticeID: The dunning notice.
func (r *InvoicingDunningService) SendDefaults(ctx context.Context, noticeID string, query *InvoicingDunningSendDefaultsQuery, opts ...CallOption) (*InvoicingDunningSendDefaultsResponse, error) {
	var out *InvoicingDunningSendDefaultsResponse
	err := r.c.call(ctx, "GET", "/api/invoicing/dunning/notices/"+pathSeg(noticeID)+"/send", query, nil, false, "", &out, opts)
	return out, err
}

// SendNotice — Sends a proposed dunning notice by e-mail, or records that it went out on paper.
//
// Needs `invoicing:payments` and the plan's dunning feature. Its amounts are brought to today and its PDF archived; then it never changes.
//
// POST /api/invoicing/dunning/notices/{noticeId}/send — Token permission: write:invoicing.
//
// noticeID: The dunning notice.
func (r *InvoicingDunningService) SendNotice(ctx context.Context, noticeID string, body *InvoicingDunningSendNoticeBody, query *InvoicingDunningSendNoticeQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/dunning/notices/"+pathSeg(noticeID)+"/send", query, body, false, "notice", &out, opts)
	return out, err
}

// State — An invoice's dunning: the levels, the notices so far, the next level and when it is due, and what a notice would claim today.
//
// `state` is null for anything but an issued invoice.
//
// GET /api/invoicing/documents/{documentId}/dunning — Token permission: read:invoicing.
//
// documentID: The document (quote, invoice, credit note…).
func (r *InvoicingDunningService) State(ctx context.Context, documentID string, query *InvoicingDunningStateQuery, opts ...CallOption) (*InvoicingDunningStateResponse, error) {
	var out *InvoicingDunningStateResponse
	err := r.c.call(ctx, "GET", "/api/invoicing/documents/"+pathSeg(documentID)+"/dunning", query, nil, false, "", &out, opts)
	return out, err
}

// InvoicingInboundService — invoicing.inbound
type InvoicingInboundService struct {
	c *Client

	Mailboxes *InvoicingInboundMailboxesService
}

func newInvoicingInboundService(c *Client) *InvoicingInboundService {
	return &InvoicingInboundService{c: c, Mailboxes: newInvoicingInboundMailboxesService(c)}
}

// Delete — Deletes a rejected received document, or one that could not be read as an e-invoice, with its original.
//
// 409 `invoicingInboundNotDeletable` otherwise: an approved or booked one is kept. Its uploader or `invoicing:payments`.
//
// DELETE /api/invoicing/inbound/{inboundId} — Token permission: write:invoicing.
//
// inboundID: The received document.
func (r *InvoicingInboundService) Delete(ctx context.Context, inboundID string, query *InvoicingInboundDeleteQuery, opts ...CallOption) (*InvoicingInboundDeleteResponse, error) {
	var out *InvoicingInboundDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/invoicing/inbound/"+pathSeg(inboundID), query, nil, false, "", &out, opts)
	return out, err
}

// Get — One received e-invoice with its EN 16931 data, validation report, history and duplicates.
//
// GET /api/invoicing/inbound/{inboundId} — Token permission: read:invoicing.
//
// inboundID: The received document.
func (r *InvoicingInboundService) Get(ctx context.Context, inboundID string, query *InvoicingInboundGetQuery, opts ...CallOption) (*InvoicingInboundGetResponse, error) {
	var out *InvoicingInboundGetResponse
	err := r.c.call(ctx, "GET", "/api/invoicing/inbound/"+pathSeg(inboundID), query, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the received e-invoices the caller may see, newest first, with totals per currency.
//
// Needs receiving invoices on the plan.
//
// GET /api/invoicing/inbound — Token permission: read:invoicing.
func (r *InvoicingInboundService) List(ctx context.Context, query *InvoicingInboundListQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/inbound", query, nil, false, "documents", &out, opts)
	return out, err
}

// Original — Downloads a received document exactly as it arrived.
//
// Checked against the hash stored on arrival (`X-Content-SHA256`). XML always as an attachment; a PDF inline unless `download=1`.
//
// GET /api/invoicing/inbound/{inboundId}/original — Token permission: read:invoicing.
//
// inboundID: The received document.
func (r *InvoicingInboundService) Original(ctx context.Context, inboundID string, query *InvoicingInboundOriginalQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/invoicing/inbound/"+pathSeg(inboundID)+"/original", query, nil, false, "", &out, opts)
	return out, err
}

// Revalidate — Validates a received e-invoice's stored original again with the KoSIT validator.
//
// Needs `invoicing:create`. For when the validator was not reachable on arrival or its rules were updated; only the XML read out of the original is validated. 503 `invoicingInboundValidatorUnavailable` (with the unchanged `document`) when it still does not answer — the earlier verdict stays. 409 `invoicingInboundNotValidatable` for a file without XML. 10 a minute.
//
// POST /api/invoicing/inbound/{inboundId}/revalidate — Token permission: write:invoicing.
//
// inboundID: The received document.
func (r *InvoicingInboundService) Revalidate(ctx context.Context, inboundID string, query *InvoicingInboundRevalidateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/inbound/"+pathSeg(inboundID)+"/revalidate", query, nil, false, "document", &out, opts)
	return out, err
}

// Transition — Checks, approves, books, rejects or reopens a received e-invoice.
//
// Needs `invoicing:payments`. check: received → checked; approve: → approved; book: approved → booked; reject (with a `note`): → rejected; reopen: rejected → received. Anything else: 409 `invoicingInboundInvalidTransition`.
//
// POST /api/invoicing/inbound/{inboundId}/status — Token permission: write:invoicing.
//
// inboundID: The received document.
func (r *InvoicingInboundService) Transition(ctx context.Context, inboundID string, body *InvoicingInboundTransitionBody, query *InvoicingInboundTransitionQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/inbound/"+pathSeg(inboundID)+"/status", query, body, false, "document", &out, opts)
	return out, err
}

// Update — Changes who sees a received e-invoice.
//
// Its uploader or `invoicing:settings`.
//
// PATCH /api/invoicing/inbound/{inboundId} — Token permission: write:invoicing.
//
// inboundID: The received document.
func (r *InvoicingInboundService) Update(ctx context.Context, inboundID string, body *InvoicingInboundUpdateBody, query *InvoicingInboundUpdateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "PATCH", "/api/invoicing/inbound/"+pathSeg(inboundID), query, body, false, "document", &out, opts)
	return out, err
}

// Upload — Receives supplier e-invoices: XRechnung, ZUGFeRD / Factur-X or a plain PDF.
//
// Needs `invoicing:create` and receiving invoices on the plan. Multipart with one or more `files` (or `file`), optional `visibility` and `workspaceId`; an API client may also send the raw file as the body (application/xml or application/pdf) with its name in `X-Filename` or `?filename=`. At most 20 files of 20 MB. Every file is stored unchangeable with its SHA-256, read, validated and flagged when it repeats an earlier one; nothing is refused for its content. 422 `invoicingInboundNotStored` when none was stored. 30 a minute.
//
// POST /api/invoicing/inbound — Token permission: write:invoicing.
func (r *InvoicingInboundService) Upload(ctx context.Context, body *InvoicingInboundUploadBody, query *InvoicingInboundUploadQuery, opts ...CallOption) ([]InvoicingInboundUploadResponseResultsItem, error) {
	var out []InvoicingInboundUploadResponseResultsItem
	err := r.c.call(ctx, "POST", "/api/invoicing/inbound", query, body, true, "results", &out, opts)
	return out, err
}

// InvoicingInboundMailboxesService — invoicing.inbound.mailboxes
type InvoicingInboundMailboxesService struct {
	c *Client
}

func newInvoicingInboundMailboxesService(c *Client) *InvoicingInboundMailboxesService {
	return &InvoicingInboundMailboxesService{c: c}
}

// Add — Receives supplier invoices from a Mail hub mailbox, automatically or by hand.
//
// Needs `invoicing:settings`, signed in as somebody who may open the mailbox in the Mail hub (an API token never reads mail). `backfillDays` (automatic only): a one-time import of the last N days.
//
// POST /api/invoicing/inbound/mailboxes — Token permission: write:invoicing.
func (r *InvoicingInboundMailboxesService) Add(ctx context.Context, body *InvoicingInboundMailboxesAddBody, query *InvoicingInboundMailboxesAddQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/inbound/mailboxes", query, body, false, "mailbox", &out, opts)
	return out, err
}

// List — The Mail hub mailboxes the organisation receives supplier invoices from, and what the newest mails came to.
//
// GET /api/invoicing/inbound/mailboxes — Token permission: read:invoicing.
func (r *InvoicingInboundMailboxesService) List(ctx context.Context, query *InvoicingInboundMailboxesListQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/inbound/mailboxes", query, nil, false, "mailboxes", &out, opts)
	return out, err
}

// Remove — Stops receiving from a mailbox; documents already taken stay.
//
// DELETE /api/invoicing/inbound/mailboxes/{mailboxId} — Token permission: write:invoicing.
//
// mailboxID: The receiving mailbox.
func (r *InvoicingInboundMailboxesService) Remove(ctx context.Context, mailboxID string, query *InvoicingInboundMailboxesRemoveQuery, opts ...CallOption) (*InvoicingInboundMailboxesRemoveResponse, error) {
	var out *InvoicingInboundMailboxesRemoveResponse
	err := r.c.call(ctx, "DELETE", "/api/invoicing/inbound/mailboxes/"+pathSeg(mailboxID), query, nil, false, "", &out, opts)
	return out, err
}

// Update — Switches a receiving mailbox between automatic and manual.
//
// Needs `invoicing:settings`. To automatic only for somebody signed in who may open the mailbox.
//
// PATCH /api/invoicing/inbound/mailboxes/{mailboxId} — Token permission: write:invoicing.
//
// mailboxID: The receiving mailbox.
func (r *InvoicingInboundMailboxesService) Update(ctx context.Context, mailboxID string, body *InvoicingInboundMailboxesUpdateBody, query *InvoicingInboundMailboxesUpdateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "PATCH", "/api/invoicing/inbound/mailboxes/"+pathSeg(mailboxID), query, body, false, "mailbox", &out, opts)
	return out, err
}

// InvoicingIssuersService — invoicing.issuers
type InvoicingIssuersService struct {
	c *Client
}

func newInvoicingIssuersService(c *Client) *InvoicingIssuersService {
	return &InvoicingIssuersService{c: c}
}

// Create — Creates an issuing company with a number range for every document type.
//
// Needs `invoicing:settings`. A second active company needs multiple companies on the plan and counts against its limit.
//
// POST /api/invoicing/issuers — Token permission: write:invoicing.
func (r *InvoicingIssuersService) Create(ctx context.Context, body *InvoicingIssuerInput, query *InvoicingIssuersCreateQuery, opts ...CallOption) (*InvoicingIssuer, error) {
	var out *InvoicingIssuer
	err := r.c.call(ctx, "POST", "/api/invoicing/issuers", query, body, false, "issuer", &out, opts)
	return out, err
}

// Delete — Deletes an issuing company that never issued anything.
//
// Needs `invoicing:settings`. 409 `invoicingHasDocuments` otherwise — archive it instead.
//
// DELETE /api/invoicing/issuers/{issuerId} — Token permission: write:invoicing.
//
// issuerID: The issuing company.
func (r *InvoicingIssuersService) Delete(ctx context.Context, issuerID string, query *InvoicingIssuersDeleteQuery, opts ...CallOption) (*InvoicingIssuersDeleteResponse, error) {
	var out *InvoicingIssuersDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/invoicing/issuers/"+pathSeg(issuerID), query, nil, false, "", &out, opts)
	return out, err
}

// Get — One issuing company with its number ranges.
//
// GET /api/invoicing/issuers/{issuerId} — Token permission: read:invoicing.
//
// issuerID: The issuing company.
func (r *InvoicingIssuersService) Get(ctx context.Context, issuerID string, query *InvoicingIssuersGetQuery, opts ...CallOption) (*InvoicingIssuer, error) {
	var out *InvoicingIssuer
	err := r.c.call(ctx, "GET", "/api/invoicing/issuers/"+pathSeg(issuerID), query, nil, false, "issuer", &out, opts)
	return out, err
}

// GetDunningSettings — A company's dunning settings: the levels, their built-in texts, and the base rate for default interest.
//
// GET /api/invoicing/issuers/{issuerId}/dunning — Token permission: read:invoicing.
//
// issuerID: The issuing company.
func (r *InvoicingIssuersService) GetDunningSettings(ctx context.Context, issuerID string, query *InvoicingIssuersGetDunningSettingsQuery, opts ...CallOption) (*InvoicingIssuersGetDunningSettingsResponse, error) {
	var out *InvoicingIssuersGetDunningSettingsResponse
	err := r.c.call(ctx, "GET", "/api/invoicing/issuers/"+pathSeg(issuerID)+"/dunning", query, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the organisation's issuing companies, the default first.
//
// GET /api/invoicing/issuers — Token permission: read:invoicing.
func (r *InvoicingIssuersService) List(ctx context.Context, query *InvoicingIssuersListQuery, opts ...CallOption) ([]InvoicingIssuer, error) {
	var out []InvoicingIssuer
	err := r.c.call(ctx, "GET", "/api/invoicing/issuers", query, nil, false, "issuers", &out, opts)
	return out, err
}

// SetDunningSettings — Replaces a company's dunning settings.
//
// Needs `invoicing:settings` and the plan's dunning feature. Checked strictly: 400 `invoicingDunningSettingsInvalid` with `field`, 400 `invoicingUnknownMergeField`.
//
// PUT /api/invoicing/issuers/{issuerId}/dunning — Token permission: write:invoicing.
//
// issuerID: The issuing company.
func (r *InvoicingIssuersService) SetDunningSettings(ctx context.Context, issuerID string, body *InvoicingIssuersSetDunningSettingsBody, query *InvoicingIssuersSetDunningSettingsQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "PUT", "/api/invoicing/issuers/"+pathSeg(issuerID)+"/dunning", query, body, false, "settings", &out, opts)
	return out, err
}

// Update — Changes an issuing company, or archives and restores it.
//
// Needs `invoicing:settings`. Restoring counts as one more active company.
//
// PATCH /api/invoicing/issuers/{issuerId} — Token permission: write:invoicing.
//
// issuerID: The issuing company.
func (r *InvoicingIssuersService) Update(ctx context.Context, issuerID string, body *InvoicingIssuersUpdateBody, query *InvoicingIssuersUpdateQuery, opts ...CallOption) (*InvoicingIssuer, error) {
	var out *InvoicingIssuer
	err := r.c.call(ctx, "PATCH", "/api/invoicing/issuers/"+pathSeg(issuerID), query, body, false, "issuer", &out, opts)
	return out, err
}

// InvoicingLedgerService — invoicing.ledger
type InvoicingLedgerService struct {
	c *Client

	Accounts *InvoicingLedgerAccountsService
	Datev    *InvoicingLedgerDatevService
	Entries  *InvoicingLedgerEntriesService
	Periods  *InvoicingLedgerPeriodsService
	Settings *InvoicingLedgerSettingsService
	YearEnd  *InvoicingLedgerYearEndService
}

func newInvoicingLedgerService(c *Client) *InvoicingLedgerService {
	return &InvoicingLedgerService{c: c, Accounts: newInvoicingLedgerAccountsService(c), Datev: newInvoicingLedgerDatevService(c), Entries: newInvoicingLedgerEntriesService(c), Periods: newInvoicingLedgerPeriodsService(c), Settings: newInvoicingLedgerSettingsService(c), YearEnd: newInvoicingLedgerYearEndService(c)}
}

// Events — The books' own history: set up, settings changed, months closed and reopened, accounts changed, DATEV exports.
//
// GET /api/invoicing/ledger/events — Token permission: read:invoicing.
func (r *InvoicingLedgerService) Events(ctx context.Context, query *InvoicingLedgerEventsQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/ledger/events", query, nil, false, "events", &out, opts)
	return out, err
}

// Opening — Posts opening balances at the first booking date, each against the carried-forward account.
//
// Needs `invoicing:bookkeeping` and the ledger on the plan.
//
// POST /api/invoicing/ledger/opening — Token permission: write:invoicing.
func (r *InvoicingLedgerService) Opening(ctx context.Context, body *InvoicingLedgerOpeningBody, query *InvoicingLedgerOpeningQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/ledger/opening", query, body, false, "entry", &out, opts)
	return out, err
}

// Report — One report of the books: EÜR, VAT return (UStVA), trial balance, account sheet, P&L, balance sheet or open items.
//
// Needs `invoicing:export` or `invoicing:bookkeeping`. JSON with the figures and a ready table; `format=csv` or `pdf` downloads it. The books complete.
//
// GET /api/invoicing/ledger/reports/{kind} — Token permission: read:invoicing.
//
// kind: euer, ustva, trial_balance, account_sheet, pnl, balance_sheet or open_items.
func (r *InvoicingLedgerService) Report(ctx context.Context, kind string, query *InvoicingLedgerReportQuery, opts ...CallOption) (*InvoicingLedgerReportResponse, error) {
	var out *InvoicingLedgerReportResponse
	err := r.c.call(ctx, "GET", "/api/invoicing/ledger/reports/"+pathSeg(kind), query, nil, false, "", &out, opts)
	return out, err
}

// Start — Starts a company's books: posts its history from the first booking date, once.
//
// Needs `invoicing:bookkeeping`. Small books run at once; larger ones are queued (`status: queued`). A second run posts only what is missing.
//
// POST /api/invoicing/ledger/start — Token permission: write:invoicing.
func (r *InvoicingLedgerService) Start(ctx context.Context, body *InvoicingLedgerStartBody, query *InvoicingLedgerStartQuery, opts ...CallOption) (*InvoicingLedgerStartResponse, error) {
	var out *InvoicingLedgerStartResponse
	err := r.c.call(ctx, "POST", "/api/invoicing/ledger/start", query, body, false, "", &out, opts)
	return out, err
}

// InvoicingLedgerAccountsService — invoicing.ledger.accounts
type InvoicingLedgerAccountsService struct {
	c *Client
}

func newInvoicingLedgerAccountsService(c *Client) *InvoicingLedgerAccountsService {
	return &InvoicingLedgerAccountsService{c: c}
}

// Create — Adds a general-ledger account inside the chart's range for its kind.
//
// Needs `invoicing:bookkeeping` and the ledger on the plan. 422 `invoicingLedgerAccountOutOfRange`, 409 `invoicingLedgerAccountNumberTaken`.
//
// POST /api/invoicing/ledger/accounts — Token permission: write:invoicing.
func (r *InvoicingLedgerAccountsService) Create(ctx context.Context, body *InvoicingLedgerAccountsCreateBody, query *InvoicingLedgerAccountsCreateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/ledger/accounts", query, body, false, "account", &out, opts)
	return out, err
}

// List — A company's chart of accounts.
//
// Needs `invoicing:export` or `invoicing:bookkeeping`.
//
// GET /api/invoicing/ledger/accounts — Token permission: read:invoicing.
func (r *InvoicingLedgerAccountsService) List(ctx context.Context, query *InvoicingLedgerAccountsListQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/ledger/accounts", query, nil, false, "accounts", &out, opts)
	return out, err
}

// Update — Renames an account, (de)activates it, or changes its default tax code or EÜR line.
//
// Needs `invoicing:bookkeeping` and the ledger on the plan. System accounts stay active.
//
// PATCH /api/invoicing/ledger/accounts/{accountId} — Token permission: write:invoicing.
//
// accountID: The ledger account.
func (r *InvoicingLedgerAccountsService) Update(ctx context.Context, accountID string, body *InvoicingLedgerAccountsUpdateBody, query *InvoicingLedgerAccountsUpdateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "PATCH", "/api/invoicing/ledger/accounts/"+pathSeg(accountID), query, body, false, "account", &out, opts)
	return out, err
}

// InvoicingLedgerDatevService — invoicing.ledger.datev
type InvoicingLedgerDatevService struct {
	c *Client
}

func newInvoicingLedgerDatevService(c *Client) *InvoicingLedgerDatevService {
	return &InvoicingLedgerDatevService{c: c}
}

// Export — Downloads a DATEV export: the booking batch (EXTF), the customer and supplier master data, or a ZIP with both and the documents' PDFs.
//
// Needs `invoicing:export` and DATEV on the plan, and the adviser and client number in the books' settings. A range inside one fiscal year. Every export is recorded.
//
// POST /api/invoicing/ledger/datev — Token permission: write:invoicing.
func (r *InvoicingLedgerDatevService) Export(ctx context.Context, body *InvoicingLedgerDatevExportBody, query *InvoicingLedgerDatevExportQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "POST", "/api/invoicing/ledger/datev", query, body, false, "", &out, opts)
	return out, err
}

// List — The DATEV exports made so far.
//
// GET /api/invoicing/ledger/datev — Token permission: read:invoicing.
func (r *InvoicingLedgerDatevService) List(ctx context.Context, query *InvoicingLedgerDatevListQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/ledger/datev", query, nil, false, "exports", &out, opts)
	return out, err
}

// InvoicingLedgerEntriesService — invoicing.ledger.entries
type InvoicingLedgerEntriesService struct {
	c *Client
}

func newInvoicingLedgerEntriesService(c *Client) *InvoicingLedgerEntriesService {
	return &InvoicingLedgerEntriesService{c: c}
}

// List — The journal, newest first: every entry with its lines and what reversed it.
//
// Needs `invoicing:export` or `invoicing:bookkeeping`. The books complete.
//
// GET /api/invoicing/ledger/entries — Token permission: read:invoicing.
func (r *InvoicingLedgerEntriesService) List(ctx context.Context, query *InvoicingLedgerEntriesListQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/ledger/entries", query, nil, false, "entries", &out, opts)
	return out, err
}

// Post — Posts a manual journal entry: balanced, on the chart, in an open month.
//
// Needs `invoicing:bookkeeping` and the ledger on the plan. With `clientKey`, sending it again answers the entry already posted.
//
// POST /api/invoicing/ledger/entries — Token permission: write:invoicing.
func (r *InvoicingLedgerEntriesService) Post(ctx context.Context, body *InvoicingLedgerEntriesPostBody, query *InvoicingLedgerEntriesPostQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/ledger/entries", query, body, false, "entry", &out, opts)
	return out, err
}

// Reverse — Reverses a manual entry or opening balance (Storno): the same lines with the sides swapped, once.
//
// Needs `invoicing:bookkeeping`. An automatic entry follows its document (409 `invoicingLedgerAutoEntry`): correct the document instead.
//
// POST /api/invoicing/ledger/entries/{entryId}/reverse — Token permission: write:invoicing.
//
// entryID: The journal entry.
func (r *InvoicingLedgerEntriesService) Reverse(ctx context.Context, entryID string, body *InvoicingLedgerEntriesReverseBody, query *InvoicingLedgerEntriesReverseQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/ledger/entries/"+pathSeg(entryID)+"/reverse", query, body, false, "entry", &out, opts)
	return out, err
}

// InvoicingLedgerPeriodsService — invoicing.ledger.periods
type InvoicingLedgerPeriodsService struct {
	c *Client
}

func newInvoicingLedgerPeriodsService(c *Client) *InvoicingLedgerPeriodsService {
	return &InvoicingLedgerPeriodsService{c: c}
}

// List — The months of a fiscal year: closed or open, and exported to DATEV.
//
// GET /api/invoicing/ledger/periods — Token permission: read:invoicing.
func (r *InvoicingLedgerPeriodsService) List(ctx context.Context, query *InvoicingLedgerPeriodsListQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/ledger/periods", query, nil, false, "periods", &out, opts)
	return out, err
}

// Set — Closes a month (nothing can be posted in it any more) or reopens it.
//
// Needs `invoicing:bookkeeping`. Reopening a month a DATEV export covered asks first (409 `invoicingLedgerPeriodExported`) until `confirmExported`. Closing answers `warnings` for mapped bank accounts that do not reconcile.
//
// POST /api/invoicing/ledger/periods — Token permission: write:invoicing.
func (r *InvoicingLedgerPeriodsService) Set(ctx context.Context, body *InvoicingLedgerPeriodsSetBody, query *InvoicingLedgerPeriodsSetQuery, opts ...CallOption) (*InvoicingLedgerPeriodsSetResponse, error) {
	var out *InvoicingLedgerPeriodsSetResponse
	err := r.c.call(ctx, "POST", "/api/invoicing/ledger/periods", query, body, false, "", &out, opts)
	return out, err
}

// InvoicingLedgerSettingsService — invoicing.ledger.settings
type InvoicingLedgerSettingsService struct {
	c *Client
}

func newInvoicingLedgerSettingsService(c *Client) *InvoicingLedgerSettingsService {
	return &InvoicingLedgerSettingsService{c: c}
}

// Create — Switches a company's books on, with an SKR03 or SKR04 chart.
//
// Needs `invoicing:bookkeeping` and the ledger or EÜR on the plan. 409 `invoicingLedgerBooksAlreadyOn`.
//
// POST /api/invoicing/ledger/settings — Token permission: write:invoicing.
func (r *InvoicingLedgerSettingsService) Create(ctx context.Context, body *InvoicingLedgerSettingsCreateBody, query *InvoicingLedgerSettingsCreateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/ledger/settings", query, body, false, "settings", &out, opts)
	return out, err
}

// Get — A company's books: chart, method, VAT mode, fiscal year; null when it keeps none.
//
// GET /api/invoicing/ledger/settings — Token permission: read:invoicing.
func (r *InvoicingLedgerSettingsService) Get(ctx context.Context, query *InvoicingLedgerSettingsGetQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/ledger/settings", query, nil, false, "settings", &out, opts)
	return out, err
}

// Update — Changes a company's books; the chart, first booking date and fiscal year start only until the first posting.
//
// PATCH /api/invoicing/ledger/settings — Token permission: write:invoicing.
func (r *InvoicingLedgerSettingsService) Update(ctx context.Context, body *InvoicingLedgerSettingsUpdateBody, query *InvoicingLedgerSettingsUpdateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "PATCH", "/api/invoicing/ledger/settings", query, body, false, "settings", &out, opts)
	return out, err
}

// InvoicingLedgerYearEndService — invoicing.ledger.yearEnd
type InvoicingLedgerYearEndService struct {
	c *Client
}

func newInvoicingLedgerYearEndService(c *Client) *InvoicingLedgerYearEndService {
	return &InvoicingLedgerYearEndService{c: c}
}

// Get — What closing a fiscal year would do, or did: the result, the carried-forward balances, whether it can be closed.
//
// GET /api/invoicing/ledger/year-end — Token permission: read:invoicing.
func (r *InvoicingLedgerYearEndService) Get(ctx context.Context, query *InvoicingLedgerYearEndGetQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/ledger/year-end", query, nil, false, "yearEnd", &out, opts)
	return out, err
}

// Set — Closes a fiscal year with its closing entries, or reopens it.
//
// Needs `invoicing:bookkeeping` and the ledger on the plan. Balance method only, every month closed.
//
// POST /api/invoicing/ledger/year-end — Token permission: write:invoicing.
func (r *InvoicingLedgerYearEndService) Set(ctx context.Context, body *InvoicingLedgerYearEndSetBody, query *InvoicingLedgerYearEndSetQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/ledger/year-end", query, body, false, "yearEnd", &out, opts)
	return out, err
}

// InvoicingNumberRangesService — invoicing.numberRanges
type InvoicingNumberRangesService struct {
	c *Client
}

func newInvoicingNumberRangesService(c *Client) *InvoicingNumberRangesService {
	return &InvoicingNumberRangesService{c: c}
}

// List — Lists the number ranges — one per company and document type, and the customer-number range — with the next number each would draw.
//
// GET /api/invoicing/number-ranges — Token permission: read:invoicing.
func (r *InvoicingNumberRangesService) List(ctx context.Context, query *InvoicingNumberRangesListQuery, opts ...CallOption) ([]InvoicingNumberRange, error) {
	var out []InvoicingNumberRange
	err := r.c.call(ctx, "GET", "/api/invoicing/number-ranges", query, nil, false, "ranges", &out, opts)
	return out, err
}

// Preview — What a number pattern would produce; nothing is drawn or stored.
//
// POST /api/invoicing/number-ranges/preview — Token permission: write:invoicing.
func (r *InvoicingNumberRangesService) Preview(ctx context.Context, body *InvoicingNumberRangesPreviewBody, query *InvoicingNumberRangesPreviewQuery, opts ...CallOption) (*InvoicingNumberRangesPreviewResponse, error) {
	var out *InvoicingNumberRangesPreviewResponse
	err := r.c.call(ctx, "POST", "/api/invoicing/number-ranges/preview", query, body, false, "", &out, opts)
	return out, err
}

// Update — Changes a number range's format, until a number was drawn from it.
//
// Needs `invoicing:settings`. 409 `invoicingRangeLocked` once a number was drawn; 422 `invoicingInvalidPattern` with `issues`.
//
// PATCH /api/invoicing/number-ranges/{rangeId} — Token permission: write:invoicing.
//
// rangeID: The number range.
func (r *InvoicingNumberRangesService) Update(ctx context.Context, rangeID string, body *InvoicingNumberRangesUpdateBody, query *InvoicingNumberRangesUpdateQuery, opts ...CallOption) (*InvoicingNumberRange, error) {
	var out *InvoicingNumberRange
	err := r.c.call(ctx, "PATCH", "/api/invoicing/number-ranges/"+pathSeg(rangeID), query, body, false, "range", &out, opts)
	return out, err
}

// InvoicingPaymentRunsService — invoicing.paymentRuns
type InvoicingPaymentRunsService struct {
	c *Client
}

func newInvoicingPaymentRunsService(c *Client) *InvoicingPaymentRunsService {
	return &InvoicingPaymentRunsService{c: c}
}

// Close — Confirms that the bank executed a payment run (each item becomes a payment), or cancels it.
//
// Needs `invoicing:payments`. 409 `invoicingPaymentRunClosed` once confirmed or cancelled.
//
// POST /api/invoicing/payment-runs/{runId} — Token permission: write:invoicing.
//
// runID: The payment run.
func (r *InvoicingPaymentRunsService) Close(ctx context.Context, runID string, body *InvoicingPaymentRunsCloseBody, query *InvoicingPaymentRunsCloseQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/payment-runs/"+pathSeg(runID), query, body, false, "run", &out, opts)
	return out, err
}

// Create — Creates a SEPA payment run: a credit transfer file (pain.001) paying purchase invoices, to upload to online banking.
//
// Needs `invoicing:payments`. Approved or booked invoices of that company with money owed, in EUR, with a payee IBAN, in no other open run; from one of the company's bank accounts; on a date not in the past. 409 `invoicingPaymentRunNotPayable` with `refused`. Nothing is sent to a bank.
//
// POST /api/invoicing/payment-runs — Token permission: write:invoicing.
func (r *InvoicingPaymentRunsService) Create(ctx context.Context, body *InvoicingPaymentRunsCreateBody, query *InvoicingPaymentRunsCreateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/payment-runs", query, body, false, "run", &out, opts)
	return out, err
}

// File — Downloads a payment run's pain.001.001.09 XML exactly as written.
//
// GET /api/invoicing/payment-runs/{runId}/file — Token permission: read:invoicing.
//
// runID: The payment run.
func (r *InvoicingPaymentRunsService) File(ctx context.Context, runID string, query *InvoicingPaymentRunsFileQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/invoicing/payment-runs/"+pathSeg(runID)+"/file", query, nil, false, "", &out, opts)
	return out, err
}

// Get — One SEPA payment run with its items.
//
// GET /api/invoicing/payment-runs/{runId} — Token permission: read:invoicing.
//
// runID: The payment run.
func (r *InvoicingPaymentRunsService) Get(ctx context.Context, runID string, query *InvoicingPaymentRunsGetQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/payment-runs/"+pathSeg(runID), query, nil, false, "run", &out, opts)
	return out, err
}

// List — Lists the SEPA payment runs, newest first.
//
// GET /api/invoicing/payment-runs — Token permission: read:invoicing.
func (r *InvoicingPaymentRunsService) List(ctx context.Context, query *InvoicingPaymentRunsListQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/payment-runs", query, nil, false, "runs", &out, opts)
	return out, err
}

// InvoicingPaymentsService — invoicing.payments
type InvoicingPaymentsService struct {
	c *Client
}

func newInvoicingPaymentsService(c *Client) *InvoicingPaymentsService {
	return &InvoicingPaymentsService{c: c}
}

// List — The payments recorded on an invoice, oldest first.
//
// GET /api/invoicing/documents/{documentId}/payments — Token permission: read:invoicing.
//
// documentID: The document (quote, invoice, credit note…).
func (r *InvoicingPaymentsService) List(ctx context.Context, documentID string, query *InvoicingPaymentsListQuery, opts ...CallOption) ([]InvoicingPayment, error) {
	var out []InvoicingPayment
	err := r.c.call(ctx, "GET", "/api/invoicing/documents/"+pathSeg(documentID)+"/payments", query, nil, false, "payments", &out, opts)
	return out, err
}

// Record — Records a payment on an issued invoice, or a refund paid back (negative).
//
// Needs `invoicing:payments`. The payment status follows (open → partially_paid → paid). 409 `invoicingNotPayable` for anything but an issued invoice.
//
// POST /api/invoicing/documents/{documentId}/payments — Token permission: write:invoicing.
//
// documentID: The document (quote, invoice, credit note…).
func (r *InvoicingPaymentsService) Record(ctx context.Context, documentID string, body *InvoicingPaymentsRecordBody, query *InvoicingPaymentsRecordQuery, opts ...CallOption) (*InvoicingPayment, error) {
	var out *InvoicingPayment
	err := r.c.call(ctx, "POST", "/api/invoicing/documents/"+pathSeg(documentID)+"/payments", query, body, false, "payment", &out, opts)
	return out, err
}

// Remove — Removes a payment recorded by mistake; the payment status is computed again.
//
// Needs `invoicing:payments`.
//
// DELETE /api/invoicing/documents/{documentId}/payments/{paymentId} — Token permission: write:invoicing.
//
// documentID: The invoice.
// paymentID: The payment.
func (r *InvoicingPaymentsService) Remove(ctx context.Context, documentID string, paymentID string, query *InvoicingPaymentsRemoveQuery, opts ...CallOption) (*InvoicingDocument, error) {
	var out *InvoicingDocument
	err := r.c.call(ctx, "DELETE", "/api/invoicing/documents/"+pathSeg(documentID)+"/payments/"+pathSeg(paymentID), query, nil, false, "document", &out, opts)
	return out, err
}

// InvoicingPurchasesService — invoicing.purchases
type InvoicingPurchasesService struct {
	c *Client
}

func newInvoicingPurchasesService(c *Client) *InvoicingPurchasesService {
	return &InvoicingPurchasesService{c: c}
}

// AIFill — Has the platform AI read a draft purchase invoice's original receipt and suggest supplier, numbers, dates, amounts, rate and IBAN.
//
// Needs `invoicing:create` and the plan's AI for invoicing. The suggestion is kept on the draft beside its fields; nothing is filled in or booked by it. 503 `invoicingPurchaseAiUnavailable` without platform AI; 429 when the organisation's AI quota is used up; 422 `invoicingPurchaseAiUnreadable` with `reason` (`heicNotReadable`, `noTextLayer`, `noOriginal`). 10 a minute.
//
// POST /api/invoicing/purchases/{purchaseId}/ai-fill — Token permission: write:invoicing.
//
// purchaseID: The purchase invoice.
func (r *InvoicingPurchasesService) AIFill(ctx context.Context, purchaseID string, query *InvoicingPurchasesAiFillQuery, opts ...CallOption) (*InvoicingPurchaseAiSuggestion, error) {
	var out *InvoicingPurchaseAiSuggestion
	err := r.c.call(ctx, "POST", "/api/invoicing/purchases/"+pathSeg(purchaseID)+"/ai-fill", query, nil, false, "suggestion", &out, opts)
	return out, err
}

// AttachOriginal — Attaches the original to a purchase invoice that has none; written once.
//
// Needs `invoicing:create`. PDF, JPEG, PNG or HEIC, 20 MB at most. 409 `invoicingPurchaseOriginalExists`.
//
// POST /api/invoicing/purchases/{purchaseId}/original — Token permission: write:invoicing.
//
// purchaseID: The purchase invoice.
func (r *InvoicingPurchasesService) AttachOriginal(ctx context.Context, purchaseID string, body *InvoicingPurchasesAttachOriginalBody, query *InvoicingPurchasesAttachOriginalQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/purchases/"+pathSeg(purchaseID)+"/original", query, body, true, "purchase", &out, opts)
	return out, err
}

// BankLines — The bank lines a purchase invoice was paid with, as the Banking hub matched them.
//
// Needs `invoicing:read`. Date and amount always; account, counterpart and purpose only where the caller may read that account's lines. `banking` is false when the caller has no Banking hub.
//
// GET /api/invoicing/purchases/{purchaseId}/bank — Token permission: read:invoicing.
//
// purchaseID: The purchase invoice.
func (r *InvoicingPurchasesService) BankLines(ctx context.Context, purchaseID string, query *InvoicingPurchasesBankLinesQuery, opts ...CallOption) ([]InvoicingBankLine, error) {
	var out []InvoicingBankLine
	err := r.c.call(ctx, "GET", "/api/invoicing/purchases/"+pathSeg(purchaseID)+"/bank", query, nil, false, "lines", &out, opts)
	return out, err
}

// Create — Enters a purchase invoice by hand, as a draft.
//
// Needs `invoicing:create` and purchases on the plan. The due date follows the supplier's payment terms, the payee IBAN the supplier. A receipt can be uploaded instead as multipart: `file` (PDF, JPEG, PNG or HEIC, 20 MB at most) and optional `data` (this JSON). 30 a minute.
//
// POST /api/invoicing/purchases — Token permission: write:invoicing.
func (r *InvoicingPurchasesService) Create(ctx context.Context, body *InvoicingPurchaseInput, query *InvoicingPurchasesCreateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/purchases", query, body, false, "purchase", &out, opts)
	return out, err
}

// Delete — Deletes a draft purchase invoice.
//
// 409 `invoicingPurchaseKept` once approved or booked.
//
// DELETE /api/invoicing/purchases/{purchaseId} — Token permission: write:invoicing.
//
// purchaseID: The purchase invoice.
func (r *InvoicingPurchasesService) Delete(ctx context.Context, purchaseID string, query *InvoicingPurchasesDeleteQuery, opts ...CallOption) (*InvoicingPurchasesDeleteResponse, error) {
	var out *InvoicingPurchasesDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/invoicing/purchases/"+pathSeg(purchaseID), query, nil, false, "", &out, opts)
	return out, err
}

// DeletePayment — Removes a payment recorded on a purchase invoice by mistake; a posted one is reversed in the journal.
//
// Needs `invoicing:payments`.
//
// DELETE /api/invoicing/purchases/{purchaseId}/payments/{paymentId} — Token permission: write:invoicing.
//
// purchaseID: The purchase invoice.
// paymentID: The payment.
func (r *InvoicingPurchasesService) DeletePayment(ctx context.Context, purchaseID string, paymentID string, query *InvoicingPurchasesDeletePaymentQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "DELETE", "/api/invoicing/purchases/"+pathSeg(purchaseID)+"/payments/"+pathSeg(paymentID), query, nil, false, "purchase", &out, opts)
	return out, err
}

// FromInbound — Makes a purchase invoice from a received e-invoice: the supplier matched or created, the lines with their tax.
//
// Needs `invoicing:create`. One per received document (409 `invoicingPurchaseExists`); 422 `invoicingPurchaseNotEinvoice` for a file that is not an e-invoice.
//
// POST /api/invoicing/purchases/from-inbound — Token permission: write:invoicing.
func (r *InvoicingPurchasesService) FromInbound(ctx context.Context, body *InvoicingPurchasesFromInboundBody, query *InvoicingPurchasesFromInboundQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/purchases/from-inbound", query, body, false, "purchase", &out, opts)
	return out, err
}

// FromInboundStatus — The purchase invoice made from a received e-invoice, if any.
//
// `purchase` is null when there is none or the caller may not see it; `exists` says which.
//
// GET /api/invoicing/purchases/from-inbound — Token permission: read:invoicing.
func (r *InvoicingPurchasesService) FromInboundStatus(ctx context.Context, query *InvoicingPurchasesFromInboundStatusQuery, opts ...CallOption) (*InvoicingPurchasesFromInboundStatusResponse, error) {
	var out *InvoicingPurchasesFromInboundStatusResponse
	err := r.c.call(ctx, "GET", "/api/invoicing/purchases/from-inbound", query, nil, false, "", &out, opts)
	return out, err
}

// Get — One purchase invoice with its lines, payments, history, supplier, approval and original.
//
// GET /api/invoicing/purchases/{purchaseId} — Token permission: read:invoicing.
//
// purchaseID: The purchase invoice.
func (r *InvoicingPurchasesService) Get(ctx context.Context, purchaseID string, query *InvoicingPurchasesGetQuery, opts ...CallOption) (*InvoicingPurchasesGetResponse, error) {
	var out *InvoicingPurchasesGetResponse
	err := r.c.call(ctx, "GET", "/api/invoicing/purchases/"+pathSeg(purchaseID), query, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the purchase invoices the caller may see, newest first, with totals and what waits for approval, is overdue or is to be paid.
//
// Needs purchases on the plan.
//
// GET /api/invoicing/purchases — Token permission: read:invoicing.
func (r *InvoicingPurchasesService) List(ctx context.Context, query *InvoicingPurchasesListQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/purchases", query, nil, false, "purchases", &out, opts)
	return out, err
}

// Options — What a purchase invoice's lines can be booked to: the issuing companies with their bank accounts, and for one company its chart's expense and asset accounts — or, with its books off, the categories already used.
//
// Needs `invoicing:read`. Without `issuerId`, the default company. `taxCodes` is the fixed list a line's `taxCode` takes.
//
// GET /api/invoicing/purchases/options — Token permission: read:invoicing.
func (r *InvoicingPurchasesService) Options(ctx context.Context, query *InvoicingPurchasesOptionsQuery, opts ...CallOption) (*InvoicingPurchasesOptionsResponse, error) {
	var out *InvoicingPurchasesOptionsResponse
	err := r.c.call(ctx, "GET", "/api/invoicing/purchases/options", query, nil, false, "", &out, opts)
	return out, err
}

// Original — Downloads a purchase invoice's original as stored: the uploaded receipt or the received e-invoice.
//
// GET /api/invoicing/purchases/{purchaseId}/original — Token permission: read:invoicing.
//
// purchaseID: The purchase invoice.
func (r *InvoicingPurchasesService) Original(ctx context.Context, purchaseID string, query *InvoicingPurchasesOriginalQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/invoicing/purchases/"+pathSeg(purchaseID)+"/original", query, nil, false, "", &out, opts)
	return out, err
}

// RecordPayment — Records a payment on an approved or booked purchase invoice — by default what is open, today, by bank transfer.
//
// Needs `invoicing:payments`. More than is open: 422 `invoicingPurchaseExceedsOpen`.
//
// POST /api/invoicing/purchases/{purchaseId}/payments — Token permission: write:invoicing.
//
// purchaseID: The purchase invoice.
func (r *InvoicingPurchasesService) RecordPayment(ctx context.Context, purchaseID string, body *InvoicingPurchasesRecordPaymentBody, query *InvoicingPurchasesRecordPaymentQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/purchases/"+pathSeg(purchaseID)+"/payments", query, body, false, "payment", &out, opts)
	return out, err
}

// Transition — Submits a purchase invoice for approval, approves or rejects it, reopens it, or books it.
//
// submit (`invoicing:create`): the approval steps are worked out and the approvers told. approve / reject (with a `note`): the step's approver; who decided step 1 does not decide step 2. reopen (`invoicing:payments`): approved → draft. book (`invoicing:payments`): approved → booked, posted to the ledger when the books are on. 422 `invoicingPurchaseIncomplete` with `issues`; 409 `invoicingPurchaseInvalidTransition`.
//
// POST /api/invoicing/purchases/{purchaseId}/status — Token permission: write:invoicing.
//
// purchaseID: The purchase invoice.
func (r *InvoicingPurchasesService) Transition(ctx context.Context, purchaseID string, body *InvoicingPurchasesTransitionBody, query *InvoicingPurchasesTransitionQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/purchases/"+pathSeg(purchaseID)+"/status", query, body, false, "purchase", &out, opts)
	return out, err
}

// Update — Changes a draft purchase invoice; `lines` replace them all.
//
// Needs `invoicing:create`. 409 `invoicingPurchaseNotDraft` once approved or booked — except `visibility` / `workspaceId` alone. A waiting approval is withdrawn.
//
// PATCH /api/invoicing/purchases/{purchaseId} — Token permission: write:invoicing.
//
// purchaseID: The purchase invoice.
func (r *InvoicingPurchasesService) Update(ctx context.Context, purchaseID string, body *InvoicingPurchaseInput, query *InvoicingPurchasesUpdateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "PATCH", "/api/invoicing/purchases/"+pathSeg(purchaseID), query, body, false, "purchase", &out, opts)
	return out, err
}

// InvoicingRecurringService — invoicing.recurring
type InvoicingRecurringService struct {
	c *Client
}

func newInvoicingRecurringService(c *Client) *InvoicingRecurringService {
	return &InvoicingRecurringService{c: c}
}

// Create — Creates a recurring invoice: an invoice repeated on a schedule, as a draft, finalized, or finalized and sent.
//
// Needs `invoicing:create` and recurring invoices on the plan; the modes `finalize` and `finalize_send` need `invoicing:finalize`.
//
// POST /api/invoicing/recurring — Token permission: write:invoicing.
func (r *InvoicingRecurringService) Create(ctx context.Context, body *InvoicingRecurringCreateBody, query *InvoicingRecurringCreateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/recurring", query, body, false, "profile", &out, opts)
	return out, err
}

// Delete — Deletes a recurring invoice and its run history; the invoices it made stay.
//
// DELETE /api/invoicing/recurring/{profileId} — Token permission: write:invoicing.
//
// profileID: The recurring invoice.
func (r *InvoicingRecurringService) Delete(ctx context.Context, profileID string, query *InvoicingRecurringDeleteQuery, opts ...CallOption) (*InvoicingRecurringDeleteResponse, error) {
	var out *InvoicingRecurringDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/invoicing/recurring/"+pathSeg(profileID), query, nil, false, "", &out, opts)
	return out, err
}

// Get — One recurring invoice with its run history, newest first.
//
// GET /api/invoicing/recurring/{profileId} — Token permission: read:invoicing.
//
// profileID: The recurring invoice.
func (r *InvoicingRecurringService) Get(ctx context.Context, profileID string, query *InvoicingRecurringGetQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/recurring/"+pathSeg(profileID), query, nil, false, "profile", &out, opts)
	return out, err
}

// List — Lists the recurring invoices the caller may see, with their next runs.
//
// GET /api/invoicing/recurring — Token permission: read:invoicing.
func (r *InvoicingRecurringService) List(ctx context.Context, query *InvoicingRecurringListQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/recurring", query, nil, false, "profiles", &out, opts)
	return out, err
}

// Prefill — What a new recurring invoice made from an invoice starts from; nothing is saved.
//
// GET /api/invoicing/recurring/from-document — Token permission: read:invoicing.
func (r *InvoicingRecurringService) Prefill(ctx context.Context, query *InvoicingRecurringPrefillQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/recurring/from-document", query, nil, false, "prefill", &out, opts)
	return out, err
}

// Update — Changes a recurring invoice, or pauses and resumes it.
//
// Needs `invoicing:create` and recurring invoices on the plan. Resuming starts from today: periods missed while paused are not billed.
//
// PATCH /api/invoicing/recurring/{profileId} — Token permission: write:invoicing.
//
// profileID: The recurring invoice.
func (r *InvoicingRecurringService) Update(ctx context.Context, profileID string, body *InvoicingRecurringUpdateBody, query *InvoicingRecurringUpdateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "PATCH", "/api/invoicing/recurring/"+pathSeg(profileID), query, body, false, "profile", &out, opts)
	return out, err
}

// InvoicingReportsService — invoicing.reports
type InvoicingReportsService struct {
	c *Client
}

func newInvoicingReportsService(c *Client) *InvoicingReportsService {
	return &InvoicingReportsService{c: c}
}

// Dataset — One dataset of the hub's numbers as rows: revenue, open items, VAT, quote conversion, expenses, payables, result.
//
// Read as the caller: the whole books with `invoicing:export`, else what they may see. Amounts in major units with each row's currency. The purchase datasets need purchases on the plan.
//
// GET /api/invoicing/reports — Token permission: read:invoicing.
func (r *InvoicingReportsService) Dataset(ctx context.Context, query *InvoicingReportsDatasetQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/reports", query, nil, false, "rows", &out, opts)
	return out, err
}

// TaxReturns — One company's EC sales list (Zusammenfassende Meldung), One-Stop-Shop figures and the €10,000 threshold for a period.
//
// Needs `invoicing:export`; the whole books. `format=ec-sales` or `oss` downloads that one as CSV.
//
// GET /api/invoicing/reports/tax-returns — Token permission: read:invoicing.
func (r *InvoicingReportsService) TaxReturns(ctx context.Context, query *InvoicingReportsTaxReturnsQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/reports/tax-returns", query, nil, false, "", &out, opts)
	return out, err
}

// InvoicingSuppliersService — invoicing.suppliers
type InvoicingSuppliersService struct {
	c *Client
}

func newInvoicingSuppliersService(c *Client) *InvoicingSuppliersService {
	return &InvoicingSuppliersService{c: c}
}

// Create — Creates a supplier.
//
// Needs `invoicing:create`. IBAN and BIC are checked; a VAT ID, tax number or IBAN another supplier has answers 409 `invoicingSupplierDuplicate` with `duplicates` unless `allowDuplicate`.
//
// POST /api/invoicing/suppliers — Token permission: write:invoicing.
func (r *InvoicingSuppliersService) Create(ctx context.Context, body *InvoicingSupplierInput, query *InvoicingSuppliersCreateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/invoicing/suppliers", query, body, false, "supplier", &out, opts)
	return out, err
}

// Delete — Deletes a supplier no purchase invoice names.
//
// 409 `invoicingSupplierHasInvoices` otherwise — archive them instead.
//
// DELETE /api/invoicing/suppliers/{supplierId} — Token permission: write:invoicing.
//
// supplierID: The supplier.
func (r *InvoicingSuppliersService) Delete(ctx context.Context, supplierID string, query *InvoicingSuppliersDeleteQuery, opts ...CallOption) (*InvoicingSuppliersDeleteResponse, error) {
	var out *InvoicingSuppliersDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/invoicing/suppliers/"+pathSeg(supplierID), query, nil, false, "", &out, opts)
	return out, err
}

// Get — One supplier, with how many purchase invoices name them and what is open per currency.
//
// GET /api/invoicing/suppliers/{supplierId} — Token permission: read:invoicing.
//
// supplierID: The supplier.
func (r *InvoicingSuppliersService) Get(ctx context.Context, supplierID string, query *InvoicingSuppliersGetQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/suppliers/"+pathSeg(supplierID), query, nil, false, "supplier", &out, opts)
	return out, err
}

// List — Lists the suppliers the caller may see, by name.
//
// GET /api/invoicing/suppliers — Token permission: read:invoicing.
func (r *InvoicingSuppliersService) List(ctx context.Context, query *InvoicingSuppliersListQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/invoicing/suppliers", query, nil, false, "suppliers", &out, opts)
	return out, err
}

// Update — Changes a supplier, or archives and restores them.
//
// PATCH /api/invoicing/suppliers/{supplierId} — Token permission: write:invoicing.
//
// supplierID: The supplier.
func (r *InvoicingSuppliersService) Update(ctx context.Context, supplierID string, body *InvoicingSuppliersUpdateBody, query *InvoicingSuppliersUpdateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "PATCH", "/api/invoicing/suppliers/"+pathSeg(supplierID), query, body, false, "supplier", &out, opts)
	return out, err
}

// InvoicingTemplatesService — invoicing.templates
type InvoicingTemplatesService struct {
	c *Client
}

func newInvoicingTemplatesService(c *Client) *InvoicingTemplatesService {
	return &InvoicingTemplatesService{c: c}
}

// Create — Creates a PDF template; the first is the default.
//
// Needs `invoicing:settings`; counts against the plan's template limit. Images must be the organisation's; the layout is checked strictly (400 `invoicingUnknownMergeField` / `invoicingLayoutInvalid`).
//
// POST /api/invoicing/templates — Token permission: write:invoicing.
func (r *InvoicingTemplatesService) Create(ctx context.Context, body *InvoicingTemplatesCreateBody, query *InvoicingTemplatesCreateQuery, opts ...CallOption) (*InvoicingTemplate, error) {
	var out *InvoicingTemplate
	err := r.c.call(ctx, "POST", "/api/invoicing/templates", query, body, false, "template", &out, opts)
	return out, err
}

// Delete — Deletes a PDF template; documents that used it render with the default.
//
// DELETE /api/invoicing/templates/{templateId} — Token permission: write:invoicing.
//
// templateID: The template.
func (r *InvoicingTemplatesService) Delete(ctx context.Context, templateID string, query *InvoicingTemplatesDeleteQuery, opts ...CallOption) (*InvoicingTemplatesDeleteResponse, error) {
	var out *InvoicingTemplatesDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/invoicing/templates/"+pathSeg(templateID), query, nil, false, "", &out, opts)
	return out, err
}

// Get — One PDF template.
//
// GET /api/invoicing/templates/{templateId} — Token permission: read:invoicing.
//
// templateID: The template.
func (r *InvoicingTemplatesService) Get(ctx context.Context, templateID string, query *InvoicingTemplatesGetQuery, opts ...CallOption) (*InvoicingTemplate, error) {
	var out *InvoicingTemplate
	err := r.c.call(ctx, "GET", "/api/invoicing/templates/"+pathSeg(templateID), query, nil, false, "template", &out, opts)
	return out, err
}

// List — Lists the organisation's PDF templates, with the layout every unset key falls back to and the fonts on offer.
//
// GET /api/invoicing/templates — Token permission: read:invoicing.
func (r *InvoicingTemplatesService) List(ctx context.Context, query *InvoicingTemplatesListQuery, opts ...CallOption) ([]InvoicingTemplate, error) {
	var out []InvoicingTemplate
	err := r.c.call(ctx, "GET", "/api/invoicing/templates", query, nil, false, "templates", &out, opts)
	return out, err
}

// Preview — Renders a template as a PDF — with a given document, or the built-in sample — without storing anything.
//
// Needs `invoicing:settings`. Answers the PDF. The config sent, else the template `templateId`, else the default. 30 a minute.
//
// POST /api/invoicing/templates/preview — Token permission: write:invoicing.
func (r *InvoicingTemplatesService) Preview(ctx context.Context, body *InvoicingTemplatesPreviewBody, query *InvoicingTemplatesPreviewQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "POST", "/api/invoicing/templates/preview", query, body, false, "", &out, opts)
	return out, err
}

// Update — Changes a PDF template; `config` and `texts` replace what was stored.
//
// Needs `invoicing:settings`. Send `config` whole, `layout` included.
//
// PATCH /api/invoicing/templates/{templateId} — Token permission: write:invoicing.
//
// templateID: The template.
func (r *InvoicingTemplatesService) Update(ctx context.Context, templateID string, body *InvoicingTemplatesUpdateBody, query *InvoicingTemplatesUpdateQuery, opts ...CallOption) (*InvoicingTemplate, error) {
	var out *InvoicingTemplate
	err := r.c.call(ctx, "PATCH", "/api/invoicing/templates/"+pathSeg(templateID), query, body, false, "template", &out, opts)
	return out, err
}

// ItemFoldersService — Item folders: Folders for the lists of databases, documents, endpoints, reports, projects and workflows — shared per workspace and per kind. Organising only: an item's address, sharing and content never change with its folder. Folders nest up to five levels; an item is in one folder or in none. Reading needs `<kind>:read` and the workspace; making, renaming and deleting folders and filing items needs `<kind>:update`.
type ItemFoldersService struct {
	c *Client
}

func newItemFoldersService(c *Client) *ItemFoldersService {
	return &ItemFoldersService{c: c}
}

// Create — Creates a folder in one kind of list of a workspace.
//
// 400 `pageFolderCycle` when the folder would sit deeper than five levels; 404 `folderNotFound` for a parent that is not a folder of that list.
//
// POST /api/item-folders
func (r *ItemFoldersService) Create(ctx context.Context, body *ItemFoldersCreateBody, opts ...CallOption) (*ItemFolder, error) {
	var out *ItemFolder
	err := r.c.call(ctx, "POST", "/api/item-folders", nil, body, false, "data.folder", &out, opts)
	return out, err
}

// Delete — Deletes a folder; its items go to no folder and its subfolders move up one level.
//
// No item is deleted.
//
// DELETE /api/item-folders/{folderId}
//
// folderID: The folder's id (a uuid).
func (r *ItemFoldersService) Delete(ctx context.Context, folderID string, opts ...CallOption) (*ItemFoldersDeleteResponseData, error) {
	var out *ItemFoldersDeleteResponseData
	err := r.c.call(ctx, "DELETE", "/api/item-folders/"+pathSeg(folderID), nil, nil, false, "data", &out, opts)
	return out, err
}

// List — Lists the folders of one kind of item in a workspace, and which item is filed where.
//
// GET /api/item-folders
func (r *ItemFoldersService) List(ctx context.Context, query *ItemFoldersListQuery, opts ...CallOption) (*ItemFolders, error) {
	var out *ItemFolders
	err := r.c.call(ctx, "GET", "/api/item-folders", query, nil, false, "data", &out, opts)
	return out, err
}

// Move — Files items of one list in a folder, or takes them out of every folder.
//
// Only items of that kind in that workspace are filed; other ids are skipped, and `moved` counts the rest. Nothing about the items themselves changes.
//
// POST /api/item-folders/move
func (r *ItemFoldersService) Move(ctx context.Context, body *ItemFoldersMoveBody, opts ...CallOption) (float64, error) {
	var out float64
	err := r.c.call(ctx, "POST", "/api/item-folders/move", nil, body, false, "data.moved", &out, opts)
	return out, err
}

// Update — Renames a folder or moves it into another folder.
//
// Only the fields sent change. 400 `pageFolderCycle` when the move would put it inside itself or deeper than five levels.
//
// PATCH /api/item-folders/{folderId}
//
// folderID: The folder's id (a uuid).
func (r *ItemFoldersService) Update(ctx context.Context, folderID string, body *ItemFoldersUpdateBody, opts ...CallOption) (*ItemFolder, error) {
	var out *ItemFolder
	err := r.c.call(ctx, "PATCH", "/api/item-folders/"+pathSeg(folderID), nil, body, false, "data.folder", &out, opts)
	return out, err
}

// LanguagesService — Languages
type LanguagesService struct {
	c *Client
}

func newLanguagesService(c *Client) *LanguagesService {
	return &LanguagesService{c: c}
}

// List — Lists the languages the platform is available in.
//
// Answers without a token. Only enabled languages, in the administrator's order; cached for five minutes.
//
// GET /api/languages — No token needed.
func (r *LanguagesService) List(ctx context.Context, opts ...CallOption) ([]PlatformLanguage, error) {
	var out []PlatformLanguage
	err := r.c.call(ctx, "GET", "/api/languages", nil, nil, false, "", &out, opts)
	return out, err
}

// MailService — Mail: The Mail hub: mailboxes (IMAP, Gmail, Outlook and the platform's own address), conversations and messages, drafts, sending now or queued, labels, signatures, snippets, rules, blocked senders, out-of-office replies and PGP. Everything needs Mail on the caller's plan (403 otherwise). A mailbox is its owner's; one shared with an organisation can be read and sent from by its members, while its settings stay the owner's.
type MailService struct {
	c *Client

	AI   *MailAiService
	IMAP *MailImapService
	SMTP *MailSmtpService
}

func newMailService(c *Client) *MailService {
	return &MailService{c: c, AI: newMailAiService(c), IMAP: newMailImapService(c), SMTP: newMailSmtpService(c)}
}

// AddPgpKey — Adds an armored PGP key: a contact's public key or the caller's own key pair.
//
// Needs PGP on the plan and `mail:manage_pgp_keys`. A private key is stored encrypted and stays locked by its own passphrase. 409 when it is already there.
//
// POST /api/mail/pgp/keys
func (r *MailService) AddPgpKey(ctx context.Context, body *MailAddPgpKeyBody, opts ...CallOption) (*MailPgpKey, error) {
	var out *MailPgpKey
	err := r.c.call(ctx, "POST", "/api/mail/pgp/keys", nil, body, false, "data.key", &out, opts)
	return out, err
}

// ApplyInviteReply — Records a guest's answer, attached to a message, on the appointment it answers.
//
// Needs Calendar too. Only when the message's sender is the guest answering and already on the guest list of an appointment the caller can edit; doing it again changes nothing.
//
// POST /api/mail/attachments/{id}/apply-reply
//
// id: The attachment.
func (r *MailService) ApplyInviteReply(ctx context.Context, id string, opts ...CallOption) (*MailApplyInviteReplyResponseData, error) {
	var out *MailApplyInviteReplyResponseData
	err := r.c.call(ctx, "POST", "/api/mail/attachments/"+pathSeg(id)+"/apply-reply", nil, nil, false, "data", &out, opts)
	return out, err
}

// ApplyRule — Runs a rule once over the mail already stored.
//
// The newest 2,000 messages of a folder of the rule's mailbox (or of the caller's own mailboxes). Moves and deletes happen on the mail server too.
//
// POST /api/mail/rules/{id}/apply
//
// id: The rule.
func (r *MailService) ApplyRule(ctx context.Context, id string, body *MailApplyRuleBody, opts ...CallOption) (*MailApplyRuleResponseData, error) {
	var out *MailApplyRuleResponseData
	err := r.c.call(ctx, "POST", "/api/mail/rules/"+pathSeg(id)+"/apply", nil, body, false, "data", &out, opts)
	return out, err
}

// BlockSender — Blocks an address or a domain: new mail from it is filed in spam.
//
// Blocking one already blocked answers it with `created: false`.
//
// POST /api/mail/blocked-senders
func (r *MailService) BlockSender(ctx context.Context, body *MailBlockSenderBody, opts ...CallOption) (*MailBlockSenderResponseDataSender, error) {
	var out *MailBlockSenderResponseDataSender
	err := r.c.call(ctx, "POST", "/api/mail/blocked-senders", nil, body, false, "data.sender", &out, opts)
	return out, err
}

// CancelScheduled — Cancels a queued send that has not gone out, or dismisses a failed one.
//
// 409 when it is already sending or gone.
//
// DELETE /api/mail/scheduled/{id}
//
// id: The queued send.
func (r *MailService) CancelScheduled(ctx context.Context, id string, opts ...CallOption) (*MailCancelScheduledResponse, error) {
	var out *MailCancelScheduledResponse
	err := r.c.call(ctx, "DELETE", "/api/mail/scheduled/"+pathSeg(id), nil, nil, false, "", &out, opts)
	return out, err
}

// CheckPlatformAddress — Whether a name is still free on the platform's own mail domain.
//
// Not in the usual envelope. Without `username` it answers only the domain. `reason` when refused: contains_plus, too_short, invalid_chars.
//
// GET /api/mail/platform-address/check
func (r *MailService) CheckPlatformAddress(ctx context.Context, query *MailCheckPlatformAddressQuery, opts ...CallOption) (*MailCheckPlatformAddressResponse, error) {
	var out *MailCheckPlatformAddressResponse
	err := r.c.call(ctx, "GET", "/api/mail/platform-address/check", query, nil, false, "", &out, opts)
	return out, err
}

// Counts — Unread conversations per folder and per mailbox, starred, drafts and queued sends.
//
// GET /api/mail/counts
func (r *MailService) Counts(ctx context.Context, query *MailCountsQuery, opts ...CallOption) (*MailCountsResponseData, error) {
	var out *MailCountsResponseData
	err := r.c.call(ctx, "GET", "/api/mail/counts", query, nil, false, "data", &out, opts)
	return out, err
}

// CreateAccount — Connects an IMAP/SMTP mailbox and starts its first sync.
//
// Needs external mailboxes on the plan and room under its mailbox limit. The servers must be reachable public hosts (400 otherwise). Gmail and Outlook are connected in a browser; the platform's own address is POST /api/mail/platform-address.
//
// POST /api/mail/accounts
func (r *MailService) CreateAccount(ctx context.Context, body *MailCreateAccountBody, opts ...CallOption) (*MailAccount, error) {
	var out *MailAccount
	err := r.c.call(ctx, "POST", "/api/mail/accounts", nil, body, false, "data", &out, opts)
	return out, err
}

// CreateAlias — Adds a send-as address to a mailbox and mails a confirmation link to it.
//
// Only the mailbox's owner; anybody else gets 404. It can be sent from once the link in that inbox is opened. At most 20 per mailbox; 10 requests an hour.
//
// POST /api/mail/accounts/{accountId}/aliases
//
// accountID: The mailbox (mail account).
func (r *MailService) CreateAlias(ctx context.Context, accountID string, body *MailCreateAliasBody, opts ...CallOption) (*MailAlias, error) {
	var out *MailAlias
	err := r.c.call(ctx, "POST", "/api/mail/accounts/"+pathSeg(accountID)+"/aliases", nil, body, false, "data.alias", &out, opts)
	return out, err
}

// CreateDraft — Saves a new draft in a mailbox the caller can send from.
//
// A reply keeps In-Reply-To / References and its conversation, when the caller may open it.
//
// POST /api/mail/drafts
func (r *MailService) CreateDraft(ctx context.Context, body *MailCreateDraftBody, opts ...CallOption) (*MailDraft, error) {
	var out *MailDraft
	err := r.c.call(ctx, "POST", "/api/mail/drafts", nil, body, false, "data", &out, opts)
	return out, err
}

// CreateFolder — Creates a folder on the mailbox's server.
//
// Only the mailbox's owner; anybody else gets 404.
//
// POST /api/mail/accounts/{accountId}/folders
//
// accountID: The mailbox (mail account).
func (r *MailService) CreateFolder(ctx context.Context, accountID string, body *MailCreateFolderBody, opts ...CallOption) (*MailCreateFolderResponseData, error) {
	var out *MailCreateFolderResponseData
	err := r.c.call(ctx, "POST", "/api/mail/accounts/"+pathSeg(accountID)+"/folders", nil, body, false, "data", &out, opts)
	return out, err
}

// CreateLabel — Creates a label.
//
// At most 100 (429).
//
// POST /api/mail/labels
func (r *MailService) CreateLabel(ctx context.Context, body *MailCreateLabelBody, opts ...CallOption) (*MailLabel, error) {
	var out *MailLabel
	err := r.c.call(ctx, "POST", "/api/mail/labels", nil, body, false, "data", &out, opts)
	return out, err
}

// CreatePlatformAddress — Creates the caller's mailbox on the platform's own domain.
//
// Needs `mail:manage_platform` and the platform address on the plan. Having one already answers it. The name may not contain `+` (reserved for sub-addresses); without one it is made from the caller's name.
//
// POST /api/mail/platform-address
func (r *MailService) CreatePlatformAddress(ctx context.Context, body *MailCreatePlatformAddressBody, opts ...CallOption) (*MailCreatePlatformAddressResponseData, error) {
	var out *MailCreatePlatformAddressResponseData
	err := r.c.call(ctx, "POST", "/api/mail/platform-address", nil, body, false, "data", &out, opts)
	return out, err
}

// CreateRule — Creates a rule that acts on new mail.
//
// On one of the caller's own mailboxes, or on all of them. At most 100 (429).
//
// POST /api/mail/rules
func (r *MailService) CreateRule(ctx context.Context, body *MailCreateRuleBody, opts ...CallOption) (*MailRule, error) {
	var out *MailRule
	err := r.c.call(ctx, "POST", "/api/mail/rules", nil, body, false, "data.rule", &out, opts)
	return out, err
}

// CreateSignature — Creates a signature.
//
// The HTML is sanitised. At most 25 (429). A new default replaces the old one.
//
// POST /api/mail/signatures
func (r *MailService) CreateSignature(ctx context.Context, body *MailCreateSignatureBody, opts ...CallOption) (*MailSignature, error) {
	var out *MailSignature
	err := r.c.call(ctx, "POST", "/api/mail/signatures", nil, body, false, "data", &out, opts)
	return out, err
}

// CreateSnippet — Creates a snippet.
//
// The HTML is sanitised. At most 200 (429); a shortcut already used answers 409.
//
// POST /api/mail/snippets
func (r *MailService) CreateSnippet(ctx context.Context, body *MailCreateSnippetBody, opts ...CallOption) (*MailSnippet, error) {
	var out *MailSnippet
	err := r.c.call(ctx, "POST", "/api/mail/snippets", nil, body, false, "data.snippet", &out, opts)
	return out, err
}

// CreateThreadNote — Adds an internal note to a conversation.
//
// The people it mentions and the assignee are notified.
//
// POST /api/mail/threads/{threadId}/notes
//
// threadID: The conversation.
func (r *MailService) CreateThreadNote(ctx context.Context, threadID string, body *MailCreateThreadNoteBody, opts ...CallOption) (*MailThreadNote, error) {
	var out *MailThreadNote
	err := r.c.call(ctx, "POST", "/api/mail/threads/"+pathSeg(threadID)+"/notes", nil, body, false, "data.note", &out, opts)
	return out, err
}

// Decrypt — Decrypts a PGP-encrypted message with the caller's own keys and checks its signature.
//
// Needs PGP on the plan. The passphrase is used once and never stored; 428 when a key needs one and none was sent, 422 when no key opens it, 403 without a private key. `signatureValid` is null when there was nothing to check against.
//
// POST /api/mail/emails/{emailId}/decrypt
//
// emailID: The message.
func (r *MailService) Decrypt(ctx context.Context, emailID string, body *MailDecryptBody, opts ...CallOption) (*MailDecryptResponseData, error) {
	var out *MailDecryptResponseData
	err := r.c.call(ctx, "POST", "/api/mail/emails/"+pathSeg(emailID)+"/decrypt", nil, body, false, "data", &out, opts)
	return out, err
}

// DeleteAccount — Disconnects a mailbox and deletes its stored mail here.
//
// Only the mailbox's owner; anybody else gets 404. The platform address cannot be deleted (400). The mail on the server is not touched.
//
// DELETE /api/mail/accounts/{id}
//
// id: The mailbox (mail account).
func (r *MailService) DeleteAccount(ctx context.Context, id string, opts ...CallOption) (*MailDeleteAccountResponse, error) {
	var out *MailDeleteAccountResponse
	err := r.c.call(ctx, "DELETE", "/api/mail/accounts/"+pathSeg(id), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteAlias — Removes a send-as address from a mailbox.
//
// Only the mailbox's owner; anybody else gets 404.
//
// DELETE /api/mail/accounts/{accountId}/aliases/{aliasId}
//
// accountID: The mailbox (mail account).
// aliasID: The send-as address.
func (r *MailService) DeleteAlias(ctx context.Context, accountID string, aliasID string, opts ...CallOption) (*MailDeleteAliasResponse, error) {
	var out *MailDeleteAliasResponse
	err := r.c.call(ctx, "DELETE", "/api/mail/accounts/"+pathSeg(accountID)+"/aliases/"+pathSeg(aliasID), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteAttachment — Deletes an upload that has not been sent yet.
//
// An attachment of a sent message cannot be deleted (403). Answers 204.
//
// DELETE /api/mail/attachments/{id}
//
// id: The attachment.
func (r *MailService) DeleteAttachment(ctx context.Context, id string, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "DELETE", "/api/mail/attachments/"+pathSeg(id), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteDraft — Deletes a draft of the caller's.
//
// DELETE /api/mail/drafts/{id}
//
// id: The draft.
func (r *MailService) DeleteDraft(ctx context.Context, id string, opts ...CallOption) (*MailDeleteDraftResponse, error) {
	var out *MailDeleteDraftResponse
	err := r.c.call(ctx, "DELETE", "/api/mail/drafts/"+pathSeg(id), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteFolder — Deletes a folder on the mailbox's server.
//
// Only the mailbox's owner; anybody else gets 404.
//
// DELETE /api/mail/accounts/{accountId}/folders
//
// accountID: The mailbox (mail account).
func (r *MailService) DeleteFolder(ctx context.Context, accountID string, query *MailDeleteFolderQuery, opts ...CallOption) (*MailDeleteFolderResponseData, error) {
	var out *MailDeleteFolderResponseData
	err := r.c.call(ctx, "DELETE", "/api/mail/accounts/"+pathSeg(accountID)+"/folders", query, nil, false, "data", &out, opts)
	return out, err
}

// DeleteLabel — Deletes a label.
//
// A system label cannot be deleted (400).
//
// DELETE /api/mail/labels/{id}
//
// id: The label.
func (r *MailService) DeleteLabel(ctx context.Context, id string, opts ...CallOption) (*MailDeleteLabelResponse, error) {
	var out *MailDeleteLabelResponse
	err := r.c.call(ctx, "DELETE", "/api/mail/labels/"+pathSeg(id), nil, nil, false, "", &out, opts)
	return out, err
}

// DeletePgpKey — Removes a PGP key from the caller's keyring for good.
//
// Needs PGP on the plan and `mail:manage_pgp_keys`.
//
// DELETE /api/mail/pgp/keys/{id}
//
// id: The key.
func (r *MailService) DeletePgpKey(ctx context.Context, id string, opts ...CallOption) (*MailDeletePgpKeyResponse, error) {
	var out *MailDeletePgpKeyResponse
	err := r.c.call(ctx, "DELETE", "/api/mail/pgp/keys/"+pathSeg(id), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteRule — Deletes a rule.
//
// DELETE /api/mail/rules/{id}
//
// id: The rule.
func (r *MailService) DeleteRule(ctx context.Context, id string, opts ...CallOption) (*MailDeleteRuleResponse, error) {
	var out *MailDeleteRuleResponse
	err := r.c.call(ctx, "DELETE", "/api/mail/rules/"+pathSeg(id), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteSignature — Deletes a signature.
//
// DELETE /api/mail/signatures/{id}
//
// id: The signature.
func (r *MailService) DeleteSignature(ctx context.Context, id string, opts ...CallOption) (*MailDeleteSignatureResponse, error) {
	var out *MailDeleteSignatureResponse
	err := r.c.call(ctx, "DELETE", "/api/mail/signatures/"+pathSeg(id), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteSnippet — Deletes a snippet.
//
// DELETE /api/mail/snippets/{id}
//
// id: The snippet.
func (r *MailService) DeleteSnippet(ctx context.Context, id string, opts ...CallOption) (*MailDeleteSnippetResponse, error) {
	var out *MailDeleteSnippetResponse
	err := r.c.call(ctx, "DELETE", "/api/mail/snippets/"+pathSeg(id), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteThread — Moves a conversation to the trash; one already in the trash is deleted for good.
//
// On the mail server too.
//
// DELETE /api/mail/threads/{threadId}
//
// threadID: The conversation.
func (r *MailService) DeleteThread(ctx context.Context, threadID string, opts ...CallOption) (*MailDeleteThreadResponse, error) {
	var out *MailDeleteThreadResponse
	err := r.c.call(ctx, "DELETE", "/api/mail/threads/"+pathSeg(threadID), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteThreadNote — Deletes a note from a conversation.
//
// Its author or the conversation's owner (403 otherwise).
//
// DELETE /api/mail/threads/{threadId}/notes/{noteId}
//
// threadID: The conversation.
// noteID: The note.
func (r *MailService) DeleteThreadNote(ctx context.Context, threadID string, noteID string, opts ...CallOption) (*MailDeleteThreadNoteResponse, error) {
	var out *MailDeleteThreadNoteResponse
	err := r.c.call(ctx, "DELETE", "/api/mail/threads/"+pathSeg(threadID)+"/notes/"+pathSeg(noteID), nil, nil, false, "", &out, opts)
	return out, err
}

// Desk — The inbox as an Action desk: the open conversations that still ask something, each in one of five sections with the reason.
//
// In the inbox, not snoozed or resolved; the newest 200.
//
// GET /api/mail/desk
func (r *MailService) Desk(ctx context.Context, query *MailDeskQuery, opts ...CallOption) (*MailDeskResponseData, error) {
	var out *MailDeskResponseData
	err := r.c.call(ctx, "GET", "/api/mail/desk", query, nil, false, "data", &out, opts)
	return out, err
}

// DownloadAttachment — Downloads an attachment of a message the caller can open.
//
// GET /api/mail/attachments/{id}
//
// id: The attachment.
func (r *MailService) DownloadAttachment(ctx context.Context, id string, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/mail/attachments/"+pathSeg(id), nil, nil, false, "", &out, opts)
	return out, err
}

// EmptyTrash — Deletes everything in the trash for good, here and on the mail server.
//
// Of one mailbox, or of every mailbox of the caller's own.
//
// POST /api/mail/trash/empty
func (r *MailService) EmptyTrash(ctx context.Context, body *MailEmptyTrashBody, opts ...CallOption) (*MailEmptyTrashResponseData, error) {
	var out *MailEmptyTrashResponseData
	err := r.c.call(ctx, "POST", "/api/mail/trash/empty", nil, body, false, "data", &out, opts)
	return out, err
}

// GetAccount — One mailbox the caller can open.
//
// Never the password; `loginUsername` only for the owner of an IMAP mailbox.
//
// GET /api/mail/accounts/{id}
//
// id: The mailbox (mail account).
func (r *MailService) GetAccount(ctx context.Context, id string, opts ...CallOption) (*MailGetAccountResponseData, error) {
	var out *MailGetAccountResponseData
	err := r.c.call(ctx, "GET", "/api/mail/accounts/"+pathSeg(id), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetAccountOrganization — Which organisation a shared mailbox is shared with, and which it could be.
//
// Only the mailbox's owner; anybody else gets 404.
//
// GET /api/mail/accounts/{accountId}/organization
//
// accountID: The mailbox (mail account).
func (r *MailService) GetAccountOrganization(ctx context.Context, accountID string, opts ...CallOption) (*MailGetAccountOrganizationResponseData, error) {
	var out *MailGetAccountOrganizationResponseData
	err := r.c.call(ctx, "GET", "/api/mail/accounts/"+pathSeg(accountID)+"/organization", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetAutoResponse — A mailbox's out-of-office reply and forwarding.
//
// Only the mailbox's owner; anybody else gets 404.
//
// GET /api/mail/accounts/{accountId}/auto-response
//
// accountID: The mailbox (mail account).
func (r *MailService) GetAutoResponse(ctx context.Context, accountID string, opts ...CallOption) (*MailAutoResponse, error) {
	var out *MailAutoResponse
	err := r.c.call(ctx, "GET", "/api/mail/accounts/"+pathSeg(accountID)+"/auto-response", nil, nil, false, "data.settings", &out, opts)
	return out, err
}

// GetContact — What the caller's mail says about one address: name, first and last mail, counts both ways and reply rate.
//
// GET /api/mail/contacts/{address}
//
// address: The e-mail address, URL-encoded.
func (r *MailService) GetContact(ctx context.Context, address string, opts ...CallOption) (*MailGetContactResponseData, error) {
	var out *MailGetContactResponseData
	err := r.c.call(ctx, "GET", "/api/mail/contacts/"+pathSeg(address), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetOpens — How often a sent message with a read receipt was opened, and when first and last.
//
// GET /api/mail/emails/{emailId}/opens
//
// emailID: The message.
func (r *MailService) GetOpens(ctx context.Context, emailID string, opts ...CallOption) (*MailGetOpensResponseData, error) {
	var out *MailGetOpensResponseData
	err := r.c.call(ctx, "GET", "/api/mail/emails/"+pathSeg(emailID)+"/opens", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetPlatformAddress — The caller's mailbox on the platform's own domain, or null.
//
// GET /api/mail/platform-address
func (r *MailService) GetPlatformAddress(ctx context.Context, opts ...CallOption) (*MailAccount, error) {
	var out *MailAccount
	err := r.c.call(ctx, "GET", "/api/mail/platform-address", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetRaw — Downloads a message's original source from the mail server, as text or as an .eml file.
//
// 404 `mailRawNotOnServer` when the server no longer has it, 413 `mailRawTooLarge`, 502 `mailRawUnreachable`.
//
// GET /api/mail/emails/{emailId}/raw
//
// emailID: The message.
func (r *MailService) GetRaw(ctx context.Context, emailID string, query *MailGetRawQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/mail/emails/"+pathSeg(emailID)+"/raw", query, nil, false, "", &out, opts)
	return out, err
}

// GetRetention — How many days a mailbox's mail is kept here.
//
// Only the mailbox's owner; anybody else gets 404. 0 keeps it for ever.
//
// GET /api/mail/accounts/{accountId}/retention
//
// accountID: The mailbox (mail account).
func (r *MailService) GetRetention(ctx context.Context, accountID string, opts ...CallOption) (*MailGetRetentionResponseData, error) {
	var out *MailGetRetentionResponseData
	err := r.c.call(ctx, "GET", "/api/mail/accounts/"+pathSeg(accountID)+"/retention", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetThread — One conversation with its messages, oldest first, and their attachments; marks it read.
//
// Each message carries `safety` (SPF/DKIM/DMARC and the sender checks) and, on received mail, a `verificationCode` found in it. Inline images are embedded in `htmlBody`.
//
// GET /api/mail/threads/{threadId}
//
// threadID: The conversation.
func (r *MailService) GetThread(ctx context.Context, threadID string, opts ...CallOption) (*MailGetThreadResponseData, error) {
	var out *MailGetThreadResponseData
	err := r.c.call(ctx, "GET", "/api/mail/threads/"+pathSeg(threadID), nil, nil, false, "data", &out, opts)
	return out, err
}

// ListAccounts — Lists the mailboxes the caller can open: their own and those shared with them.
//
// GET /api/mail/accounts
func (r *MailService) ListAccounts(ctx context.Context, opts ...CallOption) ([]MailListAccountsResponseDataItem, error) {
	var out []MailListAccountsResponseDataItem
	err := r.c.call(ctx, "GET", "/api/mail/accounts", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListAliases — Lists the addresses a mailbox may send as.
//
// Only the mailbox's owner; anybody else gets 404.
//
// GET /api/mail/accounts/{accountId}/aliases
//
// accountID: The mailbox (mail account).
func (r *MailService) ListAliases(ctx context.Context, accountID string, opts ...CallOption) ([]MailAlias, error) {
	var out []MailAlias
	err := r.c.call(ctx, "GET", "/api/mail/accounts/"+pathSeg(accountID)+"/aliases", nil, nil, false, "data.aliases", &out, opts)
	return out, err
}

// ListBlockedSenders — Lists the caller's blocked senders, oldest first.
//
// GET /api/mail/blocked-senders
func (r *MailService) ListBlockedSenders(ctx context.Context, opts ...CallOption) ([]MailListBlockedSendersResponseDataSendersItem, error) {
	var out []MailListBlockedSendersResponseDataSendersItem
	err := r.c.call(ctx, "GET", "/api/mail/blocked-senders", nil, nil, false, "data.senders", &out, opts)
	return out, err
}

// ListDrafts — Lists the caller's drafts, last changed first.
//
// GET /api/mail/drafts
func (r *MailService) ListDrafts(ctx context.Context, opts ...CallOption) ([]MailDraft, error) {
	var out []MailDraft
	err := r.c.call(ctx, "GET", "/api/mail/drafts", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListFolders — Lists a mailbox's folders on its server, and which of them are synced.
//
// Only the mailbox's owner; anybody else gets 404. Asks the mail server, so it fails (500 `actionFailed`) when the server cannot be reached.
//
// GET /api/mail/accounts/{accountId}/folders
//
// accountID: The mailbox (mail account).
func (r *MailService) ListFolders(ctx context.Context, accountID string, opts ...CallOption) (*MailListFoldersResponseData, error) {
	var out *MailListFoldersResponseData
	err := r.c.call(ctx, "GET", "/api/mail/accounts/"+pathSeg(accountID)+"/folders", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListLabels — Lists the caller's labels.
//
// GET /api/mail/labels
func (r *MailService) ListLabels(ctx context.Context, opts ...CallOption) ([]MailLabel, error) {
	var out []MailLabel
	err := r.c.call(ctx, "GET", "/api/mail/labels", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListNewsletters — Lists the newsletters in the inboxes and archives the caller can open, newest first.
//
// GET /api/mail/newsletters
func (r *MailService) ListNewsletters(ctx context.Context, query *MailListNewslettersQuery, opts ...CallOption) (*MailListNewslettersResponseData, error) {
	var out *MailListNewslettersResponseData
	err := r.c.call(ctx, "GET", "/api/mail/newsletters", query, nil, false, "data", &out, opts)
	return out, err
}

// ListPgpKeys — Lists the caller's PGP keys: their own and their contacts'.
//
// Needs PGP on the plan. Private keys are never answered.
//
// GET /api/mail/pgp/keys
func (r *MailService) ListPgpKeys(ctx context.Context, opts ...CallOption) ([]MailPgpKey, error) {
	var out []MailPgpKey
	err := r.c.call(ctx, "GET", "/api/mail/pgp/keys", nil, nil, false, "data.keys", &out, opts)
	return out, err
}

// ListRules — Lists the caller's rules in the order they run.
//
// GET /api/mail/rules
func (r *MailService) ListRules(ctx context.Context, opts ...CallOption) ([]MailRule, error) {
	var out []MailRule
	err := r.c.call(ctx, "GET", "/api/mail/rules", nil, nil, false, "data.rules", &out, opts)
	return out, err
}

// ListScheduled — Lists the caller's queued sends, latest send time first.
//
// The newest 200, in every status.
//
// GET /api/mail/scheduled
func (r *MailService) ListScheduled(ctx context.Context, opts ...CallOption) ([]MailScheduledSend, error) {
	var out []MailScheduledSend
	err := r.c.call(ctx, "GET", "/api/mail/scheduled", nil, nil, false, "data.scheduled", &out, opts)
	return out, err
}

// ListSignatures — Lists the caller's signatures.
//
// GET /api/mail/signatures
func (r *MailService) ListSignatures(ctx context.Context, opts ...CallOption) ([]MailSignature, error) {
	var out []MailSignature
	err := r.c.call(ctx, "GET", "/api/mail/signatures", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListSnippets — Lists the caller's snippets (canned replies), by name.
//
// GET /api/mail/snippets
func (r *MailService) ListSnippets(ctx context.Context, opts ...CallOption) ([]MailSnippet, error) {
	var out []MailSnippet
	err := r.c.call(ctx, "GET", "/api/mail/snippets", nil, nil, false, "data.snippets", &out, opts)
	return out, err
}

// ListSubaddresses — Lists the `+tag` addresses of the caller's mailboxes that received mail, with who used them.
//
// GET /api/mail/subaddresses
func (r *MailService) ListSubaddresses(ctx context.Context, opts ...CallOption) ([]MailListSubaddressesResponseDataSubaddressesItem, error) {
	var out []MailListSubaddressesResponseDataSubaddressesItem
	err := r.c.call(ctx, "GET", "/api/mail/subaddresses", nil, nil, false, "data.subaddresses", &out, opts)
	return out, err
}

// ListTeammates — The people a conversation of this mailbox can be assigned to.
//
// A shared inbox: every member of the organisation it is shared with. Otherwise only its owner.
//
// GET /api/mail/accounts/{accountId}/teammates
//
// accountID: The mailbox (mail account).
func (r *MailService) ListTeammates(ctx context.Context, accountID string, opts ...CallOption) ([]MailTeammate, error) {
	var out []MailTeammate
	err := r.c.call(ctx, "GET", "/api/mail/accounts/"+pathSeg(accountID)+"/teammates", nil, nil, false, "data.teammates", &out, opts)
	return out, err
}

// ListThreadNotes — Lists a conversation's internal notes, oldest first.
//
// Notes are for teammates and never sent to the customer.
//
// GET /api/mail/threads/{threadId}/notes
//
// threadID: The conversation.
func (r *MailService) ListThreadNotes(ctx context.Context, threadID string, opts ...CallOption) ([]MailThreadNote, error) {
	var out []MailThreadNote
	err := r.c.call(ctx, "GET", "/api/mail/threads/"+pathSeg(threadID)+"/notes", nil, nil, false, "data.notes", &out, opts)
	return out, err
}

// ListThreads — Lists one page of conversations of a folder, newest first, across the caller's mailboxes or of one.
//
// `view=flat` lists single messages instead (each with `_actualThreadId`); assignee and status filters apply to conversations only.
//
// GET /api/mail/threads
func (r *MailService) ListThreads(ctx context.Context, query *MailListThreadsQuery, opts ...CallOption) (*MailListThreadsResponseData, error) {
	var out *MailListThreadsResponseData
	err := r.c.call(ctx, "GET", "/api/mail/threads", query, nil, false, "data", &out, opts)
	return out, err
}

// LookupPgpKeys — Says which of the given addresses the caller holds a public key for.
//
// Needs PGP on the plan. 30 a minute.
//
// POST /api/mail/pgp/lookup
func (r *MailService) LookupPgpKeys(ctx context.Context, body *MailLookupPgpKeysBody, opts ...CallOption) (map[string]MailLookupPgpKeysResponseDataKeysValue, error) {
	var out map[string]MailLookupPgpKeysResponseDataKeysValue
	err := r.c.call(ctx, "POST", "/api/mail/pgp/lookup", nil, body, false, "data.keys", &out, opts)
	return out, err
}

// MarkAllRead — Marks every conversation of a view read, here and on the mail server.
//
// POST /api/mail/threads/mark-all-read
func (r *MailService) MarkAllRead(ctx context.Context, body *MailMarkAllReadBody, opts ...CallOption) (*MailMarkAllReadResponseData, error) {
	var out *MailMarkAllReadResponseData
	err := r.c.call(ctx, "POST", "/api/mail/threads/mark-all-read", nil, body, false, "data", &out, opts)
	return out, err
}

// ParseInvite — Reads the appointments of a calendar invitation attached to a message.
//
// Each event says whether it is already in one of the caller's calendars (`match`); `calendars` are those it could be added to. 400 when the attachment is not an .ics file.
//
// GET /api/mail/attachments/{id}/parse-ics
//
// id: The attachment.
func (r *MailService) ParseInvite(ctx context.Context, id string, opts ...CallOption) (*MailParseInviteResponseData, error) {
	var out *MailParseInviteResponseData
	err := r.c.call(ctx, "GET", "/api/mail/attachments/"+pathSeg(id)+"/parse-ics", nil, nil, false, "data", &out, opts)
	return out, err
}

// PatchThread — Changes a conversation: read, starred, folder, labels, assignee, status or desk section.
//
// Read and starred apply to its messages and to the mail server too; a new folder moves its messages there on the server. Assigning notifies the new assignee; resolving completes the tasks waiting on it.
//
// PATCH /api/mail/threads/{threadId}
//
// threadID: The conversation.
func (r *MailService) PatchThread(ctx context.Context, threadID string, body *MailPatchThreadBody, opts ...CallOption) (*MailThread, error) {
	var out *MailThread
	err := r.c.call(ctx, "PATCH", "/api/mail/threads/"+pathSeg(threadID), nil, body, false, "data", &out, opts)
	return out, err
}

// RenameFolder — Renames or moves a folder on the mailbox's server.
//
// Only the mailbox's owner; anybody else gets 404.
//
// PATCH /api/mail/accounts/{accountId}/folders
//
// accountID: The mailbox (mail account).
func (r *MailService) RenameFolder(ctx context.Context, accountID string, body *MailRenameFolderBody, opts ...CallOption) (*MailRenameFolderResponseData, error) {
	var out *MailRenameFolderResponseData
	err := r.c.call(ctx, "PATCH", "/api/mail/accounts/"+pathSeg(accountID)+"/folders", nil, body, false, "data", &out, opts)
	return out, err
}

// Reply — Replies to a message now, from its mailbox, in its conversation.
//
// To its sender; with `replyAll` the other recipients in Cc. The subject gets `Re:`.
//
// POST /api/mail/emails/{emailId}/reply
//
// emailID: The message.
func (r *MailService) Reply(ctx context.Context, emailID string, body *MailReplyBody, opts ...CallOption) (*MailReplyResponseData, error) {
	var out *MailReplyResponseData
	err := r.c.call(ctx, "POST", "/api/mail/emails/"+pathSeg(emailID)+"/reply", nil, body, false, "data", &out, opts)
	return out, err
}

// ResendAliasConfirmation — Mails a send-as address's confirmation link again.
//
// Only the mailbox's owner; anybody else gets 404.
//
// POST /api/mail/accounts/{accountId}/aliases/{aliasId}
//
// accountID: The mailbox (mail account).
// aliasID: The send-as address.
func (r *MailService) ResendAliasConfirmation(ctx context.Context, accountID string, aliasID string, opts ...CallOption) (*MailAlias, error) {
	var out *MailAlias
	err := r.c.call(ctx, "POST", "/api/mail/accounts/"+pathSeg(accountID)+"/aliases/"+pathSeg(aliasID), nil, nil, false, "data.alias", &out, opts)
	return out, err
}

// Rethread — Merges a mailbox's conversations that were split although they share a subject.
//
// Only the mailbox's owner; anybody else gets 404. Each group of conversations with the same normalised subject becomes its oldest one.
//
// POST /api/mail/accounts/{accountId}/rethread
//
// accountID: The mailbox (mail account).
func (r *MailService) Rethread(ctx context.Context, accountID string, opts ...CallOption) (*MailRethreadResponseData, error) {
	var out *MailRethreadResponseData
	err := r.c.call(ctx, "POST", "/api/mail/accounts/"+pathSeg(accountID)+"/rethread", nil, nil, false, "data", &out, opts)
	return out, err
}

// Rsvp — Answers a calendar invitation attached to a message: accepted, declined or tentative.
//
// Mails the organiser a calendar reply from the message's mailbox.
//
// POST /api/mail/attachments/{id}/rsvp
//
// id: The attachment.
func (r *MailService) Rsvp(ctx context.Context, id string, body *MailRsvpBody, opts ...CallOption) (*MailRsvpResponseData, error) {
	var out *MailRsvpResponseData
	err := r.c.call(ctx, "POST", "/api/mail/attachments/"+pathSeg(id)+"/rsvp", nil, body, false, "data", &out, opts)
	return out, err
}

// SaveAttachmentToAssets — Saves an attachment as a file in Assets.
//
// Saving the same attachment again answers the existing file (`created: false`).
//
// POST /api/mail/attachments/{id}/save-to-assets
//
// id: The attachment.
func (r *MailService) SaveAttachmentToAssets(ctx context.Context, id string, body *MailSaveAttachmentToAssetsBody, opts ...CallOption) (*MailSaveAttachmentToAssetsResponseData, error) {
	var out *MailSaveAttachmentToAssetsResponseData
	err := r.c.call(ctx, "POST", "/api/mail/attachments/"+pathSeg(id)+"/save-to-assets", nil, body, false, "data", &out, opts)
	return out, err
}

// Schedule — Queues a send: in ten seconds with a chance to undo, or at a chosen time.
//
// `kind: scheduled` needs `sendAt` between 30 seconds and 365 days ahead. The queue sends with the same rules as POST /api/mail/send. 60 a minute.
//
// POST /api/mail/scheduled
func (r *MailService) Schedule(ctx context.Context, body *MailScheduleBody, opts ...CallOption) (*MailScheduleResponseData, error) {
	var out *MailScheduleResponseData
	err := r.c.call(ctx, "POST", "/api/mail/scheduled", nil, body, false, "data", &out, opts)
	return out, err
}

// Search — Searches the messages of every mailbox the caller can open, best matches first.
//
// `q` takes web-search syntax: quoted phrases, `-excluded`, `OR`. With at least one filter `q` may be empty, and the newest matches come first.
//
// GET /api/mail/search
func (r *MailService) Search(ctx context.Context, query *MailSearchQuery, opts ...CallOption) ([]MailEmail, error) {
	var out []MailEmail
	err := r.c.call(ctx, "GET", "/api/mail/search", query, nil, false, "data.emails", &out, opts)
	return out, err
}

// SearchContacts — Suggests addresses the caller has corresponded with, most frequent first.
//
// GET /api/mail/contacts/search
func (r *MailService) SearchContacts(ctx context.Context, query *MailSearchContactsQuery, opts ...CallOption) ([]MailSearchContactsResponseDataContactsItem, error) {
	var out []MailSearchContactsResponseDataContactsItem
	err := r.c.call(ctx, "GET", "/api/mail/contacts/search", query, nil, false, "data.contacts", &out, opts)
	return out, err
}

// Send — Sends a message now and waits for the mail server.
//
// From a mailbox the caller can send from. The header `Idempotency-Key` makes a retry answer the first result for 24 hours. 60 a minute. Documents (`documentAttachments`) are attached only by a queued send (POST /api/mail/scheduled).
//
// POST /api/mail/send
func (r *MailService) Send(ctx context.Context, body *MailSendBody, opts ...CallOption) (*MailSendResponseData, error) {
	var out *MailSendResponseData
	err := r.c.call(ctx, "POST", "/api/mail/send", nil, body, false, "data", &out, opts)
	return out, err
}

// SetAccountOrganization — Chooses the organisation a shared mailbox is shared with.
//
// Only the mailbox's owner; anybody else gets 404. One of the owner's own organisations (403 otherwise), with room for the mailbox under its plan.
//
// PUT /api/mail/accounts/{accountId}/organization
//
// accountID: The mailbox (mail account).
func (r *MailService) SetAccountOrganization(ctx context.Context, accountID string, body *MailSetAccountOrganizationBody, opts ...CallOption) (*MailSetAccountOrganizationResponseData, error) {
	var out *MailSetAccountOrganizationResponseData
	err := r.c.call(ctx, "PUT", "/api/mail/accounts/"+pathSeg(accountID)+"/organization", nil, body, false, "data", &out, opts)
	return out, err
}

// SetAutoResponse — Sets a mailbox's out-of-office reply and forwarding, whole.
//
// Only the mailbox's owner; anybody else gets 404. An enabled reply needs a subject and a message, forwarding an address other than the mailbox's own.
//
// PUT /api/mail/accounts/{accountId}/auto-response
//
// accountID: The mailbox (mail account).
func (r *MailService) SetAutoResponse(ctx context.Context, accountID string, body *MailSetAutoResponseBody, opts ...CallOption) (*MailAutoResponse, error) {
	var out *MailAutoResponse
	err := r.c.call(ctx, "PUT", "/api/mail/accounts/"+pathSeg(accountID)+"/auto-response", nil, body, false, "data.settings", &out, opts)
	return out, err
}

// SetCRMSync — Shows a mailbox's conversations in the CRM, or stops.
//
// Only the mailbox's owner; anybody else gets 404. Nothing is copied: in the CRM a conversation is shown only to those who may open the mailbox.
//
// PUT /api/mail/accounts/{accountId}/crm
//
// accountID: The mailbox (mail account).
func (r *MailService) SetCRMSync(ctx context.Context, accountID string, body *MailSetCrmSyncBody, opts ...CallOption) (*MailSetCrmSyncResponseData, error) {
	var out *MailSetCrmSyncResponseData
	err := r.c.call(ctx, "PUT", "/api/mail/accounts/"+pathSeg(accountID)+"/crm", nil, body, false, "data", &out, opts)
	return out, err
}

// SetFlags — Changes one message: read, starred, archived, trashed, folder or labels.
//
// Read and starred are set on the mail server too, and the conversation's unread count follows.
//
// PATCH /api/mail/emails/{emailId}/flags
//
// emailID: The message.
func (r *MailService) SetFlags(ctx context.Context, emailID string, body *MailSetFlagsBody, opts ...CallOption) (*MailEmail, error) {
	var out *MailEmail
	err := r.c.call(ctx, "PATCH", "/api/mail/emails/"+pathSeg(emailID)+"/flags", nil, body, false, "data", &out, opts)
	return out, err
}

// SetFolders — Chooses which of a mailbox's folders are synced.
//
// Only the mailbox's owner; anybody else gets 404.
//
// PUT /api/mail/accounts/{accountId}/folders
//
// accountID: The mailbox (mail account).
func (r *MailService) SetFolders(ctx context.Context, accountID string, body *MailSetFoldersBody, opts ...CallOption) (*MailSetFoldersResponseData, error) {
	var out *MailSetFoldersResponseData
	err := r.c.call(ctx, "PUT", "/api/mail/accounts/"+pathSeg(accountID)+"/folders", nil, body, false, "data", &out, opts)
	return out, err
}

// SetRetention — Sets how many days a mailbox's mail is kept here.
//
// Only the mailbox's owner; anybody else gets 404. 0 keeps it for ever.
//
// PATCH /api/mail/accounts/{accountId}/retention
//
// accountID: The mailbox (mail account).
func (r *MailService) SetRetention(ctx context.Context, accountID string, body *MailSetRetentionBody, opts ...CallOption) (*MailSetRetentionResponseData, error) {
	var out *MailSetRetentionResponseData
	err := r.c.call(ctx, "PATCH", "/api/mail/accounts/"+pathSeg(accountID)+"/retention", nil, body, false, "data", &out, opts)
	return out, err
}

// SnoozeThread — Hides a conversation until a time, then brings it back unread to its folder.
//
// POST /api/mail/threads/{threadId}/snooze
//
// threadID: The conversation.
func (r *MailService) SnoozeThread(ctx context.Context, threadID string, body *MailSnoozeThreadBody, opts ...CallOption) (*MailSnoozeThreadResponseData, error) {
	var out *MailSnoozeThreadResponseData
	err := r.c.call(ctx, "POST", "/api/mail/threads/"+pathSeg(threadID)+"/snooze", nil, body, false, "data", &out, opts)
	return out, err
}

// Storage — Where the caller's mail storage stands against the plan, and that of each organisation sharing a mailbox with them.
//
// `limitMb` -1 is unlimited. A full pool refuses new uploads; incoming mail still arrives.
//
// GET /api/mail/storage
func (r *MailService) Storage(ctx context.Context, opts ...CallOption) (*MailStorageResponseData, error) {
	var out *MailStorageResponseData
	err := r.c.call(ctx, "GET", "/api/mail/storage", nil, nil, false, "data", &out, opts)
	return out, err
}

// SyncAccount — Starts a sync of a mailbox now, or asks for one folder to be synced.
//
// Without `folder`: the owner only, a full sync in the background (400 when sync is off). With `folder`: anybody who can open the mailbox; `syncing` says whether a sync was queued.
//
// POST /api/mail/accounts/{id}/sync
//
// id: The mailbox (mail account).
func (r *MailService) SyncAccount(ctx context.Context, id string, body *MailSyncAccountBody, opts ...CallOption) (*MailSyncAccountResponseData, error) {
	var out *MailSyncAccountResponseData
	err := r.c.call(ctx, "POST", "/api/mail/accounts/"+pathSeg(id)+"/sync", nil, body, false, "data", &out, opts)
	return out, err
}

// ThreadDocumentHTML — Renders a conversation as the HTML of one document page, for saving it as a document.
//
// It only reads: the page is created with the documents API. The dates follow the caller's regional settings. 413 when the conversation is too long for one page.
//
// POST /api/mail/threads/{threadId}/document-html
//
// threadID: The conversation.
func (r *MailService) ThreadDocumentHTML(ctx context.Context, threadID string, body *MailThreadDocumentHtmlBody, opts ...CallOption) (*MailThreadDocumentHtmlResponseData, error) {
	var out *MailThreadDocumentHtmlResponseData
	err := r.c.call(ctx, "POST", "/api/mail/threads/"+pathSeg(threadID)+"/document-html", nil, body, false, "data", &out, opts)
	return out, err
}

// UnblockSender — Unblocks a sender; mail already in spam stays there.
//
// DELETE /api/mail/blocked-senders/{id}
//
// id: The blocked sender entry.
func (r *MailService) UnblockSender(ctx context.Context, id string, opts ...CallOption) (*MailUnblockSenderResponse, error) {
	var out *MailUnblockSenderResponse
	err := r.c.call(ctx, "DELETE", "/api/mail/blocked-senders/"+pathSeg(id), nil, nil, false, "", &out, opts)
	return out, err
}

// UnsnoozeThread — Ends a conversation's snooze and puts it back in its folder.
//
// DELETE /api/mail/threads/{threadId}/snooze
//
// threadID: The conversation.
func (r *MailService) UnsnoozeThread(ctx context.Context, threadID string, opts ...CallOption) (*MailUnsnoozeThreadResponse, error) {
	var out *MailUnsnoozeThreadResponse
	err := r.c.call(ctx, "DELETE", "/api/mail/threads/"+pathSeg(threadID)+"/snooze", nil, nil, false, "", &out, opts)
	return out, err
}

// Unsubscribe — Unsubscribes from the list a message came from, in one click where the sender allows it.
//
// When the sender offers one-click unsubscribe over https it is done here (`unsubscribed: true`); otherwise the answer carries the page or address to use (`method: manual` or `link`). 404 when the message offers no way.
//
// POST /api/mail/emails/{emailId}/unsubscribe
//
// emailID: The message.
func (r *MailService) Unsubscribe(ctx context.Context, emailID string, opts ...CallOption) (*MailUnsubscribeResponseData, error) {
	var out *MailUnsubscribeResponseData
	err := r.c.call(ctx, "POST", "/api/mail/emails/"+pathSeg(emailID)+"/unsubscribe", nil, nil, false, "data", &out, opts)
	return out, err
}

// UpdateAccount — Changes a mailbox's name, sender name, sharing, sync or servers and login.
//
// Only the mailbox's owner; anybody else gets 404. Only the fields sent change. A changed IMAP login is tried once before it is saved (400 with the reason when it fails); an empty password keeps the stored one. Sharing with the organisation needs the team inbox feature and room in the pool it moves to.
//
// PUT /api/mail/accounts/{id}
//
// id: The mailbox (mail account).
func (r *MailService) UpdateAccount(ctx context.Context, id string, body *MailUpdateAccountBody, opts ...CallOption) (*MailAccount, error) {
	var out *MailAccount
	err := r.c.call(ctx, "PUT", "/api/mail/accounts/"+pathSeg(id), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateDraft — Changes a draft of the caller's.
//
// Only the fields sent change.
//
// PUT /api/mail/drafts/{id}
//
// id: The draft.
func (r *MailService) UpdateDraft(ctx context.Context, id string, body *MailUpdateDraftBody, opts ...CallOption) (*MailDraft, error) {
	var out *MailDraft
	err := r.c.call(ctx, "PUT", "/api/mail/drafts/"+pathSeg(id), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateLabel — Renames or recolours a label.
//
// A system label cannot be changed (400).
//
// PUT /api/mail/labels/{id}
//
// id: The label.
func (r *MailService) UpdateLabel(ctx context.Context, id string, body *MailUpdateLabelBody, opts ...CallOption) (*MailLabel, error) {
	var out *MailLabel
	err := r.c.call(ctx, "PUT", "/api/mail/labels/"+pathSeg(id), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdatePgpKey — Renames a PGP key or makes it the default.
//
// Needs PGP on the plan and `mail:manage_pgp_keys`.
//
// PATCH /api/mail/pgp/keys/{id}
//
// id: The key.
func (r *MailService) UpdatePgpKey(ctx context.Context, id string, body *MailUpdatePgpKeyBody, opts ...CallOption) (*MailUpdatePgpKeyResponseDataKey, error) {
	var out *MailUpdatePgpKeyResponseDataKey
	err := r.c.call(ctx, "PATCH", "/api/mail/pgp/keys/"+pathSeg(id), nil, body, false, "data.key", &out, opts)
	return out, err
}

// UpdateRule — Changes a rule.
//
// PATCH /api/mail/rules/{id}
//
// id: The rule.
func (r *MailService) UpdateRule(ctx context.Context, id string, body *MailUpdateRuleBody, opts ...CallOption) (*MailRule, error) {
	var out *MailRule
	err := r.c.call(ctx, "PATCH", "/api/mail/rules/"+pathSeg(id), nil, body, false, "data.rule", &out, opts)
	return out, err
}

// UpdateSignature — Changes a signature.
//
// Only the fields sent change; the HTML is sanitised.
//
// PUT /api/mail/signatures/{id}
//
// id: The signature.
func (r *MailService) UpdateSignature(ctx context.Context, id string, body *MailUpdateSignatureBody, opts ...CallOption) (*MailSignature, error) {
	var out *MailSignature
	err := r.c.call(ctx, "PUT", "/api/mail/signatures/"+pathSeg(id), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateSnippet — Changes a snippet.
//
// PATCH /api/mail/snippets/{id}
//
// id: The snippet.
func (r *MailService) UpdateSnippet(ctx context.Context, id string, body *MailUpdateSnippetBody, opts ...CallOption) (*MailSnippet, error) {
	var out *MailSnippet
	err := r.c.call(ctx, "PATCH", "/api/mail/snippets/"+pathSeg(id), nil, body, false, "data.snippet", &out, opts)
	return out, err
}

// UploadAttachment — Uploads a file to attach to a message; send its id in `attachmentIds`.
//
// 25 MB at most; executable types are refused (415). Counts against the plan's mail storage.
//
// POST /api/mail/attachments/upload
func (r *MailService) UploadAttachment(ctx context.Context, body *MailUploadAttachmentBody, opts ...CallOption) (*MailUploadAttachmentResponseData, error) {
	var out *MailUploadAttachmentResponseData
	err := r.c.call(ctx, "POST", "/api/mail/attachments/upload", nil, body, true, "data", &out, opts)
	return out, err
}

// MailAiService — mail.ai
type MailAiService struct {
	c *Client
}

func newMailAiService(c *Client) *MailAiService {
	return &MailAiService{c: c}
}

// DraftReply — Drafts a reply to a conversation with AI; nothing is sent.
//
// Needs AI writing help on the plan, counts against the organisation's AI allowance (429 when used up) and is refused to a person who opted out of AI. Reads the last four messages of a conversation the caller can open.
//
// POST /api/mail/ai/draft-reply
func (r *MailAiService) DraftReply(ctx context.Context, body *MailAiDraftReplyBody, opts ...CallOption) (*MailAiDraftReplyResponseData, error) {
	var out *MailAiDraftReplyResponseData
	err := r.c.call(ctx, "POST", "/api/mail/ai/draft-reply", nil, body, false, "data", &out, opts)
	return out, err
}

// Rewrite — Rewrites a message text more formal, friendlier, shorter, longer or as instructed.
//
// Needs AI writing help on the plan, counts against the organisation's AI allowance (429 when used up) and is refused to a person who opted out of AI.
//
// POST /api/mail/ai/rewrite
func (r *MailAiService) Rewrite(ctx context.Context, body *MailAiRewriteBody, opts ...CallOption) (string, error) {
	var out string
	err := r.c.call(ctx, "POST", "/api/mail/ai/rewrite", nil, body, false, "data.body", &out, opts)
	return out, err
}

// Subject — Suggests a subject line for a message text.
//
// Needs AI writing help on the plan, counts against the organisation's AI allowance (429 when used up) and is refused to a person who opted out of AI.
//
// POST /api/mail/ai/subject
func (r *MailAiService) Subject(ctx context.Context, body *MailAiSubjectBody, opts ...CallOption) (string, error) {
	var out string
	err := r.c.call(ctx, "POST", "/api/mail/ai/subject", nil, body, false, "data.subject", &out, opts)
	return out, err
}

// SummarizeThread — Summarises a conversation with AI: the gist, key points, what is asked and the tone.
//
// Needs AI writing help on the plan, counts against the organisation's AI allowance (429 when used up) and is refused to a person who opted out of AI. A conversation the caller can open.
//
// POST /api/mail/ai/summarize-thread
func (r *MailAiService) SummarizeThread(ctx context.Context, body *MailAiSummarizeThreadBody, opts ...CallOption) (*MailAiSummarizeThreadResponseData, error) {
	var out *MailAiSummarizeThreadResponseData
	err := r.c.call(ctx, "POST", "/api/mail/ai/summarize-thread", nil, body, false, "data", &out, opts)
	return out, err
}

// MailImapService — mail.imap
type MailImapService struct {
	c *Client
}

func newMailImapService(c *Client) *MailImapService {
	return &MailImapService{c: c}
}

// Action — Marks a message of a saved IMAP connection read or unread, or deletes it.
//
// Not in the usual envelope. The connection must be the caller's own.
//
// POST /api/imap/action
func (r *MailImapService) Action(ctx context.Context, body *MailImapActionBody, opts ...CallOption) (*MailImapActionResponse, error) {
	var out *MailImapActionResponse
	err := r.c.call(ctx, "POST", "/api/imap/action", nil, body, false, "", &out, opts)
	return out, err
}

// Fetch — Reads messages from the mailbox of a saved IMAP connection.
//
// Not in the usual envelope. The connection must be the caller's own (404 `connectionNotFound`); its host must be public (`imapHostRefused`).
//
// POST /api/imap/fetch
func (r *MailImapService) Fetch(ctx context.Context, body *MailImapFetchBody, opts ...CallOption) ([]MailImapFetchResponseEmailsItem, error) {
	var out []MailImapFetchResponseEmailsItem
	err := r.c.call(ctx, "POST", "/api/imap/fetch", nil, body, false, "emails", &out, opts)
	return out, err
}

// Verify — Tries an IMAP login, before a mailbox is connected with it.
//
// Not in the usual envelope. The host must be public. Ten an account in ten minutes. 500 `imapAuthFailed` for a refused login, `imapConnectFailed` when the server does not answer as one.
//
// POST /api/imap/verify
func (r *MailImapService) Verify(ctx context.Context, body *MailImapVerifyBody, opts ...CallOption) (*MailImapVerifyResponse, error) {
	var out *MailImapVerifyResponse
	err := r.c.call(ctx, "POST", "/api/imap/verify", nil, body, false, "", &out, opts)
	return out, err
}

// MailSmtpService — mail.smtp
type MailSmtpService struct {
	c *Client
}

func newMailSmtpService(c *Client) *MailSmtpService {
	return &MailSmtpService{c: c}
}

// Verify — Tries an SMTP login, before a mailbox is connected with it.
//
// Not in the usual envelope. The host must be public. Ten an account in ten minutes. Port 465 is TLS, 587 STARTTLS.
//
// POST /api/smtp/verify
func (r *MailSmtpService) Verify(ctx context.Context, body *MailSmtpVerifyBody, opts ...CallOption) (*MailSmtpVerifyResponse, error) {
	var out *MailSmtpVerifyResponse
	err := r.c.call(ctx, "POST", "/api/smtp/verify", nil, body, false, "", &out, opts)
	return out, err
}

// MarketplaceService — Marketplace: Templates, workflows, blocks, files and services that sellers offer and organizations buy: browsing the storefront, free installs, orders and purchases, reviews, questions, disputes, and a seller's own shop, sales and payouts. The marketplace can be switched off for the whole platform; while it is, every operation here answers 403 `featureDisabledByAdmin`. Paid purchases happen in the dashboard, in the payment provider's form; a seller's offers are products of the Products hub (`products.marketplace.*`).
type MarketplaceService struct {
	c *Client
}

func newMarketplaceService(c *Client) *MarketplaceService {
	return &MarketplaceService{c: c}
}

// AcceptDelivery — Accepts a service delivery as the buyer, which completes the order.
//
// The seller's payout hold starts now.
//
// POST /api/marketplace/orders/{orderId}/accept
//
// orderID: The marketplace order's id.
func (r *MarketplaceService) AcceptDelivery(ctx context.Context, orderID string, opts ...CallOption) (*MarketplaceOrder, error) {
	var out *MarketplaceOrder
	err := r.c.call(ctx, "POST", "/api/marketplace/orders/"+pathSeg(orderID)+"/accept", nil, nil, false, "data.order", &out, opts)
	return out, err
}

// AddDisputeEvidence — Adds evidence to a dispute.
//
// Needs the dispute center on the platform (`marketplace_disputes`).
//
// POST /api/marketplace/disputes/{disputeId}/evidence
//
// disputeID: The dispute's id.
func (r *MarketplaceService) AddDisputeEvidence(ctx context.Context, disputeID string, body *MarketplaceAddDisputeEvidenceBody, opts ...CallOption) (*MarketplaceAddDisputeEvidenceResponse, error) {
	var out *MarketplaceAddDisputeEvidenceResponse
	err := r.c.call(ctx, "POST", "/api/marketplace/disputes/"+pathSeg(disputeID)+"/evidence", nil, body, false, "", &out, opts)
	return out, err
}

// AddToWishlist — Puts a listing on the caller's wishlist.
//
// POST /api/marketplace/wishlist
func (r *MarketplaceService) AddToWishlist(ctx context.Context, body *MarketplaceAddToWishlistBody, opts ...CallOption) (*MarketplaceAddToWishlistResponse, error) {
	var out *MarketplaceAddToWishlistResponse
	err := r.c.call(ctx, "POST", "/api/marketplace/wishlist", nil, body, false, "", &out, opts)
	return out, err
}

// AnswerQuestion — Answers a question about the caller's listing, which publishes it, or hides an abusive one.
//
// The seller: the shop's owner or a holder of `marketplace:sell` in the listing's organization.
//
// PATCH /api/marketplace/{listingId}/questions
//
// listingID: The listing's id.
func (r *MarketplaceService) AnswerQuestion(ctx context.Context, listingID string, body *MarketplaceAnswerQuestionBody, opts ...CallOption) (*MarketplaceQuestion, error) {
	var out *MarketplaceQuestion
	err := r.c.call(ctx, "PATCH", "/api/marketplace/"+pathSeg(listingID)+"/questions", nil, body, false, "data.question", &out, opts)
	return out, err
}

// AskQuestion — Asks the seller a question about a listing.
//
// Not the listing's own seller. Five questions an hour.
//
// POST /api/marketplace/{listingId}/questions
//
// listingID: The listing's id.
func (r *MarketplaceService) AskQuestion(ctx context.Context, listingID string, body *MarketplaceAskQuestionBody, opts ...CallOption) (*MarketplaceQuestion, error) {
	var out *MarketplaceQuestion
	err := r.c.call(ctx, "POST", "/api/marketplace/"+pathSeg(listingID)+"/questions", nil, body, false, "data.question", &out, opts)
	return out, err
}

// BuiltWith — Lists up to six live sites made with a listing whose owners agreed to be shown.
//
// GET /api/marketplace/{listingId}/built-with
//
// listingID: The listing's id.
func (r *MarketplaceService) BuiltWith(ctx context.Context, listingID string, query *MarketplaceBuiltWithQuery, opts ...CallOption) ([]MarketplaceBuiltWithResponseDataSitesItem, error) {
	var out []MarketplaceBuiltWithResponseDataSitesItem
	err := r.c.call(ctx, "GET", "/api/marketplace/"+pathSeg(listingID)+"/built-with", query, nil, false, "data.sites", &out, opts)
	return out, err
}

// CancelDispute — Withdraws a dispute as its buyer; the order returns to where it was.
//
// Needs the dispute center on the platform (`marketplace_disputes`).
//
// POST /api/marketplace/disputes/{disputeId}/cancel
//
// disputeID: The dispute's id.
func (r *MarketplaceService) CancelDispute(ctx context.Context, disputeID string, opts ...CallOption) (*MarketplaceCancelDisputeResponse, error) {
	var out *MarketplaceCancelDisputeResponse
	err := r.c.call(ctx, "POST", "/api/marketplace/disputes/"+pathSeg(disputeID)+"/cancel", nil, nil, false, "", &out, opts)
	return out, err
}

// CancelSellerOrder — Cancels a service order before delivery as its seller, refunding the buyer.
//
// Only a service that is `ordered` or `in_progress`. The seller's share is not paid out, or comes off the next payout.
//
// POST /api/marketplace/seller/orders/{orderId}/cancel
//
// orderID: The marketplace order's id.
func (r *MarketplaceService) CancelSellerOrder(ctx context.Context, orderID string, body *MarketplaceCancelSellerOrderBody, opts ...CallOption) (*MarketplaceCancelSellerOrderResponseData, error) {
	var out *MarketplaceCancelSellerOrderResponseData
	err := r.c.call(ctx, "POST", "/api/marketplace/seller/orders/"+pathSeg(orderID)+"/cancel", nil, body, false, "data", &out, opts)
	return out, err
}

// CancelSubscription — Cancels a subscription the caller bought at the end of its paid period.
//
// Repeating it answers `alreadyCancelling: true`.
//
// DELETE /api/marketplace/subscriptions/{subscriptionId}
//
// subscriptionID: The order's id of an offer sold by subscription, or the id of a former tier subscription.
func (r *MarketplaceService) CancelSubscription(ctx context.Context, subscriptionID string, opts ...CallOption) (*MarketplaceCancelSubscriptionResponseData, error) {
	var out *MarketplaceCancelSubscriptionResponseData
	err := r.c.call(ctx, "DELETE", "/api/marketplace/subscriptions/"+pathSeg(subscriptionID), nil, nil, false, "data", &out, opts)
	return out, err
}

// Categories — Lists the marketplace's categories in the storefront's order, with how many listings the caller can see in each.
//
// GET /api/marketplace/categories
func (r *MarketplaceService) Categories(ctx context.Context, opts ...CallOption) ([]MarketplaceCategoriesResponseDataCategoriesItem, error) {
	var out []MarketplaceCategoriesResponseDataCategoriesItem
	err := r.c.call(ctx, "GET", "/api/marketplace/categories", nil, nil, false, "data.categories", &out, opts)
	return out, err
}

// CreateCoupon — Creates a coupon for the caller's listings, or for one of them.
//
// 409 when the code exists already.
//
// POST /api/marketplace/coupons
func (r *MarketplaceService) CreateCoupon(ctx context.Context, body *MarketplaceCreateCouponBody, opts ...CallOption) (*MarketplaceCoupon, error) {
	var out *MarketplaceCoupon
	err := r.c.call(ctx, "POST", "/api/marketplace/coupons", nil, body, false, "data.coupon", &out, opts)
	return out, err
}

// CreateReview — Reviews a listing.
//
// One review per person and listing, never of one's own. Five reviews an hour.
//
// POST /api/marketplace/{listingId}/reviews
//
// listingID: The listing's id.
func (r *MarketplaceService) CreateReview(ctx context.Context, listingID string, body *MarketplaceCreateReviewBody, opts ...CallOption) (*MarketplaceReview, error) {
	var out *MarketplaceReview
	err := r.c.call(ctx, "POST", "/api/marketplace/"+pathSeg(listingID)+"/reviews", nil, body, false, "data.review", &out, opts)
	return out, err
}

// CreateSellerListing — Starts a new offer: creates a product of the shop's organization with its marketplace offer filled in.
//
// The offer is finished and submitted for review from the product (`products.marketplace.*`).
//
// POST /api/marketplace/seller/listings
func (r *MarketplaceService) CreateSellerListing(ctx context.Context, body *MarketplaceCreateSellerListingBody, opts ...CallOption) (*MarketplaceCreateSellerListingResponseData, error) {
	var out *MarketplaceCreateSellerListingResponseData
	err := r.c.call(ctx, "POST", "/api/marketplace/seller/listings", nil, body, false, "data", &out, opts)
	return out, err
}

// CreateSellerProfile — Opens a shop: creates the caller's seller profile.
//
// One per person (409 otherwise). The shop's slug is made from the name. Being paid out needs the payout account, set up in the dashboard.
//
// POST /api/marketplace/seller
func (r *MarketplaceService) CreateSellerProfile(ctx context.Context, body *MarketplaceCreateSellerProfileBody, opts ...CallOption) (*MarketplaceSellerProfile, error) {
	var out *MarketplaceSellerProfile
	err := r.c.call(ctx, "POST", "/api/marketplace/seller", nil, body, false, "data.profile", &out, opts)
	return out, err
}

// DeleteCoupon — Deletes one of the caller's coupons.
//
// DELETE /api/marketplace/coupons
func (r *MarketplaceService) DeleteCoupon(ctx context.Context, query *MarketplaceDeleteCouponQuery, opts ...CallOption) (*MarketplaceDeleteCouponResponse, error) {
	var out *MarketplaceDeleteCouponResponse
	err := r.c.call(ctx, "DELETE", "/api/marketplace/coupons", query, nil, false, "", &out, opts)
	return out, err
}

// DeleteReview — Deletes the caller's review of a listing.
//
// DELETE /api/marketplace/{listingId}/reviews
//
// listingID: The listing's id.
func (r *MarketplaceService) DeleteReview(ctx context.Context, listingID string, query *MarketplaceDeleteReviewQuery, opts ...CallOption) (*MarketplaceDeleteReviewResponse, error) {
	var out *MarketplaceDeleteReviewResponse
	err := r.c.call(ctx, "DELETE", "/api/marketplace/"+pathSeg(listingID)+"/reviews", query, nil, false, "", &out, opts)
	return out, err
}

// DeleteSellerListing — Deletes one of the caller's listings that has no open or completed orders.
//
// DELETE /api/marketplace/seller/listings/{listingId}
//
// listingID: The listing's id.
func (r *MarketplaceService) DeleteSellerListing(ctx context.Context, listingID string, opts ...CallOption) (*MarketplaceDeleteSellerListingResponse, error) {
	var out *MarketplaceDeleteSellerListingResponse
	err := r.c.call(ctx, "DELETE", "/api/marketplace/seller/listings/"+pathSeg(listingID), nil, nil, false, "", &out, opts)
	return out, err
}

// DeliverOrder — Delivers a service order as the seller.
//
// The shop's owner, or a member of the selling organization with `products:update`.
//
// POST /api/marketplace/orders/{orderId}/deliver
//
// orderID: The marketplace order's id.
func (r *MarketplaceService) DeliverOrder(ctx context.Context, orderID string, body *MarketplaceDeliverOrderBody, opts ...CallOption) (*MarketplaceOrder, error) {
	var out *MarketplaceOrder
	err := r.c.call(ctx, "POST", "/api/marketplace/orders/"+pathSeg(orderID)+"/deliver", nil, body, false, "data.order", &out, opts)
	return out, err
}

// Download — Lists the files of a listing the caller bought, with their download links.
//
// For a buyer with access, or the listing's seller. Without `versionId`, the newest version's files.
//
// POST /api/marketplace/{listingId}/download
//
// listingID: The listing's id.
func (r *MarketplaceService) Download(ctx context.Context, listingID string, body *MarketplaceDownloadBody, opts ...CallOption) ([]MarketplaceDownloadResponseDataFilesItem, error) {
	var out []MarketplaceDownloadResponseDataFilesItem
	err := r.c.call(ctx, "POST", "/api/marketplace/"+pathSeg(listingID)+"/download", nil, body, false, "data.files", &out, opts)
	return out, err
}

// EscalateDispute — Hands a dispute to the platform's mediator.
//
// Needs the dispute center on the platform (`marketplace_disputes`). Buyer or seller; repeating it does nothing.
//
// POST /api/marketplace/disputes/{disputeId}/escalate
//
// disputeID: The dispute's id.
func (r *MarketplaceService) EscalateDispute(ctx context.Context, disputeID string, opts ...CallOption) (*MarketplaceEscalateDisputeResponse, error) {
	var out *MarketplaceEscalateDisputeResponse
	err := r.c.call(ctx, "POST", "/api/marketplace/disputes/"+pathSeg(disputeID)+"/escalate", nil, nil, false, "", &out, opts)
	return out, err
}

// FollowShop — Follows a shop.
//
// POST /api/marketplace/shop/{slug}/follow
//
// slug: The shop's slug.
func (r *MarketplaceService) FollowShop(ctx context.Context, slug string, opts ...CallOption) (*MarketplaceFollowShopResponseData, error) {
	var out *MarketplaceFollowShopResponseData
	err := r.c.call(ctx, "POST", "/api/marketplace/shop/"+pathSeg(slug)+"/follow", nil, nil, false, "data", &out, opts)
	return out, err
}

// Get — Returns a listing's product page: details, tiers priced for the caller, reviews, the seller's legal pages and what the checkout would charge.
//
// `checkout.vat` tells how VAT would apply to this buyer; when the evidence of the buyer's country conflicts or is missing it answers `ok: false` with `marketplaceCountryConflict` or `marketplaceCountryRequired`, to be settled with `confirmedCountry` or `buyerCountry`.
//
// GET /api/marketplace/{listingId}
//
// listingID: The listing's id.
func (r *MarketplaceService) Get(ctx context.Context, listingID string, query *MarketplaceGetQuery, opts ...CallOption) (*MarketplaceGetResponseData, error) {
	var out *MarketplaceGetResponseData
	err := r.c.call(ctx, "GET", "/api/marketplace/"+pathSeg(listingID), query, nil, false, "data", &out, opts)
	return out, err
}

// GetCollection — Returns a published collection with its listings in the curator's order, each with the curator's note.
//
// GET /api/marketplace/collections/{slug}
//
// slug: The collection's slug.
func (r *MarketplaceService) GetCollection(ctx context.Context, slug string, query *MarketplaceGetCollectionQuery, opts ...CallOption) (*MarketplaceGetCollectionResponseDataCollection, error) {
	var out *MarketplaceGetCollectionResponseDataCollection
	err := r.c.call(ctx, "GET", "/api/marketplace/collections/"+pathSeg(slug), query, nil, false, "data.collection", &out, opts)
	return out, err
}

// GetDispute — Returns a dispute with its messages, evidence, history and order.
//
// Needs the dispute center on the platform (`marketplace_disputes`).
//
// GET /api/marketplace/disputes/{disputeId}
//
// disputeID: The dispute's id.
func (r *MarketplaceService) GetDispute(ctx context.Context, disputeID string, opts ...CallOption) (*MarketplaceGetDisputeResponseData, error) {
	var out *MarketplaceGetDisputeResponseData
	err := r.c.call(ctx, "GET", "/api/marketplace/disputes/"+pathSeg(disputeID), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetFees — Returns the platform's commission and payout terms for sellers.
//
// GET /api/marketplace/fees
func (r *MarketplaceService) GetFees(ctx context.Context, opts ...CallOption) (*MarketplaceGetFeesResponse, error) {
	var out *MarketplaceGetFeesResponse
	err := r.c.call(ctx, "GET", "/api/marketplace/fees", nil, nil, false, "", &out, opts)
	return out, err
}

// GetFile — Downloads one file of a listing the caller bought.
//
// For a buyer with access, or the listing's seller. 60 downloads an hour; a tier's download cap applies (403 when reached).
//
// GET /api/marketplace/{listingId}/files/{fileId}
//
// listingID: The listing's id.
// fileID: The file's id, from the download list.
func (r *MarketplaceService) GetFile(ctx context.Context, listingID string, fileID string, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/marketplace/"+pathSeg(listingID)+"/files/"+pathSeg(fileID), nil, nil, false, "", &out, opts)
	return out, err
}

// GetOrder — Returns one marketplace order of the caller, as buyer or seller.
//
// GET /api/marketplace/orders/{orderId}
//
// orderID: The marketplace order's id.
func (r *MarketplaceService) GetOrder(ctx context.Context, orderID string, opts ...CallOption) (*MarketplaceGetOrderResponseData, error) {
	var out *MarketplaceGetOrderResponseData
	err := r.c.call(ctx, "GET", "/api/marketplace/orders/"+pathSeg(orderID), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetSellerListing — Returns one of the caller's listings with its versions and files.
//
// GET /api/marketplace/seller/listings/{listingId}
//
// listingID: The listing's id.
func (r *MarketplaceService) GetSellerListing(ctx context.Context, listingID string, opts ...CallOption) (*MarketplaceGetSellerListingResponseData, error) {
	var out *MarketplaceGetSellerListingResponseData
	err := r.c.call(ctx, "GET", "/api/marketplace/seller/listings/"+pathSeg(listingID), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetSellerProfile — Returns the caller's seller profile, or null.
//
// GET /api/marketplace/seller
func (r *MarketplaceService) GetSellerProfile(ctx context.Context, opts ...CallOption) (*MarketplaceSellerProfile, error) {
	var out *MarketplaceSellerProfile
	err := r.c.call(ctx, "GET", "/api/marketplace/seller", nil, nil, false, "data.profile", &out, opts)
	return out, err
}

// GetShop — Returns a seller's shop: its header with rating and trader information, its standing and its listings.
//
// The shop's listings are better browsed with GET /api/marketplace?seller={slug}, which filters and pages.
//
// GET /api/marketplace/shop/{slug}
//
// slug: The shop's slug.
func (r *MarketplaceService) GetShop(ctx context.Context, slug string, opts ...CallOption) (*MarketplaceGetShopResponseData, error) {
	var out *MarketplaceGetShopResponseData
	err := r.c.call(ctx, "GET", "/api/marketplace/shop/"+pathSeg(slug), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetShowcase — Lists the projects a purchase installed, whether each is live, and whether the product page may show it.
//
// GET /api/marketplace/purchases/{orderId}/showcase
//
// orderID: The marketplace order's id.
func (r *MarketplaceService) GetShowcase(ctx context.Context, orderID string, opts ...CallOption) (*MarketplaceGetShowcaseResponseData, error) {
	var out *MarketplaceGetShowcaseResponseData
	err := r.c.call(ctx, "GET", "/api/marketplace/purchases/"+pathSeg(orderID)+"/showcase", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetSubscriptions — Lists the caller's subscriptions to sellers' former subscription tiers.
//
// An offer sold by subscription is a purchase; see GET /api/marketplace/purchases.
//
// GET /api/marketplace/subscriptions
func (r *MarketplaceService) GetSubscriptions(ctx context.Context, opts ...CallOption) ([]MarketplaceGetSubscriptionsResponseDataSubscriptionsItem, error) {
	var out []MarketplaceGetSubscriptionsResponseDataSubscriptionsItem
	err := r.c.call(ctx, "GET", "/api/marketplace/subscriptions", nil, nil, false, "data.subscriptions", &out, opts)
	return out, err
}

// HardDeleteSellerListing — Removes one of the caller's listings for good, for an erasure request.
//
// Only a listing that never had an order; past sales are kept for tax and disputes.
//
// DELETE /api/marketplace/seller/listings/{listingId}/hard-delete
//
// listingID: The listing's id.
func (r *MarketplaceService) HardDeleteSellerListing(ctx context.Context, listingID string, opts ...CallOption) (*MarketplaceHardDeleteSellerListingResponse, error) {
	var out *MarketplaceHardDeleteSellerListingResponse
	err := r.c.call(ctx, "DELETE", "/api/marketplace/seller/listings/"+pathSeg(listingID)+"/hard-delete", nil, nil, false, "", &out, opts)
	return out, err
}

// HideFromShowcase — Stops showing a project under "Built with this".
//
// DELETE /api/marketplace/purchases/{orderId}/showcase
//
// orderID: The marketplace order's id.
func (r *MarketplaceService) HideFromShowcase(ctx context.Context, orderID string, query *MarketplaceHideFromShowcaseQuery, opts ...CallOption) (*MarketplaceHideFromShowcaseResponseData, error) {
	var out *MarketplaceHideFromShowcaseResponseData
	err := r.c.call(ctx, "DELETE", "/api/marketplace/purchases/"+pathSeg(orderID)+"/showcase", query, nil, false, "data", &out, opts)
	return out, err
}

// Install — Takes a free listing or a free tier: records an order without a payment.
//
// Installing it into a workspace is POST /api/marketplace/orders/{orderId}/install. A paid listing is bought in the dashboard.
//
// POST /api/marketplace/{listingId}/install
//
// listingID: The listing's id.
func (r *MarketplaceService) Install(ctx context.Context, listingID string, body *MarketplaceInstallBody, opts ...CallOption) (*MarketplaceInstallResponseData, error) {
	var out *MarketplaceInstallResponseData
	err := r.c.call(ctx, "POST", "/api/marketplace/"+pathSeg(listingID)+"/install", nil, body, false, "data", &out, opts)
	return out, err
}

// InstallAll — Installs every item of a bundle purchase into one workspace.
//
// Each item answers `installed`, `skipped` (`download`: files are downloaded; `alreadyInstalled`) or `failed` with the reason.
//
// POST /api/marketplace/orders/{orderId}/install-all
//
// orderID: The marketplace order's id.
func (r *MarketplaceService) InstallAll(ctx context.Context, orderID string, body *MarketplaceInstallAllBody, opts ...CallOption) ([]MarketplaceInstallAllResponseDataResultsItem, error) {
	var out []MarketplaceInstallAllResponseDataResultsItem
	err := r.c.call(ctx, "POST", "/api/marketplace/orders/"+pathSeg(orderID)+"/install-all", nil, body, false, "data.results", &out, opts)
	return out, err
}

// InstallOrder — Installs what an order bought into a workspace of the organization it was bought for.
//
// Only the buyer, while the order gives access, into a workspace where they may create the resource. The newest version the order includes unless `versionId`; installing again makes a new copy. A digital asset answers its download links; a service has nothing to install.
//
// POST /api/marketplace/orders/{orderId}/install
//
// orderID: The marketplace order's id.
func (r *MarketplaceService) InstallOrder(ctx context.Context, orderID string, body *MarketplaceInstallOrderBody, opts ...CallOption) (*MarketplaceInstallResult, error) {
	var out *MarketplaceInstallResult
	err := r.c.call(ctx, "POST", "/api/marketplace/orders/"+pathSeg(orderID)+"/install", nil, body, false, "data", &out, opts)
	return out, err
}

// InstallToWorkspace — Installs a listing the caller owns into a workspace, from their newest order that gives access.
//
// The same as POST /api/marketplace/orders/{orderId}/install with that order. The workspace must belong to the organization the purchase was made for.
//
// POST /api/marketplace/{listingId}/install-to-workspace
//
// listingID: The listing's id.
func (r *MarketplaceService) InstallToWorkspace(ctx context.Context, listingID string, body *MarketplaceInstallToWorkspaceBody, opts ...CallOption) (*MarketplaceInstallToWorkspaceResponseData, error) {
	var out *MarketplaceInstallToWorkspaceResponseData
	err := r.c.call(ctx, "POST", "/api/marketplace/"+pathSeg(listingID)+"/install-to-workspace", nil, body, false, "data", &out, opts)
	return out, err
}

// List — Browses the storefront: approved listings the caller may see, filtered, ranked and paged, with facets.
//
// `relevance` ranks by how well the search matches, sales, rating and recency; featured listings do not rank higher. Unknown filter values are ignored rather than refused.
//
// GET /api/marketplace
func (r *MarketplaceService) List(ctx context.Context, query *MarketplaceListQuery, opts ...CallOption) (*MarketplaceListResponseData, error) {
	var out *MarketplaceListResponseData
	err := r.c.call(ctx, "GET", "/api/marketplace", query, nil, false, "data", &out, opts)
	return out, err
}

// ListCollections — Lists the published collections that hold at least one listing the caller can see.
//
// GET /api/marketplace/collections
func (r *MarketplaceService) ListCollections(ctx context.Context, query *MarketplaceListCollectionsQuery, opts ...CallOption) ([]MarketplaceListCollectionsResponseDataCollectionsItem, error) {
	var out []MarketplaceListCollectionsResponseDataCollectionsItem
	err := r.c.call(ctx, "GET", "/api/marketplace/collections", query, nil, false, "data.collections", &out, opts)
	return out, err
}

// ListCoupons — Lists the caller's coupons as a seller.
//
// GET /api/marketplace/coupons
func (r *MarketplaceService) ListCoupons(ctx context.Context, opts ...CallOption) ([]MarketplaceCoupon, error) {
	var out []MarketplaceCoupon
	err := r.c.call(ctx, "GET", "/api/marketplace/coupons", nil, nil, false, "data.coupons", &out, opts)
	return out, err
}

// ListDisputes — Lists the disputes the caller is part of, as buyer, as seller or both.
//
// Needs the dispute center on the platform (`marketplace_disputes`).
//
// GET /api/marketplace/disputes
func (r *MarketplaceService) ListDisputes(ctx context.Context, query *MarketplaceListDisputesQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/marketplace/disputes", query, nil, false, "data.disputes", &out, opts)
	return out, err
}

// ListLegalPages — Lists the caller's shop legal pages: terms, return policy, privacy policy and imprint.
//
// GET /api/marketplace/seller/legal-pages
func (r *MarketplaceService) ListLegalPages(ctx context.Context, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/marketplace/seller/legal-pages", nil, nil, false, "data.legalPages", &out, opts)
	return out, err
}

// ListOrderDisputes — Lists the disputes of an order the caller bought or sold.
//
// Needs the dispute center on the platform (`marketplace_disputes`).
//
// GET /api/marketplace/orders/{orderId}/disputes
//
// orderID: The marketplace order's id.
func (r *MarketplaceService) ListOrderDisputes(ctx context.Context, orderID string, opts ...CallOption) ([]MarketplaceListOrderDisputesResponseDataDisputesItem, error) {
	var out []MarketplaceListOrderDisputesResponseDataDisputesItem
	err := r.c.call(ctx, "GET", "/api/marketplace/orders/"+pathSeg(orderID)+"/disputes", nil, nil, false, "data.disputes", &out, opts)
	return out, err
}

// ListOrders — Lists the caller's marketplace orders as buyer or as seller, newest first.
//
// GET /api/marketplace/orders
func (r *MarketplaceService) ListOrders(ctx context.Context, query *MarketplaceListOrdersQuery, opts ...CallOption) ([]MarketplaceListOrdersResponseDataOrdersItem, error) {
	var out []MarketplaceListOrdersResponseDataOrdersItem
	err := r.c.call(ctx, "GET", "/api/marketplace/orders", query, nil, false, "data.orders", &out, opts)
	return out, err
}

// ListPayouts — Lists the caller's payouts over a period, with totals; or downloads them as CSV.
//
// GET /api/marketplace/seller/payouts
func (r *MarketplaceService) ListPayouts(ctx context.Context, query *MarketplaceListPayoutsQuery, opts ...CallOption) (*MarketplaceListPayoutsResponseData, error) {
	var out *MarketplaceListPayoutsResponseData
	err := r.c.call(ctx, "GET", "/api/marketplace/seller/payouts", query, nil, false, "data", &out, opts)
	return out, err
}

// ListPurchases — Lists everything the caller bought: tier, price, access, installed version and available update, licence and invoices.
//
// GET /api/marketplace/purchases
func (r *MarketplaceService) ListPurchases(ctx context.Context, opts ...CallOption) ([]MarketplacePurchase, error) {
	var out []MarketplacePurchase
	err := r.c.call(ctx, "GET", "/api/marketplace/purchases", nil, nil, false, "data.purchases", &out, opts)
	return out, err
}

// ListQuestions — Lists a listing's answered questions and the caller's own open ones; for the seller, every open one.
//
// GET /api/marketplace/{listingId}/questions
//
// listingID: The listing's id.
func (r *MarketplaceService) ListQuestions(ctx context.Context, listingID string, opts ...CallOption) (*MarketplaceListQuestionsResponseData, error) {
	var out *MarketplaceListQuestionsResponseData
	err := r.c.call(ctx, "GET", "/api/marketplace/"+pathSeg(listingID)+"/questions", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListReviews — Lists a listing's visible reviews, newest first.
//
// GET /api/marketplace/{listingId}/reviews
//
// listingID: The listing's id.
func (r *MarketplaceService) ListReviews(ctx context.Context, listingID string, query *MarketplaceListReviewsQuery, opts ...CallOption) (*MarketplaceListReviewsResponseData, error) {
	var out *MarketplaceListReviewsResponseData
	err := r.c.call(ctx, "GET", "/api/marketplace/"+pathSeg(listingID)+"/reviews", query, nil, false, "data", &out, opts)
	return out, err
}

// ListSales — Lists the orders on the caller's listings with buyer and listing, and totals.
//
// GET /api/marketplace/seller/sales
func (r *MarketplaceService) ListSales(ctx context.Context, opts ...CallOption) (*MarketplaceListSalesResponseData, error) {
	var out *MarketplaceListSalesResponseData
	err := r.c.call(ctx, "GET", "/api/marketplace/seller/sales", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListSellerListings — Lists the caller's own listings, newest first.
//
// GET /api/marketplace/seller/listings
func (r *MarketplaceService) ListSellerListings(ctx context.Context, opts ...CallOption) ([]MarketplaceSellerListing, error) {
	var out []MarketplaceSellerListing
	err := r.c.call(ctx, "GET", "/api/marketplace/seller/listings", nil, nil, false, "data.listings", &out, opts)
	return out, err
}

// ListWishlist — Lists the caller's wishlist, newest first.
//
// GET /api/marketplace/wishlist
func (r *MarketplaceService) ListWishlist(ctx context.Context, opts ...CallOption) ([]MarketplaceListWishlistResponseDataItemsItem, error) {
	var out []MarketplaceListWishlistResponseDataItemsItem
	err := r.c.call(ctx, "GET", "/api/marketplace/wishlist", nil, nil, false, "data.items", &out, opts)
	return out, err
}

// OfferResource — Offers a project, workflow, database or endpoint in the marketplace: creates (or finds) its product with the offer filled in.
//
// Needs `products:create` and `marketplace:sell` in the resource's organization, the marketplace on the plan, and the right to change the resource. The offer is finished and submitted from the product.
//
// POST /api/marketplace/seller/listings/publish
func (r *MarketplaceService) OfferResource(ctx context.Context, body *MarketplaceOfferResourceBody, opts ...CallOption) (*MarketplaceOfferResourceResponseData, error) {
	var out *MarketplaceOfferResourceResponseData
	err := r.c.call(ctx, "POST", "/api/marketplace/seller/listings/publish", nil, body, false, "data", &out, opts)
	return out, err
}

// OpenDispute — Opens a dispute about an order as its buyer.
//
// Needs the dispute center on the platform (`marketplace_disputes`).
//
// POST /api/marketplace/orders/{orderId}/disputes
//
// orderID: The marketplace order's id.
func (r *MarketplaceService) OpenDispute(ctx context.Context, orderID string, body *MarketplaceOpenDisputeBody, opts ...CallOption) (*MarketplaceOpenDisputeResponseData, error) {
	var out *MarketplaceOpenDisputeResponseData
	err := r.c.call(ctx, "POST", "/api/marketplace/orders/"+pathSeg(orderID)+"/disputes", nil, body, false, "data", &out, opts)
	return out, err
}

// OrderService — Orders a free service; a paid one is bought in the dashboard.
//
// POST /api/marketplace/orders
func (r *MarketplaceService) OrderService(ctx context.Context, body *MarketplaceOrderServiceBody, opts ...CallOption) (*MarketplaceOrder, error) {
	var out *MarketplaceOrder
	err := r.c.call(ctx, "POST", "/api/marketplace/orders", nil, body, false, "data.order", &out, opts)
	return out, err
}

// PayoutSchedule — Tells the caller when they will be paid: what may be paid out now, what is still held, and the next payout date.
//
// GET /api/marketplace/seller/payout-schedule
func (r *MarketplaceService) PayoutSchedule(ctx context.Context, opts ...CallOption) (*MarketplacePayoutScheduleResponseData, error) {
	var out *MarketplacePayoutScheduleResponseData
	err := r.c.call(ctx, "GET", "/api/marketplace/seller/payout-schedule", nil, nil, false, "data", &out, opts)
	return out, err
}

// Preview — Returns what a read-only "Try it live" preview of an approved listing shows: a site's pages or a workflow's steps.
//
// `preview` is null for a listing with nothing to preview (a database, an endpoint, files, a service). 60 previews in ten minutes per person.
//
// GET /api/marketplace/{listingId}/preview
//
// listingID: The listing's id.
func (r *MarketplaceService) Preview(ctx context.Context, listingID string, query *MarketplacePreviewQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/marketplace/"+pathSeg(listingID)+"/preview", query, nil, false, "data.preview", &out, opts)
	return out, err
}

// RefundOrder — Refunds an order, fully or partly, as a goodwill gesture of its seller.
//
// Not for an order already refunded, disputed or cancelled. A full refund marks the order `refunded`; the refunded share comes off what the seller is owed. A bundle is refunded as a whole.
//
// POST /api/marketplace/seller/orders/{orderId}/refund
//
// orderID: The marketplace order's id.
func (r *MarketplaceService) RefundOrder(ctx context.Context, orderID string, body *MarketplaceRefundOrderBody, opts ...CallOption) (*MarketplaceRefundOrderResponseData, error) {
	var out *MarketplaceRefundOrderResponseData
	err := r.c.call(ctx, "POST", "/api/marketplace/seller/orders/"+pathSeg(orderID)+"/refund", nil, body, false, "data", &out, opts)
	return out, err
}

// RemoveFromWishlist — Takes a listing off the caller's wishlist.
//
// DELETE /api/marketplace/wishlist
func (r *MarketplaceService) RemoveFromWishlist(ctx context.Context, query *MarketplaceRemoveFromWishlistQuery, opts ...CallOption) (*MarketplaceRemoveFromWishlistResponse, error) {
	var out *MarketplaceRemoveFromWishlistResponse
	err := r.c.call(ctx, "DELETE", "/api/marketplace/wishlist", query, nil, false, "", &out, opts)
	return out, err
}

// ReportListing — Reports a listing to the platform for review.
//
// Reporting the same listing again answers `duplicate: true`.
//
// POST /api/marketplace/listings/{listingId}/report
//
// listingID: The listing's id.
func (r *MarketplaceService) ReportListing(ctx context.Context, listingID string, body *MarketplaceReportListingBody, opts ...CallOption) (*MarketplaceReportListingResponseData, error) {
	var out *MarketplaceReportListingResponseData
	err := r.c.call(ctx, "POST", "/api/marketplace/listings/"+pathSeg(listingID)+"/report", nil, body, false, "data", &out, opts)
	return out, err
}

// RequestRevision — Asks the seller to revise a service delivery, while the included revisions last.
//
// POST /api/marketplace/orders/{orderId}/revision
//
// orderID: The marketplace order's id.
func (r *MarketplaceService) RequestRevision(ctx context.Context, orderID string, body *MarketplaceRequestRevisionBody, opts ...CallOption) (*MarketplaceOrder, error) {
	var out *MarketplaceOrder
	err := r.c.call(ctx, "POST", "/api/marketplace/orders/"+pathSeg(orderID)+"/revision", nil, body, false, "data.order", &out, opts)
	return out, err
}

// SalesReport — Returns the caller's revenue and orders per day and currency over a period; or downloads it as CSV.
//
// GET /api/marketplace/seller/sales-report
func (r *MarketplaceService) SalesReport(ctx context.Context, query *MarketplaceSalesReportQuery, opts ...CallOption) (*MarketplaceSalesReportResponseData, error) {
	var out *MarketplaceSalesReportResponseData
	err := r.c.call(ctx, "GET", "/api/marketplace/seller/sales-report", query, nil, false, "data", &out, opts)
	return out, err
}

// SaveLegalPage — Creates or changes one of the caller's shop legal pages.
//
// PUT /api/marketplace/seller/legal-pages
func (r *MarketplaceService) SaveLegalPage(ctx context.Context, body *MarketplaceSaveLegalPageBody, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "PUT", "/api/marketplace/seller/legal-pages", nil, body, false, "data.legalPage", &out, opts)
	return out, err
}

// SellerUsage — Returns the selling organization's marketplace limits and what is used.
//
// `-1` means no limit; without a selling organization every limit is 0.
//
// GET /api/marketplace/seller/usage
func (r *MarketplaceService) SellerUsage(ctx context.Context, opts ...CallOption) (*MarketplaceSellerUsageResponseData, error) {
	var out *MarketplaceSellerUsageResponseData
	err := r.c.call(ctx, "GET", "/api/marketplace/seller/usage", nil, nil, false, "data", &out, opts)
	return out, err
}

// SendDisputeMessage — Sends a message on a dispute.
//
// Needs the dispute center on the platform (`marketplace_disputes`).
//
// POST /api/marketplace/disputes/{disputeId}/messages
//
// disputeID: The dispute's id.
func (r *MarketplaceService) SendDisputeMessage(ctx context.Context, disputeID string, body *MarketplaceSendDisputeMessageBody, opts ...CallOption) (*MarketplaceSendDisputeMessageResponse, error) {
	var out *MarketplaceSendDisputeMessageResponse
	err := r.c.call(ctx, "POST", "/api/marketplace/disputes/"+pathSeg(disputeID)+"/messages", nil, body, false, "", &out, opts)
	return out, err
}

// SendOrderMessage — Sends a message on an order between buyer and seller.
//
// Not on a completed or cancelled order.
//
// POST /api/marketplace/orders/{orderId}/messages
//
// orderID: The marketplace order's id.
func (r *MarketplaceService) SendOrderMessage(ctx context.Context, orderID string, body *MarketplaceSendOrderMessageBody, opts ...CallOption) (*MarketplaceSendOrderMessageResponseData, error) {
	var out *MarketplaceSendOrderMessageResponseData
	err := r.c.call(ctx, "POST", "/api/marketplace/orders/"+pathSeg(orderID)+"/messages", nil, body, false, "data", &out, opts)
	return out, err
}

// ShowInShowcase — Lets the listing's product page show a project the purchase installed under "Built with this".
//
// Only a project the order installed (422 `marketplaceShowcaseNotInstalled`); not for a refunded or cancelled order (422 `marketplaceShowcaseOrderVoid`).
//
// POST /api/marketplace/purchases/{orderId}/showcase
//
// orderID: The marketplace order's id.
func (r *MarketplaceService) ShowInShowcase(ctx context.Context, orderID string, body *MarketplaceShowInShowcaseBody, opts ...CallOption) (*MarketplaceShowInShowcaseResponseData, error) {
	var out *MarketplaceShowInShowcaseResponseData
	err := r.c.call(ctx, "POST", "/api/marketplace/purchases/"+pathSeg(orderID)+"/showcase", nil, body, false, "data", &out, opts)
	return out, err
}

// UnfollowShop — Stops following a shop.
//
// DELETE /api/marketplace/shop/{slug}/follow
//
// slug: The shop's slug.
func (r *MarketplaceService) UnfollowShop(ctx context.Context, slug string, opts ...CallOption) (*MarketplaceUnfollowShopResponseData, error) {
	var out *MarketplaceUnfollowShopResponseData
	err := r.c.call(ctx, "DELETE", "/api/marketplace/shop/"+pathSeg(slug)+"/follow", nil, nil, false, "data", &out, opts)
	return out, err
}

// UpdateOrderStatus — Moves a service order to its next status.
//
// The seller moves `ordered` → `in_progress` → `delivered` (and back to `in_progress` after a revision request) and `accepted` → `completed`; the buyer moves `delivered` → `accepted` or `revision_requested`. Cancelling is POST /api/marketplace/seller/orders/{orderId}/cancel, which refunds.
//
// PUT /api/marketplace/orders/{orderId}
//
// orderID: The marketplace order's id.
func (r *MarketplaceService) UpdateOrderStatus(ctx context.Context, orderID string, body *MarketplaceUpdateOrderStatusBody, opts ...CallOption) (*MarketplaceOrder, error) {
	var out *MarketplaceOrder
	err := r.c.call(ctx, "PUT", "/api/marketplace/orders/"+pathSeg(orderID), nil, body, false, "data.order", &out, opts)
	return out, err
}

// UpdateReview — Changes the caller's review of a listing.
//
// PUT /api/marketplace/{listingId}/reviews
//
// listingID: The listing's id.
func (r *MarketplaceService) UpdateReview(ctx context.Context, listingID string, body *MarketplaceUpdateReviewBody, opts ...CallOption) (*MarketplaceReview, error) {
	var out *MarketplaceReview
	err := r.c.call(ctx, "PUT", "/api/marketplace/"+pathSeg(listingID)+"/reviews", nil, body, false, "data.review", &out, opts)
	return out, err
}

// UpdateSellerListing — Changes a draft or rejected listing from before offers became products.
//
// A listing made from a product, or an approved one, answers 410 `marketplaceListingEditedInProducts` with its `productId`: change it through the product's offer.
//
// PUT /api/marketplace/seller/listings/{listingId}
//
// listingID: The listing's id.
func (r *MarketplaceService) UpdateSellerListing(ctx context.Context, listingID string, body *MarketplaceUpdateSellerListingBody, opts ...CallOption) (*MarketplaceSellerListing, error) {
	var out *MarketplaceSellerListing
	err := r.c.call(ctx, "PUT", "/api/marketplace/seller/listings/"+pathSeg(listingID), nil, body, false, "data.listing", &out, opts)
	return out, err
}

// UpdateSellerProfile — Changes the caller's seller profile.
//
// Only the fields sent change.
//
// PUT /api/marketplace/seller
func (r *MarketplaceService) UpdateSellerProfile(ctx context.Context, body *MarketplaceUpdateSellerProfileBody, opts ...CallOption) (*MarketplaceSellerProfile, error) {
	var out *MarketplaceSellerProfile
	err := r.c.call(ctx, "PUT", "/api/marketplace/seller", nil, body, false, "data.profile", &out, opts)
	return out, err
}

// UploadDisputeAttachment — Uploads a file to a dispute as evidence.
//
// Multipart with `file`: at most 10 MB, images, PDF, ZIP and a few other document types.
//
// POST /api/marketplace/disputes/{disputeId}/attachments
//
// disputeID: The dispute's id.
func (r *MarketplaceService) UploadDisputeAttachment(ctx context.Context, disputeID string, body *MarketplaceUploadDisputeAttachmentBody, opts ...CallOption) (*MarketplaceUploadDisputeAttachmentResponseDataAttachment, error) {
	var out *MarketplaceUploadDisputeAttachmentResponseDataAttachment
	err := r.c.call(ctx, "POST", "/api/marketplace/disputes/"+pathSeg(disputeID)+"/attachments", nil, body, true, "data.attachment", &out, opts)
	return out, err
}

// ValidateCoupon — Checks a coupon code against a listing and returns the discount it would give.
//
// POST /api/marketplace/coupons/validate
func (r *MarketplaceService) ValidateCoupon(ctx context.Context, body *MarketplaceValidateCouponBody, opts ...CallOption) (*MarketplaceValidateCouponResponseData, error) {
	var out *MarketplaceValidateCouponResponseData
	err := r.c.call(ctx, "POST", "/api/marketplace/coupons/validate", nil, body, false, "data", &out, opts)
	return out, err
}

// MessengerService — Messenger: Team chat: channels, direct messages, threads, reactions and presence. Every operation needs the Messenger feature on the caller's plan.
type MessengerService struct {
	c *Client
}

func newMessengerService(c *Client) *MessengerService {
	return &MessengerService{c: c}
}

// AddMember — Adds a person to a channel, or joins the caller to a public channel.
//
// Adding somebody else is for the channel's owners and admins, and only a person who shares an organization with the caller or has a public profile. Counts against the plan's `max_messenger_members_per_channel`. Adding a member twice returns the existing membership.
//
// POST /api/messenger/channels/{channelId}/members
//
// channelID: The channel's id.
func (r *MessengerService) AddMember(ctx context.Context, channelID string, body *MessengerAddMemberBody, opts ...CallOption) (*MessengerChannelMember, error) {
	var out *MessengerChannelMember
	err := r.c.call(ctx, "POST", "/api/messenger/channels/"+pathSeg(channelID)+"/members", nil, body, false, "data", &out, opts)
	return out, err
}

// AddReaction — Toggles the caller's reaction on a message.
//
// A person has one reaction per message: a different emoji replaces theirs, the same one removes it.
//
// POST /api/messenger/messages/{messageId}/reactions
//
// messageID: The message's id.
func (r *MessengerService) AddReaction(ctx context.Context, messageID string, body *MessengerAddReactionBody, opts ...CallOption) (*MessengerAddReactionResponseData, error) {
	var out *MessengerAddReactionResponseData
	err := r.c.call(ctx, "POST", "/api/messenger/messages/"+pathSeg(messageID)+"/reactions", nil, body, false, "data", &out, opts)
	return out, err
}

// AIDraft — Drafts a message or an answer for a channel with the platform's AI, from the channel's last 20 messages.
//
// Needs `messenger.use_ai`, membership of the channel and the `messaging:ai_compose` feature; counts against the organization's AI quota. At most 20 calls in five minutes. Nothing is sent — the text comes back for the caller to post.
//
// POST /api/messenger/ai
func (r *MessengerService) AIDraft(ctx context.Context, body *MessengerAiDraftBody, opts ...CallOption) (string, error) {
	var out string
	err := r.c.call(ctx, "POST", "/api/messenger/ai", nil, body, false, "data.text", &out, opts)
	return out, err
}

// CreateChannel — Creates a channel; the caller becomes its owner.
//
// Needs the Channels feature; counts against the plan's `max_messenger_channels`; at most 20 an hour. A `workspace` channel needs `workspaceId`, an `organization` or `public` one an organization (given, or the workspace's).
//
// POST /api/messenger/channels
func (r *MessengerService) CreateChannel(ctx context.Context, body *MessengerCreateChannelBody, opts ...CallOption) (*MessengerChannel, error) {
	var out *MessengerChannel
	err := r.c.call(ctx, "POST", "/api/messenger/channels", nil, body, false, "data", &out, opts)
	return out, err
}

// CreateChannelWebhook — Creates an incoming or outgoing webhook for a channel.
//
// The channel's owners and admins; needs the Messenger webhooks feature. The answer carries the secret `token` — for an incoming webhook the part of its URL that posts into the channel, for an outgoing one the key its calls are signed with.
//
// POST /api/messenger/channels/{channelId}/webhooks
//
// channelID: The channel's id.
func (r *MessengerService) CreateChannelWebhook(ctx context.Context, channelID string, body *MessengerCreateChannelWebhookBody, opts ...CallOption) (*MessengerCreateChannelWebhookResponseData, error) {
	var out *MessengerCreateChannelWebhookResponseData
	err := r.c.call(ctx, "POST", "/api/messenger/channels/"+pathSeg(channelID)+"/webhooks", nil, body, false, "data", &out, opts)
	return out, err
}

// DeleteChannel — Deletes a channel with all its messages.
//
// The channel's owner only. It cannot be undone.
//
// DELETE /api/messenger/channels/{id}
//
// id: The channel's id.
func (r *MessengerService) DeleteChannel(ctx context.Context, id string, opts ...CallOption) (*MessengerDeleteChannelResponseData, error) {
	var out *MessengerDeleteChannelResponseData
	err := r.c.call(ctx, "DELETE", "/api/messenger/channels/"+pathSeg(id), nil, nil, false, "data", &out, opts)
	return out, err
}

// DeleteChannelWebhook — Deletes a channel's webhook.
//
// The channel's owners and admins; needs the Messenger webhooks feature.
//
// DELETE /api/messenger/channels/{channelId}/webhooks/{webhookId}
//
// channelID: The channel's id.
// webhookID: The webhook's id.
func (r *MessengerService) DeleteChannelWebhook(ctx context.Context, channelID string, webhookID string, opts ...CallOption) (*MessengerDeleteChannelWebhookResponseData, error) {
	var out *MessengerDeleteChannelWebhookResponseData
	err := r.c.call(ctx, "DELETE", "/api/messenger/channels/"+pathSeg(channelID)+"/webhooks/"+pathSeg(webhookID), nil, nil, false, "data", &out, opts)
	return out, err
}

// DeleteMessage — Deletes a message.
//
// The sender, or the channel's owners and admins. The message stays in its thread as deleted, with its content removed.
//
// DELETE /api/messenger/messages/{messageId}
//
// messageID: The message's id.
func (r *MessengerService) DeleteMessage(ctx context.Context, messageID string, opts ...CallOption) (*MessengerDeleteMessageResponseData, error) {
	var out *MessengerDeleteMessageResponseData
	err := r.c.call(ctx, "DELETE", "/api/messenger/messages/"+pathSeg(messageID), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetChannel — Returns one channel with the caller's membership.
//
// A channel the caller is not in answers only if they could join it.
//
// GET /api/messenger/channels/{id}
//
// id: The channel's id.
func (r *MessengerService) GetChannel(ctx context.Context, id string, opts ...CallOption) (*MessengerGetChannelResponseData, error) {
	var out *MessengerGetChannelResponseData
	err := r.c.call(ctx, "GET", "/api/messenger/channels/"+pathSeg(id), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetE2E — Returns whether a direct message is end-to-end encrypted, with both participants' public keys.
//
// Direct messages between two people only.
//
// GET /api/messenger/channels/{channelId}/e2e
//
// channelID: The channel's id.
func (r *MessengerService) GetE2E(ctx context.Context, channelID string, opts ...CallOption) (*MessengerGetE2EResponseData, error) {
	var out *MessengerGetE2EResponseData
	err := r.c.call(ctx, "GET", "/api/messenger/channels/"+pathSeg(channelID)+"/e2e", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetPresence — Returns the presence of up to 200 people.
//
// A person who hides their presence from the caller shows as offline.
//
// GET /api/messenger/presence
func (r *MessengerService) GetPresence(ctx context.Context, query *MessengerGetPresenceQuery, opts ...CallOption) ([]MessengerPresence, error) {
	var out []MessengerPresence
	err := r.c.call(ctx, "GET", "/api/messenger/presence", query, nil, false, "data", &out, opts)
	return out, err
}

// GetSettings — Returns the caller's Messenger settings.
//
// GET /api/messenger/settings
func (r *MessengerService) GetSettings(ctx context.Context, opts ...CallOption) (*MessengerSettings, error) {
	var out *MessengerSettings
	err := r.c.call(ctx, "GET", "/api/messenger/settings", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetUnread — Returns how many of the caller's channels have unread messages.
//
// Channels, not messages. At most 60 calls a minute; beyond that, and without the Messenger feature, it answers 0.
//
// GET /api/messenger/unread
func (r *MessengerService) GetUnread(ctx context.Context, opts ...CallOption) (*MessengerGetUnreadResponseData, error) {
	var out *MessengerGetUnreadResponseData
	err := r.c.call(ctx, "GET", "/api/messenger/unread", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListChannels — Lists the channels and DMs the caller is in, with the public channels they can join.
//
// Most recently active first, with the caller's unread count per channel.
//
// GET /api/messenger/channels
func (r *MessengerService) ListChannels(ctx context.Context, query *MessengerListChannelsQuery, opts ...CallOption) ([]MessengerChannelListItem, error) {
	var out []MessengerChannelListItem
	err := r.c.call(ctx, "GET", "/api/messenger/channels", query, nil, false, "data", &out, opts)
	return out, err
}

// ListChannelWebhooks — Lists a channel's incoming and outgoing webhooks.
//
// Members only; needs the Messenger webhooks feature. The secret token is never listed — see `messenger.revealChannelWebhook`.
//
// GET /api/messenger/channels/{channelId}/webhooks
//
// channelID: The channel's id.
func (r *MessengerService) ListChannelWebhooks(ctx context.Context, channelID string, opts ...CallOption) ([]MessengerChannelWebhook, error) {
	var out []MessengerChannelWebhook
	err := r.c.call(ctx, "GET", "/api/messenger/channels/"+pathSeg(channelID)+"/webhooks", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListMembers — Lists a channel's members with their presence.
//
// Members only. A member's email is shown only when they share an organization with the caller.
//
// GET /api/messenger/channels/{channelId}/members
//
// channelID: The channel's id.
func (r *MessengerService) ListMembers(ctx context.Context, channelID string, opts ...CallOption) ([]MessengerListMembersResponseDataItem, error) {
	var out []MessengerListMembersResponseDataItem
	err := r.c.call(ctx, "GET", "/api/messenger/channels/"+pathSeg(channelID)+"/members", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListMessages — Lists a channel's messages, or a thread's replies, 50 at a time, oldest first.
//
// Members only (or anyone who may join a public channel). Deleted messages are left out; a plan with a message-history limit hides older messages.
//
// GET /api/messenger/channels/{channelId}/messages
//
// channelID: The channel's id.
func (r *MessengerService) ListMessages(ctx context.Context, channelID string, query *MessengerListMessagesQuery, opts ...CallOption) (*MessengerListMessagesResponseData, error) {
	var out *MessengerListMessagesResponseData
	err := r.c.call(ctx, "GET", "/api/messenger/channels/"+pathSeg(channelID)+"/messages", query, nil, false, "data", &out, opts)
	return out, err
}

// ListPinned — Lists a channel's pinned messages, most recently pinned first.
//
// GET /api/messenger/channels/{channelId}/pinned
//
// channelID: The channel's id.
func (r *MessengerService) ListPinned(ctx context.Context, channelID string, opts ...CallOption) (*MessengerListPinnedResponseData, error) {
	var out *MessengerListPinnedResponseData
	err := r.c.call(ctx, "GET", "/api/messenger/channels/"+pathSeg(channelID)+"/pinned", nil, nil, false, "data", &out, opts)
	return out, err
}

// LookupUsers — Finds people the caller can message, by name or email.
//
// At least three characters; at most 20 results and 20 searches a minute. Only people with a public profile, or who share an organization with the caller.
//
// GET /api/messenger/users
func (r *MessengerService) LookupUsers(ctx context.Context, query *MessengerLookupUsersQuery, opts ...CallOption) ([]MessengerLookupUsersResponseDataItem, error) {
	var out []MessengerLookupUsersResponseDataItem
	err := r.c.call(ctx, "GET", "/api/messenger/users", query, nil, false, "data", &out, opts)
	return out, err
}

// MarkAllRead — Marks every channel the caller is in as read.
//
// POST /api/messenger/read-all
func (r *MessengerService) MarkAllRead(ctx context.Context, opts ...CallOption) (*MessengerMarkAllReadResponseData, error) {
	var out *MessengerMarkAllReadResponseData
	err := r.c.call(ctx, "POST", "/api/messenger/read-all", nil, nil, false, "data", &out, opts)
	return out, err
}

// MarkRead — Marks a channel as read by the caller.
//
// Moves the caller's read position forward to the given message or time (default now); it never moves back. At most 120 a minute.
//
// POST /api/messenger/channels/{channelId}/read
//
// channelID: The channel's id.
func (r *MessengerService) MarkRead(ctx context.Context, channelID string, body *MessengerMarkReadBody, opts ...CallOption) (*MessengerMarkReadResponseData, error) {
	var out *MessengerMarkReadResponseData
	err := r.c.call(ctx, "POST", "/api/messenger/channels/"+pathSeg(channelID)+"/read", nil, body, false, "data", &out, opts)
	return out, err
}

// OpenDm — Opens the direct message with one or more people, creating it the first time.
//
// Needs `messenger.use_dm` and the DMs feature; at most 30 an hour and eight other people. A person with a private profile, or one visible only within their organization the caller does not share, cannot be messaged.
//
// POST /api/messenger/dms
func (r *MessengerService) OpenDm(ctx context.Context, body *MessengerOpenDmBody, opts ...CallOption) (*MessengerChannel, error) {
	var out *MessengerChannel
	err := r.c.call(ctx, "POST", "/api/messenger/dms", nil, body, false, "data", &out, opts)
	return out, err
}

// OpenRecordChannel — Opens the discussion channel of a database record, creating it the first time, and joins the caller to it.
//
// The caller needs access to the workspace and to the record.
//
// POST /api/messenger/channels/record
func (r *MessengerService) OpenRecordChannel(ctx context.Context, body *MessengerOpenRecordChannelBody, opts ...CallOption) (*MessengerOpenRecordChannelResponseData, error) {
	var out *MessengerOpenRecordChannelResponseData
	err := r.c.call(ctx, "POST", "/api/messenger/channels/record", nil, body, false, "data", &out, opts)
	return out, err
}

// Pin — Pins a message to its channel, or unpins it when it is pinned.
//
// Any member of the channel.
//
// POST /api/messenger/messages/{messageId}/pin
//
// messageID: The message's id.
func (r *MessengerService) Pin(ctx context.Context, messageID string, opts ...CallOption) (*MessengerPinResponseData, error) {
	var out *MessengerPinResponseData
	err := r.c.call(ctx, "POST", "/api/messenger/messages/"+pathSeg(messageID)+"/pin", nil, nil, false, "data", &out, opts)
	return out, err
}

// RemoveMember — Removes a person from a channel, or the caller leaves it.
//
// Removing somebody else is for the channel's owners and admins.
//
// DELETE /api/messenger/channels/{channelId}/members
//
// channelID: The channel's id.
func (r *MessengerService) RemoveMember(ctx context.Context, channelID string, query *MessengerRemoveMemberQuery, opts ...CallOption) (*MessengerRemoveMemberResponseData, error) {
	var out *MessengerRemoveMemberResponseData
	err := r.c.call(ctx, "DELETE", "/api/messenger/channels/"+pathSeg(channelID)+"/members", query, nil, false, "data", &out, opts)
	return out, err
}

// RevealChannelWebhook — Returns a channel webhook's secret token.
//
// The channel's owners and admins; at most 10 an hour, and every reveal is audit-logged.
//
// POST /api/messenger/channels/{channelId}/webhooks/{webhookId}/reveal
//
// channelID: The channel's id.
// webhookID: The webhook's id.
func (r *MessengerService) RevealChannelWebhook(ctx context.Context, channelID string, webhookID string, opts ...CallOption) (*MessengerRevealChannelWebhookResponseData, error) {
	var out *MessengerRevealChannelWebhookResponseData
	err := r.c.call(ctx, "POST", "/api/messenger/channels/"+pathSeg(channelID)+"/webhooks/"+pathSeg(webhookID)+"/reveal", nil, nil, false, "data", &out, opts)
	return out, err
}

// Search — Searches the messages of the caller's channels, newest first.
//
// A substring match, at most 50 results; at most 20 searches a minute.
//
// GET /api/messenger/search
func (r *MessengerService) Search(ctx context.Context, query *MessengerSearchQuery, opts ...CallOption) (*MessengerSearchResponseData, error) {
	var out *MessengerSearchResponseData
	err := r.c.call(ctx, "GET", "/api/messenger/search", query, nil, false, "data", &out, opts)
	return out, err
}

// SendMessage — Posts a message to a channel, or a reply to a thread.
//
// Members only; in an announcement channel, its owners and admins. At most 60 a minute and 10,000 characters. Attachments need the file-sharing feature. Mentions of people who are not members are dropped. A message the content filter blocks answers 422.
//
// POST /api/messenger/channels/{channelId}/messages
//
// channelID: The channel's id.
func (r *MessengerService) SendMessage(ctx context.Context, channelID string, body *MessengerSendMessageBody, opts ...CallOption) (*MessengerMessage, error) {
	var out *MessengerMessage
	err := r.c.call(ctx, "POST", "/api/messenger/channels/"+pathSeg(channelID)+"/messages", nil, body, false, "data", &out, opts)
	return out, err
}

// SetE2EEnabled — Turns end-to-end encryption of a direct message on or off.
//
// Direct messages between two people only. Turning it on needs the E2E feature and both participants' registered keys; once on, only the person who opened the DM may turn it off.
//
// POST /api/messenger/channels/{channelId}/e2e
//
// channelID: The channel's id.
func (r *MessengerService) SetE2EEnabled(ctx context.Context, channelID string, body *MessengerSetE2EEnabledBody, opts ...CallOption) (*MessengerSetE2EEnabledResponseData, error) {
	var out *MessengerSetE2EEnabledResponseData
	err := r.c.call(ctx, "POST", "/api/messenger/channels/"+pathSeg(channelID)+"/e2e", nil, body, false, "data", &out, opts)
	return out, err
}

// SetPresence — Sets the caller's status, or reports them online.
//
// `online` without emoji, text or expiry is a heartbeat; anything else sets the caller's own status, which wins over the automatic one (focus time, meetings, absence) until it expires.
//
// PUT /api/messenger/presence
func (r *MessengerService) SetPresence(ctx context.Context, body *MessengerSetPresenceBody, opts ...CallOption) (*MessengerSetPresenceResponseData, error) {
	var out *MessengerSetPresenceResponseData
	err := r.c.call(ctx, "PUT", "/api/messenger/presence", nil, body, false, "data", &out, opts)
	return out, err
}

// SetTyping — Tells the channel's open clients that the caller started or stopped typing.
//
// Nothing is stored. At most 6 calls in 10 seconds per channel.
//
// POST /api/messenger/channels/{channelId}/typing
//
// channelID: The channel's id.
func (r *MessengerService) SetTyping(ctx context.Context, channelID string, body *MessengerSetTypingBody, opts ...CallOption) (*MessengerSetTypingResponseData, error) {
	var out *MessengerSetTypingResponseData
	err := r.c.call(ctx, "POST", "/api/messenger/channels/"+pathSeg(channelID)+"/typing", nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateChannel — Changes a channel's name, description, topic, icon, type or visibility, or archives it.
//
// The channel's owners and admins only; widening who can find it is the owner's.
//
// PUT /api/messenger/channels/{id}
//
// id: The channel's id.
func (r *MessengerService) UpdateChannel(ctx context.Context, id string, body *MessengerUpdateChannelBody, opts ...CallOption) (*MessengerChannel, error) {
	var out *MessengerChannel
	err := r.c.call(ctx, "PUT", "/api/messenger/channels/"+pathSeg(id), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateChannelWebhook — Renames, pauses or redirects a channel's webhook.
//
// The channel's owners and admins; needs the Messenger webhooks feature.
//
// PATCH /api/messenger/channels/{channelId}/webhooks/{webhookId}
//
// channelID: The channel's id.
// webhookID: The webhook's id.
func (r *MessengerService) UpdateChannelWebhook(ctx context.Context, channelID string, webhookID string, body *MessengerUpdateChannelWebhookBody, opts ...CallOption) (*MessengerChannelWebhook, error) {
	var out *MessengerChannelWebhook
	err := r.c.call(ctx, "PATCH", "/api/messenger/channels/"+pathSeg(channelID)+"/webhooks/"+pathSeg(webhookID), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateMessage — Edits one of the caller's messages.
//
// Only the sender, while still a member of the channel; a deleted message answers 410.
//
// PUT /api/messenger/messages/{messageId}
//
// messageID: The message's id.
func (r *MessengerService) UpdateMessage(ctx context.Context, messageID string, body *MessengerUpdateMessageBody, opts ...CallOption) (*MessengerMessageRecord, error) {
	var out *MessengerMessageRecord
	err := r.c.call(ctx, "PUT", "/api/messenger/messages/"+pathSeg(messageID), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateSettings — Changes the caller's Messenger settings.
//
// PUT /api/messenger/settings
func (r *MessengerService) UpdateSettings(ctx context.Context, body *MessengerUpdateSettingsBody, opts ...CallOption) (*MessengerSettings, error) {
	var out *MessengerSettings
	err := r.c.call(ctx, "PUT", "/api/messenger/settings", nil, body, false, "data", &out, opts)
	return out, err
}

// MilestonesService — Milestones: An organisation's milestones: tracks of recorded figures (visitors, submissions, bookings, paid revenue, CRM contacts, workflow runs, years, referrals) with levels, and the rewards a level gives. The platform defines the tracks and levels; an organisation reads its progress, opts in to the public website badge, gives the shipping address of a physical reward, and may keep its milestones in one of its databases. Everything needs Milestones on the plan and `milestones:read` at organisation level; revenue figures need `invoicing:read` besides (without it the revenue track is left out).
type MilestonesService struct {
	c *Client

	Database *MilestonesDatabaseService
}

func newMilestonesService(c *Client) *MilestonesService {
	return &MilestonesService{c: c, Database: newMilestonesDatabaseService(c)}
}

// Get — Returns an organisation's milestone tracks with its figures, the levels reached and the rewards waiting.
//
// A surprise level not reached shows nothing but its place. Reading the figures also refreshes the organisation's evaluation, at most every ten minutes.
//
// GET /api/milestones
func (r *MilestonesService) Get(ctx context.Context, query *MilestonesGetQuery, opts ...CallOption) (*MilestoneOverview, error) {
	var out *MilestoneOverview
	err := r.c.call(ctx, "GET", "/api/milestones", query, nil, false, "", &out, opts)
	return out, err
}

// SubmitShippingAddress — Gives or corrects the shipping address of a physical reward that has not been shipped yet.
//
// Needs `milestones:share` besides `milestones:read`. The rewards waiting for one are `awaitingAddress` of milestones.get. 400 `rewardAddressInvalid` for an incomplete address; 404 `rewardItemNotFound` for an item that is not this organisation's or already shipped.
//
// PUT /api/milestones/fulfilment/{itemId}
//
// itemID: The reward item's id, from `awaitingAddress`.
func (r *MilestonesService) SubmitShippingAddress(ctx context.Context, itemID string, body *MilestonesSubmitShippingAddressBody, query *MilestonesSubmitShippingAddressQuery, opts ...CallOption) (*MilestonesSubmitShippingAddressResponseItem, error) {
	var out *MilestonesSubmitShippingAddressResponseItem
	err := r.c.call(ctx, "PUT", "/api/milestones/fulfilment/"+pathSeg(itemID), query, body, false, "item", &out, opts)
	return out, err
}

// UpdateSettings — Switches the public website badge with the organisation's latest milestone on or off.
//
// Needs `milestones:share` besides `milestones:read`. Switched off, the badge's address answers 404 at once. It never shows revenue.
//
// PUT /api/milestones/settings
func (r *MilestonesService) UpdateSettings(ctx context.Context, body *MilestonesUpdateSettingsBody, query *MilestonesUpdateSettingsQuery, opts ...CallOption) (bool, error) {
	var out bool
	err := r.c.call(ctx, "PUT", "/api/milestones/settings", query, body, false, "shareBadge", &out, opts)
	return out, err
}

// MilestonesDatabaseService — milestones.database
type MilestonesDatabaseService struct {
	c *Client
}

func newMilestonesDatabaseService(c *Client) *MilestonesDatabaseService {
	return &MilestonesDatabaseService{c: c}
}

// Get — Returns which database the organisation's milestones are kept in, and the databases and workspaces they could go to.
//
// Needs `milestones:share` besides `milestones:read`. `databases` lists the organisation's databases the caller may edit that are not managed by something else; `workspaces` those where a new one can be created.
//
// GET /api/milestones/database
func (r *MilestonesDatabaseService) Get(ctx context.Context, query *MilestonesDatabaseGetQuery, opts ...CallOption) (*MilestonesDatabaseGetResponse, error) {
	var out *MilestonesDatabaseGetResponse
	err := r.c.call(ctx, "GET", "/api/milestones/database", query, nil, false, "", &out, opts)
	return out, err
}

// Stop — Stops keeping the organisation's milestones in a database.
//
// Needs `milestones:share` besides `milestones:read`. The rows already written stay.
//
// DELETE /api/milestones/database
func (r *MilestonesDatabaseService) Stop(ctx context.Context, query *MilestonesDatabaseStopQuery, opts ...CallOption) (*MilestonesDatabaseStopResponse, error) {
	var out *MilestonesDatabaseStopResponse
	err := r.c.call(ctx, "DELETE", "/api/milestones/database", query, nil, false, "", &out, opts)
	return out, err
}

// Sync — Writes the organisation's milestones into a database, or tells what it would write.
//
// Needs `milestones:share` besides `milestones:read`. With `apply: true` the rows are written as the caller and the database is kept in sync from then on, with every new level reached; without it nothing is written. The rows go through the Databases hub's own rules (access, approvals, automations, the record limit): 403, 404 `databaseNotFound`, 409 `milestoneDatabaseIncompatible` or `limitReached` when they refuse.
//
// POST /api/milestones/database
func (r *MilestonesDatabaseService) Sync(ctx context.Context, body *MilestonesDatabaseSyncBody, query *MilestonesDatabaseSyncQuery, opts ...CallOption) (*MilestoneDatabaseResult, error) {
	var out *MilestoneDatabaseResult
	err := r.c.call(ctx, "POST", "/api/milestones/database", query, body, false, "result", &out, opts)
	return out, err
}

// NotificationsService — Notifications
type NotificationsService struct {
	c *Client
}

func newNotificationsService(c *Client) *NotificationsService {
	return &NotificationsService{c: c}
}

// Bulk — Marks several of the caller's notifications read or unread, or deletes them.
//
// At most 500 ids at a time. Ids that are not the caller's are skipped; `changed` lists the ones acted on.
//
// POST /api/notifications/bulk
func (r *NotificationsService) Bulk(ctx context.Context, body *NotificationsBulkBody, opts ...CallOption) ([]string, error) {
	var out []string
	err := r.c.call(ctx, "POST", "/api/notifications/bulk", nil, body, false, "changed", &out, opts)
	return out, err
}

// Delete — Deletes one of the caller's notifications.
//
// DELETE /api/notifications
func (r *NotificationsService) Delete(ctx context.Context, query *NotificationsDeleteQuery, opts ...CallOption) (*NotificationsDeleteResponse, error) {
	var out *NotificationsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/notifications", query, nil, false, "", &out, opts)
	return out, err
}

// GetSettings — Returns the caller's notification settings: every kind of notification by category, with its channels and whether it is on.
//
// Without saved settings the platform's defaults come back; categories added to the platform since the caller saved appear with their defaults.
//
// GET /api/settings/notifications
func (r *NotificationsService) GetSettings(ctx context.Context, opts ...CallOption) ([]NotificationsGetSettingsResponseCategoriesItem, error) {
	var out []NotificationsGetSettingsResponseCategoriesItem
	err := r.c.call(ctx, "GET", "/api/settings/notifications", nil, nil, false, "categories", &out, opts)
	return out, err
}

// List — Lists the caller's notifications, newest first, a page at a time.
//
// At most 30 calls a minute. For the next page pass the answer's `nextCursor` as `cursor`; it is null on the last page.
//
// GET /api/notifications
func (r *NotificationsService) List(ctx context.Context, query *NotificationsListQuery, opts ...CallOption) ([]Notification, error) {
	var out []Notification
	err := r.c.call(ctx, "GET", "/api/notifications", query, nil, false, "notifications", &out, opts)
	return out, err
}

// MarkAllRead — Marks all of the caller's notifications as read.
//
// POST /api/notifications/mark-all-read
func (r *NotificationsService) MarkAllRead(ctx context.Context, opts ...CallOption) (float64, error) {
	var out float64
	err := r.c.call(ctx, "POST", "/api/notifications/mark-all-read", nil, nil, false, "marked", &out, opts)
	return out, err
}

// MarkRead — Marks one of the caller's notifications as read.
//
// PUT /api/notifications
func (r *NotificationsService) MarkRead(ctx context.Context, body *NotificationsMarkReadBody, opts ...CallOption) (*Notification, error) {
	var out *Notification
	err := r.c.call(ctx, "PUT", "/api/notifications", nil, body, false, "", &out, opts)
	return out, err
}

// UnreadCount — Returns how many of the caller's notifications are unread.
//
// At most 60 calls a minute; beyond that it answers 0.
//
// GET /api/notifications/unread-count
func (r *NotificationsService) UnreadCount(ctx context.Context, opts ...CallOption) (float64, error) {
	var out float64
	err := r.c.call(ctx, "GET", "/api/notifications/unread-count", nil, nil, false, "count", &out, opts)
	return out, err
}

// OauthService — OAuth: Who a token belongs to, and the OAuth apps (Zapier, Make, an MCP client …) the caller has connected to their account.
type OauthService struct {
	c *Client

	Connections *OauthConnectionsService
}

func newOauthService(c *Client) *OauthService {
	return &OauthService{c: c, Connections: newOauthConnectionsService(c)}
}

// Userinfo — Returns who the token or session belongs to, in OpenID Connect UserInfo form.
//
// A token holding only OAuth client scopes (`openid`, `profile`, `email`, `offline_access`, `mcp:*`) gets `sub`, plus `email` with the `email` scope and `name` and `picture` with `profile` — never the organizations. Any other token or a session gets every field.
//
// GET /api/oauth/userinfo — Token permission: read:projects.
func (r *OauthService) Userinfo(ctx context.Context, opts ...CallOption) (*OauthUserinfoResponseData, error) {
	var out *OauthUserinfoResponseData
	err := r.c.call(ctx, "GET", "/api/oauth/userinfo", nil, nil, false, "data", &out, opts)
	return out, err
}

// OauthConnectionsService — oauth.connections
type OauthConnectionsService struct {
	c *Client
}

func newOauthConnectionsService(c *Client) *OauthConnectionsService {
	return &OauthConnectionsService{c: c}
}

// List — Lists the OAuth apps the caller has connected to their account and that still hold a live token.
//
// GET /api/oauth/my-connections
func (r *OauthConnectionsService) List(ctx context.Context, opts ...CallOption) ([]OAuthConnection, error) {
	var out []OAuthConnection
	err := r.c.call(ctx, "GET", "/api/oauth/my-connections", nil, nil, false, "data.connections", &out, opts)
	return out, err
}

// Revoke — Disconnects an OAuth app from the caller's account.
//
// Revokes every refresh and access token the app holds for the caller and switches off its webhook subscriptions, at once; the caller gets a security email.
//
// DELETE /api/oauth/my-connections/{clientId}
//
// clientID: The app's client id.
func (r *OauthConnectionsService) Revoke(ctx context.Context, clientID string, opts ...CallOption) (*OauthConnectionsRevokeResponseData, error) {
	var out *OauthConnectionsRevokeResponseData
	err := r.c.call(ctx, "DELETE", "/api/oauth/my-connections/"+pathSeg(clientID), nil, nil, false, "data", &out, opts)
	return out, err
}

// OrganizationsService — Organizations
type OrganizationsService struct {
	c *Client

	EmailTemplates *OrganizationsEmailTemplatesService
}

func newOrganizationsService(c *Client) *OrganizationsService {
	return &OrganizationsService{c: c, EmailTemplates: newOrganizationsEmailTemplatesService(c)}
}

// AddMember — Invites a person to an organization by email, or adds them at once when the organization adds directly.
//
// Needs `organization.manage_members`. In `invite` mode (the default, see `organizations.getMemberAdding`) every address gets an invitation and the answer is the same whether or not it has an account; in `direct` mode an existing account joins at once (`added`). The seat is checked against the plan.
//
// POST /api/organizations/{organizationId}/members
//
// organizationID: The organization's id.
func (r *OrganizationsService) AddMember(ctx context.Context, organizationID string, body *OrganizationsAddMemberBody, opts ...CallOption) (*OrganizationsAddMemberResponseData, error) {
	var out *OrganizationsAddMemberResponseData
	err := r.c.call(ctx, "POST", "/api/organizations/"+pathSeg(organizationID)+"/members", nil, body, false, "data", &out, opts)
	return out, err
}

// AssignTeam — Gives a team a role across the whole organization.
//
// Needs `team.assign`.
//
// POST /api/organizations/{organizationId}/teams
//
// organizationID: The organization's id.
func (r *OrganizationsService) AssignTeam(ctx context.Context, organizationID string, body *OrganizationsAssignTeamBody, opts ...CallOption) (*TeamRoleAssignment, error) {
	var out *TeamRoleAssignment
	err := r.c.call(ctx, "POST", "/api/organizations/"+pathSeg(organizationID)+"/teams", nil, body, false, "data", &out, opts)
	return out, err
}

// CancelDeletion — Cancels an organization's scheduled deletion.
//
// Needs `organization.delete`.
//
// DELETE /api/organizations/{organizationId}/scheduled-deletion
//
// organizationID: The organization's id.
func (r *OrganizationsService) CancelDeletion(ctx context.Context, organizationID string, opts ...CallOption) (*OrganizationsCancelDeletionResponseData, error) {
	var out *OrganizationsCancelDeletionResponseData
	err := r.c.call(ctx, "DELETE", "/api/organizations/"+pathSeg(organizationID)+"/scheduled-deletion", nil, nil, false, "data", &out, opts)
	return out, err
}

// Create — Creates an organization with the caller as its owner, a default workspace and the standard roles.
//
// POST /api/organizations
func (r *OrganizationsService) Create(ctx context.Context, body *OrganizationsCreateBody, opts ...CallOption) (*OrganizationsCreateResponseData, error) {
	var out *OrganizationsCreateResponseData
	err := r.c.call(ctx, "POST", "/api/organizations", nil, body, false, "data", &out, opts)
	return out, err
}

// CreateFreezeWindow — Adds a publish freeze window, one-off or recurring.
//
// Organization administrators only. Give `cronExpr`, or `startsAt` and `endsAt`.
//
// POST /api/organizations/{organizationId}/freeze-windows
//
// organizationID: The organization's id.
func (r *OrganizationsService) CreateFreezeWindow(ctx context.Context, organizationID string, body *FreezeWindowInput, opts ...CallOption) (*OrganizationsCreateFreezeWindowResponseData, error) {
	var out *OrganizationsCreateFreezeWindowResponseData
	err := r.c.call(ctx, "POST", "/api/organizations/"+pathSeg(organizationID)+"/freeze-windows", nil, body, false, "data", &out, opts)
	return out, err
}

// Delete — Schedules an organization for deletion in 14 days.
//
// Needs `organization.delete`. Until then the deletion can be cancelled with `organizations.cancelDeletion`.
//
// DELETE /api/organizations/{organizationId}
//
// organizationID: The organization's id.
func (r *OrganizationsService) Delete(ctx context.Context, organizationID string, opts ...CallOption) (*OrganizationsDeleteResponseData, error) {
	var out *OrganizationsDeleteResponseData
	err := r.c.call(ctx, "DELETE", "/api/organizations/"+pathSeg(organizationID), nil, nil, false, "data", &out, opts)
	return out, err
}

// DeleteFreezeWindow — Removes a publish freeze window.
//
// DELETE /api/organizations/{organizationId}/freeze-windows/{windowId}
//
// organizationID: The organization's id.
// windowID: The freeze window's id.
func (r *OrganizationsService) DeleteFreezeWindow(ctx context.Context, organizationID string, windowID string, opts ...CallOption) (*OrganizationsDeleteFreezeWindowResponse, error) {
	var out *OrganizationsDeleteFreezeWindowResponse
	err := r.c.call(ctx, "DELETE", "/api/organizations/"+pathSeg(organizationID)+"/freeze-windows/"+pathSeg(windowID), nil, nil, false, "", &out, opts)
	return out, err
}

// Get — Returns one organization; billing details only to those who may read billing.
//
// GET /api/organizations/{organizationId} — Token permission: read:projects.
//
// organizationID: The organization's id.
func (r *OrganizationsService) Get(ctx context.Context, organizationID string, opts ...CallOption) (*Organization, error) {
	var out *Organization
	err := r.c.call(ctx, "GET", "/api/organizations/"+pathSeg(organizationID), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetMemberAdding — Returns how the organization adds a person by email: always by invitation, or directly when they have an account.
//
// GET /api/organizations/{organizationId}/member-adding
//
// organizationID: The organization's id.
func (r *OrganizationsService) GetMemberAdding(ctx context.Context, organizationID string, opts ...CallOption) (*OrganizationsGetMemberAddingResponseData, error) {
	var out *OrganizationsGetMemberAddingResponseData
	err := r.c.call(ctx, "GET", "/api/organizations/"+pathSeg(organizationID)+"/member-adding", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetReportRecipients — Returns whether scheduled reports may go to addresses outside the organization.
//
// GET /api/organizations/{organizationId}/report-recipients
//
// organizationID: The organization's id.
func (r *OrganizationsService) GetReportRecipients(ctx context.Context, organizationID string, opts ...CallOption) (*OrganizationsGetReportRecipientsResponseData, error) {
	var out *OrganizationsGetReportRecipientsResponseData
	err := r.c.call(ctx, "GET", "/api/organizations/"+pathSeg(organizationID)+"/report-recipients", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetScheduledDeletion — Returns when an organization is due to be deleted, if it is.
//
// GET /api/organizations/{organizationId}/scheduled-deletion
//
// organizationID: The organization's id.
func (r *OrganizationsService) GetScheduledDeletion(ctx context.Context, organizationID string, opts ...CallOption) (*OrganizationsGetScheduledDeletionResponseData, error) {
	var out *OrganizationsGetScheduledDeletionResponseData
	err := r.c.call(ctx, "GET", "/api/organizations/"+pathSeg(organizationID)+"/scheduled-deletion", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetSeatCapacity — Returns the organization's seats: included, bought and taken, per tier on a tiered plan.
//
// GET /api/organizations/{organizationId}/seat-capacity
//
// organizationID: The organization's id.
func (r *OrganizationsService) GetSeatCapacity(ctx context.Context, organizationID string, opts ...CallOption) (*SeatCapacity, error) {
	var out *SeatCapacity
	err := r.c.call(ctx, "GET", "/api/organizations/"+pathSeg(organizationID)+"/seat-capacity", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetSharingSettings — Returns the organization's file-sharing defaults and limits.
//
// GET /api/organizations/{organizationId}/sharing-settings
//
// organizationID: The organization's id.
func (r *OrganizationsService) GetSharingSettings(ctx context.Context, organizationID string, opts ...CallOption) (*AssetSharingSettings, error) {
	var out *AssetSharingSettings
	err := r.c.call(ctx, "GET", "/api/organizations/"+pathSeg(organizationID)+"/sharing-settings", nil, nil, false, "settings", &out, opts)
	return out, err
}

// List — Lists the organizations the caller belongs to or has a role in.
//
// GET /api/organizations — Token permission: read:projects.
func (r *OrganizationsService) List(ctx context.Context, query *OrganizationsListQuery, opts ...CallOption) ([]Organization, error) {
	var out []Organization
	err := r.c.call(ctx, "GET", "/api/organizations", query, nil, false, "data", &out, opts)
	return out, err
}

// ListFreezeWindows — Lists an organization's publish freeze windows — times nothing may go live.
//
// GET /api/organizations/{organizationId}/freeze-windows
//
// organizationID: The organization's id.
func (r *OrganizationsService) ListFreezeWindows(ctx context.Context, organizationID string, opts ...CallOption) ([]FreezeWindow, error) {
	var out []FreezeWindow
	err := r.c.call(ctx, "GET", "/api/organizations/"+pathSeg(organizationID)+"/freeze-windows", nil, nil, false, "data.items", &out, opts)
	return out, err
}

// ListInvitations — Lists an organization's open invitations.
//
// GET /api/organizations/{organizationId}/invitations
//
// organizationID: The organization's id.
func (r *OrganizationsService) ListInvitations(ctx context.Context, organizationID string, opts ...CallOption) ([]OrganizationInvitation, error) {
	var out []OrganizationInvitation
	err := r.c.call(ctx, "GET", "/api/organizations/"+pathSeg(organizationID)+"/invitations", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListMembers — Lists an organization's members with their user details and seat.
//
// GET /api/organizations/{organizationId}/members
//
// organizationID: The organization's id.
func (r *OrganizationsService) ListMembers(ctx context.Context, organizationID string, opts ...CallOption) ([]OrganizationMember, error) {
	var out []OrganizationMember
	err := r.c.call(ctx, "GET", "/api/organizations/"+pathSeg(organizationID)+"/members", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListMyMemberships — Lists the caller's own organization memberships.
//
// GET /api/organization-members
func (r *OrganizationsService) ListMyMemberships(ctx context.Context, opts ...CallOption) ([]OrganizationMember, error) {
	var out []OrganizationMember
	err := r.c.call(ctx, "GET", "/api/organization-members", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListTeams — Lists an organization's teams with their member counts and the roles they hold in it.
//
// GET /api/organizations/{organizationId}/teams
//
// organizationID: The organization's id.
func (r *OrganizationsService) ListTeams(ctx context.Context, organizationID string, opts ...CallOption) ([]OrganizationsListTeamsResponseDataItem, error) {
	var out []OrganizationsListTeamsResponseDataItem
	err := r.c.call(ctx, "GET", "/api/organizations/"+pathSeg(organizationID)+"/teams", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListWorkspaces — Lists the workspaces of an organization the caller can open.
//
// GET /api/organizations/{organizationId}/workspaces
//
// organizationID: The organization's id.
func (r *OrganizationsService) ListWorkspaces(ctx context.Context, organizationID string, opts ...CallOption) ([]Workspace, error) {
	var out []Workspace
	err := r.c.call(ctx, "GET", "/api/organizations/"+pathSeg(organizationID)+"/workspaces", nil, nil, false, "data", &out, opts)
	return out, err
}

// RemoveMember — Removes a member from an organization.
//
// DELETE /api/organizations/{organizationId}/members/{memberId}
//
// organizationID: The organization's id.
// memberID: The membership's id.
func (r *OrganizationsService) RemoveMember(ctx context.Context, organizationID string, memberID string, opts ...CallOption) (*OrganizationsRemoveMemberResponse, error) {
	var out *OrganizationsRemoveMemberResponse
	err := r.c.call(ctx, "DELETE", "/api/organizations/"+pathSeg(organizationID)+"/members/"+pathSeg(memberID), nil, nil, false, "", &out, opts)
	return out, err
}

// RemoveTeam — Takes away the roles a team holds across the organization.
//
// DELETE /api/organizations/{organizationId}/teams/{teamId}
//
// organizationID: The organization's id.
// teamID: The team's id.
func (r *OrganizationsService) RemoveTeam(ctx context.Context, organizationID string, teamID string, opts ...CallOption) (*OrganizationsRemoveTeamResponse, error) {
	var out *OrganizationsRemoveTeamResponse
	err := r.c.call(ctx, "DELETE", "/api/organizations/"+pathSeg(organizationID)+"/teams/"+pathSeg(teamID), nil, nil, false, "", &out, opts)
	return out, err
}

// ResendInvitation — Sends an invitation's email again and extends it.
//
// POST /api/organizations/{organizationId}/invitations/{invitationId}/resend
//
// organizationID: The organization's id.
// invitationID: The invitation's id.
func (r *OrganizationsService) ResendInvitation(ctx context.Context, organizationID string, invitationID string, opts ...CallOption) (*OrganizationsResendInvitationResponseData, error) {
	var out *OrganizationsResendInvitationResponseData
	err := r.c.call(ctx, "POST", "/api/organizations/"+pathSeg(organizationID)+"/invitations/"+pathSeg(invitationID)+"/resend", nil, nil, false, "data", &out, opts)
	return out, err
}

// RevokeInvitation — Revokes an invitation, so its link no longer works.
//
// DELETE /api/organizations/{organizationId}/invitations/{invitationId}
//
// organizationID: The organization's id.
// invitationID: The invitation's id.
func (r *OrganizationsService) RevokeInvitation(ctx context.Context, organizationID string, invitationID string, opts ...CallOption) (*OrganizationsRevokeInvitationResponseData, error) {
	var out *OrganizationsRevokeInvitationResponseData
	err := r.c.call(ctx, "DELETE", "/api/organizations/"+pathSeg(organizationID)+"/invitations/"+pathSeg(invitationID), nil, nil, false, "data", &out, opts)
	return out, err
}

// SetMemberAdding — Chooses how the organization adds a person by email.
//
// `direct` is faster but tells whoever adds whether the address has an account.
//
// PUT /api/organizations/{organizationId}/member-adding
//
// organizationID: The organization's id.
func (r *OrganizationsService) SetMemberAdding(ctx context.Context, organizationID string, body *OrganizationsSetMemberAddingBody, opts ...CallOption) (*OrganizationsSetMemberAddingResponseData, error) {
	var out *OrganizationsSetMemberAddingResponseData
	err := r.c.call(ctx, "PUT", "/api/organizations/"+pathSeg(organizationID)+"/member-adding", nil, body, false, "data", &out, opts)
	return out, err
}

// SetReportRecipients — Allows or stops scheduled reports to addresses outside the organization.
//
// Needs `organization.update`. Switching it off stops the outside addresses of every schedule at their next send.
//
// PUT /api/organizations/{organizationId}/report-recipients
//
// organizationID: The organization's id.
func (r *OrganizationsService) SetReportRecipients(ctx context.Context, organizationID string, body *OrganizationsSetReportRecipientsBody, opts ...CallOption) (*OrganizationsSetReportRecipientsResponseData, error) {
	var out *OrganizationsSetReportRecipientsResponseData
	err := r.c.call(ctx, "PUT", "/api/organizations/"+pathSeg(organizationID)+"/report-recipients", nil, body, false, "data", &out, opts)
	return out, err
}

// Update — Changes an organization's name, look, settings or billing details.
//
// Needs `organization.update`; the billing fields need `billing.read` as well.
//
// PUT /api/organizations/{organizationId}
//
// organizationID: The organization's id.
func (r *OrganizationsService) Update(ctx context.Context, organizationID string, body *OrganizationsUpdateBody, opts ...CallOption) (*Organization, error) {
	var out *Organization
	err := r.c.call(ctx, "PUT", "/api/organizations/"+pathSeg(organizationID), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateFreezeWindow — Replaces a publish freeze window.
//
// Send the whole window, as when creating it.
//
// PATCH /api/organizations/{organizationId}/freeze-windows/{windowId}
//
// organizationID: The organization's id.
// windowID: The freeze window's id.
func (r *OrganizationsService) UpdateFreezeWindow(ctx context.Context, organizationID string, windowID string, body *FreezeWindowInput, opts ...CallOption) (*OrganizationsUpdateFreezeWindowResponseData, error) {
	var out *OrganizationsUpdateFreezeWindowResponseData
	err := r.c.call(ctx, "PATCH", "/api/organizations/"+pathSeg(organizationID)+"/freeze-windows/"+pathSeg(windowID), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateMember — Changes a member's role or seat tier.
//
// Needs `organization.manage_members`. Only an owner makes or unmakes an owner, and the last owner stays one.
//
// PATCH /api/organizations/{organizationId}/members/{memberId}
//
// organizationID: The organization's id.
// memberID: The membership's id.
func (r *OrganizationsService) UpdateMember(ctx context.Context, organizationID string, memberID string, body *OrganizationsUpdateMemberBody, opts ...CallOption) (*OrganizationMember, error) {
	var out *OrganizationMember
	err := r.c.call(ctx, "PATCH", "/api/organizations/"+pathSeg(organizationID)+"/members/"+pathSeg(memberID), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateSeatSettings — Sets the seat a newly added member gets, and the order seats are handed out in after a plan change.
//
// Needs `organization.manage_members`.
//
// PATCH /api/organizations/{organizationId}/seat-settings
//
// organizationID: The organization's id.
func (r *OrganizationsService) UpdateSeatSettings(ctx context.Context, organizationID string, body *OrganizationsUpdateSeatSettingsBody, opts ...CallOption) (*OrganizationsUpdateSeatSettingsResponseData, error) {
	var out *OrganizationsUpdateSeatSettingsResponseData
	err := r.c.call(ctx, "PATCH", "/api/organizations/"+pathSeg(organizationID)+"/seat-settings", nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateSharingSettings — Changes the organization's file-sharing defaults and limits.
//
// Needs `organization.update`.
//
// PUT /api/organizations/{organizationId}/sharing-settings
//
// organizationID: The organization's id.
func (r *OrganizationsService) UpdateSharingSettings(ctx context.Context, organizationID string, body *OrganizationsUpdateSharingSettingsBody, opts ...CallOption) (*AssetSharingSettings, error) {
	var out *AssetSharingSettings
	err := r.c.call(ctx, "PUT", "/api/organizations/"+pathSeg(organizationID)+"/sharing-settings", nil, body, false, "settings", &out, opts)
	return out, err
}

// VersionStats — Returns how much storage the organization's file versions take and how long they are kept.
//
// GET /api/organizations/{organizationId}/version-stats
//
// organizationID: The organization's id.
func (r *OrganizationsService) VersionStats(ctx context.Context, organizationID string, opts ...CallOption) (*OrganizationsVersionStatsResponseData, error) {
	var out *OrganizationsVersionStatsResponseData
	err := r.c.call(ctx, "GET", "/api/organizations/"+pathSeg(organizationID)+"/version-stats", nil, nil, false, "data", &out, opts)
	return out, err
}

// OrganizationsEmailTemplatesService — organizations.emailTemplates
type OrganizationsEmailTemplatesService struct {
	c *Client
}

func newOrganizationsEmailTemplatesService(c *Client) *OrganizationsEmailTemplatesService {
	return &OrganizationsEmailTemplatesService{c: c}
}

// List — Lists the visitor emails the organization may word itself, in each language its projects use, with the wording that is sent.
//
// A project's own wording wins over the organization's.
//
// GET /api/organizations/{organizationId}/email-templates
//
// organizationID: The organization's id.
func (r *OrganizationsEmailTemplatesService) List(ctx context.Context, organizationID string, opts ...CallOption) (*OrganizationsEmailTemplatesListResponseData, error) {
	var out *OrganizationsEmailTemplatesListResponseData
	err := r.c.call(ctx, "GET", "/api/organizations/"+pathSeg(organizationID)+"/email-templates", nil, nil, false, "data", &out, opts)
	return out, err
}

// Reset — Goes back to the platform's wording of a visitor email in one language.
//
// DELETE /api/organizations/{organizationId}/email-templates/{type}
//
// organizationID: The organization's id.
// typeArg: The email: `waitlist_confirm` or `waitlist_welcome`.
func (r *OrganizationsEmailTemplatesService) Reset(ctx context.Context, organizationID string, typeArg string, query *OrganizationsEmailTemplatesResetQuery, opts ...CallOption) (*OrganizationsEmailTemplatesResetResponse, error) {
	var out *OrganizationsEmailTemplatesResetResponse
	err := r.c.call(ctx, "DELETE", "/api/organizations/"+pathSeg(organizationID)+"/email-templates/"+pathSeg(typeArg), query, nil, false, "", &out, opts)
	return out, err
}

// Set — Saves the organization's own wording of a visitor email in one language, for every project that has none of its own.
//
// Needs `organization.update` and the plan's custom email templates. The HTML is cleaned before it is stored.
//
// PUT /api/organizations/{organizationId}/email-templates/{type}
//
// organizationID: The organization's id.
// typeArg: The email: `waitlist_confirm` or `waitlist_welcome`.
func (r *OrganizationsEmailTemplatesService) Set(ctx context.Context, organizationID string, typeArg string, body *EmailTemplateInput, opts ...CallOption) (*OrganizationsEmailTemplatesSetResponse, error) {
	var out *OrganizationsEmailTemplatesSetResponse
	err := r.c.call(ctx, "PUT", "/api/organizations/"+pathSeg(organizationID)+"/email-templates/"+pathSeg(typeArg), nil, body, false, "", &out, opts)
	return out, err
}

// PagesService — Pages
type PagesService struct {
	c *Client
}

func newPagesService(c *Client) *PagesService {
	return &PagesService{c: c}
}

// AbTestAction — Starts, updates, pauses, resumes or stops the page's A/B test.
//
// `start` and `update` put variant B live through the publish rules — with approval required the answer is a 202 waiting for it. At most 30 calls a minute.
//
// POST /api/pages/{pageId}/ab-test — Token permission: write:pages.
//
// pageID: The page's id.
func (r *PagesService) AbTestAction(ctx context.Context, pageID string, body *PagesAbTestActionBody, opts ...CallOption) (*PagesAbTestActionResponseData, error) {
	var out *PagesAbTestActionResponseData
	err := r.c.call(ctx, "POST", "/api/pages/"+pathSeg(pageID)+"/ab-test", nil, body, false, "data", &out, opts)
	return out, err
}

// Create — Creates a page in a project.
//
// Needs `page.create`. The page starts as a draft.
//
// POST /api/pages — Token permission: write:pages.
func (r *PagesService) Create(ctx context.Context, body *CreatePageInput, opts ...CallOption) (*Page, error) {
	var out *Page
	err := r.c.call(ctx, "POST", "/api/pages", nil, body, false, "data", &out, opts)
	return out, err
}

// CreateCollection — Creates a list page and a detail page that show a database's records on a project's site.
//
// Needs `page.create`, access to the workspace and project, and a database of the project's organization the caller can open; the two pages count against `max_pages_per_project`. The detail page answers at `basePath` with each record's slug appended, the list page at `basePath` + `-list`. Both are drafts until published.
//
// POST /api/collection-pages/generate
func (r *PagesService) CreateCollection(ctx context.Context, body *PagesCreateCollectionBody, opts ...CallOption) (*PagesCreateCollectionResponse, error) {
	var out *PagesCreateCollectionResponse
	err := r.c.call(ctx, "POST", "/api/collection-pages/generate", nil, body, false, "", &out, opts)
	return out, err
}

// CreatePreview — Returns a preview link to the page's draft, creating one when it has none.
//
// POST /api/pages/{pageId}/preview — Token permission: write:pages.
//
// pageID: The page's id.
func (r *PagesService) CreatePreview(ctx context.Context, pageID string, query *PagesCreatePreviewQuery, opts ...CallOption) (*PagesCreatePreviewResponseData, error) {
	var out *PagesCreatePreviewResponseData
	err := r.c.call(ctx, "POST", "/api/pages/"+pathSeg(pageID)+"/preview", query, nil, false, "data", &out, opts)
	return out, err
}

// CreateVersion — Saves a named version of a page.
//
// POST /api/pages/{pageId}/versions — Token permission: write:pages.
//
// pageID: The page's id.
func (r *PagesService) CreateVersion(ctx context.Context, pageID string, body *PagesCreateVersionBody, opts ...CallOption) (*PageVersion, error) {
	var out *PageVersion
	err := r.c.call(ctx, "POST", "/api/pages/"+pathSeg(pageID)+"/versions", nil, body, false, "data", &out, opts)
	return out, err
}

// Delete — Moves a page to the trash.
//
// DELETE /api/pages/{pageId} — Token permission: write:pages.
//
// pageID: The page's id.
func (r *PagesService) Delete(ctx context.Context, pageID string, opts ...CallOption) (*PagesDeleteResponseData, error) {
	var out *PagesDeleteResponseData
	err := r.c.call(ctx, "DELETE", "/api/pages/"+pathSeg(pageID), nil, nil, false, "data", &out, opts)
	return out, err
}

// DeleteVersion — Deletes a saved version of a page.
//
// DELETE /api/pages/{pageId}/versions/{versionId} — Token permission: write:pages.
//
// pageID: The page's id.
// versionID: The version's id.
func (r *PagesService) DeleteVersion(ctx context.Context, pageID string, versionID string, opts ...CallOption) (*PagesDeleteVersionResponse, error) {
	var out *PagesDeleteVersionResponse
	err := r.c.call(ctx, "DELETE", "/api/pages/"+pathSeg(pageID)+"/versions/"+pathSeg(versionID), nil, nil, false, "", &out, opts)
	return out, err
}

// Discard — Throws away a published page's unpublished changes, putting the draft back to what is live.
//
// POST /api/pages/{pageId}/discard — Token permission: write:pages.
//
// pageID: The page's id.
func (r *PagesService) Discard(ctx context.Context, pageID string, opts ...CallOption) (*Page, error) {
	var out *Page
	err := r.c.call(ctx, "POST", "/api/pages/"+pathSeg(pageID)+"/discard", nil, nil, false, "data", &out, opts)
	return out, err
}

// Duplicate — Copies a page within its project as a new draft.
//
// Needs `page.create`. The copy gets a unique name and path, is never the home page, has no translations, versions or A/B test, and starts as a draft.
//
// POST /api/pages/{pageId}/duplicate — Token permission: write:pages.
//
// pageID: The page's id.
func (r *PagesService) Duplicate(ctx context.Context, pageID string, body *PagesDuplicateBody, query *PagesDuplicateQuery, opts ...CallOption) (*Page, error) {
	var out *Page
	err := r.c.call(ctx, "POST", "/api/pages/"+pathSeg(pageID)+"/duplicate", query, body, false, "data.page", &out, opts)
	return out, err
}

// Get — Returns a page with its draft tree and whether a publish is waiting for approval.
//
// GET /api/pages/{pageId} — Token permission: read:pages.
//
// pageID: The page's id.
func (r *PagesService) Get(ctx context.Context, pageID string, opts ...CallOption) (*PagesGetResponseData, error) {
	var out *PagesGetResponseData
	err := r.c.call(ctx, "GET", "/api/pages/"+pathSeg(pageID), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetAbTest — Returns the page's A/B test on the live site: status, goal, split and each variant's results.
//
// GET /api/pages/{pageId}/ab-test — Token permission: read:pages.
//
// pageID: The page's id.
func (r *PagesService) GetAbTest(ctx context.Context, pageID string, opts ...CallOption) (*PagesGetAbTestResponseData, error) {
	var out *PagesGetAbTestResponseData
	err := r.c.call(ctx, "GET", "/api/pages/"+pathSeg(pageID)+"/ab-test", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetComponents — Returns a page's draft component tree.
//
// GET /api/pages/{pageId}/components — Token permission: read:pages.
//
// pageID: The page's id.
func (r *PagesService) GetComponents(ctx context.Context, pageID string, opts ...CallOption) ([]any, error) {
	var out []any
	err := r.c.call(ctx, "GET", "/api/pages/"+pathSeg(pageID)+"/components", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetVersion — Returns one version of a page, with its tree.
//
// GET /api/pages/{pageId}/versions/{versionId} — Token permission: read:pages.
//
// pageID: The page's id.
// versionID: The version's id.
func (r *PagesService) GetVersion(ctx context.Context, pageID string, versionID string, opts ...CallOption) (*PageVersion, error) {
	var out *PageVersion
	err := r.c.call(ctx, "GET", "/api/pages/"+pathSeg(pageID)+"/versions/"+pathSeg(versionID), nil, nil, false, "data", &out, opts)
	return out, err
}

// ListTranslations — Lists the language versions of a page, the page itself included.
//
// GET /api/pages/{pageId}/translations — Token permission: read:pages.
//
// pageID: The page's id.
func (r *PagesService) ListTranslations(ctx context.Context, pageID string, opts ...CallOption) ([]PageTranslationVariant, error) {
	var out []PageTranslationVariant
	err := r.c.call(ctx, "GET", "/api/pages/"+pathSeg(pageID)+"/translations", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListVersions — Lists a page's saved versions, newest first, without their trees.
//
// GET /api/pages/{pageId}/versions — Token permission: read:pages.
//
// pageID: The page's id.
func (r *PagesService) ListVersions(ctx context.Context, pageID string, opts ...CallOption) ([]PageVersionSummary, error) {
	var out []PageVersionSummary
	err := r.c.call(ctx, "GET", "/api/pages/"+pathSeg(pageID)+"/versions", nil, nil, false, "data", &out, opts)
	return out, err
}

// RestoreVersion — Puts a version back into the page's draft, and publishes it with `publish: true`.
//
// Needs `version.restore`. Publishing goes through the publish rules — with approval required the answer says `pendingApproval`.
//
// POST /api/pages/{pageId}/versions/{versionId}/restore — Token permission: write:pages.
//
// pageID: The page's id.
// versionID: The version's id.
func (r *PagesService) RestoreVersion(ctx context.Context, pageID string, versionID string, body *PagesRestoreVersionBody, opts ...CallOption) (*Page, error) {
	var out *Page
	err := r.c.call(ctx, "POST", "/api/pages/"+pathSeg(pageID)+"/versions/"+pathSeg(versionID)+"/restore", nil, body, false, "data", &out, opts)
	return out, err
}

// RestoreVersionLegacy — Puts a version back into the page's draft — the older address of `pages.restoreVersion`, with the same checks.
//
// POST /api/pages/{pageId}/versions/{versionId} — Token permission: write:pages.
//
// pageID: The page's id.
// versionID: The version's id.
//
// Deprecated: Use pages.restoreVersion.
func (r *PagesService) RestoreVersionLegacy(ctx context.Context, pageID string, versionID string, body *PagesRestoreVersionLegacyBody, opts ...CallOption) (*Page, error) {
	deprecated("pages.restoreVersionLegacy")
	var out *Page
	err := r.c.call(ctx, "POST", "/api/pages/"+pathSeg(pageID)+"/versions/"+pathSeg(versionID), nil, body, false, "data", &out, opts)
	return out, err
}

// RevokePreview — Revokes the page's preview link.
//
// DELETE /api/pages/{pageId}/preview — Token permission: write:pages.
//
// pageID: The page's id.
func (r *PagesService) RevokePreview(ctx context.Context, pageID string, opts ...CallOption) (*PagesRevokePreviewResponse, error) {
	var out *PagesRevokePreviewResponse
	err := r.c.call(ctx, "DELETE", "/api/pages/"+pathSeg(pageID)+"/preview", nil, nil, false, "", &out, opts)
	return out, err
}

// Scan — Checks one page for outdated content and broken links.
//
// A published page is checked as visitors see it and the findings are kept for Site health; a draft is checked as it is. At most 4 calls an hour per page.
//
// POST /api/pages/{pageId}/scan — Token permission: write:pages.
//
// pageID: The page's id.
func (r *PagesService) Scan(ctx context.Context, pageID string, opts ...CallOption) (*PagesScanResponseData, error) {
	var out *PagesScanResponseData
	err := r.c.call(ctx, "POST", "/api/pages/"+pathSeg(pageID)+"/scan", nil, nil, false, "data", &out, opts)
	return out, err
}

// SetCollectionMode — Makes a page a collection list or detail page of a database, or a static page again.
//
// Needs `page.update`. `list` and `detail` need `databaseId` and a `slugField` that is a column of it; a list page may name its `detailPageId` (a page of the same project, made a detail page of the same database if it is not one yet). `none` removes every collection setting.
//
// POST /api/collection-pages/convert
func (r *PagesService) SetCollectionMode(ctx context.Context, body *PagesSetCollectionModeBody, opts ...CallOption) (*Page, error) {
	var out *Page
	err := r.c.call(ctx, "POST", "/api/collection-pages/convert", nil, body, false, "data", &out, opts)
	return out, err
}

// Translate — Creates a language version of a page, machine-translated by default.
//
// Needs `page.create`. The new page joins the original's translation group; with `translate: false` its tree is copied as it is.
//
// POST /api/pages/{pageId}/translate — Token permission: write:pages.
//
// pageID: The page's id.
func (r *PagesService) Translate(ctx context.Context, pageID string, body *PagesTranslateBody, opts ...CallOption) (*Page, error) {
	var out *Page
	err := r.c.call(ctx, "POST", "/api/pages/"+pathSeg(pageID)+"/translate", nil, body, false, "data", &out, opts)
	return out, err
}

// Unverify — Removes the page's content-checked mark.
//
// DELETE /api/pages/{pageId}/verify — Token permission: write:pages.
//
// pageID: The page's id.
func (r *PagesService) Unverify(ctx context.Context, pageID string, opts ...CallOption) (*PagesUnverifyResponse, error) {
	var out *PagesUnverifyResponse
	err := r.c.call(ctx, "DELETE", "/api/pages/"+pathSeg(pageID)+"/verify", nil, nil, false, "", &out, opts)
	return out, err
}

// Update — Changes a page's draft — its tree, name, path or SEO fields — and publishes or schedules it with `status`.
//
// Needs `page.update`. Everything but `status: "published"` changes the draft only. Publishing applies freeze windows and the organization's approval requirement: when approval is required the answer is a 202 with `pendingApproval: true` and the page stays as it is live. A password or sign-in requirement added in the draft applies to the live page at once.
//
// PUT /api/pages/{pageId} — Token permission: write:pages.
//
// pageID: The page's id.
func (r *PagesService) Update(ctx context.Context, pageID string, body *UpdatePageInput, opts ...CallOption) (*Page, error) {
	var out *Page
	err := r.c.call(ctx, "PUT", "/api/pages/"+pathSeg(pageID), nil, body, false, "data", &out, opts)
	return out, err
}

// Verify — Marks the page's content as checked by the caller today.
//
// POST /api/pages/{pageId}/verify — Token permission: write:pages.
//
// pageID: The page's id.
func (r *PagesService) Verify(ctx context.Context, pageID string, opts ...CallOption) (*PagesVerifyResponseData, error) {
	var out *PagesVerifyResponseData
	err := r.c.call(ctx, "POST", "/api/pages/"+pathSeg(pageID)+"/verify", nil, nil, false, "data", &out, opts)
	return out, err
}

// PdfService — PDF: Tools for the PDF files of an organisation's asset library — editing, merging, splitting, compressing, password protection, summaries and explanations, form filling — and e-signature requests that send a PDF to be signed. A tool reads its source by asset id; one that produces a PDF saves it as a new file of the library, or replaces the source with `overwrite`.
type PdfService struct {
	c *Client

	Esign           *PdfEsignService
	RedactionAudits *PdfRedactionAuditsService
}

func newPdfService(c *Client) *PdfService {
	return &PdfService{c: c, Esign: newPdfEsignService(c), RedactionAudits: newPdfRedactionAuditsService(c)}
}

// AIFill — Proposes values for a PDF form's fields from a description, with the platform's AI.
//
// Needs PDF AI on the organisation's plan (403 when switched off) and its AI allowance (429); the file must be within the plan's PDF size (413). Nothing is written: apply the values with pdf.edit (`set_form_field`). 422 for a PDF without form fields.
//
// POST /api/pdf/ai-fill
func (r *PdfService) AIFill(ctx context.Context, body *PdfAiFillBody, opts ...CallOption) (*PdfAiFillResponseData, error) {
	var out *PdfAiFillResponseData
	err := r.c.call(ctx, "POST", "/api/pdf/ai-fill", nil, body, false, "data", &out, opts)
	return out, err
}

// Compress — Writes a PDF again compactly, dropping what earlier edits left behind.
//
// Needs the PDF editor on the organisation's plan (403 when it is switched off, 429 when the plan leaves it out), a file within the plan's PDF size (413; 25 MB unless the plan says otherwise), room in its storage (403) and AI allowance left (429). Images are not resampled.
//
// POST /api/pdf/compress
func (r *PdfService) Compress(ctx context.Context, body *PdfCompressBody, opts ...CallOption) (*Asset, error) {
	var out *Asset
	err := r.c.call(ctx, "POST", "/api/pdf/compress", nil, body, false, "data.asset", &out, opts)
	return out, err
}

// DetectSignatureLines — Finds where a PDF asks for a signature, from words like "Signature" or "Unterschrift" and lines to sign on.
//
// Needs PDF AI on the organisation's plan (403 when switched off) and its AI allowance (429); the file must be within the plan's PDF size (413). A heuristic over the text layer: it misses scanned pages without one.
//
// POST /api/pdf/detect-signature-lines
func (r *PdfService) DetectSignatureLines(ctx context.Context, body *PdfDetectSignatureLinesBody, opts ...CallOption) ([]PdfDetectSignatureLinesResponseDataLinesItem, error) {
	var out []PdfDetectSignatureLinesResponseDataLinesItem
	err := r.c.call(ctx, "POST", "/api/pdf/detect-signature-lines", nil, body, false, "data.lines", &out, opts)
	return out, err
}

// Edit — Applies a list of edits to a PDF — text, images, signatures, highlights, drawings, redactions, page changes, form fields, watermarks — in order.
//
// Needs the PDF editor on the organisation's plan (403 when it is switched off, 429 when the plan leaves it out), a file within the plan's PDF size (413; 25 MB unless the plan says otherwise), room in its storage (403) and AI allowance left (429). At most 500 edits. A redaction that is to remove the text and not only cover it pairs each `redact` with a `flatten_page` of that page; `redaction` then files the redaction record.
//
// POST /api/pdf/edit
func (r *PdfService) Edit(ctx context.Context, body *PdfEditBody, opts ...CallOption) (*Asset, error) {
	var out *Asset
	err := r.c.call(ctx, "POST", "/api/pdf/edit", nil, body, false, "data.asset", &out, opts)
	return out, err
}

// Explain — Explains one page of a PDF in plain language with the platform's AI.
//
// Needs PDF AI on the organisation's plan (403 when switched off) and its AI allowance (429); the file must be within the plan's PDF size (413). 422 when the page has no text layer. Nothing is stored.
//
// POST /api/pdf/explain
func (r *PdfService) Explain(ctx context.Context, body *PdfExplainBody, opts ...CallOption) (*PdfExplainResponseData, error) {
	var out *PdfExplainResponseData
	err := r.c.call(ctx, "POST", "/api/pdf/explain", nil, body, false, "data", &out, opts)
	return out, err
}

// Merge — Merges 2 to 20 PDFs, in the order given, into one.
//
// Needs the PDF editor on the organisation's plan (403 when it is switched off, 429 when the plan leaves it out), a file within the plan's PDF size (413; 25 MB unless the plan says otherwise), room in its storage (403) and AI allowance left (429). With `overwrite` the first file is replaced.
//
// POST /api/pdf/merge
func (r *PdfService) Merge(ctx context.Context, body *PdfMergeBody, opts ...CallOption) (*Asset, error) {
	var out *Asset
	err := r.c.call(ctx, "POST", "/api/pdf/merge", nil, body, false, "data.asset", &out, opts)
	return out, err
}

// Protect — Encrypts a PDF with AES-256 so it opens only with the password.
//
// Needs the PDF editor on the organisation's plan (403 when it is switched off, 429 when the plan leaves it out), a file within the plan's PDF size (413; 25 MB unless the plan says otherwise), room in its storage (403) and AI allowance left (429).
//
// POST /api/pdf/protect
func (r *PdfService) Protect(ctx context.Context, body *PdfProtectBody, opts ...CallOption) (*Asset, error) {
	var out *Asset
	err := r.c.call(ctx, "POST", "/api/pdf/protect", nil, body, false, "data.asset", &out, opts)
	return out, err
}

// Split — Makes a PDF of the pages given, in that order.
//
// Needs the PDF editor on the organisation's plan (403 when it is switched off, 429 when the plan leaves it out), a file within the plan's PDF size (413; 25 MB unless the plan says otherwise), room in its storage (403) and AI allowance left (429).
//
// POST /api/pdf/split
func (r *PdfService) Split(ctx context.Context, body *PdfSplitBody, opts ...CallOption) (*Asset, error) {
	var out *Asset
	err := r.c.call(ctx, "POST", "/api/pdf/split", nil, body, false, "data.asset", &out, opts)
	return out, err
}

// Summarize — Summarises a PDF in plain prose with the platform's AI.
//
// Needs PDF AI on the organisation's plan (403 when switched off) and its AI allowance (429); the file must be within the plan's PDF size (413). Reads the text of the first 30 pages (`truncated` says when that was not all). 422 for a scan without a text layer. Nothing is stored.
//
// POST /api/pdf/summarize
func (r *PdfService) Summarize(ctx context.Context, body *PdfSummarizeBody, opts ...CallOption) (*PdfSummarizeResponseData, error) {
	var out *PdfSummarizeResponseData
	err := r.c.call(ctx, "POST", "/api/pdf/summarize", nil, body, false, "data", &out, opts)
	return out, err
}

// Unlock — Removes the password of a protected PDF.
//
// Needs the PDF editor on the organisation's plan (403 when it is switched off, 429 when the plan leaves it out), a file within the plan's PDF size (413; 25 MB unless the plan says otherwise), room in its storage (403) and AI allowance left (429). 403 for a wrong password.
//
// POST /api/pdf/unlock
func (r *PdfService) Unlock(ctx context.Context, body *PdfUnlockBody, opts ...CallOption) (*Asset, error) {
	var out *Asset
	err := r.c.call(ctx, "POST", "/api/pdf/unlock", nil, body, false, "data.asset", &out, opts)
	return out, err
}

// PdfEsignService — pdf.esign
type PdfEsignService struct {
	c *Client
}

func newPdfEsignService(c *Client) *PdfEsignService {
	return &PdfEsignService{c: c}
}

// AuditTrail — Downloads a signature request's audit trail as a PDF: the request, its signers and every event with time, address and browser.
//
// Only the person who sent the request. Teammates who see a shared request cannot download its evidence.
//
// GET /api/pdf/esign/requests/{requestId}/audit
//
// requestID: The signature request's id.
func (r *PdfEsignService) AuditTrail(ctx context.Context, requestID string, query *PdfEsignAuditTrailQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/pdf/esign/requests/"+pathSeg(requestID)+"/audit", query, nil, false, "", &out, opts)
	return out, err
}

// Cancel — Cancels a signature request; its signing links stop working at once.
//
// Only the person who sent the request. Cancelling again changes nothing.
//
// POST /api/pdf/esign/requests/{requestId}/cancel
//
// requestID: The signature request's id.
func (r *PdfEsignService) Cancel(ctx context.Context, requestID string, body *PdfEsignCancelBody, opts ...CallOption) (*PdfEsignCancelResponseData, error) {
	var out *PdfEsignCancelResponseData
	err := r.c.call(ctx, "POST", "/api/pdf/esign/requests/"+pathSeg(requestID)+"/cancel", nil, body, false, "data", &out, opts)
	return out, err
}

// Create — Sends a PDF of the asset library to up to 20 people to sign, one after the other or all at once.
//
// The invitations are emailed (sequential: the next signer once the one before signed). Each signer's `signUrl` is a secret link to their signing page.
//
// POST /api/pdf/esign/request
func (r *PdfEsignService) Create(ctx context.Context, body *PdfEsignCreateBody, opts ...CallOption) (*PdfEsignCreateResponseData, error) {
	var out *PdfEsignCreateResponseData
	err := r.c.call(ctx, "POST", "/api/pdf/esign/request", nil, body, false, "data", &out, opts)
	return out, err
}

// List — Lists the signature requests the caller can see in an organisation, newest first.
//
// The caller's own, those shared with a workspace they can open, and those shared with the organisation; `isMine` tells which are theirs.
//
// GET /api/pdf/esign/requests
func (r *PdfEsignService) List(ctx context.Context, query *PdfEsignListQuery, opts ...CallOption) ([]SignatureRequestSummary, error) {
	var out []SignatureRequestSummary
	err := r.c.call(ctx, "GET", "/api/pdf/esign/requests", query, nil, false, "data.requests", &out, opts)
	return out, err
}

// Resend — Emails the invitation of a pending signature request again, with the same link.
//
// Only the person who sent the request.
//
// POST /api/pdf/esign/requests/{requestId}/resend
//
// requestID: The signature request's id.
func (r *PdfEsignService) Resend(ctx context.Context, requestID string, body *PdfEsignResendBody, opts ...CallOption) (*PdfEsignResendResponseData, error) {
	var out *PdfEsignResendResponseData
	err := r.c.call(ctx, "POST", "/api/pdf/esign/requests/"+pathSeg(requestID)+"/resend", nil, body, false, "data", &out, opts)
	return out, err
}

// SetVisibility — Changes who else sees a signature request: nobody, a workspace or the whole organisation.
//
// Only the person who sent the request. `workspace` needs a workspace the caller can open.
//
// PATCH /api/pdf/esign/requests/{requestId}
//
// requestID: The signature request's id.
func (r *PdfEsignService) SetVisibility(ctx context.Context, requestID string, body *PdfEsignSetVisibilityBody, opts ...CallOption) (*PdfEsignSetVisibilityResponseData, error) {
	var out *PdfEsignSetVisibilityResponseData
	err := r.c.call(ctx, "PATCH", "/api/pdf/esign/requests/"+pathSeg(requestID), nil, body, false, "data", &out, opts)
	return out, err
}

// PdfRedactionAuditsService — pdf.redactionAudits
type PdfRedactionAuditsService struct {
	c *Client
}

func newPdfRedactionAuditsService(c *Client) *PdfRedactionAuditsService {
	return &PdfRedactionAuditsService{c: c}
}

// Get — Returns the record of a redaction: which file, how many spans of which kinds, and the SHA-256 before and after.
//
// Needs `asset:redact` and membership of the record's organisation. With `pdf=1` it answers a one-page PDF of the record to download instead.
//
// GET /api/pdf/redaction-audits/{auditId}
//
// auditID: The redaction record's id.
func (r *PdfRedactionAuditsService) Get(ctx context.Context, auditID string, query *PdfRedactionAuditsGetQuery, opts ...CallOption) (*RedactionAudit, error) {
	var out *RedactionAudit
	err := r.c.call(ctx, "GET", "/api/pdf/redaction-audits/"+pathSeg(auditID), query, nil, false, "data.audit", &out, opts)
	return out, err
}

// PermissionsService — Permissions
type PermissionsService struct {
	c *Client
}

func newPermissionsService(c *Client) *PermissionsService {
	return &PermissionsService{c: c}
}

// Create — Creates a custom permission in an organization, for its own roles to carry.
//
// POST /api/permissions
func (r *PermissionsService) Create(ctx context.Context, body *CreatePermissionInput, opts ...CallOption) (*Permission, error) {
	var out *Permission
	err := r.c.call(ctx, "POST", "/api/permissions", nil, body, false, "", &out, opts)
	return out, err
}

// Delete — Deletes a custom permission.
//
// DELETE /api/permissions/{permissionId}
//
// permissionID: The custom permission's id.
func (r *PermissionsService) Delete(ctx context.Context, permissionID string, opts ...CallOption) (*PermissionsDeleteResponse, error) {
	var out *PermissionsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/permissions/"+pathSeg(permissionID), nil, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the platform's permissions and the custom permissions of the caller's organizations.
//
// With `organizationId`, only what that organization can hold and the caller could grant.
//
// GET /api/permissions
func (r *PermissionsService) List(ctx context.Context, query *PermissionsListQuery, opts ...CallOption) ([]Permission, error) {
	var out []Permission
	err := r.c.call(ctx, "GET", "/api/permissions", query, nil, false, "data", &out, opts)
	return out, err
}

// Update — Changes a custom permission's label, category or description.
//
// PUT /api/permissions/{permissionId}
//
// permissionID: The custom permission's id.
func (r *PermissionsService) Update(ctx context.Context, permissionID string, body *PermissionsUpdateBody, opts ...CallOption) (*PermissionsUpdateResponse, error) {
	var out *PermissionsUpdateResponse
	err := r.c.call(ctx, "PUT", "/api/permissions/"+pathSeg(permissionID), nil, body, false, "", &out, opts)
	return out, err
}

// PrivacyRequestsService — Privacy requests: Requests by a person who is not a platform user — a contact, a form sender, a guest — to see (access, Art. 15 GDPR) or erase (erasure, Art. 17) what an organisation holds about them, named by their email addresses. A request always covers the whole organisation, hub by hub, so the permissions are asked at organisation level: `privacy:read` sees requests with the addresses masked, `privacy:manage` everything and runs them. Needs privacy requests on the plan. Every step is audited. A request of an organisation where the caller holds no privacy permission answers 404.
type PrivacyRequestsService struct {
	c *Client
}

func newPrivacyRequestsService(c *Client) *PrivacyRequestsService {
	return &PrivacyRequestsService{c: c}
}

// Cancel — Cancels a privacy request that has not run yet.
//
// Needs `privacy:manage` in the request's organisation. Only a request in `draft` or `in_review`; otherwise 409.
//
// PATCH /api/privacy-requests/{requestId}
//
// requestID: The privacy request's id.
func (r *PrivacyRequestsService) Cancel(ctx context.Context, requestID string, body *PrivacyRequestsCancelBody, opts ...CallOption) (*PrivacyRequest, error) {
	var out *PrivacyRequest
	err := r.c.call(ctx, "PATCH", "/api/privacy-requests/"+pathSeg(requestID), nil, body, false, "data.request", &out, opts)
	return out, err
}

// Create — Records an access or erasure request for one to ten email addresses.
//
// Needs `privacy:manage`. Nothing is counted yet: that is privacyRequests.preview. It is due 30 days after it was received. Answers 201.
//
// POST /api/privacy-requests
func (r *PrivacyRequestsService) Create(ctx context.Context, body *PrivacyRequestsCreateBody, opts ...CallOption) (*PrivacyRequest, error) {
	var out *PrivacyRequest
	err := r.c.call(ctx, "POST", "/api/privacy-requests", nil, body, false, "data.request", &out, opts)
	return out, err
}

// DownloadExport — Downloads the export of an access request as JSON, or laid out for the person as HTML or PDF.
//
// Needs `privacy:manage` in the request's organisation. — never `privacy:read`: this is the person's data. Every download is audited and counted. 404 before one was generated; 410 once it expired.
//
// GET /api/privacy-requests/{requestId}/export
//
// requestID: The privacy request's id.
func (r *PrivacyRequestsService) DownloadExport(ctx context.Context, requestID string, query *PrivacyRequestsDownloadExportQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/privacy-requests/"+pathSeg(requestID)+"/export", query, nil, false, "", &out, opts)
	return out, err
}

// Erase — Runs the erasure of an erasure request, hub by hub.
//
// Needs `privacy:manage` in the request's organisation. Runs only once the requester's identity was recorded as verified (privacyRequests.verifyIdentity) — otherwise 409. It cannot be undone: `confirm` must be the request's first address. A step that fails returns the request to `in_review`, and running it again repeats only the steps not done; what must be kept (signed evidence, finalized invoices, mail for a person to review …) is counted under `retained`. A completed request answers its result again and changes nothing. 409 while it is running.
//
// POST /api/privacy-requests/{requestId}/erase
//
// requestID: The privacy request's id.
func (r *PrivacyRequestsService) Erase(ctx context.Context, requestID string, body *PrivacyRequestsEraseBody, opts ...CallOption) (*PrivacyRequestsEraseResponseData, error) {
	var out *PrivacyRequestsEraseResponseData
	err := r.c.call(ctx, "POST", "/api/privacy-requests/"+pathSeg(requestID)+"/erase", nil, body, false, "data", &out, opts)
	return out, err
}

// GenerateExport — Generates the export of an access request: what every hub holds about the addresses, in one document.
//
// Needs `privacy:manage` in the request's organisation. Runs only once the requester's identity was recorded as verified (privacyRequests.verifyIdentity) — otherwise 409. The export is kept sealed for seven days and the request is completed; its creator is told. When a hub cannot be read nothing is stored: 502 `privacyHubsFailed` with the hubs named in `hubs`.
//
// POST /api/privacy-requests/{requestId}/export
//
// requestID: The privacy request's id.
func (r *PrivacyRequestsService) GenerateExport(ctx context.Context, requestID string, opts ...CallOption) (*PrivacyRequest, error) {
	var out *PrivacyRequest
	err := r.c.call(ctx, "POST", "/api/privacy-requests/"+pathSeg(requestID)+"/export", nil, nil, false, "data.request", &out, opts)
	return out, err
}

// Get — Returns one privacy request, with its counts, its result and its latest export.
//
// Needs `privacy:read`; the addresses and the identity note only with `privacy:manage`.
//
// GET /api/privacy-requests/{requestId}
//
// requestID: The privacy request's id.
func (r *PrivacyRequestsService) Get(ctx context.Context, requestID string, opts ...CallOption) (*PrivacyRequestsGetResponseData, error) {
	var out *PrivacyRequestsGetResponseData
	err := r.c.call(ctx, "GET", "/api/privacy-requests/"+pathSeg(requestID), nil, nil, false, "data", &out, opts)
	return out, err
}

// List — Lists an organisation's privacy requests.
//
// Needs `privacy:read`; the addresses are masked without `privacy:manage`.
//
// GET /api/privacy-requests
func (r *PrivacyRequestsService) List(ctx context.Context, query *PrivacyRequestsListQuery, opts ...CallOption) (*PrivacyRequestsListResponseData, error) {
	var out *PrivacyRequestsListResponseData
	err := r.c.call(ctx, "GET", "/api/privacy-requests", query, nil, false, "data", &out, opts)
	return out, err
}

// Preview — Counts, per hub, what the organisation holds about the request's addresses.
//
// Needs `privacy:manage` in the request's organisation. Counts only, never contents; counted again on every call. Moves a `draft` to `in_review`.
//
// POST /api/privacy-requests/{requestId}/preview
//
// requestID: The privacy request's id.
func (r *PrivacyRequestsService) Preview(ctx context.Context, requestID string, opts ...CallOption) (*PrivacyRequest, error) {
	var out *PrivacyRequest
	err := r.c.call(ctx, "POST", "/api/privacy-requests/"+pathSeg(requestID)+"/preview", nil, nil, false, "data.request", &out, opts)
	return out, err
}

// VerifyIdentity — Records how the requester's identity was verified.
//
// Needs `privacy:manage` in the request's organisation. Required before the export is generated or the erasure runs. An open request may record it again as a correction; one running or answered only while it has none; a cancelled one never (409). The audit keeps the method, never the note.
//
// POST /api/privacy-requests/{requestId}/identity
//
// requestID: The privacy request's id.
func (r *PrivacyRequestsService) VerifyIdentity(ctx context.Context, requestID string, body *PrivacyRequestsVerifyIdentityBody, opts ...CallOption) (*PrivacyRequest, error) {
	var out *PrivacyRequest
	err := r.c.call(ctx, "POST", "/api/privacy-requests/"+pathSeg(requestID)+"/identity", nil, body, false, "data.request", &out, opts)
	return out, err
}

// ProductsService — Products: The organisation's product catalogue — products, categories, variants, price lists, the CSV import and export — its stock, and its offers in the marketplace. Every operation works in one organisation (`organizationId`) and needs the role's `products:*` permissions held at organisation level.
type ProductsService struct {
	c *Client

	Categories     *ProductsCategoriesService
	CustomerGroups *ProductsCustomerGroupsService
	Inventory      *ProductsInventoryService
	Marketplace    *ProductsMarketplaceService
	PriceLists     *ProductsPriceListsService
	Settings       *ProductsSettingsService
	Variants       *ProductsVariantsService
}

func newProductsService(c *Client) *ProductsService {
	return &ProductsService{c: c, Categories: newProductsCategoriesService(c), CustomerGroups: newProductsCustomerGroupsService(c), Inventory: newProductsInventoryService(c), Marketplace: newProductsMarketplaceService(c), PriceLists: newProductsPriceListsService(c), Settings: newProductsSettingsService(c), Variants: newProductsVariantsService(c)}
}

// Create — Creates a product.
//
// Needs `products:create`. Counts against the plan's product limit (archived ones do not).
//
// POST /api/products — Token permission: write:products.
func (r *ProductsService) Create(ctx context.Context, body *ProductInput, query *ProductsCreateQuery, opts ...CallOption) (*Product, error) {
	var out *Product
	err := r.c.call(ctx, "POST", "/api/products", query, body, false, "product", &out, opts)
	return out, err
}

// Delete — Moves a product to the trash.
//
// Needs `products:delete`. Allowed while documents name the product: an invoicing line keeps its own copy and only loses the link. Its marketplace offer pauses at once; its number stays reserved.
//
// DELETE /api/products/{productId} — Token permission: write:products.
//
// productID: The product.
func (r *ProductsService) Delete(ctx context.Context, productID string, query *ProductsDeleteQuery, opts ...CallOption) (*ProductsDeleteResponse, error) {
	var out *ProductsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/products/"+pathSeg(productID), query, nil, false, "", &out, opts)
	return out, err
}

// Export — Downloads the products the caller sees as a CSV file, filtered like the list.
//
// Needs `products:export`. Semicolon-separated with a BOM; the category as its path; at most 20,000 rows, by name.
//
// GET /api/products/export — Token permission: read:products.
func (r *ProductsService) Export(ctx context.Context, query *ProductsExportQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/products/export", query, nil, false, "", &out, opts)
	return out, err
}

// Get — One product.
//
// GET /api/products/{productId} — Token permission: read:products.
//
// productID: The product.
func (r *ProductsService) Get(ctx context.Context, productID string, query *ProductsGetQuery, opts ...CallOption) (*Product, error) {
	var out *Product
	err := r.c.call(ctx, "GET", "/api/products/"+pathSeg(productID), query, nil, false, "product", &out, opts)
	return out, err
}

// Import — Imports products from a CSV file: a preview first, then the write with `apply`.
//
// Needs `products:create`; a file that updates existing products (matched by number) also needs `products:update`. With `apply`, 422 `productImportInvalid` while rows have errors unless `skipInvalid`, and the plan's product limit applies.
//
// POST /api/products/import — Token permission: write:products.
func (r *ProductsService) Import(ctx context.Context, body *ProductsImportBody, query *ProductsImportQuery, opts ...CallOption) (*ProductsImportResponse, error) {
	var out *ProductsImportResponse
	err := r.c.call(ctx, "POST", "/api/products/import", query, body, false, "", &out, opts)
	return out, err
}

// List — Lists one page of the products the caller sees, with the total and how many sit in each category.
//
// Needs `products:read`.
//
// GET /api/products — Token permission: read:products.
func (r *ProductsService) List(ctx context.Context, query *ProductsListQuery, opts ...CallOption) ([]Product, error) {
	var out []Product
	err := r.c.call(ctx, "GET", "/api/products", query, nil, false, "products", &out, opts)
	return out, err
}

// NextNumber — Suggests a number for a new product: the highest in use with Settings → Products' prefix, plus one.
//
// Needs `products:create`. Not a reservation — creating with a number already taken answers 409 `productNumberTaken`.
//
// GET /api/products/next-number — Token permission: read:products.
func (r *ProductsService) NextNumber(ctx context.Context, query *ProductsNextNumberQuery, opts ...CallOption) (string, error) {
	var out string
	err := r.c.call(ctx, "GET", "/api/products/next-number", query, nil, false, "number", &out, opts)
	return out, err
}

// Update — Changes a product, or archives and restores it with `archived`.
//
// Needs `products:update`. Restoring counts against the plan's product limit again. Who sees it (`visibility`, `workspaceId`) is changed only by its creator or a holder of `products:delete`.
//
// PATCH /api/products/{productId} — Token permission: write:products.
//
// productID: The product.
func (r *ProductsService) Update(ctx context.Context, productID string, body *ProductsUpdateBody, query *ProductsUpdateQuery, opts ...CallOption) (*Product, error) {
	var out *Product
	err := r.c.call(ctx, "PATCH", "/api/products/"+pathSeg(productID), query, body, false, "product", &out, opts)
	return out, err
}

// ProductsCategoriesService — products.categories
type ProductsCategoriesService struct {
	c *Client
}

func newProductsCategoriesService(c *Client) *ProductsCategoriesService {
	return &ProductsCategoriesService{c: c}
}

// Create — Creates a product category, at the end of its parent's children unless `position` says otherwise.
//
// Needs `products:create`.
//
// POST /api/products/categories — Token permission: write:products.
func (r *ProductsCategoriesService) Create(ctx context.Context, body *ProductsCategoriesCreateBody, query *ProductsCategoriesCreateQuery, opts ...CallOption) (*ProductCategory, error) {
	var out *ProductCategory
	err := r.c.call(ctx, "POST", "/api/products/categories", query, body, false, "category", &out, opts)
	return out, err
}

// Delete — Deletes a product category; its children move up to its parent, its products stay without a category.
//
// Needs `products:delete`.
//
// DELETE /api/products/categories/{categoryId} — Token permission: write:products.
//
// categoryID: The category.
func (r *ProductsCategoriesService) Delete(ctx context.Context, categoryID string, query *ProductsCategoriesDeleteQuery, opts ...CallOption) (*ProductsCategoriesDeleteResponse, error) {
	var out *ProductsCategoriesDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/products/categories/"+pathSeg(categoryID), query, nil, false, "", &out, opts)
	return out, err
}

// List — Lists every product category the caller sees, flat; `parentId` and `position` make the tree.
//
// GET /api/products/categories — Token permission: read:products.
func (r *ProductsCategoriesService) List(ctx context.Context, query *ProductsCategoriesListQuery, opts ...CallOption) ([]ProductCategory, error) {
	var out []ProductCategory
	err := r.c.call(ctx, "GET", "/api/products/categories", query, nil, false, "categories", &out, opts)
	return out, err
}

// Reorder — Puts the children of a category (or the top level) in this order; children not named follow.
//
// Needs `products:update`. Answers the children in their new order.
//
// POST /api/products/categories/reorder — Token permission: write:products.
func (r *ProductsCategoriesService) Reorder(ctx context.Context, body *ProductsCategoriesReorderBody, query *ProductsCategoriesReorderQuery, opts ...CallOption) ([]ProductCategory, error) {
	var out []ProductCategory
	err := r.c.call(ctx, "POST", "/api/products/categories/reorder", query, body, false, "categories", &out, opts)
	return out, err
}

// Update — Renames, moves or re-places a product category.
//
// Needs `products:update`. Never under itself or one of its own children (409 `productCategoryCycle`).
//
// PATCH /api/products/categories/{categoryId} — Token permission: write:products.
//
// categoryID: The category.
func (r *ProductsCategoriesService) Update(ctx context.Context, categoryID string, body *ProductsCategoriesUpdateBody, query *ProductsCategoriesUpdateQuery, opts ...CallOption) (*ProductCategory, error) {
	var out *ProductCategory
	err := r.c.call(ctx, "PATCH", "/api/products/categories/"+pathSeg(categoryID), query, body, false, "category", &out, opts)
	return out, err
}

// ProductsCustomerGroupsService — products.customerGroups
type ProductsCustomerGroupsService struct {
	c *Client
}

func newProductsCustomerGroupsService(c *Client) *ProductsCustomerGroupsService {
	return &ProductsCustomerGroupsService{c: c}
}

// Create — Creates a customer group.
//
// Needs `products:update` and price lists on the plan. 409 `productDuplicateName` for a name already used.
//
// POST /api/products/customer-groups — Token permission: write:products.
func (r *ProductsCustomerGroupsService) Create(ctx context.Context, body *ProductsCustomerGroupsCreateBody, query *ProductsCustomerGroupsCreateQuery, opts ...CallOption) (*ProductCustomerGroup, error) {
	var out *ProductCustomerGroup
	err := r.c.call(ctx, "POST", "/api/products/customer-groups", query, body, false, "customerGroup", &out, opts)
	return out, err
}

// Delete — Deletes a customer group; customers lose it and price lists stop naming it.
//
// Needs `products:update` and price lists on the plan.
//
// DELETE /api/products/customer-groups/{groupId} — Token permission: write:products.
//
// groupID: The customer group.
func (r *ProductsCustomerGroupsService) Delete(ctx context.Context, groupID string, query *ProductsCustomerGroupsDeleteQuery, opts ...CallOption) (*ProductsCustomerGroupsDeleteResponse, error) {
	var out *ProductsCustomerGroupsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/products/customer-groups/"+pathSeg(groupID), query, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the organisation's customer groups, which customers carry and price lists name.
//
// GET /api/products/customer-groups — Token permission: read:products.
func (r *ProductsCustomerGroupsService) List(ctx context.Context, query *ProductsCustomerGroupsListQuery, opts ...CallOption) ([]ProductCustomerGroup, error) {
	var out []ProductCustomerGroup
	err := r.c.call(ctx, "GET", "/api/products/customer-groups", query, nil, false, "customerGroups", &out, opts)
	return out, err
}

// Update — Renames or re-places a customer group.
//
// Needs `products:update` and price lists on the plan.
//
// PATCH /api/products/customer-groups/{groupId} — Token permission: write:products.
//
// groupID: The customer group.
func (r *ProductsCustomerGroupsService) Update(ctx context.Context, groupID string, body *ProductsCustomerGroupsUpdateBody, query *ProductsCustomerGroupsUpdateQuery, opts ...CallOption) (*ProductCustomerGroup, error) {
	var out *ProductCustomerGroup
	err := r.c.call(ctx, "PATCH", "/api/products/customer-groups/"+pathSeg(groupID), query, body, false, "customerGroup", &out, opts)
	return out, err
}

// ProductsInventoryService — products.inventory
type ProductsInventoryService struct {
	c *Client

	Counts       *ProductsInventoryCountsService
	Locations    *ProductsInventoryLocationsService
	Movements    *ProductsInventoryMovementsService
	ReorderRules *ProductsInventoryReorderRulesService
	Settings     *ProductsInventorySettingsService
	Warehouses   *ProductsInventoryWarehousesService
}

func newProductsInventoryService(c *Client) *ProductsInventoryService {
	return &ProductsInventoryService{c: c, Counts: newProductsInventoryCountsService(c), Locations: newProductsInventoryLocationsService(c), Movements: newProductsInventoryMovementsService(c), ReorderRules: newProductsInventoryReorderRulesService(c), Settings: newProductsInventorySettingsService(c), Warehouses: newProductsInventoryWarehousesService(c)}
}

// Availability — For document lines: whether each product keeps stock, how, and what is on hand and available in the line's warehouse.
//
// Read only. Without a warehouse, the default one.
//
// POST /api/products/inventory/availability — Token permission: write:products.
func (r *ProductsInventoryService) Availability(ctx context.Context, body *ProductsInventoryAvailabilityBody, query *ProductsInventoryAvailabilityQuery, opts ...CallOption) ([]ProductsInventoryAvailabilityResponseLinesItem, error) {
	var out []ProductsInventoryAvailabilityResponseLinesItem
	err := r.c.call(ctx, "POST", "/api/products/inventory/availability", query, body, false, "lines", &out, opts)
	return out, err
}

// Batches — A stock item's batches with best-before date and what is in stock.
//
// Needs `products:read` and batches on the plan.
//
// GET /api/products/inventory/batches — Token permission: read:products.
func (r *ProductsInventoryService) Batches(ctx context.Context, query *ProductsInventoryBatchesQuery, opts ...CallOption) ([]InventoryBatch, error) {
	var out []InventoryBatch
	err := r.c.call(ctx, "GET", "/api/products/inventory/batches", query, nil, false, "batches", &out, opts)
	return out, err
}

// Item — A product's stock per warehouse, its batches and serial numbers in stock, and its newest movements.
//
// GET /api/products/inventory/items/{productId} — Token permission: read:products.
//
// productID: The product.
func (r *ProductsInventoryService) Item(ctx context.Context, productID string, query *ProductsInventoryItemQuery, opts ...CallOption) (*InventoryItem, error) {
	var out *InventoryItem
	err := r.c.call(ctx, "GET", "/api/products/inventory/items/"+pathSeg(productID), query, nil, false, "item", &out, opts)
	return out, err
}

// PurchaseDrafts — Creates purchase invoice drafts for items below their reorder point: one per supplier last bought from.
//
// Needs `invoicing:create` with purchases on the plan, `products:read` and inventory on the plan. Each line carries the reorder quantity and the last net unit cost.
//
// POST /api/products/inventory/reorder/purchase-draft — Token permission: write:products.
func (r *ProductsInventoryService) PurchaseDrafts(ctx context.Context, body *ProductsInventoryPurchaseDraftsBody, query *ProductsInventoryPurchaseDraftsQuery, opts ...CallOption) ([]ProductsInventoryPurchaseDraftsResponseDraftsItem, error) {
	var out []ProductsInventoryPurchaseDraftsResponseDraftsItem
	err := r.c.call(ctx, "POST", "/api/products/inventory/reorder/purchase-draft", query, body, false, "drafts", &out, opts)
	return out, err
}

// ReleaseReservation — Releases what an order confirmation still reserves, because the order will not be delivered.
//
// Needs `products:stock`.
//
// POST /api/products/inventory/reservations/release — Token permission: write:products.
func (r *ProductsInventoryService) ReleaseReservation(ctx context.Context, body *ProductsInventoryReleaseReservationBody, query *ProductsInventoryReleaseReservationQuery, opts ...CallOption) (float64, error) {
	var out float64
	err := r.c.call(ctx, "POST", "/api/products/inventory/reservations/release", query, body, false, "released", &out, opts)
	return out, err
}

// Reorder — Stock items at or below their reorder point, with the supplier last bought from and the last net unit cost.
//
// Needs `products:read`. Only products the caller may see; 200 at most.
//
// GET /api/products/inventory/reorder — Token permission: read:products.
func (r *ProductsInventoryService) Reorder(ctx context.Context, query *ProductsInventoryReorderQuery, opts ...CallOption) ([]InventoryReorderItem, error) {
	var out []InventoryReorderItem
	err := r.c.call(ctx, "GET", "/api/products/inventory/reorder", query, nil, false, "items", &out, opts)
	return out, err
}

// Serials — A stock item's serial numbers and where they are.
//
// Needs `products:read` and serial numbers on the plan.
//
// GET /api/products/inventory/serials — Token permission: read:products.
func (r *ProductsInventoryService) Serials(ctx context.Context, query *ProductsInventorySerialsQuery, opts ...CallOption) ([]InventorySerial, error) {
	var out []InventorySerial
	err := r.c.call(ctx, "GET", "/api/products/inventory/serials", query, nil, false, "serials", &out, opts)
	return out, err
}

// SetTracking — Keeps stock for a product (goods only), and how: by quantity, batch or serial number.
//
// Needs `products:stock`; batches and serial numbers need them on the plan. The way changes only while the product holds no stock.
//
// PUT /api/products/inventory/items/{productId} — Token permission: write:products.
//
// productID: The product.
func (r *ProductsInventoryService) SetTracking(ctx context.Context, productID string, body *ProductsInventorySetTrackingBody, query *ProductsInventorySetTrackingQuery, opts ...CallOption) (*ProductsInventorySetTrackingResponseItem, error) {
	var out *ProductsInventorySetTrackingResponseItem
	err := r.c.call(ctx, "PUT", "/api/products/inventory/items/"+pathSeg(productID), query, body, false, "item", &out, opts)
	return out, err
}

// Stock — Stock per item and warehouse: quantity, reserved, available, average cost, value and reorder point.
//
// Needs `products:read` and inventory on the plan. `format=csv` downloads the same as a CSV file.
//
// GET /api/products/inventory/stock — Token permission: read:products.
func (r *ProductsInventoryService) Stock(ctx context.Context, query *ProductsInventoryStockQuery, opts ...CallOption) ([]InventoryStockRow, error) {
	var out []InventoryStockRow
	err := r.c.call(ctx, "GET", "/api/products/inventory/stock", query, nil, false, "rows", &out, opts)
	return out, err
}

// Transfer — Moves stock between two warehouses, or between two storage locations of one warehouse, at the source's average cost.
//
// Needs `products:stock` and warehouses on the plan. Nothing is posted to the books.
//
// POST /api/products/inventory/transfers — Token permission: write:products.
func (r *ProductsInventoryService) Transfer(ctx context.Context, body *ProductsInventoryTransferBody, query *ProductsInventoryTransferQuery, opts ...CallOption) ([]InventoryMovement, error) {
	var out []InventoryMovement
	err := r.c.call(ctx, "POST", "/api/products/inventory/transfers", query, body, false, "movements", &out, opts)
	return out, err
}

// Valuation — The stock's value at a date — quantity times moving average cost per item and warehouse — its total, and per company the value against the stock account in the books.
//
// Needs `products:read`. Every product counts in the totals; one the caller may not see is listed without its name. `format=csv` downloads the rows as a CSV file.
//
// GET /api/products/inventory/valuation — Token permission: read:products.
func (r *ProductsInventoryService) Valuation(ctx context.Context, query *ProductsInventoryValuationQuery, opts ...CallOption) (*InventoryValuationReport, error) {
	var out *InventoryValuationReport
	err := r.c.call(ctx, "GET", "/api/products/inventory/valuation", query, nil, false, "valuation", &out, opts)
	return out, err
}

// ValuationClosing — Posts the periodic closing entry of one company at a date: the value of its warehouses against the stock account.
//
// Needs `products:stock` and `invoicing:bookkeeping`. Once per company and day; balance-method books only (409 `inventoryPerpetualNeedsBalance`).
//
// POST /api/products/inventory/valuation/closing — Token permission: write:products.
func (r *ProductsInventoryService) ValuationClosing(ctx context.Context, body *ProductsInventoryValuationClosingBody, query *ProductsInventoryValuationClosingQuery, opts ...CallOption) (*ProductsInventoryValuationClosingResponseClosing, error) {
	var out *ProductsInventoryValuationClosingResponseClosing
	err := r.c.call(ctx, "POST", "/api/products/inventory/valuation/closing", query, body, false, "closing", &out, opts)
	return out, err
}

// ProductsInventoryCountsService — products.inventory.counts
type ProductsInventoryCountsService struct {
	c *Client
}

func newProductsInventoryCountsService(c *Client) *ProductsInventoryCountsService {
	return &ProductsInventoryCountsService{c: c}
}

// Cancel — Cancels an open stocktaking; nothing was booked.
//
// Needs `products:stock`.
//
// DELETE /api/products/inventory/counts/{countId} — Token permission: write:products.
//
// countID: The stocktaking.
func (r *ProductsInventoryCountsService) Cancel(ctx context.Context, countID string, query *ProductsInventoryCountsCancelQuery, opts ...CallOption) (*InventoryCount, error) {
	var out *InventoryCount
	err := r.c.call(ctx, "DELETE", "/api/products/inventory/counts/"+pathSeg(countID), query, nil, false, "count", &out, opts)
	return out, err
}

// Create — Starts a stocktaking: the expected quantities and average costs of every tracked item in scope are frozen now.
//
// Needs `products:stock`. Counting one storage location needs warehouses on the plan. One open count per warehouse (409 `inventoryCountOpen`); counts of different locations may run side by side.
//
// POST /api/products/inventory/counts — Token permission: write:products.
func (r *ProductsInventoryCountsService) Create(ctx context.Context, body *ProductsInventoryCountsCreateBody, query *ProductsInventoryCountsCreateQuery, opts ...CallOption) (*InventoryCount, error) {
	var out *InventoryCount
	err := r.c.call(ctx, "POST", "/api/products/inventory/counts", query, body, false, "count", &out, opts)
	return out, err
}

// Enter — Enters counted quantities while the stocktaking is open.
//
// Needs `products:stock`. 409 `inventoryCountClosed` once it is posted or cancelled.
//
// PATCH /api/products/inventory/counts/{countId} — Token permission: write:products.
//
// countID: The stocktaking.
func (r *ProductsInventoryCountsService) Enter(ctx context.Context, countID string, body *ProductsInventoryCountsEnterBody, query *ProductsInventoryCountsEnterQuery, opts ...CallOption) (*ProductsInventoryCountsEnterResponse, error) {
	var out *ProductsInventoryCountsEnterResponse
	err := r.c.call(ctx, "PATCH", "/api/products/inventory/counts/"+pathSeg(countID), query, body, false, "", &out, opts)
	return out, err
}

// Get — A stocktaking with its lines.
//
// GET /api/products/inventory/counts/{countId} — Token permission: read:products.
//
// countID: The stocktaking.
func (r *ProductsInventoryCountsService) Get(ctx context.Context, countID string, query *ProductsInventoryCountsGetQuery, opts ...CallOption) (*ProductsInventoryCountsGetResponse, error) {
	var out *ProductsInventoryCountsGetResponse
	err := r.c.call(ctx, "GET", "/api/products/inventory/counts/"+pathSeg(countID), query, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the stocktakings, newest first.
//
// GET /api/products/inventory/counts — Token permission: read:products.
func (r *ProductsInventoryCountsService) List(ctx context.Context, query *ProductsInventoryCountsListQuery, opts ...CallOption) ([]InventoryCount, error) {
	var out []InventoryCount
	err := r.c.call(ctx, "GET", "/api/products/inventory/counts", query, nil, false, "counts", &out, opts)
	return out, err
}

// Post — Books a stocktaking: one movement per difference between counted and expected, and the count is closed.
//
// Needs `products:stock`. Lines nobody counted are skipped, or counted as zero. With perpetual valuation each difference is posted against the stock-change account.
//
// POST /api/products/inventory/counts/{countId}/post — Token permission: write:products.
//
// countID: The stocktaking.
func (r *ProductsInventoryCountsService) Post(ctx context.Context, countID string, body *ProductsInventoryCountsPostBody, query *ProductsInventoryCountsPostQuery, opts ...CallOption) (*ProductsInventoryCountsPostResponse, error) {
	var out *ProductsInventoryCountsPostResponse
	err := r.c.call(ctx, "POST", "/api/products/inventory/counts/"+pathSeg(countID)+"/post", query, body, false, "", &out, opts)
	return out, err
}

// Sheet — Downloads a stocktaking's count sheet as CSV or PDF, with or without the expected quantities.
//
// The CSV uses dot decimals and `;`; the PDF follows the caller's regional settings.
//
// GET /api/products/inventory/counts/{countId}/sheet — Token permission: read:products.
//
// countID: The stocktaking.
func (r *ProductsInventoryCountsService) Sheet(ctx context.Context, countID string, query *ProductsInventoryCountsSheetQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/products/inventory/counts/"+pathSeg(countID)+"/sheet", query, nil, false, "", &out, opts)
	return out, err
}

// ProductsInventoryLocationsService — products.inventory.locations
type ProductsInventoryLocationsService struct {
	c *Client
}

func newProductsInventoryLocationsService(c *Client) *ProductsInventoryLocationsService {
	return &ProductsInventoryLocationsService{c: c}
}

// Create — Adds a storage location (a bin or shelf) to a warehouse.
//
// Needs `products:stock` and warehouses on the plan. A code already used in the warehouse is refused.
//
// POST /api/products/inventory/warehouses/{warehouseId}/locations — Token permission: write:products.
//
// warehouseID: The warehouse.
func (r *ProductsInventoryLocationsService) Create(ctx context.Context, warehouseID string, body *ProductsInventoryLocationsCreateBody, query *ProductsInventoryLocationsCreateQuery, opts ...CallOption) (*InventoryLocation, error) {
	var out *InventoryLocation
	err := r.c.call(ctx, "POST", "/api/products/inventory/warehouses/"+pathSeg(warehouseID)+"/locations", query, body, false, "location", &out, opts)
	return out, err
}

// Update — Changes or archives a storage location.
//
// Needs `products:stock` and warehouses on the plan.
//
// PATCH /api/products/inventory/locations/{locationId} — Token permission: write:products.
//
// locationID: The storage location.
func (r *ProductsInventoryLocationsService) Update(ctx context.Context, locationID string, body *ProductsInventoryLocationsUpdateBody, query *ProductsInventoryLocationsUpdateQuery, opts ...CallOption) (*InventoryLocation, error) {
	var out *InventoryLocation
	err := r.c.call(ctx, "PATCH", "/api/products/inventory/locations/"+pathSeg(locationID), query, body, false, "location", &out, opts)
	return out, err
}

// ProductsInventoryMovementsService — products.inventory.movements
type ProductsInventoryMovementsService struct {
	c *Client
}

func newProductsInventoryMovementsService(c *Client) *ProductsInventoryMovementsService {
	return &ProductsInventoryMovementsService{c: c}
}

// Book — Books a stock movement by hand: goods in, goods out, a correction or an opening stock.
//
// Needs `products:stock`. Another warehouse than the default, or a storage location, needs warehouses on the plan; batches and serial numbers need them on the plan. A correction gives `quantity` (the change) or `setTo` (the new stock). Refusals: 409 `inventoryInsufficientStock`, 422 `inventorySerialsRequired` / `inventoryBatchRequired`, 409 `inventorySerialNotInStock` / `inventorySerialExists` / `inventoryNotTracked`.
//
// POST /api/products/inventory/movements — Token permission: write:products.
func (r *ProductsInventoryMovementsService) Book(ctx context.Context, body *ProductsInventoryMovementsBookBody, query *ProductsInventoryMovementsBookQuery, opts ...CallOption) ([]InventoryMovement, error) {
	var out []InventoryMovement
	err := r.c.call(ctx, "POST", "/api/products/inventory/movements", query, body, false, "movements", &out, opts)
	return out, err
}

// List — The stock movements journal, newest first.
//
// `format=csv` downloads the first 500 as a CSV file.
//
// GET /api/products/inventory/movements — Token permission: read:products.
func (r *ProductsInventoryMovementsService) List(ctx context.Context, query *ProductsInventoryMovementsListQuery, opts ...CallOption) ([]InventoryMovement, error) {
	var out []InventoryMovement
	err := r.c.call(ctx, "GET", "/api/products/inventory/movements", query, nil, false, "movements", &out, opts)
	return out, err
}

// ProductsInventoryReorderRulesService — products.inventory.reorderRules
type ProductsInventoryReorderRulesService struct {
	c *Client
}

func newProductsInventoryReorderRulesService(c *Client) *ProductsInventoryReorderRulesService {
	return &ProductsInventoryReorderRulesService{c: c}
}

// Delete — Removes a stock item's reorder point in a warehouse.
//
// Needs `products:stock`.
//
// DELETE /api/products/inventory/reorder-rules/{ruleId} — Token permission: write:products.
//
// ruleID: The reorder rule.
func (r *ProductsInventoryReorderRulesService) Delete(ctx context.Context, ruleID string, query *ProductsInventoryReorderRulesDeleteQuery, opts ...CallOption) (*ProductsInventoryReorderRulesDeleteResponse, error) {
	var out *ProductsInventoryReorderRulesDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/products/inventory/reorder-rules/"+pathSeg(ruleID), query, nil, false, "", &out, opts)
	return out, err
}

// Save — Sets the reorder point of a stock item in a warehouse; a second call changes it.
//
// Needs `products:stock`.
//
// POST /api/products/inventory/reorder-rules — Token permission: write:products.
func (r *ProductsInventoryReorderRulesService) Save(ctx context.Context, body *ProductsInventoryReorderRulesSaveBody, query *ProductsInventoryReorderRulesSaveQuery, opts ...CallOption) (*InventoryReorderRule, error) {
	var out *InventoryReorderRule
	err := r.c.call(ctx, "POST", "/api/products/inventory/reorder-rules", query, body, false, "rule", &out, opts)
	return out, err
}

// ProductsInventorySettingsService — products.inventory.settings
type ProductsInventorySettingsService struct {
	c *Client
}

func newProductsInventorySettingsService(c *Client) *ProductsInventorySettingsService {
	return &ProductsInventorySettingsService{c: c}
}

// Get — Settings → Products → Inventory: the negative-stock policy, returns, and per company when goods go out and how stock is valued.
//
// GET /api/products/inventory/settings — Token permission: read:products.
func (r *ProductsInventorySettingsService) Get(ctx context.Context, query *ProductsInventorySettingsGetQuery, opts ...CallOption) (*InventorySettings, error) {
	var out *InventorySettings
	err := r.c.call(ctx, "GET", "/api/products/inventory/settings", query, nil, false, "settings", &out, opts)
	return out, err
}

// Update — Changes the inventory settings.
//
// Needs `products:stock`. Perpetual valuation only with balance-method books (409 `inventoryPerpetualNeedsBalance`).
//
// PUT /api/products/inventory/settings — Token permission: write:products.
func (r *ProductsInventorySettingsService) Update(ctx context.Context, body *ProductsInventorySettingsUpdateBody, query *ProductsInventorySettingsUpdateQuery, opts ...CallOption) (*InventorySettings, error) {
	var out *InventorySettings
	err := r.c.call(ctx, "PUT", "/api/products/inventory/settings", query, body, false, "settings", &out, opts)
	return out, err
}

// ProductsInventoryWarehousesService — products.inventory.warehouses
type ProductsInventoryWarehousesService struct {
	c *Client
}

func newProductsInventoryWarehousesService(c *Client) *ProductsInventoryWarehousesService {
	return &ProductsInventoryWarehousesService{c: c}
}

// Create — Creates another warehouse.
//
// Needs `products:stock` and warehouses on the plan.
//
// POST /api/products/inventory/warehouses — Token permission: write:products.
func (r *ProductsInventoryWarehousesService) Create(ctx context.Context, body *ProductsInventoryWarehousesCreateBody, query *ProductsInventoryWarehousesCreateQuery, opts ...CallOption) (*InventoryWarehouse, error) {
	var out *InventoryWarehouse
	err := r.c.call(ctx, "POST", "/api/products/inventory/warehouses", query, body, false, "warehouse", &out, opts)
	return out, err
}

// List — Lists the organisation's warehouses with their storage locations; the default one is made on first use.
//
// GET /api/products/inventory/warehouses — Token permission: read:products.
func (r *ProductsInventoryWarehousesService) List(ctx context.Context, query *ProductsInventoryWarehousesListQuery, opts ...CallOption) ([]InventoryWarehouse, error) {
	var out []InventoryWarehouse
	err := r.c.call(ctx, "GET", "/api/products/inventory/warehouses", query, nil, false, "warehouses", &out, opts)
	return out, err
}

// Update — Changes or archives a warehouse.
//
// Needs `products:stock`. Never archives the default warehouse (409 `inventoryDefaultWarehouse`) or one that holds stock (409 `inventoryWarehouseHasStock`). Making another the default needs warehouses on the plan.
//
// PATCH /api/products/inventory/warehouses/{warehouseId} — Token permission: write:products.
//
// warehouseID: The warehouse.
func (r *ProductsInventoryWarehousesService) Update(ctx context.Context, warehouseID string, body *ProductsInventoryWarehousesUpdateBody, query *ProductsInventoryWarehousesUpdateQuery, opts ...CallOption) (*InventoryWarehouse, error) {
	var out *InventoryWarehouse
	err := r.c.call(ctx, "PATCH", "/api/products/inventory/warehouses/"+pathSeg(warehouseID), query, body, false, "warehouse", &out, opts)
	return out, err
}

// ProductsMarketplaceService — products.marketplace
type ProductsMarketplaceService struct {
	c *Client

	Coupons       *ProductsMarketplaceCouponsService
	LegalPages    *ProductsMarketplaceLegalPagesService
	Orders        *ProductsMarketplaceOrdersService
	PayoutAccount *ProductsMarketplacePayoutAccountService
	Reviews       *ProductsMarketplaceReviewsService
	Settings      *ProductsMarketplaceSettingsService
}

func newProductsMarketplaceService(c *Client) *ProductsMarketplaceService {
	return &ProductsMarketplaceService{c: c, Coupons: newProductsMarketplaceCouponsService(c), LegalPages: newProductsMarketplaceLegalPagesService(c), Orders: newProductsMarketplaceOrdersService(c), PayoutAccount: newProductsMarketplacePayoutAccountService(c), Reviews: newProductsMarketplaceReviewsService(c), Settings: newProductsMarketplaceSettingsService(c)}
}

// BundleItems — What a bundle of the organisation can include: its own listings, and whether each may be bundled now.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:update`, with listings on the plan.
//
// GET /api/products/marketplace/bundle-items — Token permission: read:products.
func (r *ProductsMarketplaceService) BundleItems(ctx context.Context, query *ProductsMarketplaceBundleItemsQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/products/marketplace/bundle-items", query, nil, false, "listings", &out, opts)
	return out, err
}

// GetOffer — A product's marketplace offer and tiers, where it stands in review, and what still stops a submission.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:update`, with listings on the plan.
//
// GET /api/products/{productId}/marketplace — Token permission: read:products.
//
// productID: The product.
func (r *ProductsMarketplaceService) GetOffer(ctx context.Context, productID string, query *ProductsMarketplaceGetOfferQuery, opts ...CallOption) (*ProductsMarketplaceGetOfferResponse, error) {
	var out *ProductsMarketplaceGetOfferResponse
	err := r.c.call(ctx, "GET", "/api/products/"+pathSeg(productID)+"/marketplace", query, nil, false, "", &out, opts)
	return out, err
}

// Offers — The organisation's products that have a marketplace offer, with where each stands, its from-price, sales and rating.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:read`.
//
// GET /api/products/marketplace/offers — Token permission: read:products.
func (r *ProductsMarketplaceService) Offers(ctx context.Context, query *ProductsMarketplaceOffersQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/products/marketplace/offers", query, nil, false, "offers", &out, opts)
	return out, err
}

// Overview — Seller Central's overview: this month's sales, the next payout, rating, refund rate, the onboarding checklist and the to-dos.
//
// Needs the marketplace switched on by the platform, `products:read`, and `marketplace:sell` or `marketplace:view_payouts` (payout figures only with the latter).
//
// GET /api/products/marketplace/overview — Token permission: read:products.
func (r *ProductsMarketplaceService) Overview(ctx context.Context, query *ProductsMarketplaceOverviewQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/products/marketplace/overview", query, nil, false, "overview", &out, opts)
	return out, err
}

// Pause — Takes a product out of the marketplace without a review.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:update`, with listings on the plan. 409 `productMarketplaceNotSubmitted` without an offer.
//
// POST /api/products/{productId}/marketplace/pause — Token permission: write:products.
//
// productID: The product.
func (r *ProductsMarketplaceService) Pause(ctx context.Context, productID string, query *ProductsMarketplacePauseQuery, opts ...CallOption) (string, error) {
	var out string
	err := r.c.call(ctx, "POST", "/api/products/"+pathSeg(productID)+"/marketplace/pause", query, nil, false, "state", &out, opts)
	return out, err
}

// Payouts — The seller's money: per currency what may be paid out now, what is on hold and what is owed back, the schedule, the payouts and the ledger rows not paid out yet.
//
// Needs the marketplace switched on by the platform, `products:read` and `marketplace:view_payouts`.
//
// GET /api/products/marketplace/payouts — Token permission: read:products.
func (r *ProductsMarketplaceService) Payouts(ctx context.Context, query *ProductsMarketplacePayoutsQuery, opts ...CallOption) (*ProductsMarketplacePayoutsResponse, error) {
	var out *ProductsMarketplacePayoutsResponse
	err := r.c.call(ctx, "GET", "/api/products/marketplace/payouts", query, nil, false, "", &out, opts)
	return out, err
}

// Preview — The offer as a buyer would see it now, in the storefront's shape.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:update`, with listings on the plan. `preview` is null without an offer.
//
// GET /api/products/{productId}/marketplace/preview — Token permission: read:products.
//
// productID: The product.
func (r *ProductsMarketplaceService) Preview(ctx context.Context, productID string, query *ProductsMarketplacePreviewQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/products/"+pathSeg(productID)+"/marketplace/preview", query, nil, false, "preview", &out, opts)
	return out, err
}

// Questions — Buyers' questions on all of the organisation's listings, newest first.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:read`.
//
// GET /api/products/marketplace/questions — Token permission: read:products.
func (r *ProductsMarketplaceService) Questions(ctx context.Context, query *ProductsMarketplaceQuestionsQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/products/marketplace/questions", query, nil, false, "questions", &out, opts)
	return out, err
}

// Resume — Shows a paused listing again, without a review.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:update`, with listings on the plan, room under the plan's listing limit and a product that is not archived.
//
// POST /api/products/{productId}/marketplace/resume — Token permission: write:products.
//
// productID: The product.
func (r *ProductsMarketplaceService) Resume(ctx context.Context, productID string, query *ProductsMarketplaceResumeQuery, opts ...CallOption) (string, error) {
	var out string
	err := r.c.call(ctx, "POST", "/api/products/"+pathSeg(productID)+"/marketplace/resume", query, nil, false, "state", &out, opts)
	return out, err
}

// Sales — The seller's sales per day and currency, gross and the seller's share.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:read`. `format=csv` downloads them as a CSV file.
//
// GET /api/products/marketplace/sales — Token permission: read:products.
func (r *ProductsMarketplaceService) Sales(ctx context.Context, query *ProductsMarketplaceSalesQuery, opts ...CallOption) (*ProductsMarketplaceSalesResponse, error) {
	var out *ProductsMarketplaceSalesResponse
	err := r.c.call(ctx, "GET", "/api/products/marketplace/sales", query, nil, false, "", &out, opts)
	return out, err
}

// SaveOffer — Creates or changes a product's marketplace offer; an incomplete one may be saved.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:update`, with listings on the plan. Never goods. `enabled` is not taken here: submit, pause and resume switch it. `tiers` replaces the whole set.
//
// PUT /api/products/{productId}/marketplace — Token permission: write:products.
//
// productID: The product.
func (r *ProductsMarketplaceService) SaveOffer(ctx context.Context, productID string, body map[string]any, query *ProductsMarketplaceSaveOfferQuery, opts ...CallOption) (*ProductsMarketplaceSaveOfferResponse, error) {
	var out *ProductsMarketplaceSaveOfferResponse
	err := r.c.call(ctx, "PUT", "/api/products/"+pathSeg(productID)+"/marketplace", query, body, false, "", &out, opts)
	return out, err
}

// Sources — What an offer of the organisation can deliver — its projects, pages, blocks, workflows, databases or endpoints — by name.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:update`, with listings on the plan. A first filter only: whether the caller may offer the one picked is checked on submission.
//
// GET /api/products/marketplace/sources — Token permission: read:products.
func (r *ProductsMarketplaceService) Sources(ctx context.Context, query *ProductsMarketplaceSourcesQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/products/marketplace/sources", query, nil, false, "sources", &out, opts)
	return out, err
}

// Statement — The settlement statement of one payout: every ledger row it settled and their sums.
//
// Needs the marketplace switched on by the platform, `products:read` and `marketplace:view_payouts`. `format=csv` downloads it for the books; `pdf` the archived statement; `fee-pdf` the platform's fee invoice.
//
// GET /api/products/marketplace/statements/{payoutId} — Token permission: read:products.
//
// payoutID: The payout.
func (r *ProductsMarketplaceService) Statement(ctx context.Context, payoutID string, query *ProductsMarketplaceStatementQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/products/marketplace/statements/"+pathSeg(payoutID), query, nil, false, "statement", &out, opts)
	return out, err
}

// Submit — Submits a product's offer for the marketplace's review.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:update`, with listings on the plan, and a verified e-mail address. 422 `productMarketplaceOfferInvalid` with `issues` while anything is missing; 409 `productMarketplaceNothingChanged` when what is live is what would be submitted. One open submission per offer: a second replaces the first.
//
// POST /api/products/{productId}/marketplace/submit — Token permission: write:products.
//
// productID: The product.
func (r *ProductsMarketplaceService) Submit(ctx context.Context, productID string, body *ProductsMarketplaceSubmitBody, query *ProductsMarketplaceSubmitQuery, opts ...CallOption) (*ProductsMarketplaceSubmitResponse, error) {
	var out *ProductsMarketplaceSubmitResponse
	err := r.c.call(ctx, "POST", "/api/products/"+pathSeg(productID)+"/marketplace/submit", query, body, false, "", &out, opts)
	return out, err
}

// ProductsMarketplaceCouponsService — products.marketplace.coupons
type ProductsMarketplaceCouponsService struct {
	c *Client
}

func newProductsMarketplaceCouponsService(c *Client) *ProductsMarketplaceCouponsService {
	return &ProductsMarketplaceCouponsService{c: c}
}

// Create — Creates a discount code of the organisation's seller profile, for all its listings or one of them.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:update`. 404 `listingNotFound` for a listing that is not the profile's; 409 `codeTaken`.
//
// POST /api/products/marketplace/coupons — Token permission: write:products.
func (r *ProductsMarketplaceCouponsService) Create(ctx context.Context, body *ProductsMarketplaceCouponsCreateBody, query *ProductsMarketplaceCouponsCreateQuery, opts ...CallOption) (*MarketplaceCoupon, error) {
	var out *MarketplaceCoupon
	err := r.c.call(ctx, "POST", "/api/products/marketplace/coupons", query, body, false, "coupon", &out, opts)
	return out, err
}

// Delete — Deletes a discount code of the organisation's seller profile.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:update`. 404 `marketplaceCouponNotFound` for another seller's code.
//
// DELETE /api/products/marketplace/coupons — Token permission: write:products.
func (r *ProductsMarketplaceCouponsService) Delete(ctx context.Context, query *ProductsMarketplaceCouponsDeleteQuery, opts ...CallOption) (*ProductsMarketplaceCouponsDeleteResponse, error) {
	var out *ProductsMarketplaceCouponsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/products/marketplace/coupons", query, nil, false, "", &out, opts)
	return out, err
}

// List — The discount codes of the organisation's seller profile, newest first.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:read`. Every holder of `marketplace:sell` in the organisation sees the same codes. 200 at most.
//
// GET /api/products/marketplace/coupons — Token permission: read:products.
func (r *ProductsMarketplaceCouponsService) List(ctx context.Context, query *ProductsMarketplaceCouponsListQuery, opts ...CallOption) ([]MarketplaceCoupon, error) {
	var out []MarketplaceCoupon
	err := r.c.call(ctx, "GET", "/api/products/marketplace/coupons", query, nil, false, "coupons", &out, opts)
	return out, err
}

// ProductsMarketplaceLegalPagesService — products.marketplace.legalPages
type ProductsMarketplaceLegalPagesService struct {
	c *Client
}

func newProductsMarketplaceLegalPagesService(c *Client) *ProductsMarketplaceLegalPagesService {
	return &ProductsMarketplaceLegalPagesService{c: c}
}

// List — The shop's legal pages: terms, return policy, privacy, imprint.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:read`. 404 `sellerProfileRequired` before the shop exists.
//
// GET /api/products/marketplace/legal-pages — Token permission: read:products.
func (r *ProductsMarketplaceLegalPagesService) List(ctx context.Context, query *ProductsMarketplaceLegalPagesListQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/products/marketplace/legal-pages", query, nil, false, "data.legalPages", &out, opts)
	return out, err
}

// Save — Creates or changes one of the shop's legal pages.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:update`. The content is HTML, at most 500 KB.
//
// PUT /api/products/marketplace/legal-pages — Token permission: write:products.
func (r *ProductsMarketplaceLegalPagesService) Save(ctx context.Context, body *ProductsMarketplaceLegalPagesSaveBody, query *ProductsMarketplaceLegalPagesSaveQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "PUT", "/api/products/marketplace/legal-pages", query, body, false, "data.legalPage", &out, opts)
	return out, err
}

// ProductsMarketplaceOrdersService — products.marketplace.orders
type ProductsMarketplaceOrdersService struct {
	c *Client
}

func newProductsMarketplaceOrdersService(c *Client) *ProductsMarketplaceOrdersService {
	return &ProductsMarketplaceOrdersService{c: c}
}

// Get — One order of the organisation's listings, with its messages, disputes, refunds and what may still be refunded.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:read`.
//
// GET /api/products/marketplace/orders/{orderId} — Token permission: read:products.
//
// orderID: The order.
func (r *ProductsMarketplaceOrdersService) Get(ctx context.Context, orderID string, query *ProductsMarketplaceOrdersGetQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/products/marketplace/orders/"+pathSeg(orderID), query, nil, false, "order", &out, opts)
	return out, err
}

// List — The orders of the organisation's listings, newest first (at most 200).
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:read`. The buyer is named by name and organisation only.
//
// GET /api/products/marketplace/orders — Token permission: read:products.
func (r *ProductsMarketplaceOrdersService) List(ctx context.Context, query *ProductsMarketplaceOrdersListQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/products/marketplace/orders", query, nil, false, "orders", &out, opts)
	return out, err
}

// Message — Writes to the buyer of an order, as the seller.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:update`. 1–5,000 characters; a completed or cancelled order takes no more messages.
//
// POST /api/products/marketplace/orders/{orderId}/messages — Token permission: write:products.
//
// orderID: The order.
func (r *ProductsMarketplaceOrdersService) Message(ctx context.Context, orderID string, body *ProductsMarketplaceOrdersMessageBody, query *ProductsMarketplaceOrdersMessageQuery, opts ...CallOption) (*ProductsMarketplaceOrdersMessageResponseMessage, error) {
	var out *ProductsMarketplaceOrdersMessageResponseMessage
	err := r.c.call(ctx, "POST", "/api/products/marketplace/orders/"+pathSeg(orderID)+"/messages", query, body, false, "message", &out, opts)
	return out, err
}

// Refund — Refunds an order, in full or in part, as a goodwill gesture of the seller.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:update`. The money goes back from the platform's charge and the seller's share is taken off what it is owed. Refused for an order that is refunded, cancelled, disputed or unpaid, or beyond what is left.
//
// POST /api/products/marketplace/orders/{orderId}/refund — Token permission: write:products.
//
// orderID: The order.
func (r *ProductsMarketplaceOrdersService) Refund(ctx context.Context, orderID string, body *ProductsMarketplaceOrdersRefundBody, query *ProductsMarketplaceOrdersRefundQuery, opts ...CallOption) (*ProductsMarketplaceOrdersRefundResponse, error) {
	var out *ProductsMarketplaceOrdersRefundResponse
	err := r.c.call(ctx, "POST", "/api/products/marketplace/orders/"+pathSeg(orderID)+"/refund", query, body, false, "", &out, opts)
	return out, err
}

// ProductsMarketplacePayoutAccountService — products.marketplace.payoutAccount
type ProductsMarketplacePayoutAccountService struct {
	c *Client
}

func newProductsMarketplacePayoutAccountService(c *Client) *ProductsMarketplacePayoutAccountService {
	return &ProductsMarketplacePayoutAccountService{c: c}
}

// Get — The organisation's payout account, read fresh from the payment provider: none, pending, ready or restricted.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:read`.
//
// GET /api/products/marketplace/payout-account — Token permission: read:products.
func (r *ProductsMarketplacePayoutAccountService) Get(ctx context.Context, query *ProductsMarketplacePayoutAccountGetQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/products/marketplace/payout-account", query, nil, false, "account", &out, opts)
	return out, err
}

// Onboard — Opens or continues the payout account's onboarding: answers the URL a person must visit in a browser.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:update`. `url` is null when the account is ready. 404 `sellerProfileRequired` before the shop exists.
//
// POST /api/products/marketplace/payout-account — Token permission: write:products.
func (r *ProductsMarketplacePayoutAccountService) Onboard(ctx context.Context, query *ProductsMarketplacePayoutAccountOnboardQuery, opts ...CallOption) (*ProductsMarketplacePayoutAccountOnboardResponse, error) {
	var out *ProductsMarketplacePayoutAccountOnboardResponse
	err := r.c.call(ctx, "POST", "/api/products/marketplace/payout-account", query, nil, false, "", &out, opts)
	return out, err
}

// ProductsMarketplaceReviewsService — products.marketplace.reviews
type ProductsMarketplaceReviewsService struct {
	c *Client
}

func newProductsMarketplaceReviewsService(c *Client) *ProductsMarketplaceReviewsService {
	return &ProductsMarketplaceReviewsService{c: c}
}

// List — The visible reviews of the organisation's listings, newest first, with the seller's responses.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:read`.
//
// GET /api/products/marketplace/reviews — Token permission: read:products.
func (r *ProductsMarketplaceReviewsService) List(ctx context.Context, query *ProductsMarketplaceReviewsListQuery, opts ...CallOption) ([]map[string]any, error) {
	var out []map[string]any
	err := r.c.call(ctx, "GET", "/api/products/marketplace/reviews", query, nil, false, "reviews", &out, opts)
	return out, err
}

// Respond — Responds publicly to a review; an empty response takes it back.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:update`. 1–2,000 characters.
//
// PATCH /api/products/marketplace/reviews — Token permission: write:products.
func (r *ProductsMarketplaceReviewsService) Respond(ctx context.Context, body *ProductsMarketplaceReviewsRespondBody, query *ProductsMarketplaceReviewsRespondQuery, opts ...CallOption) (*ProductsMarketplaceReviewsRespondResponseReview, error) {
	var out *ProductsMarketplaceReviewsRespondResponseReview
	err := r.c.call(ctx, "PATCH", "/api/products/marketplace/reviews", query, body, false, "review", &out, opts)
	return out, err
}

// ProductsMarketplaceSettingsService — products.marketplace.settings
type ProductsMarketplaceSettingsService struct {
	c *Client
}

func newProductsMarketplaceSettingsService(c *Client) *ProductsMarketplaceSettingsService {
	return &ProductsMarketplaceSettingsService{c: c}
}

// Get — The organisation's seller settings: the shop, the trader information, the price mode, what is still missing, and the payout account's state.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:read`.
//
// GET /api/products/marketplace/settings — Token permission: read:products.
func (r *ProductsMarketplaceSettingsService) Get(ctx context.Context, query *ProductsMarketplaceSettingsGetQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/products/marketplace/settings", query, nil, false, "settings", &out, opts)
	return out, err
}

// Update — Saves the seller settings; the first save creates the organisation's seller profile.
//
// Needs the marketplace switched on by the platform, the `marketplace:sell` permission and `products:update`. 422 `marketplaceSellerFieldInvalid` with `field` and `issue`; 409 `marketplaceShopSlugTaken` for a shop address another shop has.
//
// PUT /api/products/marketplace/settings — Token permission: write:products.
func (r *ProductsMarketplaceSettingsService) Update(ctx context.Context, body *ProductsMarketplaceSettingsUpdateBody, query *ProductsMarketplaceSettingsUpdateQuery, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "PUT", "/api/products/marketplace/settings", query, body, false, "settings", &out, opts)
	return out, err
}

// ProductsPriceListsService — products.priceLists
type ProductsPriceListsService struct {
	c *Client
}

func newProductsPriceListsService(c *Client) *ProductsPriceListsService {
	return &ProductsPriceListsService{c: c}
}

// Create — Creates a price list.
//
// Needs `products:update` and price lists on the plan.
//
// POST /api/products/price-lists — Token permission: write:products.
func (r *ProductsPriceListsService) Create(ctx context.Context, body *ProductPriceListInput, query *ProductsPriceListsCreateQuery, opts ...CallOption) (*ProductPriceList, error) {
	var out *ProductPriceList
	err := r.c.call(ctx, "POST", "/api/products/price-lists", query, body, false, "priceList", &out, opts)
	return out, err
}

// Delete — Deletes a price list and its prices; document lines it priced keep their price.
//
// Needs `products:delete` and price lists on the plan.
//
// DELETE /api/products/price-lists/{priceListId} — Token permission: write:products.
//
// priceListID: The price list.
func (r *ProductsPriceListsService) Delete(ctx context.Context, priceListID string, query *ProductsPriceListsDeleteQuery, opts ...CallOption) (*ProductsPriceListsDeleteResponse, error) {
	var out *ProductsPriceListsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/products/price-lists/"+pathSeg(priceListID), query, nil, false, "", &out, opts)
	return out, err
}

// Export — Downloads a price list's prices as a CSV file.
//
// Needs `products:export`. Columns: product_number; product_name; variant_sku; variant; min_quantity; price.
//
// GET /api/products/price-lists/{priceListId}/export — Token permission: read:products.
//
// priceListID: The price list.
func (r *ProductsPriceListsService) Export(ctx context.Context, priceListID string, query *ProductsPriceListsExportQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/products/price-lists/"+pathSeg(priceListID)+"/export", query, nil, false, "", &out, opts)
	return out, err
}

// Get — A price list and its prices for the products the caller sees.
//
// GET /api/products/price-lists/{priceListId} — Token permission: read:products.
//
// priceListID: The price list.
func (r *ProductsPriceListsService) Get(ctx context.Context, priceListID string, query *ProductsPriceListsGetQuery, opts ...CallOption) (*ProductsPriceListsGetResponse, error) {
	var out *ProductsPriceListsGetResponse
	err := r.c.call(ctx, "GET", "/api/products/price-lists/"+pathSeg(priceListID), query, nil, false, "", &out, opts)
	return out, err
}

// Import — Imports prices into a price list from a CSV file: a preview first, then the write with `apply`.
//
// Needs `products:update` and price lists on the plan. With `apply`, the tiers of every item the file names are replaced by the file's; every other price stays.
//
// POST /api/products/price-lists/{priceListId}/import — Token permission: write:products.
//
// priceListID: The price list.
func (r *ProductsPriceListsService) Import(ctx context.Context, priceListID string, body *ProductCsvImportInput, query *ProductsPriceListsImportQuery, opts ...CallOption) (*ProductsPriceListsImportResponse, error) {
	var out *ProductsPriceListsImportResponse
	err := r.c.call(ctx, "POST", "/api/products/price-lists/"+pathSeg(priceListID)+"/import", query, body, false, "", &out, opts)
	return out, err
}

// List — Lists the organisation's price lists, with how many prices each holds, and its customer groups.
//
// GET /api/products/price-lists — Token permission: read:products.
func (r *ProductsPriceListsService) List(ctx context.Context, query *ProductsPriceListsListQuery, opts ...CallOption) ([]ProductPriceList, error) {
	var out []ProductPriceList
	err := r.c.call(ctx, "GET", "/api/products/price-lists", query, nil, false, "priceLists", &out, opts)
	return out, err
}

// SetPrices — Replaces a price list's prices for every product the caller sees.
//
// Needs `products:update` and price lists on the plan. Prices of products the caller does not see stay. All or nothing: the same item and from-quantity twice is 409 `productDuplicateName`.
//
// PUT /api/products/price-lists/{priceListId}/prices — Token permission: write:products.
//
// priceListID: The price list.
func (r *ProductsPriceListsService) SetPrices(ctx context.Context, priceListID string, body *ProductsPriceListsSetPricesBody, query *ProductsPriceListsSetPricesQuery, opts ...CallOption) ([]ProductPrice, error) {
	var out []ProductPrice
	err := r.c.call(ctx, "PUT", "/api/products/price-lists/"+pathSeg(priceListID)+"/prices", query, body, false, "prices", &out, opts)
	return out, err
}

// Update — Changes a price list.
//
// Needs `products:update` and price lists on the plan.
//
// PATCH /api/products/price-lists/{priceListId} — Token permission: write:products.
//
// priceListID: The price list.
func (r *ProductsPriceListsService) Update(ctx context.Context, priceListID string, body *ProductPriceListInput, query *ProductsPriceListsUpdateQuery, opts ...CallOption) (*ProductPriceList, error) {
	var out *ProductPriceList
	err := r.c.call(ctx, "PATCH", "/api/products/price-lists/"+pathSeg(priceListID), query, body, false, "priceList", &out, opts)
	return out, err
}

// ProductsSettingsService — products.settings
type ProductsSettingsService struct {
	c *Client
}

func newProductsSettingsService(c *Client) *ProductsSettingsService {
	return &ProductsSettingsService{c: c}
}

// Get — Settings → Products: what a new product starts with, and the number prefix.
//
// GET /api/products/settings — Token permission: read:products.
func (r *ProductsSettingsService) Get(ctx context.Context, query *ProductsSettingsGetQuery, opts ...CallOption) (*ProductSettings, error) {
	var out *ProductSettings
	err := r.c.call(ctx, "GET", "/api/products/settings", query, nil, false, "settings", &out, opts)
	return out, err
}

// Update — Changes Settings → Products.
//
// Needs `products:update`.
//
// PATCH /api/products/settings — Token permission: write:products.
func (r *ProductsSettingsService) Update(ctx context.Context, body *ProductsSettingsUpdateBody, query *ProductsSettingsUpdateQuery, opts ...CallOption) (*ProductSettings, error) {
	var out *ProductSettings
	err := r.c.call(ctx, "PATCH", "/api/products/settings", query, body, false, "settings", &out, opts)
	return out, err
}

// ProductsVariantsService — products.variants
type ProductsVariantsService struct {
	c *Client
}

func newProductsVariantsService(c *Client) *ProductsVariantsService {
	return &ProductsVariantsService{c: c}
}

// Export — Downloads every variant of the products the caller sees as a CSV file.
//
// Needs `products:export`. Columns: product_number; product_name; sku; options; gtin; unit_price; weight_kg; archived.
//
// GET /api/products/variants/export — Token permission: read:products.
func (r *ProductsVariantsService) Export(ctx context.Context, query *ProductsVariantsExportQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/products/variants/export", query, nil, false, "", &out, opts)
	return out, err
}

// Get — A product's options with their values, and its variants (archived ones too).
//
// GET /api/products/{productId}/variants — Token permission: read:products.
//
// productID: The product.
func (r *ProductsVariantsService) Get(ctx context.Context, productID string, query *ProductsVariantsGetQuery, opts ...CallOption) (*ProductsVariantsGetResponse, error) {
	var out *ProductsVariantsGetResponse
	err := r.c.call(ctx, "GET", "/api/products/"+pathSeg(productID)+"/variants", query, nil, false, "", &out, opts)
	return out, err
}

// Import — Imports variants from a CSV file: a preview first, then the write with `apply`.
//
// Needs `products:update` and variants on the plan. With `apply`, 422 `productImportInvalid` while rows have errors unless `skipInvalid`.
//
// POST /api/products/variants/import — Token permission: write:products.
func (r *ProductsVariantsService) Import(ctx context.Context, body *ProductCsvImportInput, query *ProductsVariantsImportQuery, opts ...CallOption) (*ProductsVariantsImportResponse, error) {
	var out *ProductsVariantsImportResponse
	err := r.c.call(ctx, "POST", "/api/products/variants/import", query, body, false, "", &out, opts)
	return out, err
}

// SetOptions — Replaces a product's options; the variants follow, and the missing combinations are created.
//
// Needs `products:update` and variants on the plan. A removed value archives the variants with it; a new option gives every variant its first value. At most 100 active variants per product (422 `productVariantLimit`).
//
// PUT /api/products/{productId}/variants — Token permission: write:products.
//
// productID: The product.
func (r *ProductsVariantsService) SetOptions(ctx context.Context, productID string, body *ProductsVariantsSetOptionsBody, query *ProductsVariantsSetOptionsQuery, opts ...CallOption) (*ProductsVariantsSetOptionsResponse, error) {
	var out *ProductsVariantsSetOptionsResponse
	err := r.c.call(ctx, "PUT", "/api/products/"+pathSeg(productID)+"/variants", query, body, false, "", &out, opts)
	return out, err
}

// Update — Changes several variants of a product at once, all or nothing.
//
// Needs `products:update` and variants on the plan. A SKU must be free among the organisation's product numbers and SKUs (409 `productSkuTaken`); a GTIN valid (422 `productInvalidGtin`).
//
// PATCH /api/products/{productId}/variants — Token permission: write:products.
//
// productID: The product.
func (r *ProductsVariantsService) Update(ctx context.Context, productID string, body *ProductsVariantsUpdateBody, query *ProductsVariantsUpdateQuery, opts ...CallOption) (*ProductsVariantsUpdateResponse, error) {
	var out *ProductsVariantsUpdateResponse
	err := r.c.call(ctx, "PATCH", "/api/products/"+pathSeg(productID)+"/variants", query, body, false, "", &out, opts)
	return out, err
}

// ProjectsService — Projects
type ProjectsService struct {
	c *Client

	ContentGaps    *ProjectsContentGapsService
	Cookies        *ProjectsCookiesService
	EmailTemplates *ProjectsEmailTemplatesService
	Journeys       *ProjectsJourneysService
	PageFolders    *ProjectsPageFoldersService
	Redirects      *ProjectsRedirectsService
	Results        *ProjectsResultsService
	SharedBlocks   *ProjectsSharedBlocksService
	SiteHealth     *ProjectsSiteHealthService
	SpamProtection *ProjectsSpamProtectionService
	Waitlist       *ProjectsWaitlistService
}

func newProjectsService(c *Client) *ProjectsService {
	return &ProjectsService{c: c, ContentGaps: newProjectsContentGapsService(c), Cookies: newProjectsCookiesService(c), EmailTemplates: newProjectsEmailTemplatesService(c), Journeys: newProjectsJourneysService(c), PageFolders: newProjectsPageFoldersService(c), Redirects: newProjectsRedirectsService(c), Results: newProjectsResultsService(c), SharedBlocks: newProjectsSharedBlocksService(c), SiteHealth: newProjectsSiteHealthService(c), SpamProtection: newProjectsSpamProtectionService(c), Waitlist: newProjectsWaitlistService(c)}
}

// AddDomain — Claims a custom hostname for a project.
//
// The hostname is served only after verification (`projects.verifyDomain`): the answer's `dnsRecords` say what to create at the DNS provider. Refused for a hostname another project holds or one under the platform's own domains; needs a verified email.
//
// POST /api/projects/{projectId}/domains — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsService) AddDomain(ctx context.Context, projectID string, body *AddDomainInput, opts ...CallOption) (*ProjectDomain, error) {
	var out *ProjectDomain
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/domains", nil, body, false, "data", &out, opts)
	return out, err
}

// AddLanguage — Adds a language to a project.
//
// POST /api/projects/{projectId}/languages — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsService) AddLanguage(ctx context.Context, projectID string, body *ProjectsAddLanguageBody, opts ...CallOption) (*ProjectLanguage, error) {
	var out *ProjectLanguage
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/languages", nil, body, false, "data", &out, opts)
	return out, err
}

// Analytics — Returns a project's page views, clicks, conversions and per-page figures for a period.
//
// GET /api/projects/{projectId}/analytics — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsService) Analytics(ctx context.Context, projectID string, query *ProjectsAnalyticsQuery, opts ...CallOption) (*ProjectAnalytics, error) {
	var out *ProjectAnalytics
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/analytics", query, nil, false, "", &out, opts)
	return out, err
}

// ClaimPlatformDomain — Gives the project its free platform subdomain, if it has none yet.
//
// POST /api/projects/{projectId}/domains/platform — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsService) ClaimPlatformDomain(ctx context.Context, projectID string, opts ...CallOption) (*ProjectDomain, error) {
	var out *ProjectDomain
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/domains/platform", nil, nil, false, "data", &out, opts)
	return out, err
}

// Create — Creates a project in a workspace.
//
// Needs `project.create` and access to the workspace; counts against the plan's `max_projects`. The project gets its free platform subdomain and the default cookie categories.
//
// POST /api/projects — Token permission: write:projects.
func (r *ProjectsService) Create(ctx context.Context, body *CreateProjectInput, opts ...CallOption) (*Project, error) {
	var out *Project
	err := r.c.call(ctx, "POST", "/api/projects", nil, body, false, "data", &out, opts)
	return out, err
}

// CreateComment — Adds a comment or a reply to a project.
//
// People mentioned as `@[Name](userId)` or `@name` who can open the project are notified.
//
// POST /api/projects/{projectId}/comments — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsService) CreateComment(ctx context.Context, projectID string, body *ProjectsCreateCommentBody, opts ...CallOption) (*ProjectComment, error) {
	var out *ProjectComment
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/comments", nil, body, false, "comment", &out, opts)
	return out, err
}

// CreatePreview — Creates a seven-day preview link to the project's draft.
//
// Replaces any earlier preview link of the project.
//
// POST /api/projects/{projectId}/preview — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsService) CreatePreview(ctx context.Context, projectID string, body *ProjectsCreatePreviewBody, opts ...CallOption) (*ProjectPreviewLink, error) {
	var out *ProjectPreviewLink
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/preview", nil, body, false, "data", &out, opts)
	return out, err
}

// Delete — Moves a project to the trash.
//
// Needs `project.delete`. The project can be restored from the trash until it is deleted for good.
//
// DELETE /api/projects/{projectId} — Token permission: delete:projects.
//
// projectID: The project's id.
func (r *ProjectsService) Delete(ctx context.Context, projectID string, opts ...CallOption) (*ProjectsDeleteResponseData, error) {
	var out *ProjectsDeleteResponseData
	err := r.c.call(ctx, "DELETE", "/api/projects/"+pathSeg(projectID), nil, nil, false, "data", &out, opts)
	return out, err
}

// DeleteComment — Deletes a comment.
//
// DELETE /api/projects/{projectId}/comments/{commentId} — Token permission: delete:projects.
//
// projectID: The project's id.
// commentID: The comment's id.
func (r *ProjectsService) DeleteComment(ctx context.Context, projectID string, commentID string, opts ...CallOption) (*ProjectsDeleteCommentResponse, error) {
	var out *ProjectsDeleteCommentResponse
	err := r.c.call(ctx, "DELETE", "/api/projects/"+pathSeg(projectID)+"/comments/"+pathSeg(commentID), nil, nil, false, "", &out, opts)
	return out, err
}

// DownloadExport — Downloads a finished export as a ZIP file.
//
// Answers the file itself (`application/zip`), not JSON. Only the person who started the export can download it.
//
// GET /api/projects/{projectId}/export/{exportId}/download — Token permission: read:projects.
//
// projectID: The project's id.
// exportID: The export's id.
func (r *ProjectsService) DownloadExport(ctx context.Context, projectID string, exportID string, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/export/"+pathSeg(exportID)+"/download", nil, nil, false, "", &out, opts)
	return out, err
}

// Export — Starts an export of the project as a ZIP file that another installation can import.
//
// Needs `project.export`. The export runs in the background (202); poll `projects.listExports` and download it when `downloadable`.
//
// POST /api/projects/{projectId}/export — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsService) Export(ctx context.Context, projectID string, opts ...CallOption) (*ProjectsExportResponseDataExport, error) {
	var out *ProjectsExportResponseDataExport
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/export", nil, nil, false, "data.export", &out, opts)
	return out, err
}

// Get — Returns one project.
//
// GET /api/projects/{projectId} — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsService) Get(ctx context.Context, projectID string, opts ...CallOption) (*Project, error) {
	var out *Project
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetBrain — Returns what the AI knows about the project: its business, audience, goals, tone and facts.
//
// GET /api/projects/{projectId}/brain — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsService) GetBrain(ctx context.Context, projectID string, opts ...CallOption) (*ProjectBrainRecord, error) {
	var out *ProjectBrainRecord
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/brain", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetComingSoon — Returns the project's coming-soon mode: whether it is on, the teaser page and the paths that stay reachable.
//
// GET /api/projects/{projectId}/coming-soon — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsService) GetComingSoon(ctx context.Context, projectID string, opts ...CallOption) (*ProjectsGetComingSoonResponseData, error) {
	var out *ProjectsGetComingSoonResponseData
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/coming-soon", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetNotFoundPage — Returns which page the live site shows for a URL that does not exist.
//
// GET /api/projects/{projectId}/not-found-page — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsService) GetNotFoundPage(ctx context.Context, projectID string, opts ...CallOption) (*ProjectsGetNotFoundPageResponseData, error) {
	var out *ProjectsGetNotFoundPageResponseData
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/not-found-page", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetStyles — Returns a project's global styles (colors, fonts, spacing, radii).
//
// GET /api/projects/{projectId}/styles — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsService) GetStyles(ctx context.Context, projectID string, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/styles", nil, nil, false, "data", &out, opts)
	return out, err
}

// List — Lists the projects the caller can open, optionally of one workspace or organization.
//
// Without `workspaceId` or `organizationId` it lists the projects of every organization the caller belongs to. A private project is listed only to its creator.
//
// GET /api/projects — Token permission: read:projects.
func (r *ProjectsService) List(ctx context.Context, query *ProjectsListQuery, opts ...CallOption) ([]Project, error) {
	var out []Project
	err := r.c.call(ctx, "GET", "/api/projects", query, nil, false, "projects", &out, opts)
	return out, err
}

// ListCollaborators — Lists the people of the project's organization, for mentioning them in comments.
//
// GET /api/projects/{projectId}/collaborators — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsService) ListCollaborators(ctx context.Context, projectID string, opts ...CallOption) ([]ProjectCollaborator, error) {
	var out []ProjectCollaborator
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/collaborators", nil, nil, false, "members", &out, opts)
	return out, err
}

// ListComments — Lists a project's top-level comments with their replies.
//
// GET /api/projects/{projectId}/comments — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsService) ListComments(ctx context.Context, projectID string, query *ProjectsListCommentsQuery, opts ...CallOption) ([]ProjectCommentThread, error) {
	var out []ProjectCommentThread
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/comments", query, nil, false, "comments", &out, opts)
	return out, err
}

// ListDomains — Lists the hostnames of a project, with their verification status.
//
// GET /api/projects/{projectId}/domains — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsService) ListDomains(ctx context.Context, projectID string, opts ...CallOption) ([]ProjectDomain, error) {
	var out []ProjectDomain
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/domains", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListExports — Lists the caller's last ten exports of the project.
//
// GET /api/projects/{projectId}/export — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsService) ListExports(ctx context.Context, projectID string, opts ...CallOption) ([]ProjectExport, error) {
	var out []ProjectExport
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/export", nil, nil, false, "data.exports", &out, opts)
	return out, err
}

// ListLanguages — Lists the languages of a project.
//
// GET /api/projects/{projectId}/languages — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsService) ListLanguages(ctx context.Context, projectID string, opts ...CallOption) ([]ProjectLanguage, error) {
	var out []ProjectLanguage
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/languages", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListPages — Lists a project's pages.
//
// Without `include=components` the pages come without their component trees.
//
// GET /api/projects/{projectId}/pages — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsService) ListPages(ctx context.Context, projectID string, query *ProjectsListPagesQuery, opts ...CallOption) ([]PageListItem, error) {
	var out []PageListItem
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/pages", query, nil, false, "data", &out, opts)
	return out, err
}

// Publish — Publishes every draft page and every live page with unpublished changes.
//
// Needs `page.update` and a verified email. Each page goes through the publish rules — freeze windows and the organization's approval requirement — so a page may be sent for approval or refused instead. At most 6 calls an hour per project.
//
// POST /api/projects/{projectId}/publish — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsService) Publish(ctx context.Context, projectID string, opts ...CallOption) (*ProjectPublishResult, error) {
	var out *ProjectPublishResult
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/publish", nil, nil, false, "data", &out, opts)
	return out, err
}

// RemoveDomain — Releases a hostname from a project.
//
// DELETE /api/projects/{projectId}/domains/{domainId} — Token permission: delete:projects.
//
// projectID: The project's id.
// domainID: The domain's id.
func (r *ProjectsService) RemoveDomain(ctx context.Context, projectID string, domainID string, opts ...CallOption) ([]ProjectDomain, error) {
	var out []ProjectDomain
	err := r.c.call(ctx, "DELETE", "/api/projects/"+pathSeg(projectID)+"/domains/"+pathSeg(domainID), nil, nil, false, "data", &out, opts)
	return out, err
}

// RemoveLanguage — Removes a language from a project.
//
// DELETE /api/projects/{projectId}/languages — Token permission: delete:projects.
//
// projectID: The project's id.
func (r *ProjectsService) RemoveLanguage(ctx context.Context, projectID string, query *ProjectsRemoveLanguageQuery, opts ...CallOption) (*ProjectsRemoveLanguageResponseData, error) {
	var out *ProjectsRemoveLanguageResponseData
	err := r.c.call(ctx, "DELETE", "/api/projects/"+pathSeg(projectID)+"/languages", query, nil, false, "data", &out, opts)
	return out, err
}

// RequestApproval — Starts an approval of the project with the given steps and approvers.
//
// The approvers of the first step are notified.
//
// POST /api/projects/{projectId}/request-approval — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsService) RequestApproval(ctx context.Context, projectID string, body *ProjectsRequestApprovalBody, opts ...CallOption) (*Approval, error) {
	var out *Approval
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/request-approval", nil, body, false, "data.workflow", &out, opts)
	return out, err
}

// RevokePreview — Revokes the project's preview link.
//
// DELETE /api/projects/{projectId}/preview — Token permission: delete:projects.
//
// projectID: The project's id.
func (r *ProjectsService) RevokePreview(ctx context.Context, projectID string, opts ...CallOption) (*ProjectsRevokePreviewResponse, error) {
	var out *ProjectsRevokePreviewResponse
	err := r.c.call(ctx, "DELETE", "/api/projects/"+pathSeg(projectID)+"/preview", nil, nil, false, "", &out, opts)
	return out, err
}

// SetDefaultLanguage — Makes one of the project's languages its default.
//
// Pages without a language of their own are tagged with it; `retagged` says how many.
//
// PATCH /api/projects/{projectId}/languages — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsService) SetDefaultLanguage(ctx context.Context, projectID string, body *ProjectsSetDefaultLanguageBody, opts ...CallOption) (*ProjectsSetDefaultLanguageResponseData, error) {
	var out *ProjectsSetDefaultLanguageResponseData
	err := r.c.call(ctx, "PATCH", "/api/projects/"+pathSeg(projectID)+"/languages", nil, body, false, "data", &out, opts)
	return out, err
}

// SetHomePage — Makes a page the project's home page.
//
// PUT /api/projects/{projectId}/pages/{pageId}/set-home — Token permission: write:projects.
//
// projectID: The project's id.
// pageID: The page's id.
func (r *ProjectsService) SetHomePage(ctx context.Context, projectID string, pageID string, opts ...CallOption) (*Page, error) {
	var out *Page
	err := r.c.call(ctx, "PUT", "/api/projects/"+pathSeg(projectID)+"/pages/"+pathSeg(pageID)+"/set-home", nil, nil, false, "page", &out, opts)
	return out, err
}

// SetPrimaryDomain — Makes a hostname the project's primary one.
//
// PUT /api/projects/{projectId}/domains/{domainId}/primary — Token permission: write:projects.
//
// projectID: The project's id.
// domainID: The domain's id.
func (r *ProjectsService) SetPrimaryDomain(ctx context.Context, projectID string, domainID string, opts ...CallOption) ([]ProjectDomain, error) {
	var out []ProjectDomain
	err := r.c.call(ctx, "PUT", "/api/projects/"+pathSeg(projectID)+"/domains/"+pathSeg(domainID)+"/primary", nil, nil, false, "data", &out, opts)
	return out, err
}

// SuggestBrain — Drafts the project's brain from its pages with AI, without saving it.
//
// Counts against the organization's AI budget; save the draft with `projects.updateBrain`.
//
// POST /api/projects/{projectId}/brain/suggest — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsService) SuggestBrain(ctx context.Context, projectID string, opts ...CallOption) (*ProjectsSuggestBrainResponseData, error) {
	var out *ProjectsSuggestBrainResponseData
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/brain/suggest", nil, nil, false, "data", &out, opts)
	return out, err
}

// Traffic — Returns where the live site's visitors came from, per day and by source, referrer, campaign, country, device and browser.
//
// GET /api/projects/{projectId}/traffic — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsService) Traffic(ctx context.Context, projectID string, query *ProjectsTrafficQuery, opts ...CallOption) (*ProjectTraffic, error) {
	var out *ProjectTraffic
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/traffic", query, nil, false, "data", &out, opts)
	return out, err
}

// UnsetHomePage — Removes a page's home-page designation.
//
// DELETE /api/projects/{projectId}/pages/{pageId}/set-home — Token permission: delete:projects.
//
// projectID: The project's id.
// pageID: The page's id.
func (r *ProjectsService) UnsetHomePage(ctx context.Context, projectID string, pageID string, opts ...CallOption) (*Page, error) {
	var out *Page
	err := r.c.call(ctx, "DELETE", "/api/projects/"+pathSeg(projectID)+"/pages/"+pathSeg(pageID)+"/set-home", nil, nil, false, "page", &out, opts)
	return out, err
}

// Update — Changes a project's name, description, status, visibility or settings.
//
// PUT /api/projects/{projectId} — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsService) Update(ctx context.Context, projectID string, body *UpdateProjectInput, opts ...CallOption) (*Project, error) {
	var out *Project
	err := r.c.call(ctx, "PUT", "/api/projects/"+pathSeg(projectID), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateBrain — Saves what the AI knows about the project.
//
// PUT /api/projects/{projectId}/brain — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsService) UpdateBrain(ctx context.Context, projectID string, body *ProjectsUpdateBrainBody, opts ...CallOption) (*ProjectBrainRecord, error) {
	var out *ProjectBrainRecord
	err := r.c.call(ctx, "PUT", "/api/projects/"+pathSeg(projectID)+"/brain", nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateComingSoon — Turns coming-soon mode on or off, or changes its teaser page and reachable paths.
//
// Needs `page.update`. Turning it on needs a published teaser page of this project. Keys left out keep their value.
//
// PUT /api/projects/{projectId}/coming-soon — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsService) UpdateComingSoon(ctx context.Context, projectID string, body *ProjectsUpdateComingSoonBody, opts ...CallOption) (*ProjectsUpdateComingSoonResponseData, error) {
	var out *ProjectsUpdateComingSoonResponseData
	err := r.c.call(ctx, "PUT", "/api/projects/"+pathSeg(projectID)+"/coming-soon", nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateComment — Edits a comment's text or resolves it.
//
// Only its author changes the text; anyone who can open the project changes the status.
//
// PATCH /api/projects/{projectId}/comments/{commentId} — Token permission: write:projects.
//
// projectID: The project's id.
// commentID: The comment's id.
func (r *ProjectsService) UpdateComment(ctx context.Context, projectID string, commentID string, body *ProjectsUpdateCommentBody, opts ...CallOption) (*ProjectComment, error) {
	var out *ProjectComment
	err := r.c.call(ctx, "PATCH", "/api/projects/"+pathSeg(projectID)+"/comments/"+pathSeg(commentID), nil, body, false, "comment", &out, opts)
	return out, err
}

// UpdateNotFoundPage — Chooses the page the live site shows for a URL that does not exist.
//
// Needs `page.update`. A null `pageId` goes back to the platform's own not-found page.
//
// PUT /api/projects/{projectId}/not-found-page — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsService) UpdateNotFoundPage(ctx context.Context, projectID string, body *ProjectsUpdateNotFoundPageBody, opts ...CallOption) (*ProjectsUpdateNotFoundPageResponseData, error) {
	var out *ProjectsUpdateNotFoundPageResponseData
	err := r.c.call(ctx, "PUT", "/api/projects/"+pathSeg(projectID)+"/not-found-page", nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateStyles — Replaces a project's global styles.
//
// PUT /api/projects/{projectId}/styles — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsService) UpdateStyles(ctx context.Context, projectID string, body map[string]any, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "PUT", "/api/projects/"+pathSeg(projectID)+"/styles", nil, body, false, "data", &out, opts)
	return out, err
}

// VerifyDomain — Checks a hostname's DNS records and marks it verified when they are right.
//
// POST /api/projects/{projectId}/domains/{domainId}/verify — Token permission: write:projects.
//
// projectID: The project's id.
// domainID: The domain's id.
func (r *ProjectsService) VerifyDomain(ctx context.Context, projectID string, domainID string, opts ...CallOption) (*ProjectDomain, error) {
	var out *ProjectDomain
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/domains/"+pathSeg(domainID)+"/verify", nil, nil, false, "data", &out, opts)
	return out, err
}

// VideoAnalytics — Returns plays, viewers and completion of each video on the project's pages.
//
// GET /api/projects/{projectId}/video-analytics — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsService) VideoAnalytics(ctx context.Context, projectID string, opts ...CallOption) ([]VideoStats, error) {
	var out []VideoStats
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/video-analytics", nil, nil, false, "data.videos", &out, opts)
	return out, err
}

// ProjectsContentGapsService — projects.contentGaps
type ProjectsContentGapsService struct {
	c *Client
}

func newProjectsContentGapsService(c *Client) *ProjectsContentGapsService {
	return &ProjectsContentGapsService{c: c}
}

// Faq — Suggests an FAQ entry for a content-gap topic, or adds one to the page's draft.
//
// Without `apply` the AI drafts a question and answer (rate-limited). With `apply: true` the given `question` and `answer` are added to the page's draft and the topic is marked `faq_added`.
//
// POST /api/projects/{projectId}/content-gaps/{topicId}/faq — Token permission: write:projects.
//
// projectID: The project's id.
// topicID: The topic's id.
func (r *ProjectsContentGapsService) Faq(ctx context.Context, projectID string, topicID string, body *ProjectsContentGapsFaqBody, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/content-gaps/"+pathSeg(topicID)+"/faq", nil, body, false, "data", &out, opts)
	return out, err
}

// List — Lists the questions visitors ask that the site does not answer yet.
//
// GET /api/projects/{projectId}/content-gaps — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsContentGapsService) List(ctx context.Context, projectID string, query *ProjectsContentGapsListQuery, opts ...CallOption) (*ProjectsContentGapsListResponseData, error) {
	var out *ProjectsContentGapsListResponseData
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/content-gaps", query, nil, false, "data", &out, opts)
	return out, err
}

// Update — Dismisses a content-gap topic, marks it answered, or reopens it.
//
// POST /api/projects/{projectId}/content-gaps/{topicId} — Token permission: write:projects.
//
// projectID: The project's id.
// topicID: The topic's id.
func (r *ProjectsContentGapsService) Update(ctx context.Context, projectID string, topicID string, body *ProjectsContentGapsUpdateBody, opts ...CallOption) (*ProjectsContentGapsUpdateResponseData, error) {
	var out *ProjectsContentGapsUpdateResponseData
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/content-gaps/"+pathSeg(topicID), nil, body, false, "data", &out, opts)
	return out, err
}

// ProjectsCookiesService — projects.cookies
type ProjectsCookiesService struct {
	c *Client
}

func newProjectsCookiesService(c *Client) *ProjectsCookiesService {
	return &ProjectsCookiesService{c: c}
}

// Create — Declares a cookie on the project's banner.
//
// POST /api/projects/{projectId}/cookies/cookies — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsCookiesService) Create(ctx context.Context, projectID string, body *ProjectCookieInput, opts ...CallOption) (*ProjectCookie, error) {
	var out *ProjectCookie
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/cookies/cookies", nil, body, false, "data", &out, opts)
	return out, err
}

// CreateCategory — Adds a cookie category to the project's banner.
//
// POST /api/projects/{projectId}/cookies/categories — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsCookiesService) CreateCategory(ctx context.Context, projectID string, body *CookieCategoryInput, opts ...CallOption) (*CookieCategory, error) {
	var out *CookieCategory
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/cookies/categories", nil, body, false, "data", &out, opts)
	return out, err
}

// Delete — Removes a declared cookie.
//
// DELETE /api/projects/{projectId}/cookies/cookies/{cookieId} — Token permission: delete:projects.
//
// projectID: The project's id.
// cookieID: The cookie's id.
func (r *ProjectsCookiesService) Delete(ctx context.Context, projectID string, cookieID string, opts ...CallOption) (*ProjectsCookiesDeleteResponseData, error) {
	var out *ProjectsCookiesDeleteResponseData
	err := r.c.call(ctx, "DELETE", "/api/projects/"+pathSeg(projectID)+"/cookies/cookies/"+pathSeg(cookieID), nil, nil, false, "data", &out, opts)
	return out, err
}

// DeleteCategory — Removes a cookie category.
//
// DELETE /api/projects/{projectId}/cookies/categories/{categoryId} — Token permission: delete:projects.
//
// projectID: The project's id.
// categoryID: The category's id.
func (r *ProjectsCookiesService) DeleteCategory(ctx context.Context, projectID string, categoryID string, opts ...CallOption) (*ProjectsCookiesDeleteCategoryResponseData, error) {
	var out *ProjectsCookiesDeleteCategoryResponseData
	err := r.c.call(ctx, "DELETE", "/api/projects/"+pathSeg(projectID)+"/cookies/categories/"+pathSeg(categoryID), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetScan — Returns how a cookie scan is going, and its findings when done.
//
// GET /api/projects/{projectId}/cookies/scan/{jobId} — Token permission: read:projects.
//
// projectID: The project's id.
// jobID: The scan's job id.
func (r *ProjectsCookiesService) GetScan(ctx context.Context, projectID string, jobID string, opts ...CallOption) (*ProjectsCookiesGetScanResponseData, error) {
	var out *ProjectsCookiesGetScanResponseData
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/cookies/scan/"+pathSeg(jobID), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetSettings — Returns the project's cookie banner settings.
//
// GET /api/projects/{projectId}/cookies/settings — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsCookiesService) GetSettings(ctx context.Context, projectID string, opts ...CallOption) (*CookieBannerSettings, error) {
	var out *CookieBannerSettings
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/cookies/settings", nil, nil, false, "data", &out, opts)
	return out, err
}

// List — Lists the cookies and storage the project's banner declares.
//
// GET /api/projects/{projectId}/cookies/cookies — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsCookiesService) List(ctx context.Context, projectID string, opts ...CallOption) ([]ProjectCookie, error) {
	var out []ProjectCookie
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/cookies/cookies", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListCategories — Lists the cookie categories visitors consent to on the project's banner.
//
// GET /api/projects/{projectId}/cookies/categories — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsCookiesService) ListCategories(ctx context.Context, projectID string, opts ...CallOption) ([]CookieCategory, error) {
	var out []CookieCategory
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/cookies/categories", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListConsents — Lists the consent records visitors left on the project's banner, with totals.
//
// GET /api/projects/{projectId}/cookies/consents — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsCookiesService) ListConsents(ctx context.Context, projectID string, query *ProjectsCookiesListConsentsQuery, opts ...CallOption) (*ProjectsCookiesListConsentsResponseData, error) {
	var out *ProjectsCookiesListConsentsResponseData
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/cookies/consents", query, nil, false, "data", &out, opts)
	return out, err
}

// Scan — Starts a scan of the live site's pages for the cookies and storage they set.
//
// Runs in the background; poll `projects.cookies.getScan` with the `jobId`.
//
// POST /api/projects/{projectId}/cookies/scan — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsCookiesService) Scan(ctx context.Context, projectID string, body *ProjectsCookiesScanBody, opts ...CallOption) (*ProjectsCookiesScanResponseData, error) {
	var out *ProjectsCookiesScanResponseData
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/cookies/scan", nil, body, false, "data", &out, opts)
	return out, err
}

// Seed — Adds the standard cookie categories and the platform's own cookies where they are missing.
//
// POST /api/projects/{projectId}/cookies/seed — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsCookiesService) Seed(ctx context.Context, projectID string, opts ...CallOption) (*ProjectsCookiesSeedResponseData, error) {
	var out *ProjectsCookiesSeedResponseData
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/cookies/seed", nil, nil, false, "data", &out, opts)
	return out, err
}

// Update — Changes a declared cookie; only the fields sent change.
//
// PUT /api/projects/{projectId}/cookies/cookies/{cookieId} — Token permission: write:projects.
//
// projectID: The project's id.
// cookieID: The cookie's id.
func (r *ProjectsCookiesService) Update(ctx context.Context, projectID string, cookieID string, body *ProjectsCookiesUpdateBody, opts ...CallOption) (*ProjectCookie, error) {
	var out *ProjectCookie
	err := r.c.call(ctx, "PUT", "/api/projects/"+pathSeg(projectID)+"/cookies/cookies/"+pathSeg(cookieID), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateCategory — Changes a cookie category; only the fields sent change.
//
// PUT /api/projects/{projectId}/cookies/categories/{categoryId} — Token permission: write:projects.
//
// projectID: The project's id.
// categoryID: The category's id.
func (r *ProjectsCookiesService) UpdateCategory(ctx context.Context, projectID string, categoryID string, body *ProjectsCookiesUpdateCategoryBody, opts ...CallOption) (*CookieCategory, error) {
	var out *CookieCategory
	err := r.c.call(ctx, "PUT", "/api/projects/"+pathSeg(projectID)+"/cookies/categories/"+pathSeg(categoryID), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateSettings — Changes the project's cookie banner settings; only the fields sent change.
//
// Raise `consentVersion` to ask every visitor again.
//
// PUT /api/projects/{projectId}/cookies/settings — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsCookiesService) UpdateSettings(ctx context.Context, projectID string, body *ProjectsCookiesUpdateSettingsBody, opts ...CallOption) (*CookieBannerSettings, error) {
	var out *CookieBannerSettings
	err := r.c.call(ctx, "PUT", "/api/projects/"+pathSeg(projectID)+"/cookies/settings", nil, body, false, "data", &out, opts)
	return out, err
}

// ProjectsEmailTemplatesService — projects.emailTemplates
type ProjectsEmailTemplatesService struct {
	c *Client
}

func newProjectsEmailTemplatesService(c *Client) *ProjectsEmailTemplatesService {
	return &ProjectsEmailTemplatesService{c: c}
}

// List — Lists the visitor emails the project may word itself, in each of its languages, with the wording that is sent.
//
// GET /api/projects/{projectId}/email-templates — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsEmailTemplatesService) List(ctx context.Context, projectID string, opts ...CallOption) (*ProjectsEmailTemplatesListResponseData, error) {
	var out *ProjectsEmailTemplatesListResponseData
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/email-templates", nil, nil, false, "data", &out, opts)
	return out, err
}

// Reset — Goes back to the platform's wording of a visitor email in one language.
//
// DELETE /api/projects/{projectId}/email-templates/{type} — Token permission: delete:projects.
//
// projectID: The project's id.
// typeArg: The email: `waitlist_confirm` or `waitlist_welcome`.
func (r *ProjectsEmailTemplatesService) Reset(ctx context.Context, projectID string, typeArg string, query *ProjectsEmailTemplatesResetQuery, opts ...CallOption) (*ProjectsEmailTemplatesResetResponse, error) {
	var out *ProjectsEmailTemplatesResetResponse
	err := r.c.call(ctx, "DELETE", "/api/projects/"+pathSeg(projectID)+"/email-templates/"+pathSeg(typeArg), query, nil, false, "", &out, opts)
	return out, err
}

// Set — Saves the project's own wording of a visitor email in one language.
//
// Needs the plan's custom email templates. The HTML is cleaned: no scripts, no event handlers, only http(s) and mailto links.
//
// PUT /api/projects/{projectId}/email-templates/{type} — Token permission: write:projects.
//
// projectID: The project's id.
// typeArg: The email: `waitlist_confirm` or `waitlist_welcome`.
func (r *ProjectsEmailTemplatesService) Set(ctx context.Context, projectID string, typeArg string, body *EmailTemplateInput, opts ...CallOption) (*ProjectsEmailTemplatesSetResponse, error) {
	var out *ProjectsEmailTemplatesSetResponse
	err := r.c.call(ctx, "PUT", "/api/projects/"+pathSeg(projectID)+"/email-templates/"+pathSeg(typeArg), nil, body, false, "", &out, opts)
	return out, err
}

// ProjectsJourneysService — projects.journeys
type ProjectsJourneysService struct {
	c *Client
}

func newProjectsJourneysService(c *Client) *ProjectsJourneysService {
	return &ProjectsJourneysService{c: c}
}

// GetSettings — Returns the project's visitor-journey settings, what applies in each region, and the platform's limits.
//
// Needs `project.tracking`.
//
// GET /api/projects/{projectId}/journeys/settings — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsJourneysService) GetSettings(ctx context.Context, projectID string, opts ...CallOption) (*ProjectsJourneysGetSettingsResponseData, error) {
	var out *ProjectsJourneysGetSettingsResponseData
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/journeys/settings", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetVisitor — Returns one visitor's journey: visits, touches, consent and conversions.
//
// GET /api/projects/{projectId}/journeys/visitors/{visitorId} — Token permission: read:projects.
//
// projectID: The project's id.
// visitorID: The visitor's id.
func (r *ProjectsJourneysService) GetVisitor(ctx context.Context, projectID string, visitorID string, opts ...CallOption) (*JourneyVisitor, error) {
	var out *JourneyVisitor
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/journeys/visitors/"+pathSeg(visitorID), nil, nil, false, "data", &out, opts)
	return out, err
}

// ListDeliveries — Lists the conversions sent, or due to be sent, to ad networks, with each network's answer.
//
// GET /api/projects/{projectId}/journeys/deliveries — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsJourneysService) ListDeliveries(ctx context.Context, projectID string, query *ProjectsJourneysListDeliveriesQuery, opts ...CallOption) (*ProjectsJourneysListDeliveriesResponseData, error) {
	var out *ProjectsJourneysListDeliveriesResponseData
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/journeys/deliveries", query, nil, false, "data", &out, opts)
	return out, err
}

// ListVisitors — Lists the visitors of a period with their first and last touch and conversions.
//
// GET /api/projects/{projectId}/journeys/visitors — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsJourneysService) ListVisitors(ctx context.Context, projectID string, query *ProjectsJourneysListVisitorsQuery, opts ...CallOption) (*ProjectsJourneysListVisitorsResponseData, error) {
	var out *ProjectsJourneysListVisitorsResponseData
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/journeys/visitors", query, nil, false, "data", &out, opts)
	return out, err
}

// Report — Returns visitors, conversions and revenue of a period, credited to channels or campaigns by an attribution model.
//
// GET /api/projects/{projectId}/journeys/report — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsJourneysService) Report(ctx context.Context, projectID string, query *ProjectsJourneysReportQuery, opts ...CallOption) (*JourneyReport, error) {
	var out *JourneyReport
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/journeys/report", query, nil, false, "data", &out, opts)
	return out, err
}

// UpdateSettings — Changes the project's visitor-journey settings.
//
// Needs `project.tracking`. Send only the keys to change, in the shape of `settings` as read; the platform's locks still apply. `changed` lists what moved.
//
// PUT /api/projects/{projectId}/journeys/settings — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsJourneysService) UpdateSettings(ctx context.Context, projectID string, body map[string]any, opts ...CallOption) (*ProjectsJourneysUpdateSettingsResponseData, error) {
	var out *ProjectsJourneysUpdateSettingsResponseData
	err := r.c.call(ctx, "PUT", "/api/projects/"+pathSeg(projectID)+"/journeys/settings", nil, body, false, "data", &out, opts)
	return out, err
}

// ProjectsPageFoldersService — projects.pageFolders
type ProjectsPageFoldersService struct {
	c *Client
}

func newProjectsPageFoldersService(c *Client) *ProjectsPageFoldersService {
	return &ProjectsPageFoldersService{c: c}
}

// Create — Creates a folder in a project's pages list; no URL changes.
//
// POST /api/projects/{projectId}/page-folders — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsPageFoldersService) Create(ctx context.Context, projectID string, body *ProjectsPageFoldersCreateBody, opts ...CallOption) (*PageFolder, error) {
	var out *PageFolder
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/page-folders", nil, body, false, "data.folder", &out, opts)
	return out, err
}

// Delete — Deletes a folder of a project's pages list; its pages move to no folder and nothing else changes.
//
// DELETE /api/projects/{projectId}/page-folders/{folderId} — Token permission: delete:projects.
//
// projectID: The project's id.
// folderID: The folder's id.
func (r *ProjectsPageFoldersService) Delete(ctx context.Context, projectID string, folderID string, opts ...CallOption) (*ProjectsPageFoldersDeleteResponse, error) {
	var out *ProjectsPageFoldersDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/projects/"+pathSeg(projectID)+"/page-folders/"+pathSeg(folderID), nil, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the folders of a project's pages list.
//
// GET /api/projects/{projectId}/page-folders — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsPageFoldersService) List(ctx context.Context, projectID string, opts ...CallOption) ([]PageFolder, error) {
	var out []PageFolder
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/page-folders", nil, nil, false, "data.folders", &out, opts)
	return out, err
}

// Move — Files up to 500 pages, with their language versions, in a folder — or in none.
//
// POST /api/projects/{projectId}/page-folders/move — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsPageFoldersService) Move(ctx context.Context, projectID string, body *ProjectsPageFoldersMoveBody, opts ...CallOption) (float64, error) {
	var out float64
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/page-folders/move", nil, body, false, "data.moved", &out, opts)
	return out, err
}

// Update — Renames or moves a folder of a project's pages list.
//
// PATCH /api/projects/{projectId}/page-folders/{folderId} — Token permission: write:projects.
//
// projectID: The project's id.
// folderID: The folder's id.
func (r *ProjectsPageFoldersService) Update(ctx context.Context, projectID string, folderID string, body *ProjectsPageFoldersUpdateBody, opts ...CallOption) (*PageFolder, error) {
	var out *PageFolder
	err := r.c.call(ctx, "PATCH", "/api/projects/"+pathSeg(projectID)+"/page-folders/"+pathSeg(folderID), nil, body, false, "data.folder", &out, opts)
	return out, err
}

// ProjectsRedirectsService — projects.redirects
type ProjectsRedirectsService struct {
	c *Client
}

func newProjectsRedirectsService(c *Client) *ProjectsRedirectsService {
	return &ProjectsRedirectsService{c: c}
}

// Create — Adds a redirect from an old path to a page, a path or another site.
//
// POST /api/projects/{projectId}/redirects — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsRedirectsService) Create(ctx context.Context, projectID string, body *ProjectRedirectInput, opts ...CallOption) (*ProjectRedirect, error) {
	var out *ProjectRedirect
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/redirects", nil, body, false, "data.redirect", &out, opts)
	return out, err
}

// Delete — Removes a redirect.
//
// DELETE /api/projects/{projectId}/redirects/{redirectId} — Token permission: delete:projects.
//
// projectID: The project's id.
// redirectID: The redirect's id.
func (r *ProjectsRedirectsService) Delete(ctx context.Context, projectID string, redirectID string, opts ...CallOption) (*ProjectsRedirectsDeleteResponseData, error) {
	var out *ProjectsRedirectsDeleteResponseData
	err := r.c.call(ctx, "DELETE", "/api/projects/"+pathSeg(projectID)+"/redirects/"+pathSeg(redirectID), nil, nil, false, "data", &out, opts)
	return out, err
}

// List — Lists a project's redirects with their hits, and the pages one can point at.
//
// Redirects answer only where no live page does; at most 1000 per project.
//
// GET /api/projects/{projectId}/redirects — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsRedirectsService) List(ctx context.Context, projectID string, opts ...CallOption) ([]ProjectRedirect, error) {
	var out []ProjectRedirect
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/redirects", nil, nil, false, "data.redirects", &out, opts)
	return out, err
}

// Update — Changes a redirect; an automatic one given a new target or old path becomes a manual one.
//
// PATCH /api/projects/{projectId}/redirects/{redirectId} — Token permission: write:projects.
//
// projectID: The project's id.
// redirectID: The redirect's id.
func (r *ProjectsRedirectsService) Update(ctx context.Context, projectID string, redirectID string, body *ProjectsRedirectsUpdateBody, opts ...CallOption) (*ProjectRedirect, error) {
	var out *ProjectRedirect
	err := r.c.call(ctx, "PATCH", "/api/projects/"+pathSeg(projectID)+"/redirects/"+pathSeg(redirectID), nil, body, false, "data.redirect", &out, opts)
	return out, err
}

// ProjectsResultsService — projects.results
type ProjectsResultsService struct {
	c *Client
}

func newProjectsResultsService(c *Client) *ProjectsResultsService {
	return &ProjectsResultsService{c: c}
}

// Audiences — Returns the site results split by audience (new or returning, device, source), for every page or one.
//
// GET /api/projects/{projectId}/results/audiences — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsResultsService) Audiences(ctx context.Context, projectID string, query *ProjectsResultsAudiencesQuery, opts ...CallOption) (*ProjectsResultsAudiencesResponseData, error) {
	var out *ProjectsResultsAudiencesResponseData
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/results/audiences", query, nil, false, "data", &out, opts)
	return out, err
}

// Engagement — Returns how long visitors stayed on a page and how far they scrolled.
//
// GET /api/projects/{projectId}/results/engagement — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsResultsService) Engagement(ctx context.Context, projectID string, query *ProjectsResultsEngagementQuery, opts ...CallOption) (*ProjectsResultsEngagementResponseData, error) {
	var out *ProjectsResultsEngagementResponseData
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/results/engagement", query, nil, false, "data", &out, opts)
	return out, err
}

// Get — Returns the site results — views, blocks seen, clicks and outcomes — of every page, or of one page's blocks.
//
// GET /api/projects/{projectId}/results — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsResultsService) Get(ctx context.Context, projectID string, query *ProjectsResultsGetQuery, opts ...CallOption) (*ProjectsResultsGetResponseData, error) {
	var out *ProjectsResultsGetResponseData
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/results", query, nil, false, "data", &out, opts)
	return out, err
}

// Impact — Returns how a page's results changed after each of its last publishes, or after one.
//
// GET /api/projects/{projectId}/results/impact — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsResultsService) Impact(ctx context.Context, projectID string, query *ProjectsResultsImpactQuery, opts ...CallOption) (*ProjectsResultsImpactResponseData, error) {
	var out *ProjectsResultsImpactResponseData
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/results/impact", query, nil, false, "data", &out, opts)
	return out, err
}

// UpdateSettings — Turns audience results on or off for every page of the project.
//
// PUT /api/projects/{projectId}/site-results-settings — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsResultsService) UpdateSettings(ctx context.Context, projectID string, body *ProjectsResultsUpdateSettingsBody, opts ...CallOption) (*ProjectsResultsUpdateSettingsResponseData, error) {
	var out *ProjectsResultsUpdateSettingsResponseData
	err := r.c.call(ctx, "PUT", "/api/projects/"+pathSeg(projectID)+"/site-results-settings", nil, body, false, "data", &out, opts)
	return out, err
}

// ProjectsSharedBlocksService — projects.sharedBlocks
type ProjectsSharedBlocksService struct {
	c *Client
}

func newProjectsSharedBlocksService(c *Client) *ProjectsSharedBlocksService {
	return &ProjectsSharedBlocksService{c: c}
}

// Create — Creates a shared block from a component tree.
//
// Needs `component.create`. The block is a draft until published with `projects.sharedBlocks.publish`.
//
// POST /api/projects/{projectId}/shared-blocks — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsSharedBlocksService) Create(ctx context.Context, projectID string, body *ProjectsSharedBlocksCreateBody, opts ...CallOption) (*SharedBlock, error) {
	var out *SharedBlock
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/shared-blocks", nil, body, false, "data", &out, opts)
	return out, err
}

// Delete — Moves a shared block to the trash.
//
// Refused with 409 while pages still use it, unless `force=true` — then those pages lose the block.
//
// DELETE /api/projects/{projectId}/shared-blocks/{blockId} — Token permission: delete:projects.
//
// projectID: The project's id.
// blockID: The shared block's id.
func (r *ProjectsSharedBlocksService) Delete(ctx context.Context, projectID string, blockID string, query *ProjectsSharedBlocksDeleteQuery, opts ...CallOption) (*ProjectsSharedBlocksDeleteResponseData, error) {
	var out *ProjectsSharedBlocksDeleteResponseData
	err := r.c.call(ctx, "DELETE", "/api/projects/"+pathSeg(projectID)+"/shared-blocks/"+pathSeg(blockID), query, nil, false, "data", &out, opts)
	return out, err
}

// Get — Returns a shared block with its published versions and how many pages use it.
//
// GET /api/projects/{projectId}/shared-blocks/{blockId} — Token permission: read:projects.
//
// projectID: The project's id.
// blockID: The shared block's id.
func (r *ProjectsSharedBlocksService) Get(ctx context.Context, projectID string, blockID string, opts ...CallOption) (*ProjectsSharedBlocksGetResponseData, error) {
	var out *ProjectsSharedBlocksGetResponseData
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/shared-blocks/"+pathSeg(blockID), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetVersion — Returns one published version of a shared block, with its tree.
//
// GET /api/projects/{projectId}/shared-blocks/{blockId}/versions/{versionId} — Token permission: read:projects.
//
// projectID: The project's id.
// blockID: The shared block's id.
// versionID: The version's id.
func (r *ProjectsSharedBlocksService) GetVersion(ctx context.Context, projectID string, blockID string, versionID string, opts ...CallOption) (*SharedBlockVersion, error) {
	var out *SharedBlockVersion
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/shared-blocks/"+pathSeg(blockID)+"/versions/"+pathSeg(versionID), nil, nil, false, "data", &out, opts)
	return out, err
}

// List — Lists a project's shared blocks — a header, footer or section used on many pages.
//
// GET /api/projects/{projectId}/shared-blocks — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsSharedBlocksService) List(ctx context.Context, projectID string, query *ProjectsSharedBlocksListQuery, opts ...CallOption) ([]SharedBlockSummary, error) {
	var out []SharedBlockSummary
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/shared-blocks", query, nil, false, "data", &out, opts)
	return out, err
}

// Publish — Publishes a shared block's draft as a new version, which the pages using it then show.
//
// POST /api/projects/{projectId}/shared-blocks/{blockId}/publish — Token permission: write:projects.
//
// projectID: The project's id.
// blockID: The shared block's id.
func (r *ProjectsSharedBlocksService) Publish(ctx context.Context, projectID string, blockID string, body *ProjectsSharedBlocksPublishBody, opts ...CallOption) (*ProjectsSharedBlocksPublishResponseData, error) {
	var out *ProjectsSharedBlocksPublishResponseData
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/shared-blocks/"+pathSeg(blockID)+"/publish", nil, body, false, "data", &out, opts)
	return out, err
}

// Update — Changes a shared block's name, description or draft tree.
//
// Nothing live changes until the block is published.
//
// PATCH /api/projects/{projectId}/shared-blocks/{blockId} — Token permission: write:projects.
//
// projectID: The project's id.
// blockID: The shared block's id.
func (r *ProjectsSharedBlocksService) Update(ctx context.Context, projectID string, blockID string, body *ProjectsSharedBlocksUpdateBody, opts ...CallOption) (*SharedBlock, error) {
	var out *SharedBlock
	err := r.c.call(ctx, "PATCH", "/api/projects/"+pathSeg(projectID)+"/shared-blocks/"+pathSeg(blockID), nil, body, false, "data", &out, opts)
	return out, err
}

// ProjectsSiteHealthService — projects.siteHealth
type ProjectsSiteHealthService struct {
	c *Client
}

func newProjectsSiteHealthService(c *Client) *ProjectsSiteHealthService {
	return &ProjectsSiteHealthService{c: c}
}

// Check — Checks the project's site now and returns the findings.
//
// Needs `page.update`; rate-limited.
//
// POST /api/projects/{projectId}/site-health/check — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsSiteHealthService) Check(ctx context.Context, projectID string, opts ...CallOption) (*ProjectsSiteHealthCheckResponseData, error) {
	var out *ProjectsSiteHealthCheckResponseData
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/site-health/check", nil, nil, false, "data", &out, opts)
	return out, err
}

// Dismiss — Dismisses a site-health finding.
//
// POST /api/projects/{projectId}/site-health/{findingId}/dismiss — Token permission: write:projects.
//
// projectID: The project's id.
// findingID: The finding's id.
func (r *ProjectsSiteHealthService) Dismiss(ctx context.Context, projectID string, findingID string, opts ...CallOption) (*ProjectsSiteHealthDismissResponseData, error) {
	var out *ProjectsSiteHealthDismissResponseData
	err := r.c.call(ctx, "POST", "/api/projects/"+pathSeg(projectID)+"/site-health/"+pathSeg(findingID)+"/dismiss", nil, nil, false, "data", &out, opts)
	return out, err
}

// Get — Returns the problems found on the project's live site — broken links, missing texts and the like.
//
// GET /api/projects/{projectId}/site-health — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsSiteHealthService) Get(ctx context.Context, projectID string, query *ProjectsSiteHealthGetQuery, opts ...CallOption) (*ProjectsSiteHealthGetResponseData, error) {
	var out *ProjectsSiteHealthGetResponseData
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/site-health", query, nil, false, "data", &out, opts)
	return out, err
}

// ProjectsSpamProtectionService — projects.spamProtection
type ProjectsSpamProtectionService struct {
	c *Client
}

func newProjectsSpamProtectionService(c *Client) *ProjectsSpamProtectionService {
	return &ProjectsSpamProtectionService{c: c}
}

// Get — Returns how the project's forms are protected from spam, what applies, and what the platform allows.
//
// A provider's secret is never answered, only whether one is stored.
//
// GET /api/projects/{projectId}/spam-protection — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsSpamProtectionService) Get(ctx context.Context, projectID string, opts ...CallOption) (*ProjectsSpamProtectionGetResponseData, error) {
	var out *ProjectsSpamProtectionGetResponseData
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/spam-protection", nil, nil, false, "data", &out, opts)
	return out, err
}

// Update — Changes how the project's forms are protected from spam; only the fields sent change.
//
// PUT /api/projects/{projectId}/spam-protection — Token permission: write:projects.
//
// projectID: The project's id.
func (r *ProjectsSpamProtectionService) Update(ctx context.Context, projectID string, body *ProjectsSpamProtectionUpdateBody, opts ...CallOption) (*ProjectsSpamProtectionUpdateResponseData, error) {
	var out *ProjectsSpamProtectionUpdateResponseData
	err := r.c.call(ctx, "PUT", "/api/projects/"+pathSeg(projectID)+"/spam-protection", nil, body, false, "data", &out, opts)
	return out, err
}

// ProjectsWaitlistService — projects.waitlist
type ProjectsWaitlistService struct {
	c *Client
}

func newProjectsWaitlistService(c *Client) *ProjectsWaitlistService {
	return &ProjectsWaitlistService{c: c}
}

// DeleteSignup — Erases a waitlist sign-up with its consent record.
//
// DELETE /api/projects/{projectId}/waitlist/{signupId} — Token permission: delete:projects.
//
// projectID: The project's id.
// signupID: The sign-up's id.
func (r *ProjectsWaitlistService) DeleteSignup(ctx context.Context, projectID string, signupID string, opts ...CallOption) (*ProjectsWaitlistDeleteSignupResponse, error) {
	var out *ProjectsWaitlistDeleteSignupResponse
	err := r.c.call(ctx, "DELETE", "/api/projects/"+pathSeg(projectID)+"/waitlist/"+pathSeg(signupID), nil, nil, false, "", &out, opts)
	return out, err
}

// Export — Downloads the project's waitlist sign-ups, with their consent records, as a CSV file.
//
// Needs `project.update`; at most 50,000 rows; written to the audit log. Answers the file itself (`text/csv`).
//
// GET /api/projects/{projectId}/waitlist/export — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsWaitlistService) Export(ctx context.Context, projectID string, query *ProjectsWaitlistExportQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/waitlist/export", query, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the project's waitlist blocks with their counts, and one page of sign-ups.
//
// GET /api/projects/{projectId}/waitlist — Token permission: read:projects.
//
// projectID: The project's id.
func (r *ProjectsWaitlistService) List(ctx context.Context, projectID string, query *ProjectsWaitlistListQuery, opts ...CallOption) (*ProjectsWaitlistListResponseData, error) {
	var out *ProjectsWaitlistListResponseData
	err := r.c.call(ctx, "GET", "/api/projects/"+pathSeg(projectID)+"/waitlist", query, nil, false, "data", &out, opts)
	return out, err
}

// PushService — Push
type PushService struct {
	c *Client
}

func newPushService(c *Client) *PushService {
	return &PushService{c: c}
}

// GetPublicKey — Returns the platform's public VAPID key, which a browser needs to subscribe to its push notifications.
//
// Null when push notifications are not set up on this platform.
//
// GET /api/push/send — No token needed.
func (r *PushService) GetPublicKey(ctx context.Context, opts ...CallOption) (string, error) {
	var out string
	err := r.c.call(ctx, "GET", "/api/push/send", nil, nil, false, "publicKey", &out, opts)
	return out, err
}

// Subscribe — Registers a browser's push subscription, so the caller's notifications reach that device.
//
// Pass the browser's `PushSubscription` as JSON, made with the key from `push.getPublicKey`. Registering the same endpoint again updates its keys; an endpoint already registered by somebody else is left as it is. At most 30 calls in five minutes.
//
// POST /api/push/subscribe
func (r *PushService) Subscribe(ctx context.Context, body *PushSubscribeBody, opts ...CallOption) (*PushSubscribeResponse, error) {
	var out *PushSubscribeResponse
	err := r.c.call(ctx, "POST", "/api/push/subscribe", nil, body, false, "", &out, opts)
	return out, err
}

// Unsubscribe — Removes a device's push subscription of the caller.
//
// DELETE /api/push/subscribe
func (r *PushService) Unsubscribe(ctx context.Context, query *PushUnsubscribeQuery, opts ...CallOption) (*PushUnsubscribeResponse, error) {
	var out *PushUnsubscribeResponse
	err := r.c.call(ctx, "DELETE", "/api/push/subscribe", query, nil, false, "", &out, opts)
	return out, err
}

// ReferralsService — Referrals: The platform's referral programme as one organisation takes part in it: each member with `referrals:share` has a personal link; a referred organisation that pays and stays past the refund window qualifies, and the referrer's organisation earns the programme's rewards and tiers. The platform sets the programme; everything here needs Referrals on the plan and membership of the organisation.
type ReferralsService struct {
	c *Client
}

func newReferralsService(c *Client) *ReferralsService {
	return &ReferralsService{c: c}
}

// Get — Returns the programme as the caller's organisation takes part in it: the caller's link, the rewards, the tiers and, with `referrals:read`, the organisation's referrals.
//
// Calling it gives the caller a code in this organisation when they hold `referrals:share` and have none yet. A referral is masked: its initial and dates, never a name or an address.
//
// GET /api/referrals
func (r *ReferralsService) Get(ctx context.Context, query *ReferralsGetQuery, opts ...CallOption) (*ReferralOverview, error) {
	var out *ReferralOverview
	err := r.c.call(ctx, "GET", "/api/referrals", query, nil, false, "data", &out, opts)
	return out, err
}

// RegenerateCode — Gives the caller a new personal referral code in the organisation; the old link stops working at once.
//
// Needs `referrals:share`. The referrals made through the old link keep their history. 409 while the programme is not running.
//
// POST /api/referrals/code
func (r *ReferralsService) RegenerateCode(ctx context.Context, query *ReferralsRegenerateCodeQuery, opts ...CallOption) (*ReferralsRegenerateCodeResponseData, error) {
	var out *ReferralsRegenerateCodeResponseData
	err := r.c.call(ctx, "POST", "/api/referrals/code", query, nil, false, "data", &out, opts)
	return out, err
}

// RegionalSettingsService — Regional settings
type RegionalSettingsService struct {
	c *Client
}

func newRegionalSettingsService(c *Client) *RegionalSettingsService {
	return &RegionalSettingsService{c: c}
}

// Get — Returns the platform's default regional settings: time zone, date, time and number formats, first day of the week, currency and measurement system.
//
// These are the platform-wide defaults a person's own regional settings fall back to, not the caller's own choices.
//
// GET /api/regional-settings
func (r *RegionalSettingsService) Get(ctx context.Context, opts ...CallOption) (*RegionalSettings, error) {
	var out *RegionalSettings
	err := r.c.call(ctx, "GET", "/api/regional-settings", nil, nil, false, "", &out, opts)
	return out, err
}

// ReportsService — Reports: Reports of widgets — charts, tables and figures over databases, endpoints and the hubs' numbers — with their shares, mail schedules, public links and snapshots. Reports need `analytics:reports` on the organization's plan. A report the caller may not open answers 404, like one that does not exist.
type ReportsService struct {
	c *Client

	PublicLinks *ReportsPublicLinksService
	Schedules   *ReportsSchedulesService
	Shares      *ReportsSharesService
	Snapshots   *ReportsSnapshotsService
	Widgets     *ReportsWidgetsService
}

func newReportsService(c *Client) *ReportsService {
	return &ReportsService{c: c, PublicLinks: newReportsPublicLinksService(c), Schedules: newReportsSchedulesService(c), Shares: newReportsSharesService(c), Snapshots: newReportsSnapshotsService(c), Widgets: newReportsWidgetsService(c)}
}

// Annotations — Lists the organization's events in a period — page publishes, deploys, notable audit events, calendar events — to mark on a time chart.
//
// Needs `report.read` and `analytics:report_annotations` on the plan. Calendar events only from calendars the caller may open, and a private one only when they may see its content. At most 400, oldest first.
//
// GET /api/reports/{reportId}/annotations
//
// reportID: The report's id.
func (r *ReportsService) Annotations(ctx context.Context, reportID string, query *ReportsAnnotationsQuery, opts ...CallOption) ([]ReportsAnnotationsResponseDataItemsItem, error) {
	var out []ReportsAnnotationsResponseDataItemsItem
	err := r.c.call(ctx, "GET", "/api/reports/"+pathSeg(reportID)+"/annotations", query, nil, false, "data.items", &out, opts)
	return out, err
}

// Consumers — Lists the dashboards that show a widget of this report, to know what a change would affect.
//
// Needs `report.read` and edit access. Dashboards are named by id and owner only.
//
// GET /api/reports/{reportId}/consumers
//
// reportID: The report's id.
func (r *ReportsService) Consumers(ctx context.Context, reportID string, opts ...CallOption) ([]ReportsConsumersResponseDataItemsItem, error) {
	var out []ReportsConsumersResponseDataItemsItem
	err := r.c.call(ctx, "GET", "/api/reports/"+pathSeg(reportID)+"/consumers", nil, nil, false, "data.items", &out, opts)
	return out, err
}

// Create — Creates a report, empty or from a template.
//
// Needs `report.create` in the organization (and access to the workspace, which must be the organization's). `crm_sales` makes the CRM's Sales report (needs the CRM forecast on the plan; over `crmSpaceId`, or every deals space of `workspaceId`); `invoicing` makes the Invoicing overview (needs Invoicing access). 30 a minute per person.
//
// POST /api/reports
func (r *ReportsService) Create(ctx context.Context, body *ReportsCreateBody, opts ...CallOption) (*Report, error) {
	var out *Report
	err := r.c.call(ctx, "POST", "/api/reports", nil, body, false, "report", &out, opts)
	return out, err
}

// Data — Reads the rows of a widget's data source, as the caller.
//
// Needs `report.read`. Name a stored widget (`widgetId`) to run its saved source, or send a `dataSource` to try one. Sources of the hubs are read with the caller's access. 120 a minute per person.
//
// POST /api/reports/{reportId}/data
//
// reportID: The report's id.
func (r *ReportsService) Data(ctx context.Context, reportID string, body *ReportsDataBody, opts ...CallOption) (*ReportsDataResponseData, error) {
	var out *ReportsDataResponseData
	err := r.c.call(ctx, "POST", "/api/reports/"+pathSeg(reportID)+"/data", nil, body, false, "data", &out, opts)
	return out, err
}

// Delete — Deletes a report with its widgets.
//
// Needs `report.delete` and the right to delete it. Permanent.
//
// DELETE /api/reports/{reportId}
//
// reportID: The report's id.
func (r *ReportsService) Delete(ctx context.Context, reportID string, opts ...CallOption) (*ReportsDeleteResponse, error) {
	var out *ReportsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/reports/"+pathSeg(reportID), nil, nil, false, "", &out, opts)
	return out, err
}

// Export — Downloads a report as a PDF, or its tables as a CSV file, made with the caller's data, language and regional settings.
//
// Needs `report.read`. Answers the file itself (`application/pdf` or `text/csv`). 20 a minute per person.
//
// GET /api/reports/{reportId}/export
//
// reportID: The report's id.
func (r *ReportsService) Export(ctx context.Context, reportID string, query *ReportsExportQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/reports/"+pathSeg(reportID)+"/export", query, nil, false, "", &out, opts)
	return out, err
}

// Get — Returns a report with its widgets and what the caller may do with it.
//
// Needs `report.read`.
//
// GET /api/reports/{reportId}
//
// reportID: The report's id.
func (r *ReportsService) Get(ctx context.Context, reportID string, opts ...CallOption) (*ReportsGetResponse, error) {
	var out *ReportsGetResponse
	err := r.c.call(ctx, "GET", "/api/reports/"+pathSeg(reportID), nil, nil, false, "", &out, opts)
	return out, err
}

// List — Lists the reports of an organization the caller may open, most recently changed first.
//
// Needs `report.read`. The caller's private reports, the workspace reports of workspaces they can open, the organization's reports, and those shared with them or their teams.
//
// GET /api/reports
func (r *ReportsService) List(ctx context.Context, query *ReportsListQuery, opts ...CallOption) ([]Report, error) {
	var out []Report
	err := r.c.call(ctx, "GET", "/api/reports", query, nil, false, "reports", &out, opts)
	return out, err
}

// SourceHealth — Returns each widget's data-source health — green, amber or red as last read.
//
// Needs `report.read`. Only database, endpoint and URL sources are tracked; the others are `unknown`.
//
// GET /api/reports/{reportId}/source-health
//
// reportID: The report's id.
func (r *ReportsService) SourceHealth(ctx context.Context, reportID string, opts ...CallOption) ([]ReportsSourceHealthResponseDataItemsItem, error) {
	var out []ReportsSourceHealthResponseDataItemsItem
	err := r.c.call(ctx, "GET", "/api/reports/"+pathSeg(reportID)+"/source-health", nil, nil, false, "data.items", &out, opts)
	return out, err
}

// SourceHealthRollup — Counts the data sources of an organization's reports by health and lists up to 20 that need attention.
//
// Needs `report.read`, membership and `analytics:report_sources_dashboard` on the plan. Only reports the caller may open count.
//
// GET /api/reports/source-health-rollup
func (r *ReportsService) SourceHealthRollup(ctx context.Context, query *ReportsSourceHealthRollupQuery, opts ...CallOption) (*ReportsSourceHealthRollupResponseData, error) {
	var out *ReportsSourceHealthRollupResponseData
	err := r.c.call(ctx, "GET", "/api/reports/source-health-rollup", query, nil, false, "data", &out, opts)
	return out, err
}

// Update — Changes a report's name, description, visibility, workspace, tags or favourite mark.
//
// Needs `report.update` and edit access. Send `expectedUpdatedAt` (or an `If-Unmodified-Since` header) with the `updatedAt` you read: a report saved meanwhile answers 409 `conflictReload` with `currentUpdatedAt`.
//
// PATCH /api/reports/{reportId}
//
// reportID: The report's id.
func (r *ReportsService) Update(ctx context.Context, reportID string, body *ReportsUpdateBody, opts ...CallOption) (*Report, error) {
	var out *Report
	err := r.c.call(ctx, "PATCH", "/api/reports/"+pathSeg(reportID), nil, body, false, "report", &out, opts)
	return out, err
}

// ReportsPublicLinksService — reports.publicLinks
type ReportsPublicLinksService struct {
	c *Client
}

func newReportsPublicLinksService(c *Client) *ReportsPublicLinksService {
	return &ReportsPublicLinksService{c: c}
}

// Create — Makes a read-only public link to a report and answers its URL, once.
//
// Needs `report.read`, and only the report's owner may — a link shows the owner's view of the data. Only the token's hash is kept: `url` cannot be shown again. Whoever opens it sees the report as its owner would; the hubs' figures only with `includeHubData`.
//
// POST /api/reports/{reportId}/public-links
//
// reportID: The report's id.
func (r *ReportsPublicLinksService) Create(ctx context.Context, reportID string, body *ReportsPublicLinksCreateBody, opts ...CallOption) (*ReportsPublicLinksCreateResponse, error) {
	var out *ReportsPublicLinksCreateResponse
	err := r.c.call(ctx, "POST", "/api/reports/"+pathSeg(reportID)+"/public-links", nil, body, false, "", &out, opts)
	return out, err
}

// List — Lists a report's read-only public links, without their tokens.
//
// Needs `report.read`, and only the report's owner may — a link shows the owner's view of the data.
//
// GET /api/reports/{reportId}/public-links
//
// reportID: The report's id.
func (r *ReportsPublicLinksService) List(ctx context.Context, reportID string, opts ...CallOption) ([]ReportPublicLink, error) {
	var out []ReportPublicLink
	err := r.c.call(ctx, "GET", "/api/reports/"+pathSeg(reportID)+"/public-links", nil, nil, false, "links", &out, opts)
	return out, err
}

// Revoke — Revokes a public link; it answers 404 from then on.
//
// Needs `report.read`, and only the report's owner may — a link shows the owner's view of the data.
//
// DELETE /api/reports/{reportId}/public-links/{linkId}
//
// reportID: The report's id.
// linkID: The public link's id.
func (r *ReportsPublicLinksService) Revoke(ctx context.Context, reportID string, linkID string, opts ...CallOption) (*ReportsPublicLinksRevokeResponse, error) {
	var out *ReportsPublicLinksRevokeResponse
	err := r.c.call(ctx, "DELETE", "/api/reports/"+pathSeg(reportID)+"/public-links/"+pathSeg(linkID), nil, nil, false, "", &out, opts)
	return out, err
}

// ReportsSchedulesService — reports.schedules
type ReportsSchedulesService struct {
	c *Client
}

func newReportsSchedulesService(c *Client) *ReportsSchedulesService {
	return &ReportsSchedulesService{c: c}
}

// Create — Mails a report on a schedule — daily, weekly or monthly — to members and, where allowed, external addresses.
//
// Needs `report.read` and the right to share the report: its owner or a platform admin. Sent on the caller's clock and in their name; each member gets the report as they may see it at that moment, and who lost access is left out. At most 10 schedules per report (409 `reportScheduleLimit`).
//
// POST /api/reports/{reportId}/schedules
//
// reportID: The report's id.
func (r *ReportsSchedulesService) Create(ctx context.Context, reportID string, body *ReportScheduleInput, opts ...CallOption) (*ReportSchedule, error) {
	var out *ReportSchedule
	err := r.c.call(ctx, "POST", "/api/reports/"+pathSeg(reportID)+"/schedules", nil, body, false, "schedule", &out, opts)
	return out, err
}

// Delete — Deletes a schedule.
//
// Needs `report.read` and the right to share the report: its owner or a platform admin.
//
// DELETE /api/reports/{reportId}/schedules/{scheduleId}
//
// reportID: The report's id.
// scheduleID: The schedule's id.
func (r *ReportsSchedulesService) Delete(ctx context.Context, reportID string, scheduleID string, opts ...CallOption) (*ReportsSchedulesDeleteResponse, error) {
	var out *ReportsSchedulesDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/reports/"+pathSeg(reportID)+"/schedules/"+pathSeg(scheduleID), nil, nil, false, "", &out, opts)
	return out, err
}

// List — Lists a report's mail schedules, with the organization's members to send to and whether external addresses are allowed.
//
// Needs `report.read` and the right to share the report: its owner or a platform admin.
//
// GET /api/reports/{reportId}/schedules
//
// reportID: The report's id.
func (r *ReportsSchedulesService) List(ctx context.Context, reportID string, opts ...CallOption) (*ReportsSchedulesListResponse, error) {
	var out *ReportsSchedulesListResponse
	err := r.c.call(ctx, "GET", "/api/reports/"+pathSeg(reportID)+"/schedules", nil, nil, false, "", &out, opts)
	return out, err
}

// SendNow — Sends a schedule's report now, to its recipients as they may read it at this moment.
//
// Needs `report.read` and the right to share the report: its owner or a platform admin. Does not count as the day's scheduled send. Usually queued (202, `queued: true`); without a queue it is sent at once and `result` says what happened. 3 per schedule in 10 minutes.
//
// POST /api/reports/{reportId}/schedules/{scheduleId}/send
//
// reportID: The report's id.
// scheduleID: The schedule's id.
func (r *ReportsSchedulesService) SendNow(ctx context.Context, reportID string, scheduleID string, opts ...CallOption) (*ReportsSchedulesSendNowResponse, error) {
	var out *ReportsSchedulesSendNowResponse
	err := r.c.call(ctx, "POST", "/api/reports/"+pathSeg(reportID)+"/schedules/"+pathSeg(scheduleID)+"/send", nil, nil, false, "", &out, opts)
	return out, err
}

// Update — Changes a schedule; `enabled` alone pauses or resumes it.
//
// Needs `report.read` and the right to share the report: its owner or a platform admin. Only the fields sent change; the time stays on the clock of whoever made it.
//
// PATCH /api/reports/{reportId}/schedules/{scheduleId}
//
// reportID: The report's id.
// scheduleID: The schedule's id.
func (r *ReportsSchedulesService) Update(ctx context.Context, reportID string, scheduleID string, body *ReportsSchedulesUpdateBody, opts ...CallOption) (*ReportSchedule, error) {
	var out *ReportSchedule
	err := r.c.call(ctx, "PATCH", "/api/reports/"+pathSeg(reportID)+"/schedules/"+pathSeg(scheduleID), nil, body, false, "schedule", &out, opts)
	return out, err
}

// ReportsSharesService — reports.shares
type ReportsSharesService struct {
	c *Client
}

func newReportsSharesService(c *Client) *ReportsSharesService {
	return &ReportsSharesService{c: c}
}

// Create — Shares a report with a member or a team of its organization, to view or to edit.
//
// Needs `report.read` and the right to share the report: its owner or a platform admin. Exactly one of `userId` and `teamId`. Sharing again with the same person or team changes the permission (200 instead of 201). Shares are read on every open: who leaves the organization loses it.
//
// POST /api/reports/{reportId}/shares
//
// reportID: The report's id.
func (r *ReportsSharesService) Create(ctx context.Context, reportID string, body *ReportsSharesCreateBody, opts ...CallOption) (*ReportShare, error) {
	var out *ReportShare
	err := r.c.call(ctx, "POST", "/api/reports/"+pathSeg(reportID)+"/shares", nil, body, false, "share", &out, opts)
	return out, err
}

// Delete — Ends a share.
//
// Needs `report.read` and the right to share the report: its owner or a platform admin.
//
// DELETE /api/reports/{reportId}/shares/{shareId}
//
// reportID: The report's id.
// shareID: The share's id.
func (r *ReportsSharesService) Delete(ctx context.Context, reportID string, shareID string, opts ...CallOption) (*ReportsSharesDeleteResponse, error) {
	var out *ReportsSharesDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/reports/"+pathSeg(reportID)+"/shares/"+pathSeg(shareID), nil, nil, false, "", &out, opts)
	return out, err
}

// List — Lists who a report is shared with, and the organization's members and teams to choose from.
//
// Needs `report.read` and the right to share the report: its owner or a platform admin.
//
// GET /api/reports/{reportId}/shares
//
// reportID: The report's id.
func (r *ReportsSharesService) List(ctx context.Context, reportID string, opts ...CallOption) (*ReportsSharesListResponse, error) {
	var out *ReportsSharesListResponse
	err := r.c.call(ctx, "GET", "/api/reports/"+pathSeg(reportID)+"/shares", nil, nil, false, "", &out, opts)
	return out, err
}

// Update — Changes a share's permission.
//
// Needs `report.read` and the right to share the report: its owner or a platform admin.
//
// PATCH /api/reports/{reportId}/shares/{shareId}
//
// reportID: The report's id.
// shareID: The share's id.
func (r *ReportsSharesService) Update(ctx context.Context, reportID string, shareID string, body *ReportsSharesUpdateBody, opts ...CallOption) (*ReportShare, error) {
	var out *ReportShare
	err := r.c.call(ctx, "PATCH", "/api/reports/"+pathSeg(reportID)+"/shares/"+pathSeg(shareID), nil, body, false, "share", &out, opts)
	return out, err
}

// ReportsSnapshotsService — reports.snapshots
type ReportsSnapshotsService struct {
	c *Client
}

func newReportsSnapshotsService(c *Client) *ReportsSnapshotsService {
	return &ReportsSnapshotsService{c: c}
}

// Create — Freezes every widget's current data as a snapshot, to compare against later.
//
// Needs `report.read` and edit access. Widgets whose source is read as the viewer (CRM, Invoicing, Products, Banking) are recorded without data. 6 an hour per person and report.
//
// POST /api/reports/{reportId}/snapshots
//
// reportID: The report's id.
func (r *ReportsSnapshotsService) Create(ctx context.Context, reportID string, body *ReportsSnapshotsCreateBody, opts ...CallOption) (*ReportsSnapshotsCreateResponseData, error) {
	var out *ReportsSnapshotsCreateResponseData
	err := r.c.call(ctx, "POST", "/api/reports/"+pathSeg(reportID)+"/snapshots", nil, body, false, "data", &out, opts)
	return out, err
}

// Get — Returns a snapshot with each widget's captured data.
//
// Needs `report.read`. Data read as the person who captured it is shown to them alone.
//
// GET /api/reports/{reportId}/snapshots/{snapshotId}
//
// reportID: The report's id.
// snapshotID: The snapshot's id.
func (r *ReportsSnapshotsService) Get(ctx context.Context, reportID string, snapshotID string, opts ...CallOption) (*ReportSnapshot, error) {
	var out *ReportSnapshot
	err := r.c.call(ctx, "GET", "/api/reports/"+pathSeg(reportID)+"/snapshots/"+pathSeg(snapshotID), nil, nil, false, "data", &out, opts)
	return out, err
}

// List — Lists a report's last 50 snapshots, newest first.
//
// Needs `report.read`.
//
// GET /api/reports/{reportId}/snapshots
//
// reportID: The report's id.
func (r *ReportsSnapshotsService) List(ctx context.Context, reportID string, opts ...CallOption) ([]ReportSnapshotSummary, error) {
	var out []ReportSnapshotSummary
	err := r.c.call(ctx, "GET", "/api/reports/"+pathSeg(reportID)+"/snapshots", nil, nil, false, "data.items", &out, opts)
	return out, err
}

// Sign — Signs a snapshot, so later changes to its captured data can be detected.
//
// Needs `report.read`, edit access and `analytics:report_signed_snapshots` on the plan. Once per snapshot (409 `snapshotAlreadySigned`).
//
// POST /api/reports/{reportId}/snapshots/{snapshotId}/sign
//
// reportID: The report's id.
// snapshotID: The snapshot's id.
func (r *ReportsSnapshotsService) Sign(ctx context.Context, reportID string, snapshotID string, opts ...CallOption) (*ReportsSnapshotsSignResponseData, error) {
	var out *ReportsSnapshotsSignResponseData
	err := r.c.call(ctx, "POST", "/api/reports/"+pathSeg(reportID)+"/snapshots/"+pathSeg(snapshotID)+"/sign", nil, nil, false, "data", &out, opts)
	return out, err
}

// Verify — Checks a snapshot's signature against its captured data.
//
// Needs `report.read`.
//
// GET /api/reports/{reportId}/snapshots/{snapshotId}/sign
//
// reportID: The report's id.
// snapshotID: The snapshot's id.
func (r *ReportsSnapshotsService) Verify(ctx context.Context, reportID string, snapshotID string, opts ...CallOption) (*ReportsSnapshotsVerifyResponseData, error) {
	var out *ReportsSnapshotsVerifyResponseData
	err := r.c.call(ctx, "GET", "/api/reports/"+pathSeg(reportID)+"/snapshots/"+pathSeg(snapshotID)+"/sign", nil, nil, false, "data", &out, opts)
	return out, err
}

// ReportsWidgetsService — reports.widgets
type ReportsWidgetsService struct {
	c *Client
}

func newReportsWidgetsService(c *Client) *ReportsWidgetsService {
	return &ReportsWidgetsService{c: c}
}

// Create — Adds a widget to a report.
//
// Needs `report.update` and edit access. The position is clamped to the 12-column grid. 60 a minute per person.
//
// POST /api/reports/{reportId}/widgets
//
// reportID: The report's id.
func (r *ReportsWidgetsService) Create(ctx context.Context, reportID string, body *CreateReportWidgetInput, opts ...CallOption) (*ReportWidget, error) {
	var out *ReportWidget
	err := r.c.call(ctx, "POST", "/api/reports/"+pathSeg(reportID)+"/widgets", nil, body, false, "widget", &out, opts)
	return out, err
}

// Delete — Removes a widget from a report.
//
// Needs `report.update` and edit access.
//
// DELETE /api/reports/{reportId}/widgets/{widgetId}
//
// reportID: The report's id.
// widgetID: The widget's id.
func (r *ReportsWidgetsService) Delete(ctx context.Context, reportID string, widgetID string, opts ...CallOption) (*ReportsWidgetsDeleteResponse, error) {
	var out *ReportsWidgetsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/reports/"+pathSeg(reportID)+"/widgets/"+pathSeg(widgetID), nil, nil, false, "", &out, opts)
	return out, err
}

// History — Lists the last 50 changes to a widget's settings, newest first, field by field.
//
// Needs `report.read` and edit access.
//
// GET /api/reports/{reportId}/widgets/{widgetId}/history
//
// reportID: The report's id.
// widgetID: The widget's id.
func (r *ReportsWidgetsService) History(ctx context.Context, reportID string, widgetID string, opts ...CallOption) ([]ReportsWidgetsHistoryResponseDataItemsItem, error) {
	var out []ReportsWidgetsHistoryResponseDataItemsItem
	err := r.c.call(ctx, "GET", "/api/reports/"+pathSeg(reportID)+"/widgets/"+pathSeg(widgetID)+"/history", nil, nil, false, "data.items", &out, opts)
	return out, err
}

// List — Lists a report's widgets in grid order.
//
// Needs `report.read`.
//
// GET /api/reports/{reportId}/widgets
//
// reportID: The report's id.
func (r *ReportsWidgetsService) List(ctx context.Context, reportID string, opts ...CallOption) ([]ReportWidget, error) {
	var out []ReportWidget
	err := r.c.call(ctx, "GET", "/api/reports/"+pathSeg(reportID)+"/widgets", nil, nil, false, "widgets", &out, opts)
	return out, err
}

// Observe — Has the platform's AI read a widget's data against its prompt and keeps the one-line observation on the widget.
//
// Needs `report.read`, a `prompt` on the widget and `analytics:report_ai_insights` on the plan; counts against the organization's AI quota (402 `aiQuotaExceeded`). Send the rows the widget shows. 10 a minute per person and widget.
//
// POST /api/reports/{reportId}/widgets/{widgetId}/observe
//
// reportID: The report's id.
// widgetID: The widget's id.
func (r *ReportsWidgetsService) Observe(ctx context.Context, reportID string, widgetID string, body *ReportsWidgetsObserveBody, opts ...CallOption) (*ReportsWidgetsObserveResponseData, error) {
	var out *ReportsWidgetsObserveResponseData
	err := r.c.call(ctx, "POST", "/api/reports/"+pathSeg(reportID)+"/widgets/"+pathSeg(widgetID)+"/observe", nil, body, false, "data", &out, opts)
	return out, err
}

// Update — Changes a widget's title, type, data source, options, position or AI prompt.
//
// Needs `report.update` and edit access. Each change is kept in the widget's history.
//
// PATCH /api/reports/{reportId}/widgets/{widgetId}
//
// reportID: The report's id.
// widgetID: The widget's id.
func (r *ReportsWidgetsService) Update(ctx context.Context, reportID string, widgetID string, body *UpdateReportWidgetInput, opts ...CallOption) (*ReportWidget, error) {
	var out *ReportWidget
	err := r.c.call(ctx, "PATCH", "/api/reports/"+pathSeg(reportID)+"/widgets/"+pathSeg(widgetID), nil, body, false, "widget", &out, opts)
	return out, err
}

// UpdateLayout — Moves and resizes several widgets at once, all or none.
//
// Needs `report.update` and edit access. Widgets not of this report are left alone.
//
// PUT /api/reports/{reportId}/widgets
//
// reportID: The report's id.
func (r *ReportsWidgetsService) UpdateLayout(ctx context.Context, reportID string, body *ReportsWidgetsUpdateLayoutBody, opts ...CallOption) (*ReportsWidgetsUpdateLayoutResponse, error) {
	var out *ReportsWidgetsUpdateLayoutResponse
	err := r.c.call(ctx, "PUT", "/api/reports/"+pathSeg(reportID)+"/widgets", nil, body, false, "", &out, opts)
	return out, err
}

// RetentionPoliciesService — Retention policies: How long an organisation keeps each kind of record: after `durationDays` from its creation or last change a record is either deleted by the nightly sweep (`autoDelete`) or locked until then. A workspace policy overrides the organisation-wide one for its type. Business records (content entries, bookings, booking pages, webhooks) are only counted until somebody confirms the count.
type RetentionPoliciesService struct {
	c *Client
}

func newRetentionPoliciesService(c *Client) *RetentionPoliciesService {
	return &RetentionPoliciesService{c: c}
}

// Check — Tells whether a retention policy applies to a record with the given dates, and its deadline.
//
// Needs membership of the organisation. The workspace's policy wins over the organisation-wide one.
//
// GET /api/retention-policies/check
func (r *RetentionPoliciesService) Check(ctx context.Context, query *RetentionPoliciesCheckQuery, opts ...CallOption) (*RetentionPoliciesCheckResponseData, error) {
	var out *RetentionPoliciesCheckResponseData
	err := r.c.call(ctx, "GET", "/api/retention-policies/check", query, nil, false, "data", &out, opts)
	return out, err
}

// CheckCascade — Tells whether deleting a record is blocked because records under it are still locked by a policy.
//
// Covers an endpoint's API keys and webhooks and a workflow's executions. Needs membership of the organisation.
//
// GET /api/retention-policies/check-cascade
func (r *RetentionPoliciesService) CheckCascade(ctx context.Context, query *RetentionPoliciesCheckCascadeQuery, opts ...CallOption) (*RetentionPoliciesCheckCascadeResponseData, error) {
	var out *RetentionPoliciesCheckCascadeResponseData
	err := r.c.call(ctx, "GET", "/api/retention-policies/check-cascade", query, nil, false, "data", &out, opts)
	return out, err
}

// Confirm — Confirms a business-record policy's count, so the nightly sweep starts deleting.
//
// Needs `retention_policies:manage` in the policy's organisation (and workspace, for a workspace policy). Answers 409 while the policy has no count yet or is not waiting for confirmation. Written to the audit log.
//
// POST /api/retention-policies/{id}/confirm
//
// id: The retention policy's id.
func (r *RetentionPoliciesService) Confirm(ctx context.Context, id string, opts ...CallOption) (*RetentionPolicy, error) {
	var out *RetentionPolicy
	err := r.c.call(ctx, "POST", "/api/retention-policies/"+pathSeg(id)+"/confirm", nil, nil, false, "data", &out, opts)
	return out, err
}

// Create — Creates a retention policy for one kind of record.
//
// Needs `retention_policies:manage` in the policy's organisation (and workspace, for a workspace policy). A policy for a business-record type starts counting (`deletionGate: "preview"`) and deletes nothing until `retentionPolicies.confirm`. Pages are never deleted by a policy.
//
// POST /api/retention-policies
func (r *RetentionPoliciesService) Create(ctx context.Context, body *RetentionPoliciesCreateBody, opts ...CallOption) (*RetentionPolicy, error) {
	var out *RetentionPolicy
	err := r.c.call(ctx, "POST", "/api/retention-policies", nil, body, false, "data", &out, opts)
	return out, err
}

// Delete — Deletes a retention policy.
//
// Needs `retention_policies:manage` in the policy's organisation (and workspace, for a workspace policy).
//
// DELETE /api/retention-policies/{id}
//
// id: The retention policy's id.
func (r *RetentionPoliciesService) Delete(ctx context.Context, id string, opts ...CallOption) (*RetentionPoliciesDeleteResponseData, error) {
	var out *RetentionPoliciesDeleteResponseData
	err := r.c.call(ctx, "DELETE", "/api/retention-policies/"+pathSeg(id), nil, nil, false, "data", &out, opts)
	return out, err
}

// DryRun — Counts what a policy would delete now, without deleting, and keeps the result.
//
// Needs `retention_policies:manage` and the platform's `retention_dry_run` feature.
//
// POST /api/retention-policies/{id}/dry-run
//
// id: The retention policy's id.
func (r *RetentionPoliciesService) DryRun(ctx context.Context, id string, body *RetentionPoliciesDryRunBody, opts ...CallOption) (*RetentionPoliciesDryRunResponseData, error) {
	var out *RetentionPoliciesDryRunResponseData
	err := r.c.call(ctx, "POST", "/api/retention-policies/"+pathSeg(id)+"/dry-run", nil, body, false, "data", &out, opts)
	return out, err
}

// Get — Returns one retention policy.
//
// Needs `retention_policies:manage` in the policy's organisation (and workspace, for a workspace policy).
//
// GET /api/retention-policies/{id}
//
// id: The retention policy's id.
func (r *RetentionPoliciesService) Get(ctx context.Context, id string, opts ...CallOption) (*RetentionPolicy, error) {
	var out *RetentionPolicy
	err := r.c.call(ctx, "GET", "/api/retention-policies/"+pathSeg(id), nil, nil, false, "data", &out, opts)
	return out, err
}

// List — Lists one page of an organisation's retention policies.
//
// Needs `retention_policies:manage`. With `workspaceId`, the workspace's policies and the organisation-wide ones.
//
// GET /api/retention-policies
func (r *RetentionPoliciesService) List(ctx context.Context, query *RetentionPoliciesListQuery, opts ...CallOption) ([]RetentionPolicy, error) {
	var out []RetentionPolicy
	err := r.c.call(ctx, "GET", "/api/retention-policies", query, nil, false, "data.data", &out, opts)
	return out, err
}

// ListDryRuns — Lists a policy's last 20 dry runs.
//
// Needs `retention_policies:manage` and the platform's `retention_dry_run` feature.
//
// GET /api/retention-policies/{id}/dry-run
//
// id: The retention policy's id.
func (r *RetentionPoliciesService) ListDryRuns(ctx context.Context, id string, opts ...CallOption) ([]RetentionPoliciesListDryRunsResponseDataDryRunsItem, error) {
	var out []RetentionPoliciesListDryRunsResponseDataDryRunsItem
	err := r.c.call(ctx, "GET", "/api/retention-policies/"+pathSeg(id)+"/dry-run", nil, nil, false, "data.dryRuns", &out, opts)
	return out, err
}

// Update — Changes a retention policy's period, basis, mode, state or description.
//
// Needs `retention_policies:manage` in the policy's organisation (and workspace, for a workspace policy). Changing the period, basis or mode of a business-record policy resets its confirmation and counts again.
//
// PUT /api/retention-policies/{id}
//
// id: The retention policy's id.
func (r *RetentionPoliciesService) Update(ctx context.Context, id string, body *RetentionPoliciesUpdateBody, opts ...CallOption) (*RetentionPolicy, error) {
	var out *RetentionPolicy
	err := r.c.call(ctx, "PUT", "/api/retention-policies/"+pathSeg(id), nil, body, false, "data", &out, opts)
	return out, err
}

// RolesService — Roles
type RolesService struct {
	c *Client

	Assignments *RolesAssignmentsService
}

func newRolesService(c *Client) *RolesService {
	return &RolesService{c: c, Assignments: newRolesAssignmentsService(c)}
}

// Create — Creates a role in an organization.
//
// Needs `role.create`. Permissions the organization's plan does not include, or the caller does not hold, are refused.
//
// POST /api/roles
func (r *RolesService) Create(ctx context.Context, body *CreateRoleInput, opts ...CallOption) (*Role, error) {
	var out *Role
	err := r.c.call(ctx, "POST", "/api/roles", nil, body, false, "data", &out, opts)
	return out, err
}

// Delete — Deletes a role.
//
// DELETE /api/roles/{roleId}
//
// roleID: The role's id.
func (r *RolesService) Delete(ctx context.Context, roleID string, opts ...CallOption) (*RolesDeleteResponse, error) {
	var out *RolesDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/roles/"+pathSeg(roleID), nil, nil, false, "", &out, opts)
	return out, err
}

// Get — Returns one role with its permissions.
//
// GET /api/roles/{roleId}
//
// roleID: The role's id.
func (r *RolesService) Get(ctx context.Context, roleID string, opts ...CallOption) (*Role, error) {
	var out *Role
	err := r.c.call(ctx, "GET", "/api/roles/"+pathSeg(roleID), nil, nil, false, "data", &out, opts)
	return out, err
}

// Grantable — Lists the permissions the caller may put on a role in an organization.
//
// What the organization's plan includes and, unless the caller is a platform admin, what the caller holds.
//
// GET /api/roles/grantable
func (r *RolesService) Grantable(ctx context.Context, query *RolesGrantableQuery, opts ...CallOption) (*RolesGrantableResponseData, error) {
	var out *RolesGrantableResponseData
	err := r.c.call(ctx, "GET", "/api/roles/grantable", query, nil, false, "data", &out, opts)
	return out, err
}

// List — Lists an organization's roles, or the platform's system roles without `organizationId`.
//
// GET /api/roles
func (r *RolesService) List(ctx context.Context, query *RolesListQuery, opts ...CallOption) ([]Role, error) {
	var out []Role
	err := r.c.call(ctx, "GET", "/api/roles", query, nil, false, "data.roles", &out, opts)
	return out, err
}

// Update — Changes a role's name, description or permissions.
//
// Needs `role.update`; the same limits on permissions as when creating.
//
// PUT /api/roles/{roleId}
//
// roleID: The role's id.
func (r *RolesService) Update(ctx context.Context, roleID string, body *UpdateRoleInput, opts ...CallOption) (*Role, error) {
	var out *Role
	err := r.c.call(ctx, "PUT", "/api/roles/"+pathSeg(roleID), nil, body, false, "data", &out, opts)
	return out, err
}

// RolesAssignmentsService — roles.assignments
type RolesAssignmentsService struct {
	c *Client
}

func newRolesAssignmentsService(c *Client) *RolesAssignmentsService {
	return &RolesAssignmentsService{c: c}
}

// Create — Gives a person a role in an organization, workspace or team.
//
// Needs `role.assign`; only roles within the caller's own permissions.
//
// POST /api/role-assignments
func (r *RolesAssignmentsService) Create(ctx context.Context, body *RolesAssignmentsCreateBody, opts ...CallOption) (*RoleAssignment, error) {
	var out *RoleAssignment
	err := r.c.call(ctx, "POST", "/api/role-assignments", nil, body, false, "data", &out, opts)
	return out, err
}

// Delete — Takes a role assignment away.
//
// DELETE /api/role-assignments/{assignmentId}
//
// assignmentID: The role assignment's id.
func (r *RolesAssignmentsService) Delete(ctx context.Context, assignmentID string, opts ...CallOption) (*RolesAssignmentsDeleteResponse, error) {
	var out *RolesAssignmentsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/role-assignments/"+pathSeg(assignmentID), nil, nil, false, "", &out, opts)
	return out, err
}

// List — Lists role assignments — who holds which role where — in an organization or workspace, or of one person.
//
// GET /api/role-assignments
func (r *RolesAssignmentsService) List(ctx context.Context, query *RolesAssignmentsListQuery, opts ...CallOption) ([]RoleAssignment, error) {
	var out []RoleAssignment
	err := r.c.call(ctx, "GET", "/api/role-assignments", query, nil, false, "data", &out, opts)
	return out, err
}

// ShareService — Share links
type ShareService struct {
	c *Client
}

func newShareService(c *Client) *ShareService {
	return &ShareService{c: c}
}

// GetAsset — Opens a share link of a file or folder and returns what it shares.
//
// Needs no token. A password-protected link answers 401 with `passwordRequired: true` until the password is sent in the `X-Share-Password` header; a revoked or expired link answers 410.
//
// GET /api/share/asset/{token} — No token needed.
//
// token: The share link's token.
func (r *ShareService) GetAsset(ctx context.Context, token string, opts ...CallOption) (*ShareGetAssetResponse, error) {
	var out *ShareGetAssetResponse
	err := r.c.call(ctx, "GET", "/api/share/asset/"+pathSeg(token), nil, nil, false, "", &out, opts)
	return out, err
}

// SocialService — Social: The Social hub: connected accounts, posts to one or more networks (drafted, scheduled or published now), their comments, assignments and repeats, the inbox of replies and mentions, analytics and the AI writing helps. Everything needs the `social` permission named on each operation and the Social hub on the plan. A row is seen by whom its visibility says: `private` its creator only, `workspace` that workspace's members, `organization` everyone in it.
type SocialService struct {
	c *Client
}

func newSocialService(c *Client) *SocialService {
	return &SocialService{c: c}
}

// AccountInsights — Returns a connected account's insights from its network: followers, reach and the like.
//
// Needs `social:view_analytics`. Instagram, Threads, Bluesky, Pinterest and YouTube; other networks answer `supported: false`. Cached for 15 minutes.
//
// GET /api/social/insights/account/{accountId}
//
// accountID: The connected account's id.
func (r *SocialService) AccountInsights(ctx context.Context, accountID string, query *SocialAccountInsightsQuery, opts ...CallOption) (*SocialAccountInsightsResponse, error) {
	var out *SocialAccountInsightsResponse
	err := r.c.call(ctx, "GET", "/api/social/insights/account/"+pathSeg(accountID), query, nil, false, "", &out, opts)
	return out, err
}

// AddComment — Adds an internal comment to a post; @-mentioned colleagues are notified.
//
// Needs `social:read`.
//
// POST /api/social/posts/{postId}/comments
//
// postID: The social post's id.
func (r *SocialService) AddComment(ctx context.Context, postID string, body *SocialAddCommentBody, opts ...CallOption) (*SocialPostComment, error) {
	var out *SocialPostComment
	err := r.c.call(ctx, "POST", "/api/social/posts/"+pathSeg(postID)+"/comments", nil, body, false, "comment", &out, opts)
	return out, err
}

// AICaption — Writes a caption for a topic with the platform AI, fitted to the strictest network's length.
//
// Needs `social:create` and AI captions on the plan; refused when the person opted out of AI. Counts against the AI tokens.
//
// POST /api/social/ai/caption
func (r *SocialService) AICaption(ctx context.Context, body *SocialAiCaptionBody, opts ...CallOption) (*SocialAiCaptionResponse, error) {
	var out *SocialAiCaptionResponse
	err := r.c.call(ctx, "POST", "/api/social/ai/caption", nil, body, false, "", &out, opts)
	return out, err
}

// AIHashtags — Suggests hashtags for a caption or topic with the platform AI.
//
// Needs `social:create` and AI captions on the plan; counts against the AI tokens.
//
// POST /api/social/ai/hashtags
func (r *SocialService) AIHashtags(ctx context.Context, body *SocialAiHashtagsBody, opts ...CallOption) ([]string, error) {
	var out []string
	err := r.c.call(ctx, "POST", "/api/social/ai/hashtags", nil, body, false, "hashtags", &out, opts)
	return out, err
}

// AIVariations — Writes one caption per network with the platform AI, each in that network's voice and length.
//
// Needs `social:create` and AI captions on the plan; counts against the AI tokens. Use each as the network's `perNetwork.textOverride`.
//
// POST /api/social/ai/variations
func (r *SocialService) AIVariations(ctx context.Context, body *SocialAiVariationsBody, opts ...CallOption) (*SocialAiVariationsResponse, error) {
	var out *SocialAiVariationsResponse
	err := r.c.call(ctx, "POST", "/api/social/ai/variations", nil, body, false, "", &out, opts)
	return out, err
}

// AnalyticsSummary — Returns an organization's posting figures for a period: outcomes per network, the busiest hours and weekdays, and the networks' reach and engagement.
//
// Needs `social:view_analytics` and analytics on the plan. Hours and weekdays are counted in the caller's time zone (their profile's, else `timeZone`). `provider` comes from the insights last fetched for each post.
//
// GET /api/social/analytics/summary
func (r *SocialService) AnalyticsSummary(ctx context.Context, query *SocialAnalyticsSummaryQuery, opts ...CallOption) (*SocialAnalyticsSummaryResponse, error) {
	var out *SocialAnalyticsSummaryResponse
	err := r.c.call(ctx, "GET", "/api/social/analytics/summary", query, nil, false, "", &out, opts)
	return out, err
}

// Assign — Assigns a post to a member of its organization, replacing the current assignment.
//
// Needs `social:read`. The assignee is notified in the app and by email.
//
// POST /api/social/posts/{postId}/assignment
//
// postID: The social post's id.
func (r *SocialService) Assign(ctx context.Context, postID string, body *SocialAssignBody, opts ...CallOption) (*SocialPostAssignment, error) {
	var out *SocialPostAssignment
	err := r.c.call(ctx, "POST", "/api/social/posts/"+pathSeg(postID)+"/assignment", nil, body, false, "assignment", &out, opts)
	return out, err
}

// BestTime — Suggests three good times to post on a network, from the organization's history, with the platform AI.
//
// Needs `social:view_analytics`; counts against the organization's AI tokens. Without history the suggestion rests on the network's usual audience.
//
// POST /api/social/analytics/best-time
func (r *SocialService) BestTime(ctx context.Context, body *SocialBestTimeBody, opts ...CallOption) (*SocialBestTimeResponse, error) {
	var out *SocialBestTimeResponse
	err := r.c.call(ctx, "POST", "/api/social/analytics/best-time", nil, body, false, "", &out, opts)
	return out, err
}

// BulkCreate — Creates up to 200 posts at once, each a draft or scheduled, through the same accounts.
//
// Needs `social:create`, and `social:schedule` for rows with `scheduledAt`. Each row succeeds or fails on its own; `results` follow the order of `rows`.
//
// POST /api/social/posts/bulk
func (r *SocialService) BulkCreate(ctx context.Context, body *SocialBulkCreateBody, opts ...CallOption) (*SocialBulkCreateResponse, error) {
	var out *SocialBulkCreateResponse
	err := r.c.call(ctx, "POST", "/api/social/posts/bulk", nil, body, false, "", &out, opts)
	return out, err
}

// BulkDelete — Moves up to 200 posts to the trash.
//
// Needs `social:delete`. A post the caller cannot see, or one being published, is skipped.
//
// POST /api/social/posts/bulk-delete
func (r *SocialService) BulkDelete(ctx context.Context, body *SocialBulkDeleteBody, opts ...CallOption) (*SocialBulkDeleteResponse, error) {
	var out *SocialBulkDeleteResponse
	err := r.c.call(ctx, "POST", "/api/social/posts/bulk-delete", nil, body, false, "", &out, opts)
	return out, err
}

// BulkUpdate — Sets the campaign or the assignee of up to 200 posts.
//
// Needs `social:update`. A post the caller cannot see is skipped; the rest go on.
//
// POST /api/social/posts/bulk-update
func (r *SocialService) BulkUpdate(ctx context.Context, body *SocialBulkUpdateBody, opts ...CallOption) (*SocialBulkUpdateResponse, error) {
	var out *SocialBulkUpdateResponse
	err := r.c.call(ctx, "POST", "/api/social/posts/bulk-update", nil, body, false, "", &out, opts)
	return out, err
}

// Cancel — Cancels a scheduled post before it is published.
//
// Needs `social:schedule`.
//
// POST /api/social/posts/{postId}/cancel
//
// postID: The social post's id.
func (r *SocialService) Cancel(ctx context.Context, postID string, opts ...CallOption) (*SocialPost, error) {
	var out *SocialPost
	err := r.c.call(ctx, "POST", "/api/social/posts/"+pathSeg(postID)+"/cancel", nil, nil, false, "post", &out, opts)
	return out, err
}

// ClonePost — Copies a post into a new draft with the same text, media, options and accounts.
//
// Needs `social:create`. The copy is not scheduled.
//
// POST /api/social/posts/{postId}/clone
//
// postID: The social post's id.
func (r *SocialService) ClonePost(ctx context.Context, postID string, opts ...CallOption) (*SocialPost, error) {
	var out *SocialPost
	err := r.c.call(ctx, "POST", "/api/social/posts/"+pathSeg(postID)+"/clone", nil, nil, false, "post", &out, opts)
	return out, err
}

// ConnectBluesky — Connects a Bluesky account with an app password.
//
// Needs `social:connect_account` and counts against the plan's social accounts. The app password is used once to sign in and never stored; connecting the same account again updates it. The other networks are connected by signing in, in the dashboard.
//
// POST /api/social/accounts/connect/bluesky
func (r *SocialService) ConnectBluesky(ctx context.Context, body *SocialConnectBlueskyBody, opts ...CallOption) (*SocialConnectBlueskyResponse, error) {
	var out *SocialConnectBlueskyResponse
	err := r.c.call(ctx, "POST", "/api/social/accounts/connect/bluesky", nil, body, false, "", &out, opts)
	return out, err
}

// CreateAbTest — Creates two posts that differ only in their text, and a test that picks the better one after a while.
//
// Needs `social:create`. Both variants share accounts, media and campaign and publish like any post; the winner is decided by `metric` once `evaluateAfterHours` have passed.
//
// POST /api/social/posts/ab-test
func (r *SocialService) CreateAbTest(ctx context.Context, body *SocialCreateAbTestBody, query *SocialCreateAbTestQuery, opts ...CallOption) (*SocialCreateAbTestResponse, error) {
	var out *SocialCreateAbTestResponse
	err := r.c.call(ctx, "POST", "/api/social/posts/ab-test", query, body, false, "", &out, opts)
	return out, err
}

// CreateCampaign — Creates a campaign to tag posts with.
//
// Needs `social:create`.
//
// POST /api/social/campaigns
func (r *SocialService) CreateCampaign(ctx context.Context, body *SocialCreateCampaignBody, opts ...CallOption) (*SocialCampaign, error) {
	var out *SocialCampaign
	err := r.c.call(ctx, "POST", "/api/social/campaigns", nil, body, false, "campaign", &out, opts)
	return out, err
}

// CreateHashtagSet — Saves a set of hashtags.
//
// Needs `social:create`. Each tag gets a leading `#`; empty and repeated ones are dropped.
//
// POST /api/social/hashtag-sets
func (r *SocialService) CreateHashtagSet(ctx context.Context, body *SocialCreateHashtagSetBody, opts ...CallOption) (*SocialHashtagSet, error) {
	var out *SocialHashtagSet
	err := r.c.call(ctx, "POST", "/api/social/hashtag-sets", nil, body, false, "hashtagSet", &out, opts)
	return out, err
}

// CreatePost — Creates a social post as a draft, or scheduled when `scheduledAt` is given.
//
// Needs `social:create`, and `social:schedule` to schedule. A `private` account is posted through only by the person who connected it. Publishing now is POST /api/social/posts/{postId}/publish. `warnings` names networks whose length limit the text exceeds; it is not refused.
//
// POST /api/social/posts
func (r *SocialService) CreatePost(ctx context.Context, body *CreateSocialPostInput, query *SocialCreatePostQuery, opts ...CallOption) (*SocialPost, error) {
	var out *SocialPost
	err := r.c.call(ctx, "POST", "/api/social/posts", query, body, false, "post", &out, opts)
	return out, err
}

// CreateSnippet — Saves a reusable text snippet.
//
// Needs `social:create`.
//
// POST /api/social/content-library
func (r *SocialService) CreateSnippet(ctx context.Context, body *SocialCreateSnippetBody, query *SocialCreateSnippetQuery, opts ...CallOption) (*SocialSnippet, error) {
	var out *SocialSnippet
	err := r.c.call(ctx, "POST", "/api/social/content-library", query, body, false, "snippet", &out, opts)
	return out, err
}

// CreateTemplate — Saves a post template.
//
// Needs `social:create`.
//
// POST /api/social/templates
func (r *SocialService) CreateTemplate(ctx context.Context, body *SocialCreateTemplateBody, opts ...CallOption) (*SocialPostTemplate, error) {
	var out *SocialPostTemplate
	err := r.c.call(ctx, "POST", "/api/social/templates", nil, body, false, "template", &out, opts)
	return out, err
}

// DeleteAutoRepost — Stops repeating a post.
//
// Needs `social:update`.
//
// DELETE /api/social/posts/{postId}/auto-repost
//
// postID: The social post's id.
func (r *SocialService) DeleteAutoRepost(ctx context.Context, postID string, opts ...CallOption) (*SocialDeleteAutoRepostResponse, error) {
	var out *SocialDeleteAutoRepostResponse
	err := r.c.call(ctx, "DELETE", "/api/social/posts/"+pathSeg(postID)+"/auto-repost", nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteCampaign — Deletes a campaign; its posts keep existing without it.
//
// Needs `social:delete`.
//
// DELETE /api/social/campaigns/{campaignId}
//
// campaignID: The campaign's id.
func (r *SocialService) DeleteCampaign(ctx context.Context, campaignID string, opts ...CallOption) (*SocialDeleteCampaignResponse, error) {
	var out *SocialDeleteCampaignResponse
	err := r.c.call(ctx, "DELETE", "/api/social/campaigns/"+pathSeg(campaignID), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteComment — Deletes an internal comment.
//
// Needs `social:delete`; only the comment's author or an organization owner or admin.
//
// DELETE /api/social/posts/{postId}/comments/{commentId}
//
// postID: The social post's id.
// commentID: The comment's id.
func (r *SocialService) DeleteComment(ctx context.Context, postID string, commentID string, opts ...CallOption) (*SocialDeleteCommentResponse, error) {
	var out *SocialDeleteCommentResponse
	err := r.c.call(ctx, "DELETE", "/api/social/posts/"+pathSeg(postID)+"/comments/"+pathSeg(commentID), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteHashtagSet — Deletes a saved hashtag set.
//
// Needs `social:delete`.
//
// DELETE /api/social/hashtag-sets/{hashtagSetId}
//
// hashtagSetID: The hashtag set's id.
func (r *SocialService) DeleteHashtagSet(ctx context.Context, hashtagSetID string, opts ...CallOption) (*SocialDeleteHashtagSetResponse, error) {
	var out *SocialDeleteHashtagSetResponse
	err := r.c.call(ctx, "DELETE", "/api/social/hashtag-sets/"+pathSeg(hashtagSetID), nil, nil, false, "", &out, opts)
	return out, err
}

// DeletePost — Moves a social post to the trash.
//
// Needs `social:delete`. A scheduled post is taken off the schedule at once.
//
// DELETE /api/social/posts/{postId}
//
// postID: The social post's id.
func (r *SocialService) DeletePost(ctx context.Context, postID string, opts ...CallOption) (*SocialDeletePostResponse, error) {
	var out *SocialDeletePostResponse
	err := r.c.call(ctx, "DELETE", "/api/social/posts/"+pathSeg(postID), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteSnippet — Deletes a text snippet.
//
// Needs `social:delete`.
//
// DELETE /api/social/content-library/{snippetId}
//
// snippetID: The snippet's id.
func (r *SocialService) DeleteSnippet(ctx context.Context, snippetID string, opts ...CallOption) (*SocialDeleteSnippetResponse, error) {
	var out *SocialDeleteSnippetResponse
	err := r.c.call(ctx, "DELETE", "/api/social/content-library/"+pathSeg(snippetID), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteTemplate — Deletes a post template.
//
// Needs `social:delete`.
//
// DELETE /api/social/templates/{templateId}
//
// templateID: The template's id.
func (r *SocialService) DeleteTemplate(ctx context.Context, templateID string, opts ...CallOption) (*SocialDeleteTemplateResponse, error) {
	var out *SocialDeleteTemplateResponse
	err := r.c.call(ctx, "DELETE", "/api/social/templates/"+pathSeg(templateID), nil, nil, false, "", &out, opts)
	return out, err
}

// DisconnectAccount — Disconnects a social account.
//
// Needs `social:connect_account`. Posts already published stay; nothing more can be posted through the account.
//
// DELETE /api/social/accounts/{accountId}
//
// accountID: The connected account's id.
func (r *SocialService) DisconnectAccount(ctx context.Context, accountID string, opts ...CallOption) (*SocialDisconnectAccountResponse, error) {
	var out *SocialDisconnectAccountResponse
	err := r.c.call(ctx, "DELETE", "/api/social/accounts/"+pathSeg(accountID), nil, nil, false, "", &out, opts)
	return out, err
}

// DraftFromContent — Writes a draft post about a live page or database record, with a text for each network.
//
// Needs `social:create`. The text is written from what visitors see (AI when available, else a template); nothing is scheduled or posted. 422 with `reason` when the page or record may not be shared.
//
// POST /api/social/drafts/from-content
func (r *SocialService) DraftFromContent(ctx context.Context, body *SocialDraftFromContentBody, opts ...CallOption) (*SocialDraftFromContentResponse, error) {
	var out *SocialDraftFromContentResponse
	err := r.c.call(ctx, "POST", "/api/social/drafts/from-content", nil, body, false, "", &out, opts)
	return out, err
}

// GetAssignment — Returns who a post is assigned to, or null.
//
// Needs `social:read`.
//
// GET /api/social/posts/{postId}/assignment
//
// postID: The social post's id.
func (r *SocialService) GetAssignment(ctx context.Context, postID string, opts ...CallOption) (*SocialPostAssignment, error) {
	var out *SocialPostAssignment
	err := r.c.call(ctx, "GET", "/api/social/posts/"+pathSeg(postID)+"/assignment", nil, nil, false, "assignment", &out, opts)
	return out, err
}

// GetAutoRepost — Returns a post's repeat schedule, or null.
//
// Needs `social:update`.
//
// GET /api/social/posts/{postId}/auto-repost
//
// postID: The social post's id.
func (r *SocialService) GetAutoRepost(ctx context.Context, postID string, opts ...CallOption) (*SocialAutoRepost, error) {
	var out *SocialAutoRepost
	err := r.c.call(ctx, "GET", "/api/social/posts/"+pathSeg(postID)+"/auto-repost", nil, nil, false, "schedule", &out, opts)
	return out, err
}

// GetPost — Returns one social post with its targets.
//
// Needs `social:read`.
//
// GET /api/social/posts/{postId}
//
// postID: The social post's id.
func (r *SocialService) GetPost(ctx context.Context, postID string, opts ...CallOption) (*SocialPost, error) {
	var out *SocialPost
	err := r.c.call(ctx, "GET", "/api/social/posts/"+pathSeg(postID), nil, nil, false, "post", &out, opts)
	return out, err
}

// Limits — Returns the organization's daily post quota and connected-account quota with what is used.
//
// Needs `social:read`. `-1` means no limit.
//
// GET /api/social/limits
func (r *SocialService) Limits(ctx context.Context, query *SocialLimitsQuery, opts ...CallOption) (*SocialLimitsResponse, error) {
	var out *SocialLimitsResponse
	err := r.c.call(ctx, "GET", "/api/social/limits", query, nil, false, "", &out, opts)
	return out, err
}

// ListAccounts — Lists the connected social accounts of an organization the caller can see, without their tokens.
//
// Needs `social:read`. A private account is listed to the person who connected it only — organization owners and admins included. Also answers which networks the platform offers for connecting.
//
// GET /api/social/accounts
func (r *SocialService) ListAccounts(ctx context.Context, query *SocialListAccountsQuery, opts ...CallOption) ([]SocialAccount, error) {
	var out []SocialAccount
	err := r.c.call(ctx, "GET", "/api/social/accounts", query, nil, false, "accounts", &out, opts)
	return out, err
}

// ListCampaigns — Lists an organization's social campaigns the caller can see.
//
// Needs `social:read`.
//
// GET /api/social/campaigns
func (r *SocialService) ListCampaigns(ctx context.Context, query *SocialListCampaignsQuery, opts ...CallOption) ([]SocialCampaign, error) {
	var out []SocialCampaign
	err := r.c.call(ctx, "GET", "/api/social/campaigns", query, nil, false, "campaigns", &out, opts)
	return out, err
}

// ListComments — Lists the team's internal comments on a post, oldest first.
//
// Needs `social:read`. These are notes between colleagues, not the network's comments (those are in the inbox).
//
// GET /api/social/posts/{postId}/comments
//
// postID: The social post's id.
func (r *SocialService) ListComments(ctx context.Context, postID string, opts ...CallOption) ([]SocialPostComment, error) {
	var out []SocialPostComment
	err := r.c.call(ctx, "GET", "/api/social/posts/"+pathSeg(postID)+"/comments", nil, nil, false, "comments", &out, opts)
	return out, err
}

// ListHashtagSets — Lists an organization's saved hashtag sets the caller can see.
//
// Needs `social:read`.
//
// GET /api/social/hashtag-sets
func (r *SocialService) ListHashtagSets(ctx context.Context, query *SocialListHashtagSetsQuery, opts ...CallOption) ([]SocialHashtagSet, error) {
	var out []SocialHashtagSet
	err := r.c.call(ctx, "GET", "/api/social/hashtag-sets", query, nil, false, "hashtagSets", &out, opts)
	return out, err
}

// ListInbox — Lists the comments, replies, mentions and messages received on an organization's accounts, newest first.
//
// Needs `social:read`. Also answers the unread count and, per connected network, whether the inbox can read it (`coverage`).
//
// GET /api/social/inbox
func (r *SocialService) ListInbox(ctx context.Context, query *SocialListInboxQuery, opts ...CallOption) ([]SocialInboxItem, error) {
	var out []SocialInboxItem
	err := r.c.call(ctx, "GET", "/api/social/inbox", query, nil, false, "items", &out, opts)
	return out, err
}

// ListPosts — Lists an organization's social posts the caller can see, newest first.
//
// Needs `social:read`. At most 500 posts; each carries `assignedTo`, the user it is assigned to.
//
// GET /api/social/posts
func (r *SocialService) ListPosts(ctx context.Context, query *SocialListPostsQuery, opts ...CallOption) ([]SocialListPostsResponsePostsItem, error) {
	var out []SocialListPostsResponsePostsItem
	err := r.c.call(ctx, "GET", "/api/social/posts", query, nil, false, "posts", &out, opts)
	return out, err
}

// ListSnippets — Lists the reusable text snippets of an organization the caller can see, most recently used first.
//
// Needs `social:read`.
//
// GET /api/social/content-library
func (r *SocialService) ListSnippets(ctx context.Context, query *SocialListSnippetsQuery, opts ...CallOption) ([]SocialSnippet, error) {
	var out []SocialSnippet
	err := r.c.call(ctx, "GET", "/api/social/content-library", query, nil, false, "snippets", &out, opts)
	return out, err
}

// ListTemplates — Lists an organization's post templates the caller can see.
//
// Needs `social:read`.
//
// GET /api/social/templates
func (r *SocialService) ListTemplates(ctx context.Context, query *SocialListTemplatesQuery, opts ...CallOption) ([]SocialPostTemplate, error) {
	var out []SocialPostTemplate
	err := r.c.call(ctx, "GET", "/api/social/templates", query, nil, false, "templates", &out, opts)
	return out, err
}

// MarkInboxRead — Marks an inbox item read.
//
// Needs `social:read`.
//
// POST /api/social/inbox/{itemId}/read
//
// itemID: The inbox item's id.
func (r *SocialService) MarkInboxRead(ctx context.Context, itemID string, opts ...CallOption) (*SocialInboxItem, error) {
	var out *SocialInboxItem
	err := r.c.call(ctx, "POST", "/api/social/inbox/"+pathSeg(itemID)+"/read", nil, nil, false, "item", &out, opts)
	return out, err
}

// MarkInboxUnread — Marks an inbox item unread.
//
// Needs `social:read`.
//
// DELETE /api/social/inbox/{itemId}/read
//
// itemID: The inbox item's id.
func (r *SocialService) MarkInboxUnread(ctx context.Context, itemID string, opts ...CallOption) (*SocialInboxItem, error) {
	var out *SocialInboxItem
	err := r.c.call(ctx, "DELETE", "/api/social/inbox/"+pathSeg(itemID)+"/read", nil, nil, false, "item", &out, opts)
	return out, err
}

// PatchPost — Changes a draft or scheduled post.
//
// Needs `social:update`. A post that was published, failed or cancelled cannot be changed (400) — clone it instead. New `targetAccountIds` reset every target to pending.
//
// PATCH /api/social/posts/{postId}
//
// postID: The social post's id.
func (r *SocialService) PatchPost(ctx context.Context, postID string, body *SocialPatchPostBody, opts ...CallOption) (*SocialPost, error) {
	var out *SocialPost
	err := r.c.call(ctx, "PATCH", "/api/social/posts/"+pathSeg(postID), nil, body, false, "post", &out, opts)
	return out, err
}

// PollInbox — Fetches new inbox items from the networks now instead of waiting for the hourly run.
//
// Needs `social:create`. Each network's per-account limit applies as on the schedule.
//
// POST /api/social/inbox/poll
func (r *SocialService) PollInbox(ctx context.Context, query *SocialPollInboxQuery, opts ...CallOption) (*SocialPollInboxResponse, error) {
	var out *SocialPollInboxResponse
	err := r.c.call(ctx, "POST", "/api/social/inbox/poll", query, nil, false, "", &out, opts)
	return out, err
}

// PostInsights — Returns each published target's insights from its network: impressions, likes, comments and the like.
//
// Needs `social:view_analytics`. A network without insights answers `supported: false` for its target. A fetch with `refresh=1` stores a snapshot, which the analytics summary reads.
//
// GET /api/social/insights/post/{postId}
//
// postID: The social post's id.
func (r *SocialService) PostInsights(ctx context.Context, postID string, query *SocialPostInsightsQuery, opts ...CallOption) ([]SocialPostInsightsResponseTargetsItem, error) {
	var out []SocialPostInsightsResponseTargetsItem
	err := r.c.call(ctx, "GET", "/api/social/insights/post/"+pathSeg(postID), query, nil, false, "targets", &out, opts)
	return out, err
}

// PublishNow — Publishes a post now, or retries the targets that failed.
//
// Needs `social:publish_now` and a verified email. Targets already published are never published again, so it is safe to repeat. Counts once against the daily post quota whatever the number of networks. 403 `publishNeedsApproval` while an approval holds the post; 409 while it is publishing; 400 for a cancelled post.
//
// POST /api/social/posts/{postId}/publish
//
// postID: The social post's id.
func (r *SocialService) PublishNow(ctx context.Context, postID string, opts ...CallOption) (*PublishResult, error) {
	var out *PublishResult
	err := r.c.call(ctx, "POST", "/api/social/posts/"+pathSeg(postID)+"/publish", nil, nil, false, "", &out, opts)
	return out, err
}

// ReplyInbox — Replies to an inbox item on its network.
//
// Needs `social:create`. The reply is posted through the account that received the item.
//
// POST /api/social/inbox/{itemId}/reply
//
// itemID: The inbox item's id.
func (r *SocialService) ReplyInbox(ctx context.Context, itemID string, body *SocialReplyInboxBody, opts ...CallOption) (*SocialReplyInboxResponse, error) {
	var out *SocialReplyInboxResponse
	err := r.c.call(ctx, "POST", "/api/social/inbox/"+pathSeg(itemID)+"/reply", nil, body, false, "", &out, opts)
	return out, err
}

// RetryStory — Publishes a story post again to the targets that failed.
//
// Needs `social:publish_now` and stories on the plan. Only a post of kind `story`; targets already published are left alone.
//
// POST /api/social/posts/{postId}/retry-story
//
// postID: The social post's id.
func (r *SocialService) RetryStory(ctx context.Context, postID string, opts ...CallOption) (*SocialRetryStoryResponse, error) {
	var out *SocialRetryStoryResponse
	err := r.c.call(ctx, "POST", "/api/social/posts/"+pathSeg(postID)+"/retry-story", nil, nil, false, "", &out, opts)
	return out, err
}

// Schedule — Schedules a draft post, or moves a scheduled one to another time.
//
// Needs `social:schedule`. A failed or partly published post cannot be scheduled again; clone it.
//
// POST /api/social/posts/{postId}/schedule
//
// postID: The social post's id.
func (r *SocialService) Schedule(ctx context.Context, postID string, body *SocialScheduleBody, opts ...CallOption) (*SocialPost, error) {
	var out *SocialPost
	err := r.c.call(ctx, "POST", "/api/social/posts/"+pathSeg(postID)+"/schedule", nil, body, false, "post", &out, opts)
	return out, err
}

// SetAutoRepost — Repeats a published post at an interval, creating or changing its schedule.
//
// Needs `social:update`. A draft cannot be repeated (400); publish it first.
//
// PUT /api/social/posts/{postId}/auto-repost
//
// postID: The social post's id.
func (r *SocialService) SetAutoRepost(ctx context.Context, postID string, body *SocialSetAutoRepostBody, opts ...CallOption) (*SocialAutoRepost, error) {
	var out *SocialAutoRepost
	err := r.c.call(ctx, "PUT", "/api/social/posts/"+pathSeg(postID)+"/auto-repost", nil, body, false, "schedule", &out, opts)
	return out, err
}

// ShareCheck — Tells whether a live page or database record can be shared, what the post would link to, and through which accounts.
//
// Needs `social:create` and access to the page or record (404 otherwise).
//
// GET /api/social/drafts/from-content
func (r *SocialService) ShareCheck(ctx context.Context, query *SocialShareCheckQuery, opts ...CallOption) (*SocialShareCheckResponse, error) {
	var out *SocialShareCheckResponse
	err := r.c.call(ctx, "GET", "/api/social/drafts/from-content", query, nil, false, "", &out, opts)
	return out, err
}

// Unassign — Releases a post's current assignment.
//
// Needs `social:read`.
//
// DELETE /api/social/posts/{postId}/assignment
//
// postID: The social post's id.
func (r *SocialService) Unassign(ctx context.Context, postID string, opts ...CallOption) (*SocialUnassignResponse, error) {
	var out *SocialUnassignResponse
	err := r.c.call(ctx, "DELETE", "/api/social/posts/"+pathSeg(postID)+"/assignment", nil, nil, false, "", &out, opts)
	return out, err
}

// UpdateAccount — Changes a connected account's visibility, workspace or default Pinterest board.
//
// Needs `social:connect_account`. Tokens are never taken here; reconnect the account instead.
//
// PATCH /api/social/accounts/{accountId}
//
// accountID: The connected account's id.
func (r *SocialService) UpdateAccount(ctx context.Context, accountID string, body *SocialUpdateAccountBody, opts ...CallOption) (*SocialAccount, error) {
	var out *SocialAccount
	err := r.c.call(ctx, "PATCH", "/api/social/accounts/"+pathSeg(accountID), nil, body, false, "account", &out, opts)
	return out, err
}

// UpdateSnippet — Changes a text snippet.
//
// Needs `social:update`.
//
// PATCH /api/social/content-library/{snippetId}
//
// snippetID: The snippet's id.
func (r *SocialService) UpdateSnippet(ctx context.Context, snippetID string, body *SocialUpdateSnippetBody, opts ...CallOption) (*SocialSnippet, error) {
	var out *SocialSnippet
	err := r.c.call(ctx, "PATCH", "/api/social/content-library/"+pathSeg(snippetID), nil, body, false, "snippet", &out, opts)
	return out, err
}

// SubmissionsService — Submissions
type SubmissionsService struct {
	c *Client
}

func newSubmissionsService(c *Client) *SubmissionsService {
	return &SubmissionsService{c: c}
}

// Analytics — Returns a project's submission figures: totals by status, category and form, response times and SLA breaches per day.
//
// GET /api/submissions/analytics
func (r *SubmissionsService) Analytics(ctx context.Context, query *SubmissionsAnalyticsQuery, opts ...CallOption) (*SubmissionsAnalyticsResponseData, error) {
	var out *SubmissionsAnalyticsResponseData
	err := r.c.call(ctx, "GET", "/api/submissions/analytics", query, nil, false, "data", &out, opts)
	return out, err
}

// BulkDelete — Moves up to 500 submissions to the trash at once.
//
// Only submissions of projects the caller can open are touched; the others are skipped. They stay restorable in the trash for the organisation's retention days.
//
// DELETE /api/submissions/bulk
func (r *SubmissionsService) BulkDelete(ctx context.Context, query *SubmissionsBulkDeleteQuery, opts ...CallOption) (*SubmissionsBulkDeleteResponseData, error) {
	var out *SubmissionsBulkDeleteResponseData
	err := r.c.call(ctx, "DELETE", "/api/submissions/bulk", query, nil, false, "data", &out, opts)
	return out, err
}

// BulkUpdate — Changes the status, priority or assignee of up to 500 submissions, or adds a tag to them.
//
// Submissions the caller may not see are left alone.
//
// POST /api/submissions/bulk
func (r *SubmissionsService) BulkUpdate(ctx context.Context, body *SubmissionsBulkUpdateBody, opts ...CallOption) (float64, error) {
	var out float64
	err := r.c.call(ctx, "POST", "/api/submissions/bulk", nil, body, false, "data.updated", &out, opts)
	return out, err
}

// Create — Records a form submission, as if a visitor had sent the form.
//
// Needs `submissions.manage`.
//
// POST /api/submissions
func (r *SubmissionsService) Create(ctx context.Context, body *SubmissionsCreateBody, opts ...CallOption) (*Submission, error) {
	var out *Submission
	err := r.c.call(ctx, "POST", "/api/submissions", nil, body, false, "data.submission", &out, opts)
	return out, err
}

// CreateCannedResponse — Saves a reply for reuse, for the organization or one project.
//
// POST /api/submissions/canned-responses
func (r *SubmissionsService) CreateCannedResponse(ctx context.Context, body *SubmissionsCreateCannedResponseBody, opts ...CallOption) (*SubmissionCannedResponse, error) {
	var out *SubmissionCannedResponse
	err := r.c.call(ctx, "POST", "/api/submissions/canned-responses", nil, body, false, "data.template", &out, opts)
	return out, err
}

// Delete — Moves a submission and its thread to the trash.
//
// DELETE /api/submissions/{id}
//
// id: The submission's id.
func (r *SubmissionsService) Delete(ctx context.Context, id string, opts ...CallOption) (*SubmissionsDeleteResponseData, error) {
	var out *SubmissionsDeleteResponseData
	err := r.c.call(ctx, "DELETE", "/api/submissions/"+pathSeg(id), nil, nil, false, "data", &out, opts)
	return out, err
}

// DeleteCannedResponse — Deletes a saved reply.
//
// DELETE /api/submissions/canned-responses/{cannedResponseId}
//
// cannedResponseID: The saved reply's id.
func (r *SubmissionsService) DeleteCannedResponse(ctx context.Context, cannedResponseID string, opts ...CallOption) (*SubmissionsDeleteCannedResponseResponse, error) {
	var out *SubmissionsDeleteCannedResponseResponse
	err := r.c.call(ctx, "DELETE", "/api/submissions/canned-responses/"+pathSeg(cannedResponseID), nil, nil, false, "", &out, opts)
	return out, err
}

// Enrich — Runs the AI summary, category and priority of a submission again.
//
// Needs the plan's AI enrichment; counts against the organization's AI budget.
//
// POST /api/submissions/{id}/enrich
//
// id: The submission's id.
func (r *SubmissionsService) Enrich(ctx context.Context, id string, opts ...CallOption) (*Submission, error) {
	var out *Submission
	err := r.c.call(ctx, "POST", "/api/submissions/"+pathSeg(id)+"/enrich", nil, nil, false, "data.submission", &out, opts)
	return out, err
}

// ErasePersonalData — Erases every submission one email address sent, for good (a data subject's right to erasure).
//
// Immediate and permanent — erased submissions do not go to the trash. Only projects the caller can open are searched; `confirm` must be `true`.
//
// DELETE /api/submissions/gdpr
func (r *SubmissionsService) ErasePersonalData(ctx context.Context, query *SubmissionsErasePersonalDataQuery, opts ...CallOption) (*SubmissionsErasePersonalDataResponseData, error) {
	var out *SubmissionsErasePersonalDataResponseData
	err := r.c.call(ctx, "DELETE", "/api/submissions/gdpr", query, nil, false, "data", &out, opts)
	return out, err
}

// Export — Downloads the submissions the caller may see as a CSV file.
//
// Needs `submissions.export`. Answers the file itself (`text/csv`).
//
// GET /api/submissions/export
func (r *SubmissionsService) Export(ctx context.Context, query *SubmissionsExportQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/submissions/export", query, nil, false, "", &out, opts)
	return out, err
}

// ExportPersonalData — Exports every submission and message from one email address — a data subject's access request.
//
// Needs `submissions.export`.
//
// POST /api/submissions/gdpr
func (r *SubmissionsService) ExportPersonalData(ctx context.Context, body *SubmissionsExportPersonalDataBody, opts ...CallOption) (*SubmissionsExportPersonalDataResponseData, error) {
	var out *SubmissionsExportPersonalDataResponseData
	err := r.c.call(ctx, "POST", "/api/submissions/gdpr", nil, body, false, "data", &out, opts)
	return out, err
}

// Get — Returns a submission with its reply thread; opening a new one marks it read when the caller may update submissions.
//
// GET /api/submissions/{id}
//
// id: The submission's id.
func (r *SubmissionsService) Get(ctx context.Context, id string, opts ...CallOption) (*SubmissionsGetResponseData, error) {
	var out *SubmissionsGetResponseData
	err := r.c.call(ctx, "GET", "/api/submissions/"+pathSeg(id), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetFormSettings — Returns a form's inbox settings: notifications, auto-tags, auto-reply, SLA.
//
// GET /api/submissions/forms/{projectId}/{formId}
//
// projectID: The project's id.
// formID: The form's id on its page.
func (r *SubmissionsService) GetFormSettings(ctx context.Context, projectID string, formID string, opts ...CallOption) (*SubmissionFormSettings, error) {
	var out *SubmissionFormSettings
	err := r.c.call(ctx, "GET", "/api/submissions/forms/"+pathSeg(projectID)+"/"+pathSeg(formID), nil, nil, false, "data.settings", &out, opts)
	return out, err
}

// List — Lists the form submissions the caller may see, newest first, a page at a time.
//
// GET /api/submissions
func (r *SubmissionsService) List(ctx context.Context, query *SubmissionsListQuery, opts ...CallOption) ([]Submission, error) {
	var out []Submission
	err := r.c.call(ctx, "GET", "/api/submissions", query, nil, false, "data.submissions", &out, opts)
	return out, err
}

// ListCannedResponses — Lists the saved replies of the caller's organizations, and of a project.
//
// GET /api/submissions/canned-responses
func (r *SubmissionsService) ListCannedResponses(ctx context.Context, query *SubmissionsListCannedResponsesQuery, opts ...CallOption) ([]SubmissionCannedResponse, error) {
	var out []SubmissionCannedResponse
	err := r.c.call(ctx, "GET", "/api/submissions/canned-responses", query, nil, false, "data.templates", &out, opts)
	return out, err
}

// ListForms — Lists the forms of a project that have submissions, with counts and settings.
//
// GET /api/submissions/forms
func (r *SubmissionsService) ListForms(ctx context.Context, query *SubmissionsListFormsQuery, opts ...CallOption) ([]SubmissionsListFormsResponseDataFormsItem, error) {
	var out []SubmissionsListFormsResponseDataFormsItem
	err := r.c.call(ctx, "GET", "/api/submissions/forms", query, nil, false, "data.forms", &out, opts)
	return out, err
}

// Reply — Replies to a submission, or adds an internal note to it.
//
// A reply (`outbound`, the default) marks the submission replied; with `sendEmail` it is also emailed to the submitter. An `internal_note` changes nothing else. Rate-limited per submission.
//
// POST /api/submissions/{submissionId}/messages
//
// submissionID: The submission's id.
func (r *SubmissionsService) Reply(ctx context.Context, submissionID string, body *SubmissionsReplyBody, opts ...CallOption) (*SubmissionMessage, error) {
	var out *SubmissionMessage
	err := r.c.call(ctx, "POST", "/api/submissions/"+pathSeg(submissionID)+"/messages", nil, body, false, "data.message", &out, opts)
	return out, err
}

// SetFormVisibility — Chooses who sees a form's submissions from now on — the caller only, the workspace, or the project.
//
// Existing submissions keep theirs unless `applyToExisting` — and then only those the caller can see.
//
// PUT /api/submissions/forms/{projectId}/{formId}/visibility
//
// projectID: The project's id.
// formID: The form's id on its page.
func (r *SubmissionsService) SetFormVisibility(ctx context.Context, projectID string, formID string, body *SubmissionsSetFormVisibilityBody, opts ...CallOption) (*SubmissionsSetFormVisibilityResponseData, error) {
	var out *SubmissionsSetFormVisibilityResponseData
	err := r.c.call(ctx, "PUT", "/api/submissions/forms/"+pathSeg(projectID)+"/"+pathSeg(formID)+"/visibility", nil, body, false, "data", &out, opts)
	return out, err
}

// Thread — Lists the other submissions of the same conversation the caller may see.
//
// GET /api/submissions/{id}/thread
//
// id: The submission's id.
func (r *SubmissionsService) Thread(ctx context.Context, id string, opts ...CallOption) ([]SubmissionsThreadResponseDataRelatedItem, error) {
	var out []SubmissionsThreadResponseDataRelatedItem
	err := r.c.call(ctx, "GET", "/api/submissions/"+pathSeg(id)+"/thread", nil, nil, false, "data.related", &out, opts)
	return out, err
}

// Update — Changes a submission's status, priority, tags or assignee.
//
// PATCH /api/submissions/{id}
//
// id: The submission's id.
func (r *SubmissionsService) Update(ctx context.Context, id string, body *SubmissionsUpdateBody, opts ...CallOption) (*Submission, error) {
	var out *Submission
	err := r.c.call(ctx, "PATCH", "/api/submissions/"+pathSeg(id), nil, body, false, "data.submission", &out, opts)
	return out, err
}

// UpdateCannedResponse — Changes a saved reply.
//
// PATCH /api/submissions/canned-responses/{cannedResponseId}
//
// cannedResponseID: The saved reply's id.
func (r *SubmissionsService) UpdateCannedResponse(ctx context.Context, cannedResponseID string, body *SubmissionsUpdateCannedResponseBody, opts ...CallOption) (*SubmissionCannedResponse, error) {
	var out *SubmissionCannedResponse
	err := r.c.call(ctx, "PATCH", "/api/submissions/canned-responses/"+pathSeg(cannedResponseID), nil, body, false, "data.template", &out, opts)
	return out, err
}

// UpdateFormSettings — Changes a form's inbox settings; only the fields sent change.
//
// Needs `submissions.manage`. Notified and auto-assigned people must be able to see the form's submissions.
//
// PATCH /api/submissions/forms/{projectId}/{formId}
//
// projectID: The project's id.
// formID: The form's id on its page.
func (r *SubmissionsService) UpdateFormSettings(ctx context.Context, projectID string, formID string, body *SubmissionsUpdateFormSettingsBody, opts ...CallOption) (*SubmissionFormSettings, error) {
	var out *SubmissionFormSettings
	err := r.c.call(ctx, "PATCH", "/api/submissions/forms/"+pathSeg(projectID)+"/"+pathSeg(formID), nil, body, false, "data.settings", &out, opts)
	return out, err
}

// SubscriptionsService — Subscriptions: Cancelling an organization's paid plan with the platform, and taking a scheduled cancellation back.
type SubscriptionsService struct {
	c *Client
}

func newSubscriptionsService(c *Client) *SubscriptionsService {
	return &SubscriptionsService{c: c}
}

// Cancel — Cancels an organization's paid plan at the end of the paid period, or at once.
//
// Needs to be the organization's owner or hold `billing:manage` in it. By default the plan keeps running until the period ends; `immediate: true` ends it now. A scheduled downgrade is withdrawn. 404 `noActiveSubscription` without a paid plan.
//
// POST /api/subscriptions/cancel
func (r *SubscriptionsService) Cancel(ctx context.Context, body *SubscriptionsCancelBody, opts ...CallOption) (*SubscriptionsCancelResponse, error) {
	var out *SubscriptionsCancelResponse
	err := r.c.call(ctx, "POST", "/api/subscriptions/cancel", nil, body, false, "", &out, opts)
	return out, err
}

// Resume — Takes back a cancellation scheduled for the end of the paid period.
//
// Needs to be the organization's owner or hold `billing:manage` in it. Works only while the plan has not ended; 404 `subscriptionNotCanceling` when nothing is scheduled.
//
// DELETE /api/subscriptions/cancel
func (r *SubscriptionsService) Resume(ctx context.Context, query *SubscriptionsResumeQuery, opts ...CallOption) (*SubscriptionsResumeResponse, error) {
	var out *SubscriptionsResumeResponse
	err := r.c.call(ctx, "DELETE", "/api/subscriptions/cancel", query, nil, false, "", &out, opts)
	return out, err
}

// SupportService — Support
type SupportService struct {
	c *Client
}

func newSupportService(c *Client) *SupportService {
	return &SupportService{c: c}
}

// CreateTicket — Opens a support ticket with the platform's support team.
//
// The platform's admins are notified. The ticket belongs to the caller and the first organization they are a member of.
//
// POST /api/support/tickets
func (r *SupportService) CreateTicket(ctx context.Context, body *SupportCreateTicketBody, opts ...CallOption) (*SupportTicket, error) {
	var out *SupportTicket
	err := r.c.call(ctx, "POST", "/api/support/tickets", nil, body, false, "data", &out, opts)
	return out, err
}

// GetTicket — Returns one of the caller's support tickets with its messages, oldest first.
//
// Internal notes of the support team are left out. `ticketUser` and `assignedAdmin` are only returned to the platform's support staff.
//
// GET /api/support/tickets/{id}
//
// id: The ticket's id.
func (r *SupportService) GetTicket(ctx context.Context, id string, opts ...CallOption) (*SupportGetTicketResponseData, error) {
	var out *SupportGetTicketResponseData
	err := r.c.call(ctx, "GET", "/api/support/tickets/"+pathSeg(id), nil, nil, false, "data", &out, opts)
	return out, err
}

// ListMessages — Lists the messages of one of the caller's support tickets, oldest first.
//
// Internal notes of the support team are left out.
//
// GET /api/support/tickets/{ticketId}/messages
//
// ticketID: The ticket's id.
func (r *SupportService) ListMessages(ctx context.Context, ticketID string, opts ...CallOption) ([]SupportMessage, error) {
	var out []SupportMessage
	err := r.c.call(ctx, "GET", "/api/support/tickets/"+pathSeg(ticketID)+"/messages", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListTickets — Lists the support tickets the caller opened, newest first.
//
// GET /api/support/tickets
func (r *SupportService) ListTickets(ctx context.Context, query *SupportListTicketsQuery, opts ...CallOption) ([]SupportTicket, error) {
	var out []SupportTicket
	err := r.c.call(ctx, "GET", "/api/support/tickets", query, nil, false, "data", &out, opts)
	return out, err
}

// PostMessage — Adds the caller's reply to one of their support tickets.
//
// A ticket waiting on the caller (or resolved) goes back to waiting on the support team, and the assigned admin is notified.
//
// POST /api/support/tickets/{ticketId}/messages
//
// ticketID: The ticket's id.
func (r *SupportService) PostMessage(ctx context.Context, ticketID string, body *SupportPostMessageBody, opts ...CallOption) (*SupportMessage, error) {
	var out *SupportMessage
	err := r.c.call(ctx, "POST", "/api/support/tickets/"+pathSeg(ticketID)+"/messages", nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateTicket — Closes one of the caller's support tickets.
//
// The ticket's author can only close it (`status: "closed"`); every other change belongs to the platform's support staff.
//
// PATCH /api/support/tickets/{id}
//
// id: The ticket's id.
func (r *SupportService) UpdateTicket(ctx context.Context, id string, body *SupportUpdateTicketBody, opts ...CallOption) (*SupportTicket, error) {
	var out *SupportTicket
	err := r.c.call(ctx, "PATCH", "/api/support/tickets/"+pathSeg(id), nil, body, false, "data", &out, opts)
	return out, err
}

// Upload — Uploads a screenshot to attach to a support ticket.
//
// PNG, JPEG, GIF or WebP, at most 10 MB. Pass the returned `url` as `screenshotUrl` when opening the ticket.
//
// POST /api/support/upload
func (r *SupportService) Upload(ctx context.Context, body *SupportUploadBody, opts ...CallOption) (string, error) {
	var out string
	err := r.c.call(ctx, "POST", "/api/support/upload", nil, body, true, "url", &out, opts)
	return out, err
}

// TasksService — Tasks
type TasksService struct {
	c *Client

	Automation *TasksAutomationService
	Checklist  *TasksChecklistService
	Conditions *TasksConditionsService
	DealRoom   *TasksDealRoomService
	Delegation *TasksDelegationService
	Details    *TasksDetailsService
	Feed       *TasksFeedService
	Focus      *TasksFocusService
	Nudges     *TasksNudgesService
	PlanDay    *TasksPlanDayService
	QuickAdd   *TasksQuickAddService
	Recurrence *TasksRecurrenceService
	Settings   *TasksSettingsService
	SLA        *TasksSlaService
	Views      *TasksViewsService
}

func newTasksService(c *Client) *TasksService {
	return &TasksService{c: c, Automation: newTasksAutomationService(c), Checklist: newTasksChecklistService(c), Conditions: newTasksConditionsService(c), DealRoom: newTasksDealRoomService(c), Delegation: newTasksDelegationService(c), Details: newTasksDetailsService(c), Feed: newTasksFeedService(c), Focus: newTasksFocusService(c), Nudges: newTasksNudgesService(c), PlanDay: newTasksPlanDayService(c), QuickAdd: newTasksQuickAddService(c), Recurrence: newTasksRecurrenceService(c), Settings: newTasksSettingsService(c), SLA: newTasksSlaService(c), Views: newTasksViewsService(c)}
}

// Bulk — Completes, reschedules, reprioritises or reassigns several tasks at once; each is checked on its own.
//
// POST /api/tasks/bulk — Token permission: write:tasks.
func (r *TasksService) Bulk(ctx context.Context, body *TaskBulkInput, opts ...CallOption) (*TasksBulkResponse, error) {
	var out *TasksBulkResponse
	err := r.c.call(ctx, "POST", "/api/tasks/bulk", nil, body, false, "", &out, opts)
	return out, err
}

// Complete — Ticks an item off, or opens it again with `done: false`, the way its own hub does it.
//
// For a CRM task: `{ source: "crm", databaseId, recordId }`; a deal-room step adds `stepId`; a calendar follow-up names its `eventId`.
//
// POST /api/tasks/complete — Token permission: write:tasks.
//
// body: one of *TasksCompleteBody, *TasksCompleteBody2, *TasksCompleteBody3.
func (r *TasksService) Complete(ctx context.Context, body any, opts ...CallOption) (*TasksCompleteResponse, error) {
	var out *TasksCompleteResponse
	err := r.c.call(ctx, "POST", "/api/tasks/complete", nil, body, false, "", &out, opts)
	return out, err
}

// Create — Quick add: creates a CRM task in the workspace's Tasks space, assigned to the caller.
//
// Needs edit level on the Tasks space. With several Tasks spaces and no `spaceId`: 409 `tasksSpaceRequired` with `spaces`; with none: 409 `tasksSetupRequired` or `tasksNoSpace`.
//
// POST /api/tasks — Token permission: write:tasks.
func (r *TasksService) Create(ctx context.Context, body *TaskCreateInput, opts ...CallOption) (*TaskItem, error) {
	var out *TaskItem
	err := r.c.call(ctx, "POST", "/api/tasks", nil, body, false, "task", &out, opts)
	return out, err
}

// Estimate — Sets how long a CRM task takes; null removes it (Plan my day then counts 30 minutes).
//
// Needs edit access to the task.
//
// PUT /api/tasks/estimate — Token permission: write:tasks.
func (r *TasksService) Estimate(ctx context.Context, body *TasksEstimateBody, opts ...CallOption) (float64, error) {
	var out float64
	err := r.c.call(ctx, "PUT", "/api/tasks/estimate", nil, body, false, "minutes", &out, opts)
	return out, err
}

// FromSource — What "Make task" starts with for a mail, message, comment, activity, booking or submission.
//
// GET /api/tasks/from-source — Token permission: read:tasks.
func (r *TasksService) FromSource(ctx context.Context, query *TasksFromSourceQuery, opts ...CallOption) (*TasksFromSourceResponse, error) {
	var out *TasksFromSourceResponse
	err := r.c.call(ctx, "GET", "/api/tasks/from-source", query, nil, false, "", &out, opts)
	return out, err
}

// List — What the caller has to do in a workspace, as one list: their CRM tasks and what other hubs wait on them for.
//
// Assigned to them (`mine`, the default) or created by them (`created`). Besides CRM tasks: approvals to decide, mail threads and submissions assigned to them, questions of their agent runs. Every source applies its own rules; `failedSources` names the ones that could not be read.
//
// GET /api/tasks — Token permission: read:tasks.
func (r *TasksService) List(ctx context.Context, query *TasksListQuery, opts ...CallOption) (*TasksListResponse, error) {
	var out *TasksListResponse
	err := r.c.call(ctx, "GET", "/api/tasks", query, nil, false, "", &out, opts)
	return out, err
}

// NextUp — "Next up": the open items in one order, each with why it is there.
//
// GET /api/tasks/next-up — Token permission: read:tasks.
func (r *TasksService) NextUp(ctx context.Context, query *TasksNextUpQuery, opts ...CallOption) (*TasksNextUpResponse, error) {
	var out *TasksNextUpResponse
	err := r.c.call(ctx, "GET", "/api/tasks/next-up", query, nil, false, "", &out, opts)
	return out, err
}

// Pin — Keeps an item at the top of "Next up", or lets it go again.
//
// POST /api/tasks/pin — Token permission: write:tasks.
func (r *TasksService) Pin(ctx context.Context, body *TasksPinBody, opts ...CallOption) (bool, error) {
	var out bool
	err := r.c.call(ctx, "POST", "/api/tasks/pin", nil, body, false, "pinned", &out, opts)
	return out, err
}

// Snooze — Holds the caller's notices about a CRM task until a time; then one reminder comes.
//
// Only the caller's own notices change. `until` within the next 90 days (400 `tasksSnoozeInvalid`); null ends the snooze.
//
// POST /api/tasks/snooze — Token permission: write:tasks.
func (r *TasksService) Snooze(ctx context.Context, body *TasksSnoozeBody, opts ...CallOption) (*TasksSnoozeResponse, error) {
	var out *TasksSnoozeResponse
	err := r.c.call(ctx, "POST", "/api/tasks/snooze", nil, body, false, "", &out, opts)
	return out, err
}

// Update — Changes a CRM task's due date, status, priority or assignee.
//
// POST /api/tasks/update — Token permission: write:tasks.
func (r *TasksService) Update(ctx context.Context, body *TaskUpdateInput, opts ...CallOption) (*TasksUpdateResponse, error) {
	var out *TasksUpdateResponse
	err := r.c.call(ctx, "POST", "/api/tasks/update", nil, body, false, "", &out, opts)
	return out, err
}

// Workload — The team view: the workspace's members and their open, overdue and coming tasks.
//
// GET /api/tasks/workload — Token permission: read:tasks.
func (r *TasksService) Workload(ctx context.Context, query *TasksWorkloadQuery, opts ...CallOption) (*TasksWorkloadResponse, error) {
	var out *TasksWorkloadResponse
	err := r.c.call(ctx, "GET", "/api/tasks/workload", query, nil, false, "", &out, opts)
	return out, err
}

// TasksAutomationService — tasks.automation
type TasksAutomationService struct {
	c *Client
}

func newTasksAutomationService(c *Client) *TasksAutomationService {
	return &TasksAutomationService{c: c}
}

// Get — The caller's meeting guardian and waiting-for nudges, and the Tasks spaces they may add to.
//
// GET /api/tasks/automation-settings — Token permission: read:tasks.
func (r *TasksAutomationService) Get(ctx context.Context, query *TasksAutomationGetQuery, opts ...CallOption) (*TasksAutomationGetResponse, error) {
	var out *TasksAutomationGetResponse
	err := r.c.call(ctx, "GET", "/api/tasks/automation-settings", query, nil, false, "", &out, opts)
	return out, err
}

// Update — Changes the caller's meeting guardian and waiting-for nudges.
//
// The guardian needs a Tasks space the caller may add to (400 `tasksGuardianSpaceRequired`).
//
// PUT /api/tasks/automation-settings — Token permission: write:tasks.
func (r *TasksAutomationService) Update(ctx context.Context, body *TasksAutomationUpdateBody, opts ...CallOption) (*TasksAutomationUpdateResponse, error) {
	var out *TasksAutomationUpdateResponse
	err := r.c.call(ctx, "PUT", "/api/tasks/automation-settings", nil, body, false, "", &out, opts)
	return out, err
}

// TasksChecklistService — tasks.checklist
type TasksChecklistService struct {
	c *Client
}

func newTasksChecklistService(c *Client) *TasksChecklistService {
	return &TasksChecklistService{c: c}
}

// Convert — Turns a step of a CRM task's checklist into a task of its own, linked from the step.
//
// Needs edit access to the task; counts against the plan's record limit. The step's text becomes the title, its person and day the assignee and due date. A step that names a task already answers that task (`created: false`). Ticking the new task done ticks the step.
//
// POST /api/tasks/checklist/convert — Token permission: write:tasks.
func (r *TasksChecklistService) Convert(ctx context.Context, body *TasksChecklistConvertBody, opts ...CallOption) (*TasksChecklistConvertResponse, error) {
	var out *TasksChecklistConvertResponse
	err := r.c.call(ctx, "POST", "/api/tasks/checklist/convert", nil, body, false, "", &out, opts)
	return out, err
}

// TasksConditionsService — tasks.conditions
type TasksConditionsService struct {
	c *Client
}

func newTasksConditionsService(c *Client) *TasksConditionsService {
	return &TasksConditionsService{c: c}
}

// Add — Adds a "done when" condition to a CRM task; a thing already done completes the task at once.
//
// Needs edit access to the task, and the caller must be able to open the watched thing.
//
// POST /api/tasks/conditions — Token permission: write:tasks.
func (r *TasksConditionsService) Add(ctx context.Context, body *TaskConditionInput, opts ...CallOption) (bool, error) {
	var out bool
	err := r.c.call(ctx, "POST", "/api/tasks/conditions", nil, body, false, "completed", &out, opts)
	return out, err
}

// List — A CRM task's "done when" conditions; what each watches is named only for a reader who may open it.
//
// GET /api/tasks/conditions — Token permission: read:tasks.
func (r *TasksConditionsService) List(ctx context.Context, query *TasksConditionsListQuery, opts ...CallOption) (*TasksConditionsListResponse, error) {
	var out *TasksConditionsListResponse
	err := r.c.call(ctx, "GET", "/api/tasks/conditions", query, nil, false, "", &out, opts)
	return out, err
}

// Options — What a "done when" condition of this kind may watch — only things the caller may open.
//
// Needs edit access to the task.
//
// GET /api/tasks/conditions/options — Token permission: read:tasks.
func (r *TasksConditionsService) Options(ctx context.Context, query *TasksConditionsOptionsQuery, opts ...CallOption) (*TasksConditionsOptionsResponse, error) {
	var out *TasksConditionsOptionsResponse
	err := r.c.call(ctx, "GET", "/api/tasks/conditions/options", query, nil, false, "", &out, opts)
	return out, err
}

// Remove — Removes a "done when" condition from a task.
//
// DELETE /api/tasks/conditions — Token permission: write:tasks.
func (r *TasksConditionsService) Remove(ctx context.Context, query *TasksConditionsRemoveQuery, opts ...CallOption) (*TasksConditionsRemoveResponse, error) {
	var out *TasksConditionsRemoveResponse
	err := r.c.call(ctx, "DELETE", "/api/tasks/conditions", query, nil, false, "", &out, opts)
	return out, err
}

// TasksDealRoomService — tasks.dealRoom
type TasksDealRoomService struct {
	c *Client
}

func newTasksDealRoomService(c *Client) *TasksDealRoomService {
	return &TasksDealRoomService{c: c}
}

// Get — The deal-room step a CRM task is shown as to the customer.
//
// The deal is named only for a reader who may open it.
//
// GET /api/tasks/deal-room — Token permission: read:tasks.
func (r *TasksDealRoomService) Get(ctx context.Context, query *TasksDealRoomGetQuery, opts ...CallOption) (*TasksDealRoomGetResponse, error) {
	var out *TasksDealRoomGetResponse
	err := r.c.call(ctx, "GET", "/api/tasks/deal-room", query, nil, false, "", &out, opts)
	return out, err
}

// Link — Shows a CRM task to the customer as a deal-room step: an existing step, or a new one titled like the task.
//
// Needs edit access to the task and the deal, and the deal's room switched on.
//
// POST /api/tasks/deal-room — Token permission: write:tasks.
func (r *TasksDealRoomService) Link(ctx context.Context, body *TasksDealRoomLinkBody, opts ...CallOption) (*TasksDealRoomLinkResponse, error) {
	var out *TasksDealRoomLinkResponse
	err := r.c.call(ctx, "POST", "/api/tasks/deal-room", nil, body, false, "", &out, opts)
	return out, err
}

// Options — The open deals a CRM task could be linked to — their room on, changeable by the caller — each with its customer steps.
//
// GET /api/tasks/deal-room/options — Token permission: read:tasks.
func (r *TasksDealRoomService) Options(ctx context.Context, query *TasksDealRoomOptionsQuery, opts ...CallOption) (*TasksDealRoomOptionsResponse, error) {
	var out *TasksDealRoomOptionsResponse
	err := r.c.call(ctx, "GET", "/api/tasks/deal-room/options", query, nil, false, "", &out, opts)
	return out, err
}

// Unlink — Stops showing a task in its deal's deal room.
//
// DELETE /api/tasks/deal-room — Token permission: write:tasks.
func (r *TasksDealRoomService) Unlink(ctx context.Context, query *TasksDealRoomUnlinkQuery, opts ...CallOption) (*TasksDealRoomUnlinkResponse, error) {
	var out *TasksDealRoomUnlinkResponse
	err := r.c.call(ctx, "DELETE", "/api/tasks/deal-room", query, nil, false, "", &out, opts)
	return out, err
}

// TasksDelegationService — tasks.delegation
type TasksDelegationService struct {
	c *Client
}

func newTasksDelegationService(c *Client) *TasksDelegationService {
	return &TasksDelegationService{c: c}
}

// Cancel — Calls off the agent working on a task.
//
// DELETE /api/tasks/delegate — Token permission: write:tasks.
func (r *TasksDelegationService) Cancel(ctx context.Context, query *TasksDelegationCancelQuery, opts ...CallOption) (*TasksDelegationCancelResponse, error) {
	var out *TasksDelegationCancelResponse
	err := r.c.call(ctx, "DELETE", "/api/tasks/delegate", query, nil, false, "", &out, opts)
	return out, err
}

// Get — The caller's latest delegation of a CRM task to an agent, and the agents they may run.
//
// GET /api/tasks/delegate — Token permission: read:tasks.
func (r *TasksDelegationService) Get(ctx context.Context, query *TasksDelegationGetQuery, opts ...CallOption) (*TasksDelegationGetResponse, error) {
	var out *TasksDelegationGetResponse
	err := r.c.call(ctx, "GET", "/api/tasks/delegate", query, nil, false, "", &out, opts)
	return out, err
}

// Start — Starts an agent on a CRM task, or answers the finished agent's "mark done?".
//
// Starting (`agentId`) needs edit access to the task, the Agent Hub, `agent:execute` and an active agent the caller may see; the run acts as the caller, and steps that need confirming become questions in their Tasks list. Answering (`answer`): `done` ticks the task as the caller — never automatic.
//
// POST /api/tasks/delegate — Token permission: write:tasks.
//
// body: one of *TasksDelegationStartBody, *TasksDelegationStartBody2.
func (r *TasksDelegationService) Start(ctx context.Context, body any, opts ...CallOption) (*TasksDelegationStartResponse, error) {
	var out *TasksDelegationStartResponse
	err := r.c.call(ctx, "POST", "/api/tasks/delegate", nil, body, false, "", &out, opts)
	return out, err
}

// TasksDetailsService — tasks.details
type TasksDetailsService struct {
	c *Client
}

func newTasksDetailsService(c *Client) *TasksDetailsService {
	return &TasksDetailsService{c: c}
}

// Get — What a CRM task carries beside its cells: how it repeats, its own reminders, the caller's snooze, and whether the caller may change them.
//
// GET /api/tasks/details — Token permission: read:tasks.
func (r *TasksDetailsService) Get(ctx context.Context, query *TasksDetailsGetQuery, opts ...CallOption) (*TasksDetailsGetResponse, error) {
	var out *TasksDetailsGetResponse
	err := r.c.call(ctx, "GET", "/api/tasks/details", query, nil, false, "", &out, opts)
	return out, err
}

// Update — Sets how a CRM task repeats, or replaces its own reminders.
//
// Needs edit access to the task. `recurrence: null` stops it repeating. 400 `tasksRecurrenceInvalid` / `tasksRemindersInvalid`.
//
// PUT /api/tasks/details — Token permission: write:tasks.
func (r *TasksDetailsService) Update(ctx context.Context, body *TasksDetailsUpdateBody, opts ...CallOption) (*TasksDetailsUpdateResponse, error) {
	var out *TasksDetailsUpdateResponse
	err := r.c.call(ctx, "PUT", "/api/tasks/details", nil, body, false, "", &out, opts)
	return out, err
}

// TasksFeedService — tasks.feed
type TasksFeedService struct {
	c *Client
}

func newTasksFeedService(c *Client) *TasksFeedService {
	return &TasksFeedService{c: c}
}

// Create — Creates a calendar feed of the caller's tasks in a workspace; the previous feed stops working.
//
// The feed's URL is answered once: only a hash is kept, so a lost URL is replaced by making a new one. The URL is a secret — whoever has it reads the tasks.
//
// POST /api/tasks/settings/feed — Token permission: write:tasks.
func (r *TasksFeedService) Create(ctx context.Context, body *TasksFeedCreateBody, opts ...CallOption) (*TasksFeedCreateResponse, error) {
	var out *TasksFeedCreateResponse
	err := r.c.call(ctx, "POST", "/api/tasks/settings/feed", nil, body, false, "", &out, opts)
	return out, err
}

// Revoke — Ends the caller's calendar feed; every calendar subscribed to it gets 404 from then on.
//
// DELETE /api/tasks/settings/feed — Token permission: write:tasks.
func (r *TasksFeedService) Revoke(ctx context.Context, opts ...CallOption) (*TaskUserSettings, error) {
	var out *TaskUserSettings
	err := r.c.call(ctx, "DELETE", "/api/tasks/settings/feed", nil, nil, false, "settings", &out, opts)
	return out, err
}

// TasksFocusService — tasks.focus
type TasksFocusService struct {
	c *Client
}

func newTasksFocusService(c *Client) *TasksFocusService {
	return &TasksFocusService{c: c}
}

// End — Ends the caller's focus session now; the profile before comes back.
//
// DELETE /api/tasks/focus — Token permission: write:tasks.
func (r *TasksFocusService) End(ctx context.Context, opts ...CallOption) (*TasksFocusEndResponse, error) {
	var out *TasksFocusEndResponse
	err := r.c.call(ctx, "DELETE", "/api/tasks/focus", nil, nil, false, "", &out, opts)
	return out, err
}

// Get — The caller's running focus session and their focus profiles.
//
// GET /api/tasks/focus — Token permission: read:tasks.
func (r *TasksFocusService) Get(ctx context.Context, opts ...CallOption) (*TasksFocusGetResponse, error) {
	var out *TasksFocusGetResponse
	err := r.c.call(ctx, "GET", "/api/tasks/focus", nil, nil, false, "", &out, opts)
	return out, err
}

// Start — Focuses on a task for 5–240 minutes, optionally switching to one of the caller's focus profiles meanwhile.
//
// POST /api/tasks/focus — Token permission: write:tasks.
func (r *TasksFocusService) Start(ctx context.Context, body *TasksFocusStartBody, opts ...CallOption) (*TasksFocusStartResponse, error) {
	var out *TasksFocusStartResponse
	err := r.c.call(ctx, "POST", "/api/tasks/focus", nil, body, false, "", &out, opts)
	return out, err
}

// TasksNudgesService — tasks.nudges
type TasksNudgesService struct {
	c *Client
}

func newTasksNudgesService(c *Client) *TasksNudgesService {
	return &TasksNudgesService{c: c}
}

// Remind — Nudges the people a task or mail waits on: a notification to a task's other assignees, or a reply draft in a mail thread.
//
// A task: once a day (409 `tasksAlreadyReminded`), by who gave it or may change it. A mail: a draft to the people the last message went to — never sent by itself.
//
// POST /api/tasks/nudges/remind — Token permission: write:tasks.
//
// body: one of *TasksNudgesRemindBody, *TasksNudgesRemindBody2.
func (r *TasksNudgesService) Remind(ctx context.Context, body any, opts ...CallOption) (*TasksNudgesRemindResponse, error) {
	var out *TasksNudgesRemindResponse
	err := r.c.call(ctx, "POST", "/api/tasks/nudges/remind", nil, body, false, "", &out, opts)
	return out, err
}

// TasksPlanDayService — tasks.planDay
type TasksPlanDayService struct {
	c *Client
}

func newTasksPlanDayService(c *Client) *TasksPlanDayService {
	return &TasksPlanDayService{c: c}
}

// Blocks — The caller's booked blocks of a task that are still ahead.
//
// GET /api/tasks/plan-day/blocks — Token permission: read:tasks.
func (r *TasksPlanDayService) Blocks(ctx context.Context, query *TasksPlanDayBlocksQuery, opts ...CallOption) ([]TasksPlanDayBlocksResponseBlocksItem, error) {
	var out []TasksPlanDayBlocksResponseBlocksItem
	err := r.c.call(ctx, "GET", "/api/tasks/plan-day/blocks", query, nil, false, "blocks", &out, opts)
	return out, err
}

// Book — Books the planned blocks as tentative events in the caller's own default calendar.
//
// A block for a task that already has one moves that event. Only tasks of the caller's own list. Each block answers on its own (`failed`). Needs `calendar:create`; 409 `calendarNoCalendar` without a calendar.
//
// POST /api/tasks/plan-day/book — Token permission: write:tasks.
func (r *TasksPlanDayService) Book(ctx context.Context, body *TasksPlanDayBookBody, opts ...CallOption) (*TasksPlanDayBookResponse, error) {
	var out *TasksPlanDayBookResponse
	err := r.c.call(ctx, "POST", "/api/tasks/plan-day/book", nil, body, false, "", &out, opts)
	return out, err
}

// FreeBlocks — Frees a task's planned time: its block events are deleted from the caller's own calendar.
//
// DELETE /api/tasks/plan-day/blocks — Token permission: write:tasks.
func (r *TasksPlanDayService) FreeBlocks(ctx context.Context, query *TasksPlanDayFreeBlocksQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "DELETE", "/api/tasks/plan-day/blocks", query, nil, false, "freed", &out, opts)
	return out, err
}

// Get — "Plan my day": the caller's working window, what is in their own calendars, and a proposal placing their open tasks in the free time.
//
// Nothing is written; booking is `tasks.planDay.book`. Needs the calendar on the plan and `calendar:read`; only the caller's own calendars are read.
//
// GET /api/tasks/plan-day — Token permission: read:tasks.
func (r *TasksPlanDayService) Get(ctx context.Context, query *TasksPlanDayGetQuery, opts ...CallOption) (*TasksPlanDayGetResponse, error) {
	var out *TasksPlanDayGetResponse
	err := r.c.call(ctx, "GET", "/api/tasks/plan-day", query, nil, false, "", &out, opts)
	return out, err
}

// TasksQuickAddService — tasks.quickAdd
type TasksQuickAddService struct {
	c *Client
}

func newTasksQuickAddService(c *Client) *TasksQuickAddService {
	return &TasksQuickAddService{c: c}
}

// AI — Asks the platform AI to read a quick-add line for the parts the plain reader missed: the due date, the priority, the repeat.
//
// A suggestion only; nothing is written (the task is made by `tasks.create`). The person's AI opt-out (403 `ai_opted_out`) and the organisation's AI allowance (429 `aiQuotaExceeded`) apply.
//
// POST /api/tasks/quick-add/ai — Token permission: write:tasks.
func (r *TasksQuickAddService) AI(ctx context.Context, body *TasksQuickAddAiBody, opts ...CallOption) (*TasksQuickAddAiResponse, error) {
	var out *TasksQuickAddAiResponse
	err := r.c.call(ctx, "POST", "/api/tasks/quick-add/ai", nil, body, false, "", &out, opts)
	return out, err
}

// Records — The CRM records a quick-add mention ("+Acme", "about deal Acme") may mean, by name or e-mail; nothing is linked.
//
// GET /api/tasks/quick-add/records — Token permission: read:tasks.
func (r *TasksQuickAddService) Records(ctx context.Context, query *TasksQuickAddRecordsQuery, opts ...CallOption) (*TasksQuickAddRecordsResponse, error) {
	var out *TasksQuickAddRecordsResponse
	err := r.c.call(ctx, "GET", "/api/tasks/quick-add/records", query, nil, false, "", &out, opts)
	return out, err
}

// Space — What a Tasks space takes beyond a title: its label options and the kinds of CRM records it can link.
//
// GET /api/tasks/quick-add/space — Token permission: read:tasks.
func (r *TasksQuickAddService) Space(ctx context.Context, query *TasksQuickAddSpaceQuery, opts ...CallOption) (*TasksQuickAddSpaceResponse, error) {
	var out *TasksQuickAddSpaceResponse
	err := r.c.call(ctx, "GET", "/api/tasks/quick-add/space", query, nil, false, "", &out, opts)
	return out, err
}

// TasksRecurrenceService — tasks.recurrence
type TasksRecurrenceService struct {
	c *Client
}

func newTasksRecurrenceService(c *Client) *TasksRecurrenceService {
	return &TasksRecurrenceService{c: c}
}

// Skip — Skips this occurrence of a repeating CRM task: its due date moves to the rule's next date without it being done.
//
// Needs edit access to the task. 409 `tasksNotRecurring`, `tasksSeriesEnded`, `tasksNoDueColumn`.
//
// POST /api/tasks/recurrence/skip — Token permission: write:tasks.
func (r *TasksRecurrenceService) Skip(ctx context.Context, body *TasksRecurrenceSkipBody, opts ...CallOption) (*TasksRecurrenceSkipResponse, error) {
	var out *TasksRecurrenceSkipResponse
	err := r.c.call(ctx, "POST", "/api/tasks/recurrence/skip", nil, body, false, "", &out, opts)
	return out, err
}

// TasksSettingsService — tasks.settings
type TasksSettingsService struct {
	c *Client
}

func newTasksSettingsService(c *Client) *TasksSettingsService {
	return &TasksSettingsService{c: c}
}

// Get — The caller's Tasks settings: default reminders and the daily digest.
//
// GET /api/tasks/settings — Token permission: read:tasks.
func (r *TasksSettingsService) Get(ctx context.Context, opts ...CallOption) (*TasksSettingsGetResponse, error) {
	var out *TasksSettingsGetResponse
	err := r.c.call(ctx, "GET", "/api/tasks/settings", nil, nil, false, "", &out, opts)
	return out, err
}

// Update — Changes the caller's Tasks settings.
//
// PUT /api/tasks/settings — Token permission: write:tasks.
func (r *TasksSettingsService) Update(ctx context.Context, body *TaskSettingsPatch, opts ...CallOption) (*TaskUserSettings, error) {
	var out *TaskUserSettings
	err := r.c.call(ctx, "PUT", "/api/tasks/settings", nil, body, false, "settings", &out, opts)
	return out, err
}

// TasksSlaService — tasks.sla
type TasksSlaService struct {
	c *Client
}

func newTasksSlaService(c *Client) *TasksSlaService {
	return &TasksSlaService{c: c}
}

// Get — A Tasks space's service levels and escalation ladder, and whether the caller may change them.
//
// GET /api/tasks/sla — Token permission: read:tasks.
func (r *TasksSlaService) Get(ctx context.Context, query *TasksSlaGetQuery, opts ...CallOption) (*TasksSlaGetResponse, error) {
	var out *TasksSlaGetResponse
	err := r.c.call(ctx, "GET", "/api/tasks/sla", query, nil, false, "", &out, opts)
	return out, err
}

// Put — Sets a Tasks space's service levels and escalation ladder.
//
// Needs `crm:manage` and manage level on the space. 400 `tasksSlaInvalid` for a policy it does not take, `tasksSlaMemberUnknown` for a person outside the organisation.
//
// PUT /api/tasks/sla — Token permission: write:tasks.
func (r *TasksSlaService) Put(ctx context.Context, body *TasksSlaPutBody, opts ...CallOption) (*TasksSlaPutResponse, error) {
	var out *TasksSlaPutResponse
	err := r.c.call(ctx, "PUT", "/api/tasks/sla", nil, body, false, "", &out, opts)
	return out, err
}

// TasksViewsService — tasks.views
type TasksViewsService struct {
	c *Client
}

func newTasksViewsService(c *Client) *TasksViewsService {
	return &TasksViewsService{c: c}
}

// Create — Saves the hub's current state — scope, filters, grouping, sort, layout, search — as a named view.
//
// 409 `tasksViewLimit` at the maximum number of views.
//
// POST /api/tasks/views — Token permission: write:tasks.
func (r *TasksViewsService) Create(ctx context.Context, body *TasksViewsCreateBody, opts ...CallOption) (*TasksViewsCreateResponse, error) {
	var out *TasksViewsCreateResponse
	err := r.c.call(ctx, "POST", "/api/tasks/views", nil, body, false, "", &out, opts)
	return out, err
}

// Delete — Deletes a saved view; when it was the default, the hub's own default applies again.
//
// DELETE /api/tasks/views/{viewId} — Token permission: write:tasks.
//
// viewID: The saved view.
func (r *TasksViewsService) Delete(ctx context.Context, viewID string, opts ...CallOption) (*TasksViewsDeleteResponse, error) {
	var out *TasksViewsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/tasks/views/"+pathSeg(viewID), nil, nil, false, "", &out, opts)
	return out, err
}

// List — The caller's saved views of the Tasks hub, and the one it opens with.
//
// GET /api/tasks/views — Token permission: read:tasks.
func (r *TasksViewsService) List(ctx context.Context, opts ...CallOption) (*TasksViewsListResponse, error) {
	var out *TasksViewsListResponse
	err := r.c.call(ctx, "GET", "/api/tasks/views", nil, nil, false, "", &out, opts)
	return out, err
}

// SetDefault — Which view the hub opens with (null: the hub's own default).
//
// PUT /api/tasks/views — Token permission: write:tasks.
func (r *TasksViewsService) SetDefault(ctx context.Context, body *TasksViewsSetDefaultBody, opts ...CallOption) (*TasksViewsSetDefaultResponse, error) {
	var out *TasksViewsSetDefaultResponse
	err := r.c.call(ctx, "PUT", "/api/tasks/views", nil, body, false, "", &out, opts)
	return out, err
}

// Update — Renames a saved view, saves a new state into it, or makes it (or stops it being) the default.
//
// PATCH /api/tasks/views/{viewId} — Token permission: write:tasks.
//
// viewID: The saved view.
func (r *TasksViewsService) Update(ctx context.Context, viewID string, body *TasksViewsUpdateBody, opts ...CallOption) (*TasksViewsUpdateResponse, error) {
	var out *TasksViewsUpdateResponse
	err := r.c.call(ctx, "PATCH", "/api/tasks/views/"+pathSeg(viewID), nil, body, false, "", &out, opts)
	return out, err
}

// TaxRatesService — Tax rates
type TaxRatesService struct {
	c *Client
}

func newTaxRatesService(c *Client) *TaxRatesService {
	return &TaxRatesService{c: c}
}

// List — Lists the platform's tax rates, including those that only applied in the past.
//
// Answers without a token. Only enabled rates, sorted by name; `effectiveFrom` and `effectiveTo` say when each applied. Cached for five minutes.
//
// GET /api/tax-rates — No token needed.
func (r *TaxRatesService) List(ctx context.Context, opts ...CallOption) ([]TaxRate, error) {
	var out []TaxRate
	err := r.c.call(ctx, "GET", "/api/tax-rates", nil, nil, false, "", &out, opts)
	return out, err
}

// TeamsService — Teams
type TeamsService struct {
	c *Client
}

func newTeamsService(c *Client) *TeamsService {
	return &TeamsService{c: c}
}

// AddMember — Adds a person to a team, optionally with a role in it.
//
// Needs `team.update`.
//
// POST /api/teams/{teamId}/members — Token permission: write:teams.
//
// teamID: The team's id.
func (r *TeamsService) AddMember(ctx context.Context, teamID string, body *AddTeamMemberInput, opts ...CallOption) (*TeamMember, error) {
	var out *TeamMember
	err := r.c.call(ctx, "POST", "/api/teams/"+pathSeg(teamID)+"/members", nil, body, false, "data", &out, opts)
	return out, err
}

// Create — Creates a team in a workspace.
//
// Needs `team.create`.
//
// POST /api/teams — Token permission: write:teams.
func (r *TeamsService) Create(ctx context.Context, body *CreateTeamInput, opts ...CallOption) (*Team, error) {
	var out *Team
	err := r.c.call(ctx, "POST", "/api/teams", nil, body, false, "data", &out, opts)
	return out, err
}

// Delete — Deletes a team and its memberships.
//
// DELETE /api/teams/{teamId} — Token permission: write:teams.
//
// teamID: The team's id.
func (r *TeamsService) Delete(ctx context.Context, teamID string, opts ...CallOption) (*TeamsDeleteResponse, error) {
	var out *TeamsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/teams/"+pathSeg(teamID), nil, nil, false, "", &out, opts)
	return out, err
}

// Get — Returns a team with its members.
//
// GET /api/teams/{teamId} — Token permission: read:teams.
//
// teamID: The team's id.
func (r *TeamsService) Get(ctx context.Context, teamID string, opts ...CallOption) (*TeamsGetResponseData, error) {
	var out *TeamsGetResponseData
	err := r.c.call(ctx, "GET", "/api/teams/"+pathSeg(teamID), nil, nil, false, "data", &out, opts)
	return out, err
}

// List — Lists the teams the caller can see, optionally of one organization.
//
// GET /api/teams — Token permission: read:teams.
func (r *TeamsService) List(ctx context.Context, query *TeamsListQuery, opts ...CallOption) ([]Team, error) {
	var out []Team
	err := r.c.call(ctx, "GET", "/api/teams", query, nil, false, "data", &out, opts)
	return out, err
}

// ListMembers — Lists a team's members with their user details and roles in the team.
//
// GET /api/teams/{teamId}/members — Token permission: read:teams.
//
// teamID: The team's id.
func (r *TeamsService) ListMembers(ctx context.Context, teamID string, opts ...CallOption) ([]TeamMember, error) {
	var out []TeamMember
	err := r.c.call(ctx, "GET", "/api/teams/"+pathSeg(teamID)+"/members", nil, nil, false, "data", &out, opts)
	return out, err
}

// RemoveMember — Removes a person from a team.
//
// DELETE /api/teams/{teamId}/members/{memberId} — Token permission: write:teams.
//
// teamID: The team's id.
// memberID: The team membership's id.
func (r *TeamsService) RemoveMember(ctx context.Context, teamID string, memberID string, opts ...CallOption) (*TeamsRemoveMemberResponse, error) {
	var out *TeamsRemoveMemberResponse
	err := r.c.call(ctx, "DELETE", "/api/teams/"+pathSeg(teamID)+"/members/"+pathSeg(memberID), nil, nil, false, "", &out, opts)
	return out, err
}

// RemoveMemberRole — Takes a role in the team away from a member.
//
// DELETE /api/teams/{teamId}/members/{memberId}/assign-role — Token permission: write:teams.
//
// teamID: The team's id.
// memberID: The team membership's id.
func (r *TeamsService) RemoveMemberRole(ctx context.Context, teamID string, memberID string, query *TeamsRemoveMemberRoleQuery, opts ...CallOption) (*TeamsRemoveMemberRoleResponse, error) {
	var out *TeamsRemoveMemberRoleResponse
	err := r.c.call(ctx, "DELETE", "/api/teams/"+pathSeg(teamID)+"/members/"+pathSeg(memberID)+"/assign-role", query, nil, false, "", &out, opts)
	return out, err
}

// Update — Changes a team's name, description or settings.
//
// PUT /api/teams/{teamId} — Token permission: write:teams.
//
// teamID: The team's id.
func (r *TeamsService) Update(ctx context.Context, teamID string, body *UpdateTeamInput, opts ...CallOption) (*Team, error) {
	var out *Team
	err := r.c.call(ctx, "PUT", "/api/teams/"+pathSeg(teamID), nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateMember — Gives a team member a role in the team.
//
// Needs `team.manage_members`; only roles within the caller's own permissions.
//
// POST /api/teams/{teamId}/members/{memberId}/assign-role — Token permission: write:teams.
//
// teamID: The team's id.
// memberID: The team membership's id.
func (r *TeamsService) UpdateMember(ctx context.Context, teamID string, memberID string, body *TeamsUpdateMemberBody, opts ...CallOption) (*RoleAssignment, error) {
	var out *RoleAssignment
	err := r.c.call(ctx, "POST", "/api/teams/"+pathSeg(teamID)+"/members/"+pathSeg(memberID)+"/assign-role", nil, body, false, "data", &out, opts)
	return out, err
}

// TelephonyService — Telephony
type TelephonyService struct {
	c *Client

	Calls       *TelephonyCallsService
	Connections *TelephonyConnectionsService
	Lines       *TelephonyLinesService
}

func newTelephonyService(c *Client) *TelephonyService {
	return &TelephonyService{c: c, Calls: newTelephonyCallsService(c), Connections: newTelephonyConnectionsService(c), Lines: newTelephonyLinesService(c)}
}

// Dial — Click-to-call: the provider rings the caller's own phone on that line, then the number.
//
// Needs a line the caller may use (404 otherwise). Nothing is stored; the call arrives in the log through the provider's webhook. Never retried on a server error, so a phone is not rung twice. Rate-limited.
//
// POST /api/telephony/dial — Token permission: write:telephony.
func (r *TelephonyService) Dial(ctx context.Context, body *TelephonyDialBody, opts ...CallOption) (*TelephonyDialResponse, error) {
	var out *TelephonyDialResponse
	err := r.c.call(ctx, "POST", "/api/telephony/dial", nil, body, false, "", &out, opts)
	return out, err
}

// Overview — A workspace's telephony: the providers that can be connected, the connections and the lines.
//
// A telephony manager sees every line and each connection's webhook URL; anybody else only the lines they own or were added to.
//
// GET /api/telephony/overview — Token permission: read:telephony.
func (r *TelephonyService) Overview(ctx context.Context, query *TelephonyOverviewQuery, opts ...CallOption) (*TelephonyOverviewResponse, error) {
	var out *TelephonyOverviewResponse
	err := r.c.call(ctx, "GET", "/api/telephony/overview", query, nil, false, "", &out, opts)
	return out, err
}

// TelephonyCallsService — telephony.calls
type TelephonyCallsService struct {
	c *Client
}

func newTelephonyCallsService(c *Client) *TelephonyCallsService {
	return &TelephonyCallsService{c: c}
}

// Get — One call with every CRM record it matched, the call-back task it made, and the lines that can call back.
//
// A call on no line the caller may use answers 404, like a missing one.
//
// GET /api/telephony/calls/{callId} — Token permission: read:telephony.
//
// callID: The call.
func (r *TelephonyCallsService) Get(ctx context.Context, callID string, query *TelephonyCallsGetQuery, opts ...CallOption) (*PhoneCallDetail, error) {
	var out *PhoneCallDetail
	err := r.c.call(ctx, "GET", "/api/telephony/calls/"+pathSeg(callID), query, nil, false, "call", &out, opts)
	return out, err
}

// List — The call log of the lines the caller may use in a workspace, newest first, each with the CRM record its number matched.
//
// A telephony manager sees no more than anybody else: configuring a line is not reading it. Page on with `cursor` = the answer's `nextCursor`.
//
// GET /api/telephony/calls — Token permission: read:telephony.
func (r *TelephonyCallsService) List(ctx context.Context, query *TelephonyCallsListQuery, opts ...CallOption) ([]PhoneCallLogEntry, error) {
	var out []PhoneCallLogEntry
	err := r.c.call(ctx, "GET", "/api/telephony/calls", query, nil, false, "calls", &out, opts)
	return out, err
}

// TelephonyConnectionsService — telephony.connections
type TelephonyConnectionsService struct {
	c *Client
}

func newTelephonyConnectionsService(c *Client) *TelephonyConnectionsService {
	return &TelephonyConnectionsService{c: c}
}

// Create — Connects the customer's own phone provider account to a workspace; the credentials are tested first.
//
// Needs `telephony:manage` in the workspace. The credentials are stored only when they work, never shown again. The connection gets its own webhook URL; where the provider can list the account's lines, they are read at once.
//
// POST /api/telephony/connections — Token permission: write:telephony.
func (r *TelephonyConnectionsService) Create(ctx context.Context, body *PhoneConnectionInput, opts ...CallOption) (*PhoneConnection, error) {
	var out *PhoneConnection
	err := r.c.call(ctx, "POST", "/api/telephony/connections", nil, body, false, "connection", &out, opts)
	return out, err
}

// Delete — Deletes a connection with its lines and its call log.
//
// Needs `telephony:manage`.
//
// DELETE /api/telephony/connections/{connectionId} — Token permission: write:telephony.
//
// connectionID: The connection.
func (r *TelephonyConnectionsService) Delete(ctx context.Context, connectionID string, opts ...CallOption) (*TelephonyConnectionsDeleteResponse, error) {
	var out *TelephonyConnectionsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/telephony/connections/"+pathSeg(connectionID), nil, nil, false, "", &out, opts)
	return out, err
}

// SyncLines — Reads the account's lines from the provider: new ones are added, known ones refreshed, none deleted.
//
// Needs `telephony:manage`. Only where the provider can list lines (409 `telephonyLinesNotSupported` otherwise — add them by hand). New lines only while the plan has room; the rest are counted in `overLimit`. Owners, members and switches are never touched.
//
// POST /api/telephony/connections/{connectionId}/sync — Token permission: write:telephony.
//
// connectionID: The connection.
func (r *TelephonyConnectionsService) SyncLines(ctx context.Context, connectionID string, opts ...CallOption) (*TelephonyConnectionsSyncLinesResponse, error) {
	var out *TelephonyConnectionsSyncLinesResponse
	err := r.c.call(ctx, "POST", "/api/telephony/connections/"+pathSeg(connectionID)+"/sync", nil, nil, false, "", &out, opts)
	return out, err
}

// Test — Makes one authenticated read of the provider with the stored credentials.
//
// Needs `telephony:manage`. Works: the connection is active again; refused: its status becomes `error`. For a provider that issues the webhook secret itself (Aircall) it is read here once the webhook exists — `webhookSecret` says whether it was found.
//
// POST /api/telephony/connections/{connectionId}/test — Token permission: write:telephony.
//
// connectionID: The connection.
func (r *TelephonyConnectionsService) Test(ctx context.Context, connectionID string, opts ...CallOption) (*TelephonyConnectionsTestResponse, error) {
	var out *TelephonyConnectionsTestResponse
	err := r.c.call(ctx, "POST", "/api/telephony/connections/"+pathSeg(connectionID)+"/test", nil, nil, false, "", &out, opts)
	return out, err
}

// Update — Renames a connection, gives it new credentials, pauses or resumes it, or rotates its webhook URL.
//
// Needs `telephony:manage`. New credentials are tested before they replace the old ones. A paused connection acknowledges and drops the provider's deliveries; after a rotation the old URL answers 404.
//
// PATCH /api/telephony/connections/{connectionId} — Token permission: write:telephony.
//
// connectionID: The connection.
func (r *TelephonyConnectionsService) Update(ctx context.Context, connectionID string, body *PhoneConnectionPatch, opts ...CallOption) (*PhoneConnection, error) {
	var out *PhoneConnection
	err := r.c.call(ctx, "PATCH", "/api/telephony/connections/"+pathSeg(connectionID), nil, body, false, "connection", &out, opts)
	return out, err
}

// TelephonyLinesService — telephony.lines
type TelephonyLinesService struct {
	c *Client
}

func newTelephonyLinesService(c *Client) *TelephonyLinesService {
	return &TelephonyLinesService{c: c}
}

// Create — Adds a line by hand, for a provider or account whose lines cannot be listed.
//
// Needs `telephony:manage` in the connection's workspace and room under the plan's line limit. `providerLineId` is what the provider's webhook names the line by (sipgate: the user id `w0`; Placetel: the SIP user's sipuid); `providerDeviceId` is sipgate's device for click-to-call (`e0`).
//
// POST /api/telephony/lines — Token permission: write:telephony.
func (r *TelephonyLinesService) Create(ctx context.Context, body *PhoneLineInput, opts ...CallOption) (*TelephonyLinesCreateResponseLine, error) {
	var out *TelephonyLinesCreateResponseLine
	err := r.c.call(ctx, "POST", "/api/telephony/lines", nil, body, false, "line", &out, opts)
	return out, err
}

// Delete — Deletes a line.
//
// Needs `telephony:manage`.
//
// DELETE /api/telephony/lines/{lineId} — Token permission: write:telephony.
//
// lineID: The line.
func (r *TelephonyLinesService) Delete(ctx context.Context, lineID string, opts ...CallOption) (*TelephonyLinesDeleteResponse, error) {
	var out *TelephonyLinesDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/telephony/lines/"+pathSeg(lineID), nil, nil, false, "", &out, opts)
	return out, err
}

// Mine — The lines the caller may use right now, with whether click-to-call works on each.
//
// An empty list (never an error) when the caller has none.
//
// GET /api/telephony/lines/mine — Token permission: read:telephony.
func (r *TelephonyLinesService) Mine(ctx context.Context, query *TelephonyLinesMineQuery, opts ...CallOption) ([]PhoneUsableLine, error) {
	var out []PhoneUsableLine
	err := r.c.call(ctx, "GET", "/api/telephony/lines/mine", query, nil, false, "lines", &out, opts)
	return out, err
}

// Update — Changes a line: its label, number, owner and members, whether its calls show in the CRM, and the missed-call task.
//
// Label, number, device, owner and members need `telephony:manage` (owner and members must be in the workspace). `crmSync` is the line owner's switch (a telephony manager's for a line nobody owns). `missedCallTask` and `taskSpaceId`: the owner or a telephony manager.
//
// PATCH /api/telephony/lines/{lineId} — Token permission: write:telephony.
//
// lineID: The line.
func (r *TelephonyLinesService) Update(ctx context.Context, lineID string, body *PhoneLinePatch, opts ...CallOption) (*TelephonyLinesUpdateResponse, error) {
	var out *TelephonyLinesUpdateResponse
	err := r.c.call(ctx, "PATCH", "/api/telephony/lines/"+pathSeg(lineID), nil, body, false, "", &out, opts)
	return out, err
}

// TemplatesService — Templates
type TemplatesService struct {
	c *Client
}

func newTemplatesService(c *Client) *TemplatesService {
	return &TemplatesService{c: c}
}

// Create — Saves a component, page or workflow as a template.
//
// Needs `template.create`.
//
// POST /api/templates
func (r *TemplatesService) Create(ctx context.Context, body *CreateTemplateInput, opts ...CallOption) (*Template, error) {
	var out *Template
	err := r.c.call(ctx, "POST", "/api/templates", nil, body, false, "data", &out, opts)
	return out, err
}

// Delete — Deletes a template.
//
// DELETE /api/templates/{templateId}
//
// templateID: The template's id.
func (r *TemplatesService) Delete(ctx context.Context, templateID string, opts ...CallOption) (*TemplatesDeleteResponse, error) {
	var out *TemplatesDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/templates/"+pathSeg(templateID), nil, nil, false, "", &out, opts)
	return out, err
}

// Get — Returns a template with its content.
//
// GET /api/templates/{templateId}
//
// templateID: The template's id.
func (r *TemplatesService) Get(ctx context.Context, templateID string, opts ...CallOption) (*Template, error) {
	var out *Template
	err := r.c.call(ctx, "GET", "/api/templates/"+pathSeg(templateID), nil, nil, false, "data", &out, opts)
	return out, err
}

// List — Lists the templates the caller may use in an organization — their private ones, the organization's and its workspaces', and public ones.
//
// The list leaves out each template's content; `templates.get` has it.
//
// GET /api/templates
func (r *TemplatesService) List(ctx context.Context, query *TemplatesListQuery, opts ...CallOption) ([]Template, error) {
	var out []Template
	err := r.c.call(ctx, "GET", "/api/templates", query, nil, false, "data", &out, opts)
	return out, err
}

// Replace — Changes a template, exactly as `templates.update` does.
//
// PUT /api/templates/{templateId}
//
// templateID: The template's id.
func (r *TemplatesService) Replace(ctx context.Context, templateID string, body *UpdateTemplateInput, opts ...CallOption) (*Template, error) {
	var out *Template
	err := r.c.call(ctx, "PUT", "/api/templates/"+pathSeg(templateID), nil, body, false, "data", &out, opts)
	return out, err
}

// Update — Changes a template's details or content; only the fields sent change.
//
// Needs `template.update`; making it public needs `template.publish`.
//
// PATCH /api/templates/{templateId}
//
// templateID: The template's id.
func (r *TemplatesService) Update(ctx context.Context, templateID string, body *UpdateTemplateInput, opts ...CallOption) (*Template, error) {
	var out *Template
	err := r.c.call(ctx, "PATCH", "/api/templates/"+pathSeg(templateID), nil, body, false, "data", &out, opts)
	return out, err
}

// Use — Records that the template was used and returns it with its content.
//
// POST /api/templates/{templateId}
//
// templateID: The template's id.
func (r *TemplatesService) Use(ctx context.Context, templateID string, body *TemplatesUseBody, opts ...CallOption) (*Template, error) {
	var out *Template
	err := r.c.call(ctx, "POST", "/api/templates/"+pathSeg(templateID), nil, body, false, "data", &out, opts)
	return out, err
}

// TrashService — Trash: What was deleted across every hub, until it is deleted for good: restore it, or delete it now.
type TrashService struct {
	c *Client

	Settings *TrashSettingsService
}

func newTrashService(c *Client) *TrashService {
	return &TrashService{c: c, Settings: newTrashSettingsService(c)}
}

// Apply — Restores items from the trash, or deletes them for good.
//
// Answers per item. A restore that clashes with something live (a page at its address, a file of the same name) answers `conflict` with what clashes; send that item again with `resolution`: `keepBoth` restores it under the suggested free name, `replace` moves the live one to the trash instead. Items the caller may not touch answer `forbidden` and are left alone. Deleting for good cannot be undone.
//
// POST /api/trash
func (r *TrashService) Apply(ctx context.Context, body *TrashApplyBody, opts ...CallOption) ([]TrashActionResult, error) {
	var out []TrashActionResult
	err := r.c.call(ctx, "POST", "/api/trash", nil, body, false, "data.results", &out, opts)
	return out, err
}

// Empty — Empties the trash: deletes for good everything in it the caller may delete.
//
// Cannot be undone. With `kind`, only the items of that kind.
//
// DELETE /api/trash
func (r *TrashService) Empty(ctx context.Context, query *TrashEmptyQuery, opts ...CallOption) (*TrashEmptyResponseData, error) {
	var out *TrashEmptyResponseData
	err := r.c.call(ctx, "DELETE", "/api/trash", query, nil, false, "data", &out, opts)
	return out, err
}

// List — Lists what the caller may restore from the organisation's trash, newest first, each with the day it goes for good.
//
// Per kind, an item is listed to whoever could have deleted it in the first place.
//
// GET /api/trash
func (r *TrashService) List(ctx context.Context, query *TrashListQuery, opts ...CallOption) ([]TrashItem, error) {
	var out []TrashItem
	err := r.c.call(ctx, "GET", "/api/trash", query, nil, false, "data.items", &out, opts)
	return out, err
}

// TrashSettingsService — trash.settings
type TrashSettingsService struct {
	c *Client
}

func newTrashSettingsService(c *Client) *TrashSettingsService {
	return &TrashSettingsService{c: c}
}

// Get — How many days the organisation's trash keeps things, and the platform's limits.
//
// GET /api/trash/settings
func (r *TrashSettingsService) Get(ctx context.Context, query *TrashSettingsGetQuery, opts ...CallOption) (*TrashSettingsGetResponseData, error) {
	var out *TrashSettingsGetResponseData
	err := r.c.call(ctx, "GET", "/api/trash/settings", query, nil, false, "data", &out, opts)
	return out, err
}

// Update — Sets how many days the organisation's trash keeps things.
//
// Needs the `retention_policies:manage` permission. The days are held to the platform's limits; the answer says what was applied.
//
// PUT /api/trash/settings
func (r *TrashSettingsService) Update(ctx context.Context, body *TrashSettingsUpdateBody, opts ...CallOption) (*TrashSettingsUpdateResponseData, error) {
	var out *TrashSettingsUpdateResponseData
	err := r.c.call(ctx, "PUT", "/api/trash/settings", nil, body, false, "data", &out, opts)
	return out, err
}

// UserService — Current user
type UserService struct {
	c *Client
}

func newUserService(c *Client) *UserService {
	return &UserService{c: c}
}

// AddAIMemory — Adds a note to what the AI assistant remembers.
//
// Workspace and organization notes need the right to manage that level's AI persona. With memory switched off nothing is stored: `stored: false` says why.
//
// POST /api/user/ai-memory
func (r *UserService) AddAIMemory(ctx context.Context, body *UserAddAiMemoryBody, opts ...CallOption) (*UserAddAiMemoryResponse, error) {
	var out *UserAddAiMemoryResponse
	err := r.c.call(ctx, "POST", "/api/user/ai-memory", nil, body, false, "", &out, opts)
	return out, err
}

// ClaimHandle — Claims a public creator handle for the caller.
//
// PUT /api/user/handle
func (r *UserService) ClaimHandle(ctx context.Context, body *UserClaimHandleBody, opts ...CallOption) (*UserClaimHandleResponse, error) {
	var out *UserClaimHandleResponse
	err := r.c.call(ctx, "PUT", "/api/user/handle", nil, body, false, "", &out, opts)
	return out, err
}

// ClearPresence — Clears the presence the caller set, back to automatic.
//
// DELETE /api/user/presence
func (r *UserService) ClearPresence(ctx context.Context, opts ...CallOption) (*UserClearPresenceResponse, error) {
	var out *UserClearPresenceResponse
	err := r.c.call(ctx, "DELETE", "/api/user/presence", nil, nil, false, "", &out, opts)
	return out, err
}

// CountAIEvents — Counts the records of the caller's AI generations kept for audit.
//
// GET /api/user/ai-events
func (r *UserService) CountAIEvents(ctx context.Context, opts ...CallOption) (float64, error) {
	var out float64
	err := r.c.call(ctx, "GET", "/api/user/ai-events", nil, nil, false, "data.count", &out, opts)
	return out, err
}

// CreateSignature — Saves a signature image for signing documents.
//
// At most 8 per person, each a PNG data URL of at most 200 KB.
//
// POST /api/user/signatures
func (r *UserService) CreateSignature(ctx context.Context, body *UserCreateSignatureBody, opts ...CallOption) (*UserCreateSignatureResponseDataSignature, error) {
	var out *UserCreateSignatureResponseDataSignature
	err := r.c.call(ctx, "POST", "/api/user/signatures", nil, body, false, "data.signature", &out, opts)
	return out, err
}

// DeleteAIEvents — Erases the records of the caller's AI generations.
//
// Only the audit records: what was generated (files, drafts) stays.
//
// DELETE /api/user/ai-events
func (r *UserService) DeleteAIEvents(ctx context.Context, opts ...CallOption) (float64, error) {
	var out float64
	err := r.c.call(ctx, "DELETE", "/api/user/ai-events", nil, nil, false, "data.deleted", &out, opts)
	return out, err
}

// DeleteAIMemory — Makes the AI assistant forget a note.
//
// DELETE /api/user/ai-memory
func (r *UserService) DeleteAIMemory(ctx context.Context, query *UserDeleteAiMemoryQuery, opts ...CallOption) (*UserDeleteAiMemoryResponse, error) {
	var out *UserDeleteAiMemoryResponse
	err := r.c.call(ctx, "DELETE", "/api/user/ai-memory", query, nil, false, "", &out, opts)
	return out, err
}

// DeleteAvatar — Removes the caller's profile picture.
//
// DELETE /api/user/avatar
func (r *UserService) DeleteAvatar(ctx context.Context, opts ...CallOption) (*UserDeleteAvatarResponse, error) {
	var out *UserDeleteAvatarResponse
	err := r.c.call(ctx, "DELETE", "/api/user/avatar", nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteSignature — Deletes a saved signature; the newest other one becomes the default.
//
// DELETE /api/user/signatures/{signatureId}
//
// signatureID: The signature's id.
func (r *UserService) DeleteSignature(ctx context.Context, signatureID string, opts ...CallOption) (*UserDeleteSignatureResponse, error) {
	var out *UserDeleteSignatureResponse
	err := r.c.call(ctx, "DELETE", "/api/user/signatures/"+pathSeg(signatureID), nil, nil, false, "", &out, opts)
	return out, err
}

// ExportActivityLogs — Downloads the activity the caller may see as a JSON or CSV file.
//
// Answers the file itself (`application/json` or `text/csv`).
//
// GET /api/user/activity-logs/export
func (r *UserService) ExportActivityLogs(ctx context.Context, query *UserExportActivityLogsQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/user/activity-logs/export", query, nil, false, "", &out, opts)
	return out, err
}

// ForgetRecognizedDevice — Forgets a recognized device, so the next sign-in from it is reported as new.
//
// DELETE /api/user/recognized-devices
func (r *UserService) ForgetRecognizedDevice(ctx context.Context, query *UserForgetRecognizedDeviceQuery, opts ...CallOption) (*UserForgetRecognizedDeviceResponse, error) {
	var out *UserForgetRecognizedDeviceResponse
	err := r.c.call(ctx, "DELETE", "/api/user/recognized-devices", query, nil, false, "", &out, opts)
	return out, err
}

// GetAIMemory — Lists what the AI assistant remembers — about the caller, and at the workspaces and organizations they belong to.
//
// GET /api/user/ai-memory
func (r *UserService) GetAIMemory(ctx context.Context, opts ...CallOption) (*UserGetAiMemoryResponse, error) {
	var out *UserGetAiMemoryResponse
	err := r.c.call(ctx, "GET", "/api/user/ai-memory", nil, nil, false, "", &out, opts)
	return out, err
}

// GetAIMemorySwitch — Returns whether the AI assistant may remember things about the caller.
//
// GET /api/user/ai-memory/switch
func (r *UserService) GetAIMemorySwitch(ctx context.Context, opts ...CallOption) (*UserGetAiMemorySwitchResponse, error) {
	var out *UserGetAiMemorySwitchResponse
	err := r.c.call(ctx, "GET", "/api/user/ai-memory/switch", nil, nil, false, "", &out, opts)
	return out, err
}

// GetHandoff — Returns until when the caller is away, and their away message.
//
// GET /api/user/handoff
func (r *UserService) GetHandoff(ctx context.Context, opts ...CallOption) (*UserGetHandoffResponse, error) {
	var out *UserGetHandoffResponse
	err := r.c.call(ctx, "GET", "/api/user/handoff", nil, nil, false, "", &out, opts)
	return out, err
}

// GetPreferences — Returns the caller's preferences — look, language, regional formats, notifications, defaults.
//
// GET /api/user/preferences
func (r *UserService) GetPreferences(ctx context.Context, opts ...CallOption) (*UserPreferences, error) {
	var out *UserPreferences
	err := r.c.call(ctx, "GET", "/api/user/preferences", nil, nil, false, "", &out, opts)
	return out, err
}

// GetPresence — Returns the presence — online, away, do not disturb — of the caller and up to 200 other people.
//
// GET /api/user/presence
func (r *UserService) GetPresence(ctx context.Context, query *UserGetPresenceQuery, opts ...CallOption) ([]UserGetPresenceResponseDataItem, error) {
	var out []UserGetPresenceResponseDataItem
	err := r.c.call(ctx, "GET", "/api/user/presence", query, nil, false, "data", &out, opts)
	return out, err
}

// GetPresenceSettings — Returns what sets the caller's presence automatically and who sees it.
//
// GET /api/user/presence/settings
func (r *UserService) GetPresenceSettings(ctx context.Context, opts ...CallOption) (*UserGetPresenceSettingsResponseData, error) {
	var out *UserGetPresenceSettingsResponseData
	err := r.c.call(ctx, "GET", "/api/user/presence/settings", nil, nil, false, "data", &out, opts)
	return out, err
}

// GetProfile — Returns the caller's own profile.
//
// GET /api/user/profile
func (r *UserService) GetProfile(ctx context.Context, opts ...CallOption) (*UserProfile, error) {
	var out *UserProfile
	err := r.c.call(ctx, "GET", "/api/user/profile", nil, nil, false, "profile", &out, opts)
	return out, err
}

// GetSecurity — Returns whether the caller has two-factor sign-in on.
//
// GET /api/user/security
func (r *UserService) GetSecurity(ctx context.Context, opts ...CallOption) (*UserGetSecurityResponse, error) {
	var out *UserGetSecurityResponse
	err := r.c.call(ctx, "GET", "/api/user/security", nil, nil, false, "", &out, opts)
	return out, err
}

// ListActivityLogs — Lists the activity the caller may see in the audit log — their own, and their organizations' for administrators.
//
// GET /api/user/activity-logs
func (r *UserService) ListActivityLogs(ctx context.Context, query *UserListActivityLogsQuery, opts ...CallOption) ([]ActivityLog, error) {
	var out []ActivityLog
	err := r.c.call(ctx, "GET", "/api/user/activity-logs", query, nil, false, "logs", &out, opts)
	return out, err
}

// ListRecognizedDevices — Lists the devices and places the caller has signed in from.
//
// GET /api/user/recognized-devices
func (r *UserService) ListRecognizedDevices(ctx context.Context, opts ...CallOption) ([]UserListRecognizedDevicesResponseDevicesItem, error) {
	var out []UserListRecognizedDevicesResponseDevicesItem
	err := r.c.call(ctx, "GET", "/api/user/recognized-devices", nil, nil, false, "devices", &out, opts)
	return out, err
}

// ListSignatures — Lists the caller's saved signatures for signing documents.
//
// GET /api/user/signatures
func (r *UserService) ListSignatures(ctx context.Context, opts ...CallOption) ([]UserListSignaturesResponseDataSignaturesItem, error) {
	var out []UserListSignaturesResponseDataSignaturesItem
	err := r.c.call(ctx, "GET", "/api/user/signatures", nil, nil, false, "data.signatures", &out, opts)
	return out, err
}

// PatchPreferences — Merges the given keys into the caller's preferences, exactly as `user.updatePreferences` does.
//
// PATCH /api/user/preferences
func (r *UserService) PatchPreferences(ctx context.Context, body *UserPreferences, opts ...CallOption) (*UserPreferences, error) {
	var out *UserPreferences
	err := r.c.call(ctx, "PATCH", "/api/user/preferences", nil, body, false, "", &out, opts)
	return out, err
}

// ReleaseHandle — Gives up the caller's public creator handle.
//
// DELETE /api/user/handle
func (r *UserService) ReleaseHandle(ctx context.Context, opts ...CallOption) (*UserReleaseHandleResponse, error) {
	var out *UserReleaseHandleResponse
	err := r.c.call(ctx, "DELETE", "/api/user/handle", nil, nil, false, "", &out, opts)
	return out, err
}

// SetAIMemorySwitch — Lets the AI assistant remember things about the caller, or stops it.
//
// PUT /api/user/ai-memory/switch
func (r *UserService) SetAIMemorySwitch(ctx context.Context, body *UserSetAiMemorySwitchBody, opts ...CallOption) (*UserSetAiMemorySwitchResponse, error) {
	var out *UserSetAiMemorySwitchResponse
	err := r.c.call(ctx, "PUT", "/api/user/ai-memory/switch", nil, body, false, "", &out, opts)
	return out, err
}

// SetHandoff — Marks the caller away until a date, with a message — or back with `until: null`.
//
// PUT /api/user/handoff
func (r *UserService) SetHandoff(ctx context.Context, body *UserSetHandoffBody, opts ...CallOption) (*UserSetHandoffResponse, error) {
	var out *UserSetHandoffResponse
	err := r.c.call(ctx, "PUT", "/api/user/handoff", nil, body, false, "", &out, opts)
	return out, err
}

// SetPresence — Sets the caller's own presence and status text, optionally until a time.
//
// PUT /api/user/presence
func (r *UserService) SetPresence(ctx context.Context, body *UserSetPresenceBody, opts ...CallOption) (*UserSetPresenceResponse, error) {
	var out *UserSetPresenceResponse
	err := r.c.call(ctx, "PUT", "/api/user/presence", nil, body, false, "", &out, opts)
	return out, err
}

// UpdateAIMemory — Changes a note the AI assistant remembers.
//
// PATCH /api/user/ai-memory
func (r *UserService) UpdateAIMemory(ctx context.Context, body *UserUpdateAiMemoryBody, query *UserUpdateAiMemoryQuery, opts ...CallOption) (*UserUpdateAiMemoryResponse, error) {
	var out *UserUpdateAiMemoryResponse
	err := r.c.call(ctx, "PATCH", "/api/user/ai-memory", query, body, false, "", &out, opts)
	return out, err
}

// UpdateAvatar — Uploads a new profile picture for the caller.
//
// An image of at most 5 MB; SVG is refused.
//
// POST /api/user/avatar
func (r *UserService) UpdateAvatar(ctx context.Context, body *UserUpdateAvatarBody, opts ...CallOption) (string, error) {
	var out string
	err := r.c.call(ctx, "POST", "/api/user/avatar", nil, body, true, "avatarUrl", &out, opts)
	return out, err
}

// UpdatePreferences — Merges the given keys into the caller's preferences and returns them all.
//
// A shallow merge: keys left out keep their value. Keys the platform does not know are dropped, and security notifications cannot be switched off.
//
// PUT /api/user/preferences
func (r *UserService) UpdatePreferences(ctx context.Context, body *UserPreferences, opts ...CallOption) (*UserPreferences, error) {
	var out *UserPreferences
	err := r.c.call(ctx, "PUT", "/api/user/preferences", nil, body, false, "", &out, opts)
	return out, err
}

// UpdatePresenceSettings — Changes what sets the caller's presence automatically and who sees it.
//
// PUT /api/user/presence/settings
func (r *UserService) UpdatePresenceSettings(ctx context.Context, body *UserUpdatePresenceSettingsBody, opts ...CallOption) (*UserUpdatePresenceSettingsResponseData, error) {
	var out *UserUpdatePresenceSettingsResponseData
	err := r.c.call(ctx, "PUT", "/api/user/presence/settings", nil, body, false, "data", &out, opts)
	return out, err
}

// UpdateProfile — Changes the caller's profile.
//
// `fullName` is required. The email address is changed elsewhere, with confirmation. Send `If-Unmodified-Since` with the profile's `updatedAt` to get a 409 instead of overwriting a newer save.
//
// PUT /api/user/profile
func (r *UserService) UpdateProfile(ctx context.Context, body *UpdateProfileInput, opts ...CallOption) (*UserProfile, error) {
	var out *UserProfile
	err := r.c.call(ctx, "PUT", "/api/user/profile", nil, body, false, "profile", &out, opts)
	return out, err
}

// UpdateSignature — Renames a saved signature or makes it the default.
//
// PATCH /api/user/signatures/{signatureId}
//
// signatureID: The signature's id.
func (r *UserService) UpdateSignature(ctx context.Context, signatureID string, body *UserUpdateSignatureBody, opts ...CallOption) (*UserUpdateSignatureResponse, error) {
	var out *UserUpdateSignatureResponse
	err := r.c.call(ctx, "PATCH", "/api/user/signatures/"+pathSeg(signatureID), nil, body, false, "", &out, opts)
	return out, err
}

// UsersService — Users
type UsersService struct {
	c *Client
}

func newUsersService(c *Client) *UsersService {
	return &UsersService{c: c}
}

// ListOrganizations — Lists the organizations a user belongs to — only for the caller.
//
// GET /api/users/{userId}/organizations
//
// userID: The caller's own user id; any other answers 403.
func (r *UsersService) ListOrganizations(ctx context.Context, userID string, opts ...CallOption) ([]Organization, error) {
	var out []Organization
	err := r.c.call(ctx, "GET", "/api/users/"+pathSeg(userID)+"/organizations", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListRoles — Lists the roles a user holds, directly and through teams — only for the caller.
//
// GET /api/users/{userId}/roles
//
// userID: The caller's own user id; any other answers 403.
func (r *UsersService) ListRoles(ctx context.Context, userID string, opts ...CallOption) ([]Role, error) {
	var out []Role
	err := r.c.call(ctx, "GET", "/api/users/"+pathSeg(userID)+"/roles", nil, nil, false, "roles", &out, opts)
	return out, err
}

// ListWorkspaces — Lists the workspaces a user is a member of — only for the caller.
//
// GET /api/users/{userId}/workspaces
//
// userID: The caller's own user id; any other answers 403.
func (r *UsersService) ListWorkspaces(ctx context.Context, userID string, opts ...CallOption) ([]Workspace, error) {
	var out []Workspace
	err := r.c.call(ctx, "GET", "/api/users/"+pathSeg(userID)+"/workspaces", nil, nil, false, "data", &out, opts)
	return out, err
}

// VariablesService — Workflow variables: Values workflows store and read between runs.
type VariablesService struct {
	c *Client
}

func newVariablesService(c *Client) *VariablesService {
	return &VariablesService{c: c}
}

// Get — Reads a workflow variable by its key.
//
// Needs `workflow.read` in the variable's workspace; an expired variable reads as not found.
//
// GET /api/variables/get
func (r *VariablesService) Get(ctx context.Context, query *VariablesGetQuery, opts ...CallOption) (*VariablesGetResponse, error) {
	var out *VariablesGetResponse
	err := r.c.call(ctx, "GET", "/api/variables/get", query, nil, false, "", &out, opts)
	return out, err
}

// Save — Creates or replaces a workflow variable.
//
// Needs `workflow.update` in the workspace. `key` must be the one `workspaceId`, `name` and (for `workflow` scope) `workflowId` make.
//
// POST /api/variables/save
func (r *VariablesService) Save(ctx context.Context, body *VariablesSaveBody, opts ...CallOption) (*VariablesSaveResponse, error) {
	var out *VariablesSaveResponse
	err := r.c.call(ctx, "POST", "/api/variables/save", nil, body, false, "", &out, opts)
	return out, err
}

// WebhooksService — Webhooks: Subscriptions that post platform events to your URL.
type WebhooksService struct {
	c *Client
}

func newWebhooksService(c *Client) *WebhooksService {
	return &WebhooksService{c: c}
}

// GetSubscription — Returns one of the caller's webhook subscriptions, without its secret.
//
// GET /api/webhooks/subscribe/{id}
//
// id: The subscription's id.
func (r *WebhooksService) GetSubscription(ctx context.Context, id string, opts ...CallOption) (*WebhookSubscription, error) {
	var out *WebhookSubscription
	err := r.c.call(ctx, "GET", "/api/webhooks/subscribe/"+pathSeg(id), nil, nil, false, "data", &out, opts)
	return out, err
}

// ListSubscriptions — Lists the caller's webhook subscriptions, without their secrets.
//
// GET /api/webhooks/subscribe
func (r *WebhooksService) ListSubscriptions(ctx context.Context, query *WebhooksListSubscriptionsQuery, opts ...CallOption) ([]WebhookSubscription, error) {
	var out []WebhookSubscription
	err := r.c.call(ctx, "GET", "/api/webhooks/subscribe", query, nil, false, "data", &out, opts)
	return out, err
}

// Subscribe — Subscribes a URL to an event of an organization; the answer carries the signing secret, once.
//
// Counts against the plan's `max_webhook_subscriptions`.
//
// POST /api/webhooks/subscribe
func (r *WebhooksService) Subscribe(ctx context.Context, body *WebhooksSubscribeBody, opts ...CallOption) (*WebhooksSubscribeResponseData, error) {
	var out *WebhooksSubscribeResponseData
	err := r.c.call(ctx, "POST", "/api/webhooks/subscribe", nil, body, false, "data", &out, opts)
	return out, err
}

// Unsubscribe — Deletes a webhook subscription.
//
// DELETE /api/webhooks/subscribe/{id}
//
// id: The subscription's id.
func (r *WebhooksService) Unsubscribe(ctx context.Context, id string, opts ...CallOption) (*WebhooksUnsubscribeResponse, error) {
	var out *WebhooksUnsubscribeResponse
	err := r.c.call(ctx, "DELETE", "/api/webhooks/subscribe/"+pathSeg(id), nil, nil, false, "", &out, opts)
	return out, err
}

// WorkflowsService — Workflows
type WorkflowsService struct {
	c *Client
}

func newWorkflowsService(c *Client) *WorkflowsService {
	return &WorkflowsService{c: c}
}

// Create — Creates a workflow.
//
// Needs `workflow.create`; counts against the plan's `max_workflows`. Connections a node names must be ones the caller may use.
//
// POST /api/workflows
func (r *WorkflowsService) Create(ctx context.Context, body *WorkflowsCreateBody, opts ...CallOption) (*Workflow, error) {
	var out *Workflow
	err := r.c.call(ctx, "POST", "/api/workflows", nil, body, false, "", &out, opts)
	return out, err
}

// CreateAlert — Emails the given addresses when a workflow fails a number of times within a window.
//
// POST /api/workflows/{workflowId}/alerts
//
// workflowID: The workflow's id.
func (r *WorkflowsService) CreateAlert(ctx context.Context, workflowID string, body *WorkflowsCreateAlertBody, opts ...CallOption) (*WorkflowsCreateAlertResponseAlert, error) {
	var out *WorkflowsCreateAlertResponseAlert
	err := r.c.call(ctx, "POST", "/api/workflows/"+pathSeg(workflowID)+"/alerts", nil, body, false, "alert", &out, opts)
	return out, err
}

// CreateVersion — Saves a version of a workflow.
//
// POST /api/workflows/{workflowId}/versions
//
// workflowID: The workflow's id.
func (r *WorkflowsService) CreateVersion(ctx context.Context, workflowID string, body *WorkflowsCreateVersionBody, opts ...CallOption) (*WorkflowVersion, error) {
	var out *WorkflowVersion
	err := r.c.call(ctx, "POST", "/api/workflows/"+pathSeg(workflowID)+"/versions", nil, body, false, "", &out, opts)
	return out, err
}

// Delete — Moves a workflow to the trash; it stops running.
//
// DELETE /api/workflows/{workflowId}
//
// workflowID: The workflow's id.
func (r *WorkflowsService) Delete(ctx context.Context, workflowID string, opts ...CallOption) (*WorkflowsDeleteResponse, error) {
	var out *WorkflowsDeleteResponse
	err := r.c.call(ctx, "DELETE", "/api/workflows/"+pathSeg(workflowID), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteAlert — Removes a workflow's failure alert.
//
// DELETE /api/workflows/{workflowId}/alerts
//
// workflowID: The workflow's id.
func (r *WorkflowsService) DeleteAlert(ctx context.Context, workflowID string, query *WorkflowsDeleteAlertQuery, opts ...CallOption) (*WorkflowsDeleteAlertResponse, error) {
	var out *WorkflowsDeleteAlertResponse
	err := r.c.call(ctx, "DELETE", "/api/workflows/"+pathSeg(workflowID)+"/alerts", query, nil, false, "", &out, opts)
	return out, err
}

// DeleteDeadLetter — Removes an entry from the dead-letter queue.
//
// DELETE /api/workflows/dead-letter-queue
func (r *WorkflowsService) DeleteDeadLetter(ctx context.Context, query *WorkflowsDeleteDeadLetterQuery, opts ...CallOption) (*WorkflowsDeleteDeadLetterResponse, error) {
	var out *WorkflowsDeleteDeadLetterResponse
	err := r.c.call(ctx, "DELETE", "/api/workflows/dead-letter-queue", query, nil, false, "", &out, opts)
	return out, err
}

// DeleteVersion — Deletes one saved version of a workflow.
//
// DELETE /api/workflows/{workflowId}/versions/{versionId}
//
// workflowID: The workflow's id.
// versionID: The version's id.
func (r *WorkflowsService) DeleteVersion(ctx context.Context, workflowID string, versionID string, opts ...CallOption) (*WorkflowsDeleteVersionResponse, error) {
	var out *WorkflowsDeleteVersionResponse
	err := r.c.call(ctx, "DELETE", "/api/workflows/"+pathSeg(workflowID)+"/versions/"+pathSeg(versionID), nil, nil, false, "", &out, opts)
	return out, err
}

// DeleteVersions — Deletes every saved version of a workflow.
//
// DELETE /api/workflows/{workflowId}/versions
//
// workflowID: The workflow's id.
func (r *WorkflowsService) DeleteVersions(ctx context.Context, workflowID string, opts ...CallOption) (*WorkflowsDeleteVersionsResponse, error) {
	var out *WorkflowsDeleteVersionsResponse
	err := r.c.call(ctx, "DELETE", "/api/workflows/"+pathSeg(workflowID)+"/versions", nil, nil, false, "", &out, opts)
	return out, err
}

// Duplicate — Copies a workflow as a new draft.
//
// Counts against the plan's `max_workflows`.
//
// POST /api/workflows/{workflowId}/duplicate
//
// workflowID: The workflow's id.
func (r *WorkflowsService) Duplicate(ctx context.Context, workflowID string, opts ...CallOption) (*Workflow, error) {
	var out *Workflow
	err := r.c.call(ctx, "POST", "/api/workflows/"+pathSeg(workflowID)+"/duplicate", nil, nil, false, "", &out, opts)
	return out, err
}

// Execute — Runs a workflow now with the given trigger data; the run is queued.
//
// Needs `workflow.execute` or `workflow.update`. Poll `workflows.listExecutions` for the outcome.
//
// POST /api/workflows/{workflowId}/executions
//
// workflowID: The workflow's id.
func (r *WorkflowsService) Execute(ctx context.Context, workflowID string, body *WorkflowsExecuteBody, opts ...CallOption) (*WorkflowsExecuteResponse, error) {
	var out *WorkflowsExecuteResponse
	err := r.c.call(ctx, "POST", "/api/workflows/"+pathSeg(workflowID)+"/executions", nil, body, false, "", &out, opts)
	return out, err
}

// Export — Exports a workflow as JSON for `workflows.import`, without stored credentials.
//
// GET /api/workflows/{workflowId}/export
//
// workflowID: The workflow's id.
func (r *WorkflowsService) Export(ctx context.Context, workflowID string, opts ...CallOption) (*WorkflowExport, error) {
	var out *WorkflowExport
	err := r.c.call(ctx, "GET", "/api/workflows/"+pathSeg(workflowID)+"/export", nil, nil, false, "", &out, opts)
	return out, err
}

// Get — Returns a workflow — trigger, nodes and edges — with its run counts.
//
// GET /api/workflows/{workflowId}
//
// workflowID: The workflow's id.
func (r *WorkflowsService) Get(ctx context.Context, workflowID string, opts ...CallOption) (*Workflow, error) {
	var out *Workflow
	err := r.c.call(ctx, "GET", "/api/workflows/"+pathSeg(workflowID), nil, nil, false, "", &out, opts)
	return out, err
}

// GetVersion — Returns one saved version of a workflow.
//
// GET /api/workflows/{workflowId}/versions/{versionId}
//
// workflowID: The workflow's id.
// versionID: The version's id.
func (r *WorkflowsService) GetVersion(ctx context.Context, workflowID string, versionID string, opts ...CallOption) (*WorkflowVersion, error) {
	var out *WorkflowVersion
	err := r.c.call(ctx, "GET", "/api/workflows/"+pathSeg(workflowID)+"/versions/"+pathSeg(versionID), nil, nil, false, "", &out, opts)
	return out, err
}

// Import — Creates a workflow from an export.
//
// Takes the export as it is (`{ workflow: {…} }`) or its fields at the top level.
//
// POST /api/workflows/import
func (r *WorkflowsService) Import(ctx context.Context, body *WorkflowsImportBody, opts ...CallOption) (*Workflow, error) {
	var out *Workflow
	err := r.c.call(ctx, "POST", "/api/workflows/import", nil, body, false, "workflow", &out, opts)
	return out, err
}

// InstallRecipe — Creates a draft workflow from a recipe and the answers to its inputs.
//
// POST /api/workflows/recipes
func (r *WorkflowsService) InstallRecipe(ctx context.Context, body *WorkflowsInstallRecipeBody, opts ...CallOption) (*WorkflowsInstallRecipeResponseData, error) {
	var out *WorkflowsInstallRecipeResponseData
	err := r.c.call(ctx, "POST", "/api/workflows/recipes", nil, body, false, "data", &out, opts)
	return out, err
}

// List — Lists the workflows the caller may open, with their run counts.
//
// GET /api/workflows
func (r *WorkflowsService) List(ctx context.Context, query *WorkflowsListQuery, opts ...CallOption) ([]Workflow, error) {
	var out []Workflow
	err := r.c.call(ctx, "GET", "/api/workflows", query, nil, false, "workflows", &out, opts)
	return out, err
}

// ListAlerts — Lists a workflow's failure alerts.
//
// GET /api/workflows/{workflowId}/alerts
//
// workflowID: The workflow's id.
func (r *WorkflowsService) ListAlerts(ctx context.Context, workflowID string, opts ...CallOption) ([]WorkflowsListAlertsResponseAlertsItem, error) {
	var out []WorkflowsListAlertsResponseAlertsItem
	err := r.c.call(ctx, "GET", "/api/workflows/"+pathSeg(workflowID)+"/alerts", nil, nil, false, "alerts", &out, opts)
	return out, err
}

// ListAllDeadLetters — Lists the runs that failed for good, of every workflow the caller may open.
//
// GET /api/workflows/dead-letter-queue
func (r *WorkflowsService) ListAllDeadLetters(ctx context.Context, opts ...CallOption) ([]DeadLetterJob, error) {
	var out []DeadLetterJob
	err := r.c.call(ctx, "GET", "/api/workflows/dead-letter-queue", nil, nil, false, "", &out, opts)
	return out, err
}

// ListDeadLetterQueue — Lists a workflow's runs that failed for good.
//
// GET /api/workflows/{workflowId}/dead-letter-queue
//
// workflowID: The workflow's id.
func (r *WorkflowsService) ListDeadLetterQueue(ctx context.Context, workflowID string, opts ...CallOption) ([]DeadLetterJob, error) {
	var out []DeadLetterJob
	err := r.c.call(ctx, "GET", "/api/workflows/"+pathSeg(workflowID)+"/dead-letter-queue", nil, nil, false, "", &out, opts)
	return out, err
}

// ListExecutions — Lists a workflow's latest runs with their steps.
//
// With an execution id (`exec-…`) in place of the workflow id it answers that one run instead.
//
// GET /api/workflows/{workflowId}/executions
//
// workflowID: The workflow's id.
func (r *WorkflowsService) ListExecutions(ctx context.Context, workflowID string, query *WorkflowsListExecutionsQuery, opts ...CallOption) ([]WorkflowExecution, error) {
	var out []WorkflowExecution
	err := r.c.call(ctx, "GET", "/api/workflows/"+pathSeg(workflowID)+"/executions", query, nil, false, "", &out, opts)
	return out, err
}

// ListRecipes — Lists the ready-made workflows that can be installed with a few answers.
//
// GET /api/workflows/recipes
func (r *WorkflowsService) ListRecipes(ctx context.Context, opts ...CallOption) ([]WorkflowsListRecipesResponseDataRecipesItem, error) {
	var out []WorkflowsListRecipesResponseDataRecipesItem
	err := r.c.call(ctx, "GET", "/api/workflows/recipes", nil, nil, false, "data.recipes", &out, opts)
	return out, err
}

// ListVersions — Lists a workflow's saved versions, newest first.
//
// GET /api/workflows/{workflowId}/versions
//
// workflowID: The workflow's id.
func (r *WorkflowsService) ListVersions(ctx context.Context, workflowID string, opts ...CallOption) ([]WorkflowVersion, error) {
	var out []WorkflowVersion
	err := r.c.call(ctx, "GET", "/api/workflows/"+pathSeg(workflowID)+"/versions", nil, nil, false, "", &out, opts)
	return out, err
}

// Queue — Lists a workflow's queued runs, or with `action=rateLimit` its rate-limit state.
//
// GET /api/workflows/{workflowId}/queue
//
// workflowID: The workflow's id.
func (r *WorkflowsService) Queue(ctx context.Context, workflowID string, query *WorkflowsQueueQuery, opts ...CallOption) (any, error) {
	var out any
	err := r.c.call(ctx, "GET", "/api/workflows/"+pathSeg(workflowID)+"/queue", query, nil, false, "", &out, opts)
	return out, err
}

// ReplayExecution — Runs a past execution again from one of its nodes, with the data it had there.
//
// POST /api/workflows/{workflowId}/executions/{executionId}/replay
//
// workflowID: The workflow's id.
// executionID: The execution's id.
func (r *WorkflowsService) ReplayExecution(ctx context.Context, workflowID string, executionID string, body *WorkflowsReplayExecutionBody, opts ...CallOption) (map[string]any, error) {
	var out map[string]any
	err := r.c.call(ctx, "POST", "/api/workflows/"+pathSeg(workflowID)+"/executions/"+pathSeg(executionID)+"/replay", nil, body, false, "data.execution", &out, opts)
	return out, err
}

// RestoreVersion — Puts a saved version back as the workflow.
//
// POST /api/workflows/{workflowId}/versions/{versionId}
//
// workflowID: The workflow's id.
// versionID: The version's id.
func (r *WorkflowsService) RestoreVersion(ctx context.Context, workflowID string, versionID string, body *WorkflowsRestoreVersionBody, opts ...CallOption) (*Workflow, error) {
	var out *Workflow
	err := r.c.call(ctx, "POST", "/api/workflows/"+pathSeg(workflowID)+"/versions/"+pathSeg(versionID), nil, body, false, "workflow", &out, opts)
	return out, err
}

// RetryDeadJob — Runs a failed run of the dead-letter queue again.
//
// POST /api/workflows/{workflowId}/dead-letter-queue
//
// workflowID: The workflow's id.
func (r *WorkflowsService) RetryDeadJob(ctx context.Context, workflowID string, body *WorkflowsRetryDeadJobBody, opts ...CallOption) (*WorkflowExecution, error) {
	var out *WorkflowExecution
	err := r.c.call(ctx, "POST", "/api/workflows/"+pathSeg(workflowID)+"/dead-letter-queue", nil, body, false, "executionLog", &out, opts)
	return out, err
}

// SetStatus — Activates, pauses or drafts a workflow.
//
// PATCH /api/workflows/{workflowId}/status
//
// workflowID: The workflow's id.
func (r *WorkflowsService) SetStatus(ctx context.Context, workflowID string, body *WorkflowsSetStatusBody, opts ...CallOption) (*Workflow, error) {
	var out *Workflow
	err := r.c.call(ctx, "PATCH", "/api/workflows/"+pathSeg(workflowID)+"/status", nil, body, false, "workflow", &out, opts)
	return out, err
}

// ToggleStatus — Switches a workflow between active and paused.
//
// POST /api/workflows/{workflowId}/status
//
// workflowID: The workflow's id.
//
// Deprecated: Toggling depends on the state the caller last saw; set the status explicitly. Use workflows.setStatus.
func (r *WorkflowsService) ToggleStatus(ctx context.Context, workflowID string, opts ...CallOption) (*Workflow, error) {
	deprecated("workflows.toggleStatus")
	var out *Workflow
	err := r.c.call(ctx, "POST", "/api/workflows/"+pathSeg(workflowID)+"/status", nil, nil, false, "workflow", &out, opts)
	return out, err
}

// Update — Changes a workflow; only the fields sent change.
//
// Needs `workflow.update`. A trigger or status change reaches the scheduler at once.
//
// PUT /api/workflows/{workflowId}
//
// workflowID: The workflow's id.
func (r *WorkflowsService) Update(ctx context.Context, workflowID string, body *WorkflowsUpdateBody, opts ...CallOption) (*Workflow, error) {
	var out *Workflow
	err := r.c.call(ctx, "PUT", "/api/workflows/"+pathSeg(workflowID), nil, body, false, "", &out, opts)
	return out, err
}

// Usage — Returns an organization's workflow runs per day against its daily cap.
//
// GET /api/workflows/usage
func (r *WorkflowsService) Usage(ctx context.Context, query *WorkflowsUsageQuery, opts ...CallOption) (*WorkflowsUsageResponse, error) {
	var out *WorkflowsUsageResponse
	err := r.c.call(ctx, "GET", "/api/workflows/usage", query, nil, false, "", &out, opts)
	return out, err
}

// WorkspacesService — Workspaces
type WorkspacesService struct {
	c *Client
}

func newWorkspacesService(c *Client) *WorkspacesService {
	return &WorkspacesService{c: c}
}

// AddMember — Adds a member of the organization to a workspace, optionally with a role in it.
//
// Needs `workspace.manage_members`.
//
// POST /api/workspaces/{workspaceId}/members
//
// workspaceID: The workspace's id.
func (r *WorkspacesService) AddMember(ctx context.Context, workspaceID string, body *WorkspacesAddMemberBody, opts ...CallOption) (*WorkspaceMember, error) {
	var out *WorkspaceMember
	err := r.c.call(ctx, "POST", "/api/workspaces/"+pathSeg(workspaceID)+"/members", nil, body, false, "data", &out, opts)
	return out, err
}

// AssignMemberRole — Gives a workspace member a role in the workspace.
//
// Needs `role.assign`; only roles within the caller's own permissions.
//
// POST /api/workspaces/{workspaceId}/members/{memberId}/assign-role
//
// workspaceID: The workspace's id.
// memberID: The membership's id.
func (r *WorkspacesService) AssignMemberRole(ctx context.Context, workspaceID string, memberID string, body *WorkspacesAssignMemberRoleBody, opts ...CallOption) (*RoleAssignment, error) {
	var out *RoleAssignment
	err := r.c.call(ctx, "POST", "/api/workspaces/"+pathSeg(workspaceID)+"/members/"+pathSeg(memberID)+"/assign-role", nil, body, false, "data", &out, opts)
	return out, err
}

// AssignTeam — Gives a team a role in a workspace.
//
// Needs `team.assign`.
//
// POST /api/workspaces/{workspaceId}/teams
//
// workspaceID: The workspace's id.
func (r *WorkspacesService) AssignTeam(ctx context.Context, workspaceID string, body *WorkspacesAssignTeamBody, opts ...CallOption) (*TeamRoleAssignment, error) {
	var out *TeamRoleAssignment
	err := r.c.call(ctx, "POST", "/api/workspaces/"+pathSeg(workspaceID)+"/teams", nil, body, false, "data", &out, opts)
	return out, err
}

// CancelDeletion — Cancels a workspace's scheduled deletion.
//
// DELETE /api/workspaces/{workspaceId}/scheduled-deletion
//
// workspaceID: The workspace's id.
func (r *WorkspacesService) CancelDeletion(ctx context.Context, workspaceID string, opts ...CallOption) (*WorkspacesCancelDeletionResponseData, error) {
	var out *WorkspacesCancelDeletionResponseData
	err := r.c.call(ctx, "DELETE", "/api/workspaces/"+pathSeg(workspaceID)+"/scheduled-deletion", nil, nil, false, "data", &out, opts)
	return out, err
}

// Create — Creates a workspace in an organization, with the caller as its first member.
//
// Needs `workspace.create`; counts against the plan's `max_workspaces`.
//
// POST /api/workspaces
func (r *WorkspacesService) Create(ctx context.Context, body *WorkspacesCreateBody, opts ...CallOption) (*Workspace, error) {
	var out *Workspace
	err := r.c.call(ctx, "POST", "/api/workspaces", nil, body, false, "data", &out, opts)
	return out, err
}

// Delete — Schedules a workspace for deletion in 14 days.
//
// Needs `workspace.delete`. Until then it can be cancelled with `workspaces.cancelDeletion`.
//
// DELETE /api/workspaces/{workspaceId}
//
// workspaceID: The workspace's id.
func (r *WorkspacesService) Delete(ctx context.Context, workspaceID string, opts ...CallOption) (*WorkspacesDeleteResponseData, error) {
	var out *WorkspacesDeleteResponseData
	err := r.c.call(ctx, "DELETE", "/api/workspaces/"+pathSeg(workspaceID), nil, nil, false, "data", &out, opts)
	return out, err
}

// DeleteEnvVar — Removes a workspace variable.
//
// DELETE /api/workspaces/{workspaceId}/env-vars
//
// workspaceID: The workspace's id.
func (r *WorkspacesService) DeleteEnvVar(ctx context.Context, workspaceID string, query *WorkspacesDeleteEnvVarQuery, opts ...CallOption) (*WorkspacesDeleteEnvVarResponse, error) {
	var out *WorkspacesDeleteEnvVarResponse
	err := r.c.call(ctx, "DELETE", "/api/workspaces/"+pathSeg(workspaceID)+"/env-vars", query, nil, false, "", &out, opts)
	return out, err
}

// Get — Returns one workspace.
//
// GET /api/workspaces/{workspaceId} — Token permission: read:projects.
//
// workspaceID: The workspace's id.
func (r *WorkspacesService) Get(ctx context.Context, workspaceID string, opts ...CallOption) (*Workspace, error) {
	var out *Workspace
	err := r.c.call(ctx, "GET", "/api/workspaces/"+pathSeg(workspaceID), nil, nil, false, "data", &out, opts)
	return out, err
}

// GetScheduledDeletion — Returns when a workspace is due to be deleted, if it is.
//
// GET /api/workspaces/{workspaceId}/scheduled-deletion
//
// workspaceID: The workspace's id.
func (r *WorkspacesService) GetScheduledDeletion(ctx context.Context, workspaceID string, opts ...CallOption) (*WorkspacesGetScheduledDeletionResponseData, error) {
	var out *WorkspacesGetScheduledDeletionResponseData
	err := r.c.call(ctx, "GET", "/api/workspaces/"+pathSeg(workspaceID)+"/scheduled-deletion", nil, nil, false, "data", &out, opts)
	return out, err
}

// List — Lists the workspaces the caller can open, optionally of one organization.
//
// GET /api/workspaces — Token permission: read:projects.
func (r *WorkspacesService) List(ctx context.Context, query *WorkspacesListQuery, opts ...CallOption) ([]Workspace, error) {
	var out []Workspace
	err := r.c.call(ctx, "GET", "/api/workspaces", query, nil, false, "data", &out, opts)
	return out, err
}

// ListEnvVars — Lists a workspace's variables, which its workflows read as `{{env.NAME}}`.
//
// Needs `workflow.read` in the workspace. A secret's value is never answered.
//
// GET /api/workspaces/{workspaceId}/env-vars
//
// workspaceID: The workspace's id.
func (r *WorkspacesService) ListEnvVars(ctx context.Context, workspaceID string, opts ...CallOption) ([]WorkspaceEnvVar, error) {
	var out []WorkspaceEnvVar
	err := r.c.call(ctx, "GET", "/api/workspaces/"+pathSeg(workspaceID)+"/env-vars", nil, nil, false, "envVars", &out, opts)
	return out, err
}

// ListMembers — Lists a workspace's members with their user details and roles.
//
// GET /api/workspaces/{workspaceId}/members
//
// workspaceID: The workspace's id.
func (r *WorkspacesService) ListMembers(ctx context.Context, workspaceID string, opts ...CallOption) ([]WorkspaceMember, error) {
	var out []WorkspaceMember
	err := r.c.call(ctx, "GET", "/api/workspaces/"+pathSeg(workspaceID)+"/members", nil, nil, false, "data", &out, opts)
	return out, err
}

// ListTeams — Lists the teams with a role in a workspace.
//
// GET /api/workspaces/{workspaceId}/teams
//
// workspaceID: The workspace's id.
func (r *WorkspacesService) ListTeams(ctx context.Context, workspaceID string, opts ...CallOption) ([]WorkspacesListTeamsResponseDataItem, error) {
	var out []WorkspacesListTeamsResponseDataItem
	err := r.c.call(ctx, "GET", "/api/workspaces/"+pathSeg(workspaceID)+"/teams", nil, nil, false, "data", &out, opts)
	return out, err
}

// RemoveMember — Removes a member from a workspace.
//
// DELETE /api/workspaces/{workspaceId}/members/{memberId}
//
// workspaceID: The workspace's id.
// memberID: The membership's id.
func (r *WorkspacesService) RemoveMember(ctx context.Context, workspaceID string, memberID string, opts ...CallOption) (*WorkspacesRemoveMemberResponse, error) {
	var out *WorkspacesRemoveMemberResponse
	err := r.c.call(ctx, "DELETE", "/api/workspaces/"+pathSeg(workspaceID)+"/members/"+pathSeg(memberID), nil, nil, false, "", &out, opts)
	return out, err
}

// RemoveMemberRole — Takes a role in the workspace away from a member.
//
// DELETE /api/workspaces/{workspaceId}/members/{memberId}/assign-role
//
// workspaceID: The workspace's id.
// memberID: The membership's id.
func (r *WorkspacesService) RemoveMemberRole(ctx context.Context, workspaceID string, memberID string, query *WorkspacesRemoveMemberRoleQuery, opts ...CallOption) (*WorkspacesRemoveMemberRoleResponse, error) {
	var out *WorkspacesRemoveMemberRoleResponse
	err := r.c.call(ctx, "DELETE", "/api/workspaces/"+pathSeg(workspaceID)+"/members/"+pathSeg(memberID)+"/assign-role", query, nil, false, "", &out, opts)
	return out, err
}

// RemoveTeam — Takes a team's roles in a workspace away — all of them, or one with `roleId`.
//
// DELETE /api/workspaces/{workspaceId}/teams/{teamId}
//
// workspaceID: The workspace's id.
// teamID: The team's id.
func (r *WorkspacesService) RemoveTeam(ctx context.Context, workspaceID string, teamID string, query *WorkspacesRemoveTeamQuery, opts ...CallOption) (*WorkspacesRemoveTeamResponse, error) {
	var out *WorkspacesRemoveTeamResponse
	err := r.c.call(ctx, "DELETE", "/api/workspaces/"+pathSeg(workspaceID)+"/teams/"+pathSeg(teamID), query, nil, false, "", &out, opts)
	return out, err
}

// SetEnvVar — Creates or replaces a workspace variable.
//
// Needs `workflow.update` in the workspace. A secret is sealed at rest; sending the mask back as its value keeps the stored one.
//
// POST /api/workspaces/{workspaceId}/env-vars
//
// workspaceID: The workspace's id.
func (r *WorkspacesService) SetEnvVar(ctx context.Context, workspaceID string, body *WorkspacesSetEnvVarBody, opts ...CallOption) (*WorkspaceEnvVar, error) {
	var out *WorkspaceEnvVar
	err := r.c.call(ctx, "POST", "/api/workspaces/"+pathSeg(workspaceID)+"/env-vars", nil, body, false, "envVar", &out, opts)
	return out, err
}

// Update — Changes a workspace's name, description, icon, color or settings.
//
// PUT /api/workspaces/{workspaceId}
//
// workspaceID: The workspace's id.
func (r *WorkspacesService) Update(ctx context.Context, workspaceID string, body *WorkspacesUpdateBody, opts ...CallOption) (*Workspace, error) {
	var out *Workspace
	err := r.c.call(ctx, "PUT", "/api/workspaces/"+pathSeg(workspaceID), nil, body, false, "data", &out, opts)
	return out, err
}
