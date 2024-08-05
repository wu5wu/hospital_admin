<script setup>
import { ref, reactive ,onMounted} from "vue";
import { Search } from "@element-plus/icons-vue";
import request from "@/utils/request";
import { ElDrawer, ElMessage, ElMessageBox } from "element-plus";
import { Gtedoctorschedule} from "@/api/doctor.js";
import { Check, Delete, Edit, Message, Star,Close } from "@element-plus/icons-vue";
import { Plus } from "@element-plus/icons-vue";
import { doctorStore } from '@/stores/doctor.js'
const doctorData = doctorStore();


//分页条数据模型
const pageNum = ref(1); //当前页
const total = ref(20); //总条数
const pageSize = ref(10); //每页条数

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
    Doctor:"",
    reason:"",
    ApplicationId:"",
    
});
//值班列表数据模型
const shifttList = ref([]);

const search = async () => {
  let params = {
    doctorId : doctorData.Doctor.DoctorId,
    pageNum: pageNum.value,
    pageSize: pageSize.value,
  };

  let result = await Gtedoctorschedule(params);
  //渲染列表数据

  shifttList.value = result.message;

  total.value = result.Total;
};
search()


</script>





<template>
  <div>
   
    <div class="crad">
      <!--表单-->
      <div>
        <el-table
          :data="shifttList"
          :default-sort="{ prop: 'id', order: 'descending' }"
          style="width: 100%"
        >
          <!--          <el-table-column type="selection" width="55" align="center"></el-table-column>-->
          <el-table-column
            prop="ScheduleId"
            label="排班编号"
            show-overflow-tooltip
          ></el-table-column>

           <el-table-column
            prop="Data"
            label="值班日期"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            prop="start_time"
            label="开始时间"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            prop="end_time"
            label="结束时间"
            show-overflow-tooltip
          ></el-table-column>

          <el-table-column
            prop="Departments"
            label="科室"
            show-overflow-tooltip
          ></el-table-column>

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