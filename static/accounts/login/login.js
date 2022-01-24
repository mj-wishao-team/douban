// 假超链接
document.querySelector('.douban-logo a')['onclick']=x=>{
    alert('之后应该是跳转到豆瓣电影首页')
}
document.querySelector('.wechat')['onclick']=x=>{
    alert('目前不支持第三方登录')
}
document.querySelector('.weibo')['onclick']=x=>{
    alert('目前不支持第三方登录')
}
document.querySelector('.qq')['onclick']=x=>{
    alert('目前不支持第三方登录')
}
document.querySelector('.password a')['onclick']=x=>{
    alert('目前不支持找回密码')
}

// 悬停"下载豆瓣App"按钮右边的二维码,显示提示文字,2秒后消失
const downloadText = document.querySelector('.download-text');
document.querySelector('.download-picture')['onmouseover']=x=>{
    downloadText.style.display = 'block';
    setTimeout(()=>{downloadText.style.display = 'none'},2000)
};

// 短信登录和密码登录的切换
const phoneWay = document.querySelector('.phone-way');
const passwordWay = document.querySelector('.password-way');
const loginPhoneWay = document.querySelector('#login-phone-way');
const loginPasswordWay = document.querySelector('#login-password-way');
loginPhoneWay.onclick = () => {
    phoneWay.style.display = 'block';
    passwordWay.style.display = 'none';
    loginPhoneWay.className = 'login-select-on';
    loginPasswordWay.className = 'login-no-select-on';
};
loginPasswordWay.onclick = () => {
    phoneWay.style.display = 'none';
    passwordWay.style.display = 'block';
    loginPhoneWay.className = 'login-no-select-on';
    loginPasswordWay.className = 'login-select-on';
};

// 扫码登录和账号登录切换
const QRcodePictureHover = document.querySelector('.QRcode-picture-hover');
const computerPictureHover = document.querySelector('.computer-picture-hover');
const loginText = document.querySelector('.login-text');
const loginQRcode = document.querySelector('.login-QRcode');
QRcodePictureHover.onclick = () => {
    QRcodePictureHover.style.display = 'none';
    computerPictureHover.style.display = 'block';
    loginText.style.display = 'none';
    loginQRcode.style.display = 'block';
};
computerPictureHover.onclick = () => {
    QRcodePictureHover.style.display = 'block';
    computerPictureHover.style.display = 'none';
    loginText.style.display = 'block';
    loginQRcode.style.display = 'none';
};

// 点击下方短信验证登录跳转回去
document.querySelector('.change-password-way')['onclick']=x=>{
    loginText.style.display = 'block';
    loginQRcode.style.display = 'none';
    QRcodePictureHover.style.display = 'block';
    computerPictureHover.style.display = 'none';
}

// 点击+86切换号码地区功能
// 先是打开关闭列表
const phoneArea = document.querySelector('.phone-area');
const allPhoneAreas = document.querySelector('.all-phone-areas');
const maskLayer = document.querySelector('#mask-layer')
phoneArea.onclick = () => {
    maskLayer.className = 'mask-layer'
    allPhoneAreas.style.display = 'block';
}
document.querySelector('.areas-select-cancel')['onclick']=x=>{
    allPhoneAreas.style.display = 'none';
    maskLayer.className = ''
}
// 点击改变号码地区
const globalPhoneListLis = document.querySelectorAll('#global-phone-list-li')
const globalPhoneListNumber = document.querySelectorAll('.GlobalPhoneItem-number')
for(let i = 0; i < globalPhoneListLis.length; i++){  // 为所有li加上一个点击事件
    globalPhoneListLis[i].onclick = () => {
        for(let i = 0; i < globalPhoneListLis.length; i++){
            globalPhoneListLis[i].className = '';
        } // 思路是先清除所有li的类名，然后再为所点击的这个li加上类名
        globalPhoneListLis[i].className = 'is-active';
        let newPhoneAreas = globalPhoneListNumber[i].innerHTML; // 获取海外区位码并且加进去
        maskLayer.className = '';
        phoneArea.innerHTML = `${newPhoneAreas}`;
        allPhoneAreas.style.display = 'none';
    }
}

// 点击发送验证码的部分
const reg = /^1[0-9]{10}$/
let phoneNumber = document.querySelector('.phone-numbers-input')
const errorMessage1Text = document.querySelector('#error-message1-text')
const mpanel4 = document.querySelector('#mpanel4')
document.querySelector('#get-SMS-verification-code')['onclick']=x=>{
    if (phoneArea.innerHTML != '+86' ) {
        alert('目前没法支持海外手机号')
    } else if ( reg.test(phoneNumber.value) == false ) {
        errorMessage1Text.innerHTML = '请正确填写手机号';
        errorMessage1Text.style.display = 'block';
        setTimeout(() => {errorMessage1Text.style.display = 'none'},5000)
    } else { //符合1开头的11位数字（整的弱校验），开始验证码
        mpanel4.style.display = 'block';
        maskLayer.className = 'mask-layer';
    }
}

