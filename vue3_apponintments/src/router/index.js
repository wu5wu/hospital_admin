//导入vue-router
import { createRouter, createWebHistory } from 'vue-router'
//导入组件
import LoginVue from '@/views/Login.vue'
import RegisterVue from '@/views/Register.vue'
import CollectVUe from '@/views/Collect.vue'
import HomeVue from '@/views/manager/Home.vue'
import ManagerVue from '@/views/Manager.vue'
import CourseVue from '@/views/manager/Course.vue'
import DepartmentVue from '@/views/manager/Department.vue'
import PatientVue from '@/views/Patient.vue'
import AppointmentVue from '@/views/patient/Appointment.vue'
import InformationVue from '@/views/patient/Information.vue'
import ReserveMyVue from '@/views/patient/ReserveMy.vue'
import ReserveMyVue2 from '@/views/manager/ReserveMy.vue'
import TupianceshiVue from '@/views/patient/tupianceshi.vue'
import DoctorHomeVue from '@/views/doctor/DoctorHome.vue'
import InformationVue2 from '@/views/doctor/Information.vue'
import DoctorInfVue from '@/views/manager/DoctorInf.vue'
import PatientInfVue from '@/views/manager/PatientInf.vue'
import AppointmentVue2 from '@/views/doctor/Appointment.vue'
import ApplicationVue from '@/views/doctor/Application.vue'
import ShiftVue from '@/views/doctor/Shift.vue'
import ApplicationInfVue from '@/views/manager/ApplicationInf.vue'
//定义路由关系

const routes = [
    { path: '/login', component: LoginVue },
    { path: '/register', component: RegisterVue },
    { path: '/', component: ManagerVue ,redirect:'/patient/appointment',children:[
        {path:'/home', component: HomeVue}, //二级路由
        {path:'/course', component: CourseVue},
        {path:'/Collect', component: CollectVUe},
        {path:'/department', component: DepartmentVue},
        {path:'/admin/doctorInf', component: DoctorInfVue},
        {path:'/admin/patientInf', component: PatientInfVue},
        {path:'/admin/getappointment', component: ReserveMyVue2},
        {path:'/admin/getappointment', component: ReserveMyVue2},
        {path:'/admin/applicationInf', component: ApplicationInfVue},

    ]},
    

    { path: '/patient', component: PatientVue ,redirect:'/patient/home',children:[
        {path:'/patient/home', component: HomeVue}, //二级路由
        {path:'/patient/appointment', component: AppointmentVue},
        {path:'/Collect', component: CollectVUe},
        {path:'/department', component: DepartmentVue},
        {path:'/patient/information', component: InformationVue},
        {path:'/patient/reservemy', component: ReserveMyVue},
        {path:'/patient/tupian', component: TupianceshiVue},
    ]},



    { path: '/doctor', component: DoctorHomeVue ,redirect:'/doctor/home',children:[
        {path:'/doctor/home', component: HomeVue},
        {path:'/doctor/Information', component: InformationVue2}, //二级路由
        {path:'/doctor/appointment', component: AppointmentVue2},
        {path:'/doctor/application', component: ApplicationVue},
        {path:'/doctor/shift', component: ShiftVue},
        // {path:'/patient/information', component: InformationVue},
        // {path:'/patient/reservemy', component: ReserveMyVue},
        // {path:'/patient/tupian', component: TupianceshiVue},
    ]},


]
//创建路由器
const router = createRouter({
    history: createWebHistory(),
    routes: routes
});

export default router