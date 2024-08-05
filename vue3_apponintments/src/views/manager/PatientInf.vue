<script setup>
import { ref, reactive ,onMounted} from "vue";
import { Search } from "@element-plus/icons-vue";
import request from "@/utils/request";
import { ElDrawer, ElMessage, ElMessageBox } from "element-plus";
import { GetpatientList,Deletepatient} from "@/api/patient.js";
import { Check, Delete, Edit, Message, Star } from "@element-plus/icons-vue";
import { Plus } from "@element-plus/icons-vue";

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

  Name: "",
 
});
//医生列表数据模型
const patientList = ref([]);

const search = async () => {
  let params = {
    Name :data.Name ,
    pageNum: pageNum.value,
    pageSize: pageSize.value,
  };

  let result = await GetpatientList(params);
  //渲染列表数据

  patientList.value = result.message;

  total.value = result.Total;
};
search()

const reset = () => {
  data.Name = "";
  search();
};


const handleDelete = (Data) => {
  ElMessageBox.confirm("删除数据后无法恢复，您确认删除吗？", "删除确认", {
    type: "warning",
  }).then(async () => {
  console.log(Data)
    let patientId = {
        patientId : Data
    }
    let result = await Deletepatient(patientId);
    if (result.code === 200) {
      search(); // 重新获取数据
      ElMessage.success(result.message);
    } else {
      ElMessage.error(result.msg);
    }
  });
};


</script>





<template>
  <div>
    <div class="card" style="margin-bottom: 10px">
      <el-input
        style="width: 200px; margin-right: 10px"
        v-model="data.Name"
        placeholder="请输入用户姓名"
        :prefix-icon="Search"
      ></el-input>

      <el-button type="primary" tyle="margin-left:10px" @click="search">查询</el-button>
      <el-button type="info" @click="reset">重置</el-button>
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
            prop="PatientId"
            label="用户编号"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            prop="Username"
            label="用户名"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            prop="Name"
            label="姓名"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            prop="Identity"
            label="身份证"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            prop="Phone"
            label="电话号码"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            prop="Balance"
            label="余额"
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
              <el-button
                type="danger"
                @click="handleDelete(row.PatientId)"
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


   
  </div>
</template>