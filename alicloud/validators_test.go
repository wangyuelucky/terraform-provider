package alicloud

import "testing"

func TestValidateInstancePort(t *testing.T) {
	validPorts := []int{1, 22, 80, 100, 8088, 65535}
	for _, v := range validPorts {
		_, errors := validateInstancePort(v, "instance_port")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid instance port number between 1 and 65535: %q", v, errors)
		}
	}

	invalidPorts := []int{-10, -1, 0}
	for _, v := range invalidPorts {
		_, errors := validateInstancePort(v, "instance_port")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid instance port number", v)
		}
	}
}

func TestValidateInstanceProtocol(t *testing.T) {
	validProtocals := []string{"http", "tcp", "https", "udp"}
	for _, v := range validProtocals {
		_, errors := validateInstanceProtocol(v, "instance_protocal")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid instance protocol: %q", v, errors)
		}
	}

	invalidProtocals := []string{"HTTP", "abc", "ecmp", "dubbo"}
	for _, v := range invalidProtocals {
		_, errors := validateInstanceProtocol(v, "instance_protocal")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid instance protocol", v)
		}
	}
}

func TestValidateInstanceDiskCategory(t *testing.T) {
	validDiskCategory := []string{"cloud", "cloud_efficiency", "cloud_ssd"}
	for _, v := range validDiskCategory {
		_, errors := validateDiskCategory(v, "instance_disk_category")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid instance disk category: %q", v, errors)
		}
	}

	invalidDiskCategory := []string{"all", "ephemeral", "ephemeral_ssd", "ALL", "efficiency"}
	for _, v := range invalidDiskCategory {
		_, errors := validateDiskCategory(v, "instance_disk_category")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid instance disk category", v)
		}
	}
}

func TestValidateInstanceName(t *testing.T) {
	validInstanceName := []string{"hi", "hi http://", "some word + any word &", "http", "中文"}
	for _, v := range validInstanceName {
		_, errors := validateInstanceName(v, "instance_name")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid instance name: %q", v, errors)
		}
	}

	invalidInstanceName := []string{"y", "http://", "https://", "+"}
	for _, v := range invalidInstanceName {
		_, errors := validateInstanceName(v, "instance_name")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid instance name", v)
		}
	}
}

func TestValidateInstanceDescription(t *testing.T) {
	validInstanceDescription := []string{"hi", "hi http://", "some word + any word &", "http://", "中文"}
	for _, v := range validInstanceDescription {
		_, errors := validateInstanceDescription(v, "instance_description")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid instance description: %q", v, errors)
		}
	}

	invalidvalidInstanceDescription := []string{"y", ""}
	for _, v := range invalidvalidInstanceDescription {
		_, errors := validateInstanceName(v, "instance_description")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid instance description", v)
		}
	}
}

func TestValidateSecurityGroupName(t *testing.T) {
	validSecurityGroupName := []string{"hi", "hi http://", "some word + any word &", "http", "中文", "12345"}
	for _, v := range validSecurityGroupName {
		_, errors := validateSecurityGroupName(v, "security_group_name")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid security group name: %q", v, errors)
		}
	}

	invalidSecurityGroupName := []string{"y", "http://", "https://", "+"}
	for _, v := range invalidSecurityGroupName {
		_, errors := validateSecurityGroupName(v, "security_group_name")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid security group name", v)
		}
	}
}

func TestValidateSecurityGroupDescription(t *testing.T) {
	validSecurityGroupDescription := []string{"hi", "hi http://", "some word + any word &", "http://", "中文"}
	for _, v := range validSecurityGroupDescription {
		_, errors := validateSecurityGroupDescription(v, "security_group_description")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid security group description: %q", v, errors)
		}
	}

	invalidSecurityGroupDescription := []string{"y", ""}
	for _, v := range invalidSecurityGroupDescription {
		_, errors := validateSecurityGroupDescription(v, "security_group_description")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid security group description", v)
		}
	}
}

func TestValidateCIDRNetworkAddress(t *testing.T) {
	validCIDRNetworkAddress := []string{"192.168.10.0/24", "0.0.0.0/0", "10.121.10.0/24"}
	for _, v := range validCIDRNetworkAddress {
		_, errors := validateCIDRNetworkAddress(v, "cidr_network_address")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid cidr network address: %q", v, errors)
		}
	}

	invalidCIDRNetworkAddress := []string{"1.2.3.4", "0x38732/21"}
	for _, v := range invalidCIDRNetworkAddress {
		_, errors := validateCIDRNetworkAddress(v, "cidr_network_address")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid cidr network address", v)
		}
	}
}

func TestValidateSwitchCIDRNetworkAddress(t *testing.T) {
	validSwitchCIDRNetworkAddress := []string{"192.168.10.0/24", "0.0.0.0/16", "127.0.0.0/29", "10.121.10.0/24"}
	for _, v := range validSwitchCIDRNetworkAddress {
		_, errors := validateSwitchCIDRNetworkAddress(v, "switch_cidr_network_address")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid switch cidr network address: %q", v, errors)
		}
	}

	invalidSwitchCIDRNetworkAddress := []string{"1.2.3.4", "0x38732/21", "10.121.10.0/15", "10.121.10.0/30", "256.121.10.0/22"}
	for _, v := range invalidSwitchCIDRNetworkAddress {
		_, errors := validateSwitchCIDRNetworkAddress(v, "switch_cidr_network_address")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid switch cidr network address", v)
		}
	}
}

