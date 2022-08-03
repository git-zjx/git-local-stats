# git-local-stats

使用 Go 可视化本地 Git 提交信息

## 安装
```shell
go install github.com/git-zjx/git-local-stats
```

## 使用
```shell
# 添加需要统计的库的文件夹
git-local-stats add /path/to/folder
# 可视化某一贡献者的数据
git-local-stats stats 977904037@qq.com
```

## 效果
![stats](stats.png)

## 来源
[这里这里](https://flaviocopes.com/go-git-contributions/)   
大体没有修改，只是改了显示的效果和命令执行的方式