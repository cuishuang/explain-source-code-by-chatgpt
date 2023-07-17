# File: pkg/securitycontext/accessors.go

pkg/securitycontext/accessors.go文件的作用是提供对容器和Pod安全上下文的访问和修改功能。

以下是各个结构体和函数的详细介绍：

1. PodSecurityContextAccessor：提供对Pod安全上下文的访问功能。
2. PodSecurityContextMutator：提供对Pod安全上下文的修改功能。
3. podSecurityContextWrapper：封装了Pod的安全上下文，提供了方便的访问和修改方法。
4. ContainerSecurityContextAccessor：提供对容器安全上下文的访问功能。
5. ContainerSecurityContextMutator：提供对容器安全上下文的修改功能。
6. containerSecurityContextWrapper：封装了容器的安全上下文，提供了方便的访问和修改方法。
7. effectiveContainerSecurityContextWrapper：封装了有效容器的安全上下文，提供了方便的访问和修改方法。

以下是每个函数的具体作用：

1. NewPodSecurityContextAccessor：创建一个新的PodSecurityContextAccessor对象，用于访问Pod的安全上下文。
2. NewPodSecurityContextMutator：创建一个新的PodSecurityContextMutator对象，用于修改Pod的安全上下文。
3. PodSecurityContext：返回Pod的安全上下文。
4. ensurePodSC：确保Pod的安全上下文是非空的。
5. HostNetwork：返回指示Pod是否使用主机网络的布尔值。
6. SetHostNetwork：设置Pod是否使用主机网络。
7. HostPID：返回指示Pod是否使用主机PID命名空间的布尔值。
8. SetHostPID：设置Pod是否使用主机PID命名空间。
9. HostIPC：返回指示Pod是否使用主机IPC命名空间的布尔值。
10. SetHostIPC：设置Pod是否使用主机IPC命名空间。
11. SELinuxOptions：返回Pod的SELinux选项。
12. SetSELinuxOptions：设置Pod的SELinux选项。
13. RunAsUser：返回容器运行时的用户ID。
14. SetRunAsUser：设置容器运行时的用户ID。
15. RunAsGroup：返回容器运行时的组ID。
16. SetRunAsGroup：设置容器运行时的组ID。
17. RunAsNonRoot：返回指示容器是否以非root用户身份运行的布尔值。
18. SetRunAsNonRoot：设置容器是否以非root用户身份运行。
19. SeccompProfile：返回容器的Seccomp配置文件。
20. SetSeccompProfile：设置容器的Seccomp配置文件。
21. SupplementalGroups：返回容器的附加组列表。
22. SetSupplementalGroups：设置容器的附加组列表。
23. FSGroup：返回容器的文件系统组。
24. SetFSGroup：设置容器的文件系统组。
25. NewContainerSecurityContextAccessor：创建一个新的ContainerSecurityContextAccessor对象，用于访问容器的安全上下文。
26. NewContainerSecurityContextMutator：创建一个新的ContainerSecurityContextMutator对象，用于修改容器的安全上下文。
27. ContainerSecurityContext：返回容器的安全上下文。
28. ensureContainerSC：确保容器的安全上下文是非空的。
29. Capabilities：返回容器的能力集。
30. SetCapabilities：设置容器的能力集。
31. Privileged：返回指示容器是否为特权容器的布尔值。
32. SetPrivileged：设置容器是否为特权容器。
33. ProcMount：返回容器的/proc挂载点。
34. ReadOnlyRootFilesystem：返回指示容器的根文件系统是否为只读的布尔值。
35. SetReadOnlyRootFilesystem：设置容器的根文件系统是否为只读。
36. AllowPrivilegeEscalation：返回指示容器是否允许特权提升的布尔值。
37. SetAllowPrivilegeEscalation：设置容器是否允许特权提升。
38. NewEffectiveContainerSecurityContextAccessor：创建一个新的EffectiveContainerSecurityContextAccessor对象，用于访问有效容器的安全上下文。
39. NewEffectiveContainerSecurityContextMutator：创建一个新的EffectiveContainerSecurityContextMutator对象，用于修改有效容器的安全上下文。

这些结构体和函数提供了对容器和Pod安全上下文的灵活访问和修改功能，可以根据需要设置和管理安全策略。

