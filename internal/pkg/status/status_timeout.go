package status

type (
	Timeout struct {
		Status `yaml:",inline" json:",inline"`
	}
)

func (s Timeout) IsTimeout() bool {
	return true
}
