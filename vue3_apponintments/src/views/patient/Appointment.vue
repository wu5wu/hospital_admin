<script setup>
import { ref, reactive, onMounted } from "vue";

import { ElDrawer, ElMessage, ElMessageBox } from "element-plus";

import { getdoctorschedule, getDepartmentList, getAppointmentDate, getAppointmentAdd } from "@/api/appointment.js";

import dayjs from 'dayjs'; // 日期库
import 'dayjs/locale/zh-cn'; // 日期库的语言包

import { useUserStore } from '@/stores/t_user.js'
const userData = useUserStore();

//分页条数据模型
const pageNum = ref(1); //当前页
const total = ref(20); //总条数
const pageSize = ref(5); //每页条数

//当每页条数发生了变化，调用此函数
const onSizeChange = (size) => {
  pageSize.value = size;
  getDoctorsChedule();
};
//当前页码发生变化，调用此函数
const onCurrentChange = (num) => {
  pageNum.value = num;
  getDoctorsChedule();
};


// reactive 数据对象
const data = reactive({
  department_id: null, // 添加 DepartmentId 并初始化，
  ScheduleId: "",
  doctorName: "",
  Data: ref(''), // 用户选择的日期
  patient_id: "",// 获取了 patient_id 和 schedule_id
  schedule_id: "",
  status: "",
  tableData: [],
  total: 0,
  formVisible: false,
  formVisible2: false,
  form: [],
});
// 生成未来七天的日期列表
const futureDates = ref([]);

onMounted(() => {
  const today = dayjs().startOf('day');
  let dates = [];
  for (let i = 0; i <= 7; i++) {
    dates.push(today.add(i, 'day').format('YYYY-MM-DD'));
  }
  futureDates.value = dates; // 确保 futureDates 是响应式的
  console.log(futureDates.value); // 调试输出

});

// 科室下拉列表
const departments = ref([]);

onMounted(async () => {
  try {
    const response = await getDepartmentList();
    // 假设后端返回的是一个包含科室对象的数组
    console.log(response)
    departments.value = response.message;
    console.log(departments.value)
  } catch (error) {
    console.error('获取科室列表失败:', error);
    ElMessage.error('获取科室列表失败');
  }
});
// 获取当前日期
const today = new Date();
console.log(today)
let day = today.getDate();
let month = today.getMonth() + 1; // January is 0!
let year = today.getFullYear();

if(day < 10) {
  day = '0' + day;
}
if(month < 10) {
  month = '0' + month;
}

// 格式化日期
const formattedDate = ref(`${year}-${month}-${day}`);

//值班列表数据模型
const cheduleList = ref([]);

const getDoctorsChedule = async () => {
  let params = {
    DepartmentId: data.department_id,
    Data: formattedDate.value,
    pageNum: pageNum.value,
    pageSize: pageSize.value,
  };
  let result = await getdoctorschedule(params);
  //渲染列表数据
  console.log(result);
  cheduleList.value = result.message;
  console.log(cheduleList);
  total.value = result.Total;
};
getDoctorsChedule();






/**
 * 查询
 */
const search = async () => {
  let params = {
    DepartmentId: data.department_id, // 使用 reactive 对象的 department_id 且 确保使用 .value
    Data: data.Data.value, // 确保 Data 是一个 ref 并且使用 .value访问其值
    pageNum: pageNum.value,
    pageSize: pageSize.value,
  };
  console.log("params:", params);
  let result = await getdoctorschedule(params);
  //渲染列表数据
  console.log(result);
  cheduleList.value = result.message;
  console.log(cheduleList);
  total.value = result.Total;
};

/**
 * 预约
 */
// 进行预约的函数
// 获取当前日期
const today2 = new Date();
console.log(today2)
let day2 = today2.getDate();
let month2 = today2.getMonth() + 1; // January is 0!
let year2 = today2.getFullYear();

if(day2 < 10) {
  day2 = '0' + day2;
}
if(month2 < 10) {
  month2 = '0' + month2;
}

// 格式化日期
const formattedDate2 = ref(`${year2}-${month2}-${day2} ${today2.toLocaleTimeString()}`);

const makeAppointment = async (scope) => {
  
if(userData.user.PatientId == ""){
  ElMessage.error('请先登录');
}else{

   try {
    // 准备请求数据
    const appointmentData = {
      patient_id: userData.user.PatientId,
      schedule_id: scope.value.ScheduleId,
      created: formattedDate2.value,
    };

    // 发送 POST 请求到后端
    const response = await getAppointmentAdd(appointmentData);

    // 处理响应
    if (response.code === 200) {
      ElMessage.success(response.message);
      console.log(response.balance)
      userData.user.Balance = response.balance
      formVisible.value = false
    }else if(response.code === 201){
      ElMessage.error(response.message);
    }else {
      ElMessage.error('预约失败');
    }
  } catch (error) {
    console.error('预约请求出错:', error);
    ElMessage.error('预约请求出错');
  }
}

let params = {
    DepartmentId: data.department_id,
    Data: formattedDate.value,
    pageNum: pageNum.value,
    pageSize: pageSize.value,
  };
  let result = await getdoctorschedule(params);
  //渲染列表数据
  console.log(result);
  cheduleList.value = result.message;
  console.log(cheduleList);
  total.value = result.Total;
};


