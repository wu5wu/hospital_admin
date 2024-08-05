<script setup>
import {ref,reactive} from "vue";
import {Search} from '@element-plus/icons-vue';
import request from "@/utils/request";
import {ElDrawer, ElMessage, ElMessageBox} from 'element-plus';
import {
  Check,
  Delete,
  Edit,
  Message,
  Star,
} from '@element-plus/icons-vue';
import {Plus} from "@element-plus/icons-vue";
import { AddDepartmentService,DepartmentDeleteService,DepartmentUpdateService,DepartmentFindService,DepartmentDeleteServices} from "@/api/departments.js";


// 分页条数据模型
const pageNum = ref(1); //当前页
const total = ref(15); //总条数
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

//数据列表（科室）
const departments = ref([]);


//添加表单数据模型
const adddepartments = ref({
  id: '',
  Name: '',
  Class: ''
});


const add = async () => {
  let result = await AddDepartmentService(adddepartments.value);
  if (result.code == 200) {
    adddepartments.value = {}
    data.formVisible = false
    search()
    
  } else {
    adddepartments.value = {}
    ElMessage.error(result.message);
    data.formVisible = false
  }
};


// const  baseUrl = '/department'

const data = reactive({
  id: '',
  Name: '',
  Class: '',
  tableData: [],
  
  
})

/**
 * 查询
 */

 const finglist = ref({
  id: '',
  Name: '',
  Class: ''
});
const search =async () => {
  let data ={
    Name: finglist.value.Name,
    pageNum: pageNum.value,
    pageSize: pageSize.value,
  }
  let result = await DepartmentFindService(data);
  departments.value = result.message
  total.value = result.Total
}
search()


// const handleCurrentChange = (pageNum) => {
//   //当翻页的时候重新加载数据即可
//   search()
// }

// /**
//  * 重置
//  */
const reset = () => {
  finglist.value.Name = ""
  search()
}



/**
 * 添加
 */
const handleAdd = () => {
  
  data.form = {}
  data.formVisible = true
}




// /**
//  * 修改保存数据到后台
//  */
const handleEdit = (row) => {
  adddepartments.value = row
  data.formVisibles = true
}
const update = async () => {
  let result = await DepartmentUpdateService(adddepartments.value);
  console.log("22222222");
  console.log(result.code);
  if (result.code == 200) {
    ElMessage.success("修改成功");
    adddepartments.value = {}
    data.formVisible = false
    search()
  } else {
    adddepartments.value = {}
    ElMessage.error("修改失败");
    data.formVisible = false
  }
};

// /**
//  * 删除数据
//  */


const deleteDepartment = (row) => {
  ElMessageBox.confirm("删除数据后无法恢复，您确认删除吗？？", "温馨提示", {
    confirmButtonText: "确认",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      //用户点击了确认
      let result = await DepartmentDeleteService(row);
      ElMessage.success(result.message);
      //再次调用search，获取所有
      search();
    })
    .catch(() => {
      //用户点击了取消
      ElMessage({
        type: "info",
        message: "取消删除",
      });
    });
};


const selectedDepartments = ref([]);
const DepartmentId = ref([]);
// 处理选中行变化
const handleSelectionChange = (selection) => {
  selectedDepartments.value = selection;
  console.log(selectedDepartments.value)
};
// 获取批量删除的数据
const delBatch = () => {
  if (selectedDepartments.value.length === 0) {
    ElMessage.warning('请先选择要删除的科室');
    return;
  }

  ElMessageBox.confirm('确认删除选中的科室吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(() => {
    selectedDepartments.value.forEach(Department => {
      console.log(Department.DepartmentId)
      DepartmentId.value.push(Department.DepartmentId)
    });
    deleteDepartments()
    
    selectedDepartments.value = []; // 清空选中状态
  }).catch(() => {
    ElMessage.info('取消删除');
  });
};
//批量删除
const deleteDepartments = async() => {
let result = await DepartmentDeleteServices(DepartmentId.value);
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
    <div class="card" style="margin-bottom: 10px;">
      <el-input style="width: 200px; margin-right: 10px" v-model="finglist.Name" placeholder="请输入科室名称"
                :prefix-icon="Search"></el-input>
      <el-button type="primary" tyle="margin-left:10px" @click="search">查询</el-button>
      <el-button type="info" @click="reset">重置</el-button>
    </div>

    <!--添加-->
    <div class="crad">
      <div style="margin-bottom:20px ">
        <el-button type="primary" plain @click="handleAdd">新增</el-button>
        <el-button type="danger" plain @click="delBatch">批量删除</el-button>
      </div>
      <!--表单-->
      <div>
        <el-table :data="departments"  :default-sort="{ prop: 'DepartmentId', order: 'ascending' }" 
        @selection-change="handleSelectionChange"
        style="width: 100%">
          <el-table-column type="selection" width="55" align="center"></el-table-column>
          <el-table-column prop="DepartmentId" label="序号" width="80" align="center" sortable></el-table-column>
          <el-table-column prop="Name" label="科室名称" show-overflow-tooltip></el-table-column>
          <el-table-column prop="Class" label="科室类别" show-overflow-tooltip></el-table-column>

          <el-table-column label="操作" align="center" width="160">
            <template #default="{ row }">
              <el-button type="primary" @click="handleEdit(row)" :icon="Edit" circle/>
              <el-button type="danger" @click="deleteDepartment(row)" :icon="Delete" circle/>
            </template>
          </el-table-column>
        </el-table>
      </div>
      
    </div>
    
   
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

    <el-dialog width="40%" v-model="data.formVisible" title="创建科室" destroy-on-close>
      <el-form :model="data.form" :rules="rules" ref="formRef" label-width="100px" label-position="right" stylr="padding-right:40px">
                <!-- <el-form-item label="ID">
                  <el-input v-model="data.form.id" prop="departmentId" autocomplete="off"/>
                </el-form-item> -->
        <el-form-item label="科室名称">
          <el-input v-model="adddepartments.Name" prop="Name" autocomplete="off"/>
        </el-form-item>
        <el-form-item label="科室类别">
          <el-input v-model="adddepartments.Class" prop="Class" autocomplete="off"/>
        </el-form-item>
      </el-form>
      <template #footer>
      <span class="dialog-footer">
        <el-button @click="data.formVisible = false">取 消</el-button>
        <el-button type="primary" @click="add">保 存</el-button>
      </span>
      </template>
    </el-dialog>


     <!--修改-->
     <el-dialog width="40%" v-model="data.formVisibles" title="修改科室" destroy-on-close>
      <el-form :model="data.form" :rules="rules" ref="formRef" label-width="100px" label-position="right" stylr="padding-right:40px">
                <!-- <el-form-item label="ID">
                  <el-input v-model="data.form.id" prop="departmentId" autocomplete="off"/>
                </el-form-item> -->
        <el-form-item label="科室名称">
          <el-input v-model="adddepartments.Name" prop="Name" autocomplete="off"/>
        </el-form-item>
        <el-form-item label="科室类别">
          <el-input v-model="adddepartments.Class" prop="Class" autocomplete="off"/>
        </el-form-item>
      </el-form>
      <template #footer>
      <span class="dialog-footer">
        <el-button @click="data.formVisibles = false">取 消</el-button>
        <el-button type="primary" @click="data.formVisibles = false;update()">保 存</el-button>
      </span>
      </template>
    </el-dialog>
  </div>
</template>

