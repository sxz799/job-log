import request from '../utils/request'


export function list(){
    return request.get("/api/todo/list")
}

export function add(data:{}){
    return request.post("/api/todo/", data)
}

export function update(id:number,data:{}){
    return request.put("/api/todo/"+id, data)
}

export function del(id:number){
    return request.delete("/api/todo/" + id)
}


export function get(id:number){
    return request.get("/api/todo/" + id)
}



