![](static-files/newbee-mall.png)

## 代码已迁移至[newbee-mall仓库](https://github.com/newbee-ltd/newbee-mall-api-go),本地址后续不维护
### newbee-mall商城go版本

本项目为[新蜂商城后端接口 newbee-mall-api](https://github.com/newbee-ltd/newbee-mall-api) 的go版本,使用原版本的所有数据结构 
技术栈为 Gin，主要面向服务端开发人员，前端 Vue 页面源码在另外三个 Vue 仓库。
前端项目：
- [新蜂商城 Vue2 版本 newbee-mall-vue-app](https://github.com/newbee-ltd/newbee-mall-vue-app)
- [新蜂商城 Vue3 版本 newbee-mall-vue3-app](https://github.com/newbee-ltd/newbee-mall-vue3-app)
- [新蜂商城后台管理系统 Vue3 版本 vue3-admin](https://github.com/newbee-ltd/vue3-admin)

**如果觉得项目还不错的话可以给项目一个 Star 吧，**
## 联系作者

> 大家有任何问题或者建议都可以在 [issues](https://github.com/627886474/newbee-mall-api-go/issues) 中反馈给我

### 项目讲解
-- --
- [【go商城】gin+gorm实现CRUD](https://blog.csdn.net/zxc19854/article/details/125267635)

### 本地启动
-- --
#### 后端项目启动
首先执行static-files 中的sql文件

```bash
# 克隆项目
git clone https://github.com/627886474/newbee-mall-api-go

# 使用 go mod 并安装go依赖包
go generate
# 编译 
go build -o server main.go (windows编译命令为go build -o server.exe main.go )
# 运行二进制
./server (windows运行命令为 server.exe)
```
#### 前端项目启动

然后按照原项目的部署说明部署即可

[后台管理项目](https://github.com/newbee-ltd/vue3-admin)

测试用户名：admin  测试密码：123456


[前台商城](https://github.com/newbee-ltd/newbee-mall-vue3-app)

直接注册账号就可以了
-- --


## 页面展示

以下为新蜂商城 Vue 版本的页面预览：

- 登录页

![](static-files/登录.png)

- 首页

![](static-files/首页.png)

- 商品搜索

![](static-files/商品搜索.png)

- 商品详情页

![](static-files/详情页.png)

- 购物车

![](static-files/购物车.png)

- 生成订单

![](static-files/生成订单.png)

- 地址管理

![](static-files/地址管理.png)

- 订单列表

![](static-files/订单列表.png)

- 订单详情

![](static-files/订单详情.png)

## 感谢
- [newbee-ltd](https://github.com/newbee-ltd)

- [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin)
