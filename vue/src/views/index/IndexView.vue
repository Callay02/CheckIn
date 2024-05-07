<!--
 * @Author: Callay 2415993100@qq.com
 * @Date: 2024-05-07 12:42:17
 * @LastEditors: Callay 2415993100@qq.com
 * @LastEditTime: 2024-05-07 19:15:32
 * @FilePath: \checkIn\vue\src\views\index\IndexView.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
    <div class="common-layout">
        <el-container>
            <el-header>
                <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal"  :ellipsis="false" router>
                    <el-menu-item index="home">首页</el-menu-item>
                    <el-menu-item index="history">历史</el-menu-item>
                    <el-menu-item index="organization">组织</el-menu-item>
                    <el-menu-item index="userInfo">个人信息</el-menu-item>
                </el-menu>
            </el-header>
            <el-main>
                <RouterView/>
            </el-main>
        </el-container>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const activeIndex = ref('home')
</script>

<style scoped></style>