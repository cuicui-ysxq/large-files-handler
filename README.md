# 大文件处理工具

大文件处理工具。将大文件分割成小文件以便上传到GitHub。

## 拆分大文件

```bash
cd ./src/main/split
go run . -i <input_file> -s <chunk_size> -o <output_dir>
```

* `input_file`：输入文件（**必填**）
* `chunk_size`：分块大小，单位：字节（**选填**，如未填则为 49MB，比 GitHub 单文件大小限制 50MB 小 1MB）
* `output_dir`：输出目录（**选填**，如未填则为当前目录）
