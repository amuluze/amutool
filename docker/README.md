## docker

### container

-   GetContainerIDByName： 根据名称获取指定容器ID（**Done**）
-   ListContainer： 获取所有容器（**Done**）
-   CreateContainer：创建容器（**Done**）
-   StartContainer： 启动容器（**Done**）
-   StopContainer： 停止容器（**Done**）
-   RestartContainer： 重启容器（**Done**）
-   DeleteContainer：删除容器（**Done**）
-   CopyFileToContainer： 向容器中拷贝文件（**Done**）
-   GetContainerMem：获取指定容器的内存使用情况，单位MB（**Done**）
-   GetContainerCPU： 获取指定容器 CPU 使用情况，单位百分比（**Done**）
-   ContainerLogs： 查看指定容器的日志（**Done**）
-   RenameContainer： 重命名容器（**Done**）

### image

-   ListImage：获取本地所有的镜像信息，类似 docker images（**Done**）
-   GetImageByName：根据 imageName 获取 Image 详情（**Done**）
-   GetImageByID： 根据 imageID 获取 Image 详情（**Done**）
-   RemoveImage： 删除镜像（**Done**）
-   PruneImages： 清理虚悬镜像（**Done**）
-   ImportImage： 镜像导入（**Done**）
-   ExportImage： 镜像导出（**Done**）
-   TagImage： 修改镜像 tag（**Done**）
-   SearchImage： 镜像查找（**Done**）

### network

-   ListNetwork：获取所有 Docker 网络（**Done**）
-   QueryNetwork: 根据 network ID 获取 network 详情（**Done**）
-   CreateNetwork： 创建 Docker 网络（**Done**）
-   DeleteNetwork：删除 Docker 网络（**Done**）
-   PruneNetWork： 清理虚悬 Docker 网络（**Done**）

### volume

-   ListVolume： 获取所有 Docker 数据卷
-   CreateVolume： 创建 Docker 数据卷
-   DeleteVolume： 删除 Docker 数据卷
-   PruneVolume： 清理 Docker 数据卷

## compose

docker-compose.yaml 文件生成。