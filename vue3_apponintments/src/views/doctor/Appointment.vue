<script setup>
import { ref, reactive ,onMounted} from "vue";
import { Search } from "@element-plus/icons-vue";
import request from "@/utils/request";
import { ElDrawer, ElMessage, ElMessageBox } from "element-plus";
import { Gtepatient,Setappointment} from "@/api/doctor.js";
import { Check, Delete, Edit, Message, Star } from "@element-plus/icons-vue";
import { Plus } from "@element-plus/icons-vue";
import { doctorStore } from '@/stores/doctor.js'
const doctorData = doctorStore();


//分页条数据模型
const pageNum = ref(1); //当前页
const total = ref(20); //总条数
const pageSize = ref(5); //每页条数

//当每页条数发生了变化，调用此函数
const onSizeChange = (size) => {
  pageSize.value = size;
  search()
};
//当前页码发生变化，调用此函数
const onCurrentChange = (num) => {
  pageNum.value = num;
  search()
};

const data = reactive({
    Data:"",
  Name: "",
 Status:""
});
//医生列表数据模型
const patientList = ref([]);

const search = async () => {
  let params = {
    doctorId : doctorData.Doctor.DoctorId,
    Date: data.Data,
    Status : data.Status,
    pageNum: pageNum.value,
    pageSize: pageSize.value,
  };

  let result = await Gtepatient(params);
  //渲染列表数据

  patientList.value = result.message;

  total.value = result.Total;
};
search()

//按姓名进行查询
const searchN = async () => {  
  let params = {
    doctorId : doctorData.Doctor.DoctorId,
    Name : data.Name,
    pageNum: pageNum.value,
    pageSize: pageSize.value,
  };

  let result = await Gtepatient(params);
  //渲染列表数据

  patientList.value = result.message;

  
};

const endappointment = (Data) => {
  ElMessageBox.confirm("确认后数据后无法恢复", "完成确认", {
    type: "warning",
  }).then(async () => {
  console.log(Data)
    let AppointmentId = {
        appointment_id : Data,
        status: "已完成"
    }
    let result = await Setappointment(AppointmentId);
    if (result.code === 200) {
      search(); // 重新获取数据
      ElMessage.success(result.message);
    } else {
      ElMessage.error(result.message);
    }
  });
};
</script>





<template>
  <div>
    <div class="card" style="margin-bottom: 10px">


       

      <div class="block">
            <el-date-picker
              v-model="data.Data"
              type="date"
              placeholder="选择日期"
              format="YYYY/MM/DD"
              value-format="YYYY-MM-DD"/>

 <el-select v-model="data.Status" placeholder="请选择状态" style="width: 200px; margin-right: 10px">
        <el-option :key="null"  :value="null"></el-option>
        <el-option :key="1" :label="'已签到'" :value="'已签到'"></el-option>
        <el-option :key="2" :label="'已预约'" :value="'已预约'"></el-option>
        <el-option :key="2" :label="'已完成'" :value="'已完成'"></el-option>
        <el-option :key="2" :label="'已取消'" :value="'已取消'"></el-option>
      </el-select>

              <el-button type="primary" tyle="margin-left:10px" @click="search">查询</el-button>
          </div>

      <!-- <el-button type="primary" tyle="margin-left:10px" @click="searchS">查询</el-button> -->
          
      <el-input
        style="width: 200px; margin-right: 10px"
        v-model="data.Name"
        placeholder="请输入用户姓名"
        :prefix-icon="Search"
      ></el-input>
      <el-button type="primary" tyle="margin-left:10px" @click="searchN">查询</el-button>


      
          
    </div>

  
    <div class="crad">
      <!--表单-->
      <div>
        <el-table
          :data="patientList"
          :default-sort="{ prop: 'id', order: 'descending' }"
          style="width: 100%"
        >
          <!--          <el-table-column type="selection" width="55" align="center"></el-table-column>-->
          <el-table-column
            prop="AppointmentId"
            label="预约编号"
            show-overflow-tooltip
          ></el-table-column>

           <el-table-column
            prop="Name"
            label="患者姓名"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            prop="Identity"
            label="身份证号"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            prop="Phone"
            label="电话号码"
            show-overflow-tooltip
          ></el-table-column>

          <el-table-column
            prop="Data"
            label="预约日期"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            prop="SignTime"
            label="签到时间"
            show-overflow-tooltip
          ></el-table-column>
         
          <el-table-column
            prop="Status"
            label="状态"
            show-overflow-tooltip
          ></el-table-column>
         
          <el-table-column
            prop="Img"
            label="头像"
            show-overflow-tooltip
          ><template #default="{ row }">
        <el-image
          :src="row.Img"
          style="width: 50px; height: 50px; border-radius: 50%; object-fit: cover;"
        ></el-image>
      </template></el-table-column>
          

          <el-table-column label="操作" align="center" width="160">
            <template #default="{row}">
             
              <el-button type="primary" tyle="margin-left:10px" @click="endappointment(row.AppointmentId)">完成</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
    <!-- 分页条 -->
    <el-pagination
      v-model:current-page="pageNum"
      v-model:page-size="pageSize"
      :page-sizes="[3, 5, 10, 15]"
      layout="jumper, total, sizes, prev, pager, next"
      background
      :total="total"
      @size-change="onSizeChange"
      @current-change="onCurrentChange"
      style="margin-top: 20px; justify-content: flex-end"
    />


   
  </div>
</template>