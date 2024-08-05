<script setup>
import {
    Edit,
    Delete
} from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from "element-plus";
import { ref } from 'vue'
import { collectListService,collectDeleteService} from "@/api/collect.js";
import { useUserStore } from "@/stores/t_user.js";
const userData = useUserStore();

//用户搜索时数据绑定
const goodsName=ref('')

//文章列表数据模型
const collectList = ref([
    
])

//分页条数据模型
const pageNum = ref(1)//当前页
const total = ref(20)//总条数
const pageSize = ref(5)//每页条数

//当每页条数发生了变化，调用此函数
const onSizeChange = (size) => {
    pageSize.value = size
    getOrderList()
}
//当前页码发生变化，调用此函数
const onCurrentChange = (num) => {
    pageNum.value = num
    getOrderList()
}


const getOrderList = async () => {
  let params = {
    pageNum: pageNum.value,
    pageSize: pageSize.value,
  };
  let collect = {
    user_id: userData.user.student_id,
    good_name: goodsName.value,
  };
  let result = await collectListService(params, collect);
  //渲染列表数据
  console.log(result);
  collectList.value = result.list;

  //渲染总条数
  total.value = result.total;
};
getOrderList();



const deleteCollect = (row) => {
    ElMessageBox.confirm(
        '你确认取消收藏吗？',
        '温馨提示',
        {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning',
        }
    )
        .then(async () => {
            //用户点击了确认
            let result = await collectDeleteService(row.goods_id)
            ElMessage.success(result.msg)
            getOrderList()
        })
        .catch(() => {
            ElMessage({
                type: 'info',
                message: '取消删除',
            })
        })
}
</script>
<template>
    <el-card class="page-container">
        <template #header>
            <div class="header">
                <span>收藏列表</span>
            </div>
        </template>
        <!-- 搜索表单 -->
        <el-form inline>
            <el-form-item label="搜索：">
                <el-input v-model="goodsName" placeholder="求输入商品名称" clearable />
            </el-form-item>
            
            <el-form-item>
                <el-button type="primary" @click="getOrderList">搜索</el-button>
                <el-button @click="(goodsName = '')">重置</el-button>
            </el-form-item>
        </el-form>
        <!-- 文章列表 -->
        <el-table :data="collectList" style="width: 100%">
            <el-table-column label="商品编号"  prop="goods_id"></el-table-column>
            <el-table-column label="图片">
        <template #default="{ row }">
          <img :src="row.img" style="width: 50px; height: 50px" />
        </template>
      </el-table-column>
            <el-table-column label="商品名称" prop="good_name"></el-table-column>
            <el-table-column label="介绍" prop="describle"> </el-table-column>
            <el-table-column label="价格" prop="price"></el-table-column>
            <el-table-column label="操作" width="100">
                <template #default="{ row }">
                    <el-button :icon="Delete" circle plain type="danger" @click="deleteCollect(row)"></el-button>
                </template>
            </el-table-column>
            <template #empty>
                <el-empty description="没有数据" />
            </template>
        </el-table>
        <!-- 分页条 -->
        <el-pagination v-model:current-page="pageNum" v-model:page-size="pageSize" :page-sizes="[3, 5 ,10, 15]"
            layout="jumper, total, sizes, prev, pager, next" background :total="total" @size-change="onSizeChange"
            @current-change="onCurrentChange" style="margin-top: 20px; justify-content: flex-end" />
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