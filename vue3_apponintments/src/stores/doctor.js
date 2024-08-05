import { defineStore } from "pinia";
import { ref } from 'vue';

export const doctorStore = defineStore('Doctor', () => {
    // 定义描述用户的对象
    const Doctor = ref({
        DoctorId: '',
        Name: '',
        Username: '',
        Department: '',
        Introduction: '',
        Img: '',
        State:'',
    });


    // 定义设置用户信息的方法
    const setDoctor = (newDoctor) => {
        Doctor.value = { ...newDoctor };
    };

    // 定义移除用户信息的方法
    const removeDoctor = () => {
        user.value = {
        DoctorId: '',
        Name: '',
        Username: '',
        Department: '',
        Introduction: '',
        Img: '',
        State: '',
            // 重置其他用户信息字段
        };
    };

    return {
        Doctor,
        setDoctor,
        removeDoctor
    };
},
{
    persist:true
});
