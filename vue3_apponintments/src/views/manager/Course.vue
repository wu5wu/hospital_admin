<script setup>
import { ref, reactive } from "vue";
import { Search } from "@element-plus/icons-vue";
import request from "@/utils/request";
import { ElDrawer, ElMessage, ElMessageBox } from "element-plus";
import {
  getDoctorsCheduleList,
  finfdoctorschedule,
  intelligentScheduling,
  deletedoctorschedule,
  adddoctorschedule,
  Setdoctorschedule,
  deletedoctorschedules,
} from "@/api/doctorsChedule.js";

import { Check, Delete, Edit, Message, Star } from "@element-plus/icons-vue";
import { Plus } from "@element-plus/icons-vue";

//分页条数据模型
const pageNum = ref(1); //当前页
const total = ref(20); //总条数
const pageSize = ref(5); //每页条数

//当每页条数发生了变化，调用此函数
const onSizeChange = (size) => {
  pageSize.value = size;
  search();
};
//当前页码发生变化，调用此函数
const onCurrentChange = (num) => {
  pageNum.value = num;
  search();
};

//排班列表数据模型
const cheduleList = ref([]);

const data = reactive({
  schedule_id: "",
  doctor_id: "",
  data: "",
  timing:  ref(0),
  start_time: "",
  end_time: "",
  state: "",
  name: "",
  department: "",
  tableData: [],
  total: 0,
  formVisible: false,
  formVisible2: false,
  form: [],
});
// //添加表单数据模型
// const adds = ref({
//   data: '',
//   name: '',
//   timing: '',
//   start_time: "",
//   end_time: "",
// });

/**
 * 查询
 */
const search = async () => {
  let params = {
    Name: data.doctorName,
    Data: data.data,
    pageNum: pageNum.value,
    pageSize: pageSize.value,
  };
  console.log("params:", params);
  let result = await finfdoctorschedule(params);
  //渲染列表数据
  console.log(result);
  cheduleList.value = result.message;
  console.log(cheduleList);
  total.value = result.Total;
};
search();
/**
 * 重置
 */
const reset = () => {
  data.doctorName = "";
  data.data = "";
  search();
};

/**
 * 定义表单规则
 */
const rules = reactive({
  schedule_id: [{ required: true, message: "账号不可为空", trigger: "blur" }],
});

/**
 * ref引用的声明
 */
const formRef = ref();

/**
 * 添加
 */
const handleAdd = async () => {
  let result = await intelligentScheduling();
  if (result.code == 200) {
    ElMessage.success(result.message);
  } else {
    ElMessage.error("排班失败");
  }
  search();
};

/**
 * 添加、修改保存数据到后台
 */
const save = async () => {
  let datas={
    ScheduleId: data.form.ScheduleId,
    Data : data.form.Data,
    Timing :parseInt(data.form.Timing, 10),
    start_time: data.form.start_time,
    end_time:data.form.end_time
  }
  let result = await Setdoctorschedule(datas);
if(result.code==200){
  ElMessage.success(result.message);
  search()

}
data.formVisible = false
  // formRef.value.validate((valid) => {
  //   // 表单的验证
  //   if (valid) {
  //     request
  //       .request({
  //         url: data.form.schedule_id ? baseUrl + "/update" : baseUrl + "/add",
  //         method: data.form.schedule_id ? "PUT" : "POST",
  //         data: data.form,
  //       })
  //       .then((res) => {
  //         if (res.code === "200") {
  //           search(); // 重新获取数据
  //           data.formVisible = false; // 关闭弹窗
  //           ElMessage.success("操作成功");
  //         } else {
  //           ElMessage.error(res.msg);
  //         }
  //       });
  //   }
  // });
};

/**
 * 修改保存数据到后台
 */
const handleEdit = (row) => {
  data.form = JSON.parse(JSON.stringify(row));
  data.formVisible = true;
};

const handleadd  = () => {
  data.name = "";
  data.data = "";
  data.timing = "",
  data.start_time = "";
  data.end_time = "";
  data.formVisible2 = true;
};


