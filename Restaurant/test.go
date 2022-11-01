package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "strconv"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/gorilla/handler"
	"github.com/gorilla/mux"
)

type Order struct {
	Order_id    int64
	Cust_name   string
	Phone       string
	Address     string
	Total_price int64
	Date        string
}

type Item struct {
	Item_id int64
	Name    string
	Price   int64
	//address string
}
type orderItem struct {
	orderId int64
	itemId  int64
}

//insert

// func select_Gust(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("updategustForm.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	n, _ := strconv.Atoi(r.FormValue("SSID"))
// 	fmt.Println(n)
// 	Gust1 := oneGust(n)
// 	t.Execute(w, struct {
// 		Gust1 Gust
// 	}{Gust1})
// }

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}
func add_Order(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	db, err := sql.Open("mysql", "root:roooooot@tcp(127.0.0.1:3306)/d")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("connect to mysql")
	defer db.Close()
	insert, err := db.Query("insert into Orders values(" + r.FormValue("order_id") + ",'" + r.FormValue("cust_name") + "'," + r.FormValue("phone") + ",'" + r.FormValue("address") + "'," + r.FormValue("total_price") + ",'" + r.FormValue("Date") + "')")
	if err != nil {

		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("insert success")
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, struct {
		Mas string
	}{"added succesfuly"})
}

// func befor_add_gust(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("addgustForm.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	t.Execute(w, nil)
// }
// func add_Emplyee(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	}
// 	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hotel_db")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("connect to mysql")
// 	defer db.Close()
// 	insert, err := db.Query("call insert_into_Employee(" + r.FormValue("SSID") + ",'" + r.FormValue("Name") + "','" + r.FormValue("Gender") + "'," + r.FormValue("Salary") + ",'" + r.FormValue("City") + "','" + r.FormValue("Birthday") + "'," + r.FormValue("Phone") + ")")
// 	if err != nil {

// 		panic(err.Error())
// 	}
// 	defer insert.Close()
// 	fmt.Println("insert success")
// 	t, err := template.ParseFiles("AddemployeeForm.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	t.Execute(w, struct {
// 		Mas string
// 	}{"added succesfuly"})
// }
// func befor_add_employee(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("AddemployeeForm.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	t.Execute(w, nil)
// }
// func add_room(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	}
// 	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hotel_db")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("connect to mysql")
// 	defer db.Close()
// 	insert, err := db.Query("call insert_into_Room(" + r.FormValue("RoomNum") + ",'" + r.FormValue("RoomType") + "'," + r.FormValue("Price") + ",'" + r.FormValue("Description") + "','" + r.FormValue("Image") + "')")
// 	if err != nil {

// 		panic(err.Error())
// 	}
// 	defer insert.Close()
// 	fmt.Println("insert success")
// 	t, err := template.ParseFiles("addroomForm.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	t.Execute(w, struct {
// 		Mas string
// 	}{"added succesfuly"})
// }
// func befor_add_room(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("addroomForm.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	t.Execute(w, nil)
// }

// func add_booking(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	}
// 	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hotel_db")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("connect to mysql")
// 	defer db.Close()
// 	insert, err := db.Query("call insert_into_Booking(" + r.FormValue("bookingid") + "," + r.FormValue("GustSSID") + "," + r.FormValue("RoomNum") + ",'" + r.FormValue("inputdata") + "','" + r.FormValue("outputdata") + "')")
// 	if err != nil {

// 		panic(err.Error())
// 	}
// 	defer insert.Close()
// 	fmt.Println("insert success" + r.FormValue("bookingid") + "," + r.FormValue("GustSSID") + "," + r.FormValue("RoomNum") + ",'" + r.FormValue("inputdata") + "','" + r.FormValue("outputdata"))
// 	update1, err := db.Query("update gust set reserved ='yes' where Gust_ID=" + r.FormValue("GustSSID") + ";")
// 	if err != nil {

// 		panic(err.Error())
// 	}
// 	defer update1.Close()
// 	update2, err := db.Query("update room set state ='busy' where Room_NO=" + r.FormValue("RoomNum") + ";")
// 	if err != nil {

