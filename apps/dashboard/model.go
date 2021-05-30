package dashboard

type ChartStockTrend struct {
	ID             int
	ChartName      string
	ChartFrequency ChartFrequency
}

// ChartFrequency 图表更新频率
type ChartFrequency struct {
	ID              int
	UpdateFrequency string // month: 每个月最后一天
}
