import { defineStore } from "pinia";
import { ref } from 'vue';

export const useUserStore = defineStore('user', () => {
    // 定义描述用户的对象
    const user = ref({
        Username: '',
        Password: '',
        identity: '',
        Name: '',
        Phone: '',
        Balance: '',
        Img: '',
        PatientId: ''
    });


    // 定义设置用户信息的方法
    const setUser = (newUser) => {
        user.value = { ...newUser };
    };

    // 定义移除用户信息的方法
    const removeUser = () => {
        user.value = {
            PatientId: '',
            Username: '',
            Password: '',
            identity: '',
            Name: '',
            Phone: '',
            Balance: '',
            Img: ''
            // 重置其他用户信息字段
        };
    };

    return {
        user,
        setUser,
        removeUser
    };
},
{
    persist:true
});
