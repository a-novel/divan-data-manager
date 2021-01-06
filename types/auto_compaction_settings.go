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
	FromHour     uint `json:"fromHour"`
	FromMinute   uint `json:"fromMinute"`
	ToHour       uint `json:"toHour"`
	ToMinute     uint `json:"toMinute"`
	AbortOutside bool `json:"abortOutside"`
}

type Threshold struct {
	Percentage interface{} `json:"percentage"`
	Size       interface{} `json:"size"`
}

type ThresholdD struct {
	Percentage uint `json:"percentage"`
	Size       uint `json:"size"`
}

func (t *Threshold) Determine() *ThresholdD {
	var output ThresholdD

	if fval, ok := t.Percentage.(float64); ok {
		output.Percentage = uint(fval)
	}

	if fval, ok := t.Size.(float64); ok {
		output.Size = uint(fval)
	}

	return &output
}
