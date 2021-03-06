package divan_types

type AutoCompactionSettings struct {
	ParallelDBAndViewCompaction    bool `json:"parallelDBAndViewCompaction"`
	AllowedTimePeriod              *AutoCompactionTimePeriodD
	DatabaseFragmentationThreshold *ThresholdD
	ViewFragmentationThreshold     *ThresholdD

	UndDatabaseFragmentationThreshold Threshold                `json:"databaseFragmentationThreshold"`
	UndViewFragmentationThreshold     Threshold                `json:"viewFragmentationThreshold"`
	UndAllowedTimePeriod              AutoCompactionTimePeriod `json:"allowedTimePeriod"`
}

func (acs *AutoCompactionSettings) Parse() {
	acs.DatabaseFragmentationThreshold = acs.UndDatabaseFragmentationThreshold.Determine()
	acs.ViewFragmentationThreshold = acs.UndViewFragmentationThreshold.Determine()
	acs.AllowedTimePeriod = acs.UndAllowedTimePeriod.Determine()
}

type AutoCompactionTimePeriodD struct {
	FromHour     uint64 `json:"fromHour"`
	FromMinute   uint64 `json:"fromMinute"`
	ToHour       uint64 `json:"toHour"`
	ToMinute     uint64 `json:"toMinute"`
	AbortOutside bool   `json:"abortOutside"`
}

type AutoCompactionTimePeriod struct {
	FromHour     interface{} `json:"fromHour"`
	FromMinute   interface{} `json:"fromMinute"`
	ToHour       interface{} `json:"toHour"`
	ToMinute     interface{} `json:"toMinute"`
	AbortOutside bool        `json:"abortOutside"`
}

func (a *AutoCompactionTimePeriod) Determine() *AutoCompactionTimePeriodD {
	output := &AutoCompactionTimePeriodD{
		AbortOutside: a.AbortOutside,
	}

	if fval, ok := a.FromHour.(float64); ok {
		output.FromHour = uint64(fval)
	}

	if fval, ok := a.FromMinute.(float64); ok {
		output.FromMinute = uint64(fval)
	}

	if fval, ok := a.ToHour.(float64); ok {
		output.ToHour = uint64(fval)
	}

	if fval, ok := a.ToMinute.(float64); ok {
		output.ToMinute = uint64(fval)
	}

	return output
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
