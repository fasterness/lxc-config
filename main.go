// Package lxc-config provides a structure for working with
// lxc container configurations in Go. man(5) page
// http://man7.org/linux/man-pages/man5/lxc.container.conf.5.html
// used as reference

package lxcconfig

import (
	"fmt"
	"net"
	"reflect"
)

//
// Usage
//  package main
//  import (
//    "fmt"
//    "github.com/fasterness/lxc-config"
//  )
//  func main() {
//    c := lxcconfig.New()
//    str := c.String()
//    fmt.Println(str)
//  }

const (
	VETH    NetworkType = "veth"
	VLAN    NetworkType = "vlan"
	MACVLAN NetworkType = "macvlan"
	PHYS    NetworkType = "phys"
	X86     Arch        = "x86"
	I686    Arch        = "i686"
	X86_64  Arch        = "x86_64"
	AMD64   Arch        = "amd64"
)

// NetworkType specifies what kind of network virtualization to be used  for  the
// container.  one of the following:
//   empty: will create only the loopback interface.
//
//   veth: a peer network device is created with one side assigned to
//   the container and  the  other  side  is  attached  to  a  bridge
//   specified   by  the  lxc.network.link.
//
//   macvlan:  a  macvlan  interface  is  linked  with  the interface
//   specified by the lxc.network.link and assigned to the container.
//
//   phys:  an  already   existing   interface   specified   by   the
//   lxc.network.link is assigned to the container.
type NetworkType string

//Arch typifies a platform architecture
type Arch string

//Arch typifies a platform architecture
type IdMap string

//Config defines an lxc configuration
//    ** _IdMap_ "lxc.id_map [u|g] 0 100000 100000" Four values must be provided. First a character, either 'u', or 'g', to specify whether user or group ids are being mapped. Next is the first userid as seen in the user namespace of the container. Next is the userid as seen on the host. Finally, a range indicating the number of consecutive ids to map.
//    ** _NetworkType_ "lxc.network.type [none|empty|veth|vlan|macvlan|phys]" specify what kind of network virtualization to be used for the container.
//    ** _NetworkLink_ "lxc.network.link lxcbr0" specify the interface to be used for real network traffic.
//    ** _NetworkFlags_ "lxc.network.flags up" specify an action to do for the network.
//    ** _NetworkName_ "lxc.network.name eth0" The name to be used for the network link inside the container
//    ** _NetworkMacAddress_ "lxc.network.hwaddr" the interface mac address is dynamically allocated by default to the virtual interface, but in  some cases, this is needed to resolve a mac address conflict or to always have the same  link-local ipv6 address
//    ** _AddressV4_ "lxc.network.ipv4 192.168.1.123/24" specify the ipv4 address to assign to the virtualized interface. Several lines specify several ipv4 addresses.  The address is in format x.y.z.t/m, eg. 192.168.1.123/24.
//    ** _AddressV6_ "lxc.network.ipv6 2003:db8:1:0:214:1234:fe0b:3596/64" specify the ipv6 address to assign to the virtualized interface. Several lines specify several ipv6 addresses.  The address is in format x::y/m, eg. 2003:db8:1:0:214:1234:fe0b:3596/64
//    ** _MacVlanMode_ "lxc.network.macvlan.mode [private|vepa|bridge]"
//    ** _AppArmorProfile_ "lxc.aa_profile unconfined" Specify the apparmor profile under which the container should be run.
//    ** _Rootfs_ "lxc.rootfs /mnt/lxc" specify a file location containing the new file tree for a root file system.
//    ** _Utsname_ "lxc.utsname fasterness" hostname for the container
//    ** _Arch_ "lxc.arch [x86|i686|x86_64|amd64]" Specify the architecture for the container.
//    ** _Include_ "lxc.include /path/to/filename" Specify the file to be included. The included file must be in the same valid lxc configuration file format.
//    ** _Pts_ "lxc.pts 1" If set, the container will have a new pseudo tty instance, making this private to it.  The value specifies the maximum number of pseudo ttys allowed for a pts instance (this limitation is not implemented yet).
//    ** _Tty_ "lxc.tty 3" Specify the number of tty to make available to the container.
//    ** _Mount_ "lxc.mount /path/to/fstab.lxc" specify a file location in  the  fstab  format, containing the mount informations.
//    ** _MountEntry_ "lxc.mount.entry proc proc proc nodev,noexec,nosuid 0 0" specify a mount point corresponding to a line in the fstab format.
//    ** _CapDrop_ "lxc.cap.drop" TODO: incomplete
//    ** _Cgroup_ "lxc.cgroup" TODO: incomplete
type Config struct {
	IdMap             []IdMap           `conf:"lxc.id_map",format:"%1s %d %d %d"`
	NetworkType       NetworkType       `conf:"lxc.network.type"`
	NetworkLink       string            `conf:"lxc.network.link"`
	NetworkFlags      string            `conf:"lxc.network.flags"`
	NetworkName       string            `conf:"lxc.network.name"`
	NetworkMacAddress net.HardwareAddr  `conf:"lxc.network.hwaddr"`
	AddressV4         []net.IPNet       `conf:"lxc.network.ipv4"`
	AddressV6         []net.IPNet       `conf:"lxc.network.ipv6"`
	MacVlanMode       string            `conf:"lxc.network.macvlan.mode"`
	AppArmorProfile   string            `conf:"lxc.aa_profile"`
	Rootfs            string            `conf:"lxc.rootfs"`
	Utsname           string            `conf:"lxc.utsname"`
	Arch              Arch              `conf:"lxc.arch"`
	Include           []string          `conf:"lxc.include"`
	Pts               int               `conf:"lxc.pts"`
	Tty               int               `conf:"lxc.tty"`
	Mount             []string          `conf:"lxc.mount"`
	MountEntry        int               `conf:"lxc.mount.entry"`
	CapDrop           string            `conf:"lxc.cap.drop"`
	Cgroup            map[string]string `conf:"lxc.cgroup."`
}

//New creates a default configuration
func New() *Config {
	var c Config
	c.IdMap = []IdMap{"u 0 100000 100000", "g 0 100000 100000"}
	c.NetworkType = VETH
	c.NetworkLink = "lxcbr0"
	c.AppArmorProfile = "unconfined"
	c.Arch = "x86"
	return &c
}

//isZero returns true if variable's value is equal to its type's zero value
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

//String marshals the configuration to a string
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
