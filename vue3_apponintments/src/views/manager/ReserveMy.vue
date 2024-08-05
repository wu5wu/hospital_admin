<script setup>
import { ElMessage, ElMessageBox } from "element-plus";
import { ref, reactive } from 'vue';
import { getAppointment ,deleteAppointment} from "@/api/appointment.js";
import { useUserStore } from "@/stores/t_user.js";
import { Check, Delete, Edit, Message, Star } from "@element-plus/icons-vue";

const userData = useUserStore();

// 初始化状态
const collectList = ref([]);
// 用于跟踪按钮的禁用状态
const buttonDisabled = reactive({});

//分页条数据模型
const pageNum = ref(1); //当前页
const total = ref(20); //总条数
const pageSize = ref(5); //每页条数

//当每页条数发生了变化，调用此函数
const onSizeChange = (size) => {
  pageSize.value = size;
  getOrderList()
};
//当前页码发生变化，调用此函数
const onCurrentChange = (num) => {
  pageNum.value = num;
  getOrderList()
};


// 获取订单列表
const getOrderList = async () => {
    let params = {
        patientId: 0,
        pageNum: pageNum.value,
        pageSize: pageSize.value,
    };
    let result = await getAppointment(params);
    collectList.value = result.message;
    total.value = result.Total;
    // 初始化所有按钮为可用状态
    result.message.forEach(item => {
        buttonDisabled[item.appointment_id] = {
            collect: false,
            delete: false,
        };
    });
};
getOrderList();

// 删除预约
const deleteCollect = async (appointment_id) => {
    
    try {
        await ElMessageBox.confirm(
            '你确认删除该预约吗？',
            '温馨提示',
            {
                confirmButtonText: '确认',
                cancelButtonText: '取消',
                type: 'warning',
            }
        );

        let data = {
            
            appointment_id: appointment_id,
           
        };

        let result = await deleteAppointment(data);
        if (result.code == 200){

            userData.user.balance = result.balance;
            ElMessage.success("已取消预约");
            getOrderList();
        }else{
            ElMessage.error(result.message);
        }




       
    } catch (error) {
        ElMessage({
            type: 'info',
            message: '取消操作',
        });
    } finally {
        buttonDisabled[appointment_id].delete = false;
    }
}



</script>

<template>
    <el-card class="page-container">
        <template #header>
            <div class="header">
                <span>患者挂号</span>
            </div>
        </template>

        <el-table :data="collectList" style="width: 100%">
            <el-table-column label="预约ID" prop="appointment_id"></el-table-column>
            <el-table-column label="医生" prop="Name"> </el-table-column>
            <el-table-column label="科室" prop="Departments"></el-table-column>
            <el-table-column label="开始时间" prop="StartTime"></el-table-column>
            <el-table-column label="结束时间" prop="EndTime"></el-table-column>
            <el-table-column label="状态" prop="status"></el-table-column>
            <el-table-column label="创建时间" prop="created"></el-table-column>
            <el-table-column label="操作" width="100">
                <template #default="{ row }">
                <el-button
                type="danger"
                @click="deleteCollect(row.appointment_id)"
                :icon="Delete"
                circle
              />
                </template>
            </el-table-column>
            <template #empty>
                <el-empty description="没有数据" />
            </template>
        </el-table>
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
    </el-card>
</template>

<style lang="scss" scoped>
.page-container {
    min-height: 100%;
    box-sizing: border-box;

    .header {
        display: flex;
        align-items: center;
        justify-content: space-between;
    }
}
</style>
