<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" label-width="100">

        <el-form-item label="运单审核状态">
          <el-select v-model="searchInfo.orderAuditStatus" placeholder="全选" clearable style="width: 120px">
            <el-option v-for="(item, index) in orderAuditOptions" :key="index" :label="item.label"
                       :value="item.value" :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="开票状态">
          <el-select v-model="searchInfo.invoiceStatus" placeholder="全选" clearable style="width: 120px">
            <el-option v-for="(item, index) in invoiceStatusOptions" :key="index" :label="item.label"
                       :value="item.value" :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="目的地">
          <el-input placeholder="搜索条件" v-model="searchInfo.end"></el-input>
        </el-form-item>
        <el-form-item label="收货客户">
          <el-input placeholder="搜索条件" v-model="searchInfo.endCustomer"></el-input>
        </el-form-item>

        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增</el-button>
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
      id="orderTable"
      height="700"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="40"></el-table-column>
    <el-table-column label="序号" type="index" fixed prop="scope.$index" align="center" width = "50"></el-table-column>

    <el-table-column label="运输单号" prop="transportId" center width="100"></el-table-column>

    <el-table-column label="派车时间" prop="departureTime" align="center" width="120" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="起始地" prop="start" width="120" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="目的地" prop="end" width="120" align="center" :show-overflow-tooltip="true"></el-table-column>

      <el-table-column label="总运费" prop="totalFare" width="80" align="center" :show-overflow-tooltip="true"></el-table-column>

      <el-table-column label="预付运费" prop="preFare" width="80" align="center" :show-overflow-tooltip="true"></el-table-column>

      <el-table-column label="已支付" prop="paidFare" width="80" align="center" :show-overflow-tooltip="true"></el-table-column>

      <el-table-column label="运费尾款" prop="dueFare" width="80" align="center" :show-overflow-tooltip="true"></el-table-column>

      <el-table-column label="运单状态" prop="transportStatus" width="80" align="center" :show-overflow-tooltip="true" :formatter="transportStatusFormat">
      </el-table-column>

      <el-table-column label="运费审核状态" prop="fareAuditStatus" width="110" align="center" :show-overflow-tooltip="true" :formatter="fareAuditFormat">
      </el-table-column>

      <el-table-column label="运单审核状态" prop="orderAuditStatus" width="110" align="center" :show-overflow-tooltip="true" :formatter="orderAuditFormat">
      </el-table-column>

      <el-table-column label="运费支付状态" prop="farePayStatus" width="110" align="center" :show-overflow-tooltip="true" :formatter="farePayFormat">
      </el-table-column>

      <el-table-column label="尾款支付状态" prop="dueFareStatus" width="110" align="center" :show-overflow-tooltip="true" :formatter="dueFareFormat">
      </el-table-column>

      <el-table-column label="开票状态" prop="invoiceStatus" width="80" align="center" :show-overflow-tooltip="true" :formatter="invoiceStatusFormat">
      </el-table-column>

    <el-table-column label="货物" prop="good" width="80" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="司机" prop="driver" width="80" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="司机电话" prop="driverPhone" width="100" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="车牌号" prop="carNumber" width="100" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="装车数据" prop="startWeight" width="80" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="卸车数据" prop="endWeight" width="80" align="center" :show-overflow-tooltip="true"></el-table-column>



    <el-table-column label="发货客户" prop="startCustomer" width="120" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="发货电话" prop="startPhone" width="120" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="收货客户" prop="endCustomer" width="120" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="收货电话" prop="endPhone" width="120" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="服务站名称" prop="station" width="100" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="服务站电话" prop="stationPhone" width="100" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="计划单号" prop="planId" width="100" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="订单编号" prop="orderId" width="100" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="承运方" prop="agent" width="100" align="center" :show-overflow-tooltip="true"></el-table-column>



    <el-table-column label="创造人" prop="createBy" width="80" align="center" :show-overflow-tooltip="true"></el-table-column>

    <el-table-column label="操作" width="120" align="center">
      <template slot-scope="scope">
        <el-button @click="updateOrder(scope.row)" size="small" type="primary">变更</el-button>
        <el-popover placement="top" width="200" v-model="scope.row.visible">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin: 0">
            <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
            <el-button type="primary" size="mini" @click="deleteOrder(scope.row)">确定</el-button>
          </div>
          <el-button type="danger" size="mini" slot="reference">删除</el-button>
        </el-popover>
      </template>
    </el-table-column>
    </el-table>
    <!-- 导出按钮 -->
    <div class="toexcel">
      <el-button  @click="exportExcel" type="primary">导出</el-button>
    </div>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="创建运单">
      <el-form ref="elForm" :model="createOrder" :rules="rules" size="large" label-width="100px">
        <el-col :span="8">
          <el-form-item label="派车时间" prop="departureTime">
            <el-date-picker v-model="createOrder.departureTime" type="datetime"
                            :style="{width: '100%'}" placeholder="选择派车时间" clearable></el-date-picker>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="订单单号" prop="orderId">
            <el-input v-model="createOrder.orderId" placeholder="订单单号" clearable :style="{width: '100%'}">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="承运方" prop="agent">
            <el-select v-model="createOrder.agent" placeholder="承运方" clearable :style="{width: '100%'}">
              <el-option v-for="(item, index) in agentOptions" :key="index" :label="item.label"
                         :value="item.value" :disabled="item.disabled"></el-option>
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label="发货客户" prop="startCustomer">
            <el-input v-model="createOrder.startCustomer" placeholder="请输入发货客户" clearable
                      :style="{width: '100%'} "></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label="联系电话" prop="startPhone">
            <el-input v-model="createOrder.startPhone" placeholder="请输入联系电话" clearable
                      :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="发货地址" prop="start">
            <el-input v-model="createOrder.start" placeholder="请输入发货地址" clearable :style="{width: '100%'}">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label="收货客户" prop="endCustomer">
            <el-input v-model="createOrder.endCustomer" placeholder="请输入收货客户" clearable
                      :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label="联系电话" prop="endPhone">
            <el-input v-model="createOrder.endPhone" placeholder="请输入联电话" clearable :style="{width: '100%'}">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="收货地址" prop="end">
            <el-input v-model="createOrder.end" placeholder="请输入收货地址" clearable :style="{width: '100%'}">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label="货物名称" prop="good">
            <el-input v-model="createOrder.good" placeholder="请输入货物名称" clearable :style="{width: '100%'}">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="7">
          <el-form-item label="重量(吨）" prop="startWeight">
            <el-input-number v-model="createOrder.startWeight" placeholder="请输入重量(吨)" clearable
                      :style="{width: '100%'}"></el-input-number>
          </el-form-item>
        </el-col>
        <el-col :span="11">
          <el-form-item label="每吨运费" prop="totalFare">
            <el-input v-model="createOrder.totalFare" placeholder="请输入每吨运费" clearable :style="{width: '100%'}">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="备注" prop="extra">
            <el-input v-model="createOrder.extra" type="textarea" placeholder="请输入备注"
                      :autosize="{minRows: 2, maxRows: 4}" :style="{width: '100%'}"></el-input>
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
    createOrder,
    deleteOrder,
    deleteOrderByIds,
    updateOrder,
    findOrder,
    getOrderList
} from "@/api/order";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";
import FileSaver from "file-saver";
import XLSX from "xlsx";

