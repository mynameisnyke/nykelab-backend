package identify

type MagickInfo struct {
	Image Image `json:"image"`
}

type Image struct {
	Name              string            `json:"name"`
	Format            string            `json:"format"`
	FormatDescription string            `json:"formatDescription"`
	MIMEType          string            `json:"mimeType"`
	Class             string            `json:"class"`
	Geometry          Geometry          `json:"geometry"`
	Resolution        PrintSize         `json:"resolution"`
	PrintSize         PrintSize         `json:"printSize"`
	Units             string            `json:"units"`
	Type              string            `json:"type"`
	Endianess         string            `json:"endianess"`
	Colorspace        string            `json:"colorspace"`
	Depth             int64             `json:"depth"`
	BaseDepth         int64             `json:"baseDepth"`
	ChannelDepth      ChannelDepth      `json:"channelDepth"`
	Pixels            int64             `json:"pixels"`
	ImageStatistics   ImageStatistics   `json:"imageStatistics"`
	ChannelStatistics ChannelStatistics `json:"channelStatistics"`
	RenderingIntent   string            `json:"renderingIntent"`
	Gamma             float64           `json:"gamma"`
	Chromaticity      Chromaticity      `json:"chromaticity"`
	BackgroundColor   string            `json:"backgroundColor"`
	BorderColor       string            `json:"borderColor"`
	MatteColor        string            `json:"matteColor"`
	TransparentColor  string            `json:"transparentColor"`
	Interlace         string            `json:"interlace"`
	Intensity         string            `json:"intensity"`
	Compose           string            `json:"compose"`
	PageGeometry      Geometry          `json:"pageGeometry"`
	Dispose           string            `json:"dispose"`
	Iterations        int64             `json:"iterations"`
	Compression       string            `json:"compression"`
	Quality           int64             `json:"quality"`
	Orientation       string            `json:"orientation"`
	Properties        map[string]string `json:"properties"`
	Profiles          Profiles          `json:"profiles"`
	Artifacts         Artifacts         `json:"artifacts"`
	Tainted           bool              `json:"tainted"`
	Filesize          string            `json:"filesize"`
	NumberPixels      string            `json:"numberPixels"`
	PixelsPerSecond   string            `json:"pixelsPerSecond"`
	UserTime          string            `json:"userTime"`
	ElapsedTime       string            `json:"elapsedTime"`
	Version           string            `json:"version"`
}

type Artifacts struct {
	Filename string `json:"filename"`
}

type ChannelDepth struct {
	Red   int64 `json:"red"`
	Green int64 `json:"green"`
	Blue  int64 `json:"blue"`
}

type ChannelStatistics struct {
	Red   Blue `json:"red"`
	Green Blue `json:"green"`
	Blue  Blue `json:"blue"`
}

type Blue struct {
	Min               string `json:"min"`
	Max               string `json:"max"`
	Mean              string `json:"mean"`
	StandardDeviation string `json:"standardDeviation"`
	Kurtosis          string `json:"kurtosis"`
	Skewness          string `json:"skewness"`
}

type Chromaticity struct {
	RedPrimary   Primary `json:"redPrimary"`
	GreenPrimary Primary `json:"greenPrimary"`
	BluePrimary  Primary `json:"bluePrimary"`
	WhitePrimary Primary `json:"whitePrimary"`
}

type Primary struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Geometry struct {
	Width  int64 `json:"width"`
	Height int64 `json:"height"`
	X      int64 `json:"x"`
	Y      int64 `json:"y"`
}

type ImageStatistics struct {
	All Blue `json:"all"`
}

type PrintSize struct {
	X string `json:"x"`
	Y string `json:"y"`
}

type Profiles struct {
	The8Bim The8_Bim `json:"8bim"`
	Exif    The8_Bim `json:"exif"`
	Icc     The8_Bim `json:"icc"`
	Iptc    Iptc     `json:"iptc"`
	Xmp     The8_Bim `json:"xmp"`
}

type The8_Bim struct {
	Length string `json:"length"`
}

type Iptc struct {
	City190             []string      `json:"City[1,90]"`
	Unknown20           []interface{} `json:"unknown[2,0]"`
	CreatedDate255      []string      `json:"Created Date[2,55]"`
	CreatedTime260      []string      `json:"Created Time[2,60]"`
	Unknown262          []string      `json:"unknown[2,62]"`
	Unknown263          []string      `json:"unknown[2,63]"`
	Byline280           []string      `json:"Byline[2,80]"`
	CopyrightString2116 []string      `json:"Copyright String[2,116]"`
	Length              string        `json:"length"`
}
