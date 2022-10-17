package core

import (
	"sync"
	"time"
)

//Autogenerated https://mholt.github.io/json-to-go/
type Object struct {
	sync.Mutex
	Inflections *Inflections `json:"inflections"`
	Connections *Connections `json:"connections"`
	User   bool          `json:"user"`
	Status string        `json:"status"`
	Tasks  []interface{} `json:"tasks"`
	Type   string        `json:"type"`
	ID     string        `json:"_id"`
	Name   string        `json:"name"`
	Fields []*Field `json:"fields"`
	Template   string `json:"template,omitempty"`
	Key        string `json:"key"`
	Identifier string `json:"identifier"`
	ProfileKey string `json:"profile_key,omitempty"`
}

type Inflections struct {
	Singular string `json:"singular"`
	Plural   string `json:"plural"`
} 

type Connections struct {
	Inbound  []interface{} `json:"inbound"`
	Outbound []interface{} `json:"outbound"`
}

type Field struct {
	Type        string        `json:"type"`
	Required    bool          `json:"required"`
	Unique      bool          `json:"unique"`
	User        bool          `json:"user"`
	Conditional bool          `json:"conditional"`
	Rules       []interface{} `json:"rules"`
	Validation  []interface{} `json:"validation"`
	ID          string        `json:"_id"`
	Key         string        `json:"key"`
	Name        string        `json:"name"`
	Format      *Format `json:"format,omitempty"`
	Format0 *FormatX `json:"format,omitempty"`
	Format1 *FormatX `json:"format,omitempty"`
	Default bool `json:"default,omitempty"`
}

type Format struct {
	Source string        `json:"source"`
	Thumbs []interface{} `json:"thumbs"`
}

type FormatX struct {
	Default  bool   `json:"default,omitempty"`
	Format   string `json:"format,omitempty"`
	Input    string `json:"input,omitempty"`
	Required bool   `json:"required,omitempty"`
}

type Scene struct {
	sync.Mutex
	Groups []*Group `json:"groups"`
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Key   string `json:"key"`
	Views []*View `json:"views"`
	Parent             string        `json:"parent,omitempty"`
	Object             interface{}   `json:"object,omitempty"`
	Type               string        `json:"type,omitempty"`
	LimitProfileAccess bool          `json:"limit_profile_access,omitempty"`
	AllowedProfiles    []interface{} `json:"allowed_profiles,omitempty"`
}

type Group struct {
	Columns []*Column `json:"columns"`
} 

type Column struct {
	Keys  []string `json:"keys"`
	Width int      `json:"width"`
}

type View struct {
	Columns []interface{} `json:"columns"`
	Links   []*Link `json:"links"`
	Inputs      []interface{} `json:"inputs"`
	ID          string        `json:"_id"`
	Groups      []interface{} `json:"groups"`
	Format      string        `json:"format"`
	Label       string        `json:"label"`
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Source      *Source `json:"source"`
	Key string `json:"key"`
}

type Link struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Scene string `json:"scene"`
}

type Source struct {
	Criteria *Criteria `json:"criteria"`
	Limit string        `json:"limit"`
	Sort  []interface{} `json:"sort"`
	Type  string        `json:"type"`
}

type Criteria struct {
	Match  string        `json:"match"`
	Rules  []interface{} `json:"rules"`
	Groups []interface{} `json:"groups"`
} 

type Users struct {
	Enabled      bool   `json:"enabled"`
	Scope        string `json:"scope"`
	Registration string `json:"registration"`
}

type Ecommerce  struct {
	Enabled bool `json:"enabled"`
}

