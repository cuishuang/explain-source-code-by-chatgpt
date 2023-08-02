# File: runc/libcontainer/user/lookup_unix.go

在runc项目中，`lookup_unix.go` 文件的作用是实现了在Unix系统上查找用户和组的功能。

以下是对每个函数的详细介绍：

- `LookupUser`：通过用户名查找用户，返回一个 `User` 结构体。
- `LookupUid`：通过用户ID查找用户，返回一个 `User` 结构体。
- `lookupUserFunc`：内部函数，用于根据提供的查询函数查找用户。
- `LookupGroup`：通过组名查找组，返回一个 `Group` 结构体。
- `LookupGid`：通过组ID查找组，返回一个 `Group` 结构体。
- `lookupGroupFunc`：内部函数，用于根据提供的查询函数查找组。
- `GetPasswdPath`：获取密码文件的路径（通常是 `/etc/passwd`）。
- `GetPasswd`：从密码文件中读取并返回用户信息的切片。
- `GetGroupPath`：获取组文件的路径（通常是 `/etc/group`）。
- `GetGroup`：从组文件中读取并返回组信息的切片。
- `CurrentUser`：返回当前进程的有效用户ID（UID）和组ID（GID）。
- `CurrentGroup`：返回当前进程的真实组ID（GID）。
- `currentUserSubIDs`：返回当前进程的所有附属用户ID（SUID）。
- `CurrentUserSubUIDs`：返回当前进程的所有附属用户ID（SUID）的切片。
- `CurrentUserSubGIDs`：返回当前进程的所有附属组ID（SGID）的切片。
- `CurrentProcessUIDMap`：返回当前进程的用户ID（UID）映射状态。
- `CurrentProcessGIDMap`：返回当前进程的组ID（GID）映射状态。

以上这些函数都是在 runc 项目中用于处理用户和组相关操作的辅助函数。可以通过调用这些函数来查找、获取和处理用户和组的信息。

