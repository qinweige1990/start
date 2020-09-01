package request

import "gin-vue-admin/model"

type CargoOrderSearch struct{
    model.CargoOrder
    PageInfo
}