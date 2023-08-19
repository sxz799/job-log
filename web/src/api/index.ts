import request from '../utils/request'



const api = {
    list() {
        return request.get("/api/todo/list")
    },
    get(id: number) {
        return request.get("/api/todo/" + id)
    },
    add(data: {}) {
        return request.post("/api/todo/", data)
    },
    del(id: number) {
        return request.delete("/api/todo/" + id)
    },
}

export default api