const Add = async () => {
  let datas = {
    name: data.name,
    data: data.data,
    Timing: Number(data.timing),
    start_time: data.start_time,
    end_time: data.end_time,
  };
  console.log(datas)
  let result = await adddoctorschedule(datas);
  if (result.code == 200) {
    ElMessage.success(result.message);
  } else {
    ElMessage.error(result.message);
  }
  search();
};

/**
 * 删除数据
 */
const handleDelete = (schedule_id) => {
  ElMessageBox.confirm("删除数据后无法恢复，您确认删除吗？", "删除确认", {
    type: "warning",
  }).then(async () => {
    // .then((res) => {
    // request.delete(baseUrl + "/delete/" + schedule_id).then((res) => {
    //   if (res.code === "200") {
    //     search(); // 重新获取数据
    //     ElMessage.success("操作成功");
    //   } else {
    //     ElMessage.error(res.msg);
    //   }
    // });
    console.log(schedule_id);
    let result = await deletedoctorschedule(schedule_id);
    if (result.code === 200) {
      search(); // 重新获取数据
      ElMessage.success(result.message);
    } else {
      ElMessage.error(result.msg);
    }
  });
};


const selectedScheduleIds = ref([]);
const ScheduleId = ref([]);
// 处理选中行变化
const handleSelectionChange = (selection) => {
  selectedScheduleIds.value = selection;
  
};
// 获取批量删除的数据
const delBatch = () => {
  if (selectedScheduleIds.value.length === 0) {
    ElMessage.warning('请先选择要删除的科室');
    return;
  }

  ElMessageBox.confirm('确认删除选中的科室吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(() => {
    selectedScheduleIds.value.forEach(ScheduleIds => {
      console.log(ScheduleIds.ScheduleId)
      ScheduleId.value.push(ScheduleIds.ScheduleId)
    });
    deleteDepartments()
    
    selectedScheduleIds.value = []; // 清空选中状态
  }).catch(() => {
    ElMessage.info('取消删除');
  });
};
//批量删除
const deleteDepartments = async() => {
let result = await deletedoctorschedules(ScheduleId.value);
if (result.code == 200) {
  search()
  ElMessage.success(result.message);
}else{
  ElMessage.err(result.error);
}
}
</script>





