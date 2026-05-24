# q
YOLOv8-pose 标注文件中，每个关键点的数据由哪三个部分组成？
# a
每个关键点由三个数字组成：归一化的 x 坐标（范围 [0, 1]）、归一化的 y 坐标（范围 [0, 1]）和可见性标志。可见性中，`1` 表示关键点可见，`2` 表示被遮挡，`0` 表示不存在或未标记。

# q
YOLOv8 训练输出的 `Weights` 文件（.pt）和 `Args.yaml` 文件分别包含什么信息？
# a
- `Weights` 文件（.pt）：保存训练后神经网络的所有参数和权重，是模型训练的直接产物，用于后续推理。
- `Args.yaml` 文件：记录训练时使用的所有配置参数（如学习率、批大小、训练轮数等），便于复现训练过程。

# q
YOLOv8 训练产生的 `confusion_matrix.png` 和 `confusion_matrix_normalized.png` 有什么区别？
# a
`confusion_matrix.png` 是原始混淆矩阵，直接展示预测结果与真实标签的对比数值；`confusion_matrix_normalized.png` 是归一化后的混淆矩阵，将各条目转换为比例，便于比较不同类别之间的预测准确率。

# q
在 YOLOv8-pose 训练输出中，`PosePR_curve.png` 和 `PoseR_curve.png` 分别表示什么？
# a
- `PosePR_curve.png`：姿势检测任务的精确度-召回率曲线，反映模型在不同阈值下精确度与召回率的权衡。
- `PoseR_curve.png`：姿势检测任务的召回率-置信度曲线，展示召回率随置信度阈值的变化，用于寻找最优置信度阈值。

