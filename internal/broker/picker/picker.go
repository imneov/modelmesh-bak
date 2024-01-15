package picker

const WeightAttributeKey = "customize_weight"

// PickerBuilder creates Picker.
type PickerBuilder interface {
	// Build returns a picker that will be used by gRPC to pick a SubConn.
	Build(info PickerBuildInfo) Picker
}

// PickerBuildInfo contains information needed by the picker builder to
// construct a picker.
// 提供排序所需的基础信息
type PickerBuildInfo struct {
	// ReadySCs is a map from all ready SubConns to the Addresses used to
	// create them.
	ReadySCs map[PickerKey]PickerInfo
}

// PickerInfo contains information about a SubConn created by the base
// balancer.
// 提供排序所需的信息
type PickerInfo struct {
	//Address resolver.Address // the address used to create this SubConn
	Address struct {
		Weight int32
	}
}