// 		panic(err.Error())
// 	}
// 	defer update2.Close()
// 	t, err := template.ParseFiles("addbookingForm.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	t.Execute(w, struct {
// 		Mas string
// 	}{"added succesfuly"})
// }
// func befor_add_booking(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("addbookingForm.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	Rooms := allRoomNotbook()
// 	Gusts := allGustsNotbook()
// 	t.Execute(w, struct {
// 		Gusts []Gust
// 		Rooms []Room
// 	}{Gusts, Rooms})
// }

// //select
func Show_Items(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	Items := allItems()
	Orders := allOrders()
	t.Execute(w, struct {
		Orders []Order
		Items  []Item
	}{Orders, Items})
}

// fun items

func allItems() []Item {

	db, err := sql.Open("mysql", "root:roooooot@tcp(127.0.0.1:3306)/d")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("connect to mysql")
	defer db.Close()
	sel, err := db.Query("SELECT * FROM d.Item;")
	if err != nil {
		panic(err.Error())
	}
	var Items []Item
	for sel.Next() {
		var item Item
		err = sel.Scan(&item.Item_id, &item.Name, &item.Price)

		if err != nil {
			panic(err.Error())
		}
		Items = append(Items, item)
	}

	defer sel.Close()
	return Items
}

//Show Orders
// //select
func Show_Orders(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	Orders := allOrders()
	t.Execute(w, struct {
		Orders []Order
	}{Orders})
}

// fun orders

func allOrders() []Order {

	db, err := sql.Open("mysql", "root:roooooot@tcp(127.0.0.1:3306)/d")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("connect to mysql")
	defer db.Close()
	sel, err := db.Query("SELECT * FROM d.Orders;")
	if err != nil {
		panic(err.Error())
	}
	var Orders []Order
	for sel.Next() {
		var order Order
		err = sel.Scan(&order.Order_id, &order.Cust_name, &order.Phone, &order.Address, &order.Total_price, &order.Date)

		if err != nil {
			panic(err.Error())
		}
		Orders = append(Orders, order)
	}

	defer sel.Close()
	return Orders
}

// func select_Employees(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("selectallemployee.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	Employees := allEmployee()
// 	t.Execute(w, Employees)
// }
// func select_Rooms(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("selectallroom.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	Rooms := allRooms()
// 	t.Execute(w, Rooms)
// }
// func select_Bookings(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("selectallbooking.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	Bookings := allBooking()
// 	t.Execute(w, Bookings)
// }

// //updata
// func before_updeta_Gusts(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("updategustForm.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	Gusts := allGustsNotbook()
// 	t.Execute(w, Gusts)
// }
// func updata_Gust(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	}

// 	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hotel_db")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("connect to mysql")
// 	defer db.Close()
// 	//",reserved=" + r.FormValue("reserved") +
// 	updata, err := db.Query("update gust set Name='" + r.FormValue("Name") + "', Gender='" + r.FormValue("Gender") + "',Phone_Number=" + r.FormValue("Phone") + ",  city= '" + r.FormValue("City") + "', birthday='" + r.FormValue("Birthday") + "' where Gust_ID=" + r.FormValue("SSID"))
// 	if err != nil {

// 		panic(err.Error())
// 	}
// 	defer updata.Close()
// 	fmt.Println("update success")
// 	t, err := template.ParseFiles("updategustForm.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	t.Execute(w, struct {
// 		Mas string
// 	}{"update succesfuly"})
// }

// func before_updeta_Employee(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("updetaemployee.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	Employees := allEmployee()
// 	t.Execute(w, Employees)
// }
// func updata_Employee(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	}

// 	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hotel_db")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("connect to mysql")
// 	defer db.Close()
// 	updata, err := db.Query("update employee set Name='" + r.FormValue("Name") + "' ,salary=" + r.FormValue("Salary") + ", Gender='" + r.FormValue("Gender") + "',phone=" + r.FormValue("Phone") + ",  city= '" + r.FormValue("City") + "', birthday='" + r.FormValue("Birthday") + "' where em_ID=" + r.FormValue("SSID"))
// 	if err != nil {

