## windows CMD:

### 测试标签管理功能

> curl http://localhost:8000/api/v1/tags -X POST -F name=Go -F created_by=eddycjy

> curl -X POST http://127.0.0.1:8000/api/v1/tags -F name=PHP -F created_by=eddycjy

> curl -X POST http://127.0.0.1:8000/api/v1/tags -F name=Rust -F created_by=eddycjy

> curl -X GET http://127.0.0.1:8000/api/v1/tags -G -d page=1 -d page_size=2

> curl -X GET http://127.0.0.1:8000/api/v1/tags -G -d page=2 -d page_size=2

### 测试文件上传

> curl -X POST http://127.0.0.1:8000/upload/file -F file=@{file_path} -F type=1