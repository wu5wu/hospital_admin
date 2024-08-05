<script setup>
import { ref ,onMounted} from "vue";
import { doctorStore } from '@/stores/doctor.js'
const doctorData = doctorStore();

const doctorInfo = ref({
  DoctorId: doctorData.Doctor.DoctorId,
  Name: doctorData.Doctor.Name,
  Username: doctorData.Doctor.Username,
  DepartmentId: doctorData.Doctor.DepartmentId,
  Introduction: doctorData.Doctor.Introduction,
  Img: doctorData.Doctor.Img,
  State:  doctorData.Doctor.State,
});
import { getDepartmentList, } from "@/api/appointment.js";
import {UpdateDoctors} from '@/api/doctor.js'
import { ElMessage } from 'element-plus';


//修改个人信息
const updateUserInfo = async ()=>{
  let data ={
    Name : doctorInfo.value.Name,
    DoctorId : doctorInfo.value.DoctorId,
    DepartmentId :  departmentId.value,
    Introduction :doctorInfo.value.Introduction,
    img : imgUrl.value,
    state: doctorInfo.value.State
  }
    let result = await UpdateDoctors(data)
    
    if(result.code==200){
        ElMessage.success('修改成功')
        //更新pinia中的数据
        doctorData.Doctor.State = doctorInfo.value.State,
        doctorData.Doctor.Name=doctorInfo.value.Name,
        doctorData.Doctor.DepartmentId=departmentId.value,
        doctorData.Doctor.Introduction=doctorInfo.value.Introduction
    }else{
        ElMessage.error("修改失败")
    }
    
}

const imgUrl= ref(doctorData.Doctor.Img)
const uploadSuccess = (result)=>{
    //回显图片
    console.log(result)
    console.log("--------------------")
    imgUrl.value = result.path
    doctorData.Doctor.Img = result.path
    console.log(imgUrl.value)
}

const departments = ref([]);
const departmentId = ref(doctorData.Doctor.DepartmentId);
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

</script>
<template>
  <el-card class="page-container">
    <template #header>
      <div class="header">
        <span>基本资料</span>
      </div>
    </template>

    <el-row>
      <el-col :span="4">
        <div class="avatar-upload-container">
          <el-upload ref="uploadRef" class="avatar-upload" :show-file-list="false" :auto-upload="true"
            action="/api/patient/doupload" name="file" :on-success="uploadSuccess">
            <img v-if="imgUrl" :src="imgUrl" class="avatar" width="200" />
            <img v-else src="https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png" width="200" />
          </el-upload>
        </div>
          <br />
      </el-col>

      <el-col :span="12">
        <el-form :model="doctorInfo" label-width="100px" size="large" class="left-aligned-form">
          <el-form-item label="用户名">
            <el-input v-model="doctorInfo.Username" disabled></el-input>
          </el-form-item>
          <el-form-item label="姓名">
            <el-input v-model="doctorInfo.Name"></el-input>
          </el-form-item>
          <el-form-item label="科室">

            <el-select v-model="departmentId" placeholder="请选择科室" style="width: 200px; margin-right: 10px">
              <el-option v-for="dept in departments" :key="dept.DepartmentId" :label="dept.Name"
                :value="dept.DepartmentId">
              </el-option>
            </el-select>


          </el-form-item>
          <el-form-item label="个人简介">
            <el-input v-model="doctorInfo.Introduction"></el-input>
          </el-form-item>
          <!-- <el-form-item label="余额" >
            <el-input v-model="doctorInfo.balance" disabled></el-input> 
          </el-form-item> -->
          <el-form-item label="状态">
          
          <el-select placeholder="请选择" v-model="doctorInfo.State">
            <el-option label="空闲" value="空闲"></el-option>
            <el-option label="值班中" value="值班中"></el-option>
            <el-option label="手术中" value="手术中"></el-option>
            <el-option label="忙碌" value="忙碌"></el-option>
          </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="updateUserInfo">提交修改</el-button>
            <!-- <el-button type="primary" @click="handleadd">充值</el-button> -->
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
  </el-card>
</template> 


<style lang="scss" scoped>
.avatar-uploader {
    :deep() {
        .avatar {
            width: 278px;
            height: 278px;
            display: block;
        }

        .el-upload {
            border: 1px dashed var(--el-border-color);
            border-radius: 6px;
            cursor: pointer;
            position: relative;
            overflow: hidden;
            transition: var(--el-transition-duration-fast);
        }

        .el-upload:hover {
            border-color: var(--el-color-primary);
        }

        .el-icon.avatar-uploader-icon {
            font-size: 28px;
            color: #8c939d;
            width: 278px;
            height: 278px;
            text-align: center;
        }
    }
}
</style>