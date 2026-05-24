# q
KVM是什么，全拼是什么
# a
KVM（Kernel-based Virtual Machine）是Linux内核中的一个虚拟化模块，能将Linux变成Hypervisor，让宿主机直接运行多个虚拟机。

# q
如何具体理解KVM和QEMU的关系与分工
# a
QEMU像一个“超级模拟器”，负责模拟各种硬件设备，充当虚拟机的“管家”；KVM是Linux自带的“虚拟化引擎”，利用硬件辅助虚拟化让虚拟机运行速度接近物理机。通常QEMU负责I/O和平台管理，KVM负责CPU和内存的硬件加速。

# q
如何从内核层面理解KVM
# a
KVM本质上是一组内核模块（主要是`kvm.ko`以及架构相关的`kvm-intel.ko`或`kvm-amd.ko`），作为Linux内核的可加载驱动，在内核空间中提供虚拟化核心功能。

# q
虚拟机的上下文切换主要指什么
# a
指CPU在虚拟机客体（Guest）和宿主机（Host）之间切换执行状态的过程。当发生高权限硬件中断时，虚拟机会通过VM-exit将控制权交还给宿主机内核处理，处理完毕后再通过VM-entry恢复虚拟机的运行。

# q
CPU的Ring0、Ring1等是什么
# a
CPU的Ring是操作系统和应用程序访问硬件资源的权限级别，又称特权级（Privilege Level）或环。现代操作系统通常只用Ring 0（内核态）和Ring 3（用户态），Ring 1和Ring 2很少使用。

