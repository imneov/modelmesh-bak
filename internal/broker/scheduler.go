package broker

import (
	"fmt"
	"github.com/imneov/modelmesh/internal/broker/config"
	"github.com/imneov/modelmesh/internal/broker/picker"
	"k8s.io/klog/v2"
)

type SchedulerPick func() (picker.PickResult, error)

type Scheduler struct {
	builder   picker.PickerBuilder
	picker    picker.Picker
	errPicker picker.Picker
	pickInfo  picker.PickInfo
	reoucres  []*config.ServiceGroup
}

func NewScheduler(cfg *config.Schedule) (*Scheduler, error) {
	klog.Errorf("cfg: %v", cfg.Method)
	builder := picker.Builder(cfg.Method)
	if builder == nil {
		return nil, fmt.Errorf("builder is nil")
	}

	errPicker := picker.NewErrPicker(fmt.Errorf("picker is nil"))
	return &Scheduler{builder: builder, errPicker: errPicker, picker: nil}, nil
}

func (s *Scheduler) RegeneratePicker(groups []*config.ServiceGroup) {
	resource := map[picker.Resource]picker.ResourceInfo{}
	for _, group := range groups {
		resource[group] = picker.ResourceInfo{}
	}
	picker := s.builder.Build(picker.PickerBuildInfo{
		Resources: resource,
	})

	s.picker = picker
}

func (s *Scheduler) Pick(info picker.PickInfo) (picker.PickResult, error) {
	if s.picker == nil {
		return s.errPicker.Pick(info)
	}
	return s.picker.Pick(info)
}
