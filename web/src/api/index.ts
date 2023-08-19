import request from '../utils/request'


export function list(){
    return request.get("/api/todo/list")
}
export function get(id:number){
    return request.get("/api/todo/" + id)
}
export function add(data:{}){
    return request.post("/api/todo/add", data)
}
export function del(id:number){
    return request.delete("/api/todo/" + id)
}



