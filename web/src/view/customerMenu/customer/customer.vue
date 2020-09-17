<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="名称">
          <el-input placeholder="搜索条件" v-model="searchInfo.name"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增客户</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>

    <el-table-column label="名称" prop="name" width="120"></el-table-column>

    <el-table-column label="类别" prop="type" width="120" :formatter="typeFormat"></el-table-column>

    <el-table-column label="所属公司" prop="company" width="120"></el-table-column>

    <el-table-column label="联系人" prop="contactMan" width="120"></el-table-column>

    <el-table-column label="联系电话" prop="phone" width="120"></el-table-column>

    <el-table-column label="地址" prop="address" width="120"></el-table-column>

    <el-table-column label="备注" prop="extra" width="120"></el-table-column>

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updateCustomer(scope.row)" size="small" type="primary">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteCustomer(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px">
        <el-form-item label="名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入名称" clearable :style="{width: '100%'}"></el-input>
        </el-form-item>
        <el-form-item label="类别" prop="type">
          <el-select v-model="formData.type" placeholder="请输入类别" clearable :style="{width: '100%'}">
            <el-option v-for="(item, index) in typeOptions" :key="index" :label="item.label"
                       :value="item.value" :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="所属公司" prop="company">
          <el-input v-model="formData.company" placeholder="请输入所属公司" clearable :style="{width: '100%'}">
          </el-input>
        </el-form-item>
        <el-form-item label="联系人" prop="contactMan">
          <el-input v-model="formData.contactMan" placeholder="请输入联系人" clearable :style="{width: '100%'}">
          </el-input>
        </el-form-item>
        <el-form-item label="联系电话" prop="phone">
          <el-input v-model="formData.phone" placeholder="请输入联系电话" clearable :style="{width: '100%'}">
          </el-input>
        </el-form-item>
        <el-form-item label="地址" prop="address">
          <el-input v-model="formData.address" placeholder="请输入地址" clearable :style="{width: '100%'}">
          </el-input>
        </el-form-item>
        <el-form-item label="备注" prop="extra">
          <el-input v-model="formData.extra" type="textarea" placeholder="请输入备注"
                    :autosize="{minRows: 4, maxRows: 4}" :style="{width: '100%'}"></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createCustomer,
    deleteCustomer,
    deleteCustomerByIds,
    updateCustomer,
    findCustomer,
    getCustomerList
} from "@/api/customerMenu";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "Customer",
  mixins: [infoList],
  data() {
    return {
      listApi: getCustomerList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        name: undefined,
        type: undefined,
        company: undefined,
        contactMan: undefined,
        phone: undefined,
        address: undefined,
        extra: undefined,
      },
      rules: {
        name: [{
          required: true,
          message: '请输入名称',
          trigger: 'blur'
        }],
        type: [{
          required: true,
          message: '请输入类别',
          trigger: 'change'
        }],
        company: [{
          required: true,
          message: '请输入所属公司',
          trigger: 'blur'
        }],
        contactMan: [{
          required: true,
          message: '请输入联系人',
          trigger: 'blur'
        }],
        phone: [{
          required: true,
          message: '请输入联系电话',
          trigger: 'blur'
        }],
        address: [{
          required: true,
          message: '请输入地址',
          trigger: 'blur'
        }],
      },
      typeOptions: [{
        "label": "发货客户",
        "value": 1
      }, {
        "label": "收货客户",
        "value": 2
      }],
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
    typeFormat(row) {
      switch (row.type) {
        case 1: return "发货客户"
        case 2: return "收货客户"
      }
    },
    onOpen() {},
    onClose() {
      this.$refs['elForm'].resetFields()
    },
    close() {
      this.$emit('update:visible', false)
    },
    handelConfirm() {
      this.$refs['elForm'].validate(valid => {
        if (!valid) return
        this.close()
      })
    },
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      async onDelete() {
        const ids = []
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteCustomerByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateCustomer(row) {
      const res = await findCustomer({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.recustomer;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {

          name:null,
          type:null,
          company:null,
          contactMan:null,
          phone:null,
          address:null,
          extra:null,
      };
    },
    async deleteCustomer(row) {
      this.visible = false;
      const res = await deleteCustomer({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createCustomer(this.formData);
          break;
        case "update":
          res = await updateCustomer(this.formData);
          break;
        default:
          res = await createCustomer(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();}
};
</script>

<style>
</style>
