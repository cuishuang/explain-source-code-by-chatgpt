# File: excelize/crypt.go

在Go生态Excelize项目中，excelize/crypt.go文件的作用是实现Excel文件的加密和解密功能。

blockKey：Excel文件加密时用于生成密钥的块密钥
oleIdentifier：OLE标识符，用于标识Excel文件是否加密
headerCLSID：加密头部的CLSID
difSect：差异扇区
endOfChain：结束扇区
fatSect：FAT扇区
iterCount：迭代次数
packageEncryptionChunkSize：加密块大小
packageOffset：文件偏移量
sheetProtectionSpinCount：工作表保护自旋次数
workbookProtectionSpinCount：工作簿保护自旋次数

Encryption：表示加密算法的类型
KeyData：密钥数据
DataIntegrity：表示数据完整性的类型
KeyEncryptors：密钥加密器
KeyEncryptor：单个密钥加密器
EncryptedKey：加密密钥
StandardEncryptionHeader：标准加密头
StandardEncryptionVerifier：标准加密验签器
encryption：加密对象
cfb：密码块链接模式
sector：扇区对象

Decrypt：解密Excel文件
Encrypt：加密Excel文件
extractPart：提取部分
encryptionMechanism：加密机制
standardDecrypt：标准解密
standardEncryptionVerifier：标准加密验签
standardConvertPasswdToKey：标准将密码转换为密钥
standardXORBytes：标准异或字节
encrypt：加密
standardKeyEncryption：标准密钥加密
agileDecrypt：灵活解密
convertPasswdToKey：将密码转换为密钥
hashing：哈希算法
createUInt32LEBuffer：创建32位无符号整数缓冲区
parseEncryptionInfo：解析加密信息
decrypt：解密
decryptPackage：解密包
createIV：创建初始化向量
randomBytes：生成随机字节
genISOPasswdHash：生成ISO密码哈希
writeBytes：写入字节
writeUint16：写入16位无符号整数
writeUint32：写入32位无符号整数
writeUint64：写入64位无符号整数
writeStrings：写入字符串
put：放置
compare：比较
prepare：准备
locate：定位
writeMSAT：写入MSAT
writeDirectoryEntry：写入目录条目
writeSectorChains：写入扇区链
write：写入数据

这些函数和变量共同实现了Excel文件的加密和解密流程，包括密钥生成、加密头部的设置、加密数据的填充、加密算法的选择和实现、加密文件的写入等。

