# File: istio/pkg/workloadapi/security/authorization.pb.go

istio/pkg/workloadapi/security/authorization.pb.go文件是Istio授权API的Protocol Buffers定义文件，用于定义授权规则和策略。

以下是每个变量和结构体的作用：
- Scope_name、Scope_value、Action_name、Action_value：这些是定义在Scope和Action结构体中的常量，用于表示授权规则的作用域和动作。
- File_workloadapi_security_authorization_proto、file_workloadapi_security_authorization_proto_rawDesc、file_workloadapi_security_authorization_proto_rawDescOnce、file_workloadapi_security_authorization_proto_rawDescData、file_workloadapi_security_authorization_proto_enumTypes、file_workloadapi_security_authorization_proto_msgTypes、file_workloadapi_security_authorization_proto_goTypes、file_workloadapi_security_authorization_proto_depIdxs：这些变量是Protocol Buffers自动生成的文件描述符和枚举/消息/Go类型的定义，用于序列化和反序列化授权规则。

以下是每个结构体的作用：
- Scope：表示授权规则的作用域，可以是Namespace、Service、ServiceAccount或Host。
- Action：表示授权规则的动作，可以是Allow、Deny或Log。
- Authorization：表示一个授权规则，包含作用域、动作和其他匹配条件。
- Group：表示一个用户组。
- Rules：表示一组授权规则。
- Match：表示匹配条件，可以根据IP地址、Port等进行匹配。
- Address：表示一个IP地址范围。
- StringMatch：表示字符串匹配条件，可以根据精确匹配、前缀匹配、后缀匹配和存在性匹配进行匹配。

以下是每个函数的作用：
- Enum：返回授权规则枚举的定义。
- String：返回指定规则或匹配的字符串表示。
- Descriptor：返回授权规则的描述符。
- Type：返回授权规则的类型。
- Number：返回授权规则的编号。
- EnumDescriptor：返回授权规则枚举的描述符。
- Reset：重置授权规则对象的字段。
- ProtoMessage：实现了Protocol Buffers的消息接口。
- ProtoReflect：返回授权规则的反射接口。
- GetName：返回授权规则的名称。
- GetNamespace：返回授权规则的命名空间。
- GetScope：返回授权规则的作用域。
- GetAction：返回授权规则的动作。
- GetGroups：返回授权规则的用户组。
- GetRules：返回授权规则的规则列表。
- GetMatches：返回授权规则的匹配条件列表。
- GetNamespaces：返回授权规则的命名空间列表。
- GetNotNamespaces：返回授权规则的非命名空间列表。
- GetPrincipals：返回授权规则的主体列表。
- GetNotPrincipals：返回授权规则的非主体列表。
- GetSourceIps：返回授权规则的源IP列表。
- GetNotSourceIps：返回授权规则的非源IP列表。
- GetDestinationIps：返回授权规则的目标IP列表。
- GetNotDestinationIps：返回授权规则的非目标IP列表。
- GetDestinationPorts：返回授权规则的目标端口列表。
- GetNotDestinationPorts：返回授权规则的非目标端口列表。
- GetAddress：返回授权规则的地址。
- GetLength：返回授权规则的长度。
- GetMatchType：返回字符串匹配条件的匹配类型。
- GetExact：返回字符串匹配条件的精确匹配值。
- GetPrefix：返回字符串匹配条件的前缀匹配值。
- GetSuffix：返回字符串匹配条件的后缀匹配值。
- GetPresence：返回字符串匹配条件的存在性匹配类型。
- isStringMatch_MatchType：判断字符串匹配类型是否符合要求。
- file_workloadapi_security_authorization_proto_rawDescGZIP：返回文件描述符的GZIP压缩内容。
- init：初始化授权规则的文件描述符。
- file_workloadapi_security_authorization_proto_init：初始化Istio授权API的文件描述符。

