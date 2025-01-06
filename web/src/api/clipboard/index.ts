import request from '../../utils/request.ts'


export const list = (params: {}) => {
    return request({
        url: "/api/clipboard/list",
        method: "get",
        params
    })
}

export const get = () => {
    return request({
        url: "/api/clipboard",
        method: "get",
    })
}

export const add = (data: {}) => {
    return request({
        url: "/api/clipboard",
        method: "post",
        data
    })
}

export const del = (id: number) => {
    return request({
        url: "/api/clipboard/" + id,
        method: "delete",
    })
}






