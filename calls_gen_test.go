// Code generated from sdks/openapi.json by `pnpm sdk:generate`. DO NOT EDIT.
// Generated from sdks/openapi.json by `pnpm sdk:generate` — do not edit.

package crafvy

import (
	"context"
	"encoding/json"
)

// conformanceCalls calls each operation with its arguments decoded from the case.
var conformanceCalls = []struct {
	id   string
	call func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error)
}{
	{"abTests.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CreateAbTestInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AbTests.Create(ctx, bodyArg)
	}},
	{"abTests.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AbTests.Delete(ctx, a[0])
	}},
	{"abTests.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AbTests.Get(ctx, a[0])
	}},
	{"abTests.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AbTestsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.AbTests.List(ctx, queryArg)
	}},
	{"abTests.promote", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AbTestsPromoteBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AbTests.Promote(ctx, a[0], bodyArg)
	}},
	{"abTests.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UpdateAbTestInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AbTests.Update(ctx, a[0], bodyArg)
	}},
	{"ads.connections.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AdsConnectionsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Ads.Connections.Create(ctx, bodyArg)
	}},
	{"ads.connections.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Ads.Connections.Delete(ctx, a[0])
	}},
	{"ads.connections.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AdsConnectionsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Ads.Connections.List(ctx, queryArg)
	}},
	{"ads.connections.test", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Ads.Connections.Test(ctx, a[0])
	}},
	{"ads.connections.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AdsConnectionsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Ads.Connections.Update(ctx, a[0], bodyArg)
	}},
	{"agents.activity", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Agents.Activity(ctx, a[0])
	}},
	{"agents.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AgentsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Agents.Create(ctx, bodyArg)
	}},
	{"agents.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Agents.Delete(ctx, a[0])
	}},
	{"agents.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AgentsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Agents.List(ctx, queryArg)
	}},
	{"agents.queue", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AgentsQueueQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Agents.Queue(ctx, queryArg)
	}},
	{"agents.tools", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AgentsToolsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Agents.Tools(ctx, queryArg)
	}},
	{"agents.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AgentsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Agents.Update(ctx, a[0], bodyArg)
	}},
	{"agents.groups.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AgentsGroupsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Agents.Groups.Create(ctx, bodyArg)
	}},
	{"agents.groups.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Agents.Groups.Delete(ctx, a[0])
	}},
	{"agents.groups.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AgentsGroupsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Agents.Groups.List(ctx, queryArg)
	}},
	{"agents.groups.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AgentsGroupsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Agents.Groups.Update(ctx, a[0], bodyArg)
	}},
	{"agents.memory.add", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AgentsMemoryAddBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Agents.Memory.Add(ctx, bodyArg)
	}},
	{"agents.memory.forget", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AgentsMemoryForgetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Agents.Memory.Forget(ctx, queryArg)
	}},
	{"agents.memory.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AgentsMemoryListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Agents.Memory.List(ctx, queryArg)
	}},
	{"agents.questions.answer", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AgentsQuestionsAnswerBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Agents.Questions.Answer(ctx, a[0], bodyArg)
	}},
	{"agents.questions.decline", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Agents.Questions.Decline(ctx, a[0])
	}},
	{"agents.questions.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AgentsQuestionsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Agents.Questions.List(ctx, queryArg)
	}},
	{"agents.runs.cancel", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Agents.Runs.Cancel(ctx, a[0])
	}},
	{"agents.runs.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Agents.Runs.Delete(ctx, a[0])
	}},
	{"agents.runs.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Agents.Runs.Get(ctx, a[0])
	}},
	{"agents.runs.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AgentsRunsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Agents.Runs.List(ctx, queryArg)
	}},
	{"agents.runs.retry", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Agents.Runs.Retry(ctx, a[0])
	}},
	{"agents.runs.start", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AgentsRunsStartBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Agents.Runs.Start(ctx, bodyArg)
	}},
	{"agents.schedules.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AgentsSchedulesCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Agents.Schedules.Create(ctx, bodyArg)
	}},
	{"agents.schedules.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Agents.Schedules.Delete(ctx, a[0])
	}},
	{"agents.schedules.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AgentsSchedulesListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Agents.Schedules.List(ctx, queryArg)
	}},
	{"agents.schedules.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AgentsSchedulesUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Agents.Schedules.Update(ctx, a[0], bodyArg)
	}},
	{"ai.analyzeImage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiAnalyzeImageBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.AnalyzeImage(ctx, bodyArg)
	}},
	{"ai.analyzeUrl", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiAnalyzeUrlBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.AnalyzeURL(ctx, bodyArg)
	}},
	{"ai.bulkOperation", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiBulkOperationBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.BulkOperation(ctx, bodyArg)
	}},
	{"ai.checkAccessibility", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiCheckAccessibilityBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.CheckAccessibility(ctx, bodyArg)
	}},
	{"ai.componentAction", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiComponentActionBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.ComponentAction(ctx, bodyArg)
	}},
	{"ai.generateContent", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiGenerateContentBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.GenerateContent(ctx, bodyArg)
	}},
	{"ai.generateImage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiGenerateImageBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.GenerateImage(ctx, bodyArg)
	}},
	{"ai.generateProject", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiGenerateProjectBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.GenerateProject(ctx, bodyArg)
	}},
	{"ai.generateSeo", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiGenerateSeoBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.GenerateSEO(ctx, bodyArg)
	}},
	{"ai.getTimeSaved", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AiGetTimeSavedQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.AI.GetTimeSaved(ctx, queryArg)
	}},
	{"ai.getUsage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AiGetUsageQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.AI.GetUsage(ctx, queryArg)
	}},
	{"ai.queryDatabases", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiQueryDatabasesBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.QueryDatabases(ctx, bodyArg)
	}},
	{"ai.rewriteDocumentText", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiRewriteDocumentTextBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.RewriteDocumentText(ctx, bodyArg)
	}},
	{"ai.subtitles", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiSubtitlesBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.Subtitles(ctx, bodyArg)
	}},
	{"ai.chats.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiChatsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.Chats.Create(ctx, bodyArg)
	}},
	{"ai.chats.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AI.Chats.Delete(ctx, a[0])
	}},
	{"ai.chats.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AI.Chats.Get(ctx, a[0])
	}},
	{"ai.chats.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AiChatsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.AI.Chats.List(ctx, queryArg)
	}},
	{"ai.chats.rename", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiChatsRenameBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.Chats.Rename(ctx, a[0], bodyArg)
	}},
	{"ai.marking.exportActivity", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AiMarkingExportActivityQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.AI.Marking.ExportActivity(ctx, queryArg)
	}},
	{"ai.marking.getPolicy", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AiMarkingGetPolicyQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.AI.Marking.GetPolicy(ctx, queryArg)
	}},
	{"ai.marking.listActivity", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AiMarkingListActivityQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.AI.Marking.ListActivity(ctx, queryArg)
	}},
	{"ai.marking.setPolicy", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiMarkingSetPolicyBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.Marking.SetPolicy(ctx, bodyArg)
	}},
	{"ai.personas.clear", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AI.Personas.Clear(ctx)
	}},
	{"ai.personas.clearOrganization", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AI.Personas.ClearOrganization(ctx, a[0])
	}},
	{"ai.personas.clearWorkspace", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AI.Personas.ClearWorkspace(ctx, a[0])
	}},
	{"ai.personas.deriveFromSamples", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiPersonasDeriveFromSamplesBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.Personas.DeriveFromSamples(ctx, bodyArg)
	}},
	{"ai.personas.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AI.Personas.Get(ctx)
	}},
	{"ai.personas.getOrganization", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AI.Personas.GetOrganization(ctx, a[0])
	}},
	{"ai.personas.getWorkspace", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AI.Personas.GetWorkspace(ctx, a[0])
	}},
	{"ai.personas.preview", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiPersonasPreviewBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.Personas.Preview(ctx, bodyArg)
	}},
	{"ai.personas.set", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AIPersonaInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.Personas.Set(ctx, bodyArg)
	}},
	{"ai.personas.setOrganization", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AISharedPersonaInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.Personas.SetOrganization(ctx, a[0], bodyArg)
	}},
	{"ai.personas.setWorkspace", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AISharedPersonaInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.Personas.SetWorkspace(ctx, a[0], bodyArg)
	}},
	{"ai.receipts.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AiReceiptsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.AI.Receipts.Get(ctx, a[0], queryArg)
	}},
	{"ai.receipts.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AiReceiptsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.AI.Receipts.List(ctx, queryArg)
	}},
	{"ai.receipts.share", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AiReceiptsShareQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.AI.Receipts.Share(ctx, a[0], queryArg)
	}},
	{"ai.receipts.unshare", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AiReceiptsUnshareQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.AI.Receipts.Unshare(ctx, a[0], queryArg)
	}},
	{"ai.skills.addTemplate", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiSkillsAddTemplateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.Skills.AddTemplate(ctx, bodyArg)
	}},
	{"ai.skills.approve", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiSkillsApproveBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.Skills.Approve(ctx, a[0], bodyArg)
	}},
	{"ai.skills.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AICreateSkillInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.Skills.Create(ctx, bodyArg)
	}},
	{"ai.skills.createOrganization", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AICreateSharedSkillInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.Skills.CreateOrganization(ctx, a[0], bodyArg)
	}},
	{"ai.skills.createWorkspace", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AICreateSharedSkillInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.Skills.CreateWorkspace(ctx, a[0], bodyArg)
	}},
	{"ai.skills.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AI.Skills.Delete(ctx, a[0])
	}},
	{"ai.skills.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AI.Skills.Get(ctx, a[0])
	}},
	{"ai.skills.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AI.Skills.List(ctx)
	}},
	{"ai.skills.listOrganization", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AI.Skills.ListOrganization(ctx, a[0])
	}},
	{"ai.skills.listTemplates", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AiSkillsListTemplatesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.AI.Skills.ListTemplates(ctx, queryArg)
	}},
	{"ai.skills.listTools", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AiSkillsListToolsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.AI.Skills.ListTools(ctx, queryArg)
	}},
	{"ai.skills.listWorkspace", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AI.Skills.ListWorkspace(ctx, a[0])
	}},
	{"ai.skills.setOverride", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiSkillsSetOverrideBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.Skills.SetOverride(ctx, a[0], bodyArg)
	}},
	{"ai.skills.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AIUpdateSkillInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.Skills.Update(ctx, a[0], bodyArg)
	}},
	{"ai.video.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AiVideoCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AI.Video.Create(ctx, bodyArg)
	}},
	{"ai.video.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AI.Video.Get(ctx, a[0])
	}},
	{"analytics.aggregate", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AnalyticsAggregateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Analytics.Aggregate(ctx, queryArg)
	}},
	{"analytics.heatmap", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AnalyticsHeatmapQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Analytics.Heatmap(ctx, queryArg)
	}},
	{"analytics.track", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AnalyticsEvent
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Analytics.Track(ctx, bodyArg)
	}},
	{"announcements.dismiss", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Announcements.Dismiss(ctx, a[0])
	}},
	{"announcements.listActive", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Announcements.ListActive(ctx)
	}},
	{"approvals.approve", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ApprovalsApproveBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Approvals.Approve(ctx, a[0], bodyArg)
	}},
	{"approvals.bulk", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ApprovalsBulkBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Approvals.Bulk(ctx, bodyArg)
	}},
	{"approvals.cancel", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Approvals.Cancel(ctx, a[0])
	}},
	{"approvals.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CreateApprovalInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Approvals.Create(ctx, bodyArg)
	}},
	{"approvals.createDelegation", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ApprovalsCreateDelegationBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Approvals.CreateDelegation(ctx, bodyArg)
	}},
	{"approvals.deleteDelegation", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ApprovalsDeleteDelegationQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Approvals.DeleteDelegation(ctx, queryArg)
	}},
	{"approvals.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Approvals.Get(ctx, a[0])
	}},
	{"approvals.getSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ApprovalsGetSettingsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Approvals.GetSettings(ctx, queryArg)
	}},
	{"approvals.history", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Approvals.History(ctx, a[0])
	}},
	{"approvals.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ApprovalsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Approvals.List(ctx, queryArg)
	}},
	{"approvals.listDelegations", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Approvals.ListDelegations(ctx)
	}},
	{"approvals.reject", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ApprovalsRejectBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Approvals.Reject(ctx, a[0], bodyArg)
	}},
	{"approvals.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ApprovalsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Approvals.Update(ctx, a[0], bodyArg)
	}},
	{"approvals.updateSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ApprovalsUpdateSettingsBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Approvals.UpdateSettings(ctx, bodyArg)
	}},
	{"apps.builds.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AppsBuildsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Apps.Builds.Create(ctx, a[0], bodyArg)
	}},
	{"apps.builds.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AppsBuildsDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Apps.Builds.Delete(ctx, a[0], queryArg)
	}},
	{"apps.builds.download", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AppsBuildsDownloadQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Apps.Builds.Download(ctx, a[0], queryArg)
	}},
	{"apps.builds.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AppsBuildsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Apps.Builds.List(ctx, a[0], queryArg)
	}},
	{"apps.builds.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AppsBuildsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Apps.Builds.Update(ctx, a[0], bodyArg)
	}},
	{"apps.publications.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AppsPublicationsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Apps.Publications.Create(ctx, bodyArg)
	}},
	{"apps.publications.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AppsPublicationsDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Apps.Publications.Delete(ctx, a[0], queryArg)
	}},
	{"apps.publications.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AppsPublicationsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Apps.Publications.Get(ctx, a[0], queryArg)
	}},
	{"apps.publications.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AppsPublicationsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Apps.Publications.List(ctx, queryArg)
	}},
	{"apps.publications.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AppsPublicationsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Apps.Publications.Update(ctx, a[0], bodyArg)
	}},
	{"apps.signingCredentials.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AppsSigningCredentialsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Apps.SigningCredentials.Create(ctx, bodyArg)
	}},
	{"apps.signingCredentials.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AppsSigningCredentialsDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Apps.SigningCredentials.Delete(ctx, a[0], queryArg)
	}},
	{"apps.signingCredentials.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AppsSigningCredentialsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Apps.SigningCredentials.List(ctx, queryArg)
	}},
	{"assetFolders.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CreateAssetFolderInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AssetFolders.Create(ctx, bodyArg)
	}},
	{"assetFolders.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AssetFolders.Delete(ctx, a[0])
	}},
	{"assetFolders.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AssetFoldersListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.AssetFolders.List(ctx, queryArg)
	}},
	{"assetFolders.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UpdateAssetFolderInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AssetFolders.Update(ctx, a[0], bodyArg)
	}},
	{"assets.bulk", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AssetsBulkBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Assets.Bulk(ctx, bodyArg)
	}},
	{"assets.createText", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AssetsCreateTextBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Assets.CreateText(ctx, bodyArg)
	}},
	{"assets.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Assets.Delete(ctx, a[0])
	}},
	{"assets.emptyTrash", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AssetsEmptyTrashQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Assets.EmptyTrash(ctx, queryArg)
	}},
	{"assets.generateImage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AssetsGenerateImageBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Assets.GenerateImage(ctx, bodyArg)
	}},
	{"assets.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Assets.Get(ctx, a[0])
	}},
	{"assets.intelligence", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AssetsIntelligenceQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Assets.Intelligence(ctx, queryArg)
	}},
	{"assets.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AssetsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Assets.List(ctx, queryArg)
	}},
	{"assets.listTrash", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AssetsListTrashQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Assets.ListTrash(ctx, queryArg)
	}},
	{"assets.storageUsage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AssetsStorageUsageQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Assets.StorageUsage(ctx, queryArg)
	}},
	{"assets.sync", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AssetsSyncQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Assets.Sync(ctx, queryArg)
	}},
	{"assets.transcribe", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AssetsTranscribeBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Assets.Transcribe(ctx, bodyArg)
	}},
	{"assets.trashAction", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AssetsTrashActionBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Assets.TrashAction(ctx, bodyArg)
	}},
	{"assets.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UpdateAssetInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Assets.Update(ctx, a[0], bodyArg)
	}},
	{"assets.upload", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AssetsUploadBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Assets.Upload(ctx, bodyArg)
	}},
	{"assets.usage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AssetsUsageQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Assets.Usage(ctx, a[0], queryArg)
	}},
	{"assets.zip", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AssetsZipBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Assets.Zip(ctx, bodyArg)
	}},
	{"assets.versions.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AssetsVersionsDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Assets.Versions.Delete(ctx, a[0], a[1], queryArg)
	}},
	{"assets.versions.diff", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AssetsVersionsDiffBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Assets.Versions.Diff(ctx, a[0], bodyArg)
	}},
	{"assets.versions.download", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AssetsVersionsDownloadQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Assets.Versions.Download(ctx, a[0], a[1], queryArg)
	}},
	{"assets.versions.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AssetsVersionsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Assets.Versions.Get(ctx, a[0], a[1], queryArg)
	}},
	{"assets.versions.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AssetsVersionsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Assets.Versions.List(ctx, a[0], queryArg)
	}},
	{"assets.versions.restore", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AssetsVersionsRestoreBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Assets.Versions.Restore(ctx, a[0], a[1], bodyArg)
	}},
	{"assets.versions.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AssetsVersionsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Assets.Versions.Update(ctx, a[0], a[1], bodyArg)
	}},
	{"assetShares.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CreateAssetShareInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AssetShares.Create(ctx, a[0], bodyArg)
	}},
	{"assetShares.createFolder", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CreateAssetShareInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AssetShares.CreateFolder(ctx, a[0], bodyArg)
	}},
	{"assetShares.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AssetShares.Delete(ctx, a[0], a[1])
	}},
	{"assetShares.deleteById", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AssetShares.DeleteByID(ctx, a[0])
	}},
	{"assetShares.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AssetShares.List(ctx, a[0])
	}},
	{"assetShares.listFolder", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.AssetShares.ListFolder(ctx, a[0])
	}},
	{"assetShares.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UpdateAssetShareInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AssetShares.Update(ctx, a[0], a[1], bodyArg)
	}},
	{"assetShares.updateById", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UpdateAssetShareInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.AssetShares.UpdateByID(ctx, a[0], bodyArg)
	}},
	{"auditLogs.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *AuditLogsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.AuditLogs.List(ctx, queryArg)
	}},
	{"banking.forecast", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingForecastQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Forecast(ctx, queryArg)
	}},
	{"banking.importPreview", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BankingImportPreviewBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingImportPreviewQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.ImportPreview(ctx, bodyArg, queryArg)
	}},
	{"banking.reports", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingReportsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Reports(ctx, queryArg)
	}},
	{"banking.accounts.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BankingAccountsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingAccountsCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Accounts.Create(ctx, bodyArg, queryArg)
	}},
	{"banking.accounts.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingAccountsDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Accounts.Delete(ctx, a[0], queryArg)
	}},
	{"banking.accounts.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingAccountsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Accounts.Get(ctx, a[0], queryArg)
	}},
	{"banking.accounts.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingAccountsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Accounts.List(ctx, queryArg)
	}},
	{"banking.accounts.reconciliation", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingAccountsReconciliationQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Accounts.Reconciliation(ctx, a[0], queryArg)
	}},
	{"banking.accounts.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BankingAccountsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingAccountsUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Accounts.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"banking.accounts.grants.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingAccountsGrantsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Accounts.Grants.List(ctx, a[0], queryArg)
	}},
	{"banking.accounts.grants.remove", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingAccountsGrantsRemoveQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Accounts.Grants.Remove(ctx, a[0], a[1], queryArg)
	}},
	{"banking.accounts.grants.set", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BankingAccountsGrantsSetBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingAccountsGrantsSetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Accounts.Grants.Set(ctx, a[0], bodyArg, queryArg)
	}},
	{"banking.accounts.imports.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingAccountsImportsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Accounts.Imports.List(ctx, a[0], queryArg)
	}},
	{"banking.accounts.imports.upload", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BankingAccountsImportsUploadBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingAccountsImportsUploadQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Accounts.Imports.Upload(ctx, a[0], bodyArg, queryArg)
	}},
	{"banking.accounts.ledger.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingAccountsLedgerGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Accounts.Ledger.Get(ctx, a[0], queryArg)
	}},
	{"banking.accounts.ledger.remove", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingAccountsLedgerRemoveQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Accounts.Ledger.Remove(ctx, a[0], queryArg)
	}},
	{"banking.accounts.ledger.set", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BankingAccountsLedgerSetBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingAccountsLedgerSetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Accounts.Ledger.Set(ctx, a[0], bodyArg, queryArg)
	}},
	{"banking.anomalies.review", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg any
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingAnomaliesReviewQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Anomalies.Review(ctx, a[0], bodyArg, queryArg)
	}},
	{"banking.expected.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingExpectedListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Expected.List(ctx, queryArg)
	}},
	{"banking.expected.record", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingExpectedRecordQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Expected.Record(ctx, queryArg)
	}},
	{"banking.expected.review", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg any
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingExpectedReviewQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Expected.Review(ctx, a[0], bodyArg, queryArg)
	}},
	{"banking.expected.plans.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BankPaymentPlanInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingExpectedPlansCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Expected.Plans.Create(ctx, bodyArg, queryArg)
	}},
	{"banking.expected.plans.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingExpectedPlansDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Expected.Plans.Delete(ctx, a[0], queryArg)
	}},
	{"banking.expected.plans.generate", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingExpectedPlansGenerateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Expected.Plans.Generate(ctx, a[0], queryArg)
	}},
	{"banking.expected.plans.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingExpectedPlansGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Expected.Plans.Get(ctx, a[0], queryArg)
	}},
	{"banking.expected.plans.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingExpectedPlansListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Expected.Plans.List(ctx, queryArg)
	}},
	{"banking.expected.plans.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BankingExpectedPlansUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingExpectedPlansUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Expected.Plans.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"banking.matches.undo", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingMatchesUndoQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Matches.Undo(ctx, a[0], queryArg)
	}},
	{"banking.matching.count", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingMatchingCountQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Matching.Count(ctx, queryArg)
	}},
	{"banking.matching.queue", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingMatchingQueueQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Matching.Queue(ctx, queryArg)
	}},
	{"banking.matching.run", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BankingMatchingRunBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingMatchingRunQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Matching.Run(ctx, bodyArg, queryArg)
	}},
	{"banking.payments.act", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg any
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingPaymentsActQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Payments.Act(ctx, a[0], bodyArg, queryArg)
	}},
	{"banking.payments.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BankingPaymentsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingPaymentsCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Payments.Create(ctx, bodyArg, queryArg)
	}},
	{"banking.payments.fromPurchases", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BankingPaymentsFromPurchasesBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingPaymentsFromPurchasesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Payments.FromPurchases(ctx, bodyArg, queryArg)
	}},
	{"banking.payments.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingPaymentsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Payments.Get(ctx, a[0], queryArg)
	}},
	{"banking.payments.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingPaymentsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Payments.List(ctx, queryArg)
	}},
	{"banking.payments.options", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingPaymentsOptionsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Payments.Options(ctx, queryArg)
	}},
	{"banking.payments.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BankingPaymentsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingPaymentsUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Payments.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"banking.payments.settings.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingPaymentsSettingsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Payments.Settings.Get(ctx, queryArg)
	}},
	{"banking.payments.settings.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BankingPaymentsSettingsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingPaymentsSettingsUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Payments.Settings.Update(ctx, bodyArg, queryArg)
	}},
	{"banking.receipts.book", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingReceiptsBookQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Receipts.Book(ctx, a[0], queryArg)
	}},
	{"banking.receipts.decide", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BankingReceiptsDecideBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingReceiptsDecideQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Receipts.Decide(ctx, a[0], bodyArg, queryArg)
	}},
	{"banking.receipts.file", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingReceiptsFileQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Receipts.File(ctx, a[0], queryArg)
	}},
	{"banking.receipts.missing", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingReceiptsMissingQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Receipts.Missing(ctx, queryArg)
	}},
	{"banking.receipts.run", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingReceiptsRunQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Receipts.Run(ctx, queryArg)
	}},
	{"banking.receipts.unlink", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingReceiptsUnlinkQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Receipts.Unlink(ctx, a[0], queryArg)
	}},
	{"banking.receipts.settings.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingReceiptsSettingsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Receipts.Settings.Get(ctx, queryArg)
	}},
	{"banking.receipts.settings.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BankingReceiptsSettingsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingReceiptsSettingsUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Receipts.Settings.Update(ctx, bodyArg, queryArg)
	}},
	{"banking.rules.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BankingRulesCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingRulesCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Rules.Create(ctx, bodyArg, queryArg)
	}},
	{"banking.rules.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingRulesDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Rules.Delete(ctx, a[0], queryArg)
	}},
	{"banking.rules.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingRulesListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Rules.List(ctx, queryArg)
	}},
	{"banking.rules.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BankingRulesUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingRulesUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Rules.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"banking.settings.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingSettingsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Settings.Get(ctx, queryArg)
	}},
	{"banking.settings.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BankingSettingsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingSettingsUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Settings.Update(ctx, bodyArg, queryArg)
	}},
	{"banking.transactions.addReceipt", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg any
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingTransactionsAddReceiptQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Transactions.AddReceipt(ctx, a[0], bodyArg, queryArg)
	}},
	{"banking.transactions.annotate", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BankingTransactionsAnnotateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingTransactionsAnnotateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Transactions.Annotate(ctx, a[0], bodyArg, queryArg)
	}},
	{"banking.transactions.export", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingTransactionsExportQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Transactions.Export(ctx, queryArg)
	}},
	{"banking.transactions.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingTransactionsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Transactions.Get(ctx, a[0], queryArg)
	}},
	{"banking.transactions.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingTransactionsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Transactions.List(ctx, queryArg)
	}},
	{"banking.transactions.match", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg any
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BankingTransactionsMatchQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Transactions.Match(ctx, a[0], bodyArg, queryArg)
	}},
	{"banking.transactions.matchCandidates", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingTransactionsMatchCandidatesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Transactions.MatchCandidates(ctx, a[0], queryArg)
	}},
	{"banking.transactions.matches", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingTransactionsMatchesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Transactions.Matches(ctx, a[0], queryArg)
	}},
	{"banking.transactions.receiptCandidates", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingTransactionsReceiptCandidatesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Transactions.ReceiptCandidates(ctx, a[0], queryArg)
	}},
	{"banking.transactions.receipts", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingTransactionsReceiptsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Transactions.Receipts(ctx, a[0], queryArg)
	}},
	{"banking.transactions.suggest", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BankingTransactionsSuggestQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Banking.Transactions.Suggest(ctx, a[0], queryArg)
	}},
	{"billing.documentPdf", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BillingDocumentPdfQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Billing.DocumentPDF(ctx, a[0], queryArg)
	}},
	{"billing.documentXml", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Billing.DocumentXml(ctx, a[0])
	}},
	{"billing.explainDocument", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Billing.ExplainDocument(ctx, a[0])
	}},
	{"billing.getAddress", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BillingGetAddressQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Billing.GetAddress(ctx, queryArg)
	}},
	{"billing.getStatus", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BillingGetStatusQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Billing.GetStatus(ctx, queryArg)
	}},
	{"billing.listDocuments", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BillingListDocumentsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Billing.ListDocuments(ctx, queryArg)
	}},
	{"billing.listHistory", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BillingListHistoryQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Billing.ListHistory(ctx, queryArg)
	}},
	{"billing.listMarketplaceInvoices", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BillingListMarketplaceInvoicesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Billing.ListMarketplaceInvoices(ctx, queryArg)
	}},
	{"billing.marketplaceInvoicePdf", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Billing.MarketplaceInvoicePDF(ctx, a[0])
	}},
	{"billing.updateAddress", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UpdateBillingAddressInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Billing.UpdateAddress(ctx, bodyArg)
	}},
	{"billing.usage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BillingUsageQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Billing.Usage(ctx, queryArg)
	}},
	{"builderBlocks.catalog", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BuilderBlocksCatalogQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.BuilderBlocks.Catalog(ctx, queryArg)
	}},
	{"builderBlocks.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BuilderBlocksCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BuilderBlocksCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.BuilderBlocks.Create(ctx, bodyArg, queryArg)
	}},
	{"builderBlocks.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BuilderBlocksDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.BuilderBlocks.Delete(ctx, a[0], queryArg)
	}},
	{"builderBlocks.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BuilderBlocksListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.BuilderBlocks.List(ctx, queryArg)
	}},
	{"builderBlocks.offerInMarketplace", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BuilderBlocksOfferInMarketplaceBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BuilderBlocksOfferInMarketplaceQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.BuilderBlocks.OfferInMarketplace(ctx, a[0], bodyArg, queryArg)
	}},
	{"builderBlocks.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BuilderBlocksUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BuilderBlocksUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.BuilderBlocks.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"builderBlocks.usage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BuilderBlocksUsageQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.BuilderBlocks.Usage(ctx, a[0], queryArg)
	}},
	{"builderBlocks.versions.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *BuilderBlocksVersionsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *BuilderBlocksVersionsCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.BuilderBlocks.Versions.Create(ctx, a[0], bodyArg, queryArg)
	}},
	{"builderBlocks.versions.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *BuilderBlocksVersionsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.BuilderBlocks.Versions.List(ctx, a[0], queryArg)
	}},
	{"calendar.addAttendee", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CalendarAddAttendeeBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Calendar.AddAttendee(ctx, a[0], bodyArg)
	}},
	{"calendar.aiCommand", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CalendarAiCommandBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Calendar.AICommand(ctx, bodyArg)
	}},
	{"calendar.cancelBooking", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CalendarCancelBookingBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Calendar.CancelBooking(ctx, a[0], bodyArg)
	}},
	{"calendar.checkConflicts", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CalendarCheckConflictsBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Calendar.CheckConflicts(ctx, bodyArg)
	}},
	{"calendar.confirmBooking", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.ConfirmBooking(ctx, a[0])
	}},
	{"calendar.createBookingPage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CalendarCreateBookingPageBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Calendar.CreateBookingPage(ctx, bodyArg)
	}},
	{"calendar.createCalendar", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CalendarCreateCalendarBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Calendar.CreateCalendar(ctx, bodyArg)
	}},
	{"calendar.createEvent", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CalendarCreateEventBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Calendar.CreateEvent(ctx, a[0], bodyArg)
	}},
	{"calendar.createEventConference", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.CreateEventConference(ctx, a[0])
	}},
	{"calendar.createProposal", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CalendarCreateProposalBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Calendar.CreateProposal(ctx, bodyArg)
	}},
	{"calendar.createShare", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CalendarCreateShareBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Calendar.CreateShare(ctx, a[0], bodyArg)
	}},
	{"calendar.decideSpamBooking", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CalendarDecideSpamBookingBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Calendar.DecideSpamBooking(ctx, a[0], bodyArg)
	}},
	{"calendar.deleteBooking", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.DeleteBooking(ctx, a[0])
	}},
	{"calendar.deleteBookingPage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.DeleteBookingPage(ctx, a[0])
	}},
	{"calendar.deleteCalendar", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CalendarDeleteCalendarQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Calendar.DeleteCalendar(ctx, a[0], queryArg)
	}},
	{"calendar.deleteEvent", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.DeleteEvent(ctx, a[0])
	}},
	{"calendar.deleteShare", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.DeleteShare(ctx, a[0], a[1])
	}},
	{"calendar.disableCaldavSync", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.DisableCaldavSync(ctx, a[0])
	}},
	{"calendar.disconnectIntegration", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.DisconnectIntegration(ctx, a[0])
	}},
	{"calendar.enableCaldavSync", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.EnableCaldavSync(ctx, a[0])
	}},
	{"calendar.eventBriefing", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.EventBriefing(ctx, a[0])
	}},
	{"calendar.eventMeetingNeed", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.EventMeetingNeed(ctx, a[0])
	}},
	{"calendar.eventOutcome", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg any
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Calendar.EventOutcome(ctx, a[0], bodyArg)
	}},
	{"calendar.eventRecurrence", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CalendarEventRecurrenceBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Calendar.EventRecurrence(ctx, a[0], bodyArg)
	}},
	{"calendar.finalizeProposal", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CalendarFinalizeProposalBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *CalendarFinalizeProposalQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Calendar.FinalizeProposal(ctx, a[0], bodyArg, queryArg)
	}},
	{"calendar.getBooking", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.GetBooking(ctx, a[0])
	}},
	{"calendar.getBookingPage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.GetBookingPage(ctx, a[0])
	}},
	{"calendar.getCaldavSync", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.GetCaldavSync(ctx, a[0])
	}},
	{"calendar.getCalendar", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.GetCalendar(ctx, a[0])
	}},
	{"calendar.getCalendarFeedToken", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.GetCalendarFeedToken(ctx, a[0])
	}},
	{"calendar.getEvent", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.GetEvent(ctx, a[0])
	}},
	{"calendar.getEventIcs", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.GetEventIcs(ctx, a[0])
	}},
	{"calendar.getEventOutcome", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.GetEventOutcome(ctx, a[0])
	}},
	{"calendar.getSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.GetSettings(ctx)
	}},
	{"calendar.holidays", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CalendarHolidaysQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Calendar.Holidays(ctx, queryArg)
	}},
	{"calendar.listAttendees", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.ListAttendees(ctx, a[0])
	}},
	{"calendar.listBookingPages", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.ListBookingPages(ctx)
	}},
	{"calendar.listBookings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CalendarListBookingsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Calendar.ListBookings(ctx, queryArg)
	}},
	{"calendar.listCalendars", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.ListCalendars(ctx)
	}},
	{"calendar.listEvents", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CalendarListEventsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Calendar.ListEvents(ctx, a[0], queryArg)
	}},
	{"calendar.listIntegrationCalendars", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.ListIntegrationCalendars(ctx, a[0])
	}},
	{"calendar.listIntegrations", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.ListIntegrations(ctx)
	}},
	{"calendar.listProposals", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.ListProposals(ctx)
	}},
	{"calendar.listShares", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.ListShares(ctx, a[0])
	}},
	{"calendar.listWaitlist", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CalendarListWaitlistQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Calendar.ListWaitlist(ctx, queryArg)
	}},
	{"calendar.pullCaldavSync", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.PullCaldavSync(ctx, a[0])
	}},
	{"calendar.refresh", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.Refresh(ctx)
	}},
	{"calendar.removeAttendee", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CalendarRemoveAttendeeQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Calendar.RemoveAttendee(ctx, a[0], queryArg)
	}},
	{"calendar.removeWaitlistEntry", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.RemoveWaitlistEntry(ctx, a[0])
	}},
	{"calendar.rescheduleBooking", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CalendarRescheduleBookingBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Calendar.RescheduleBooking(ctx, a[0], bodyArg)
	}},
	{"calendar.revokeCalendarFeed", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.RevokeCalendarFeed(ctx, a[0])
	}},
	{"calendar.rotateCalendarFeed", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.RotateCalendarFeed(ctx, a[0])
	}},
	{"calendar.search", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CalendarSearchQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Calendar.Search(ctx, queryArg)
	}},
	{"calendar.setIntegrationCalendars", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CalendarSetIntegrationCalendarsBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Calendar.SetIntegrationCalendars(ctx, a[0], bodyArg)
	}},
	{"calendar.syncIntegration", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Calendar.SyncIntegration(ctx, a[0])
	}},
	{"calendar.travelEstimate", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CalendarTravelEstimateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Calendar.TravelEstimate(ctx, bodyArg)
	}},
	{"calendar.updateBookingPage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CalendarUpdateBookingPageBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Calendar.UpdateBookingPage(ctx, a[0], bodyArg)
	}},
	{"calendar.updateCalendar", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CalendarUpdateCalendarBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Calendar.UpdateCalendar(ctx, a[0], bodyArg)
	}},
	{"calendar.updateEvent", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CalendarUpdateEventBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Calendar.UpdateEvent(ctx, a[0], bodyArg)
	}},
	{"calendar.updateSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CalendarUpdateSettingsBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Calendar.UpdateSettings(ctx, bodyArg)
	}},
	{"calendar.updateShare", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CalendarUpdateShareBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Calendar.UpdateShare(ctx, a[0], a[1], bodyArg)
	}},
	{"code.storage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CodeStorageQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Code.Storage(ctx, queryArg)
	}},
	{"code.assets.getText", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Code.Assets.GetText(ctx, a[0])
	}},
	{"code.assets.saveText", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CodeAssetsSaveTextBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Code.Assets.SaveText(ctx, a[0], bodyArg)
	}},
	{"code.data.bind", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CodeDataBindBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Code.Data.Bind(ctx, a[0], bodyArg)
	}},
	{"code.data.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Code.Data.List(ctx, a[0])
	}},
	{"code.data.unbind", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CodeDataUnbindQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Code.Data.Unbind(ctx, a[0], queryArg)
	}},
	{"code.data.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CodeDataUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Code.Data.Update(ctx, a[0], bodyArg)
	}},
	{"code.deployments.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CodeDeploymentsDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Code.Deployments.Delete(ctx, a[0], queryArg)
	}},
	{"code.deployments.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Code.Deployments.List(ctx, a[0])
	}},
	{"code.deployments.logs", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CodeDeploymentsLogsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Code.Deployments.Logs(ctx, a[0], queryArg)
	}},
	{"code.deployments.publish", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CodeDeploymentsPublishBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Code.Deployments.Publish(ctx, a[0], bodyArg)
	}},
	{"code.deployments.removeDomain", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CodeDeploymentsRemoveDomainQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Code.Deployments.RemoveDomain(ctx, a[0], queryArg)
	}},
	{"code.deployments.setDomain", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CodeDeploymentsSetDomainBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *CodeDeploymentsSetDomainQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Code.Deployments.SetDomain(ctx, a[0], bodyArg, queryArg)
	}},
	{"code.deployments.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CodeDeploymentsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Code.Deployments.Update(ctx, a[0], bodyArg)
	}},
	{"code.deployments.verifyDomain", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CodeDeploymentsVerifyDomainQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Code.Deployments.VerifyDomain(ctx, a[0], queryArg)
	}},
	{"code.env.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CodeEnvDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Code.Env.Delete(ctx, a[0], queryArg)
	}},
	{"code.env.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Code.Env.List(ctx, a[0])
	}},
	{"code.env.set", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CodeEnvSetBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Code.Env.Set(ctx, a[0], bodyArg)
	}},
	{"code.files.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CodeFilesCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Code.Files.Create(ctx, a[0], bodyArg)
	}},
	{"code.files.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CodeFilesDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Code.Files.Delete(ctx, a[0], queryArg)
	}},
	{"code.files.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CodeFilesGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Code.Files.Get(ctx, a[0], queryArg)
	}},
	{"code.files.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CodeFilesUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Code.Files.Update(ctx, a[0], bodyArg)
	}},
	{"code.git.connect", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CodeGitConnectBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Code.Git.Connect(ctx, a[0], bodyArg)
	}},
	{"code.git.history", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CodeGitHistoryQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Code.Git.History(ctx, a[0], queryArg)
	}},
	{"code.git.repos", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CodeGitReposQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Code.Git.Repos(ctx, queryArg)
	}},
	{"code.git.status", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Code.Git.Status(ctx, a[0])
	}},
	{"code.git.sync", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CodeGitSyncBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Code.Git.Sync(ctx, a[0], bodyArg)
	}},
	{"code.git.credentials.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CodeGitCredentialsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Code.Git.Credentials.Create(ctx, bodyArg)
	}},
	{"code.git.credentials.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CodeGitCredentialsDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Code.Git.Credentials.Delete(ctx, queryArg)
	}},
	{"code.git.credentials.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CodeGitCredentialsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Code.Git.Credentials.List(ctx, queryArg)
	}},
	{"code.projects.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CodeProjectsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Code.Projects.Create(ctx, bodyArg)
	}},
	{"code.projects.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Code.Projects.Delete(ctx, a[0])
	}},
	{"code.projects.download", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Code.Projects.Download(ctx, a[0])
	}},
	{"code.projects.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Code.Projects.Get(ctx, a[0])
	}},
	{"code.projects.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CodeProjectsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Code.Projects.List(ctx, queryArg)
	}},
	{"code.projects.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CodeProjectsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Code.Projects.Update(ctx, a[0], bodyArg)
	}},
	{"community.acceptAnswer", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunityAcceptAnswerBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.AcceptAnswer(ctx, a[0], a[1], bodyArg)
	}},
	{"community.acceptInvite", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.AcceptInvite(ctx, a[0])
	}},
	{"community.aiSummarize", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.AISummarize(ctx, a[0], a[1])
	}},
	{"community.amaRecap", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.AmaRecap(ctx, a[0], a[1])
	}},
	{"community.approvePost", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunityApprovePostBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.ApprovePost(ctx, a[0], a[1], bodyArg)
	}},
	{"community.approveReply", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunityApproveReplyBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.ApproveReply(ctx, a[0], a[1], bodyArg)
	}},
	{"community.becomeMember", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunityBecomeMemberBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.BecomeMember(ctx, a[0], bodyArg)
	}},
	{"community.bookmarkPost", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.BookmarkPost(ctx, a[0], a[1])
	}},
	{"community.clearAcceptedAnswer", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.ClearAcceptedAnswer(ctx, a[0], a[1])
	}},
	{"community.createHub", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunityCreateHubBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.CreateHub(ctx, bodyArg)
	}},
	{"community.createInvite", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunityCreateInviteBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.CreateInvite(ctx, a[0], bodyArg)
	}},
	{"community.createPost", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunityCreatePostBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.CreatePost(ctx, a[0], bodyArg)
	}},
	{"community.createReply", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunityCreateReplyBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.CreateReply(ctx, a[0], a[1], bodyArg)
	}},
	{"community.createTag", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunityCreateTagBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.CreateTag(ctx, a[0], bodyArg)
	}},
	{"community.deleteDraft", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.DeleteDraft(ctx, a[0])
	}},
	{"community.deleteHub", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.DeleteHub(ctx, a[0])
	}},
	{"community.deletePost", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.DeletePost(ctx, a[0], a[1])
	}},
	{"community.deleteReply", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.DeleteReply(ctx, a[0], a[1], a[2])
	}},
	{"community.feed", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CommunityFeedQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Community.Feed(ctx, queryArg)
	}},
	{"community.getDraft", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.GetDraft(ctx, a[0])
	}},
	{"community.getHub", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.GetHub(ctx, a[0])
	}},
	{"community.getInvite", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.GetInvite(ctx, a[0])
	}},
	{"community.getMemberBadges", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.GetMemberBadges(ctx, a[0], a[1])
	}},
	{"community.getMemberProfile", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.GetMemberProfile(ctx, a[0], a[1])
	}},
	{"community.getPost", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.GetPost(ctx, a[0], a[1])
	}},
	{"community.join", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.Join(ctx, a[0])
	}},
	{"community.leaderboard", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CommunityLeaderboardQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Community.Leaderboard(ctx, a[0], queryArg)
	}},
	{"community.leave", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.Leave(ctx, a[0])
	}},
	{"community.listBookmarks", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.ListBookmarks(ctx)
	}},
	{"community.listHubs", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CommunityListHubsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Community.ListHubs(ctx, queryArg)
	}},
	{"community.listInvites", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.ListInvites(ctx, a[0])
	}},
	{"community.listJoinRequests", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.ListJoinRequests(ctx, a[0])
	}},
	{"community.listMembers", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CommunityListMembersQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Community.ListMembers(ctx, a[0], queryArg)
	}},
	{"community.listPosts", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CommunityListPostsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Community.ListPosts(ctx, a[0], queryArg)
	}},
	{"community.listReplies", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CommunityListRepliesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Community.ListReplies(ctx, a[0], a[1], queryArg)
	}},
	{"community.listReputationEvents", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CommunityListReputationEventsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Community.ListReputationEvents(ctx, a[0], a[1], queryArg)
	}},
	{"community.listTags", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.ListTags(ctx, a[0])
	}},
	{"community.lockPost", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.LockPost(ctx, a[0], a[1])
	}},
	{"community.moderation", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CommunityModerationQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Community.Moderation(ctx, queryArg)
	}},
	{"community.pinPost", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.PinPost(ctx, a[0], a[1])
	}},
	{"community.reactPost", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunityReactPostBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.ReactPost(ctx, a[0], a[1], bodyArg)
	}},
	{"community.reactReply", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunityReactReplyBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.ReactReply(ctx, a[0], a[1], a[2], bodyArg)
	}},
	{"community.removeMember", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.RemoveMember(ctx, a[0], a[1])
	}},
	{"community.report", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunityReportBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.Report(ctx, bodyArg)
	}},
	{"community.resolveJoinRequest", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunityResolveJoinRequestBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.ResolveJoinRequest(ctx, a[0], a[1], bodyArg)
	}},
	{"community.resolveReport", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunityResolveReportBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.ResolveReport(ctx, a[0], bodyArg)
	}},
	{"community.saveDraft", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunitySaveDraftBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.SaveDraft(ctx, a[0], bodyArg)
	}},
	{"community.search", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CommunitySearchQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Community.Search(ctx, queryArg)
	}},
	{"community.transferHub", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunityTransferHubBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.TransferHub(ctx, a[0], bodyArg)
	}},
	{"community.unbookmarkPost", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.UnbookmarkPost(ctx, a[0], a[1])
	}},
	{"community.updateHub", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunityUpdateHubBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.UpdateHub(ctx, a[0], bodyArg)
	}},
	{"community.updateMember", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunityUpdateMemberBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.UpdateMember(ctx, a[0], a[1], bodyArg)
	}},
	{"community.updatePost", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunityUpdatePostBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.UpdatePost(ctx, a[0], a[1], bodyArg)
	}},
	{"community.updateReply", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunityUpdateReplyBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.UpdateReply(ctx, a[0], a[1], a[2], bodyArg)
	}},
	{"community.upvotePost", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.UpvotePost(ctx, a[0], a[1])
	}},
	{"community.upvoteReply", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Community.UpvoteReply(ctx, a[0], a[1], a[2])
	}},
	{"community.votePoll", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CommunityVotePollBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Community.VotePoll(ctx, a[0], a[1], bodyArg)
	}},
	{"connections.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CreateConnectionInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Connections.Create(ctx, bodyArg)
	}},
	{"connections.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Connections.Delete(ctx, a[0])
	}},
	{"connections.dependents", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Connections.Dependents(ctx, a[0])
	}},
	{"connections.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Connections.Get(ctx, a[0])
	}},
	{"connections.getOAuthUrl", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ConnectionsGetOAuthUrlQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Connections.GetOAuthURL(ctx, a[0], queryArg)
	}},
	{"connections.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ConnectionsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Connections.List(ctx, queryArg)
	}},
	{"connections.listRequirements", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Connections.ListRequirements(ctx)
	}},
	{"connections.refresh", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Connections.Refresh(ctx, a[0])
	}},
	{"connections.test", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Connections.Test(ctx, a[0])
	}},
	{"connections.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UpdateConnectionInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Connections.Update(ctx, a[0], bodyArg)
	}},
	{"countries.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Countries.List(ctx)
	}},
	{"creators.faces", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Creators.Faces(ctx, a[0])
	}},
	{"creators.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Creators.Get(ctx, a[0])
	}},
	{"creators.regenerateAbout", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Creators.RegenerateAbout(ctx, a[0])
	}},
	{"creators.reputation", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Creators.Reputation(ctx, a[0])
	}},
	{"creators.setAboutHidden", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CreatorsSetAboutHiddenBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Creators.SetAboutHidden(ctx, a[0], bodyArg)
	}},
	{"crm.connected", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CrmConnectedQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.CRM.Connected(ctx, a[0], a[1], queryArg)
	}},
	{"crm.mailboxes", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CrmMailboxesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.CRM.Mailboxes(ctx, queryArg)
	}},
	{"crm.actions.bookingLink", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmActionsBookingLinkBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Actions.BookingLink(ctx, a[0], a[1], bodyArg)
	}},
	{"crm.actions.document", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmActionsDocumentBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Actions.Document(ctx, a[0], a[1], bodyArg)
	}},
	{"crm.actions.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Actions.Get(ctx, a[0], a[1])
	}},
	{"crm.actions.requestSignature", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmActionsRequestSignatureBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Actions.RequestSignature(ctx, a[0], a[1], bodyArg)
	}},
	{"crm.actions.signatureOptions", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CrmActionsSignatureOptionsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.CRM.Actions.SignatureOptions(ctx, a[0], a[1], queryArg)
	}},
	{"crm.activities.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Activities.Delete(ctx, a[0], a[1], a[2])
	}},
	{"crm.activities.log", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmLogActivityInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Activities.Log(ctx, a[0], a[1], bodyArg)
	}},
	{"crm.activities.logCall", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmLogCallInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Activities.LogCall(ctx, a[0], a[1], bodyArg)
	}},
	{"crm.activities.timeline", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CrmActivitiesTimelineQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.CRM.Activities.Timeline(ctx, a[0], a[1], queryArg)
	}},
	{"crm.activities.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmActivitiesUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Activities.Update(ctx, a[0], a[1], a[2], bodyArg)
	}},
	{"crm.brief.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmBriefGetBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Brief.Get(ctx, a[0], a[1], bodyArg)
	}},
	{"crm.brief.saveNote", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmBriefSaveNoteBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Brief.SaveNote(ctx, a[0], a[1], bodyArg)
	}},
	{"crm.captureRules.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmCaptureRulesCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.CaptureRules.Create(ctx, bodyArg)
	}},
	{"crm.captureRules.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.CaptureRules.Delete(ctx, a[0])
	}},
	{"crm.captureRules.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CrmCaptureRulesListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.CRM.CaptureRules.List(ctx, queryArg)
	}},
	{"crm.captureRules.sources", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CrmCaptureRulesSourcesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.CRM.CaptureRules.Sources(ctx, queryArg)
	}},
	{"crm.captureRules.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmCaptureRulesUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.CaptureRules.Update(ctx, a[0], bodyArg)
	}},
	{"crm.consents.forSpace", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Consents.ForSpace(ctx, a[0])
	}},
	{"crm.consents.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Consents.Get(ctx, a[0], a[1])
	}},
	{"crm.consents.grant", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg any
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Consents.Grant(ctx, a[0], a[1], bodyArg)
	}},
	{"crm.consents.reconsent", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmConsentsReconsentBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Consents.Reconsent(ctx, a[0], a[1], bodyArg)
	}},
	{"crm.dealRoom.addStep", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmDealRoomStepInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.DealRoom.AddStep(ctx, a[0], a[1], bodyArg)
	}},
	{"crm.dealRoom.deleteStep", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.DealRoom.DeleteStep(ctx, a[0], a[1], a[2])
	}},
	{"crm.dealRoom.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.DealRoom.Get(ctx, a[0], a[1])
	}},
	{"crm.dealRoom.message", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg any
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.DealRoom.Message(ctx, a[0], a[1], bodyArg)
	}},
	{"crm.dealRoom.options", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CrmDealRoomOptionsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.CRM.DealRoom.Options(ctx, a[0], a[1], queryArg)
	}},
	{"crm.dealRoom.share", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmDealRoomShareBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.DealRoom.Share(ctx, a[0], a[1], bodyArg)
	}},
	{"crm.dealRoom.unshare", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.DealRoom.Unshare(ctx, a[0], a[1], a[2])
	}},
	{"crm.dealRoom.updateStep", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmDealRoomStepPatch
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.DealRoom.UpdateStep(ctx, a[0], a[1], a[2], bodyArg)
	}},
	{"crm.duplicates.check", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg any
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Duplicates.Check(ctx, a[0], bodyArg)
	}},
	{"crm.duplicates.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Duplicates.List(ctx, a[0])
	}},
	{"crm.duplicates.merge", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmMergeRequest
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Duplicates.Merge(ctx, a[0], bodyArg)
	}},
	{"crm.duplicates.previewMerge", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmMergeRequest
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Duplicates.PreviewMerge(ctx, a[0], bodyArg)
	}},
	{"crm.enrichment.bulk", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmEnrichmentBulkBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Enrichment.Bulk(ctx, a[0], bodyArg)
	}},
	{"crm.enrichment.configure", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmEnrichmentConfigureBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Enrichment.Configure(ctx, a[0], bodyArg)
	}},
	{"crm.enrichment.decide", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmEnrichmentDecideBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Enrichment.Decide(ctx, a[0], a[1], bodyArg)
	}},
	{"crm.enrichment.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Enrichment.Get(ctx, a[0], a[1])
	}},
	{"crm.enrichment.run", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Enrichment.Run(ctx, a[0], a[1])
	}},
	{"crm.enrichment.settings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Enrichment.Settings(ctx, a[0])
	}},
	{"crm.groups.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Groups.Delete(ctx, a[0])
	}},
	{"crm.groups.rename", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmGroupsRenameBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Groups.Rename(ctx, a[0], bodyArg)
	}},
	{"crm.journey.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Journey.Get(ctx, a[0], a[1])
	}},
	{"crm.journey.summarise", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmJourneySummariseBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Journey.Summarise(ctx, a[0], a[1], bodyArg)
	}},
	{"crm.matches.forRecord", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Matches.ForRecord(ctx, a[0], a[1])
	}},
	{"crm.matches.forSpace", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Matches.ForSpace(ctx, a[0])
	}},
	{"crm.matches.propose", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmMatchesProposeBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Matches.Propose(ctx, a[0], a[1], bodyArg)
	}},
	{"crm.matching.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CrmMatchingGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.CRM.Matching.Get(ctx, a[0], queryArg)
	}},
	{"crm.matching.put", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmMatchingPutBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Matching.Put(ctx, a[0], bodyArg)
	}},
	{"crm.meetings.apply", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmMeetingsApplyBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Meetings.Apply(ctx, a[0], bodyArg)
	}},
	{"crm.meetings.proposal", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CrmMeetingsProposalQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.CRM.Meetings.Proposal(ctx, a[0], queryArg)
	}},
	{"crm.offlineConversions.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.OfflineConversions.Get(ctx, a[0])
	}},
	{"crm.offlineConversions.put", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmOfflineConversionsPutBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.OfflineConversions.Put(ctx, a[0], bodyArg)
	}},
	{"crm.pipeline.add", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmPipelineAddBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Pipeline.Add(ctx, a[0], bodyArg)
	}},
	{"crm.pipeline.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmPipelineUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Pipeline.Update(ctx, a[0], bodyArg)
	}},
	{"crm.playbooks.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Playbooks.Get(ctx, a[0])
	}},
	{"crm.playbooks.put", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmPlaybooksPutBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Playbooks.Put(ctx, a[0], bodyArg)
	}},
	{"crm.playbooks.retry", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Playbooks.Retry(ctx, a[0], a[1], a[2])
	}},
	{"crm.playbooks.runs", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Playbooks.Runs(ctx, a[0], a[1])
	}},
	{"crm.portal.invitations", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Portal.Invitations(ctx, a[0], a[1])
	}},
	{"crm.portal.invite", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmPortalInviteBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Portal.Invite(ctx, a[0], a[1], bodyArg)
	}},
	{"crm.portal.inviteOptions", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CrmPortalInviteOptionsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.CRM.Portal.InviteOptions(ctx, a[0], a[1], queryArg)
	}},
	{"crm.portal.record", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Portal.Record(ctx, a[0], a[1])
	}},
	{"crm.portal.resend", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Portal.Resend(ctx, a[0])
	}},
	{"crm.portal.setSitePage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmPortalSetSitePageBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Portal.SetSitePage(ctx, a[0], bodyArg)
	}},
	{"crm.portal.siteInvitations", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Portal.SiteInvitations(ctx, a[0])
	}},
	{"crm.portal.sites", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CrmPortalSitesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.CRM.Portal.Sites(ctx, queryArg)
	}},
	{"crm.portal.spaces", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CrmPortalSpacesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.CRM.Portal.Spaces(ctx, queryArg)
	}},
	{"crm.portal.withdraw", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Portal.Withdraw(ctx, a[0])
	}},
	{"crm.records.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg any
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Records.Create(ctx, a[0], bodyArg)
	}},
	{"crm.records.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Records.Delete(ctx, a[0], a[1])
	}},
	{"crm.records.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Records.Get(ctx, a[0], a[1])
	}},
	{"crm.records.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CrmRecordsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.CRM.Records.List(ctx, a[0], queryArg)
	}},
	{"crm.records.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmRecordsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Records.Update(ctx, a[0], a[1], bodyArg)
	}},
	{"crm.sales.dealSignals", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Sales.DealSignals(ctx, a[0])
	}},
	{"crm.sales.flowFunnel", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CrmSalesFlowFunnelQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.CRM.Sales.FlowFunnel(ctx, queryArg)
	}},
	{"crm.sales.report", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CrmSalesReportQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.CRM.Sales.Report(ctx, queryArg)
	}},
	{"crm.spaces.convert", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmSpacesConvertBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Spaces.Convert(ctx, bodyArg)
	}},
	{"crm.spaces.convertible", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CrmSpacesConvertibleQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.CRM.Spaces.Convertible(ctx, queryArg)
	}},
	{"crm.spaces.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmSpacesCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Spaces.Create(ctx, bodyArg)
	}},
	{"crm.spaces.design", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmSpacesDesignBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Spaces.Design(ctx, bodyArg)
	}},
	{"crm.spaces.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Spaces.Get(ctx, a[0])
	}},
	{"crm.spaces.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *CrmSpacesListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.CRM.Spaces.List(ctx, queryArg)
	}},
	{"crm.spaces.setGroup", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg any
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Spaces.SetGroup(ctx, a[0], bodyArg)
	}},
	{"crm.spaces.setup", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg any
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Spaces.Setup(ctx, bodyArg)
	}},
	{"crm.spaces.summary", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Spaces.Summary(ctx, a[0])
	}},
	{"crm.spaces.updateConfig", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmSpacesUpdateConfigBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Spaces.UpdateConfig(ctx, a[0], bodyArg)
	}},
	{"crm.tasks.setDone", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmTasksSetDoneBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Tasks.SetDone(ctx, a[0], a[1], bodyArg)
	}},
	{"crm.website.createPages", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmWebsiteCreatePagesBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Website.CreatePages(ctx, a[0], bodyArg)
	}},
	{"crm.website.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Website.Get(ctx, a[0])
	}},
	{"crm.website.preview", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CrmWebsitePreviewBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.CRM.Website.Preview(ctx, a[0], bodyArg)
	}},
	{"crm.website.record", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.CRM.Website.Record(ctx, a[0], a[1])
	}},
	{"currencies.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Currencies.List(ctx)
	}},
	{"currencies.rates", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Currencies.Rates(ctx)
	}},
	{"dashboards.addWidget", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DashboardsAddWidgetBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Dashboards.AddWidget(ctx, a[0], bodyArg)
	}},
	{"dashboards.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DashboardsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Dashboards.Create(ctx, bodyArg)
	}},
	{"dashboards.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Dashboards.Delete(ctx, a[0])
	}},
	{"dashboards.deleteWidget", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Dashboards.DeleteWidget(ctx, a[0], a[1])
	}},
	{"dashboards.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Dashboards.Get(ctx, a[0])
	}},
	{"dashboards.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *DashboardsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Dashboards.List(ctx, queryArg)
	}},
	{"dashboards.listWidgets", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Dashboards.ListWidgets(ctx, a[0])
	}},
	{"dashboards.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DashboardsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Dashboards.Update(ctx, a[0], bodyArg)
	}},
	{"dashboards.updateWidget", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DashboardsUpdateWidgetBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Dashboards.UpdateWidget(ctx, a[0], a[1], bodyArg)
	}},
	{"data.export", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *DataExportQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Data.Export(ctx, queryArg)
	}},
	{"data.exports.download", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *DataExportsDownloadQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Data.Exports.Download(ctx, queryArg)
	}},
	{"data.exports.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *DataExportsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Data.Exports.List(ctx, queryArg)
	}},
	{"data.exports.start", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DataExportsStartBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Data.Exports.Start(ctx, bodyArg)
	}},
	{"databases.activity", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *DatabasesActivityQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Databases.Activity(ctx, a[0], queryArg)
	}},
	{"databases.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DatabasesCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Databases.Create(ctx, bodyArg)
	}},
	{"databases.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Databases.Delete(ctx, a[0])
	}},
	{"databases.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *DatabasesGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Databases.Get(ctx, a[0], queryArg)
	}},
	{"databases.import", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DatabasesImportBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Databases.Import(ctx, a[0], bodyArg)
	}},
	{"databases.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *DatabasesListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Databases.List(ctx, queryArg)
	}},
	{"databases.people", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Databases.People(ctx, a[0])
	}},
	{"databases.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DatabasesUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Databases.Update(ctx, a[0], bodyArg)
	}},
	{"databases.columns.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Databases.Columns.Delete(ctx, a[0], a[1])
	}},
	{"databases.columns.disableCalendarFeed", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Databases.Columns.DisableCalendarFeed(ctx, a[0], a[1])
	}},
	{"databases.columns.enableCalendarFeed", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DatabasesColumnsEnableCalendarFeedBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Databases.Columns.EnableCalendarFeed(ctx, a[0], a[1], bodyArg)
	}},
	{"databases.columns.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Databases.Columns.Get(ctx, a[0], a[1])
	}},
	{"databases.columns.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DatabasesColumnsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Databases.Columns.Update(ctx, a[0], a[1], bodyArg)
	}},
	{"databases.records.activity", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *DatabasesRecordsActivityQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Databases.Records.Activity(ctx, a[0], a[1], queryArg)
	}},
	{"databases.records.aiFormula", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DatabasesRecordsAiFormulaBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Databases.Records.AIFormula(ctx, a[0], a[1], bodyArg)
	}},
	{"databases.records.bulk", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DatabasesRecordsBulkBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Databases.Records.Bulk(ctx, a[0], bodyArg)
	}},
	{"databases.records.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CreateRecordInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Databases.Records.Create(ctx, a[0], bodyArg)
	}},
	{"databases.records.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Databases.Records.Delete(ctx, a[0], a[1])
	}},
	{"databases.records.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Databases.Records.Get(ctx, a[0], a[1])
	}},
	{"databases.records.links", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Databases.Records.Links(ctx, a[0], a[1])
	}},
	{"databases.records.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *DatabasesRecordsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Databases.Records.List(ctx, a[0], queryArg)
	}},
	{"databases.records.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UpdateRecordInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Databases.Records.Update(ctx, a[0], a[1], bodyArg)
	}},
	{"databases.shares.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *DatabasesSharesListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Databases.Shares.List(ctx, a[0], queryArg)
	}},
	{"databases.shares.remove", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Databases.Shares.Remove(ctx, a[0], a[1])
	}},
	{"databases.shares.set", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DatabasesSharesSetBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Databases.Shares.Set(ctx, a[0], bodyArg)
	}},
	{"documents.bulk", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DocumentsBulkBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Documents.Bulk(ctx, bodyArg)
	}},
	{"documents.compare", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *DocumentsCompareQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Documents.Compare(ctx, a[0], queryArg)
	}},
	{"documents.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DocumentsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Documents.Create(ctx, bodyArg)
	}},
	{"documents.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *DocumentsDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Documents.Delete(ctx, a[0], queryArg)
	}},
	{"documents.duplicate", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DocumentsDuplicateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Documents.Duplicate(ctx, a[0], bodyArg)
	}},
	{"documents.export", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *DocumentsExportQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Documents.Export(ctx, a[0], queryArg)
	}},
	{"documents.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Documents.Get(ctx, a[0])
	}},
	{"documents.google", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *DocumentsGoogleQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Documents.Google(ctx, queryArg)
	}},
	{"documents.import", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DocumentsImportBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Documents.Import(ctx, a[0], bodyArg)
	}},
	{"documents.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *DocumentsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Documents.List(ctx, queryArg)
	}},
	{"documents.move", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DocumentsMoveBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Documents.Move(ctx, a[0], bodyArg)
	}},
	{"documents.restore", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Documents.Restore(ctx, a[0])
	}},
	{"documents.search", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *DocumentsSearchQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Documents.Search(ctx, queryArg)
	}},
	{"documents.setFavorite", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DocumentsSetFavoriteBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Documents.SetFavorite(ctx, a[0], bodyArg)
	}},
	{"documents.templates", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *DocumentsTemplatesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Documents.Templates(ctx, queryArg)
	}},
	{"documents.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DocumentsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Documents.Update(ctx, a[0], bodyArg)
	}},
	{"documents.comments.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DocumentsCommentsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Documents.Comments.Create(ctx, a[0], a[1], bodyArg)
	}},
	{"documents.comments.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Documents.Comments.Delete(ctx, a[0], a[1], a[2])
	}},
	{"documents.comments.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Documents.Comments.List(ctx, a[0], a[1])
	}},
	{"documents.comments.react", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DocumentsCommentsReactBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Documents.Comments.React(ctx, a[0], a[1], a[2], bodyArg)
	}},
	{"documents.comments.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DocumentsCommentsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Documents.Comments.Update(ctx, a[0], a[1], a[2], bodyArg)
	}},
	{"documents.mailMerge.clear", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Documents.MailMerge.Clear(ctx, a[0])
	}},
	{"documents.mailMerge.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *DocumentsMailMergeGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Documents.MailMerge.Get(ctx, a[0], queryArg)
	}},
	{"documents.mailMerge.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DocumentsMailMergeUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Documents.MailMerge.Update(ctx, a[0], bodyArg)
	}},
	{"documents.pages.backlinks", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Documents.Pages.Backlinks(ctx, a[0], a[1])
	}},
	{"documents.pages.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DocumentsPagesCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Documents.Pages.Create(ctx, a[0], bodyArg)
	}},
	{"documents.pages.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Documents.Pages.Delete(ctx, a[0], a[1])
	}},
	{"documents.pages.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Documents.Pages.Get(ctx, a[0], a[1])
	}},
	{"documents.pages.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Documents.Pages.List(ctx, a[0])
	}},
	{"documents.pages.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DocumentsPagesUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Documents.Pages.Update(ctx, a[0], a[1], bodyArg)
	}},
	{"documents.shares.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *DocumentsSharesCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Documents.Shares.Create(ctx, a[0], bodyArg)
	}},
	{"documents.shares.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Documents.Shares.List(ctx, a[0])
	}},
	{"documents.shares.revoke", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Documents.Shares.Revoke(ctx, a[0], a[1])
	}},
	{"documents.versions.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Documents.Versions.Get(ctx, a[0], a[1], a[2])
	}},
	{"documents.versions.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Documents.Versions.List(ctx, a[0], a[1])
	}},
	{"documents.versions.restore", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Documents.Versions.Restore(ctx, a[0], a[1], a[2])
	}},
	{"endpoints.consumers", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Endpoints.Consumers(ctx, a[0])
	}},
	{"endpoints.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CreateEndpointInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Endpoints.Create(ctx, bodyArg)
	}},
	{"endpoints.createApiKey", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CreateEndpointApiKeyInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Endpoints.CreateAPIKey(ctx, a[0], bodyArg)
	}},
	{"endpoints.createEnvVar", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CreateEndpointEnvVarInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Endpoints.CreateEnvVar(ctx, a[0], bodyArg)
	}},
	{"endpoints.createWebhook", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CreateEndpointWebhookInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Endpoints.CreateWebhook(ctx, a[0], bodyArg)
	}},
	{"endpoints.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Endpoints.Delete(ctx, a[0])
	}},
	{"endpoints.deleteApiKey", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Endpoints.DeleteAPIKey(ctx, a[0], a[1])
	}},
	{"endpoints.deleteEnvVar", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Endpoints.DeleteEnvVar(ctx, a[0], a[1])
	}},
	{"endpoints.deleteWebhook", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Endpoints.DeleteWebhook(ctx, a[0], a[1])
	}},
	{"endpoints.dependencies", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Endpoints.Dependencies(ctx, a[0])
	}},
	{"endpoints.explainCapture", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *EndpointsExplainCaptureBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Endpoints.ExplainCapture(ctx, a[0], bodyArg)
	}},
	{"endpoints.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Endpoints.Get(ctx, a[0])
	}},
	{"endpoints.getActivity", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *EndpointsGetActivityQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Endpoints.GetActivity(ctx, a[0], queryArg)
	}},
	{"endpoints.getAnalytics", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *EndpointsGetAnalyticsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Endpoints.GetAnalytics(ctx, a[0], queryArg)
	}},
	{"endpoints.getAuthConfig", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Endpoints.GetAuthConfig(ctx, a[0])
	}},
	{"endpoints.getCapture", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Endpoints.GetCapture(ctx, a[0], a[1])
	}},
	{"endpoints.getDocs", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *EndpointsGetDocsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Endpoints.GetDocs(ctx, a[0], queryArg)
	}},
	{"endpoints.getTimings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Endpoints.GetTimings(ctx, a[0], a[1])
	}},
	{"endpoints.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *EndpointsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Endpoints.List(ctx, queryArg)
	}},
	{"endpoints.listApiKeys", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Endpoints.ListAPIKeys(ctx, a[0])
	}},
	{"endpoints.listCaptures", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *EndpointsListCapturesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Endpoints.ListCaptures(ctx, a[0], queryArg)
	}},
	{"endpoints.listEnvVars", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Endpoints.ListEnvVars(ctx, a[0])
	}},
	{"endpoints.listWebhooks", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Endpoints.ListWebhooks(ctx, a[0])
	}},
	{"endpoints.logActivity", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *EndpointsLogActivityBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Endpoints.LogActivity(ctx, a[0], bodyArg)
	}},
	{"endpoints.replace", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UpdateEndpointInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Endpoints.Replace(ctx, a[0], bodyArg)
	}},
	{"endpoints.shadowResults", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *EndpointsShadowResultsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Endpoints.ShadowResults(ctx, a[0], queryArg)
	}},
	{"endpoints.suspendApiKey", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *EndpointsSuspendApiKeyBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Endpoints.SuspendAPIKey(ctx, a[0], a[1], bodyArg)
	}},
	{"endpoints.testAuth", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *EndpointsTestAuthBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Endpoints.TestAuth(ctx, a[0], bodyArg)
	}},
	{"endpoints.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UpdateEndpointInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Endpoints.Update(ctx, a[0], bodyArg)
	}},
	{"endpoints.updateAuthConfig", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *EndpointsUpdateAuthConfigBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Endpoints.UpdateAuthConfig(ctx, a[0], bodyArg)
	}},
	{"endpoints.updateEnvVar", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *EndpointsUpdateEnvVarBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Endpoints.UpdateEnvVar(ctx, a[0], a[1], bodyArg)
	}},
	{"endpoints.updateWebhook", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *EndpointsUpdateWebhookBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Endpoints.UpdateWebhook(ctx, a[0], a[1], bodyArg)
	}},
	{"featureGates.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *FeatureGatesGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.FeatureGates.Get(ctx, queryArg)
	}},
	{"featureRequests.addComment", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *FeatureRequestsAddCommentBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.FeatureRequests.AddComment(ctx, a[0], bodyArg)
	}},
	{"featureRequests.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *FeatureRequestsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.FeatureRequests.Create(ctx, bodyArg)
	}},
	{"featureRequests.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.FeatureRequests.Get(ctx, a[0])
	}},
	{"featureRequests.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *FeatureRequestsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.FeatureRequests.List(ctx, queryArg)
	}},
	{"featureRequests.listComments", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.FeatureRequests.ListComments(ctx, a[0])
	}},
	{"featureRequests.unvote", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.FeatureRequests.Unvote(ctx, a[0])
	}},
	{"featureRequests.vote", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.FeatureRequests.Vote(ctx, a[0])
	}},
	{"help.getArticle", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Help.GetArticle(ctx, a[0])
	}},
	{"help.listArticles", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *HelpListArticlesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Help.ListArticles(ctx, queryArg)
	}},
	{"invoicing.approvalRules.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingApprovalRulesCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingApprovalRulesCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.ApprovalRules.Create(ctx, bodyArg, queryArg)
	}},
	{"invoicing.approvalRules.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingApprovalRulesDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.ApprovalRules.Delete(ctx, a[0], queryArg)
	}},
	{"invoicing.approvalRules.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingApprovalRulesListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.ApprovalRules.List(ctx, queryArg)
	}},
	{"invoicing.approvalRules.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingApprovalRulesUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingApprovalRulesUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.ApprovalRules.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.customers.checkVat", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingCustomersCheckVatQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Customers.CheckVAT(ctx, a[0], queryArg)
	}},
	{"invoicing.customers.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingCustomerInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingCustomersCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Customers.Create(ctx, bodyArg, queryArg)
	}},
	{"invoicing.customers.createFromCrm", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingCustomersCreateFromCrmBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingCustomersCreateFromCrmQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Customers.CreateFromCRM(ctx, bodyArg, queryArg)
	}},
	{"invoicing.customers.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingCustomersDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Customers.Delete(ctx, a[0], queryArg)
	}},
	{"invoicing.customers.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingCustomersGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Customers.Get(ctx, a[0], queryArg)
	}},
	{"invoicing.customers.import", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingCustomersImportBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingCustomersImportQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Customers.Import(ctx, bodyArg, queryArg)
	}},
	{"invoicing.customers.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingCustomersListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Customers.List(ctx, queryArg)
	}},
	{"invoicing.customers.previewFromCrm", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingCustomersPreviewFromCrmQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Customers.PreviewFromCRM(ctx, queryArg)
	}},
	{"invoicing.customers.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingCustomersUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingCustomersUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Customers.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.documents.answerQuote", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingDocumentsAnswerQuoteBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingDocumentsAnswerQuoteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Documents.AnswerQuote(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.documents.bankLines", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingDocumentsBankLinesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Documents.BankLines(ctx, a[0], queryArg)
	}},
	{"invoicing.documents.checkEinvoice", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingDocumentsCheckEinvoiceQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Documents.CheckEinvoice(ctx, a[0], queryArg)
	}},
	{"invoicing.documents.convert", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingDocumentsConvertBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingDocumentsConvertQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Documents.Convert(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.documents.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingDraftInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingDocumentsCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Documents.Create(ctx, bodyArg, queryArg)
	}},
	{"invoicing.documents.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingDocumentsDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Documents.Delete(ctx, a[0], queryArg)
	}},
	{"invoicing.documents.duplicate", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingDocumentsDuplicateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Documents.Duplicate(ctx, a[0], queryArg)
	}},
	{"invoicing.documents.export", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingDocumentsExportQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Documents.Export(ctx, queryArg)
	}},
	{"invoicing.documents.finalize", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingDocumentsFinalizeQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Documents.Finalize(ctx, a[0], queryArg)
	}},
	{"invoicing.documents.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingDocumentsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Documents.Get(ctx, a[0], queryArg)
	}},
	{"invoicing.documents.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingDocumentsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Documents.List(ctx, queryArg)
	}},
	{"invoicing.documents.markSent", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingDocumentsMarkSentBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingDocumentsMarkSentQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Documents.MarkSent(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.documents.pdf", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingDocumentsPdfQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Documents.PDF(ctx, a[0], queryArg)
	}},
	{"invoicing.documents.send", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingDocumentsSendBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingDocumentsSendQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Documents.Send(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.documents.sendDefaults", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingDocumentsSendDefaultsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Documents.SendDefaults(ctx, a[0], queryArg)
	}},
	{"invoicing.documents.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingDraftPatch
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingDocumentsUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Documents.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.documents.xml", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingDocumentsXmlQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Documents.Xml(ctx, a[0], queryArg)
	}},
	{"invoicing.dunning.getNotice", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingDunningGetNoticeQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Dunning.GetNotice(ctx, a[0], queryArg)
	}},
	{"invoicing.dunning.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingDunningListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Dunning.List(ctx, queryArg)
	}},
	{"invoicing.dunning.noticePdf", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingDunningNoticePdfQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Dunning.NoticePDF(ctx, a[0], queryArg)
	}},
	{"invoicing.dunning.pauseCustomer", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingDunningPauseCustomerBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingDunningPauseCustomerQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Dunning.PauseCustomer(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.dunning.pauseDocument", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingDunningPauseDocumentBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingDunningPauseDocumentQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Dunning.PauseDocument(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.dunning.propose", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingDunningProposeBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingDunningProposeQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Dunning.Propose(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.dunning.removeNotice", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingDunningRemoveNoticeQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Dunning.RemoveNotice(ctx, a[0], queryArg)
	}},
	{"invoicing.dunning.sendDefaults", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingDunningSendDefaultsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Dunning.SendDefaults(ctx, a[0], queryArg)
	}},
	{"invoicing.dunning.sendNotice", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingDunningSendNoticeBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingDunningSendNoticeQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Dunning.SendNotice(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.dunning.state", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingDunningStateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Dunning.State(ctx, a[0], queryArg)
	}},
	{"invoicing.inbound.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingInboundDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Inbound.Delete(ctx, a[0], queryArg)
	}},
	{"invoicing.inbound.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingInboundGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Inbound.Get(ctx, a[0], queryArg)
	}},
	{"invoicing.inbound.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingInboundListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Inbound.List(ctx, queryArg)
	}},
	{"invoicing.inbound.original", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingInboundOriginalQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Inbound.Original(ctx, a[0], queryArg)
	}},
	{"invoicing.inbound.revalidate", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingInboundRevalidateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Inbound.Revalidate(ctx, a[0], queryArg)
	}},
	{"invoicing.inbound.transition", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingInboundTransitionBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingInboundTransitionQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Inbound.Transition(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.inbound.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingInboundUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingInboundUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Inbound.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.inbound.upload", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingInboundUploadBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingInboundUploadQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Inbound.Upload(ctx, bodyArg, queryArg)
	}},
	{"invoicing.inbound.mailboxes.add", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingInboundMailboxesAddBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingInboundMailboxesAddQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Inbound.Mailboxes.Add(ctx, bodyArg, queryArg)
	}},
	{"invoicing.inbound.mailboxes.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingInboundMailboxesListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Inbound.Mailboxes.List(ctx, queryArg)
	}},
	{"invoicing.inbound.mailboxes.remove", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingInboundMailboxesRemoveQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Inbound.Mailboxes.Remove(ctx, a[0], queryArg)
	}},
	{"invoicing.inbound.mailboxes.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingInboundMailboxesUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingInboundMailboxesUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Inbound.Mailboxes.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.issuers.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingIssuerInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingIssuersCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Issuers.Create(ctx, bodyArg, queryArg)
	}},
	{"invoicing.issuers.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingIssuersDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Issuers.Delete(ctx, a[0], queryArg)
	}},
	{"invoicing.issuers.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingIssuersGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Issuers.Get(ctx, a[0], queryArg)
	}},
	{"invoicing.issuers.getDunningSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingIssuersGetDunningSettingsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Issuers.GetDunningSettings(ctx, a[0], queryArg)
	}},
	{"invoicing.issuers.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingIssuersListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Issuers.List(ctx, queryArg)
	}},
	{"invoicing.issuers.setDunningSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingIssuersSetDunningSettingsBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingIssuersSetDunningSettingsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Issuers.SetDunningSettings(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.issuers.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingIssuersUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingIssuersUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Issuers.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.ledger.events", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingLedgerEventsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Ledger.Events(ctx, queryArg)
	}},
	{"invoicing.ledger.opening", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingLedgerOpeningBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingLedgerOpeningQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Ledger.Opening(ctx, bodyArg, queryArg)
	}},
	{"invoicing.ledger.report", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingLedgerReportQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Ledger.Report(ctx, a[0], queryArg)
	}},
	{"invoicing.ledger.start", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingLedgerStartBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingLedgerStartQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Ledger.Start(ctx, bodyArg, queryArg)
	}},
	{"invoicing.ledger.accounts.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingLedgerAccountsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingLedgerAccountsCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Ledger.Accounts.Create(ctx, bodyArg, queryArg)
	}},
	{"invoicing.ledger.accounts.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingLedgerAccountsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Ledger.Accounts.List(ctx, queryArg)
	}},
	{"invoicing.ledger.accounts.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingLedgerAccountsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingLedgerAccountsUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Ledger.Accounts.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.ledger.datev.export", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingLedgerDatevExportBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingLedgerDatevExportQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Ledger.Datev.Export(ctx, bodyArg, queryArg)
	}},
	{"invoicing.ledger.datev.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingLedgerDatevListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Ledger.Datev.List(ctx, queryArg)
	}},
	{"invoicing.ledger.entries.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingLedgerEntriesListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Ledger.Entries.List(ctx, queryArg)
	}},
	{"invoicing.ledger.entries.post", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingLedgerEntriesPostBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingLedgerEntriesPostQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Ledger.Entries.Post(ctx, bodyArg, queryArg)
	}},
	{"invoicing.ledger.entries.reverse", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingLedgerEntriesReverseBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingLedgerEntriesReverseQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Ledger.Entries.Reverse(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.ledger.periods.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingLedgerPeriodsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Ledger.Periods.List(ctx, queryArg)
	}},
	{"invoicing.ledger.periods.set", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingLedgerPeriodsSetBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingLedgerPeriodsSetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Ledger.Periods.Set(ctx, bodyArg, queryArg)
	}},
	{"invoicing.ledger.settings.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingLedgerSettingsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingLedgerSettingsCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Ledger.Settings.Create(ctx, bodyArg, queryArg)
	}},
	{"invoicing.ledger.settings.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingLedgerSettingsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Ledger.Settings.Get(ctx, queryArg)
	}},
	{"invoicing.ledger.settings.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingLedgerSettingsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingLedgerSettingsUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Ledger.Settings.Update(ctx, bodyArg, queryArg)
	}},
	{"invoicing.ledger.yearEnd.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingLedgerYearEndGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Ledger.YearEnd.Get(ctx, queryArg)
	}},
	{"invoicing.ledger.yearEnd.set", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingLedgerYearEndSetBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingLedgerYearEndSetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Ledger.YearEnd.Set(ctx, bodyArg, queryArg)
	}},
	{"invoicing.numberRanges.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingNumberRangesListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.NumberRanges.List(ctx, queryArg)
	}},
	{"invoicing.numberRanges.preview", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingNumberRangesPreviewBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingNumberRangesPreviewQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.NumberRanges.Preview(ctx, bodyArg, queryArg)
	}},
	{"invoicing.numberRanges.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingNumberRangesUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingNumberRangesUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.NumberRanges.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.paymentRuns.close", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingPaymentRunsCloseBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingPaymentRunsCloseQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.PaymentRuns.Close(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.paymentRuns.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingPaymentRunsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingPaymentRunsCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.PaymentRuns.Create(ctx, bodyArg, queryArg)
	}},
	{"invoicing.paymentRuns.file", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingPaymentRunsFileQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.PaymentRuns.File(ctx, a[0], queryArg)
	}},
	{"invoicing.paymentRuns.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingPaymentRunsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.PaymentRuns.Get(ctx, a[0], queryArg)
	}},
	{"invoicing.paymentRuns.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingPaymentRunsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.PaymentRuns.List(ctx, queryArg)
	}},
	{"invoicing.payments.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingPaymentsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Payments.List(ctx, a[0], queryArg)
	}},
	{"invoicing.payments.record", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingPaymentsRecordBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingPaymentsRecordQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Payments.Record(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.payments.remove", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingPaymentsRemoveQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Payments.Remove(ctx, a[0], a[1], queryArg)
	}},
	{"invoicing.purchases.aiFill", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingPurchasesAiFillQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Purchases.AIFill(ctx, a[0], queryArg)
	}},
	{"invoicing.purchases.attachOriginal", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingPurchasesAttachOriginalBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingPurchasesAttachOriginalQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Purchases.AttachOriginal(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.purchases.bankLines", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingPurchasesBankLinesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Purchases.BankLines(ctx, a[0], queryArg)
	}},
	{"invoicing.purchases.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingPurchaseInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingPurchasesCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Purchases.Create(ctx, bodyArg, queryArg)
	}},
	{"invoicing.purchases.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingPurchasesDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Purchases.Delete(ctx, a[0], queryArg)
	}},
	{"invoicing.purchases.deletePayment", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingPurchasesDeletePaymentQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Purchases.DeletePayment(ctx, a[0], a[1], queryArg)
	}},
	{"invoicing.purchases.fromInbound", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingPurchasesFromInboundBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingPurchasesFromInboundQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Purchases.FromInbound(ctx, bodyArg, queryArg)
	}},
	{"invoicing.purchases.fromInboundStatus", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingPurchasesFromInboundStatusQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Purchases.FromInboundStatus(ctx, queryArg)
	}},
	{"invoicing.purchases.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingPurchasesGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Purchases.Get(ctx, a[0], queryArg)
	}},
	{"invoicing.purchases.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingPurchasesListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Purchases.List(ctx, queryArg)
	}},
	{"invoicing.purchases.options", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingPurchasesOptionsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Purchases.Options(ctx, queryArg)
	}},
	{"invoicing.purchases.original", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingPurchasesOriginalQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Purchases.Original(ctx, a[0], queryArg)
	}},
	{"invoicing.purchases.recordPayment", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingPurchasesRecordPaymentBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingPurchasesRecordPaymentQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Purchases.RecordPayment(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.purchases.transition", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingPurchasesTransitionBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingPurchasesTransitionQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Purchases.Transition(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.purchases.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingPurchaseInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingPurchasesUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Purchases.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.recurring.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingRecurringCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingRecurringCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Recurring.Create(ctx, bodyArg, queryArg)
	}},
	{"invoicing.recurring.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingRecurringDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Recurring.Delete(ctx, a[0], queryArg)
	}},
	{"invoicing.recurring.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingRecurringGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Recurring.Get(ctx, a[0], queryArg)
	}},
	{"invoicing.recurring.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingRecurringListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Recurring.List(ctx, queryArg)
	}},
	{"invoicing.recurring.prefill", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingRecurringPrefillQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Recurring.Prefill(ctx, queryArg)
	}},
	{"invoicing.recurring.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingRecurringUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingRecurringUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Recurring.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.reports.dataset", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingReportsDatasetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Reports.Dataset(ctx, queryArg)
	}},
	{"invoicing.reports.taxReturns", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingReportsTaxReturnsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Reports.TaxReturns(ctx, queryArg)
	}},
	{"invoicing.suppliers.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingSupplierInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingSuppliersCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Suppliers.Create(ctx, bodyArg, queryArg)
	}},
	{"invoicing.suppliers.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingSuppliersDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Suppliers.Delete(ctx, a[0], queryArg)
	}},
	{"invoicing.suppliers.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingSuppliersGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Suppliers.Get(ctx, a[0], queryArg)
	}},
	{"invoicing.suppliers.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingSuppliersListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Suppliers.List(ctx, queryArg)
	}},
	{"invoicing.suppliers.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingSuppliersUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingSuppliersUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Suppliers.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"invoicing.templates.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingTemplatesCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingTemplatesCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Templates.Create(ctx, bodyArg, queryArg)
	}},
	{"invoicing.templates.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingTemplatesDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Templates.Delete(ctx, a[0], queryArg)
	}},
	{"invoicing.templates.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingTemplatesGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Templates.Get(ctx, a[0], queryArg)
	}},
	{"invoicing.templates.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *InvoicingTemplatesListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Templates.List(ctx, queryArg)
	}},
	{"invoicing.templates.preview", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingTemplatesPreviewBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingTemplatesPreviewQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Templates.Preview(ctx, bodyArg, queryArg)
	}},
	{"invoicing.templates.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *InvoicingTemplatesUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *InvoicingTemplatesUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Invoicing.Templates.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"itemFolders.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ItemFoldersCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.ItemFolders.Create(ctx, bodyArg)
	}},
	{"itemFolders.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.ItemFolders.Delete(ctx, a[0])
	}},
	{"itemFolders.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ItemFoldersListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.ItemFolders.List(ctx, queryArg)
	}},
	{"itemFolders.move", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ItemFoldersMoveBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.ItemFolders.Move(ctx, bodyArg)
	}},
	{"itemFolders.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ItemFoldersUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.ItemFolders.Update(ctx, a[0], bodyArg)
	}},
	{"languages.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Languages.List(ctx)
	}},
	{"mail.addPgpKey", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailAddPgpKeyBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.AddPgpKey(ctx, bodyArg)
	}},
	{"mail.applyInviteReply", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.ApplyInviteReply(ctx, a[0])
	}},
	{"mail.applyRule", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailApplyRuleBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.ApplyRule(ctx, a[0], bodyArg)
	}},
	{"mail.blockSender", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailBlockSenderBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.BlockSender(ctx, bodyArg)
	}},
	{"mail.cancelScheduled", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.CancelScheduled(ctx, a[0])
	}},
	{"mail.checkPlatformAddress", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MailCheckPlatformAddressQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Mail.CheckPlatformAddress(ctx, queryArg)
	}},
	{"mail.counts", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MailCountsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Mail.Counts(ctx, queryArg)
	}},
	{"mail.createAccount", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailCreateAccountBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.CreateAccount(ctx, bodyArg)
	}},
	{"mail.createAlias", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailCreateAliasBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.CreateAlias(ctx, a[0], bodyArg)
	}},
	{"mail.createDraft", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailCreateDraftBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.CreateDraft(ctx, bodyArg)
	}},
	{"mail.createFolder", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailCreateFolderBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.CreateFolder(ctx, a[0], bodyArg)
	}},
	{"mail.createLabel", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailCreateLabelBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.CreateLabel(ctx, bodyArg)
	}},
	{"mail.createPlatformAddress", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailCreatePlatformAddressBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.CreatePlatformAddress(ctx, bodyArg)
	}},
	{"mail.createRule", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailCreateRuleBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.CreateRule(ctx, bodyArg)
	}},
	{"mail.createSignature", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailCreateSignatureBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.CreateSignature(ctx, bodyArg)
	}},
	{"mail.createSnippet", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailCreateSnippetBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.CreateSnippet(ctx, bodyArg)
	}},
	{"mail.createThreadNote", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailCreateThreadNoteBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.CreateThreadNote(ctx, a[0], bodyArg)
	}},
	{"mail.decrypt", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailDecryptBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.Decrypt(ctx, a[0], bodyArg)
	}},
	{"mail.deleteAccount", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.DeleteAccount(ctx, a[0])
	}},
	{"mail.deleteAlias", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.DeleteAlias(ctx, a[0], a[1])
	}},
	{"mail.deleteAttachment", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.DeleteAttachment(ctx, a[0])
	}},
	{"mail.deleteDraft", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.DeleteDraft(ctx, a[0])
	}},
	{"mail.deleteFolder", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MailDeleteFolderQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Mail.DeleteFolder(ctx, a[0], queryArg)
	}},
	{"mail.deleteLabel", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.DeleteLabel(ctx, a[0])
	}},
	{"mail.deletePgpKey", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.DeletePgpKey(ctx, a[0])
	}},
	{"mail.deleteRule", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.DeleteRule(ctx, a[0])
	}},
	{"mail.deleteSignature", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.DeleteSignature(ctx, a[0])
	}},
	{"mail.deleteSnippet", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.DeleteSnippet(ctx, a[0])
	}},
	{"mail.deleteThread", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.DeleteThread(ctx, a[0])
	}},
	{"mail.deleteThreadNote", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.DeleteThreadNote(ctx, a[0], a[1])
	}},
	{"mail.desk", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MailDeskQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Mail.Desk(ctx, queryArg)
	}},
	{"mail.downloadAttachment", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.DownloadAttachment(ctx, a[0])
	}},
	{"mail.emptyTrash", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailEmptyTrashBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.EmptyTrash(ctx, bodyArg)
	}},
	{"mail.getAccount", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.GetAccount(ctx, a[0])
	}},
	{"mail.getAccountOrganization", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.GetAccountOrganization(ctx, a[0])
	}},
	{"mail.getAutoResponse", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.GetAutoResponse(ctx, a[0])
	}},
	{"mail.getContact", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.GetContact(ctx, a[0])
	}},
	{"mail.getOpens", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.GetOpens(ctx, a[0])
	}},
	{"mail.getPlatformAddress", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.GetPlatformAddress(ctx)
	}},
	{"mail.getRaw", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MailGetRawQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Mail.GetRaw(ctx, a[0], queryArg)
	}},
	{"mail.getRetention", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.GetRetention(ctx, a[0])
	}},
	{"mail.getThread", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.GetThread(ctx, a[0])
	}},
	{"mail.listAccounts", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.ListAccounts(ctx)
	}},
	{"mail.listAliases", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.ListAliases(ctx, a[0])
	}},
	{"mail.listBlockedSenders", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.ListBlockedSenders(ctx)
	}},
	{"mail.listDrafts", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.ListDrafts(ctx)
	}},
	{"mail.listFolders", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.ListFolders(ctx, a[0])
	}},
	{"mail.listLabels", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.ListLabels(ctx)
	}},
	{"mail.listNewsletters", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MailListNewslettersQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Mail.ListNewsletters(ctx, queryArg)
	}},
	{"mail.listPgpKeys", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.ListPgpKeys(ctx)
	}},
	{"mail.listRules", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.ListRules(ctx)
	}},
	{"mail.listScheduled", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.ListScheduled(ctx)
	}},
	{"mail.listSignatures", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.ListSignatures(ctx)
	}},
	{"mail.listSnippets", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.ListSnippets(ctx)
	}},
	{"mail.listSubaddresses", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.ListSubaddresses(ctx)
	}},
	{"mail.listTeammates", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.ListTeammates(ctx, a[0])
	}},
	{"mail.listThreadNotes", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.ListThreadNotes(ctx, a[0])
	}},
	{"mail.listThreads", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MailListThreadsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Mail.ListThreads(ctx, queryArg)
	}},
	{"mail.lookupPgpKeys", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailLookupPgpKeysBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.LookupPgpKeys(ctx, bodyArg)
	}},
	{"mail.markAllRead", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailMarkAllReadBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.MarkAllRead(ctx, bodyArg)
	}},
	{"mail.parseInvite", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.ParseInvite(ctx, a[0])
	}},
	{"mail.patchThread", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailPatchThreadBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.PatchThread(ctx, a[0], bodyArg)
	}},
	{"mail.renameFolder", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailRenameFolderBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.RenameFolder(ctx, a[0], bodyArg)
	}},
	{"mail.reply", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailReplyBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.Reply(ctx, a[0], bodyArg)
	}},
	{"mail.resendAliasConfirmation", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.ResendAliasConfirmation(ctx, a[0], a[1])
	}},
	{"mail.rethread", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.Rethread(ctx, a[0])
	}},
	{"mail.rsvp", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailRsvpBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.Rsvp(ctx, a[0], bodyArg)
	}},
	{"mail.saveAttachmentToAssets", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailSaveAttachmentToAssetsBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.SaveAttachmentToAssets(ctx, a[0], bodyArg)
	}},
	{"mail.schedule", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailScheduleBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.Schedule(ctx, bodyArg)
	}},
	{"mail.search", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MailSearchQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Mail.Search(ctx, queryArg)
	}},
	{"mail.searchContacts", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MailSearchContactsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Mail.SearchContacts(ctx, queryArg)
	}},
	{"mail.send", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailSendBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.Send(ctx, bodyArg)
	}},
	{"mail.setAccountOrganization", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailSetAccountOrganizationBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.SetAccountOrganization(ctx, a[0], bodyArg)
	}},
	{"mail.setAutoResponse", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailSetAutoResponseBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.SetAutoResponse(ctx, a[0], bodyArg)
	}},
	{"mail.setCrmSync", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailSetCrmSyncBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.SetCRMSync(ctx, a[0], bodyArg)
	}},
	{"mail.setFlags", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailSetFlagsBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.SetFlags(ctx, a[0], bodyArg)
	}},
	{"mail.setFolders", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailSetFoldersBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.SetFolders(ctx, a[0], bodyArg)
	}},
	{"mail.setRetention", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailSetRetentionBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.SetRetention(ctx, a[0], bodyArg)
	}},
	{"mail.snoozeThread", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailSnoozeThreadBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.SnoozeThread(ctx, a[0], bodyArg)
	}},
	{"mail.storage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.Storage(ctx)
	}},
	{"mail.syncAccount", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailSyncAccountBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.SyncAccount(ctx, a[0], bodyArg)
	}},
	{"mail.threadDocumentHtml", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailThreadDocumentHtmlBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.ThreadDocumentHTML(ctx, a[0], bodyArg)
	}},
	{"mail.unblockSender", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.UnblockSender(ctx, a[0])
	}},
	{"mail.unsnoozeThread", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.UnsnoozeThread(ctx, a[0])
	}},
	{"mail.unsubscribe", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Mail.Unsubscribe(ctx, a[0])
	}},
	{"mail.updateAccount", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailUpdateAccountBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.UpdateAccount(ctx, a[0], bodyArg)
	}},
	{"mail.updateDraft", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailUpdateDraftBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.UpdateDraft(ctx, a[0], bodyArg)
	}},
	{"mail.updateLabel", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailUpdateLabelBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.UpdateLabel(ctx, a[0], bodyArg)
	}},
	{"mail.updatePgpKey", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailUpdatePgpKeyBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.UpdatePgpKey(ctx, a[0], bodyArg)
	}},
	{"mail.updateRule", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailUpdateRuleBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.UpdateRule(ctx, a[0], bodyArg)
	}},
	{"mail.updateSignature", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailUpdateSignatureBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.UpdateSignature(ctx, a[0], bodyArg)
	}},
	{"mail.updateSnippet", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailUpdateSnippetBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.UpdateSnippet(ctx, a[0], bodyArg)
	}},
	{"mail.uploadAttachment", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailUploadAttachmentBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.UploadAttachment(ctx, bodyArg)
	}},
	{"mail.ai.draftReply", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailAiDraftReplyBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.AI.DraftReply(ctx, bodyArg)
	}},
	{"mail.ai.rewrite", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailAiRewriteBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.AI.Rewrite(ctx, bodyArg)
	}},
	{"mail.ai.subject", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailAiSubjectBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.AI.Subject(ctx, bodyArg)
	}},
	{"mail.ai.summarizeThread", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailAiSummarizeThreadBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.AI.SummarizeThread(ctx, bodyArg)
	}},
	{"mail.imap.action", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailImapActionBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.IMAP.Action(ctx, bodyArg)
	}},
	{"mail.imap.fetch", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailImapFetchBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.IMAP.Fetch(ctx, bodyArg)
	}},
	{"mail.imap.verify", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailImapVerifyBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.IMAP.Verify(ctx, bodyArg)
	}},
	{"mail.smtp.verify", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MailSmtpVerifyBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Mail.SMTP.Verify(ctx, bodyArg)
	}},
	{"marketplace.acceptDelivery", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.AcceptDelivery(ctx, a[0])
	}},
	{"marketplace.addDisputeEvidence", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceAddDisputeEvidenceBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.AddDisputeEvidence(ctx, a[0], bodyArg)
	}},
	{"marketplace.addToWishlist", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceAddToWishlistBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.AddToWishlist(ctx, bodyArg)
	}},
	{"marketplace.answerQuestion", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceAnswerQuestionBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.AnswerQuestion(ctx, a[0], bodyArg)
	}},
	{"marketplace.askQuestion", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceAskQuestionBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.AskQuestion(ctx, a[0], bodyArg)
	}},
	{"marketplace.builtWith", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MarketplaceBuiltWithQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Marketplace.BuiltWith(ctx, a[0], queryArg)
	}},
	{"marketplace.cancelDispute", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.CancelDispute(ctx, a[0])
	}},
	{"marketplace.cancelSellerOrder", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceCancelSellerOrderBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.CancelSellerOrder(ctx, a[0], bodyArg)
	}},
	{"marketplace.cancelSubscription", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.CancelSubscription(ctx, a[0])
	}},
	{"marketplace.categories", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.Categories(ctx)
	}},
	{"marketplace.createCoupon", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceCreateCouponBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.CreateCoupon(ctx, bodyArg)
	}},
	{"marketplace.createReview", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceCreateReviewBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.CreateReview(ctx, a[0], bodyArg)
	}},
	{"marketplace.createSellerListing", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceCreateSellerListingBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.CreateSellerListing(ctx, bodyArg)
	}},
	{"marketplace.createSellerProfile", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceCreateSellerProfileBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.CreateSellerProfile(ctx, bodyArg)
	}},
	{"marketplace.deleteCoupon", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MarketplaceDeleteCouponQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Marketplace.DeleteCoupon(ctx, queryArg)
	}},
	{"marketplace.deleteReview", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MarketplaceDeleteReviewQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Marketplace.DeleteReview(ctx, a[0], queryArg)
	}},
	{"marketplace.deleteSellerListing", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.DeleteSellerListing(ctx, a[0])
	}},
	{"marketplace.deliverOrder", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceDeliverOrderBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.DeliverOrder(ctx, a[0], bodyArg)
	}},
	{"marketplace.download", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceDownloadBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.Download(ctx, a[0], bodyArg)
	}},
	{"marketplace.escalateDispute", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.EscalateDispute(ctx, a[0])
	}},
	{"marketplace.followShop", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.FollowShop(ctx, a[0])
	}},
	{"marketplace.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MarketplaceGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Marketplace.Get(ctx, a[0], queryArg)
	}},
	{"marketplace.getCollection", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MarketplaceGetCollectionQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Marketplace.GetCollection(ctx, a[0], queryArg)
	}},
	{"marketplace.getDispute", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.GetDispute(ctx, a[0])
	}},
	{"marketplace.getFees", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.GetFees(ctx)
	}},
	{"marketplace.getFile", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.GetFile(ctx, a[0], a[1])
	}},
	{"marketplace.getOrder", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.GetOrder(ctx, a[0])
	}},
	{"marketplace.getSellerListing", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.GetSellerListing(ctx, a[0])
	}},
	{"marketplace.getSellerProfile", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.GetSellerProfile(ctx)
	}},
	{"marketplace.getShop", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.GetShop(ctx, a[0])
	}},
	{"marketplace.getShowcase", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.GetShowcase(ctx, a[0])
	}},
	{"marketplace.getSubscriptions", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.GetSubscriptions(ctx)
	}},
	{"marketplace.hardDeleteSellerListing", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.HardDeleteSellerListing(ctx, a[0])
	}},
	{"marketplace.hideFromShowcase", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MarketplaceHideFromShowcaseQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Marketplace.HideFromShowcase(ctx, a[0], queryArg)
	}},
	{"marketplace.install", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceInstallBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.Install(ctx, a[0], bodyArg)
	}},
	{"marketplace.installAll", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceInstallAllBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.InstallAll(ctx, a[0], bodyArg)
	}},
	{"marketplace.installOrder", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceInstallOrderBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.InstallOrder(ctx, a[0], bodyArg)
	}},
	{"marketplace.installToWorkspace", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceInstallToWorkspaceBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.InstallToWorkspace(ctx, a[0], bodyArg)
	}},
	{"marketplace.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MarketplaceListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Marketplace.List(ctx, queryArg)
	}},
	{"marketplace.listCollections", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MarketplaceListCollectionsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Marketplace.ListCollections(ctx, queryArg)
	}},
	{"marketplace.listCoupons", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.ListCoupons(ctx)
	}},
	{"marketplace.listDisputes", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MarketplaceListDisputesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Marketplace.ListDisputes(ctx, queryArg)
	}},
	{"marketplace.listLegalPages", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.ListLegalPages(ctx)
	}},
	{"marketplace.listOrderDisputes", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.ListOrderDisputes(ctx, a[0])
	}},
	{"marketplace.listOrders", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MarketplaceListOrdersQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Marketplace.ListOrders(ctx, queryArg)
	}},
	{"marketplace.listPayouts", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MarketplaceListPayoutsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Marketplace.ListPayouts(ctx, queryArg)
	}},
	{"marketplace.listPurchases", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.ListPurchases(ctx)
	}},
	{"marketplace.listQuestions", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.ListQuestions(ctx, a[0])
	}},
	{"marketplace.listReviews", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MarketplaceListReviewsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Marketplace.ListReviews(ctx, a[0], queryArg)
	}},
	{"marketplace.listSales", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.ListSales(ctx)
	}},
	{"marketplace.listSellerListings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.ListSellerListings(ctx)
	}},
	{"marketplace.listWishlist", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.ListWishlist(ctx)
	}},
	{"marketplace.offerResource", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceOfferResourceBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.OfferResource(ctx, bodyArg)
	}},
	{"marketplace.openDispute", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceOpenDisputeBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.OpenDispute(ctx, a[0], bodyArg)
	}},
	{"marketplace.orderService", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceOrderServiceBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.OrderService(ctx, bodyArg)
	}},
	{"marketplace.payoutSchedule", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.PayoutSchedule(ctx)
	}},
	{"marketplace.preview", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MarketplacePreviewQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Marketplace.Preview(ctx, a[0], queryArg)
	}},
	{"marketplace.refundOrder", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceRefundOrderBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.RefundOrder(ctx, a[0], bodyArg)
	}},
	{"marketplace.removeFromWishlist", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MarketplaceRemoveFromWishlistQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Marketplace.RemoveFromWishlist(ctx, queryArg)
	}},
	{"marketplace.reportListing", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceReportListingBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.ReportListing(ctx, a[0], bodyArg)
	}},
	{"marketplace.requestRevision", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceRequestRevisionBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.RequestRevision(ctx, a[0], bodyArg)
	}},
	{"marketplace.salesReport", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MarketplaceSalesReportQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Marketplace.SalesReport(ctx, queryArg)
	}},
	{"marketplace.saveLegalPage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceSaveLegalPageBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.SaveLegalPage(ctx, bodyArg)
	}},
	{"marketplace.sellerUsage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.SellerUsage(ctx)
	}},
	{"marketplace.sendDisputeMessage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceSendDisputeMessageBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.SendDisputeMessage(ctx, a[0], bodyArg)
	}},
	{"marketplace.sendOrderMessage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceSendOrderMessageBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.SendOrderMessage(ctx, a[0], bodyArg)
	}},
	{"marketplace.showInShowcase", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceShowInShowcaseBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.ShowInShowcase(ctx, a[0], bodyArg)
	}},
	{"marketplace.unfollowShop", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Marketplace.UnfollowShop(ctx, a[0])
	}},
	{"marketplace.updateOrderStatus", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceUpdateOrderStatusBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.UpdateOrderStatus(ctx, a[0], bodyArg)
	}},
	{"marketplace.updateReview", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceUpdateReviewBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.UpdateReview(ctx, a[0], bodyArg)
	}},
	{"marketplace.updateSellerListing", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceUpdateSellerListingBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.UpdateSellerListing(ctx, a[0], bodyArg)
	}},
	{"marketplace.updateSellerProfile", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceUpdateSellerProfileBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.UpdateSellerProfile(ctx, bodyArg)
	}},
	{"marketplace.uploadDisputeAttachment", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceUploadDisputeAttachmentBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.UploadDisputeAttachment(ctx, a[0], bodyArg)
	}},
	{"marketplace.validateCoupon", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MarketplaceValidateCouponBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Marketplace.ValidateCoupon(ctx, bodyArg)
	}},
	{"messenger.addMember", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MessengerAddMemberBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Messenger.AddMember(ctx, a[0], bodyArg)
	}},
	{"messenger.addReaction", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MessengerAddReactionBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Messenger.AddReaction(ctx, a[0], bodyArg)
	}},
	{"messenger.aiDraft", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MessengerAiDraftBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Messenger.AIDraft(ctx, bodyArg)
	}},
	{"messenger.createChannel", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MessengerCreateChannelBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Messenger.CreateChannel(ctx, bodyArg)
	}},
	{"messenger.createChannelWebhook", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MessengerCreateChannelWebhookBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Messenger.CreateChannelWebhook(ctx, a[0], bodyArg)
	}},
	{"messenger.deleteChannel", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Messenger.DeleteChannel(ctx, a[0])
	}},
	{"messenger.deleteChannelWebhook", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Messenger.DeleteChannelWebhook(ctx, a[0], a[1])
	}},
	{"messenger.deleteMessage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Messenger.DeleteMessage(ctx, a[0])
	}},
	{"messenger.getChannel", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Messenger.GetChannel(ctx, a[0])
	}},
	{"messenger.getE2E", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Messenger.GetE2E(ctx, a[0])
	}},
	{"messenger.getPresence", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MessengerGetPresenceQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Messenger.GetPresence(ctx, queryArg)
	}},
	{"messenger.getSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Messenger.GetSettings(ctx)
	}},
	{"messenger.getUnread", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Messenger.GetUnread(ctx)
	}},
	{"messenger.listChannels", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MessengerListChannelsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Messenger.ListChannels(ctx, queryArg)
	}},
	{"messenger.listChannelWebhooks", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Messenger.ListChannelWebhooks(ctx, a[0])
	}},
	{"messenger.listMembers", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Messenger.ListMembers(ctx, a[0])
	}},
	{"messenger.listMessages", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MessengerListMessagesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Messenger.ListMessages(ctx, a[0], queryArg)
	}},
	{"messenger.listPinned", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Messenger.ListPinned(ctx, a[0])
	}},
	{"messenger.lookupUsers", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MessengerLookupUsersQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Messenger.LookupUsers(ctx, queryArg)
	}},
	{"messenger.markAllRead", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Messenger.MarkAllRead(ctx)
	}},
	{"messenger.markRead", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MessengerMarkReadBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Messenger.MarkRead(ctx, a[0], bodyArg)
	}},
	{"messenger.openDm", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MessengerOpenDmBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Messenger.OpenDm(ctx, bodyArg)
	}},
	{"messenger.openRecordChannel", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MessengerOpenRecordChannelBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Messenger.OpenRecordChannel(ctx, bodyArg)
	}},
	{"messenger.pin", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Messenger.Pin(ctx, a[0])
	}},
	{"messenger.removeMember", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MessengerRemoveMemberQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Messenger.RemoveMember(ctx, a[0], queryArg)
	}},
	{"messenger.revealChannelWebhook", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Messenger.RevealChannelWebhook(ctx, a[0], a[1])
	}},
	{"messenger.search", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MessengerSearchQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Messenger.Search(ctx, queryArg)
	}},
	{"messenger.sendMessage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MessengerSendMessageBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Messenger.SendMessage(ctx, a[0], bodyArg)
	}},
	{"messenger.setE2EEnabled", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MessengerSetE2EEnabledBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Messenger.SetE2EEnabled(ctx, a[0], bodyArg)
	}},
	{"messenger.setPresence", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MessengerSetPresenceBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Messenger.SetPresence(ctx, bodyArg)
	}},
	{"messenger.setTyping", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MessengerSetTypingBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Messenger.SetTyping(ctx, a[0], bodyArg)
	}},
	{"messenger.updateChannel", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MessengerUpdateChannelBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Messenger.UpdateChannel(ctx, a[0], bodyArg)
	}},
	{"messenger.updateChannelWebhook", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MessengerUpdateChannelWebhookBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Messenger.UpdateChannelWebhook(ctx, a[0], a[1], bodyArg)
	}},
	{"messenger.updateMessage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MessengerUpdateMessageBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Messenger.UpdateMessage(ctx, a[0], bodyArg)
	}},
	{"messenger.updateSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MessengerUpdateSettingsBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Messenger.UpdateSettings(ctx, bodyArg)
	}},
	{"milestones.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MilestonesGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Milestones.Get(ctx, queryArg)
	}},
	{"milestones.submitShippingAddress", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MilestonesSubmitShippingAddressBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *MilestonesSubmitShippingAddressQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Milestones.SubmitShippingAddress(ctx, a[0], bodyArg, queryArg)
	}},
	{"milestones.updateSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MilestonesUpdateSettingsBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *MilestonesUpdateSettingsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Milestones.UpdateSettings(ctx, bodyArg, queryArg)
	}},
	{"milestones.database.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MilestonesDatabaseGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Milestones.Database.Get(ctx, queryArg)
	}},
	{"milestones.database.stop", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *MilestonesDatabaseStopQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Milestones.Database.Stop(ctx, queryArg)
	}},
	{"milestones.database.sync", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *MilestonesDatabaseSyncBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *MilestonesDatabaseSyncQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Milestones.Database.Sync(ctx, bodyArg, queryArg)
	}},
	{"notifications.bulk", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *NotificationsBulkBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Notifications.Bulk(ctx, bodyArg)
	}},
	{"notifications.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *NotificationsDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Notifications.Delete(ctx, queryArg)
	}},
	{"notifications.getSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Notifications.GetSettings(ctx)
	}},
	{"notifications.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *NotificationsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Notifications.List(ctx, queryArg)
	}},
	{"notifications.markAllRead", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Notifications.MarkAllRead(ctx)
	}},
	{"notifications.markRead", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *NotificationsMarkReadBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Notifications.MarkRead(ctx, bodyArg)
	}},
	{"notifications.unreadCount", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Notifications.UnreadCount(ctx)
	}},
	{"oauth.userinfo", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.OAuth.Userinfo(ctx)
	}},
	{"oauth.connections.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.OAuth.Connections.List(ctx)
	}},
	{"oauth.connections.revoke", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.OAuth.Connections.Revoke(ctx, a[0])
	}},
	{"organizations.addMember", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *OrganizationsAddMemberBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Organizations.AddMember(ctx, a[0], bodyArg)
	}},
	{"organizations.assignTeam", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *OrganizationsAssignTeamBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Organizations.AssignTeam(ctx, a[0], bodyArg)
	}},
	{"organizations.cancelDeletion", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.CancelDeletion(ctx, a[0])
	}},
	{"organizations.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *OrganizationsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Organizations.Create(ctx, bodyArg)
	}},
	{"organizations.createFreezeWindow", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *FreezeWindowInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Organizations.CreateFreezeWindow(ctx, a[0], bodyArg)
	}},
	{"organizations.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.Delete(ctx, a[0])
	}},
	{"organizations.deleteFreezeWindow", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.DeleteFreezeWindow(ctx, a[0], a[1])
	}},
	{"organizations.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.Get(ctx, a[0])
	}},
	{"organizations.getMemberAdding", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.GetMemberAdding(ctx, a[0])
	}},
	{"organizations.getReportRecipients", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.GetReportRecipients(ctx, a[0])
	}},
	{"organizations.getScheduledDeletion", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.GetScheduledDeletion(ctx, a[0])
	}},
	{"organizations.getSeatCapacity", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.GetSeatCapacity(ctx, a[0])
	}},
	{"organizations.getSharingSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.GetSharingSettings(ctx, a[0])
	}},
	{"organizations.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *OrganizationsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Organizations.List(ctx, queryArg)
	}},
	{"organizations.listFreezeWindows", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.ListFreezeWindows(ctx, a[0])
	}},
	{"organizations.listInvitations", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.ListInvitations(ctx, a[0])
	}},
	{"organizations.listMembers", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.ListMembers(ctx, a[0])
	}},
	{"organizations.listMyMemberships", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.ListMyMemberships(ctx)
	}},
	{"organizations.listTeams", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.ListTeams(ctx, a[0])
	}},
	{"organizations.listWorkspaces", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.ListWorkspaces(ctx, a[0])
	}},
	{"organizations.removeMember", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.RemoveMember(ctx, a[0], a[1])
	}},
	{"organizations.removeTeam", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.RemoveTeam(ctx, a[0], a[1])
	}},
	{"organizations.resendInvitation", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.ResendInvitation(ctx, a[0], a[1])
	}},
	{"organizations.revokeInvitation", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.RevokeInvitation(ctx, a[0], a[1])
	}},
	{"organizations.setMemberAdding", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *OrganizationsSetMemberAddingBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Organizations.SetMemberAdding(ctx, a[0], bodyArg)
	}},
	{"organizations.setReportRecipients", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *OrganizationsSetReportRecipientsBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Organizations.SetReportRecipients(ctx, a[0], bodyArg)
	}},
	{"organizations.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *OrganizationsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Organizations.Update(ctx, a[0], bodyArg)
	}},
	{"organizations.updateFreezeWindow", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *FreezeWindowInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Organizations.UpdateFreezeWindow(ctx, a[0], a[1], bodyArg)
	}},
	{"organizations.updateMember", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *OrganizationsUpdateMemberBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Organizations.UpdateMember(ctx, a[0], a[1], bodyArg)
	}},
	{"organizations.updateSeatSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *OrganizationsUpdateSeatSettingsBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Organizations.UpdateSeatSettings(ctx, a[0], bodyArg)
	}},
	{"organizations.updateSharingSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *OrganizationsUpdateSharingSettingsBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Organizations.UpdateSharingSettings(ctx, a[0], bodyArg)
	}},
	{"organizations.versionStats", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.VersionStats(ctx, a[0])
	}},
	{"organizations.emailTemplates.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Organizations.EmailTemplates.List(ctx, a[0])
	}},
	{"organizations.emailTemplates.reset", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *OrganizationsEmailTemplatesResetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Organizations.EmailTemplates.Reset(ctx, a[0], a[1], queryArg)
	}},
	{"organizations.emailTemplates.set", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *EmailTemplateInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Organizations.EmailTemplates.Set(ctx, a[0], a[1], bodyArg)
	}},
	{"pages.abTestAction", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PagesAbTestActionBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Pages.AbTestAction(ctx, a[0], bodyArg)
	}},
	{"pages.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CreatePageInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Pages.Create(ctx, bodyArg)
	}},
	{"pages.createCollection", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PagesCreateCollectionBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Pages.CreateCollection(ctx, bodyArg)
	}},
	{"pages.createPreview", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *PagesCreatePreviewQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Pages.CreatePreview(ctx, a[0], queryArg)
	}},
	{"pages.createVersion", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PagesCreateVersionBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Pages.CreateVersion(ctx, a[0], bodyArg)
	}},
	{"pages.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Pages.Delete(ctx, a[0])
	}},
	{"pages.deleteVersion", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Pages.DeleteVersion(ctx, a[0], a[1])
	}},
	{"pages.discard", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Pages.Discard(ctx, a[0])
	}},
	{"pages.duplicate", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PagesDuplicateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *PagesDuplicateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Pages.Duplicate(ctx, a[0], bodyArg, queryArg)
	}},
	{"pages.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Pages.Get(ctx, a[0])
	}},
	{"pages.getAbTest", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Pages.GetAbTest(ctx, a[0])
	}},
	{"pages.getComponents", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Pages.GetComponents(ctx, a[0])
	}},
	{"pages.getVersion", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Pages.GetVersion(ctx, a[0], a[1])
	}},
	{"pages.listTranslations", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Pages.ListTranslations(ctx, a[0])
	}},
	{"pages.listVersions", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Pages.ListVersions(ctx, a[0])
	}},
	{"pages.restoreVersion", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PagesRestoreVersionBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Pages.RestoreVersion(ctx, a[0], a[1], bodyArg)
	}},
	{"pages.restoreVersionLegacy", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PagesRestoreVersionLegacyBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Pages.RestoreVersionLegacy(ctx, a[0], a[1], bodyArg)
	}},
	{"pages.revokePreview", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Pages.RevokePreview(ctx, a[0])
	}},
	{"pages.scan", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Pages.Scan(ctx, a[0])
	}},
	{"pages.setCollectionMode", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PagesSetCollectionModeBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Pages.SetCollectionMode(ctx, bodyArg)
	}},
	{"pages.translate", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PagesTranslateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Pages.Translate(ctx, a[0], bodyArg)
	}},
	{"pages.unverify", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Pages.Unverify(ctx, a[0])
	}},
	{"pages.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UpdatePageInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Pages.Update(ctx, a[0], bodyArg)
	}},
	{"pages.verify", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Pages.Verify(ctx, a[0])
	}},
	{"pdf.aiFill", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PdfAiFillBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.PDF.AIFill(ctx, bodyArg)
	}},
	{"pdf.compress", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PdfCompressBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.PDF.Compress(ctx, bodyArg)
	}},
	{"pdf.detectSignatureLines", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PdfDetectSignatureLinesBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.PDF.DetectSignatureLines(ctx, bodyArg)
	}},
	{"pdf.edit", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PdfEditBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.PDF.Edit(ctx, bodyArg)
	}},
	{"pdf.explain", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PdfExplainBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.PDF.Explain(ctx, bodyArg)
	}},
	{"pdf.merge", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PdfMergeBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.PDF.Merge(ctx, bodyArg)
	}},
	{"pdf.protect", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PdfProtectBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.PDF.Protect(ctx, bodyArg)
	}},
	{"pdf.split", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PdfSplitBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.PDF.Split(ctx, bodyArg)
	}},
	{"pdf.summarize", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PdfSummarizeBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.PDF.Summarize(ctx, bodyArg)
	}},
	{"pdf.unlock", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PdfUnlockBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.PDF.Unlock(ctx, bodyArg)
	}},
	{"pdf.esign.auditTrail", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *PdfEsignAuditTrailQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.PDF.Esign.AuditTrail(ctx, a[0], queryArg)
	}},
	{"pdf.esign.cancel", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PdfEsignCancelBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.PDF.Esign.Cancel(ctx, a[0], bodyArg)
	}},
	{"pdf.esign.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PdfEsignCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.PDF.Esign.Create(ctx, bodyArg)
	}},
	{"pdf.esign.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *PdfEsignListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.PDF.Esign.List(ctx, queryArg)
	}},
	{"pdf.esign.resend", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PdfEsignResendBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.PDF.Esign.Resend(ctx, a[0], bodyArg)
	}},
	{"pdf.esign.setVisibility", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PdfEsignSetVisibilityBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.PDF.Esign.SetVisibility(ctx, a[0], bodyArg)
	}},
	{"pdf.redactionAudits.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *PdfRedactionAuditsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.PDF.RedactionAudits.Get(ctx, a[0], queryArg)
	}},
	{"permissions.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CreatePermissionInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Permissions.Create(ctx, bodyArg)
	}},
	{"permissions.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Permissions.Delete(ctx, a[0])
	}},
	{"permissions.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *PermissionsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Permissions.List(ctx, queryArg)
	}},
	{"permissions.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PermissionsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Permissions.Update(ctx, a[0], bodyArg)
	}},
	{"privacyRequests.cancel", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PrivacyRequestsCancelBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.PrivacyRequests.Cancel(ctx, a[0], bodyArg)
	}},
	{"privacyRequests.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PrivacyRequestsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.PrivacyRequests.Create(ctx, bodyArg)
	}},
	{"privacyRequests.downloadExport", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *PrivacyRequestsDownloadExportQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.PrivacyRequests.DownloadExport(ctx, a[0], queryArg)
	}},
	{"privacyRequests.erase", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PrivacyRequestsEraseBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.PrivacyRequests.Erase(ctx, a[0], bodyArg)
	}},
	{"privacyRequests.generateExport", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.PrivacyRequests.GenerateExport(ctx, a[0])
	}},
	{"privacyRequests.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.PrivacyRequests.Get(ctx, a[0])
	}},
	{"privacyRequests.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *PrivacyRequestsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.PrivacyRequests.List(ctx, queryArg)
	}},
	{"privacyRequests.preview", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.PrivacyRequests.Preview(ctx, a[0])
	}},
	{"privacyRequests.verifyIdentity", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PrivacyRequestsVerifyIdentityBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.PrivacyRequests.VerifyIdentity(ctx, a[0], bodyArg)
	}},
	{"products.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Create(ctx, bodyArg, queryArg)
	}},
	{"products.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Delete(ctx, a[0], queryArg)
	}},
	{"products.export", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsExportQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Export(ctx, queryArg)
	}},
	{"products.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Get(ctx, a[0], queryArg)
	}},
	{"products.import", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsImportBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsImportQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Import(ctx, bodyArg, queryArg)
	}},
	{"products.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.List(ctx, queryArg)
	}},
	{"products.nextNumber", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsNextNumberQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.NextNumber(ctx, queryArg)
	}},
	{"products.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"products.categories.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsCategoriesCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsCategoriesCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Categories.Create(ctx, bodyArg, queryArg)
	}},
	{"products.categories.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsCategoriesDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Categories.Delete(ctx, a[0], queryArg)
	}},
	{"products.categories.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsCategoriesListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Categories.List(ctx, queryArg)
	}},
	{"products.categories.reorder", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsCategoriesReorderBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsCategoriesReorderQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Categories.Reorder(ctx, bodyArg, queryArg)
	}},
	{"products.categories.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsCategoriesUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsCategoriesUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Categories.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"products.customerGroups.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsCustomerGroupsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsCustomerGroupsCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.CustomerGroups.Create(ctx, bodyArg, queryArg)
	}},
	{"products.customerGroups.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsCustomerGroupsDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.CustomerGroups.Delete(ctx, a[0], queryArg)
	}},
	{"products.customerGroups.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsCustomerGroupsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.CustomerGroups.List(ctx, queryArg)
	}},
	{"products.customerGroups.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsCustomerGroupsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsCustomerGroupsUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.CustomerGroups.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"products.inventory.availability", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsInventoryAvailabilityBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsInventoryAvailabilityQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Availability(ctx, bodyArg, queryArg)
	}},
	{"products.inventory.batches", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsInventoryBatchesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Batches(ctx, queryArg)
	}},
	{"products.inventory.item", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsInventoryItemQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Item(ctx, a[0], queryArg)
	}},
	{"products.inventory.purchaseDrafts", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsInventoryPurchaseDraftsBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsInventoryPurchaseDraftsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.PurchaseDrafts(ctx, bodyArg, queryArg)
	}},
	{"products.inventory.releaseReservation", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsInventoryReleaseReservationBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsInventoryReleaseReservationQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.ReleaseReservation(ctx, bodyArg, queryArg)
	}},
	{"products.inventory.reorder", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsInventoryReorderQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Reorder(ctx, queryArg)
	}},
	{"products.inventory.serials", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsInventorySerialsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Serials(ctx, queryArg)
	}},
	{"products.inventory.setTracking", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsInventorySetTrackingBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsInventorySetTrackingQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.SetTracking(ctx, a[0], bodyArg, queryArg)
	}},
	{"products.inventory.stock", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsInventoryStockQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Stock(ctx, queryArg)
	}},
	{"products.inventory.transfer", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsInventoryTransferBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsInventoryTransferQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Transfer(ctx, bodyArg, queryArg)
	}},
	{"products.inventory.valuation", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsInventoryValuationQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Valuation(ctx, queryArg)
	}},
	{"products.inventory.valuationClosing", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsInventoryValuationClosingBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsInventoryValuationClosingQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.ValuationClosing(ctx, bodyArg, queryArg)
	}},
	{"products.inventory.counts.cancel", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsInventoryCountsCancelQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Counts.Cancel(ctx, a[0], queryArg)
	}},
	{"products.inventory.counts.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsInventoryCountsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsInventoryCountsCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Counts.Create(ctx, bodyArg, queryArg)
	}},
	{"products.inventory.counts.enter", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsInventoryCountsEnterBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsInventoryCountsEnterQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Counts.Enter(ctx, a[0], bodyArg, queryArg)
	}},
	{"products.inventory.counts.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsInventoryCountsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Counts.Get(ctx, a[0], queryArg)
	}},
	{"products.inventory.counts.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsInventoryCountsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Counts.List(ctx, queryArg)
	}},
	{"products.inventory.counts.post", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsInventoryCountsPostBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsInventoryCountsPostQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Counts.Post(ctx, a[0], bodyArg, queryArg)
	}},
	{"products.inventory.counts.sheet", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsInventoryCountsSheetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Counts.Sheet(ctx, a[0], queryArg)
	}},
	{"products.inventory.locations.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsInventoryLocationsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsInventoryLocationsCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Locations.Create(ctx, a[0], bodyArg, queryArg)
	}},
	{"products.inventory.locations.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsInventoryLocationsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsInventoryLocationsUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Locations.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"products.inventory.movements.book", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsInventoryMovementsBookBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsInventoryMovementsBookQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Movements.Book(ctx, bodyArg, queryArg)
	}},
	{"products.inventory.movements.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsInventoryMovementsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Movements.List(ctx, queryArg)
	}},
	{"products.inventory.reorderRules.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsInventoryReorderRulesDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.ReorderRules.Delete(ctx, a[0], queryArg)
	}},
	{"products.inventory.reorderRules.save", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsInventoryReorderRulesSaveBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsInventoryReorderRulesSaveQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.ReorderRules.Save(ctx, bodyArg, queryArg)
	}},
	{"products.inventory.settings.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsInventorySettingsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Settings.Get(ctx, queryArg)
	}},
	{"products.inventory.settings.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsInventorySettingsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsInventorySettingsUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Settings.Update(ctx, bodyArg, queryArg)
	}},
	{"products.inventory.warehouses.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsInventoryWarehousesCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsInventoryWarehousesCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Warehouses.Create(ctx, bodyArg, queryArg)
	}},
	{"products.inventory.warehouses.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsInventoryWarehousesListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Warehouses.List(ctx, queryArg)
	}},
	{"products.inventory.warehouses.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsInventoryWarehousesUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsInventoryWarehousesUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Inventory.Warehouses.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"products.marketplace.bundleItems", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplaceBundleItemsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.BundleItems(ctx, queryArg)
	}},
	{"products.marketplace.getOffer", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplaceGetOfferQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.GetOffer(ctx, a[0], queryArg)
	}},
	{"products.marketplace.offers", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplaceOffersQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Offers(ctx, queryArg)
	}},
	{"products.marketplace.overview", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplaceOverviewQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Overview(ctx, queryArg)
	}},
	{"products.marketplace.pause", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplacePauseQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Pause(ctx, a[0], queryArg)
	}},
	{"products.marketplace.payouts", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplacePayoutsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Payouts(ctx, queryArg)
	}},
	{"products.marketplace.preview", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplacePreviewQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Preview(ctx, a[0], queryArg)
	}},
	{"products.marketplace.questions", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplaceQuestionsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Questions(ctx, queryArg)
	}},
	{"products.marketplace.resume", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplaceResumeQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Resume(ctx, a[0], queryArg)
	}},
	{"products.marketplace.sales", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplaceSalesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Sales(ctx, queryArg)
	}},
	{"products.marketplace.saveOffer", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg map[string]any
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsMarketplaceSaveOfferQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.SaveOffer(ctx, a[0], bodyArg, queryArg)
	}},
	{"products.marketplace.sources", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplaceSourcesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Sources(ctx, queryArg)
	}},
	{"products.marketplace.statement", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplaceStatementQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Statement(ctx, a[0], queryArg)
	}},
	{"products.marketplace.submit", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsMarketplaceSubmitBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsMarketplaceSubmitQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Submit(ctx, a[0], bodyArg, queryArg)
	}},
	{"products.marketplace.coupons.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsMarketplaceCouponsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsMarketplaceCouponsCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Coupons.Create(ctx, bodyArg, queryArg)
	}},
	{"products.marketplace.coupons.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplaceCouponsDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Coupons.Delete(ctx, queryArg)
	}},
	{"products.marketplace.coupons.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplaceCouponsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Coupons.List(ctx, queryArg)
	}},
	{"products.marketplace.legalPages.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplaceLegalPagesListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.LegalPages.List(ctx, queryArg)
	}},
	{"products.marketplace.legalPages.save", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsMarketplaceLegalPagesSaveBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsMarketplaceLegalPagesSaveQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.LegalPages.Save(ctx, bodyArg, queryArg)
	}},
	{"products.marketplace.orders.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplaceOrdersGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Orders.Get(ctx, a[0], queryArg)
	}},
	{"products.marketplace.orders.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplaceOrdersListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Orders.List(ctx, queryArg)
	}},
	{"products.marketplace.orders.message", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsMarketplaceOrdersMessageBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsMarketplaceOrdersMessageQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Orders.Message(ctx, a[0], bodyArg, queryArg)
	}},
	{"products.marketplace.orders.refund", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsMarketplaceOrdersRefundBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsMarketplaceOrdersRefundQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Orders.Refund(ctx, a[0], bodyArg, queryArg)
	}},
	{"products.marketplace.payoutAccount.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplacePayoutAccountGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.PayoutAccount.Get(ctx, queryArg)
	}},
	{"products.marketplace.payoutAccount.onboard", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplacePayoutAccountOnboardQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.PayoutAccount.Onboard(ctx, queryArg)
	}},
	{"products.marketplace.reviews.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplaceReviewsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Reviews.List(ctx, queryArg)
	}},
	{"products.marketplace.reviews.respond", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsMarketplaceReviewsRespondBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsMarketplaceReviewsRespondQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Reviews.Respond(ctx, bodyArg, queryArg)
	}},
	{"products.marketplace.settings.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsMarketplaceSettingsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Settings.Get(ctx, queryArg)
	}},
	{"products.marketplace.settings.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsMarketplaceSettingsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsMarketplaceSettingsUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Marketplace.Settings.Update(ctx, bodyArg, queryArg)
	}},
	{"products.priceLists.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductPriceListInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsPriceListsCreateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.PriceLists.Create(ctx, bodyArg, queryArg)
	}},
	{"products.priceLists.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsPriceListsDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.PriceLists.Delete(ctx, a[0], queryArg)
	}},
	{"products.priceLists.export", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsPriceListsExportQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.PriceLists.Export(ctx, a[0], queryArg)
	}},
	{"products.priceLists.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsPriceListsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.PriceLists.Get(ctx, a[0], queryArg)
	}},
	{"products.priceLists.import", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductCsvImportInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsPriceListsImportQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.PriceLists.Import(ctx, a[0], bodyArg, queryArg)
	}},
	{"products.priceLists.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsPriceListsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.PriceLists.List(ctx, queryArg)
	}},
	{"products.priceLists.setPrices", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsPriceListsSetPricesBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsPriceListsSetPricesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.PriceLists.SetPrices(ctx, a[0], bodyArg, queryArg)
	}},
	{"products.priceLists.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductPriceListInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsPriceListsUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.PriceLists.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"products.settings.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsSettingsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Settings.Get(ctx, queryArg)
	}},
	{"products.settings.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsSettingsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsSettingsUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Settings.Update(ctx, bodyArg, queryArg)
	}},
	{"products.variants.export", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsVariantsExportQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Variants.Export(ctx, queryArg)
	}},
	{"products.variants.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProductsVariantsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Variants.Get(ctx, a[0], queryArg)
	}},
	{"products.variants.import", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductCsvImportInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsVariantsImportQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Variants.Import(ctx, bodyArg, queryArg)
	}},
	{"products.variants.setOptions", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsVariantsSetOptionsBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsVariantsSetOptionsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Variants.SetOptions(ctx, a[0], bodyArg, queryArg)
	}},
	{"products.variants.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProductsVariantsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *ProductsVariantsUpdateQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Products.Variants.Update(ctx, a[0], bodyArg, queryArg)
	}},
	{"projects.addDomain", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AddDomainInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.AddDomain(ctx, a[0], bodyArg)
	}},
	{"projects.addLanguage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsAddLanguageBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.AddLanguage(ctx, a[0], bodyArg)
	}},
	{"projects.analytics", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsAnalyticsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.Analytics(ctx, a[0], queryArg)
	}},
	{"projects.claimPlatformDomain", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.ClaimPlatformDomain(ctx, a[0])
	}},
	{"projects.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CreateProjectInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.Create(ctx, bodyArg)
	}},
	{"projects.createComment", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsCreateCommentBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.CreateComment(ctx, a[0], bodyArg)
	}},
	{"projects.createPreview", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsCreatePreviewBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.CreatePreview(ctx, a[0], bodyArg)
	}},
	{"projects.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.Delete(ctx, a[0])
	}},
	{"projects.deleteComment", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.DeleteComment(ctx, a[0], a[1])
	}},
	{"projects.downloadExport", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.DownloadExport(ctx, a[0], a[1])
	}},
	{"projects.export", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.Export(ctx, a[0])
	}},
	{"projects.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.Get(ctx, a[0])
	}},
	{"projects.getBrain", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.GetBrain(ctx, a[0])
	}},
	{"projects.getComingSoon", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.GetComingSoon(ctx, a[0])
	}},
	{"projects.getNotFoundPage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.GetNotFoundPage(ctx, a[0])
	}},
	{"projects.getStyles", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.GetStyles(ctx, a[0])
	}},
	{"projects.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.List(ctx, queryArg)
	}},
	{"projects.listCollaborators", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.ListCollaborators(ctx, a[0])
	}},
	{"projects.listComments", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsListCommentsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.ListComments(ctx, a[0], queryArg)
	}},
	{"projects.listDomains", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.ListDomains(ctx, a[0])
	}},
	{"projects.listExports", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.ListExports(ctx, a[0])
	}},
	{"projects.listLanguages", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.ListLanguages(ctx, a[0])
	}},
	{"projects.listPages", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsListPagesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.ListPages(ctx, a[0], queryArg)
	}},
	{"projects.publish", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.Publish(ctx, a[0])
	}},
	{"projects.removeDomain", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.RemoveDomain(ctx, a[0], a[1])
	}},
	{"projects.removeLanguage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsRemoveLanguageQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.RemoveLanguage(ctx, a[0], queryArg)
	}},
	{"projects.requestApproval", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsRequestApprovalBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.RequestApproval(ctx, a[0], bodyArg)
	}},
	{"projects.revokePreview", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.RevokePreview(ctx, a[0])
	}},
	{"projects.setDefaultLanguage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsSetDefaultLanguageBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.SetDefaultLanguage(ctx, a[0], bodyArg)
	}},
	{"projects.setHomePage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.SetHomePage(ctx, a[0], a[1])
	}},
	{"projects.setPrimaryDomain", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.SetPrimaryDomain(ctx, a[0], a[1])
	}},
	{"projects.suggestBrain", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.SuggestBrain(ctx, a[0])
	}},
	{"projects.traffic", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsTrafficQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.Traffic(ctx, a[0], queryArg)
	}},
	{"projects.unsetHomePage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.UnsetHomePage(ctx, a[0], a[1])
	}},
	{"projects.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UpdateProjectInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.Update(ctx, a[0], bodyArg)
	}},
	{"projects.updateBrain", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsUpdateBrainBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.UpdateBrain(ctx, a[0], bodyArg)
	}},
	{"projects.updateComingSoon", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsUpdateComingSoonBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.UpdateComingSoon(ctx, a[0], bodyArg)
	}},
	{"projects.updateComment", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsUpdateCommentBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.UpdateComment(ctx, a[0], a[1], bodyArg)
	}},
	{"projects.updateNotFoundPage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsUpdateNotFoundPageBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.UpdateNotFoundPage(ctx, a[0], bodyArg)
	}},
	{"projects.updateStyles", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg map[string]any
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.UpdateStyles(ctx, a[0], bodyArg)
	}},
	{"projects.verifyDomain", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.VerifyDomain(ctx, a[0], a[1])
	}},
	{"projects.videoAnalytics", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.VideoAnalytics(ctx, a[0])
	}},
	{"projects.contentGaps.faq", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsContentGapsFaqBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.ContentGaps.Faq(ctx, a[0], a[1], bodyArg)
	}},
	{"projects.contentGaps.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsContentGapsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.ContentGaps.List(ctx, a[0], queryArg)
	}},
	{"projects.contentGaps.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsContentGapsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.ContentGaps.Update(ctx, a[0], a[1], bodyArg)
	}},
	{"projects.cookies.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectCookieInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.Cookies.Create(ctx, a[0], bodyArg)
	}},
	{"projects.cookies.createCategory", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CookieCategoryInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.Cookies.CreateCategory(ctx, a[0], bodyArg)
	}},
	{"projects.cookies.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.Cookies.Delete(ctx, a[0], a[1])
	}},
	{"projects.cookies.deleteCategory", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.Cookies.DeleteCategory(ctx, a[0], a[1])
	}},
	{"projects.cookies.getScan", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.Cookies.GetScan(ctx, a[0], a[1])
	}},
	{"projects.cookies.getSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.Cookies.GetSettings(ctx, a[0])
	}},
	{"projects.cookies.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.Cookies.List(ctx, a[0])
	}},
	{"projects.cookies.listCategories", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.Cookies.ListCategories(ctx, a[0])
	}},
	{"projects.cookies.listConsents", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsCookiesListConsentsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.Cookies.ListConsents(ctx, a[0], queryArg)
	}},
	{"projects.cookies.scan", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsCookiesScanBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.Cookies.Scan(ctx, a[0], bodyArg)
	}},
	{"projects.cookies.seed", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.Cookies.Seed(ctx, a[0])
	}},
	{"projects.cookies.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsCookiesUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.Cookies.Update(ctx, a[0], a[1], bodyArg)
	}},
	{"projects.cookies.updateCategory", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsCookiesUpdateCategoryBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.Cookies.UpdateCategory(ctx, a[0], a[1], bodyArg)
	}},
	{"projects.cookies.updateSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsCookiesUpdateSettingsBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.Cookies.UpdateSettings(ctx, a[0], bodyArg)
	}},
	{"projects.emailTemplates.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.EmailTemplates.List(ctx, a[0])
	}},
	{"projects.emailTemplates.reset", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsEmailTemplatesResetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.EmailTemplates.Reset(ctx, a[0], a[1], queryArg)
	}},
	{"projects.emailTemplates.set", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *EmailTemplateInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.EmailTemplates.Set(ctx, a[0], a[1], bodyArg)
	}},
	{"projects.journeys.getSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.Journeys.GetSettings(ctx, a[0])
	}},
	{"projects.journeys.getVisitor", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.Journeys.GetVisitor(ctx, a[0], a[1])
	}},
	{"projects.journeys.listDeliveries", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsJourneysListDeliveriesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.Journeys.ListDeliveries(ctx, a[0], queryArg)
	}},
	{"projects.journeys.listVisitors", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsJourneysListVisitorsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.Journeys.ListVisitors(ctx, a[0], queryArg)
	}},
	{"projects.journeys.report", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsJourneysReportQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.Journeys.Report(ctx, a[0], queryArg)
	}},
	{"projects.journeys.updateSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg map[string]any
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.Journeys.UpdateSettings(ctx, a[0], bodyArg)
	}},
	{"projects.pageFolders.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsPageFoldersCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.PageFolders.Create(ctx, a[0], bodyArg)
	}},
	{"projects.pageFolders.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.PageFolders.Delete(ctx, a[0], a[1])
	}},
	{"projects.pageFolders.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.PageFolders.List(ctx, a[0])
	}},
	{"projects.pageFolders.move", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsPageFoldersMoveBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.PageFolders.Move(ctx, a[0], bodyArg)
	}},
	{"projects.pageFolders.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsPageFoldersUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.PageFolders.Update(ctx, a[0], a[1], bodyArg)
	}},
	{"projects.redirects.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectRedirectInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.Redirects.Create(ctx, a[0], bodyArg)
	}},
	{"projects.redirects.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.Redirects.Delete(ctx, a[0], a[1])
	}},
	{"projects.redirects.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.Redirects.List(ctx, a[0])
	}},
	{"projects.redirects.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsRedirectsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.Redirects.Update(ctx, a[0], a[1], bodyArg)
	}},
	{"projects.results.audiences", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsResultsAudiencesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.Results.Audiences(ctx, a[0], queryArg)
	}},
	{"projects.results.engagement", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsResultsEngagementQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.Results.Engagement(ctx, a[0], queryArg)
	}},
	{"projects.results.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsResultsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.Results.Get(ctx, a[0], queryArg)
	}},
	{"projects.results.impact", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsResultsImpactQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.Results.Impact(ctx, a[0], queryArg)
	}},
	{"projects.results.updateSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsResultsUpdateSettingsBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.Results.UpdateSettings(ctx, a[0], bodyArg)
	}},
	{"projects.sharedBlocks.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsSharedBlocksCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.SharedBlocks.Create(ctx, a[0], bodyArg)
	}},
	{"projects.sharedBlocks.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsSharedBlocksDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.SharedBlocks.Delete(ctx, a[0], a[1], queryArg)
	}},
	{"projects.sharedBlocks.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.SharedBlocks.Get(ctx, a[0], a[1])
	}},
	{"projects.sharedBlocks.getVersion", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.SharedBlocks.GetVersion(ctx, a[0], a[1], a[2])
	}},
	{"projects.sharedBlocks.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsSharedBlocksListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.SharedBlocks.List(ctx, a[0], queryArg)
	}},
	{"projects.sharedBlocks.publish", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsSharedBlocksPublishBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.SharedBlocks.Publish(ctx, a[0], a[1], bodyArg)
	}},
	{"projects.sharedBlocks.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsSharedBlocksUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.SharedBlocks.Update(ctx, a[0], a[1], bodyArg)
	}},
	{"projects.siteHealth.check", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.SiteHealth.Check(ctx, a[0])
	}},
	{"projects.siteHealth.dismiss", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.SiteHealth.Dismiss(ctx, a[0], a[1])
	}},
	{"projects.siteHealth.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsSiteHealthGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.SiteHealth.Get(ctx, a[0], queryArg)
	}},
	{"projects.spamProtection.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.SpamProtection.Get(ctx, a[0])
	}},
	{"projects.spamProtection.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ProjectsSpamProtectionUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Projects.SpamProtection.Update(ctx, a[0], bodyArg)
	}},
	{"projects.waitlist.deleteSignup", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Projects.Waitlist.DeleteSignup(ctx, a[0], a[1])
	}},
	{"projects.waitlist.export", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsWaitlistExportQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.Waitlist.Export(ctx, a[0], queryArg)
	}},
	{"projects.waitlist.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ProjectsWaitlistListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Projects.Waitlist.List(ctx, a[0], queryArg)
	}},
	{"push.getPublicKey", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Push.GetPublicKey(ctx)
	}},
	{"push.subscribe", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PushSubscribeBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Push.Subscribe(ctx, bodyArg)
	}},
	{"push.unsubscribe", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *PushUnsubscribeQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Push.Unsubscribe(ctx, queryArg)
	}},
	{"referrals.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ReferralsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Referrals.Get(ctx, queryArg)
	}},
	{"referrals.regenerateCode", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ReferralsRegenerateCodeQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Referrals.RegenerateCode(ctx, queryArg)
	}},
	{"regionalSettings.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.RegionalSettings.Get(ctx)
	}},
	{"reports.annotations", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ReportsAnnotationsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Reports.Annotations(ctx, a[0], queryArg)
	}},
	{"reports.consumers", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Reports.Consumers(ctx, a[0])
	}},
	{"reports.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ReportsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Reports.Create(ctx, bodyArg)
	}},
	{"reports.data", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ReportsDataBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Reports.Data(ctx, a[0], bodyArg)
	}},
	{"reports.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Reports.Delete(ctx, a[0])
	}},
	{"reports.export", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ReportsExportQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Reports.Export(ctx, a[0], queryArg)
	}},
	{"reports.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Reports.Get(ctx, a[0])
	}},
	{"reports.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ReportsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Reports.List(ctx, queryArg)
	}},
	{"reports.sourceHealth", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Reports.SourceHealth(ctx, a[0])
	}},
	{"reports.sourceHealthRollup", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *ReportsSourceHealthRollupQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Reports.SourceHealthRollup(ctx, queryArg)
	}},
	{"reports.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ReportsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Reports.Update(ctx, a[0], bodyArg)
	}},
	{"reports.publicLinks.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ReportsPublicLinksCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Reports.PublicLinks.Create(ctx, a[0], bodyArg)
	}},
	{"reports.publicLinks.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Reports.PublicLinks.List(ctx, a[0])
	}},
	{"reports.publicLinks.revoke", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Reports.PublicLinks.Revoke(ctx, a[0], a[1])
	}},
	{"reports.schedules.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ReportScheduleInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Reports.Schedules.Create(ctx, a[0], bodyArg)
	}},
	{"reports.schedules.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Reports.Schedules.Delete(ctx, a[0], a[1])
	}},
	{"reports.schedules.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Reports.Schedules.List(ctx, a[0])
	}},
	{"reports.schedules.sendNow", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Reports.Schedules.SendNow(ctx, a[0], a[1])
	}},
	{"reports.schedules.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ReportsSchedulesUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Reports.Schedules.Update(ctx, a[0], a[1], bodyArg)
	}},
	{"reports.shares.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ReportsSharesCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Reports.Shares.Create(ctx, a[0], bodyArg)
	}},
	{"reports.shares.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Reports.Shares.Delete(ctx, a[0], a[1])
	}},
	{"reports.shares.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Reports.Shares.List(ctx, a[0])
	}},
	{"reports.shares.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ReportsSharesUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Reports.Shares.Update(ctx, a[0], a[1], bodyArg)
	}},
	{"reports.snapshots.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ReportsSnapshotsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Reports.Snapshots.Create(ctx, a[0], bodyArg)
	}},
	{"reports.snapshots.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Reports.Snapshots.Get(ctx, a[0], a[1])
	}},
	{"reports.snapshots.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Reports.Snapshots.List(ctx, a[0])
	}},
	{"reports.snapshots.sign", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Reports.Snapshots.Sign(ctx, a[0], a[1])
	}},
	{"reports.snapshots.verify", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Reports.Snapshots.Verify(ctx, a[0], a[1])
	}},
	{"reports.widgets.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CreateReportWidgetInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Reports.Widgets.Create(ctx, a[0], bodyArg)
	}},
	{"reports.widgets.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Reports.Widgets.Delete(ctx, a[0], a[1])
	}},
	{"reports.widgets.history", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Reports.Widgets.History(ctx, a[0], a[1])
	}},
	{"reports.widgets.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Reports.Widgets.List(ctx, a[0])
	}},
	{"reports.widgets.observe", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ReportsWidgetsObserveBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Reports.Widgets.Observe(ctx, a[0], a[1], bodyArg)
	}},
	{"reports.widgets.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UpdateReportWidgetInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Reports.Widgets.Update(ctx, a[0], a[1], bodyArg)
	}},
	{"reports.widgets.updateLayout", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *ReportsWidgetsUpdateLayoutBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Reports.Widgets.UpdateLayout(ctx, a[0], bodyArg)
	}},
	{"retentionPolicies.check", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *RetentionPoliciesCheckQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.RetentionPolicies.Check(ctx, queryArg)
	}},
	{"retentionPolicies.checkCascade", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *RetentionPoliciesCheckCascadeQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.RetentionPolicies.CheckCascade(ctx, queryArg)
	}},
	{"retentionPolicies.confirm", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.RetentionPolicies.Confirm(ctx, a[0])
	}},
	{"retentionPolicies.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *RetentionPoliciesCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.RetentionPolicies.Create(ctx, bodyArg)
	}},
	{"retentionPolicies.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.RetentionPolicies.Delete(ctx, a[0])
	}},
	{"retentionPolicies.dryRun", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *RetentionPoliciesDryRunBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.RetentionPolicies.DryRun(ctx, a[0], bodyArg)
	}},
	{"retentionPolicies.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.RetentionPolicies.Get(ctx, a[0])
	}},
	{"retentionPolicies.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *RetentionPoliciesListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.RetentionPolicies.List(ctx, queryArg)
	}},
	{"retentionPolicies.listDryRuns", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.RetentionPolicies.ListDryRuns(ctx, a[0])
	}},
	{"retentionPolicies.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *RetentionPoliciesUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.RetentionPolicies.Update(ctx, a[0], bodyArg)
	}},
	{"roles.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CreateRoleInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Roles.Create(ctx, bodyArg)
	}},
	{"roles.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Roles.Delete(ctx, a[0])
	}},
	{"roles.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Roles.Get(ctx, a[0])
	}},
	{"roles.grantable", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *RolesGrantableQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Roles.Grantable(ctx, queryArg)
	}},
	{"roles.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *RolesListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Roles.List(ctx, queryArg)
	}},
	{"roles.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UpdateRoleInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Roles.Update(ctx, a[0], bodyArg)
	}},
	{"roles.assignments.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *RolesAssignmentsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Roles.Assignments.Create(ctx, bodyArg)
	}},
	{"roles.assignments.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Roles.Assignments.Delete(ctx, a[0])
	}},
	{"roles.assignments.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *RolesAssignmentsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Roles.Assignments.List(ctx, queryArg)
	}},
	{"share.getAsset", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Share.GetAsset(ctx, a[0])
	}},
	{"social.accountInsights", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SocialAccountInsightsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Social.AccountInsights(ctx, a[0], queryArg)
	}},
	{"social.addComment", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialAddCommentBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Social.AddComment(ctx, a[0], bodyArg)
	}},
	{"social.aiCaption", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialAiCaptionBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Social.AICaption(ctx, bodyArg)
	}},
	{"social.aiHashtags", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialAiHashtagsBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Social.AIHashtags(ctx, bodyArg)
	}},
	{"social.aiVariations", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialAiVariationsBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Social.AIVariations(ctx, bodyArg)
	}},
	{"social.analyticsSummary", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SocialAnalyticsSummaryQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Social.AnalyticsSummary(ctx, queryArg)
	}},
	{"social.assign", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialAssignBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Social.Assign(ctx, a[0], bodyArg)
	}},
	{"social.bestTime", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialBestTimeBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Social.BestTime(ctx, bodyArg)
	}},
	{"social.bulkCreate", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialBulkCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Social.BulkCreate(ctx, bodyArg)
	}},
	{"social.bulkDelete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialBulkDeleteBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Social.BulkDelete(ctx, bodyArg)
	}},
	{"social.bulkUpdate", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialBulkUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Social.BulkUpdate(ctx, bodyArg)
	}},
	{"social.cancel", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Social.Cancel(ctx, a[0])
	}},
	{"social.clonePost", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Social.ClonePost(ctx, a[0])
	}},
	{"social.connectBluesky", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialConnectBlueskyBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Social.ConnectBluesky(ctx, bodyArg)
	}},
	{"social.createAbTest", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialCreateAbTestBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *SocialCreateAbTestQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Social.CreateAbTest(ctx, bodyArg, queryArg)
	}},
	{"social.createCampaign", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialCreateCampaignBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Social.CreateCampaign(ctx, bodyArg)
	}},
	{"social.createHashtagSet", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialCreateHashtagSetBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Social.CreateHashtagSet(ctx, bodyArg)
	}},
	{"social.createPost", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CreateSocialPostInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *SocialCreatePostQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Social.CreatePost(ctx, bodyArg, queryArg)
	}},
	{"social.createSnippet", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialCreateSnippetBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *SocialCreateSnippetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Social.CreateSnippet(ctx, bodyArg, queryArg)
	}},
	{"social.createTemplate", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialCreateTemplateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Social.CreateTemplate(ctx, bodyArg)
	}},
	{"social.deleteAutoRepost", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Social.DeleteAutoRepost(ctx, a[0])
	}},
	{"social.deleteCampaign", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Social.DeleteCampaign(ctx, a[0])
	}},
	{"social.deleteComment", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Social.DeleteComment(ctx, a[0], a[1])
	}},
	{"social.deleteHashtagSet", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Social.DeleteHashtagSet(ctx, a[0])
	}},
	{"social.deletePost", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Social.DeletePost(ctx, a[0])
	}},
	{"social.deleteSnippet", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Social.DeleteSnippet(ctx, a[0])
	}},
	{"social.deleteTemplate", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Social.DeleteTemplate(ctx, a[0])
	}},
	{"social.disconnectAccount", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Social.DisconnectAccount(ctx, a[0])
	}},
	{"social.draftFromContent", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialDraftFromContentBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Social.DraftFromContent(ctx, bodyArg)
	}},
	{"social.getAssignment", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Social.GetAssignment(ctx, a[0])
	}},
	{"social.getAutoRepost", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Social.GetAutoRepost(ctx, a[0])
	}},
	{"social.getPost", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Social.GetPost(ctx, a[0])
	}},
	{"social.limits", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SocialLimitsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Social.Limits(ctx, queryArg)
	}},
	{"social.listAccounts", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SocialListAccountsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Social.ListAccounts(ctx, queryArg)
	}},
	{"social.listCampaigns", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SocialListCampaignsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Social.ListCampaigns(ctx, queryArg)
	}},
	{"social.listComments", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Social.ListComments(ctx, a[0])
	}},
	{"social.listHashtagSets", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SocialListHashtagSetsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Social.ListHashtagSets(ctx, queryArg)
	}},
	{"social.listInbox", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SocialListInboxQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Social.ListInbox(ctx, queryArg)
	}},
	{"social.listPosts", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SocialListPostsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Social.ListPosts(ctx, queryArg)
	}},
	{"social.listSnippets", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SocialListSnippetsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Social.ListSnippets(ctx, queryArg)
	}},
	{"social.listTemplates", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SocialListTemplatesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Social.ListTemplates(ctx, queryArg)
	}},
	{"social.markInboxRead", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Social.MarkInboxRead(ctx, a[0])
	}},
	{"social.markInboxUnread", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Social.MarkInboxUnread(ctx, a[0])
	}},
	{"social.patchPost", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialPatchPostBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Social.PatchPost(ctx, a[0], bodyArg)
	}},
	{"social.pollInbox", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SocialPollInboxQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Social.PollInbox(ctx, queryArg)
	}},
	{"social.postInsights", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SocialPostInsightsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Social.PostInsights(ctx, a[0], queryArg)
	}},
	{"social.publishNow", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Social.PublishNow(ctx, a[0])
	}},
	{"social.replyInbox", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialReplyInboxBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Social.ReplyInbox(ctx, a[0], bodyArg)
	}},
	{"social.retryStory", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Social.RetryStory(ctx, a[0])
	}},
	{"social.schedule", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialScheduleBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Social.Schedule(ctx, a[0], bodyArg)
	}},
	{"social.setAutoRepost", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialSetAutoRepostBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Social.SetAutoRepost(ctx, a[0], bodyArg)
	}},
	{"social.shareCheck", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SocialShareCheckQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Social.ShareCheck(ctx, queryArg)
	}},
	{"social.unassign", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Social.Unassign(ctx, a[0])
	}},
	{"social.updateAccount", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialUpdateAccountBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Social.UpdateAccount(ctx, a[0], bodyArg)
	}},
	{"social.updateSnippet", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SocialUpdateSnippetBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Social.UpdateSnippet(ctx, a[0], bodyArg)
	}},
	{"submissions.analytics", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SubmissionsAnalyticsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Submissions.Analytics(ctx, queryArg)
	}},
	{"submissions.bulkDelete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SubmissionsBulkDeleteQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Submissions.BulkDelete(ctx, queryArg)
	}},
	{"submissions.bulkUpdate", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SubmissionsBulkUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Submissions.BulkUpdate(ctx, bodyArg)
	}},
	{"submissions.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SubmissionsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Submissions.Create(ctx, bodyArg)
	}},
	{"submissions.createCannedResponse", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SubmissionsCreateCannedResponseBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Submissions.CreateCannedResponse(ctx, bodyArg)
	}},
	{"submissions.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Submissions.Delete(ctx, a[0])
	}},
	{"submissions.deleteCannedResponse", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Submissions.DeleteCannedResponse(ctx, a[0])
	}},
	{"submissions.enrich", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Submissions.Enrich(ctx, a[0])
	}},
	{"submissions.erasePersonalData", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SubmissionsErasePersonalDataQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Submissions.ErasePersonalData(ctx, queryArg)
	}},
	{"submissions.export", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SubmissionsExportQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Submissions.Export(ctx, queryArg)
	}},
	{"submissions.exportPersonalData", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SubmissionsExportPersonalDataBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Submissions.ExportPersonalData(ctx, bodyArg)
	}},
	{"submissions.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Submissions.Get(ctx, a[0])
	}},
	{"submissions.getFormSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Submissions.GetFormSettings(ctx, a[0], a[1])
	}},
	{"submissions.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SubmissionsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Submissions.List(ctx, queryArg)
	}},
	{"submissions.listCannedResponses", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SubmissionsListCannedResponsesQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Submissions.ListCannedResponses(ctx, queryArg)
	}},
	{"submissions.listForms", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SubmissionsListFormsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Submissions.ListForms(ctx, queryArg)
	}},
	{"submissions.reply", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SubmissionsReplyBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Submissions.Reply(ctx, a[0], bodyArg)
	}},
	{"submissions.setFormVisibility", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SubmissionsSetFormVisibilityBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Submissions.SetFormVisibility(ctx, a[0], a[1], bodyArg)
	}},
	{"submissions.thread", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Submissions.Thread(ctx, a[0])
	}},
	{"submissions.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SubmissionsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Submissions.Update(ctx, a[0], bodyArg)
	}},
	{"submissions.updateCannedResponse", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SubmissionsUpdateCannedResponseBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Submissions.UpdateCannedResponse(ctx, a[0], bodyArg)
	}},
	{"submissions.updateFormSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SubmissionsUpdateFormSettingsBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Submissions.UpdateFormSettings(ctx, a[0], a[1], bodyArg)
	}},
	{"subscriptions.cancel", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SubscriptionsCancelBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Subscriptions.Cancel(ctx, bodyArg)
	}},
	{"subscriptions.resume", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SubscriptionsResumeQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Subscriptions.Resume(ctx, queryArg)
	}},
	{"support.createTicket", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SupportCreateTicketBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Support.CreateTicket(ctx, bodyArg)
	}},
	{"support.getTicket", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Support.GetTicket(ctx, a[0])
	}},
	{"support.listMessages", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Support.ListMessages(ctx, a[0])
	}},
	{"support.listTickets", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *SupportListTicketsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Support.ListTickets(ctx, queryArg)
	}},
	{"support.postMessage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SupportPostMessageBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Support.PostMessage(ctx, a[0], bodyArg)
	}},
	{"support.updateTicket", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SupportUpdateTicketBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Support.UpdateTicket(ctx, a[0], bodyArg)
	}},
	{"support.upload", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *SupportUploadBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Support.Upload(ctx, bodyArg)
	}},
	{"tasks.bulk", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TaskBulkInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.Bulk(ctx, bodyArg)
	}},
	{"tasks.complete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg any
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.Complete(ctx, bodyArg)
	}},
	{"tasks.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TaskCreateInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.Create(ctx, bodyArg)
	}},
	{"tasks.estimate", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TasksEstimateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.Estimate(ctx, bodyArg)
	}},
	{"tasks.fromSource", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TasksFromSourceQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Tasks.FromSource(ctx, queryArg)
	}},
	{"tasks.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TasksListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Tasks.List(ctx, queryArg)
	}},
	{"tasks.nextUp", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TasksNextUpQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Tasks.NextUp(ctx, queryArg)
	}},
	{"tasks.pin", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TasksPinBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.Pin(ctx, bodyArg)
	}},
	{"tasks.snooze", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TasksSnoozeBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.Snooze(ctx, bodyArg)
	}},
	{"tasks.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TaskUpdateInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.Update(ctx, bodyArg)
	}},
	{"tasks.workload", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TasksWorkloadQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Tasks.Workload(ctx, queryArg)
	}},
	{"tasks.automation.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TasksAutomationGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Tasks.Automation.Get(ctx, queryArg)
	}},
	{"tasks.automation.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TasksAutomationUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.Automation.Update(ctx, bodyArg)
	}},
	{"tasks.checklist.convert", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TasksChecklistConvertBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.Checklist.Convert(ctx, bodyArg)
	}},
	{"tasks.conditions.add", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TaskConditionInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.Conditions.Add(ctx, bodyArg)
	}},
	{"tasks.conditions.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TasksConditionsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Tasks.Conditions.List(ctx, queryArg)
	}},
	{"tasks.conditions.options", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TasksConditionsOptionsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Tasks.Conditions.Options(ctx, queryArg)
	}},
	{"tasks.conditions.remove", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TasksConditionsRemoveQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Tasks.Conditions.Remove(ctx, queryArg)
	}},
	{"tasks.dealRoom.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TasksDealRoomGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Tasks.DealRoom.Get(ctx, queryArg)
	}},
	{"tasks.dealRoom.link", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TasksDealRoomLinkBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.DealRoom.Link(ctx, bodyArg)
	}},
	{"tasks.dealRoom.options", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TasksDealRoomOptionsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Tasks.DealRoom.Options(ctx, queryArg)
	}},
	{"tasks.dealRoom.unlink", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TasksDealRoomUnlinkQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Tasks.DealRoom.Unlink(ctx, queryArg)
	}},
	{"tasks.delegation.cancel", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TasksDelegationCancelQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Tasks.Delegation.Cancel(ctx, queryArg)
	}},
	{"tasks.delegation.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TasksDelegationGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Tasks.Delegation.Get(ctx, queryArg)
	}},
	{"tasks.delegation.start", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg any
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.Delegation.Start(ctx, bodyArg)
	}},
	{"tasks.details.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TasksDetailsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Tasks.Details.Get(ctx, queryArg)
	}},
	{"tasks.details.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TasksDetailsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.Details.Update(ctx, bodyArg)
	}},
	{"tasks.feed.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TasksFeedCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.Feed.Create(ctx, bodyArg)
	}},
	{"tasks.feed.revoke", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Tasks.Feed.Revoke(ctx)
	}},
	{"tasks.focus.end", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Tasks.Focus.End(ctx)
	}},
	{"tasks.focus.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Tasks.Focus.Get(ctx)
	}},
	{"tasks.focus.start", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TasksFocusStartBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.Focus.Start(ctx, bodyArg)
	}},
	{"tasks.nudges.remind", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg any
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.Nudges.Remind(ctx, bodyArg)
	}},
	{"tasks.planDay.blocks", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TasksPlanDayBlocksQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Tasks.PlanDay.Blocks(ctx, queryArg)
	}},
	{"tasks.planDay.book", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TasksPlanDayBookBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.PlanDay.Book(ctx, bodyArg)
	}},
	{"tasks.planDay.freeBlocks", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TasksPlanDayFreeBlocksQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Tasks.PlanDay.FreeBlocks(ctx, queryArg)
	}},
	{"tasks.planDay.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TasksPlanDayGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Tasks.PlanDay.Get(ctx, queryArg)
	}},
	{"tasks.quickAdd.ai", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TasksQuickAddAiBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.QuickAdd.AI(ctx, bodyArg)
	}},
	{"tasks.quickAdd.records", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TasksQuickAddRecordsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Tasks.QuickAdd.Records(ctx, queryArg)
	}},
	{"tasks.quickAdd.space", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TasksQuickAddSpaceQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Tasks.QuickAdd.Space(ctx, queryArg)
	}},
	{"tasks.recurrence.skip", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TasksRecurrenceSkipBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.Recurrence.Skip(ctx, bodyArg)
	}},
	{"tasks.settings.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Tasks.Settings.Get(ctx)
	}},
	{"tasks.settings.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TaskSettingsPatch
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.Settings.Update(ctx, bodyArg)
	}},
	{"tasks.sla.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TasksSlaGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Tasks.SLA.Get(ctx, queryArg)
	}},
	{"tasks.sla.put", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TasksSlaPutBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.SLA.Put(ctx, bodyArg)
	}},
	{"tasks.views.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TasksViewsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.Views.Create(ctx, bodyArg)
	}},
	{"tasks.views.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Tasks.Views.Delete(ctx, a[0])
	}},
	{"tasks.views.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Tasks.Views.List(ctx)
	}},
	{"tasks.views.setDefault", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TasksViewsSetDefaultBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.Views.SetDefault(ctx, bodyArg)
	}},
	{"tasks.views.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TasksViewsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Tasks.Views.Update(ctx, a[0], bodyArg)
	}},
	{"taxRates.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.TaxRates.List(ctx)
	}},
	{"teams.addMember", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *AddTeamMemberInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Teams.AddMember(ctx, a[0], bodyArg)
	}},
	{"teams.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CreateTeamInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Teams.Create(ctx, bodyArg)
	}},
	{"teams.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Teams.Delete(ctx, a[0])
	}},
	{"teams.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Teams.Get(ctx, a[0])
	}},
	{"teams.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TeamsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Teams.List(ctx, queryArg)
	}},
	{"teams.listMembers", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Teams.ListMembers(ctx, a[0])
	}},
	{"teams.removeMember", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Teams.RemoveMember(ctx, a[0], a[1])
	}},
	{"teams.removeMemberRole", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TeamsRemoveMemberRoleQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Teams.RemoveMemberRole(ctx, a[0], a[1], queryArg)
	}},
	{"teams.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UpdateTeamInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Teams.Update(ctx, a[0], bodyArg)
	}},
	{"teams.updateMember", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TeamsUpdateMemberBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Teams.UpdateMember(ctx, a[0], a[1], bodyArg)
	}},
	{"telephony.dial", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TelephonyDialBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Telephony.Dial(ctx, bodyArg)
	}},
	{"telephony.overview", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TelephonyOverviewQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Telephony.Overview(ctx, queryArg)
	}},
	{"telephony.calls.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TelephonyCallsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Telephony.Calls.Get(ctx, a[0], queryArg)
	}},
	{"telephony.calls.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TelephonyCallsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Telephony.Calls.List(ctx, queryArg)
	}},
	{"telephony.connections.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PhoneConnectionInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Telephony.Connections.Create(ctx, bodyArg)
	}},
	{"telephony.connections.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Telephony.Connections.Delete(ctx, a[0])
	}},
	{"telephony.connections.syncLines", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Telephony.Connections.SyncLines(ctx, a[0])
	}},
	{"telephony.connections.test", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Telephony.Connections.Test(ctx, a[0])
	}},
	{"telephony.connections.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PhoneConnectionPatch
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Telephony.Connections.Update(ctx, a[0], bodyArg)
	}},
	{"telephony.lines.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PhoneLineInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Telephony.Lines.Create(ctx, bodyArg)
	}},
	{"telephony.lines.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Telephony.Lines.Delete(ctx, a[0])
	}},
	{"telephony.lines.mine", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TelephonyLinesMineQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Telephony.Lines.Mine(ctx, queryArg)
	}},
	{"telephony.lines.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *PhoneLinePatch
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Telephony.Lines.Update(ctx, a[0], bodyArg)
	}},
	{"templates.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *CreateTemplateInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Templates.Create(ctx, bodyArg)
	}},
	{"templates.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Templates.Delete(ctx, a[0])
	}},
	{"templates.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Templates.Get(ctx, a[0])
	}},
	{"templates.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TemplatesListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Templates.List(ctx, queryArg)
	}},
	{"templates.replace", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UpdateTemplateInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Templates.Replace(ctx, a[0], bodyArg)
	}},
	{"templates.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UpdateTemplateInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Templates.Update(ctx, a[0], bodyArg)
	}},
	{"templates.use", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TemplatesUseBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Templates.Use(ctx, a[0], bodyArg)
	}},
	{"trash.apply", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TrashApplyBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Trash.Apply(ctx, bodyArg)
	}},
	{"trash.empty", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TrashEmptyQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Trash.Empty(ctx, queryArg)
	}},
	{"trash.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TrashListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Trash.List(ctx, queryArg)
	}},
	{"trash.settings.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *TrashSettingsGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Trash.Settings.Get(ctx, queryArg)
	}},
	{"trash.settings.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *TrashSettingsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Trash.Settings.Update(ctx, bodyArg)
	}},
	{"user.addAiMemory", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UserAddAiMemoryBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.User.AddAIMemory(ctx, bodyArg)
	}},
	{"user.claimHandle", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UserClaimHandleBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.User.ClaimHandle(ctx, bodyArg)
	}},
	{"user.clearPresence", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.User.ClearPresence(ctx)
	}},
	{"user.countAiEvents", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.User.CountAIEvents(ctx)
	}},
	{"user.createSignature", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UserCreateSignatureBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.User.CreateSignature(ctx, bodyArg)
	}},
	{"user.deleteAiEvents", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.User.DeleteAIEvents(ctx)
	}},
	{"user.deleteAiMemory", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *UserDeleteAiMemoryQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.User.DeleteAIMemory(ctx, queryArg)
	}},
	{"user.deleteAvatar", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.User.DeleteAvatar(ctx)
	}},
	{"user.deleteSignature", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.User.DeleteSignature(ctx, a[0])
	}},
	{"user.exportActivityLogs", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *UserExportActivityLogsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.User.ExportActivityLogs(ctx, queryArg)
	}},
	{"user.forgetRecognizedDevice", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *UserForgetRecognizedDeviceQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.User.ForgetRecognizedDevice(ctx, queryArg)
	}},
	{"user.getAiMemory", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.User.GetAIMemory(ctx)
	}},
	{"user.getAiMemorySwitch", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.User.GetAIMemorySwitch(ctx)
	}},
	{"user.getHandoff", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.User.GetHandoff(ctx)
	}},
	{"user.getPreferences", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.User.GetPreferences(ctx)
	}},
	{"user.getPresence", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *UserGetPresenceQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.User.GetPresence(ctx, queryArg)
	}},
	{"user.getPresenceSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.User.GetPresenceSettings(ctx)
	}},
	{"user.getProfile", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.User.GetProfile(ctx)
	}},
	{"user.getSecurity", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.User.GetSecurity(ctx)
	}},
	{"user.listActivityLogs", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *UserListActivityLogsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.User.ListActivityLogs(ctx, queryArg)
	}},
	{"user.listRecognizedDevices", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.User.ListRecognizedDevices(ctx)
	}},
	{"user.listSignatures", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.User.ListSignatures(ctx)
	}},
	{"user.patchPreferences", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UserPreferences
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.User.PatchPreferences(ctx, bodyArg)
	}},
	{"user.releaseHandle", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.User.ReleaseHandle(ctx)
	}},
	{"user.setAiMemorySwitch", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UserSetAiMemorySwitchBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.User.SetAIMemorySwitch(ctx, bodyArg)
	}},
	{"user.setHandoff", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UserSetHandoffBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.User.SetHandoff(ctx, bodyArg)
	}},
	{"user.setPresence", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UserSetPresenceBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.User.SetPresence(ctx, bodyArg)
	}},
	{"user.updateAiMemory", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UserUpdateAiMemoryBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		var queryArg *UserUpdateAiMemoryQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.User.UpdateAIMemory(ctx, bodyArg, queryArg)
	}},
	{"user.updateAvatar", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UserUpdateAvatarBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.User.UpdateAvatar(ctx, bodyArg)
	}},
	{"user.updatePreferences", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UserPreferences
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.User.UpdatePreferences(ctx, bodyArg)
	}},
	{"user.updatePresenceSettings", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UserUpdatePresenceSettingsBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.User.UpdatePresenceSettings(ctx, bodyArg)
	}},
	{"user.updateProfile", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UpdateProfileInput
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.User.UpdateProfile(ctx, bodyArg)
	}},
	{"user.updateSignature", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *UserUpdateSignatureBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.User.UpdateSignature(ctx, a[0], bodyArg)
	}},
	{"users.listOrganizations", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Users.ListOrganizations(ctx, a[0])
	}},
	{"users.listRoles", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Users.ListRoles(ctx, a[0])
	}},
	{"users.listWorkspaces", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Users.ListWorkspaces(ctx, a[0])
	}},
	{"variables.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *VariablesGetQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Variables.Get(ctx, queryArg)
	}},
	{"variables.save", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *VariablesSaveBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Variables.Save(ctx, bodyArg)
	}},
	{"webhooks.getSubscription", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Webhooks.GetSubscription(ctx, a[0])
	}},
	{"webhooks.listSubscriptions", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *WebhooksListSubscriptionsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Webhooks.ListSubscriptions(ctx, queryArg)
	}},
	{"webhooks.subscribe", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *WebhooksSubscribeBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Webhooks.Subscribe(ctx, bodyArg)
	}},
	{"webhooks.unsubscribe", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Webhooks.Unsubscribe(ctx, a[0])
	}},
	{"workflows.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *WorkflowsCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Workflows.Create(ctx, bodyArg)
	}},
	{"workflows.createAlert", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *WorkflowsCreateAlertBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Workflows.CreateAlert(ctx, a[0], bodyArg)
	}},
	{"workflows.createVersion", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *WorkflowsCreateVersionBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Workflows.CreateVersion(ctx, a[0], bodyArg)
	}},
	{"workflows.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workflows.Delete(ctx, a[0])
	}},
	{"workflows.deleteAlert", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *WorkflowsDeleteAlertQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Workflows.DeleteAlert(ctx, a[0], queryArg)
	}},
	{"workflows.deleteDeadLetter", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *WorkflowsDeleteDeadLetterQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Workflows.DeleteDeadLetter(ctx, queryArg)
	}},
	{"workflows.deleteVersion", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workflows.DeleteVersion(ctx, a[0], a[1])
	}},
	{"workflows.deleteVersions", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workflows.DeleteVersions(ctx, a[0])
	}},
	{"workflows.duplicate", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workflows.Duplicate(ctx, a[0])
	}},
	{"workflows.execute", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *WorkflowsExecuteBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Workflows.Execute(ctx, a[0], bodyArg)
	}},
	{"workflows.export", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workflows.Export(ctx, a[0])
	}},
	{"workflows.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workflows.Get(ctx, a[0])
	}},
	{"workflows.getVersion", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workflows.GetVersion(ctx, a[0], a[1])
	}},
	{"workflows.import", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *WorkflowsImportBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Workflows.Import(ctx, bodyArg)
	}},
	{"workflows.installRecipe", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *WorkflowsInstallRecipeBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Workflows.InstallRecipe(ctx, bodyArg)
	}},
	{"workflows.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *WorkflowsListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Workflows.List(ctx, queryArg)
	}},
	{"workflows.listAlerts", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workflows.ListAlerts(ctx, a[0])
	}},
	{"workflows.listAllDeadLetters", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workflows.ListAllDeadLetters(ctx)
	}},
	{"workflows.listDeadLetterQueue", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workflows.ListDeadLetterQueue(ctx, a[0])
	}},
	{"workflows.listExecutions", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *WorkflowsListExecutionsQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Workflows.ListExecutions(ctx, a[0], queryArg)
	}},
	{"workflows.listRecipes", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workflows.ListRecipes(ctx)
	}},
	{"workflows.listVersions", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workflows.ListVersions(ctx, a[0])
	}},
	{"workflows.queue", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *WorkflowsQueueQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Workflows.Queue(ctx, a[0], queryArg)
	}},
	{"workflows.replayExecution", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *WorkflowsReplayExecutionBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Workflows.ReplayExecution(ctx, a[0], a[1], bodyArg)
	}},
	{"workflows.restoreVersion", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *WorkflowsRestoreVersionBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Workflows.RestoreVersion(ctx, a[0], a[1], bodyArg)
	}},
	{"workflows.retryDeadJob", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *WorkflowsRetryDeadJobBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Workflows.RetryDeadJob(ctx, a[0], bodyArg)
	}},
	{"workflows.setStatus", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *WorkflowsSetStatusBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Workflows.SetStatus(ctx, a[0], bodyArg)
	}},
	{"workflows.toggleStatus", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workflows.ToggleStatus(ctx, a[0])
	}},
	{"workflows.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *WorkflowsUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Workflows.Update(ctx, a[0], bodyArg)
	}},
	{"workflows.usage", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *WorkflowsUsageQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Workflows.Usage(ctx, queryArg)
	}},
	{"workspaces.addMember", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *WorkspacesAddMemberBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Workspaces.AddMember(ctx, a[0], bodyArg)
	}},
	{"workspaces.assignMemberRole", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *WorkspacesAssignMemberRoleBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Workspaces.AssignMemberRole(ctx, a[0], a[1], bodyArg)
	}},
	{"workspaces.assignTeam", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *WorkspacesAssignTeamBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Workspaces.AssignTeam(ctx, a[0], bodyArg)
	}},
	{"workspaces.cancelDeletion", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workspaces.CancelDeletion(ctx, a[0])
	}},
	{"workspaces.create", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *WorkspacesCreateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Workspaces.Create(ctx, bodyArg)
	}},
	{"workspaces.delete", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workspaces.Delete(ctx, a[0])
	}},
	{"workspaces.deleteEnvVar", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *WorkspacesDeleteEnvVarQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Workspaces.DeleteEnvVar(ctx, a[0], queryArg)
	}},
	{"workspaces.get", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workspaces.Get(ctx, a[0])
	}},
	{"workspaces.getScheduledDeletion", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workspaces.GetScheduledDeletion(ctx, a[0])
	}},
	{"workspaces.list", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *WorkspacesListQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Workspaces.List(ctx, queryArg)
	}},
	{"workspaces.listEnvVars", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workspaces.ListEnvVars(ctx, a[0])
	}},
	{"workspaces.listMembers", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workspaces.ListMembers(ctx, a[0])
	}},
	{"workspaces.listTeams", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workspaces.ListTeams(ctx, a[0])
	}},
	{"workspaces.removeMember", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		return c.Workspaces.RemoveMember(ctx, a[0], a[1])
	}},
	{"workspaces.removeMemberRole", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *WorkspacesRemoveMemberRoleQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Workspaces.RemoveMemberRole(ctx, a[0], a[1], queryArg)
	}},
	{"workspaces.removeTeam", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var queryArg *WorkspacesRemoveTeamQuery
		if err := decodeArg(query, &queryArg); err != nil {
			return nil, err
		}
		return c.Workspaces.RemoveTeam(ctx, a[0], a[1], queryArg)
	}},
	{"workspaces.setEnvVar", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *WorkspacesSetEnvVarBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Workspaces.SetEnvVar(ctx, a[0], bodyArg)
	}},
	{"workspaces.update", func(ctx context.Context, c *Client, a []string, body, query json.RawMessage) (any, error) {
		var bodyArg *WorkspacesUpdateBody
		if err := decodeArg(body, &bodyArg); err != nil {
			return nil, err
		}
		return c.Workspaces.Update(ctx, a[0], bodyArg)
	}},
}