// 网上偷的滑动拼图验证码
$('#mpanel4').slideVerify({
    type : 2,		//类型
      vOffset : 5,	//误差量，根据需求自行调整
      vSpace : 5,	//间隔
      imgName : ['1.jpg', '2.jpg', '3.jpg', '4.jpg', '5.jpg', '6.jpg', '7.jpg', '8.jpg', '9.jpg', '10.png', '11.png','12.jpg','13.jpg', '14.png'],
      imgSize : {
        width: '400px',
        height: '200px',
      },
      blockSize : {
        width: '40px',
        height: '40px',
      },
      barSize : {
        width : '400px',
        height : '40px',
      },
      ready : function() {
    },
      success : function() {
        // 验证成功，添加你自己的代码
        mpanel4.style.display = 'none';
        maskLayer.className = '';
        smsVerifiyThen()
      },
      error : function() {
          alert('验证失败！');
      }
      
  });	

// 拼图验证码成功后进行的部分
const countdownDiv = document.querySelector('#countdown');
const getSmsVerificationCode = document.querySelector('#get-SMS-verification-code')
// 倒计时函数
function countdown() {
    getSmsVerificationCode.style.display = 'none'
    countdownDiv.style.display = 'block';
    let current = 59;
    setTimeout(function go() {
        countdownDiv.innerHTML = `${current}s 后重新发送`
        if (current >= 1) {
            setTimeout(go, 1000);
        };
        current--;
    }, 1000);
    setTimeout(() => {
        getSmsVerificationCode.style.display = 'inline-block'
        countdownDiv.style.display = 'none';
    }, 60000);
};
// 拼图验证码完成之后要执行的东西
function smsVerifiyThen(){
    countdown(); // 先是发送验证码的倒计时
    smsLogin1() // 然后开始调用发送短信的函数
}

// 此处之后写一个弹出错误提示函数降低代码重复

// 此处之后写一个fetch模板函数降低代码重复

// 密码登录
const accountInput = document.querySelector('.account-input');
const passwordInput = document.querySelector('.password-input');
const loginButton2 = document.querySelector('.login-button2');
const errorMessage2Text = document.querySelector('.error-message2-text');
loginButton2.onclick = () => {
    (async () => {
        let message = new FormData();
        message.append('loginAccount', `${accountInput.value}`);
        message.append('password', `${passwordInput.value}`);
        let result = await fetch('http://106.55.225.88:9090/api/user/login/pw', {
            method: 'POST',
            body: message,
        });
        let jsonObj = await result.json();
        console.log(jsonObj.data);       // 此处是控制台输出状态供调试

        if(jsonObj.status == 'ture'){
            localStorage.setItem('access_token', jsonObj.access_token);   //把jwt存在本地
            localStorage.setItem('refresh_token', jsonObj.refresh_token);
            alert('成功登录');
        }else if(jsonObj.data == '账号或密码错误'){
            errorMessage2Text.innerHTML = '用户名或密码错误'
            errorMessage2Text.style.display = 'block' 
            setTimeout(() =>{errorMessage2Text.style.display = 'none'},5000)
        }else if(jsonObj.data == '请输入密码'){
            errorMessage2Text.innerHTML = '请输入密码';
            errorMessage2Text.style.display = 'block';
            setTimeout(() =>{errorMessage2Text.style.display = 'none'},5000)
        }else if(jsonObj.data == '请输入注册时用的邮箱或者手机号'){
            errorMessage2Text.innerHTML = '请填写账号';
            errorMessage2Text.style.display = 'block';
            setTimeout(() =>{errorMessage2Text.style.display = 'none'},5000)
        }else{
            console.log('服务器繁忙，请稍后再试');
            alert('服务器繁忙，请稍后再试');
        }
    })()     
}

