# polestar_app

A new Flutter project.

## Getting Started

This project is a starting point for a Flutter application.

A few resources to get you started if this is your first Flutter project:

- [Lab: Write your first Flutter app](https://flutter.dev/docs/get-started/codelab)
- [Cookbook: Useful Flutter samples](https://flutter.dev/docs/cookbook)

For help getting started with Flutter, view our
[online documentation](https://flutter.dev/docs), which offers tutorials,
samples, guidance on mobile development, and a full API reference.


# 编译错误解决方案

## 1. IntelliJ Idea 创建flutter工程一直卡在creating flutter project界面问题

解决办法：进入 $FLUTTER_HOME/cache目录 删除lockfile文件 重新创建或打开工程即可（其实已经创建成功但lockfile被占用了导致无法结束）  

## 2. Could not determine the dependencies of task ':app:compileDebugJavaWithJavac'

- 报错信息

```log
      FAILURE: Build failed with an exception.
      
      * What went wrong:
      Could not determine the dependencies of task ':app:compileDebugJavaWithJavac'.
      > Could not resolve all task dependencies for configuration ':app:debugCompileClasspath'.
         > Could not resolve io.flutter:flutter_embedding_debug:1.0.0-c9506cb8e93e5e8879152ff5c948b175abb5b997.
```

- 解决办法
      
修改flutter SDK目录中的
flutter/packages/flutter_tools/gradle/resolve_dependencies.gradle      
flutter/packages/flutter_tools/gradle/aar_init_script.gradle    
flutter/packages/flutter_tools/gradle/flutter.gradle

三个文件中的：https://storage.googleapis.com/download.flutter.io   替换为：http://download.flutter.io   重新编译就好了


## 3. 运行flutter channel dev 切换分支时报错

- 报错信息

```log
Downloading the Dart SDK using the BITS service failed, retrying with WebRequest...
Invoke-WebRequest : The underlying connection was closed: An unexpected error occurred on a send.
At C:\src\flutter\bin\internal\update_dart_sdk.ps1:68 char:5
+     Invoke-WebRequest -Uri $dartSdkUrl -OutFile $dartSdkZip
+     ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidOperation: (System.Net.HttpWebRequest:HttpWebRequest) [Invoke-WebRequest], WebExc
   eption
    + FullyQualifiedErrorId : WebCmdletWebResponseException,Microsoft.PowerShell.Commands.InvokeWebRequestCommand
```

- 解决办法

* 添加如下环境变量重启终端再执行一次

export PUB_HOSTED_URL=https://pub.flutter-io.cn
export FLUTTER_STORAGE_BASE_URL=https://storage.flutter-io.cn

