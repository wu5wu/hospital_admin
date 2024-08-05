<script setup>
import {reactive, ref} from "vue";

import {ElMessage} from "element-plus";

import {PatientLoginService} from "@/api/login.js";

import { useRouter } from 'vue-router'
const router = useRouter();
const rules = reactive({

  Username: [{required: true, message: '账号不可为空', trigger: 'blur'},{min: 6, max: 18, message: '长度至少6位', trigger: 'blur'},],
  Password: [
    {required: true, message: '密码不可为空', trigger: 'blur'},
    {min: 6, max: 18, message: '长度至少6位', trigger: 'blur'},
  ],
  Name: [{required: true, message: '姓名不可为空', trigger: 'blur'},
  {min: 2, max: 18, message: '长度至少为2', trigger: 'blur'},
  ],
  Identity: [{required: true, message: '身份证号码不可为空', trigger: 'blur'},
  {min: 18, max: 18, message: '长度为18位', trigger: 'blur'},
  ],
  sex: [{required: true, message: '性别不可为空', trigger: 'blur'}],

})


const patientdata = ref({
  Username :'',
  Password :'',
  Name :'',
  Identity :'',
  Img :''
})



const formRef = ref()
const register = async() => {
  
  let result = await PatientLoginService(patientdata.value);
  if (result.code == 200) {
    ElMessage.success(result.message);
    console.log(result);
    // categorys.value = result.data;
    router.push('/login')
    //可以将数据放入浏览器中，方便后续使用
  } else {
    // alert("登录失败")
    ElMessage.error(result.error);
  }


}
</script>


<template>
  <div>
    <div class="login-container">
      <div style="width: 350px" class="login-box">
        <div style="font-weight: bold;font-size: 24px;text-align: center;margin-bottom: 30px">欢迎注册医院预约挂号系统</div>
        <el-form :model="patientdata" ref="formRef" :rules=rules>
          <el-form-item prop="Username">
            <el-input prefix-icon="UserFilled" v-model="patientdata.Username" placeholder="请输入用户名"/>
          </el-form-item>
          <el-form-item prop="Password">
            <el-input show-password  prefix-icon="Lock" v-model="patientdata.Password" placeholder="请输入密码"/>
          </el-form-item>
          <el-form-item prop="Name">
            <el-input  prefix-icon="Name" v-model="patientdata.Name" placeholder="姓名"/>
          </el-form-item>
          <el-form-item prop="Identity">
            <el-input  prefix-icon="Folder" v-model="patientdata.Identity" placeholder="请输入身份证号码"/>
          </el-form-item>
          
          <el-form-item>
            <el-button type="primary" style="width: 100%" @click="register">注册</el-button>
          </el-form-item>
        </el-form>
        <div style="margin-top: 30px;text-align: right">
          已有账号？请<a href="/login">登录</a>
        </div>
      </div>
    </div>
  </div>
</template>





<style scoped>
.login-container {
  min-height: 100vh;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  background-image: url("@/assets/imgs/loginbg.jpg");
  background-size: cover;
}

.login-box {
  background-color: rgba(255, 255, 255, .5);
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  padding: 30px;
  border-radius: 5px;
}

</style>

