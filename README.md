# fis-receiver

fis receiver with golang

## 安装
```
bash -c "$(curl -k https://raw.githubusercontent.com/lrenc/fis-receiver/master/install.sh)"
```

## 启动

为积极响应组织号召，在启动server前需添加一个用于上传验证的token参数：

```
$ cd fis-receiver

# token 是启动服务者自定义的一串字符串（可以理解为密码）在fis-conf中需要指定token参数为相同的字符串，才能正常上传
$ nohup ./main token port & # 默认端口为8527
```

## Test

浏览器访问 http://host:port

## 其他

机器无法访问外网？可以先将代码安装到可以访问的机器，然后执行scp，或者直接联系我。
