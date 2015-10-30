package consul_provisioner

import (
	"github.com/hashicorp/consul/api"
	"log"
	"os"
	"testing"
)

var consulConfig = struct {
	Address         string
	Datacenter      string
	User            string
	Pass            string
	Scheme          string
	Token           string
	TestProvisioner ConsulProvisionerInterface
}{}

func init() {
	var err error

	consulConfig.Address = os.Getenv("CONSUL_ADDRESS")
	consulConfig.Datacenter = os.Getenv("CONSUL_DATACENTER")
	consulConfig.User = os.Getenv("CONSUL_USERNAME")
	consulConfig.Pass = os.Getenv("CONSUL_PASSWORD")
	consulConfig.Scheme = os.Getenv("CONSUL_SCHEME")
	consulConfig.Token = os.Getenv("CONSUL_TOKEN")

	var config api.Config

	config.Address = consulConfig.Address
	config.Datacenter = consulConfig.Datacenter
	config.Scheme = consulConfig.Scheme
	config.Token = consulConfig.Token

	if consulConfig.User != "" {
		var auth api.HttpBasicAuth
		auth.Username = consulConfig.User
		auth.Password = consulConfig.Pass
		config.HttpAuth = &auth
	}

	consulConfig.TestProvisioner, err = New(&config)
	if err != nil {
		log.Println(err)
	}
}

func Test_AddKeyValue(t *testing.T) {
	if consulConfig.Address == "" {
		t.Skip("Skipping test as Consul env vars are not set: CONSUL_ADDRESS, CONSUL_DATACENTER, (CONSUL_USERNAME, CONSUL_PASSWORD) / CONSUL_TOKEN")
	}
	log.Println("Testing add key-value")
	err := consulConfig.TestProvisioner.AddKV("testKey", []byte("testValue"), nil)

	if err != nil {
		log.Fatalln(err)
	}
}

func Test_GetValue(t *testing.T) {
	if consulConfig.Address == "" {
		t.Skip("Skipping test as Consul env vars are not set: CONSUL_ADDRESS, CONSUL_DATACENTER, (CONSUL_USERNAME, CONSUL_PASSWORD) / CONSUL_TOKEN")
	}
	log.Println("Testing get key-value")
	val, err := consulConfig.TestProvisioner.GetValue("testKey")
	if err != nil {
		log.Fatalln(err)
	}
	if val != nil {
		log.Println("Test successful, retrieved value: ", string(val))
	}
}

func Test_AddKVList(t *testing.T) {
	var list api.KVPairs

	if consulConfig.Address == "" {
		t.Skip("Skipping test as Consul env vars are not set: CONSUL_ADDRESS, CONSUL_DATACENTER, (CONSUL_USERNAME, CONSUL_PASSWORD) / CONSUL_TOKEN")
	}
	log.Println("Testing put key-value list")

	list = append(list, &api.KVPair{Key: "usb/loglevel", Value: []byte("debug")})

	list = append(list, &api.KVPair{Key: "usb/broker_api", Value: []byte("{\"listen\":\":54054\",\"credentials\":{\"username\":\"demouser\",\"password\":\"demopassword\"}}")})

	list = append(list, &api.KVPair{Key: "usb/management_api", Value: []byte("{\"listen\":\":54053\",\"uaa_secret\":\"myuaasecret\",\"uaa_client\":\"myuaaclient\",\"authentication\":{\"uaa\":{\"adminscope\":\"usb.management.admin\",\"public_key\":\"\"}}}")})

	list = append(list, &api.KVPair{Key: "usb/drivers/mysql", Value: []byte("mysql")})

	list = append(list, &api.KVPair{Key: "usb/drivers/mysql/instances/00000000-0000-0000-0000-0000000000M1/Name", Value: []byte("mysql-local")})

	list = append(list, &api.KVPair{Key: "usb/drivers/mysql/instances/00000000-0000-0000-0000-0000000000M1/Configuration", Value: []byte("{\"server\":\"127.0.0.1\",\"port\":\"3306\",\"userid\":\"root\",\"password\":\"password\"}")})

	list = append(list, &api.KVPair{Key: "usb/drivers/mysql/instances/00000000-0000-0000-0000-0000000000M1/dials/B0000000-0000-0000-0000-000000000LM1", Value: []byte("{\"id\":\"B0000000-0000-0000-0000-000000000LM1\",\"configuration\":{},\"plan\":{\"name\":\"free\",\"id\":\"53425178-F731-49E7-9E53-5CF4BE9D807L\",\"description\":\"This is the first plan\",\"free\":true}}")})

	list = append(list, &api.KVPair{Key: "usb/drivers/mysql/instances/00000000-0000-0000-0000-0000000000M1/service", Value: []byte("{\"id\":\"83E94C97-C755-46A5-8653-461517EB442L\",\"bindable\":true,\"name\":\"MysqlLocalService\",\"description\":\"Mysql Local Service\",\"tags\":[\"mysql\"],\"metadata\":{\"providerDisplayName\":\"Mysql Local Service\"}}")})

	err := consulConfig.TestProvisioner.PutKVs(&list, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func Test_GetAllKvs(t *testing.T) {
	if consulConfig.Address == "" {
		t.Skip("Skipping test as Consul env vars are not set: CONSUL_ADDRESS, CONSUL_DATACENTER, (CONSUL_USERNAME, CONSUL_PASSWORD) / CONSUL_TOKEN")
	}
	log.Println("Testing get key-value list")
	val, err := consulConfig.TestProvisioner.GetAllKVs("usb", nil)
	if err != nil {
		log.Fatalln(err)
	}
	if val != nil {
		for _, kv := range val {
			log.Println("Retrieved value: ", kv.Key, string(kv.Value))
		}
	}
}

func Test_RemoveKey(t *testing.T) {
	if consulConfig.Address == "" {
		t.Skip("Skipping test as Consul env vars are not set: CONSUL_ADDRESS, CONSUL_DATACENTER, (CONSUL_USERNAME, CONSUL_PASSWORD) / CONSUL_TOKEN")
	}
	log.Println("Testing delete key-value")
	var options api.WriteOptions
	err := consulConfig.TestProvisioner.DeleteKV("testKey", &options)
	if err != nil {
		log.Fatalln(err)
	}
}

func Test_GetKeys(t *testing.T) {
	if consulConfig.Address == "" {
		t.Skip("Skipping test as Consul env vars are not set: CONSUL_ADDRESS, CONSUL_DATACENTER, (CONSUL_USERNAME, CONSUL_PASSWORD) / CONSUL_TOKEN")
	}
	log.Println("Testing get keys")
	result, err := consulConfig.TestProvisioner.GetAllKeys("usb/", "", nil)

	if err != nil {
		log.Fatalln(err)
	}
	if result != nil {
		for _, val := range result {
			t.Log("Received ", val)
		}
	}
}

func Test_RemoveKeyList(t *testing.T) {
	if consulConfig.Address == "" {
		t.Skip("Skipping test as Consul env vars are not set: CONSUL_ADDRESS, CONSUL_DATACENTER, (CONSUL_USERNAME, CONSUL_PASSWORD) / CONSUL_TOKEN")
	}
	log.Println("Testing delete key-value list with prefix")
	err := consulConfig.TestProvisioner.DelteKVs("usb", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