// 		panic(err.Error())
// 	}
// 	defer updata.Close()
// 	fmt.Println("update success")
// 	t, err := template.ParseFiles("updateemployeeForm.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	t.Execute(w, struct {
// 		Mas string
// 	}{"update succesfuly"})
// }

// func before_updeta_Rooms(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("updetaroom.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	Rooms := allRoomNotbook()
// 	t.Execute(w, Rooms)
// }
// func updata_room(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	}

// 	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hotel_db")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("connect to mysql")
// 	defer db.Close()
// 	updata, err := db.Query("UPDATE hotel_db.room SET  Room_Type ='" + r.FormValue("RoomType") + "', state ='" + r.FormValue("RoomType") + "' ,price_room = " + r.FormValue("Price") + ",Descripition ='" + r.FormValue("Description") + "', img ='" + r.FormValue("Image") + "' WHERE Room_NO " + r.FormValue("RoomNum"))
// 	if err != nil {

// 		panic(err.Error())
// 	}
// 	defer updata.Close()
// 	fmt.Println("update success")
// 	t, err := template.ParseFiles("updateroomForm.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	t.Execute(w, struct {
// 		Mas string
// 	}{"update succesfuly"})
// }

// //delete
// func before_delete_Gust(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("index.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	Items := allGustsNotbook()
// 	t.Execute(w, Items)
// }
func delete_Gust(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	db, err := sql.Open("mysql", "root:roooooot@tcp(127.0.0.1:3306)/d")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("connect to mysql")
	defer db.Close()
	updata, err := db.Query("DELETE FROM d.Orders where order_id=" + r.FormValue("itemnum"))
	if err != nil {

		panic(err.Error())
	}
	defer updata.Close()
	fmt.Println("delete success")
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, struct {
		Mas string
	}{"delete succesfuly"})
}

// func before_delete_employee(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("deleteemployee.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	Employees := allEmployee()
// 	t.Execute(w, Employees)
// }
// func delete_employee(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	}

// 	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hotel_db")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("connect to mysql")
// 	defer db.Close()
// 	updata, err := db.Query("DELETE FROM hotel_db.employee where em_ID=" + r.FormValue("SSID"))
// 	if err != nil {

// 		panic(err.Error())
// 	}
// 	defer updata.Close()
// 	fmt.Println("delete success")
// 	t, err := template.ParseFiles("deleteemployeeForm.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	t.Execute(w, struct {
// 		Mas string
// 	}{"delete succesfuly"})
// }

// func before_delete_room(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("deleteroom.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	Rooms := allRoomNotbook()
// 	t.Execute(w, Rooms)
// }
// func delete_room(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	}

// 	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hotel_db")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("connect to mysql")
// 	defer db.Close()
// 	updata, err := db.Query("DELETE FROM hotel_db.room where Room_NO=" + r.FormValue("RoomNum"))
// 	if err != nil {

// 		panic(err.Error())
// 	}
// 	defer updata.Close()
// 	fmt.Println("delete success")
// 	t, err := template.ParseFiles("deleteroomForm.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	t.Execute(w, struct {
// 		Mas string
// 	}{"delete succesfuly"})
// }
// func before_delete_booking(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("deletebooking.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	Rooms := allRoomNotbook()
// 	t.Execute(w, Rooms)
// }
// func delete_booking(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	}

// 	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hotel_db")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("connect to mysql")
// 	defer db.Close()
// 	updata, err := db.Query("DELETE FROM hotel_db.booking where Booking_ID=" + r.FormValue("BookingId"))
// 	if err != nil {

// 		panic(err.Error())
// 	}
// 	defer updata.Close()
// 	update1, err := db.Query("update gust set reserved ='no' where Gust_ID=" + r.FormValue("GustSSID") + ";")
// 	if err != nil {

// 		panic(err.Error())
// 	}
// 	defer update1.Close()
// 	update2, err := db.Query("update room set state ='empty' where Room_NO=" + r.FormValue("RoomNum") + ";")
// 	if err != nil {

