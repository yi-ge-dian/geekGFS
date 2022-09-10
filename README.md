# GeekGFS
A Simple Implementation For GFS（Google File System）

GeekGFS具有客户端client、元数据服务器master和多个数据服务器chunkserver三个基本组件

GFSClient需要实现read、write、append、exist和delete函数功能。函数实现要求：从master获取元数据（包含chunk ID和chunk位置），更新master上的元数据，最后与chunkserver进行实际的数据传输。

GFSMasterserver模拟GFS元数据服务器，在内存中存储所有的元数据，且需要周期备份到硬盘。

GFSChunkserver实现分块存储，块固定为64KB。为简化起见，可使用同一个节点的多个路径模拟多个物理节点。

编写测试程序，进行测试验证。开发语言不限。
备注：可不考虑元数据加锁、chunk租约、复制，master故障转移、chunkserver心跳、垃圾回收等复杂机制。
