import request from '../utils/request'


export function list(){
    return request.get("/api/todo/list")
}

export function add(data:{}){
    return request.post("/api/todo/add", data)
}

export function update(data:{}){
    return request.post("/api/todo/update", data)
}

export function finish(id:number){
    return request.get("/api/todo/finish/"+ id)
}
export function del(id:number){
    return request.delete("/api/todo/" + id)
}


export function get(id:number){
    return request.get("/api/todo/" + id)
}



