package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "🐳 राहुल भाई की पहली लाइव GitOps ऐप एकदम फर्स्ट क्लास चल रही है! 🚀")
}

func main() {
	http.HandleFunc("/", helloHandler)
	
	// यहाँ हमने पोर्ट बदलकर 9090 कर दिया है
	fmt.Println("Server starting on port 9090...")
	http.ListenAndServe(":9090", nil)
}