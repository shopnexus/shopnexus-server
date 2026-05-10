ShopNexus thức giấc giữa đêm sâu,
Trong repo lớn ánh terminal màu,
Không chia cắt trăm miền repository,
Một monorepo giữ chung nhịp cầu.

Microservice đứng thành từng phân khu,
`account`, `catalog` nối lời viễn du,
`order` gọi `inventory` lặng lẽ,
Qua Restate bền bỉ giữa sương mù.

Không handler rối tung vì event,
Không truy vết quanh co giữa nghìn thread,
Flow xuôi dòng như con sông có hướng,
Retry hoài cho tới lúc thành công.

Restate như người giữ cổng âm thầm,
Ghi journal từng nhịp gọi xa xăm,
Nếu service ngủ quên trong lỗi mạng,
Ngày mai về vẫn tiếp tục đang làm.

Proxy đứng như chiếc bóng vô hình,
Method gọi nghe gần tựa tiến trình,
Nhưng phía dưới là ingress đang chảy,
HTTP mang tín hiệu đăng trình.

```go
inventories, err := orderbiz.inventory.ReserveInventory(...)
```

Một dòng code tưởng giản dị biết bao,
Sau lưng nó là orchestration cao,
Service nọ chẳng nhìn service khác,
Mọi con đường đều hội tụ một vào.

Redis lock canh cửa giữa đêm dài,
TTL trôi như cát cuốn không ngừng trôi,
Nên goroutine âm thầm gia hạn,
Giữ chiếc khóa chưa mục giữa dòng đời.

`unlock()` khép lại cuộc hành quân,
Xóa key đi như gió cuốn phù vân,
Không deadlock nằm hoang trong bóng tối,
Chỉ còn log cháy đỏ giữa machine.

SQLC rèn struct từ câu SQL,
Type-safe như kiếm thép giữa compiler,
`pgtempl` viết CRUD từ migration,
Đỡ dev ngồi gõ lặp đến tàn hơi.

`genrestate` sinh interface tựa mơ,
Proxy mọc lên từ định nghĩa chờ,
Từ Go interface thành service thật,
Code generation vang vọng từng giờ.

Trong `catalog` là muôn vàn sản phẩm,
Tag và search như sao sáng xa xăm,
`promotion` gom giảm giá chồng lớp,
`analytic` đếm từng lượt ghé thăm.

`chat` giữ lại tin nhắn người mua,
`common` chở file vượt biển object store,
`inventory` canh từng serial nhỏ,
`order` đợi seller xác nhận đơn.

Có những đêm migrate vừa chạy xong,
Schema đổi như thủy triều trên sông,
Version nối tiếp bằng tên migration,
Lặng lẽ mà nâng đỡ cả hệ thống.

ShopNexus không chỉ là backend,
Mà là bản đồ của những lần fail,
Của retry, rollback và transaction,
Của những bug thức trắng suốt nhiều đêm.

Mai sau nếu service cần tách riêng,
Không refactor đổ nát cả công trình,
Chỉ đổi config, đổi nơi deployment,
Kiến trúc xưa vẫn giữ vững thân mình.

Và giữa tiếng quạt quay cùng bàn phím,
Một developer lặng ngắm log im,
Thấy từng module như thành phố nhỏ,
Đang sống nhờ orchestration bền tim.
