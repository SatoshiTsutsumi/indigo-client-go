package indigo

import (
	"fmt"
	"testing"
	"time"
)

/*
    (1, 1, 1)
	Ubuntu 18.04 + vCPU1
	{"success":true,"message":"Instance has been created successfully","vms":{"id":319016,"instance_name":"VM00","set_no":14,"vps_kind":"10","sequence_id":319016,"user_id":2130,"service_id":"idg-00012097","status":"READY","sshkey_id":20164,"snapshot_id":0,"created_at":{"date":"2022-04-10 13:05:53.109426","timezone_type":3,"timezone":"UTC"},"host_id":105,"plan_id":1,"plan":"1CR1GB","disk_point":0,"memsize":1,"cpus":1,"os_id":1,"otherstatus":10,"uuid":null,"uidgid":329016,"vnc_port":0,"vnc_passwd":"x4AxZOeElaWoi0dt","arpaname":null,"arpadate":0,"started_at":null,"closed_at":null,"vm_revert":0,"ipaddress":null,"macaddress":null,"instancetype_id":1,"import_instance":0,"container_id":null,"daemonstatus":"","outofstock":0,"ipaddress_type":"ipv4"}}

    (1, 2, 1)
	CentOS 7.5 + vCPU1
    {"success":true,"message":"Instance has been created successfully","vms":{"id":319024,"instance_name":"VM00","set_no":14,"vps_kind":"10","sequence_id":319024,"user_id":2130,"service_id":"idg-00012097","status":"READY","sshkey_id":20168,"snapshot_id":0,"created_at":{"date":"2022-04-10 13:24:17.396810","timezone_type":3,"timezone":"UTC"},"host_id":104,"plan_id":1,"plan":"1CR1GB","disk_point":0,"memsize":1,"cpus":1,"os_id":2,"otherstatus":10,"uuid":null,"uidgid":329024,"vnc_port":0,"vnc_passwd":"dqY5PI2UrhcOzc3o","arpaname":null,"arpadate":0,"started_at":null,"closed_at":null,"vm_revert":0,"ipaddress":null,"macaddress":null,"instancetype_id":1,"import_instance":0,"container_id":null,"daemonstatus":"","outofstock":0,"ipaddress_type":"ipv4"}}
	GetList
	[{"id":319024,"instance_name":"VM00","set_no":14,"vps_kind":10,"sequence_id":319024,"user_id":2130,"service_id":"idg-00012097","status":"OPEN","sshkey_id":20168,"snapshot_id":0,"created_at":"2022-04-10 13:24:17","host_id":104,"plan_id":1,"plan":"1CR1GB","disk_point":0,"memsize":1,"cpus":1,"os_id":2,"otherstatus":20,"uuid":"ecf5a61f-9c79-4587-8ecb-7035950449a6","uidgid":329024,"vnc_port":11987,"vnc_passwd":"dqY5PI2UrhcOzc3o","arpaname":"164-70-124-240.pro.static.arena.ne.jp","arpadate":1649597196,"started_at":"2022-04-10 13:26:36","closed_at":null,"vm_revert":0,"ipaddress":"164.70.124.240","macaddress":"00:13:5D:13:69:61","instancetype_id":1,"import_instance":0,"container_id":null,"daemonstatus":"","outofstock":0,"ipaddress_type":"ipv4","VEID":"14100000319024","os":{"id":2,"categoryid":2,"code":"CentOS","name":"CentOS7.5","viewname":"CentOS 7.5","instancetype_id":1},"ip":"164.70.124.240","instancecreatedbyuser":"satoshi.tsutsumi ","regionname":"Tokyo","instancestatus":"Stopped","instancetype":{"id":1,"name":"instance","display_name":"KVM Instance","created_at":"2019-10-09 13:11:20","updated_at":"2019-10-09 13:11:20"}}]

	(1, 3, 1)
	Import ULR required

	(1, 4, 1)
	OS was not found

	(Ubuntu 20.04 + IPv6 only)
	GetList
	[{"id":319025,"instance_name":"Ubuntu-1vCPU768MB20GB-03","set_no":18,"vps_kind":10,"sequence_id":319025,"user_id":2130,"service_id":"idg-00012097","status":"READY","sshkey_id":20170,"snapshot_id":0,"created_at":"2022-04-10 13:30:46","host_id":192,"plan_id":13,"plan":"1CR768MB","disk_point":0,"memsize":0.768,"cpus":1,"os_id":6,"otherstatus":10,"uuid":null,"uidgid":329025,"vnc_port":0,"vnc_passwd":"lHEQUKbMPDeX1eIx","arpaname":null,"arpadate":0,"started_at":null,"closed_at":null,"vm_revert":0,"ipaddress":null,"macaddress":null,"instancetype_id":1,"import_instance":0,"container_id":null,"daemonstatus":"","outofstock":0,"ipaddress_type":"ipv6","VEID":"18100000319025","os":{"id":6,"categoryid":1,"code":"Ubuntu","name":"Ubuntu20.04","viewname":"Ubuntu 20.04","instancetype_id":1},"ip":"","instancecreatedbyuser":"satoshi.tsutsumi ","regionname":"Tokyo","instancestatus":"OS installation In Progress","instancetype":{"id":1,"name":"instance","display_name":"KVM Instance","created_at":"2019-10-09 13:11:20","updated_at":"2019-10-09 13:11:20"}}]


	(Rocky Linux 8.4 + vCPU 2)
	[{"id":319029,"instance_name":"Rocky-2vCPU2GB40GB-04","set_no":14,"vps_kind":10,"sequence_id":319029,"user_id":2130,"service_id":"idg-00012097","status":"READY","sshkey_id":20170,"snapshot_id":0,"created_at":"2022-04-10 13:32:40","host_id":111,"plan_id":2,"plan":"2CR2GB","disk_point":0,"memsize":2,"cpus":2,"os_id":10,"otherstatus":10,"uuid":null,"uidgid":329029,"vnc_port":0,"vnc_passwd":"ilf2tU5hPooY7YH2","arpaname":null,"arpadate":0,"started_at":null,"closed_at":null,"vm_revert":0,"ipaddress":null,"macaddress":null,"instancetype_id":1,"import_instance":0,"container_id":null,"daemonstatus":"","outofstock":0,"ipaddress_type":"ipv4","VEID":"14100000319029","os":{"id":10,"categoryid":5,"code":"Rocky","name":"RockyLinux8.4","viewname":"Rocky Linux 8.4","instancetype_id":1},"ip":"","instancecreatedbyuser":"satoshi.tsutsumi ","regionname":"Tokyo","instancestatus":"OS installation In Progress","instancetype":{"id":1,"name":"instance","display_name":"KVM Instance","created_at":"2019-10-09 13:11:20","updated_at":"2019-10-09 13:11:20"}},{"id":319025,"instance_name":"Ubuntu-1vCPU768MB20GB-03","set_no":18,"vps_kind":10,"sequence_id":319025,"user_id":2130,"service_id":"idg-00012097","status":"OPEN","sshkey_id":20170,"snapshot_id":0,"created_at":"2022-04-10 13:30:46","host_id":192,"plan_id":13,"plan":"1CR768MB","disk_point":0,"memsize":0.768,"cpus":1,"os_id":6,"otherstatus":20,"uuid":"5147353a-054d-4cff-9b11-2854b7368d96","uidgid":329025,"vnc_port":10049,"vnc_passwd":"lHEQUKbMPDeX1eIx","arpaname":"2001:02C0:0100:0417:0018:cafe:008d:0001.pro.static.arena.ne.","arpadate":1649597581,"started_at":"2022-04-10 13:33:01","closed_at":null,"vm_revert":0,"ipaddress":"2001:02C0:0100:0417:0018:cafe:008d:0001","macaddress":"00:13:5D:13:69:8A","instancetype_id":1,"import_instance":0,"container_id":null,"daemonstatus":"","outofstock":0,"ipaddress_type":"ipv6","VEID":"18100000319025","os":{"id":6,"categoryid":1,"code":"Ubuntu","name":"Ubuntu20.04","viewname":"Ubuntu 20.04","instancetype_id":1},"ip":"2001:02C0:0100:0417:0018:cafe:008d:0001","instancecreatedbyuser":"satoshi.tsutsumi ","regionname":"Tokyo","instancestatus":"Stopped","instancetype":{"id":1,"name":"instance","display_name":"KVM Instance","created_at":"2019-10-09 13:11:20","updated_at":"2019-10-09 13:11:20"}}]
*/