// 短信登录与注册
const phoneNumbersInput = document.querySelector('.phone-numbers-input');
const smsVerificationCodeInput = document.querySelector('.SMS-verification-code-input');
const welcomeRegister = document.querySelector('.welcomeRegister');
const loginButton1 = document.querySelector('.login-button1');
// 第一次，发送手机号请求验证码的函数
async function smsLogin1() {
    let message = new FormData();
    message.append('phone', `${phoneNumbersInput.value}`)
    let result = await fetch('http://106.55.225.88:9090/api/verify/sms', {
        method: 'POST',
        body: message
    })
    let jsonObj = await result.json();
    console.log(jsonObj.data);// 此处是控制台输出状态供调试
    console.log('info内容为' + jsonObj.info);// 此处是控制台输出状态供调试

    if(jsonObj.data == '电话号码不能为空'){
        errorMessage1Text.innerHTML = '电话号码不能为空';
        errorMessage1Text.style.display = 'block';
        setTimeout(() =>{errorMessage1Text.style.display = 'none'},5000)
    }else if(jsonObj.data == '手机号格式错误'){
        errorMessage1Text.innerHTML = '请正确填写手机号';
        errorMessage1Text.style.display = 'block';
        setTimeout(() =>{errorMessage1Text.style.display = 'none'},5000)
    }else if(jsonObj.info == '服务器出错'){
        console.log('服务器繁忙，请稍后再试');
        alert('服务器繁忙，请稍后再试');
    }

    // 按下按钮，如果后端返回短信发送成功,调用第二次发送的函数
    loginButton1.onclick = () => {
        if(jsonObj.info == '短信发送成功' && smsVerificationCodeInput.value != ''){
            smsLongin2();
            loginButton1.innerHTML = '登录中...';
        }    
    }
}

// 第二次，请求验证码成功后，调用此函数与另一个接口发送验证码与手机号
async function smsLongin2(){                   
    let message = new FormData();
    message.append('phone', `${phoneNumbersInput.value}`);
    message.append('verify_code', `${smsVerificationCodeInput.value}`)
    let result = await fetch('http://106.55.225.88:9090/api/user/login/sms', {
        method: 'POST',
            body: message,
        })
    let jsonObj = await result.json();
    console.log(jsonObj.data);              // 此处是控制台输出状态供调试
    console.log('info内容为' + jsonObj.info);// 此处是控制台输出状态供调试
    if(jsonObj.info == '新用户'){
        smsRegister();             // 判断是新用户，调用下面的函数来注册
    }else if(jsonObj.status == 'ture'){
        localStorage.setItem('access_token', jsonObj.access_token);   //成功登录,把jwt存在本地
        localStorage.setItem('refresh_token', jsonObj.refresh_token);
        alert('成功登录');
        loginButton1.innerHTML = '登录';
    }else if(jsonObj.data == '未发送验证码'){
        errorMessage1Text.innerHTML = '验证码不能为空';
        errorMessage1Text.style.display = 'block';
        setTimeout(() =>{errorMessage1Text.style.display = 'none'},5000)
        loginButton1.innerHTML = '登录';
    }else if(jsonObj.data == '电话号码不能为空'){
        errorMessage1Text.innerHTML = '电话号码不能为空';
        errorMessage1Text.style.display = 'block';
        setTimeout(() =>{errorMessage1Text.style.display = 'none'},5000)
        loginButton1.innerHTML = '登录';
    }else if(jsonObj.data == '手机号格式错误'){
        errorMessage1Text.innerHTML = '请正确填写手机号';
        errorMessage1Text.style.display = 'block';
        setTimeout(() =>{errorMessage1Text.style.display = 'none'},5000)
        loginButton1.innerHTML = '登录';
    }else if(jsonObj.data == '验证码错误或者过期'){
        errorMessage1Text.innerHTML = '验证码错误或者过期';
        errorMessage1Text.style.display = 'block';
        setTimeout(() =>{errorMessage1Text.style.display = 'none'},5000)
        loginButton1.innerHTML = '登录';
    }else if(jsonObj.data == '验证码不能为空'){
        errorMessage1Text.innerHTML = '验证码不能为空';
        errorMessage1Text.style.display = 'block';
        setTimeout(() =>{errorMessage1Text.style.display = 'none'},5000)
        loginButton1.innerHTML = '登录';
    }
    // else {
    //     console.log('服务器繁忙，请稍后再试');
    //     alert('服务器繁忙，请稍后再试');
    //     smsVerificationCodeInput.innerHTML = '登录';
    // }
            
}
// 第三次，如果是第一次使用的新用户，执行这个函数，把手机号存储到本地，并弹出注册协议窗口
function smsRegister(){
    localStorage.setItem('registerPhoneNumber', `${phoneNumbersInput.value}`)
    maskLayer.className = 'mask-layer';
    welcomeRegister.style.display = 'block';
}

// 新用户短信注册弹出的注册协议窗口部分
document.querySelector('.welcomeRegister-cancel')['onclick']=x=>{
    maskLayer.className = '';
    welcomeRegister.style.display = 'none';
}
document.querySelector('#register-button')['onclick']=x=>{
    maskLayer.className = '';
    welcomeRegister.style.display = 'none';
    window.location.href="../loginFirst/loginFirst.html"
    smsVerificationCodeInput.innerHTML = '登录';
}
