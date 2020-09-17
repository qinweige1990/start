package request

import "gin-vue-admin/model"

type InformationSearch struct{
    model.Information
    PageInfo
}