<template>
  <div>
    <div class="card" style="margin-bottom: 10px">
      <el-input
        style="width: 200px; margin-right: 10px"
        v-model="data.doctorName"
        placeholder="请输入医生姓名"
        :prefix-icon="Search"
      ></el-input>
      <!-- <el-input
        style="width: 200px; margin-right: 10px"
        v-model="data.data"
        placeholder="请输入日期"
        :prefix-icon="Search"
      ></el-input> -->
     
            <el-date-picker
              v-model="data.data"
              type="date"
              placeholder="请选择日期"
              format="YYYY/MM/DD"
              value-format="YYYY-MM-DD"/>
        
      <el-button type="primary" tyle="margin-left:10px" @click="search"
        >查询</el-button
      >
      <el-button type="info" @click="reset">重置</el-button>
      <el-button type="primary" tyle="margin-left:10px" @click="handleadd">添加</el-button>
    </div>

    <!--添加-->
    <div class="crad">
      <div style="margin-bottom: 20px">
        <el-button type="primary" plain @click="handleAdd">智能排班</el-button>
        <el-button type="danger" plain @click="delBatch">批量删除</el-button>
      </div>
      <!--表单-->
      <div>
        <el-table
          :data="cheduleList"
          :default-sort="{ prop: 'id', order: 'descending' }"
          @selection-change="handleSelectionChange"
          style="width: 100%"
        >
                   <el-table-column type="selection" width="55" align="center"></el-table-column>
          <el-table-column
            prop="ScheduleId"
            label="排班序号"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            prop="DoctorId"
            label="医生ID"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            prop="Data"
            label="日期"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            prop="start_time"
            label="起始时间"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            prop="end_time"
            label="结束时间"
            show-overflow-tooltip
          ></el-table-column>
          
          <el-table-column
            prop="Name"
            label="医生"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            prop="Departments"
            label="科室"
            show-overflow-tooltip
          ></el-table-column>

          <el-table-column
            prop="Number"
            label="剩余名额"
            show-overflow-tooltip
          ></el-table-column>

          <el-table-column label="操作" align="center" width="160">
            <template #default="scope">
              <el-button
                type="primary"
                @click="handleEdit(scope.row)"
                :icon="Edit"
                circle
              />
              <el-button
                type="danger"
                @click="handleDelete(scope.row.ScheduleId)"
                :icon="Delete"
                circle
              />
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

    <!--修改弹窗-->
    <el-dialog
      width="40%"
      v-model="data.formVisible"
      title="信息"
      destroy-on-close
    >
      <el-form
        :model="data.form"
        :rules="rules"
        ref="formRef"
        label-width="100px"
        label-position="right"
        stylr="padding-right:40px"
      >
        <el-form-item label="排班序号">
          <el-input
            v-model="data.form.ScheduleId"
            prop="schedule_id"
            autocomplete="off"
            disabled
          />
        </el-form-item>
        <el-form-item label="医生编号">
          <el-input
            v-model="data.form.DoctorId"
            prop="schedule_id"
            autocomplete="off"
          />
        </el-form-item>
      
        <el-form-item label="日期" prop="Data">
          
          <div class="block">
            <el-date-picker
              v-model="data.form.Data"
              type="date"
              placeholder="Pick a Date"
              format="YYYY/MM/DD"
              value-format="YYYY-MM-DD"/>
          </div>
    
        </el-form-item>
        
        <el-form-item label="timing">
        <el-select placeholder="请选择" v-model="data.form.Timing">
          <el-option label="早上" value=0></el-option>
          <el-option label="下午" value=1></el-option>
        </el-select>
      </el-form-item>
        <el-form-item label="起始时间">
          <el-input
            v-model="data.form.start_time"
            type="string"
            placeholder="选择结束时间"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item label="结束时间">
          <el-input
            v-model="data.form.end_time"
            type="datetime"
            placeholder="选择结束时间"
            autocomplete="off"
          />
        </el-form-item>
       
      </el-form>

      <template #footer>
        <span class="dialog-footer">
          <el-button @click="data.formVisible = false">取 消</el-button>
          <el-button type="primary" @click="save">保 存</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 添加弹窗 -->
    <el-dialog
      width="40%"
      v-model="data.formVisible2"
      title="创建科室"
      destroy-on-close
    >
      <el-form
        :model="data.form"
        :rules="rules"
        ref="formRef"
        label-width="100px"
        label-position="right"
        stylr="padding-right:40px"
      >
        <el-form-item label="医生姓名">
          <el-input v-model="data.name" prop="Name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="日期" prop="data">
          <!-- <div class="block">
            <el-date-picker
              v-model="data.data"
              type="date"
              placeholder="Pick a day"
              :size="size"
            />
          </div> -->
          <el-date-picker
        v-model="data.data"
        type="date"
        placeholder="Pick a Date"
        format="YYYY/MM/DD"
        value-format="YYYY-MM-DD"
      />
        </el-form-item>
        
        <el-form-item label="timing">
        <el-select placeholder="请选择" v-model="data.timing">
          <el-option label="早上" value=0></el-option>
          <el-option label="下午" value=1></el-option>
        </el-select>
      </el-form-item>
        <el-form-item label="起始时间">
          <el-select placeholder="请选择" v-model="data.start_time">
          <el-option label="09:00:00" value="09:00:00"></el-option>
          <el-option label="14:00:00" value="14:00:00"></el-option>
        </el-select>
        </el-form-item>
        <el-form-item label="结束时间">
          <el-select placeholder="请选择" v-model="data.end_time">
          <el-option label="12:00:00" value="12:00:00"></el-option>
          <el-option label="18:00:00" value="18:00:00"></el-option>
        </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="data.formVisible2 = false">取 消</el-button>
          <el-button type="primary" @click="Add">保 存</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>