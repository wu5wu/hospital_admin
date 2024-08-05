  <script setup>
  import { useRoute } from 'vue-router'
  const $route = useRoute()
  console.log($route.path)
  
  // const logout = () => {
  //   localStorage.removeItem('student-user')
  // }

  import { useUserStore } from '@/stores/t_user.js'
const userData = useUserStore();
import {ElMessage,ElMessageBox} from 'element-plus'
import { useTokenStore } from '@/stores/token.js'
const tokenStore = useTokenStore()

import {useRouter} from 'vue-router'
const router = useRouter()
const handleCommand = ()=>{
        ElMessageBox.confirm(
            '你确认退出登录码？',
            '温馨提示',
            {
                confirmButtonText: '确认',
                cancelButtonText: '取消',
                type: 'warning',
            }
        )
            .then(async () => {
                //用户点击了确认
                //清空pinia中的token和个人信息
                userData.user={}
                tokenStore.token=''
                //跳转到登录页
                router.push('/login')
            })
            .catch(() => {
                //用户点击了取消
                ElMessage({
                    type: 'info',
                    message: '取消退出',
                })
            })
    
}


  </script>
  
  <template>
    <div>
      <div style="height: 60px; background-color: #fff; display: flex; align-items: center; border-bottom: 1px solid #ddd">
        <div style="flex: 1">
          <div style="padding-left: 20px; display: flex; align-items: center">
            <img src="@/assets/imgs/logo.png" alt="" style="width: 40px">
            <div style="font-weight: bold; font-size: 24px; margin-left: 5px">医院预约挂号系统</div>
          </div>
        </div>
        <div style="width: fit-content; padding-right: 10px; display: flex; align-items: center;">
          <el-avatar> <img :src="userData.user.Img " alt="" style="width: 40px; height: 40px"> </el-avatar>
          <span style="margin-left: 5px">姓名:{{ userData.user.Name  }}</span>
        </div>
      </div>
  
      <div style="display: flex">
        <div style="width: 200px; border-right: 1px solid #ddd; min-height: calc(100vh - 60px)">
          <el-menu
              router
              style="border: none"
              :default-active="$route.path"
              :default-openeds="['/patient/home']"
          >
            <el-menu-item index="/patient/home">
              <el-icon><HomeFilled /></el-icon>
              <span>系统首页</span>
            </el-menu-item>
  
        
  
            <el-sub-menu index="3">
              <template #title>
                <el-icon><el-icon-menu /></el-icon>
                <span>预约就诊</span>
              </template>
              <el-menu-item index="/patient/appointment">
  
                <span>预约挂号</span>
              </el-menu-item>
          
            </el-sub-menu>
  
  
            <el-sub-menu index="4">
              <template #title>
                <el-icon><el-icon-menu /></el-icon>
                <span>用户管理</span>
              </template>
            
              <el-menu-item index="/patient/reservemy">
                <span>个人预约</span>
              </el-menu-item>
              <el-menu-item index="/patient/information">
                <span>个人信息</span>
              </el-menu-item>
            </el-sub-menu>
  
            <el-menu-item @click=handleCommand>
              <el-icon><SwitchButton /></el-icon>
              <span>退出系统</span>
            </el-menu-item>
  
          </el-menu>
        </div>
  
        <div style="flex: 1; width: 0; background-color: #f8f8ff; padding: 10px">
          <router-view />
        </div>
      </div>
  
    </div>
  </template>
  

  <style scoped>
  .el-menu-item.is-active {
    background-color: #dcede9 !important;
  }
  .el-menu-item:hover {
    color: #11A983;
  }
  :deep(th)  {
    color: #333;
  }
  </style>