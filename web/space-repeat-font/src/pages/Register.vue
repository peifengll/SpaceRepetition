<script setup lang="ts">
import request from "@/utils/request.ts";

const email = ref('')
const password = ref('')
const password2 = ref('')
const toRegister = () => {
  if (password.value=="" || password2.value==""|| email.value==""){
    ElMessage({
      message: 'Warning',
      type: 'warning',
    })
    return;
  }
  if (password.value != password2.value) {
    ElMessage({
      message: 'Warning, 两次输入的密码不一样,请重新输入',
      type: 'warning',
    })
    return
  }

  request.post("v1/register",{
    email: email.value,
    password: password.value
  }).then(res=>{
    ElMessage({
      message: '注册成功，自己后边改名字',
      type: 'success',
    })
    console.log(res)
    location.href('login')
  }).catch(error=>{
    if (error.response.data.code==1001){
      ElMessage.error('这个邮箱注册过了！')
    }else{
      ElMessage.error('未知错误，请联系管理员')
    }
    console.log("error",error)
  })

}

</script>

<template>
  <div class="center-grid">
    <form class="form">
      <p class="title">Register </p>
      <p class="message">Signup now and get full access to our app. </p>
      <label>
        <input v-model="email" required placeholder="" type="email" class="input">
        <span>Email</span>
      </label>

      <label>
        <input v-model="password" required placeholder="" type="password" class="input">
        <span>Password</span>
      </label>
      <label>
        <input v-model="password2" required placeholder="" type="password" class="input">
        <span>Confirm password</span>
      </label>
      <button class="submit" type="button" @click="toRegister">Submit</button>
      <p class="signin">Already have an acount ? <a href="login">Signin</a></p>
    </form>

  </div>
</template>

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