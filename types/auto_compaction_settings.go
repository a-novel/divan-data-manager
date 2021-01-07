package divan_types

type AutoCompactionSettings struct {
	ParallelDBAndViewCompaction    bool                      `json:"parallelDBAndViewCompaction"`
	AllowedTimePeriod              *AutoCompactionTimePeriod `json:"allowedTimePeriod"`
	DatabaseFragmentationThreshold *ThresholdD
	ViewFragmentationThreshold     *ThresholdD

	UndDatabaseFragmentationThreshold Threshold `json:"databaseFragmentationThreshold"`
	UndViewFragmentationThreshold     Threshold `json:"viewFragmentationThreshold"`
}

func (acs *AutoCompactionSettings) Parse() {
	acs.DatabaseFragmentationThreshold = acs.UndDatabaseFragmentationThreshold.Determine()
	acs.ViewFragmentationThreshold = acs.UndViewFragmentationThreshold.Determine()
}

type AutoCompactionTimePeriod struct {
	FromHour     uint64 `json:"fromHour"`
	FromMinute   uint64 `json:"fromMinute"`
	ToHour       uint64 `json:"toHour"`
	ToMinute     uint64 `json:"toMinute"`
	AbortOutside bool   `json:"abortOutside"`
}

type Threshold struct {
	Percentage interface{} `json:"percentage"`
	Size       interface{} `json:"size"`
}

type ThresholdD struct {
	Percentage uint64 `json:"percentage"`
	Size       uint64 `json:"size"`
}

func (t *Threshold) Determine() *ThresholdD {
	var output ThresholdD

	if fval, ok := t.Percentage.(float64); ok {
		output.Percentage = uint64(fval)
	}

	if fval, ok := t.Size.(float64); ok {
		output.Size = uint64(fval)
	}

	return &output
}
