// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mindspore_serving/proto/ms_master.proto

package mindspore_serving_proto

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ServableRegSpec) Validate() error {
	for _, item := range this.Methods {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Methods", err)
			}
		}
	}
	if this.ModelInfos != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ModelInfos); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ModelInfos", err)
		}
	}
	return nil
}
func (this *ServableRegSpec_MethodInfo) Validate() error {
	return nil
}
func (this *WorkerRegSpec) Validate() error {
	if this.ServableSpec != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ServableSpec); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ServableSpec", err)
		}
	}
	return nil
}
func (this *RegisterRequest) Validate() error {
	if this.WorkerSpec != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WorkerSpec); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WorkerSpec", err)
		}
	}
	return nil
}
func (this *RegisterReply) Validate() error {
	if this.ErrorMsg != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ErrorMsg); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ErrorMsg", err)
		}
	}
	return nil
}
func (this *ExitRequest) Validate() error {
	return nil
}
func (this *ExitReply) Validate() error {
	if this.ErrorMsg != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ErrorMsg); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ErrorMsg", err)
		}
	}
	return nil
}
func (this *NotifyFailedRequest) Validate() error {
	return nil
}
func (this *NotifyFailedReply) Validate() error {
	return nil
}
func (this *GetModelInfoRequest) Validate() error {
	return nil
}
func (this *ModelSubGraphInfo) Validate() error {
	for _, item := range this.Inputs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Inputs", err)
			}
		}
	}
	for _, item := range this.Outputs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Outputs", err)
			}
		}
	}
	return nil
}
func (this *ModelInfo) Validate() error {
	for _, item := range this.SubgraphInfos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SubgraphInfos", err)
			}
		}
	}
	return nil
}
func (this *ModelInfos) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *GetModelInfoReply) Validate() error {
	if this.ModelInfos != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ModelInfos); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ModelInfos", err)
		}
	}
	if this.ErrorMsg != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ErrorMsg); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ErrorMsg", err)
		}
	}
	return nil
}
