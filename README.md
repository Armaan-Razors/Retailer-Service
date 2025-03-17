## **Retailer Service API (Golang + GORM)**  

### **📌 Overview**  
This is a REST API for a **retailer service** that allows:  
✅ **Retailers** to add, update, and list products.  
✅ **Customers** to place orders and view order details.  

---

### **🛠 Tech Stack**  
- **Golang** (Gin framework)  
- **GORM** (ORM for MySQL)  

---

### **🚀 Setup Instructions**  
#### **1️⃣ Install Dependencies**  
```sh
go mod tidy
```
#### **2️⃣ Set Up MySQL**  
```sql
CREATE DATABASE retailer;
USE retailer;
```
#### **3️⃣ Configure `.env` File**  
```env
DB_USER=root
DB_PASS=yourpassword
DB_HOST=127.0.0.1:3306
DB_NAME=retailer
```
#### **4️⃣ Run the Server**  
```sh
go run main.go
```
✅ **Server runs on** `http://localhost:8080`

---

### **📌 API Endpoints**  
#### **🛒 Products**  
- **Add Product** → `POST /product`  
- **Get Product by ID** → `GET /product/:id`  
- **List All Products** → `GET /products`  
- **Update Product** → `PATCH /product/:id`  

#### **📦 Orders**  
- **Place Order** → `POST /order`  
- **Get Order by ID** → `GET /order/:id`  

---
