<!--
 * @Author: Callay 2415993100@qq.com
 * @Date: 2024-05-06 22:43:44
 * @LastEditors: Callay 2415993100@qq.com
 * @LastEditTime: 2024-05-07 18:52:04
 * @FilePath: \checkIn\vue\src\views\LoginView.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
    <div>
        <el-card style="margin-top: 12rem;height: 18rem;">
            <div style="width: auto;height: auto;">
                <div style="display: flex;justify-content: center;margin-top: 15px;">
                    <p style="font-size: 2rem;font-family:'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;font-weight: bold;">
                        欢迎登录
                    </p>
                </div>
                <div style="display: flex;justify-content: center;margin-top: 20px;">
                    <el-form ref="loginFormRef" style="max-width: 600px" :model="loginForm" status-icon :rules="rules"
                        label-width="auto" class="demo-ruleForm">
                        <el-form-item label="用户ID" prop="id">
                            <el-input v-model="loginForm.id" type="text" autocomplete="off" />
                        </el-form-item>
                        <el-form-item label="密码" prop="password">
                            <el-input v-model="loginForm.password" type="password" autocomplete="off" />
                        </el-form-item>
                        <div style="display: flex;justify-content: center;">
                            <el-form-item>
                                <el-button type="primary" @click="submitForm(loginFormRef)">
                                    登录
                                </el-button>
                                <el-button @click="resetForm(loginFormRef)">重置</el-button>
                            </el-form-item>
                        </div>

                    </el-form>
                </div>
            </div>
        </el-card>
    </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { type FormInstance, type FormRules } from 'element-plus'
import HRequest from '@/utils/request'
import router from '@/router'

const request = new HRequest()
const loginFormRef = ref<FormInstance>()
const loginForm = reactive({
    id: '',
    password: '',
})

const rules = reactive<FormRules<typeof loginForm>>({
    id: [{ required: true, message: "请输入用户ID", trigger: 'blur' }],
    password: [{ required: true, message: "请输入密码", trigger: 'blur' }],
})

const submitForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.validate((valid, fields) => {
        if (valid) {
            request.post({ url: "user/login", data: loginForm }).then(res => {
                console.log(res)
                if (res.code == 200) {
                    ElMessage.success(res.message)
                    sessionStorage.setItem("jwt-token",res.data)
                    router.push({path:'/index/home'})
                } else {
                    ElMessage.warning(res.message)
                }
            })
        } else {
            ElMessage.warning("登录信息不能为空")
        }
    })
}

const resetForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
}
</script>

<style scoped></style>