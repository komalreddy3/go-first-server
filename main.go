package main
import ("fmt"
"log"
"net/http"
)
func formHandler(w http.ResponseWriter, r *http.Request){
	if err:= r.ParseForm(); err!=nil{
     fmt.Fprintf(w,"Parse form error %v",err);
	 return
	}
	fmt.Fprintf(w,"Post request successful")
	name:= r.FormValue("name")
	address:= r.FormValue("address")
	fmt.Fprintf(w,"Name is %s ",name)
	fmt.Fprintf(w,"address is %s ",address)
}
func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path!="/hello"{
		http.Error(w,"404 not found",http.StatusNotFound)
	}
	if r.Method !="GET"{
		http.Error(w,"method is not supported",http.StatusNotFound)
	}
	fmt.Fprintf(w,"hello")
}
func main(){
	fileServer:= http.FileServer(http.Dir("./static")) // telling go lang to check the directory for rendering
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Println("Server started at port 8080");
	if err:=http.ListenAndServe("localhost:8080",nil); err!=nil{
		log.Fatal(err);
	}
}