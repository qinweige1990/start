package request

import "gin-vue-admin/model"

type DriverSearch struct{
    model.Driver
    PageInfo
}