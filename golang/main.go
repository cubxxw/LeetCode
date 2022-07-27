package main
// GO语言中的哈希表
import "fmt"

func main() {
   var countryCapitalmap map[string]string
   /* 创建集合 */
   countryCapitalmap = make(map[string]string)

   /* map 插入 key-value 对，各个国家对应的首都 */
   countryCapitalmap["France"] = "Paris"
   countryCapitalmap["Italy"] = "Rome"
   countryCapitalmap["Japan"] = "Tokyo"
   countryCapitalmap["India"] = "New Delhi"

   /* 使用 key 输出 map 值 */
   for country := range countryCapitalmap {
      fmt.Println("Capital of",country,"is",countryCapitalmap[country])
   }

   /* 查看元素在集合中是否存在 */
   captial, ok := countryCapitalmap["United States"]
   /* 如果 ok 是 true, 则存在，否则不存在 */
   if(ok){
      fmt.Println("Capital of United States is", captial)  
   }else {
      fmt.Println("Capital of United States is not present") 
   }
}