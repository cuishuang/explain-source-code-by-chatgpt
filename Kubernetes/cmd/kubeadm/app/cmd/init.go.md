# File: cmd/kubeadm/app/cmd/init.go

在kubernetes项目中，`cmd/kubeadm/app/cmd/init.go`文件是kubeadm init命令的入口点，它负责初始化Kubernetes集群。



- `_`：变量`_`是通配符变量，用于忽略导入的包或声明的变量，以避免引用未使用的包或变量而导致编译错误。



- `initOptions`：结构体`initOptions`定义了执行初始化命令时可用的命令行选项，例如是否跳过预检查、证书相关选项等。



- `initData`：结构体`initData`定义了执行初始化命令时需要的数据，包括配置文件路径、证书目录、kubelet数据目录等。



- `newCmdInit`：函数`newCmdInit`创建了`kubeadm init`命令的实例。



- `AddInitConfigFlags`：函数`AddInitConfigFlags`向`kubeadm init`命令添加配置标志，例如`--config`。



- `AddClusterConfigFlags`：函数`AddClusterConfigFlags`向`kubeadm init`命令添加集群配置标志，例如`--apiserver-advertise-address`。



- `AddInitOtherFlags`：函数`AddInitOtherFlags`向`kubeadm init`命令添加其他标志，例如`--skip-token-print`。



- `newInitOptions`：函数`newInitOptions`创建了`initOptions`结构体的实例。



- `newInitData`：函数`newInitData`创建了`initData`结构体的实例。



- `UploadCerts`：函数`UploadCerts`负责上传证书到指定目录。





- `CertificateKey`：函数`CertificateKey`返回证书的密钥。


- `SetCertificateKey`：函数`SetCertificateKey`设置证书的密钥。


- `SkipCertificateKeyPrint`：函数`SkipCertificateKeyPrint`设置是否跳过打印证书密钥的标志。


- `Cfg`：变量`Cfg`是一个全局变量，代表初始化命令的配置。


- `DryRun`：变量`DryRun`控制是否只运行预检查。


- `SkipTokenPrint`：变量`SkipTokenPrint`控制是否跳过打印令牌。


- `IgnorePreflightErrors`：变量`IgnorePreflightErrors`控制是否忽略预检查错误。


- `CertificateWriteDir`：变量`CertificateWriteDir`是证书写入的目录。


- `CertificateDir`：变量`CertificateDir`是证书的目录。


- `KubeConfigDir`：变量`KubeConfigDir`是kubeconfig的目录。


- `KubeConfigPath`：变量`KubeConfigPath`是kubeconfig的路径。


- `ManifestDir`：变量`ManifestDir`是manifest文件的目录。


- `KubeletDir`：变量`KubeletDir`是kubelet的数据目录。


- `ExternalCA`：变量`ExternalCA`控制是否使用外部CA。


- `OutputWriter`：变量`OutputWriter`控制输出的写入目标。


- `Client`：变量`Client`是与APIServer进行通信的客户端。


- `Tokens`：变量`Tokens`存储生成的令牌。


- `PatchesDir`：变量`PatchesDir`是补丁目录。

