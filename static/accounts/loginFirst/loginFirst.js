// 引入接口文件
import {
    API
} from '../../API/API.js'

const password = document.querySelector('.password input');
const username = document.querySelector('.username input');
const button = document.querySelector('.button');
const errorText = document.querySelector('.error-text');
button.onclick = () => {
    if (username.value == '') {
        errorText.innerHTML = '请填写昵称';
        setTimeout(() => {
            errorText.innerHTML = ''
        }, 5000)
    } else if (password.value == '') {
        errorText.innerHTML = '请填写密码';
        setTimeout(() => {
            errorText.innerHTML = ''
        }, 5000)
    } else if (password.value.length <= 6) {
        errorText.innerHTML = '请填写至少六个字符的密码';
        setTimeout(() => {
            errorText.innerHTML = ''
        }, 5000)
    } else {
        Register() // 简单判断昵称密码合规后，开始注册
    }
}
// 注册函数
let phone = localStorage.getItem('registerPhoneNumber');
async function Register() {
    let message = new FormData();
    message.append('phone', `${phone}`)
    message.append('username', `${username.value}`);
    message.append('password', `${password.value}`)
    let result = await fetch(`${API.serverAddress}${API.smsRegisterAPI}`, {
        method: 'POST',
        body: message,
    })
    let jsonObj = await result.json();
    console.log(jsonObj); // 控制台输出供调试
    if (jsonObj.status == 'ture') {
        alert('注册成功，5秒后跳转到登录页面...')
        console.log(jsonObj.data); // 控制台输出供调试
        setTimeout(() => {
            window.location.href = "../login/login.html"
        }, 5000)
    } else {
        alert(`注册失败，${jsonObj.data}`)
    }
}