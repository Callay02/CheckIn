<!--
 * @Author: Callay 2415993100@qq.com
 * @Date: 2024-05-07 19:08:45
 * @LastEditors: Callay 2415993100@qq.com
 * @LastEditTime: 2024-05-07 22:49:21
 * @FilePath: \checkIn\vue\src\views\index\UserInfoView.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
    <div>
        <div style="display: flex;justify-content: center;margin-top: 15px;">
            <el-avatar class="avatarClass"> {{ user.id }} </el-avatar>
        </div>
        <div style="margin-top: 20px;">
            <div style="display: flex;justify-content: center;">
                <p style="width: 80px;">用户名：</p><el-input v-model="user.name" style="width: 120px;" placeholder="Please input"></el-input>
            </div>
            <div style="display: flex;justify-content: center;margin-top: 15px;">
                <p style="width: 120px;">上一次修改时间：</p><el-input v-model="user.updateTime" style="width: 120px;" placeholder="Please input"></el-input>
            </div>
        </div>

    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { onBeforeMount } from 'vue';
import HRequest from '@/utils/request';
const request = new HRequest()

// var id = ref()
// var name = ref()
// var updateTime = ref()
var user =ref<any>()
onBeforeMount(() => {
    request.get({ url: 'user/getUserInfo' }).then(res => {
        if (res.code == 200) {
            // id.value=res.data.id;
            // name.value=res.data.name;
            // updateTime.value=res.data.updateTime
            user.value=res.data
        }
    })
})
</script>

<style scoped>
.avatarClass {
    width: 8rem;
    height: 8rem;
    font-size: 2rem;
}
</style>