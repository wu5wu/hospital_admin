<script setup>
import { ref, reactive ,onMounted} from "vue";
import { Search } from "@element-plus/icons-vue";
import request from "@/utils/request";
import { ElDrawer, ElMessage, ElMessageBox } from "element-plus";
import { getDepartmentList, } from "@/api/appointment.js";
import { GetDoctors,DeleteDoctors ,UpdateDoctors,AddDoctors} from "@/api/doctor.js";
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


const departments = ref([]);
const departmentId = ref(null);
onMounted(async () => {
  try {
    const response = await getDepartmentList();
    // 假设后端返回的是一个包含科室对象的数组
    
    departments.value = response.message;
    console.log(departments.value)
  } catch (error) {
    console.error('获取科室列表失败:', error);
    ElMessage.error('获取科室列表失败');
  }
});




//医生列表数据模型
const doctorList = ref([]);

const getDoctors = async () => {
  let params = {
    pageNum: pageNum.value,
    pageSize: pageSize.value,
  };
  let result = await GetDoctors(params);
  //渲染列表数据
  console.log(result);
  doctorList.value = result.message;
 
  total.value = result.Total;
};
getDoctors();


const data = reactive({
  
  Username: "",
  Password: "",
  Name: "",
  DepartmentId: "",
  Introduction: "",
  Price: "",
  formVisible: false,
  formVisible2: false,
  form: [],
});


/**
 * 查询
 */
const search = async () => {
  let params = {
    Name :data.Name ,
    DepartmentId :departmentId.value ,
    pageNum: pageNum.value,
    pageSize: pageSize.value,
  };

  let result = await GetDoctors(params);
  //渲染列表数据

  doctorList.value = result.message;

  total.value = result.Total;
};

/**
 * 重置
 */
const reset = () => {
  data.Name = "";
  departmentId.value = null;
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

/**
 * 添加、修改保存数据到后台
 */
const save = async () => {
  let datas={
    DoctorId :  data.form.DoctorId,
    Name : data.form.Name,
    DepartmentId :data.form.DepartmentId ,
    Introduction : data.form.Introduction,
    Price :  parseInt(data.form.Price, 10),
    State : data.form.State
  }
  let result = await UpdateDoctors(datas);
if(result.code==200){
  ElMessage.success(result.message);
  search()

}
data.formVisible = false
  
};

/**
 * 修改保存数据到后台
 */
const handleEdit = (row) => {
  data.form = JSON.parse(JSON.stringify(row));
  data.formVisible = true;
};

const handleadd  = () => {
  Department_id.value = ""
  data.formVisible2 = true;
};

const Department_id = ref()
const Add = async () => {
  let datas = {
    Username:data.Username,
    Password :  data.Password,
    Name : data.Name,
    DepartmentId :Department_id.value ,
    Introduction : data.Introduction,
    Price :parseInt(data.Price, 10) ,
  };
  console.log(datas)
  let result = await AddDoctors(datas);
  if (result.code == 200) {
    ElMessage.success(result.message);
    data.formVisible2 = false
  } else {
    ElMessage.error(result.message);
  }
  search();

};

/**
 * 删除数据
 */
const handleDelete = (Data) => {
  ElMessageBox.confirm("删除数据后无法恢复，您确认删除吗？", "删除确认", {
    type: "warning",
  }).then(async () => {
  
    let DoctorId = {
      DoctorId : Data
    }
    let result = await DeleteDoctors(DoctorId);
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
        placeholder="请输入医生姓名"
        :prefix-icon="Search"
      ></el-input>

      <el-select v-model="departmentId" placeholder="请选择科室" style="width: 200px; margin-right: 10px">
        <el-option  :key=null :label="请选择科室" :value=null></el-option>
        <el-option v-for="dept in departments" :key="dept.DepartmentId" :label="dept.Name" :value="dept.DepartmentId">
        </el-option>
      </el-select>


      <el-button type="primary" tyle="margin-left:10px" @click="search"
        >查询</el-button
      >
      <el-button type="info" @click="reset">重置</el-button>
      <el-button type="primary" tyle="margin-left:10px" @click="handleadd">添加</el-button>
    </div>

  
    <div class="crad">
      <!--表单-->
      <div>
        <el-table
          :data="doctorList"
          :default-sort="{ prop: 'id', order: 'descending' }"
          @selection-change="handleSelectionChange"
          style="width: 100%"
        >
          <!--          <el-table-column type="selection" width="55" align="center"></el-table-column>-->
          <el-table-column
            prop="DoctorId"
            label="医生编号"
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
            prop="Department"
            label="所属科室"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            prop="Introduction"
            label="个人简介"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            prop="Price"
            label="价格"
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
      <el-table-column
            prop="State"
            label="状态"
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
                @click="handleDelete(scope.row.DoctorId)"
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
        <el-form-item label="医生编号">
          <el-input
            v-model="data.form.DoctorId"
            prop="schedule_id"
            autocomplete="off"
            disabled
          />
        </el-form-item>
        <el-form-item label="用户名">
          <el-input
            v-model="data.form.Username"
            prop="schedule_id"
            autocomplete="off"
            disabled
          />
        </el-form-item>
        <el-form-item label="姓名">
          <el-input
            v-model="data.form.Name"
            
            autocomplete="off"
          />
        </el-form-item>
        
        <el-form-item label="所属科室">
          <!-- <el-input
            v-model="data.form.Department"
            autocomplete="off"
          /> -->
          <el-select v-model="data.form.DepartmentId" placeholder="请选择科室" style="width: 200px; margin-right: 10px">
        
        <el-option v-for="dept in departments" :key="dept.DepartmentId" :label="dept.Name" :value="dept.DepartmentId">
        </el-option>
      </el-select>
        </el-form-item>
        <el-form-item label="个人简介">
          <el-input
            v-model="data.form.Introduction"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item label="价格">
          <el-input
            v-model="data.form.Price"
            autocomplete="off"
          />
        </el-form-item>
    
        <el-form-item label="状态">
       
        <el-select placeholder="请选择" v-model="data.form.State">
          <el-option label="空闲" value="空闲"></el-option>
          <el-option label="值班中" value="值班中"></el-option>
          <el-option label="手术中" value="手术中"></el-option>
          <el-option label="忙碌" value="忙碌"></el-option>
        </el-select>
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
      title="添加医生"
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
      
        <el-form-item label="用户名">
          <el-input
            v-model="data.Username"
            prop="schedule_id"
            autocomplete="off"
          
          />
        </el-form-item>
        <el-form-item label="密码">
          <el-input
            v-model="data.Password"
            prop="schedule_id"
            autocomplete="off"
        
          />
        </el-form-item>

        <el-form-item label="姓名">
          <el-input
            v-model="data.Name"
            
            autocomplete="off"
          />
        </el-form-item>
        
        <el-form-item label="所属科室">
        
          <el-select v-model="Department_id" placeholder="请选择科室" style="width: 200px; margin-right: 10px">
        
            <el-option v-for="dept in departments" :key="dept.DepartmentId" :label="dept.Name" :value="dept.DepartmentId">
            </el-option>
          </el-select>
        </el-form-item>


        <el-form-item label="个人简介">
          <el-input
            v-model="data.Introduction"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item label="价格">
          <el-input
            v-model="data.Price"
            autocomplete="off"
          />
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