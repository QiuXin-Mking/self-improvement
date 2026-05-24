# q
KVM XML 中 `<interface type='bridge'>` 配置包含哪些核心元素？
# a
核心元素包括：
- `<source bridge='...'/>`：指定宿主机上的网桥名称，如 `br-openstack`。
- `<virtualport type='openvswitch'>`：标注该桥接网卡可与 OpenvSwitch 集成，适用于 SDN 场景（如 OpenStack Neutron）。
- `<target dev='...'/>`：定义虚拟网卡在虚拟机内部的设备名称，如 `net30_vpxe`。
- `<model type='virtio'/>`：指定网卡驱动模型为 `virtio`（高性能准虚拟化驱动）。
- `<alias name='...'/>`：该网卡的内部别称，如 `net0`。

# q
KVM 网卡模型中 `virtio` 的作用是什么？
# a
`virtio` 是 KVM/虚拟化环境下的一种高性能准虚拟化驱动，可显著提升网络 I/O 性能，降低开销。在 XML 中通过 `<model type='virtio'/>` 启用。

