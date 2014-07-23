# lxcconfig

Package lxc-config provides a structure for working with
lxc container configurations in Go. man(5) page
http://man7.org/linux/man-pages/man5/lxc.container.conf.5.html
used as reference

--
    import "github.com/fasterness/lxc-config"


## Usage

```go
Usage
 package main
 import (
   "fmt"
   "github.com/fasterness/lxc-config"
 )
 func main() {
   c := lxcconfig.New()
   str := c.String()
   fmt.Println(str)
 }
 ```

#### type Arch

```go
type Arch string
```

Arch typifies a platform architecture

#### type Config

```go
type Config struct {
	IdMap             []string          `conf:"lxc.id_map",format:"%1s %d %d %d"`
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
```

Config defines an lxc configuration
+ *IdMap*			          ```lxc.id_map [u|g] 0 100000 100000``` Four values must be provided. First a character, either 'u', or 'g', to specify whether user or group ids are being mapped. Next is the first userid as seen in the user namespace of the container. Next is the userid as seen on the host. Finally, a range indicating the number of consecutive ids to map.
+ *NetworkType*			    ```lxc.network.type [none|empty|veth|vlan|macvlan|phys]``` specify what kind of network virtualization to be used for the container.
+ *NetworkLink*			    ```lxc.network.link lxcbr0``` specify the interface to be used for real network traffic.
+ *NetworkFlags*			  ```lxc.network.flags up``` specify an action to do for the network.
+ *NetworkName*			    ```lxc.network.name eth0``` The name to be used for the network link inside the container
+ *NetworkMacAddress*	  ```lxc.network.hwaddr``` the interface mac address is dynamically allocated by default to the virtual interface, but in  some cases, this is needed to resolve a mac address conflict or to always have the same  link-local ipv6 address
+ *AddressV4*			      ```lxc.network.ipv4 192.168.1.123/24``` specify the ipv4 address to assign to the virtualized interface. Several lines specify several ipv4 addresses.  The address is in format x.y.z.t/m, eg. 192.168.1.123/24.
+ *AddressV6*			      ```lxc.network.ipv6 2003:db8:1:0:214:1234:fe0b:3596/64``` specify the ipv6 address to assign to the virtualized interface. Several lines specify several ipv6 addresses.  The address is in format x::y/m, eg. 2003:db8:1:0:214:1234:fe0b:3596/64
+ *MacVlanMode*			    ```lxc.network.macvlan.mode [private|vepa|bridge]```
+ *AppArmorProfile*			```lxc.aa_profile unconfined``` Specify the apparmor profile under which the container should be run.
+ *Rootfs*			        ```lxc.rootfs /mnt/lxc``` specify a file location containing the new file tree for a root file system.
+ *Utsname*			        ```lxc.utsname fasterness``` hostname for the container
+ *Arch*			          ```lxc.arch [x86|i686|x86_64|amd64]``` Specify the architecture for the container.
+ *Include*			        ```lxc.include /path/to/filename``` Specify the file to be included. The included file must be in the same valid lxc configuration file format.
+ *Pts*			            ```lxc.pts 1``` If set, the container will have a new pseudo tty instance, making this private to it.  The value specifies the maximum number of pseudo ttys allowed for a pts instance (this limitation is not implemented yet).
+ *Tty*			            ```lxc.tty 3``` Specify the number of tty to make available to the container.
+ *Mount*			          ```lxc.mount /path/to/fstab.lxc``` specify a file location in  the  fstab  format, containing the mount informations.
+ *MountEntry*			    ```lxc.mount.entry proc proc proc nodev,noexec,nosuid 0 0``` specify a mount point corresponding to a line in the fstab format.
+ *CapDrop*			        ```lxc.cap.drop``` TODO: incomplete
+ *Cgroup*			        ```lxc.cgroup``` TODO: incomplete

#### func  New

```go
func New() *Config
```
New creates a default configuration

#### func (*Config) String

```go
func (c *Config) String() string
```
String marshals the configuration to a string

#### type IdMap

```go
type IdMap string
```

Arch typifies a platform architecture

#### type NetworkType

```go
type NetworkType string
```

NetworkType specifies what kind of network virtualization to be used for the
container. one of the following:

    empty: will create only the loopback interface.

    veth: a peer network device is created with one side assigned to
    the container and  the  other  side  is  attached  to  a  bridge
    specified   by  the  lxc.network.link.

    macvlan:  a  macvlan  interface  is  linked  with  the interface
    specified by the lxc.network.link and assigned to the container.

    phys:  an  already   existing   interface   specified   by   the
    lxc.network.link is assigned to the container.
