<script setup>
  import {reactive, ref} from "vue";
  // import request from "@/utils/request";
  import {ElMessage} from "element-plus";
  // import router from "@/router";
  import {AdminRegisterService,DoctorRegisterService,PatientRegisterService} from "@/api/login.js";
  import {ForgotPassword,Update,GenCaptcha} from "@/api/patient.js";
  import { useUserStore } from '@/stores/t_user.js'
import { useTokenStore } from '@/stores/token.js'
import { doctorStore } from '@/stores/doctor.js'
import {useRouter} from 'vue-router'
const router = useRouter()
//调用useTokenStore得到状态
const tokenStore = useTokenStore();
const userData = useUserStore();
const doctorData = doctorStore();
  const data = reactive({
    form: {role:'管理员用户'}
  })
  


  const registerData = ref({
    username: "",
    password: "",
    role: "",
    captcha_id:'',
    answer:'',
})
//清空数据
const clearRegisterData = () => {
    registerData.value = {
      username: '',
      password: '',
      role: "",
    }
}


  const login = async () => {
    let result 
    if(registerData.value.role=="管理员"){
        result = await AdminRegisterService(registerData.value)
        if(result.code==200){
        ElMessage.success("登录成功")
        tokenStore.setToken(result.token)
        router.push('/home')
    }else{
        // alert("登录失败")
        ElMessage.error(result.error)
    }
    }else if(registerData.value.role=="医生"){
      result = await DoctorRegisterService(registerData.value)
      if(result.code==200){
        ElMessage.success("登录成功")
        console.log(result)
        doctorData.setDoctor(result.data)
        tokenStore.setToken(result.token)
        router.push('/doctor')
    }else{
        // alert("登录失败")
        ElMessage.error(result.error)
    }
    }else if(registerData.value.role=="患者"){
      result = await PatientRegisterService(registerData.value)
      if(result.code==200){
        ElMessage.success("登录成功")

        userData.setUser(result.data)
        tokenStore.setToken(result.token)
        console.log(result)
        router.push('/patient/home')
    }else{
        // alert("登录失败")
        ElMessage.error(result.error)
    }
    }else{
      ElMessage.error("请选择角色")
    }  
   
}
const datas = reactive({
    formVisible: false,
    identity:'',
    name :'',
    username:'',
    password:'',
    password2:'',
    patient:[],
    formVisible2: false,
  })
  
const hand = () => {
datas.formVisible = true
}

const handleForgotPassword = async () => {
  let data ={
    identity :datas.identity,
    username : datas.username,
  }
  let result = await ForgotPassword(data)
      if(result.code==200){
        console.log(result.patient)
        ElMessage.success(result.message)
        datas.patient = result.patient
        datas.formVisible = false
        datas.formVisible2 = true
    }else{
        
        ElMessage.error(result.error)
    }
}


const updateUserInfo = async ()=>{

if (datas.password == datas.password2){
  let data ={
    
    PatientId :datas.patient.PatientId,
    Password : datas.password
   
  }
    let result = await Update(data)
    if(result.code==200){
        ElMessage.success('修改成功')
        datas.formVisible2 = false
    }else{
        ElMessage.error("修改失败")
    }
}else{
  ElMessage.error("两次密码输入不一致")
}
  
    
}
// 获取验证码
const captchaData = ref({ data: '', captcha_id: '' });
const getCaptcha = async () => {

  let result = await GenCaptcha()
      if(result.code==200){
        console.log(result)
        captchaData.value.data = result.message.data;
        registerData.value.captcha_id = result.message.captcha_id;
        
    }else{
        
        ElMessage.error(result.error)
    }
}
getCaptcha()


  </script>
  
  
  
  
  
  <template>
    <div>
      <div class="login-container">
        <div style="width: 350px" class="login-box">
          <div style="font-weight: bold;font-size: 24px;text-align: center;margin-bottom: 30px">欢迎登录医院预约挂号系统</div>
          <el-form :model="data.form" ref="formRef" >
            <el-form-item >
              <el-input prefix-icon="UserFilled" v-model="registerData.username" placeholder="请输入用户名"/>
            </el-form-item>
            <el-form-item >
              <el-input show-password  prefix-icon="Lock" v-model="registerData.password" placeholder="请输入密码"/>
            </el-form-item>
        <el-form-item>
          <div class="captcha-container">
            <img :src="captchaData.data" alt="captcha" @click="getCaptcha" style="float: right; margin-left: 10px;"/>
            <el-input v-model="registerData.answer" placeholder="请输入验证码" style="width: 140px; height: 40px; cursor: pointer;"/>
          </div>
        </el-form-item>
            <el-form-item prop="role">
              <el-select style="width:100%"  v-model = "registerData.role">
                <el-option value="管理员" label="管理员"> </el-option>
                <el-option value="医生" label="医生"> </el-option>
                <el-option value="患者" label="患者"> </el-option>
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" style="width: 100%" @click="login">登录</el-button>
            </el-form-item>
          </el-form>
          <div style="display: flex; justify-content: space-between; margin-top: 25px;">
          <div @click="hand" style="cursor: pointer">
           找回密码
          </div>
          <div>
            还没有账号？请<a href="/register">注册</a>
          </div>
          </div>
        </div>
      </div>
    </div>

    <el-dialog
      width="40%"
      v-model="datas.formVisible"
      title="身份认证"
      destroy-on-close>
      <el-form
        :rules="rules"
        ref="formRef"
        label-width="100px"
        label-position="right"
        stylr="padding-right:40px"
      >
        <el-form-item label="用户名">
          <el-input v-model="datas.username" prop="Name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="身份证">
          <el-input v-model="datas.identity" prop="Name" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="datas.formVisible = false">取 消</el-button>
          <el-button type="primary" @click="handleForgotPassword">确认</el-button>
        </span>
      </template>
    </el-dialog>



    <el-dialog
      width="40%"
      v-model="datas.formVisible2"
      title="确认新密码"
      destroy-on-close>
      <el-form
        :rules="rules"
        ref="formRef"
        label-width="100px"
        label-position="right"
        stylr="padding-right:40px"
      >
        <el-form-item label="新密码">
          <el-input v-model="datas.password" prop="Name" autocomplete="off" placeholder="请输入新密码"/>
        </el-form-item>
        <el-form-item label="新密码">
          <el-input v-model="datas.password2" prop="Name" autocomplete="off" placeholder="请再次输入新密码"/>
        </el-form-item>
        
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="datas.formVisible2 = false">取 消</el-button>
          <el-button type="primary" @click="updateUserInfo">确认</el-button>
        </span>
      </template>
    </el-dialog>

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
    background-color: rgba(255, 255, 255, 0);
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    padding: 30px;
    border-radius: 5px;
  }
  .captcha-image {
  cursor: pointer;
  height: 50px;
}
  </style>
  
  
  
  
  
  
  
  
  
