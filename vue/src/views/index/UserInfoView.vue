<!--
 * @Author: Callay 2415993100@qq.com
 * @Date: 2024-05-07 19:08:45
 * @LastEditors: Callay 2415993100@qq.com
 * @LastEditTime: 2024-05-07 22:49:21
 * @FilePath: \checkIn\vue\src\views\index\UserInfoView.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
    <div v-if="user">
        <div style="display: flex;justify-content: center;margin-top: 15px;">
            <el-avatar class="avatarClass"> {{ user.id }} </el-avatar>
        </div>
        <div style="margin-top: 20px;display: flex;justify-content: center;align-items: center;">
            <el-form label-width="auto" :model="user" style="max-width: 600px">
                <el-form-item>
                    <p>姓名</p>
                    <el-input v-model="user.name" />
                </el-form-item>
                <el-form-item>
                    <p>密码</p>
                    <el-input type="password" v-model="user.password" />
                </el-form-item>
                <el-form-item>
                    <p>上一次更新时间</p>
                    <el-input disabled v-model="user.updateTime" />
                </el-form-item>
                <el-form-item>
                    <div style="display: flex;justify-content: center;align-items: center;width: 100%;">
                        <el-button type="primary" @click="updateUserInfo">
                            修改
                        </el-button>
                    </div>

                </el-form-item>
            </el-form>
        </div>

    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { onBeforeMount,toRaw } from 'vue';
import HRequest from '@/utils/request';
const request = new HRequest()

var user = ref<any>()
onBeforeMount(() => {
    request.get({ url: 'user/getUserInfo' }).then(res => {
        if (res.code == 200) {
            user.value = res.data
        }
    })
})
function updateUserInfo(){
    request.post({url:'user/updateUserInfo',data:user.value}).then(res=>{
        if(res.code==200) ElMessage.success(res.message)
        else ElMessage.warning(res.message)
    })
}
</script>

<style scoped>
.avatarClass {
    width: 8rem;
    height: 8rem;
    font-size: 2rem;
}
</style>