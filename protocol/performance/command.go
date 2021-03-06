// Code generated by cdpgen. DO NOT EDIT.

package performance

// EnableArgs represents the arguments for Enable in the Performance domain.
type EnableArgs struct {
	// TimeDomain Time domain to use for collecting and reporting duration
	// metrics.
	//
	// Values: "timeTicks", "threadTicks".
	TimeDomain *string `json:"timeDomain,omitempty"`
}

// NewEnableArgs initializes EnableArgs with the required arguments.
func NewEnableArgs() *EnableArgs {
	args := new(EnableArgs)

	return args
}

// SetTimeDomain sets the TimeDomain optional argument. Time domain to
// use for collecting and reporting duration metrics.
//
// Values: "timeTicks", "threadTicks".
func (a *EnableArgs) SetTimeDomain(timeDomain string) *EnableArgs {
	a.TimeDomain = &timeDomain
	return a
}

// SetTimeDomainArgs represents the arguments for SetTimeDomain in the Performance domain.
type SetTimeDomainArgs struct {
	// TimeDomain Time domain
	//
	// Values: "timeTicks", "threadTicks".
	TimeDomain string `json:"timeDomain"`
}

// NewSetTimeDomainArgs initializes SetTimeDomainArgs with the required arguments.
func NewSetTimeDomainArgs(timeDomain string) *SetTimeDomainArgs {
	args := new(SetTimeDomainArgs)
	args.TimeDomain = timeDomain
	return args
}

// GetMetricsReply represents the return values for GetMetrics in the Performance domain.
type GetMetricsReply struct {
	Metrics []Metric `json:"metrics"` // Current values for run-time metrics.
}
