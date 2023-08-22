import request from '../utils/request'


export const list = (params: {}) => {
    return request({
        url: "/api/todo/list",
        method: "get",
        params
    })

}

export const add = (data: {}) => {
    return request({
        url: "/api/todo/",
        method: "post",
        data
    })
}

export const update = (id: number, data: {}) => {
    return request({
        url: "/api/todo/" + id,
        method: "put",
        data
    })
}

export const del = (id: number) => {
    return request({
        url: "/api/todo/" + id,
        method: "delete",
    })
}

export const get = (id: number) => {
    return request({
        url: "/api/todo/" + id,
        method: "get",
    })
}



