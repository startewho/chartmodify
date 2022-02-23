package chartmodify

/**
Building
**/
type Builing struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Systems []Systems `json:"systems"`
	Regions int       `json:"regions"`
}
type Systems struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Regions struct {
	ID    string `json:"id"`
	Left  int    `json:"left"`
	Right int    `json:"right"`
}

/**
ChartSetting
**/

type ChartSetting struct {
	Type   string   `json:"type"`
	Series []Series `json:"series"`
	Format Format   `json:"format"`
	Legend Legend   `json:"legend"`
	Title  Title    `json:"title"`
}
type Series struct {
	Name       string `json:"name"`
	Categories string `json:"categories"`
	Values     string `json:"values"`
}
type Format struct {
	XScale          float64 `json:"x_scale"`
	YScale          float64 `json:"y_scale"`
	XOffset         int     `json:"x_offset"`
	YOffset         int     `json:"y_offset"`
	PrintObj        bool    `json:"print_obj"`
	LockAspectRatio bool    `json:"lock_aspect_ratio"`
	Locked          bool    `json:"locked"`
}
type Legend struct {
	Position      string `json:"position"`
	ShowLegendKey bool   `json:"show_legend_key"`
}
type Title struct {
	Name string `json:"name"`
}
