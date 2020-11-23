package jpush

import "encoding/json"

type Platform string
type Target string

const (
	PlatformAndroid  Platform = "android"
	PlatformIOS      Platform = "ios"
	PlatformWinPhone Platform = "winphone"
	PlatformQuickApp Platform = "quickapp"
)

type PushAudience struct {
	Tag            []string `json:"tag,omitempty"`
	TagAnd         []string `json:"tag_and,omitempty"`
	TagNot         []string `json:"tag_not,omitempty"`
	Alias          []string `json:"alias,omitempty"`
	RegistrationId []string `json:"registration_id,omitempty"`
	Segment        []string `json:"segment,omitempty"`
	ABTest         []string `json:"abtest,omitempty"`
}

type PushNotification struct {
	Alert    string                `json:"alert,omitempty"`
	Android  *NotificationAndroid  `json:"android,omitempty"`
	IOS      *NotificationIOS      `json:"ios,omitempty"`
	VIOS     *NotificationVIOS     `json:"vios,omitempty"`
	WinPhone *NotificationWinPhone `json:"winphone,omitempty"`
	QuickApp *NotificationQuickApp `json:"quickapp,omitempty"`
}

type NotificationAndroid struct {
	Alert             string                 `json:"alert"`
	Title             string                 `json:"title,omitempty"`
	BuilderId         int                    `json:"builder_id,int,omitempty"`
	ChannelId         int                    `json:"channel_id,int,omitempty"`
	Priority          int                    `json:"priority,omitempty"`
	Category          string                 `json:"category,omitempty"`
	Style             int                    `json:"style,int,omitempty"`
	AlertType         int                    `json:"alert_type,int,omitempty"`
	BigText           string                 `json:"big_text,omitempty"`
	Inbox             map[string]interface{} `json:"inbox,omitempty"`
	BigPicPath        string                 `json:"big_pic_path,omitempty"`
	Extras            map[string]interface{} `json:"extras,omitempty"`
	LargeIcon         string                 `json:"large_icon,omitempty"`
	Intent            map[string]interface{} `json:"intent,omitempty"`
	UriActivity       string                 `json:"uri_activity,omitempty"` // 兼容华为
	UriAction         string                 `json:"uri_action,omitempty"`   //兼容小米
	BadgeAddNum       int                    `json:"badge_add_num,int,omitempty"`
	BadgeClass        string                 `json:"badge_class,omitempty"`
	Sound             string                 `json:"sound,omitempty"`
	ShowBeginTime     string                 `json:"show_begin_time,omitempty"`
	ShowEndTime       string                 `json:"show_end_time,omitempty"`
	DisplayForeground string                 `json:"display_foreground,omitempty"`
}

type NotificationIOS struct {
	Alert            interface{}            `json:"alert"`
	Sound            string                 `json:"sound,omitempty"`
	Badge            int                    `json:"badge,int,omitempty"`
	ContentAvailable bool                   `json:"content-available,omitempty"`
	MutableContent   bool                   `json:"mutable-content,omitempty"`
	Category         string                 `json:"category,omitempty"`
	Extras           map[string]interface{} `json:"extras,omitempty"`
	ThreadId         string                 `json:"thread-id,omitempty"`
}
type NotificationVIOS struct {
	Key string `json:"key,omitempty"`
}

type NotificationQuickApp struct {
	Title  string                 `json:"title,omitempty"`
	Alert  string                 `json:"alert"`
	Page   string                 `json:"page,omitempty"`
	Extras map[string]interface{} `json:"extras,omitempty"`
}
type NotificationWinPhone struct {
	Alert    string                 `json:"alert"`
	Title    string                 `json:"title,omitempty"`
	OpenPage string                 `json:"_open_page,omitempty"`
	Extras   map[string]interface{} `json:"extras,omitempty"`
}

type PushMessage struct {
	MsgContent  string                 `json:"msg_content"`
	Title       string                 `json:"title,omitempty"`
	ContentType string                 `json:"content_type,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}

type SmsMessage struct {
	Content      string `json:"content"`
	DelayTime    int    `json:"delay_time,int,omitempty"`
	SignId       int    `json:"signid,int,omitempty"`
	TempId       int64  `json:"temp_id,long,omitempty"`
	TempPara     string `json:"temp_para,long,omitempty"`
	ActiveFilter bool   `json:"active_filter,bool,omitempty"`
}

type PushOptions struct {
	SendNo            int                    `json:"sendno,int,omitempty"`
	TimeToLive        int                    `json:"time_to_live,int,omitempty"`
	OverrideMsgId     int64                  `json:"override_msg_id,int64,omitempty"`
	ApnsProduction    bool                   `json:"apns_production"`
	ApnsCollapseId    string                 `json:"apns_collapse_id,omitempty"`
	BigPushDuration   int                    `json:"big_push_duration,int,omitempty"`
	ThirdPartyChannel map[string]interface{} `json:"third_party_channel,omitempty"`
}

type PushList struct {
	Cid struct {
		Platform     Platform          `json:"platform"`
		Target       Target            `json:"target"`
		Notification *PushNotification `json:"notification,omitempty"`
		Message      *PushMessage      `json:"message,omitempty"`
		SmsMessage   *SmsMessage       `json:"sms_message,omitempty"`
		Options      *PushOptions      `json:"options,omitempty"`
	}
}
type PushRequest struct {
	Cid          string            `json:"cid,omitempty"`
	Platform     Platform          `json:"platform"`
	Audience     *PushAudience     `json:"audience,omitempty"`
	Notification *PushNotification `json:"notification,omitempty"`
	Message      *PushMessage      `json:"message,omitempty"`
	SmsMessage   *SmsMessage       `json:"sms_message,omitempty"`
	Options      *PushOptions      `json:"options,omitempty"`
}

type PushSingleRequest struct {
	PushList PushList
}

type Response struct {
	data []byte
}

func (res *Response) Array() ([]interface{}, error) {
	list := make([]interface{}, 0)
	err := json.Unmarshal(res.data, &list)
	return list, err
}

func (res *Response) Map() (map[string]interface{}, error) {
	result := make(map[string]interface{})
	err := json.Unmarshal(res.data, &result)
	return result, err
}

func (res *Response) Bytes() []byte {
	return res.data
}
