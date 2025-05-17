from datetime import datetime
import pytz

# Ví dụ: start_date là chuỗi "2025-05-13"
start_date = "2025-05-9"

# Chuyển chuỗi start_date thành datetime đối tượng và gán timezone UTC
start_date_dt = datetime.strptime(start_date, "%Y-%m-%d")

# Thêm timezone UTC vào đối tượng datetime
start_date_dt_utc = pytz.utc.localize(start_date_dt)

# Chuyển đổi thành timestamp
end_date = start_date_dt_utc.timestamp()

print(end_date)

time = 1746748800.0

dt = datetime.fromtimestamp(end_date)
print(dt)