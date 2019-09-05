package domain


type SlackWebhookRequest struct {
	Token string `json:"token"`
	TeamId string `json:"team_id"`
	TeamDomain string `json:"team_domain"`
	EnterpriseId string `json:"enterprise_id"`
	EnterpriseName string `json:"enterprise_name"`
	ChannelId string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	UserId string `json:"user_id"`
	Username string `json:"user_name"`
	Command string `json:"command"`
	Text string `json:"text"`
	ResponseUrl string `json:"response_url"`
	TriggerId string `json:"trigger_id"'`
}
