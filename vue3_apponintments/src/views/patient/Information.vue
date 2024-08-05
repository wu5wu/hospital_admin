<script setup>
import { ref } from "vue";
import { useUserStore } from '@/stores/t_user.js'
const userData = useUserStore();
console.log(userData.user);

const userInfo = ref({
  user: userData.user.Username,
  identity: userData.user.Identity,
  name: userData.user.Name,
  phone: userData.user.Phone,
  balance: userData.user.Balance,
});
import {TopUp,Update,Getbalance} from '@/api/patient.js'
import { ElMessage } from 'element-plus';


//修改个人信息
const updateUserInfo = async ()=>{
  let data ={
    identity : userInfo.value.identity,
    phone : userInfo.value.phone,
    name :  userInfo.value.name,
    PatientId :userData.user.PatientId,
    img : imgUrl.value
  }
    let result = await Update(data)
    if(result.code==200){
        ElMessage.success('修改成功')
        //更新pinia中的数据
   
  userData.user.Name=userInfo.value.name,
   userData.user.Phone=userInfo.value.phone,
   userData.user.Identity=userInfo.value.identity
    }else{
        ElMessage.error("修改失败")
    }
    
}

//弹窗的控制
 let formVisible = ref(false)
 let balance = ref(0)
 const handleadd  = () => {
  formVisible.value = true;
};

//充值方法
const chozhi = async ()=>{
  let data = {
        patientId: userData.user.PatientId,
       balance: parseInt(balance.value, 10)
    };
    let result = await TopUp(data)
    if(result.code==200){

      window.open(result.result,"_blank")
        //更新pinia中的数据
  //  userData.user.Balance=result.balance
  //  userInfo.value.balance = result.balance
   formVisible.value = false;
    }else{
        ElMessage.error("充值失败")
    }
}

//充值方法
const getbalance = async ()=>{
  let data = {
        patientId: userData.user.PatientId,
    };
    let result = await Getbalance(data)
    if(result.code==200){
        //更新pinia中的数据
   userData.user.Balance=result.balance
   userInfo.value.balance = result.balance
  
    }else{
        ElMessage.error("错误")
    }
}
getbalance();

const imgUrl= ref(userData.user.Img)
const uploadSuccess = (result)=>{
    //回显图片
    console.log(result)
    console.log("--------------------")
    imgUrl.value = result.path
    userData.user.Img = result.path
    console.log(imgUrl.value)
}



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
        <!-- 左侧：图片上传区域 -->
        <div class="avatar-upload-container">
          <!-- 用于文件上传功能 -->
          <el-upload ref="uploadRef" class="avatar-upload" :show-file-list="false" :auto-upload="true"
            action="/api/patient/doupload" name="file" :on-success="uploadSuccess">
            <img v-if="imgUrl" :src="imgUrl" class="avatar" alt="用户头像" width="200" />
            <img v-else src="" width="200" />
            <div  class="avatar-uploader-icon">
              <i class="el-icon-plus"></i>
            </div>
          </el-upload>
        </div>
        <br />

      </el-col>
    
      <el-col :span="12">
        <el-form :model="userInfo" label-width="100px" size="large">
          <el-form-item label="用户名">
            <el-input v-model="userInfo.user" disabled></el-input>
          </el-form-item>
          <el-form-item label="姓名">
            <el-input v-model="userInfo.name"></el-input>
          </el-form-item>
          <el-form-item label="身份证">
            <el-input v-model="userInfo.identity"></el-input>
          </el-form-item>
          <el-form-item label="电话号码">
            <el-input v-model="userInfo.phone"></el-input>
          </el-form-item>
          <el-form-item label="余额">
            <el-input v-model="userInfo.balance" disabled></el-input>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" @click="updateUserInfo">提交修改</el-button>
            <el-button type="primary" @click="handleadd">充值</el-button>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
  </el-card>
  
  <el-dialog
      width="40%"
      v-model="formVisible"
      title="充值"
      destroy-on-close>
      <el-form
        :rules="rules"
        ref="formRef"
        label-width="100px"
        label-position="right"
        stylr="padding-right:40px"
      >
        <el-form-item label="充值金额">
          <el-input v-model="balance" prop="Name" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="formVisible = false">取 消</el-button>
          <el-button type="primary" @click="chozhi">保 存</el-button>
        </span>
      </template>
    </el-dialog>

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