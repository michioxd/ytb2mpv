package ytdlp

type YtdlpOutput struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Formats []Format `json:"formats"`
}

type Format struct {
	FormatID         string        `json:"format_id"`
	FormatNote       *string       `json:"format_note,omitempty"`
	EXT              EXT           `json:"ext"`
	Protocol         Protocol      `json:"protocol"`
	Acodec           *Acodec       `json:"acodec,omitempty"`
	Vcodec           string        `json:"vcodec"`
	URL              string        `json:"url"`
	Width            *int64        `json:"width,omitempty"`
	Height           *int64        `json:"height,omitempty"`
	FPS              *float64      `json:"fps,omitempty"`
	Rows             *int64        `json:"rows,omitempty"`
	Columns          *int64        `json:"columns,omitempty"`
	Fragments        []Fragment    `json:"fragments,omitempty"`
	AudioEXT         Acodec        `json:"audio_ext"`
	VideoEXT         Acodec        `json:"video_ext"`
	Vbr              float64       `json:"vbr"`
	ABR              *int64        `json:"abr"`
	Tbr              *float64      `json:"tbr"`
	Resolution       string        `json:"resolution"`
	AspectRatio      *float64      `json:"aspect_ratio"`
	FilesizeApprox   any           `json:"filesize_approx"`
	HTTPHeaders      any           `json:"http_headers"`
	Format           string        `json:"format"`
	FormatIndex      any           `json:"format_index"`
	ManifestURL      *string       `json:"manifest_url,omitempty"`
	Language         *string       `json:"language,omitempty"`
	Preference       any           `json:"preference"`
	Quality          *int64        `json:"quality,omitempty"`
	HasDRM           *bool         `json:"has_drm,omitempty"`
	SourcePreference *int64        `json:"source_preference,omitempty"`
	DynamicRange     *DynamicRange `json:"dynamic_range,omitempty"`
}

type Fragment struct {
	URL      string  `json:"url"`
	Duration float64 `json:"duration"`
}

type Acodec string

const (
	AcodecMp4 Acodec = "mp4"
	None      Acodec = "none"
)

type DynamicRange string

const (
	SDR DynamicRange = "SDR"
)

type EXT string

const (
	EXTMhtml EXT = "mhtml"
	EXTMp4   EXT = "mp4"
)

type Accept string

type AcceptLanguage string

const (
	EnUsEnQ05 AcceptLanguage = "en-us,en;q=0.5"
)

type SECFetchMode string

const (
	Navigate SECFetchMode = "navigate"
)

type Protocol string

const (
	M3U8Native    Protocol = "m3u8_native"
	ProtocolMhtml Protocol = "mhtml"
)
