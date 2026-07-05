# Set up NetBSD for Mac OS Apple Silicon

```bash
qemu-system-x86_64
  -m 4096 \
  -smp 6 \
  -drive file=netbsd.img,format=qcow2,if=virtio \
  -netdev user,id=net0 \
  -device virtio-net-pci,netdev=net0 \
  -device virtio-rng-pci \
  -rtc base=utc \          
  -display cocoa \                 
  -serial file:serial.log \
  -serial mon:stdio \
  -device usb-ehci,id=ehci \
  -device usb-mouse,bus=ehci.0 \
  -device usb-kbd,bus=ehci.0 \
  -accel tcg
```

```bash
qemu-system-x86_64 \
  -m 4096 \
  -smp 6 \
  -drive file=netbsd.img,format=qcow2,if=virtio \
  -netdev user,id=net0 \
  -device virtio-net-pci,netdev=net0 \
  -device virtio-rng-pci \
  -rtc base=utc \
  -nographic
```

```bash
qemu-system-x86_64  
  -m 4096 \
  -smp 6 \
  -drive file=netbsd.img,format=qcow2,if=virtio \
  -netdev user,id=net0 \
  -device virtio-net-pci,netdev=net0 \
  -device virtio-rng-pci \
  -rtc base=utc \
  -nographic \
  -serial mon:stdio \
  -accel tcg
```