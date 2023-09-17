import request from '../../utils/request.ts'


export const list = (params: {}) => {
    return request({
        url: "/api/clipboard/list",
        method: "get",
        params
    })
}

export const update = ( data: {}) => {
    return request({
        url: "/api/clipboard/" ,
        method: "put",
        data
    })
}

export const del = (id: number) => {
    return request({
        url: "/api/clipboard/" + id,
        method: "delete",
    })
}

export const get = () => {
    return request({
        url: "/api/clipboard/",
        method: "get",
    })
}



