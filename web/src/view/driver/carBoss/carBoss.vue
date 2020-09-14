<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="姓名">
          <el-input placeholder="搜索条件" v-model="searchInfo.name"></el-input>
        </el-form-item>
        <el-form-item label="状态">
          <el-input placeholder="搜索条件" v-model="searchInfo.status"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增车老板</el-button>
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

    <el-table-column label="姓名" prop="name" width="120"></el-table-column>

    <el-table-column label="电话" prop="phone" width="120"></el-table-column>

    <el-table-column label="身份证" prop="identity" width="120"></el-table-column>

    <el-table-column label="绑定银行" prop="bank" width="120"></el-table-column>

    <el-table-column label="银行卡号" prop="bankNumber" width="120"></el-table-column>

    <el-table-column label="状态" prop="status" width="120"></el-table-column>

    <el-table-column label="照片" prop="img" width="120"></el-table-column>

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updateCarBoss(scope.row)" size="small" type="primary">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteCarBoss(scope.row)">确定</el-button>
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
      <el-form ref="driverForm" :model="formData" :rules="rules" size="medium" label-width="100px">
        <el-col :span="11">
          <el-form-item label="姓名" prop="name">
            <el-input v-model="formData.name" placeholder="请输入姓名" clearable :style="{width: '100%'}">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="电话" prop="phone">
            <el-input v-model="formData.phone" placeholder="请输入电话" clearable :style="{width: '100%'}">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="身份证号" prop="identity">
            <el-input v-model="formData.identity" placeholder="请输入身份证号" clearable :style="{width: '100%'}">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="绑定银行" prop="bank">
            <el-input v-model="formData.bank" placeholder="请输入绑定银行" clearable :style="{width: '100%'}">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="银行卡号" prop="bankNumber">
            <el-input v-model="formData.bankNumber" placeholder="请选择银行卡号" clearable
                      :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="状态" prop="status">
            <el-select v-model="formData.status" placeholder="请选择状态" clearable :style="{width: '100%'}">
              <el-option v-for="(item, index) in statusOptions" :key="index" :label="item.label"
                         :value="item.value" :disabled="item.disabled"></el-option>
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="证件照片" prop="img">
            <el-upload ref="img" :file-list="imgfileList" :action="imgAction"
                       :before-upload="imgBeforeUpload" list-type="picture-card" accept="image/*" name="img">
              <i class="el-icon-plus"></i>
            </el-upload>
          </el-form-item>
        </el-col>
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
    createCarBoss,
    deleteCarBoss,
    deleteCarBossByIds,
    updateCarBoss,
    findCarBoss,
    getCarBossList
} from "@/api/carBoss";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "CarBoss",
  mixins: [infoList],
  data() {
    return {
      listApi: getCarBossList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
        name:null,phone:null,identity:null,bank:null,bankNumber:null,status:null,img:null,
      },
      formData: {
        name: undefined,
        phone: undefined,
        identity: undefined,
        bank: undefined,
        bankNumber: undefined,
        status: undefined,
        img: null,
      },
      rules: {
        name: [{
          required: true,
          message: '请输入姓名',
          trigger: 'blur'
        }],
        phone: [{
          required: true,
          message: '请输入电话',
          trigger: 'blur'
        }],
        identity: [{
          required: true,
          message: '请输入身份证号',
          trigger: 'blur'
        }],
        bank: [{
          required: true,
          message: '请输入绑定银行',
          trigger: 'blur'
        }],
        bankNumber: [],
        status: [{
          required: true,
          message: '请选择状态',
          trigger: 'change'
        }],
      },
      imgAction: 'https://jsonplaceholder.typicode.com/posts/',
      imgfileList: [],
      statusOptions: [{
        "label": "审核中",
        "value": 0
      }, {
        "label": "审核通过",
        "value": 1
      }, {
        "label": "审核未通过",
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
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
    onOpen() {},
    onClose() {
      this.$refs['driverForm'].resetFields()
    },
    close() {
      this.$emit('update:visible', false)
    },
    handelConfirm() {
      this.$refs['driverForm'].validate(valid => {
        if (!valid) return
        this.close()
      })
    },
    imgBeforeUpload(file) {
      let isRightSize = file.size / 1024 / 1024 < 10
      if (!isRightSize) {
        this.$message.error('文件大小超过 10MB')
      }
      let isAccept = new RegExp('image/*').test(file.type)
      if (!isAccept) {
        this.$message.error('应该选择image/*类型的文件')
      }
      return isRightSize && isAccept
    },

      async onDelete() {
        const ids = []
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteCarBossByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateCarBoss(row) {
      const res = await findCarBoss({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.recarBoss;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {

          name:null,
          phone:null,
          identity:null,
          bank:null,
          bankNumber:null,
          status:null,
          img:null,
      };
    },
    async deleteCarBoss(row) {
      this.visible = false;
      const res = await deleteCarBoss({ ID: row.ID });
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
          res = await createCarBoss(this.formData);
          break;
        case "update":
          res = await updateCarBoss(this.formData);
          break;
        default:
          res = await createCarBoss(this.formData);
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