type DBSchema struct {
	ID    	  string `json:"_id"`
	Users 	  *Users `json:"users"`
	Ecommerce *Ecommerce `json:"ecommerce"`
	Counts Counts `json:"counts"`
	FieldCount      int    `json:"field_count"`
	ThumbCount      int    `json:"thumb_count"`
	ObjectCount     int    `json:"object_count"`
	TaskCount       int    `json:"task_count"`
	ViewCount       int    `json:"view_count"`
	SceneCount      int    `json:"scene_count"`
	CredentialCount int    `json:"credential_count"`
	Status          string `json:"status"`
	Settings        *Settings `json:"settings"`
	Tasks             []interface{} `json:"tasks"`
	PaymentProcessors []interface{} `json:"payment_processors"`
	Design           *Design `json:"design"`
	Layout *Layout `json:"layout"`
	Copying       bool          `json:"copying"`
	FeatureFlags  []bool        `json:"feature_flags"`
	Name          string        `json:"name"`
	Slug          string        `json:"slug"`
	Distributions []interface{} `json:"distributions"`
	Versions      []*Version `json:"versions"`
	FirstCreated time.Time `json:"first_created"`
	AccountID    string    `json:"account_id"`
	UserID       string    `json:"user_id"`
}

type Design  struct {
	General *General `json:"general"`
	Regions *Regions `json:"regions"`
	Components *Components `json:"components"`
}

type General struct {
	Theme string `json:"theme"`
	Font  string `json:"font"`
	Links *Links `json:"links"`
	Buttons *Buttons `json:"buttons"`
	Tables *Tables `json:"tables"`
} 

type Links struct {
	Color string `json:"color"`
}

type Buttons struct {
	Color   string `json:"color"`
	BgColor string `json:"bg_color"`
} 
type Tables struct {
	Style    string `json:"style"`
	Dividers bool   `json:"dividers"`
	Striped  bool   `json:"striped"`
	Hover    bool   `json:"hover"`
	Spacing  string `json:"spacing"`
} 

type Regions struct {
	Body *Body  `json:"body"`
	Header *Header `json:"header"`
}

type Body struct {
	FullWidth bool `json:"full_width"`
}
type Header struct {
	BgColor string `json:"bg_color"`
	Menu   *Menu `json:"menu"`
	Title *Title `json:"title"`
}

type Menu  struct {
	Show               bool   `json:"show"`
	UserAuthBased      bool   `json:"user_auth_based"`
	Format             string `json:"format"`
	Color              string `json:"color"`
	OutlineOrFillColor string `json:"outline_or_fill_color"`
} 

type Title  struct {
	Color    string `json:"color"`
	ShowLogo bool   `json:"show_logo"`
}

type Components struct {
	Notifications *Notifications `json:"notifications"`
}

type Notifications struct {
	Color   string `json:"color"`
	BgColor string `json:"bg_color"`
} 

type Layout struct {
	EntrySceneMenu bool   `json:"entry_scene_menu"`
	AppMenuAuth    bool   `json:"app_menu_auth"`
	Theme          string `json:"theme"`
}

type Version struct {
	sync.Mutex //added for concurrency
	ID      string `json:"_id"`
	Status  string `json:"status"`
	Objects []*Object `json:"objects"`
	Scenes []*Scene `json:"scenes"`
}

type Counts struct {
	TotalEntries int `json:"total_entries"`
	AssetCount   int `json:"asset_count"`
	AssetSize    int `json:"asset_size"`
} 

type Settings struct {
	Geo                         bool   `json:"geo"`
	Timezone                    string `json:"timezone"`
	TimezoneOffset              string `json:"timezone_offset"`
	TimezoneDst                 string `json:"timezone_dst"`
	Cluster                     string `json:"cluster"`
	PoweredByLink               bool   `json:"powered_by_link"`
	NewCount                    bool   `json:"new_count"`
	RegistrationMigrationStatus string `json:"registration_migration_status"`
	HTTPSRedirect               bool   `json:"https_redirect"`
	ShouldPurgeRecordHistory    bool   `json:"should_purge_record_history"`
	SupportAccess               bool   `json:"support_access"`
	UseMultipleAPISubdomains    bool   `json:"use_multiple_api_subdomains"`
	ScriptProtectionEnabled     bool   `json:"scriptProtectionEnabled"`
	EmbedLoginMethod            string `json:"embed_login_method"`
	IgnoreBilling               bool   `json:"ignoreBilling"`
	V3Beta                      bool   `json:"v3_beta"`
	V3OpenBeta                  bool   `json:"v3_open_beta"`
	Mongo                       string `json:"mongo"`
	Solr                        string `json:"solr"`
}