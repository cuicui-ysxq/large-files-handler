# 大文件处理工具

大文件处理工具。将大文件分割成小文件以便上传到 GitHub。

## 构建

执行

```bash
./build.sh
```

构建所有可执行文件。

该脚本会创建 `./bin` 目录，并在其中放置构建好的可执行文件，包含的内容有：

* `split`：[拆分大文件](#拆分大文件)
* `merge`：[合并小文件](#合并小文件)

## 清除构建文件

执行

```bash
./clean-up.sh
```

以清除构建的可执行文件。

此操作会将 `./bin` 目录整个删除。

## 主要功能

### 拆分大文件

```bash
./bin/split . -i <input_file> -s <chunk_size> -o <output_dir>
```

* `input_file`：输入文件（**必填**）
* `chunk_size`：分块大小，单位：字节（**选填**，如未填则为 49 MB，比 GitHub 单文件大小限制 50 MB 小 1 MB）
* `output_dir`：输出目录（**选填**，如未填则为当前目录）

### 合并小文件

```bash
./bin/merge . -i <input_file> [-i <input_file>...] [-b <buffer_size>] -o <output_file>
```

* `input_file`：输入文件（**至少填 1 个**，如果有多个则重复输入 `-i <input_file>`）
* `buffer_size`：缓冲区大小，单位：字节（**选填**，如未填则为 16 MB）
* `output_file`：输出文件（**必填**）