const doctor = reactive({
  
  
});
// const doctor = ref();
//弹窗的控制
let formVisible = ref(false)
 const handleadd  = (row) => {
  doctor.value = row
  console.log(doctor.value.Class)
  formVisible.value = true;
};



/**
 * 重置
 */
// const reset = () => {
//   data.department_id = 0;
//   data.Data = "";
//   getdoctorschedule();
// };


/**
 * ref引用的声明
 */
const formRef = ref();



</script>



<template>
  <div>
    <div class="card" style="margin-bottom: 10px">
      <el-select v-model="data.department_id" placeholder="请选择科室" style="width: 200px; margin-right: 10px">
        <el-option  :key=null :label="请选择科室" :value=null></el-option>
        <el-option v-for="dept in departments" :key="dept.DepartmentId" :label="dept.Name" :value="dept.DepartmentId">
        </el-option>
      </el-select>

      <!-- 日期选择下拉列表 -->
      <el-select v-model="formattedDate" placeholder="请选择日期" style="width: 200px; margin-right: 10px">
        <el-option v-for="date in futureDates" :key="date" :label="date" :value="date">
        </el-option>
      </el-select>

      <el-button type="primary" style="margin-left:10px" @click="getDoctorsChedule">查询</el-button>
      <!-- <el-button type="info" @click="reset">重置</el-button> -->
    </div>

    <!--添加-->
    <div class="crad">
      <!--表单-->
      <div>
        <el-table :data="cheduleList" :default-sort="{ prop: 'id', order: 'descending' }" style="width: 100%">
          <!-- <el-table-column type="selection" width="55" align="center"></el-table-column> -->
          <el-table-column prop="ScheduleId" label="排班ID" show-overflow-tooltip></el-table-column>
          <el-table-column prop="Name" label="医生姓名" show-overflow-tooltip></el-table-column>
          <el-table-column label="图片" show-overflow-tooltip>
      <template #default="{ row }">
        <el-image
          :src="row.Img"
          style="width: 50px; height: 50px; border-radius: 50%; object-fit: cover;"
        ></el-image>
      </template>
    </el-table-column>
          <el-table-column prop="Departments" label="科室" show-overflow-tooltip></el-table-column>
          <el-table-column prop="Class" label="类别" show-overflow-tooltip></el-table-column>
          <el-table-column prop="Introduction" label="医生简介" show-overflow-tooltip></el-table-column>
          <el-table-column prop="Data" label="日期" show-overflow-tooltip></el-table-column>
          <el-table-column label="开始时间" prop="start_time"></el-table-column>
          <el-table-column label="结束时间" prop="end_time"></el-table-column>
          <el-table-column prop="Price" label="价格" show-overflow-tooltip></el-table-column>
          <el-table-column prop="Number" label="剩余名额" show-overflow-tooltip></el-table-column>

          <el-table-column label="操作" align="center" width="160">
            <template #default="scope">
              <!-- <div v-if="isPatient"> -->
              <!-- <el-button type="primary" @click="makeAppointment(scope.row)">预约</el-button> -->
              <el-button type="primary" @click="handleadd(scope.row)">预约</el-button>
              <!-- </div> -->
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
    <!-- 分页条 -->
    <el-pagination v-model:current-page="pageNum" v-model:page-size="pageSize" :page-sizes="[3, 5, 10, 15]"
      layout="jumper, total, sizes, prev, pager, next" background :total="total" @size-change="onSizeChange"
      @current-change="onCurrentChange" style="margin-top: 20px; justify-content: flex-end" />

    <!--修改弹窗-->


    <!-- 预约弹窗 -->
    <el-dialog
      width="40%"
      v-model="formVisible"
      title="预约"
      destroy-on-close>
      <el-form
        :rules="rules"
        ref="formRef"
        label-width="100px"
        label-position="right"
        stylr="padding-right:40px"
      >
      <el-form
    :rules="rules"
    ref="formRef"
    label-width="100px"
    label-position="right"
    stylr="padding-right:40px"
  >
  <el-form-item label="医生姓名:">
      <span>{{ doctor.value.Name }}</span>
    </el-form-item>
    <el-form-item label="医生图片:">
      <el-avatar shape="square" :size="150" :src="doctor.value.Img" />
    </el-form-item>
    <el-form-item label="所属科室:">
      <span>{{ doctor.value.Class }}</span>
    </el-form-item>
    <el-form-item label="医生简介:">
      <span>{{ doctor.value.Introduction }}</span>
    </el-form-item>
    <el-form-item label="预约日期:">
      <span>{{ doctor.value.Data }}</span>
    </el-form-item>
    <el-form-item label="开始时间:">
      <span>{{ doctor.value.start_time }}</span>
    </el-form-item>
    <el-form-item label="结束时间:">
      <span>{{ doctor.value.end_time }}</span>
    </el-form-item>
    <el-form-item label="价格:">
      <span>{{ doctor.value.Price }}</span>
    </el-form-item>
    <el-form-item label="剩余名额:">
      <span>{{ doctor.value.Number }}</span>
    </el-form-item>
  </el-form>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="formVisible = false">取 消</el-button>
          <el-button type="primary" @click="makeAppointment(doctor)">确认</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
