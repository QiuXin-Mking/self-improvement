# q
已知基带信号 $x(t)$ 的奈奎斯特采样率为 $\omega_s$，用载波 $\cos(\omega_s t)$ 调制得到 $y(t)=x(t)\cos(\omega_s t)$，$y(t)$ 的奈奎斯特采样率是多少？
# a
$y(t)$ 的奈奎斯特采样率为 $3\omega_s$（弧度/秒）。  
因为 $x(t)$ 的最高频率分量为 $\frac{\omega_s}{2}$，调制后频谱被搬移到 $\pm\omega_s$ 处，最高频率变为 $\omega_s + \frac{\omega_s}{2} = \frac{3\omega_s}{2}$。根据奈奎斯特定理，采样率至少为最高频率的 2 倍，即 $2 \times \frac{3\omega_s}{2} = 3\omega_s$。

