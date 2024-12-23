CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,            -- ใช้ SERIAL สำหรับการเพิ่มเลข ID อัตโนมัติ
    username VARCHAR(255) NOT NULL,   -- ชื่อผู้ใช้
    email VARCHAR(255) NOT NULL UNIQUE, -- อีเมล ต้องไม่ซ้ำ
    password_hash TEXT NOT NULL,      -- รหัสผ่านที่เข้ารหัส
    first_name VARCHAR(255),          -- ชื่อ
    last_name VARCHAR(255),           -- นามสกุล
    address_id INTEGER,               -- เชื่อมโยงกับตาราง addresses
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP, -- เวลาในการสร้าง
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP, -- เวลาในการอัพเดท
    CONSTRAINT fk_address FOREIGN KEY(address_id) REFERENCES addresses(id) -- กำหนด foreign key
);