// 		panic(err.Error())
// 	}
// 	defer update2.Close()
// 	fmt.Println("delete success")
// 	t, err := template.ParseFiles("deletebookingForm.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	t.Execute(w, struct {
// 		Mas string
// 	}{"delete succesfuly"})
// }
func main() {

	r := mux.NewRouter()
	fmt.Println("select seccusfully")
	// r.HandleFunc("/", index).Methods("GET", "POST")
	//insert
	r.HandleFunc("/addorder", add_Order).Methods("POST")
	// r.HandleFunc("/addroom", add_room).Methods("POST")
	// r.HandleFunc("/addEmployee", add_Emplyee).Methods("POST")
	// r.HandleFunc("/addBooking", add_booking).Methods("POST")
	// r.HandleFunc("/beforeaddbooking", befor_add_booking).Methods("GET", "POST")
	// r.HandleFunc("/beforeaddgust", befor_add_gust).Methods("GET", "POST")
	// r.HandleFunc("/beforeaddemployee", befor_add_employee).Methods("GET", "POST")
	// r.HandleFunc("/beforeaddroom", befor_add_room).Methods("GET", "POST")
	// //select
	// r.HandleFunc("/", Show_Items).Methods("GET", "POST")
	r.HandleFunc("/", Show_Items).Methods("GET", "POST")
	// r.HandleFunc("/selectoneGust", select_Gust).Methods("GET", "POST")
	// r.HandleFunc("/SelectEmployees", select_Employees).Methods("GET", "POST")
	// r.HandleFunc("/SelectRooms", select_Rooms).Methods("GET", "POST")
	// r.HandleFunc("/SelectBookings", select_Bookings).Methods("GET", "POST")
	// //updata
	// r.HandleFunc("/beforeupdateGusts", before_updeta_Gusts).Methods("GET", "POST")
	// r.HandleFunc("/beforeupdaterooms", before_updeta_Rooms).Methods("GET", "POST")
	// r.HandleFunc("/beforeupdateemployees", before_updeta_Employee).Methods("GET", "POST")
	// r.HandleFunc("/updateGusts", updata_Gust).Methods("GET", "POST")
	// r.HandleFunc("/updaterooms", updata_room).Methods("GET", "POST")
	// r.HandleFunc("/updateemployees", updata_Employee).Methods("GET", "POST")
	// //dalate
	// r.HandleFunc("/beforedeleteGusts", before_delete_Gust).Methods("GET", "POST")
	// r.HandleFunc("/beforedeleterooms", before_delete_room).Methods("GET", "POST")
	// r.HandleFunc("/beforedeleteemployees", before_delete_employee).Methods("GET", "POST")
	r.HandleFunc("/deleteGust", delete_Gust).Methods("GET", "POST")
	// r.HandleFunc("/deleteroom", delete_room).Methods("GET", "POST")
	// r.HandleFunc("/deleteemployee", delete_employee).Methods("GET", "POST")
	fs := http.FileServer(http.Dir("./css/"))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", fs))
	http.ListenAndServe(":8080", r)
}

// //gust

// func allGustsNotbook() []Gust {

// 	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hotel_db")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("connect to mysql")
// 	defer db.Close()
// 	sel, err := db.Query("SELECT * FROM hotel_db.gust where reserved='no';")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	var Gusts []Gust
// 	for sel.Next() {
// 		var gust Gust
// 		err = sel.Scan(&gust.Gust_ID, &gust.Name, &gust.Gender, &gust.Phone, &gust.City, &gust.Birthday, &gust.Reserved)

// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		Gusts = append(Gusts, gust)
// 	}

// 	defer sel.Close()
// 	return Gusts
// }

// func oneGust(id int) Gust {

// 	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hotel_db")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("connect to mysql", id)
// 	defer db.Close()
// 	sel, err := db.Query("SELECT * FROM hotel_db.gust where Gust_ID =" + strconv.Itoa(id))
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	var gust Gust
// 	for sel.Next() {

// 		err = sel.Scan(&gust.Gust_ID, &gust.Name, &gust.Gender, &gust.Phone, &gust.City, &gust.Birthday, &gust.Reserved)

// 		if err != nil {
// 			panic(err.Error())
// 		}
// 	}

// 	defer sel.Close()
// 	return gust
// }

// //employee
// func allEmployee() []Employee {

