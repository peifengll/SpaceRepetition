<template>
  <div class="center-grid">
    <form class="form">
      <p class="title">Login </p>
      <p class="message">Use Email and Password to login </p>


      <label>
        <input v-model="email" required placeholder="" type="email" class="input">
        <span>Email</span>
      </label>

      <label>
        <input v-model="password" required placeholder="" type="password" class="input">
        <span>Password</span>
      </label>

      <button class="submit" type="button" @click="getLoginData()">login</button>
      <p class="signin"> No an acount ? <a href="register">Register</a></p>
    </form>

  </div>
</template>

<script setup lang="ts">
// import { ElMessage } from 'element-plus'

import request, { SelfUrl } from "../utils/request.ts";
import useStore from "../stores";

const email = ref('');
const password = ref('');
const store = useStore();

const getLoginData = () => {
  // store.increment()
  // console.log(store.count)
  // return
  request.post("/v1/login", {
    email: email.value,
    password: password.value
  }).then(res => {
    console.log(res)
    if (res.status == 200) {
      console.log(res.data.data.accessToken)
      ElMessage({
        showClose: true,
        message: '登录成功',
        type: 'success',
      })
      store.token=res.data.data.accessToken
      window.location.href=SelfUrl
      // console.log("走，去其他页面")
    }
  }).catch(error => {
    console.log(error)
    ElMessage.error('认证失败，请检查邮箱和密码是否正确')

  })

};

</script>


<style scoped>
.center-grid {
  display: grid;
  margin: 90px;
  place-items: center; /* 水平和垂直居中 */
}

.form {
  min-width: 320px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  max-width: 350px;
  background-color: #fff;
  padding: 20px;
  border-radius: 20px;
  position: relative;
}

.title {
  font-size: 28px;
  color: royalblue;
  font-weight: 600;
  letter-spacing: -1px;
  position: relative;
  display: flex;
  align-items: center;
  padding-left: 30px;
}

.title::before, .title::after {
  position: absolute;
  content: "";
  height: 16px;
  width: 16px;
  border-radius: 50%;
  left: 0px;
  background-color: royalblue;
}

.title::before {
  width: 18px;
  height: 18px;
  background-color: royalblue;
}

.title::after {
  width: 18px;
  height: 18px;
  animation: pulse 1s linear infinite;
}

.message, .signin {
  color: rgba(88, 87, 87, 0.822);
  font-size: 14px;
}

.signin {
  text-align: center;
}

.signin a {
  color: royalblue;
}

.signin a:hover {
  text-decoration: underline royalblue;
}

.flex {
  display: flex;
  width: 100%;
  gap: 6px;
}

.form label {
  position: relative;
}

.form label .input {
  width: 100%;
  padding: 10px 10px 20px 10px;
  outline: 0;
  border: 1px solid rgba(105, 105, 105, 0.397);
  border-radius: 10px;
}

.form label .input + span {
  position: absolute;
  left: 10px;
  top: 15px;
  color: grey;
  font-size: 0.9em;
  cursor: text;
  transition: 0.3s ease;
}

.form label .input:placeholder-shown + span {
  top: 15px;
  font-size: 0.9em;
}

.form label .input:focus + span, .form label .input:valid + span {
  top: 30px;
  font-size: 0.7em;
  font-weight: 600;
}

.form label .input:valid + span {
  color: green;
}

.submit {
  border: none;
  outline: none;
  background-color: royalblue;
  padding: 10px;
  border-radius: 10px;
  color: #fff;
  font-size: 16px;
  transform: .3s ease;
}

.submit:hover {
  background-color: rgb(56, 90, 194);
}

@keyframes pulse {
  from {
    transform: scale(0.9);
    opacity: 1;
  }

  to {
    transform: scale(1.8);
    opacity: 0;
  }
}
</style>