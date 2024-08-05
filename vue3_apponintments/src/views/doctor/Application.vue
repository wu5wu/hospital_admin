<script setup>
import { ref, reactive } from "vue";
import { Search } from "@element-plus/icons-vue";
import request from "@/utils/request";
import { ElDrawer, ElMessage, ElMessageBox } from "element-plus";
import { Getapplication,Addapplication,Setapplication} from "@/api/doctor.js";
import { Check, Delete, Edit, Message, Star,Close } from "@element-plus/icons-vue";
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
    reason:"",
    ApplicationId:"",
    formVisible:false,
    formVisible2:false
});
//医生列表数据模型
const patientList = ref([]);

const search = async () => {
  let params = {
    doctorId : doctorData.Doctor.DoctorId,
    Status : data.Status,
    pageNum: pageNum.value,
    pageSize: pageSize.value,
  };

  let result = await Getapplication(params);
  //渲染列表数据

  patientList.value = result.message;

  total.value = result.Total;
};
search()

const endappointment = (Data) => {
  ElMessageBox.confirm("撤销数据后无法恢复，您确认取消申请吗？", "取消确认", {
    type: "warning",
  }).then(async () => {
  console.log(Data)
    let ApplicationId = {
      ApplicationId : Data,
      State: "撤销"
    }
    let result = await Setapplication(ApplicationId);
    if (result.code === 200) {
      search(); // 重新获取数据
      ElMessage.success(result.message);
    } else {
      ElMessage.error(result.message);
    }
  });
};

const Add = () => {
  data.formVisible = true
}

const addapplication = async () => {
  let params = {
    doctorId : doctorData.Doctor.DoctorId,
    data : data.Data,
    reason : data.reason,
    pageNum: pageNum.value,
    pageSize: pageSize.value,
  };

  let result = await Addapplication(params);
  //渲染列表数据
  if (result.code === 200) {
      search(); // 重新获取数据
      ElMessage.success(result.message);
    } else {
      ElMessage.error(result.message);
    }

data.formVisible = false

};
const handleEdit = (row) => {
    data.formVisible2 = true
data.Data = row.Data
data.reason = row.Reason
data.ApplicationId = row.ApplicationId
};

const setapplication = async () => {
  let params = {
    ApplicationId : data.ApplicationId,
    data : data.Data,
    reason : data.reason,
  };

  let result = await Setapplication(params);
  //渲染列表数据
  if (result.code === 200) {
      search(); // 重新获取数据
      ElMessage.success(result.message);
    } else {
      ElMessage.error(result.message);
    }

data.formVisible2 = false
};

</script>





<template>
  <div>
    <div class="card" style="margin-bottom: 10px">
      <el-button type="primary" tyle="margin-left:10px" @click="Add">请假</el-button>
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
            prop="ApplicationId"
            label="请假编号"
            show-overflow-tooltip
          ></el-table-column>

           <el-table-column
            prop="Data"
            label="请假日期"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            prop="Reason"
            label="理由"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            prop="State"
            label="状态"
            show-overflow-tooltip
          ></el-table-column>

          <el-table-column
            prop="CTime"
            label="创建时间"
            show-overflow-tooltip
          ></el-table-column>


          <el-table-column label="操作" align="center" width="160">
            <template #default="{row}">
              <el-button
                type=""
                @click="handleEdit(row)"
                :icon="Edit"
                circle
              />
              <el-button
                type=""
                @click="endappointment(row.ApplicationId)"
                :icon="Close"
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
  </div>


<!-- 请假弹窗 -->
  <el-dialog
      width="40%"
      v-model="data.formVisible"
      title="请假申请表"
      destroy-on-close
    >
      <el-form
        :model="data"
        :rules="rules"
        ref="formRef"
        label-width="100px"
        label-position="right"
        stylr="padding-right:40px"
      >
      
        <el-form-item label="请假日期">
            <el-date-picker
              v-model="data.Data"
              type="date"
              placeholder="选择日期"
              format="YYYY/MM/DD"
              value-format="YYYY-MM-DD"/>

        </el-form-item>
        <el-form-item label="理由">
          <el-input
            v-model="data.reason"
            prop="schedule_id"
            autocomplete="off"/>
        </el-form-item>

       
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="data.formVisible = false">取 消</el-button>
          <el-button type="primary" @click="addapplication">保 存</el-button>
        </span>
      </template>
    </el-dialog>
  

<!-- 修改弹窗 -->
    <el-dialog
      width="40%"
      v-model="data.formVisible2"
      title="请假申请表"
      destroy-on-close
    >
      <el-form
        :model="data"
        :rules="rules"
        ref="formRef"
        label-width="100px"
        label-position="right"
        stylr="padding-right:40px"
      >
      
        <el-form-item label="请假日期">
            <el-date-picker
              v-model="data.Data"
              type="date"
              placeholder="选择日期"
              format="YYYY/MM/DD"
              value-format="YYYY-MM-DD"/>

        </el-form-item>
        <el-form-item label="理由">
          <el-input
            v-model="data.reason"
            prop="schedule_id"
            autocomplete="off"/>
        </el-form-item>

       
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="data.formVisible2 = false">取 消</el-button>
          <el-button type="primary" @click="setapplication">保 存</el-button>
        </span>
      </template>
    </el-dialog>



</template>