func TestValidateIoOptimized(t *testing.T) {
	validIoOptimized := []string{"", "none", "optimized"}
	for _, v := range validIoOptimized {
		_, errors := validateIoOptimized(v, "ioOptimized")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid IoOptimized value: %q", v, errors)
		}
	}

	invalidIoOptimized := []string{"true", "ioOptimized"}
	for _, v := range invalidIoOptimized {
		_, errors := validateIoOptimized(v, "ioOptimized")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid IoOptimized value", v)
		}
	}
}

func TestValidateInstanceNetworkType(t *testing.T) {
	validInstanceNetworkType := []string{"", "Classic", "Vpc"}
	for _, v := range validInstanceNetworkType {
		_, errors := validateInstanceNetworkType(v, "instance_network_type")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid instance network type value: %q", v, errors)
		}
	}

	invalidInstanceNetworkType := []string{"classic", "vswitch", "123"}
	for _, v := range invalidInstanceNetworkType {
		_, errors := validateInstanceNetworkType(v, "instance_network_type")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid instance network type value", v)
		}
	}
}

func TestValidateInstanceChargeType(t *testing.T) {
	validInstanceChargeType := []string{"", "PrePaid", "PostPaid"}
	for _, v := range validInstanceChargeType {
		_, errors := validateInstanceChargeType(v, "instance_charge_type")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid instance charge type value: %q", v, errors)
		}
	}

	invalidInstanceChargeType := []string{"prepay", "yearly", "123"}
	for _, v := range invalidInstanceChargeType {
		_, errors := validateInstanceChargeType(v, "instance_charge_type")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid instance charge type value", v)
		}
	}
}

func TestValidateInternetChargeType(t *testing.T) {
	validInternetChargeType := []string{"", "PayByBandwidth", "PayByTraffic"}
	for _, v := range validInternetChargeType {
		_, errors := validateInternetChargeType(v, "internet_charge_type")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid internet charge type value: %q", v, errors)
		}
	}

	invalidInternetChargeType := []string{"paybybandwidth", "paybytraffic", "123"}
	for _, v := range invalidInternetChargeType {
		_, errors := validateInternetChargeType(v, "internet_charge_type")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid internet charge type value", v)
		}
	}
}

func TestValidateInternetMaxBandWidthOut(t *testing.T) {
	validInternetMaxBandWidthOut := []int{1, 22, 100}
	for _, v := range validInternetMaxBandWidthOut {
		_, errors := validateInternetMaxBandWidthOut(v, "internet_max_bandwidth_out")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid internet max bandwidth out value: %q", v, errors)
		}
	}

	invalidInternetMaxBandWidthOut := []int{-2, 0, 101, 123}
	for _, v := range invalidInternetMaxBandWidthOut {
		_, errors := validateInternetMaxBandWidthOut(v, "internet_max_bandwidth_out")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid internet max bandwidth out value", v)
		}
	}
}

func TestValidateSlbName(t *testing.T) {
	validSlbName := []string{"h", "http://", "123", "hello, aliyun! "}
	for _, v := range validSlbName {
		_, errors := validateSlbName(v, "slb_name")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid slb name: %q", v, errors)
		}
	}

	// todo: add invalid case
}

func TestValidateSlbInternetChargeType(t *testing.T) {
	validSlbInternetChargeType := []string{"paybybandwidth", "paybytraffic"}
	for _, v := range validSlbInternetChargeType {
		_, errors := validateSlbInternetChargeType(v, "slb_internet_charge_type")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid slb internet charge type value: %q", v, errors)
		}
	}

	invalidSlbInternetChargeType := []string{"PayByBandwidth", "PayByTraffic"}
	for _, v := range invalidSlbInternetChargeType {
		_, errors := validateSlbInternetChargeType(v, "slb_internet_charge_type")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid slb internet charge type value", v)
		}
	}
}

func TestValidateSlbBandwidth(t *testing.T) {
	validSlbBandwidth := []int{1, 22, 1000}
	for _, v := range validSlbBandwidth {
		_, errors := validateSlbBandwidth(v, "slb_bandwidth")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid slb bandwidth value: %q", v, errors)
		}
	}

	invalidSlbBandwidth := []int{-2, 0, 1001}
	for _, v := range invalidSlbBandwidth {
		_, errors := validateSlbBandwidth(v, "slb_bandwidth")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid slb bandwidth value", v)
		}
	}
}

func TestValidateSlbListenerBandwidth(t *testing.T) {
	validSlbListenerBandwidth := []int{-1, 1, 22, 1000}
	for _, v := range validSlbListenerBandwidth {
		_, errors := validateSlbListenerBandwidth(v, "slb_bandwidth")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid slb listener bandwidth value: %q", v, errors)
		}
	}

	invalidSlbListenerBandwidth := []int{-2, 0, -10, 1001}
	for _, v := range invalidSlbListenerBandwidth {
		_, errors := validateSlbListenerBandwidth(v, "slb_bandwidth")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid slb listener bandwidth value", v)
		}
	}
}
