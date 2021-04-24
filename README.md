## 实现区块结构以及与区块相关功能
1. 区块结构分析
2. 新建区块
3. 如何生成哈希
4. 类型转换
## 实现区块链基本结构
1. 实现链表（通过切片进行缓存）-区块链基本结构
2. 实现创世区块与区块链初始化功能
2. 实战上链功能
## 实现POW共识算法
1. pow结构分析
2. 设置目标难度值分析
3. 哈希碰撞
4. 数据准备
## 实现coinbase交易
1. coinbase生成函数实现
2. 交易哈希的实现
3. coinbase交易函数的实现
## 实现CLI发起转账
1. 添加命令行转账功能
2. 实现通过挖矿生成新的区块
3. 实现生成普通交易
4. 修改挖矿函数，调用NewSimpleTransaction()
5. 通过CLI实现普通转账交易调用
6. 实现余额查询与UTXO查询
7. 实现余额查询cli端的封装
8. 实现UTXO查找封装
9. 实现输入输出验证功能

## UTXO查找的内部实现
1. 实现查找数据库指定地址所有已花费输出函数
2. 实现coinbase交易判断函数
3. 实现查找指定地址所有UTXO的函数

## 命令行文件分离

## 转账逻辑完善与UTXO查找优化
1. 实现查找可用UTXO的函数
2. 实现通过UTXO查询进行转账
3. 实现多笔交易转账

## 实现钱包模块集成
1. 将钱包创建功能加入命令行操作
2. 实现获取地址列表功能
3. 钱包功能持久化

## 实现钱包与输入输出功能的结合
1. 实现输入结构与钱包功能结合
2. 实现输出结构与钱包功能结合
3. 调用方修改

## 交易签名实现
1. 在交易生成的时候对交易进行签名
2. 在交易被打包进入区块之前进行验证

## 交易签名验证
1. 在交易被打包进入区块之前进行验证

## 实现挖矿奖励
1. 默认情况下，谁发起交易，谁就得到奖励

## 实现UTXO查找优化
1. 添加UTXOSet结构
2. 添加重置utxo table的功能
3. 实现查找区块链中所有utxo的功能
4. 实现查找区块链中所有已花费输出的功能

## 实现余额查找功能优化
1. 实现通过utxo table查找指定地址的UTXO
2. 实现UTXO的余额查找函数
3. 命令行获取指定地址余额调用函数修改

## 实现utxo table实时更新
1. 实现update更新函数

## Merkle树的实现
1. Merkle节点的实现
2. Merkle树的实现
3. 交易哈希与Merkle树的联结

## 网络实现
1. 模拟两个节点，一个节点进行创建、转账等操作，另一个节点进行数据同步操作
2. 通过不同端口来模拟不同节点
3. 节点与程序进行关联-数据同步
4. 钱包文件与blockchain数据库文件与节点ID号进行关联
