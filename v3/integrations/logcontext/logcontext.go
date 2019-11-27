package logcontext

import newrelic "github.com/newrelic/go-agent/v3/newrelic"

// Keys used for logging context JSON.
const (
	KeyFile       = "file.name"
	KeyLevel      = "log.level"
	KeyLine       = "line.number"
	KeyMessage    = "message"
	KeyMethod     = "method.name"
	KeyTimestamp  = "timestamp"
	KeyTraceID    = "trace.id"
	KeySpanID     = "span.id"
	KeyEntityName = "entity.name"
	KeyEntityType = "entity.type"
	KeyEntityGUID = "entity.guid"
	KeyHostname   = "hostname"
)

func metadataMapField(m map[string]interface{}, key, val string) {
	if val != "" {
		m[key] = val
	}
}

// AddLinkingMetadata adds the LinkingMetadata into a map.  Only non-empty
// string fields are included in the map.  The specific key names facilitate
// agent logs in context.  These keys are: "trace.id", "span.id",
// "entity.name", "entity.type", "entity.guid", and "hostname".
func AddLinkingMetadata(m map[string]interface{}, md newrelic.LinkingMetadata) {
	metadataMapField(m, KeyTraceID, md.TraceID)
	metadataMapField(m, KeySpanID, md.SpanID)
	metadataMapField(m, KeyEntityName, md.EntityName)
	metadataMapField(m, KeyEntityType, md.EntityType)
	metadataMapField(m, KeyEntityGUID, md.EntityGUID)
	metadataMapField(m, KeyHostname, md.Hostname)
}
