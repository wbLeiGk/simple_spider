<script setup lang="ts">
import { ref } from 'vue'
import {  Search } from '@element-plus/icons'
import axios from "axios"
axios.defaults.headers.post['Content-Type'] = 'application/json';
// axios.defaults.headers.post['Content-Type'] = 'text/plain';

defineProps<{ 
  }>()

const url_input = ref("")
let artical_content= ref("")
const url_get_get = async ()=>{
  console.log(url_input.value); 
  const instance = axios.create({baseURL: '/api'});
  let response = await instance.post("/url",{url:url_input.value})
  // artical_content = ref(response.data.message)
  artical_content.value = response.data.message
}
</script>

<template>
  <el-row :gutter="10">
    <el-col :span="12">      
      <el-input
        v-model="url_input"
        placeholder="please input spider url"
        :prefix-icon="Search"
      /></el-col>
    <el-col :span="3">
          <el-button type="primary" :icon="Search" 
          @click="url_get_get">Search</el-button>
    </el-col>
  </el-row>

  <el-row :gutter="10" style="margin-top: 20px">
    <el-col :span="23"> 
      <!-- {{artical_content}} -->
      <el-input
        v-model="artical_content"
        :autosize="{ minRows: 30, maxRows: 30 }"
        type="textarea"
        placeholder="Please input"
      /> 
    </el-col>    

  </el-row>
</template>

<style scoped>
a {
  color: #42b983;
}

label {
  margin: 0 0.5em;
  font-weight: bold;
}

code {
  background-color: #eee;
  padding: 2px 4px;
  border-radius: 4px;
  color: #304455;
}
</style>