export default {
  name: "Order",
  mixins: [infoList],
  data() {
    return {
      listApi: getOrderList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
            intOptions:[],
          formData: {
        transportId:null,departureTime:null,start:null,end:null,good:null,driver:null,driverPhone:null,carNumber:null,endWeight:null,startWeight:null,totalFare:null,preFare:null,paidFare:null,dueFare:null,startCustomer:null,startPhone:null,endCustomer:null,endPhone:null,station:null,stationPhone:null,planId:null,orderId:null,agent:null,transportStatus:null,fareAuditStatus:null,orderAuditStatus:null,farePayStatus:null,dueFareStatus:null,invoiceStatus:null,createBy:null,
      },
      createOrder: {
        departureTime: null,
        orderId: undefined,
        agent: undefined,
        startCustomer: undefined,
        startPhone: undefined,
        start: undefined,
        endCustomer: undefined,
        endPhone: undefined,
        end: undefined,
        good: undefined,
        startWeight: undefined,
        totalFare: undefined,
        extra: undefined,
      },
      rules: {
        departureTime: [{
          required: true,
          message: '选择派车时间',
          trigger: 'change'
        }],
        orderId: [{
          required: true,
          message: '订单单号',
          trigger: 'blur'
        }],
        agent: [{
          required: true,
          message: '承运方',
          trigger: 'change'
        }],
        startCustomer: [{
          required: true,
          message: '请输入发货客户',
          trigger: 'blur'
        }],
        startPhone: [{
          required: true,
          message: '请输入联系电话',
          trigger: 'blur'
        }],
        start: [{
          required: true,
          message: '请输入发货地址',
          trigger: 'blur'
        }],
        endCustomer: [{
          required: true,
          message: '请输入收货客户',
          trigger: 'blur'
        }],
        endPhone: [{
          required: true,
          message: '请输入联系人',
          trigger: 'blur'
        }],
        end: [{
          required: true,
          message: '请输入收货地址',
          trigger: 'blur'
        }],
        good: [{
          required: true,
          message: '请输入货物名称',
          trigger: 'blur'
        }],
        startWeight: [{
          type: 'number',
          required: true,
          message: '请输入重量(吨)',
          trigger: 'blur'
        }],
        totalFare: [{
          required: true,
          message: '请输入每吨运费',
          trigger: 'blur'
        }],
        extra: [],
      },
      agentOptions: [{
        "label": "大牛",
        "value": "大牛"
      }, {
        "label": "自运",
        "value": "自运"
      }],
      orderAuditOptions :[{
        "label": "全选",
        "value": ""
      },{
        "label": "待审核",
        "value": "0"
      }, {
        "label": "审核通过",
        "value": "1"
      },{
        "label": "审核不通过",
        "value": "2"
      }],
      invoiceStatusOptions :[{
        "label": "全选",
        "value": ""
      },{
        "label": "未开票",
        "value": "0"
      }, {
        "label": "已开票",
        "value": "1"
      },{
        "label": "审核中",
        "value": "2"
      },{
        "label": "已驳回",
        "value": "3"
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
    // 导出表格所用
    exportExcel() {
      // 从表格生成workbook
      let fix = document.querySelector('.el-table__fixed');
        let wb = XLSX.utils.table_to_book(document.querySelector('#orderTable').removeChild(fix));
        document.querySelector('#orderTable').appendChild(fix);

      let wbout = XLSX.write(wb, {
        bookType: "xlsx",
        bookSST: true,
        type: "array"
      });
      try {
        // 下载
        let b = new Blob([wbout], { type: "application/octet-stream" });
        FileSaver.saveAs(b, "运单报表.xlsx");
      } catch (e) {
        if (typeof console !== "undefined") {
          console.log(e, wbout);
        }
      }
      return wbout;
    },

      orderAuditFormat(row) {
        switch (row.orderAuditStatus) {
          case 0: return "待审核"
          case 1: return "审核通过"
          case 2: return "审核不通过"
        }
      },
    fareAuditFormat(row) {
      switch (row.orderAuditStatus) {
        case 0: return "待审核"
        case 1: return "审核通过"
        case 2: return "已驳回"
      }
    },
    invoiceStatusFormat(row) {
      switch (row.orderAuditStatus) {
        case 0: return "未开票"
        case 1: return "开票中"
        case 2: return "已开票"
        case 3: return "审核中"
        case 4: return "已驳回"
      }
    },
    dueFareFormat(row) {
      switch (row.orderAuditStatus) {
        case 0: return "未支付"
        case 1: return "已支付"
      }
    },
    farePayFormat(row) {
      switch (row.orderAuditStatus) {
        case 0: return "未支付"
        case 1: return "部分支付"
        case 2: return "全部支付"
      }
    },
    transportStatusFormat(row) {
      switch (row.orderAuditStatus) {
        case 0: return "待装货"
        case 1: return "待送达"
        case 2: return "待签收"
        case 3: return "已签收"
        case 4: return "已删除"
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
        const res = await deleteOrderByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateOrder(row) {
      const res = await findOrder({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reorder;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {

          transportId:null,
          departureTime:null,
          start:null,
          end:null,
          good:null,
          driver:null,
          driverPhone:null,
          carNumber:null,
          startWeight:null,
          endWeight:null,
          totalFare:null,
          preFare:null,
          paidFare:null,
          dueFare:null,
          startCustomer:null,
          startPhone:null,
          endCustomer:null,
          endPhone:null,
          station:null,
          stationPhone:null,
          planId:null,
          orderId:null,
          agent:null,
          transportStatus:null,
          fareAuditStatus:null,
          orderAuditStatus:null,
          farePayStatus:null,
          dueFareStatus:null,
          invoiceStatus:null,
          createBy:null,
      };
    },
    async deleteOrder(row) {
      this.visible = false;
      const res = await deleteOrder({ ID: row.ID });
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
          res = await createOrder(this.createOrder);
          break;
        case "update":
          res = await updateOrder(this.formData);
          break;
        default:
          res = await createOrder(this.formData);
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
  },

  async created() {
    await this.getTableData();await this.getDict("int");await this.getDict("int");await this.getDict("int");await this.getDict("int");await this.getDict("int");await this.getDict("int")}
};
</script>

<style scoped lang="css">

</style>
