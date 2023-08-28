package models

// Settings represents the app settings stored in the DB.
type Settings struct {
	BounceActions map[string]struct {
		Action string `json:"action"`
		Count  int    `json:"count"`
	} `json:"bounce.actions"`
	AppSiteName   string `json:"app.site_name"`
	AppRootURL    string `json:"app.root_url"`
	AppLogoURL    string `json:"app.logo_url"`
	AppFaviconURL string `json:"app.favicon_url"`
	AppFromEmail  string `json:"app.from_email"`
	AppLang       string `json:"app.lang"`

	AppMessageSlidingWindowDuration string `json:"app.message_sliding_window_duration"`
	SecurityCaptchaKey              string `json:"security.captcha_key"`
	SecurityCaptchaSecret           string `json:"security.captcha_secret"`

	UploadProvider             string `json:"upload.provider"`
	UploadFilesystemUploadPath string `json:"upload.filesystem.upload_path"`
	UploadFilesystemUploadURI  string `json:"upload.filesystem.upload_uri"`
	UploadS3URL                string `json:"upload.s3.url"`
	UploadS3PublicURL          string `json:"upload.s3.public_url"`
	UploadS3AwsAccessKeyID     string `json:"upload.s3.aws_access_key_id"`
	UploadS3AwsDefaultRegion   string `json:"upload.s3.aws_default_region"`
	UploadS3AwsSecretAccessKey string `json:"upload.s3.aws_secret_access_key,omitempty"`
	UploadS3Bucket             string `json:"upload.s3.bucket"`
	UploadS3BucketDomain       string `json:"upload.s3.bucket_domain"`
	UploadS3BucketPath         string `json:"upload.s3.bucket_path"`
	UploadS3BucketType         string `json:"upload.s3.bucket_type"`
	UploadS3Expiry             string `json:"upload.s3.expiry"`

	SendgridKey string `json:"bounce.sendgrid_key"`

	AdminCustomCSS    string   `json:"appearance.admin.custom_css"`
	AdminCustomJS     string   `json:"appearance.admin.custom_js"`
	PublicCustomCSS   string   `json:"appearance.public.custom_css"`
	PublicCustomJS    string   `json:"appearance.public.custom_js"`
	AppNotifyEmails   []string `json:"app.notify_emails"`
	PrivacyExportable []string `json:"privacy.exportable"`
	DomainBlocklist   []string `json:"privacy.domain_blocklist"`

	UploadExtensions []string `json:"upload.extensions"`

	SMTP []struct {
		UUID          string              `json:"uuid"`
		Host          string              `json:"host"`
		HelloHostname string              `json:"hello_hostname"`
		AuthProtocol  string              `json:"auth_protocol"`
		Username      string              `json:"username"`
		Password      string              `json:"password,omitempty"`
		IdleTimeout   string              `json:"idle_timeout"`
		WaitTimeout   string              `json:"wait_timeout"`
		TLSType       string              `json:"tls_type"`
		EmailHeaders  []map[string]string `json:"email_headers"`
		Port          int                 `json:"port"`
		MaxConns      int                 `json:"max_conns"`
		MaxMsgRetries int                 `json:"max_msg_retries"`
		Enabled       bool                `json:"enabled"`
		TLSSkipVerify bool                `json:"tls_skip_verify"`
	} `json:"smtp"`

	Messengers []struct {
		UUID          string `json:"uuid"`
		Name          string `json:"name"`
		RootURL       string `json:"root_url"`
		Username      string `json:"username"`
		Password      string `json:"password,omitempty"`
		Timeout       string `json:"timeout"`
		MaxConns      int    `json:"max_conns"`
		MaxMsgRetries int    `json:"max_msg_retries"`
		Enabled       bool   `json:"enabled"`
	} `json:"messengers"`

	BounceBoxes []struct {
		UUID          string `json:"uuid"`
		Type          string `json:"type"`
		Host          string `json:"host"`
		AuthProtocol  string `json:"auth_protocol"`
		ReturnPath    string `json:"return_path"`
		Username      string `json:"username"`
		Password      string `json:"password,omitempty"`
		ScanInterval  string `json:"scan_interval"`
		Port          int    `json:"port"`
		Enabled       bool   `json:"enabled"`
		TLSEnabled    bool   `json:"tls_enabled"`
		TLSSkipVerify bool   `json:"tls_skip_verify"`
	} `json:"bounce.mailboxes"`

	AppBatchSize     int `json:"app.batch_size"`
	AppConcurrency   int `json:"app.concurrency"`
	AppMaxSendErrors int `json:"app.max_send_errors"`
	AppMessageRate   int `json:"app.message_rate"`

	AppMessageSlidingWindowRate int `json:"app.message_sliding_window_rate"`

	EnablePublicSubPage           bool `json:"app.enable_public_subscription_page"`
	EnablePublicArchive           bool `json:"app.enable_public_archive"`
	EnablePublicArchiveRSSContent bool `json:"app.enable_public_archive_rss_content"`
	SendOptinConfirmation         bool `json:"app.send_optin_confirmation"`
	CheckUpdates                  bool `json:"app.check_updates"`

	AppMessageSlidingWindow bool `json:"app.message_sliding_window"`

	PrivacyIndividualTracking bool `json:"privacy.individual_tracking"`
	PrivacyUnsubHeader        bool `json:"privacy.unsubscribe_header"`
	PrivacyAllowBlocklist     bool `json:"privacy.allow_blocklist"`
	PrivacyAllowPreferences   bool `json:"privacy.allow_preferences"`
	PrivacyAllowExport        bool `json:"privacy.allow_export"`
	PrivacyAllowWipe          bool `json:"privacy.allow_wipe"`
	PrivacyRecordOptinIP      bool `json:"privacy.record_optin_ip"`

	SecurityEnableCaptcha bool `json:"security.enable_captcha"`

	BounceEnabled        bool `json:"bounce.enabled"`
	BounceEnableWebhooks bool `json:"bounce.webhooks_enabled"`
	SESEnabled           bool `json:"bounce.ses_enabled"`
	SendgridEnabled      bool `json:"bounce.sendgrid_enabled"`
}
