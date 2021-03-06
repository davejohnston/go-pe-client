package puppetdb

const nodes = "/pdb/query/v4/nodes"

// Nodes will return all nodes matching the given query. Deactivated and expired nodes aren’t included in the response.
func (c *Client) Nodes(query string, pagination *Pagination, orderBy *OrderBy) ([]Node, error) {
	payload := []Node{}
	err := getRequest(c, nodes, query, pagination, orderBy, &payload)
	return payload, err
}

// Node is a PuppetDB node
type Node struct {
	Deactivated                  interface{} `json:"deactivated"`
	LatestReportHash             string      `json:"latest_report_hash"`
	FactsEnvironment             string      `json:"facts_environment"`
	CachedCatalogStatus          string      `json:"cached_catalog_status"`
	ReportEnvironment            string      `json:"report_environment"`
	LatestReportCorrectiveChange bool        `json:"latest_report_corrective_change"`
	CatalogEnvironment           string      `json:"catalog_environment"`
	FactsTimestamp               string      `json:"facts_timestamp"`
	LatestReportNoop             bool        `json:"latest_report_noop"`
	Expired                      interface{} `json:"expired"`
	LatestReportNoopPending      bool        `json:"latest_report_noop_pending"`
	ReportTimestamp              string      `json:"report_timestamp"`
	Certname                     string      `json:"certname"`
	CatalogTimestamp             string      `json:"catalog_timestamp"`
	LatestReportJobID            string      `json:"latest_report_job_id"`
	LatestReportStatus           string      `json:"latest_report_status"`
	Count                        int         `json:"count"`
}