// 	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hotel_db")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("connect to mysql")
// 	defer db.Close()
// 	sel, err := db.Query("SELECT * FROM hotel_db.employee;")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	var Employees []Employee
// 	for sel.Next() {
// 		var emp Employee
// 		err = sel.Scan(&emp.Id, &emp.Hotel_code, &emp.Name, &emp.Salary, &emp.City, &emp.Gender, &emp.Birthday, &emp.Phone)

// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		Employees = append(Employees, emp)
// 	}

// 	defer sel.Close()
// 	return Employees
// }

// func oneEmployee(id int) Employee {

// 	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hotel_db")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("connect to mysql")
// 	defer db.Close()
// 	sel, err := db.Query("SELECT * FROM hotel_db.employee where em_ID =" + strconv.Itoa(id))
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	var emp Employee

// 	for sel.Next() {
// 		err = sel.Scan(&emp.Id, &emp.Hotel_code, &emp.Name, &emp.Salary, &emp.City, &emp.Gender, &emp.Birthday, &emp.Phone)

// 		if err != nil {
// 			panic(err.Error())
// 		}

// 	}

// 	defer sel.Close()
// 	return emp
// }

// //room
// func allRooms() []Room {

// 	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hotel_db")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("connect to mysql")
// 	defer db.Close()
// 	sel, err := db.Query("SELECT * FROM hotel_db.room ;")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	var Rooms []Room
// 	for sel.Next() {
// 		var room Room
// 		err = sel.Scan(&room.Room_NO, &room.Room_Type, &room.Price_Room, &room.Descripition, &room.State, &room.Img, &room.Hotel_Code)

// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		Rooms = append(Rooms, room)
// 	}

// 	defer sel.Close()
// 	return Rooms
// }

// func allRoomNotbook() []Room {

// 	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hotel_db")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("connect to mysql")
// 	defer db.Close()
// 	sel, err := db.Query("SELECT * FROM hotel_db.room where state='empty';")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	var Rooms []Room
// 	for sel.Next() {
// 		var room Room
// 		err = sel.Scan(&room.Room_NO, &room.Room_Type, &room.Price_Room, &room.Descripition, &room.State, &room.Img, &room.Hotel_Code)

// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		Rooms = append(Rooms, room)
// 	}

// 	defer sel.Close()
// 	return Rooms
// }
// func oneRoom(id int) Room {

// 	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hotel_db")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("connect to mysql")
// 	defer db.Close()
// 	sel, err := db.Query("SELECT * FROM hotel_db.room where Room_NO=" + strconv.Itoa(id))
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	var room Room
// 	for sel.Next() {

// 		err = sel.Scan(&room.Room_NO, &room.Room_Type, &room.Price_Room, &room.Descripition, &room.State, &room.Img, &room.Hotel_Code)

// 		if err != nil {
// 			panic(err.Error())
// 		}

// 	}

// 	defer sel.Close()
// 	return room
// }

// //BOOking
// func oneBooking(id int) Booking {

// 	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hotel_db")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("connect to mysql")
// 	defer db.Close()
// 	sel, err := db.Query("SELECT * FROM hotel_db.booking where Booking_ID =" + strconv.Itoa(id))
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	var bok Booking

// 	for sel.Next() {

// 		err = sel.Scan(&bok.Booking_ID, &bok.Gust_ID, &bok.Room_NO, &bok.Booking_Date, &bok.Check_In_Date, &bok.Check_Out_Date)

// 		if err != nil {
// 			panic(err.Error())
// 		}

// 	}

// 	defer sel.Close()
// 	return bok
// }

// func allBooking() []Booking {

// 	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/hotel_db")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("connect to mysql")
// 	defer db.Close()
// 	sel, err := db.Query("SELECT * FROM hotel_db.booking ;")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	var bookings []Booking
// 	for sel.Next() {
// 		var bok Booking
// 		err = sel.Scan(&bok.Booking_ID, &bok.Gust_ID, &bok.Room_NO, &bok.Booking_Date, &bok.Check_In_Date, &bok.Check_Out_Date)

// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		bookings = append(bookings, bok)

// 	}

// 	defer sel.Close()
// 	return bookings
// }
