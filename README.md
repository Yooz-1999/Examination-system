## Agnes 风格 AI 任务平台（推断版）

本 README 基于你提供的页面截图与交互元素，对系统功能与接口进行合理推断与整理，供原型评审、接口对齐与后续实现参考。若与真实实现不符，请以实际服务为准并在实现阶段同步更新。

### 目录
- 概览
- 系统角色与核心能力
- 运行与开发
- 统一约定（鉴权、错误、分页、幂等、速率限制）
- 接口一览（按领域分组）
- 示例请求与响应
- 变更记录与版本策略
- 未来工作

## 概览

这是一个面向内容生产/研究的 AI 任务平台，支持：
- 创建任务（研究、报告、AI 幻灯片、AI 设计等）并异步执行
- 查看任务/作业进度与结果
- 上传附件/素材作为上下文
- 浏览模板与一键使用
- 管理个人知识库
- 账户、通知、计费/额度与移动端下载等

页面要点（来自截图推断）：
- 顶部任务输入框，支持选择“研究 / AI 幻灯片 / AI 设计”等模式
- 回形针图标支持上传文件
- 首页存在各类可复用模板卡片（趋势报告、行程指南、比赛总结等）
- 左侧存在“新任务”和任务列表；底部有“知识库”入口
- 右下角二维码用于移动端 App 获取

## 系统角色与核心能力

- 用户（终端使用者）：创建任务、上传素材、查看结果
- 任务编排器：接收任务并生成作业（jobs），异步执行与进度追踪
- 生成/推理引擎：文本/图像/多媒体生成
- 模板中心：模板的检索、查看与一键使用
- 知识库：文档空间与素材管理

## 运行与开发

本仓库包含 Go 微服务骨架（Api/Common/Srv）。README 聚焦 API 产品层规范，不代表已有代码即具备全部能力。

本地运行（示例）：
```bash
# 1) 启用 go work 空间
go work use .

# 2) 启动学生服务（示例服务）
cd Srv/stu_srv && go run stu_main.go

# 3) 启动网关/API（若存在）
cd Api && go run api_main.go
```

依赖配置：请根据 `Common/appconfig/config.yaml` 调整 MySQL/Redis/Nacos 等。

## 统一约定

### 鉴权
- 推荐使用 Bearer Token（`Authorization: Bearer <token>`）
- 登录成功后返回 `accessToken` 与到期时间

### 错误格式
```json
{
  "code": "RESOURCE_NOT_FOUND",
  "message": "task not found",
  "requestId": "a1b2c3"
}
```

### 分页
- 查询：`page`、`pageSize`；响应返回 `total`、`items`

### 幂等
- 对创建/取消类操作建议支持 `Idempotency-Key` 头

### 速率限制（可选）
- `429 Too Many Requests`，响应头返回 `X-RateLimit-*`

## 接口一览（推断）

> 以下为根据 UI 推断出的 REST 风格接口。路径与字段仅用于方案评审与对齐。

### 1. 认证与用户
- POST `/auth/login`
- POST `/auth/logout`
- GET `/me`（我的信息）
- PUT `/me`（更新个人资料）

### 2. 任务与作业执行
- POST `/tasks` 创建任务（如：研究/报告/设计）
- GET `/tasks?status=pending|running|done` 任务列表
- GET `/tasks/{taskId}` 任务详情
- POST `/tasks/{taskId}/cancel` 取消任务
- POST `/jobs` 创建作业（长耗时执行单元）
- GET `/jobs/{jobId}/status` 作业进度/状态

### 3. 研究/检索
- POST `/research/search` 通用检索
- POST `/research/deep_report` 深度研究报告
- POST `/research/summary` 快速总结
- POST `/research/brainstorm` 头脑风暴/创意
- POST `/research/review` 复盘/评审
- GET `/research/results/{id}` 结果获取

### 4. 文档与幻灯片生成
- POST `/generate/doc` 生成长文/报告
- POST `/generate/slides` 生成 AI 幻灯片（PPTX）
- GET `/exports/{exportId}?format=pdf|pptx|md` 导出

### 5. 设计与多媒体
- POST `/design/generate` AI 设计
- POST `/media/images` 生成图像
- POST `/media/videos` 生成视频

### 6. 模板中心
- GET `/templates?type=report|trend|guide&page=...`
- GET `/templates/{templateId}` 模板详情
- POST `/templates/{templateId}/use` 一键使用

### 7. 附件与素材
- POST `/uploads` 上传（回形针）
- GET `/assets/{assetId}` 获取素材
- DELETE `/assets/{assetId}` 删除素材

### 8. 知识库
- GET `/kb/spaces` 空间列表
- POST `/kb/spaces` 创建空间
- GET `/kb/spaces/{spaceId}/documents` 文档列表
- POST `/kb/spaces/{spaceId}/documents` 上传/新增文档

### 9. 收藏、推荐与通知
- POST `/content/{id}/favorite` 收藏
- GET `/favorites` 我的收藏
- GET `/feed/trending` 趋势/推荐流
- GET `/notifications` 通知列表

### 10. 计费与额度
- GET `/billing/plan` 当前套餐
- GET `/billing/usage` 用量统计

### 11. 设备/移动端
- GET `/mobile/app-link` 返回 iOS/Android 下载信息或二维码数据

### 12. 审计与日志
- GET `/audit/events?taskId=...` 操作轨迹与系统事件

## 示例

### 创建研究任务
```http
POST /tasks
Authorization: Bearer <token>
Content-Type: application/json

{
  "type": "research",
  "title": "到第100页检索并生成趋势报告",
  "params": {
    "depth": 100,
    "modes": ["search", "summary"],
    "keywords": ["VR", "fitness"]
  }
}
```

响应：
```json
{
  "taskId": "t_123",
  "status": "pending",
  "createdAt": "2025-09-29T10:00:00Z"
}
```

### 轮询作业进度
```http
GET /jobs/j_456/status
Authorization: Bearer <token>
```

响应：
```json
{
  "jobId": "j_456",
  "status": "running",
  "progress": 42,
  "etaSeconds": 120
}
```

### 上传素材
```http
POST /uploads
Authorization: Bearer <token>
Content-Type: multipart/form-data

files: ["doc1.pdf", "notes.txt"]
```

响应：
```json
{
  "assets": [
    { "assetId": "a1", "name": "doc1.pdf" },
    { "assetId": "a2", "name": "notes.txt" }
  ]
}
```

## 版本与变更

- 版本策略：推荐采用语义化版本 `MAJOR.MINOR.PATCH`
- 变更记录：请在 `docs/tasks/CHANGELOG.md` 或本 README 中维护关键接口变化

## 未来工作

- 将推断接口与实际后端代码对齐，补充 OpenAPI/Swagger 文档
- 为各接口补充鉴权范围、速率限制与错误码清单
- 覆盖自动化测试（单元/集成/E2E）与性能指标

---

维护者：项目团队（产品/架构/后端/前端/测试/数据）



