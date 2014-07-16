package lxcconfig

import (
	"fmt"
	"net"
	"reflect"
)

type NetworkType string
type Arch string
type IdMap string
type Config struct {
	IdMap             []IdMap           `conf:"lxc.id_map"`         //= "u 0 100000 100000"
	NetworkType       NetworkType       `conf:"lxc.network.type"`   //= "veth"
	NetworkLink       string            `conf:"lxc.network.link"`   //= "lxcbr0"
	NetworkFlags      string            `conf:"lxc.network.flags"`  // up
	NetworkName       string            `conf:"lxc.network.name"`   //eg: eth0, inside the container
	NetworkMacAddress string            `conf:"lxc.network.hwaddr"` //
	AddressV4         []net.IPNet       `conf:"lxc.network.ipv4"`
	AddressVr         []net.IPNet       `conf:"lxc.network.ipv6"`
	MacVlanMode       string            `conf:"lxc.network.macvlan.mode"`
	AppArmorProfile   string            `conf:"lxc.aa_profile"` //= "unconfined"
	Rootfs            string            `conf:"lxc.rootfs"`
	Utsname           string            `conf:"lxc.utsname"`
	Arch              string            `conf:"lxc.arch"`
	Include           []string          `conf:"lxc.include"`
	Pts               int               `conf:"lxc.pts"`
	Tty               int               `conf:"lxc.tty"`
	Mount             []string          `conf:"lxc.mount"`
	MountEntry        int               `conf:"lxc.mount.entry"`
	CapDrop           string            `conf:"lxc.cap.drop"`
	Cgroup            map[string]string `conf:"lxc.cgroup"`
}

const (
	VETH    NetworkType = "veth"
	VLAN    NetworkType = "vlan"
	MACVLAN NetworkType = "macvlan"
	PHYS    NetworkType = "phys"
)

func New() *Config {
	var c Config
	c.IdMap = []IdMap{"u 0 100000 100000", "g 0 100000 100000"}
	c.NetworkType = VETH
	c.NetworkLink = "lxcbr0"
	c.AppArmorProfile = "unconfined"
	c.Arch = "x86"
	return &c
}
func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && isZero(v.Index(i))
		}
		return z
	case reflect.Struct:
		z := true
		for i := 0; i < v.NumField(); i++ {
			z = z && isZero(v.Field(i))
		}
		return z
	}
	z := reflect.Zero(v.Type())
	return v.Interface() == z.Interface()
}
func (c *Config) String() string {
	config := ""
	t := reflect.TypeOf(c).Elem()
	v := reflect.ValueOf(*c)
	num := t.NumField()
	for i := 0; i < num; i++ {
		f := t.Field(i)
		fv := v.Field(i)
		if !isZero(fv) {
			switch v := fv.Interface().(type) {
			case []string:
			case []net.IPNet:
			case []IdMap:
				for _, x := range v {
					config = fmt.Sprintf("%s\n%s: %s", config, f.Tag.Get("conf"), x)
				}
			case map[string]string:
				for k, x := range v {
					config = fmt.Sprintf("%s\n%s.%s: %s", config, f.Tag.Get("conf"), k, v, x)
				}
			default:
				config = fmt.Sprintf("%s\n%s: %s", config, f.Tag.Get("conf"), v)
			}
		}
	}
	return config
}
