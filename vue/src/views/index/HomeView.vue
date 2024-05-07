<!--
 * @Author: Callay 2415993100@qq.com
 * @Date: 2024-05-07 18:41:54
 * @LastEditors: Callay 2415993100@qq.com
 * @LastEditTime: 2024-05-07 19:50:48
 * @FilePath: \checkIn\vue\src\views\index\homeView.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE   
-->
<template>
    <div>
        <div style="display: flex;justify-content: center;">
            <p style="font-size: x-large;font-family:Verdana, Geneva, Tahoma, sans-serif">
                组织：
            </p>
            <el-select v-model="value" placeholder="Select" size="default" style="width: 240px">
                <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" />
            </el-select>
        </div>
        <div style="margin-top: 8rem;">
            <div style="display: flex;justify-content: center;">
                <el-button class="buttonClass" type="primary" circle>打卡</el-button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const value = ref('')
var time = ref(Date.now())

const options = [
    {
        value: 'Option1',
        label: 'Option1',
    },
    {
        value: 'Option2',
        label: 'Option2',
    },
    {
        value: 'Option3',
        label: 'Option3',
    },
    {
        value: 'Option4',
        label: 'Option4',
    },
    {
        value: 'Option5',
        label: 'Option5',
    },
]

</script>

<style scoped>
.buttonClass {
    width: 10rem;
    height: 10rem;
    padding: 0;
    font-size: 2rem;
}
</style>