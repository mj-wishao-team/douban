import {
    API
} from '../../../API/API.js';

// 头部导航栏部分
const noneAll = document.querySelectorAll('.nav-left a');
const movie = document.querySelector('.movie');
const settings = document.querySelector('#settings');
const myPage = document.querySelector('#my-page');
const none1 = document.querySelector('#none1');
const none2 = document.querySelector('#none2');
const accountSetting = document.querySelector('#account-setting');
const logout = document.querySelector('#logout');
document.querySelector('.account')['onclick'] = () => {
    if (settings.style.display == 'none') {
        settings.style.display = 'block';
    } else {
        settings.style.display = 'none';
    }
}
for (let i = 0; i < noneAll.length; i++) {
    noneAll[i].onclick = () => {
        alert('没有这个功能😭')
    }
}
movie.onclick = () => {
    window.location.href = "../../../index/main/main.html"
}
myPage.onclick = () => {

}
none1.onclick = () => {
    alert('没有这个功能😭')
}
none2.onclick = () => {
    alert('没有这个功能😭')
}
accountSetting.onclick = () => {
    window.location.href = "../accountSetting.html"
}
logout.onclick = () => {
    localStorage.clear();
    window.location.href = "../../../index/main/main.html"
}


const email = document.querySelector('.email-input');
const verifyCode = document.querySelector('.verification-code-input');
// 获取邮箱验证码的函数
async function sendEmail() {
    let message = new FormData();
    message.append('email', `${email.value}`)
    let result = await fetch(`${API.serverAddress}${API.sendEmailAPI}`, {
        method: 'POST',
        body: message
    })
    let jsonObj = await result.json();
    console.log(jsonObj.data);
    console.log(jsonObj.status);
    if (jsonObj.status == 'ture' && verifyCode.value != '') {
        // document.querySelector('.next-step-button')['onclick'] = () => {
        //     nextStep() //获取成功则调用下面验证验证码函数
        // }
    }
}

const reg = /^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/
const emailInput = document.querySelector('.email-input');
const getCodeButton = document.querySelector('.get-code-button');
const error = document.querySelector('.error');
// 获取验证码按钮
document.querySelector('.get-code-button')['onclick'] = () => {
    if (reg.test(emailInput.value) == true) { // 简单校验下邮箱格式
        sendEmail(); //点击获取验证码调用上面函数
        countdown();
    } else {
        error.innerHTML = '请正确填写邮箱';
        setTimeout(() => {
            error.innerHTML = '';
        }, 3000)
    }
}

const getCodeButtonThen = document.querySelector('.get-code-button-then');
// 倒计时函数
function countdown() {
    getCodeButton.style.display = 'none'
    getCodeButtonThen.style.display = 'inline-block';
    let current = 59;
    setTimeout(function go() {
        getCodeButtonThen.innerHTML = `${current}s 后重新发送`
        if (current >= 1) {
            setTimeout(go, 1000);
        };
        current--;
    }, 1000);
    setTimeout(() => {
        getCodeButton.style.display = 'inline-block'
        getCodeButtonThen.style.display = 'none';
    }, 60000);
};

document.querySelector('.next-step-button')['onclick'] = () => {
    nextStep() //获取成功则调用下面验证验证码函数
}

// 验证验证码的函数
const done = document.querySelector('.done')
async function nextStep() {
    let accessToken = localStorage.getItem('access_token');
    let refreshToken = localStorage.getItem('refresh_token');
    let message = new FormData();
    message.append('email', `${email.value}`);
    message.append('verify_code', `${verifyCode.value}`)
    let result = await fetch(`${API.serverAddress}${API.emailVerifyAPI}`, {
        method: 'PUT',
        headers: {
            'Authorization': `Bearer Token ${accessToken}\xa0${refreshToken}`,
        },
        body: message
    })
    let jsonObj = await result.json();
    console.log(`${accessToken}\xa0${refreshToken}`);
    console.log(jsonObj);
    if (jsonObj.data == '修改成功') {
        done.style.display = 'block'
        setTimeout(() => {
            window.location.href = "../accountSetting.html"
        }, 5000)
    }
}

document.querySelector('.done-button')['onclick'] = () => {
    window.location.href = "../accountSetting.html"
}