func TestGetRegionList(t *testing.T) {
	time.Sleep(time.Second * 4)
	regions, err := client.GetRegionList()
	if err != nil || len(regions) == 0 {
		t.Errorf("GetRegionList() = %v, want %v", err, "'not nil'")
	}
}

func TestGetInstanceTypeList(t *testing.T) {
	time.Sleep(time.Second * 4)
	types, err := client.GetInstanceTypeList()
	if err != nil {
		t.Fatalf("GetInstanceTypeList() = %v, want %v", err, "'nil'")
	}
	fmt.Printf("%v", types)
}

func TestGetOSList(t *testing.T) {
	time.Sleep(time.Second * 4)
	oss, err := client.GetOSList()
	if err != nil {
		t.Fatalf("GetOSList() = %v, want %v", err, "'nil'")
	}
	fmt.Printf("%v", oss)
}

func TestGetInstanceSpecList(t *testing.T) {
	time.Sleep(time.Second * 4)
	oss, err := client.GetInstanceSpecList()
	if err != nil {
		t.Fatalf("GetInstanceSpecList() = %v, want %v", err, "'nil'")
	}
	for i, os := range oss {
		fmt.Printf("%d %v\n", i, os)
	}
}

var instance *NewInstance

func TestCreateInstance(t *testing.T) {
	var err error
	time.Sleep(time.Second * 4)
	instance, err = client.CreateInstance(key.ID, 1, 4, 1, "VM00")
	if err != nil {
		t.Fatalf("CreateInstance() = %v, want %v", err, "'nil'")
	}
	fmt.Printf("%v", instance)
}

func TestGetInstanceList(t *testing.T) {
	time.Sleep(time.Second * 4)
	instances, err := client.GetInstanceList()
	if err != nil {
		t.Fatalf("GetInstanceList() = %v, want %v", err, "'nil'")
	}
	fmt.Printf("%v", instances)
}

func TestUpdateInstanceStatus(t *testing.T) {
	time.Sleep(time.Second * 4)
	status, err := client.UpdateInstanceStatus(instance.ID, "destroy")
	if err != nil {
		t.Fatalf("UpdateInstanceStatus() = %v, want %v", err, "'nil'")
	}
	fmt.Printf("%v", status)
}
