<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="车牌">
          <el-input placeholder="搜索条件" v-model="searchInfo.license"></el-input>
        </el-form-item>
        <el-form-item label="司机">
          <el-input placeholder="搜索条件" v-model="searchInfo.name"></el-input>
        </el-form-item>
        <el-form-item label="司机审核状态">
          <el-select v-model="searchInfo.driverReview" placeholder="全选" clearable style="width: 120px">
            <el-option v-for="(item, index) in driverReviewOptions" :key="index" :label="item.label"
                       :value="item.value" :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="车辆审核状态">
          <el-select v-model="searchInfo.carReview" placeholder="全选" clearable style="width: 120px">
            <el-option v-for="(item, index) in carReviewOptions" :key="index" :label="item.label"
                       :value="item.value" :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增熟车</el-button>
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
    <el-table-column type="selection" width="45"></el-table-column>
    <el-table-column label="序号" type="index" fixed prop="scope.$index" align="center" width = "50"></el-table-column>

    <el-table-column label="车牌" prop="license" width="100" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="司机" prop="name" width="100" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="电话" prop="phone" width="100" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="运输证号" prop="transportNumber" align="center" :show-overflow-tooltip="true" width="100"></el-table-column>

    <el-table-column label="载重" prop="loading" width="100" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="品牌" prop="brand" width="100" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="出厂日期" prop="date" width="120" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="认证来源" prop="identify" width="100" align="center" :show-overflow-tooltip="true" :formatter="identifyFrom"></el-table-column>

    <el-table-column label="司机审核状态" prop="driverReview" width="120" align="center" :show-overflow-tooltip="true" :formatter="driverAuditFormat"></el-table-column>

    <el-table-column label="车辆审核状态" prop="carReview" width="120" align="center" :show-overflow-tooltip="true" :formatter="carAuditFormat"></el-table-column>


      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updateDriver(scope.row)" size="small" type="primary">变更</el-button>
          <el-popover placement="top" width="300" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteDriver(scope.row)">确定</el-button>
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
        <el-col :span="24">
          <el-form-item label="车牌" prop="license">
            <el-input v-model="formData.license" placeholder="请输入车牌" clearable :style="{width: '100%'}">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="11">
          <el-form-item label="司机姓名" prop="name">
            <el-input v-model="formData.name" placeholder="请输入司机姓名" clearable :style="{width: '100%'}">
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
          <el-form-item label="运输证号" prop="transportNumber">
            <el-input v-model="formData.transportNumber" placeholder="请输入运输证号" clearable
                      :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="15">
          <el-form-item label="品牌" prop="brand">
            <el-input v-model="formData.brand" placeholder="请输入品牌" clearable :style="{width: '100%'}">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="9">
          <el-form-item label="载重" prop="loading">
            <el-input v-model="formData.loading" placeholder="请输入载重" clearable :style="{width: '100%'}">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="出厂日期" prop="date">
            <el-date-picker v-model="formData.date"
                            :style="{width: '100%'}" placeholder="请选择出厂日期" clearable></el-date-picker>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="认证来源" prop="identify">
            <el-select v-model="formData.identify" placeholder="请选择认证来源" clearable :style="{width: '100%'}">
              <el-option v-for="(item, index) in identifyOptions" :key="myDriver" :label="item.label"
                         :value="item.value" :disabled="item.disabled"></el-option>
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="司机状态" prop="driverReview">
            <el-select v-model="formData.driverReview" placeholder="请选择司机状态" clearable
                       :style="{width: '100%'}">
              <el-option v-for="(item, index) in driverReviewOptions" :key="myDriver" :label="item.label"
                         :value="item.value" :disabled="item.disabled"></el-option>
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="车辆状态" prop="carReview">
            <el-select v-model="formData.carReview" placeholder="请选择车辆状态" clearable
                       :style="{width: '100%'}">
              <el-option v-for="(item, index) in carReviewOptions" :key="myDriver" :label="item.label"
                         :value="item.value" :disabled="item.disabled"></el-option>
            </el-select>
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
    createDriver,
    deleteDriver,
    deleteDriverByIds,
    updateDriver,
    findDriver,
    getDriverList
} from "@/api/driver";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "Driver",
  mixins: [infoList],
  data() {
    return {
      listApi: getDriverList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        license: undefined,
        name: undefined,
        phone: undefined,
        transportNumber: undefined,
        brand: undefined,
        loading: undefined,
        date: null,
        identify: undefined,
        driverReview: undefined,
        carReview: undefined,
      },
      rules: {
        license: [{
          required: true,
          message: '请输入车牌',
          trigger: 'blur'
        }],
        name: [{
          required: true,
          message: '请输入司机姓名',
          trigger: 'blur'
        }],
        phone: [{
          required: true,
          message: '请输入电话',
          trigger: 'blur'
        }],
        transportNumber: [{
          required: true,
          message: '请输入运输证号',
          trigger: 'blur'
        }],
        brand: [],
        loading: [],
        date: [],
        identify: [{
          required: true,
          message: '请选择认证来源',
          trigger: 'change'
        }],
        driverReview: [{
          required: true,
          message: '请选择司机状态',
          trigger: 'change'
        }],
        carReview: [{
          required: true,
          message: '请选择车辆状态',
          trigger: 'change'
        }],
      },
      identifyOptions: [ {
        "label": "公司认证",
        "value": 1
      }, {
        "label": "ETC认证",
        "value": 2
      }],
      driverReviewOptions: [ {
        "label": "审核中",
        "value": 1
      }, {
        "label": "审核通过",
        "value": 2
      }, {
        "label": "审核未通过",
        "value": 3
      }],
      carReviewOptions: [ {
        "label": "审核中",
        "value": 1
      }, {
        "label": "审核通过",
        "value": 2
      }, {
        "label": "审核未通过",
        "value": 3
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
    driverAuditFormat(row) {
      switch (row.driverReview) {
        case 1: return "审核中"
        case 2: return "审核通过"
        case 3: return "审核不通过"
      }
    },
    carAuditFormat(row) {
      switch (row.carReview) {
        case 1: return "审核中"
        case 2: return "审核通过"
        case 3: return "审核不通过"
      }
    },
    identifyFrom(row) {
      switch(row.identify) {
        case 1: return "公司认证"
        case 2: return "ETC认证"
      }
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
        const res = await deleteDriverByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateDriver(row) {
      const res = await findDriver({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.redriver;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {

          license:null,
          name:null,
          phone:null,
          transportNumber:null,
          loading:null,
          brand:null,
          date:null,
          identify:null,
          driverReview:null,
          carReview:null,
          driverLicenseImg:null,
          carLicenseImg:null,
      };
    },
    async deleteDriver(row) {
      this.visible = false;
      const res = await deleteDriver({ ID: row.ID });
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
          res = await createDriver(this.formData);
          break;
        case "update":
          res = await updateDriver(this.formData);
          break;
        default:
          res = await createDriver(this.formData);
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
  created() {
    this.getTableData();
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
};
</script>

<style>
</style>
