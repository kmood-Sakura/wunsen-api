# แนวทางการ run wunsen api
1. ทำการ run file run.sh สำหรับ Linux และ run.bat สำหรับ windows

# แนวทางการทดสอบระบบ API
1. ลองใช้ postman request ไปที่ 0.0.0.0:4000/health เพื่อ check health
2. เมื่อ health ปกติให้ทำการยิง api 0.0.0.0:4000/api/bmi โดยใช้ method POST โดยที่ body เป็นแบบ json โดย field ที่ require คือ {"sex": "male|female", "weight": (float), "height" (float)} โดยที่ weight และ height มีหน่วยเป็น kg และ cm ตามลำดับ
3. รับผล
